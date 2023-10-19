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

type Gltestsrc struct {
	*base.GstPushSrc
}

func NewGltestsrc() (*Gltestsrc, error) {
	e, err := gst.NewElement("gltestsrc")
	if err != nil {
		return nil, err
	}
	return &Gltestsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewGltestsrcWithName(name string) (*Gltestsrc, error) {
	e, err := gst.NewElementWithName("gltestsrc", name)
	if err != nil {
		return nil, err
	}
	return &Gltestsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// is-live (bool)
//
// Whether to act as a live source

func (e *Gltestsrc) GetIsLive() (bool, error) {
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

func (e *Gltestsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// pattern (GstGltestSrcPattern)
//
// Type of test pattern to generate

func (e *Gltestsrc) GetPattern() (GstGltestSrcPattern, error) {
	var value GstGltestSrcPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGltestSrcPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGltestSrcPattern")
	}
	return value, nil
}

func (e *Gltestsrc) SetPattern(value GstGltestSrcPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

// ----- Constants -----

type GstGltestSrcPattern string

const (
	GstGltestSrcPatternSmpte GstGltestSrcPattern = "smpte" // smpte (0) – SMPTE 100%% color bars
	GstGltestSrcPatternSnow GstGltestSrcPattern = "snow" // snow (1) – Random (television snow)
	GstGltestSrcPatternBlack GstGltestSrcPattern = "black" // black (2) – 100%% Black
	GstGltestSrcPatternWhite GstGltestSrcPattern = "white" // white (3) – 100%% White
	GstGltestSrcPatternRed GstGltestSrcPattern = "red" // red (4) – Red
	GstGltestSrcPatternGreen GstGltestSrcPattern = "green" // green (5) – Green
	GstGltestSrcPatternBlue GstGltestSrcPattern = "blue" // blue (6) – Blue
	GstGltestSrcPatternCheckers1 GstGltestSrcPattern = "checkers-1" // checkers-1 (7) – Checkers 1px
	GstGltestSrcPatternCheckers2 GstGltestSrcPattern = "checkers-2" // checkers-2 (8) – Checkers 2px
	GstGltestSrcPatternCheckers4 GstGltestSrcPattern = "checkers-4" // checkers-4 (9) – Checkers 4px
	GstGltestSrcPatternCheckers8 GstGltestSrcPattern = "checkers-8" // checkers-8 (10) – Checkers 8px
	GstGltestSrcPatternCircular GstGltestSrcPattern = "circular" // circular (11) – Circular
	GstGltestSrcPatternBlink GstGltestSrcPattern = "blink" // blink (12) – Blink
	GstGltestSrcPatternMandelbrot GstGltestSrcPattern = "mandelbrot" // mandelbrot (13) – Mandelbrot Fractal
)

