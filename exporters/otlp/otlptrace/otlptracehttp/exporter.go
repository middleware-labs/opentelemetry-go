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

package otlptracehttp // import "github.com/middleware-labs/otel/exporters/otlp/otlptrace/otlptracehttp"

import (
	"context"

	"github.com/middleware-labs/otel/exporters/otlp/otlptrace"
)

// New constructs a new Exporter and starts it.
func New(ctx context.Context, opts ...Option) (*otlptrace.Exporter, error) {
	return otlptrace.New(ctx, NewClient(opts...))
}

// NewUnstarted constructs a new Exporter and does not start it.
func NewUnstarted(opts ...Option) *otlptrace.Exporter {
	return otlptrace.NewUnstarted(NewClient(opts...))
}
