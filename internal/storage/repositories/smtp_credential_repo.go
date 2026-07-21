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

package repositories

import (
	"time"

	"github.com/goposta/posta/internal/models"
	"gorm.io/gorm"
)

type SMTPCredentialRepository struct {
	db *gorm.DB
}

func NewSMTPCredentialRepository(db *gorm.DB) *SMTPCredentialRepository {
	return &SMTPCredentialRepository{db: db}
}

func (r *SMTPCredentialRepository) Create(cred *models.SMTPCredential) error {
	return r.db.Create(cred).Error
}

func (r *SMTPCredentialRepository) FindByUsername(username string) (*models.SMTPCredential, error) {
	var cred models.SMTPCredential
	if err := r.db.Where("username = ?", username).First(&cred).Error; err != nil {
		return nil, err
	}
	return &cred, nil
}

func (r *SMTPCredentialRepository) FindByWorkspaceID(workspaceID uint, limit, offset int) ([]models.SMTPCredential, int64, error) {
	var creds []models.SMTPCredential
	var total int64

	r.db.Model(&models.SMTPCredential{}).Where("workspace_id = ?", workspaceID).Count(&total)

	if err := r.db.Where("workspace_id = ?", workspaceID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&creds).Error; err != nil {
		return nil, 0, err
	}
	return creds, total, nil
}

func (r *SMTPCredentialRepository) FindByID(id uint) (*models.SMTPCredential, error) {
	var cred models.SMTPCredential
	if err := r.db.First(&cred, id).Error; err != nil {
		return nil, err
	}
	return &cred, nil
}

func (r *SMTPCredentialRepository) UpdateLastUsed(id uint) error {
	now := time.Now()
	return r.db.Model(&models.SMTPCredential{}).Where("id = ?", id).Update("last_used_at", now).Error
}

func (r *SMTPCredentialRepository) Revoke(id uint) error {
	return r.db.Model(&models.SMTPCredential{}).Where("id = ?", id).Update("revoked", true).Error
}

func (r *SMTPCredentialRepository) Delete(id uint) error {
	return r.db.Delete(&models.SMTPCredential{}, id).Error
}
