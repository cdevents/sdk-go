# CDEvents Go SDK Docs

This folder contains example of how to use this SDK.

> **Note** For simplicity, the code below does not include error handling. The go files in example folders includes it.

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
	"os"

	examples "github.com/cdevents/sdk-go/docs/examples"
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
examples.PanicOnError(err, "could not create a cdevent")
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

// The event schema needs to be loaded, so the SDK may validate
// In this example, the schema is located in the same folder as
// the go code
customSchema, err := os.ReadFile("myregistry-quotaexceeded_schema.json")
examples.PanicOnError(err, "cannot load schema file")

err = cdevents.LoadJsonSchema(customSchemaUri, customSchema)
examples.PanicOnError(err, "cannot load the custom schema file")
```

To see the event, let's render it as JSON and log it:

```golang
// Render the event as JSON
eventJson, err := cdevents.AsJSONString(event)
examples.PanicOnError(err, "failed to marshal the CDEvent")
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

The whole code of is available under [`examples/custom`](./examples/custom/main.go):

```shell
➜ cd examples/custom
➜ go run main.go | jq .
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

## Consume a CDEvent with a Custom Schema

CDEvents producers may include a `schemaUri` in their events. The extra schema **must** comply with the CDEvents schema and may add additional rules on top.
The `schemaUri` field includes the `$id` field of the custom schema and can be used for different purposes:
* specify the format of the data included in the `customData` field
* specify the format of the subject content of custom events
* refine the format of one or more fields of a specific CDEvent

In this examples, the custom schema is used to define the format of the `customData` for a `change.created` events, which corresponds to the following golang `struct`:

```golang
type ChangeData struct {
	User     string `json:"user"`               // The user that created the PR
	Assignee string `json:"assignee,omitempty"` // The user assigned to the PR (optional)
	Head     string `json:"head"`               // The head commit (sha) of the PR
	Base     string `json:"base"`               // The base commit (sha) for the PR
}
```

The goal of this example is to consume (parse) an event with a custom schema and validate it. In the example we load the event from disk. In real life the event will be typically received over the network or extracted from a database.

For this scenario we will need a few imports:

```golang
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
```

Before parsing an event with a custom schema, it's required to load the schema into the SDK. This avoids having to download and compile the schema every time a message is parsed.

```golang
// Load and register the custom schema
customSchema, err := os.ReadFile("changecreated_schema.json")

// Unmarshal the schema to extract the $id. The $id can also be hardcoded as a const
eventAux := &struct {
    Id string `json:"$id"`
}{}
err = json.Unmarshal(customSchema, eventAux)
err = cdevents.LoadJsonSchema(eventAux.Id, customSchema)
```

Once the schema is loaded, it's possible to parse the event itself.
In this case we know that the event is in the v0.4 version format, so we use the corresponding API.

```golang
// Load, unmarshal and validate the event
eventBytes, err := os.ReadFile("changecreated.json")
event, err := cdevents04.NewFromJSONBytes(eventBytes)

err = cdevent.Validate(event)
if err != nil {
	log.Fatalf("cannot validate event %v: %v", event, err)
}

// Print the event
eventJson, err := cdevents.AsJSONString(event)
examples.PanicOnError(err, "failed to marshal the CDEvent")
fmt.Printf("%s\n\n", eventJson)
```
