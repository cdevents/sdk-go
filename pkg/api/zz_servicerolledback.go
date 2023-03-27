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
	_ "embed"
	"fmt"
	"time"
)

//go:embed spec/schemas/servicerolledback.json
var servicerolledbackschema string

var (
	// ServiceRolledback event v0.1.1
	ServiceRolledbackEventV1 CDEventType = CDEventType{
		Subject:   "service",
		Predicate: "rolledback",
		Version:   "0.1.1",
	}
)

type ServiceRolledbackSubjectContent struct {
	ArtifactId string `json:"artifactId" validate:"purl"`

	Environment Reference `json:"environment"`
}

type ServiceRolledbackSubject struct {
	SubjectBase
	Content ServiceRolledbackSubjectContent `json:"content"`
}

func (sc ServiceRolledbackSubject) GetEventType() CDEventType {
	return ServiceRolledbackEventV1
}

func (sc ServiceRolledbackSubject) GetSubjectType() SubjectType {
	return "service"
}

type ServiceRolledbackEvent struct {
	Context Context                  `json:"context"`
	Subject ServiceRolledbackSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ServiceRolledbackEvent) GetType() CDEventType {
	return ServiceRolledbackEventV1
}

func (e ServiceRolledbackEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceRolledbackEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceRolledbackEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceRolledbackEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceRolledbackEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceRolledbackEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceRolledbackEvent) GetSubject() Subject {
	return e.Subject
}

func (e ServiceRolledbackEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ServiceRolledbackEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ServiceRolledbackEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ServiceRolledbackEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ServiceRolledbackEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceRolledbackEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceRolledbackEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceRolledbackEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceRolledbackEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ServiceRolledbackEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ServiceRolledbackEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), servicerolledbackschema
}

// Set subject custom fields

func (e *ServiceRolledbackEvent) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func (e *ServiceRolledbackEvent) SetSubjectEnvironment(environment Reference) {
	e.Subject.Content.Environment = environment
}

// New creates a new ServiceRolledbackEvent
func NewServiceRolledbackEvent() (*ServiceRolledbackEvent, error) {
	e := &ServiceRolledbackEvent{
		Context: Context{
			Type:    ServiceRolledbackEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceRolledbackSubject{
			SubjectBase: SubjectBase{
				Type: "service",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}