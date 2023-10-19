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

type Pulsesrc struct {
	*base.GstPushSrc
}

func NewPulsesrc() (*Pulsesrc, error) {
	e, err := gst.NewElement("pulsesrc")
	if err != nil {
		return nil, err
	}
	return &Pulsesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewPulsesrcWithName(name string) (*Pulsesrc, error) {
	e, err := gst.NewElementWithName("pulsesrc", name)
	if err != nil {
		return nil, err
	}
	return &Pulsesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// client-name (string)
//
// The PulseAudio client name to use.

func (e *Pulsesrc) GetClientName() (string, error) {
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

func (e *Pulsesrc) SetClientName(value string) error {
	return e.Element.SetProperty("client-name", value)
}

// current-device (string)
//
// The current PulseAudio source device

func (e *Pulsesrc) GetCurrentDevice() (string, error) {
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
// The PulseAudio source device to connect to

func (e *Pulsesrc) GetDevice() (string, error) {
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

func (e *Pulsesrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Pulsesrc) GetDeviceName() (string, error) {
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
// Whether the stream is muted or not.

func (e *Pulsesrc) GetMute() (bool, error) {
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

func (e *Pulsesrc) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// server (string)
//
// The PulseAudio server to connect to

func (e *Pulsesrc) GetServer() (string, error) {
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

func (e *Pulsesrc) SetServer(value string) error {
	return e.Element.SetProperty("server", value)
}

// source-output-index (uint)
//
// The index of the PulseAudio source output corresponding to this element.

func (e *Pulsesrc) GetSourceOutputIndex() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("source-output-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// stream-properties (GstStructure)
//
// List of pulseaudio stream properties. A list of defined properties can be
// found in the pulseaudio api docs.

// volume (float64)
//
// The volume of the record stream.

func (e *Pulsesrc) GetVolume() (float64, error) {
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

func (e *Pulsesrc) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

