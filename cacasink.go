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

type Cacasink struct {
	*base.GstBaseSink
}

func NewCacasink() (*Cacasink, error) {
	e, err := gst.NewElement("cacasink")
	if err != nil {
		return nil, err
	}
	return &Cacasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewCacasinkWithName(name string) (*Cacasink, error) {
	e, err := gst.NewElementWithName("cacasink", name)
	if err != nil {
		return nil, err
	}
	return &Cacasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// anti-aliasing (bool)
//
// Enables Anti-Aliasing

func (e *Cacasink) GetAntiAliasing() (bool, error) {
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

func (e *Cacasink) SetAntiAliasing(value bool) error {
	return e.Element.SetProperty("anti-aliasing", value)
}

// dither (GstCacasinkDithering)
//
// Set type of Dither

func (e *Cacasink) GetDither() (GstCacasinkDithering, error) {
	var value GstCacasinkDithering
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCacasinkDithering)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCacasinkDithering")
	}
	return value, nil
}

func (e *Cacasink) SetDither(value GstCacasinkDithering) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

// screen-height (int)
//
// The height of the screen

func (e *Cacasink) GetScreenHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("screen-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// screen-width (int)
//
// The width of the screen

func (e *Cacasink) GetScreenWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("screen-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ----- Constants -----

type GstCacasinkDithering string

const (
	GstCacasinkDitheringNone GstCacasinkDithering = "none" // none (49) – No dithering
	GstCacasinkDithering2X2 GstCacasinkDithering = "2x2" // 2x2 (50) – Ordered 2x2 Bayer dithering
	GstCacasinkDithering4X4 GstCacasinkDithering = "4x4" // 4x4 (51) – Ordered 4x4 Bayer dithering
	GstCacasinkDithering8X8 GstCacasinkDithering = "8x8" // 8x8 (52) – Ordered 8x8 Bayer dithering
	GstCacasinkDitheringRandom GstCacasinkDithering = "random" // random (53) – Random dithering
)

