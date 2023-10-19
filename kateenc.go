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
)

type Kateenc struct {
	Element *gst.Element
}

func NewKateenc() (*Kateenc, error) {
	e, err := gst.NewElement("kateenc")
	if err != nil {
		return nil, err
	}
	return &Kateenc{Element: e}, nil
}

func NewKateencWithName(name string) (*Kateenc, error) {
	e, err := gst.NewElementWithName("kateenc", name)
	if err != nil {
		return nil, err
	}
	return &Kateenc{Element: e}, nil
}

// ----- Properties -----

// category (string)
//
// The category of the stream

func (e *Kateenc) GetCategory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("category")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Kateenc) SetCategory(value string) error {
	return e.Element.SetProperty("category", value)
}

// default-spu-duration (float32)
//
// The assumed max duration (in seconds) of SPUs with no duration specified

func (e *Kateenc) GetDefaultSpuDuration() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("default-spu-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Kateenc) SetDefaultSpuDuration(value float32) error {
	return e.Element.SetProperty("default-spu-duration", value)
}

// granule-rate-denominator (int)
//
// The denominator of the granule rate

func (e *Kateenc) GetGranuleRateDenominator() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("granule-rate-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kateenc) SetGranuleRateDenominator(value int) error {
	return e.Element.SetProperty("granule-rate-denominator", value)
}

// granule-rate-numerator (int)
//
// The numerator of the granule rate

func (e *Kateenc) GetGranuleRateNumerator() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("granule-rate-numerator")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kateenc) SetGranuleRateNumerator(value int) error {
	return e.Element.SetProperty("granule-rate-numerator", value)
}

// granule-shift (int)
//
// The granule shift

func (e *Kateenc) GetGranuleShift() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("granule-shift")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kateenc) SetGranuleShift(value int) error {
	return e.Element.SetProperty("granule-shift", value)
}

// keepalive-min-time (float32)
//
// Minimum time to emit keepalive packets (0 disables keepalive packets)

func (e *Kateenc) GetKeepaliveMinTime() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("keepalive-min-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Kateenc) SetKeepaliveMinTime(value float32) error {
	return e.Element.SetProperty("keepalive-min-time", value)
}

// language (string)
//
// The language of the stream (e.g. "fr" or "fr_FR" for French)

func (e *Kateenc) GetLanguage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("language")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Kateenc) SetLanguage(value string) error {
	return e.Element.SetProperty("language", value)
}

// original-canvas-height (int)
//
// The height of the canvas this stream was authored for (0 is unspecified)

func (e *Kateenc) GetOriginalCanvasHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kateenc) SetOriginalCanvasHeight(value int) error {
	return e.Element.SetProperty("original-canvas-height", value)
}

// original-canvas-width (int)
//
// The width of the canvas this stream was authored for (0 is unspecified)

func (e *Kateenc) GetOriginalCanvasWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original-canvas-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Kateenc) SetOriginalCanvasWidth(value int) error {
	return e.Element.SetProperty("original-canvas-width", value)
}

