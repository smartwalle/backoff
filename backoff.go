/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This is a fork of google.golang.org/grpc/internal/backoff written in a
// more extendable way

package backoff

import (
	"time"
)

// Backoff defines the methodology for backing off.
type Backoff interface {
	// Duration returns the amount of time to wait before the next retry given
	// the number of consecutive failures.
	Duration(retries int) time.Duration
}
