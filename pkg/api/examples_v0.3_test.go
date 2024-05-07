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

package api_test

import (
	"os"
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
	testArtifactSubjectId = "pkg:golang/mygit.com/myorg/myapp@234fd47e07d1004f0aed9c"
	testChangeSource      = "my-git.example/an-org/a-repo"
	testPipeline          = "myPipeline"
	testSubjecturl        = "https://www.example.com/mySubject123"
	testPipelineOutcome   = "failure"
	testPipelineErrors    = "Something went wrong\nWith some more details"
	testTaskName          = "myTask"
	testTaskOutcome       = "failure"
	testTaskRunErrors     = "Something went wrong\nWith some more details"
	testRepo              = "TestRepo"
	testOwner             = "TestOrg"
	testUrl               = "https://example.org/TestOrg/TestRepo"
	testViewUrl           = "https://example.org/view/TestOrg/TestRepo"
	testEnvironmentId     = "test123"
	testEnvironmentName   = "testEnv"
	testEnvironmentUrl    = "https://example.org/testEnv"
	testRepoReference     = &api.Reference{
		Id: "TestRepo/TestOrg", Source: "https://example.org"}
	testIncidentSubjectId    = "incident-123"
	testIncidentSource       = "/monitoring/prod1"
	testEnvironmentReference = &api.Reference{
		Id: "prod1", Source: "/iaas/geo1"}
	testServiceReference = &api.Reference{
		Id: "myApp", Source: "/clusterA/namespaceB"}
	testTestRunId       = "myTestCaseRun123"
	testSignature       = "MEYCIQCBT8U5ypDXWCjlNKfzTV4KH516/SK13NZSh8znnSMNkQIhAJ3XiQlc9PM1KyjITcZXHotdMB+J3NGua5T/yshmiPmp"
	testTestEnvironment = &api.Reference{
		Id: "dev", Source: "testkube-dev-123"}
	testTestCaseStarted = &api.TestCaseRunStartedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseFinished = &api.TestCaseRunFinishedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseQueued = &api.TestCaseRunQueuedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestTriggerQueued = &api.TestCaseRunQueuedSubjectContentTrigger{
		Type: "schedule"}
	testTestTriggerStarted = &api.TestCaseRunStartedSubjectContentTrigger{
		Type: "schedule"}
	testTestOutcome             = "pass"
	testTestOutputSubjectId     = "testrunreport-12123"
	testTestOutputSubjectSource = "/event/source/testrunreport-12123"
	testTestOutputFormat        = "video/quicktime"
	testTestOutputOutputType    = "video"
	testTestCaseRun             = &api.Reference{Id: testTestRunId, Source: "testkube-dev-123"}
	testTestSuiteRunId          = "myTestSuiteRun123"
	testTestSuiteStarted        = &api.TestSuiteRunStartedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteQueued = &api.TestSuiteRunQueuedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteFinished = &api.TestSuiteRunFinishedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteOutcome        = "fail"
	testTestSuiteReason         = "Host 123.34.23.32 not found"
	testTestSuiteSeverity       = "critical"
	testTestSuiteTriggerQueued  = &api.TestSuiteRunQueuedSubjectContentTrigger{Type: "pipeline"}
	testTestSuiteTriggerStarted = &api.TestSuiteRunStartedSubjectContentTrigger{Type: "pipeline"}

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

func exampleArtifactPackagedEvent(e *api.ArtifactPackagedEventV0_1_1) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectChange(&api.Reference{Id: testChangeId, Source: testChangeSource})
}

func exampleArtifactPublishedEvent(e *api.ArtifactPublishedEventV0_1_1) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
}

func exampleArtifactSignedEvent(e *api.ArtifactSignedEventV0_1_0) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectSignature(testSignature)
}

func exampleBranchCreatedEvent(e *api.BranchCreatedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBranchDeletedEvent(e *api.BranchDeletedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBuildFinishedEvent(e *api.BuildFinishedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleBuildQueuedEvent(e *api.BuildQueuedEventV0_1_1) {
	// Set example specific fields
}

func exampleBuildStartedEvent(e *api.BuildStartedEventV0_1_1) {
	// Set example specific fields
}

func exampleChangeAbandonedEvent(e *api.ChangeAbandonedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeCreatedEvent(e *api.ChangeCreatedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeMergedEvent(e *api.ChangeMergedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeReviewedEvent(e *api.ChangeReviewedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeUpdatedEvent(e *api.ChangeUpdatedEventV0_1_2) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleEnvironmentCreatedEvent(e *api.EnvironmentCreatedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUrl(testEnvironmentUrl)
}

func exampleEnvironmentDeletedEvent(e *api.EnvironmentDeletedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
}

func exampleEnvironmentModifiedEvent(e *api.EnvironmentModifiedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUrl(testEnvironmentUrl)
}

func exampleIncidentDetectedEvent(e *api.IncidentDetectedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectId)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId(testArtifactId)
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time above threshold of 100ms")
}

func exampleIncidentReportedEvent(e *api.IncidentReportedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectId)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId(testArtifactId)
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time above threshold of 100ms")
	e.SetSubjectTicketURI("https://my-issues.example/incidents/ticket-345")
}

func exampleIncidentResolvedEvent(e *api.IncidentResolvedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectId)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId("pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93439")
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time restored below 100ms")
}

func examplePipelineRunFinishedEvent(e *api.PipelineRunFinishedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjecturl)
	e.SetSubjectOutcome(testPipelineOutcome)
	e.SetSubjectErrors(testPipelineErrors)
}

func examplePipelineRunQueuedEvent(e *api.PipelineRunQueuedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjecturl)
}

func examplePipelineRunStartedEvent(e *api.PipelineRunStartedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjecturl)
}

func exampleRepositoryCreatedEvent(e *api.RepositoryCreatedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testUrl)
	e.SetSubjectViewUrl(testViewUrl)
}

func exampleRepositoryDeletedEvent(e *api.RepositoryDeletedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testUrl)
	e.SetSubjectViewUrl(testViewUrl)
}

func exampleRepositoryModifiedEvent(e *api.RepositoryModifiedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testUrl)
	e.SetSubjectViewUrl(testViewUrl)
}

func exampleServiceDeployedEvent(e *api.ServiceDeployedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServicePublishedEvent(e *api.ServicePublishedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
}

func exampleServiceRemovedEvent(e *api.ServiceRemovedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
}

func exampleServiceRolledbackEvent(e *api.ServiceRolledbackEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServiceUpgradedEvent(e *api.ServiceUpgradedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleTaskRunFinishedEvent(e *api.TaskRunFinishedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUrl(testSubjecturl)
	e.SetSubjectPipelineRun(&api.Reference{Id: testSubjectId})
	e.SetSubjectOutcome(testTaskOutcome)
	e.SetSubjectErrors(testTaskRunErrors)
}

func exampleTaskRunStartedEvent(e *api.TaskRunStartedEventV0_1_1) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUrl(testSubjecturl)
	e.SetSubjectPipelineRun(&api.Reference{Id: testSubjectId})
}

func exampleTestCaseRunFinishedEvent(e *api.TestCaseRunFinishedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestRunId)
	e.SetSubjectId(testTestRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseFinished)
	e.SetSubjectOutcome(testTestOutcome)
}

func exampleTestCaseRunQueuedEvent(e *api.TestCaseRunQueuedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestRunId)
	e.SetSubjectId(testTestRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseQueued)
	e.SetSubjectTrigger(testTestTriggerQueued)
}

func exampleTestCaseRunStartedEvent(e *api.TestCaseRunStartedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseStarted)
	e.SetSubjectTrigger(testTestTriggerStarted)
}

func exampleTestSuiteRunFinishedEvent(e *api.TestSuiteRunFinishedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteFinished)
	e.SetSubjectOutcome(testTestSuiteOutcome)
	e.SetSubjectSeverity(testTestSuiteSeverity)
	e.SetSubjectReason(testTestSuiteReason)
}

func exampleTestSuiteRunStartedEvent(e *api.TestSuiteRunStartedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteStarted)
	e.SetSubjectTrigger(testTestSuiteTriggerStarted)
}

func exampleTestSuiteRunQueuedEvent(e *api.TestSuiteRunQueuedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteQueued)
	e.SetSubjectTrigger(testTestSuiteTriggerQueued)
}

func exampleTestOutputPublishedEvent(e *api.TestOutputPublishedEventV0_1_0) {
	// Set example specific fields
	e.SetSubjectId(testTestOutputSubjectId)
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
		examplesConsumed[short], err = os.ReadFile(examplesFolder + string(os.PathSeparator) + short + ".json")
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
			consumed, err := apiv03.NewFromJsonBytes(exampleConsumed)
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
		})
	}
}
