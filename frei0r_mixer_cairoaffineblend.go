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
)

type Frei0RMixerCairoaffineblend struct {
	Element *gst.Element
}

func NewFrei0RMixerCairoaffineblend() (*Frei0RMixerCairoaffineblend, error) {
	e, err := gst.NewElement("frei0r-mixer-cairoaffineblend")
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerCairoaffineblend{Element: e}, nil
}

func NewFrei0RMixerCairoaffineblendWithName(name string) (*Frei0RMixerCairoaffineblend, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-cairoaffineblend", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerCairoaffineblend{Element: e}, nil
}

// ----- Properties -----

// anchor-x (float64)
//
// X position of rotation center within the second input

func (e *Frei0RMixerCairoaffineblend) GetAnchorX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("anchor-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetAnchorX(value float64) error {
	return e.Element.SetProperty("anchor-x", value)
}

// anchor-y (float64)
//
// Y position of rotation center within the second input

func (e *Frei0RMixerCairoaffineblend) GetAnchorY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("anchor-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetAnchorY(value float64) error {
	return e.Element.SetProperty("anchor-y", value)
}

// blend-mode (string)
//
// Blend mode used to compose image. Accepted values: 'normal', 'add', 'saturate', 'multiply', 'screen', 'overlay', 'darken', 'lighten', 'colordodge', 'colorburn', 'hardlight', 'softlight', 'difference', 'exclusion', 'hslhue', 'hslsaturation', 'hslcolor', 'hslluminosity'

func (e *Frei0RMixerCairoaffineblend) GetBlendMode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("blend-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetBlendMode(value string) error {
	return e.Element.SetProperty("blend-mode", value)
}

// opacity (float64)
//
// Opacity of second input

func (e *Frei0RMixerCairoaffineblend) GetOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetOpacity(value float64) error {
	return e.Element.SetProperty("opacity", value)
}

// rotation (float64)
//
// Rotation of second input, value interperted as range 0 - 360

func (e *Frei0RMixerCairoaffineblend) GetRotation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetRotation(value float64) error {
	return e.Element.SetProperty("rotation", value)
}

// x (float64)
//
// X position of second input, value interperted as range -2width - 3width

func (e *Frei0RMixerCairoaffineblend) GetX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetX(value float64) error {
	return e.Element.SetProperty("x", value)
}

// x-scale (float64)
//
// X scale of second input, value interperted as range 0 - 5

func (e *Frei0RMixerCairoaffineblend) GetXScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetXScale(value float64) error {
	return e.Element.SetProperty("x-scale", value)
}

// y (float64)
//
// Y position of second input, value interperted as range -2height - 3height

func (e *Frei0RMixerCairoaffineblend) GetY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetY(value float64) error {
	return e.Element.SetProperty("y", value)
}

// y-scale (float64)
//
// Y scale of second input, value interperted as range 0 - 5

func (e *Frei0RMixerCairoaffineblend) GetYScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoaffineblend) SetYScale(value float64) error {
	return e.Element.SetProperty("y-scale", value)
}

