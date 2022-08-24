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
	// RepositoryCreated event
	RepositoryCreatedEventV1    CDEventType = "dev.cdevents.repository.created.v1"
	repositoryCreatedSchemaFile string      = "repositorycreated"
)

type RepositoryCreatedSubjectContent struct{}

type RepositoryCreatedSubject struct {
	SubjectBase
	Content RepositoryCreatedSubjectContent `json:"content"`
}

func (sc RepositoryCreatedSubject) GetEventType() CDEventType {
	return RepositoryCreatedEventV1
}

func (sc RepositoryCreatedSubject) GetSubjectType() SubjectType {
	return RepositorySubjectType
}

type RepositoryCreatedEvent struct {
	Context Context                  `json:"context"`
	Subject RepositoryCreatedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e RepositoryCreatedEvent) GetType() CDEventType {
	return RepositoryCreatedEventV1
}

func (e RepositoryCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e RepositoryCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e RepositoryCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryCreatedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *RepositoryCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e RepositoryCreatedEvent) GetSchema() string {
	return repositoryCreatedSchemaFile
}

func NewRepositoryCreatedEvent() (*RepositoryCreatedEvent, error) {
	e := &RepositoryCreatedEvent{
		Context: Context{
			Type:    RepositoryCreatedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: RepositoryCreatedSubject{},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
