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
	"net/http"
	"time"

	"github.com/middleware-labs/otel"
	"github.com/middleware-labs/otel/example/passthrough/handler"
	"github.com/middleware-labs/otel/exporters/stdout/stdouttrace"
	"github.com/middleware-labs/otel/propagation"
	sdktrace "github.com/middleware-labs/otel/sdk/trace"
	"github.com/middleware-labs/otel/trace"
)

func main() {
	ctx := context.Background()

	initPassthroughGlobals()
	tp, err := nonGlobalTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = tp.Shutdown(ctx) }()

	// make an initial http request
	r, err := http.NewRequest("", "", nil)
	if err != nil {
		panic(err)
	}

	// This is roughly what an instrumented http client does.
	log.Println("The \"make outer request\" span should be recorded, because it is recorded with a Tracer from the SDK TracerProvider")
	var span trace.Span
	ctx, span = tp.Tracer("example/passthrough/outer").Start(ctx, "make outer request")
	defer span.End()
	r = r.WithContext(ctx)
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	backendFunc := func(r *http.Request) {
		// This is roughly what an instrumented http server does.
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		log.Println("The \"handle inner request\" span should be recorded, because it is recorded with a Tracer from the SDK TracerProvider")
		_, span := tp.Tracer("example/passthrough/inner").Start(ctx, "handle inner request")
		defer span.End()

		// Do "backend work"
		time.Sleep(time.Second)
	}
	// This handler will be a passthrough, since we didn't set a global TracerProvider
	passthroughHandler := handler.New(backendFunc)
	passthroughHandler.HandleHTTPReq(r)
}

func initPassthroughGlobals() {
	// We explicitly DO NOT set the global TracerProvider using otel.SetTracerProvider().
	// The unset TracerProvider returns a "non-recording" span, but still passes through context.
	log.Println("Register a global TextMapPropagator, but do not register a global TracerProvider to be in \"passthrough\" mode.")
	log.Println("The \"passthrough\" mode propagates the TraceContext and Baggage, but does not record spans.")
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

// nonGlobalTracer creates a trace provider instance for testing, but doesn't
// set it as the global tracer provider.
func nonGlobalTracer() (*sdktrace.TracerProvider, error) {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize stdouttrace exporter: %w", err)
	}
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	return tp, nil
}
