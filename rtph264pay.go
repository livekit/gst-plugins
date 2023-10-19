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

type Rtph264Pay struct {
	Element *gst.Element
}

func NewRtph264Pay() (*Rtph264Pay, error) {
	e, err := gst.NewElement("rtph264pay")
	if err != nil {
		return nil, err
	}
	return &Rtph264Pay{Element: e}, nil
}

func NewRtph264PayWithName(name string) (*Rtph264Pay, error) {
	e, err := gst.NewElementWithName("rtph264pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtph264Pay{Element: e}, nil
}

// ----- Properties -----

// aggregate-mode (GstRtpH264AggregateMode)
//
// Bundle suitable SPS/PPS NAL units into STAP-A aggregate packets.

// config-interval (int)
//
// Send SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)

func (e *Rtph264Pay) GetConfigInterval() (int, error) {
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

func (e *Rtph264Pay) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

// sprop-parameter-sets (string)
//
// The base64 sprop-parameter-sets to set in out caps (set to NULL to extract from stream)

func (e *Rtph264Pay) GetSpropParameterSets() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sprop-parameter-sets")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtph264Pay) SetSpropParameterSets(value string) error {
	return e.Element.SetProperty("sprop-parameter-sets", value)
}

// ----- Constants -----

type GstRtpH264AggregateMode string

const (
	GstRtpH264AggregateModeNone GstRtpH264AggregateMode = "none" // none (0) – Do not aggregate NAL units
	GstRtpH264AggregateModeZeroLatency GstRtpH264AggregateMode = "zero-latency" // zero-latency (1) – Aggregate NAL units until a VCL unit is included
	GstRtpH264AggregateModeMaxStap GstRtpH264AggregateMode = "max-stap" // max-stap (2) – Aggregate all NAL units with the same timestamp (adds one frame of latency)
)

