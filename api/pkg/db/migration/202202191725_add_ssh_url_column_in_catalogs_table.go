// Copyright © 2022 The Tekton Authors.
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

package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tektoncd/hub/api/gen/log"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"gorm.io/gorm"
)

func addSSHURLColumnInCatalogsTable(log *log.Logger) *gormigrate.Migration {

	return &gormigrate.Migration{
		ID: "202202191725_add_ssh_url_column_in_catalogs_table",
		Migrate: func(db *gorm.DB) error {
			if err := db.Migrator().AddColumn(&model.Catalog{}, "ssh_url"); err != nil {
				log.Error(err)
				return err
			}
			return nil
		},
	}
}