// Copyright 2024 Kozmoai Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"encoding/json"
	"time"
)

type AppEditedBy struct {
	UserID   int       `json:"userID"`
	EditedAt time.Time `json:"editedAt"`
}

func NewAppEditedByUser(user *User) *AppEditedBy {
	return &AppEditedBy{
		UserID:   user.ID,
		EditedAt: time.Now(),
	}
}

func NewAppEditedByUserID(userID int) *AppEditedBy {
	return &AppEditedBy{
		UserID:   userID,
		EditedAt: time.Now(),
	}
}

func (a *AppEditedBy) ExportToJSONString() string {
	r, _ := json.Marshal(a)
	return string(r)
}
