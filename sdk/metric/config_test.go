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

package metric

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/middleware-labs/otel/sdk/metric/aggregation"
	"github.com/middleware-labs/otel/sdk/metric/metricdata"
	"github.com/middleware-labs/otel/sdk/resource"
)

type reader struct {
	producer          sdkProducer
	externalProducers []Producer
	temporalityFunc   TemporalitySelector
	aggregationFunc   AggregationSelector
	collectFunc       func(context.Context, *metricdata.ResourceMetrics) error
	forceFlushFunc    func(context.Context) error
	shutdownFunc      func(context.Context) error
}

var _ Reader = (*reader)(nil)

func (r *reader) aggregation(kind InstrumentKind) aggregation.Aggregation { // nolint:revive  // import-shadow for method scoped by type.
	return r.aggregationFunc(kind)
}

func (r *reader) register(p sdkProducer)      { r.producer = p }
func (r *reader) RegisterProducer(p Producer) { r.externalProducers = append(r.externalProducers, p) }
func (r *reader) temporality(kind InstrumentKind) metricdata.Temporality {
	return r.temporalityFunc(kind)
}
func (r *reader) Collect(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	return r.collectFunc(ctx, rm)
}
func (r *reader) ForceFlush(ctx context.Context) error { return r.forceFlushFunc(ctx) }
func (r *reader) Shutdown(ctx context.Context) error   { return r.shutdownFunc(ctx) }

func TestConfigReaderSignalsEmpty(t *testing.T) {
	f, s := config{}.readerSignals()

	require.NotNil(t, f)
	require.NotNil(t, s)

	ctx := context.Background()
	assert.Nil(t, f(ctx))
	assert.Nil(t, s(ctx))
	assert.ErrorIs(t, s(ctx), ErrReaderShutdown)
}

func TestConfigReaderSignalsForwarded(t *testing.T) {
	var flush, sdown int
	r := &reader{
		forceFlushFunc: func(ctx context.Context) error {
			flush++
			return nil
		},
		shutdownFunc: func(ctx context.Context) error {
			sdown++
			return nil
		},
	}
	c := newConfig([]Option{WithReader(r)})
	f, s := c.readerSignals()

	require.NotNil(t, f)
	require.NotNil(t, s)

	ctx := context.Background()
	assert.NoError(t, f(ctx))
	assert.NoError(t, f(ctx))
	assert.NoError(t, s(ctx))
	assert.ErrorIs(t, s(ctx), ErrReaderShutdown)

	assert.Equal(t, 2, flush, "flush not called 2 times")
	assert.Equal(t, 1, sdown, "shutdown not called 1 time")
}

func TestConfigReaderSignalsForwardedErrors(t *testing.T) {
	r := &reader{
		forceFlushFunc: func(ctx context.Context) error { return assert.AnError },
		shutdownFunc:   func(ctx context.Context) error { return assert.AnError },
	}
	c := newConfig([]Option{WithReader(r)})
	f, s := c.readerSignals()

	require.NotNil(t, f)
	require.NotNil(t, s)

	ctx := context.Background()
	assert.ErrorIs(t, f(ctx), assert.AnError)
	assert.ErrorIs(t, s(ctx), assert.AnError)
	assert.ErrorIs(t, s(ctx), ErrReaderShutdown)
}

func TestUnifyMultiError(t *testing.T) {
	f := func(context.Context) error { return assert.AnError }
	funcs := []func(context.Context) error{f, f, f}
	errs := []error{assert.AnError, assert.AnError, assert.AnError}
	target := fmt.Errorf("%v", errs)
	assert.Equal(t, unify(funcs)(context.Background()), target)
}

func TestWithResource(t *testing.T) {
	res := resource.NewSchemaless()
	c := newConfig([]Option{WithResource(res)})
	assert.Same(t, res, c.res)
}

func TestWithReader(t *testing.T) {
	r := &reader{}
	c := newConfig([]Option{WithReader(r)})
	require.Len(t, c.readers, 1)
	assert.Same(t, r, c.readers[0])
}

func TestWithView(t *testing.T) {
	c := newConfig([]Option{WithView(
		NewView(
			Instrument{Kind: InstrumentKindObservableCounter},
			Stream{Name: "a"},
		),
		NewView(
			Instrument{Kind: InstrumentKindCounter},
			Stream{Name: "b"},
		),
	)})
	assert.Len(t, c.views, 2)
}
