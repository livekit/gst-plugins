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

type Frei0RFilterPr0File struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterPr0File() (*Frei0RFilterPr0File, error) {
	e, err := gst.NewElement("frei0r-filter-pr0file")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPr0File{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterPr0FileWithName(name string) (*Frei0RFilterPr0File, error) {
	e, err := gst.NewElementWithName("frei0r-filter-pr0file", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPr0File{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha-trace (bool)
//
// Show Alpha trace on scope

func (e *Frei0RFilterPr0File) GetAlphaTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("alpha-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetAlphaTrace(value bool) error {
	return e.Element.SetProperty("alpha-trace", value)
}

// b-trace (bool)
//
// Show B trace on scope

func (e *Frei0RFilterPr0File) GetBTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetBTrace(value bool) error {
	return e.Element.SetProperty("b-trace", value)
}

// channel (float64)
//
// Channel to numerically display

func (e *Frei0RFilterPr0File) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

// color (float64)
//
// rec 601 or rec 709

func (e *Frei0RFilterPr0File) GetColor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetColor(value float64) error {
	return e.Element.SetProperty("color", value)
}

// crosshair-color (float64)
//
// Color of the profile marker

func (e *Frei0RFilterPr0File) GetCrosshairColor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("crosshair-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetCrosshairColor(value float64) error {
	return e.Element.SetProperty("crosshair-color", value)
}

// display-average (bool)
//
// e

func (e *Frei0RFilterPr0File) GetDisplayAverage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-average")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetDisplayAverage(value bool) error {
	return e.Element.SetProperty("display-average", value)
}

// display-maximum (bool)

func (e *Frei0RFilterPr0File) GetDisplayMaximum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-maximum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetDisplayMaximum(value bool) error {
	return e.Element.SetProperty("display-maximum", value)
}

// display-minimum (bool)

func (e *Frei0RFilterPr0File) GetDisplayMinimum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-minimum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetDisplayMinimum(value bool) error {
	return e.Element.SetProperty("display-minimum", value)
}

// display-rms (bool)

func (e *Frei0RFilterPr0File) GetDisplayRms() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-rms")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetDisplayRms(value bool) error {
	return e.Element.SetProperty("display-rms", value)
}

// g-trace (bool)
//
// Show G trace on scope

func (e *Frei0RFilterPr0File) GetGTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("g-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetGTrace(value bool) error {
	return e.Element.SetProperty("g-trace", value)
}

// length (float64)
//
// Length of profile

func (e *Frei0RFilterPr0File) GetLength() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetLength(value float64) error {
	return e.Element.SetProperty("length", value)
}

// marker-1 (float64)
//
// Position of marker 1

func (e *Frei0RFilterPr0File) GetMarker1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("marker-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetMarker1(value float64) error {
	return e.Element.SetProperty("marker-1", value)
}

// marker-2 (float64)
//
// Position of marker 2

func (e *Frei0RFilterPr0File) GetMarker2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("marker-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetMarker2(value float64) error {
	return e.Element.SetProperty("marker-2", value)
}

// param-256-scale (bool)
//
// use 0-255 instead of 0.0-1.0

func (e *Frei0RFilterPr0File) GetParam256Scale() (bool, error) {
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

func (e *Frei0RFilterPr0File) SetParam256Scale(value bool) error {
	return e.Element.SetProperty("param-256-scale", value)
}

// pb-trace (bool)
//
// Show Pb trace on scope

func (e *Frei0RFilterPr0File) GetPbTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pb-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetPbTrace(value bool) error {
	return e.Element.SetProperty("pb-trace", value)
}

// pr-trace (bool)
//
// Show Pr trace on scope

func (e *Frei0RFilterPr0File) GetPrTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pr-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetPrTrace(value bool) error {
	return e.Element.SetProperty("pr-trace", value)
}

// r-trace (bool)
//
// Show R trace on scope

func (e *Frei0RFilterPr0File) GetRTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("r-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetRTrace(value bool) error {
	return e.Element.SetProperty("r-trace", value)
}

// tilt (float64)
//
// Tilt of profile

func (e *Frei0RFilterPr0File) GetTilt() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

// x (float64)
//
// X position of profile

func (e *Frei0RFilterPr0File) GetX() (float64, error) {
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

func (e *Frei0RFilterPr0File) SetX(value float64) error {
	return e.Element.SetProperty("x", value)
}

// y (float64)
//
// Y position of profile

func (e *Frei0RFilterPr0File) GetY() (float64, error) {
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

func (e *Frei0RFilterPr0File) SetY(value float64) error {
	return e.Element.SetProperty("y", value)
}

// y-trace (bool)
//
// Show Y' trace on scope

func (e *Frei0RFilterPr0File) GetYTrace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("y-trace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPr0File) SetYTrace(value bool) error {
	return e.Element.SetProperty("y-trace", value)
}

