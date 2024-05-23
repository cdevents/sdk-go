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
	// ArtifactSigned event type v0.2.0
	ArtifactSignedEventTypeV0_2_0 CDEventType = CDEventType{
		Subject:   "artifact",
		Predicate: "signed",
		Version:   "0.2.0",
	}
)

type ArtifactSignedSubjectContentV0_2_0 struct {
	Signature string `json:"signature"`
}

type ArtifactSignedSubjectV0_2_0 struct {
	SubjectBase
	Content ArtifactSignedSubjectContentV0_2_0 `json:"content"`
}

func (sc ArtifactSignedSubjectV0_2_0) GetSubjectType() SubjectType {
	return "artifact"
}

type ArtifactSignedEventV0_2_0 struct {
	Context Context                     `json:"context"`
	Subject ArtifactSignedSubjectV0_2_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ArtifactSignedEventV0_2_0) GetType() CDEventType {
	return ArtifactSignedEventTypeV0_2_0
}

func (e ArtifactSignedEventV0_2_0) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ArtifactSignedEventV0_2_0) GetId() string {
	return e.Context.Id
}

func (e ArtifactSignedEventV0_2_0) GetSource() string {
	return e.Context.Source
}

func (e ArtifactSignedEventV0_2_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactSignedEventV0_2_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactSignedEventV0_2_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactSignedEventV0_2_0) GetSubject() Subject {
	return e.Subject
}

func (e ArtifactSignedEventV0_2_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactSignedEventV0_2_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ArtifactSignedEventV0_2_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactSignedEventV0_2_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ArtifactSignedEventV0_2_0) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactSignedEventV0_2_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactSignedEventV0_2_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactSignedEventV0_2_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactSignedEventV0_2_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ArtifactSignedEventV0_2_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ArtifactSignedEventV0_2_0) GetSchema() (string, string) {
	eType := e.GetType()
	id, schema, _ := GetSchemaBySpecSubjectPredicate(CDEventsSpecVersion, eType.Subject, eType.Predicate)
	return id, schema
}

// Set subject custom fields

func (e *ArtifactSignedEventV0_2_0) SetSubjectSignature(signature string) {
	e.Subject.Content.Signature = signature
}

// New creates a new ArtifactSignedEventV0_2_0
func NewArtifactSignedEventV0_2_0(specVersion string) (*ArtifactSignedEventV0_2_0, error) {
	e := &ArtifactSignedEventV0_2_0{
		Context: Context{
			Type:    ArtifactSignedEventTypeV0_2_0,
			Version: specVersion,
		},
		Subject: ArtifactSignedSubjectV0_2_0{
			SubjectBase: SubjectBase{
				Type: "artifact",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
