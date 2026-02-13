/*
Copyright 2026 The CDEvents Authors

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

package v05_test

import (
	"testing"

	"github.com/cdevents/sdk-go/pkg/api"
	cdevents "github.com/cdevents/sdk-go/pkg/api/v05"
	"github.com/google/go-cmp/cmp"
)

func TestNewFromJsonString(t *testing.T) {
	// Create a valid event JSON string
	event, err := cdevents.NewBuildStartedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetSource("/test/source")
	event.SetSubjectId("test-build-123")

	// Convert to JSON string
	jsonBytes, err := api.AsJsonBytes(event)
	if err != nil {
		t.Fatalf("failed to convert event to JSON: %v", err)
	}
	jsonString := string(jsonBytes)

	// Test NewFromJsonString
	parsedEvent, err := cdevents.NewFromJsonString(jsonString)
	if err != nil {
		t.Fatalf("NewFromJsonString failed: %v", err)
	}

	// Verify the parsed event matches the original
	if d := cmp.Diff(event.GetId(), parsedEvent.GetId()); d != "" {
		t.Errorf("ID mismatch: diff(-want,+got):\n%s", d)
	}
	if d := cmp.Diff(event.GetSource(), parsedEvent.GetSource()); d != "" {
		t.Errorf("Source mismatch: diff(-want,+got):\n%s", d)
	}
	if d := cmp.Diff(event.GetType(), parsedEvent.GetType()); d != "" {
		t.Errorf("Type mismatch: diff(-want,+got):\n%s", d)
	}
	if d := cmp.Diff(event.GetSubjectId(), parsedEvent.GetSubjectId()); d != "" {
		t.Errorf("SubjectId mismatch: diff(-want,+got):\n%s", d)
	}
}

func TestNewFromJsonStringInvalid(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr   string
		wantError bool
	}{
		{
			name:      "empty string",
			jsonStr:   "",
			wantError: true,
		},
		{
			name:      "invalid JSON",
			jsonStr:   "{invalid json}",
			wantError: true,
		},
		{
			name:      "valid JSON but not a CDEvent",
			jsonStr:   `{"foo": "bar"}`,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := cdevents.NewFromJsonString(tc.jsonStr)
			if tc.wantError && err == nil {
				t.Error("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
