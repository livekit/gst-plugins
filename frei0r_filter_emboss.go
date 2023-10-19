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

type Frei0RFilterEmboss struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterEmboss() (*Frei0RFilterEmboss, error) {
	e, err := gst.NewElement("frei0r-filter-emboss")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterEmboss{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterEmbossWithName(name string) (*Frei0RFilterEmboss, error) {
	e, err := gst.NewElementWithName("frei0r-filter-emboss", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterEmboss{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// azimuth (float64)
//
// Light direction

func (e *Frei0RFilterEmboss) GetAzimuth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("azimuth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterEmboss) SetAzimuth(value float64) error {
	return e.Element.SetProperty("azimuth", value)
}

// elevation (float64)
//
// Background lightness

func (e *Frei0RFilterEmboss) GetElevation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("elevation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterEmboss) SetElevation(value float64) error {
	return e.Element.SetProperty("elevation", value)
}

// width45 (float64)
//
// Bump height

func (e *Frei0RFilterEmboss) GetWidth45() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width45")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterEmboss) SetWidth45(value float64) error {
	return e.Element.SetProperty("width45", value)
}

