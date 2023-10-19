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

type Cutter struct {
	Element *gst.Element
}

func NewCutter() (*Cutter, error) {
	e, err := gst.NewElement("cutter")
	if err != nil {
		return nil, err
	}
	return &Cutter{Element: e}, nil
}

func NewCutterWithName(name string) (*Cutter, error) {
	e, err := gst.NewElementWithName("cutter", name)
	if err != nil {
		return nil, err
	}
	return &Cutter{Element: e}, nil
}

// ----- Properties -----

// leaky (bool)
//
// do we leak buffers when below threshold ?

func (e *Cutter) GetLeaky() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("leaky")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cutter) SetLeaky(value bool) error {
	return e.Element.SetProperty("leaky", value)
}

// pre-length (uint64)
//
// Length of pre-recording buffer (in nanoseconds)

func (e *Cutter) GetPreLength() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("pre-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Cutter) SetPreLength(value uint64) error {
	return e.Element.SetProperty("pre-length", value)
}

// run-length (uint64)
//
// Length of drop below threshold before cut_stop (in nanoseconds)

func (e *Cutter) GetRunLength() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("run-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Cutter) SetRunLength(value uint64) error {
	return e.Element.SetProperty("run-length", value)
}

// threshold (float64)
//
// Volume threshold before trigger

func (e *Cutter) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Cutter) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

// threshold-dB (float64)
//
// Volume threshold before trigger (in dB)

func (e *Cutter) GetThresholdDB() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold-dB")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Cutter) SetThresholdDB(value float64) error {
	return e.Element.SetProperty("threshold-dB", value)
}

