/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loader

import (
	"testing"
)

func TestIsRepoURL(t *testing.T) {

	testcases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "https://github.com/org/repo",
			expected: true,
		},
		{
			input:    "github.com/org/repo",
			expected: true,
		},
		{
			input:    "/github.com/org/repo",
			expected: false,
		},
		{
			input:    "/abs/path/to/file",
			expected: false,
		},
		{
			input:    "../relative",
			expected: false,
		},
		{
			input:    "git::https://gitlab.com/org/repo",
			expected: true,
		},
	}
	for _, tc := range testcases {
		actual := isRepoUrl(tc.input)
		if actual != tc.expected {
			t.Errorf("unexpected error: unexpected result %t for input %s", actual, tc.input)
		}
	}
}