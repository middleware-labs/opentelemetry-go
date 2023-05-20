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

package global // import "github.com/middleware-labs/otel/internal/global"

import (
	"context"
	"sync/atomic"

	"github.com/middleware-labs/otel/attribute"
	"github.com/middleware-labs/otel/metric"
	"github.com/middleware-labs/otel/metric/embedded"
	"github.com/middleware-labs/otel/metric/instrument"
)

// unwrapper unwraps to return the underlying instrument implementation.
type unwrapper interface {
	Unwrap() instrument.Observable
}

type afCounter struct {
	embedded.Float64ObservableCounter
	instrument.Float64Observable

	name string
	opts []instrument.Float64ObservableCounterOption

	delegate atomic.Value //instrument.Float64ObservableCounter
}

var _ unwrapper = (*afCounter)(nil)
var _ instrument.Float64ObservableCounter = (*afCounter)(nil)

func (i *afCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Float64ObservableCounter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *afCounter) Unwrap() instrument.Observable {
	if ctr := i.delegate.Load(); ctr != nil {
		return ctr.(instrument.Float64ObservableCounter)
	}
	return nil
}

type afUpDownCounter struct {
	embedded.Float64ObservableUpDownCounter
	instrument.Float64Observable

	name string
	opts []instrument.Float64ObservableUpDownCounterOption

	delegate atomic.Value //instrument.Float64ObservableUpDownCounter
}

var _ unwrapper = (*afUpDownCounter)(nil)
var _ instrument.Float64ObservableUpDownCounter = (*afUpDownCounter)(nil)

func (i *afUpDownCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Float64ObservableUpDownCounter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *afUpDownCounter) Unwrap() instrument.Observable {
	if ctr := i.delegate.Load(); ctr != nil {
		return ctr.(instrument.Float64ObservableUpDownCounter)
	}
	return nil
}

type afGauge struct {
	embedded.Float64ObservableGauge
	instrument.Float64Observable

	name string
	opts []instrument.Float64ObservableGaugeOption

	delegate atomic.Value //instrument.Float64ObservableGauge
}

var _ unwrapper = (*afGauge)(nil)
var _ instrument.Float64ObservableGauge = (*afGauge)(nil)

func (i *afGauge) setDelegate(m metric.Meter) {
	ctr, err := m.Float64ObservableGauge(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *afGauge) Unwrap() instrument.Observable {
	if ctr := i.delegate.Load(); ctr != nil {
		return ctr.(instrument.Float64ObservableGauge)
	}
	return nil
}

type aiCounter struct {
	embedded.Int64ObservableCounter
	instrument.Int64Observable

	name string
	opts []instrument.Int64ObservableCounterOption

	delegate atomic.Value //instrument.Int64ObservableCounter
}

var _ unwrapper = (*aiCounter)(nil)
var _ instrument.Int64ObservableCounter = (*aiCounter)(nil)

func (i *aiCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Int64ObservableCounter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *aiCounter) Unwrap() instrument.Observable {
	if ctr := i.delegate.Load(); ctr != nil {
		return ctr.(instrument.Int64ObservableCounter)
	}
	return nil
}

type aiUpDownCounter struct {
	embedded.Int64ObservableUpDownCounter
	instrument.Int64Observable

	name string
	opts []instrument.Int64ObservableUpDownCounterOption

	delegate atomic.Value //instrument.Int64ObservableUpDownCounter
}

var _ unwrapper = (*aiUpDownCounter)(nil)
var _ instrument.Int64ObservableUpDownCounter = (*aiUpDownCounter)(nil)

func (i *aiUpDownCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Int64ObservableUpDownCounter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *aiUpDownCounter) Unwrap() instrument.Observable {
	if ctr := i.delegate.Load(); ctr != nil {
		return ctr.(instrument.Int64ObservableUpDownCounter)
	}
	return nil
}

type aiGauge struct {
	embedded.Int64ObservableGauge
	instrument.Int64Observable

	name string
	opts []instrument.Int64ObservableGaugeOption

	delegate atomic.Value //instrument.Int64ObservableGauge
}

var _ unwrapper = (*aiGauge)(nil)
var _ instrument.Int64ObservableGauge = (*aiGauge)(nil)

func (i *aiGauge) setDelegate(m metric.Meter) {
	ctr, err := m.Int64ObservableGauge(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *aiGauge) Unwrap() instrument.Observable {
	if ctr := i.delegate.Load(); ctr != nil {
		return ctr.(instrument.Int64ObservableGauge)
	}
	return nil
}

// Sync Instruments.
type sfCounter struct {
	embedded.Float64Counter

	name string
	opts []instrument.Float64CounterOption

	delegate atomic.Value //instrument.Float64Counter
}

var _ instrument.Float64Counter = (*sfCounter)(nil)

func (i *sfCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Float64Counter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *sfCounter) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	if ctr := i.delegate.Load(); ctr != nil {
		ctr.(instrument.Float64Counter).Add(ctx, incr, attrs...)
	}
}

type sfUpDownCounter struct {
	embedded.Float64UpDownCounter

	name string
	opts []instrument.Float64UpDownCounterOption

	delegate atomic.Value //instrument.Float64UpDownCounter
}

var _ instrument.Float64UpDownCounter = (*sfUpDownCounter)(nil)

func (i *sfUpDownCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Float64UpDownCounter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *sfUpDownCounter) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	if ctr := i.delegate.Load(); ctr != nil {
		ctr.(instrument.Float64UpDownCounter).Add(ctx, incr, attrs...)
	}
}

type sfHistogram struct {
	embedded.Float64Histogram

	name string
	opts []instrument.Float64HistogramOption

	delegate atomic.Value //instrument.Float64Histogram
}

var _ instrument.Float64Histogram = (*sfHistogram)(nil)

func (i *sfHistogram) setDelegate(m metric.Meter) {
	ctr, err := m.Float64Histogram(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *sfHistogram) Record(ctx context.Context, x float64, attrs ...attribute.KeyValue) {
	if ctr := i.delegate.Load(); ctr != nil {
		ctr.(instrument.Float64Histogram).Record(ctx, x, attrs...)
	}
}

type siCounter struct {
	embedded.Int64Counter

	name string
	opts []instrument.Int64CounterOption

	delegate atomic.Value //instrument.Int64Counter
}

var _ instrument.Int64Counter = (*siCounter)(nil)

func (i *siCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Int64Counter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *siCounter) Add(ctx context.Context, x int64, attrs ...attribute.KeyValue) {
	if ctr := i.delegate.Load(); ctr != nil {
		ctr.(instrument.Int64Counter).Add(ctx, x, attrs...)
	}
}

type siUpDownCounter struct {
	embedded.Int64UpDownCounter

	name string
	opts []instrument.Int64UpDownCounterOption

	delegate atomic.Value //instrument.Int64UpDownCounter
}

var _ instrument.Int64UpDownCounter = (*siUpDownCounter)(nil)

func (i *siUpDownCounter) setDelegate(m metric.Meter) {
	ctr, err := m.Int64UpDownCounter(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *siUpDownCounter) Add(ctx context.Context, x int64, attrs ...attribute.KeyValue) {
	if ctr := i.delegate.Load(); ctr != nil {
		ctr.(instrument.Int64UpDownCounter).Add(ctx, x, attrs...)
	}
}

type siHistogram struct {
	embedded.Int64Histogram

	name string
	opts []instrument.Int64HistogramOption

	delegate atomic.Value //instrument.Int64Histogram
}

var _ instrument.Int64Histogram = (*siHistogram)(nil)

func (i *siHistogram) setDelegate(m metric.Meter) {
	ctr, err := m.Int64Histogram(i.name, i.opts...)
	if err != nil {
		GetErrorHandler().Handle(err)
		return
	}
	i.delegate.Store(ctr)
}

func (i *siHistogram) Record(ctx context.Context, x int64, attrs ...attribute.KeyValue) {
	if ctr := i.delegate.Load(); ctr != nil {
		ctr.(instrument.Int64Histogram).Record(ctx, x, attrs...)
	}
}
