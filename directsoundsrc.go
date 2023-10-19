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

type Directsoundsrc struct {
	*base.GstPushSrc
}

func NewDirectsoundsrc() (*Directsoundsrc, error) {
	e, err := gst.NewElement("directsoundsrc")
	if err != nil {
		return nil, err
	}
	return &Directsoundsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDirectsoundsrcWithName(name string) (*Directsoundsrc, error) {
	e, err := gst.NewElementWithName("directsoundsrc", name)
	if err != nil {
		return nil, err
	}
	return &Directsoundsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// DirectSound playback device as a GUID string (volume and mute will not work!)

func (e *Directsoundsrc) GetDevice() (string, error) {
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

func (e *Directsoundsrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Directsoundsrc) GetDeviceName() (string, error) {
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

func (e *Directsoundsrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

// mute (bool)
//
// Mute state of this stream

func (e *Directsoundsrc) GetMute() (bool, error) {
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

func (e *Directsoundsrc) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// volume (float64)
//
// Volume of this stream

func (e *Directsoundsrc) GetVolume() (float64, error) {
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

func (e *Directsoundsrc) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

