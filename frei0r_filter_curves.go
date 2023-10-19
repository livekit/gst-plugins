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

type Frei0RFilterCurves struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterCurves() (*Frei0RFilterCurves, error) {
	e, err := gst.NewElement("frei0r-filter-curves")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCurves{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterCurvesWithName(name string) (*Frei0RFilterCurves, error) {
	e, err := gst.NewElementWithName("frei0r-filter-curves", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCurves{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// b--zier-spline (string)
//
// Use cubic Bézier spline. Has to be a sorted list of points in the format "handle1x;handle1y#pointx;pointy#handle2x;handle2y"(pointx = in, pointy = out). Points are separated by a "|".The values can have "double" precision. x, y for points should be in the range 0-1. x,y for handles might also be out of this range.

func (e *Frei0RFilterCurves) GetBZierSpline() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("b--zier-spline")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetBZierSpline(value string) error {
	return e.Element.SetProperty("b--zier-spline", value)
}

// channel (float64)
//
// Channel to adjust (0 = red, 0.1 = green, 0.2 = blue, 0.3 = alpha, 0.4 = luma, 0.5 = rgb, 0.6 = hue, 0.7 = saturation)

func (e *Frei0RFilterCurves) GetChannel() (float64, error) {
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

func (e *Frei0RFilterCurves) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

// curve-point-number (float64)
//
// Number of points to use to build curve (/10 to fit [0,1] parameter range). Minimum 2 (0.2), Maximum 5 (0.5). Not relevant for Bézier spline.

func (e *Frei0RFilterCurves) GetCurvePointNumber() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("curve-point-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetCurvePointNumber(value float64) error {
	return e.Element.SetProperty("curve-point-number", value)
}

// graph-position (float64)
//
// Output image corner where curve graph will be drawn (0.1 = TOP,LEFT; 0.2 = TOP,RIGHT; 0.3 = BOTTOM,LEFT; 0.4 = BOTTOM, RIGHT)

func (e *Frei0RFilterCurves) GetGraphPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("graph-position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetGraphPosition(value float64) error {
	return e.Element.SetProperty("graph-position", value)
}

// luma-formula (bool)
//
// Use Rec. 601 (false) or Rec. 709 (true)

func (e *Frei0RFilterCurves) GetLumaFormula() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("luma-formula")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetLumaFormula(value bool) error {
	return e.Element.SetProperty("luma-formula", value)
}

// point-1-input-value (float64)
//
// Point 1 input value

func (e *Frei0RFilterCurves) GetPoint1InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-1-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint1InputValue(value float64) error {
	return e.Element.SetProperty("point-1-input-value", value)
}

// point-1-output-value (float64)
//
// Point 1 output value

func (e *Frei0RFilterCurves) GetPoint1OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-1-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint1OutputValue(value float64) error {
	return e.Element.SetProperty("point-1-output-value", value)
}

// point-2-input-value (float64)
//
// Point 2 input value

func (e *Frei0RFilterCurves) GetPoint2InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-2-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint2InputValue(value float64) error {
	return e.Element.SetProperty("point-2-input-value", value)
}

// point-2-output-value (float64)
//
// Point 2 output value

func (e *Frei0RFilterCurves) GetPoint2OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-2-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint2OutputValue(value float64) error {
	return e.Element.SetProperty("point-2-output-value", value)
}

// point-3-input-value (float64)
//
// Point 3 input value

func (e *Frei0RFilterCurves) GetPoint3InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-3-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint3InputValue(value float64) error {
	return e.Element.SetProperty("point-3-input-value", value)
}

// point-3-output-value (float64)
//
// Point 3 output value

func (e *Frei0RFilterCurves) GetPoint3OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-3-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint3OutputValue(value float64) error {
	return e.Element.SetProperty("point-3-output-value", value)
}

// point-4-input-value (float64)
//
// Point 4 input value

func (e *Frei0RFilterCurves) GetPoint4InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-4-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint4InputValue(value float64) error {
	return e.Element.SetProperty("point-4-input-value", value)
}

// point-4-output-value (float64)
//
// Point 4 output value

func (e *Frei0RFilterCurves) GetPoint4OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-4-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint4OutputValue(value float64) error {
	return e.Element.SetProperty("point-4-output-value", value)
}

// point-5-input-value (float64)
//
// Point 5 input value

func (e *Frei0RFilterCurves) GetPoint5InputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-5-input-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint5InputValue(value float64) error {
	return e.Element.SetProperty("point-5-input-value", value)
}

// point-5-output-value (float64)
//
// Point 5 output value

func (e *Frei0RFilterCurves) GetPoint5OutputValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("point-5-output-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetPoint5OutputValue(value float64) error {
	return e.Element.SetProperty("point-5-output-value", value)
}

// show-curves (bool)
//
// Draw curve graph on output image

func (e *Frei0RFilterCurves) GetShowCurves() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-curves")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterCurves) SetShowCurves(value bool) error {
	return e.Element.SetProperty("show-curves", value)
}

