/*
 * Copyright © 2021 Andrew Roth <roth.andy@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yaml

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// GenericDocument is used to decode any YAML document without having to specify its schema
type GenericDocument interface{}

// RawValues is the struct that gets marshaled to YAML that the incubator/raw helm chart can use
type RawValues struct {
	Resources []GenericDocument `yaml:"resources"`
}

// ConvertToHelmRawValues takes in a YAML string, converts it to a format that
// the incubator/raw helm chart can use, and outputs the converted string
func ConvertToHelmRawValues(infile string) (out string, err error) {
	// Build array of GenericDocuments
	var docs []GenericDocument
	decoder := yaml.NewDecoder(strings.NewReader(infile))
	for {
		var doc GenericDocument
		decodeErr := decoder.Decode(&doc)
		if decodeErr != nil {
			if decodeErr.Error() == "EOF" {
				break
			} else {
				out = ""
				err = decodeErr
				return
			}
		}
		docs = append(docs, doc)
	}

	// Convert to RawValues
	var rv RawValues
	rv.Resources = docs
	// We can safely ignore the error here since the only values that would cause
	// Marshal to fail would already be caught by the Decoder above
	rvBytes, _ := yaml.Marshal(rv)

	// Return
	out = string(rvBytes)
	return
}
