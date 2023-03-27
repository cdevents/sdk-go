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

package api

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testXmlString = "<xml>testData</xml>"
)

var (
	testObject                               = testType{TestData: "testValue"}
	testJsonString                           []byte
	eventWithNonJsonCustomData               *ArtifactPackagedEvent
	eventWithInterfaceJsonCustomData         *ArtifactPackagedEvent
	eventWithInterfaceJsonImplicitCustomData *ArtifactPackagedEvent
	eventWithJsonCustomData                  *ArtifactPackagedEvent
	eventWithJsonImplicitCustomData          *ArtifactPackagedEvent
)

func init() {
	var err error
	testJsonString, err = json.Marshal(testObject)
	panicOnError(err)

	eventWithNonJsonCustomData, _ = NewArtifactPackagedEvent()
	eventWithNonJsonCustomData.CustomDataContentType = "application/xml"
	eventWithNonJsonCustomData.CustomData = []byte(testXmlString)

	eventWithJsonCustomData, _ = NewArtifactPackagedEvent()
	eventWithJsonCustomData.CustomDataContentType = "application/json"
	eventWithJsonCustomData.CustomData = testJsonString

	eventWithJsonImplicitCustomData, _ = NewArtifactPackagedEvent()
	eventWithJsonImplicitCustomData.CustomDataContentType = ""
	eventWithJsonImplicitCustomData.CustomData = testJsonString

	eventWithInterfaceJsonCustomData, _ = NewArtifactPackagedEvent()
	eventWithInterfaceJsonCustomData.CustomDataContentType = "application/json"
	eventWithInterfaceJsonCustomData.CustomData = testObject

	eventWithInterfaceJsonImplicitCustomData, _ = NewArtifactPackagedEvent()
	eventWithInterfaceJsonImplicitCustomData.CustomDataContentType = ""
	eventWithInterfaceJsonImplicitCustomData.CustomData = testObject
}

type testType struct {
	TestData string `json:"testData,omitempty"`
}

type testWrongType struct {
	WrongTestData string `json:"wrongTestData,omitempty"`
}

func TestGetCustomDataAsNonJson(t *testing.T) {

	receiver := &testType{}
	expectedError := "cannot unmarshal content-type application/xml"

	err := GetCustomDataAs(eventWithNonJsonCustomData, receiver)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if d := cmp.Diff(expectedError, err.Error()); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

func TestGetCustomDataAsJson(t *testing.T) {

	receiver := &testType{}
	expectedValue := "testValue"

	tests := []struct {
		name  string
		event CDEvent
	}{{
		name:  "json bytes",
		event: eventWithJsonCustomData,
	}, {
		name:  "json bytes, implicit",
		event: eventWithJsonImplicitCustomData,
	}, {
		name:  "interface",
		event: eventWithInterfaceJsonCustomData,
	}, {
		name:  "interface, implicit",
		event: eventWithInterfaceJsonImplicitCustomData,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := GetCustomDataAs(tc.event, receiver)
			if err != nil {
				t.Fatalf("did not expect an error, got %v", err)
			}

			if d := cmp.Diff(expectedValue, receiver.TestData); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestGetCustomDataAsJsonInvalidReceiver(t *testing.T) {

	receiver := &testWrongType{}
	expectedReceiver := &testWrongType{}

	tests := []struct {
		name  string
		event CDEvent
	}{{
		name:  "json bytes",
		event: eventWithJsonCustomData,
	}, {
		name:  "json bytes, implicit",
		event: eventWithJsonImplicitCustomData,
	}, {
		name:  "interface",
		event: eventWithInterfaceJsonCustomData,
	}, {
		name:  "interface, implicit",
		event: eventWithInterfaceJsonImplicitCustomData,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := GetCustomDataAs(tc.event, receiver)
			if err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}

			if d := cmp.Diff(*expectedReceiver, *receiver); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestSetCustomData(t *testing.T) {

	tests := []struct {
		name         string
		contentType  string
		data         interface{}
		expectedData interface{}
	}{{
		name:         "json, json bytes",
		contentType:  "application/json",
		data:         testJsonString,
		expectedData: testJsonString,
	}, {
		name:         "xml, xml bytes",
		contentType:  "application/xml",
		data:         []byte(testXmlString),
		expectedData: []byte(testXmlString),
	}, {
		name:         "json, interface",
		contentType:  "application/json",
		data:         testObject,
		expectedData: testObject,
	}, {
		name:         "empty, interface",
		contentType:  "",
		data:         testType{TestData: "testValue"},
		expectedData: testType{TestData: "testValue"},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e, _ := NewArtifactPackagedEvent()
			err := e.SetCustomData(tc.contentType, tc.data)
			if err != nil {
				t.Fatalf("expected to set the custom data, but got %v", err)
			}

			if d := cmp.Diff(tc.expectedData, e.CustomData); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestSetCustomDataInvalid(t *testing.T) {
	e, _ := NewArtifactPackagedEvent()
	err := e.SetCustomData("application/xml", testType{TestData: "testValue"})
	if err == nil {
		t.Fatalf("did not expect this to work, but it did")
	}
}

func TestGetCustomData(t *testing.T) {

	tests := []struct {
		name         string
		contentType  string
		data         interface{}
		expectedData interface{}
	}{{
		name:         "json, json bytes",
		contentType:  "application/json",
		data:         testJsonString,
		expectedData: map[string]any{"testData": string("testValue")},
	}, {
		name:         "xml, xml bytes",
		contentType:  "application/xml",
		data:         []byte(testXmlString),
		expectedData: []byte(testXmlString),
	}, {
		name:         "json, interface",
		contentType:  "application/json",
		data:         testObject,
		expectedData: testObject,
	}, {
		name:         "empty, json bytes",
		contentType:  "",
		data:         testJsonString,
		expectedData: map[string]any{"testData": string("testValue")},
	}, {
		name:         "empty, interface",
		contentType:  "application/json",
		data:         testObject,
		expectedData: testObject,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e, _ := NewArtifactPackagedEvent()
			err := e.SetCustomData(tc.contentType, tc.data)
			if err != nil {
				t.Fatalf("expected to set the custom data, but got %v", err)
			}
			data, err := e.GetCustomData()
			if err != nil {
				t.Fatalf("%v", err)
			}

			if d := cmp.Diff(tc.expectedData, data); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestGetCustomDataInvalidJson(t *testing.T) {
	e, _ := NewArtifactPackagedEvent()
	data := testType{TestData: "testValue"}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("somehow could not marshal %v: %v", data, err)
	}
	err = e.SetCustomData("application/json", dataBytes[:len(dataBytes)-2])
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = e.GetCustomData()
	if err == nil {
		t.Errorf("expected error from broken data, got nil")
	}
}

func TestGetCustomDataXmlNotBytes(t *testing.T) {
	e, _ := NewArtifactPackagedEvent()
	data := testType{TestData: "testValue"}
	// Set using "application/json", else it won't be allowed
	err := e.SetCustomData("application/json", data)
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Override content type to XML
	e.CustomDataContentType = "application/xml"
	_, err = e.GetCustomData()
	if err == nil {
		t.Errorf("expected error from broken data, got nil")
	}
}

func TestGetCustomDataRaw(t *testing.T) {

	tests := []struct {
		name         string
		contentType  string
		data         interface{}
		expectedData interface{}
	}{{
		name:         "json, json bytes",
		contentType:  "application/json",
		data:         testJsonString,
		expectedData: testJsonString,
	}, {
		name:         "json, xml bytes (invalid)",
		contentType:  "application/json",
		data:         []byte(testXmlString),
		expectedData: []byte(testXmlString),
	}, {
		name:         "xml, xml bytes",
		contentType:  "application/xml",
		data:         []byte(testXmlString),
		expectedData: []byte(testXmlString),
	}, {
		name:         "json, interface",
		contentType:  "application/json",
		data:         testType{TestData: "testValue"},
		expectedData: []byte(testJsonString),
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e, _ := NewArtifactPackagedEvent()
			err := e.SetCustomData(tc.contentType, tc.data)
			if err != nil {
				t.Fatalf("expected to set the custom data, but got %v", err)
			}
			data, err := e.GetCustomDataRaw()
			if err != nil {
				t.Fatalf("%v", err)
			}

			if d := cmp.Diff(tc.expectedData, data); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestGetCustomDataRawXmlNotBytes(t *testing.T) {
	e, _ := NewArtifactPackagedEvent()
	data := testType{TestData: "testValue"}
	// Set using "application/json", else it won't be allowed
	err := e.SetCustomData("application/json", data)
	if err != nil {
		t.Fatalf("%v", err)
	}
	// Override content type to XML
	e.CustomDataContentType = "application/xml"
	_, err = e.GetCustomDataRaw()
	if err == nil {
		t.Fatalf("expected error from broken data, got nil")
	}
}

func TestCDEventTypeFromString(t *testing.T) {
	tests := []struct {
		name      string
		eventType string
		want      *CDEventType
		wantError bool
	}{{
		name:      "parses",
		eventType: "dev.cdevents.a.b.123.a-da#@#",
		want: &CDEventType{
			Subject:   "a",
			Predicate: "b",
			Version:   "123.a-da#@#",
		},
		wantError: false,
	}, {
		name:      "missing version",
		eventType: "dev.cdevents.a.b",
		want:      nil,
		wantError: true,
	}, {
		name:      "wrong beginning",
		eventType: "dev.wrong.a.b.version",
		want:      nil,
		wantError: true,
	}, {
		name:      "invalid character",
		eventType: "dev.cdevents.a1.b2.version",
		want:      nil,
		wantError: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CDEventTypeFromString(tc.eventType)
			if tc.wantError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("did not expect error, but got %s", err)
			}

			if d := cmp.Diff(tc.want, got); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}
