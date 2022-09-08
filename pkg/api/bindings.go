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

	cloudevents "github.com/cloudevents/sdk-go/v2"
	schemaproducer "github.com/invopop/jsonschema"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

var (

	// All event types
	allEvents = []CDEvent{
		&PipelineRunQueuedEvent{},
		&PipelineRunStartedEvent{},
		&PipelineRunFinishedEvent{},
		&TaskRunStartedEvent{},
		&TaskRunFinishedEvent{},
		&ChangeCreatedEvent{},
		&ChangeUpdatedEvent{},
		&ChangeReviewedEvent{},
		&ChangeMergedEvent{},
		&ChangeAbandonedEvent{},
		&RepositoryCreatedEvent{},
		&RepositoryModifiedEvent{},
		&RepositoryDeletedEvent{},
		&BranchCreatedEvent{},
		&BranchDeletedEvent{},
		&TestSuiteStartedEvent{},
		&TestSuiteFinishedEvent{},
		&TestCaseQueuedEvent{},
		&TestCaseStartedEvent{},
		&TestCaseFinishedEvent{},
		&BuildQueuedEvent{},
		&BuildStartedEvent{},
		&BuildFinishedEvent{},
		&ArtifactPackagedEvent{},
		&ArtifactPublishedEvent{},
		&EnvironmentCreatedEvent{},
		&EnvironmentModifiedEvent{},
		&EnvironmentDeletedEvent{},
		&ServiceDeployedEvent{},
		&ServiceUpgradedEvent{},
		&ServiceRolledbackEvent{},
		&ServiceRemovedEvent{},
		&ServicePublishedEvent{},
	}

	// Map schema names to schema strings
	allEventSchemas map[string]string

	// Map CDEventType to empty event objects
	cdeventsByTypes map[CDEventType]CDEvent
)

func init() {

	// Init the schema map
	allEventSchemas = make(map[string]string)
	cdeventsByTypes = make(map[CDEventType]CDEvent)

	// Setup a reflector
	id := schemaproducer.EmptyID
	id = id.Add(fmt.Sprintf("https://cdevents.dev/%s/schema", CDEventsSpecVersion))
	reflector := schemaproducer.Reflector{
		BaseSchemaID:   id,
		DoNotReference: true,
	}

	for _, eventType := range allEvents {
		// Setup schema strings
		s := reflector.Reflect(eventType)
		data, err := json.MarshalIndent(s, "", "  ")
		panicOnError(err)
		allEventSchemas[eventType.GetSchema()] = string(data)

		// Set type to receiver map
		cdeventsByTypes[eventType.GetType()] = eventType
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// TODO(afrittoli) we may want to define something like:
// const CDEventsContentType = "application/cdevents+json"
// but it's not yet in the spec

// ParseType returns a CDEventType is eventType is a valid type
func ParseType(eventType string) (CDEventType, error) {
	t := CDEventType(eventType)
	_, ok := cdeventsByTypes[t]
	if !ok {
		return "", fmt.Errorf("unknown event type %s", eventType)
	}
	return t, nil
}

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
	sch, err := jsonschema.CompileString(fmt.Sprintf("%s.json", schemaName), allEventSchemas[schemaName])
	if err != nil {
		return fmt.Errorf("cannot compile jsonschema %s, %s", schemaName, err)
	}
	var v interface{}
	jsonString, err := AsJsonString(event)
	if err != nil {
		return fmt.Errorf("cannot render the event %s as json %s", event, err)
	}
	if err := json.Unmarshal([]byte(jsonString), &v); err != nil {
		return fmt.Errorf("cannot unmarshal event json: %v", err)
	}
	return sch.Validate(v)
}

// Build a new CDEventReader from a JSON string
func NewFromJsonString(event string) (CDEvent, error) {
	return NewFromJsonBytes([]byte(event))
}

// Build a new CDEventReader from a JSON string as []bytes
func NewFromJsonBytes(event []byte) (CDEvent, error) {
	eventAux := &struct {
		Context Context `json:"context"`
	}{}
	err := json.Unmarshal(event, eventAux)
	if err != nil {
		return nil, err
	}
	receiver, ok := cdeventsByTypes[CDEventType(eventAux.Context.Type)]
	if !ok {
		return nil, fmt.Errorf("unknown event type %s", eventAux.Context.Type)
	}
	err = json.Unmarshal(event, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}
