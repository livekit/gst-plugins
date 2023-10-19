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

type Frei0RSrcIsing0R struct {
	*base.GstPushSrc
}

func NewFrei0RSrcIsing0R() (*Frei0RSrcIsing0R, error) {
	e, err := gst.NewElement("frei0r-src-ising0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcIsing0R{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcIsing0RWithName(name string) (*Frei0RSrcIsing0R, error) {
	e, err := gst.NewElementWithName("frei0r-src-ising0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcIsing0R{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// border-growth (float64)
//
// Border Growth

func (e *Frei0RSrcIsing0R) GetBorderGrowth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("border-growth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcIsing0R) SetBorderGrowth(value float64) error {
	return e.Element.SetProperty("border-growth", value)
}

// spontaneous-growth (float64)
//
// Spontaneous Growth

func (e *Frei0RSrcIsing0R) GetSpontaneousGrowth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spontaneous-growth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcIsing0R) SetSpontaneousGrowth(value float64) error {
	return e.Element.SetProperty("spontaneous-growth", value)
}

// temperature (float64)
//
// Noise Temperature

func (e *Frei0RSrcIsing0R) GetTemperature() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("temperature")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcIsing0R) SetTemperature(value float64) error {
	return e.Element.SetProperty("temperature", value)
}

