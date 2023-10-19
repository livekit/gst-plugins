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

type Speexenc struct {
	Element *gst.Element
}

func NewSpeexenc() (*Speexenc, error) {
	e, err := gst.NewElement("speexenc")
	if err != nil {
		return nil, err
	}
	return &Speexenc{Element: e}, nil
}

func NewSpeexencWithName(name string) (*Speexenc, error) {
	e, err := gst.NewElementWithName("speexenc", name)
	if err != nil {
		return nil, err
	}
	return &Speexenc{Element: e}, nil
}

// ----- Properties -----

// abr (int)
//
// Enable average bit-rate (0 = disabled)

func (e *Speexenc) GetAbr() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("abr")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Speexenc) SetAbr(value int) error {
	return e.Element.SetProperty("abr", value)
}

// bitrate (int)
//
// Specify an encoding bit-rate (in bps). (0 = automatic)

func (e *Speexenc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Speexenc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// complexity (int)
//
// Set encoding complexity

func (e *Speexenc) GetComplexity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("complexity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Speexenc) SetComplexity(value int) error {
	return e.Element.SetProperty("complexity", value)
}

// dtx (bool)
//
// Enable discontinuous transmission

func (e *Speexenc) GetDtx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dtx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Speexenc) SetDtx(value bool) error {
	return e.Element.SetProperty("dtx", value)
}

// last-message (string)
//
// The last status message

func (e *Speexenc) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// mode (GstSpeexEncMode)
//
// The encoding mode

func (e *Speexenc) GetMode() (GstSpeexEncMode, error) {
	var value GstSpeexEncMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSpeexEncMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSpeexEncMode")
	}
	return value, nil
}

func (e *Speexenc) SetMode(value GstSpeexEncMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// nframes (int)
//
// Number of frames per buffer

func (e *Speexenc) GetNframes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("nframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Speexenc) SetNframes(value int) error {
	return e.Element.SetProperty("nframes", value)
}

// quality (float32)
//
// Encoding quality

func (e *Speexenc) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Speexenc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

// vad (bool)
//
// Enable voice activity detection

func (e *Speexenc) GetVad() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vad")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Speexenc) SetVad(value bool) error {
	return e.Element.SetProperty("vad", value)
}

// vbr (bool)
//
// Enable variable bit-rate

func (e *Speexenc) GetVbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Speexenc) SetVbr(value bool) error {
	return e.Element.SetProperty("vbr", value)
}

// ----- Constants -----

type GstSpeexEncMode string

const (
	GstSpeexEncModeAuto GstSpeexEncMode = "auto" // auto (0) – Auto
	GstSpeexEncModeUwb GstSpeexEncMode = "uwb" // uwb (1) – Ultra Wide Band
	GstSpeexEncModeWb GstSpeexEncMode = "wb" // wb (2) – Wide Band
	GstSpeexEncModeNb GstSpeexEncMode = "nb" // nb (3) – Narrow Band
)

