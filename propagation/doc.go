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

/*
Package propagation contains OpenTelemetry context propagators.

OpenTelemetry propagators are used to extract and inject context data from and
into messages exchanged by applications. The propagator supported by this
package is the W3C Trace Context encoding
(https://www.w3.org/TR/trace-context/), and W3C Baggage
(https://www.w3.org/TR/baggage/).
*/
package propagation // import "github.com/middleware-labs/otel/propagation"
