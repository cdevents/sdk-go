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

package v03_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cdevents/sdk-go/pkg/api"
	apiv03 "github.com/cdevents/sdk-go/pkg/api/v03"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

const examplesFolder = "spec-v0.3/examples"

var (
	// Examples Data
	testArtifactSubjectID = "pkg:golang/mygit.com/myorg/myapp@234fd47e07d1004f0aed9c"
	testChangeSource      = "my-git.example/an-org/a-repo"
	testPipeline          = "myPipeline"
	testSubjectURL        = "https://www.example.com/mySubject123"
	testPipelineOutcome   = "failure"
	testPipelineErrors    = "Something went wrong\nWith some more details"
	testTaskName          = "myTask"
	testTaskOutcome       = "failure"
	testTaskRunErrors     = "Something went wrong\nWith some more details"
	testRepo              = "TestRepo"
	testOwner             = "TestOrg"
	testURL               = "https://example.org/TestOrg/TestRepo"
	testViewURL           = "https://example.org/view/TestOrg/TestRepo"
	testEnvironmentID     = "test123"
	testEnvironmentName   = "testEnv"
	testEnvironmentURL    = "https://example.org/testEnv"
	testRepoReference     = &api.Reference{
		ID: "TestRepo/TestOrg", Source: "https://example.org"}
	testIncidentSubjectID    = "incident-123"
	testIncidentSource       = "/monitoring/prod1"
	testEnvironmentReference = &api.Reference{
		ID: "prod1", Source: "/iaas/geo1"}
	testServiceReference = &api.Reference{
		ID: "myApp", Source: "/clusterA/namespaceB"}
	testTestRunID       = "myTestCaseRun123"
	testSignature       = "MEYCIQCBT8U5ypDXWCjlNKfzTV4KH516/SK13NZSh8znnSMNkQIhAJ3XiQlc9PM1KyjITcZXHotdMB+J3NGua5T/yshmiPmp"
	testTestEnvironment = &api.Reference{
		ID: "dev", Source: "testkube-dev-123"}
	testTestCaseStarted = &apiv03.TestCaseRunStartedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseFinished = &apiv03.TestCaseRunFinishedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseQueued = &apiv03.TestCaseRunQueuedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestTriggerQueued = &apiv03.TestCaseRunQueuedSubjectContentTrigger{
		Type: "schedule"}
	testTestTriggerStarted = &apiv03.TestCaseRunStartedSubjectContentTrigger{
		Type: "schedule"}
	testTestOutcome             = "pass"
	testTestOutputSubjectID     = "testrunreport-12123"
	testTestOutputSubjectSource = "/event/source/testrunreport-12123"
	testTestOutputFormat        = "video/quicktime"
	testTestOutputOutputType    = "video"
	testTestCaseRun             = &api.Reference{ID: testTestRunID, Source: "testkube-dev-123"}
	testTestSuiteRunID          = "myTestSuiteRun123"
	testTestSuiteStarted        = &apiv03.TestSuiteRunStartedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteQueued = &apiv03.TestSuiteRunQueuedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteFinished = &apiv03.TestSuiteRunFinishedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteOutcome        = "fail"
	testTestSuiteReason         = "Host 123.34.23.32 not found"
	testTestSuiteSeverity       = "critical"
	testTestSuiteTriggerQueued  = &apiv03.TestSuiteRunQueuedSubjectContentTrigger{Type: "pipeline"}
	testTestSuiteTriggerStarted = &apiv03.TestSuiteRunStartedSubjectContentTrigger{Type: "pipeline"}

	examplesConsumed map[string][]byte
	examplesProduced map[string]api.CDEvent
	err              error
)

func init() {
	// Get the time once
	t, _ := time.Parse(time.RFC3339Nano, "2023-03-20T14:27:05.315384Z")
	timeNow = func() time.Time {
		return t
	}

	// Set the UUID to a fixed one
	u, _ := uuid.Parse("271069a8-fc18-44f1-b38f-9d70a1695819")
	uuidNewRandom = func() (uuid.UUID, error) {
		return u, nil
	}
}

func exampleArtifactPackagedEvent(e *apiv03.ArtifactPackagedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectChange(&api.Reference{ID: testChangeId, Source: testChangeSource})
}

func exampleArtifactPublishedEvent(e *apiv03.ArtifactPublishedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
}

func exampleArtifactSignedEvent(e *apiv03.ArtifactSignedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectSignature(testSignature)
}

func exampleBranchCreatedEvent(e *apiv03.BranchCreatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBranchDeletedEvent(e *apiv03.BranchDeletedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBuildFinishedEvent(e *apiv03.BuildFinishedEvent) {
	// Set example specific fields
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleBuildQueuedEvent(_ *apiv03.BuildQueuedEvent) {
	// Set example specific fields
}

func exampleBuildStartedEvent(_ *apiv03.BuildStartedEvent) {
	// Set example specific fields
}

func exampleChangeAbandonedEvent(e *apiv03.ChangeAbandonedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeCreatedEvent(e *apiv03.ChangeCreatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeMergedEvent(e *apiv03.ChangeMergedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeReviewedEvent(e *apiv03.ChangeReviewedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeUpdatedEvent(e *apiv03.ChangeUpdatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleEnvironmentCreatedEvent(e *apiv03.EnvironmentCreatedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUrl(testEnvironmentURL)
}

func exampleEnvironmentDeletedEvent(e *apiv03.EnvironmentDeletedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
}

func exampleEnvironmentModifiedEvent(e *apiv03.EnvironmentModifiedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUrl(testEnvironmentURL)
}

func exampleIncidentDetectedEvent(e *apiv03.IncidentDetectedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectID)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId(testArtifactId)
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time above threshold of 100ms")
}

func exampleIncidentReportedEvent(e *apiv03.IncidentReportedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectID)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId(testArtifactId)
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time above threshold of 100ms")
	e.SetSubjectTicketURI("https://my-issues.example/incidents/ticket-345")
}

func exampleIncidentResolvedEvent(e *apiv03.IncidentResolvedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectID)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId("pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93439")
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time restored below 100ms")
}

func examplePipelineRunFinishedEvent(e *apiv03.PipelineRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjectURL)
	e.SetSubjectOutcome(testPipelineOutcome)
	e.SetSubjectErrors(testPipelineErrors)
}

func examplePipelineRunQueuedEvent(e *apiv03.PipelineRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjectURL)
}

func examplePipelineRunStartedEvent(e *apiv03.PipelineRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjectURL)
}

func exampleRepositoryCreatedEvent(e *apiv03.RepositoryCreatedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testURL)
	e.SetSubjectViewUrl(testViewURL)
}

func exampleRepositoryDeletedEvent(e *apiv03.RepositoryDeletedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testURL)
	e.SetSubjectViewUrl(testViewURL)
}

func exampleRepositoryModifiedEvent(e *apiv03.RepositoryModifiedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testURL)
	e.SetSubjectViewUrl(testViewURL)
}

func exampleServiceDeployedEvent(e *apiv03.ServiceDeployedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServicePublishedEvent(e *apiv03.ServicePublishedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
}

func exampleServiceRemovedEvent(e *apiv03.ServiceRemovedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
}

func exampleServiceRolledbackEvent(e *apiv03.ServiceRolledbackEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServiceUpgradedEvent(e *apiv03.ServiceUpgradedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleTaskRunFinishedEvent(e *apiv03.TaskRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUrl(testSubjectURL)
	e.SetSubjectPipelineRun(&api.Reference{ID: testSubjectId})
	e.SetSubjectOutcome(testTaskOutcome)
	e.SetSubjectErrors(testTaskRunErrors)
}

func exampleTaskRunStartedEvent(e *apiv03.TaskRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUrl(testSubjectURL)
	e.SetSubjectPipelineRun(&api.Reference{ID: testSubjectId})
}

func exampleTestCaseRunFinishedEvent(e *apiv03.TestCaseRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunID)
	e.SetSubjectId(testTestRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseFinished)
	e.SetSubjectOutcome(testTestOutcome)
}

func exampleTestCaseRunQueuedEvent(e *apiv03.TestCaseRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunID)
	e.SetSubjectId(testTestRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseQueued)
	e.SetSubjectTrigger(testTestTriggerQueued)
}

func exampleTestCaseRunStartedEvent(e *apiv03.TestCaseRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseStarted)
	e.SetSubjectTrigger(testTestTriggerStarted)
}

func exampleTestSuiteRunFinishedEvent(e *apiv03.TestSuiteRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteFinished)
	e.SetSubjectOutcome(testTestSuiteOutcome)
	e.SetSubjectSeverity(testTestSuiteSeverity)
	e.SetSubjectReason(testTestSuiteReason)
}

func exampleTestSuiteRunStartedEvent(e *apiv03.TestSuiteRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteStarted)
	e.SetSubjectTrigger(testTestSuiteTriggerStarted)
}

func exampleTestSuiteRunQueuedEvent(e *apiv03.TestSuiteRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteQueued)
	e.SetSubjectTrigger(testTestSuiteTriggerQueued)
}

func exampleTestOutputPublishedEvent(e *apiv03.TestOutputPublishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestOutputSubjectID)
	e.SetSubjectSource(testTestOutputSubjectSource)
	e.SetSubjectOutputType(testTestOutputOutputType)
	e.SetSubjectFormat(testTestOutputFormat)
	e.SetSubjectTestCaseRun(testTestCaseRun)
}

func init() {
	// Load event examples from the spec
	examplesConsumed = make(map[string][]byte)

	for _, event := range apiv03.CDEventsTypes {
		short := event.GetType().Short()
		examplesConsumed[short], err = os.ReadFile(filepath.Join("..", examplesFolder, short+".json"))
		panicOnError(err)
	}
}

// TestExamples verifies that the SDK can produce events like those
// included in the specification examples folder.
// To do so:
// - it produces a CDEvent from scratch, and sets the values like in the examples
// - it parses the examples into a CDEvent and
// - it verifies that produced and consumed CDEvent match
func TestExamples(t *testing.T) {
	for name, exampleConsumed := range examplesConsumed {
		t.Run(name, func(t *testing.T) {
			produced, ok := examplesProduced[name]
			if !ok {
				t.Fatalf("missing produced event for event type: %v", name)
			}
			// Check that the produced event is valid
			err := api.Validate(produced)
			if err != nil {
				t.Errorf("produced event failed to validate: %v", err)
			}
			consumed, err := apiv03.NewFromJSONBytes(exampleConsumed)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			// Check the context, except for ID and Timestamp
			if d := cmp.Diff(consumed.GetVersion(), produced.GetVersion()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(consumed.GetSource(), produced.GetSource()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(consumed.GetType(), produced.GetType()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check the subject
			if d := cmp.Diff(consumed.GetSubject(), produced.GetSubject()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Coverage for GetSubjectContent
			if d := cmp.Diff(consumed.GetSubjectContent(), produced.GetSubjectContent()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}
