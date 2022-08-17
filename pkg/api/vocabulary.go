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

const (

	// Repository events
	RepositoryCreatedEventV1  CDEventType = "dev.cdevents.repository.created.v1"
	RepositoryModifiedEventV1 CDEventType = "dev.cdevents.repository.modified.v1"
	RepositoryDeletedEventV1  CDEventType = "dev.cdevents.repository.deleted.v1"

	BranchCreatedEventV1 CDEventType = "dev.cdevents.repository.branch.created.v1"
	BranchDeletedEventV1 CDEventType = "dev.cdevents.repository.branch.deleted.v1"

	// Change Events
	ChangeCreatedEventV1   CDEventType = "dev.cdevents.repository.change.created.v1"
	ChangeUpdatedEventV1   CDEventType = "dev.cdevents.repository.change.updated.v1"
	ChangeReviewedEventV1  CDEventType = "dev.cdevents.repository.change.reviewed.v1"
	ChangeMergedEventV1    CDEventType = "dev.cdevents.repository.change.merged.v1"
	ChangeAbandonedEventV1 CDEventType = "dev.cdevents.repository.change.abandoned.v1"

	// Build Events
	BuildStartedEventV1  CDEventType = "dev.cdevents.build.started.v1"
	BuildQueuedEventV1   CDEventType = "dev.cdevents.build.queued.v1"
	BuildFinishedEventV1 CDEventType = "dev.cdevents.build.finished.v1"

	// Test Events
	TestCaseStartedEventV1  CDEventType = "dev.cdevents.test.case.started.v1"
	TestCaseQueuedEventV1   CDEventType = "dev.cdevents.test.case.queued.v1"
	TestCaseFinishedEventV1 CDEventType = "dev.cdevents.test.case.finished.v1"

	TestSuiteStartedEventV1  CDEventType = "dev.cdevents.test.suite.started.v1"
	TestSuiteQueuedEventV1   CDEventType = "dev.cdevents.test.suite.queued.v1"
	TestSuiteFinishedEventV1 CDEventType = "dev.cdevents.test.suite.finished.v1"

	// Artifact Events
	ArtifactPackagedEventV1  CDEventType = "dev.cdevents.artifact.packaged.v1"
	ArtifactPublishedEventV1 CDEventType = "dev.cdevents.artifact.published.v1"

	// Environment Events
	EnvironmentCreatedEventV1  CDEventType = "dev.cdevents.environment.created.v1"
	EnvironmentModifiedEventV1 CDEventType = "dev.cdevents.environment.modified.v1"
	EnvironmentDeletedEventV1  CDEventType = "dev.cdevents.environment.deleted.v1"

	// Service Events
	ServiceDeployedEventV1   CDEventType = "dev.cdevents.service.deployed.v1"
	ServiceUpgradedEventV1   CDEventType = "dev.cdevents.service.upgraded.v1"
	ServiceRolledbackEventV1 CDEventType = "dev.cdevents.service.rolledback.v1"
	ServiceRemovedEventV1    CDEventType = "dev.cdevents.service.removed.v1"
)
