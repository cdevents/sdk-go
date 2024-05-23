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

import "time"

var (
	// IncidentReported event type v0.1.0
	IncidentReportedEventTypeV0_1_0 CDEventType = CDEventType{
		Subject:   "incident",
		Predicate: "reported",
		Version:   "0.1.0",
	}
)

type IncidentReportedSubjectContentV0_1_0 struct {
	ArtifactId string `json:"artifactId,omitempty" validate:"purl"`

	Description string `json:"description,omitempty"`

	Environment *Reference `json:"environment"`

	Service *Reference `json:"service,omitempty"`

	TicketURI string `json:"ticketURI"`
}

type IncidentReportedSubjectV0_1_0 struct {
	SubjectBase
	Content IncidentReportedSubjectContentV0_1_0 `json:"content"`
}

func (sc IncidentReportedSubjectV0_1_0) GetSubjectType() SubjectType {
	return "incident"
}

type IncidentReportedEventV0_1_0 struct {
	Context Context                       `json:"context"`
	Subject IncidentReportedSubjectV0_1_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e IncidentReportedEventV0_1_0) GetType() CDEventType {
	return IncidentReportedEventTypeV0_1_0
}

func (e IncidentReportedEventV0_1_0) GetVersion() string {
	return CDEventsSpecVersion
}

func (e IncidentReportedEventV0_1_0) GetId() string {
	return e.Context.Id
}

func (e IncidentReportedEventV0_1_0) GetSource() string {
	return e.Context.Source
}

func (e IncidentReportedEventV0_1_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e IncidentReportedEventV0_1_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e IncidentReportedEventV0_1_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e IncidentReportedEventV0_1_0) GetSubject() Subject {
	return e.Subject
}

func (e IncidentReportedEventV0_1_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e IncidentReportedEventV0_1_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e IncidentReportedEventV0_1_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e IncidentReportedEventV0_1_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *IncidentReportedEventV0_1_0) SetId(id string) {
	e.Context.Id = id
}

func (e *IncidentReportedEventV0_1_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *IncidentReportedEventV0_1_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *IncidentReportedEventV0_1_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *IncidentReportedEventV0_1_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *IncidentReportedEventV0_1_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e IncidentReportedEventV0_1_0) GetSchema() (string, string) {
	eType := e.GetType()
	id, schema, _ := GetSchemaBySpecSubjectPredicate(CDEventsSpecVersion, eType.Subject, eType.Predicate)
	return id, schema
}

// Set subject custom fields

func (e *IncidentReportedEventV0_1_0) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func (e *IncidentReportedEventV0_1_0) SetSubjectDescription(description string) {
	e.Subject.Content.Description = description
}

func (e *IncidentReportedEventV0_1_0) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

func (e *IncidentReportedEventV0_1_0) SetSubjectService(service *Reference) {
	e.Subject.Content.Service = service
}

func (e *IncidentReportedEventV0_1_0) SetSubjectTicketURI(ticketURI string) {
	e.Subject.Content.TicketURI = ticketURI
}

// New creates a new IncidentReportedEventV0_1_0
func NewIncidentReportedEventV0_1_0(specVersion string) (*IncidentReportedEventV0_1_0, error) {
	e := &IncidentReportedEventV0_1_0{
		Context: Context{
			Type:    IncidentReportedEventTypeV0_1_0,
			Version: specVersion,
		},
		Subject: IncidentReportedSubjectV0_1_0{
			SubjectBase: SubjectBase{
				Type: "incident",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
