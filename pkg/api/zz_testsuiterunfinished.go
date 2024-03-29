// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

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
	"fmt"
	"time"
)

var testsuiterunfinishedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/test-suite-finished-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1},"type":{"type":"string","enum":["dev.cdevents.testsuiterun.finished.0.1.0"],"default":"dev.cdevents.testsuiterun.finished.0.1.0"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string"},"type":{"type":"string","minLength":1,"enum":["testSuiteRun"],"default":"testSuiteRun"},"content":{"properties":{"environment":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]},"testSuite":{"type":"object","additionalProperties":false,"required":["id"],"properties":{"id":{"type":"string","minLength":1},"version":{"type":"string"},"name":{"type":"string"},"uri":{"type":"string","format":"uri"}}},"outcome":{"type":"string","enum":["pass","fail","cancel","error"]},"severity":{"type":"string","enum":["low","medium","high","critical"]},"reason":{"type":"string"}},"additionalProperties":false,"type":"object","required":["outcome","environment"]}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// TestSuiteRunFinished event v0.1.0
	TestSuiteRunFinishedEventV1 CDEventType = CDEventType{
		Subject:   "testsuiterun",
		Predicate: "finished",
		Version:   "0.1.0",
	}
)

type TestSuiteRunFinishedSubjectContent struct {
	Environment *Reference `json:"environment"`

	Outcome string `json:"outcome"`

	Reason string `json:"reason,omitempty"`

	Severity string `json:"severity,omitempty"`

	TestSuite *TestSuiteRunFinishedSubjectContentTestSuite `json:"testSuite,omitempty"`
}

type TestSuiteRunFinishedSubject struct {
	SubjectBase
	Content TestSuiteRunFinishedSubjectContent `json:"content"`
}

func (sc TestSuiteRunFinishedSubject) GetSubjectType() SubjectType {
	return "testSuiteRun"
}

type TestSuiteRunFinishedEvent struct {
	Context Context                     `json:"context"`
	Subject TestSuiteRunFinishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TestSuiteRunFinishedEvent) GetType() CDEventType {
	return TestSuiteRunFinishedEventV1
}

func (e TestSuiteRunFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestSuiteRunFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e TestSuiteRunFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestSuiteRunFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestSuiteRunFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestSuiteRunFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestSuiteRunFinishedEvent) GetSubject() Subject {
	return e.Subject
}

func (e TestSuiteRunFinishedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TestSuiteRunFinishedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e TestSuiteRunFinishedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TestSuiteRunFinishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *TestSuiteRunFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestSuiteRunFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestSuiteRunFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestSuiteRunFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestSuiteRunFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TestSuiteRunFinishedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e TestSuiteRunFinishedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), testsuiterunfinishedschema
}

// Set subject custom fields

func (e *TestSuiteRunFinishedEvent) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

func (e *TestSuiteRunFinishedEvent) SetSubjectOutcome(outcome string) {
	e.Subject.Content.Outcome = outcome
}

func (e *TestSuiteRunFinishedEvent) SetSubjectReason(reason string) {
	e.Subject.Content.Reason = reason
}

func (e *TestSuiteRunFinishedEvent) SetSubjectSeverity(severity string) {
	e.Subject.Content.Severity = severity
}

func (e *TestSuiteRunFinishedEvent) SetSubjectTestSuite(testSuite *TestSuiteRunFinishedSubjectContentTestSuite) {
	e.Subject.Content.TestSuite = testSuite
}

// New creates a new TestSuiteRunFinishedEvent
func NewTestSuiteRunFinishedEvent() (*TestSuiteRunFinishedEvent, error) {
	e := &TestSuiteRunFinishedEvent{
		Context: Context{
			Type:    TestSuiteRunFinishedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: TestSuiteRunFinishedSubject{
			SubjectBase: SubjectBase{
				Type: "testSuiteRun",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// TestSuiteRunFinishedSubjectContentTestSuite holds the content of a TestSuite field in the content
type TestSuiteRunFinishedSubjectContentTestSuite struct {
	Id string `json:"id"`

	Name string `json:"name,omitempty"`

	Uri string `json:"uri,omitempty"`

	Version string `json:"version,omitempty"`
}
