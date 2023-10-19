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

type Audioparse struct {
	Element *gst.Element
}

func NewAudioparse() (*Audioparse, error) {
	e, err := gst.NewElement("audioparse")
	if err != nil {
		return nil, err
	}
	return &Audioparse{Element: e}, nil
}

func NewAudioparseWithName(name string) (*Audioparse, error) {
	e, err := gst.NewElementWithName("audioparse", name)
	if err != nil {
		return nil, err
	}
	return &Audioparse{Element: e}, nil
}

// ----- Properties -----

// channel-positions (GstGvalueArray)
//
// Channel positions used on the output

func (e *Audioparse) GetChannelPositions() (interface{}, error) {
	return e.Element.GetProperty("channel-positions")
}

func (e *Audioparse) SetChannelPositions(value interface{}) error {
	return e.Element.SetProperty("channel-positions", value)
}

// channels (int)
//
// Number of channels in raw stream

func (e *Audioparse) GetChannels() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audioparse) SetChannels(value int) error {
	return e.Element.SetProperty("channels", value)
}

// format (GstAudioParseFormat)
//
// Format of audio samples in raw stream

func (e *Audioparse) GetFormat() (GstAudioParseFormat, error) {
	var value GstAudioParseFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioParseFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioParseFormat")
	}
	return value, nil
}

func (e *Audioparse) SetFormat(value GstAudioParseFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

// interleaved (bool)
//
// True if audio has interleaved layout

func (e *Audioparse) GetInterleaved() (bool, error) {
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

func (e *Audioparse) SetInterleaved(value bool) error {
	return e.Element.SetProperty("interleaved", value)
}

// rate (int)
//
// Rate of audio samples in raw stream

func (e *Audioparse) GetRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audioparse) SetRate(value int) error {
	return e.Element.SetProperty("rate", value)
}

// raw-format (GstAudioFormat)
//
// Format of audio samples in raw stream

func (e *Audioparse) GetRawFormat() (interface{}, error) {
	return e.Element.GetProperty("raw-format")
}

func (e *Audioparse) SetRawFormat(value interface{}) error {
	return e.Element.SetProperty("raw-format", value)
}

// use-sink-caps (bool)
//
// Use the sink caps for the format, only performing timestamping

func (e *Audioparse) GetUseSinkCaps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-sink-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audioparse) SetUseSinkCaps(value bool) error {
	return e.Element.SetProperty("use-sink-caps", value)
}

// ----- Constants -----

type GstAudioParseFormat string

const (
	GstAudioParseFormatRaw GstAudioParseFormat = "raw" // raw (0) – Raw
	GstAudioParseFormatAlaw GstAudioParseFormat = "alaw" // alaw (2) – A-Law
	GstAudioParseFormatMulaw GstAudioParseFormat = "mulaw" // mulaw (1) – µ-Law
)

