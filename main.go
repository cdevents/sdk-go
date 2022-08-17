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

package main

import (
	"fmt"

	"github.com/cdevents/sdk-go/pkg/api"
)

var (
	prq *api.PipelineRunQueuedEvent
	prs *api.PipelineRunStartedEvent
	prf *api.PipelineRunFinishedEvent

	source    = "TestAsCloudEvent"
	subjectid = "mySubject123"
	pipeline  = "myPipeline"
	url       = "https://www.example.com/mySubject123"
	outcome   = api.PipelineRunOutcomeFailed
	errors    = "Something went wrong\nWith some more details"
)

func makecde(eventType api.CDEventType) api.CDEvent {
	event, _ := api.NewCDEvent(eventType)
	event.SetSource(source)
	event.SetSubjectId(subjectid)
	return event
}

func main() {

	e := makecde(api.PipelineRunQueuedEventV1)
	prq, _ = e.(*api.PipelineRunQueuedEvent)
	prq.SetSubjectPipelineName(pipeline)
	prq.SetSubjectURL(url)

	e = makecde(api.PipelineRunStartedEventV1)
	prs, _ = e.(*api.PipelineRunStartedEvent)
	prs.SetSubjectPipelineName(pipeline)
	prs.SetSubjectURL(url)

	e = makecde(api.PipelineRunFinishedEventV1)
	prf, _ = e.(*api.PipelineRunFinishedEvent)
	prf.SetSubjectPipelineName(pipeline)
	prf.SetSubjectURL(url)
	prf.SetSubjectOutcome(outcome)
	prf.SetSubjectErrors(errors)

	ej, _ := api.AsJsonString(prq)
	fmt.Printf("%v\n", ej)
	ej, _ = api.AsJsonString(prs)
	fmt.Printf("%v\n", ej)
	ej, _ = api.AsJsonString(prf)
	fmt.Printf("%v\n", ej)
}
