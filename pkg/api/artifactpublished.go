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
	// ArtifactPublished event
	ArtifactPublishedEventV1    CDEventType = "dev.cdevents.artifact.published.v1"
	artifactPublishedSchemaFile string      = "artifactpublished"
)

type ArtifactPublishedSubjectContent struct{}

type ArtifactPublishedSubject struct {
	SubjectBase
	Content ArtifactPublishedSubjectContent `json:"content"`
}

func (sc ArtifactPublishedSubject) GetEventType() CDEventType {
	return ArtifactPublishedEventV1
}

func (sc ArtifactPublishedSubject) GetSubjectType() SubjectType {
	return ArtifactSubjectType
}

type ArtifactPublishedEvent struct {
	Context Context                  `json:"context"`
	Subject ArtifactPublishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ArtifactPublishedEvent) GetType() CDEventType {
	return ArtifactPublishedEventV1
}

func (e ArtifactPublishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ArtifactPublishedEvent) GetId() string {
	return e.Context.Id
}

func (e ArtifactPublishedEvent) GetSource() string {
	return e.Context.Source
}

func (e ArtifactPublishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactPublishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactPublishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactPublishedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ArtifactPublishedEvent) GetCustomData() []byte {
	return e.CustomData
}

func (e ArtifactPublishedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ArtifactPublishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ArtifactPublishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactPublishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactPublishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactPublishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactPublishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ArtifactPublishedEvent) SetCustomData(contentType string, data interface{}) error {
	dataBytes, err := customDataBytes(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = dataBytes
	return nil
}

func (e ArtifactPublishedEvent) GetSchema() string {
	return artifactPublishedSchemaFile
}

func NewArtifactPublishedEvent() (*ArtifactPublishedEvent, error) {
	e := &ArtifactPublishedEvent{
		Context: Context{
			Type:    ArtifactPublishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ArtifactPublishedSubject{
			SubjectBase: SubjectBase{
				Type: ArtifactSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
