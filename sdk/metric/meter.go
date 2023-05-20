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

package metric // import "github.com/middleware-labs/otel/sdk/metric"

import (
	"context"
	"errors"
	"fmt"

	"github.com/middleware-labs/otel/attribute"
	"github.com/middleware-labs/otel/internal/global"
	"github.com/middleware-labs/otel/metric"
	"github.com/middleware-labs/otel/metric/embedded"
	"github.com/middleware-labs/otel/metric/instrument"
	"github.com/middleware-labs/otel/sdk/instrumentation"
	"github.com/middleware-labs/otel/sdk/metric/internal"
)

// meter handles the creation and coordination of all metric instruments. A
// meter represents a single instrumentation scope; all metric telemetry
// produced by an instrumentation scope will use metric instruments from a
// single meter.
type meter struct {
	embedded.Meter

	scope instrumentation.Scope
	pipes pipelines

	int64IP   *instProvider[int64]
	float64IP *instProvider[float64]
}

func newMeter(s instrumentation.Scope, p pipelines) *meter {
	// viewCache ensures instrument conflicts, including number conflicts, this
	// meter is asked to create are logged to the user.
	var viewCache cache[string, streamID]

	return &meter{
		scope:     s,
		pipes:     p,
		int64IP:   newInstProvider[int64](s, p, &viewCache),
		float64IP: newInstProvider[float64](s, p, &viewCache),
	}
}

// Compile-time check meter implements metric.Meter.
var _ metric.Meter = (*meter)(nil)

// Int64Counter returns a new instrument identified by name and configured with
// options. The instrument is used to synchronously record increasing int64
// measurements during a computational operation.
func (m *meter) Int64Counter(name string, options ...instrument.Int64CounterOption) (instrument.Int64Counter, error) {
	cfg := instrument.NewInt64CounterConfig(options...)
	const kind = InstrumentKindCounter
	return m.int64IP.lookup(kind, name, cfg.Description(), cfg.Unit())
}

// Int64UpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// int64 measurements during a computational operation.
func (m *meter) Int64UpDownCounter(name string, options ...instrument.Int64UpDownCounterOption) (instrument.Int64UpDownCounter, error) {
	cfg := instrument.NewInt64UpDownCounterConfig(options...)
	const kind = InstrumentKindUpDownCounter
	return m.int64IP.lookup(kind, name, cfg.Description(), cfg.Unit())
}

// Int64Histogram returns a new instrument identified by name and configured
// with options. The instrument is used to synchronously record the
// distribution of int64 measurements during a computational operation.
func (m *meter) Int64Histogram(name string, options ...instrument.Int64HistogramOption) (instrument.Int64Histogram, error) {
	cfg := instrument.NewInt64HistogramConfig(options...)
	const kind = InstrumentKindHistogram
	return m.int64IP.lookup(kind, name, cfg.Description(), cfg.Unit())
}

// Int64ObservableCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// increasing int64 measurements once per a measurement collection cycle.
func (m *meter) Int64ObservableCounter(name string, options ...instrument.Int64ObservableCounterOption) (instrument.Int64ObservableCounter, error) {
	cfg := instrument.NewInt64ObservableCounterConfig(options...)
	const kind = InstrumentKindObservableCounter
	p := int64ObservProvider{m.int64IP}
	inst, err := p.lookup(kind, name, cfg.Description(), cfg.Unit())
	if err != nil {
		return nil, err
	}
	p.registerCallbacks(inst, cfg.Callbacks())
	return inst, nil
}

// Int64ObservableUpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// int64 measurements once per a measurement collection cycle.
func (m *meter) Int64ObservableUpDownCounter(name string, options ...instrument.Int64ObservableUpDownCounterOption) (instrument.Int64ObservableUpDownCounter, error) {
	cfg := instrument.NewInt64ObservableUpDownCounterConfig(options...)
	const kind = InstrumentKindObservableUpDownCounter
	p := int64ObservProvider{m.int64IP}
	inst, err := p.lookup(kind, name, cfg.Description(), cfg.Unit())
	if err != nil {
		return nil, err
	}
	p.registerCallbacks(inst, cfg.Callbacks())
	return inst, nil
}

// Int64ObservableGauge returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// instantaneous int64 measurements once per a measurement collection cycle.
func (m *meter) Int64ObservableGauge(name string, options ...instrument.Int64ObservableGaugeOption) (instrument.Int64ObservableGauge, error) {
	cfg := instrument.NewInt64ObservableGaugeConfig(options...)
	const kind = InstrumentKindObservableGauge
	p := int64ObservProvider{m.int64IP}
	inst, err := p.lookup(kind, name, cfg.Description(), cfg.Unit())
	if err != nil {
		return nil, err
	}
	p.registerCallbacks(inst, cfg.Callbacks())
	return inst, nil
}

// Float64Counter returns a new instrument identified by name and configured
// with options. The instrument is used to synchronously record increasing
// float64 measurements during a computational operation.
func (m *meter) Float64Counter(name string, options ...instrument.Float64CounterOption) (instrument.Float64Counter, error) {
	cfg := instrument.NewFloat64CounterConfig(options...)
	const kind = InstrumentKindCounter
	return m.float64IP.lookup(kind, name, cfg.Description(), cfg.Unit())
}

// Float64UpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// float64 measurements during a computational operation.
func (m *meter) Float64UpDownCounter(name string, options ...instrument.Float64UpDownCounterOption) (instrument.Float64UpDownCounter, error) {
	cfg := instrument.NewFloat64UpDownCounterConfig(options...)
	const kind = InstrumentKindUpDownCounter
	return m.float64IP.lookup(kind, name, cfg.Description(), cfg.Unit())
}

// Float64Histogram returns a new instrument identified by name and configured
// with options. The instrument is used to synchronously record the
// distribution of float64 measurements during a computational operation.
func (m *meter) Float64Histogram(name string, options ...instrument.Float64HistogramOption) (instrument.Float64Histogram, error) {
	cfg := instrument.NewFloat64HistogramConfig(options...)
	const kind = InstrumentKindHistogram
	return m.float64IP.lookup(kind, name, cfg.Description(), cfg.Unit())
}

// Float64ObservableCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// increasing float64 measurements once per a measurement collection cycle.
func (m *meter) Float64ObservableCounter(name string, options ...instrument.Float64ObservableCounterOption) (instrument.Float64ObservableCounter, error) {
	cfg := instrument.NewFloat64ObservableCounterConfig(options...)
	const kind = InstrumentKindObservableCounter
	p := float64ObservProvider{m.float64IP}
	inst, err := p.lookup(kind, name, cfg.Description(), cfg.Unit())
	if err != nil {
		return nil, err
	}
	p.registerCallbacks(inst, cfg.Callbacks())
	return inst, nil
}

// Float64ObservableUpDownCounter returns a new instrument identified by name
// and configured with options. The instrument is used to asynchronously record
// float64 measurements once per a measurement collection cycle.
func (m *meter) Float64ObservableUpDownCounter(name string, options ...instrument.Float64ObservableUpDownCounterOption) (instrument.Float64ObservableUpDownCounter, error) {
	cfg := instrument.NewFloat64ObservableUpDownCounterConfig(options...)
	const kind = InstrumentKindObservableUpDownCounter
	p := float64ObservProvider{m.float64IP}
	inst, err := p.lookup(kind, name, cfg.Description(), cfg.Unit())
	if err != nil {
		return nil, err
	}
	p.registerCallbacks(inst, cfg.Callbacks())
	return inst, nil
}

// Float64ObservableGauge returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// instantaneous float64 measurements once per a measurement collection cycle.
func (m *meter) Float64ObservableGauge(name string, options ...instrument.Float64ObservableGaugeOption) (instrument.Float64ObservableGauge, error) {
	cfg := instrument.NewFloat64ObservableGaugeConfig(options...)
	const kind = InstrumentKindObservableGauge
	p := float64ObservProvider{m.float64IP}
	inst, err := p.lookup(kind, name, cfg.Description(), cfg.Unit())
	if err != nil {
		return nil, err
	}
	p.registerCallbacks(inst, cfg.Callbacks())
	return inst, nil
}

// RegisterCallback registers f to be called each collection cycle so it will
// make observations for insts during those cycles.
//
// The only instruments f can make observations for are insts. All other
// observations will be dropped and an error will be logged.
//
// Only instruments from this meter can be registered with f, an error is
// returned if other instrument are provided.
//
// The returned Registration can be used to unregister f.
func (m *meter) RegisterCallback(f metric.Callback, insts ...instrument.Observable) (metric.Registration, error) {
	if len(insts) == 0 {
		// Don't allocate a observer if not needed.
		return noopRegister{}, nil
	}

	reg := newObserver()
	var errs multierror
	for _, inst := range insts {
		// Unwrap any global.
		if u, ok := inst.(interface {
			Unwrap() instrument.Observable
		}); ok {
			inst = u.Unwrap()
		}

		switch o := inst.(type) {
		case int64Observable:
			if err := o.registerable(m.scope); err != nil {
				if !errors.Is(err, errEmptyAgg) {
					errs.append(err)
				}
				continue
			}
			reg.registerInt64(o.observablID)
		case float64Observable:
			if err := o.registerable(m.scope); err != nil {
				if !errors.Is(err, errEmptyAgg) {
					errs.append(err)
				}
				continue
			}
			reg.registerFloat64(o.observablID)
		default:
			// Instrument external to the SDK.
			return nil, fmt.Errorf("invalid observable: from different implementation")
		}
	}

	if err := errs.errorOrNil(); err != nil {
		return nil, err
	}

	if reg.len() == 0 {
		// All insts use drop aggregation.
		return noopRegister{}, nil
	}

	cback := func(ctx context.Context) error {
		return f(ctx, reg)
	}
	return m.pipes.registerMultiCallback(cback), nil
}

type observer struct {
	embedded.Observer

	float64 map[observablID[float64]]struct{}
	int64   map[observablID[int64]]struct{}
}

func newObserver() observer {
	return observer{
		float64: make(map[observablID[float64]]struct{}),
		int64:   make(map[observablID[int64]]struct{}),
	}
}

func (r observer) len() int {
	return len(r.float64) + len(r.int64)
}

func (r observer) registerFloat64(id observablID[float64]) {
	r.float64[id] = struct{}{}
}

func (r observer) registerInt64(id observablID[int64]) {
	r.int64[id] = struct{}{}
}

var (
	errUnknownObserver = errors.New("unknown observable instrument")
	errUnregObserver   = errors.New("observable instrument not registered for callback")
)

func (r observer) ObserveFloat64(o instrument.Float64Observable, v float64, a ...attribute.KeyValue) {
	var oImpl float64Observable
	switch conv := o.(type) {
	case float64Observable:
		oImpl = conv
	case interface {
		Unwrap() instrument.Observable
	}:
		// Unwrap any global.
		async := conv.Unwrap()
		var ok bool
		if oImpl, ok = async.(float64Observable); !ok {
			global.Error(errUnknownObserver, "failed to record asynchronous")
			return
		}
	default:
		global.Error(errUnknownObserver, "failed to record")
		return
	}

	if _, registered := r.float64[oImpl.observablID]; !registered {
		global.Error(errUnregObserver, "failed to record",
			"name", oImpl.name,
			"description", oImpl.description,
			"unit", oImpl.unit,
			"number", fmt.Sprintf("%T", float64(0)),
		)
		return
	}
	oImpl.observe(v, a)
}

func (r observer) ObserveInt64(o instrument.Int64Observable, v int64, a ...attribute.KeyValue) {
	var oImpl int64Observable
	switch conv := o.(type) {
	case int64Observable:
		oImpl = conv
	case interface {
		Unwrap() instrument.Observable
	}:
		// Unwrap any global.
		async := conv.Unwrap()
		var ok bool
		if oImpl, ok = async.(int64Observable); !ok {
			global.Error(errUnknownObserver, "failed to record asynchronous")
			return
		}
	default:
		global.Error(errUnknownObserver, "failed to record")
		return
	}

	if _, registered := r.int64[oImpl.observablID]; !registered {
		global.Error(errUnregObserver, "failed to record",
			"name", oImpl.name,
			"description", oImpl.description,
			"unit", oImpl.unit,
			"number", fmt.Sprintf("%T", int64(0)),
		)
		return
	}
	oImpl.observe(v, a)
}

type noopRegister struct{ embedded.Registration }

func (noopRegister) Unregister() error {
	return nil
}

// instProvider provides all OpenTelemetry instruments.
type instProvider[N int64 | float64] struct {
	scope   instrumentation.Scope
	pipes   pipelines
	resolve resolver[N]
}

func newInstProvider[N int64 | float64](s instrumentation.Scope, p pipelines, c *cache[string, streamID]) *instProvider[N] {
	return &instProvider[N]{scope: s, pipes: p, resolve: newResolver[N](p, c)}
}

func (p *instProvider[N]) aggs(kind InstrumentKind, name, desc, u string) ([]internal.Aggregator[N], error) {
	inst := Instrument{
		Name:        name,
		Description: desc,
		Unit:        u,
		Kind:        kind,
		Scope:       p.scope,
	}
	return p.resolve.Aggregators(inst)
}

// lookup returns the resolved instrumentImpl.
func (p *instProvider[N]) lookup(kind InstrumentKind, name, desc, u string) (*instrumentImpl[N], error) {
	aggs, err := p.aggs(kind, name, desc, u)
	return &instrumentImpl[N]{aggregators: aggs}, err
}

type int64ObservProvider struct{ *instProvider[int64] }

func (p int64ObservProvider) lookup(kind InstrumentKind, name, desc, u string) (int64Observable, error) {
	aggs, err := p.aggs(kind, name, desc, u)
	return newInt64Observable(p.scope, kind, name, desc, u, aggs), err
}

func (p int64ObservProvider) registerCallbacks(inst int64Observable, cBacks []instrument.Int64Callback) {
	if inst.observable == nil || len(inst.aggregators) == 0 {
		// Drop aggregator.
		return
	}

	for _, cBack := range cBacks {
		p.pipes.registerCallback(p.callback(inst, cBack))
	}
}

func (p int64ObservProvider) callback(i int64Observable, f instrument.Int64Callback) func(context.Context) error {
	inst := int64Observer{int64Observable: i}
	return func(ctx context.Context) error { return f(ctx, inst) }
}

type int64Observer struct {
	embedded.Int64Observer
	int64Observable
}

func (o int64Observer) Observe(val int64, attrs ...attribute.KeyValue) {
	o.observe(val, attrs)
}

type float64ObservProvider struct{ *instProvider[float64] }

func (p float64ObservProvider) lookup(kind InstrumentKind, name, desc, u string) (float64Observable, error) {
	aggs, err := p.aggs(kind, name, desc, u)
	return newFloat64Observable(p.scope, kind, name, desc, u, aggs), err
}

func (p float64ObservProvider) registerCallbacks(inst float64Observable, cBacks []instrument.Float64Callback) {
	if inst.observable == nil || len(inst.aggregators) == 0 {
		// Drop aggregator.
		return
	}

	for _, cBack := range cBacks {
		p.pipes.registerCallback(p.callback(inst, cBack))
	}
}

func (p float64ObservProvider) callback(i float64Observable, f instrument.Float64Callback) func(context.Context) error {
	inst := float64Observer{float64Observable: i}
	return func(ctx context.Context) error { return f(ctx, inst) }
}

type float64Observer struct {
	embedded.Float64Observer
	float64Observable
}

func (o float64Observer) Observe(val float64, attrs ...attribute.KeyValue) {
	o.observe(val, attrs)
}
