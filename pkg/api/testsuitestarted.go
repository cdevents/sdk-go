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
	// TestSuiteStarted event
	TestSuiteStartedEventV1    CDEventType = "dev.cdevents.testsuite.started.v1"
	testSuiteStartedSchemaFile string      = "testsuitestarted"
)

type TestSuiteStartedSubjectContent struct{}

type TestSuiteStartedSubject struct {
	SubjectBase
	Content TestSuiteStartedSubjectContent `json:"content"`
}

func (sc TestSuiteStartedSubject) GetEventType() CDEventType {
	return TestSuiteStartedEventV1
}

func (sc TestSuiteStartedSubject) GetSubjectType() SubjectType {
	return TestSuiteSubjectType
}

type TestSuiteStartedEvent struct {
	Context Context                 `json:"context"`
	Subject TestSuiteStartedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TestSuiteStartedEvent) GetType() CDEventType {
	return TestSuiteStartedEventV1
}

func (e TestSuiteStartedEvent) GetVersion() string {
	return string(CDEventsSpecVersion)
}

func (e TestSuiteStartedEvent) GetId() string {
	return e.Context.Id
}

func (e TestSuiteStartedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestSuiteStartedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestSuiteStartedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestSuiteStartedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestSuiteStartedEvent) GetSubject() Subject {
	return e.Subject
}

func (e TestSuiteStartedEvent) GetCustomData() []byte {
	return e.CustomData
}

func (e TestSuiteStartedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e TestSuiteStartedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *TestSuiteStartedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestSuiteStartedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestSuiteStartedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestSuiteStartedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestSuiteStartedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TestSuiteStartedEvent) SetCustomData(contentType string, data interface{}) error {
	dataBytes, err := customDataBytes(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = dataBytes
	return nil
}

func (e TestSuiteStartedEvent) GetSchema() string {
	return testSuiteStartedSchemaFile
}

func NewTestSuiteStartedEvent() (*TestSuiteStartedEvent, error) {
	e := &TestSuiteStartedEvent{
		Context: Context{
			Type:    TestSuiteStartedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TestSuiteStartedSubject{
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
