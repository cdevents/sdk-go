// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

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

import "github.com/cdevents/sdk-go/pkg/api"

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type testData struct {
	TestValues []map[string]string `json:"testValues"`
}

var (
	// Examples Data
	testSource               = "/event/source/123"
	testSubjectId            = "mySubject123"
	testValue                = "testValue"
	testArtifactId           = "pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
	testDataJson             = testData{TestValues: []map[string]string{{"k1": "v1"}, {"k2": "v2"}}}
	testDataJsonUnmarshalled = map[string]any{
		"testValues": []any{map[string]any{"k1": string("v1")}, map[string]any{"k2": string("v2")}},
	}
	testDataXml  = []byte("<xml>testData</xml>")
	testChangeId = "myChange123"

	// V04+ Examples Data
	testChainId   = "4c8cb7dd-3448-41de-8768-eec704e2829b"
	testLinks     api.EmbeddedLinksArray
	testContextId = "5328c37f-bb7e-4bb7-84ea-9f5f85e4a7ce"
	testSchemaUri = "https://myorg.com/schema/custom"
)

func init() {
	// Set up test links
	tags := api.Tags{
		"foo1": "bar",
		"foo2": "bar",
	}
	reference := api.EventReference{
		ContextId: testContextId,
	}
	elr := api.NewEmbeddedLinkRelation()
	elr.SetTags(tags)
	elr.SetLinkKind("TRIGGER")
	elr.SetTarget(reference)
	elp := api.NewEmbeddedLinkPath()
	elp.SetTags(tags)
	elp.SetFrom(reference)
	ele := api.NewEmbeddedLinkEnd()
	ele.SetTags(tags)
	ele.SetFrom(reference)
	testLinks = api.EmbeddedLinksArray{
		elr, elp, ele,
	}
}

func setContext(event api.CDEventWriter, subjectId string) {
	event.SetSource(testSource)
	event.SetSubjectId(subjectId)
}

func setContextV04(event api.CDEventWriterV04, chainId, schemaUri bool) {
	if chainId {
		event.SetChainId(testChainId)
	}
	if schemaUri {
		event.SetSchemaUri(testSchemaUri)
	}
	event.SetLinks(testLinks)
}
