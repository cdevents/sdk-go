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
	// TestCaseFinished event
	TestCaseFinishedEventV1    CDEventType = "dev.cdevents.testcase.finished.v1"
	testCaseFinishedSchemaFile string      = "testcasefinished"
)

type TestCaseFinishedSubjectContent struct{}

type TestCaseFinishedSubject struct {
	SubjectBase
	Content TestCaseFinishedSubjectContent `json:"content"`
}

func (sc TestCaseFinishedSubject) GetEventType() CDEventType {
	return TestCaseFinishedEventV1
}

func (sc TestCaseFinishedSubject) GetSubjectType() SubjectType {
	return TestCaseSubjectType
}

type TestCaseFinishedEvent struct {
	Context Context                 `json:"context"`
	Subject TestCaseFinishedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e TestCaseFinishedEvent) GetType() CDEventType {
	return TestCaseFinishedEventV1
}

func (e TestCaseFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestCaseFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e TestCaseFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestCaseFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestCaseFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestCaseFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestCaseFinishedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *TestCaseFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestCaseFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestCaseFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestCaseFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestCaseFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e TestCaseFinishedEvent) GetSchema() string {
	return testCaseFinishedSchemaFile
}

func NewTestCaseFinishedEvent() (*TestCaseFinishedEvent, error) {
	e := &TestCaseFinishedEvent{
		Context: Context{
			Type:    TestCaseFinishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TestCaseFinishedSubject{},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
