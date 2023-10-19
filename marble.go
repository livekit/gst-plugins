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

type Marble struct {
	*base.GstBaseTransform
}

func NewMarble() (*Marble, error) {
	e, err := gst.NewElement("marble")
	if err != nil {
		return nil, err
	}
	return &Marble{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewMarbleWithName(name string) (*Marble, error) {
	e, err := gst.NewElementWithName("marble", name)
	if err != nil {
		return nil, err
	}
	return &Marble{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amount (float64)
//
// Amount of effect

func (e *Marble) GetAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Marble) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

// turbulence (float64)
//
// Turbulence of the effect

func (e *Marble) GetTurbulence() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("turbulence")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Marble) SetTurbulence(value float64) error {
	return e.Element.SetProperty("turbulence", value)
}

// x-scale (float64)
//
// X scale of the texture

func (e *Marble) GetXScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Marble) SetXScale(value float64) error {
	return e.Element.SetProperty("x-scale", value)
}

// y-scale (float64)
//
// Y scale of the texture

func (e *Marble) GetYScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Marble) SetYScale(value float64) error {
	return e.Element.SetProperty("y-scale", value)
}

