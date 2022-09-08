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
	// RepositoryModified event
	RepositoryModifiedEventV1    CDEventType = "dev.cdevents.repository.modified.v1"
	repositoryModifiedSchemaFile string      = "repositorymodified"
)

type RepositoryModifiedSubjectContent struct {

	// The name of the repository, like "sdk-go", "spec" or "a-repo"
	Name string `json:"name,omitempty"`

	// The owner of the repository, like "cdevents", "an-org" or "an-user"
	Owner string `json:"owner,omitempty"`

	// The URL to programmatically access repository, for instance via "git clone"
	Url string `json:"url,omitempty"`

	// The URL for a human to browse the repository, for instance a Web UI to the repo
	ViewUrl string `json:"viewUrl,omitempty"`
}

type RepositoryModifiedSubject struct {
	SubjectBase
	Content RepositoryModifiedSubjectContent `json:"content"`
}

func (sc RepositoryModifiedSubject) GetEventType() CDEventType {
	return RepositoryModifiedEventV1
}

func (sc RepositoryModifiedSubject) GetSubjectType() SubjectType {
	return RepositorySubjectType
}

type RepositoryModifiedEvent struct {
	Context Context                   `json:"context"`
	Subject RepositoryModifiedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e RepositoryModifiedEvent) GetType() CDEventType {
	return RepositoryModifiedEventV1
}

func (e RepositoryModifiedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryModifiedEvent) GetId() string {
	return e.Context.Id
}

func (e RepositoryModifiedEvent) GetSource() string {
	return e.Context.Source
}

func (e RepositoryModifiedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryModifiedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryModifiedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryModifiedEvent) GetSubject() Subject {
	return e.Subject
}

func (e RepositoryModifiedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryModifiedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e RepositoryModifiedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryModifiedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *RepositoryModifiedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryModifiedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryModifiedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryModifiedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryModifiedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *RepositoryModifiedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryModifiedEvent) GetSchema() string {
	return repositoryModifiedSchemaFile
}

// Subject field setters
func (e *RepositoryModifiedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryModifiedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryModifiedEvent) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryModifiedEvent) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

func NewRepositoryModifiedEvent() (*RepositoryModifiedEvent, error) {
	e := &RepositoryModifiedEvent{
		Context: Context{
			Type:    RepositoryModifiedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: RepositoryModifiedSubject{
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
