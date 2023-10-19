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

type Frei0RFilterColorDistance struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterColorDistance() (*Frei0RFilterColorDistance, error) {
	e, err := gst.NewElement("frei0r-filter-color-distance")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColorDistance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterColorDistanceWithName(name string) (*Frei0RFilterColorDistance, error) {
	e, err := gst.NewElementWithName("frei0r-filter-color-distance", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColorDistance{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color-b (float32)
//
// The Source Color

func (e *Frei0RFilterColorDistance) GetColorB() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterColorDistance) SetColorB(value float32) error {
	return e.Element.SetProperty("color-b", value)
}

// color-g (float32)
//
// The Source Color

func (e *Frei0RFilterColorDistance) GetColorG() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterColorDistance) SetColorG(value float32) error {
	return e.Element.SetProperty("color-g", value)
}

// color-r (float32)
//
// The Source Color

func (e *Frei0RFilterColorDistance) GetColorR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterColorDistance) SetColorR(value float32) error {
	return e.Element.SetProperty("color-r", value)
}

