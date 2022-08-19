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
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

var (
	testSource          = "TestAsCloudEvent"
	testSubjectId       = "mySubject123"
	testPipeline        = "myPipeline"
	testSubjectURL      = "https://www.example.com/mySubject123"
	testPipelineOutcome = PipelineRunOutcomeFailed
	testPipelineErrors  = "Something went wrong\nWith some more details"
	testTaskName        = "myTask"
	testTaskOutcome     = TaskRunOutcomeFailed
	testTaskRunErrors   = "Something went wrong\nWith some more details"

	pipelineRunQueuedEvent   *PipelineRunQueuedEvent
	pipelineRunStartedEvent  *PipelineRunStartedEvent
	pipelineRunFinishedEvent *PipelineRunFinishedEvent
	taskRunStartedEvent      *TaskRunStartedEvent
	taskRunFinishedEvent     *TaskRunFinishedEvent
	changeCreatedEvent       *ChangeCreatedEvent
	changeUpdatedEvent       *ChangeUpdatedEvent
	changeReviewedEvent      *ChangeReviewedEvent
	changeMergedEvent        *ChangeMergedEvent
	changeAbandonedEvent     *ChangeAbandonedEvent

	pipelineRunQueuedEventJsonTemplate = `
{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.pipelinerun.queued.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "pipelineRun",
		"content": {
			"pipelineName": "myPipeline",
			"url": "https://www.example.com/mySubject123"
		}
	}
}`

	pipelineRunStartedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.pipelinerun.started.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "pipelineRun",
		"content": {
			"pipelineName": "myPipeline",
			"url": "https://www.example.com/mySubject123"
		}
	}
}`

	pipelineRunFinishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.pipelinerun.finished.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "pipelineRun",
		"content": {
			"pipelineName": "myPipeline",
			"url": "https://www.example.com/mySubject123",
			"outcome": "failure",
			"errors": "Something went wrong\nWith some more details"
		}
	}
}`

	taskRunStartedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.taskrun.started.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "taskRun",
		"content": {
			"taskName": "myTask",
			"url": "https://www.example.com/mySubject123",
			"pipelineRun": {
				"id": "mySubject123"
			}
		}
	}
}`

	taskRunFinishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.taskrun.finished.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "taskRun",
		"content": {
			"taskName": "myTask",
			"url": "https://www.example.com/mySubject123",
			"pipelineRun": {
				"id": "mySubject123"
			},
			"outcome": "failure",
			"errors": "Something went wrong\nWith some more details"
		}
	}
}`

	changeCreatedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.created.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {}
	}
}`

	changeUpdatedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.updated.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {}
	}
}`

	changeReviewedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.reviewed.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {}
	}
}`

	changeMergedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.merged.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {}
	}
}`

	changeAbandonedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.abandoned.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {}
	}
}`

	pipelineRunQueuedEventJson   string
	pipelineRunStartedEventJson  string
	pipelineRunFinishedEventJson string
	taskRunStartedEventJson      string
	taskRunFinishedEventJson     string
	changeCreateEventJson        string
	changeUpdatedEventJson       string
	changeReviewedEventJson      string
	changeMergedEventJson        string
	changeAbandonedEventJson     string
)

func init() {

	// Get the time once
	t := time.Now().Round(0)
	timeNow = func() time.Time {
		return t
	}

	// Get the UUID once
	u, _ := uuid.NewRandom()
	uuidNewRandom = func() (uuid.UUID, error) {
		return u, nil
	}
}

func makeCDEvent(eventType CDEventType) CDEvent {
	event, _ := NewCDEvent(eventType)
	event.SetSource(testSource)
	event.SetSubjectId(testSubjectId)
	return event
}

func init() {
	e := makeCDEvent(PipelineRunQueuedEventV1)
	pipelineRunQueuedEvent, _ = e.(*PipelineRunQueuedEvent)
	pipelineRunQueuedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunQueuedEvent.SetSubjectURL(testSubjectURL)

	e = makeCDEvent(PipelineRunStartedEventV1)
	pipelineRunStartedEvent, _ = e.(*PipelineRunStartedEvent)
	pipelineRunStartedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunStartedEvent.SetSubjectURL(testSubjectURL)

	e = makeCDEvent(PipelineRunFinishedEventV1)
	pipelineRunFinishedEvent, _ = e.(*PipelineRunFinishedEvent)
	pipelineRunFinishedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunFinishedEvent.SetSubjectURL(testSubjectURL)
	pipelineRunFinishedEvent.SetSubjectOutcome(testPipelineOutcome)
	pipelineRunFinishedEvent.SetSubjectErrors(testPipelineErrors)

	e = makeCDEvent(TaskRunStartedEventV1)
	taskRunStartedEvent, _ = e.(*TaskRunStartedEvent)
	taskRunStartedEvent.SetSubjectTaskName(testTaskName)
	taskRunStartedEvent.SetSubjectURL(testSubjectURL)
	taskRunStartedEvent.SetSubjectPipelineRun(Reference{Id: testSubjectId})

	e = makeCDEvent(TaskRunFinishedEventV1)
	taskRunFinishedEvent, _ = e.(*TaskRunFinishedEvent)
	taskRunFinishedEvent.SetSubjectTaskName(testTaskName)
	taskRunFinishedEvent.SetSubjectURL(testSubjectURL)
	taskRunFinishedEvent.SetSubjectPipelineRun(Reference{Id: testSubjectId})
	taskRunFinishedEvent.SetSubjectOutcome(testTaskOutcome)
	taskRunFinishedEvent.SetSubjectErrors(testTaskRunErrors)

	e = makeCDEvent(ChangeCreatedEventV1)
	changeCreatedEvent, _ = e.(*ChangeCreatedEvent)

	e = makeCDEvent(ChangeUpdatedEventV1)
	changeUpdatedEvent, _ = e.(*ChangeUpdatedEvent)

	e = makeCDEvent(ChangeReviewedEventV1)
	changeReviewedEvent, _ = e.(*ChangeReviewedEvent)

	e = makeCDEvent(ChangeMergedEventV1)
	changeMergedEvent, _ = e.(*ChangeMergedEvent)

	e = makeCDEvent(ChangeAbandonedEventV1)
	changeAbandonedEvent, _ = e.(*ChangeAbandonedEvent)

	newUUID, _ := uuidNewRandom()
	newTime := timeNow().Format(time.RFC3339Nano)
	pipelineRunQueuedEventJson = fmt.Sprintf(pipelineRunQueuedEventJsonTemplate, newUUID, newTime)
	pipelineRunStartedEventJson = fmt.Sprintf(pipelineRunStartedEventJsonTemplate, newUUID, newTime)
	pipelineRunFinishedEventJson = fmt.Sprintf(pipelineRunFinishedEventJsonTemplate, newUUID, newTime)
	taskRunStartedEventJson = fmt.Sprintf(taskRunStartedEventJsonTemplate, newUUID, newTime)
	taskRunFinishedEventJson = fmt.Sprintf(taskRunFinishedEventJsonTemplate, newUUID, newTime)
	changeCreateEventJson = fmt.Sprintf(changeCreatedEventJsonTemplate, newUUID, newTime)
	changeUpdatedEventJson = fmt.Sprintf(changeUpdatedEventJsonTemplate, newUUID, newTime)
	changeReviewedEventJson = fmt.Sprintf(changeReviewedEventJsonTemplate, newUUID, newTime)
	changeMergedEventJson = fmt.Sprintf(changeMergedEventJsonTemplate, newUUID, newTime)
	changeAbandonedEventJson = fmt.Sprintf(changeAbandonedEventJsonTemplate, newUUID, newTime)
}

func TestAsCloudEvent(t *testing.T) {

	tests := []struct {
		name            string
		event           CDEvent
		payloadReceiver interface{}
	}{{
		name:            "pipelinerun queued",
		event:           pipelineRunQueuedEvent,
		payloadReceiver: &PipelineRunQueuedEvent{},
	}, {
		name:            "pipelinerun started",
		event:           pipelineRunStartedEvent,
		payloadReceiver: &PipelineRunStartedEvent{},
	}, {
		name:            "pipelinerun finished",
		event:           pipelineRunFinishedEvent,
		payloadReceiver: &PipelineRunFinishedEvent{},
	}, {
		name:            "taskrun started",
		event:           taskRunStartedEvent,
		payloadReceiver: &TaskRunStartedEvent{},
	}, {
		name:            "taskrun finished",
		event:           taskRunFinishedEvent,
		payloadReceiver: &TaskRunFinishedEvent{},
	}, {
		name:            "change created",
		event:           changeCreatedEvent,
		payloadReceiver: &ChangeCreatedEvent{},
	}, {
		name:            "change updated",
		event:           changeUpdatedEvent,
		payloadReceiver: &ChangeUpdatedEvent{},
	}, {
		name:            "change reviewed",
		event:           changeReviewedEvent,
		payloadReceiver: &ChangeReviewedEvent{},
	}, {
		name:            "change merged",
		event:           changeMergedEvent,
		payloadReceiver: &ChangeMergedEvent{},
	}, {
		name:            "change abandoned",
		event:           changeAbandonedEvent,
		payloadReceiver: &ChangeAbandonedEvent{},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ce, err := AsCloudEvent(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(testSubjectId, ce.Context.GetSubject()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(testSource, ce.Context.GetSource()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetType().String(), ce.Context.GetType()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			err = ce.DataAs(tc.payloadReceiver)
			if err != nil {
				t.Fatalf("somehow cannot unmarshal test event %v", ce)
			}
			if d := cmp.Diff(tc.event, tc.payloadReceiver); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestAsCloudEventInvalid(t *testing.T) {
	_, err := AsCloudEvent(nil)
	if err == nil {
		t.Fatalf("expected it to fail, but it didn't")
	}
}

func TestAsJsonString(t *testing.T) {

	compiler := jsonschema.NewCompiler()

	tests := []struct {
		name       string
		event      CDEvent
		jsonString string
		schemaName string
	}{{
		name:       "pipelinerun queued",
		event:      pipelineRunQueuedEvent,
		jsonString: pipelineRunQueuedEventJson,
		schemaName: "pipelinerunqueued",
	}, {
		name:       "pipelinerun started",
		event:      pipelineRunStartedEvent,
		jsonString: pipelineRunStartedEventJson,
		schemaName: "pipelinerunstarted",
	}, {
		name:       "pipelinerun finished",
		event:      pipelineRunFinishedEvent,
		jsonString: pipelineRunFinishedEventJson,
		schemaName: "pipelinerunfinished",
	}, {
		name:       "taskrun started",
		event:      taskRunStartedEvent,
		jsonString: taskRunStartedEventJson,
		schemaName: "taskrunstarted",
	}, {
		name:       "taskrun finished",
		event:      taskRunFinishedEvent,
		jsonString: taskRunFinishedEventJson,
		schemaName: "taskrunfinished",
	}, {
		name:       "change created",
		event:      changeCreatedEvent,
		jsonString: changeCreateEventJson,
		schemaName: "changecreated",
	}, {
		name:       "change updated",
		event:      changeUpdatedEvent,
		jsonString: changeUpdatedEventJson,
		schemaName: "changeupdated",
	}, {
		name:       "change reviewed",
		event:      changeReviewedEvent,
		jsonString: changeReviewedEventJson,
		schemaName: "changereviewed",
	}, {
		name:       "change merged",
		event:      changeMergedEvent,
		jsonString: changeMergedEventJson,
		schemaName: "changemerged",
	}, {
		name:       "change abandoned",
		event:      changeAbandonedEvent,
		jsonString: changeAbandonedEventJson,
		schemaName: "changeabandoned",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// First validate that the test JSON compiles against the schema
			sch, err := compiler.Compile(fmt.Sprintf("../../jsonschema/%s.json", tc.schemaName))
			if err != nil {
				t.Fatalf("Cannot compile jsonschema %s", tc.schemaName)
			}
			var v interface{}
			if err := json.Unmarshal([]byte(tc.jsonString), &v); err != nil {
				t.Fatalf("Cannot unmarshal test json: %v", err)
			}
			err = sch.Validate(v)
			if err != nil {
				t.Fatalf("Failed to validate events %s", err)
			}
			// Then test that AsJsonString produces a matching JSON from the event
			obtainedJsonString, err := AsJsonString(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			expectedJsonString := &bytes.Buffer{}
			if err := json.Compact(expectedJsonString, []byte(tc.jsonString)); err != nil {
				t.Fatalf("somehow cannot compact test json %s", tc.jsonString)
			}
			if d := cmp.Diff(expectedJsonString.String(), obtainedJsonString); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestInvalidEvent(t *testing.T) {

	eventNoSource, _ := NewCDEvent(ChangeAbandonedEventV1)
	eventNoSource.SetSubjectId(testSubjectId)

	eventNoSubjectId, _ := NewCDEvent(ChangeAbandonedEventV1)
	eventNoSubjectId.SetSource(testSource)

	eventSubjectIncomplete := makeCDEvent(PipelineRunQueuedEventV1)
	eventSubjectIncomplete, _ = eventSubjectIncomplete.(*PipelineRunQueuedEvent)

	tests := []struct {
		name  string
		event CDEvent
	}{{
		name:  "missing source",
		event: eventNoSource,
	}, {
		name:  "missing subject id",
		event: eventNoSubjectId,
	}, {
		name:  "missing mandatory subject details",
		event: eventSubjectIncomplete,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// First validate that the test JSON compiles against the schema
			err := Validate(tc.event)
			if err == nil {
				t.Fatalf("Expected validation to fail, but it succeeded instead")
			}
		})
	}
}

func TestAsJsonStringEmpty(t *testing.T) {
	obtainedJsonString, err := AsJsonString(nil)
	if err != nil {
		t.Fatalf("didn't expected it to fail, but it did: %v", err)
	}
	if d := cmp.Diff("", obtainedJsonString); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}
