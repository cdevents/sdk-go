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

type TaskRunOutcome string

func (t TaskRunOutcome) String() string {
	return string(t)
}

const (
	// TaskRunFinished event
	TaskRunFinishedEventV1 CDEventType = "dev.cdevents.taskrun.finished.v1"

	TaskRunOutcomeSuccessful TaskRunOutcome = "success"

	// PipelineRun failed
	TaskRunOutcomeFailed TaskRunOutcome = "failure"

	// PipelineRun errored
	TaskRunOutcomeErrored TaskRunOutcome = "error"

	taskRunFinishedSchemaFile string = "taskrunfinished"
)

type TaskRunFinishedSubjectContent struct {

	// The name of the task executed in the task run
	TaskName string `json:"taskName,omitempty"`

	// A URL to the pipeline run
	URL string `json:"url,omitempty"`

	// The PipelineRun associated to the task run
	PipelineRun Reference `json:"pipelineRun,omitempty"`

	// The PipelineRun outcome
	Outcome TaskRunOutcome `json:"outcome,omitempty"`

	// A string with eventual error descriptions
	Errors string `json:"errors,omitempty"`
}

type TaskRunFinishedSubject struct {
	SubjectBase
	Content TaskRunFinishedSubjectContent `json:"content"`
}

func (sc TaskRunFinishedSubject) GetEventType() CDEventType {
	return TaskRunFinishedEventV1
}

func (sc TaskRunFinishedSubject) GetSubjectType() SubjectType {
	return TaskRunSubjectType
}

type TaskRunFinishedEvent struct {
	Context Context                `json:"context"`
	Subject TaskRunFinishedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e TaskRunFinishedEvent) GetType() CDEventType {
	return TaskRunFinishedEventV1
}

func (e TaskRunFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TaskRunFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e TaskRunFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e TaskRunFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TaskRunFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TaskRunFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TaskRunFinishedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation
// TODO(afrittoli) Add stricter validation where relevant

func (e *TaskRunFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TaskRunFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TaskRunFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TaskRunFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TaskRunFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

// Subject field setters
func (e *TaskRunFinishedEvent) SetSubjectTaskName(pipelineName string) {
	e.Subject.Content.TaskName = pipelineName
}

func (e *TaskRunFinishedEvent) SetSubjectURL(url string) {
	e.Subject.Content.URL = url
}

func (e *TaskRunFinishedEvent) SetSubjectPipelineRun(pipelineRun Reference) {
	e.Subject.Content.PipelineRun = pipelineRun
}

func (e *TaskRunFinishedEvent) SetSubjectOutcome(outcome TaskRunOutcome) {
	e.Subject.Content.Outcome = outcome
}

func (e *TaskRunFinishedEvent) SetSubjectErrors(errors string) {
	e.Subject.Content.Errors = errors
}

func (e TaskRunFinishedEvent) GetSchema() string {
	return taskRunFinishedSchemaFile
}

func NewTaskRunFinishedEvent() (*TaskRunFinishedEvent, error) {
	e := &TaskRunFinishedEvent{
		Context: Context{
			Type:    TaskRunFinishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TaskRunFinishedSubject{
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
