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

//go:embed spec/schemas/serviceremoved.json
var serviceremovedschema string

var (
	// ServiceRemoved event v0.1.0
	ServiceRemovedEventV1 CDEventType = CDEventType{
		Subject:   "service",
		Predicate: "removed",
		Version:   "0.1.0",
	}
)

type ServiceRemovedSubjectContent struct {

	// The Environment where the service is deployed
	Environment Reference `json:"environment,omitempty"`
}

type ServiceRemovedSubject struct {
	SubjectBase
	Content ServiceRemovedSubjectContent `json:"content"`
}

func (sc ServiceRemovedSubject) GetEventType() CDEventType {
	return ServiceRemovedEventV1
}

func (sc ServiceRemovedSubject) GetSubjectType() SubjectType {
	return ServiceSubjectType
}

type ServiceRemovedEvent struct {
	Context Context               `json:"context"`
	Subject ServiceRemovedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ServiceRemovedEvent) GetType() CDEventType {
	return ServiceRemovedEventV1
}

func (e ServiceRemovedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceRemovedEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceRemovedEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceRemovedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceRemovedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceRemovedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceRemovedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ServiceRemovedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ServiceRemovedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ServiceRemovedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ServiceRemovedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ServiceRemovedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceRemovedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceRemovedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceRemovedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceRemovedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ServiceRemovedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ServiceRemovedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), serviceremovedschema
}

// Subject field setters
func (e *ServiceRemovedEvent) SetSubjectEnvironment(environment Reference) {
	e.Subject.Content.Environment = environment
}

func NewServiceRemovedEvent() (*ServiceRemovedEvent, error) {
	e := &ServiceRemovedEvent{
		Context: Context{
			Type:    ServiceRemovedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceRemovedSubject{
			SubjectBase: SubjectBase{
				Type: ServiceSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
