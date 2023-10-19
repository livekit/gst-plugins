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

type Theoradec struct {
	Element *gst.Element
}

func NewTheoradec() (*Theoradec, error) {
	e, err := gst.NewElement("theoradec")
	if err != nil {
		return nil, err
	}
	return &Theoradec{Element: e}, nil
}

func NewTheoradecWithName(name string) (*Theoradec, error) {
	e, err := gst.NewElementWithName("theoradec", name)
	if err != nil {
		return nil, err
	}
	return &Theoradec{Element: e}, nil
}

// ----- Properties -----

// visualize-bit-usage (int)
//
// Sets the bitstream breakdown visualization mode. Values influence the width of the bit usage bars to show

func (e *Theoradec) GetVisualizeBitUsage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-bit-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoradec) SetVisualizeBitUsage(value int) error {
	return e.Element.SetProperty("visualize-bit-usage", value)
}

// visualize-macroblock-modes (int)
//
// Show macroblock mode selection overlaid on image. Value gives a mask for macroblock (MB) modes to show

func (e *Theoradec) GetVisualizeMacroblockModes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-macroblock-modes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoradec) SetVisualizeMacroblockModes(value int) error {
	return e.Element.SetProperty("visualize-macroblock-modes", value)
}

// visualize-motion-vectors (int)
//
// Show motion vector selection overlaid on image. Value gives a mask for motion vector (MV) modes to show

func (e *Theoradec) GetVisualizeMotionVectors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-motion-vectors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoradec) SetVisualizeMotionVectors(value int) error {
	return e.Element.SetProperty("visualize-motion-vectors", value)
}

// visualize-quantization-modes (int)
//
// Show adaptive quantization mode selection overlaid on image. Value gives a mask for quantization (QI) modes to show

func (e *Theoradec) GetVisualizeQuantizationModes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-quantization-modes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoradec) SetVisualizeQuantizationModes(value int) error {
	return e.Element.SetProperty("visualize-quantization-modes", value)
}

