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

package instrument // import "github.com/middleware-labs/otel/metric/instrument"

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64Configuration(t *testing.T) {
	const (
		token  int64 = 43
		desc         = "Instrument description."
		uBytes       = "By"
	)

	run := func(got int64Config) func(*testing.T) {
		return func(t *testing.T) {
			assert.Equal(t, desc, got.Description(), "description")
			assert.Equal(t, uBytes, got.Unit(), "unit")
		}
	}

	t.Run("Int64Counter", run(
		NewInt64CounterConfig(WithDescription(desc), WithUnit(uBytes)),
	))

	t.Run("Int64UpDownCounter", run(
		NewInt64UpDownCounterConfig(WithDescription(desc), WithUnit(uBytes)),
	))

	t.Run("Int64Histogram", run(
		NewInt64HistogramConfig(WithDescription(desc), WithUnit(uBytes)),
	))
}

type int64Config interface {
	Description() string
	Unit() string
}
