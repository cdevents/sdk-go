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
	// TestCaseStarted event
	TestCaseStartedEventV1    CDEventType = "dev.cdevents.testcase.started.v1"
	testCaseStartedSchemaFile string      = "testcasestarted"
)

type TestCaseStartedSubjectContent struct{}

type TestCaseStartedSubject struct {
	SubjectBase
	Content TestCaseStartedSubjectContent `json:"content"`
}

func (sc TestCaseStartedSubject) GetEventType() CDEventType {
	return TestCaseStartedEventV1
}

func (sc TestCaseStartedSubject) GetSubjectType() SubjectType {
	return TestCaseSubjectType
}

type TestCaseStartedEvent struct {
	Context Context                `json:"context"`
	Subject TestCaseStartedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e TestCaseStartedEvent) GetType() CDEventType {
	return TestCaseStartedEventV1
}

func (e TestCaseStartedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestCaseStartedEvent) GetId() string {
	return e.Context.Id
}

func (e TestCaseStartedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestCaseStartedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestCaseStartedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestCaseStartedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestCaseStartedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *TestCaseStartedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestCaseStartedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestCaseStartedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestCaseStartedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestCaseStartedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e TestCaseStartedEvent) GetSchema() string {
	return testCaseStartedSchemaFile
}

func NewTestCaseStartedEvent() (*TestCaseStartedEvent, error) {
	e := &TestCaseStartedEvent{
		Context: Context{
			Type:    TestCaseStartedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TestCaseStartedSubject{
			SubjectBase: SubjectBase{
				Type: TestCaseSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
