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
	// ChangeAbandoned event
	ChangeAbandonedEventV1    CDEventType = "dev.cdevents.change.abandoned.0.1.0-draft"
	changeAbandonedSchemaFile string      = "changeabandoned"
)

type ChangeAbandonedSubjectContent struct {

	// Repository where the change occurrence happened
	Repository Reference `json:"repository"`
}

type ChangeAbandonedSubject struct {
	SubjectBase
	Content ChangeAbandonedSubjectContent `json:"content"`
}

func (sc ChangeAbandonedSubject) GetEventType() CDEventType {
	return ChangeAbandonedEventV1
}

func (sc ChangeAbandonedSubject) GetSubjectType() SubjectType {
	return ChangeSubjectType
}

type ChangeAbandonedEvent struct {
	Context Context                `json:"context"`
	Subject ChangeAbandonedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ChangeAbandonedEvent) GetType() CDEventType {
	return ChangeAbandonedEventV1
}

func (e ChangeAbandonedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeAbandonedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeAbandonedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeAbandonedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeAbandonedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeAbandonedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeAbandonedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ChangeAbandonedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ChangeAbandonedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ChangeAbandonedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ChangeAbandonedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ChangeAbandonedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeAbandonedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeAbandonedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeAbandonedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeAbandonedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ChangeAbandonedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ChangeAbandonedEvent) GetSchema() string {
	return changeAbandonedSchemaFile
}

// Subject field setters
func (e *ChangeAbandonedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

func NewChangeAbandonedEvent() (*ChangeAbandonedEvent, error) {
	e := &ChangeAbandonedEvent{
		Context: Context{
			Type:    ChangeAbandonedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeAbandonedSubject{
			SubjectBase: SubjectBase{
				Type: ChangeSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
