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

type Frei0RFilterTint0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterTint0R() (*Frei0RFilterTint0R, error) {
	e, err := gst.NewElement("frei0r-filter-tint0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterTint0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterTint0RWithName(name string) (*Frei0RFilterTint0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-tint0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterTint0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// map-black-to-b (float32)
//
// The color to map source color with null luminance

func (e *Frei0RFilterTint0R) GetMapBlackToB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-black-to-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetMapBlackToB(value float32) error {
	return e.Element.SetProperty("map-black-to-b", value)
}

// map-black-to-g (float32)
//
// The color to map source color with null luminance

func (e *Frei0RFilterTint0R) GetMapBlackToG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-black-to-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetMapBlackToG(value float32) error {
	return e.Element.SetProperty("map-black-to-g", value)
}

// map-black-to-r (float32)
//
// The color to map source color with null luminance

func (e *Frei0RFilterTint0R) GetMapBlackToR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-black-to-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetMapBlackToR(value float32) error {
	return e.Element.SetProperty("map-black-to-r", value)
}

// map-white-to-b (float32)
//
// The color to map source color with full luminance

func (e *Frei0RFilterTint0R) GetMapWhiteToB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-white-to-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetMapWhiteToB(value float32) error {
	return e.Element.SetProperty("map-white-to-b", value)
}

// map-white-to-g (float32)
//
// The color to map source color with full luminance

func (e *Frei0RFilterTint0R) GetMapWhiteToG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-white-to-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetMapWhiteToG(value float32) error {
	return e.Element.SetProperty("map-white-to-g", value)
}

// map-white-to-r (float32)
//
// The color to map source color with full luminance

func (e *Frei0RFilterTint0R) GetMapWhiteToR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("map-white-to-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetMapWhiteToR(value float32) error {
	return e.Element.SetProperty("map-white-to-r", value)
}

// tint-amount (float64)
//
// Amount of color

func (e *Frei0RFilterTint0R) GetTintAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tint-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterTint0R) SetTintAmount(value float64) error {
	return e.Element.SetProperty("tint-amount", value)
}

