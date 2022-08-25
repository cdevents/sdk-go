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
	}, {
		name:      "change created",
		eventType: ChangeCreatedEventV1,
		expectedEvent: &ChangeCreatedEvent{
			Context: Context{
				Type:      ChangeCreatedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: ChangeCreatedSubject{
				SubjectBase: SubjectBase{
					Type: ChangeSubjectType,
				},
			},
		},
	}, {
		name:      "change updated",
		eventType: ChangeUpdatedEventV1,
		expectedEvent: &ChangeUpdatedEvent{
			Context: Context{
				Type:      ChangeUpdatedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: ChangeUpdatedSubject{
				SubjectBase: SubjectBase{
					Type: ChangeSubjectType,
				},
			},
		},
	}, {
		name:      "change reviewed",
		eventType: ChangeReviewedEventV1,
		expectedEvent: &ChangeReviewedEvent{
			Context: Context{
				Type:      ChangeReviewedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: ChangeReviewedSubject{
				SubjectBase: SubjectBase{
					Type: ChangeSubjectType,
				},
			},
		},
	}, {
		name:      "change merged",
		eventType: ChangeMergedEventV1,
		expectedEvent: &ChangeMergedEvent{
			Context: Context{
				Type:      ChangeMergedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: ChangeMergedSubject{
				SubjectBase: SubjectBase{
					Type: ChangeSubjectType,
				},
			},
		},
	}, {
		name:      "change abandoned",
		eventType: ChangeAbandonedEventV1,
		expectedEvent: &ChangeAbandonedEvent{
			Context: Context{
				Type:      ChangeAbandonedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: ChangeAbandonedSubject{
				SubjectBase: SubjectBase{
					Type: ChangeSubjectType,
				},
			},
		},
	}, {
		name:      "repository created",
		eventType: RepositoryCreatedEventV1,
		expectedEvent: &RepositoryCreatedEvent{
			Context: Context{
				Type:      RepositoryCreatedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: RepositoryCreatedSubject{
				SubjectBase: SubjectBase{
					Type: RepositorySubjectType,
				},
			},
		},
	}, {
		name:      "repository modified",
		eventType: RepositoryModifiedEventV1,
		expectedEvent: &RepositoryModifiedEvent{
			Context: Context{
				Type:      RepositoryModifiedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: RepositoryModifiedSubject{
				SubjectBase: SubjectBase{
					Type: RepositorySubjectType,
				},
			},
		},
	}, {
		name:      "repository deleted",
		eventType: RepositoryDeletedEventV1,
		expectedEvent: &RepositoryDeletedEvent{
			Context: Context{
				Type:      RepositoryDeletedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: RepositoryDeletedSubject{
				SubjectBase: SubjectBase{
					Type: RepositorySubjectType,
				},
			},
		},
	}, {
		name:      "branch created",
		eventType: BranchCreatedEventV1,
		expectedEvent: &BranchCreatedEvent{
			Context: Context{
				Type:      BranchCreatedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: BranchCreatedSubject{
				SubjectBase: SubjectBase{
					Type: BranchSubjectType,
				},
			},
		},
	}, {
		name:      "branch deleted",
		eventType: BranchDeletedEventV1,
		expectedEvent: &BranchDeletedEvent{
			Context: Context{
				Type:      BranchDeletedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: BranchDeletedSubject{
				SubjectBase: SubjectBase{
					Type: BranchSubjectType,
				},
			},
		},
	}, {
		name:      "testcase queued",
		eventType: TestCaseQueuedEventV1,
		expectedEvent: &TestCaseQueuedEvent{
			Context: Context{
				Type:      TestCaseQueuedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TestCaseQueuedSubject{
				SubjectBase: SubjectBase{
					Type: TestCaseSubjectType,
				},
			},
		},
	}, {
		name:      "testcase started",
		eventType: TestCaseStartedEventV1,
		expectedEvent: &TestCaseStartedEvent{
			Context: Context{
				Type:      TestCaseStartedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TestCaseStartedSubject{
				SubjectBase: SubjectBase{
					Type: TestCaseSubjectType,
				},
			},
		},
	}, {
		name:      "testcase finished",
		eventType: TestCaseFinishedEventV1,
		expectedEvent: &TestCaseFinishedEvent{
			Context: Context{
				Type:      TestCaseFinishedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TestCaseFinishedSubject{
				SubjectBase: SubjectBase{
					Type: TestCaseSubjectType,
				},
			},
		},
	}, {
		name:      "testsuite started",
		eventType: TestSuiteStartedEventV1,
		expectedEvent: &TestSuiteStartedEvent{
			Context: Context{
				Type:      TestSuiteStartedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TestSuiteStartedSubject{
				SubjectBase: SubjectBase{
					Type: TestSuiteSubjectType,
				},
			},
		},
	}, {
		name:      "testsuite finished",
		eventType: TestSuiteFinishedEventV1,
		expectedEvent: &TestSuiteFinishedEvent{
			Context: Context{
				Type:      TestSuiteFinishedEventV1,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   CDEventsSpecVersion,
			},
			Subject: TestSuiteFinishedSubject{
				SubjectBase: SubjectBase{
					Type: TestSuiteSubjectType,
				},
			},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			event, err := NewCDEvent(tc.eventType)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(tc.expectedEvent, event); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestNewCDEventFailed(t *testing.T) {

	_, err := NewCDEvent("not supported")
	if err == nil {
		t.Fatalf("expected it to fail, but it didn't")
	}
}
