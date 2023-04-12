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
	"strings"
	"testing"

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
