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

type Wasapi2Sink struct {
	*base.GstBaseSink
}

func NewWasapi2Sink() (*Wasapi2Sink, error) {
	e, err := gst.NewElement("wasapi2sink")
	if err != nil {
		return nil, err
	}
	return &Wasapi2Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewWasapi2SinkWithName(name string) (*Wasapi2Sink, error) {
	e, err := gst.NewElementWithName("wasapi2sink", name)
	if err != nil {
		return nil, err
	}
	return &Wasapi2Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// Audio device ID as provided by Windows.Devices.Enumeration.DeviceInformation.Id

func (e *Wasapi2Sink) GetDevice() (string, error) {
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

func (e *Wasapi2Sink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// dispatcher (GstGpointer)
//
// ICoreDispatcher COM object used for activating device from UI thread.

func (e *Wasapi2Sink) GetDispatcher() (interface{}, error) {
	return e.Element.GetProperty("dispatcher")
}

func (e *Wasapi2Sink) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

// low-latency (bool)
//
// Optimize all settings for lowest latency. Always safe to enable.

func (e *Wasapi2Sink) GetLowLatency() (bool, error) {
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

func (e *Wasapi2Sink) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// mute (bool)
//
// Mute state of this stream

func (e *Wasapi2Sink) GetMute() (bool, error) {
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

func (e *Wasapi2Sink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// volume (float64)
//
// Volume of this stream

func (e *Wasapi2Sink) GetVolume() (float64, error) {
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

func (e *Wasapi2Sink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

