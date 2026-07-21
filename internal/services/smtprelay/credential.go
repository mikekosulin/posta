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
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/goposta/posta/internal/models"
	"github.com/goposta/posta/internal/storage/repositories"
)

// SMTPUsernamePrefix identifies a generated SMTP Relay username.
const SMTPUsernamePrefix = "smtp_"

type CredentialService struct {
	repo *repositories.SMTPCredentialRepository
}

func NewCredentialService(repo *repositories.SMTPCredentialRepository) *CredentialService {
	return &CredentialService{repo: repo}
}

func (s *CredentialService) GenerateCredential(userID, workspaceID uint, name string, allowedIPs []string) (username, password string, cred *models.SMTPCredential, err error) {
	userBytes := make([]byte, 8)
	if _, err = rand.Read(userBytes); err != nil {
		return "", "", nil, fmt.Errorf("failed to generate username: %w", err)
	}
	username = SMTPUsernamePrefix + hex.EncodeToString(userBytes)

	passBytes := make([]byte, 32)
	if _, err = rand.Read(passBytes); err != nil {
		return "", "", nil, fmt.Errorf("failed to generate password: %w", err)
	}
	password = hex.EncodeToString(passBytes)

	cred = &models.SMTPCredential{
		WorkspaceID:  workspaceID,
		UserID:       userID,
		Name:         name,
		Username:     username,
		PasswordHash: hashPassword(password),
		AllowedIPs:   allowedIPs,
	}

	if err = s.repo.Create(cred); err != nil {
		return "", "", nil, err
	}

	return username, password, cred, nil
}

// VerifyPassword reports whether password matches cred's stored hash.
func VerifyPassword(cred *models.SMTPCredential, password string) bool {
	return cred.PasswordHash == hashPassword(password)
}

func hashPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return hex.EncodeToString(h[:])
}
