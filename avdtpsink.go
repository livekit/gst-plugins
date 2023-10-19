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

type Avdtpsink struct {
	*base.GstBaseSink
}

func NewAvdtpsink() (*Avdtpsink, error) {
	e, err := gst.NewElement("avdtpsink")
	if err != nil {
		return nil, err
	}
	return &Avdtpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAvdtpsinkWithName(name string) (*Avdtpsink, error) {
	e, err := gst.NewElementWithName("avdtpsink", name)
	if err != nil {
		return nil, err
	}
	return &Avdtpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// auto-connect (bool)
//
// Automatically attempt to connect to device

func (e *Avdtpsink) GetAutoConnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avdtpsink) SetAutoConnect(value bool) error {
	return e.Element.SetProperty("auto-connect", value)
}

// device (string)
//
// Bluetooth remote device address

func (e *Avdtpsink) GetDevice() (string, error) {
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

func (e *Avdtpsink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// transport (string)
//
// Use configured transport

func (e *Avdtpsink) GetTransport() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Avdtpsink) SetTransport(value string) error {
	return e.Element.SetProperty("transport", value)
}

