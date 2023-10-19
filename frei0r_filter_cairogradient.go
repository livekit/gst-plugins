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

type Frei0RFilterCairogradient struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterCairogradient() (*Frei0RFilterCairogradient, error) {
	e, err := gst.NewElement("frei0r-filter-cairogradient")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCairogradient{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterCairogradientWithName(name string) (*Frei0RFilterCairogradient, error) {
	e, err := gst.NewElementWithName("frei0r-filter-cairogradient", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCairogradient{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// blend-mode (string)
//
// Blend mode used to compose gradient on image. Accepted values: 'normal', 'add', 'saturate', 'multiply', 'screen', 'overlay', 'darken', 'lighten', 'colordodge', 'colorburn', 'hardlight', 'softlight', 'difference', 'exclusion', 'hslhue', 'hslsaturation', 'hslcolor', 'hslluminosity'

func (e *Frei0RFilterCairogradient) GetBlendMode() (string, error) {
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

func (e *Frei0RFilterCairogradient) SetBlendMode(value string) error {
	return e.Element.SetProperty("blend-mode", value)
}

// end-color-b (float32)
//
// Second color of the gradient

func (e *Frei0RFilterCairogradient) GetEndColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("end-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetEndColorB(value float32) error {
	return e.Element.SetProperty("end-color-b", value)
}

// end-color-g (float32)
//
// Second color of the gradient

func (e *Frei0RFilterCairogradient) GetEndColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("end-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetEndColorG(value float32) error {
	return e.Element.SetProperty("end-color-g", value)
}

// end-color-r (float32)
//
// Second color of the gradient

func (e *Frei0RFilterCairogradient) GetEndColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("end-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetEndColorR(value float32) error {
	return e.Element.SetProperty("end-color-r", value)
}

// end-opacity (float64)
//
// Opacity of the second color of the gradient

func (e *Frei0RFilterCairogradient) GetEndOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("end-opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetEndOpacity(value float64) error {
	return e.Element.SetProperty("end-opacity", value)
}

// end-x (float64)
//
// X position of the end point of the gradient

func (e *Frei0RFilterCairogradient) GetEndX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("end-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetEndX(value float64) error {
	return e.Element.SetProperty("end-x", value)
}

// end-y (float64)
//
// Y position of the end point of the gradient

func (e *Frei0RFilterCairogradient) GetEndY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("end-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetEndY(value float64) error {
	return e.Element.SetProperty("end-y", value)
}

// offset (float64)
//
// Position of first color in the line connecting gradient ends, really useful only for radial gradient

func (e *Frei0RFilterCairogradient) GetOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetOffset(value float64) error {
	return e.Element.SetProperty("offset", value)
}

// pattern (string)
//
// Linear or radial gradient

func (e *Frei0RFilterCairogradient) GetPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetPattern(value string) error {
	return e.Element.SetProperty("pattern", value)
}

// start-color-b (float32)
//
// First color of the gradient

func (e *Frei0RFilterCairogradient) GetStartColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("start-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetStartColorB(value float32) error {
	return e.Element.SetProperty("start-color-b", value)
}

// start-color-g (float32)
//
// First color of the gradient

func (e *Frei0RFilterCairogradient) GetStartColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("start-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetStartColorG(value float32) error {
	return e.Element.SetProperty("start-color-g", value)
}

// start-color-r (float32)
//
// First color of the gradient

func (e *Frei0RFilterCairogradient) GetStartColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("start-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetStartColorR(value float32) error {
	return e.Element.SetProperty("start-color-r", value)
}

// start-opacity (float64)
//
// Opacity of the first color of the gradient

func (e *Frei0RFilterCairogradient) GetStartOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("start-opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetStartOpacity(value float64) error {
	return e.Element.SetProperty("start-opacity", value)
}

// start-x (float64)
//
// X position of the start point of the gradient

func (e *Frei0RFilterCairogradient) GetStartX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("start-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetStartX(value float64) error {
	return e.Element.SetProperty("start-x", value)
}

// start-y (float64)
//
// Y position of the start point of the gradient

func (e *Frei0RFilterCairogradient) GetStartY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("start-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairogradient) SetStartY(value float64) error {
	return e.Element.SetProperty("start-y", value)
}

