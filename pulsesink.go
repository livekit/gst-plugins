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

type Pulsesink struct {
	*base.GstBaseSink
}

func NewPulsesink() (*Pulsesink, error) {
	e, err := gst.NewElement("pulsesink")
	if err != nil {
		return nil, err
	}
	return &Pulsesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewPulsesinkWithName(name string) (*Pulsesink, error) {
	e, err := gst.NewElementWithName("pulsesink", name)
	if err != nil {
		return nil, err
	}
	return &Pulsesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// client-name (string)
//
// The PulseAudio client name to use.

func (e *Pulsesink) GetClientName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("client-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Pulsesink) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

// current-device (string)
//
// The current PulseAudio sink device

func (e *Pulsesink) GetCurrentDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// device (string)
//
// The PulseAudio sink device to connect to

func (e *Pulsesink) GetDevice() (string, error) {
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

func (e *Pulsesink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Pulsesink) GetDeviceName() (string, error) {
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

func (e *Pulsesink) GetMute() (bool, error) {
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

func (e *Pulsesink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// server (string)
//
// The PulseAudio server to connect to

func (e *Pulsesink) GetServer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("server")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Pulsesink) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

// stream-properties (GstStructure)
//
// List of pulseaudio stream properties. A list of defined properties can be
// found in the pulseaudio api docs.

// volume (float64)
//
// Linear volume of this stream, 1.0=100%%

func (e *Pulsesink) GetVolume() (float64, error) {
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

func (e *Pulsesink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

