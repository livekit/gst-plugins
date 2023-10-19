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

type Frei0RFilterDistort0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterDistort0R() (*Frei0RFilterDistort0R, error) {
	e, err := gst.NewElement("frei0r-filter-distort0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterDistort0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterDistort0RWithName(name string) (*Frei0RFilterDistort0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-distort0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterDistort0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amplitude (float64)
//
// The amplitude of the plasma signal

func (e *Frei0RFilterDistort0R) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDistort0R) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

// frequency (float64)
//
// The frequency of the plasma signal

func (e *Frei0RFilterDistort0R) GetFrequency() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDistort0R) SetFrequency(value float64) error {
	return e.Element.SetProperty("frequency", value)
}

// use-velocity (bool)
//
// 'Time Based' or 'Adjustable Velocity'

func (e *Frei0RFilterDistort0R) GetUseVelocity() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-velocity")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterDistort0R) SetUseVelocity(value bool) error {
	return e.Element.SetProperty("use-velocity", value)
}

// velocity (float64)
//
// Changing speed of the plasma signal

func (e *Frei0RFilterDistort0R) GetVelocity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("velocity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDistort0R) SetVelocity(value float64) error {
	return e.Element.SetProperty("velocity", value)
}

