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
	// ChangeMerged event
	ChangeMergedEventV1 CDEventType = "dev.cdevents.change.merged.v1"
)

type ChangeMergedSubjectContent struct{}

type ChangeMergedSubject struct {
	SubjectBase
	Content ChangeMergedSubjectContent `json:"content"`
}

func (sc ChangeMergedSubject) GetEventType() CDEventType {
	return ChangeMergedEventV1
}

func (sc ChangeMergedSubject) GetSubjectType() SubjectType {
	return ChangeSubjectType
}

type ChangeMergedEvent struct {
	Context Context             `json:"context"`
	Subject ChangeMergedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ChangeMergedEvent) GetType() CDEventType {
	return ChangeMergedEventV1
}

func (e ChangeMergedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeMergedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeMergedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeMergedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeMergedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeMergedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeMergedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ChangeMergedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeMergedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeMergedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeMergedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeMergedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newChangeMergedEvent() CDEvent {
	return &ChangeMergedEvent{
		Context: Context{
			Type:    ChangeMergedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeMergedSubject{
			SubjectBase: SubjectBase{
				Type: ChangeSubjectType,
			},
		},
	}
}
