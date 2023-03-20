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

//go:embed spec/schemas/changeupdated.json
var changeupdatedschema string

var (
	// ChangeUpdated event v0.1.0
	ChangeUpdatedEventV1 CDEventType = CDEventType{
		Subject:   "change",
		Predicate: "updated",
		Version:   "0.1.0",
	}
)

type ChangeUpdatedSubjectContent struct {

	// Repository where the change occurrence happened
	Repository Reference `json:"repository"`
}

type ChangeUpdatedSubject struct {
	SubjectBase
	Content ChangeUpdatedSubjectContent `json:"content"`
}

func (sc ChangeUpdatedSubject) GetEventType() CDEventType {
	return ChangeUpdatedEventV1
}

func (sc ChangeUpdatedSubject) GetSubjectType() SubjectType {
	return ChangeSubjectType
}

type ChangeUpdatedEvent struct {
	Context Context              `json:"context"`
	Subject ChangeUpdatedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ChangeUpdatedEvent) GetType() CDEventType {
	return ChangeUpdatedEventV1
}

func (e ChangeUpdatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeUpdatedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeUpdatedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeUpdatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeUpdatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeUpdatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeUpdatedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ChangeUpdatedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ChangeUpdatedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ChangeUpdatedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ChangeUpdatedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ChangeUpdatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeUpdatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeUpdatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeUpdatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeUpdatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ChangeUpdatedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ChangeUpdatedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), changeupdatedschema
}

// Subject field setters
func (e *ChangeUpdatedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

func NewChangeUpdatedEvent() (*ChangeUpdatedEvent, error) {
	e := &ChangeUpdatedEvent{
		Context: Context{
			Type:    ChangeUpdatedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeUpdatedSubject{
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
