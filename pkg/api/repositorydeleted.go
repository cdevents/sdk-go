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
	// RepositoryDeleted event
	RepositoryDeletedEventV1    CDEventType = "dev.cdevents.repository.deleted.v1"
	repositoryDeletedSchemaFile string      = "repositorydeleted"
)

type RepositoryDeletedSubjectContent struct {

	// The name of the repository, like "sdk-go", "spec" or "a-repo"
	Name string `json:"name,omitempty"`

	// The owner of the repository, like "cdevents", "an-org" or "an-user"
	Owner string `json:"owner,omitempty"`

	// The URL to programmatically access repository, for instance via "git clone"
	Url string `json:"url,omitempty"`

	// The URL for a human to browse the repository, for instance a Web UI to the repo
	ViewUrl string `json:"viewUrl,omitempty"`
}

type RepositoryDeletedSubject struct {
	SubjectBase
	Content RepositoryDeletedSubjectContent `json:"content"`
}

func (sc RepositoryDeletedSubject) GetEventType() CDEventType {
	return RepositoryDeletedEventV1
}

func (sc RepositoryDeletedSubject) GetSubjectType() SubjectType {
	return RepositorySubjectType
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
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e RepositoryDeletedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
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
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryDeletedEvent) GetSchema() string {
	return repositoryDeletedSchemaFile
}

// Subject field setters
func (e *RepositoryDeletedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryDeletedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryDeletedEvent) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryDeletedEvent) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

func NewRepositoryDeletedEvent() (*RepositoryDeletedEvent, error) {
	e := &RepositoryDeletedEvent{
		Context: Context{
			Type:    RepositoryDeletedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: RepositoryDeletedSubject{
			SubjectBase: SubjectBase{
				Type: RepositorySubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
