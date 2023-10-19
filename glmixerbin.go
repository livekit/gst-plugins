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

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type Glmixerbin struct {
	Element *gst.Element
}

func NewGlmixerbin() (*Glmixerbin, error) {
	e, err := gst.NewElement("glmixerbin")
	if err != nil {
		return nil, err
	}
	return &Glmixerbin{Element: e}, nil
}

func NewGlmixerbinWithName(name string) (*Glmixerbin, error) {
	e, err := gst.NewElementWithName("glmixerbin", name)
	if err != nil {
		return nil, err
	}
	return &Glmixerbin{Element: e}, nil
}

// ----- Properties -----

// context (GstGLContext)
//
// Get OpenGL context

func (e *Glmixerbin) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// force-live (bool)
//
// Causes the element to aggregate on a timeout even when no live source is
// connected to its sinks. See min-upstream-latency for a
// companion property: in the vast majority of cases where you plan to plug in
// live sources with a non-zero latency, you should set it to a non-zero value.

func (e *Glmixerbin) GetForceLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Glmixerbin) SetForceLive(value bool) error {
	return e.Element.SetProperty("force-live", value)
}

// latency (uint64)
//
// Additional latency in live mode to allow upstream to take longer to produce buffers for the current position (in nanoseconds)

func (e *Glmixerbin) GetLatency() (uint64, error) {
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

func (e *Glmixerbin) SetLatency(value uint64) error {
	return e.Element.SetProperty("latency", value)
}

// min-upstream-latency (uint64)
//
// Force minimum upstream latency (in nanoseconds). When sources with a
// higher latency are expected to be plugged in dynamically after the
// aggregator has started playing, this allows overriding the minimum
// latency reported by the initial source(s). This is only taken into
// account when larger than the actually reported minimum latency.

func (e *Glmixerbin) GetMinUpstreamLatency() (uint64, error) {
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

func (e *Glmixerbin) SetMinUpstreamLatency(value uint64) error {
	return e.Element.SetProperty("min-upstream-latency", value)
}

// mixer (GstElement)
//
// The GL mixer chain to use

func (e *Glmixerbin) GetMixer() (interface{}, error) {
	return e.Element.GetProperty("mixer")
}

func (e *Glmixerbin) SetMixer(value interface{}) error {
	return e.Element.SetProperty("mixer", value)
}

// start-time (uint64)
//
// Start time to use if start-time-selection=set

func (e *Glmixerbin) GetStartTime() (uint64, error) {
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

func (e *Glmixerbin) SetStartTime(value uint64) error {
	return e.Element.SetProperty("start-time", value)
}

// start-time-selection (GstGlmixerBinStartTimeSelection)
//
// Decides which start time is output

func (e *Glmixerbin) GetStartTimeSelection() (GstGlmixerBinStartTimeSelection, error) {
	var value GstGlmixerBinStartTimeSelection
	var ok bool
	v, err := e.Element.GetProperty("start-time-selection")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGlmixerBinStartTimeSelection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGlmixerBinStartTimeSelection")
	}
	return value, nil
}

func (e *Glmixerbin) SetStartTimeSelection(value GstGlmixerBinStartTimeSelection) error {
	e.Element.SetArg("start-time-selection", string(value))
	return nil
}

// ----- Constants -----

type GstGlmixerBinStartTimeSelection string

const (
	GstGlmixerBinStartTimeSelectionZero GstGlmixerBinStartTimeSelection = "zero" // zero (0) – Start at 0 running time (default)
	GstGlmixerBinStartTimeSelectionFirst GstGlmixerBinStartTimeSelection = "first" // first (1) – Start at first observed input running time
	GstGlmixerBinStartTimeSelectionSet GstGlmixerBinStartTimeSelection = "set" // set (2) – Set start time with start-time property
)

