package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	examples "github.com/cdevents/sdk-go/docs/examples"
	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cdevents04 "github.com/cdevents/sdk-go/pkg/api/v04"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type ChangeData struct {
	User     string `json:"user"`               // The user that created the PR
	Assignee string `json:"assignee,omitempty"` // The user assigned to the PR (optional)
	Head     string `json:"head"`               // The head commit (sha) of the PR
	Base     string `json:"base"`               // The base commit (sha) for the PR
}

func main() {
	var ce *cloudevents.Event
	var c cloudevents.Client

	// Load and register the custom schema
	customSchema, err := os.ReadFile("changecreated_schema.json")
	examples.PanicOnError(err, "cannot load schema file")

	// Unmarshal the schema to extract the $id. The $id can also be hardcoded as a const
	eventAux := &struct {
		Id string `json:"$id"`
	}{}
	err = json.Unmarshal(customSchema, eventAux)
	examples.PanicOnError(err, "cannot get $id from schema file")
	err = cdevents.LoadJsonSchema(eventAux.Id, customSchema)
	examples.PanicOnError(err, "cannot load the custom schema file")

	// Load and unmarshal the event
	eventBytes, err := os.ReadFile("changecreated.json")
	examples.PanicOnError(err, "cannot load event file")
	event, err := cdevents04.NewFromJsonBytes(eventBytes)
	examples.PanicOnError(err, "failed to unmarshal the CDEvent")

	// Print the event
	eventJson, err := cdevents.AsJsonString(event)
	examples.PanicOnError(err, "failed to marshal the CDEvent")
	fmt.Printf("%s\n\n", eventJson)

	ce, err = cdevents.AsCloudEvent(event)
	examples.PanicOnError(err, "failed to create cloudevent")

	// Set send options
	source, err := examples.CreateSmeeChannel()
	examples.PanicOnError(err, "failed to create a smee channel")

	ctx := cloudevents.ContextWithTarget(context.Background(), *source)
	ctx = cloudevents.WithEncodingBinary(ctx)

	c, err = cloudevents.NewClientHTTP()
	examples.PanicOnError(err, "failed to create a CloudEvents client")

	// Send the CloudEvent
	// c is a CloudEvent client
	if result := c.Send(ctx, *ce); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	}
}
