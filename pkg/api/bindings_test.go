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

type testData struct {
	TestValues []map[string]string `json:"testValues"`
}

var (
	testSource          = "TestAsCloudEvent"
	testSubjectId       = "mySubject123"
	testPipeline        = "myPipeline"
	testSubjecturl      = "https://www.example.com/mySubject123"
	testPipelineOutcome = PipelineRunOutcomeFailed
	testPipelineErrors  = "Something went wrong\nWith some more details"
	testTaskName        = "myTask"
	testTaskOutcome     = TaskRunOutcomeFailed
	testTaskRunErrors   = "Something went wrong\nWith some more details"
	testRepo            = "TestRepo"
	testOwner           = "TestOrg"
	testUrl             = "https://example.org/TestOrg/TestRepo"
	testViewUrl         = "https://example.org/view/TestOrg/TestRepo"
	testArtifactId      = "0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
	testEnvironmentId   = "test123"
	testEnvironmentName = "testEnv"
	testEnvironmentUrl  = "https://example.org/testEnv"
	testDataJson        = testData{TestValues: []map[string]string{{"k1": "v1"}, {"k2": "v2"}}}
	testDataXml         = []byte("<xml>testData</xml>")

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
	repositoryCreatedEvent   *RepositoryCreatedEvent
	repositoryModifiedEvent  *RepositoryModifiedEvent
	repositoryDeletedEvent   *RepositoryDeletedEvent
	branchCreatedEvent       *BranchCreatedEvent
	branchDeletedEvent       *BranchDeletedEvent
	testCaseQueuedEvent      *TestCaseQueuedEvent
	testCaseStartedEvent     *TestCaseStartedEvent
	testCaseFinishedEvent    *TestCaseFinishedEvent
	testSuiteStartedEvent    *TestSuiteStartedEvent
	testSuiteFinishedEvent   *TestSuiteFinishedEvent
	buildQueuedEvent         *BuildQueuedEvent
	buildStartedEvent        *BuildStartedEvent
	buildFinishedEvent       *BuildFinishedEvent
	artifactPackagedEvent    *ArtifactPackagedEvent
	artifactPublishedEvent   *ArtifactPublishedEvent
	environmentCreatedEvent  *EnvironmentCreatedEvent
	environmentModifiedEvent *EnvironmentModifiedEvent
	environmentDeletedEvent  *EnvironmentDeletedEvent
	serviceDeployedEvent     *ServiceDeployedEvent
	serviceUpgradedEvent     *ServiceUpgradedEvent
	serviceRolledBackEvent   *ServiceRolledbackEvent
	serviceRemovedEvent      *ServiceRemovedEvent
	servicePublishedEvent    *ServicePublishedEvent

	eventJsonCustomData    *ArtifactPackagedEvent
	eventNonJsonCustomData *ArtifactPackagedEvent

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

	repositoryCreatedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.repository.created.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "repository",
		"content": {
			"name": "TestRepo",
			"owner": "TestOrg",
			"url": "https://example.org/TestOrg/TestRepo",
			"viewUrl": "https://example.org/view/TestOrg/TestRepo"
		}
	}
}`

	repositoryModifiedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.repository.modified.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "repository",
		"content": {
			"name": "TestRepo",
			"owner": "TestOrg",
			"url": "https://example.org/TestOrg/TestRepo",
			"viewUrl": "https://example.org/view/TestOrg/TestRepo"
		}
	}
}`

	repositoryDeletedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.repository.deleted.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "repository",
		"content": {
			"name": "TestRepo",
			"owner": "TestOrg",
			"url": "https://example.org/TestOrg/TestRepo",
			"viewUrl": "https://example.org/view/TestOrg/TestRepo"
		}
	}
}`

	branchCreatedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.branch.created.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "branch",
		"content": {}
	}
}`

	branchDeletedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.branch.deleted.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "branch",
		"content": {}
	}
}`

	testCaseQueuedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testcase.queued.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testCase",
		"content": {}
	}
}`

	testCaseStartedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testcase.started.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testCase",
		"content": {}
	}
}`

	testCaseFinishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testcase.finished.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testCase",
		"content": {}
	}
}`

	testSuiteStartedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testsuite.started.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testSuite",
		"content": {}
	}
}`

	testSuiteFinishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testsuite.finished.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testSuite",
		"content": {}
	}
}`

	buildQueuedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.build.queued.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "build",
		"content": {}
	}
}`

	buildStartedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.build.started.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "build",
		"content": {}
	}
}`

	buildFinishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.build.finished.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "build",
		"content": {
			"artifactId": "0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
		}
	}
}`

	artifactPackagedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {}
	}
}`

	artifactPublishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.published.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {}
	}
}`

	environmentCreatedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.environment.created.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "environment",
		"content": {
			"name": "testEnv",
			"url": "https://example.org/testEnv"
		}
	}
}`

	environmentModifiedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.environment.modified.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "environment",
		"content": {
			"name": "testEnv",
			"url": "https://example.org/testEnv"
		}
	}
}`

	environmentDeletedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.environment.deleted.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "environment",
		"content": {
			"name": "testEnv"
		}
	}
}`

	serviceDeployedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.deployed.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	serviceUpgradedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.upgraded.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	serviceRolledBackEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.rolledback.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	serviceRemovedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.removed.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	servicePublishedEventJsonTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.published.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	eventJsonCustomDataTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {}
	},
	"customData": {
		"testValues": [
			{"k1": "v1"},
			{"k2": "v2"}
		]
	},
	"customDataContentType": "application/json"
}`

	eventImplicitJsonCustomDataTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {}
	},
	"customData": {
		"testValues": [
			{"k1": "v1"},
			{"k2": "v2"}
		]
	},
	"customDataContentType": "application/json"
}`

	eventNonJsonCustomDataTemplate = `{
	"context": {
		"version": "draft",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.v1",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {}
	},
	"customData": "PHhtbD50ZXN0RGF0YTwveG1sPg==",
	"customDataContentType": "application/xml"
}`

	pipelineRunQueuedEventJson      string
	pipelineRunStartedEventJson     string
	pipelineRunFinishedEventJson    string
	taskRunStartedEventJson         string
	taskRunFinishedEventJson        string
	changeCreateEventJson           string
	changeUpdatedEventJson          string
	changeReviewedEventJson         string
	changeMergedEventJson           string
	changeAbandonedEventJson        string
	repositoryCreatedEventJson      string
	repositoryModifiedEventJson     string
	repositoryDeletedEventJson      string
	branchCreatedEventJson          string
	branchDeletedEventJson          string
	testCaseQueuedEventJson         string
	testCaseStartedEventJson        string
	testCaseFinishedEventJson       string
	testSuiteStartedEventJson       string
	testSuiteFinishedEventJson      string
	buildQueuedEventJson            string
	buildStartedEventJson           string
	buildFinishedEventJson          string
	artifactPackagedEventJson       string
	artifactPublishedEventJson      string
	environmentCreatedEventJson     string
	environmentModifiedEventJson    string
	environmentDeletedEventJson     string
	serviceDeployedEventJson        string
	serviceUpgradedEventJson        string
	serviceRolledBackEventJson      string
	serviceRemovedEventJson         string
	servicePublishedEventJson       string
	eventJsonCustomDataJson         string
	eventImplicitJsonCustomDataJson string
	eventNonJsonCustomDataJson      string
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

func setContext(event CDEventWriter) {
	event.SetSource(testSource)
	event.SetSubjectId(testSubjectId)
}

func init() {
	pipelineRunQueuedEvent, _ = NewPipelineRunQueuedEvent()
	setContext(pipelineRunQueuedEvent)
	pipelineRunQueuedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunQueuedEvent.SetSubjectUrl(testSubjecturl)

	pipelineRunStartedEvent, _ = NewPipelineRunStartedEvent()
	setContext(pipelineRunStartedEvent)
	pipelineRunStartedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunStartedEvent.SetSubjectUrl(testSubjecturl)

	pipelineRunFinishedEvent, _ = NewPipelineRunFinishedEvent()
	setContext(pipelineRunFinishedEvent)
	pipelineRunFinishedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunFinishedEvent.SetSubjectUrl(testSubjecturl)
	pipelineRunFinishedEvent.SetSubjectOutcome(testPipelineOutcome)
	pipelineRunFinishedEvent.SetSubjectErrors(testPipelineErrors)

	taskRunStartedEvent, _ = NewTaskRunStartedEvent()
	setContext(taskRunStartedEvent)
	taskRunStartedEvent.SetSubjectTaskName(testTaskName)
	taskRunStartedEvent.SetSubjectUrl(testSubjecturl)
	taskRunStartedEvent.SetSubjectPipelineRun(Reference{Id: testSubjectId})

	taskRunFinishedEvent, _ = NewTaskRunFinishedEvent()
	setContext(taskRunFinishedEvent)
	taskRunFinishedEvent.SetSubjectTaskName(testTaskName)
	taskRunFinishedEvent.SetSubjectUrl(testSubjecturl)
	taskRunFinishedEvent.SetSubjectPipelineRun(Reference{Id: testSubjectId})
	taskRunFinishedEvent.SetSubjectOutcome(testTaskOutcome)
	taskRunFinishedEvent.SetSubjectErrors(testTaskRunErrors)

	changeCreatedEvent, _ = NewChangeCreatedEvent()
	setContext(changeCreatedEvent)

	changeUpdatedEvent, _ = NewChangeUpdatedEvent()
	setContext(changeUpdatedEvent)

	changeReviewedEvent, _ = NewChangeReviewedEvent()
	setContext(changeReviewedEvent)

	changeMergedEvent, _ = NewChangeMergedEvent()
	setContext(changeMergedEvent)

	changeAbandonedEvent, _ = NewChangeAbandonedEvent()
	setContext(changeAbandonedEvent)

	repositoryCreatedEvent, _ = NewRepositoryCreatedEvent()
	setContext(repositoryCreatedEvent)
	repositoryCreatedEvent.SetSubjectName(testRepo)
	repositoryCreatedEvent.SetSubjectOwner(testOwner)
	repositoryCreatedEvent.SetSubjectUrl(testUrl)
	repositoryCreatedEvent.SetSubjectViewUrl(testViewUrl)

	repositoryModifiedEvent, _ = NewRepositoryModifiedEvent()
	setContext(repositoryModifiedEvent)
	repositoryModifiedEvent.SetSubjectName(testRepo)
	repositoryModifiedEvent.SetSubjectOwner(testOwner)
	repositoryModifiedEvent.SetSubjectUrl(testUrl)
	repositoryModifiedEvent.SetSubjectViewUrl(testViewUrl)

	repositoryDeletedEvent, _ = NewRepositoryDeletedEvent()
	setContext(repositoryDeletedEvent)
	repositoryDeletedEvent.SetSubjectName(testRepo)
	repositoryDeletedEvent.SetSubjectOwner(testOwner)
	repositoryDeletedEvent.SetSubjectUrl(testUrl)
	repositoryDeletedEvent.SetSubjectViewUrl(testViewUrl)

	branchCreatedEvent, _ = NewBranchCreatedEvent()
	setContext(branchCreatedEvent)

	branchDeletedEvent, _ = NewBranchDeletedEvent()
	setContext(branchDeletedEvent)

	testCaseQueuedEvent, _ = NewTestCaseQueuedEvent()
	setContext(testCaseQueuedEvent)

	testCaseStartedEvent, _ = NewTestCaseStartedEvent()
	setContext(testCaseStartedEvent)

	testCaseFinishedEvent, _ = NewTestCaseFinishedEvent()
	setContext(testCaseFinishedEvent)

	testSuiteStartedEvent, _ = NewTestSuiteStartedEvent()
	setContext(testSuiteStartedEvent)

	testSuiteFinishedEvent, _ = NewTestSuiteFinishedEvent()
	setContext(testSuiteFinishedEvent)

	buildQueuedEvent, _ = NewBuildQueuedEvent()
	setContext(buildQueuedEvent)

	buildStartedEvent, _ = NewBuildStartedEvent()
	setContext(buildStartedEvent)

	buildFinishedEvent, _ = NewBuildFinishedEvent()
	setContext(buildFinishedEvent)
	buildFinishedEvent.SetSubjectArtifactId(testArtifactId)

	artifactPackagedEvent, _ = NewArtifactPackagedEvent()
	setContext(artifactPackagedEvent)

	artifactPublishedEvent, _ = NewArtifactPublishedEvent()
	setContext(artifactPublishedEvent)

	environmentCreatedEvent, _ = NewEnvironmentCreatedEvent()
	setContext(environmentCreatedEvent)
	environmentCreatedEvent.SetSubjectName(testEnvironmentName)
	environmentCreatedEvent.SetSubjectUrl(testEnvironmentUrl)

	environmentModifiedEvent, _ = NewEnvironmentModifiedEvent()
	setContext(environmentModifiedEvent)
	environmentModifiedEvent.SetSubjectName(testEnvironmentName)
	environmentModifiedEvent.SetSubjectUrl(testEnvironmentUrl)

	environmentDeletedEvent, _ = NewEnvironmentDeletedEvent()
	setContext(environmentDeletedEvent)
	environmentDeletedEvent.SetSubjectName(testEnvironmentName)

	serviceDeployedEvent, _ = NewServiceDeployedEvent()
	setContext(serviceDeployedEvent)
	serviceDeployedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	serviceUpgradedEvent, _ = NewServiceUpgradedEvent()
	setContext(serviceUpgradedEvent)
	serviceUpgradedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	serviceRolledBackEvent, _ = NewServiceRolledbackEvent()
	setContext(serviceRolledBackEvent)
	serviceRolledBackEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	serviceRemovedEvent, _ = NewServiceRemovedEvent()
	setContext(serviceRemovedEvent)
	serviceRemovedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	servicePublishedEvent, _ = NewServicePublishedEvent()
	setContext(servicePublishedEvent)
	servicePublishedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	eventJsonCustomData, _ = NewArtifactPackagedEvent()
	setContext(eventJsonCustomData)
	err := eventJsonCustomData.SetCustomData("application/json", testDataJson)
	panicOnError(err)

	eventNonJsonCustomData, _ = NewArtifactPackagedEvent()
	setContext(eventNonJsonCustomData)
	err = eventNonJsonCustomData.SetCustomData("application/xml", testDataXml)
	panicOnError(err)

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
	repositoryCreatedEventJson = fmt.Sprintf(repositoryCreatedEventJsonTemplate, newUUID, newTime)
	repositoryModifiedEventJson = fmt.Sprintf(repositoryModifiedEventJsonTemplate, newUUID, newTime)
	repositoryDeletedEventJson = fmt.Sprintf(repositoryDeletedEventJsonTemplate, newUUID, newTime)
	branchCreatedEventJson = fmt.Sprintf(branchCreatedEventJsonTemplate, newUUID, newTime)
	branchDeletedEventJson = fmt.Sprintf(branchDeletedEventJsonTemplate, newUUID, newTime)
	testCaseQueuedEventJson = fmt.Sprintf(testCaseQueuedEventJsonTemplate, newUUID, newTime)
	testCaseStartedEventJson = fmt.Sprintf(testCaseStartedEventJsonTemplate, newUUID, newTime)
	testCaseFinishedEventJson = fmt.Sprintf(testCaseFinishedEventJsonTemplate, newUUID, newTime)
	testSuiteStartedEventJson = fmt.Sprintf(testSuiteStartedEventJsonTemplate, newUUID, newTime)
	testSuiteFinishedEventJson = fmt.Sprintf(testSuiteFinishedEventJsonTemplate, newUUID, newTime)
	buildQueuedEventJson = fmt.Sprintf(buildQueuedEventJsonTemplate, newUUID, newTime)
	buildStartedEventJson = fmt.Sprintf(buildStartedEventJsonTemplate, newUUID, newTime)
	buildFinishedEventJson = fmt.Sprintf(buildFinishedEventJsonTemplate, newUUID, newTime)
	artifactPackagedEventJson = fmt.Sprintf(artifactPackagedEventJsonTemplate, newUUID, newTime)
	artifactPublishedEventJson = fmt.Sprintf(artifactPublishedEventJsonTemplate, newUUID, newTime)
	environmentCreatedEventJson = fmt.Sprintf(environmentCreatedEventJsonTemplate, newUUID, newTime)
	environmentModifiedEventJson = fmt.Sprintf(environmentModifiedEventJsonTemplate, newUUID, newTime)
	environmentDeletedEventJson = fmt.Sprintf(environmentDeletedEventJsonTemplate, newUUID, newTime)
	serviceDeployedEventJson = fmt.Sprintf(serviceDeployedEventJsonTemplate, newUUID, newTime)
	serviceUpgradedEventJson = fmt.Sprintf(serviceUpgradedEventJsonTemplate, newUUID, newTime)
	serviceRolledBackEventJson = fmt.Sprintf(serviceRolledBackEventJsonTemplate, newUUID, newTime)
	serviceRemovedEventJson = fmt.Sprintf(serviceRemovedEventJsonTemplate, newUUID, newTime)
	servicePublishedEventJson = fmt.Sprintf(servicePublishedEventJsonTemplate, newUUID, newTime)
	eventJsonCustomDataJson = fmt.Sprintf(eventJsonCustomDataTemplate, newUUID, newTime)
	eventImplicitJsonCustomDataJson = fmt.Sprintf(eventImplicitJsonCustomDataTemplate, newUUID, newTime)
	eventNonJsonCustomDataJson = fmt.Sprintf(eventNonJsonCustomDataTemplate, newUUID, newTime)
}

func TestAsCloudEvent(t *testing.T) {

	tests := []struct {
		name            string
		event           CDEventReader
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
	}, {
		name:            "repository created",
		event:           repositoryCreatedEvent,
		payloadReceiver: &RepositoryCreatedEvent{},
	}, {
		name:            "repository modified",
		event:           repositoryModifiedEvent,
		payloadReceiver: &RepositoryModifiedEvent{},
	}, {
		name:            "repository deleted",
		event:           repositoryDeletedEvent,
		payloadReceiver: &RepositoryDeletedEvent{},
	}, {
		name:            "branch created",
		event:           branchCreatedEvent,
		payloadReceiver: &BranchCreatedEvent{},
	}, {
		name:            "branch deleted",
		event:           branchDeletedEvent,
		payloadReceiver: &BranchDeletedEvent{},
	}, {
		name:            "testcase queued",
		event:           testCaseQueuedEvent,
		payloadReceiver: &TestCaseQueuedEvent{},
	}, {
		name:            "testcase started",
		event:           testCaseStartedEvent,
		payloadReceiver: &TestCaseStartedEvent{},
	}, {
		name:            "testcase finished",
		event:           testCaseFinishedEvent,
		payloadReceiver: &TestCaseFinishedEvent{},
	}, {
		name:            "testsuite started",
		event:           testSuiteStartedEvent,
		payloadReceiver: &TestSuiteStartedEvent{},
	}, {
		name:            "testsuite finished",
		event:           testSuiteFinishedEvent,
		payloadReceiver: &TestSuiteFinishedEvent{},
	}, {
		name:            "build queued",
		event:           buildQueuedEvent,
		payloadReceiver: &BuildQueuedEvent{},
	}, {
		name:            "build started",
		event:           buildStartedEvent,
		payloadReceiver: &BuildStartedEvent{},
	}, {
		name:            "build finished",
		event:           buildFinishedEvent,
		payloadReceiver: &BuildFinishedEvent{},
	}, {
		name:            "artifact packaged",
		event:           artifactPackagedEvent,
		payloadReceiver: &ArtifactPackagedEvent{},
	}, {
		name:            "artifact published",
		event:           artifactPublishedEvent,
		payloadReceiver: &ArtifactPublishedEvent{},
	}, {
		name:            "environment created",
		event:           environmentCreatedEvent,
		payloadReceiver: &EnvironmentCreatedEvent{},
	}, {
		name:            "environment modified",
		event:           environmentModifiedEvent,
		payloadReceiver: &EnvironmentModifiedEvent{},
	}, {
		name:            "environment deleted",
		event:           environmentDeletedEvent,
		payloadReceiver: &EnvironmentDeletedEvent{},
	}, {
		name:            "service deployed",
		event:           serviceDeployedEvent,
		payloadReceiver: &ServiceDeployedEvent{},
	}, {
		name:            "service upgraded",
		event:           serviceUpgradedEvent,
		payloadReceiver: &ServiceUpgradedEvent{},
	}, {
		name:            "service rolledback",
		event:           serviceRolledBackEvent,
		payloadReceiver: &ServiceRolledbackEvent{},
	}, {
		name:            "service removed",
		event:           serviceRemovedEvent,
		payloadReceiver: &ServiceRemovedEvent{},
	}, {
		name:            "service published",
		event:           servicePublishedEvent,
		payloadReceiver: &ServicePublishedEvent{},
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
				t.Fatalf("somehow cannot unmarshal test event %v, %v", ce, err)
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
	}, {
		name:       "repository created",
		event:      repositoryCreatedEvent,
		jsonString: repositoryCreatedEventJson,
		schemaName: "repositorycreated",
	}, {
		name:       "repository modified",
		event:      repositoryModifiedEvent,
		jsonString: repositoryModifiedEventJson,
		schemaName: "repositorymodified",
	}, {
		name:       "repository deleted",
		event:      repositoryDeletedEvent,
		jsonString: repositoryDeletedEventJson,
		schemaName: "repositorydeleted",
	}, {
		name:       "branch created",
		event:      branchCreatedEvent,
		jsonString: branchCreatedEventJson,
		schemaName: "branchcreated",
	}, {
		name:       "branch deleted",
		event:      branchDeletedEvent,
		jsonString: branchDeletedEventJson,
		schemaName: "branchdeleted",
	}, {
		name:       "testcase queued",
		event:      testCaseQueuedEvent,
		jsonString: testCaseQueuedEventJson,
		schemaName: "testcasequeued",
	}, {
		name:       "testcase started",
		event:      testCaseStartedEvent,
		jsonString: testCaseStartedEventJson,
		schemaName: "testcasestarted",
	}, {
		name:       "testcase finished",
		event:      testCaseFinishedEvent,
		jsonString: testCaseFinishedEventJson,
		schemaName: "testcasefinished",
	}, {
		name:       "testsuite started",
		event:      testSuiteStartedEvent,
		jsonString: testSuiteStartedEventJson,
		schemaName: "testsuitestarted",
	}, {
		name:       "testsuite finished",
		event:      testSuiteFinishedEvent,
		jsonString: testSuiteFinishedEventJson,
		schemaName: "testsuitefinished",
	}, {
		name:       "build queued",
		event:      buildQueuedEvent,
		jsonString: buildQueuedEventJson,
		schemaName: "buildqueued",
	}, {
		name:       "build started",
		event:      buildStartedEvent,
		jsonString: buildStartedEventJson,
		schemaName: "buildstarted",
	}, {
		name:       "build finished",
		event:      buildFinishedEvent,
		jsonString: buildFinishedEventJson,
		schemaName: "buildfinished",
	}, {
		name:       "artifact packaged",
		event:      artifactPackagedEvent,
		jsonString: artifactPackagedEventJson,
		schemaName: "artifactpackaged",
	}, {
		name:       "artifact published",
		event:      artifactPublishedEvent,
		jsonString: artifactPublishedEventJson,
		schemaName: "artifactpublished",
	}, {
		name:       "environment created",
		event:      environmentCreatedEvent,
		jsonString: environmentCreatedEventJson,
		schemaName: "environmentcreated",
	}, {
		name:       "environment modified",
		event:      environmentModifiedEvent,
		jsonString: environmentModifiedEventJson,
		schemaName: "environmentmodified",
	}, {
		name:       "environment deleted",
		event:      environmentDeletedEvent,
		jsonString: environmentDeletedEventJson,
		schemaName: "environmentdeleted",
	}, {
		name:       "service deployed",
		event:      serviceDeployedEvent,
		jsonString: serviceDeployedEventJson,
		schemaName: "servicedeployed",
	}, {
		name:       "service upgraded",
		event:      serviceUpgradedEvent,
		jsonString: serviceUpgradedEventJson,
		schemaName: "serviceupgraded",
	}, {
		name:       "service rolledback",
		event:      serviceRolledBackEvent,
		jsonString: serviceRolledBackEventJson,
		schemaName: "servicerolledback",
	}, {
		name:       "service removed",
		event:      serviceRemovedEvent,
		jsonString: serviceRemovedEventJson,
		schemaName: "serviceremoved",
	}, {
		name:       "service published",
		event:      servicePublishedEvent,
		jsonString: servicePublishedEventJson,
		schemaName: "servicepublished",
	}, {
		name:       "json custom data",
		event:      eventJsonCustomData,
		jsonString: eventJsonCustomDataJson,
		schemaName: "artifactpackaged",
	}, {
		name:       "json custom data implicit",
		event:      eventJsonCustomData,
		jsonString: eventImplicitJsonCustomDataJson,
		schemaName: "artifactpackaged",
	}, {
		name:       "xml custom data",
		event:      eventNonJsonCustomData,
		jsonString: eventNonJsonCustomDataJson,
		schemaName: "artifactpackaged",
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

	// mandatory source missing
	eventNoSource, _ := NewCDEvent(ChangeAbandonedEventV1)
	eventNoSource.SetSubjectId(testSubjectId)

	// mandatory subject id missing
	eventNoSubjectId, _ := NewCDEvent(ChangeAbandonedEventV1)
	eventNoSubjectId.SetSource(testSource)

	// forced invalid version
	eventBadVersion, _ := NewChangeAbandonedEvent()
	eventBadVersion.Context.Version = "invalid"

	// mandatory subject url missing
	eventIncompleteSubject, _ := NewRepositoryCreatedEvent()
	eventIncompleteSubject.SetSource(testSource)
	eventIncompleteSubject.SetSubjectId(testSubjectId)
	eventIncompleteSubject.SetSubjectName(testRepo)

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
		name:  "invalid version",
		event: eventBadVersion,
	}, {
		name:  "missing subject url",
		event: eventIncompleteSubject,
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
