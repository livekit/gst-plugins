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

type Frei0RFilterPr0Be struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterPr0Be() (*Frei0RFilterPr0Be, error) {
	e, err := gst.NewElement("frei0r-filter-pr0be")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPr0Be{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterPr0BeWithName(name string) (*Frei0RFilterPr0Be, error) {
	e, err := gst.NewElementWithName("frei0r-filter-pr0be", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPr0Be{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// big-window (bool)
//
// Display more data

func (e *Frei0RFilterPr0Be) GetBigWindow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("big-window")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetBigWindow(value bool) error {
	return e.Element.SetProperty("big-window", value)
}

// measurement (float64)
//
// What measurement to display

func (e *Frei0RFilterPr0Be) GetMeasurement() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("measurement")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetMeasurement(value float64) error {
	return e.Element.SetProperty("measurement", value)
}

// param-256-scale (bool)
//
// use 0-255 instead of 0.0-1.0

func (e *Frei0RFilterPr0Be) GetParam256Scale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("param-256-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetParam256Scale(value bool) error {
	return e.Element.SetProperty("param-256-scale", value)
}

// show-alpha (bool)
//
// Display alpha value too

func (e *Frei0RFilterPr0Be) GetShowAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetShowAlpha(value bool) error {
	return e.Element.SetProperty("show-alpha", value)
}

// x (float64)
//
// X position of probe

func (e *Frei0RFilterPr0Be) GetX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetX(value float64) error {
	return e.Element.SetProperty("x", value)
}

// x-size (float64)
//
// X size of probe

func (e *Frei0RFilterPr0Be) GetXSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetXSize(value float64) error {
	return e.Element.SetProperty("x-size", value)
}

// y (float64)
//
// Y position of probe

func (e *Frei0RFilterPr0Be) GetY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetY(value float64) error {
	return e.Element.SetProperty("y", value)
}

// y-size (float64)
//
// Y size of probe

func (e *Frei0RFilterPr0Be) GetYSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0Be) SetYSize(value float64) error {
	return e.Element.SetProperty("y-size", value)
}

