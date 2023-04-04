// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

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
	"fmt"
	"time"
)

var changereviewedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.2.0/schema/change-reviewed-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.change.reviewed.0.1.2"],"default":"dev.cdevents.change.reviewed.0.1.2"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["change"],"default":"change"},"content":{"properties":{"repository":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]}},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// ChangeReviewed event v0.1.2
	ChangeReviewedEventV1 CDEventType = CDEventType{
		Subject:   "change",
		Predicate: "reviewed",
		Version:   "0.1.2",
	}
)

type ChangeReviewedSubjectContent struct {
	Repository Reference `json:"repository"`
}

type ChangeReviewedSubject struct {
	SubjectBase
	Content ChangeReviewedSubjectContent `json:"content"`
}

func (sc ChangeReviewedSubject) GetSubjectType() SubjectType {
	return "change"
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
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ChangeReviewedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ChangeReviewedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
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
	err := CheckCustomData(contentType, data)
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

// Set subject custom fields

func (e *ChangeReviewedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

// New creates a new ChangeReviewedEvent
func NewChangeReviewedEvent() (*ChangeReviewedEvent, error) {
	e := &ChangeReviewedEvent{
		Context: Context{
			Type:    ChangeReviewedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeReviewedSubject{
			SubjectBase: SubjectBase{
				Type: "change",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
