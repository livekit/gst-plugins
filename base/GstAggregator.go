// Copyright 2023 LiveKit, Inc.
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

package base

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type GstAggregator struct {
	Element *gst.Element
}

// ----- Properties -----

// emit-signals (bool)
//
// Enables the emission of signals such as samples-selected

func (e *GstAggregator) GetEmitSignals() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("emit-signals")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *GstAggregator) SetEmitSignals(value bool) error {
	return e.Element.SetProperty("emit-signals", value)
}

// latency (uint64)

func (e *GstAggregator) GetLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *GstAggregator) SetLatency(value uint64) error {
	return e.Element.SetProperty("latency", value)
}

// min-upstream-latency (uint64)
//
// Force minimum upstream latency (in nanoseconds). When sources with a
// higher latency are expected to be plugged in dynamically after the
// aggregator has started playing, this allows overriding the minimum
// latency reported by the initial source(s). This is only taken into
// account when larger than the actually reported minimum latency.

func (e *GstAggregator) GetMinUpstreamLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-upstream-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *GstAggregator) SetMinUpstreamLatency(value uint64) error {
	return e.Element.SetProperty("min-upstream-latency", value)
}

// start-time (uint64)

func (e *GstAggregator) GetStartTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("start-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *GstAggregator) SetStartTime(value uint64) error {
	return e.Element.SetProperty("start-time", value)
}

// start-time-selection (GstAggregatorStartTimeSelection)

func (e *GstAggregator) GetStartTimeSelection() (interface{}, error) {
	return e.Element.GetProperty("start-time-selection")
}

func (e *GstAggregator) SetStartTimeSelection(value interface{}) error {
	return e.Element.SetProperty("start-time-selection", value)
}

