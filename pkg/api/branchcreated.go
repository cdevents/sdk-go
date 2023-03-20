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
	_ "embed"
	"fmt"
	"time"
)

//go:embed spec/schemas/branchcreated.json
var branchcreatedschema string

var (
	// BranchCreated event v0.1.1
	BranchCreatedEventV1 CDEventType = CDEventType{
		Subject:   "branch",
		Predicate: "created",
		Version:   "0.1.1",
	}
)

type BranchCreatedSubjectContent struct {

	// Repository where the branch occurrence happened
	Repository Reference `json:"repository"`
}

type BranchCreatedSubject struct {
	SubjectBase
	Content BranchCreatedSubjectContent `json:"content"`
}

func (sc BranchCreatedSubject) GetEventType() CDEventType {
	return BranchCreatedEventV1
}

func (sc BranchCreatedSubject) GetSubjectType() SubjectType {
	return BranchSubjectType
}

type BranchCreatedEvent struct {
	Context Context              `json:"context"`
	Subject BranchCreatedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e BranchCreatedEvent) GetType() CDEventType {
	return BranchCreatedEventV1
}

func (e BranchCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e BranchCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e BranchCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e BranchCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BranchCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e BranchCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BranchCreatedEvent) GetSubject() Subject {
	return e.Subject
}

func (e BranchCreatedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e BranchCreatedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e BranchCreatedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e BranchCreatedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *BranchCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *BranchCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BranchCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BranchCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BranchCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *BranchCreatedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e BranchCreatedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), branchcreatedschema
}

// Subject field setters
func (e *BranchCreatedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

func NewBranchCreatedEvent() (*BranchCreatedEvent, error) {
	e := &BranchCreatedEvent{
		Context: Context{
			Type:    BranchCreatedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: BranchCreatedSubject{
			SubjectBase: SubjectBase{
				Type: BranchSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
