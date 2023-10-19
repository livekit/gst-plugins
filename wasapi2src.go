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

type Wasapi2Src struct {
	*base.GstPushSrc
}

func NewWasapi2Src() (*Wasapi2Src, error) {
	e, err := gst.NewElement("wasapi2src")
	if err != nil {
		return nil, err
	}
	return &Wasapi2Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewWasapi2SrcWithName(name string) (*Wasapi2Src, error) {
	e, err := gst.NewElementWithName("wasapi2src", name)
	if err != nil {
		return nil, err
	}
	return &Wasapi2Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// Audio device ID as provided by Windows.Devices.Enumeration.DeviceInformation.Id

func (e *Wasapi2Src) GetDevice() (string, error) {
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

func (e *Wasapi2Src) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// dispatcher (GstGpointer)
//
// ICoreDispatcher COM object used for activating device from UI thread.

func (e *Wasapi2Src) GetDispatcher() (interface{}, error) {
	return e.Element.GetProperty("dispatcher")
}

func (e *Wasapi2Src) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

// loopback (bool)
//
// Open render device for loopback recording

func (e *Wasapi2Src) GetLoopback() (bool, error) {
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

func (e *Wasapi2Src) SetLoopback(value bool) error {
	return e.Element.SetProperty("loopback", value)
}

// loopback-mode (GstWasapi2SrcLoopbackMode)
//
// Loopback mode. "target-process-id" must be specified in case of
// process loopback modes.

// loopback-silence-on-device-mute (bool)
//
// When loopback recording, if the device is muted, inject silence in the pipeline

func (e *Wasapi2Src) GetLoopbackSilenceOnDeviceMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loopback-silence-on-device-mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wasapi2Src) SetLoopbackSilenceOnDeviceMute(value bool) error {
	return e.Element.SetProperty("loopback-silence-on-device-mute", value)
}

// loopback-target-pid (uint)
//
// Target process id to be recorded or excluded depending on loopback mode

// low-latency (bool)
//
// Optimize all settings for lowest latency. Always safe to enable.

func (e *Wasapi2Src) GetLowLatency() (bool, error) {
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

func (e *Wasapi2Src) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// mute (bool)
//
// Mute state of this stream

func (e *Wasapi2Src) GetMute() (bool, error) {
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

func (e *Wasapi2Src) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// volume (float64)
//
// Volume of this stream

func (e *Wasapi2Src) GetVolume() (float64, error) {
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

func (e *Wasapi2Src) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

// ----- Constants -----

type GstWasapi2SrcLoopbackMode string

const (
	GstWasapi2SrcLoopbackModeDefault GstWasapi2SrcLoopbackMode = "default" // default (0) – Default
	GstWasapi2SrcLoopbackModeIncludeProcessTree GstWasapi2SrcLoopbackMode = "include-process-tree" // include-process-tree (1) – Include process and its child processes
	GstWasapi2SrcLoopbackModeExcludeProcessTree GstWasapi2SrcLoopbackMode = "exclude-process-tree" // exclude-process-tree (2) – Exclude process and its child processes
)

