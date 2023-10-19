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

type Tsparse struct {
	Element *gst.Element
}

func NewTsparse() (*Tsparse, error) {
	e, err := gst.NewElement("tsparse")
	if err != nil {
		return nil, err
	}
	return &Tsparse{Element: e}, nil
}

func NewTsparseWithName(name string) (*Tsparse, error) {
	e, err := gst.NewElementWithName("tsparse", name)
	if err != nil {
		return nil, err
	}
	return &Tsparse{Element: e}, nil
}

// ----- Properties -----

// alignment (uint)
//
// Number of packets per buffer (padded with dummy packets on EOS) (0 = auto)

func (e *Tsparse) GetAlignment() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Tsparse) SetAlignment(value uint) error {
	return e.Element.SetProperty("alignment", value)
}

// pcr-pid (int)
//
// Set the PID to use for PCR values (-1 for auto)

func (e *Tsparse) GetPcrPid() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pcr-pid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tsparse) SetPcrPid(value int) error {
	return e.Element.SetProperty("pcr-pid", value)
}

// set-timestamps (bool)
//
// If set, timestamps will be set on the output buffers using PCRs and smoothed over the smoothing-latency period

func (e *Tsparse) GetSetTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("set-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tsparse) SetSetTimestamps(value bool) error {
	return e.Element.SetProperty("set-timestamps", value)
}

// smoothing-latency (uint)
//
// Additional latency in microseconds for smoothing jitter in input timestamps on live capture

func (e *Tsparse) GetSmoothingLatency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("smoothing-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Tsparse) SetSmoothingLatency(value uint) error {
	return e.Element.SetProperty("smoothing-latency", value)
}

// split-on-rai (bool)
//
// If set, buffers sized smaller than the alignment will be sent so that RAI packets are at the start of a new buffer

func (e *Tsparse) GetSplitOnRai() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("split-on-rai")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tsparse) SetSplitOnRai(value bool) error {
	return e.Element.SetProperty("split-on-rai", value)
}

