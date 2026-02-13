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

package api_test

import (
	"encoding/json"
	"testing"
	"time"

	api "github.com/cdevents/sdk-go/pkg/api"
	"github.com/google/go-cmp/cmp"
)

// Test Context.GetVersion
func TestContextGetVersion(t *testing.T) {
	ctx := api.Context{
		SharedContext: api.SharedContext{
			Id:        "test-id",
			Source:    "test-source",
			Type:      api.CDEventType{Subject: "test", Predicate: "event", Version: "0.1.0"},
			Timestamp: time.Now(),
		},
		Version: "0.4.0",
	}

	got := ctx.GetVersion()
	want := "0.4.0"

	if got != want {
		t.Errorf("GetVersion() = %v, want %v", got, want)
	}
}

// Test Context.GetType
func TestContextGetType(t *testing.T) {
	eventType := api.CDEventType{Subject: "test", Predicate: "event", Version: "0.1.0"}
	ctx := api.Context{
		SharedContext: api.SharedContext{
			Id:        "test-id",
			Source:    "test-source",
			Type:      eventType,
			Timestamp: time.Now(),
		},
		Version: "0.4.0",
	}

	got := ctx.GetType()

	if d := cmp.Diff(eventType, got); d != "" {
		t.Errorf("GetType() diff(-want,+got):\n%s", d)
	}
}

// Test ContextV05.GetVersion
func TestContextV05GetVersion(t *testing.T) {
	ctx := api.ContextV05{
		SharedContext: api.SharedContext{
			Id:        "test-id",
			Source:    "test-source",
			Type:      api.CDEventType{Subject: "test", Predicate: "event", Version: "0.1.0"},
			Timestamp: time.Now(),
		},
		SpecVersion: "0.5.0",
	}

	got := ctx.GetVersion()
	want := "0.5.0"

	if got != want {
		t.Errorf("GetVersion() = %v, want %v", got, want)
	}
}

// Test ContextV05.GetType
func TestContextV05GetType(t *testing.T) {
	eventType := api.CDEventType{Subject: "test", Predicate: "event", Version: "0.1.0"}
	ctx := api.ContextV05{
		SharedContext: api.SharedContext{
			Id:        "test-id",
			Source:    "test-source",
			Type:      eventType,
			Timestamp: time.Now(),
		},
		SpecVersion: "0.5.0",
	}

	got := ctx.GetType()

	if d := cmp.Diff(eventType, got); d != "" {
		t.Errorf("GetType() diff(-want,+got):\n%s", d)
	}
}

// Test ContextForUnmarshalling.GetVersion
func TestContextForUnmarshallingGetVersion(t *testing.T) {
	tests := []struct {
		name    string
		ctx     api.ContextForUnmarshalling
		want    string
		wantUse bool
	}{
		{
			name: "v0.4 format with version field",
			ctx: api.ContextForUnmarshalling{
				Version: "0.4.0",
			},
			want:    "0.4.0",
			wantUse: false,
		},
		{
			name: "v0.5 format with specversion field",
			ctx: api.ContextForUnmarshalling{
				SpecVersion: "0.5.0",
			},
			want:    "0.5.0",
			wantUse: true,
		},
		{
			name: "both fields present, specversion takes precedence",
			ctx: api.ContextForUnmarshalling{
				Version:     "0.4.0",
				SpecVersion: "0.5.0",
			},
			want:    "0.5.0",
			wantUse: true,
		},
		{
			name:    "neither field present",
			ctx:     api.ContextForUnmarshalling{},
			want:    "",
			wantUse: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ctx.GetVersion()
			if got != tc.want {
				t.Errorf("GetVersion() = %v, want %v", got, tc.want)
			}

			gotUse := tc.ctx.UsesSpecVersion()
			if gotUse != tc.wantUse {
				t.Errorf("UsesSpecVersion() = %v, want %v", gotUse, tc.wantUse)
			}
		})
	}
}

// Test ContextForUnmarshalling.GetType
func TestContextForUnmarshallingGetType(t *testing.T) {
	eventType := api.CDEventType{Subject: "test", Predicate: "event", Version: "0.1.0"}
	ctx := api.ContextForUnmarshalling{
		SharedContext: api.SharedContext{
			Type: eventType,
		},
	}

	got := ctx.GetType()

	if d := cmp.Diff(eventType, got); d != "" {
		t.Errorf("GetType() diff(-want,+got):\n%s", d)
	}
}

// Test CDEventType.Root
func TestCDEventTypeRoot(t *testing.T) {
	tests := []struct {
		name      string
		eventType api.CDEventType
		want      string
	}{
		{
			name: "standard event",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "dev.cdevents",
		},
		{
			name: "custom event",
			eventType: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: "dev.cdeventsx",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.eventType.Root()
			if got != tc.want {
				t.Errorf("Root() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Test CDEventType.FQSubject
func TestCDEventTypeFQSubject(t *testing.T) {
	tests := []struct {
		name      string
		eventType api.CDEventType
		want      string
	}{
		{
			name: "standard event",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "pipelinerun",
		},
		{
			name: "custom event",
			eventType: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: "jenkins-build",
		},
		{
			name: "undefined subject",
			eventType: api.CDEventType{
				Subject:   "",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "<undefined-subject>",
		},
		{
			name: "custom with undefined subject",
			eventType: api.CDEventType{
				Subject:   "",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: "jenkins-<undefined-subject>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.eventType.FQSubject()
			if got != tc.want {
				t.Errorf("FQSubject() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Test CDEventType.String
func TestCDEventTypeString(t *testing.T) {
	tests := []struct {
		name      string
		eventType api.CDEventType
		want      string
	}{
		{
			name: "complete standard event",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "dev.cdevents.pipelinerun.started.0.3.0",
		},
		{
			name: "complete custom event",
			eventType: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: "dev.cdeventsx.jenkins-build.completed.1.0.0",
		},
		{
			name: "missing predicate",
			eventType: api.CDEventType{
				Subject: "pipelinerun",
				Version: "0.3.0",
			},
			want: "dev.cdevents.pipelinerun.<undefined-predicate>.0.3.0",
		},
		{
			name: "missing version",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
			},
			want: "dev.cdevents.pipelinerun.started.<undefined-version>",
		},
		{
			name:      "all fields missing",
			eventType: api.CDEventType{},
			want:      "dev.cdevents.<undefined-subject>.<undefined-predicate>.<undefined-version>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.eventType.String()
			if got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Test CDEventType.UnversionedString
func TestCDEventTypeUnversionedString(t *testing.T) {
	tests := []struct {
		name      string
		eventType api.CDEventType
		want      string
	}{
		{
			name: "standard event",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "dev.cdevents.pipelinerun.started",
		},
		{
			name: "custom event",
			eventType: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: "dev.cdeventsx.jenkins-build.completed",
		},
		{
			name: "missing predicate",
			eventType: api.CDEventType{
				Subject: "pipelinerun",
				Version: "0.3.0",
			},
			want: "dev.cdevents.pipelinerun.<undefined-predicate>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.eventType.UnversionedString()
			if got != tc.want {
				t.Errorf("UnversionedString() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Test CDEventType.Short
func TestCDEventTypeShort(t *testing.T) {
	tests := []struct {
		name      string
		eventType api.CDEventType
		want      string
	}{
		{
			name: "valid standard event",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "pipelinerun_started",
		},
		{
			name: "valid custom event",
			eventType: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: "jenkins-build_completed",
		},
		{
			name: "missing subject",
			eventType: api.CDEventType{
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: "<undefined-subject>_started",
		},
		{
			name: "missing predicate",
			eventType: api.CDEventType{
				Subject: "pipelinerun",
				Version: "0.3.0",
			},
			want: "",
		},
		{
			name:      "both missing",
			eventType: api.CDEventType{},
			want:      "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.eventType.Short()
			if got != tc.want {
				t.Errorf("Short() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Test CDEventType.IsCompatible
func TestCDEventTypeIsCompatible(t *testing.T) {
	tests := []struct {
		name       string
		type1      api.CDEventType
		type2      api.CDEventType
		compatible bool
	}{
		{
			name: "same major version",
			type1: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "1.2.0",
			},
			type2: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "1.3.0",
			},
			compatible: true,
		},
		{
			name: "different major version",
			type1: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "1.2.0",
			},
			type2: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "2.0.0",
			},
			compatible: false,
		},
		{
			name: "different subject",
			type1: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "1.2.0",
			},
			type2: api.CDEventType{
				Subject:   "taskrun",
				Predicate: "started",
				Version:   "1.2.0",
			},
			compatible: false,
		},
		{
			name: "different predicate",
			type1: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "1.2.0",
			},
			type2: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "finished",
				Version:   "1.2.0",
			},
			compatible: false,
		},
		{
			name: "v0 versions are compatible",
			type1: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.2.0",
			},
			type2: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			compatible: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.type1.IsCompatible(tc.type2)
			if got != tc.compatible {
				t.Errorf("IsCompatible() = %v, want %v", got, tc.compatible)
			}
		})
	}
}

// Test CDEventType.MarshalJSON
func TestCDEventTypeMarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		eventType api.CDEventType
		want      string
	}{
		{
			name: "standard event",
			eventType: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			want: `"dev.cdevents.pipelinerun.started.0.3.0"`,
		},
		{
			name: "custom event",
			eventType: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			want: `"dev.cdeventsx.jenkins-build.completed.1.0.0"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.eventType)
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			if string(got) != tc.want {
				t.Errorf("MarshalJSON() = %v, want %v", string(got), tc.want)
			}
		})
	}
}

// Test CDEventType.UnmarshalJSON
func TestCDEventTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		json      string
		want      api.CDEventType
		wantError bool
	}{
		{
			name: "valid standard event",
			json: `"dev.cdevents.pipelinerun.started.0.3.0"`,
			want: api.CDEventType{
				Subject:   "pipelinerun",
				Predicate: "started",
				Version:   "0.3.0",
			},
			wantError: false,
		},
		{
			name: "valid custom event",
			json: `"dev.cdeventsx.jenkins-build.completed.1.0.0"`,
			want: api.CDEventType{
				Subject:   "build",
				Predicate: "completed",
				Version:   "1.0.0",
				Custom:    "jenkins",
			},
			wantError: false,
		},
		{
			name:      "null value",
			json:      `null`,
			want:      api.CDEventType{},
			wantError: false,
		},
		{
			name:      "empty string",
			json:      `""`,
			want:      api.CDEventType{},
			wantError: false,
		},
		{
			name:      "invalid format",
			json:      `"invalid.format"`,
			want:      api.CDEventType{},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got api.CDEventType
			err := json.Unmarshal([]byte(tc.json), &got)
			if tc.wantError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.wantError {
				if d := cmp.Diff(tc.want, got); d != "" {
					t.Errorf("UnmarshalJSON() diff(-want,+got):\n%s", d)
				}
			}
		})
	}
}

// Test NewEmbeddedLinkPath
func TestNewEmbeddedLinkPath(t *testing.T) {
	link := api.NewEmbeddedLinkPath()

	if link.GetLinkType() != api.LinkTypePath {
		t.Errorf("GetLinkType() = %v, want %v", link.GetLinkType(), api.LinkTypePath)
	}

	// Test setters
	tags := api.Tags{"key": "value"}
	link.SetTags(tags)
	if d := cmp.Diff(tags, link.GetTags()); d != "" {
		t.Errorf("GetTags() diff(-want,+got):\n%s", d)
	}

	from := api.EventReference{ContextId: "test-id"}
	link.SetFrom(from)
	if d := cmp.Diff(from, link.GetFrom()); d != "" {
		t.Errorf("GetFrom() diff(-want,+got):\n%s", d)
	}
}

// Test NewEmbeddedLinkEnd
func TestNewEmbeddedLinkEnd(t *testing.T) {
	link := api.NewEmbeddedLinkEnd()

	if link.GetLinkType() != api.LinkTypeEnd {
		t.Errorf("GetLinkType() = %v, want %v", link.GetLinkType(), api.LinkTypeEnd)
	}

	// Test setters
	tags := api.Tags{"key": "value"}
	link.SetTags(tags)
	if d := cmp.Diff(tags, link.GetTags()); d != "" {
		t.Errorf("GetTags() diff(-want,+got):\n%s", d)
	}

	from := api.EventReference{ContextId: "test-id"}
	link.SetFrom(from)
	if d := cmp.Diff(from, link.GetFrom()); d != "" {
		t.Errorf("GetFrom() diff(-want,+got):\n%s", d)
	}
}

// Test NewEmbeddedLinkRelation
func TestNewEmbeddedLinkRelation(t *testing.T) {
	link := api.NewEmbeddedLinkRelation()

	if link.GetLinkType() != api.LinkTypeRelation {
		t.Errorf("GetLinkType() = %v, want %v", link.GetLinkType(), api.LinkTypeRelation)
	}

	// Test setters
	tags := api.Tags{"key": "value"}
	link.SetTags(tags)
	if d := cmp.Diff(tags, link.GetTags()); d != "" {
		t.Errorf("GetTags() diff(-want,+got):\n%s", d)
	}

	target := api.EventReference{ContextId: "test-id"}
	link.SetTarget(target)
	if d := cmp.Diff(target, link.GetTarget()); d != "" {
		t.Errorf("GetTarget() diff(-want,+got):\n%s", d)
	}

	linkKind := "depends"
	link.SetLinkKind(linkKind)
	if link.GetLinkKind() != linkKind {
		t.Errorf("GetLinkKind() = %v, want %v", link.GetLinkKind(), linkKind)
	}
}

// Test EmbeddedLinksArray.UnmarshalJSON
func TestEmbeddedLinksArrayUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		json      string
		wantLen   int
		wantTypes []api.LinkType
		wantError bool
	}{
		{
			name: "mixed link types",
			json: `[
				{"linkType":"PATH","from":{"contextId":"id1"},"tags":{}},
				{"linkType":"END","from":{"contextId":"id2"},"tags":{}},
				{"linkType":"RELATION","target":{"contextId":"id3"},"linkKind":"depends","tags":{}}
			]`,
			wantLen:   3,
			wantTypes: []api.LinkType{api.LinkTypePath, api.LinkTypeEnd, api.LinkTypeRelation},
			wantError: false,
		},
		{
			name: "path only",
			json: `[
				{"linkType":"PATH","from":{"contextId":"id1"},"tags":{}}
			]`,
			wantLen:   1,
			wantTypes: []api.LinkType{api.LinkTypePath},
			wantError: false,
		},
		{
			name: "end only",
			json: `[
				{"linkType":"END","from":{"contextId":"id1"},"tags":{}}
			]`,
			wantLen:   1,
			wantTypes: []api.LinkType{api.LinkTypeEnd},
			wantError: false,
		},
		{
			name: "relation only",
			json: `[
				{"linkType":"RELATION","target":{"contextId":"id1"},"linkKind":"depends","tags":{}}
			]`,
			wantLen:   1,
			wantTypes: []api.LinkType{api.LinkTypeRelation},
			wantError: false,
		},
		{
			name:      "empty array",
			json:      `[]`,
			wantLen:   0,
			wantTypes: []api.LinkType{},
			wantError: false,
		},
		{
			name: "invalid link type",
			json: `[
				{"linkType":"INVALID","from":{"contextId":"id1"},"tags":{}}
			]`,
			wantError: true,
		},
		{
			name:      "malformed json",
			json:      `[{"linkType":"PATH"`,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var links api.EmbeddedLinksArray
			err := json.Unmarshal([]byte(tc.json), &links)

			if tc.wantError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.wantError {
				if len(links) != tc.wantLen {
					t.Errorf("len(links) = %v, want %v", len(links), tc.wantLen)
				}

				for i, link := range links {
					if i < len(tc.wantTypes) && link.GetLinkType() != tc.wantTypes[i] {
						t.Errorf("link[%d].GetLinkType() = %v, want %v", i, link.GetLinkType(), tc.wantTypes[i])
					}
				}
			}
		})
	}
}

// Test GetCustomData with base64 encoding
func TestGetCustomDataBase64(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		data        interface{}
		want        interface{}
		wantError   bool
	}{
		{
			name:        "base64 encoded string",
			contentType: "application/octet-stream",
			data:        "SGVsbG8gV29ybGQ=", // "Hello World" in base64
			want:        []byte("Hello World"),
			wantError:   false,
		},
		{
			name:        "invalid base64",
			contentType: "application/octet-stream",
			data:        "not-valid-base64!!!",
			want:        nil,
			wantError:   true,
		},
		{
			name:        "json string should error",
			contentType: "application/json",
			data:        "string-data",
			want:        nil,
			wantError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := api.GetCustomData(tc.contentType, tc.data)

			if tc.wantError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.wantError {
				if d := cmp.Diff(tc.want, got); d != "" {
					t.Errorf("GetCustomData() diff(-want,+got):\n%s", d)
				}
			}
		})
	}
}

// Test GetCustomDataRaw with nil data
func TestGetCustomDataRawNil(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		data        interface{}
		wantError   bool
	}{
		{
			name:        "nil with json content type",
			contentType: "application/json",
			data:        nil,
			wantError:   false,
		},
		{
			name:        "nil with xml content type",
			contentType: "application/xml",
			data:        nil,
			wantError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := api.GetCustomDataRaw(tc.contentType, tc.data)

			if tc.wantError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

// Test CheckCustomData with empty content type
func TestCheckCustomDataEmptyContentType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		data        interface{}
		wantError   bool
	}{
		{
			name:        "empty content type with interface",
			contentType: "",
			data:        map[string]string{"key": "value"},
			wantError:   false,
		},
		{
			name:        "empty content type with bytes",
			contentType: "",
			data:        []byte("data"),
			wantError:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := api.CheckCustomData(tc.contentType, tc.data)

			if tc.wantError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

// Test SubjectType.String
func TestSubjectTypeString(t *testing.T) {
	st := api.SubjectType("testSubject")
	want := "testSubject"

	got := st.String()
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

// Test CDEventCustomDataEncoding.String
func TestCDEventCustomDataEncodingString(t *testing.T) {
	encoding := api.CDEventCustomDataEncoding("base64")
	want := "base64"

	got := encoding.String()
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
