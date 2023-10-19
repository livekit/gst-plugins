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

type Gloverlay struct {
	*base.GstBaseTransform
}

func NewGloverlay() (*Gloverlay, error) {
	e, err := gst.NewElement("gloverlay")
	if err != nil {
		return nil, err
	}
	return &Gloverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGloverlayWithName(name string) (*Gloverlay, error) {
	e, err := gst.NewElementWithName("gloverlay", name)
	if err != nil {
		return nil, err
	}
	return &Gloverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float64)
//
// Global alpha of overlay image

func (e *Gloverlay) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Gloverlay) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// location (string)
//
// Location of image file to overlay

func (e *Gloverlay) GetLocation() (string, error) {
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

func (e *Gloverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// offset-x (int)
//
// For positive value, horizontal offset of overlay image in pixels from left of video image. For negative value, horizontal offset of overlay image in pixels from right of video image

func (e *Gloverlay) GetOffsetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Gloverlay) SetOffsetX(value int) error {
	return e.Element.SetProperty("offset-x", value)
}

// offset-y (int)
//
// For positive value, vertical offset of overlay image in pixels from top of video image. For negative value, vertical offset of overlay image in pixels from bottom of video image

func (e *Gloverlay) GetOffsetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Gloverlay) SetOffsetY(value int) error {
	return e.Element.SetProperty("offset-y", value)
}

// overlay-height (int)
//
// Height of overlay image in pixels (0 = same as overlay image)

func (e *Gloverlay) GetOverlayHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Gloverlay) SetOverlayHeight(value int) error {
	return e.Element.SetProperty("overlay-height", value)
}

// overlay-width (int)
//
// Width of overlay image in pixels (0 = same as overlay image)

func (e *Gloverlay) GetOverlayWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Gloverlay) SetOverlayWidth(value int) error {
	return e.Element.SetProperty("overlay-width", value)
}

// relative-x (float64)
//
// Horizontal offset of overlay image in fractions of video image width, from top-left corner of video image

func (e *Gloverlay) GetRelativeX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("relative-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Gloverlay) SetRelativeX(value float64) error {
	return e.Element.SetProperty("relative-x", value)
}

// relative-y (float64)
//
// Vertical offset of overlay image in fractions of video image height, from top-left corner of video image

func (e *Gloverlay) GetRelativeY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("relative-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Gloverlay) SetRelativeY(value float64) error {
	return e.Element.SetProperty("relative-y", value)
}

