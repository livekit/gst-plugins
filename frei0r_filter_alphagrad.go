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

type Frei0RFilterAlphagrad struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterAlphagrad() (*Frei0RFilterAlphagrad, error) {
	e, err := gst.NewElement("frei0r-filter-alphagrad")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterAlphagrad{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterAlphagradWithName(name string) (*Frei0RFilterAlphagrad, error) {
	e, err := gst.NewElementWithName("frei0r-filter-alphagrad", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterAlphagrad{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// max (float64)

func (e *Frei0RFilterAlphagrad) GetMax() (float64, error) {
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

func (e *Frei0RFilterAlphagrad) SetMax(value float64) error {
	return e.Element.SetProperty("max", value)
}

// min (float64)

func (e *Frei0RFilterAlphagrad) GetMin() (float64, error) {
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

func (e *Frei0RFilterAlphagrad) SetMin(value float64) error {
	return e.Element.SetProperty("min", value)
}

// operation (float64)

func (e *Frei0RFilterAlphagrad) GetOperation() (float64, error) {
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

func (e *Frei0RFilterAlphagrad) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

// position (float64)

func (e *Frei0RFilterAlphagrad) GetPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlphagrad) SetPosition(value float64) error {
	return e.Element.SetProperty("position", value)
}

// tilt (float64)

func (e *Frei0RFilterAlphagrad) GetTilt() (float64, error) {
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

func (e *Frei0RFilterAlphagrad) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

// transition-width (float64)

func (e *Frei0RFilterAlphagrad) GetTransitionWidth() (float64, error) {
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

func (e *Frei0RFilterAlphagrad) SetTransitionWidth(value float64) error {
	return e.Element.SetProperty("transition-width", value)
}

