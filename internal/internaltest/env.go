// Copyright The OpenTelemetry Authors
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

package internaltest // import "github.com/middleware-labs/otel/internal/internaltest"

import (
	"os"
)

type Env struct {
	Name   string
	Value  string
	Exists bool
}

// EnvStore stores and recovers environment variables.
type EnvStore interface {
	// Records the environment variable into the store.
	Record(key string)

	// Restore recovers the environment variables in the store.
	Restore() error
}

var _ EnvStore = (*envStore)(nil)

type envStore struct {
	store map[string]Env
}

func (s *envStore) add(env Env) {
	s.store[env.Name] = env
}

func (s *envStore) Restore() error {
	var err error
	for _, v := range s.store {
		if v.Exists {
			err = os.Setenv(v.Name, v.Value)
		} else {
			err = os.Unsetenv(v.Name)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *envStore) setEnv(key, value string) error {
	s.Record(key)

	err := os.Setenv(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *envStore) Record(key string) {
	originValue, exists := os.LookupEnv(key)
	s.add(Env{
		Name:   key,
		Value:  originValue,
		Exists: exists,
	})
}

func NewEnvStore() EnvStore {
	return newEnvStore()
}

func newEnvStore() *envStore {
	return &envStore{store: make(map[string]Env)}
}

func SetEnvVariables(env map[string]string) (EnvStore, error) {
	envStore := newEnvStore()

	for k, v := range env {
		err := envStore.setEnv(k, v)
		if err != nil {
			return nil, err
		}
	}
	return envStore, nil
}
