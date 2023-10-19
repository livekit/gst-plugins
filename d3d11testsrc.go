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

type D3D11Testsrc struct {
	*base.GstBaseSrc
}

func NewD3D11Testsrc() (*D3D11Testsrc, error) {
	e, err := gst.NewElement("d3d11testsrc")
	if err != nil {
		return nil, err
	}
	return &D3D11Testsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewD3D11TestsrcWithName(name string) (*D3D11Testsrc, error) {
	e, err := gst.NewElementWithName("d3d11testsrc", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Testsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// adapter (int)
//
// DXGI Adapter index (-1 for any device)

func (e *D3D11Testsrc) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *D3D11Testsrc) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

// alpha (float32)
//
// Global alpha value to apply

func (e *D3D11Testsrc) GetAlpha() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *D3D11Testsrc) SetAlpha(value float32) error {
	return e.Element.SetProperty("alpha", value)
}

// alpha-mode (GstD3D11AlphaMode)
//
// Alpha mode to use

func (e *D3D11Testsrc) GetAlphaMode() (GstD3D11AlphaMode, error) {
	var value GstD3D11AlphaMode
	var ok bool
	v, err := e.Element.GetProperty("alpha-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11AlphaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11AlphaMode")
	}
	return value, nil
}

func (e *D3D11Testsrc) SetAlphaMode(value GstD3D11AlphaMode) error {
	e.Element.SetArg("alpha-mode", string(value))
	return nil
}

// is-live (bool)
//
// Whether to act as a live source

func (e *D3D11Testsrc) GetIsLive() (bool, error) {
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

func (e *D3D11Testsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// pattern (GstD3D11TestSrcPattern)
//
// Type of test pattern to generate

func (e *D3D11Testsrc) GetPattern() (GstD3D11TestSrcPattern, error) {
	var value GstD3D11TestSrcPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11TestSrcPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11TestSrcPattern")
	}
	return value, nil
}

func (e *D3D11Testsrc) SetPattern(value GstD3D11TestSrcPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

// ----- Constants -----

type GstD3D11AlphaMode string

const (
	GstD3D11AlphaModeUnspecified GstD3D11AlphaMode = "unspecified" // unspecified (0) – Unspecified
	GstD3D11AlphaModePremultiplied GstD3D11AlphaMode = "premultiplied" // premultiplied (1) – Premultiplied
	GstD3D11AlphaModeStraight GstD3D11AlphaMode = "straight" // straight (2) – Straight
)

type GstD3D11TestSrcPattern string

const (
	GstD3D11TestSrcPatternSmpte GstD3D11TestSrcPattern = "smpte" // smpte (0) – SMPTE 100%% color bars
	GstD3D11TestSrcPatternSnow GstD3D11TestSrcPattern = "snow" // snow (1) – Random (television snow)
	GstD3D11TestSrcPatternBlack GstD3D11TestSrcPattern = "black" // black (2) – 100%% Black
	GstD3D11TestSrcPatternWhite GstD3D11TestSrcPattern = "white" // white (3) – 100%% White
	GstD3D11TestSrcPatternRed GstD3D11TestSrcPattern = "red" // red (4) – Red
	GstD3D11TestSrcPatternGreen GstD3D11TestSrcPattern = "green" // green (5) – Green
	GstD3D11TestSrcPatternBlue GstD3D11TestSrcPattern = "blue" // blue (6) – Blue
	GstD3D11TestSrcPatternCheckers1 GstD3D11TestSrcPattern = "checkers-1" // checkers-1 (7) – Checkers 1px
	GstD3D11TestSrcPatternCheckers2 GstD3D11TestSrcPattern = "checkers-2" // checkers-2 (8) – Checkers 2px
	GstD3D11TestSrcPatternCheckers4 GstD3D11TestSrcPattern = "checkers-4" // checkers-4 (9) – Checkers 4px
	GstD3D11TestSrcPatternCheckers8 GstD3D11TestSrcPattern = "checkers-8" // checkers-8 (10) – Checkers 8px
	GstD3D11TestSrcPatternCircular GstD3D11TestSrcPattern = "circular" // circular (11) – Circular
	GstD3D11TestSrcPatternBlink GstD3D11TestSrcPattern = "blink" // blink (12) – Blink
	GstD3D11TestSrcPatternBall GstD3D11TestSrcPattern = "ball" // ball (18) – Moving ball
)

