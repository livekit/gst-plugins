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

type Square struct {
	*base.GstBaseTransform
}

func NewSquare() (*Square, error) {
	e, err := gst.NewElement("square")
	if err != nil {
		return nil, err
	}
	return &Square{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSquareWithName(name string) (*Square, error) {
	e, err := gst.NewElementWithName("square", name)
	if err != nil {
		return nil, err
	}
	return &Square{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// height (float64)
//
// Height of the square, relative to the frame height

func (e *Square) GetHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Square) SetHeight(value float64) error {
	return e.Element.SetProperty("height", value)
}

// width (float64)
//
// Width of the square, relative to the frame width

func (e *Square) GetWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Square) SetWidth(value float64) error {
	return e.Element.SetProperty("width", value)
}

// zoom (float64)
//
// Zoom amount in the center region

func (e *Square) GetZoom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Square) SetZoom(value float64) error {
	return e.Element.SetProperty("zoom", value)
}

