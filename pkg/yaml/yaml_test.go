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

import "testing"

func TestConvertToHelmRawValues(t *testing.T) {
	type args struct {
		infile string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{"simple-single", args{infile: simpleSingleIn}, simpleSingleOut, false},
		{"bad-format", args{badFormatIn}, "", true},
		{"simple-multi", args{simpleMultiIn}, simpleMultiOut, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := ConvertToHelmRawValues(tt.args.infile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToHelmRawValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("ConvertToHelmRawValues() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

const simpleSingleIn string = `a: Easy!
b:
  c: 2
  d: [3, 4]
`

const simpleSingleOut string = `resources:
    - a: Easy!
      b:
        c: 2
        d:
            - 3
            - 4
`

const badFormatIn string = `a Easy!
b:
  c: 2
  d: [3, 4]
`

const simpleMultiIn string = `a: Easy!
b:
  c: 2
  d: [3, 4]
---
a: Easy!
b:
  c: 2
  d: [3, 4]
`

const simpleMultiOut string = `resources:
    - a: Easy!
      b:
        c: 2
        d:
            - 3
            - 4
    - a: Easy!
      b:
        c: 2
        d:
            - 3
            - 4
`
