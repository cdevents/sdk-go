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
	"encoding/base64"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testXmlString  = "<xml>testData</xml>"
	testJsonString = "{\"testData\":\"testValue\"}"
)

var (
	eventWithNonJsonCustomData *ArtifactPackagedEvent
	eventWithJsonCustomData    *ArtifactPackagedEvent
	testXmlBytesB64            []byte
)

func init() {

	eventWithNonJsonCustomData, _ = NewArtifactPackagedEvent()
	eventWithNonJsonCustomData.CustomDataContentType = "application/xml"
	eventWithNonJsonCustomData.CustomData = []byte(testXmlString)

	eventWithJsonCustomData, _ = NewArtifactPackagedEvent()
	eventWithJsonCustomData.CustomDataContentType = "application/json"
	eventWithJsonCustomData.CustomData = []byte(testJsonString)

	testXmlBytes := []byte(testXmlString)
	testXmlBytesB64 = make([]byte, base64.StdEncoding.EncodedLen(len(testXmlBytes)))
	base64.StdEncoding.Encode(testXmlBytesB64, testXmlBytes)
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

	err := getCustomDataAs(eventWithNonJsonCustomData, receiver)
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

	err := getCustomDataAs(eventWithJsonCustomData, receiver)
	if err != nil {
		t.Fatalf("did not expect an error, got %v", err)
	}

	if d := cmp.Diff(expectedValue, receiver.TestData); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

func TestGetCustomDataAsJsonInvalidReceiver(t *testing.T) {

	receiver := &testWrongType{}
	expectedReceiver := &testWrongType{}

	err := getCustomDataAs(eventWithJsonCustomData, receiver)
	if err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if d := cmp.Diff(*expectedReceiver, *receiver); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

func TestSetCustomData(t *testing.T) {

	tests := []struct {
		name         string
		contentType  string
		data         interface{}
		expectedData []byte
	}{{
		name:         "json, bytes",
		contentType:  "application/json",
		data:         []byte(testJsonString),
		expectedData: []byte(testJsonString),
	}, {
		name:         "xml, bytes",
		contentType:  "application/xml",
		data:         []byte(testXmlString),
		expectedData: testXmlBytesB64,
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
