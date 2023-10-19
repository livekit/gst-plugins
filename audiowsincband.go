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

type Audiowsincband struct {
	*base.GstBaseTransform
}

func NewAudiowsincband() (*Audiowsincband, error) {
	e, err := gst.NewElement("audiowsincband")
	if err != nil {
		return nil, err
	}
	return &Audiowsincband{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiowsincbandWithName(name string) (*Audiowsincband, error) {
	e, err := gst.NewElementWithName("audiowsincband", name)
	if err != nil {
		return nil, err
	}
	return &Audiowsincband{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// length (int)
//
// Filter kernel length, will be rounded to the next odd number

func (e *Audiowsincband) GetLength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audiowsincband) SetLength(value int) error {
	return e.Element.SetProperty("length", value)
}

// lower-frequency (float32)
//
// Cut-off lower frequency (Hz)

func (e *Audiowsincband) GetLowerFrequency() (float32, error) {
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

func (e *Audiowsincband) SetLowerFrequency(value float32) error {
	return e.Element.SetProperty("lower-frequency", value)
}

// mode (GstAudioWsincBandMode)
//
// Band pass or band reject mode

func (e *Audiowsincband) GetMode() (GstAudioWsincBandMode, error) {
	var value GstAudioWsincBandMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWsincBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWsincBandMode")
	}
	return value, nil
}

func (e *Audiowsincband) SetMode(value GstAudioWsincBandMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// upper-frequency (float32)
//
// Cut-off upper frequency (Hz)

func (e *Audiowsincband) GetUpperFrequency() (float32, error) {
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

func (e *Audiowsincband) SetUpperFrequency(value float32) error {
	return e.Element.SetProperty("upper-frequency", value)
}

// window (GstAudioWsincBandWindow)
//
// Window function to use

func (e *Audiowsincband) GetWindow() (GstAudioWsincBandWindow, error) {
	var value GstAudioWsincBandWindow
	var ok bool
	v, err := e.Element.GetProperty("window")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWsincBandWindow)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWsincBandWindow")
	}
	return value, nil
}

func (e *Audiowsincband) SetWindow(value GstAudioWsincBandWindow) error {
	e.Element.SetArg("window", string(value))
	return nil
}

// ----- Constants -----

type GstAudioWsincBandMode string

const (
	GstAudioWsincBandModeBandPass GstAudioWsincBandMode = "band-pass" // band-pass (0) – Band pass (default)
	GstAudioWsincBandModeBandReject GstAudioWsincBandMode = "band-reject" // band-reject (1) – Band reject
)

type GstAudioWsincBandWindow string

const (
	GstAudioWsincBandWindowHamming GstAudioWsincBandWindow = "hamming" // hamming (0) – Hamming window (default)
	GstAudioWsincBandWindowBlackman GstAudioWsincBandWindow = "blackman" // blackman (1) – Blackman window
	GstAudioWsincBandWindowGaussian GstAudioWsincBandWindow = "gaussian" // gaussian (2) – Gaussian window
	GstAudioWsincBandWindowCosine GstAudioWsincBandWindow = "cosine" // cosine (3) – Cosine window
	GstAudioWsincBandWindowHann GstAudioWsincBandWindow = "hann" // hann (4) – Hann window
)

