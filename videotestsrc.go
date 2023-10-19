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

type Videotestsrc struct {
	*base.GstPushSrc
}

func NewVideotestsrc() (*Videotestsrc, error) {
	e, err := gst.NewElement("videotestsrc")
	if err != nil {
		return nil, err
	}
	return &Videotestsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewVideotestsrcWithName(name string) (*Videotestsrc, error) {
	e, err := gst.NewElementWithName("videotestsrc", name)
	if err != nil {
		return nil, err
	}
	return &Videotestsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// animation-mode (GstVideoTestSrcAnimationMode)
//
// For pattern=ball, which counter defines the position of the ball.

func (e *Videotestsrc) GetAnimationMode() (GstVideoTestSrcAnimationMode, error) {
	var value GstVideoTestSrcAnimationMode
	var ok bool
	v, err := e.Element.GetProperty("animation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoTestSrcAnimationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoTestSrcAnimationMode")
	}
	return value, nil
}

func (e *Videotestsrc) SetAnimationMode(value GstVideoTestSrcAnimationMode) error {
	e.Element.SetArg("animation-mode", string(value))
	return nil
}

// background-color (uint)
//
// Color to use for background color of some patterns.  Default is
// black (0xff000000).

func (e *Videotestsrc) GetBackgroundColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("background-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Videotestsrc) SetBackgroundColor(value uint) error {
	return e.Element.SetProperty("background-color", value)
}

// flip (bool)
//
// For pattern=ball, invert colors every second.

func (e *Videotestsrc) GetFlip() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flip")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videotestsrc) SetFlip(value bool) error {
	return e.Element.SetProperty("flip", value)
}

// foreground-color (uint)
//
// Color to use for solid-color pattern and foreground color of other
// patterns.  Default is white (0xffffffff).

func (e *Videotestsrc) GetForegroundColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("foreground-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Videotestsrc) SetForegroundColor(value uint) error {
	return e.Element.SetProperty("foreground-color", value)
}

// horizontal-speed (int)
//
// Scroll image number of pixels per frame (positive is scroll to the left)

func (e *Videotestsrc) GetHorizontalSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("horizontal-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetHorizontalSpeed(value int) error {
	return e.Element.SetProperty("horizontal-speed", value)
}

// is-live (bool)
//
// Whether to act as a live source

func (e *Videotestsrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videotestsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// k0 (int)
//
// Zoneplate zero order phase, for generating plain fields or phase offsets

func (e *Videotestsrc) GetK0() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("k0")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetK0(value int) error {
	return e.Element.SetProperty("k0", value)
}

// kt (int)
//
// Zoneplate 1st order t phase, for generating phase rotation as a function of time

func (e *Videotestsrc) GetKt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKt(value int) error {
	return e.Element.SetProperty("kt", value)
}

// kt2 (int)
//
// Zoneplate 2nd order t phase, t*t/256 cycles per picture

func (e *Videotestsrc) GetKt2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kt2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKt2(value int) error {
	return e.Element.SetProperty("kt2", value)
}

// kx (int)
//
// Zoneplate 1st order x phase, for generating constant horizontal frequencies

func (e *Videotestsrc) GetKx() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kx")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKx(value int) error {
	return e.Element.SetProperty("kx", value)
}

// kx2 (int)
//
// Zoneplate 2nd order x phase, normalised to kx2/256 cycles per horizontal pixel at width/2 from origin

func (e *Videotestsrc) GetKx2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kx2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKx2(value int) error {
	return e.Element.SetProperty("kx2", value)
}

// kxt (int)
//
// Zoneplate x*t product phase, normalised to kxy/256 cycles per vertical pixel at width/2 from origin

func (e *Videotestsrc) GetKxt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kxt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKxt(value int) error {
	return e.Element.SetProperty("kxt", value)
}

// kxy (int)
//
// Zoneplate x*y product phase

func (e *Videotestsrc) GetKxy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKxy(value int) error {
	return e.Element.SetProperty("kxy", value)
}

// ky (int)
//
// Zoneplate 1st order y phase, for generating content vertical frequencies

func (e *Videotestsrc) GetKy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ky")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKy(value int) error {
	return e.Element.SetProperty("ky", value)
}

// ky2 (int)
//
// Zoneplate 2nd order y phase, normailsed to ky2/256 cycles per vertical pixel at height/2 from origin

func (e *Videotestsrc) GetKy2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ky2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKy2(value int) error {
	return e.Element.SetProperty("ky2", value)
}

// kyt (int)
//
// Zoneplate y*t product phase

func (e *Videotestsrc) GetKyt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kyt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetKyt(value int) error {
	return e.Element.SetProperty("kyt", value)
}

// motion (GstVideoTestSrcMotionType)
//
// For pattern=ball, what motion the ball does

func (e *Videotestsrc) GetMotion() (GstVideoTestSrcMotionType, error) {
	var value GstVideoTestSrcMotionType
	var ok bool
	v, err := e.Element.GetProperty("motion")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoTestSrcMotionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoTestSrcMotionType")
	}
	return value, nil
}

func (e *Videotestsrc) SetMotion(value GstVideoTestSrcMotionType) error {
	e.Element.SetArg("motion", string(value))
	return nil
}

// pattern (GstVideoTestSrcPattern)
//
// Type of test pattern to generate

func (e *Videotestsrc) GetPattern() (GstVideoTestSrcPattern, error) {
	var value GstVideoTestSrcPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoTestSrcPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoTestSrcPattern")
	}
	return value, nil
}

func (e *Videotestsrc) SetPattern(value GstVideoTestSrcPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

// timestamp-offset (int64)
//
// An offset added to timestamps set on buffers (in ns)

func (e *Videotestsrc) GetTimestampOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Videotestsrc) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

// xoffset (int)
//
// Zoneplate 2nd order products x offset

func (e *Videotestsrc) GetXoffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetXoffset(value int) error {
	return e.Element.SetProperty("xoffset", value)
}

// yoffset (int)
//
// Zoneplate 2nd order products y offset

func (e *Videotestsrc) GetYoffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("yoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videotestsrc) SetYoffset(value int) error {
	return e.Element.SetProperty("yoffset", value)
}

// ----- Constants -----

type GstVideoTestSrcPattern string

const (
	GstVideoTestSrcPatternSmpte GstVideoTestSrcPattern = "smpte" // smpte (0) – SMPTE 100%% color bars
	GstVideoTestSrcPatternSnow GstVideoTestSrcPattern = "snow" // snow (1) – Random (television snow)
	GstVideoTestSrcPatternBlack GstVideoTestSrcPattern = "black" // black (2) – 100%% Black
	GstVideoTestSrcPatternWhite GstVideoTestSrcPattern = "white" // white (3) – 100%% White
	GstVideoTestSrcPatternRed GstVideoTestSrcPattern = "red" // red (4) – Red
	GstVideoTestSrcPatternGreen GstVideoTestSrcPattern = "green" // green (5) – Green
	GstVideoTestSrcPatternBlue GstVideoTestSrcPattern = "blue" // blue (6) – Blue
	GstVideoTestSrcPatternCheckers1 GstVideoTestSrcPattern = "checkers-1" // checkers-1 (7) – Checkers 1px
	GstVideoTestSrcPatternCheckers2 GstVideoTestSrcPattern = "checkers-2" // checkers-2 (8) – Checkers 2px
	GstVideoTestSrcPatternCheckers4 GstVideoTestSrcPattern = "checkers-4" // checkers-4 (9) – Checkers 4px
	GstVideoTestSrcPatternCheckers8 GstVideoTestSrcPattern = "checkers-8" // checkers-8 (10) – Checkers 8px
	GstVideoTestSrcPatternCircular GstVideoTestSrcPattern = "circular" // circular (11) – Circular
	GstVideoTestSrcPatternBlink GstVideoTestSrcPattern = "blink" // blink (12) – Blink
	GstVideoTestSrcPatternSmpte75 GstVideoTestSrcPattern = "smpte75" // smpte75 (13) – SMPTE 75%% color bars
	GstVideoTestSrcPatternZonePlate GstVideoTestSrcPattern = "zone-plate" // zone-plate (14) – Zone plate
	GstVideoTestSrcPatternGamut GstVideoTestSrcPattern = "gamut" // gamut (15) – Gamut checkers
	GstVideoTestSrcPatternChromaZonePlate GstVideoTestSrcPattern = "chroma-zone-plate" // chroma-zone-plate (16) – Chroma zone plate
	GstVideoTestSrcPatternSolidColor GstVideoTestSrcPattern = "solid-color" // solid-color (17) – Solid color
	GstVideoTestSrcPatternBall GstVideoTestSrcPattern = "ball" // ball (18) – Moving ball
	GstVideoTestSrcPatternSmpte100 GstVideoTestSrcPattern = "smpte100" // smpte100 (19) – SMPTE 100%% color bars
	GstVideoTestSrcPatternBar GstVideoTestSrcPattern = "bar" // bar (20) – Bar
	GstVideoTestSrcPatternPinwheel GstVideoTestSrcPattern = "pinwheel" // pinwheel (21) – Pinwheel
	GstVideoTestSrcPatternSpokes GstVideoTestSrcPattern = "spokes" // spokes (22) – Spokes
	GstVideoTestSrcPatternGradient GstVideoTestSrcPattern = "gradient" // gradient (23) – Gradient
	GstVideoTestSrcPatternColors GstVideoTestSrcPattern = "colors" // colors (24) – Colors
	GstVideoTestSrcPatternSmpteRp219 GstVideoTestSrcPattern = "smpte-rp-219" // smpte-rp-219 (25) – SMPTE test pattern, RP 219 conformant
)

type GstVideoTestSrcAnimationMode string

const (
	GstVideoTestSrcAnimationModeFrames GstVideoTestSrcAnimationMode = "frames" // frames (0) – frame count
	GstVideoTestSrcAnimationModeWallTime GstVideoTestSrcAnimationMode = "wall-time" // wall-time (1) – wall clock time
	GstVideoTestSrcAnimationModeRunningTime GstVideoTestSrcAnimationMode = "running-time" // running-time (2) – running time
)

type GstVideoTestSrcMotionType string

const (
	GstVideoTestSrcMotionTypeWavy GstVideoTestSrcMotionType = "wavy" // wavy (0) – Ball waves back and forth, up and down
	GstVideoTestSrcMotionTypeSweep GstVideoTestSrcMotionType = "sweep" // sweep (1) – 1 revolution per second
	GstVideoTestSrcMotionTypeHsweep GstVideoTestSrcMotionType = "hsweep" // hsweep (2) – 1/2 revolution per second, then reset to top
)

