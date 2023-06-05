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
	"math/rand"
	"sync"
	"time"
)

var (
	DefaultExponential = NewExponential()
)

// Exponential implements Exponential backoff algorithm as defined in
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
type Exponential struct {
	r  *rand.Rand
	mu sync.Mutex

	// BaseDelay is the amount of time to backoff after the first failure.
	BaseDelay time.Duration
	// Multiplier is the factor with which to multiply backoffs after a
	// failed retry. Should ideally be greater than 1.
	Multiplier float64
	// Jitter is the factor with which backoffs are randomized.
	Jitter float64
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

func NewExponential() *Exponential {
	var e = &Exponential{}
	e.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	e.BaseDelay = 1.0 * time.Second
	e.Multiplier = 1.6
	e.Jitter = 0.2
	e.MaxDelay = 120 * time.Second
	return e
}

// Duration returns the amount of time to wait before the next retry given the
// number of retries.
func (e *Exponential) Duration(retries int) time.Duration {
	if retries == 0 {
		return e.BaseDelay
	}
	backoff, max := float64(e.BaseDelay), float64(e.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= e.Multiplier
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + e.Jitter*(e.float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}

func (e *Exponential) float64() float64 {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.r.Float64()
}
