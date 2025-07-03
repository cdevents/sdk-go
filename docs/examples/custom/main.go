package main

import (
	"context"
	"fmt"
	"log"
	"os"

	examples "github.com/cdevents/sdk-go/docs/examples"
	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

const customSchemaURI = "https://myregistry.dev/schemas/cdevents/quota-exceeded/0_1_0"

type Quota struct {
	User      string `json:"user"`      // The use the quota applies ot
	Limit     string `json:"limit"`     // The limit enforced by the quota e.g. 100Gb
	Current   int    `json:"current"`   // The current % of the quota used e.g. 90%
	Threshold int    `json:"threshold"` // The threshold for warning event e.g. 85%
	Level     string `json:"level"`     // INFO: <threshold, WARNING: >threshold, <quota, CRITICAL: >quota
}

func main() {
	var ce *cloudevents.Event
	var c cloudevents.Client

	// Define the event type
	eventType := cdevents.CDEventType{
		Subject:   "quota",
		Predicate: "exceeded",
		Version:   "0.1.0",
		Custom:    "myregistry",
	}

	// Define the content
	quotaRule123 := Quota{
		User:      "heavy_user",
		Limit:     "50Tb",
		Current:   90,
		Threshold: 85,
		Level:     "WARNING",
	}

	// Create the base event
	event, err := cdeventsv04.NewCustomTypeEvent()
	examples.PanicOnError(err, "could not create a cdevent")
	event.SetEventType(eventType)

	// Set the required context fields
	event.SetSubjectId("quotaRule123")
	event.SetSource("my/first/cdevent/program")

	// Set the required subject fields
	event.SetSubjectContent(quotaRule123)
	event.SetSchemaUri(customSchemaURI)

	// Print the event
	eventJSON, err := cdevents.AsJSONString(event)
	examples.PanicOnError(err, "failed to marshal the CDEvent")
	fmt.Printf("%s", eventJSON)

	// To validate the event, we need to load its custom schema
	customSchema, err := os.ReadFile("myregistry-quotaexceeded_schema.json")
	examples.PanicOnError(err, "cannot load schema file")

	err = cdevents.LoadJsonSchema(customSchemaURI, customSchema)
	examples.PanicOnError(err, "cannot load the custom schema file")

	ce, err = cdevents.AsCloudEvent(event)
	examples.PanicOnError(err, "failed to create cloudevent")

	// Set send options
	source, err := examples.CreateSmeeChannel()
	examples.PanicOnError(err, "failed to create a smee channel")
	ctx := cloudevents.ContextWithTarget(context.Background(), *source)
	ctx = cloudevents.WithEncodingBinary(ctx)

	c, err = cloudevents.NewClientHTTP()
	examples.PanicOnError(err, "failed to create the CloudEvents client")

	// Send the CloudEvent
	// c is a CloudEvent client
	if result := c.Send(ctx, *ce); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	}
}
