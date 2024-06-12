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
	// RepositoryDeleted event type v0.2.0
	RepositoryDeletedEventTypeV0_2_0 CDEventType = CDEventType{
		Subject:   "repository",
		Predicate: "deleted",
		Version:   "0.2.0",
	}
)

type RepositoryDeletedSubjectContentV0_2_0 struct {
	Name string `json:"name,omitempty"`

	Owner string `json:"owner,omitempty"`

	Url string `json:"url,omitempty"`

	ViewUrl string `json:"viewUrl,omitempty"`
}

type RepositoryDeletedSubjectV0_2_0 struct {
	SubjectBase
	Content RepositoryDeletedSubjectContentV0_2_0 `json:"content"`
}

func (sc RepositoryDeletedSubjectV0_2_0) GetSubjectType() SubjectType {
	return "repository"
}

type RepositoryDeletedEventV0_2_0 struct {
	Context ContextV04                     `json:"context"`
	Subject RepositoryDeletedSubjectV0_2_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e RepositoryDeletedEventV0_2_0) GetType() CDEventType {
	return RepositoryDeletedEventTypeV0_2_0
}

func (e RepositoryDeletedEventV0_2_0) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryDeletedEventV0_2_0) GetId() string {
	return e.Context.Id
}

func (e RepositoryDeletedEventV0_2_0) GetSource() string {
	return e.Context.Source
}

func (e RepositoryDeletedEventV0_2_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryDeletedEventV0_2_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryDeletedEventV0_2_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryDeletedEventV0_2_0) GetSubject() Subject {
	return e.Subject
}

func (e RepositoryDeletedEventV0_2_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEventV0_2_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e RepositoryDeletedEventV0_2_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEventV0_2_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsReaderV04 implementation

func (e RepositoryDeletedEventV0_2_0) GetChainId() string {
	return e.Context.ChainId
}

func (e RepositoryDeletedEventV0_2_0) GetLinks() EmbeddedLinksArray {
	return e.Context.Links
}

func (e RepositoryDeletedEventV0_2_0) GetSchemaUri() string {
	return e.Context.SchemaUri
}

// CDEventsWriter implementation

func (e *RepositoryDeletedEventV0_2_0) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryDeletedEventV0_2_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryDeletedEventV0_2_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryDeletedEventV0_2_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryDeletedEventV0_2_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *RepositoryDeletedEventV0_2_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryDeletedEventV0_2_0) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.4.1", eType.Subject, eType.Predicate)
}

// CDEventsWriterV04 implementation

func (e *RepositoryDeletedEventV0_2_0) SetChainId(chainId string) {
	e.Context.ChainId = chainId
}

func (e *RepositoryDeletedEventV0_2_0) SetLinks(links EmbeddedLinksArray) {
	e.Context.Links = links
}

func (e *RepositoryDeletedEventV0_2_0) SetSchemaUri(schema string) {
	e.Context.SchemaUri = schema
}

// Set subject custom fields

func (e *RepositoryDeletedEventV0_2_0) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryDeletedEventV0_2_0) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryDeletedEventV0_2_0) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryDeletedEventV0_2_0) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

// New creates a new RepositoryDeletedEventV0_2_0
func NewRepositoryDeletedEventV0_2_0(specVersion string) (*RepositoryDeletedEventV0_2_0, error) {
	e := &RepositoryDeletedEventV0_2_0{
		Context: ContextV04{
			Context{
				Type:    RepositoryDeletedEventTypeV0_2_0,
				Version: specVersion,
			},
			ContextLinks{},
			ContextCustom{},
		},
		Subject: RepositoryDeletedSubjectV0_2_0{
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
