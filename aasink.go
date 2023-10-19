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

type Aasink struct {
	*base.GstBaseSink
}

func NewAasink() (*Aasink, error) {
	e, err := gst.NewElement("aasink")
	if err != nil {
		return nil, err
	}
	return &Aasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAasinkWithName(name string) (*Aasink, error) {
	e, err := gst.NewElementWithName("aasink", name)
	if err != nil {
		return nil, err
	}
	return &Aasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// brightness (int)
//
// brightness

func (e *Aasink) GetBrightness() (int, error) {
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

func (e *Aasink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (int)
//
// contrast

func (e *Aasink) GetContrast() (int, error) {
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

func (e *Aasink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// dither (GstAasinkDitherers)
//
// dither

func (e *Aasink) GetDither() (GstAasinkDitherers, error) {
	var value GstAasinkDitherers
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAasinkDitherers)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAasinkDitherers")
	}
	return value, nil
}

func (e *Aasink) SetDither(value GstAasinkDitherers) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

// driver (GstAasinkDrivers)
//
// driver

func (e *Aasink) GetDriver() (GstAasinkDrivers, error) {
	var value GstAasinkDrivers
	var ok bool
	v, err := e.Element.GetProperty("driver")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAasinkDrivers)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAasinkDrivers")
	}
	return value, nil
}

func (e *Aasink) SetDriver(value GstAasinkDrivers) error {
	e.Element.SetArg("driver", string(value))
	return nil
}

// frame-time (int)
//
// frame time

func (e *Aasink) GetFrameTime() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frame-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// frames-displayed (int)
//
// frames displayed

func (e *Aasink) GetFramesDisplayed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frames-displayed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// gamma (float32)
//
// gamma

func (e *Aasink) GetGamma() (float32, error) {
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

func (e *Aasink) SetGamma(value float32) error {
	return e.Element.SetProperty("gamma", value)
}

// height (int)
//
// height

func (e *Aasink) GetHeight() (int, error) {
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

func (e *Aasink) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// inversion (bool)
//
// inversion

func (e *Aasink) GetInversion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("inversion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Aasink) SetInversion(value bool) error {
	return e.Element.SetProperty("inversion", value)
}

// randomval (int)
//
// randomval

func (e *Aasink) GetRandomval() (int, error) {
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

func (e *Aasink) SetRandomval(value int) error {
	return e.Element.SetProperty("randomval", value)
}

// width (int)
//
// width

func (e *Aasink) GetWidth() (int, error) {
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

func (e *Aasink) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

// ----- Constants -----

type GstAasinkDitherers string

const (
	GstAasinkDitherersNoDithering GstAasinkDitherers = "no-dithering" // no-dithering (0) – no dithering
	GstAasinkDitherersErrorDistribution GstAasinkDitherers = "error-distribution" // error-distribution (1) – error-distribution
	GstAasinkDitherersFloydSteelbergDithering GstAasinkDitherers = "floyd-steelberg-dithering" // floyd-steelberg-dithering (2) – floyd-steelberg dithering
)

type GstAasinkDrivers string

const (
	GstAasinkDriversX11 GstAasinkDrivers = "x11" // x11 (0) – X11 driver 1.1
	GstAasinkDriversLinux GstAasinkDrivers = "linux" // linux (1) – Linux pc console driver 1.0
	GstAasinkDriversSlang GstAasinkDrivers = "slang" // slang (2) – Slang driver 1.0
	GstAasinkDriversCurses GstAasinkDrivers = "curses" // curses (3) – Curses driver 1.0
	GstAasinkDriversStdout GstAasinkDrivers = "stdout" // stdout (4) – Standard output driver
	GstAasinkDriversStderr GstAasinkDrivers = "stderr" // stderr (5) – Standard error driver
)

