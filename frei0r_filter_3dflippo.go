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

type Frei0RFilter3Dflippo struct {
	*base.GstBaseTransform
}

func NewFrei0RFilter3Dflippo() (*Frei0RFilter3Dflippo, error) {
	e, err := gst.NewElement("frei0r-filter-3dflippo")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilter3Dflippo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilter3DflippoWithName(name string) (*Frei0RFilter3Dflippo, error) {
	e, err := gst.NewElementWithName("frei0r-filter-3dflippo", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilter3Dflippo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// center-position--x- (float64)
//
// Position of the center of rotation on the X axis

func (e *Frei0RFilter3Dflippo) GetCenterPositionX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("center-position--x-")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetCenterPositionX(value float64) error {
	return e.Element.SetProperty("center-position--x-", value)
}

// center-position--y- (float64)
//
// Position of the center of rotation on the Y axis

func (e *Frei0RFilter3Dflippo) GetCenterPositionY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("center-position--y-")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetCenterPositionY(value float64) error {
	return e.Element.SetProperty("center-position--y-", value)
}

// don-t-blank-mask (bool)
//
// Mask for frame transposition is not blanked, so a trace of old transpositions is maintained

func (e *Frei0RFilter3Dflippo) GetDonTBlankMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("don-t-blank-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetDonTBlankMask(value bool) error {
	return e.Element.SetProperty("don-t-blank-mask", value)
}

// fill-with-image-or-black (bool)
//
// If true, pixels that are not transposed are black, otherwise, they are copied with the original

func (e *Frei0RFilter3Dflippo) GetFillWithImageOrBlack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fill-with-image-or-black")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetFillWithImageOrBlack(value bool) error {
	return e.Element.SetProperty("fill-with-image-or-black", value)
}

// invert-rotation-assignment (bool)
//
// If true, when mapping rotation, make inverted (wrong) assignment

func (e *Frei0RFilter3Dflippo) GetInvertRotationAssignment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert-rotation-assignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetInvertRotationAssignment(value bool) error {
	return e.Element.SetProperty("invert-rotation-assignment", value)
}

// x-axis-rotation (float64)
//
// Rotation on the X axis

func (e *Frei0RFilter3Dflippo) GetXAxisRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-axis-rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetXAxisRotation(value float64) error {
	return e.Element.SetProperty("x-axis-rotation", value)
}

// x-axis-rotation-rate (float64)
//
// Rotation rate on the X axis

func (e *Frei0RFilter3Dflippo) GetXAxisRotationRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-axis-rotation-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetXAxisRotationRate(value float64) error {
	return e.Element.SetProperty("x-axis-rotation-rate", value)
}

// y-axis-rotation (float64)
//
// Rotation on the Y axis

func (e *Frei0RFilter3Dflippo) GetYAxisRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-axis-rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetYAxisRotation(value float64) error {
	return e.Element.SetProperty("y-axis-rotation", value)
}

// y-axis-rotation-rate (float64)
//
// Rotation rate on the Y axis

func (e *Frei0RFilter3Dflippo) GetYAxisRotationRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-axis-rotation-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetYAxisRotationRate(value float64) error {
	return e.Element.SetProperty("y-axis-rotation-rate", value)
}

// z-axis-rotation (float64)
//
// Rotation on the Z axis

func (e *Frei0RFilter3Dflippo) GetZAxisRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("z-axis-rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetZAxisRotation(value float64) error {
	return e.Element.SetProperty("z-axis-rotation", value)
}

// z-axis-rotation-rate (float64)
//
// Rotation rate on the Z axis

func (e *Frei0RFilter3Dflippo) GetZAxisRotationRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("z-axis-rotation-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilter3Dflippo) SetZAxisRotationRate(value float64) error {
	return e.Element.SetProperty("z-axis-rotation-rate", value)
}

