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
	"reflect"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

// TODO(afrittoli) we may want to define something like:
// const CDEventsContentType = "application/cdevents+json"
// but it's not yet in the spec

// AsCloudEvent renders a CDEvent as a CloudEvent
func AsCloudEvent(event CDEventReader) (*cloudevents.Event, error) {
	if event == nil {
		return nil, fmt.Errorf("nil CDEvent cannot be rendered as CloudEvent")
	}
	err := Validate(event)
	if err != nil {
		return nil, fmt.Errorf("cannot validate CDEvent %v", err)
	}
	ce := cloudevents.NewEvent()
	// TODO(afrittoli) We should have some validation in place
	ce.SetSource(event.GetSource())
	ce.SetSubject(event.GetSubjectId())
	ce.SetType(event.GetType().String())
	err = ce.SetData(cloudevents.ApplicationJSON, event)
	return &ce, err
}

// AsJsonString renders a CDEvent as a JSON string
func AsJsonString(event CDEventReader) (string, error) {
	if event == nil {
		return "", nil
	}
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// Validate checks the CDEvent against the JSON schema
func Validate(event CDEventReader) error {
	schemaName := event.GetSchema()
	compiler := jsonschema.NewCompiler()
	sch, err := compiler.Compile(fmt.Sprintf("../../jsonschema/%s.json", schemaName))
	if err != nil {
		return fmt.Errorf("cannot compile jsonschema %s, %s", schemaName, err)
	}
	eventType := reflect.TypeOf(event)
	sch.Types = append(sch.Types, eventType.String())
	// var v interface{}
	// jsonString, err := AsJsonString(event)
	// if err != nil {
	// 	return fmt.Errorf("cannot render the event %s as json %s", event, err)
	// }
	// if err := json.Unmarshal([]byte(jsonString), &v); err != nil {
	// 	return fmt.Errorf("cannot unmarshal event json: %v", err)
	// }
	return sch.Validate(event)
}
