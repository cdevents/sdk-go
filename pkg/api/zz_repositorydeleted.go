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

//go:embed spec/schemas/repositorydeleted.json
var repositorydeletedschema string

var (
	// RepositoryDeleted event v0.1.1
	RepositoryDeletedEventV1 CDEventType = CDEventType{
		Subject:   "repository",
		Predicate: "deleted",
		Version:   "0.1.1",
	}
)

type RepositoryDeletedSubjectContent struct {
	Name string `json:"name"`

	Owner string `json:"owner"`

	Url string `json:"url"`

	ViewUrl string `json:"viewUrl"`
}

type RepositoryDeletedSubject struct {
	SubjectBase
	Content RepositoryDeletedSubjectContent `json:"content"`
}

func (sc RepositoryDeletedSubject) GetEventType() CDEventType {
	return RepositoryDeletedEventV1
}

func (sc RepositoryDeletedSubject) GetSubjectType() SubjectType {
	return "repository"
}

type RepositoryDeletedEvent struct {
	Context Context                  `json:"context"`
	Subject RepositoryDeletedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e RepositoryDeletedEvent) GetType() CDEventType {
	return RepositoryDeletedEventV1
}

func (e RepositoryDeletedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryDeletedEvent) GetId() string {
	return e.Context.Id
}

func (e RepositoryDeletedEvent) GetSource() string {
	return e.Context.Source
}

func (e RepositoryDeletedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryDeletedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryDeletedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryDeletedEvent) GetSubject() Subject {
	return e.Subject
}

func (e RepositoryDeletedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e RepositoryDeletedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *RepositoryDeletedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryDeletedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryDeletedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryDeletedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryDeletedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *RepositoryDeletedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryDeletedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), repositorydeletedschema
}

// Set subject custom fields

func (e *RepositoryDeletedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryDeletedEvent) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryDeletedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryDeletedEvent) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

// New creates a new RepositoryDeletedEvent
func NewRepositoryDeletedEvent() (*RepositoryDeletedEvent, error) {
	e := &RepositoryDeletedEvent{
		Context: Context{
			Type:    RepositoryDeletedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: RepositoryDeletedSubject{
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