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

type Faceoverlay struct {
	Element *gst.Element
}

func NewFaceoverlay() (*Faceoverlay, error) {
	e, err := gst.NewElement("faceoverlay")
	if err != nil {
		return nil, err
	}
	return &Faceoverlay{Element: e}, nil
}

func NewFaceoverlayWithName(name string) (*Faceoverlay, error) {
	e, err := gst.NewElementWithName("faceoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Faceoverlay{Element: e}, nil
}

// ----- Properties -----

// h (float32)
//
// Specify image height relative to face height.

func (e *Faceoverlay) GetH() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("h")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Faceoverlay) SetH(value float32) error {
	return e.Element.SetProperty("h", value)
}

// location (string)
//
// Location of SVG file to use for face overlay

func (e *Faceoverlay) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Faceoverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// w (float32)
//
// Specify image width relative to face width.

func (e *Faceoverlay) GetW() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("w")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Faceoverlay) SetW(value float32) error {
	return e.Element.SetProperty("w", value)
}

// x (float32)
//
// Specify image x relative to detected face x.

func (e *Faceoverlay) GetX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Faceoverlay) SetX(value float32) error {
	return e.Element.SetProperty("x", value)
}

// y (float32)
//
// Specify image y relative to detected face y.

func (e *Faceoverlay) GetY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Faceoverlay) SetY(value float32) error {
	return e.Element.SetProperty("y", value)
}

