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

type Gdkpixbufoverlay struct {
	*base.GstBaseTransform
}

func NewGdkpixbufoverlay() (*Gdkpixbufoverlay, error) {
	e, err := gst.NewElement("gdkpixbufoverlay")
	if err != nil {
		return nil, err
	}
	return &Gdkpixbufoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGdkpixbufoverlayWithName(name string) (*Gdkpixbufoverlay, error) {
	e, err := gst.NewElementWithName("gdkpixbufoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Gdkpixbufoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float64)
//
// Global alpha of overlay image

func (e *Gdkpixbufoverlay) GetAlpha() (float64, error) {
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

func (e *Gdkpixbufoverlay) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// coef-x (float64)
//
// In absolute positioning mode, the x coordinate of overlay image's
// top-left corner is now given by
// offset-x + (relative-x * overlay_width) + (coef-x * video_width).
// This allows to align the image absolutely and relatively
// to any edge or center position.

func (e *Gdkpixbufoverlay) GetCoefX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("coef-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Gdkpixbufoverlay) SetCoefX(value float64) error {
	return e.Element.SetProperty("coef-x", value)
}

// coef-y (float64)
//
// In absolute positioning mode, the y coordinate of overlay image's
// top-left corner is now given by
// offset-y + (relative-y * overlay_height) + (coef-y * video_height).
// This allows to align the image absolutely and relatively
// to any edge or center position.

func (e *Gdkpixbufoverlay) GetCoefY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("coef-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Gdkpixbufoverlay) SetCoefY(value float64) error {
	return e.Element.SetProperty("coef-y", value)
}

// location (string)
//
// Location of image file to overlay

func (e *Gdkpixbufoverlay) GetLocation() (string, error) {
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

func (e *Gdkpixbufoverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// offset-x (int)
//
// For positive value, horizontal offset of overlay image in pixels from left of video image. For negative value, horizontal offset of overlay image in pixels from right of video image

func (e *Gdkpixbufoverlay) GetOffsetX() (int, error) {
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

func (e *Gdkpixbufoverlay) SetOffsetX(value int) error {
	return e.Element.SetProperty("offset-x", value)
}

// offset-y (int)
//
// For positive value, vertical offset of overlay image in pixels from top of video image. For negative value, vertical offset of overlay image in pixels from bottom of video image

func (e *Gdkpixbufoverlay) GetOffsetY() (int, error) {
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

func (e *Gdkpixbufoverlay) SetOffsetY(value int) error {
	return e.Element.SetProperty("offset-y", value)
}

// overlay-height (int)
//
// Height of overlay image in pixels (0 = same as overlay image)

func (e *Gdkpixbufoverlay) GetOverlayHeight() (int, error) {
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

func (e *Gdkpixbufoverlay) SetOverlayHeight(value int) error {
	return e.Element.SetProperty("overlay-height", value)
}

// overlay-width (int)
//
// Width of overlay image in pixels (0 = same as overlay image)

func (e *Gdkpixbufoverlay) GetOverlayWidth() (int, error) {
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

func (e *Gdkpixbufoverlay) SetOverlayWidth(value int) error {
	return e.Element.SetProperty("overlay-width", value)
}

// pixbuf (GstGdkPixbuf)
//
// GdkPixbuf object to render.

func (e *Gdkpixbufoverlay) GetPixbuf() (interface{}, error) {
	return e.Element.GetProperty("pixbuf")
}

func (e *Gdkpixbufoverlay) SetPixbuf(value interface{}) error {
	return e.Element.SetProperty("pixbuf", value)
}

// positioning-mode (GstGdkPixbufPositioningMode)
//
// Positioning mode of offset-x and offset-y properties. Determines how
// negative x/y offsets will be interpreted. By default negative values
// are for positioning relative to the right/bottom edge of the video
// image, but you can use this property to select absolute positioning
// relative to a (0, 0) origin in the top-left corner. That way negative
// offsets will be to the left/above the video image, which allows you to
// smoothly slide logos into and out of the frame if desired.

func (e *Gdkpixbufoverlay) GetPositioningMode() (GstGdkPixbufPositioningMode, error) {
	var value GstGdkPixbufPositioningMode
	var ok bool
	v, err := e.Element.GetProperty("positioning-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGdkPixbufPositioningMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGdkPixbufPositioningMode")
	}
	return value, nil
}

func (e *Gdkpixbufoverlay) SetPositioningMode(value GstGdkPixbufPositioningMode) error {
	e.Element.SetArg("positioning-mode", string(value))
	return nil
}

// relative-x (float64)
//
// Horizontal offset of overlay image in fractions of video image width, from top-left corner of video image (in relative positioning)

func (e *Gdkpixbufoverlay) GetRelativeX() (float64, error) {
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

func (e *Gdkpixbufoverlay) SetRelativeX(value float64) error {
	return e.Element.SetProperty("relative-x", value)
}

// relative-y (float64)
//
// Vertical offset of overlay image in fractions of video image height, from top-left corner of video image (in relative positioning)

func (e *Gdkpixbufoverlay) GetRelativeY() (float64, error) {
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

func (e *Gdkpixbufoverlay) SetRelativeY(value float64) error {
	return e.Element.SetProperty("relative-y", value)
}

// ----- Constants -----

type GstGdkPixbufPositioningMode string

const (
	GstGdkPixbufPositioningModePixelsRelativeToEdges GstGdkPixbufPositioningMode = "pixels-relative-to-edges" // pixels-relative-to-edges (0) – pixels-relative-to-edges
	GstGdkPixbufPositioningModePixelsAbsolute GstGdkPixbufPositioningMode = "pixels-absolute" // pixels-absolute (1) – pixels-absolute
)

