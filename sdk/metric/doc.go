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

// Package metric provides an implementation of the OpenTelemetry metric SDK.
//
// See https://opentelemetry.io/docs/concepts/signals/metrics/ for information
// about the concept of OpenTelemetry metrics and
// https://opentelemetry.io/docs/concepts/components/ for more information
// about OpenTelemetry SDKs.
//
// The entry point for the metric package is the MeterProvider. It is the
// object that all API calls use to create Meters, instruments, and ultimately
// make metric measurements. Also, it is an object that should be used to
// control the life-cycle (start, flush, and shutdown) of the SDK.
//
// A MeterProvider needs to be configured to export the measured data, this is
// done by configuring it with a Reader implementation (using the WithReader
// MeterProviderOption). Readers take two forms: ones that push to an endpoint
// (NewPeriodicReader), and ones that an endpoint pulls from. See the
// github.com/middleware-labs/otel/exporters package for exporters that can be used as
// or with these Readers.
//
// Each Reader, when registered with the MeterProvider, can be augmented with a
// View. Views allow users that run OpenTelemetry instrumented code to modify
// the generated data of that instrumentation.
//
// The data generated by a MeterProvider needs to include information about its
// origin. A MeterProvider needs to be configured with a Resource, using the
// WithResource MeterProviderOption, to include this information. This Resource
// should be used to describe the unique runtime environment instrumented code
// is being run on. That way when multiple instances of the code are collected
// at a single endpoint their origin is decipherable.
package metric // import "github.com/middleware-labs/otel/sdk/metric"
