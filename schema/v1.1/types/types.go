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

package types // import "github.com/middleware-labs/otel/schema/v1.1/types"

import types10 "github.com/middleware-labs/otel/schema/v1.0/types"

// TelemetryVersion is a version number key in the schema file (e.g. "1.7.0").
type TelemetryVersion types10.TelemetryVersion

// AttributeName is an attribute name string.
type AttributeName string

// AttributeValue is an attribute value.
type AttributeValue interface{}
