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

type Alsasrc struct {
	*base.GstPushSrc
}

func NewAlsasrc() (*Alsasrc, error) {
	e, err := gst.NewElement("alsasrc")
	if err != nil {
		return nil, err
	}
	return &Alsasrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewAlsasrcWithName(name string) (*Alsasrc, error) {
	e, err := gst.NewElementWithName("alsasrc", name)
	if err != nil {
		return nil, err
	}
	return &Alsasrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// card-name (string)
//
// Human-readable name of the sound card

func (e *Alsasrc) GetCardName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("card-name")
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
// ALSA device, as defined in an asound configuration file

func (e *Alsasrc) GetDevice() (string, error) {
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

func (e *Alsasrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Alsasrc) GetDeviceName() (string, error) {
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

// use-driver-timestamps (bool)
//
// Use driver timestamps or the pipeline clock timestamps

func (e *Alsasrc) GetUseDriverTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-driver-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Alsasrc) SetUseDriverTimestamps(value bool) error {
	return e.Element.SetProperty("use-driver-timestamps", value)
}

