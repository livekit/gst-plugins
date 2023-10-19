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

type Glcolorbalance struct {
	*base.GstBaseTransform
}

func NewGlcolorbalance() (*Glcolorbalance, error) {
	e, err := gst.NewElement("glcolorbalance")
	if err != nil {
		return nil, err
	}
	return &Glcolorbalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGlcolorbalanceWithName(name string) (*Glcolorbalance, error) {
	e, err := gst.NewElementWithName("glcolorbalance", name)
	if err != nil {
		return nil, err
	}
	return &Glcolorbalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// brightness (float64)
//
// brightness

func (e *Glcolorbalance) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glcolorbalance) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (float64)
//
// contrast

func (e *Glcolorbalance) GetContrast() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glcolorbalance) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

// hue (float64)
//
// hue

func (e *Glcolorbalance) GetHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glcolorbalance) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

// saturation (float64)
//
// saturation

func (e *Glcolorbalance) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glcolorbalance) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

