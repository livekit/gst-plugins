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

type Videobalance struct {
	*base.GstBaseTransform
}

func NewVideobalance() (*Videobalance, error) {
	e, err := gst.NewElement("videobalance")
	if err != nil {
		return nil, err
	}
	return &Videobalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideobalanceWithName(name string) (*Videobalance, error) {
	e, err := gst.NewElementWithName("videobalance", name)
	if err != nil {
		return nil, err
	}
	return &Videobalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// brightness (float64)
//
// brightness

func (e *Videobalance) GetBrightness() (float64, error) {
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

func (e *Videobalance) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (float64)
//
// contrast

func (e *Videobalance) GetContrast() (float64, error) {
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

func (e *Videobalance) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

// hue (float64)
//
// hue

func (e *Videobalance) GetHue() (float64, error) {
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

func (e *Videobalance) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

// saturation (float64)
//
// saturation

func (e *Videobalance) GetSaturation() (float64, error) {
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

func (e *Videobalance) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

