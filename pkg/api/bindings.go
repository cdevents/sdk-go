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
	"net/url"
	"regexp"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/go-playground/validator/v10"
	purl "github.com/package-url/packageurl-go"
	"golang.org/x/mod/semver"
)

const (
	SCHEMA_ID_REGEX   = `^https://cdevents.dev/([0-9]\.[0-9])\.[0-9]/schema/([^ ]*)$`
	CustomEventMapKey = "custom"
)

var (
	// Validation helper as singleton
	validate              *validator.Validate
	CDEventsSchemaIdRegex = regexp.MustCompile(SCHEMA_ID_REGEX)
)

func init() {
	// Register custom validators
	validate = validator.New()
	validate.RegisterStructValidation(ValidateEventType, CDEventType{})
	err := validate.RegisterValidation("uri-reference", ValidateUriReference)
	panicOnError(err)
	err = validate.RegisterValidation("purl", ValidatePurl)
	panicOnError(err)
	err = validate.RegisterValidation("event-link-type", ValidateLinkType)
	panicOnError(err)
}

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// TODO(afrittoli) we may want to define something like:
// const CDEventsContentType = "application/cdevents+json"
// but it's not yet in the spec

// ParseType returns a CDEventType if eventType is a valid type
// Since the list of valid events is spec specific, we only validate
// in spec specific code whether this event type exists
func ParseType(eventType string) (*CDEventType, error) {
	t, err := CDEventTypeFromString(eventType)
	if err != nil {
		return nil, err
	}
	if !semver.IsValid("v" + t.Version) {
		return nil, fmt.Errorf("invalid version format %s", t.Version)
	}
	return t, nil
}

func ValidateEventType(sl validator.StructLevel) {
	_, err := ParseType(sl.Current().Interface().(CDEventType).String())
	if err != nil {
		sl.ReportError(sl.Current().Interface(), "Type", "", "", "")
	}
}

func ValidateUriReference(fl validator.FieldLevel) bool {
	_, err := url.Parse(fl.Field().String())
	return err == nil
}

func ValidatePurl(fl validator.FieldLevel) bool {
	_, err := purl.FromString(fl.Field().String())
	return err == nil
}

func ValidateLinkType(fl validator.FieldLevel) bool {
	lt := LinkType(fl.Field().String())
	_, ok := LinkTypes[lt]
	return ok
}

// AsCloudEvent renders a CDEvent as a CloudEvent
func AsCloudEvent(event CDEventReader) (*cloudevents.Event, error) {
	if event == nil {
		return nil, fmt.Errorf("nil CDEvent cannot be rendered as CloudEvent")
	}
	// Validate the event
	err := Validate(event)
	if err != nil {
		return nil, fmt.Errorf("cannot validate CDEvent %v", err)
	}
	ce := cloudevents.NewEvent()
	ce.SetID(event.GetId())
	ce.SetSource(event.GetSource())
	ce.SetSubject(event.GetSubjectId())
	ce.SetType(event.GetType().String())
	err = ce.SetData(cloudevents.ApplicationJSON, event)
	return &ce, err
}

// AsJsonBytes renders a CDEvent as a JSON string
func AsJsonBytes(event CDEventReader) ([]byte, error) {
	if event == nil {
		return nil, nil
	}
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// AsJsonString renders a CDEvent as a JSON string
func AsJsonString(event CDEventReader) (string, error) {
	jsonBytes, err := AsJsonBytes(event)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// Validate checks the CDEvent against the JSON schema and validate constraints
func Validate(event CDEventReader) error {
	_, sch, err := event.GetSchema()
	if err != nil {
		return err
	}
	var v interface{}
	jsonString, err := AsJsonString(event)
	if err != nil {
		return fmt.Errorf("cannot render the event %s as json %s", event, err)
	}
	if err := json.Unmarshal([]byte(jsonString), &v); err != nil {
		return fmt.Errorf("cannot unmarshal event json: %v", err)
	}
	// Validate the "validate" tags
	if err := validate.Struct(event); err != nil {
		return err
	}
	// Validate the "jsonschema" tags
	if err := sch.Validate(v); err != nil {
		return err
	}
	// Check if there is a custom schema
	v4event, ok := event.(CDEventReaderV04)
	if ok {
		schema, err := v4event.GetCustomSchema()
		if err != nil {
			return err
		}
		// If there is no schema defined, we're done
		if schema == nil {
			return nil
		}
		err = schema.Validate(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewFromJsonBytesContext[ContextType] builds a new CDEventReader from a JSON string as []bytes
// This works by unmarshalling the context first, extracting the event type and using
// that to unmarshal the rest of the event into the correct object.
// `ContextType` defines the type of Context that can be used to unmarshal the event.
func NewFromJsonBytesContext[CDEventType CDEvent](event []byte, cdeventsMap map[string]CDEventType) (CDEventType, error) {
	eventAux := &struct {
		Context Context `json:"context"`
	}{}
	var nilReturn, receiver CDEventType
	var ok bool
	err := json.Unmarshal(event, eventAux)
	if err != nil {
		return nilReturn, err
	}
	eventType := eventAux.Context.GetType()
	if eventType.Custom != "" {
		receiver = cdeventsMap[CustomEventMapKey] // Custom type receiver does not have a predefined type
	} else {
		receiver, ok = cdeventsMap[eventType.UnversionedString()]
		if !ok {
			// This should not happen as unmarshalling and validate checks if the type is known to the SDK
			return nilReturn, fmt.Errorf("unknown event type %s", eventAux.Context.GetType())
		}
		// Check if the receiver is compatible. It must have the same subject and predicate
		// and share the same major version.
		// If the minor version is different and the message received as a version that is
		// greater than the SDK one, some fields may be lost, as newer versions may add new
		// fields to the event specification.
		if !eventType.IsCompatible(receiver.GetType()) {
			return nilReturn, fmt.Errorf("sdk event version %s not compatible with %s", receiver.GetType().Version, eventType.Version)
		}
	}
	err = json.Unmarshal(event, receiver)
	if err != nil {
		return nilReturn, err
	}
	return receiver, nil
}
