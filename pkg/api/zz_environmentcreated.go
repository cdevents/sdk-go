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

var environmentcreatedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.2.0/schema/environment-created-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.environment.created.0.1.1"],"default":"dev.cdevents.environment.created.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["environment"],"default":"environment"},"content":{"properties":{"name":{"type":"string"},"url":{"type":"string"}},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// EnvironmentCreated event v0.1.1
	EnvironmentCreatedEventV1 CDEventType = CDEventType{
		Subject:   "environment",
		Predicate: "created",
		Version:   "0.1.1",
	}
)

type EnvironmentCreatedSubjectContent struct {
	Name string `json:"name"`

	Url string `json:"url"`
}

type EnvironmentCreatedSubject struct {
	SubjectBase
	Content EnvironmentCreatedSubjectContent `json:"content"`
}

func (sc EnvironmentCreatedSubject) GetSubjectType() SubjectType {
	return "environment"
}

type EnvironmentCreatedEvent struct {
	Context Context                   `json:"context"`
	Subject EnvironmentCreatedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e EnvironmentCreatedEvent) GetType() CDEventType {
	return EnvironmentCreatedEventV1
}

func (e EnvironmentCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e EnvironmentCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e EnvironmentCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e EnvironmentCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e EnvironmentCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e EnvironmentCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e EnvironmentCreatedEvent) GetSubject() Subject {
	return e.Subject
}

func (e EnvironmentCreatedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e EnvironmentCreatedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e EnvironmentCreatedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e EnvironmentCreatedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *EnvironmentCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *EnvironmentCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *EnvironmentCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *EnvironmentCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *EnvironmentCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *EnvironmentCreatedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e EnvironmentCreatedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), environmentcreatedschema
}

// Set subject custom fields

func (e *EnvironmentCreatedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *EnvironmentCreatedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

// New creates a new EnvironmentCreatedEvent
func NewEnvironmentCreatedEvent() (*EnvironmentCreatedEvent, error) {
	e := &EnvironmentCreatedEvent{
		Context: Context{
			Type:    EnvironmentCreatedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: EnvironmentCreatedSubject{
			SubjectBase: SubjectBase{
				Type: "environment",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
