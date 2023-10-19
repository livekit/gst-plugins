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

type Dshowaudiosrc struct {
	*base.GstPushSrc
}

func NewDshowaudiosrc() (*Dshowaudiosrc, error) {
	e, err := gst.NewElement("dshowaudiosrc")
	if err != nil {
		return nil, err
	}
	return &Dshowaudiosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDshowaudiosrcWithName(name string) (*Dshowaudiosrc, error) {
	e, err := gst.NewElementWithName("dshowaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &Dshowaudiosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// Directshow device reference (classID/name)

func (e *Dshowaudiosrc) GetDevice() (string, error) {
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

func (e *Dshowaudiosrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-index (int)
//
// Index of the enumerated audio device

func (e *Dshowaudiosrc) GetDeviceIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dshowaudiosrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

// device-name (string)
//
// Human-readable name of the sound device

func (e *Dshowaudiosrc) GetDeviceName() (string, error) {
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

func (e *Dshowaudiosrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

