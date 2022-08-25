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
	// TestSuiteFinished event
	TestSuiteFinishedEventV1    CDEventType = "dev.cdevents.testsuite.finished.v1"
	testSuiteFinishedSchemaFile string      = "testsuitefinished"
)

type TestSuiteFinishedSubjectContent struct{}

type TestSuiteFinishedSubject struct {
	SubjectBase
	Content TestSuiteFinishedSubjectContent `json:"content"`
}

func (sc TestSuiteFinishedSubject) GetEventType() CDEventType {
	return TestSuiteFinishedEventV1
}

func (sc TestSuiteFinishedSubject) GetSubjectType() SubjectType {
	return TestSuiteSubjectType
}

type TestSuiteFinishedEvent struct {
	Context Context                  `json:"context"`
	Subject TestSuiteFinishedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e TestSuiteFinishedEvent) GetType() CDEventType {
	return TestSuiteFinishedEventV1
}

func (e TestSuiteFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestSuiteFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e TestSuiteFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestSuiteFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestSuiteFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestSuiteFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestSuiteFinishedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *TestSuiteFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestSuiteFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestSuiteFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestSuiteFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestSuiteFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e TestSuiteFinishedEvent) GetSchema() string {
	return testSuiteFinishedSchemaFile
}

func NewTestSuiteFinishedEvent() (*TestSuiteFinishedEvent, error) {
	e := &TestSuiteFinishedEvent{
		Context: Context{
			Type:    TestSuiteFinishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TestSuiteFinishedSubject{
			SubjectBase: SubjectBase{
				Type: TestSuiteSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
