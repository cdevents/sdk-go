// Code generated by tools/generator. DO NOT EDIT.

//go:build testonly

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
	// FooSubjectBarPredicate event type v1.2.3
	FooSubjectBarPredicateEventTypeV1_2_3 CDEventType = CDEventType{
		Subject:   "foosubject",
		Predicate: "barpredicate",
		Version:   "1.2.3",
	}
)

type FooSubjectBarPredicateSubjectContentV1_2_3 struct {
	ArtifactId string `json:"artifactId,omitempty" validate:"purl"`

	ObjectField *FooSubjectBarPredicateSubjectContentObjectFieldV1_2_3 `json:"objectField,omitempty"`

	PlainField string `json:"plainField"`

	ReferenceField *Reference `json:"referenceField"`
}

type FooSubjectBarPredicateSubjectV1_2_3 struct {
	SubjectBase
	Content FooSubjectBarPredicateSubjectContentV1_2_3 `json:"content"`
}

func (sc FooSubjectBarPredicateSubjectV1_2_3) GetSubjectType() SubjectType {
	return "fooSubject"
}

type FooSubjectBarPredicateEventV1_2_3 struct {
	Context ContextV04                          `json:"context"`
	Subject FooSubjectBarPredicateSubjectV1_2_3 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e FooSubjectBarPredicateEventV1_2_3) GetType() CDEventType {
	return FooSubjectBarPredicateEventTypeV1_2_3
}

func (e FooSubjectBarPredicateEventV1_2_3) GetVersion() string {
	return CDEventsSpecVersion
}

func (e FooSubjectBarPredicateEventV1_2_3) GetId() string {
	return e.Context.Id
}

func (e FooSubjectBarPredicateEventV1_2_3) GetSource() string {
	return e.Context.Source
}

func (e FooSubjectBarPredicateEventV1_2_3) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e FooSubjectBarPredicateEventV1_2_3) GetSubjectId() string {
	return e.Subject.Id
}

func (e FooSubjectBarPredicateEventV1_2_3) GetSubjectSource() string {
	return e.Subject.Source
}

func (e FooSubjectBarPredicateEventV1_2_3) GetSubject() Subject {
	return e.Subject
}

func (e FooSubjectBarPredicateEventV1_2_3) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e FooSubjectBarPredicateEventV1_2_3) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e FooSubjectBarPredicateEventV1_2_3) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e FooSubjectBarPredicateEventV1_2_3) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsReaderV04 implementation

func (e FooSubjectBarPredicateEventV1_2_3) GetChainId() string {
	return e.Context.ChainId
}

func (e FooSubjectBarPredicateEventV1_2_3) GetLinks() EmbeddedLinksArray {
	return e.Context.Links
}

func (e FooSubjectBarPredicateEventV1_2_3) GetSchemaUri() string {
	return e.Context.SchemaUri
}

// CDEventsWriter implementation

func (e *FooSubjectBarPredicateEventV1_2_3) SetId(id string) {
	e.Context.Id = id
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e FooSubjectBarPredicateEventV1_2_3) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return TestCompiledSchemas.GetBySpecSubjectPredicate("99.0.0", eType.Subject, eType.Predicate)
}

// CDEventsWriterV04 implementation

func (e *FooSubjectBarPredicateEventV1_2_3) SetChainId(chainId string) {
	e.Context.ChainId = chainId
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetLinks(links EmbeddedLinksArray) {
	e.Context.Links = links
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSchemaUri(schema string) {
	e.Context.SchemaUri = schema
}

// Set subject custom fields

func (e *FooSubjectBarPredicateEventV1_2_3) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSubjectObjectField(objectField *FooSubjectBarPredicateSubjectContentObjectFieldV1_2_3) {
	e.Subject.Content.ObjectField = objectField
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSubjectPlainField(plainField string) {
	e.Subject.Content.PlainField = plainField
}

func (e *FooSubjectBarPredicateEventV1_2_3) SetSubjectReferenceField(referenceField *Reference) {
	e.Subject.Content.ReferenceField = referenceField
}

// New creates a new FooSubjectBarPredicateEventV1_2_3
func NewFooSubjectBarPredicateEventV1_2_3(specVersion string) (*FooSubjectBarPredicateEventV1_2_3, error) {
	e := &FooSubjectBarPredicateEventV1_2_3{
		Context: ContextV04{
			Context{
				Type:    FooSubjectBarPredicateEventTypeV1_2_3,
				Version: specVersion,
			},
			ContextLinks{},
			ContextCustom{},
		},
		Subject: FooSubjectBarPredicateSubjectV1_2_3{
			SubjectBase: SubjectBase{
				Type: "fooSubject",
			},
		},
	}
	// Pre-seeded with test data
	t, _ := time.Parse(time.RFC3339Nano, "2023-03-20T14:27:05.315384Z")
	e.SetTimestamp(t)
	e.SetId("271069a8-fc18-44f1-b38f-9d70a1695819")
	return e, nil
}

// FooSubjectBarPredicateSubjectContentObjectFieldV1_2_3 holds the content of a ObjectField field in the content
type FooSubjectBarPredicateSubjectContentObjectFieldV1_2_3 struct {
	Optional string `json:"optional,omitempty"`

	Required string `json:"required"`
}
