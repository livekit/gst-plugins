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

	"github.com/livekit/gstplugins/base"
)

type Interaudiosrc struct {
	*base.GstBaseSrc
}

func NewInteraudiosrc() (*Interaudiosrc, error) {
	e, err := gst.NewElement("interaudiosrc")
	if err != nil {
		return nil, err
	}
	return &Interaudiosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewInteraudiosrcWithName(name string) (*Interaudiosrc, error) {
	e, err := gst.NewElementWithName("interaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &Interaudiosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// buffer-time (uint64)
//
// Size of audio buffer

func (e *Interaudiosrc) GetBufferTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Interaudiosrc) SetBufferTime(value uint64) error {
	return e.Element.SetProperty("buffer-time", value)
}

// channel (string)
//
// Channel name to match inter src and sink elements

func (e *Interaudiosrc) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Interaudiosrc) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

// latency-time (uint64)
//
// Latency as reported by the source

func (e *Interaudiosrc) GetLatencyTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("latency-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Interaudiosrc) SetLatencyTime(value uint64) error {
	return e.Element.SetProperty("latency-time", value)
}

// period-time (uint64)
//
// The minimum amount of data to read in each iteration

func (e *Interaudiosrc) GetPeriodTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("period-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Interaudiosrc) SetPeriodTime(value uint64) error {
	return e.Element.SetProperty("period-time", value)
}

