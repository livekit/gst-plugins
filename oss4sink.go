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

type Oss4Sink struct {
	*base.GstBaseSink
}

func NewOss4Sink() (*Oss4Sink, error) {
	e, err := gst.NewElement("oss4sink")
	if err != nil {
		return nil, err
	}
	return &Oss4Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewOss4SinkWithName(name string) (*Oss4Sink, error) {
	e, err := gst.NewElementWithName("oss4sink", name)
	if err != nil {
		return nil, err
	}
	return &Oss4Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// OSS4 device (e.g. /dev/oss/hdaudio0/pcm0 or /dev/dspN) (NULL = use first available playback device)

func (e *Oss4Sink) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Oss4Sink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Oss4Sink) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// mute (bool)
//
// Mute state of this stream

func (e *Oss4Sink) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Oss4Sink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// volume (float64)
//
// Linear volume of this stream, 1.0=100%%

func (e *Oss4Sink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Oss4Sink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

