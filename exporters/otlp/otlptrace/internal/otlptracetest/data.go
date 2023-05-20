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

package otlptracetest // import "github.com/middleware-labs/otel/exporters/otlp/otlptrace/internal/otlptracetest"

import (
	"time"

	"github.com/middleware-labs/otel/attribute"
	"github.com/middleware-labs/otel/codes"
	"github.com/middleware-labs/otel/sdk/instrumentation"
	"github.com/middleware-labs/otel/sdk/resource"
	tracesdk "github.com/middleware-labs/otel/sdk/trace"
	"github.com/middleware-labs/otel/sdk/trace/tracetest"
	"github.com/middleware-labs/otel/trace"
)

// SingleReadOnlySpan returns a one-element slice with a read-only span. It
// may be useful for testing driver's trace export.
func SingleReadOnlySpan() []tracesdk.ReadOnlySpan {
	return tracetest.SpanStubs{
		{
			SpanContext: trace.NewSpanContext(trace.SpanContextConfig{
				TraceID:    trace.TraceID{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9},
				SpanID:     trace.SpanID{3, 4, 5, 6, 7, 8, 9, 0},
				TraceFlags: trace.FlagsSampled,
			}),
			Parent: trace.NewSpanContext(trace.SpanContextConfig{
				TraceID:    trace.TraceID{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9},
				SpanID:     trace.SpanID{1, 2, 3, 4, 5, 6, 7, 8},
				TraceFlags: trace.FlagsSampled,
			}),
			SpanKind:          trace.SpanKindInternal,
			Name:              "foo",
			StartTime:         time.Date(2020, time.December, 8, 20, 23, 0, 0, time.UTC),
			EndTime:           time.Date(2020, time.December, 0, 20, 24, 0, 0, time.UTC),
			Attributes:        []attribute.KeyValue{},
			Events:            []tracesdk.Event{},
			Links:             []tracesdk.Link{},
			Status:            tracesdk.Status{Code: codes.Ok},
			DroppedAttributes: 0,
			DroppedEvents:     0,
			DroppedLinks:      0,
			ChildSpanCount:    0,
			Resource:          resource.NewSchemaless(attribute.String("a", "b")),
			InstrumentationLibrary: instrumentation.Library{
				Name:    "bar",
				Version: "0.0.0",
			},
		},
	}.Snapshots()
}
