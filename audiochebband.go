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

type Audiochebband struct {
	*base.GstBaseTransform
}

func NewAudiochebband() (*Audiochebband, error) {
	e, err := gst.NewElement("audiochebband")
	if err != nil {
		return nil, err
	}
	return &Audiochebband{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiochebbandWithName(name string) (*Audiochebband, error) {
	e, err := gst.NewElementWithName("audiochebband", name)
	if err != nil {
		return nil, err
	}
	return &Audiochebband{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// lower-frequency (float32)
//
// Start frequency of the band (Hz)

func (e *Audiochebband) GetLowerFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lower-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiochebband) SetLowerFrequency(value float32) error {
	return e.Element.SetProperty("lower-frequency", value)
}

// mode (GstAudioChebBandMode)
//
// Low pass or high pass mode

func (e *Audiochebband) GetMode() (GstAudioChebBandMode, error) {
	var value GstAudioChebBandMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioChebBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioChebBandMode")
	}
	return value, nil
}

func (e *Audiochebband) SetMode(value GstAudioChebBandMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// poles (int)
//
// Number of poles to use, will be rounded up to the next multiply of four

func (e *Audiochebband) GetPoles() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("poles")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audiochebband) SetPoles(value int) error {
	return e.Element.SetProperty("poles", value)
}

// ripple (float32)
//
// Amount of ripple (dB)

func (e *Audiochebband) GetRipple() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ripple")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiochebband) SetRipple(value float32) error {
	return e.Element.SetProperty("ripple", value)
}

// type (int)
//
// Type of the chebychev filter

func (e *Audiochebband) GetType() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audiochebband) SetType(value int) error {
	return e.Element.SetProperty("type", value)
}

// upper-frequency (float32)
//
// Stop frequency of the band (Hz)

func (e *Audiochebband) GetUpperFrequency() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("upper-frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiochebband) SetUpperFrequency(value float32) error {
	return e.Element.SetProperty("upper-frequency", value)
}

// ----- Constants -----

type GstAudioChebBandMode string

const (
	GstAudioChebBandModeBandPass GstAudioChebBandMode = "band-pass" // band-pass (0) – Band pass (default)
	GstAudioChebBandModeBandReject GstAudioChebBandMode = "band-reject" // band-reject (1) – Band reject
)

