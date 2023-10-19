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

type Openalsink struct {
	*base.GstBaseSink
}

func NewOpenalsink() (*Openalsink, error) {
	e, err := gst.NewElement("openalsink")
	if err != nil {
		return nil, err
	}
	return &Openalsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewOpenalsinkWithName(name string) (*Openalsink, error) {
	e, err := gst.NewElementWithName("openalsink", name)
	if err != nil {
		return nil, err
	}
	return &Openalsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// Human-readable name of the device

func (e *Openalsink) GetDevice() (string, error) {
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

func (e *Openalsink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Human-readable name of the opened device

func (e *Openalsink) GetDeviceName() (string, error) {
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

// user-context (GstGpointer)
//
// User context

func (e *Openalsink) GetUserContext() (interface{}, error) {
	return e.Element.GetProperty("user-context")
}

func (e *Openalsink) SetUserContext(value interface{}) error {
	return e.Element.SetProperty("user-context", value)
}

// user-device (GstGpointer)
//
// User device

func (e *Openalsink) GetUserDevice() (interface{}, error) {
	return e.Element.GetProperty("user-device")
}

func (e *Openalsink) SetUserDevice(value interface{}) error {
	return e.Element.SetProperty("user-device", value)
}

// user-source (uint)
//
// User source

func (e *Openalsink) GetUserSource() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("user-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openalsink) SetUserSource(value uint) error {
	return e.Element.SetProperty("user-source", value)
}

