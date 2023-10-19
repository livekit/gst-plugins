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

type Frei0RFilterBluescreen0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterBluescreen0R() (*Frei0RFilterBluescreen0R, error) {
	e, err := gst.NewElement("frei0r-filter-bluescreen0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterBluescreen0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterBluescreen0RWithName(name string) (*Frei0RFilterBluescreen0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-bluescreen0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterBluescreen0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color-b (float32)
//
// The color to make transparent (B G R)

func (e *Frei0RFilterBluescreen0R) GetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterBluescreen0R) SetColorB(value float32) error {
	return e.Element.SetProperty("color-b", value)
}

// color-g (float32)
//
// The color to make transparent (B G R)

func (e *Frei0RFilterBluescreen0R) GetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterBluescreen0R) SetColorG(value float32) error {
	return e.Element.SetProperty("color-g", value)
}

// color-r (float32)
//
// The color to make transparent (B G R)

func (e *Frei0RFilterBluescreen0R) GetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterBluescreen0R) SetColorR(value float32) error {
	return e.Element.SetProperty("color-r", value)
}

// distance (float64)
//
// Distance to Color (127 is good)

func (e *Frei0RFilterBluescreen0R) GetDistance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterBluescreen0R) SetDistance(value float64) error {
	return e.Element.SetProperty("distance", value)
}

