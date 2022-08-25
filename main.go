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
	"encoding/json"
	"fmt"
	"os"

	"github.com/cdevents/sdk-go/pkg/api"
	"github.com/invopop/jsonschema"
)

var (
	allEvents = map[string]api.CDEvent{
		"pipelinerunqueued":   &api.PipelineRunQueuedEvent{},
		"pipelinerunstarted":  &api.PipelineRunStartedEvent{},
		"pipelinerunfinished": &api.PipelineRunFinishedEvent{},
		"taskrunstarted":      &api.TaskRunStartedEvent{},
		"taskrunfinished":     &api.TaskRunFinishedEvent{},
		"changecreated":       &api.ChangeCreatedEvent{},
		"changeupdated":       &api.ChangeUpdatedEvent{},
		"changereviewed":      &api.ChangeReviewedEvent{},
		"changemerged":        &api.ChangeMergedEvent{},
		"changeabandoned":     &api.ChangeAbandonedEvent{},
		"repositorycreated":   &api.RepositoryCreatedEvent{},
		"repositorymodified":  &api.RepositoryModifiedEvent{},
		"repositorydeleted":   &api.RepositoryDeletedEvent{},
		"branchcreated":       &api.BranchCreatedEvent{},
		"branchdeleted":       &api.BranchDeletedEvent{},
		"testsuitestarted":    &api.TestSuiteStartedEvent{},
		"testsuitefinished":   &api.TestSuiteFinishedEvent{},
		"testcasequeued":      &api.TestCaseQueuedEvent{},
		"testcasestarted":     &api.TestCaseStartedEvent{},
		"testcasefinished":    &api.TestCaseFinishedEvent{},
		"buildqueued":         &api.BuildQueuedEvent{},
		"buildstarted":        &api.BuildStartedEvent{},
		"buildfinished":       &api.BuildFinishedEvent{},
		"artifactpackaged":    &api.ArtifactPackagedEvent{},
		"artifactpublished":   &api.ArtifactPublishedEvent{},
		"environmentcreated":  &api.EnvironmentCreatedEvent{},
		"environmentmodified": &api.EnvironmentModifiedEvent{},
		"environmentdeleted":  &api.EnvironmentDeletedEvent{},
		"servicedeployed":     &api.ServiceDeployedEvent{},
		"serviceupgraded":     &api.ServiceUpgradedEvent{},
		"servicerolledback":   &api.ServiceRolledbackEvent{},
		"serviceremoved":      &api.ServiceRemovedEvent{},
		"servicepublished":    &api.ServicePublishedEvent{},
	}
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// Setup a reflector
	id := jsonschema.EmptyID
	id.Add(fmt.Sprintf("https://cdevents.dev/%s/schema", api.CDEventsSpecVersion))
	reflector := jsonschema.Reflector{
		BaseSchemaID:   id,
		DoNotReference: true,
	}
	for filename, eventType := range allEvents {
		f, err := os.Create(fmt.Sprintf("jsonschema/%s.json", filename))
		panicOnError(err)
		defer f.Close()
		s := reflector.Reflect(eventType)
		data, err := json.MarshalIndent(s, "", "  ")
		panicOnError(err)
		_, err = f.WriteString(string(data))
		panicOnError(err)
	}
}
