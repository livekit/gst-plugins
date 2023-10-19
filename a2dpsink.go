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
)

type A2Dpsink struct {
	Element *gst.Element
}

func NewA2Dpsink() (*A2Dpsink, error) {
	e, err := gst.NewElement("a2dpsink")
	if err != nil {
		return nil, err
	}
	return &A2Dpsink{Element: e}, nil
}

func NewA2DpsinkWithName(name string) (*A2Dpsink, error) {
	e, err := gst.NewElementWithName("a2dpsink", name)
	if err != nil {
		return nil, err
	}
	return &A2Dpsink{Element: e}, nil
}

// ----- Properties -----

// auto-connect (bool)
//
// Automatically attempt to connect to device

func (e *A2Dpsink) GetAutoConnect() (bool, error) {
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

func (e *A2Dpsink) SetAutoConnect(value bool) error {
	return e.Element.SetProperty("auto-connect", value)
}

// device (string)
//
// Bluetooth remote device address

func (e *A2Dpsink) GetDevice() (string, error) {
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

func (e *A2Dpsink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// transport (string)
//
// Use configured transport

func (e *A2Dpsink) GetTransport() (string, error) {
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

func (e *A2Dpsink) SetTransport(value string) error {
	return e.Element.SetProperty("transport", value)
}

