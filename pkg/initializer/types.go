// Copyright © 2018 The Gardener Authors.
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

package initializer

import (
	"github.com/gardener/etcd-backup-restore/pkg/initializer/validator"
	"github.com/gardener/etcd-backup-restore/pkg/snapshot/restorer"
	"github.com/sirupsen/logrus"
)

type InitializerConfig struct {
	StorageProvider string
	RestoreOptions  *restorer.RestoreOptions
}

type EtcdInitializer struct {
	Validator *validator.DataValidator
	Config    *InitializerConfig
	Logger    *logrus.Logger
}

type Initializer interface {
	Initialize() error
}
