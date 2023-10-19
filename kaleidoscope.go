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

type Kaleidoscope struct {
	*base.GstBaseTransform
}

func NewKaleidoscope() (*Kaleidoscope, error) {
	e, err := gst.NewElement("kaleidoscope")
	if err != nil {
		return nil, err
	}
	return &Kaleidoscope{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewKaleidoscopeWithName(name string) (*Kaleidoscope, error) {
	e, err := gst.NewElementWithName("kaleidoscope", name)
	if err != nil {
		return nil, err
	}
	return &Kaleidoscope{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// angle (float64)
//
// primary angle in radians of the kaleidoscope effect

func (e *Kaleidoscope) GetAngle() (float64, error) {
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

func (e *Kaleidoscope) SetAngle(value float64) error {
	return e.Element.SetProperty("angle", value)
}

// angle2 (float64)
//
// secondary angle in radians of the kaleidoscope effect

func (e *Kaleidoscope) GetAngle2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Kaleidoscope) SetAngle2(value float64) error {
	return e.Element.SetProperty("angle2", value)
}

// sides (int)
//
// Number of sides of the kaleidoscope

func (e *Kaleidoscope) GetSides() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sides")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kaleidoscope) SetSides(value int) error {
	return e.Element.SetProperty("sides", value)
}

