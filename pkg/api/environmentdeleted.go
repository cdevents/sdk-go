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
	// EnvironmentDeleted event
	EnvironmentDeletedEventV1    CDEventType = "dev.cdevents.environment.deleted.v1"
	environmentDeletedSchemaFile string      = "environmentdeleted"
)

type EnvironmentDeletedSubjectContent struct {

	// The name of the environment, for instance dev, prod, ci-123
	Name string `json:"name,omitempty"`
}

type EnvironmentDeletedSubject struct {
	SubjectBase
	Content EnvironmentDeletedSubjectContent `json:"content"`
}

func (sc EnvironmentDeletedSubject) GetEventType() CDEventType {
	return EnvironmentDeletedEventV1
}

func (sc EnvironmentDeletedSubject) GetSubjectType() SubjectType {
	return EnvironmentSubjectType
}

type EnvironmentDeletedEvent struct {
	Context Context                   `json:"context"`
	Subject EnvironmentDeletedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e EnvironmentDeletedEvent) GetType() CDEventType {
	return EnvironmentDeletedEventV1
}

func (e EnvironmentDeletedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e EnvironmentDeletedEvent) GetId() string {
	return e.Context.Id
}

func (e EnvironmentDeletedEvent) GetSource() string {
	return e.Context.Source
}

func (e EnvironmentDeletedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e EnvironmentDeletedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e EnvironmentDeletedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e EnvironmentDeletedEvent) GetSubject() Subject {
	return e.Subject
}

func (e EnvironmentDeletedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e EnvironmentDeletedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e EnvironmentDeletedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e EnvironmentDeletedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *EnvironmentDeletedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *EnvironmentDeletedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *EnvironmentDeletedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *EnvironmentDeletedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *EnvironmentDeletedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *EnvironmentDeletedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e EnvironmentDeletedEvent) GetSchema() string {
	return environmentDeletedSchemaFile
}

// Subject field setters
func (e *EnvironmentDeletedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func NewEnvironmentDeletedEvent() (*EnvironmentDeletedEvent, error) {
	e := &EnvironmentDeletedEvent{
		Context: Context{
			Type:    EnvironmentDeletedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: EnvironmentDeletedSubject{
			SubjectBase: SubjectBase{
				Type: EnvironmentSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
