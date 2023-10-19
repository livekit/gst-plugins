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

type Rtph265Pay struct {
	Element *gst.Element
}

func NewRtph265Pay() (*Rtph265Pay, error) {
	e, err := gst.NewElement("rtph265pay")
	if err != nil {
		return nil, err
	}
	return &Rtph265Pay{Element: e}, nil
}

func NewRtph265PayWithName(name string) (*Rtph265Pay, error) {
	e, err := gst.NewElementWithName("rtph265pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtph265Pay{Element: e}, nil
}

// ----- Properties -----

// aggregate-mode (GstRtpH265AggregateMode)
//
// Bundle suitable SPS/PPS NAL units into STAP-A aggregate packets.

// config-interval (int)
//
// Send VPS, SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)

func (e *Rtph265Pay) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtph265Pay) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

// ----- Constants -----

type GstRtpH265AggregateMode string

const (
	GstRtpH265AggregateModeNone GstRtpH265AggregateMode = "none" // none (0) – Do not aggregate NAL units
	GstRtpH265AggregateModeZeroLatency GstRtpH265AggregateMode = "zero-latency" // zero-latency (1) – Aggregate NAL units until a VCL or suffix unit is included
	GstRtpH265AggregateModeMax GstRtpH265AggregateMode = "max" // max (2) – Aggregate all NAL units with the same timestamp (adds one frame of latency)
)

