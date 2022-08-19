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
	// EnvironmentCreated event
	EnvironmentCreatedEventV1    CDEventType = "dev.cdevents.environment.created.v1"
	environmentCreatedSchemaFile string      = "environmentcreated"
)

type EnvironmentCreatedSubjectContent struct{}

type EnvironmentCreatedSubject struct {
	SubjectBase
	Content EnvironmentCreatedSubjectContent `json:"content"`
}

func (sc EnvironmentCreatedSubject) GetEventType() CDEventType {
	return EnvironmentCreatedEventV1
}

func (sc EnvironmentCreatedSubject) GetSubjectType() SubjectType {
	return EnvironmentSubjectType
}

type EnvironmentCreatedEvent struct {
	Context Context                   `json:"context"`
	Subject EnvironmentCreatedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e EnvironmentCreatedEvent) GetType() CDEventType {
	return EnvironmentCreatedEventV1
}

func (e EnvironmentCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e EnvironmentCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e EnvironmentCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e EnvironmentCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e EnvironmentCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e EnvironmentCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e EnvironmentCreatedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *EnvironmentCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *EnvironmentCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *EnvironmentCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *EnvironmentCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *EnvironmentCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *EnvironmentCreatedEvent) GetSchema() string {
	return environmentCreatedSchemaFile
}

func newEnvironmentCreatedEvent() CDEvent {
	return &EnvironmentCreatedEvent{
		Context: Context{
			Type:    EnvironmentCreatedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: EnvironmentCreatedSubject{},
	}
}
