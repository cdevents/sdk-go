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

	pipelineRunQueuedEventJson   string
	pipelineRunStartedEventJson  string
	pipelineRunFinishedEventJson string
	taskRunStartedEventJson      string
	taskRunFinishedEventJson     string
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

	newUUID, _ := uuidNewRandom()
	newTime := timeNow()
	pipelineRunQueuedEventJson = fmt.Sprintf(pipelineRunQueuedEventJsonTemplate, newUUID, newTime.Format(time.RFC3339Nano))
	pipelineRunStartedEventJson = fmt.Sprintf(pipelineRunStartedEventJsonTemplate, newUUID, newTime.Format(time.RFC3339Nano))
	pipelineRunFinishedEventJson = fmt.Sprintf(pipelineRunFinishedEventJsonTemplate, newUUID, newTime.Format(time.RFC3339Nano))
	taskRunStartedEventJson = fmt.Sprintf(taskRunStartedEventJsonTemplate, newUUID, newTime.Format(time.RFC3339Nano))
	taskRunFinishedEventJson = fmt.Sprintf(taskRunFinishedEventJsonTemplate, newUUID, newTime.Format(time.RFC3339Nano))
}

func TestAsCloudEvent(t *testing.T) {

	tests := []struct {
		name            string
		event           CDEvent
		payloadReceiver interface{}
		shouldFail      bool
	}{{
		name:            "pipelinerun queued",
		event:           pipelineRunQueuedEvent,
		payloadReceiver: &PipelineRunQueuedEvent{},
		shouldFail:      false,
	}, {
		name:            "pipelinerun started",
		event:           pipelineRunStartedEvent,
		payloadReceiver: &PipelineRunStartedEvent{},
		shouldFail:      false,
	}, {
		name:            "pipelinerun finished",
		event:           pipelineRunFinishedEvent,
		payloadReceiver: &PipelineRunFinishedEvent{},
		shouldFail:      false,
	}, {
		name:            "taskrun started",
		event:           taskRunStartedEvent,
		payloadReceiver: &TaskRunStartedEvent{},
		shouldFail:      false,
	}, {
		name:            "taskrun finished",
		event:           taskRunFinishedEvent,
		payloadReceiver: &TaskRunFinishedEvent{},
		shouldFail:      false,
	}, {
		name:       "invalid event",
		event:      nil,
		shouldFail: true,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ce, err := AsCloudEvent(tc.event)
			if err != nil && !tc.shouldFail {
				t.Fatalf("didn't expected it to fail, but it did")
			}
			if err == nil && tc.shouldFail {
				t.Fatalf("expected it to fail, but it didn't")
			}
			if tc.event != nil {
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
			}
		})
	}
}

func TestAsJsonString(t *testing.T) {

	tests := []struct {
		name       string
		event      CDEvent
		jsonString string
	}{{
		name:       "pipelinerun queued",
		event:      pipelineRunQueuedEvent,
		jsonString: pipelineRunQueuedEventJson,
	}, {
		name:       "pipelinerun started",
		event:      pipelineRunStartedEvent,
		jsonString: pipelineRunStartedEventJson,
	}, {
		name:       "pipelinerun finished",
		event:      pipelineRunFinishedEvent,
		jsonString: pipelineRunFinishedEventJson,
	}, {
		name:       "taskrun started",
		event:      taskRunStartedEvent,
		jsonString: taskRunStartedEventJson,
	}, {
		name:       "taskrun finished",
		event:      taskRunFinishedEvent,
		jsonString: taskRunFinishedEventJson,
	}, {
		name:       "nil event",
		event:      nil,
		jsonString: "",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obtainedJsonString, err := AsJsonString(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did")
			}
			if tc.jsonString != "" {
				expectedJsonString := &bytes.Buffer{}
				if err := json.Compact(expectedJsonString, []byte(tc.jsonString)); err != nil {
					t.Fatalf("somehow cannot compact test json %s", tc.jsonString)
				}
				if d := cmp.Diff(expectedJsonString.String(), obtainedJsonString); d != "" {
					t.Errorf("args: diff(-want,+got):\n%s", d)
				}
			} else {
				if d := cmp.Diff(tc.jsonString, obtainedJsonString); d != "" {
					t.Errorf("args: diff(-want,+got):\n%s", d)
				}
			}
		})
	}
}
