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

type Frei0RFilterC0Rners struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterC0Rners() (*Frei0RFilterC0Rners, error) {
	e, err := gst.NewElement("frei0r-filter-c0rners")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterC0Rners{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterC0RnersWithName(name string) (*Frei0RFilterC0Rners, error) {
	e, err := gst.NewElementWithName("frei0r-filter-c0rners", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterC0Rners{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha-operation (float64)

func (e *Frei0RFilterC0Rners) GetAlphaOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha-operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetAlphaOperation(value float64) error {
	return e.Element.SetProperty("alpha-operation", value)
}

// corner-1-x (float64)
//
// X coordinate of corner 1

func (e *Frei0RFilterC0Rners) GetCorner1X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-1-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner1X(value float64) error {
	return e.Element.SetProperty("corner-1-x", value)
}

// corner-1-y (float64)
//
// Y coordinate of corner 1

func (e *Frei0RFilterC0Rners) GetCorner1Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-1-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner1Y(value float64) error {
	return e.Element.SetProperty("corner-1-y", value)
}

// corner-2-x (float64)
//
// X coordinate of corner 2

func (e *Frei0RFilterC0Rners) GetCorner2X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-2-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner2X(value float64) error {
	return e.Element.SetProperty("corner-2-x", value)
}

// corner-2-y (float64)
//
// Y coordinate of corner 2

func (e *Frei0RFilterC0Rners) GetCorner2Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-2-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner2Y(value float64) error {
	return e.Element.SetProperty("corner-2-y", value)
}

// corner-3-x (float64)
//
// X coordinate of corner 3

func (e *Frei0RFilterC0Rners) GetCorner3X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-3-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner3X(value float64) error {
	return e.Element.SetProperty("corner-3-x", value)
}

// corner-3-y (float64)
//
// Y coordinate of corner 3

func (e *Frei0RFilterC0Rners) GetCorner3Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-3-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner3Y(value float64) error {
	return e.Element.SetProperty("corner-3-y", value)
}

// corner-4-x (float64)
//
// X coordinate of corner 4

func (e *Frei0RFilterC0Rners) GetCorner4X() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-4-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner4X(value float64) error {
	return e.Element.SetProperty("corner-4-x", value)
}

// corner-4-y (float64)
//
// Y coordinate of corner 4

func (e *Frei0RFilterC0Rners) GetCorner4Y() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("corner-4-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetCorner4Y(value float64) error {
	return e.Element.SetProperty("corner-4-y", value)
}

// enable-stretch (bool)
//
// Enable stretching

func (e *Frei0RFilterC0Rners) GetEnableStretch() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-stretch")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetEnableStretch(value bool) error {
	return e.Element.SetProperty("enable-stretch", value)
}

// feather-alpha (float64)
//
// Makes smooth transition into transparent

func (e *Frei0RFilterC0Rners) GetFeatherAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("feather-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetFeatherAlpha(value float64) error {
	return e.Element.SetProperty("feather-alpha", value)
}

// interpolator (float64)
//
// Quality of interpolation

func (e *Frei0RFilterC0Rners) GetInterpolator() (float64, error) {
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

func (e *Frei0RFilterC0Rners) SetInterpolator(value float64) error {
	return e.Element.SetProperty("interpolator", value)
}

// stretch-x (float64)
//
// Amount of stretching in X direction

func (e *Frei0RFilterC0Rners) GetStretchX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("stretch-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetStretchX(value float64) error {
	return e.Element.SetProperty("stretch-x", value)
}

// stretch-y (float64)
//
// Amount of stretching in Y direction

func (e *Frei0RFilterC0Rners) GetStretchY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("stretch-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetStretchY(value float64) error {
	return e.Element.SetProperty("stretch-y", value)
}

// transparent-background (bool)
//
// Makes background transparent

func (e *Frei0RFilterC0Rners) GetTransparentBackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("transparent-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterC0Rners) SetTransparentBackground(value bool) error {
	return e.Element.SetProperty("transparent-background", value)
}

