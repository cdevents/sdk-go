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
	// EnvironmentCreated event
	EnvironmentCreatedEventV1    CDEventType = "dev.cdevents.environment.created.v1"
	environmentCreatedSchemaFile string      = "environmentcreated"
)

type EnvironmentCreatedSubjectContent struct {

	// The name of the environment, for instance dev, prod, ci-123
	Name string `json:"name,omitempty"`

	// A URL to the environment, for instance https://my-cluster.zone.my-cloud-provider
	Url string `json:"url,omitempty"`
}

type EnvironmentCreatedSubject struct {
	SubjectBase
	Content EnvironmentCreatedSubjectContent `json:"content"`
}

func (sc EnvironmentCreatedSubject) GetEventType() CDEventType {
	return EnvironmentCreatedEventV1
}

func (sc EnvironmentCreatedSubject) GetSubjectType() SubjectType {
	return EnvironmentSubjectType
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

func (e EnvironmentCreatedEvent) GetCustomData() []byte {
	return e.CustomData
}

func (e EnvironmentCreatedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
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
	dataBytes, err := customDataBytes(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = dataBytes
	return nil
}

func (e EnvironmentCreatedEvent) GetSchema() string {
	return environmentCreatedSchemaFile
}

// Subject field setters
func (e *EnvironmentCreatedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *EnvironmentCreatedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func NewEnvironmentCreatedEvent() (*EnvironmentCreatedEvent, error) {
	e := &EnvironmentCreatedEvent{
		Context: Context{
			Type:    EnvironmentCreatedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: EnvironmentCreatedSubject{
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
