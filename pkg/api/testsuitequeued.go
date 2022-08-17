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
	// TestSuiteQueued event
	TestSuiteQueuedEventV1 CDEventType = "dev.cdevents.testsuite.queued.v1"
)

type TestSuiteQueuedSubjectContent struct{}

type TestSuiteQueuedSubject struct {
	SubjectBase
	Content TestSuiteQueuedSubjectContent `json:"content"`
}

func (sc TestSuiteQueuedSubject) GetEventType() CDEventType {
	return TestSuiteQueuedEventV1
}

func (sc TestSuiteQueuedSubject) GetSubjectType() SubjectType {
	return TestSuiteSubjectType
}

type TestSuiteQueuedEvent struct {
	Context Context                `json:"context"`
	Subject TestSuiteQueuedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e TestSuiteQueuedEvent) GetType() CDEventType {
	return TestSuiteQueuedEventV1
}

func (e TestSuiteQueuedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestSuiteQueuedEvent) GetId() string {
	return e.Context.Id
}

func (e TestSuiteQueuedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestSuiteQueuedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestSuiteQueuedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestSuiteQueuedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestSuiteQueuedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *TestSuiteQueuedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestSuiteQueuedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestSuiteQueuedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestSuiteQueuedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestSuiteQueuedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newTestSuiteQueuedEvent() CDEvent {
	return &TestSuiteQueuedEvent{
		Context: Context{
			Type:    TestSuiteQueuedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TestSuiteQueuedSubject{},
	}
}
