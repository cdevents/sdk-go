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

package v04_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cdevents/sdk-go/pkg/api"
	apiv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

const (
	examplesFolder = "spec-v0.4/conformance"
	customExample  = "spec-v0.4/custom/conformance.json"
)

var (
	// Examples Data
	testArtifactSubjectID = "pkg:golang/mygit.com/myorg/myapp@234fd47e07d1004f0aed9c"
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
	testTestCaseStarted = &apiv04.TestCaseRunStartedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseFinished = &apiv04.TestCaseRunFinishedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseQueued = &apiv04.TestCaseRunQueuedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestTriggerQueued = &apiv04.TestCaseRunQueuedSubjectContentTrigger{
		Type: "schedule"}
	testTestTriggerStarted = &apiv04.TestCaseRunStartedSubjectContentTrigger{
		Type: "schedule"}
	testTestOutcome             = "pass"
	testTestOutputSubjectID     = "testrunreport-12123"
	testTestOutputSubjectSource = "/event/source/testrunreport-12123"
	testTestOutputFormat        = "video/quicktime"
	testTestOutputOutputType    = "video"
	testTestCaseRun             = &api.Reference{ID: testTestRunID, Source: "testkube-dev-123"}
	testTestSuiteRunID          = "myTestSuiteRun123"
	testTestSuiteStarted        = &apiv04.TestSuiteRunStartedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteQueued = &apiv04.TestSuiteRunQueuedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteFinished = &apiv04.TestSuiteRunFinishedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteOutcome        = "fail"
	testTestSuiteReason         = "Host 123.34.23.32 not found"
	testTestSuiteSeverity       = "critical"
	testTestSuiteTriggerQueued  = &apiv04.TestSuiteRunQueuedSubjectContentTrigger{Type: "pipeline"}
	testTestSuiteTriggerStarted = &apiv04.TestSuiteRunStartedSubjectContentTrigger{Type: "pipeline"}
	testSubjectUser             = "mybot-myapp"
	testSbomURI                 = "https://sbom.repo/myorg/234fd47e07d1004f0aed9c.sbom"
	testChangeDescription       = "This PR address a bug from a recent PR"
	testTicketID                = "ticket-123"
	testTicketSource            = "/ticketing/system"
	testTicketAssignees         = []string{"Bob"}
	testTicketCreator           = "Alice"
	testTicketGroup             = "security"
	testTicketLabels            = []string{"bug"}
	testTicketMilestone         = "123"
	testTicketPriority          = "high"
	testTicketResolution        = "completed"
	testTicketSummary           = "New CVE-123 detected"
	testTicketType              = "task"
	testTicketUpdatedBy         = "Bob"
	testTicketURI               = "https://example.issues.com/ticket123"
	testCustomEventType         = api.CDEventType{
		// dev.cdeventsx.mytool-resource.created.0.1.0
		Subject:   "resource",
		Predicate: "created",
		Custom:    "mytool",
		Version:   "0.1.0",
	}
	testCustomContentBytes = []byte(`{
		"user": "mybot-myapp",
		"description": "a useful resource",
		"nested": {
			"key": "value",
			"list": ["data1", "data2"]
		}
    }`)
	testCustomContent            interface{}
	testCustomSchemaID           = "https://myorg.com/schema/custom"
	testCustomSchemaID2          = "https://myorg.com/schema/mytool"
	testCustomSchemaJSONTemplate = `{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"$id": "%s",
	"additionalProperties": true,
	"type": "object"
}`
	testCustomSchemaJSON  = fmt.Sprintf(testCustomSchemaJSONTemplate, testCustomSchemaID)
	testCustomSchema2JSON = fmt.Sprintf(testCustomSchemaJSONTemplate, testCustomSchemaID2)
	testCustomSchemas     = map[string][]byte{
		testCustomSchemaID:  []byte(testCustomSchemaJSON),
		testCustomSchemaID2: []byte(testCustomSchema2JSON),
	}
	examplesConsumed map[string][]byte
	examplesProduced map[string]api.CDEventV04
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

	err = json.Unmarshal(testCustomContentBytes, &testCustomContent)
	panicOnError(err)

	for id, jsonBytes := range testCustomSchemas {
		err := api.LoadJsonSchema(id, jsonBytes)
		panicOnError(err)
	}

	// Load event examples from the spec
	examplesConsumed = make(map[string][]byte)

	for _, event := range apiv04.CDEventsTypes {
		short := event.GetType().Short()
		if short != "" {
			examplesConsumed[short], err = os.ReadFile(filepath.Join("..", examplesFolder, short+".json"))
			panicOnError(err)
		} else {
			// There is no type set for custom events, and the example is in a different folder
			examplesConsumed[short], err = os.ReadFile(filepath.Join("..", customExample))
		}
	}
}

func exampleArtifactPackagedEvent(e *apiv04.ArtifactPackagedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectChange(&api.Reference{ID: testChangeId, Source: testChangeSource})
	e.SetSubjectSbom(&api.ArtifactPackagedSubjectContentSbomV0_2_0{
		Uri: testSbomURI,
	})
}

func exampleArtifactPublishedEvent(e *apiv04.ArtifactPublishedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectUser(testSubjectUser)
	e.SetSubjectSbom(&api.ArtifactPublishedSubjectContentSbomV0_2_0{
		Uri: testSbomURI,
	})
}

func exampleArtifactSignedEvent(e *apiv04.ArtifactSignedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectSignature(testSignature)
}

func exampleArtifactDeletedEvent(e *apiv04.ArtifactDeletedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectUser(testSubjectUser)
	e.SetChainId("")
}

func exampleArtifactDownloadedEvent(e *apiv04.ArtifactDownloadedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectID)
	e.SetSubjectUser(testSubjectUser)
	e.SetChainId("")
}

func exampleBranchCreatedEvent(e *apiv04.BranchCreatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBranchDeletedEvent(e *apiv04.BranchDeletedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBuildFinishedEvent(e *apiv04.BuildFinishedEvent) {
	// Set example specific fields
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleBuildQueuedEvent(_ *apiv04.BuildQueuedEvent) {
	// Set example specific fields
}

func exampleBuildStartedEvent(_ *apiv04.BuildStartedEvent) {
	// Set example specific fields
}

func exampleChangeAbandonedEvent(e *apiv04.ChangeAbandonedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeCreatedEvent(e *apiv04.ChangeCreatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
	e.SetSubjectDescription(testChangeDescription)
}

func exampleChangeMergedEvent(e *apiv04.ChangeMergedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeReviewedEvent(e *apiv04.ChangeReviewedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeUpdatedEvent(e *apiv04.ChangeUpdatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleEnvironmentCreatedEvent(e *apiv04.EnvironmentCreatedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUrl(testEnvironmentURL)
}

func exampleEnvironmentDeletedEvent(e *apiv04.EnvironmentDeletedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
}

func exampleEnvironmentModifiedEvent(e *apiv04.EnvironmentModifiedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUrl(testEnvironmentURL)
}

func exampleIncidentDetectedEvent(e *apiv04.IncidentDetectedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectID)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId(testArtifactId)
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time above threshold of 100ms")
}

func exampleIncidentReportedEvent(e *apiv04.IncidentReportedEvent) {
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

func exampleIncidentResolvedEvent(e *apiv04.IncidentResolvedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectID)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId("pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93439")
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time restored below 100ms")
}

func examplePipelineRunFinishedEvent(e *apiv04.PipelineRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjecturl)
	e.SetSubjectOutcome(testPipelineOutcome)
	e.SetSubjectErrors(testPipelineErrors)
}

func examplePipelineRunQueuedEvent(e *apiv04.PipelineRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjecturl)
}

func examplePipelineRunStartedEvent(e *apiv04.PipelineRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUrl(testSubjecturl)
}

func exampleRepositoryCreatedEvent(e *apiv04.RepositoryCreatedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testURL)
	e.SetSubjectViewUrl(testViewURL)
}

func exampleRepositoryDeletedEvent(e *apiv04.RepositoryDeletedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testURL)
	e.SetSubjectViewUrl(testViewURL)
}

func exampleRepositoryModifiedEvent(e *apiv04.RepositoryModifiedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUrl(testURL)
	e.SetSubjectViewUrl(testViewURL)
}

func exampleServiceDeployedEvent(e *apiv04.ServiceDeployedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServicePublishedEvent(e *apiv04.ServicePublishedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
}

func exampleServiceRemovedEvent(e *apiv04.ServiceRemovedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
}

func exampleServiceRolledbackEvent(e *apiv04.ServiceRolledbackEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServiceUpgradedEvent(e *apiv04.ServiceUpgradedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{ID: testEnvironmentID})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleTaskRunFinishedEvent(e *apiv04.TaskRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUrl(testSubjecturl)
	e.SetSubjectPipelineRun(&api.Reference{ID: testSubjectId})
	e.SetSubjectOutcome(testTaskOutcome)
	e.SetSubjectErrors(testTaskRunErrors)
}

func exampleTestCaseRunSkippedEvent(e *apiv04.TestCaseRunSkippedEvent) {
	// Set example specific fields
	e.SetSubjectId("myTestCaseRun123")
	e.SetSubjectEnvironment(&api.Reference{
		ID:     "dev",
		Source: "testkube-dev-123",
	})
	e.SetSubjectReason("Not running in this environment")
	e.SetSubjectTestCase(&api.TestCaseRunSkippedSubjectContentTestCaseV0_1_0{
		Id:      "92834723894",
		Name:    "Login Test",
		Type:    "integration",
		Version: "1.0",
	})
	e.SetSubjectTestSuiteRun(&api.Reference{
		ID:     "test-suite-111",
		Source: "testkube-dev-123",
	})
	e.SetChainId("")
}

func exampleTaskRunStartedEvent(e *apiv04.TaskRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUrl(testSubjecturl)
	e.SetSubjectPipelineRun(&api.Reference{ID: testSubjectId})
}

func exampleTestCaseRunFinishedEvent(e *apiv04.TestCaseRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunID)
	e.SetSubjectId(testTestRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseFinished)
	e.SetSubjectOutcome(testTestOutcome)
}

func exampleTestCaseRunQueuedEvent(e *apiv04.TestCaseRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunID)
	e.SetSubjectId(testTestRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseQueued)
	e.SetSubjectTrigger(testTestTriggerQueued)
}

func exampleTestCaseRunStartedEvent(e *apiv04.TestCaseRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseStarted)
	e.SetSubjectTrigger(testTestTriggerStarted)
}

func exampleTestSuiteRunFinishedEvent(e *apiv04.TestSuiteRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteFinished)
	e.SetSubjectOutcome(testTestSuiteOutcome)
	e.SetSubjectSeverity(testTestSuiteSeverity)
	e.SetSubjectReason(testTestSuiteReason)
}

func exampleTestSuiteRunStartedEvent(e *apiv04.TestSuiteRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteStarted)
	e.SetSubjectTrigger(testTestSuiteTriggerStarted)
}

func exampleTestSuiteRunQueuedEvent(e *apiv04.TestSuiteRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunID)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteQueued)
	e.SetSubjectTrigger(testTestSuiteTriggerQueued)
}

func exampleTestOutputPublishedEvent(e *apiv04.TestOutputPublishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestOutputSubjectID)
	e.SetSubjectSource(testTestOutputSubjectSource)
	e.SetSubjectOutputType(testTestOutputOutputType)
	e.SetSubjectFormat(testTestOutputFormat)
	e.SetSubjectTestCaseRun(testTestCaseRun)
}

func exampleTicketClosedEvent(e *apiv04.TicketClosedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTicketID)
	e.SetSource(testTicketSource)
	e.SetSubjectSource(testTicketSource)
	e.SetSubjectAssignees(testTicketAssignees)
	e.SetSubjectCreator(testTicketCreator)
	e.SetSubjectGroup(testTicketGroup)
	e.SetSubjectLabels(testTicketLabels)
	e.SetSubjectMilestone(testTicketMilestone)
	e.SetSubjectPriority(testTicketPriority)
	e.SetSubjectResolution(testTicketResolution)
	e.SetSubjectSummary(testTicketSummary)
	e.SetSubjectTicketType(testTicketType)
	e.SetSubjectUpdatedBy(testTicketUpdatedBy)
	e.SetSubjectUri(testTicketURI)
	e.SetChainId("")
}

func exampleTicketCreatedEvent(e *apiv04.TicketCreatedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTicketID)
	e.SetSource(testTicketSource)
	e.SetSubjectSource(testTicketSource)
	e.SetSubjectAssignees(testTicketAssignees)
	e.SetSubjectCreator(testTicketCreator)
	e.SetSubjectGroup(testTicketGroup)
	e.SetSubjectLabels(testTicketLabels)
	e.SetSubjectMilestone(testTicketMilestone)
	e.SetSubjectPriority(testTicketPriority)
	e.SetSubjectSummary(testTicketSummary)
	e.SetSubjectTicketType(testTicketType)
	e.SetSubjectUri(testTicketURI)
	e.SetChainId("")
}

func exampleTicketUpdatedEvent(e *apiv04.TicketUpdatedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTicketID)
	e.SetSource(testTicketSource)
	e.SetSubjectSource(testTicketSource)
	e.SetSubjectAssignees(testTicketAssignees)
	e.SetSubjectCreator(testTicketCreator)
	e.SetSubjectGroup(testTicketGroup)
	e.SetSubjectLabels(testTicketLabels)
	e.SetSubjectMilestone(testTicketMilestone)
	e.SetSubjectPriority(testTicketPriority)
	e.SetSubjectSummary(testTicketSummary)
	e.SetSubjectTicketType(testTicketType)
	e.SetSubjectUpdatedBy(testTicketUpdatedBy)
	e.SetSubjectUri(testTicketURI)
	e.SetChainId("")
}

func exampleCustomTypeEvent(e *apiv04.CustomTypeEvent) {
	// Set example specific fields
	// Set the type to dev.cdeventsx.mytool-resource.created.0.1.0
	e.SetEventType(testCustomEventType)
	e.SetSubjectContent(testCustomContent)
	e.SetSchemaUri(testCustomSchemaID2)
	e.SetSubjectId("pkg:resource/name@234fd47e07d1004f0aed9c")
	e.SetChainId("6ca3f9c5-1cef-4ce0-861c-2456a69cf137")
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
			// Check that the custom schema for the produced event what's excepted
			producedSchema, err := produced.GetCustomSchema()
			if err != nil {
				t.Errorf("failed to obtain the produced event custom schema: %v", err)
			}
			if d := cmp.Diff(producedSchema.ID, produced.GetSchemaUri()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			consumed, err := apiv04.NewFromJSONBytes(exampleConsumed)
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
			// Check v04+ attributes
			if d := cmp.Diff(consumed.GetChainId(), produced.GetChainId()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(consumed.GetSchemaUri(), produced.GetSchemaUri()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(consumed.GetLinks(), produced.GetLinks()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Coverage for GetCustomSchema
			consumedSchema, err := consumed.GetCustomSchema()
			if err != nil {
				t.Errorf("failed to obtain the consumed event custom schema: %v", err)
			}
			if d := cmp.Diff(consumedSchema.ID, producedSchema.ID); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check the case of no custom schema
			produced.SetSchemaUri("")
			producedSchema, err = produced.GetCustomSchema()
			if producedSchema != nil || err != nil {
				t.Errorf("expected nil schema and error when schema is not set, got schema %v, error %v", producedSchema, err)
			}
			// Check the case of custom schema missing from the DB
			notFoundSchema := "https://this.is.not.found/in/the/db"
			produced.SetSchemaUri(notFoundSchema)
			producedSchema, err = produced.GetCustomSchema()
			if err == nil {
				t.Errorf("expected an error when schema is not found, got schema %v, error %v", producedSchema, err)
			}
			expectedError := fmt.Sprintf("schema with id %s could not be found", notFoundSchema)
			if !strings.HasPrefix(err.Error(), expectedError) {
				t.Errorf("error %s does not start with the expected prefix %s", err.Error(), expectedError)
			}
		})
	}
}
