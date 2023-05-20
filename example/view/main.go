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
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/middleware-labs/otel/attribute"
	otelprom "github.com/middleware-labs/otel/exporters/prometheus"
	"github.com/middleware-labs/otel/metric/instrument"
	"github.com/middleware-labs/otel/sdk/instrumentation"
	"github.com/middleware-labs/otel/sdk/metric"
	"github.com/middleware-labs/otel/sdk/metric/aggregation"
)

const meterName = "github.com/open-telemetry/opentelemetry-go/example/view"

func main() {
	ctx := context.Background()

	// The exporter embeds a default OpenTelemetry Reader, allowing it to be used in WithReader.
	exporter, err := otelprom.New()
	if err != nil {
		log.Fatal(err)
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		// View to customize histogram buckets and rename a single histogram instrument.
		metric.WithView(metric.NewView(
			metric.Instrument{
				Name:  "custom_histogram",
				Scope: instrumentation.Scope{Name: meterName},
			},
			metric.Stream{
				Name: "bar",
				Aggregation: aggregation.ExplicitBucketHistogram{
					Boundaries: []float64{64, 128, 256, 512, 1024, 2048, 4096},
				},
			},
		)),
	)
	meter := provider.Meter(meterName)

	// Start the prometheus HTTP server and pass the exporter Collector to it
	go serveMetrics()

	attrs := []attribute.KeyValue{
		attribute.Key("A").String("B"),
		attribute.Key("C").String("D"),
	}

	counter, err := meter.Float64Counter("foo", instrument.WithDescription("a simple counter"))
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 5, attrs...)

	histogram, err := meter.Float64Histogram("custom_histogram", instrument.WithDescription("a histogram with custom buckets and rename"))
	if err != nil {
		log.Fatal(err)
	}
	histogram.Record(ctx, 136, attrs...)
	histogram.Record(ctx, 64, attrs...)
	histogram.Record(ctx, 701, attrs...)
	histogram.Record(ctx, 830, attrs...)

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}

func serveMetrics() {
	log.Printf("serving metrics at localhost:2222/metrics")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2222", nil)
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}
