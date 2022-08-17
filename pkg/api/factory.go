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
	default:
		return nil, fmt.Errorf("event %v not supported", eventType)
	}
}
