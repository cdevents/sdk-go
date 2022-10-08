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
	// BuildStarted event
	BuildStartedEventV1    CDEventType = "dev.cdevents.build.started.0.1.0-draft"
	buildStartedSchemaFile string      = "buildstarted"
)

type BuildStartedSubjectContent struct{}

type BuildStartedSubject struct {
	SubjectBase
	Content BuildStartedSubjectContent `json:"content"`
}

func (sc BuildStartedSubject) GetEventType() CDEventType {
	return BuildStartedEventV1
}

func (sc BuildStartedSubject) GetSubjectType() SubjectType {
	return BuildSubjectType
}

type BuildStartedEvent struct {
	Context Context             `json:"context"`
	Subject BuildStartedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e BuildStartedEvent) GetType() CDEventType {
	return BuildStartedEventV1
}

func (e BuildStartedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e BuildStartedEvent) GetId() string {
	return e.Context.Id
}

func (e BuildStartedEvent) GetSource() string {
	return e.Context.Source
}

func (e BuildStartedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BuildStartedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e BuildStartedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BuildStartedEvent) GetSubject() Subject {
	return e.Subject
}

func (e BuildStartedEvent) GetCustomData() (interface{}, error) {
	return getCustomData(e.CustomDataContentType, e.CustomData)
}

func (e BuildStartedEvent) GetCustomDataAs(receiver interface{}) error {
	return getCustomDataAs(e, receiver)
}

func (e BuildStartedEvent) GetCustomDataRaw() ([]byte, error) {
	return getCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e BuildStartedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *BuildStartedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *BuildStartedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BuildStartedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BuildStartedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BuildStartedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *BuildStartedEvent) SetCustomData(contentType string, data interface{}) error {
	err := checkCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e BuildStartedEvent) GetSchema() string {
	return buildStartedSchemaFile
}

func NewBuildStartedEvent() (*BuildStartedEvent, error) {
	e := &BuildStartedEvent{
		Context: Context{
			Type:    BuildStartedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: BuildStartedSubject{
			SubjectBase: SubjectBase{
				Type: BuildSubjectType,
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
