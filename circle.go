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

type Circle struct {
	*base.GstBaseTransform
}

func NewCircle() (*Circle, error) {
	e, err := gst.NewElement("circle")
	if err != nil {
		return nil, err
	}
	return &Circle{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCircleWithName(name string) (*Circle, error) {
	e, err := gst.NewElementWithName("circle", name)
	if err != nil {
		return nil, err
	}
	return &Circle{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// angle (float64)
//
// Angle at which the arc starts in radians

func (e *Circle) GetAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Circle) SetAngle(value float64) error {
	return e.Element.SetProperty("angle", value)
}

// height (int)
//
// Height of the arc

func (e *Circle) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Circle) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// spread-angle (float64)
//
// Length of the arc in radians

func (e *Circle) GetSpreadAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spread-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Circle) SetSpreadAngle(value float64) error {
	return e.Element.SetProperty("spread-angle", value)
}

