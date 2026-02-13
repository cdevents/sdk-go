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

package v05_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cdevents/sdk-go/pkg/api"
	apiv05 "github.com/cdevents/sdk-go/pkg/api/v05"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

const (
	examplesFolder = "spec-v0.5/conformance"
	customExample  = "spec-v0.5/custom/conformance.json"
)

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
	testTestCaseStarted = &apiv05.TestCaseRunStartedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseFinished = &apiv05.TestCaseRunFinishedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestCaseQueued = &apiv05.TestCaseRunQueuedSubjectContentTestCase{
		Id: "92834723894", Name: "Login Test", Type: "integration", Version: "1.0"}
	testTestTriggerQueued = &apiv05.TestCaseRunQueuedSubjectContentTrigger{
		Type: "schedule"}
	testTestTriggerStarted = &apiv05.TestCaseRunStartedSubjectContentTrigger{
		Type: "schedule"}
	testTestOutcome             = "success"
	testTestOutputSubjectId     = "testrunreport-12123"
	testTestOutputSubjectSource = "/event/source/testrunreport-12123"
	testTestOutputFormat        = "video/quicktime"
	testTestOutputOutputType    = "video"
	testTestCaseRun             = &api.Reference{Id: testTestRunId, Source: "testkube-dev-123"}
	testTestSuiteRunId          = "myTestSuiteRun123"
	testTestSuiteStarted        = &apiv05.TestSuiteRunStartedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteQueued = &apiv05.TestSuiteRunQueuedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteFinished = &apiv05.TestSuiteRunFinishedSubjectContentTestSuite{
		Id: "92834723894", Name: "Auth TestSuite", Version: "1.0"}
	testTestSuiteOutcome        = "failure"
	testTestSuiteReason         = "Host 123.34.23.32 not found"
	testTestSuiteSeverity       = "critical"
	testTestSuiteTriggerQueued  = &apiv05.TestSuiteRunQueuedSubjectContentTrigger{Type: "pipeline"}
	testTestSuiteTriggerStarted = &apiv05.TestSuiteRunStartedSubjectContentTrigger{Type: "pipeline"}
	testSubjectUser             = "mybot-myapp"
	testSbomUri                 = "https://sbom.repo/myorg/234fd47e07d1004f0aed9c.sbom"
	testChangeDescription       = "This PR address a bug from a recent PR"
	testTicketId                = "ticket-123"
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
	testTicketUri               = "https://example.issues.com/ticket123"
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
	testCustomSchemaId           = "https://myorg.com/schema/custom"
	testCustomSchemaId2          = "https://myorg.com/schema/mytool"
	testCustomSchemaJsonTemplate = `{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"$id": "%s",
	"additionalProperties": true,
	"type": "object"
}`
	testCustomSchemaJson  = fmt.Sprintf(testCustomSchemaJsonTemplate, testCustomSchemaId)
	testCustomSchema2Json = fmt.Sprintf(testCustomSchemaJsonTemplate, testCustomSchemaId2)
	testCustomSchemas     = map[string][]byte{
		testCustomSchemaId:  []byte(testCustomSchemaJson),
		testCustomSchemaId2: []byte(testCustomSchema2Json),
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

	for _, event := range apiv05.CDEventsTypes {
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

func exampleArtifactPackagedEvent(e *apiv05.ArtifactPackagedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectChange(&api.Reference{Id: testChangeId, Source: testChangeSource})
	e.SetSubjectSbom(&api.ArtifactPackagedSubjectContentSbomV0_3_0{
		Uri: testSbomUri,
	})
}

func exampleArtifactPublishedEvent(e *apiv05.ArtifactPublishedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectUser(testSubjectUser)
	e.SetSubjectSbom(&api.ArtifactPublishedSubjectContentSbomV0_3_0{
		Uri: testSbomUri,
	})
}

func exampleArtifactSignedEvent(e *apiv05.ArtifactSignedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectSignature(testSignature)
}

func exampleArtifactDeletedEvent(e *apiv05.ArtifactDeletedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectUser(testSubjectUser)
	e.SetChainId("")
}

func exampleArtifactDownloadedEvent(e *apiv05.ArtifactDownloadedEvent) {
	// Set example specific fields
	setContext(e, testArtifactSubjectId)
	e.SetSubjectUser(testSubjectUser)
	e.SetChainId("")
}

func exampleBranchCreatedEvent(e *apiv05.BranchCreatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBranchDeletedEvent(e *apiv05.BranchDeletedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleBuildFinishedEvent(e *apiv05.BuildFinishedEvent) {
	// Set example specific fields
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleBuildQueuedEvent(e *apiv05.BuildQueuedEvent) {
	// Set example specific fields
}

func exampleBuildStartedEvent(e *apiv05.BuildStartedEvent) {
	// Set example specific fields
}

func exampleChangeAbandonedEvent(e *apiv05.ChangeAbandonedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeCreatedEvent(e *apiv05.ChangeCreatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
	e.SetSubjectDescription(testChangeDescription)
}

func exampleChangeMergedEvent(e *apiv05.ChangeMergedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeReviewedEvent(e *apiv05.ChangeReviewedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleChangeUpdatedEvent(e *apiv05.ChangeUpdatedEvent) {
	// Set example specific fields
	e.SetSubjectRepository(testRepoReference)
}

func exampleEnvironmentCreatedEvent(e *apiv05.EnvironmentCreatedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUri(testEnvironmentUrl)
}

func exampleEnvironmentDeletedEvent(e *apiv05.EnvironmentDeletedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
}

func exampleEnvironmentModifiedEvent(e *apiv05.EnvironmentModifiedEvent) {
	// Set example specific fields
	e.SetSubjectName(testEnvironmentName)
	e.SetSubjectUri(testEnvironmentUrl)
}

func exampleIncidentDetectedEvent(e *apiv05.IncidentDetectedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectId)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId(testArtifactId)
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time above threshold of 100ms")
}

func exampleIncidentReportedEvent(e *apiv05.IncidentReportedEvent) {
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

func exampleIncidentResolvedEvent(e *apiv05.IncidentResolvedEvent) {
	// Set example specific fields
	e.SetSubjectId(testIncidentSubjectId)
	e.SetSource(testIncidentSource)
	e.SetSubjectSource(testIncidentSource)
	e.SetSubjectArtifactId("pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93439")
	e.SetSubjectService(testServiceReference)
	e.SetSubjectEnvironment(testEnvironmentReference)
	e.SetSubjectDescription("Response time restored below 100ms")
}

func examplePipelineRunFinishedEvent(e *apiv05.PipelineRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUri(testSubjecturl)
	e.SetSubjectOutcome(testPipelineOutcome)
	e.SetSubjectErrors(testPipelineErrors)
}

func examplePipelineRunQueuedEvent(e *apiv05.PipelineRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUri(testSubjecturl)
}

func examplePipelineRunStartedEvent(e *apiv05.PipelineRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectPipelineName(testPipeline)
	e.SetSubjectUri(testSubjecturl)
}

func exampleRepositoryCreatedEvent(e *apiv05.RepositoryCreatedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUri(testUrl)
	e.SetSubjectViewUrl(testViewUrl)
}

func exampleRepositoryDeletedEvent(e *apiv05.RepositoryDeletedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUri(testUrl)
	e.SetSubjectViewUrl(testViewUrl)
}

func exampleRepositoryModifiedEvent(e *apiv05.RepositoryModifiedEvent) {
	// Set example specific fields
	e.SetSubjectName(testRepo)
	e.SetSubjectOwner(testOwner)
	e.SetSubjectUri(testUrl)
	e.SetSubjectViewUrl(testViewUrl)
}

func exampleServiceDeployedEvent(e *apiv05.ServiceDeployedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServicePublishedEvent(e *apiv05.ServicePublishedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
}

func exampleServiceRemovedEvent(e *apiv05.ServiceRemovedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
}

func exampleServiceRolledbackEvent(e *apiv05.ServiceRolledbackEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleServiceUpgradedEvent(e *apiv05.ServiceUpgradedEvent) {
	// Set example specific fields
	e.SetSubjectEnvironment(&api.Reference{Id: testEnvironmentId})
	e.SetSubjectArtifactId(testArtifactId)
}

func exampleTaskRunFinishedEvent(e *apiv05.TaskRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUri(testSubjecturl)
	e.SetSubjectPipelineRun(&api.Reference{Id: testSubjectId})
	e.SetSubjectOutcome(testTaskOutcome)
	e.SetSubjectErrors(testTaskRunErrors)
}

func exampleTestCaseRunSkippedEvent(e *apiv05.TestCaseRunSkippedEvent) {
	// Set example specific fields
	e.SetSubjectId("myTestCaseRun123")
	e.SetSubjectEnvironment(&api.Reference{
		Id:     "dev",
		Source: "testkube-dev-123",
	})
	e.SetSubjectReason("Not running in this environment")
	e.SetSubjectTestCase(&api.TestCaseRunSkippedSubjectContentTestCaseV0_2_0{
		Id:      "92834723894",
		Name:    "Login Test",
		Type:    "integration",
		Version: "1.0",
	})
	e.SetSubjectTestSuiteRun(&api.Reference{
		Id:     "test-suite-111",
		Source: "testkube-dev-123",
	})
	e.SetChainId("")
}

func exampleTaskRunStartedEvent(e *apiv05.TaskRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectTaskName(testTaskName)
	e.SetSubjectUri(testSubjecturl)
	e.SetSubjectPipelineRun(&api.Reference{Id: testSubjectId})
}

func exampleTestCaseRunFinishedEvent(e *apiv05.TestCaseRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunId)
	e.SetSubjectId(testTestRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseFinished)
	e.SetSubjectOutcome(testTestOutcome)
}

func exampleTestCaseRunQueuedEvent(e *apiv05.TestCaseRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunId)
	e.SetSubjectId(testTestRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseQueued)
	e.SetSubjectTrigger(testTestTriggerQueued)
}

func exampleTestCaseRunStartedEvent(e *apiv05.TestCaseRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestRunId)
	e.SetSubjectId(testTestRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestCase(testTestCaseStarted)
	e.SetSubjectTrigger(testTestTriggerStarted)
}

func exampleTestSuiteRunFinishedEvent(e *apiv05.TestSuiteRunFinishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteFinished)
	e.SetSubjectOutcome(testTestSuiteOutcome)
	e.SetSubjectSeverity(testTestSuiteSeverity)
	e.SetSubjectReason(testTestSuiteReason)
}

func exampleTestSuiteRunStartedEvent(e *apiv05.TestSuiteRunStartedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteStarted)
	e.SetSubjectTrigger(testTestSuiteTriggerStarted)
}

func exampleTestSuiteRunQueuedEvent(e *apiv05.TestSuiteRunQueuedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestSuiteRunId)
	e.SetSubjectEnvironment(testTestEnvironment)
	e.SetSubjectTestSuite(testTestSuiteQueued)
	e.SetSubjectTrigger(testTestSuiteTriggerQueued)
}

func exampleTestOutputPublishedEvent(e *apiv05.TestOutputPublishedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTestOutputSubjectId)
	e.SetSubjectSource(testTestOutputSubjectSource)
	e.SetSubjectOutputType(testTestOutputOutputType)
	e.SetSubjectFormat(testTestOutputFormat)
	e.SetSubjectTestCaseRun(testTestCaseRun)
}

func exampleTicketClosedEvent(e *apiv05.TicketClosedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTicketId)
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
	e.SetSubjectUri(testTicketUri)
	e.SetChainId("")
}

func exampleTicketCreatedEvent(e *apiv05.TicketCreatedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTicketId)
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
	e.SetSubjectUri(testTicketUri)
	e.SetChainId("")
}

func exampleTicketUpdatedEvent(e *apiv05.TicketUpdatedEvent) {
	// Set example specific fields
	e.SetSubjectId(testTicketId)
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
	e.SetSubjectUri(testTicketUri)
	e.SetChainId("")
}

func exampleCustomTypeEvent(e *apiv05.CustomTypeEvent) {
	// Set example specific fields
	// Set the type to dev.cdeventsx.mytool-resource.created.0.1.0
	e.SetEventType(testCustomEventType)
	e.SetSubjectContent(testCustomContent)
	e.SetSchemaUri(testCustomSchemaId2)
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
			consumed, err := apiv05.NewFromJsonBytes(exampleConsumed)
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
