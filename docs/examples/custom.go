package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Quota struct {
	User      string `json:"user,omitempty"`      // The use the quota applies ot
	Limit     string `json:"limit,omitempty"`     // The limit enforced by the quota e.g. 100Gb
	Current   int    `json:"current,omitempty"`   // The current % of the quota used e.g. 90%
	Threshold int    `json:"threshold,omitempty"` // The threshold for warning event e.g. 85%
	Level     string `json:"level,omitempty"`     // INFO: <threshold, WARNING: >threshold, <quota, CRITICAL: >quota
}

// Copied from https://github.com/eswdd/go-smee/blob/33b0bac1f1ef3abef04c518ddf7552b04edbadd2/smee.go#L54C1-L67C2
func CreateSmeeChannel() (*string, error) {
	httpClient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := httpClient.Head("https://smee.io/new")
	if err != nil {
		return nil, err
	}

	loc := resp.Header.Get("Location")
	return &loc, nil
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
	if err != nil {
		log.Fatalf("could not create a cdevent, %v", err)
	}
	event.SetEventType(eventType)

	// Set the required context fields
	event.SetSubjectId("quotaRule123")
	event.SetSource("my/first/cdevent/program")

	// Set the required subject fields
	event.SetSubjectContent(quotaRule123)

	event.SetSchemaUri("https://myregistry.dev/schemas/cdevents/quota-exceeded/0_1_0")

	// Print the event
	eventJson, err := cdevents.AsJsonString(event)
	if err != nil {
		log.Fatalf("failed to marshal the CDEvent, %v", err)
	}
	fmt.Printf("%s", eventJson)

	ce, err = cdevents.AsCloudEvent(event)
	if err != nil {
		log.Fatalf("failed to create cloudevent, %v", err)
	}

	// Set send options
	source, err := CreateSmeeChannel()
	if err != nil {
		log.Fatalf("failed to create a smee channel: %v", err)
	}
	ctx := cloudevents.ContextWithTarget(context.Background(), *source)
	ctx = cloudevents.WithEncodingBinary(ctx)

	c, err = cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	// Send the CloudEvent
	// c is a CloudEvent client
	if result := c.Send(ctx, *ce); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	}
}
