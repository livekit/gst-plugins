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

type Frei0RFilterNormaliz0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterNormaliz0R() (*Frei0RFilterNormaliz0R, error) {
	e, err := gst.NewElement("frei0r-filter-normaliz0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterNormaliz0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterNormaliz0RWithName(name string) (*Frei0RFilterNormaliz0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-normaliz0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterNormaliz0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// blackpt-b (float32)
//
// Output color to which darkest input color is mapped (default black)

func (e *Frei0RFilterNormaliz0R) GetBlackptB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blackpt-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetBlackptB(value float32) error {
	return e.Element.SetProperty("blackpt-b", value)
}

// blackpt-g (float32)
//
// Output color to which darkest input color is mapped (default black)

func (e *Frei0RFilterNormaliz0R) GetBlackptG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blackpt-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetBlackptG(value float32) error {
	return e.Element.SetProperty("blackpt-g", value)
}

// blackpt-r (float32)
//
// Output color to which darkest input color is mapped (default black)

func (e *Frei0RFilterNormaliz0R) GetBlackptR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("blackpt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetBlackptR(value float32) error {
	return e.Element.SetProperty("blackpt-r", value)
}

// independence (float64)
//
// Proportion of independent to linked channel normalization (default 1.0)

func (e *Frei0RFilterNormaliz0R) GetIndependence() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("independence")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetIndependence(value float64) error {
	return e.Element.SetProperty("independence", value)
}

// smoothing (float64)
//
// Amount of temporal smoothing of the input range, to reduce flicker (default 0.0)

func (e *Frei0RFilterNormaliz0R) GetSmoothing() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("smoothing")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetSmoothing(value float64) error {
	return e.Element.SetProperty("smoothing", value)
}

// strength (float64)
//
// Strength of filter, from no effect to full normalization (default 1.0)

func (e *Frei0RFilterNormaliz0R) GetStrength() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetStrength(value float64) error {
	return e.Element.SetProperty("strength", value)
}

// whitept-b (float32)
//
// Output color to which brightest input color is mapped (default white)

func (e *Frei0RFilterNormaliz0R) GetWhiteptB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("whitept-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetWhiteptB(value float32) error {
	return e.Element.SetProperty("whitept-b", value)
}

// whitept-g (float32)
//
// Output color to which brightest input color is mapped (default white)

func (e *Frei0RFilterNormaliz0R) GetWhiteptG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("whitept-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetWhiteptG(value float32) error {
	return e.Element.SetProperty("whitept-g", value)
}

// whitept-r (float32)
//
// Output color to which brightest input color is mapped (default white)

func (e *Frei0RFilterNormaliz0R) GetWhiteptR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("whitept-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterNormaliz0R) SetWhiteptR(value float32) error {
	return e.Element.SetProperty("whitept-r", value)
}

