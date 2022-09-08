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
	// ServicePublished event
	ServicePublishedEventV1    CDEventType = "dev.cdevents.service.published.v1"
	servicePublishedSchemaFile string      = "servicepublished"
)

type ServicePublishedSubjectContent struct {

	// The Environment where the service is deployed
	Environment Reference `json:"environment,omitempty"`
}

type ServicePublishedSubject struct {
	SubjectBase
	Content ServicePublishedSubjectContent `json:"content"`
}

func (sc ServicePublishedSubject) GetEventType() CDEventType {
	return ServicePublishedEventV1
}

func (sc ServicePublishedSubject) GetSubjectType() SubjectType {
	return ServiceSubjectType
}

type ServicePublishedEvent struct {
	Context Context                 `json:"context"`
	Subject ServicePublishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ServicePublishedEvent) GetType() CDEventType {
	return ServicePublishedEventV1
}

func (e ServicePublishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServicePublishedEvent) GetId() string {
	return e.Context.Id
}

func (e ServicePublishedEvent) GetSource() string {
	return e.Context.Source
}

func (e ServicePublishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServicePublishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServicePublishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServicePublishedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ServicePublishedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ServicePublishedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ServicePublishedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ServicePublishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ServicePublishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServicePublishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServicePublishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServicePublishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServicePublishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ServicePublishedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ServicePublishedEvent) GetSchema() string {
	return servicePublishedSchemaFile
}

// Subject field setters
func (e *ServicePublishedEvent) SetSubjectEnvironment(environment Reference) {
	e.Subject.Content.Environment = environment
}

func NewServicePublishedEvent() (*ServicePublishedEvent, error) {
	e := &ServicePublishedEvent{
		Context: Context{
			Type:    ServicePublishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ServicePublishedSubject{
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
