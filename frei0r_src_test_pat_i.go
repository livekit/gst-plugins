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

type Frei0RSrcTestPatI struct {
	*base.GstPushSrc
}

func NewFrei0RSrcTestPatI() (*Frei0RSrcTestPatI, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-i")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatI{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcTestPatIWithName(name string) (*Frei0RSrcTestPatI, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-i", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatI{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// amplitude (float64)
//
// Amplitude (contrast) of the pattern

func (e *Frei0RSrcTestPatI) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatI) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

// channel (float64)
//
// Into which color channel to draw

func (e *Frei0RSrcTestPatI) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatI) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

// negative (bool)
//
// Change polarity of impulse/step

func (e *Frei0RSrcTestPatI) GetNegative() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("negative")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatI) SetNegative(value bool) error {
	return e.Element.SetProperty("negative", value)
}

// tilt (float64)
//
// Angle of step function

func (e *Frei0RSrcTestPatI) GetTilt() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatI) SetTilt(value float64) error {
	return e.Element.SetProperty("tilt", value)
}

// type (float64)
//
// Type of test pattern

func (e *Frei0RSrcTestPatI) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatI) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

// width (float64)
//
// Width of impulse

func (e *Frei0RSrcTestPatI) GetWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatI) SetWidth(value float64) error {
	return e.Element.SetProperty("width", value)
}

