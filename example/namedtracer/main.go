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

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-logr/stdr"

	"github.com/middleware-labs/otel"
	"github.com/middleware-labs/otel/attribute"
	"github.com/middleware-labs/otel/baggage"
	"github.com/middleware-labs/otel/example/namedtracer/foo"
	"github.com/middleware-labs/otel/exporters/stdout/stdouttrace"
	sdktrace "github.com/middleware-labs/otel/sdk/trace"
	"github.com/middleware-labs/otel/trace"
)

var (
	fooKey     = attribute.Key("ex.com/foo")
	barKey     = attribute.Key("ex.com/bar")
	anotherKey = attribute.Key("ex.com/another")
)

var tp *sdktrace.TracerProvider

// initTracer creates and registers trace provider instance.
func initTracer() error {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("failed to initialize stdouttrace exporter: %w", err)
	}
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tp = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)
	return nil
}

func main() {
	// Set logging level to info to see SDK status messages
	stdr.SetVerbosity(5)

	// initialize trace provider.
	if err := initTracer(); err != nil {
		log.Panic(err)
	}

	// Create a named tracer with package path as its name.
	tracer := tp.Tracer("example/namedtracer/main")
	ctx := context.Background()
	defer func() { _ = tp.Shutdown(ctx) }()

	m0, _ := baggage.NewMember(string(fooKey), "foo1")
	m1, _ := baggage.NewMember(string(barKey), "bar1")
	b, _ := baggage.New(m0, m1)
	ctx = baggage.ContextWithBaggage(ctx, b)

	var span trace.Span
	ctx, span = tracer.Start(ctx, "operation")
	defer span.End()
	span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
	span.SetAttributes(anotherKey.String("yes"))
	if err := foo.SubOperation(ctx); err != nil {
		panic(err)
	}
}
