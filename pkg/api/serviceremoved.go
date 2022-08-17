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
	// ServiceRemoved event
	ServiceRemovedEventV1 CDEventType = "dev.cdevents.service.removed.v1"
)

type ServiceRemovedSubjectContent struct{}

type ServiceRemovedSubject struct {
	SubjectBase
	Content ServiceRemovedSubjectContent `json:"content"`
}

func (sc ServiceRemovedSubject) GetEventType() CDEventType {
	return ServiceRemovedEventV1
}

func (sc ServiceRemovedSubject) GetSubjectType() SubjectType {
	return ServiceSubjectType
}

type ServiceRemovedEvent struct {
	Context Context               `json:"context"`
	Subject ServiceRemovedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ServiceRemovedEvent) GetType() CDEventType {
	return ServiceRemovedEventV1
}

func (e ServiceRemovedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceRemovedEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceRemovedEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceRemovedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceRemovedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceRemovedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceRemovedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ServiceRemovedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceRemovedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceRemovedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceRemovedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceRemovedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newServiceRemovedEvent() CDEvent {
	return &ServiceRemovedEvent{
		Context: Context{
			Type:    ServiceRemovedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceRemovedSubject{},
	}
}
