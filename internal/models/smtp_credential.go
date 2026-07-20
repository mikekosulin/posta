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

package models

import (
	"time"

	"github.com/lib/pq"
)

type SMTPCredential struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	WorkspaceID  uint           `json:"workspace_id" gorm:"index;not null"`
	UserID       uint           `json:"user_id" gorm:"index;not null"` // creator, for audit
	Name         string         `json:"name" gorm:"not null"`
	Username     string         `json:"username" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	AllowedIPs   pq.StringArray `json:"allowed_ips" gorm:"type:text[]"`
	Revoked      bool           `json:"revoked" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at"`
	LastUsedAt   *time.Time     `json:"last_used_at"`

	User User `json:"-" gorm:"foreignKey:UserID"`
}

func (k *SMTPCredential) IsValid() bool {
	return !k.Revoked
}
