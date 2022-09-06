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
	// PipelineRunStarted event
	PipelineRunStartedEventV1    CDEventType = "dev.cdevents.pipelinerun.started.v1"
	pipelineRunStartedSchemaFile string      = "pipelinerunstarted"
)

type PipelineRunStartedSubjectContent struct {

	// The name of the pipeline executed in the pipeline run
	PipelineName string `json:"pipelineName"`

	// A URL to the pipeline run
	Url string `json:"url"`
}

type PipelineRunStartedSubject struct {
	SubjectBase
	Content PipelineRunStartedSubjectContent `json:"content"`
}

func (sc PipelineRunStartedSubject) GetEventType() CDEventType {
	return PipelineRunStartedEventV1
}

func (sc PipelineRunStartedSubject) GetSubjectType() SubjectType {
	return PipelineRunSubjectType
}

type PipelineRunStartedEvent struct {
	Context Context                   `json:"context"`
	Subject PipelineRunStartedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e PipelineRunStartedEvent) GetType() CDEventType {
	return PipelineRunStartedEventV1
}

func (e PipelineRunStartedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e PipelineRunStartedEvent) GetId() string {
	return e.Context.Id
}

func (e PipelineRunStartedEvent) GetSource() string {
	return e.Context.Source
}

func (e PipelineRunStartedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e PipelineRunStartedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e PipelineRunStartedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e PipelineRunStartedEvent) GetSubject() Subject {
	return e.Subject
}

func (e PipelineRunStartedEvent) GetCustomData() []byte {
	return e.CustomData
}

func (e PipelineRunStartedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e PipelineRunStartedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation
// TODO(afrittoli) Add stricter validation where relevant

func (e *PipelineRunStartedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *PipelineRunStartedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *PipelineRunStartedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *PipelineRunStartedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *PipelineRunStartedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *PipelineRunStartedEvent) SetCustomData(contentType string, data interface{}) error {
	dataBytes, err := customDataBytes(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = dataBytes
	return nil
}

// Subject field setters
func (e *PipelineRunStartedEvent) SetSubjectPipelineName(pipelineName string) {
	e.Subject.Content.PipelineName = pipelineName
}

func (e *PipelineRunStartedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e PipelineRunStartedEvent) GetSchema() string {
	return pipelineRunStartedSchemaFile
}

func NewPipelineRunStartedEvent() (*PipelineRunStartedEvent, error) {
	e := &PipelineRunStartedEvent{
		Context: Context{
			Type:    PipelineRunStartedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: PipelineRunStartedSubject{
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
