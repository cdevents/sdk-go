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
	// ServiceUpgraded event
	ServiceUpgradedEventV1 CDEventType = "dev.cdevents.service.upgraded.v1"
)

type ServiceUpgradedSubjectContent struct{}

type ServiceUpgradedSubject struct {
	SubjectBase
	Content ServiceUpgradedSubjectContent `json:"content"`
}

func (sc ServiceUpgradedSubject) GetEventType() CDEventType {
	return ServiceUpgradedEventV1
}

func (sc ServiceUpgradedSubject) GetSubjectType() SubjectType {
	return ServiceSubjectType
}

type ServiceUpgradedEvent struct {
	Context Context                `json:"context"`
	Subject ServiceUpgradedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ServiceUpgradedEvent) GetType() CDEventType {
	return ServiceUpgradedEventV1
}

func (e ServiceUpgradedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceUpgradedEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceUpgradedEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceUpgradedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceUpgradedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceUpgradedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceUpgradedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ServiceUpgradedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceUpgradedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceUpgradedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceUpgradedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceUpgradedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newServiceUpgradedEvent() CDEvent {
	return &ServiceUpgradedEvent{
		Context: Context{
			Type:    ServiceUpgradedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceUpgradedSubject{},
	}
}
