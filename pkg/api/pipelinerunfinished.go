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

type PipelineRunOutcome string

func (t PipelineRunOutcome) String() string {
	return string(t)
}

const (
	// PipelineRunFinished event
	PipelineRunFinishedEventV1 CDEventType = "dev.cdevents.pipelinerun.finished.0.1.0"

	// PipelineRun successful
	PipelineRunOutcomeSuccessful PipelineRunOutcome = "success"

	// PipelineRun failed
	PipelineRunOutcomeFailed PipelineRunOutcome = "failure"

	// PipelineRun errored
	PipelineRunOutcomeErrored PipelineRunOutcome = "error"

	pipelineRunFinishedSchemaFile string = "pipelinerunfinished"
)

type PipelineRunFinishedSubjectContent struct {

	// The name of the pipeline executed in the pipeline run
	PipelineName string `json:"pipelineName,omitempty"`

	// A URL to the pipeline run
	Url string `json:"url,omitempty"`

	// The PipelineRun outcome
	Outcome PipelineRunOutcome `json:"outcome,omitempty"`

	// A string with eventual error descriptions
	Errors string `json:"errors,omitempty"`
}

type PipelineRunFinishedSubject struct {
	SubjectBase
	Content PipelineRunFinishedSubjectContent `json:"content"`
}

func (sc PipelineRunFinishedSubject) GetEventType() CDEventType {
	return PipelineRunFinishedEventV1
}

func (sc PipelineRunFinishedSubject) GetSubjectType() SubjectType {
	return PipelineRunSubjectType
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
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e PipelineRunFinishedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e PipelineRunFinishedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e PipelineRunFinishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation
// TODO(afrittoli) Add stricter validation where relevant

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
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

// Subject field setters
func (e *PipelineRunFinishedEvent) SetSubjectPipelineName(pipelineName string) {
	e.Subject.Content.PipelineName = pipelineName
}

func (e *PipelineRunFinishedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *PipelineRunFinishedEvent) SetSubjectOutcome(outcome PipelineRunOutcome) {
	e.Subject.Content.Outcome = outcome
}

func (e *PipelineRunFinishedEvent) SetSubjectErrors(errors string) {
	e.Subject.Content.Errors = errors
}

func (e PipelineRunFinishedEvent) GetSchema() string {
	return pipelineRunFinishedSchemaFile
}

func NewPipelineRunFinishedEvent() (*PipelineRunFinishedEvent, error) {
	e := &PipelineRunFinishedEvent{
		Context: Context{
			Type:    PipelineRunFinishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: PipelineRunFinishedSubject{
			SubjectBase: SubjectBase{
				Type: PipelineRunSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
