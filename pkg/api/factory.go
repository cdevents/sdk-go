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
	"time"

	"github.com/google/uuid"
)

var timeNow = time.Now
var uuidNewRandom = uuid.NewRandom

func initCDEvent(e CDEvent) (CDEvent, error) {
	eventUUID, err := uuidNewRandom()
	if err != nil {
		return nil, err
	}
	e.SetId(fmt.Sprintf("%v", eventUUID))
	e.SetTimestamp(timeNow())
	return e, nil
}

// NewCDEvent produces a CDEvent by type
// This function can be used by users but it's meant mainly for testing purposes
func NewCDEvent(eventType CDEventType) (CDEvent, error) {
	switch eventType {
	case PipelineRunQueuedEventV1:
		return NewPipelineRunQueuedEvent()
	case PipelineRunStartedEventV1:
		return NewPipelineRunStartedEvent()
	case PipelineRunFinishedEventV1:
		return NewPipelineRunFinishedEvent()
	case TaskRunStartedEventV1:
		return NewTaskRunStartedEvent()
	case TaskRunFinishedEventV1:
		return NewTaskRunFinishedEvent()
	case RepositoryCreatedEventV1:
		return NewRepositoryCreatedEvent()
	case RepositoryModifiedEventV1:
		return NewRepositoryModifiedEvent()
	case RepositoryDeletedEventV1:
		return NewRepositoryDeletedEvent()
	case BranchCreatedEventV1:
		return NewBranchCreatedEvent()
	case BranchDeletedEventV1:
		return NewBranchDeletedEvent()
	case ChangeCreatedEventV1:
		return NewChangeCreatedEvent()
	case ChangeUpdatedEventV1:
		return NewChangeUpdatedEvent()
	case ChangeReviewedEventV1:
		return NewChangeReviewedEvent()
	case ChangeMergedEventV1:
		return NewChangeMergedEvent()
	case ChangeAbandonedEventV1:
		return NewChangeAbandonedEvent()
	case BuildQueuedEventV1:
		return NewBuildQueuedEvent()
	case BuildStartedEventV1:
		return NewBuildStartedEvent()
	case BuildFinishedEventV1:
		return NewBuildFinishedEvent()
	case TestCaseQueuedEventV1:
		return NewTestCaseQueuedEvent()
	case TestCaseStartedEventV1:
		return NewTestCaseStartedEvent()
	case TestCaseFinishedEventV1:
		return NewTestCaseFinishedEvent()
	case TestSuiteStartedEventV1:
		return NewTestSuiteStartedEvent()
	case TestSuiteFinishedEventV1:
		return NewTestSuiteFinishedEvent()
	case ArtifactPackagedEventV1:
		return NewArtifactPackagedEvent()
	case ArtifactPublishedEventV1:
		return NewArtifactPublishedEvent()
	// case EnvironmentCreatedEventV1:
	// 	return NewEnvironmentCreatedEvent()
	// case EnvironmentModifiedEventV1:
	// 	return NewEnvironmentModifiedEvent()
	// case EnvironmentDeletedEventV1:
	// 	return NewEnvironmentDeletedEvent()
	// case ServiceDeployedEventV1:
	// 	return NewServiceDeployedEvent()
	// case ServiceUpgradedEventV1:
	// 	return NewServiceUpgradedEvent()
	// case ServiceRolledbackEventV1:
	// 	return NewServiceRolledbackEvent()
	// case ServiceRemovedEventV1:
	// 	return NewServiceRemovedEvent()
	// case ServicePublishedEventV1:
	// 	return NewServicePublishedEvent()
	default:
		return nil, fmt.Errorf("event %v not supported", eventType)
	}
}
