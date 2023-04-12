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
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	TEMPLATES            = "tools/templates/*.tmpl"
	SCHEMA_FOLDER        = "./pkg/api/spec/schemas"
	GEN_CODE_FOLDER      = "./pkg/api"
	TEST_TEMPLATES       = "tools/templates_test/*.tmpl"
	TEST_SCHEMA_FOLDER   = "./pkg/api/tests/schemas"
	TEST_GEN_CODE_FOLDER = "./pkg/api"
	TEST_OUTPUT_PREFIX   = "ztest_"

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

	// Tool
	capitalizer cases.Caser
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
	SubjectType    string
	Contents       []ContentField
	ContentTypes   []ContentType
	Prefix         string
	Schema         string
}

type AllData struct {
	Slice []Data
}

func (d Data) OutputFile() string {
	return "zz_" + d.Prefix + d.SubjectLower + d.PredicateLower + ".go"
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
	// Parse input parameters
	log.SetFlags(0)
	log.SetPrefix("generator: ")
	flag.Parse()

	var err error

	// Generate SDK files
	log.Printf("Generating SDK files from templates: %s and schemas: %s into %s", TEMPLATES, SCHEMA_FOLDER, GEN_CODE_FOLDER)
	err = generate(TEMPLATES, SCHEMA_FOLDER, GEN_CODE_FOLDER, "", GO_TYPES_NAMES)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Generate SDK test files
	log.Printf("Generating SDK files from templates: %s and schemas: %s into %s", TEST_TEMPLATES, TEST_SCHEMA_FOLDER, TEST_GEN_CODE_FOLDER)
	err = generate(TEST_TEMPLATES, TEST_SCHEMA_FOLDER, TEST_GEN_CODE_FOLDER, TEST_OUTPUT_PREFIX, GO_TYPES_TEST_NAMES)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func generate(templatesFolder, schemaFolder, genFolder, prefix string, goTypes map[string]string) error {
	// allData is used to accumulate data from all jsonschemas
	// which is then used to run shared templates
	allData := AllData{
		Slice: make([]Data, 0),
	}

	allTemplates, err := template.ParseGlob(templatesFolder)
	if err != nil {
		return err
	}

	// Walk the jsonschemas folder, process each ".json" file
	walkProcessor := getWalkProcessor(allTemplates, genFolder, goTypes, &allData, prefix)
	err = filepath.Walk(schemaFolder, walkProcessor)
	if err != nil {
		return err
	}

	// Process the types template
	outputFileName := genFolder + string(os.PathSeparator) + "zz_" + prefix + strings.TrimSuffix(typesTemplateFileName, filepath.Ext(typesTemplateFileName))
	err = executeTemplate(allTemplates, typesTemplateFileName, outputFileName, allData.Slice)
	if err != nil {
		return err
	}

	// Process example test files - only for real data
	if prefix == "" {
		for _, examplesTestsTemplateFileName := range examplesTestsTemplateFileNames {
			outputFileName := genFolder + string(os.PathSeparator) + "zz_" + prefix + strings.TrimSuffix(examplesTestsTemplateFileName, filepath.Ext(examplesTestsTemplateFileName))
			err = executeTemplate(allTemplates, examplesTestsTemplateFileName, outputFileName, allData.Slice)
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

func getWalkProcessor(allTemplates *template.Template, genFolder string, goTypes map[string]string, allData *AllData, prefix string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Nothing to do with folders
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			// This should not happen, but ignore just in case
			return nil
		}
		// Load the jsonschema from the spec
		sch, err := jsonschema.Compile(path)
		if err != nil {
			return err
		}

		// Prepare the data
		data, err := DataFromSchema(sch, goTypes)
		if err != nil {
			return err
		}
		data.Prefix = prefix
		// Load the raw schema data
		rawSchema, err := os.ReadFile(path)
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
		return executeTemplate(allTemplates, eventTemplateFileName, genFolder+string(os.PathSeparator)+data.OutputFile(), data)
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
