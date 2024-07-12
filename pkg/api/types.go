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
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
	"golang.org/x/mod/semver"
)

const (
	EventTypeRoot                   = "dev.cdevents"
	CustomEventTypeRoot             = "dev.cdeventsx"
	CDEventsSchemaURLTemplate       = "https://cdevents.dev/%s/schema/%s-%s-event"
	CDEventsCustomSchemaURLTemplate = "https://cdevents.dev/%s/schema/custom"
	CDEventsTypeRegex               = "^dev\\.cdevents\\.(?P<subject>[a-z]+)\\.(?P<predicate>[a-z]+)\\.(?P<version>.*)$"
	CDEventsCustomTypeRegex         = "^dev\\.cdeventsx\\.(?P<tool>[a-z]+)-(?P<subject>[a-z]+)\\.(?P<predicate>[a-z]+)\\.(?P<version>.*)$"

	LinkTypePath     LinkType = "PATH"
	LinkTypeEnd      LinkType = "END"
	LinkTypeRelation LinkType = "RELATION"
)

var (
	CDEventsTypeCRegex       = regexp.MustCompile(CDEventsTypeRegex)
	CDEventsCustomTypeCRegex = regexp.MustCompile(CDEventsCustomTypeRegex)
	LinkTypes                = map[LinkType]interface{}{
		LinkTypePath:     "",
		LinkTypeEnd:      "",
		LinkTypeRelation: "",
	}
)

type BaseContextReader interface {

	// GetVersion returns the CDEvents spec version
	GetVersion() string

	// GetType returns the CDEvents event type as string
	GetType() CDEventType
}

func (t Context) GetVersion() string {
	return t.Version
}

func (t Context) GetType() CDEventType {
	return t.Type
}

type Context struct {
	// Spec: https://cdevents.dev/docs/spec/#version
	// Description: The version of the CDEvents specification which the event
	// uses. This enables the interpretation of the context. Compliant event
	// producers MUST use a value of draft when referring to this version of the
	// specification.
	Version string `json:"version" jsonschema:"required"`

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
	Source string `json:"source" jsonschema:"required,minLength=1" validate:"uri-reference"`

	// Spec: https://cdevents.dev/docs/spec/#type
	// Description: defines the type of event, as combination of a subject and
	// predicate. Valid event types are defined in the vocabulary. All event
	// types should be prefixed with dev.cdevents.
	// One occurrence may have multiple events associated, as long as they have
	// different event types
	Type CDEventType `json:"type" jsonschema:"required,minLength=1" validate:"required,structonly"`

	// Spec: https://cdevents.dev/docs/spec/#timestamp
	// Description: defines the time of the occurrence. When the
	// time of the occurrence is not available, the time when the event was
	// produced MAY be used. In case the transport layer should require a
	// re-transmission of the event, the timestamp SHOULD NOT be updated, i.e.
	// it should be the same for the same source + id combination.
	Timestamp time.Time `json:"timestamp" jsonschema:"required"`
}

type Tags map[string]interface{}

type LinkType string

type EmbeddedLink interface {
	// GetLinkType returns the content of the jsonschema "linkType"
	GetLinkType() LinkType
}

type EmbeddedLinkWithTags interface {
	EmbeddedLink

	// GetTags returns the content of the jsonschema "tags" object field
	// which defines no property and allows for additional ones
	GetTags() Tags

	// SetTags sets the content of the jsonschema "tags" object field
	SetTags(tags Tags)
}

type EmbeddedLinkWithTagsAndSource interface {
	EmbeddedLinkWithTags

	// GetFrom returns the source of the link, in the "from" field
	GetFrom() EventReference

	// SetFrom sets the source of the link, in the "from" field
	SetFrom(reference EventReference)
}

type EmbeddedLinkWithTagsAndRelation interface {
	EmbeddedLinkWithTags

	// GetTarget returns the target of the link, in the "target" field
	GetTarget() EventReference

	// SetTarget sets the target of the link, in the "target" field
	SetTarget(reference EventReference)

	// GetLinkKind returns the link kind, in the "linkKind" field
	GetLinkKind() string

	// SetLinkKind sets the kind of the link, in the "linkKind" field
	SetLinkKind(kind string)
}

// EventReference contains the ID of a linked event
type EventReference struct {
	// ContextId is the ID of the linked event
	ContextId string `json:"contextId" jsonschema:"required,minLength=1"`
}

// embeddedLinkPath is private so that NewEmbeddedLinkPath must be used
// to create an object with correct defaults
type embeddedLinkPath struct {
	LinkType LinkType       `json:"linkType" jsonschema:"required,minLength=1" validate:"event-link-type"`
	From     EventReference `json:"from" jsonschema:"required,minLength=1"`
	Tags     Tags           `json:"tags"`
}

func (l embeddedLinkPath) GetLinkType() LinkType {
	return l.LinkType
}

func (l embeddedLinkPath) GetTags() Tags {
	return l.Tags
}

func (l embeddedLinkPath) GetFrom() EventReference {
	return l.From
}

func (l *embeddedLinkPath) SetTags(tags Tags) {
	l.Tags = tags
}

func (l *embeddedLinkPath) SetFrom(from EventReference) {
	l.From = from
}

func NewEmbeddedLinkPath() EmbeddedLinkWithTagsAndSource {
	return &embeddedLinkPath{
		LinkType: LinkTypePath,
	}
}

// embeddedLinkPath is private so that NewEmbeddedLinkPath must be used
// to create an object with correct defaults
type embeddedLinkEnd struct {
	LinkType LinkType       `json:"linkType" jsonschema:"required,minLength=1" validate:"event-link-type"`
	From     EventReference `json:"from" jsonschema:"required,minLength=1"`
	Tags     Tags           `json:"tags"`
}

func (l embeddedLinkEnd) GetLinkType() LinkType {
	return l.LinkType
}

func (l embeddedLinkEnd) GetTags() Tags {
	return l.Tags
}

func (l embeddedLinkEnd) GetFrom() EventReference {
	return l.From
}

func (l *embeddedLinkEnd) SetTags(tags Tags) {
	l.Tags = tags
}

func (l *embeddedLinkEnd) SetFrom(from EventReference) {
	l.From = from
}

func NewEmbeddedLinkEnd() EmbeddedLinkWithTagsAndSource {
	return &embeddedLinkEnd{
		LinkType: LinkTypeEnd,
	}
}

// embeddedLinkPath is private so that NewEmbeddedLinkPath must be used
// to create an object with correct defaults
type embeddedLinkRelation struct {
	LinkType LinkType       `json:"linkType" jsonschema:"required,minLength=1" validate:"event-link-type"`
	LinkKind string         `json:"linkKind" jsonschema:"required,minLength=1"`
	Target   EventReference `json:"target" jsonschema:"required,minLength=1"`
	Tags     Tags           `json:"tags"`
}

func (l embeddedLinkRelation) GetLinkType() LinkType {
	return l.LinkType
}

func (l embeddedLinkRelation) GetLinkKind() string {
	return l.LinkKind
}

func (l embeddedLinkRelation) GetTags() Tags {
	return l.Tags
}

func (l embeddedLinkRelation) GetTarget() EventReference {
	return l.Target
}

func (l *embeddedLinkRelation) SetLinkKind(linkKind string) {
	l.LinkKind = linkKind
}

func (l *embeddedLinkRelation) SetTags(tags Tags) {
	l.Tags = tags
}

func (l *embeddedLinkRelation) SetTarget(target EventReference) {
	l.Target = target
}

func NewEmbeddedLinkRelation() EmbeddedLinkWithTagsAndRelation {
	return &embeddedLinkRelation{
		LinkType: LinkTypeRelation,
	}
}

type EmbeddedLinksArray []EmbeddedLinkWithTags

func (ela *EmbeddedLinksArray) UnmarshalJSON(b []byte) error {
	var rawEmbeddedLinks []*json.RawMessage
	err := json.Unmarshal(b, &rawEmbeddedLinks)
	if err != nil {
		return err
	}

	m := &struct {
		LinkType LinkType `json:"linkType"`
	}{}
	receiver := make([]EmbeddedLinkWithTags, len(rawEmbeddedLinks))
	for index, rawEmbeddedLink := range rawEmbeddedLinks {
		err = json.Unmarshal(*rawEmbeddedLink, &m)
		if err != nil {
			return err
		}
		if m.LinkType == LinkTypeEnd {
			var e embeddedLinkEnd
			err = json.Unmarshal(*rawEmbeddedLink, &e)
			if err != nil {
				return err
			}
			receiver[index] = &e
		} else if m.LinkType == LinkTypePath {
			var e embeddedLinkPath
			err = json.Unmarshal(*rawEmbeddedLink, &e)
			if err != nil {
				return err
			}
			receiver[index] = &e
		} else if m.LinkType == LinkTypeRelation {
			var e embeddedLinkRelation
			err = json.Unmarshal(*rawEmbeddedLink, &e)
			if err != nil {
				return err
			}
			receiver[index] = &e
		} else {
			return fmt.Errorf("unsupported link type %s found", m.LinkType)
		}
	}
	*ela = receiver
	return nil
}

type ContextLinks struct {
	// Spec: https://cdevents.dev/docs/spec/#chain_id
	// Description: Identifier for a chain as defined in the links spec
	// https://github.com/cdevents/spec/blob/v0.4.1/links.md
	ChainId string `json:"chainId,omitempty"`

	// Spec: https://cdevents.dev/docs/spec/#links
	// Description: Identifier for an event. Subsequent delivery attempts of the
	// same event MAY share the same id. This attribute matches the syntax and
	// semantics of the id attribute of CloudEvents:
	// https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#id
	Links EmbeddedLinksArray `json:"links,omitempty" validate:"dive"`
}

type ContextCustom struct {
	// Spec: https://cdevents.dev/docs/spec/#schemauri
	// Description: ink to a jsonschema schema that further refines
	// the event schema as defined by CDEvents.
	SchemaUri string `json:"schemaUri,omitempty"`
}

type ContextV04 struct {
	Context
	ContextLinks
	ContextCustom
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
	Source string `json:"source,omitempty" validate:"uri-reference"`
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
	GetSubjectType() SubjectType
}

type CDEventType struct {
	Subject   string
	Predicate string

	// Version is a semantic version in the form <major>.<minor>.<patch>
	Version string

	// Custom holds the tool name in case of custom events
	Custom string
}

func (t CDEventType) Root() string {
	root := EventTypeRoot
	if t.Custom != "" {
		root = CustomEventTypeRoot
	}
	return root
}

// FQSubject returns the fully qualified subject, which includes
// the tool name from t.Custom in case of custom events
func (t CDEventType) FQSubject() string {
	s := t.Subject
	if s == "" {
		s = "<undefined-subject>"
	}
	if t.Custom != "" {
		s = t.Custom + "-" + s
	}
	return s
}

func (t CDEventType) String() string {
	predicate := t.Predicate
	if predicate == "" {
		predicate = "<undefined-predicate>"
	}
	version := t.Version
	if version == "" {
		version = "<undefined-version>"
	}
	return t.Root() + "." + t.FQSubject() + "." + predicate + "." + version
}

func (t CDEventType) UnversionedString() string {
	predicate := t.Predicate
	if predicate == "" {
		predicate = "<undefined-predicate>"
	}
	return t.Root() + "." + t.FQSubject() + "." + predicate
}

func (t CDEventType) Short() string {
	s := t.FQSubject()
	p := t.Predicate
	if s == "" || p == "" {
		return ""
	}
	return t.FQSubject() + "_" + t.Predicate
}

// Two CDEventTypes are compatible if the subject and predicates
// are identical and they share the same major version
func (t CDEventType) IsCompatible(other CDEventType) bool {
	return t.Predicate == other.Predicate &&
		t.Subject == other.Subject &&
		semver.Major("v"+t.Version) == semver.Major("v"+other.Version)
}

func (t *CDEventType) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	cdeventType, err := CDEventTypeFromString(strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}
	*t = *cdeventType
	return nil
}

func (t CDEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func CDEventTypeFromString(cdeventType string) (*CDEventType, error) {
	names := CDEventsTypeCRegex.SubexpNames()
	parts := CDEventsTypeCRegex.FindStringSubmatch(cdeventType)
	if len(parts) != 4 {
		names = CDEventsCustomTypeCRegex.SubexpNames()
		parts = CDEventsCustomTypeCRegex.FindStringSubmatch(cdeventType)
		if len(parts) != 5 {
			return nil, fmt.Errorf("cannot parse event type %s", cdeventType)
		}
	}
	returnType := CDEventType{}
	for i, matchName := range names {
		if i == 0 {
			continue
		}
		switch matchName {
		case "subject":
			returnType.Subject = parts[i]
		case "predicate":
			returnType.Predicate = parts[i]
		case "version":
			returnType.Version = parts[i]
		case "tool":
			returnType.Custom = parts[i]
		}
	}
	return &returnType, nil
}

type CDEventReader interface {

	// Event type and spec version readers
	BaseContextReader

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

	// The event specific subject. It is possible to use a type assertion with
	// the generic Subject to obtain an event specific implementation of Subject
	// for direct access to the content fields
	GetSubjectContent() interface{}

	// The URL and content of the schema file associated to the event type
	GetSchema() (string, *jsonschema.Schema, error)

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

type CDEventReaderV04 interface {
	CDEventReader

	// The ChainId for the event
	GetChainId() string

	// The links array for the event
	GetLinks() EmbeddedLinksArray

	// The custom schema URI
	GetSchemaUri() string

	// The compiled schemaUri for the event
	GetCustomSchema() (*jsonschema.Schema, error)
}

type CDEventWriterV04 interface {
	CDEventWriter

	// The ChainId for the event
	SetChainId(chainId string)

	// The links array for the event
	SetLinks(links EmbeddedLinksArray)

	// The custom schema URI
	SetSchemaUri(schema string)
}

type CustomCDEventReader interface {
	CDEventReaderV04
}

type CustomCDEventWriter interface {
	CDEventWriterV04

	// CustomCDEvent can represent different event types
	SetEventType(eventType CDEventType)

	// CustomCDEvent types can have different subject fields
	SetSubjectContent(subjectContent interface{})
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
//     can hold an un-marshalled golang interface{} of a specific type
//     or a marshalled byte slice
//
//   - if the CDEvent is consumed and thus un-marshalled from a []byte
//     the `CustomData` holds the data un-marshalled from []byte, into
//     a generic interface{}. It may be un-marshalled into a specific
//     golang type via the `GetCustomDataAs`
//
// - When the content type is anything else:
//
//   - if the CDEvent is produced via the golang API, the `CustomData`
//     hold an byte slice with the data passed via the API
//
//   - if the CDEvent is consumed and thus un-marshalled from a []byte
//     the `CustomData` holds the data base64 encoded
type CDEventCustomData struct {

	// CustomData added to the CDEvent. Format not specified by the SPEC.
	CustomData interface{} `json:"customData,omitempty" jsonschema:"oneof_type=object;string"`

	// CustomDataContentType for CustomData in a CDEvent.
	CustomDataContentType string `json:"customDataContentType,omitempty"`
}

type CDEvent interface {
	CDEventReader
	CDEventWriter
}

type CDEventV04 interface {
	CDEventReaderV04
	CDEventWriterV04
}

// Used to implement type specific GetCustomDataRaw()
func GetCustomDataRaw(contentType string, data interface{}) ([]byte, error) {
	switch data := data.(type) {
	case []byte:
		return data, nil
	default:
		if contentType != "application/json" && contentType != "" {
			return nil, fmt.Errorf("cannot use %v with content type %s", data, contentType)
		}
		// The content type is JSON, but the data is un-marshalled
		return json.Marshal(data)
	}
}

// Used to implement type specific GetCustomDataAs()
func GetCustomDataAs(e CDEventReader, receiver interface{}) error {
	contentType := e.GetCustomDataContentType()
	if contentType != "application/json" && contentType != "" {
		return fmt.Errorf("cannot unmarshal content-type %s", contentType)
	}
	data, err := e.GetCustomDataRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, receiver)
}

// Used to implement type specific GetCustomData()
func GetCustomData(contentType string, data interface{}) (interface{}, error) {
	var v interface{}
	if contentType == "" {
		contentType = "application/json"
	}
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
	case string:
		if contentType == "application/json" {
			return nil, fmt.Errorf("content type %s should not be a string: %s", contentType, data)
		}
		// If the data is not "application/json", and it's a string after
		// un-marshalling, we assume it's base64 encoded
		// NOTE(afrittoli) The standard un-marshaller would decode if the
		// receiving type was []byte, but we have interface because we need
		// to be able to store golang objects as well
		return b64.StdEncoding.DecodeString(data)
	default:
		if contentType != "application/json" {
			return nil, fmt.Errorf("cannot use %v with content type %s", data, contentType)
		}
		// The content type is JSON, pass through un-marshalled data
		return data, nil
	}
}

// Used to implement SetCustomData()
func CheckCustomData(contentType string, data interface{}) error {
	_, isBytes := data.([]byte)
	if !isBytes && contentType != "application/json" && contentType != "" {
		return fmt.Errorf("%s data must be set as []bytes, got %v", contentType, data)
	}
	return nil
}
