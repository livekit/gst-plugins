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

type Frei0RFilterKeyspillm0Pup struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterKeyspillm0Pup() (*Frei0RFilterKeyspillm0Pup, error) {
	e, err := gst.NewElement("frei0r-filter-keyspillm0pup")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterKeyspillm0Pup{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterKeyspillm0PupWithName(name string) (*Frei0RFilterKeyspillm0Pup, error) {
	e, err := gst.NewElementWithName("frei0r-filter-keyspillm0pup", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterKeyspillm0Pup{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amount-1 (float64)

func (e *Frei0RFilterKeyspillm0Pup) GetAmount1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetAmount1(value float64) error {
	return e.Element.SetProperty("amount-1", value)
}

// amount-2 (float64)

func (e *Frei0RFilterKeyspillm0Pup) GetAmount2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetAmount2(value float64) error {
	return e.Element.SetProperty("amount-2", value)
}

// hue-gate (float64)
//
// Restrict mask to hues close to key

func (e *Frei0RFilterKeyspillm0Pup) GetHueGate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue-gate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetHueGate(value float64) error {
	return e.Element.SetProperty("hue-gate", value)
}

// key-color-b (float32)
//
// Key color that was used for chroma keying

func (e *Frei0RFilterKeyspillm0Pup) GetKeyColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("key-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetKeyColorB(value float32) error {
	return e.Element.SetProperty("key-color-b", value)
}

// key-color-g (float32)
//
// Key color that was used for chroma keying

func (e *Frei0RFilterKeyspillm0Pup) GetKeyColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("key-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetKeyColorG(value float32) error {
	return e.Element.SetProperty("key-color-g", value)
}

// key-color-r (float32)
//
// Key color that was used for chroma keying

func (e *Frei0RFilterKeyspillm0Pup) GetKeyColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("key-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetKeyColorR(value float32) error {
	return e.Element.SetProperty("key-color-r", value)
}

// mask-to-alpha (bool)
//
// Replace alpha channel with the mask

func (e *Frei0RFilterKeyspillm0Pup) GetMaskToAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask-to-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetMaskToAlpha(value bool) error {
	return e.Element.SetProperty("mask-to-alpha", value)
}

// mask-type (string)
//
// Which mask to apply [0,1,2,3]

func (e *Frei0RFilterKeyspillm0Pup) GetMaskType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mask-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetMaskType(value string) error {
	return e.Element.SetProperty("mask-type", value)
}

// operation-1 (string)
//
// First operation 1 [0,1,2]

func (e *Frei0RFilterKeyspillm0Pup) GetOperation1() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("operation-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetOperation1(value string) error {
	return e.Element.SetProperty("operation-1", value)
}

// operation-2 (string)
//
// Second operation 2 [0,1,2]

func (e *Frei0RFilterKeyspillm0Pup) GetOperation2() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("operation-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetOperation2(value string) error {
	return e.Element.SetProperty("operation-2", value)
}

// saturation-threshold (float64)
//
// Restrict mask to saturated colors

func (e *Frei0RFilterKeyspillm0Pup) GetSaturationThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetSaturationThreshold(value float64) error {
	return e.Element.SetProperty("saturation-threshold", value)
}

// show-mask (bool)
//
// Replace image with the mask

func (e *Frei0RFilterKeyspillm0Pup) GetShowMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetShowMask(value bool) error {
	return e.Element.SetProperty("show-mask", value)
}

// slope (float64)
//
// Range of colors around the key where effect gradually decreases

func (e *Frei0RFilterKeyspillm0Pup) GetSlope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("slope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetSlope(value float64) error {
	return e.Element.SetProperty("slope", value)
}

// target-color-b (float32)
//
// Desired color to replace key residue with

func (e *Frei0RFilterKeyspillm0Pup) GetTargetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("target-color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetTargetColorB(value float32) error {
	return e.Element.SetProperty("target-color-b", value)
}

// target-color-g (float32)
//
// Desired color to replace key residue with

func (e *Frei0RFilterKeyspillm0Pup) GetTargetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("target-color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetTargetColorG(value float32) error {
	return e.Element.SetProperty("target-color-g", value)
}

// target-color-r (float32)
//
// Desired color to replace key residue with

func (e *Frei0RFilterKeyspillm0Pup) GetTargetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("target-color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetTargetColorR(value float32) error {
	return e.Element.SetProperty("target-color-r", value)
}

// tolerance (float64)
//
// Range of colors around the key, where effect is full strength

func (e *Frei0RFilterKeyspillm0Pup) GetTolerance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterKeyspillm0Pup) SetTolerance(value float64) error {
	return e.Element.SetProperty("tolerance", value)
}

