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

//go:embed spec/schemas/changereviewed.json
var changereviewedschema string

var (
	// ChangeReviewed event v0.1.0
	ChangeReviewedEventV1 CDEventType = CDEventType{
		Subject:   "change",
		Predicate: "reviewed",
		Version:   "0.1.0",
	}
)

type ChangeReviewedSubjectContent struct {

	// Repository where the change occurrence happened
	Repository Reference `json:"repository"`
}

type ChangeReviewedSubject struct {
	SubjectBase
	Content ChangeReviewedSubjectContent `json:"content"`
}

func (sc ChangeReviewedSubject) GetEventType() CDEventType {
	return ChangeReviewedEventV1
}

func (sc ChangeReviewedSubject) GetSubjectType() SubjectType {
	return ChangeSubjectType
}

type ChangeReviewedEvent struct {
	Context Context               `json:"context"`
	Subject ChangeReviewedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ChangeReviewedEvent) GetType() CDEventType {
	return ChangeReviewedEventV1
}

func (e ChangeReviewedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeReviewedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeReviewedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeReviewedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeReviewedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeReviewedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeReviewedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ChangeReviewedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ChangeReviewedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ChangeReviewedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ChangeReviewedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ChangeReviewedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeReviewedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeReviewedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeReviewedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeReviewedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ChangeReviewedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ChangeReviewedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), changereviewedschema
}

// Subject field setters
func (e *ChangeReviewedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

func NewChangeReviewedEvent() (*ChangeReviewedEvent, error) {
	e := &ChangeReviewedEvent{
		Context: Context{
			Type:    ChangeReviewedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeReviewedSubject{
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
