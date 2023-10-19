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

type CacatvexampleLaunchLine struct {
	*base.GstBaseTransform
}

func NewCacatvexampleLaunchLine() (*CacatvexampleLaunchLine, error) {
	e, err := gst.NewElement("cacatvExample launch line")
	if err != nil {
		return nil, err
	}
	return &CacatvexampleLaunchLine{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCacatvexampleLaunchLineWithName(name string) (*CacatvexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("cacatvExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &CacatvexampleLaunchLine{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// anti-aliasing (bool)
//
// Enables Anti-Aliasing

func (e *CacatvexampleLaunchLine) GetAntiAliasing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("anti-aliasing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *CacatvexampleLaunchLine) SetAntiAliasing(value bool) error {
	return e.Element.SetProperty("anti-aliasing", value)
}

// canvas-height (int)
//
// The height of the canvas in characters

func (e *CacatvexampleLaunchLine) GetCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *CacatvexampleLaunchLine) SetCanvasHeight(value int) error {
	return e.Element.SetProperty("canvas-height", value)
}

// canvas-width (int)
//
// The width of the canvas in characters

func (e *CacatvexampleLaunchLine) GetCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *CacatvexampleLaunchLine) SetCanvasWidth(value int) error {
	return e.Element.SetProperty("canvas-width", value)
}

// dither (GstCacatvDithering)
//
// Set type of Dither

func (e *CacatvexampleLaunchLine) GetDither() (GstCacatvDithering, error) {
	var value GstCacatvDithering
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCacatvDithering)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCacatvDithering")
	}
	return value, nil
}

func (e *CacatvexampleLaunchLine) SetDither(value GstCacatvDithering) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

// font (int)
//
// selected libcaca font

func (e *CacatvexampleLaunchLine) GetFont() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("font")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *CacatvexampleLaunchLine) SetFont(value int) error {
	return e.Element.SetProperty("font", value)
}

// ----- Constants -----

type GstCacatvDithering string

const (
	GstCacatvDitheringNone GstCacatvDithering = "none" // none (49) – No dither_mode
	GstCacatvDithering2X2 GstCacatvDithering = "2x2" // 2x2 (50) – Ordered 2x2 Bayer dither_mode
	GstCacatvDithering4X4 GstCacatvDithering = "4x4" // 4x4 (51) – Ordered 4x4 Bayer dither_mode
	GstCacatvDithering8X8 GstCacatvDithering = "8x8" // 8x8 (52) – Ordered 8x8 Bayer dither_mode
	GstCacatvDitheringRandom GstCacatvDithering = "random" // random (53) – Random dither_mode
)

