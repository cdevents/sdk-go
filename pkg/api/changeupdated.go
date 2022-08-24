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
	// ChangeUpdated event
	ChangeUpdatedEventV1    CDEventType = "dev.cdevents.change.updated.v1"
	changeUpdatedSchemaFile string      = "changeupdated"
)

type ChangeUpdatedSubjectContent struct{}

type ChangeUpdatedSubject struct {
	SubjectBase
	Content ChangeUpdatedSubjectContent `json:"content"`
}

func (sc ChangeUpdatedSubject) GetEventType() CDEventType {
	return ChangeUpdatedEventV1
}

func (sc ChangeUpdatedSubject) GetSubjectType() SubjectType {
	return ChangeSubjectType
}

type ChangeUpdatedEvent struct {
	Context Context              `json:"context"`
	Subject ChangeUpdatedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ChangeUpdatedEvent) GetType() CDEventType {
	return ChangeUpdatedEventV1
}

func (e ChangeUpdatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeUpdatedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeUpdatedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeUpdatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeUpdatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeUpdatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeUpdatedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ChangeUpdatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeUpdatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeUpdatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeUpdatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeUpdatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e ChangeUpdatedEvent) GetSchema() string {
	return changeUpdatedSchemaFile
}

func NewChangeUpdatedEvent() (*ChangeUpdatedEvent, error) {
	e := &ChangeUpdatedEvent{
		Context: Context{
			Type:    ChangeUpdatedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeUpdatedSubject{
			SubjectBase: SubjectBase{
				Type: ChangeSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
