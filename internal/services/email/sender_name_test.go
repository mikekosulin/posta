/*
 * Copyright 2026 Jonas Kaninda
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
 *
 */

package email

import "testing"

func TestResolveSenderWithDefault(t *testing.T) {
	cases := []struct {
		name        string
		from        string
		defaultName string
		want        string
	}{
		{"bare address gets default name", "hello@example.com", "Acme", "\"Acme\" <hello@example.com>"},
		{"existing display name is kept", "Support <hello@example.com>", "Acme", "Support <hello@example.com>"},
		{"empty default leaves bare address", "hello@example.com", "", "hello@example.com"},
		{"whitespace default leaves bare address", "hello@example.com", "   ", "hello@example.com"},
		{"unparseable input is untouched", "not-an-email", "Acme", "not-an-email"},
		{"name needing quoting is escaped", "hello@example.com", "Acme, Inc.", "\"Acme, Inc.\" <hello@example.com>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := resolveSenderWithDefault(tc.from, tc.defaultName); got != tc.want {
				t.Errorf("resolveSenderWithDefault(%q, %q) = %q, want %q", tc.from, tc.defaultName, got, tc.want)
			}
		})
	}
}
