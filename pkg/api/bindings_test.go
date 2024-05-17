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
	"os"
	"testing"

	"github.com/cdevents/sdk-go/pkg/api"
	testapi "github.com/cdevents/sdk-go/pkg/api/v990"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

const testsFolder = "tests/examples"

type testData struct {
	TestValues []map[string]string `json:"testValues"`
}

var (
	// Examples Data
	testSource               = "/event/source/123"
	testSubjectId            = "mySubject123"
	testValue                = "testValue"
	testArtifactId           = "pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
	testDataJson             = testData{TestValues: []map[string]string{{"k1": "v1"}, {"k2": "v2"}}}
	testDataJsonUnmarshalled = map[string]any{
		"testValues": []any{map[string]any{"k1": string("v1")}, map[string]any{"k2": string("v2")}},
	}
	testDataXml  = []byte("<xml>testData</xml>")
	testChangeId = "myChange123"

	eventJsonCustomData             *testapi.FooSubjectBarPredicateEvent
	eventNonJsonCustomData          *testapi.FooSubjectBarPredicateEvent
	eventJsonCustomDataUnmarshalled *testapi.FooSubjectBarPredicateEvent

	eventJsonCustomDataFile         = "json_custom_data"
	eventImplicitJsonCustomDataFile = "implicit_json_custom_data"
	eventNonJsonCustomDataFile      = "non_json_custom_data"
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func setContext(event api.CDEventWriter, subjectId string) {
	event.SetSource(testSource)
	event.SetSubjectId(subjectId)
}

func init() {
	eventJsonCustomData, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventJsonCustomData, testSubjectId)
	eventJsonCustomData.SetSubjectReferenceField(&api.Reference{Id: testChangeId})
	eventJsonCustomData.SetSubjectPlainField(testValue)
	eventJsonCustomData.SetSubjectArtifactId(testArtifactId)
	eventJsonCustomData.SetSubjectObjectField(&api.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeId, Optional: testSource})
	err := eventJsonCustomData.SetCustomData("application/json", testDataJson)
	panicOnError(err)

	eventJsonCustomDataUnmarshalled, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventJsonCustomDataUnmarshalled, testSubjectId)
	eventJsonCustomDataUnmarshalled.SetSubjectReferenceField(&api.Reference{Id: testChangeId})
	eventJsonCustomDataUnmarshalled.SetSubjectPlainField(testValue)
	eventJsonCustomDataUnmarshalled.SetSubjectArtifactId(testArtifactId)
	eventJsonCustomDataUnmarshalled.SetSubjectObjectField(&api.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeId, Optional: testSource})
	err = eventJsonCustomDataUnmarshalled.SetCustomData("application/json", testDataJsonUnmarshalled)
	panicOnError(err)

	eventNonJsonCustomData, _ = testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventNonJsonCustomData, testSubjectId)
	eventNonJsonCustomData.SetSubjectReferenceField(&api.Reference{Id: testChangeId})
	eventNonJsonCustomData.SetSubjectPlainField(testValue)
	eventNonJsonCustomData.SetSubjectArtifactId(testArtifactId)
	eventNonJsonCustomData.SetSubjectObjectField(&api.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeId, Optional: testSource})
	err = eventNonJsonCustomData.SetCustomData("application/xml", testDataXml)
	panicOnError(err)
}

// TestAsCloudEvent produces a CloudEvent from a CDEvent using `AsCloudEvent`
// and then attempts to parse the CloudEvent payload back into a specific CDEvent
func TestAsCloudEvent(t *testing.T) {

	tests := []struct {
		name  string
		event api.CDEventReader
	}{{
		name:  "event with JSON custom data",
		event: eventJsonCustomData,
	}, {
		name:  "event with non-JSON custom data",
		event: eventNonJsonCustomData,
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
			if d := cmp.Diff(testSubjectId, ce.Context.GetSubject()); d != "" {
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
	_, err := api.AsCloudEvent(nil)
	if err == nil {
		t.Fatalf("expected it to fail, but it didn't")
	}
}

// TestAsJsonBytes renders a CDEvent as JSON and verifies it matches a
// manually crafted JSON for that event. The order of the fields in the
// rendered JSON depends on a number of factors, and is not deterministic
// so we must compare events unmarshalled to an interface
func TestAsJsonBytes(t *testing.T) {

	tests := []struct {
		name     string
		event    api.CDEvent
		fileName string
	}{{
		name:     "json custom data",
		event:    eventJsonCustomData,
		fileName: eventJsonCustomDataFile,
	}, {
		name:     "xml custom data",
		event:    eventNonJsonCustomData,
		fileName: eventNonJsonCustomDataFile,
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
			schema, url := tc.event.GetSchema()
			sch, err := jsonschema.CompileString(schema, url)
			if err != nil {
				t.Fatalf("Cannot compile jsonschema %s: %v", url, err)
			}
			var v interface{}
			if err := json.Unmarshal(eventBytes, &v); err != nil {
				t.Fatalf("Cannot unmarshal test json: %v", err)
			}
			err = sch.Validate(v)
			if err != nil {
				t.Fatalf("Failed to validate events %s", err)
			}
			// Then test that AsJsonBytes produces a matching JSON from the event
			obtainedJsonString, err := api.AsJsonBytes(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			err = json.Unmarshal([]byte(obtainedJsonString), &obtainedInteface)
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
	eventNoSource.SetSubjectId(testSubjectId)

	// mandatory subject id missing
	eventNoSubjectId, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventNoSubjectId.SetSource(testSource)

	// forced invalid version
	eventBadVersion, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventBadVersion.Context.Version = "invalid"

	// mandatory plainField and referenceField missing
	eventIncompleteSubject, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventIncompleteSubject.SetSource(testSource)
	eventIncompleteSubject.SetSubjectId(testSubjectId)

	// invalid source format in context
	eventInvalidSource, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventInvalidSource.SetSource("\\--##@@")

	// invalid source format in reference
	eventInvalidSourceReference, _ := testapi.NewFooSubjectBarPredicateEvent()
	eventInvalidSourceReference.SetSubjectReferenceField(
		&api.Reference{Id: "1234", Source: "\\--##@@"})

	// invalid format of purl
	eventInvalidPurl, _ := testapi.NewFooSubjectBarPredicateEvent()
	setContext(eventInvalidPurl, testSubjectId)
	eventInvalidPurl.SetSubjectArtifactId("not-a-valid-purl")

	// invalid event type
	eventInvalidType := &testapi.FooSubjectBarPredicateEvent{
		Context: api.Context{
			Type:    api.CDEventType{Subject: "not-a-valid-type"},
			Version: api.CDEventsSpecVersion,
		},
		Subject: api.FooSubjectBarPredicateSubject{
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
		event: eventNoSubjectId,
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

func TestAsJsonStringEmpty(t *testing.T) {
	obtainedJsonString, err := api.AsJsonString(nil)
	if err != nil {
		t.Fatalf("didn't expected it to fail, but it did: %v", err)
	}
	if d := cmp.Diff("", obtainedJsonString); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

func TestNewFromJsonString(t *testing.T) {

	tests := []struct {
		name     string
		event    api.CDEvent
		fileName string
	}{{
		name:     "json custom data",
		event:    eventJsonCustomDataUnmarshalled,
		fileName: eventJsonCustomDataFile,
	}, {
		name:     "json custom data implicit",
		event:    eventJsonCustomDataUnmarshalled,
		fileName: eventImplicitJsonCustomDataFile,
	}, {
		name:     "xml custom data",
		event:    eventNonJsonCustomData,
		fileName: eventNonJsonCustomDataFile,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Load the event from file
			eventBytes, err := os.ReadFile(testsFolder + string(os.PathSeparator) + tc.fileName + ".json")
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			obtainedEvent, err := testapi.NewFromJsonBytes(eventBytes)
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
	setContext(event, testSubjectId)
	event.SetSubjectReferenceField(&api.Reference{Id: testChangeId})
	event.SetSubjectPlainField(testValue)
	event.SetSubjectArtifactId(testArtifactId)
	event.SetSubjectObjectField(&api.FooSubjectBarPredicateSubjectContentObjectField{Required: testChangeId, Optional: testSource})
	err := event.SetCustomData("application/json", testDataJsonUnmarshalled)
	panicOnError(err)
	event.Context.Type.Version = eventVersion
	event.Context.Version = specVersion
	return event
}

func TestNewFromJsonBytes(t *testing.T) {

	minorVersion := testEventWithVersion("1.999.0", testapi.SpecVersion)
	patchVersion := testEventWithVersion("1.2.999", testapi.SpecVersion)
	pastPatchVersion := testEventWithVersion("1.2.0", testapi.SpecVersion)
	pastSpecVersion := testEventWithVersion("1.2.3", "0.1.0")

	tests := []struct {
		testFile    string
		description string
		wantError   string
		wantEvent   api.CDEvent
	}{{
		testFile:    "future_event_major_version",
		description: "A newer major version in the event is backward incompatible and cannot be parsed",
		wantError:   "sdk event version 1.2.3 not compatible with 999.0.0",
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
			obtained, err := testapi.NewFromJsonBytes(eventBytes)
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
