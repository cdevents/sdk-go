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
	// PipelineRun events
	PipelineRunQueuedEventV1    CDEventType = "dev.cdevents.pipelinerun.queued.v1"
	pipelineRunQueuedSchemaFile string      = "pipelinerunqueued"
)

type PipelineRunQueuedSubjectContent struct {

	// The name of the pipeline executed in the pipeline run
	PipelineName string `json:"pipelineName,omitempty"`

	// A URL to the pipeline run
	Url string `json:"url,omitempty"`
}

type PipelineRunQueuedSubject struct {
	SubjectBase
	Content PipelineRunQueuedSubjectContent `json:"content"`
}

func (sc PipelineRunQueuedSubject) GetEventType() CDEventType {
	return PipelineRunQueuedEventV1
}

func (sc PipelineRunQueuedSubject) GetSubjectType() SubjectType {
	return PipelineRunSubjectType
}

type PipelineRunQueuedEvent struct {
	Context Context                  `json:"context"`
	Subject PipelineRunQueuedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e PipelineRunQueuedEvent) GetType() CDEventType {
	return PipelineRunQueuedEventV1
}

func (e PipelineRunQueuedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e PipelineRunQueuedEvent) GetId() string {
	return e.Context.Id
}

func (e PipelineRunQueuedEvent) GetSource() string {
	return e.Context.Source
}

func (e PipelineRunQueuedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e PipelineRunQueuedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e PipelineRunQueuedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e PipelineRunQueuedEvent) GetSubject() Subject {
	return e.Subject
}

func (e PipelineRunQueuedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e PipelineRunQueuedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e PipelineRunQueuedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e PipelineRunQueuedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation
// TODO(afrittoli) Add stricter validation where relevant

func (e *PipelineRunQueuedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *PipelineRunQueuedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *PipelineRunQueuedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *PipelineRunQueuedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *PipelineRunQueuedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *PipelineRunQueuedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

// Subject field setters
func (e *PipelineRunQueuedEvent) SetSubjectPipelineName(pipelineName string) {
	e.Subject.Content.PipelineName = pipelineName
}

func (e *PipelineRunQueuedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e PipelineRunQueuedEvent) GetSchema() string {
	return pipelineRunQueuedSchemaFile
}

func NewPipelineRunQueuedEvent() (*PipelineRunQueuedEvent, error) {
	e := &PipelineRunQueuedEvent{
		Context: Context{
			Type:    PipelineRunQueuedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: PipelineRunQueuedSubject{
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
