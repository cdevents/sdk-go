// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

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

package v03_test

import (
	"github.com/cdevents/sdk-go/pkg/api"
	apiv03 "github.com/cdevents/sdk-go/pkg/api/v03"
)

func init() {
	tests = []testNewCDEventType{}
	tests = append(tests, testNewCDEventType{
		name:      "artifact packaged",
		eventType: apiv03.ArtifactPackagedEventType.String(),
		expectedEvent: &apiv03.ArtifactPackagedEvent{
			Context: api.Context{
				Type:      apiv03.ArtifactPackagedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ArtifactPackagedSubject{
				SubjectBase: api.SubjectBase{
					Type: "artifact",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "artifact published",
		eventType: apiv03.ArtifactPublishedEventType.String(),
		expectedEvent: &apiv03.ArtifactPublishedEvent{
			Context: api.Context{
				Type:      apiv03.ArtifactPublishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ArtifactPublishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "artifact",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "artifact signed",
		eventType: apiv03.ArtifactSignedEventType.String(),
		expectedEvent: &apiv03.ArtifactSignedEvent{
			Context: api.Context{
				Type:      apiv03.ArtifactSignedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ArtifactSignedSubject{
				SubjectBase: api.SubjectBase{
					Type: "artifact",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "branch created",
		eventType: apiv03.BranchCreatedEventType.String(),
		expectedEvent: &apiv03.BranchCreatedEvent{
			Context: api.Context{
				Type:      apiv03.BranchCreatedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.BranchCreatedSubject{
				SubjectBase: api.SubjectBase{
					Type: "branch",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "branch deleted",
		eventType: apiv03.BranchDeletedEventType.String(),
		expectedEvent: &apiv03.BranchDeletedEvent{
			Context: api.Context{
				Type:      apiv03.BranchDeletedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.BranchDeletedSubject{
				SubjectBase: api.SubjectBase{
					Type: "branch",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "build finished",
		eventType: apiv03.BuildFinishedEventType.String(),
		expectedEvent: &apiv03.BuildFinishedEvent{
			Context: api.Context{
				Type:      apiv03.BuildFinishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.BuildFinishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "build",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "build queued",
		eventType: apiv03.BuildQueuedEventType.String(),
		expectedEvent: &apiv03.BuildQueuedEvent{
			Context: api.Context{
				Type:      apiv03.BuildQueuedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.BuildQueuedSubject{
				SubjectBase: api.SubjectBase{
					Type: "build",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "build started",
		eventType: apiv03.BuildStartedEventType.String(),
		expectedEvent: &apiv03.BuildStartedEvent{
			Context: api.Context{
				Type:      apiv03.BuildStartedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.BuildStartedSubject{
				SubjectBase: api.SubjectBase{
					Type: "build",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "change abandoned",
		eventType: apiv03.ChangeAbandonedEventType.String(),
		expectedEvent: &apiv03.ChangeAbandonedEvent{
			Context: api.Context{
				Type:      apiv03.ChangeAbandonedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ChangeAbandonedSubject{
				SubjectBase: api.SubjectBase{
					Type: "change",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "change created",
		eventType: apiv03.ChangeCreatedEventType.String(),
		expectedEvent: &apiv03.ChangeCreatedEvent{
			Context: api.Context{
				Type:      apiv03.ChangeCreatedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ChangeCreatedSubject{
				SubjectBase: api.SubjectBase{
					Type: "change",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "change merged",
		eventType: apiv03.ChangeMergedEventType.String(),
		expectedEvent: &apiv03.ChangeMergedEvent{
			Context: api.Context{
				Type:      apiv03.ChangeMergedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ChangeMergedSubject{
				SubjectBase: api.SubjectBase{
					Type: "change",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "change reviewed",
		eventType: apiv03.ChangeReviewedEventType.String(),
		expectedEvent: &apiv03.ChangeReviewedEvent{
			Context: api.Context{
				Type:      apiv03.ChangeReviewedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ChangeReviewedSubject{
				SubjectBase: api.SubjectBase{
					Type: "change",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "change updated",
		eventType: apiv03.ChangeUpdatedEventType.String(),
		expectedEvent: &apiv03.ChangeUpdatedEvent{
			Context: api.Context{
				Type:      apiv03.ChangeUpdatedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ChangeUpdatedSubject{
				SubjectBase: api.SubjectBase{
					Type: "change",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "environment created",
		eventType: apiv03.EnvironmentCreatedEventType.String(),
		expectedEvent: &apiv03.EnvironmentCreatedEvent{
			Context: api.Context{
				Type:      apiv03.EnvironmentCreatedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.EnvironmentCreatedSubject{
				SubjectBase: api.SubjectBase{
					Type: "environment",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "environment deleted",
		eventType: apiv03.EnvironmentDeletedEventType.String(),
		expectedEvent: &apiv03.EnvironmentDeletedEvent{
			Context: api.Context{
				Type:      apiv03.EnvironmentDeletedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.EnvironmentDeletedSubject{
				SubjectBase: api.SubjectBase{
					Type: "environment",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "environment modified",
		eventType: apiv03.EnvironmentModifiedEventType.String(),
		expectedEvent: &apiv03.EnvironmentModifiedEvent{
			Context: api.Context{
				Type:      apiv03.EnvironmentModifiedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.EnvironmentModifiedSubject{
				SubjectBase: api.SubjectBase{
					Type: "environment",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "incident detected",
		eventType: apiv03.IncidentDetectedEventType.String(),
		expectedEvent: &apiv03.IncidentDetectedEvent{
			Context: api.Context{
				Type:      apiv03.IncidentDetectedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.IncidentDetectedSubject{
				SubjectBase: api.SubjectBase{
					Type: "incident",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "incident reported",
		eventType: apiv03.IncidentReportedEventType.String(),
		expectedEvent: &apiv03.IncidentReportedEvent{
			Context: api.Context{
				Type:      apiv03.IncidentReportedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.IncidentReportedSubject{
				SubjectBase: api.SubjectBase{
					Type: "incident",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "incident resolved",
		eventType: apiv03.IncidentResolvedEventType.String(),
		expectedEvent: &apiv03.IncidentResolvedEvent{
			Context: api.Context{
				Type:      apiv03.IncidentResolvedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.IncidentResolvedSubject{
				SubjectBase: api.SubjectBase{
					Type: "incident",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "pipelinerun finished",
		eventType: apiv03.PipelineRunFinishedEventType.String(),
		expectedEvent: &apiv03.PipelineRunFinishedEvent{
			Context: api.Context{
				Type:      apiv03.PipelineRunFinishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.PipelineRunFinishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "pipelineRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "pipelinerun queued",
		eventType: apiv03.PipelineRunQueuedEventType.String(),
		expectedEvent: &apiv03.PipelineRunQueuedEvent{
			Context: api.Context{
				Type:      apiv03.PipelineRunQueuedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.PipelineRunQueuedSubject{
				SubjectBase: api.SubjectBase{
					Type: "pipelineRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "pipelinerun started",
		eventType: apiv03.PipelineRunStartedEventType.String(),
		expectedEvent: &apiv03.PipelineRunStartedEvent{
			Context: api.Context{
				Type:      apiv03.PipelineRunStartedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.PipelineRunStartedSubject{
				SubjectBase: api.SubjectBase{
					Type: "pipelineRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "repository created",
		eventType: apiv03.RepositoryCreatedEventType.String(),
		expectedEvent: &apiv03.RepositoryCreatedEvent{
			Context: api.Context{
				Type:      apiv03.RepositoryCreatedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.RepositoryCreatedSubject{
				SubjectBase: api.SubjectBase{
					Type: "repository",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "repository deleted",
		eventType: apiv03.RepositoryDeletedEventType.String(),
		expectedEvent: &apiv03.RepositoryDeletedEvent{
			Context: api.Context{
				Type:      apiv03.RepositoryDeletedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.RepositoryDeletedSubject{
				SubjectBase: api.SubjectBase{
					Type: "repository",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "repository modified",
		eventType: apiv03.RepositoryModifiedEventType.String(),
		expectedEvent: &apiv03.RepositoryModifiedEvent{
			Context: api.Context{
				Type:      apiv03.RepositoryModifiedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.RepositoryModifiedSubject{
				SubjectBase: api.SubjectBase{
					Type: "repository",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "service deployed",
		eventType: apiv03.ServiceDeployedEventType.String(),
		expectedEvent: &apiv03.ServiceDeployedEvent{
			Context: api.Context{
				Type:      apiv03.ServiceDeployedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ServiceDeployedSubject{
				SubjectBase: api.SubjectBase{
					Type: "service",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "service published",
		eventType: apiv03.ServicePublishedEventType.String(),
		expectedEvent: &apiv03.ServicePublishedEvent{
			Context: api.Context{
				Type:      apiv03.ServicePublishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ServicePublishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "service",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "service removed",
		eventType: apiv03.ServiceRemovedEventType.String(),
		expectedEvent: &apiv03.ServiceRemovedEvent{
			Context: api.Context{
				Type:      apiv03.ServiceRemovedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ServiceRemovedSubject{
				SubjectBase: api.SubjectBase{
					Type: "service",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "service rolledback",
		eventType: apiv03.ServiceRolledbackEventType.String(),
		expectedEvent: &apiv03.ServiceRolledbackEvent{
			Context: api.Context{
				Type:      apiv03.ServiceRolledbackEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ServiceRolledbackSubject{
				SubjectBase: api.SubjectBase{
					Type: "service",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "service upgraded",
		eventType: apiv03.ServiceUpgradedEventType.String(),
		expectedEvent: &apiv03.ServiceUpgradedEvent{
			Context: api.Context{
				Type:      apiv03.ServiceUpgradedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.ServiceUpgradedSubject{
				SubjectBase: api.SubjectBase{
					Type: "service",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "taskrun finished",
		eventType: apiv03.TaskRunFinishedEventType.String(),
		expectedEvent: &apiv03.TaskRunFinishedEvent{
			Context: api.Context{
				Type:      apiv03.TaskRunFinishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TaskRunFinishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "taskRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "taskrun started",
		eventType: apiv03.TaskRunStartedEventType.String(),
		expectedEvent: &apiv03.TaskRunStartedEvent{
			Context: api.Context{
				Type:      apiv03.TaskRunStartedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TaskRunStartedSubject{
				SubjectBase: api.SubjectBase{
					Type: "taskRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testcaserun finished",
		eventType: apiv03.TestCaseRunFinishedEventType.String(),
		expectedEvent: &apiv03.TestCaseRunFinishedEvent{
			Context: api.Context{
				Type:      apiv03.TestCaseRunFinishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestCaseRunFinishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testCaseRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testcaserun queued",
		eventType: apiv03.TestCaseRunQueuedEventType.String(),
		expectedEvent: &apiv03.TestCaseRunQueuedEvent{
			Context: api.Context{
				Type:      apiv03.TestCaseRunQueuedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestCaseRunQueuedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testCaseRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testcaserun started",
		eventType: apiv03.TestCaseRunStartedEventType.String(),
		expectedEvent: &apiv03.TestCaseRunStartedEvent{
			Context: api.Context{
				Type:      apiv03.TestCaseRunStartedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestCaseRunStartedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testCaseRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testoutput published",
		eventType: apiv03.TestOutputPublishedEventType.String(),
		expectedEvent: &apiv03.TestOutputPublishedEvent{
			Context: api.Context{
				Type:      apiv03.TestOutputPublishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestOutputPublishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testOutput",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testsuiterun finished",
		eventType: apiv03.TestSuiteRunFinishedEventType.String(),
		expectedEvent: &apiv03.TestSuiteRunFinishedEvent{
			Context: api.Context{
				Type:      apiv03.TestSuiteRunFinishedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestSuiteRunFinishedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testSuiteRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testsuiterun queued",
		eventType: apiv03.TestSuiteRunQueuedEventType.String(),
		expectedEvent: &apiv03.TestSuiteRunQueuedEvent{
			Context: api.Context{
				Type:      apiv03.TestSuiteRunQueuedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestSuiteRunQueuedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testSuiteRun",
				},
			},
		},
	})
	tests = append(tests, testNewCDEventType{
		name:      "testsuiterun started",
		eventType: apiv03.TestSuiteRunStartedEventType.String(),
		expectedEvent: &apiv03.TestSuiteRunStartedEvent{
			Context: api.Context{
				Type:      apiv03.TestSuiteRunStartedEventType,
				Timestamp: timeNow(),
				Id:        testUUID(),
				Version:   "0.3.0",
			},
			Subject: apiv03.TestSuiteRunStartedSubject{
				SubjectBase: api.SubjectBase{
					Type: "testSuiteRun",
				},
			},
		},
	})
}
