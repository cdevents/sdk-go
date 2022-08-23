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
	// RepositoryModified event
	RepositoryModifiedEventV1    CDEventType = "dev.cdevents.repository.modified.v1"
	repositoryModifiedSchemaFile string      = "repositorymodified"
)

type RepositoryModifiedSubjectContent struct{}

type RepositoryModifiedSubject struct {
	SubjectBase
	Content RepositoryModifiedSubjectContent `json:"content"`
}

func (sc RepositoryModifiedSubject) GetEventType() CDEventType {
	return RepositoryModifiedEventV1
}

func (sc RepositoryModifiedSubject) GetSubjectType() SubjectType {
	return RepositorySubjectType
}

type RepositoryModifiedEvent struct {
	Context Context                   `json:"context"`
	Subject RepositoryModifiedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e RepositoryModifiedEvent) GetType() CDEventType {
	return RepositoryModifiedEventV1
}

func (e RepositoryModifiedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryModifiedEvent) GetId() string {
	return e.Context.Id
}

func (e RepositoryModifiedEvent) GetSource() string {
	return e.Context.Source
}

func (e RepositoryModifiedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryModifiedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryModifiedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryModifiedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *RepositoryModifiedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryModifiedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryModifiedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryModifiedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryModifiedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e RepositoryModifiedEvent) GetSchema() string {
	return repositoryModifiedSchemaFile
}

func newRepositoryModifiedEvent() CDEvent {
	return &RepositoryModifiedEvent{
		Context: Context{
			Type:    RepositoryModifiedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: RepositoryModifiedSubject{},
	}
}
