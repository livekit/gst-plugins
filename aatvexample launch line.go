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

type AatvexampleLaunchLine struct {
	*base.GstBaseTransform
}

func NewAatvexampleLaunchLine() (*AatvexampleLaunchLine, error) {
	e, err := gst.NewElement("aatvExample launch line")
	if err != nil {
		return nil, err
	}
	return &AatvexampleLaunchLine{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAatvexampleLaunchLineWithName(name string) (*AatvexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("aatvExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &AatvexampleLaunchLine{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// brightness (int)
//
// Brightness

func (e *AatvexampleLaunchLine) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// brightness-actual (float32)
//
// Actual calculated foreground pixel fill percentage

func (e *AatvexampleLaunchLine) GetBrightnessActual() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness-actual")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// brightness-auto (bool)
//
// Automatically adjust brightness based on the previous frame's foreground pixel fill percentage

func (e *AatvexampleLaunchLine) GetBrightnessAuto() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("brightness-auto")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetBrightnessAuto(value bool) error {
	return e.Element.SetProperty("brightness-auto", value)
}

// brightness-max (float32)
//
// Maximum target foreground pixel fill percentage for automatic brightness control

func (e *AatvexampleLaunchLine) GetBrightnessMax() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetBrightnessMax(value float32) error {
	return e.Element.SetProperty("brightness-max", value)
}

// brightness-min (float32)
//
// Minimum target foreground pixel fill percentage for automatic brightness control

func (e *AatvexampleLaunchLine) GetBrightnessMin() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetBrightnessMin(value float32) error {
	return e.Element.SetProperty("brightness-min", value)
}

// color-background (uint)
//
// Color to use as the background for the ASCII text (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorBackground() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorBackground(value uint) error {
	return e.Element.SetProperty("color-background", value)
}

// color-rain (uint)
//
// Automatically sets color-rain-bold, color-rain-normal, and color-rain-dim with progressively dimmer values (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorRain() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorRain(value uint) error {
	return e.Element.SetProperty("color-rain", value)
}

// color-rain-bold (uint)
//
// Sets the brightest color to use for foreground ASCII text rain overlays (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorRainBold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain-bold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorRainBold(value uint) error {
	return e.Element.SetProperty("color-rain-bold", value)
}

// color-rain-dim (uint)
//
// Sets the dimmest brightness color to use for foreground ASCII text rain overlays (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorRainDim() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain-dim")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorRainDim(value uint) error {
	return e.Element.SetProperty("color-rain-dim", value)
}

// color-rain-normal (uint)
//
// Sets the normal brightness color to use for foreground ASCII text rain overlays (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorRainNormal() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain-normal")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorRainNormal(value uint) error {
	return e.Element.SetProperty("color-rain-normal", value)
}

// color-text (uint)
//
// Automatically sets color-test-bold, color-text-normal, and color-text-dim with progressively dimmer values (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorText() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorText(value uint) error {
	return e.Element.SetProperty("color-text", value)
}

// color-text-bold (uint)
//
// Sets the brightest color to use for foreground ASCII text (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorTextBold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text-bold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorTextBold(value uint) error {
	return e.Element.SetProperty("color-text-bold", value)
}

// color-text-dim (uint)
//
// Sets the dimmest brightness color to use for foreground ASCII text (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorTextDim() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text-dim")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorTextDim(value uint) error {
	return e.Element.SetProperty("color-text-dim", value)
}

// color-text-normal (uint)
//
// Sets the normal brightness color to use for foreground ASCII text (big-endian ARGB).

func (e *AatvexampleLaunchLine) GetColorTextNormal() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text-normal")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetColorTextNormal(value uint) error {
	return e.Element.SetProperty("color-text-normal", value)
}

// contrast (int)
//
// Contrast

func (e *AatvexampleLaunchLine) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// dither (GstAatvDitherers)
//
// Add noise to more closely approximate gray levels.

func (e *AatvexampleLaunchLine) GetDither() (GstAatvDitherers, error) {
	var value GstAatvDitherers
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAatvDitherers)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAatvDitherers")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetDither(value GstAatvDitherers) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

// font (GstAatvFonts)
//
// AAlib Font

func (e *AatvexampleLaunchLine) GetFont() (GstAatvFonts, error) {
	var value GstAatvFonts
	var ok bool
	v, err := e.Element.GetProperty("font")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAatvFonts)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAatvFonts")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetFont(value GstAatvFonts) error {
	e.Element.SetArg("font", string(value))
	return nil
}

// gamma (float32)
//
// Gamma correction

func (e *AatvexampleLaunchLine) GetGamma() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gamma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetGamma(value float32) error {
	return e.Element.SetProperty("gamma", value)
}

// height (int)
//
// Height of the ASCII canvas

func (e *AatvexampleLaunchLine) GetHeight() (int, error) {
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

func (e *AatvexampleLaunchLine) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// rain-delay-max (int)
//
// Maximum frame delay between rain motion

func (e *AatvexampleLaunchLine) GetRainDelayMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-delay-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRainDelayMax(value int) error {
	return e.Element.SetProperty("rain-delay-max", value)
}

// rain-delay-min (int)
//
// Minimum frame delay between rain motion

func (e *AatvexampleLaunchLine) GetRainDelayMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-delay-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRainDelayMin(value int) error {
	return e.Element.SetProperty("rain-delay-min", value)
}

// rain-length-max (int)
//
// Maximum length of a rain

func (e *AatvexampleLaunchLine) GetRainLengthMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-length-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRainLengthMax(value int) error {
	return e.Element.SetProperty("rain-length-max", value)
}

// rain-length-min (int)
//
// Minimum length of a rain

func (e *AatvexampleLaunchLine) GetRainLengthMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-length-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRainLengthMin(value int) error {
	return e.Element.SetProperty("rain-length-min", value)
}

// rain-mode (GstAatvRainModes)
//
// Set the direction of raindrops

func (e *AatvexampleLaunchLine) GetRainMode() (GstAatvRainModes, error) {
	var value GstAatvRainModes
	var ok bool
	v, err := e.Element.GetProperty("rain-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAatvRainModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAatvRainModes")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRainMode(value GstAatvRainModes) error {
	e.Element.SetArg("rain-mode", string(value))
	return nil
}

// rain-spawn-rate (float32)
//
// Percentage chance for a raindrop to spawn

func (e *AatvexampleLaunchLine) GetRainSpawnRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rain-spawn-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRainSpawnRate(value float32) error {
	return e.Element.SetProperty("rain-spawn-rate", value)
}

// randomval (int)
//
// Adds a random value in the range (-randomval/2,ranomval/2) to each pixel during rendering

func (e *AatvexampleLaunchLine) GetRandomval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("randomval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AatvexampleLaunchLine) SetRandomval(value int) error {
	return e.Element.SetProperty("randomval", value)
}

// width (int)
//
// Width of the ASCII canvas

func (e *AatvexampleLaunchLine) GetWidth() (int, error) {
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

func (e *AatvexampleLaunchLine) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

// ----- Constants -----

type GstAatvDitherers string

const (
	GstAatvDitherersNoDithering GstAatvDitherers = "no-dithering" // no-dithering (0) – no dithering
	GstAatvDitherersErrorDistribution GstAatvDitherers = "error-distribution" // error-distribution (1) – error-distribution
	GstAatvDitherersFloydSteelbergDithering GstAatvDitherers = "floyd-steelberg-dithering" // floyd-steelberg-dithering (2) – floyd-steelberg dithering
)

type GstAatvFonts string

const (
	GstAatvFontsStandardVga8X8Font GstAatvFonts = "Standard-vga-8x8-font" // Standard-vga-8x8-font (0) – vga8
	GstAatvFontsStandardVga8X9Font GstAatvFonts = "Standard-vga-8x9-font" // Standard-vga-8x9-font (1) – vga9
	GstAatvFontsStandardMdavga8X14Font GstAatvFonts = "Standard-mda/vga-8x14-font" // Standard-mda/vga-8x14-font (2) – mda14
	GstAatvFontsStandardVga8X16Font GstAatvFonts = "Standard-vga-8x16-font" // Standard-vga-8x16-font (3) – vga16
	GstAatvFontsX8X13Font GstAatvFonts = "X-8x13-font" // X-8x13-font (4) – X8x13
	GstAatvFontsX8X13BoldFont GstAatvFonts = "X-8x13bold-font" // X-8x13bold-font (5) – X8x13bold
	GstAatvFontsStandardX8X16Font GstAatvFonts = "Standard-X-8x16-font" // Standard-X-8x16-font (6) – X8x16
	GstAatvFontsLineFont8X8 GstAatvFonts = "line-Font-8x8" // line-Font-8x8 (7) – line
	GstAatvFontsFont8X8FromVgagl GstAatvFonts = "Font-8x8-from-vgagl" // Font-8x8-from-vgagl (8) – vgagl8
	GstAatvFontsAdobeCourier GstAatvFonts = "Adobe-courier" // Adobe-courier (9) – courier
)

type GstAatvRainModes string

const (
	GstAatvRainModesNone GstAatvRainModes = "none" // none (0) – No Rain
	GstAatvRainModesDown GstAatvRainModes = "down" // down (1) – Rain Down
	GstAatvRainModesUp GstAatvRainModes = "up" // up (2) – Rain Up
	GstAatvRainModesLeft GstAatvRainModes = "left" // left (3) – Rain Left
	GstAatvRainModesRight GstAatvRainModes = "right" // right (4) – Rain Right
)

