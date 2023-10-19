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

type Audiowsinclimit struct {
	*base.GstBaseTransform
}

func NewAudiowsinclimit() (*Audiowsinclimit, error) {
	e, err := gst.NewElement("audiowsinclimit")
	if err != nil {
		return nil, err
	}
	return &Audiowsinclimit{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiowsinclimitWithName(name string) (*Audiowsinclimit, error) {
	e, err := gst.NewElementWithName("audiowsinclimit", name)
	if err != nil {
		return nil, err
	}
	return &Audiowsinclimit{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cutoff (float32)
//
// Cut-off Frequency (Hz)

func (e *Audiowsinclimit) GetCutoff() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cutoff")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audiowsinclimit) SetCutoff(value float32) error {
	return e.Element.SetProperty("cutoff", value)
}

// length (int)
//
// Filter kernel length, will be rounded to the next odd number

func (e *Audiowsinclimit) GetLength() (int, error) {
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

func (e *Audiowsinclimit) SetLength(value int) error {
	return e.Element.SetProperty("length", value)
}

// mode (GstAudioWsincLimitMode)
//
// Low pass or high pass mode

func (e *Audiowsinclimit) GetMode() (GstAudioWsincLimitMode, error) {
	var value GstAudioWsincLimitMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWsincLimitMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWsincLimitMode")
	}
	return value, nil
}

func (e *Audiowsinclimit) SetMode(value GstAudioWsincLimitMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// window (GstAudioWsincLimitWindow)
//
// Window function to use

func (e *Audiowsinclimit) GetWindow() (GstAudioWsincLimitWindow, error) {
	var value GstAudioWsincLimitWindow
	var ok bool
	v, err := e.Element.GetProperty("window")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioWsincLimitWindow)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioWsincLimitWindow")
	}
	return value, nil
}

func (e *Audiowsinclimit) SetWindow(value GstAudioWsincLimitWindow) error {
	e.Element.SetArg("window", string(value))
	return nil
}

// ----- Constants -----

type GstAudioWsincLimitMode string

const (
	GstAudioWsincLimitModeLowPass GstAudioWsincLimitMode = "low-pass" // low-pass (0) – Low pass (default)
	GstAudioWsincLimitModeHighPass GstAudioWsincLimitMode = "high-pass" // high-pass (1) – High pass
)

type GstAudioWsincLimitWindow string

const (
	GstAudioWsincLimitWindowHamming GstAudioWsincLimitWindow = "hamming" // hamming (0) – Hamming window (default)
	GstAudioWsincLimitWindowBlackman GstAudioWsincLimitWindow = "blackman" // blackman (1) – Blackman window
	GstAudioWsincLimitWindowGaussian GstAudioWsincLimitWindow = "gaussian" // gaussian (2) – Gaussian window
	GstAudioWsincLimitWindowCosine GstAudioWsincLimitWindow = "cosine" // cosine (3) – Cosine window
	GstAudioWsincLimitWindowHann GstAudioWsincLimitWindow = "hann" // hann (4) – Hann window
)

