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
	"testing"

	"github.com/goposta/posta/internal/models"
)

func TestHashPassword(t *testing.T) {
	h1 := hashPassword("correct-horse-battery-staple")
	h2 := hashPassword("correct-horse-battery-staple")
	if h1 != h2 {
		t.Errorf("hashPassword() is not deterministic: %q != %q", h1, h2)
	}
}

func TestHashPasswordDiffers(t *testing.T) {
	h1 := hashPassword("password-one")
	h2 := hashPassword("password-two")
	if h1 == h2 {
		t.Errorf("hashPassword() produced same hash for different passwords: %q", h1)
	}
}

func TestVerifyPassword(t *testing.T) {
	cred := &models.SMTPCredential{PasswordHash: hashPassword("s3cret-pass")}

	if !VerifyPassword(cred, "s3cret-pass") {
		t.Error("VerifyPassword() = false, want true for correct password")
	}
	if VerifyPassword(cred, "wrong-pass") {
		t.Error("VerifyPassword() = true, want false for incorrect password")
	}
}
