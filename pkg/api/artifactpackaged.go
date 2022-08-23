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
	"time"
)

const (
	// ArtifactPackaged event
	ArtifactPackagedEventV1    CDEventType = "dev.cdevents.artifact.packaged.v1"
	artifactPackagedSchemaFile string      = "artifactpackaged"
)

type ArtifactPackagedSubjectContent struct{}

type ArtifactPackagedSubject struct {
	SubjectBase
	Content ArtifactPackagedSubjectContent `json:"content"`
}

func (sc ArtifactPackagedSubject) GetEventType() CDEventType {
	return ArtifactPackagedEventV1
}

func (sc ArtifactPackagedSubject) GetSubjectType() SubjectType {
	return ArtifactSubjectType
}

type ArtifactPackagedEvent struct {
	Context Context                 `json:"context"`
	Subject ArtifactPackagedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ArtifactPackagedEvent) GetType() CDEventType {
	return ArtifactPackagedEventV1
}

func (e ArtifactPackagedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ArtifactPackagedEvent) GetId() string {
	return e.Context.Id
}

func (e ArtifactPackagedEvent) GetSource() string {
	return e.Context.Source
}

func (e ArtifactPackagedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactPackagedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactPackagedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactPackagedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ArtifactPackagedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactPackagedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactPackagedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactPackagedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactPackagedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e ArtifactPackagedEvent) GetSchema() string {
	return artifactPackagedSchemaFile
}

func newArtifactPackagedEvent() CDEvent {
	return &ArtifactPackagedEvent{
		Context: Context{
			Type:    ArtifactPackagedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ArtifactPackagedSubject{},
	}
}
