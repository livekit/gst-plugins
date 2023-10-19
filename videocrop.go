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

type Videocrop struct {
	*base.GstBaseTransform
}

func NewVideocrop() (*Videocrop, error) {
	e, err := gst.NewElement("videocrop")
	if err != nil {
		return nil, err
	}
	return &Videocrop{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideocropWithName(name string) (*Videocrop, error) {
	e, err := gst.NewElementWithName("videocrop", name)
	if err != nil {
		return nil, err
	}
	return &Videocrop{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bottom (int)
//
// Pixels to crop at bottom (-1 to auto-crop)

func (e *Videocrop) GetBottom() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videocrop) SetBottom(value int) error {
	return e.Element.SetProperty("bottom", value)
}

// left (int)
//
// Pixels to crop at left (-1 to auto-crop)

func (e *Videocrop) GetLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videocrop) SetLeft(value int) error {
	return e.Element.SetProperty("left", value)
}

// right (int)
//
// Pixels to crop at right (-1 to auto-crop)

func (e *Videocrop) GetRight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("right")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videocrop) SetRight(value int) error {
	return e.Element.SetProperty("right", value)
}

// top (int)
//
// Pixels to crop at top (-1 to auto-crop)

func (e *Videocrop) GetTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videocrop) SetTop(value int) error {
	return e.Element.SetProperty("top", value)
}

