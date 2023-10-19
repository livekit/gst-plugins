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

type Frei0RFilterAlphaspot struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterAlphaspot() (*Frei0RFilterAlphaspot, error) {
	e, err := gst.NewElement("frei0r-filter-alphaspot")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterAlphaspot{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterAlphaspotWithName(name string) (*Frei0RFilterAlphaspot, error) {
	e, err := gst.NewElementWithName("frei0r-filter-alphaspot", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterAlphaspot{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// max (float64)

func (e *Frei0RFilterAlphaspot) GetMax() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("max")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetMax(value float64) error {
	return e.Element.SetProperty("max", value)
}

// min (float64)

func (e *Frei0RFilterAlphaspot) GetMin() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("min")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetMin(value float64) error {
	return e.Element.SetProperty("min", value)
}

// operation (float64)

func (e *Frei0RFilterAlphaspot) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

// position-x (float64)

func (e *Frei0RFilterAlphaspot) GetPositionX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetPositionX(value float64) error {
	return e.Element.SetProperty("position-x", value)
}

// position-y (float64)

func (e *Frei0RFilterAlphaspot) GetPositionY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetPositionY(value float64) error {
	return e.Element.SetProperty("position-y", value)
}

// shape (float64)

func (e *Frei0RFilterAlphaspot) GetShape() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shape")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetShape(value float64) error {
	return e.Element.SetProperty("shape", value)
}

// size-x (float64)

func (e *Frei0RFilterAlphaspot) GetSizeX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetSizeX(value float64) error {
	return e.Element.SetProperty("size-x", value)
}

// size-y (float64)

func (e *Frei0RFilterAlphaspot) GetSizeY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetSizeY(value float64) error {
	return e.Element.SetProperty("size-y", value)
}

// tilt (float64)

func (e *Frei0RFilterAlphaspot) GetTilt() (float64, error) {
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

func (e *Frei0RFilterAlphaspot) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

// transition-width (float64)

func (e *Frei0RFilterAlphaspot) GetTransitionWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("transition-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphaspot) SetTransitionWidth(value float64) error {
	return e.Element.SetProperty("transition-width", value)
}

