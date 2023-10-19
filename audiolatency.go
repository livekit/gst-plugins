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

type Audiolatency struct {
	Element *gst.Element
}

func NewAudiolatency() (*Audiolatency, error) {
	e, err := gst.NewElement("audiolatency")
	if err != nil {
		return nil, err
	}
	return &Audiolatency{Element: e}, nil
}

func NewAudiolatencyWithName(name string) (*Audiolatency, error) {
	e, err := gst.NewElementWithName("audiolatency", name)
	if err != nil {
		return nil, err
	}
	return &Audiolatency{Element: e}, nil
}

// ----- Properties -----

// average-latency (int64)
//
// The running average latency, in microseconds

func (e *Audiolatency) GetAverageLatency() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("average-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// last-latency (int64)
//
// The last latency that was measured, in microseconds

func (e *Audiolatency) GetLastLatency() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("last-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// print-latency (bool)
//
// Print the measured latencies on stdout

func (e *Audiolatency) GetPrintLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("print-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audiolatency) SetPrintLatency(value bool) error {
	return e.Element.SetProperty("print-latency", value)
}

// samplesperbuffer (int)
//
// The number of audio samples in each outgoing buffer.
// See also samplesperbuffer

func (e *Audiolatency) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audiolatency) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

