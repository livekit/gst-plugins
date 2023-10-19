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

type Rawaudioparse struct {
	*base.GstBaseParse
}

func NewRawaudioparse() (*Rawaudioparse, error) {
	e, err := gst.NewElement("rawaudioparse")
	if err != nil {
		return nil, err
	}
	return &Rawaudioparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewRawaudioparseWithName(name string) (*Rawaudioparse, error) {
	e, err := gst.NewElementWithName("rawaudioparse", name)
	if err != nil {
		return nil, err
	}
	return &Rawaudioparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// channel-positions (GstGvalueArray)
//
// Channel positions used on the output

func (e *Rawaudioparse) GetChannelPositions() (interface{}, error) {
	return e.Element.GetProperty("channel-positions")
}

func (e *Rawaudioparse) SetChannelPositions(value interface{}) error {
	return e.Element.SetProperty("channel-positions", value)
}

// format (GstRawAudioParseFormat)
//
// Format of the raw audio stream

func (e *Rawaudioparse) GetFormat() (GstRawAudioParseFormat, error) {
	var value GstRawAudioParseFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRawAudioParseFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRawAudioParseFormat")
	}
	return value, nil
}

func (e *Rawaudioparse) SetFormat(value GstRawAudioParseFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

// interleaved (bool)
//
// True if audio has interleaved layout

func (e *Rawaudioparse) GetInterleaved() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("interleaved")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rawaudioparse) SetInterleaved(value bool) error {
	return e.Element.SetProperty("interleaved", value)
}

// num-channels (int)
//
// Number of channels in raw stream

func (e *Rawaudioparse) GetNumChannels() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rawaudioparse) SetNumChannels(value int) error {
	return e.Element.SetProperty("num-channels", value)
}

// pcm-format (GstAudioFormat)
//
// Format of audio samples in PCM stream (ignored if format property is not set to pcm)

func (e *Rawaudioparse) GetPcmFormat() (interface{}, error) {
	return e.Element.GetProperty("pcm-format")
}

func (e *Rawaudioparse) SetPcmFormat(value interface{}) error {
	return e.Element.SetProperty("pcm-format", value)
}

// sample-rate (int)
//
// Rate of audio samples in raw stream

func (e *Rawaudioparse) GetSampleRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sample-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rawaudioparse) SetSampleRate(value int) error {
	return e.Element.SetProperty("sample-rate", value)
}

// ----- Constants -----

type GstRawAudioParseFormat string

const (
	GstRawAudioParseFormatPcm GstRawAudioParseFormat = "pcm" // pcm (0) – PCM
	GstRawAudioParseFormatAlaw GstRawAudioParseFormat = "alaw" // alaw (2) – A-Law
	GstRawAudioParseFormatMulaw GstRawAudioParseFormat = "mulaw" // mulaw (1) – µ-Law
)

