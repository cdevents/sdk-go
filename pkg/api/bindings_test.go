/*
Copyright 2022 The CDEvents Authors

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

package api_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cdevents/sdk-go/pkg/api"
	testapi "github.com/cdevents/sdk-go/pkg/api/v991"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const testsFolder = "tests-v99.1/examples"

type testData struct {
	TestValues []map[string]string `json:"testValues"`
}

var (
	// Examples Data
	testSource               = "/event/source/123"
	testSubjectID            = "mySubject123"
	testValue                = "testValue"
	testArtifactID           = "pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
	testInvalidArtifactID    = "not-in-purl-format"
	testDataJSON             = testData{TestValues: []map[string]string{{"k1": "v1"}, {"k2": "v2"}}}
	testDataJSONUnmarshalled = map[string]any{
		"testValues": []any{map[string]any{"k1": string("v1")}, map[string]any{"k2": string("v2")}},
	}
	testDataXML  = []byte("<xml>testData</xml>")
	testChangeID = "myChange123"

	// V04+ Examples Data
	testLinks                    api.EmbeddedLinksArray
	testContextID                = "5328c37f-bb7e-4bb7-84ea-9f5f85e4a7ce"
	testChainID                  = "4c8cb7dd-3448-41de-8768-eec704e2829b"
	testSchemaURI                = "https://myorg.com/schema/custom"
	testCustomSchemaJSONTemplate = `{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$id": "%s",
		"additionalProperties": true,
		"type": "object"
	}`
	testCustomSchemaJSON                 = fmt.Sprintf(testCustomSchemaJSONTemplate, testSchemaURI)
	testSchemaURIStricter                = "https://myorg.com/schema/stricter"
	testCustomSchemaJSONStricterTemplate = `{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$id": "%s",
		"additionalProperties": true,
		"type": "object",
		"properties": {
			"customData": {
				"type": "object",
				"additionalProperties": false,
				"properties": {
					"important": {
						"type": "string"
					}
				},
				"required": [
					"important"
				]
			}
		}
	}`
	testCustomSchemaJSONStricterJSON = fmt.Sprintf(testCustomSchemaJSONStricterTemplate, testSchemaURIStricter)
	testCustomSchemas                = map[string][]byte{
		testSchemaURI:         []byte(testCustomSchemaJSON),
		testSchemaURIStricter: []byte(testCustomSchemaJSONStricterJSON),
	}

	eventJSONCustomData             *testapi.FooSubjectBarPredicateEvent
	eventNonJSONCustomData          *testapi.FooSubjectBarPredicateEvent
	eventJSONCustomDataUnmarshalled *testapi.FooSubjectBarPredicateEvent
	eventJSONCustomDataCustomSchema *testapi.FooSubjectBarPredicateEvent
	eventInvalidArtifactIDFormat    *testapi.FooSubjectBarPredicateEvent

	eventJSONCustomDataFile         = "json_custom_data"
	eventImplicitJSONCustomDataFile = "implicit_json_custom_data"
	eventNonJSONCustomDataFile      = "non_json_custom_data"

	eventInvalidType = &testapi.FooSubjectBarPredicateEvent{
		Context: api.ContextV04{
			api.Context{
				Type: api.CDEventType{
					Subject:   "invalid",
					Predicate: "invalid",
					Version:   "#not@semver", // Invalid version format
				},
				Version: "9.9.9",
			},
			api.ContextLinks{},
			api.ContextCustom{},
		},
	}

	eventUnknownType = &testapi.FooSubjectBarPredicateEvent{
		Context: api.ContextV04{
			api.Context{
				Type: api.CDEventType{
					Subject:   "invalid", // Unknown subject
					Predicate: "invalid", // Unknown predicate
					Version:   "1.2.3",
				},
				Version: "9.9.9",
			},
			api.ContextLinks{},
			api.ContextCustom{},
		},
	}
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func setContext(event api.CDEventWriter, subjectID string) {
	event.SetSource(testSource)
	event.SetSubjectId(subjectID)
}

func setContextV04(event api.CDEventWriterV04, chainID, schemaURI bool) { //nolint: unparam
	if chainID {
		event.SetChainId(testChainID)
	}
	if schemaURI {
		event.SetSchemaUri(testSchemaURI)
	}
	event.SetLinks(testLinks)
}

func init() {
	// Set up test links
	tags := api.Tags{
		"foo1": "bar",
		"foo2": "bar",
	}
	reference := api.EventReference{
		ContextID: testContextID,
	}
	elr := api.NewEmbeddedLinkRelation()
	elr.SetTags(tags)
	elr.SetLinkKind("TRIGGER")
	elr.SetTarget(reference)
	elp := api.NewEmbeddedLinkPath()
	elp.SetTags(tags)
	elp.SetFrom(reference)
	ele := api.NewEmbeddedLinkEnd()
	ele.SetTags(tags)
	ele.SetFrom(reference)
	testLinks = api.EmbeddedLinksArray{
		elr, elp, ele,
	}

	setContext(eventInvalidType, testSubjectID)
	setContextV04(eventInvalidType, true, true)
	eventInvalidType.SetSubjectArtifactId(testArtifactID)

	setContext(eventUnknownType, testSubjectID)
	setContextV04(eventUnknownType, true, true)
	eventUnknownType.SetSubjectArtifactId(testArtifactID)

	eventInvalidArtifactIDFormat, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventInvalidArtifactIDFormat, testSubjectID)
	setContextV04(eventInvalidArtifactIDFormat, true, true)
	eventInvalidArtifactIDFormat.SetSubjectArtifactId(testInvalidArtifactID)

	eventJSONCustomData, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventJSONCustomData, testSubjectID)
	setContextV04(eventJSONCustomData, true, true)
	eventJSONCustomData.SetSubjectReferenceField(&api.Reference{ID: testChangeID})
	eventJSONCustomData.SetSubjectPlainField(testValue)
	eventJSONCustomData.SetSubjectArtifactId(testArtifactID)
	eventJSONCustomData.SetSubjectObjectField(&testapi.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeID, Optional: testSource})
	err := eventJSONCustomData.SetCustomData("application/json", testDataJSON)
	panicOnError(err)

	eventJSONCustomDataUnmarshalled, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventJSONCustomDataUnmarshalled, testSubjectID)
	setContextV04(eventJSONCustomDataUnmarshalled, true, true)
	eventJSONCustomDataUnmarshalled.SetSubjectReferenceField(&api.Reference{ID: testChangeID})
	eventJSONCustomDataUnmarshalled.SetSubjectPlainField(testValue)
	eventJSONCustomDataUnmarshalled.SetSubjectArtifactId(testArtifactID)
	eventJSONCustomDataUnmarshalled.SetSubjectObjectField(&testapi.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeID, Optional: testSource})
	err = eventJSONCustomDataUnmarshalled.SetCustomData("application/json", testDataJSONUnmarshalled)
	panicOnError(err)

	eventNonJSONCustomData, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventNonJSONCustomData, testSubjectID)
	setContextV04(eventNonJSONCustomData, true, true)
	eventNonJSONCustomData.SetSubjectReferenceField(&api.Reference{ID: testChangeID})
	eventNonJSONCustomData.SetSubjectPlainField(testValue)
	eventNonJSONCustomData.SetSubjectArtifactId(testArtifactID)
	eventNonJSONCustomData.SetSubjectObjectField(&testapi.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeID, Optional: testSource})
	err = eventNonJSONCustomData.SetCustomData("application/xml", testDataXML)
	panicOnError(err)

	eventJSONCustomDataCustomSchema, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventJSONCustomDataCustomSchema, testSubjectID)
	setContextV04(eventJSONCustomDataCustomSchema, true, true)
	eventJSONCustomDataCustomSchema.SetSchemaUri(testSchemaURIStricter)
	eventJSONCustomDataCustomSchema.SetSubjectReferenceField(&api.Reference{ID: testChangeID})
	eventJSONCustomDataCustomSchema.SetSubjectPlainField(testValue)
	eventJSONCustomDataCustomSchema.SetSubjectArtifactId(testArtifactID)
	eventJSONCustomDataCustomSchema.SetSubjectObjectField(&testapi.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeID, Optional: testSource})
	err = eventJSONCustomDataCustomSchema.SetCustomData("application/json", testDataJSON)
	panicOnError(err)

	for id, jsonBytes := range testCustomSchemas {
		err = api.LoadJsonSchema(id, jsonBytes)
		panicOnError(err)
	}
}

// TestAsCloudEvent produces a CloudEvent from a CDEvent using `AsCloudEvent`
// and then attempts to parse the CloudEvent payload back into a specific CDEvent
func TestAsCloudEvent(t *testing.T) {
	tests := []struct {
		name  string
		event api.CDEventReader
	}{{
		name:  "event with JSON custom data",
		event: eventJSONCustomData,
	}, {
		name:  "event with non-JSON custom data",
		event: eventNonJSONCustomData,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			payloadReceiver := &testapi.FooSubjectBarPredicateEvent{}
			ce, err := api.AsCloudEvent(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(tc.event.GetId(), ce.Context.GetID()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(testSubjectID, ce.Context.GetSubject()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(testSource, ce.Context.GetSource()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetType().String(), ce.Context.GetType()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			err = ce.DataAs(payloadReceiver)
			if err != nil {
				t.Fatalf("somehow cannot unmarshal test event %v, %v", ce, err)
			}
			if d := cmp.Diff(tc.event, payloadReceiver, cmpopts.IgnoreFields(api.CDEventCustomData{}, "CustomData")); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if tc.event.GetCustomDataContentType() == "application/json" {
				want := &testData{}
				got := &testData{}
				err = tc.event.GetCustomDataAs(want)
				if err != nil {
					t.Fatalf("didn't expected it to fail, but it did: %v", err)
				}
				err = payloadReceiver.GetCustomDataAs(got)
				if err != nil {
					t.Fatalf("cannot read custom data from parse event: %v", err)
				}
				if d := cmp.Diff(want, got); d != "" {
					t.Errorf("args: diff(-want,+got):\n%s", d)
				}
			} else {
				var want, got interface{}
				want, err = tc.event.GetCustomData()
				if err != nil {
					t.Fatalf("didn't expected it to fail, but it did: %v", err)
				}
				got, err = payloadReceiver.GetCustomData()
				if err != nil {
					t.Fatalf("cannot read custom data from parse event: %v", err)
				}
				if d := cmp.Diff(want, got); d != "" {
					t.Errorf("args: diff(-want,+got):\n%s", d)
				}
			}
		})
	}
}

func TestAsCloudEventInvalid(t *testing.T) {
	tests := []struct {
		name  string
		event api.CDEventReader
		error string
	}{{
		name:  "nil event",
		event: nil,
		error: "nil CDEvent cannot be rendered as CloudEvent",
	}, {
		name:  "event with invalid type",
		event: eventInvalidType,
		error: "cannot validate CDEvent Key: 'FooSubjectBarPredicateEventV2_2_3.Context.Context.Type.",
	}, {
		name:  "event with unknown type",
		event: eventUnknownType,
		error: "cannot validate CDEvent jsonschema validation failed with 'https://cdevents.dev/99.1.0/schema/foosubject-barpredicate-event#'",
	}, {
		name:  "event with invalid artifact id format",
		event: eventInvalidArtifactIDFormat,
		error: "cannot validate CDEvent Key: 'FooSubjectBarPredicateEventV2_2_3.Subject.Content.ArtifactId'",
	}, {
		name:  "does not match the custom schema",
		event: eventJSONCustomDataCustomSchema,
		error: "cannot validate CDEvent jsonschema validation failed with 'https://myorg.com/schema/stricter#",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := api.AsCloudEvent(tc.event)
			if err == nil {
				t.Fatalf("expected it to fail, but it didn't")
			}
			if !strings.HasPrefix(err.Error(), tc.error) {
				t.Errorf("error %s does not start with the expected prefix %s", err.Error(), tc.error)
			}
		})
	}
}

// TestAsJSONString renders a CDEvent as JSON and verifies it matches a
// manually crafted JSON for that event. The order of the fields in the
// rendered JSON depends on a number of factors, and is not deterministic
// so we must compare events unmarshalled to an interface
func TestAsJSONString(t *testing.T) {
	tests := []struct {
		name     string
		event    api.CDEvent
		fileName string
	}{{
		name:     "json custom data",
		event:    eventJSONCustomData,
		fileName: eventJSONCustomDataFile,
	}, {
		name:     "xml custom data",
		event:    eventNonJSONCustomData,
		fileName: eventNonJSONCustomDataFile,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var obtainedInteface, expectedInterface interface{}
			// Load the event from file
			eventBytes, err := os.ReadFile(testsFolder + string(os.PathSeparator) + tc.fileName + ".json")
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			// First validate that the test JSON compiles against the schema
			url, sch, err := tc.event.GetSchema()
			if err != nil {
				t.Fatalf("Cannot find jsonschema %s: %v", url, err)
			}
			var v interface{}
			if err := json.Unmarshal(eventBytes, &v); err != nil {
				t.Fatalf("Cannot unmarshal test json: %v", err)
			}
			err = sch.Validate(v)
			if err != nil {
				t.Fatalf("Failed to validate events %s", err)
			}
			// Then test that AsJSONBytes produces a matching JSON from the event
			obtainedJSONString, err := api.AsJSONBytes(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			err = json.Unmarshal(obtainedJSONString, &obtainedInteface)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			err = json.Unmarshal(eventBytes, &expectedInterface)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(expectedInterface, obtainedInteface); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestInvalidEvent(t *testing.T) {
	// mandatory source missing
	eventNoSource, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventNoSource.SetSubjectId(testSubjectID)

	// mandatory subject id missing
	eventNoSubjectID, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventNoSubjectID.SetSource(testSource)

	// forced invalid version
	eventBadVersion, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventBadVersion.Context.Version = "invalid"

	// mandatory plainField and referenceField missing
	eventIncompleteSubject, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventIncompleteSubject.SetSource(testSource)
	eventIncompleteSubject.SetSubjectId(testSubjectID)

	// invalid source format in context
	eventInvalidSource, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventInvalidSource.SetSource("\\--##@@")

	// invalid source format in reference
	eventInvalidSourceReference, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventInvalidSourceReference.SetSubjectReferenceField(
		&api.Reference{ID: "1234", Source: "\\--##@@"})

	// invalid format of purl
	eventInvalidPurl, _ := testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventInvalidPurl, testSubjectID)
	eventInvalidPurl.SetSubjectArtifactId("not-a-valid-purl")

	// invalid event type
	eventInvalidType := &testapi.FooSubjectBarPredicateEvent{
		Context: api.ContextV04{
			Context: api.Context{
				Type:    api.CDEventType{Subject: "not-a-valid-type"},
				Version: testapi.SpecVersion,
			},
		},
		Subject: testapi.FooSubjectBarPredicateSubject{
			SubjectBase: api.SubjectBase{
				Type: "notAValidSubjectType",
			},
		},
	}

	tests := []struct {
		name  string
		event api.CDEvent
	}{{
		name:  "missing source",
		event: eventNoSource,
	}, {
		name:  "missing subject id",
		event: eventNoSubjectID,
	}, {
		name:  "invalid version",
		event: eventBadVersion,
	}, {
		name:  "missing subject url",
		event: eventIncompleteSubject,
	}, {
		name:  "invalid source in context",
		event: eventInvalidSource,
	}, {
		name:  "invalid source in reference",
		event: eventInvalidSourceReference,
	}, {
		name:  "invalid purl in build finished",
		event: eventInvalidPurl,
	}, {
		name:  "invalid event type",
		event: eventInvalidType,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// First validate that the test JSON compiles against the schema
			err := api.Validate(tc.event)
			if err == nil {
				t.Fatalf("Expected validation to fail, but it succeeded instead")
			}
		})
	}
}

func TestAsJSONStringEmpty(t *testing.T) {
	obtainedJSONString, err := api.AsJSONString(nil)
	if err != nil {
		t.Fatalf("didn't expected it to fail, but it did: %v", err)
	}
	if d := cmp.Diff("", obtainedJSONString); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

func TestNewFromJsonString(t *testing.T) {
	tests := []struct {
		name     string
		event    api.CDEventV04
		fileName string
	}{{
		name:     "json custom data",
		event:    eventJSONCustomDataUnmarshalled,
		fileName: eventJSONCustomDataFile,
	}, {
		name:     "json custom data implicit",
		event:    eventJSONCustomDataUnmarshalled,
		fileName: eventImplicitJSONCustomDataFile,
	}, {
		name:     "xml custom data",
		event:    eventNonJSONCustomData,
		fileName: eventNonJSONCustomDataFile,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Load the event from file
			eventBytes, err := os.ReadFile(testsFolder + string(os.PathSeparator) + tc.fileName + ".json")
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			obtainedEvent, err := testapi.NewFromJSONBytes(eventBytes)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			// Check the context
			if d := cmp.Diff(tc.event.GetId(), obtainedEvent.GetId()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetVersion(), obtainedEvent.GetVersion()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetSource(), obtainedEvent.GetSource()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetTimestamp(), obtainedEvent.GetTimestamp()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetType(), obtainedEvent.GetType()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check V04 context
			if d := cmp.Diff(tc.event.GetChainId(), obtainedEvent.GetChainId()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetSchemaUri(), obtainedEvent.GetSchemaUri()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetLinks(), obtainedEvent.GetLinks()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check the subject
			if d := cmp.Diff(tc.event.GetSubject(), obtainedEvent.GetSubject()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check the data
			expectedData, err := tc.event.GetCustomData()
			if err != nil {
				t.Fatalf("cannot get data from test event %s", err)
			}
			obtainedData, err := obtainedEvent.GetCustomData()
			if err != nil {
				t.Fatalf("cannot get data from new event %s", err)
			}
			if d := cmp.Diff(expectedData, obtainedData); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestParseType(t *testing.T) {
	tests := []struct {
		name      string
		eventType string
		want      *api.CDEventType
		wantError string
	}{{
		name:      "valid",
		eventType: "dev.cdevents.foosubject.barpredicate.0.1.2-draft",
		want: &api.CDEventType{
			Subject:   "foosubject",
			Predicate: "barpredicate",
			Version:   "0.1.2-draft",
		},
		wantError: "",
	}, {
		name:      "invalid root",
		eventType: "foo.bar.subject.predicate.0.1.2-draft",
		want:      nil,
		wantError: "cannot parse event type foo.bar.subject.predicate.0.1.2-draft",
	}, {
		name:      "invalid format",
		eventType: "dev.cdevents.foosubject_barpredicate_0.1.2-draft",
		want:      nil,
		wantError: "cannot parse event type dev.cdevents.foosubject_barpredicate_0.1.2-draft",
	}, {
		name:      "unknown subject",
		eventType: "dev.cdevents.subject.barpredicate.0.1.2-draft",
		want: &api.CDEventType{
			Subject:   "subject",
			Predicate: "barpredicate",
			Version:   "0.1.2-draft",
		},
		wantError: "",
	}, {
		name:      "unknown predicate",
		eventType: "dev.cdevents.foosubject.predicate.0.1.2-draft",
		want: &api.CDEventType{
			Subject:   "foosubject",
			Predicate: "predicate",
			Version:   "0.1.2-draft",
		},
		wantError: "",
	}, {
		name:      "invalid version",
		eventType: "dev.cdevents.foosubject.barpredicate.0.1-draft",
		want:      nil,
		wantError: "invalid version format 0.1-draft",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obtained, err := api.ParseType(tc.eventType)
			if err != nil {
				if tc.wantError == "" {
					t.Fatalf("didn't expected it to fail, but it did: %v", err)
				} else {
					if d := cmp.Diff(tc.wantError, err.Error()); d != "" {
						t.Errorf("args: diff(-want,+got):\n%s", d)
					}
				}
			}

			// Check the subject
			if d := cmp.Diff(tc.want, obtained); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func testEventWithVersion(eventVersion string, specVersion string) *testapi.FooSubjectBarPredicateEvent {
	event, _ := testapi.NewFooSubjectBarPredicateEvent()
	setContext(event, testSubjectID)
	setContextV04(event, true, true)
	event.SetSubjectReferenceField(&api.Reference{ID: testChangeID})
	event.SetSubjectPlainField(testValue)
	event.SetSubjectArtifactId(testArtifactID)
	event.SetSubjectObjectField(&testapi.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeID, Optional: testSource})
	err := event.SetCustomData("application/json", testDataJSONUnmarshalled)
	panicOnError(err)
	event.Context.Type.Version = eventVersion
	event.Context.Version = specVersion
	return event
}

func TestNewFromJSONBytes(t *testing.T) {
	minorVersion := testEventWithVersion("2.999.1", testapi.SpecVersion)
	patchVersion := testEventWithVersion("2.2.999", testapi.SpecVersion)
	pastPatchVersion := testEventWithVersion("2.2.0", testapi.SpecVersion)
	pastSpecVersion := testEventWithVersion("2.2.3", "0.1.0")

	tests := []struct {
		testFile    string
		description string
		wantError   string
		wantEvent   api.CDEvent
	}{{
		testFile:    "future_event_major_version",
		description: "A newer major version in the event is backward incompatible and cannot be parsed",
		wantError:   "sdk event version 2.2.3 not compatible with 999.1.0",
	}, {
		testFile:    "future_event_minor_version",
		description: "A newer minor version in the event is compatible and can be parsed, data is lost",
		wantEvent:   minorVersion,
	}, {
		testFile:    "future_event_patch_version",
		description: "A newer patch version in the event is compatible and can be parsed",
		wantEvent:   patchVersion,
	}, {
		testFile:    "past_event_patch_version",
		description: "An older patch version in the event is compatible and can be parsed",
		wantEvent:   pastPatchVersion,
	}, {
		testFile:    "past_spec_patch_version",
		description: "An older patch version in the spec is compatible and can be parsed",
		wantEvent:   pastSpecVersion,
	}, {
		testFile:    "non_unmarshable",
		description: "The event has a valid context but fails to unmarshal",
		wantError:   `invalid character '@' after object key:value pair`,
	}, {
		testFile:    "unknown_type",
		description: "The event has a valid structure but unknown type",
		wantError:   "unknown event type dev.cdevents.foosubject.gazumped.0.1.0",
	}, {
		testFile:    "unparsable_context",
		description: "The context cannot be parsed, mandatory field is missing",
		wantError:   `invalid character '&' after object key:value pair`,
	}, {
		testFile:    "unparsable_type",
		description: "The context can be parsed, but the type is invalid",
		wantError:   "cannot parse event type dev.cdevents.foosubject_barpredicate_0.1.0",
	}}
	for _, tc := range tests {
		t.Run(tc.testFile, func(t *testing.T) {
			eventBytes, err := os.ReadFile(testsFolder + string(os.PathSeparator) + tc.testFile + ".json")
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			obtained, err := testapi.NewFromJSONBytes(eventBytes)
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
				} else {
					// Check the event is what is expected
					if d := cmp.Diff(tc.wantEvent, obtained, cmpopts.IgnoreFields(api.Context{}, "Id", "Timestamp")); d != "" {
						t.Errorf("args: diff(-want,+got):\n%s", d)
					}
				}
			}
		})
	}
}
