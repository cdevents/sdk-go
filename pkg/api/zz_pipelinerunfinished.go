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

var pipelinerunfinishedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/pipeline-run-finished-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.pipelinerun.finished.0.1.1"],"default":"dev.cdevents.pipelinerun.finished.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["pipelineRun"],"default":"pipelineRun"},"content":{"properties":{"pipelineName":{"type":"string"},"url":{"type":"string"},"outcome":{"type":"string"},"errors":{"type":"string"}},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// PipelineRunFinished event v0.1.1
	PipelineRunFinishedEventV1 CDEventType = CDEventType{
		Subject:   "pipelinerun",
		Predicate: "finished",
		Version:   "0.1.1",
	}
)

type PipelineRunFinishedSubjectContent struct {
	Errors string `json:"errors,omitempty"`

	Outcome string `json:"outcome,omitempty"`

	PipelineName string `json:"pipelineName,omitempty"`

	Url string `json:"url,omitempty"`
}

type PipelineRunFinishedSubject struct {
	SubjectBase
	Content PipelineRunFinishedSubjectContent `json:"content"`
}

func (sc PipelineRunFinishedSubject) GetSubjectType() SubjectType {
	return "pipelineRun"
}

type PipelineRunFinishedEvent struct {
	Context Context                    `json:"context"`
	Subject PipelineRunFinishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e PipelineRunFinishedEvent) GetType() CDEventType {
	return PipelineRunFinishedEventV1
}

func (e PipelineRunFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e PipelineRunFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e PipelineRunFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e PipelineRunFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e PipelineRunFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e PipelineRunFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e PipelineRunFinishedEvent) GetSubject() Subject {
	return e.Subject
}

func (e PipelineRunFinishedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e PipelineRunFinishedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e PipelineRunFinishedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e PipelineRunFinishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *PipelineRunFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *PipelineRunFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *PipelineRunFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *PipelineRunFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *PipelineRunFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *PipelineRunFinishedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e PipelineRunFinishedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), pipelinerunfinishedschema
}

// Set subject custom fields

func (e *PipelineRunFinishedEvent) SetSubjectErrors(errors string) {
	e.Subject.Content.Errors = errors
}

func (e *PipelineRunFinishedEvent) SetSubjectOutcome(outcome string) {
	e.Subject.Content.Outcome = outcome
}

func (e *PipelineRunFinishedEvent) SetSubjectPipelineName(pipelineName string) {
	e.Subject.Content.PipelineName = pipelineName
}

func (e *PipelineRunFinishedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

// New creates a new PipelineRunFinishedEvent
func NewPipelineRunFinishedEvent() (*PipelineRunFinishedEvent, error) {
	e := &PipelineRunFinishedEvent{
		Context: Context{
			Type:    PipelineRunFinishedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: PipelineRunFinishedSubject{
			SubjectBase: SubjectBase{
				Type: "pipelineRun",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
