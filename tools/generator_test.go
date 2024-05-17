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
	"github.com/santhosh-tekuri/jsonschema/v5"
)

const testSchemaJson = "../pkg/api/tests/schemas/foosubjectbarpredicate.json"

var (
	testSchema      *jsonschema.Schema
	testSubject     = "FooSubject"
	testSubjectType = "fooSubject"
	testPredicate   = "BarPredicate"
	testVersion     = "1.2.3"
	testVersionName = "1_2_3"
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	var err error
	testSchema, err = jsonschema.Compile(testSchemaJson)
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
			Type:      "*FooSubjectBarPredicateSubjectContentObjectField",
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
	got, err := DataFromSchema(testSchema, mappings)
	if err != nil {
		t.Fatalf(err.Error())
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
