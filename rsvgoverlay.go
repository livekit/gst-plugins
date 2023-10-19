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

type Rsvgoverlay struct {
	*base.GstBaseTransform
}

func NewRsvgoverlay() (*Rsvgoverlay, error) {
	e, err := gst.NewElement("rsvgoverlay")
	if err != nil {
		return nil, err
	}
	return &Rsvgoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRsvgoverlayWithName(name string) (*Rsvgoverlay, error) {
	e, err := gst.NewElementWithName("rsvgoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Rsvgoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// data (string)
//
// SVG data.

func (e *Rsvgoverlay) GetData() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("data")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetData(value string) error {
	return e.Element.SetProperty("data", value)
}

// fit-to-frame (bool)
//
// Fit the SVG to fill the whole frame.

func (e *Rsvgoverlay) GetFitToFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fit-to-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetFitToFrame(value bool) error {
	return e.Element.SetProperty("fit-to-frame", value)
}

// height (int)
//
// Specify a height in pixels.

func (e *Rsvgoverlay) GetHeight() (int, error) {
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

func (e *Rsvgoverlay) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// height-relative (float32)
//
// Specify a height relative to the display size.

func (e *Rsvgoverlay) GetHeightRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("height-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetHeightRelative(value float32) error {
	return e.Element.SetProperty("height-relative", value)
}

// location (string)
//
// SVG file location.

func (e *Rsvgoverlay) GetLocation() (string, error) {
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

func (e *Rsvgoverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// width (int)
//
// Specify a width in pixels.

func (e *Rsvgoverlay) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

// width-relative (float32)
//
// Specify a width relative to the display size.

func (e *Rsvgoverlay) GetWidthRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("width-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetWidthRelative(value float32) error {
	return e.Element.SetProperty("width-relative", value)
}

// x (int)
//
// Specify an x offset.

func (e *Rsvgoverlay) GetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetX(value int) error {
	return e.Element.SetProperty("x", value)
}

// x-relative (float32)
//
// Specify an x offset relative to the display size.

func (e *Rsvgoverlay) GetXRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("x-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetXRelative(value float32) error {
	return e.Element.SetProperty("x-relative", value)
}

// y (int)
//
// Specify a y offset.

func (e *Rsvgoverlay) GetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetY(value int) error {
	return e.Element.SetProperty("y", value)
}

// y-relative (float32)
//
// Specify a y offset relative to the display size.

func (e *Rsvgoverlay) GetYRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("y-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Rsvgoverlay) SetYRelative(value float32) error {
	return e.Element.SetProperty("y-relative", value)
}

