/*
Copyright 2023 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader" // loads the HTTP loader
	"golang.org/x/mod/semver"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	TEMPLATES          = "tools/templates/*.tmpl"
	CODE               = "./pkg/api"
	GEN_CODE           = "./pkg/api"
	REPO_ROOT          string
	TEMPLATES_FOLDER   string
	CODE_FOLDER        string
	GEN_CODE_FOLDER    string
	SPEC_FOLDER_PREFIX = "spec-"
	SPEC_VERSIONS      = []string{"0.3.0"}
	SCHEMA_FOLDER      = "schemas"
	TEST_SCHEMA_FOLDER = "tests"
	TEST_OUTPUT_PREFIX = "ztest_"

	GO_TYPES_NAMES = map[string]string{
		"taskrun":      "TaskRun",
		"pipelinerun":  "PipelineRun",
		"testcaserun":  "TestCaseRun",
		"testsuiterun": "TestSuiteRun",
		"testoutput":   "TestOutput",
	}

	GO_TYPES_TEST_NAMES = map[string]string{
		"foosubject":   "FooSubject",
		"barpredicate": "BarPredicate",
	}

	// Templates
	eventTemplateFileName          = "event.go.tmpl"
	typesTemplateFileName          = "types.go.tmpl"
	examplesTestsTemplateFileNames = []string{
		"examples_test.go.tmpl",
		"factory_test.go.tmpl",
	}
	specTemplateFileName = "docs.go.tmpl"

	// Tool
	capitalizer cases.Caser

	// Flags
	RESOURCES_PATH = flag.String("resources", "", "the path to the generator resources root folder")
)

const REFERENCE_TYPE = "Reference"

// ContentField holds the name and type of each content field
type ContentField struct {
	Name      string
	NameLower string
	Type      string
	Required  bool
}

// ContentType holds the data required to render any custom
// type used within the content.
type ContentType struct {
	Name   string
	Fields []ContentField
}

type Data struct {
	Subject        string
	SubjectLower   string
	Predicate      string
	PredicateLower string
	Version        string
	VersionName    string
	SubjectType    string
	Contents       []ContentField
	ContentTypes   []ContentType
	Prefix         string
	Schema         string
	IsTestData     bool
}

type AllData struct {
	Slice            []Data
	SpecVersion      string
	SpecVersionShort string
	SpecVersionName  string
	IsTestData       bool
}

func (d Data) OutputFile() string {
	return "zz_" + d.Prefix + d.SubjectLower + d.PredicateLower + "_" + d.VersionName + ".go"
}

func init() {
	capitalizer = cases.Title(language.English, cases.NoLower)
}

// GoTypeName returns the name to be used when building Go types
// Special mappings are defined in mappings.
// For other words, the "Title" caser is used.
func GoTypeName(schemaName string, mappings map[string]string) string {
	name, ok := mappings[schemaName]
	if !ok {
		return capitalizer.String(schemaName)
	} else {
		return name
	}
}

func main() {
	var err error
	var ex string

	// Parse input parameters
	log.SetFlags(0)
	log.SetPrefix("generator: ")
	flag.Parse()

	if *RESOURCES_PATH == "" {
		ex, err = os.Executable()
		if err != nil {
			panic(err)
		}
		toolPath := filepath.Clean(filepath.Dir(filepath.Join(ex, "..")))
		RESOURCES_PATH = &toolPath
	}

	// Setup folder variables
	TEMPLATES_FOLDER = filepath.Join(*RESOURCES_PATH, TEMPLATES)
	CODE_FOLDER = filepath.Join(*RESOURCES_PATH, CODE)
	GEN_CODE_FOLDER = filepath.Join(*RESOURCES_PATH, GEN_CODE)

	// Generate SDK files
	for _, version := range SPEC_VERSIONS {
		shortVersion := semver.MajorMinor("v" + version)
		versioned_schema_folder := filepath.Join(CODE_FOLDER, SPEC_FOLDER_PREFIX+shortVersion, SCHEMA_FOLDER)
		log.Printf("Generating SDK files from templates: %s and schemas: %s into %s", TEMPLATES_FOLDER, versioned_schema_folder, GEN_CODE_FOLDER)
		err = generate(versioned_schema_folder, GEN_CODE_FOLDER, "", version, GO_TYPES_NAMES, false)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
	}

	// Generate SDK test files
	test_schema_folder := filepath.Join(CODE_FOLDER, TEST_SCHEMA_FOLDER, SCHEMA_FOLDER)
	log.Printf("Generating Test SDK files from templates: %s and schemas: %s into %s", TEMPLATES_FOLDER, test_schema_folder, GEN_CODE_FOLDER)
	err = generate(test_schema_folder, GEN_CODE_FOLDER, TEST_OUTPUT_PREFIX, "99.0.0", GO_TYPES_TEST_NAMES, true)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func generate(schemaFolder, genFolder, prefix, specVersion string, goTypes map[string]string, isTestMode bool) error {
	// allData is used to accumulate data from all jsonschemas
	// which is then used to run shared templates
	shortSpecVersion := semver.MajorMinor("v" + specVersion)
	allData := AllData{
		Slice:            make([]Data, 0),
		SpecVersion:      specVersion,
		SpecVersionShort: shortSpecVersion,
		SpecVersionName:  strings.Replace(shortSpecVersion, ".", "", -1),
		IsTestData:       isTestMode,
	}

	allTemplates, err := template.ParseGlob(TEMPLATES_FOLDER)
	if err != nil {
		return err
	}

	// Walk the jsonschemas folder, process each ".json" file
	walkProcessor := getWalkProcessor(schemaFolder, allTemplates, genFolder, goTypes, &allData, prefix, isTestMode)
	err = fs.WalkDir(os.DirFS(schemaFolder), ".", walkProcessor)
	if err != nil {
		return err
	}

	// Process the spec template. Create the target folder is it doesn't exist
	specFileFolder := filepath.Join(genFolder, allData.SpecVersionName)
	err = os.MkdirAll(specFileFolder, os.ModePerm)
	if err != nil {
		return err
	}

	// Spec types (types.go)
	outputFileName := filepath.Join(genFolder, allData.SpecVersionName, strings.TrimSuffix(typesTemplateFileName, filepath.Ext(typesTemplateFileName)))
	err = executeTemplate(allTemplates, typesTemplateFileName, outputFileName, allData)
	if err != nil {
		return err
	}

	// Spec aliases (docs.go)
	specFileName := filepath.Join(genFolder, allData.SpecVersionName, strings.TrimSuffix(specTemplateFileName, filepath.Ext(specTemplateFileName)))
	err = executeTemplate(allTemplates, specTemplateFileName, specFileName, allData)
	if err != nil {
		return err
	}

	// Process example test files - only for real data
	if !isTestMode {
		for _, examplesTestsTemplateFileName := range examplesTestsTemplateFileNames {
			outputFileName := filepath.Join(genFolder, "zz_"+prefix+strings.TrimSuffix(examplesTestsTemplateFileName, filepath.Ext(examplesTestsTemplateFileName)))
			err = executeTemplate(allTemplates, examplesTestsTemplateFileName, outputFileName, allData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func executeTemplate(allTemplates *template.Template, templateName, outputFileName string, data interface{}) error {
	// Write the template output to a buffer
	generated := new(bytes.Buffer)

	err := allTemplates.ExecuteTemplate(generated, templateName, data)
	if err != nil {
		return err
	}

	src, err := format.Source(generated.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		log.Printf("%s", generated.String())
		return err
	}

	// Prepare the output file
	return os.WriteFile(outputFileName, src, 0644)
}

func getWalkProcessor(rootDir string, allTemplates *template.Template, genFolder string, goTypes map[string]string, allData *AllData, prefix string, isTestMode bool) fs.WalkDirFunc {
	return func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Do not go into sub-folders
		if info.IsDir() {
			if info.Name() == "." {
				return nil
			}
			return fs.SkipDir
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			// This should not happen, but ignore just in case
			return nil
		}
		// Set the whole path
		schemaPath := filepath.Join(rootDir, path)
		// Load the jsonschema from the spec
		sch, err := jsonschema.Compile(schemaPath)
		if err != nil {
			return err
		}

		// Prepare the data
		data, err := DataFromSchema(sch, goTypes)
		if err != nil {
			return err
		}
		data.Prefix = prefix
		data.IsTestData = isTestMode
		// Load the raw schema data
		rawSchema, err := os.ReadFile(schemaPath)
		if err != nil {
			return err
		}
		compressedRawSchema := bytes.NewBuffer([]byte{})
		err = json.Compact(compressedRawSchema, rawSchema)
		if err != nil {
			return err
		}
		data.Schema = compressedRawSchema.String()
		allData.Slice = append(allData.Slice, *data)

		// Execute the template
		return executeTemplate(allTemplates, eventTemplateFileName, filepath.Join(genFolder, data.OutputFile()), data)
	}
}

func DataFromSchema(schema *jsonschema.Schema, mappings map[string]string) (*Data, error) {
	// Parse the event type from the context
	contextSchema, ok := schema.Properties["context"]
	if !ok {
		return nil, fmt.Errorf("no context property in schema %s", schema.Location)
	}
	eventTypeSchema, ok := contextSchema.Properties["type"]
	if !ok {
		return nil, fmt.Errorf("no type property in schema %s", eventTypeSchema.Location)
	}
	if len(eventTypeSchema.Enum) == 0 {
		return nil, fmt.Errorf("no value defined for type in schema %s", eventTypeSchema.Location)
	}
	eventTypeString, ok := eventTypeSchema.Enum[0].(string)
	if !ok {
		return nil, fmt.Errorf("non-string value defined for type in schema %s", eventTypeSchema.Location)
	}
	if eventTypeString == "" {
		return nil, fmt.Errorf("empty value defined for type in schema %s", eventTypeSchema.Location)
	}
	eventType, err := cdevents.CDEventTypeFromString(string(eventTypeString))
	if err != nil {
		return nil, err
	}

	// Parse the subject type
	subjectSchema, ok := schema.Properties["subject"]
	if !ok {
		return nil, fmt.Errorf("no subject property in schema %s", schema.Location)
	}
	subjectTypeSchema, ok := subjectSchema.Properties["type"]
	if !ok {
		return nil, fmt.Errorf("no type property in schema %s", subjectSchema.Location)
	}
	if len(subjectTypeSchema.Enum) == 0 {
		return nil, fmt.Errorf("no value defined for type in schema %s", subjectTypeSchema.Location)
	}
	subjectTypeString, ok := subjectTypeSchema.Enum[0].(string)
	if !ok {
		return nil, fmt.Errorf("non-string value defined for type in schema %s", subjectTypeSchema.Location)
	}

	// Parse the subject content fields
	contentFields := []ContentField{}
	contentTypes := []ContentType{}
	contentSchema, ok := subjectSchema.Properties["content"]
	if !ok {
		return nil, fmt.Errorf("no content property in schema %s", subjectSchema.Location)
	}
	for name, propertySchema := range contentSchema.Properties {
		contentField := ContentField{}
		contentField.NameLower = name
		contentField.Name = capitalizer.String(name)
		contentField.Required = false
		for _, value := range contentSchema.Required {
			if name == value {
				contentField.Required = true
			}
		}
		if len(propertySchema.Types) != 1 {
			return nil, fmt.Errorf("only one type allowed for content property in schema %s", propertySchema.Location)
		}
		switch propertySchema.Types[0] {
		case "object":
			contentType, err := typesForSchema(name, propertySchema, mappings)
			if err != nil {
				return nil, err
			}
			namespacedType := GoTypeName(contentType.Name, mappings)
			if contentType.Name != REFERENCE_TYPE {
				// If this is not a "Reference" we need to define a new type
				contentTypes = append(contentTypes, *contentType)
				// If this is not a "Reference" we need to namespace the type name to the event
				namespacedType = GoTypeName(eventType.Subject, mappings) +
					GoTypeName(eventType.Predicate, mappings) + "SubjectContent" +
					GoTypeName(contentType.Name, mappings)
			}
			// We must use pointers here for "omitempty" to work when rendering to JSON
			contentField.Type = "*" + namespacedType
		case "string":
			contentField.Type = "string"
		default:
			return nil, fmt.Errorf("content property type %s not allowed in schema %s", contentField.Type, propertySchema.Location)
		}
		contentFields = append(contentFields, contentField)
	}
	// Sort contents for deterministic code rendering
	sort.Slice(contentFields, func(i, j int) bool {
		return contentFields[i].Name < contentFields[j].Name
	})
	sort.Slice(contentTypes, func(i, j int) bool {
		return contentTypes[i].Name < contentTypes[j].Name
	})
	return &Data{
		Subject:        GoTypeName(eventType.Subject, mappings),
		Predicate:      GoTypeName(eventType.Predicate, mappings),
		SubjectLower:   eventType.Subject,
		PredicateLower: eventType.Predicate,
		Version:        eventType.Version,
		VersionName:    strings.ReplaceAll(eventType.Version, ".", "_"),
		SubjectType:    subjectTypeString,
		Contents:       contentFields,
		ContentTypes:   contentTypes,
	}, nil
}

// typesForSchema takes a property from a jsonschema and produces
// a ContentType object, as long as all fields are of type string
func typesForSchema(name string, property *jsonschema.Schema, mappings map[string]string) (*ContentType, error) {
	fields := []ContentField{}
	otherNames := []string{}
	referenceFields := []string{}
	for name, propertySchema := range property.Properties {
		switch name {
		case "id", "source":
			referenceFields = append(referenceFields, name)
		default:
			otherNames = append(otherNames, name)
		}
		if len(propertySchema.Types) != 1 {
			return nil, fmt.Errorf("only one type allowed for content property in schema %s", propertySchema.Location)
		}
		if propertySchema.Types[0] != "string" {
			return nil, fmt.Errorf("only one string type allowed for content property in schema %s", propertySchema.Location)
		}
		field := ContentField{
			NameLower: name,
			Name:      GoTypeName(name, mappings),
			Type:      "string",
			Required:  false,
		}
		for _, value := range property.Required {
			if name == value {
				field.Required = true
			}
		}
		fields = append(fields, field)
	}
	// Check if this is a reference
	if len(referenceFields) == 2 && len(otherNames) == 0 {
		name = REFERENCE_TYPE
	}
	// Sort fields for consistent generation
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name
	})
	return &ContentType{
		Name:   GoTypeName(name, mappings),
		Fields: fields,
	}, nil
}
