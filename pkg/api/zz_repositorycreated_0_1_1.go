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
	"time"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	// RepositoryCreated event type v0.1.1
	RepositoryCreatedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "repository",
		Predicate: "created",
		Version:   "0.1.1",
	}
)

type RepositoryCreatedSubjectContentV0_1_1 struct {
	Name string `json:"name"`

	Owner string `json:"owner,omitempty"`

	Url string `json:"url"`

	ViewUrl string `json:"viewUrl,omitempty"`
}

type RepositoryCreatedSubjectV0_1_1 struct {
	SubjectBase
	Content RepositoryCreatedSubjectContentV0_1_1 `json:"content"`
}

func (sc RepositoryCreatedSubjectV0_1_1) GetSubjectType() SubjectType {
	return "repository"
}

type RepositoryCreatedEventV0_1_1 struct {
	Context Context                        `json:"context"`
	Subject RepositoryCreatedSubjectV0_1_1 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e RepositoryCreatedEventV0_1_1) GetType() CDEventType {
	return RepositoryCreatedEventTypeV0_1_1
}

func (e RepositoryCreatedEventV0_1_1) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryCreatedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e RepositoryCreatedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e RepositoryCreatedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryCreatedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryCreatedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryCreatedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e RepositoryCreatedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryCreatedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e RepositoryCreatedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryCreatedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *RepositoryCreatedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryCreatedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryCreatedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryCreatedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryCreatedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *RepositoryCreatedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryCreatedEventV0_1_1) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.3.0", eType.Subject, eType.Predicate)
}

// Set subject custom fields

func (e *RepositoryCreatedEventV0_1_1) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryCreatedEventV0_1_1) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryCreatedEventV0_1_1) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryCreatedEventV0_1_1) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

// New creates a new RepositoryCreatedEventV0_1_1
func NewRepositoryCreatedEventV0_1_1(specVersion string) (*RepositoryCreatedEventV0_1_1, error) {
	e := &RepositoryCreatedEventV0_1_1{
		Context: Context{
			Type:    RepositoryCreatedEventTypeV0_1_1,
			Version: specVersion,
		},
		Subject: RepositoryCreatedSubjectV0_1_1{
			SubjectBase: SubjectBase{
				Type: "repository",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
