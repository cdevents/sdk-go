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
	PipelineRunFinishedEventV1 CDEventType = "dev.cdevents.pipelinerun.finished.v1"

	// PipelineRun successful
	PipelineRunOutcomeSuccessful PipelineRunOutcome = "success"

	// PipelineRun failed
	PipelineRunOutcomeFailed PipelineRunOutcome = "failure"

	// PipelineRun errored
	PipelineRunOutcomeErrored PipelineRunOutcome = "error"
)

type PipelineRunFinishedSubjectContent struct {

	// The name of the pipeline executed in the pipeline run
	PipelineName string `json:"pipelineName,omitempty"`

	// A URL to the pipeline run
	URL string `json:"url,omitempty"`

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

// Subject field setters
func (e *PipelineRunFinishedEvent) SetSubjectPipelineName(pipelineName string) {
	e.Subject.Content.PipelineName = pipelineName
}

func (e *PipelineRunFinishedEvent) SetSubjectURL(url string) {
	e.Subject.Content.URL = url
}

func (e *PipelineRunFinishedEvent) SetSubjectOutcome(outcome PipelineRunOutcome) {
	e.Subject.Content.Outcome = outcome
}

func (e *PipelineRunFinishedEvent) SetSubjectErrors(errors string) {
	e.Subject.Content.Errors = errors
}

func newPipelineRunFinishedEvent() CDEvent {
	return &PipelineRunFinishedEvent{
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
}
