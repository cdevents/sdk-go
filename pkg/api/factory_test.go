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
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func init() {

	// Get the time once
	t := time.Now()
	timeNow = func() time.Time {
		return t
	}

	// Get the UUID once
	u, _ := uuid.NewRandom()
	uuidNewRandom = func() (uuid.UUID, error) {
		return u, nil
	}
}

func testUUID() string {
	u, _ := uuidNewRandom()
	return fmt.Sprintf("%v", u)
}

func TestNewCDEvent(t *testing.T) {

	tests := []struct {
		name          string
		eventType     CDEventType
		expectedEvent CDEvent
		shouldFail    bool
	}{{
		name:      "pipelinerun queued",
		eventType: PipelineRunQueuedEventV1,
		expectedEvent: &PipelineRunQueuedEvent{
			Context: Context{
				Type:      PipelineRunQueuedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: PipelineRunQueuedSubject{
				SubjectBase: SubjectBase{
					Type: PipelineRunSubjectType,
				},
			},
		},
		shouldFail: false,
	}, {
		name:      "pipelinerun started",
		eventType: PipelineRunStartedEventV1,
		expectedEvent: &PipelineRunStartedEvent{
			Context: Context{
				Type:      PipelineRunStartedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: PipelineRunStartedSubject{
				SubjectBase: SubjectBase{
					Type: PipelineRunSubjectType,
				},
			},
		},
		shouldFail: false,
	}, {
		name:      "pipelinerun finished",
		eventType: PipelineRunFinishedEventV1,
		expectedEvent: &PipelineRunFinishedEvent{
			Context: Context{
				Type:      PipelineRunFinishedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: PipelineRunFinishedSubject{
				SubjectBase: SubjectBase{
					Type: PipelineRunSubjectType,
				},
			},
		},
		shouldFail: false,
	}, {
		name:      "taskrun started",
		eventType: TaskRunStartedEventV1,
		expectedEvent: &TaskRunStartedEvent{
			Context: Context{
				Type:      TaskRunStartedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TaskRunStartedSubject{
				SubjectBase: SubjectBase{
					Type: TaskRunSubjectType,
				},
			},
		},
		shouldFail: false,
	}, {
		name:      "taskrun finished",
		eventType: TaskRunFinishedEventV1,
		expectedEvent: &TaskRunFinishedEvent{
			Context: Context{
				Type:      TaskRunFinishedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TaskRunFinishedSubject{
				SubjectBase: SubjectBase{
					Type: TaskRunSubjectType,
				},
			},
		},
		shouldFail: false,
	}, {
		name:          "not supported",
		eventType:     "not supported",
		expectedEvent: nil,
		shouldFail:    true,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			event, err := NewCDEvent(tc.eventType)
			if err != nil && !tc.shouldFail {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if err == nil && tc.shouldFail {
				t.Fatalf("expected it to fail, but it didn't")
			}
			if d := cmp.Diff(tc.expectedEvent, event); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}
