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

type Dshowvideosrc struct {
	*base.GstPushSrc
}

func NewDshowvideosrc() (*Dshowvideosrc, error) {
	e, err := gst.NewElement("dshowvideosrc")
	if err != nil {
		return nil, err
	}
	return &Dshowvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDshowvideosrcWithName(name string) (*Dshowvideosrc, error) {
	e, err := gst.NewElementWithName("dshowvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Dshowvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device (string)
//
// Directshow device path (@..classID/name)

func (e *Dshowvideosrc) GetDevice() (string, error) {
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

func (e *Dshowvideosrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-index (int)
//
// Index of the enumerated video device

func (e *Dshowvideosrc) GetDeviceIndex() (int, error) {
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

func (e *Dshowvideosrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

// device-name (string)
//
// Human-readable name of the video device

func (e *Dshowvideosrc) GetDeviceName() (string, error) {
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

func (e *Dshowvideosrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

