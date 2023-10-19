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

type Wasapisrc struct {
	*base.GstPushSrc
}

func NewWasapisrc() (*Wasapisrc, error) {
	e, err := gst.NewElement("wasapisrc")
	if err != nil {
		return nil, err
	}
	return &Wasapisrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewWasapisrcWithName(name string) (*Wasapisrc, error) {
	e, err := gst.NewElementWithName("wasapisrc", name)
	if err != nil {
		return nil, err
	}
	return &Wasapisrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// WASAPI device endpoint ID as provided by IMMDevice::GetId

func (e *Wasapisrc) GetDevice() (string, error) {
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

func (e *Wasapisrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// exclusive (bool)
//
// Open the device in exclusive mode

func (e *Wasapisrc) GetExclusive() (bool, error) {
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

func (e *Wasapisrc) SetExclusive(value bool) error {
	return e.Element.SetProperty("exclusive", value)
}

// loopback (bool)
//
// Open the sink device for loopback recording

func (e *Wasapisrc) GetLoopback() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loopback")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wasapisrc) SetLoopback(value bool) error {
	return e.Element.SetProperty("loopback", value)
}

// low-latency (bool)
//
// Optimize all settings for lowest latency. Always safe to enable.

func (e *Wasapisrc) GetLowLatency() (bool, error) {
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

func (e *Wasapisrc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// role (GstWasapiDeviceRole)
//
// Role of the device: communications, multimedia, etc

func (e *Wasapisrc) GetRole() (GstWasapiDeviceRole, error) {
	var value GstWasapiDeviceRole
	var ok bool
	v, err := e.Element.GetProperty("role")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWasapiDeviceRole)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWasapiDeviceRole")
	}
	return value, nil
}

func (e *Wasapisrc) SetRole(value GstWasapiDeviceRole) error {
	e.Element.SetArg("role", string(value))
	return nil
}

// use-audioclient3 (bool)
//
// Whether to use the Windows 10 AudioClient3 API when available

func (e *Wasapisrc) GetUseAudioclient3() (bool, error) {
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

func (e *Wasapisrc) SetUseAudioclient3(value bool) error {
	return e.Element.SetProperty("use-audioclient3", value)
}

// ----- Constants -----

type GstWasapiDeviceRole string

const (
	GstWasapiDeviceRoleConsole GstWasapiDeviceRole = "console" // console (0) – Games, system notifications, voice commands
	GstWasapiDeviceRoleMultimedia GstWasapiDeviceRole = "multimedia" // multimedia (1) – Music, movies, recorded media
	GstWasapiDeviceRoleComms GstWasapiDeviceRole = "comms" // comms (2) – Voice communications
)

