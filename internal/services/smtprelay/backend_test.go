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

package smtprelay

import (
	"errors"
	"testing"

	"github.com/emersion/go-smtp"
)

func TestMapSendError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantNil  bool
		wantCode int
	}{
		{name: "nil error accepted", err: nil, wantNil: true},
		{name: "rate limit", err: errors.New("rate_limit: too many requests"), wantCode: 452},
		{name: "domain verification", err: errors.New("domain_verification: sender domain not verified"), wantCode: 550},
		{name: "other error", err: errors.New("boom"), wantCode: 451},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapSendError(tt.err, "127.0.0.1:1234")
			if tt.wantNil {
				if got != nil {
					t.Fatalf("expected nil error, got %v", got)
				}
				return
			}
			smtpErr, ok := got.(*smtp.SMTPError)
			if !ok {
				t.Fatalf("expected *smtp.SMTPError, got %T", got)
			}
			if smtpErr.Code != tt.wantCode {
				t.Fatalf("expected code %d, got %d", tt.wantCode, smtpErr.Code)
			}
		})
	}
}
