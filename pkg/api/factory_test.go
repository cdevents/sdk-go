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
	"testing"
	"time"

	"github.com/cdevents/sdk-go/pkg/api"
	cdevents "github.com/cdevents/sdk-go/pkg/api/v03"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

func init() {
	// Get the time once
	t := time.Now()
	timeNow = func() time.Time {
		return t
	}

	// Get the UUID once
	u, _ := uuid.NewRandom()
	uuidNewRandom = func() (uuid.UUID, error) {
		return u, nil
	}
}

func testUUID() string {
	u, _ := uuidNewRandom()
	return fmt.Sprintf("%v", u)
}

type testNewCDEventType struct {
	name          string
	eventType     string
	expectedEvent api.CDEvent
}

// tests is used in TestNewCDEvents. It's content is
// generated in zz_factory_tests.go
var (
	tests           []testNewCDEventType
	testContentType = "application/json"
	timeNow         = time.Now
	uuidNewRandom   = uuid.NewRandom
	testSpecVersion = "0.3.0"
)

func TestNewCDEvent(t *testing.T) {
	testDataJsonBytes, err := json.Marshal(testDataJson)
	if err != nil {
		t.Fatalf("didn't expected it to fail, but it did: %v", err)
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			event, err := cdevents.NewCDEvent(tc.eventType, testSpecVersion)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(tc.expectedEvent, event,
				cmpopts.IgnoreFields(api.Context{}, "Timestamp"),
				cmpopts.IgnoreFields(api.Context{}, "Id")); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check GetType
			if d := cmp.Diff(tc.eventType, event.GetType().String()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// CustomData set and get
			err = event.SetCustomData(testContentType, testDataJsonBytes)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			customDataRawGot, err := event.GetCustomDataRaw()
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(testDataJsonBytes, customDataRawGot); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			customDataGot, err := event.GetCustomData()
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(testDataJsonUnmarshalled, customDataGot); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			var customDataAsGot testData
			err = event.GetCustomDataAs(&customDataAsGot)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(testDataJson, customDataAsGot); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(event.GetCustomDataContentType(), testContentType); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Subject source set and get
			event.SetSubjectSource("randomSubjectSource")
			gotSubjectSource := event.GetSubjectSource()
			if d := cmp.Diff("randomSubjectSource", gotSubjectSource); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestNewCDEventFailed(t *testing.T) {

	_, err := cdevents.NewCDEvent(api.CDEventType{Subject: "not supported"}.String(), testSpecVersion)
	if err == nil {
		t.Fatalf("expected it to fail, but it didn't")
	}
}
