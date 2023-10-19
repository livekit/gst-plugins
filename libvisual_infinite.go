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

type LibvisualInfinite struct {
	Element *gst.Element
}

func NewLibvisualInfinite() (*LibvisualInfinite, error) {
	e, err := gst.NewElement("libvisual_infinite")
	if err != nil {
		return nil, err
	}
	return &LibvisualInfinite{Element: e}, nil
}

func NewLibvisualInfiniteWithName(name string) (*LibvisualInfinite, error) {
	e, err := gst.NewElementWithName("libvisual_infinite", name)
	if err != nil {
		return nil, err
	}
	return &LibvisualInfinite{Element: e}, nil
}

// ----- Properties -----

// shade-amount (uint)
//
// Shading color to use (big-endian ARGB)

func (e *LibvisualInfinite) GetShadeAmount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shade-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *LibvisualInfinite) SetShadeAmount(value uint) error {
	return e.Element.SetProperty("shade-amount", value)
}

// shader (GstAudioVisualizerShader)
//
// Shader function to apply on each frame

func (e *LibvisualInfinite) GetShader() (interface{}, error) {
	return e.Element.GetProperty("shader")
}

func (e *LibvisualInfinite) SetShader(value interface{}) error {
	return e.Element.SetProperty("shader", value)
}

