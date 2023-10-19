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

type Alsasink struct {
	*base.GstBaseSink
}

func NewAlsasink() (*Alsasink, error) {
	e, err := gst.NewElement("alsasink")
	if err != nil {
		return nil, err
	}
	return &Alsasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAlsasinkWithName(name string) (*Alsasink, error) {
	e, err := gst.NewElementWithName("alsasink", name)
	if err != nil {
		return nil, err
	}
	return &Alsasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// card-name (string)
//
// Human-readable name of the sound card

func (e *Alsasink) GetCardName() (string, error) {
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

func (e *Alsasink) GetDevice() (string, error) {
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

func (e *Alsasink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Alsasink) GetDeviceName() (string, error) {
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

