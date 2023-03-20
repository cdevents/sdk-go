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
	_ "embed"
	"fmt"
	"time"
)

//go:embed spec/schemas/taskrunstarted.json
var taskrunstartedschema string

var (
	// TaskRunStarted event v0.1.0
	TaskRunStartedEventV1 CDEventType = CDEventType{
		Subject:   "taskrun",
		Predicate: "started",
		Version:   "0.1.0",
	}
)

type TaskRunStartedSubjectContent struct {

	// The name of the task executed in the task run
	TaskName string `json:"taskName,omitempty"`

	// A URL to the pipeline run
	Url string `json:"url,omitempty"`

	// The PipelineRun associated to the task run
	PipelineRun Reference `json:"pipelineRun,omitempty"`
}

type TaskRunStartedSubject struct {
	SubjectBase
	Content TaskRunStartedSubjectContent `json:"content"`
}

func (sc TaskRunStartedSubject) GetEventType() CDEventType {
	return TaskRunStartedEventV1
}

func (sc TaskRunStartedSubject) GetSubjectType() SubjectType {
	return TaskRunSubjectType
}

type TaskRunStartedEvent struct {
	Context Context               `json:"context"`
	Subject TaskRunStartedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TaskRunStartedEvent) GetType() CDEventType {
	return TaskRunStartedEventV1
}

func (e TaskRunStartedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TaskRunStartedEvent) GetId() string {
	return e.Context.Id
}

func (e TaskRunStartedEvent) GetSource() string {
	return e.Context.Source
}

func (e TaskRunStartedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TaskRunStartedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TaskRunStartedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TaskRunStartedEvent) GetSubject() Subject {
	return e.Subject
}

func (e TaskRunStartedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TaskRunStartedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e TaskRunStartedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TaskRunStartedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation
// TODO(afrittoli) Add stricter validation where relevant

func (e *TaskRunStartedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TaskRunStartedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TaskRunStartedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TaskRunStartedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TaskRunStartedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TaskRunStartedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

// Subject field setters
func (e *TaskRunStartedEvent) SetSubjectTaskName(pipelineName string) {
	e.Subject.Content.TaskName = pipelineName
}

func (e *TaskRunStartedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *TaskRunStartedEvent) SetSubjectPipelineRun(pipelineRun Reference) {
	e.Subject.Content.PipelineRun = pipelineRun
}

func (e TaskRunStartedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), taskrunstartedschema
}

func NewTaskRunStartedEvent() (*TaskRunStartedEvent, error) {
	e := &TaskRunStartedEvent{
		Context: Context{
			Type:    TaskRunStartedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: TaskRunStartedSubject{
			SubjectBase: SubjectBase{
				Type: TaskRunSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
