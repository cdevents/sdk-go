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
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
	"golang.org/x/mod/semver"
)

const testSchemaJSON = "../pkg/api/tests-v99.1/schemas/foosubjectbarpredicate.json"
const specVersion = "0.4.1"

var (
	testSchema      *jsonschema.Schema
	testSubject     = "FooSubject"
	testSubjectType = "fooSubject"
	testPredicate   = "BarPredicate"
	testVersion     = "2.2.3"
	testVersionName = "2_2_3"
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	var err error
	pathLoader := PathLoader{}
	loader := jsonschema.SchemeURLLoader{
		"file":  jsonschema.FileLoader{},
		"http":  pathLoader,
		"https": pathLoader,
	}
	compiler = *jsonschema.NewCompiler()
	compiler.UseLoader(loader)
	schemas = Schemas{
		IsTestData: false,
		Data:       make(map[string][]byte),
	}
	shortVersion := semver.MajorMinor("v" + specVersion)
	schemaFolder := filepath.Join("../pkg/api", SpecFolderPrefix+shortVersion, SchemaFolders[0]) // links
	err = loadSchemas(schemaFolder, &schemas)
	panicOnError(err)
	testSchema, err = compiler.Compile(testSchemaJSON)
	panicOnError(err)
}

func TestDataFromSchema(t *testing.T) {
	want := &Data{
		Subject:        testSubject,
		Predicate:      testPredicate,
		SubjectLower:   strings.ToLower(testSubject),
		PredicateLower: strings.ToLower(testPredicate),
		Version:        testVersion,
		VersionName:    testVersionName,
		SubjectType:    testSubjectType,
		Contents: []ContentField{{
			Name:      "ArtifactId",
			NameLower: "artifactId",
			Type:      "string",
		}, {
			Name:      "ReferenceField",
			NameLower: "referenceField",
			Type:      "*Reference",
			Required:  true,
		}, {
			Name:      "PlainField",
			NameLower: "plainField",
			Type:      "string",
			Required:  true,
		}, {
			Name:      "ObjectField",
			NameLower: "objectField",
			Type:      "*FooSubjectBarPredicateSubjectContentObjectFieldV2_2_3",
		}},
		ContentTypes: []ContentType{{
			Name: "ObjectField",
			Fields: []ContentField{{
				Name:      "Required",
				NameLower: "required",
				Type:      "string",
				Required:  true,
			}, {
				Name:      "Optional",
				NameLower: "optional",
				Type:      "string",
			}},
		}},
	}

	mappings := map[string]string{
		"foosubject":   "FooSubject",
		"barpredicate": "BarPredicate",
	}
	got, err := DataFromSchema(testSchema, mappings, "0.1.2")
	if err != nil {
		t.Fatal(err.Error())
	}
	less := func(a, b ContentField) bool { return a.Name < b.Name }
	if d := cmp.Diff(want, got, cmpopts.SortSlices(less)); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

// TestExecuteTemplate_Success verifies that templating and code formatting
// are both applied
func TestExecuteTemplate_Success(t *testing.T) {
	// Template code. The result will be compiled.
	helloTemplate := `package main

import "fmt"
func main() {
    fmt.Println("Hello, {{.Name}}")
}`
	// Expected output is the formatted code
	expectedOutput := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
`

	// Create a temporary directory to store the generated files
	tempDir, err := os.MkdirTemp("", "test-execute-template-*")
	if err != nil {
		t.Fatalf("error creating temporary directory: %s", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a template file
	tmpTemplateFilename := "template.txt.tmpl"
	templateFile := filepath.Join(tempDir, tmpTemplateFilename)
	err = os.WriteFile(templateFile, []byte(helloTemplate), 0644)
	if err != nil {
		t.Fatalf("error creating template file: %s", err)
	}

	// Create a template set
	allTemplates := template.New("all")
	_, err = allTemplates.ParseFiles(templateFile)
	if err != nil {
		t.Fatalf("error parsing template file: %s", err)
	}

	// Execute the template
	outputFileName := filepath.Join(tempDir, "output.txt")
	err = executeTemplate(allTemplates, tmpTemplateFilename, outputFileName, struct{ Name string }{"World"})
	if err != nil {
		t.Fatalf("error executing template: %s", err)
	}

	// Check the output file contents
	output, err := os.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("error reading output file: %s", err)
	}
	if string(output) != expectedOutput {
		t.Errorf("unexpected output: got %q, want %q", output, expectedOutput)
	}
}

// TestExecuteTemplate_Error tests the error case of executeTemplate()
func TestExecuteTemplate_Error(t *testing.T) {
	// Create a temporary directory to store the generated files
	tempDir, err := os.MkdirTemp("", "test-execute-template-*")
	if err != nil {
		t.Fatalf("error creating temporary directory: %s", err)
	}
	defer os.RemoveAll(tempDir)

	// Create an template file, valid template but invalid go code
	templateFile := filepath.Join(tempDir, "template.txt")
	err = os.WriteFile(templateFile, []byte("Hello, {{.Name}}!"), 0644)
	if err != nil {
		t.Fatalf("error creating template file: %s", err)
	}

	// Create a template set
	allTemplates := template.New("all")
	_, err = allTemplates.ParseFiles(templateFile)
	if err != nil {
		t.Fatalf("error parsing template file: %s", err)
	}

	// Execute the template
	outputFileName := filepath.Join(tempDir, "output.txt")
	err = executeTemplate(allTemplates, "template.txt", outputFileName, struct{ Name string }{"World"})
	if err == nil {
		t.Fatal("expected error executing template, got nil")
	}
}

// TestValidateStringEnumAnyOf tests the validation of the string enum anyOf case.
func TestValidateStringEnumAnyOf(t *testing.T) {
	var boolType jsonschema.Types = 4
	var stringType jsonschema.Types = 32
	tests := []struct {
		name      string
		schema    jsonschema.Schema
		wantError string
	}{{
		name: "valid",
		schema: jsonschema.Schema{
			Location: "test_schema",
			AnyOf: []*jsonschema.Schema{
				{
					Location: "test_schema#/properties/content/anyOf/0",
					Types:    &stringType, // []string{"string"},
					Enum: &jsonschema.Enum{
						Values: []interface{}{"foo", "bar"},
					},
				},
				{
					Location: "test_schema#/properties/content/anyOf/1",
					Types:    &stringType, // []string{"string"},
				},
			},
		},
		wantError: "",
	}, {
		name: "enum missing",
		schema: jsonschema.Schema{
			Location: "test_schema",
			AnyOf: []*jsonschema.Schema{
				{
					Location: "test_schema#/properties/content/anyOf/0",
					Types:    &stringType, // []string{"string"},
				},
				{
					Location: "test_schema#/properties/content/anyOf/1",
					Types:    &stringType, // []string{"string"},
				},
			},
		},
		wantError: "one enum required when using anyOf for types test_schema: <nil>",
	}, {
		name: "too many enums",
		schema: jsonschema.Schema{
			Location: "test_schema",
			AnyOf: []*jsonschema.Schema{
				{
					Location: "test_schema#/properties/content/anyOf/0",
					Types:    &stringType, // []string{"string"},
					Enum: &jsonschema.Enum{
						Values: []interface{}{"foo", "bar"},
					},
				},
				{
					Location: "test_schema#/properties/content/anyOf/1",
					Types:    &stringType, // []string{"string"},
					Enum: &jsonschema.Enum{
						Values: []interface{}{"foo", "bar"},
					},
				},
			},
		},
		wantError: "only one enum allowed when using anyOf for types test_schema#/properties/content/anyOf/1: [string]",
	}, {
		name: "too many types",
		schema: jsonschema.Schema{
			Location: "test_schema",
			AnyOf: []*jsonschema.Schema{
				{
					Location: "test_schema#/properties/content/anyOf/0",
					Types:    &stringType, // []string{"string"},
					Enum: &jsonschema.Enum{
						Values: []interface{}{"foo", "bar"},
					},
				},
				{
					Location: "test_schema#/properties/content/anyOf/1",
					Types:    &stringType, // []string{"string"},
				},
				{
					Location: "test_schema#/properties/content/anyOf/2",
					Types:    &stringType, // []string{"string"},
				},
			},
		},
		wantError: "only two types allowed when using anyOf for content property in schema test_schema: <nil>",
	}, {
		name: "wrong types",
		schema: jsonschema.Schema{
			Location: "test_schema",
			AnyOf: []*jsonschema.Schema{
				{
					Location: "test_schema#/properties/content/anyOf/0",
					Types:    &stringType, // []string{"string"},
					Enum: &jsonschema.Enum{
						Values: []interface{}{"foo", "bar"},
					},
				},
				{
					Location: "test_schema#/properties/content/anyOf/1",
					Types:    &boolType, // []string{"bool"},
				},
			},
		},
		wantError: "only string allowed when using anyOf for types test_schema#/properties/content/anyOf/1: [boolean]",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateStringEnumAnyOf(&tc.schema)
			if err != nil {
				if tc.wantError == "" {
					t.Fatalf("didn't expected it to fail, but it did: %v", err)
				} else {
					// Check the error is what is expected
					if d := cmp.Diff(tc.wantError, err.Error()); d != "" {
						t.Errorf("args: diff(-want,+got):\n%s", d)
					}
				}
			}
			if err == nil {
				if tc.wantError != "" {
					t.Fatalf("expected an error, but go none")
				}
			}
		})
	}
}
