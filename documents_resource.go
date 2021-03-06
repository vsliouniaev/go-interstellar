// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2019-present, Jet.com, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License."

package interstellar

// DocumentProperties are the well-known properties that may exist on a Document resources
type DocumentProperties struct {
	ID          string `json:"id"`
	ETag        string `json:"_etag"`
	ResourceID  string `json:"_rid"`
	Timestamp   int64  `json:"_ts"`
	Self        string `json:"_self"`
	Attachments string `json:"_attachments"`
}

// DocumentIndexingDirective determines if a document create/update should be indexed
type DocumentIndexingDirective string

const (
	// DocumentIndexingInclude specifies that the document should be included in the collection index.
	DocumentIndexingInclude = DocumentIndexingDirective("Include")
	// DocumentIndexingExclude specifies that the document should be excluded from the collection index.
	DocumentIndexingExclude = DocumentIndexingDirective("Exclude")
)
