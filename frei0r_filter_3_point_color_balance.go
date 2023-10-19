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

type Frei0RFilter3PointColorBalance struct {
	*base.GstBaseTransform
}

func NewFrei0RFilter3PointColorBalance() (*Frei0RFilter3PointColorBalance, error) {
	e, err := gst.NewElement("frei0r-filter-3-point-color-balance")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilter3PointColorBalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilter3PointColorBalanceWithName(name string) (*Frei0RFilter3PointColorBalance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-3-point-color-balance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilter3PointColorBalance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// black-color-b (float32)
//
// Black color

func (e *Frei0RFilter3PointColorBalance) GetBlackColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("black-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetBlackColorB(value float32) error {
	return e.Element.SetProperty("black-color-b", value)
}

// black-color-g (float32)
//
// Black color

func (e *Frei0RFilter3PointColorBalance) GetBlackColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("black-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetBlackColorG(value float32) error {
	return e.Element.SetProperty("black-color-g", value)
}

// black-color-r (float32)
//
// Black color

func (e *Frei0RFilter3PointColorBalance) GetBlackColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("black-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetBlackColorR(value float32) error {
	return e.Element.SetProperty("black-color-r", value)
}

// gray-color-b (float32)
//
// Gray color

func (e *Frei0RFilter3PointColorBalance) GetGrayColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gray-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetGrayColorB(value float32) error {
	return e.Element.SetProperty("gray-color-b", value)
}

// gray-color-g (float32)
//
// Gray color

func (e *Frei0RFilter3PointColorBalance) GetGrayColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gray-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetGrayColorG(value float32) error {
	return e.Element.SetProperty("gray-color-g", value)
}

// gray-color-r (float32)
//
// Gray color

func (e *Frei0RFilter3PointColorBalance) GetGrayColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gray-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetGrayColorR(value float32) error {
	return e.Element.SetProperty("gray-color-r", value)
}

// source-image-on-left-side (bool)
//
// Source image on left side

func (e *Frei0RFilter3PointColorBalance) GetSourceImageOnLeftSide() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("source-image-on-left-side")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetSourceImageOnLeftSide(value bool) error {
	return e.Element.SetProperty("source-image-on-left-side", value)
}

// split-preview (bool)
//
// Split privew

func (e *Frei0RFilter3PointColorBalance) GetSplitPreview() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("split-preview")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetSplitPreview(value bool) error {
	return e.Element.SetProperty("split-preview", value)
}

// white-color-b (float32)
//
// White color

func (e *Frei0RFilter3PointColorBalance) GetWhiteColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("white-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetWhiteColorB(value float32) error {
	return e.Element.SetProperty("white-color-b", value)
}

// white-color-g (float32)
//
// White color

func (e *Frei0RFilter3PointColorBalance) GetWhiteColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("white-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetWhiteColorG(value float32) error {
	return e.Element.SetProperty("white-color-g", value)
}

// white-color-r (float32)
//
// White color

func (e *Frei0RFilter3PointColorBalance) GetWhiteColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("white-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilter3PointColorBalance) SetWhiteColorR(value float32) error {
	return e.Element.SetProperty("white-color-r", value)
}

