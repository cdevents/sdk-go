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
	// ServiceDeployed event
	ServiceDeployedEventV1    CDEventType = "dev.cdevents.service.deployed.v1"
	serviceDeployedSchemaFile string      = "servicedeployed"
)

type ServiceDeployedSubjectContent struct {

	// The Environment where the service is deployed
	Environment Reference `json:"environment,omitempty"`

	// The Id of the artifact deployed
	ArtifactId string `json:"artifactId,omitempty"`
}

type ServiceDeployedSubject struct {
	SubjectBase
	Content ServiceDeployedSubjectContent `json:"content"`
}

func (sc ServiceDeployedSubject) GetEventType() CDEventType {
	return ServiceDeployedEventV1
}

func (sc ServiceDeployedSubject) GetSubjectType() SubjectType {
	return ServiceSubjectType
}

type ServiceDeployedEvent struct {
	Context Context                `json:"context"`
	Subject ServiceDeployedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ServiceDeployedEvent) GetType() CDEventType {
	return ServiceDeployedEventV1
}

func (e ServiceDeployedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceDeployedEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceDeployedEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceDeployedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceDeployedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceDeployedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceDeployedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ServiceDeployedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ServiceDeployedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e ServiceDeployedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ServiceDeployedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ServiceDeployedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceDeployedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceDeployedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceDeployedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceDeployedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ServiceDeployedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ServiceDeployedEvent) GetSchema() string {
	return serviceDeployedSchemaFile
}

// Subject field setters
func (e *ServiceDeployedEvent) SetSubjectEnvironment(environment Reference) {
	e.Subject.Content.Environment = environment
}

func (e *ServiceDeployedEvent) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func NewServiceDeployedEvent() (*ServiceDeployedEvent, error) {
	e := &ServiceDeployedEvent{
		Context: Context{
			Type:    ServiceDeployedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceDeployedSubject{
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
