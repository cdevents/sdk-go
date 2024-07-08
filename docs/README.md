# CDEvents Go SDK Docs

This folder contains example of how to use this SDK.

## Create a Custom CDEvent

If a tool wants to emit events that are not supported by the CDEvents specification,
they can do so via [custom events](https://github.com/cdevents/spec/tree/main/custom).

Custom events are follow the CDEvents format and can be defined via the
`CustomTypeEvent` object, available since v0.4, as well as using the `CustomCDEventReader`
and `CustomCDEventWriter` interfaces.

Let's consider the following scenario: a tool called "MyRegistry" has a concept of "Quota"
which can be "exceeded" by users of the system. We want to use events to notify when that
happens, but CDEvents does not define any quota related subject.

```golang
type Quota struct {
	User string `json:"user,omitempty"` // The use the quota applies ot
	Limit string `json:"limit,omitempty"` // The limit enforced by the quota e.g. 100Gb
	Current int `json:"current,omitempty"` // The current % of the quota used e.g. 90%
	Threshold int `json:"threshold,omitempty"` // The threshold for warning event e.g. 85%
	Level string `json:"level,omitempty"` // INFO: <threshold, WARNING: >threshold, <quota, CRITICAL: >quota
}
```
For this scenario we will need a few imports:

```golang
import (
	"context"
    "fmt"
	"log"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)
```

Let's define a custom event type for this scenario.
This is our first iteration, so the event will have version "0.1.0".

```golang
eventType := cdevents.CDEventType{
    Subject:   "quota",
    Predicate: "exceeded",
    Version:   "0.1.0",
    Custom:    "myregistry",
}
```

With a `Quota` object, let's create a CDEvent for it:

```golang
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
event.SetSource("myregistry/region/staging")

// Set the required subject content
event.SetSubjectContent(quotaRule123)

// If we host a schema for the overall custom CDEvent, we can add it
// to the event so that the receiver may validate custom fields like
// the event type and subject content
event.SetSchemaUri("https://myregistry.dev/schemas/cdevents/quota-exceeded/0_1_0")
```

To see the event, let's render it as JSON and log it:

```golang
// Render the event as JSON
eventJson, err := cdevents.AsJsonString(event)
if err != nil {
    log.Fatalf("failed to marshal the CDEvent, %v", err)
}
// Print the event
fmt.Printf("%s", eventJson)
```

The resulting CDEvents will look like:

```json
{"context":{"version":"0.4.1","id":"37fc85d9-187f-4ceb-a11d-9df30f809624","source":"my/first/cdevent/program","type":"dev.cdeventsx.myregistry-quota.exceeded.0.1.0","timestamp":"2024-07-09T14:00:54.375172+01:00","schemaUri":"https://myregistry.dev/schemas/cdevents/quota-exceeded/0_1_0"},"subject":{"id":"quotaRule123","source":"my/first/cdevent/program","type":"myregistry-quota","content":{"user":"heavy_user","limit":"50Tb","current":90,"threshold":85,"level":"WARNING"}}}
```

To send the event, let's setup a test sink, for instance using [smee.io/](https://smee.io/).
Then let's render the event as CloudEvent and send it to the sink:

```golang
ce, err = cdevents.AsCloudEvent(event)
if err != nil {
    log.Fatalf("failed to create cloudevent, %v", err)
}

// Set send options
ctx := cloudevents.ContextWithTarget(context.Background(), "https://smee.io/<you-channel-id>")
ctx = cloudevents.WithEncodingBinary(ctx)

c, err = cloudevents.NewClientHTTP()
if err != nil {
    log.Fatalf("failed to create client, %v", err)
}

// Send the CloudEvent
if result := c.Send(ctx, *ce); cloudevents.IsUndelivered(result) {
    log.Fatalf("failed to send, %v", result)
}
```

The whole code of is available under [`examples/custom.go`](./examples/custom.go):

```shell
➜ go run custom.go | jq .
{
  "context": {
    "version": "0.4.1",
    "id": "f7be8a13-8bd7-4a3a-881f-ed49cc0ebf8f",
    "source": "my/first/cdevent/program",
    "type": "dev.cdeventsx.myregistry-quota.exceeded.0.1.0",
    "timestamp": "2024-07-09T14:01:00.449264+01:00",
    "schemaUri": "https://myregistry.dev/schemas/cdevents/quota-exceeded/0_1_0"
  },
  "subject": {
    "id": "quotaRule123",
    "source": "my/first/cdevent/program",
    "type": "myregistry-quota",
    "content": {
      "user": "heavy_user",
      "limit": "50Tb",
      "current": 90,
      "threshold": 85,
      "level": "WARNING"
    }
  }
}
```