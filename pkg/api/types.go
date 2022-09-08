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
	"encoding/json"
	"fmt"
	"time"
)

const CDEventsSpecVersion = "draft"

type Context struct {
	// Spec: https://cdevents.dev/docs/spec/#version
	// Description: The version of the CDEvents specification which the event
	// uses. This enables the interpretation of the context. Compliant event
	// producers MUST use a value of draft when referring to this version of the
	// specification.
	Version string `json:"version" jsonschema:"required,enum=draft,default=draft"`

	// Spec: https://cdevents.dev/docs/spec/#id
	// Description: Identifier for an event. Subsequent delivery attempts of the
	// same event MAY share the same id. This attribute matches the syntax and
	// semantics of the id attribute of CloudEvents:
	// https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#id
	Id string `json:"id" jsonschema:"required,minLength=1"`

	// Spec: https://cdevents.dev/docs/spec/#source
	// Description: defines the context in which an event happened. The main
	// purpose of the source is to provide global uniqueness for source + id.
	// The source MAY identify a single producer or a group of producer that
	// belong to the same application.
	Source string `json:"source" jsonschema:"required,minLength=1"`

	// Spec: https://cdevents.dev/docs/spec/#type
	// Description: defines the type of event, as combination of a subject and
	// predicate. Valid event types are defined in the vocabulary. All event
	// types should be prefixed with dev.cdevents.
	// One occurrence may have multiple events associated, as long as they have
	// different event types
	Type CDEventType `json:"type" jsonschema:"required,minLength=1"`

	// Spec: https://cdevents.dev/docs/spec/#timestamp
	// Description: Description: defines the time of the occurrence. When the
	// time of the occurrence is not available, the time when the event was
	// produced MAY be used. In case the transport layer should require a
	// re-transmission of the event, the timestamp SHOULD NOT be updated, i.e.
	// it should be the same for the same source + id combination.
	Timestamp time.Time `json:"timestamp" jsonschema:"required"`
}

type Reference struct {

	// Spec: https://cdevents.dev/docs/spec/#format-of-subjects
	// Description: Uniquely identifies the subject within the source
	Id string `json:"id" jsonschema:"required,minLength=1"`

	// Spec: https://cdevents.dev/docs/spec/#format-of-subjects
	// Description: defines the context in which an event happened. The main
	// purpose of the source is to provide global uniqueness for source + id.
	// The source MAY identify a single producer or a group of producer that
	// belong to the same application.
	Source string `json:"source,omitempty"`
}

type SubjectBase struct {
	Reference

	// The type of subject. Constraints what is a valid valid SubjectContent
	Type SubjectType `json:"type" jsonschema:"required,minLength=1"`
}

type SubjectType string

func (t SubjectType) String() string {
	return string(t)
}

type Subject interface {
	GetEventType() CDEventType

	GetSubjectType() SubjectType
}

type CDEventType string

func (t CDEventType) String() string {
	return string(t)
}

type CDEventReader interface {

	// The CDEventType "dev.cdevents.*"
	GetType() CDEventType

	// The CDEvents specification version implemented
	GetVersion() string

	// The event ID, unique for this event within the event producer (source)
	GetId() string

	// The source of the event
	GetSource() string

	// The time when the occurrence described in the event happened, or when
	// the event was produced if the former is not available
	GetTimestamp() time.Time

	// The ID of the subject, unique within the event producer (source), it may
	// by used in multiple events
	GetSubjectId() string

	// The source of the subject. Usually this matches the source of the event
	// but it may also be different.
	GetSubjectSource() string

	// The event specific subject. It is possible to use a type assertion with
	// the generic Subject to obtain an event specific implementation of Subject
	// for direct access to the content fields
	GetSubject() Subject

	// The name of the schema file associated to the event type
	GetSchema() string

	// The custom data attached to the event
	// Depends on GetCustomDataContentType()
	// - When "application/json", un-marshalled data
	// - Else, raw []byte
	GetCustomData() (interface{}, error)

	// The raw custom data attached to the event
	GetCustomDataRaw() ([]byte, error)

	// Custom data un-marshalled into receiver, only if
	// GetCustomDataContentType() returns "application/json", else error
	GetCustomDataAs(receiver interface{}) error

	// Custom data content-type
	GetCustomDataContentType() string
}

type CDEventWriter interface {

	// The event ID, unique for this event within the event producer (source)
	SetId(id string)

	// The source of the event
	SetSource(source string)

	// The time when the occurrence described in the event happened, or when
	// the event was produced if the former is not available
	SetTimestamp(timestamp time.Time)

	// The ID of the subject, unique within the event producer (source), it may
	// by used in multiple events
	SetSubjectId(subjectId string)

	// The source of the subject. Usually this matches the source of the event
	// but it may also be different.
	SetSubjectSource(subjectSource string)

	// Set custom data. If contentType is "application/json", data can also be
	// anything that can be marshalled into json. For any other
	// content type, data must be passed as a []byte
	SetCustomData(contentType string, data interface{}) error
}

type CDEventCustomDataEncoding string

func (t CDEventCustomDataEncoding) String() string {
	return string(t)
}

// CDEventCustomData hosts the CDEvent custom data fields
//
// `CustomDataContentType` describes the content type of the data.
//
// `CustomData` contains the data:
//
// - When the content type is "application/json":
//
//   - if the CDEvent is produced via the golang API, the `CustomData`
//     usually holds an un-marshalled golang interface{} of some type
//
//   - if the CDEvent is consumed and thus un-marshalled from a []byte
//     the `CustomData` holds the data as a []byte, so that it may be
//     un-marshalled into a specific golang type via the `GetCustomDataAs`
//
// - When the content type is anything else:
//
//   - the content data is always stored as []byte, as the SDK does not
//     have enough knowledge about the data to un-marshal it into a
//     golang type
type CDEventCustomData struct {

	// CustomData added to the CDEvent. Format not specified by the SPEC.
	CustomData interface{} `json:"customData,omitempty" jsonschema:"oneof_type=object;string"`

	// CustomDataContentType for CustomData in a CDEvent.
	CustomDataContentType string `json:"customDataContentType,omitempty"`
}

// Customize the Unmarshal into *CDEventCustomData
// Only unmarshal the content type, and let the data as []byte.
// The unmarshal of the data is postponed to the GetCustomData* functions
func (d *CDEventCustomData) UnMarshalJSON(data []byte) error {
	// First read the content type
	cc := &struct {
		RawData     []byte `json:"customData,omitempty"`
		ContentType string `json:"customDataContentType,omitempty"`
	}{}
	if err := json.Unmarshal(data, &cc); err != nil {
		return err
	}
	d.CustomDataContentType = cc.ContentType
	d.CustomData = cc.RawData
	return nil
}

type CDEvent interface {
	CDEventReader
	CDEventWriter
}

// Used to implement GetCustomDataRaw()
func getCustomDataRaw(contentType string, data interface{}) ([]byte, error) {
	switch data := data.(type) {
	case []byte:
		return data, nil
	default:
		if contentType != "application/json" {
			return nil, fmt.Errorf("cannot use %v with content type %s", data, contentType)
		}
		// The content type is JSON, but the data is un-marshalled
		return json.Marshal(data)
	}
}

// Used to implement GetCustomDataAs()
func getCustomDataAs(e CDEventReader, receiver interface{}) error {
	if e.GetCustomDataContentType() != "application/json" {
		return fmt.Errorf("cannot unmarshal content-type %s", e.GetCustomDataContentType())
	}
	data, err := e.GetCustomDataRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, receiver)
}

// Used to implement GetCustomData()
func getCustomData(contentType string, data interface{}) (interface{}, error) {
	var v interface{}
	switch data := data.(type) {
	case []byte:
		// The data is JSON but still raw. Let's un-marshal it.
		if contentType == "application/json" {
			err := json.Unmarshal(data, &v)
			if err != nil {
				return nil, err
			}
			return v, nil
		}
		// The content type is not JSON, pass through raw data
		return data, nil
	default:
		if contentType != "application/json" {
			return nil, fmt.Errorf("cannot use %v with content type %s", data, contentType)
		}
		// The content type is JSON, pass through un-marshalled data
		return data, nil
	}
}

// Used to implement SetCustomData()
func checkCustomData(contentType string, data interface{}) error {
	_, isBytes := data.([]byte)
	if !isBytes && contentType != "application/json" {
		return fmt.Errorf("%s data must be set as []bytes, got %v", contentType, data)
	}
	return nil
}
