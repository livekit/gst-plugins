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

type Frei0RFilterDefish0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterDefish0R() (*Frei0RFilterDefish0R, error) {
	e, err := gst.NewElement("frei0r-filter-defish0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterDefish0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterDefish0RWithName(name string) (*Frei0RFilterDefish0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-defish0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterDefish0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amount (float64)
//
// Focal Ratio

func (e *Frei0RFilterDefish0R) GetAmount() (float64, error) {
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

func (e *Frei0RFilterDefish0R) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

// aspect-type (float64)
//
// Pixel aspect ratio presets

func (e *Frei0RFilterDefish0R) GetAspectType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

// defish (bool)
//
// Fish or Defish

func (e *Frei0RFilterDefish0R) GetDefish() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("defish")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetDefish(value bool) error {
	return e.Element.SetProperty("defish", value)
}

// interpolator (float64)
//
// Quality of interpolation

func (e *Frei0RFilterDefish0R) GetInterpolator() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("interpolator")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetInterpolator(value float64) error {
	return e.Element.SetProperty("interpolator", value)
}

// manual-aspect (float64)
//
// Manual Pixel Aspect ratio

func (e *Frei0RFilterDefish0R) GetManualAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

// manual-scale (float64)
//
// Manual Scale

func (e *Frei0RFilterDefish0R) GetManualScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetManualScale(value float64) error {
	return e.Element.SetProperty("manual-scale", value)
}

// scaling (float64)
//
// Scaling method

func (e *Frei0RFilterDefish0R) GetScaling() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scaling")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetScaling(value float64) error {
	return e.Element.SetProperty("scaling", value)
}

// type (float64)
//
// Mapping function

func (e *Frei0RFilterDefish0R) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDefish0R) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

