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

type Wasapisink struct {
	*base.GstBaseSink
}

func NewWasapisink() (*Wasapisink, error) {
	e, err := gst.NewElement("wasapisink")
	if err != nil {
		return nil, err
	}
	return &Wasapisink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewWasapisinkWithName(name string) (*Wasapisink, error) {
	e, err := gst.NewElementWithName("wasapisink", name)
	if err != nil {
		return nil, err
	}
	return &Wasapisink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// WASAPI device endpoint ID as provided by IMMDevice::GetId

func (e *Wasapisink) GetDevice() (string, error) {
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

func (e *Wasapisink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// exclusive (bool)
//
// Open the device in exclusive mode

func (e *Wasapisink) GetExclusive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("exclusive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wasapisink) SetExclusive(value bool) error {
	return e.Element.SetProperty("exclusive", value)
}

// low-latency (bool)
//
// Optimize all settings for lowest latency. Always safe to enable.

func (e *Wasapisink) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wasapisink) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// mute (bool)
//
// Mute state of this stream

func (e *Wasapisink) GetMute() (bool, error) {
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

func (e *Wasapisink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// role (GstWasapiDeviceRole)
//
// Role of the device: communications, multimedia, etc

func (e *Wasapisink) GetRole() (interface{}, error) {
	return e.Element.GetProperty("role")
}

func (e *Wasapisink) SetRole(value interface{}) error {
	return e.Element.SetProperty("role", value)
}

// use-audioclient3 (bool)
//
// Use the Windows 10 AudioClient3 API when available and if the low-latency property is set to TRUE

func (e *Wasapisink) GetUseAudioclient3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-audioclient3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wasapisink) SetUseAudioclient3(value bool) error {
	return e.Element.SetProperty("use-audioclient3", value)
}

