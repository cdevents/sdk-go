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
	// EnvironmentModified event
	EnvironmentModifiedEventV1    CDEventType = "dev.cdevents.environment.modified.v1"
	environmentModifiedSchemaFile string      = "environmentmodified"
)

type EnvironmentModifiedSubjectContent struct{}

type EnvironmentModifiedSubject struct {
	SubjectBase
	Content EnvironmentModifiedSubjectContent `json:"content"`
}

func (sc EnvironmentModifiedSubject) GetEventType() CDEventType {
	return EnvironmentModifiedEventV1
}

func (sc EnvironmentModifiedSubject) GetSubjectType() SubjectType {
	return EnvironmentSubjectType
}

type EnvironmentModifiedEvent struct {
	Context Context                    `json:"context"`
	Subject EnvironmentModifiedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e EnvironmentModifiedEvent) GetType() CDEventType {
	return EnvironmentModifiedEventV1
}

func (e EnvironmentModifiedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e EnvironmentModifiedEvent) GetId() string {
	return e.Context.Id
}

func (e EnvironmentModifiedEvent) GetSource() string {
	return e.Context.Source
}

func (e EnvironmentModifiedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e EnvironmentModifiedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e EnvironmentModifiedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e EnvironmentModifiedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *EnvironmentModifiedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *EnvironmentModifiedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *EnvironmentModifiedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *EnvironmentModifiedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *EnvironmentModifiedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *EnvironmentModifiedEvent) GetSchema() string {
	return environmentModifiedSchemaFile
}

func newEnvironmentModifiedEvent() CDEvent {
	return &EnvironmentModifiedEvent{
		Context: Context{
			Type:    EnvironmentModifiedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: EnvironmentModifiedSubject{},
	}
}
