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

type Frei0RFilterWhiteBalanceLmsSpace struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterWhiteBalanceLmsSpace() (*Frei0RFilterWhiteBalanceLmsSpace, error) {
	e, err := gst.NewElement("frei0r-filter-white-balance--lms-space-")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterWhiteBalanceLmsSpace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterWhiteBalanceLmsSpaceWithName(name string) (*Frei0RFilterWhiteBalanceLmsSpace, error) {
	e, err := gst.NewElementWithName("frei0r-filter-white-balance--lms-space-", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterWhiteBalanceLmsSpace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color-temperature (float64)
//
// Choose an output color temperature, if different from 6500 K.

func (e *Frei0RFilterWhiteBalanceLmsSpace) GetColorTemperature() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color-temperature")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterWhiteBalanceLmsSpace) SetColorTemperature(value float64) error {
	return e.Element.SetProperty("color-temperature", value)
}

// neutral-color-b (float32)
//
// Choose a color from the source image that should be white.

func (e *Frei0RFilterWhiteBalanceLmsSpace) GetNeutralColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterWhiteBalanceLmsSpace) SetNeutralColorB(value float32) error {
	return e.Element.SetProperty("neutral-color-b", value)
}

// neutral-color-g (float32)
//
// Choose a color from the source image that should be white.

func (e *Frei0RFilterWhiteBalanceLmsSpace) GetNeutralColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterWhiteBalanceLmsSpace) SetNeutralColorG(value float32) error {
	return e.Element.SetProperty("neutral-color-g", value)
}

// neutral-color-r (float32)
//
// Choose a color from the source image that should be white.

func (e *Frei0RFilterWhiteBalanceLmsSpace) GetNeutralColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("neutral-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterWhiteBalanceLmsSpace) SetNeutralColorR(value float32) error {
	return e.Element.SetProperty("neutral-color-r", value)
}

