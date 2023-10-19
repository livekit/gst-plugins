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

type Frei0RFilterWhiteBalance struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterWhiteBalance() (*Frei0RFilterWhiteBalance, error) {
	e, err := gst.NewElement("frei0r-filter-white-balance")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterWhiteBalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterWhiteBalanceWithName(name string) (*Frei0RFilterWhiteBalance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-white-balance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterWhiteBalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// green-tint (float64)
//
// Adjust the level of green.

func (e *Frei0RFilterWhiteBalance) GetGreenTint() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("green-tint")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterWhiteBalance) SetGreenTint(value float64) error {
	return e.Element.SetProperty("green-tint", value)
}

// neutral-color-b (float32)
//
// Choose a color from the source image that should be white.

func (e *Frei0RFilterWhiteBalance) GetNeutralColorB() (float32, error) {
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

func (e *Frei0RFilterWhiteBalance) SetNeutralColorB(value float32) error {
	return e.Element.SetProperty("neutral-color-b", value)
}

// neutral-color-g (float32)
//
// Choose a color from the source image that should be white.

func (e *Frei0RFilterWhiteBalance) GetNeutralColorG() (float32, error) {
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

func (e *Frei0RFilterWhiteBalance) SetNeutralColorG(value float32) error {
	return e.Element.SetProperty("neutral-color-g", value)
}

// neutral-color-r (float32)
//
// Choose a color from the source image that should be white.

func (e *Frei0RFilterWhiteBalance) GetNeutralColorR() (float32, error) {
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

func (e *Frei0RFilterWhiteBalance) SetNeutralColorR(value float32) error {
	return e.Element.SetProperty("neutral-color-r", value)
}

