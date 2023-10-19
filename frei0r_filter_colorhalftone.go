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

type Frei0RFilterColorhalftone struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterColorhalftone() (*Frei0RFilterColorhalftone, error) {
	e, err := gst.NewElement("frei0r-filter-colorhalftone")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColorhalftone{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterColorhalftoneWithName(name string) (*Frei0RFilterColorhalftone, error) {
	e, err := gst.NewElementWithName("frei0r-filter-colorhalftone", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColorhalftone{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cyan-angle (float64)
//
// Cyan dots angle

func (e *Frei0RFilterColorhalftone) GetCyanAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("cyan-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColorhalftone) SetCyanAngle(value float64) error {
	return e.Element.SetProperty("cyan-angle", value)
}

// dot-radius (float64)
//
// Halftone pattern dot size

func (e *Frei0RFilterColorhalftone) GetDotRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dot-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColorhalftone) SetDotRadius(value float64) error {
	return e.Element.SetProperty("dot-radius", value)
}

// magenta-angle (float64)
//
// Magenta dots angle

func (e *Frei0RFilterColorhalftone) GetMagentaAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("magenta-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColorhalftone) SetMagentaAngle(value float64) error {
	return e.Element.SetProperty("magenta-angle", value)
}

// yellow-angle (float64)
//
// Yellow dots angle

func (e *Frei0RFilterColorhalftone) GetYellowAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("yellow-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterColorhalftone) SetYellowAngle(value float64) error {
	return e.Element.SetProperty("yellow-angle", value)
}

