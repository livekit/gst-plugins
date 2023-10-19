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

type Frei0RFilterBgsubtract0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterBgsubtract0R() (*Frei0RFilterBgsubtract0R, error) {
	e, err := gst.NewElement("frei0r-filter-bgsubtract0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterBgsubtract0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterBgsubtract0RWithName(name string) (*Frei0RFilterBgsubtract0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-bgsubtract0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterBgsubtract0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// blur (float64)
//
// Blur alpha channel by given radius (to remove sharp edges)

func (e *Frei0RFilterBgsubtract0R) GetBlur() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterBgsubtract0R) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

// denoise (bool)
//
// Remove noise

func (e *Frei0RFilterBgsubtract0R) GetDenoise() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("denoise")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterBgsubtract0R) SetDenoise(value bool) error {
	return e.Element.SetProperty("denoise", value)
}

// threshold (float64)
//
// Threshold for difference

func (e *Frei0RFilterBgsubtract0R) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterBgsubtract0R) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

