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

const CDEventsSpecVersion = "draft"

type Context struct {
	// Spec: https://cdevents.dev/docs/spec/#version
	// Description: The version of the CDEvents specification which the event
	// uses. This enables the interpretation of the context. Compliant event
	// producers MUST use a value of draft when referring to this version of the
	// specification.
	Version string `json:"version"`

	// Spec: https://cdevents.dev/docs/spec/#id
	// Description: Identifier for an event. Subsequent delivery attempts of the
	// same event MAY share the same id. This attribute matches the syntax and
	// semantics of the id attribute of CloudEvents:
	// https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#id
	Id string `json:"id"`

	// Spec: https://cdevents.dev/docs/spec/#source
	// Description: defines the context in which an event happened. The main
	// purpose of the source is to provide global uniqueness for source + id.
	// The source MAY identify a single producer or a group of producer that
	// belong to the same application.
	Source string `json:"source"`

	// Spec: https://cdevents.dev/docs/spec/#type
	// Description: defines the type of event, as combination of a subject and
	// predicate. Valid event types are defined in the vocabulary. All event
	// types should be prefixed with dev.cdevents.
	// One occurrence may have multiple events associated, as long as they have
	// different event types
	Type CDEventType `json:"type"`

	// Spec: https://cdevents.dev/docs/spec/#timestamp
	// Description: Description: defines the time of the occurrence. When the
	// time of the occurrence is not available, the time when the event was
	// produced MAY be used. In case the transport layer should require a
	// re-transmission of the event, the timestamp SHOULD NOT be updated, i.e.
	// it should be the same for the same source + id combination.
	Timestamp time.Time `json:"timestamp"`
}

type Reference struct {

	// Spec: https://cdevents.dev/docs/spec/#format-of-subjects
	// Description: Uniquely identifies the subject within the source
	Id string `json:"id"`

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
	Type SubjectType `json:"type"`
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
}

type CDEvent interface {
	CDEventReader
	CDEventWriter
}
