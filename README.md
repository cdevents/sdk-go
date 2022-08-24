# Golang CDEvents SDK

Golang SDK to emit [CDEvents](https://cdevents/dev).

The SDK can be used to create CDEvents and send them as CloudEvents, as well as parse a received CloudEvent into a CDEvent.

## Disclaimer 🚧

This SDK is work in work in progress, it will be maintained in sync with the
specification but it only covers a limited number of events.
For a wider range of events, the [old SDK/CLI][old-sdk] may be used, with the
caveat that the old SDK is not aligned with the new version of the specification.

## Get started

Add the module as dependency using go mod:

```golang
go get github.com/cdevents/sdk-go
```

And import the module in your code

```golang
import cdevents "github.com/cdevents/sdk-go"
```

## Create your first CDEvent

To create a CDEvent, for instance a [*pipelineRun queued*](https://cdevents.dev/docs/core/#pipelinerun-queued) one:

```golang
func main() {

    // Create the base event
    event, err := cdevents.NewPipelineRunQueuedEvent()
    if err != nil {
      log.Fatalf("could not create a cdevent, %v", err)
    }

    // Set the required context fields
    event.SetSubjectId("myPipelineRun1")
    event.SetSource("my/first/cdevent/program")

    // Set the required subject fields
    event.SetSubjectPipelineName("myPipeline")
    event.SetSubjectURL("https://example.com/myPipeline")
}
```

## Send your first CDEvent as CloudEvent

To send a CDEvent as CloudEvent:

```golang
func main() {
    // (...) set the event first
    ce := cdevents.AsCloudEvent(event)

    // Set send options
    ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")
    ctx = cloudevents.WithEncodingBinary(ctx)

    // Sent the CloudEvent
    if result := c.Send(ctx, ce); cloudevents.IsUndelivered(result) {
        log.Fatalf("failed to send, %v", result)
    }
}
```

See the [CloudEvents](https://github.com/cloudevents/sdk-go#send-your-first-cloudevent) docs as well.

## Contributing

If you would like to contribute, see our [development](DEVELOPMENT.md) guide.

## References

- [CDEvents](https://cdevents.dev)
- [CDFoundation SIG Events Vocabulary Draft](https://github.com/cdfoundation/sig-events/tree/main/vocabulary-draft)

[old-sdk]: https://github.com/cdfoundation/sig-events/tree/main/cde/sdk/go
