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

func NewCDEvent(eventType CDEventType) (CDEvent, error) {
	switch eventType {
	case PipelineRunQueuedEventV1:
		e := newPipelineRunQueuedEvent()
		return initCDEvent(e)
	case PipelineRunStartedEventV1:
		e := newPipelineRunStartedEvent()
		return initCDEvent(e)
	case PipelineRunFinishedEventV1:
		e := newPipelineRunFinishedEvent()
		return initCDEvent(e)
	case TaskRunStartedEventV1:
		e := newTaskRunStartedEvent()
		return initCDEvent(e)
	case TaskRunFinishedEventV1:
		e := newTaskRunFinishedEvent()
		return initCDEvent(e)
	// case RepositoryCreatedEventV1:
	// 	e := newRepositoryCreatedEvent()
	// 	return initCDEvent(e)
	// case RepositoryModifiedEventV1:
	// 	e := newRepositoryModifiedEvent()
	// 	return initCDEvent(e)
	// case RepositoryDeletedEventV1:
	// 	e := newRepositoryDeletedEvent()
	// 	return initCDEvent(e)
	// case BranchCreatedEventV1:
	// 	e := newBranchCreatedEvent()
	// 	return initCDEvent(e)
	// case BranchDeletedEventV1:
	// 	e := newBranchDeletedEvent()
	// 	return initCDEvent(e)
	case ChangeCreatedEventV1:
		e := newChangeCreatedEvent()
		return initCDEvent(e)
	case ChangeUpdatedEventV1:
		e := newChangeUpdatedEvent()
		return initCDEvent(e)
	case ChangeReviewedEventV1:
		e := newChangeReviewedEvent()
		return initCDEvent(e)
	case ChangeMergedEventV1:
		e := newChangeMergedEvent()
		return initCDEvent(e)
	case ChangeAbandonedEventV1:
		e := newChangeAbandonedEvent()
		return initCDEvent(e)
	// case BuildQueuedEventV1:
	// 	e := newBuildQueuedEvent()
	// 	return initCDEvent(e)
	// case BuildStartedEventV1:
	// 	e := newBuildStartedEvent()
	// 	return initCDEvent(e)
	// case BuildFinishedEventV1:
	// 	e := newBuildFinishedEvent()
	// 	return initCDEvent(e)
	// case TestCaseQueuedEventV1:
	// 	e := newTestCaseQueuedEvent()
	// 	return initCDEvent(e)
	// case TestCaseStartedEventV1:
	// 	e := newTestCaseStartedEvent()
	// 	return initCDEvent(e)
	// case TestCaseFinishedEventV1:
	// 	e := newTestCaseFinishedEvent()
	// 	return initCDEvent(e)
	// case TestSuiteQueuedEventV1:
	// 	e := newTestSuiteQueuedEvent()
	// 	return initCDEvent(e)
	// case TestSuiteStartedEventV1:
	// 	e := newTestSuiteStartedEvent()
	// 	return initCDEvent(e)
	// case TestSuiteFinishedEventV1:
	// 	e := newTestSuiteFinishedEvent()
	// 	return initCDEvent(e)
	// case ArtifactPackagedEventV1:
	// 	e := newArtifactPackagedEvent()
	// 	return initCDEvent(e)
	// case ArtifactPublishedEventV1:
	// 	e := newArtifactPublishedEvent()
	// 	return initCDEvent(e)
	// case EnvironmentCreatedEventV1:
	// 	e := newEnvironmentCreatedEvent()
	// 	return initCDEvent(e)
	// case EnvironmentModifiedEventV1:
	// 	e := newEnvironmentModifiedEvent()
	// 	return initCDEvent(e)
	// case EnvironmentDeletedEventV1:
	// 	e := newEnvironmentDeletedEvent()
	// 	return initCDEvent(e)
	// case ServiceDeployedEventV1:
	// 	e := newServiceDeployedEvent()
	// 	return initCDEvent(e)
	// case ServiceUpgradedEventV1:
	// 	e := newServiceUpgradedEvent()
	// 	return initCDEvent(e)
	// case ServiceRolledbackEventV1:
	// 	e := newServiceRolledbackEvent()
	// 	return initCDEvent(e)
	// case ServiceRemovedEventV1:
	// 	e := newServiceRemovedEvent()
	// 	return initCDEvent(e)
	// case ServicePublishedEventV1:
	// 	e := newServicePublishedEvent()
	// 	return initCDEvent(e)
	default:
		return nil, fmt.Errorf("event %v not supported", eventType)
	}
}
