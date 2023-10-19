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

type Frei0RFilterColoradjRgb struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterColoradjRgb() (*Frei0RFilterColoradjRgb, error) {
	e, err := gst.NewElement("frei0r-filter-coloradj-rgb")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColoradjRgb{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterColoradjRgbWithName(name string) (*Frei0RFilterColoradjRgb, error) {
	e, err := gst.NewElementWithName("frei0r-filter-coloradj-rgb", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColoradjRgb{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// action (float64)
//
// Type of color adjustment

func (e *Frei0RFilterColoradjRgb) GetAction() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("action")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetAction(value float64) error {
	return e.Element.SetProperty("action", value)
}

// alpha-controlled (bool)
//
// Adjust only areas with nonzero alpha

func (e *Frei0RFilterColoradjRgb) GetAlphaControlled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("alpha-controlled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetAlphaControlled(value bool) error {
	return e.Element.SetProperty("alpha-controlled", value)
}

// b (float64)
//
// Amount of blue

func (e *Frei0RFilterColoradjRgb) GetB() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetB(value float64) error {
	return e.Element.SetProperty("b", value)
}

// g (float64)
//
// Amount of green

func (e *Frei0RFilterColoradjRgb) GetG() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetG(value float64) error {
	return e.Element.SetProperty("g", value)
}

// keep-luma (bool)
//
// Don't change brightness

func (e *Frei0RFilterColoradjRgb) GetKeepLuma() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-luma")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetKeepLuma(value bool) error {
	return e.Element.SetProperty("keep-luma", value)
}

// luma-formula (float64)

func (e *Frei0RFilterColoradjRgb) GetLumaFormula() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("luma-formula")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetLumaFormula(value float64) error {
	return e.Element.SetProperty("luma-formula", value)
}

// r (float64)
//
// Amount of red

func (e *Frei0RFilterColoradjRgb) GetR() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColoradjRgb) SetR(value float64) error {
	return e.Element.SetProperty("r", value)
}

