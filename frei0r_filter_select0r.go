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

type Frei0RFilterSelect0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterSelect0R() (*Frei0RFilterSelect0R, error) {
	e, err := gst.NewElement("frei0r-filter-select0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSelect0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterSelect0RWithName(name string) (*Frei0RFilterSelect0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-select0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSelect0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color-to-select-b (float32)

func (e *Frei0RFilterSelect0R) GetColorToSelectB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-to-select-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetColorToSelectB(value float32) error {
	return e.Element.SetProperty("color-to-select-b", value)
}

// color-to-select-g (float32)

func (e *Frei0RFilterSelect0R) GetColorToSelectG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-to-select-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetColorToSelectG(value float32) error {
	return e.Element.SetProperty("color-to-select-g", value)
}

// color-to-select-r (float32)

func (e *Frei0RFilterSelect0R) GetColorToSelectR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-to-select-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetColorToSelectR(value float32) error {
	return e.Element.SetProperty("color-to-select-r", value)
}

// delta-b---i---i (float64)

func (e *Frei0RFilterSelect0R) GetDeltaBII() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delta-b---i---i")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetDeltaBII(value float64) error {
	return e.Element.SetProperty("delta-b---i---i", value)
}

// delta-g---b---chroma (float64)

func (e *Frei0RFilterSelect0R) GetDeltaGBChroma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delta-g---b---chroma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetDeltaGBChroma(value float64) error {
	return e.Element.SetProperty("delta-g---b---chroma", value)
}

// delta-r---a---hue (float64)

func (e *Frei0RFilterSelect0R) GetDeltaRAHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("delta-r---a---hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetDeltaRAHue(value float64) error {
	return e.Element.SetProperty("delta-r---a---hue", value)
}

// edge-mode (float64)

func (e *Frei0RFilterSelect0R) GetEdgeMode() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("edge-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetEdgeMode(value float64) error {
	return e.Element.SetProperty("edge-mode", value)
}

// invert-selection (bool)

func (e *Frei0RFilterSelect0R) GetInvertSelection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert-selection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetInvertSelection(value bool) error {
	return e.Element.SetProperty("invert-selection", value)
}

// operation (float64)

func (e *Frei0RFilterSelect0R) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

// selection-subspace (float64)

func (e *Frei0RFilterSelect0R) GetSelectionSubspace() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("selection-subspace")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetSelectionSubspace(value float64) error {
	return e.Element.SetProperty("selection-subspace", value)
}

// slope (float64)

func (e *Frei0RFilterSelect0R) GetSlope() (float64, error) {
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

func (e *Frei0RFilterSelect0R) SetSlope(value float64) error {
	return e.Element.SetProperty("slope", value)
}

// subspace-shape (float64)

func (e *Frei0RFilterSelect0R) GetSubspaceShape() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("subspace-shape")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSelect0R) SetSubspaceShape(value float64) error {
	return e.Element.SetProperty("subspace-shape", value)
}

