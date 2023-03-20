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

//go:embed spec/schemas/branchdeleted.json
var branchdeletedschema string

var (
	// BranchDeleted event v0.1.1
	BranchDeletedEventV1 CDEventType = CDEventType{
		Subject:   "branch",
		Predicate: "deleted",
		Version:   "0.1.1",
	}
)

type BranchDeletedSubjectContent struct {

	// Repository where the branch occurrence happened
	Repository Reference `json:"repository"`
}

type BranchDeletedSubject struct {
	SubjectBase
	Content BranchDeletedSubjectContent `json:"content"`
}

func (sc BranchDeletedSubject) GetEventType() CDEventType {
	return BranchDeletedEventV1
}

func (sc BranchDeletedSubject) GetSubjectType() SubjectType {
	return BranchSubjectType
}

type BranchDeletedEvent struct {
	Context Context              `json:"context"`
	Subject BranchDeletedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e BranchDeletedEvent) GetType() CDEventType {
	return BranchDeletedEventV1
}

func (e BranchDeletedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e BranchDeletedEvent) GetId() string {
	return e.Context.Id
}

func (e BranchDeletedEvent) GetSource() string {
	return e.Context.Source
}

func (e BranchDeletedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BranchDeletedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e BranchDeletedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BranchDeletedEvent) GetSubject() Subject {
	return e.Subject
}

func (e BranchDeletedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e BranchDeletedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e BranchDeletedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e BranchDeletedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *BranchDeletedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *BranchDeletedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BranchDeletedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BranchDeletedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BranchDeletedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *BranchDeletedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e BranchDeletedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), branchdeletedschema
}

// Subject field setters
func (e *BranchDeletedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

func NewBranchDeletedEvent() (*BranchDeletedEvent, error) {
	e := &BranchDeletedEvent{
		Context: Context{
			Type:    BranchDeletedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: BranchDeletedSubject{
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
