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

type Audiocheblimit struct {
	*base.GstBaseTransform
}

func NewAudiocheblimit() (*Audiocheblimit, error) {
	e, err := gst.NewElement("audiocheblimit")
	if err != nil {
		return nil, err
	}
	return &Audiocheblimit{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiocheblimitWithName(name string) (*Audiocheblimit, error) {
	e, err := gst.NewElementWithName("audiocheblimit", name)
	if err != nil {
		return nil, err
	}
	return &Audiocheblimit{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cutoff (float32)
//
// Cut off frequency (Hz)

func (e *Audiocheblimit) GetCutoff() (float32, error) {
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

func (e *Audiocheblimit) SetCutoff(value float32) error {
	return e.Element.SetProperty("cutoff", value)
}

// mode (GstAudioChebLimitMode)
//
// Low pass or high pass mode

func (e *Audiocheblimit) GetMode() (GstAudioChebLimitMode, error) {
	var value GstAudioChebLimitMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioChebLimitMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioChebLimitMode")
	}
	return value, nil
}

func (e *Audiocheblimit) SetMode(value GstAudioChebLimitMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// poles (int)
//
// Number of poles to use, will be rounded up to the next even number

func (e *Audiocheblimit) GetPoles() (int, error) {
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

func (e *Audiocheblimit) SetPoles(value int) error {
	return e.Element.SetProperty("poles", value)
}

// ripple (float32)
//
// Amount of ripple (dB)

func (e *Audiocheblimit) GetRipple() (float32, error) {
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

func (e *Audiocheblimit) SetRipple(value float32) error {
	return e.Element.SetProperty("ripple", value)
}

// type (int)
//
// Type of the chebychev filter

func (e *Audiocheblimit) GetType() (int, error) {
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

func (e *Audiocheblimit) SetType(value int) error {
	return e.Element.SetProperty("type", value)
}

// ----- Constants -----

type GstAudioChebLimitMode string

const (
	GstAudioChebLimitModeLowPass GstAudioChebLimitMode = "low-pass" // low-pass (0) – Low pass (default)
	GstAudioChebLimitModeHighPass GstAudioChebLimitMode = "high-pass" // high-pass (1) – High pass
)

