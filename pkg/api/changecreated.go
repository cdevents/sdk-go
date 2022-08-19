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
	// ChangeCreated event
	ChangeCreatedEventV1    CDEventType = "dev.cdevents.change.created.v1"
	changeCreatedSchemaFile string      = "changecreated"
)

type ChangeCreatedSubjectContent struct{}

type ChangeCreatedSubject struct {
	SubjectBase
	Content ChangeCreatedSubjectContent `json:"content"`
}

func (sc ChangeCreatedSubject) GetEventType() CDEventType {
	return ChangeCreatedEventV1
}

func (sc ChangeCreatedSubject) GetSubjectType() SubjectType {
	return ChangeSubjectType
}

type ChangeCreatedEvent struct {
	Context Context              `json:"context"`
	Subject ChangeCreatedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ChangeCreatedEvent) GetType() CDEventType {
	return ChangeCreatedEventV1
}

func (e ChangeCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeCreatedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ChangeCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ChangeCreatedEvent) GetSchema() string {
	return changeCreatedSchemaFile
}

func newChangeCreatedEvent() CDEvent {
	return &ChangeCreatedEvent{
		Context: Context{
			Type:    ChangeCreatedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeCreatedSubject{
			SubjectBase: SubjectBase{
				Type: ChangeSubjectType,
			},
		},
	}
}
