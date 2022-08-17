#!/usr/bin/env bash

# This script can be used to add a new event with empty data model
# based on the subject and predicate names

# Usage:
# ./add-event.sh <subject> <predicate>
#
# Both subect and predicate should be give in camelcase

BASE_DIR="$( cd "$( dirname "$0" )/.." >/dev/null 2>&1 && pwd )"

set -e

SUBJECT=$1
PREDICATE=$2
SUBJECT_LOWER_CAMEL=${SUBJECT,}
SUBJECT_UPPER_CAMEL=${SUBJECT^}
SUBJECT_LOWER=${SUBJECT,,}
PREDICATE_LOWER_CAMEL=${PREDICATE,}
PREDICATE_UPPER_CAMEL=${PREDICATE^}
PREDICATE_LOWER=${PREDICATE,,}

cat > "${BASE_DIR}/pkg/api/${SUBJECT_LOWER}${PREDICATE_LOWER}.go" << EOF
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
	"time"
)

const (
	// ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL} event
	${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}EventV1 CDEventType = "dev.cdevents.${SUBJECT_LOWER}.${PREDICATE_LOWER}.v1"
)

type ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}SubjectContent struct {}

type ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Subject struct {
	SubjectBase
	Content ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}SubjectContent \`json:"content"\`
}

func (sc ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Subject) GetEventType() CDEventType {
	return ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}EventV1
}

func (sc ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Subject) GetSubjectType() SubjectType {
	return ${SUBJECT_UPPER_CAMEL}SubjectType
}

type ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event struct {
	Context Context                    \`json:"context"\`
	Subject ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Subject \`json:"subject"\`
}

// CDEventsReader implementation

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetType() CDEventType {
	return ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}EventV1
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetId() string {
	return e.Context.Id
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetSource() string {
	return e.Context.Source
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetSubjectId() string {
	return e.Subject.Id
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) SetId(id string) {
	e.Context.Id = id
}

func (e *${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func new${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event() CDEvent {
	return &${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Event{
		Context: Context{
			Type:    ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}EventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ${SUBJECT_UPPER_CAMEL}${PREDICATE_UPPER_CAMEL}Subject{},
	}
}
EOF

echo "Created ${BASE_DIR}/pkg/api/${SUBJECT_LOWER}${PREDICATE_LOWER}.go"