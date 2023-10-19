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

type Mfvideosrc struct {
	*base.GstPushSrc
}

func NewMfvideosrc() (*Mfvideosrc, error) {
	e, err := gst.NewElement("mfvideosrc")
	if err != nil {
		return nil, err
	}
	return &Mfvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewMfvideosrcWithName(name string) (*Mfvideosrc, error) {
	e, err := gst.NewElementWithName("mfvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Mfvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device-index (int)
//
// The zero-based device index

func (e *Mfvideosrc) GetDeviceIndex() (int, error) {
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

func (e *Mfvideosrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

// device-name (string)
//
// The human-readable device name

func (e *Mfvideosrc) GetDeviceName() (string, error) {
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

func (e *Mfvideosrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

// device-path (string)
//
// The device path

func (e *Mfvideosrc) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Mfvideosrc) SetDevicePath(value string) error {
	return e.Element.SetProperty("device-path", value)
}

// dispatcher (GstGpointer)
//
// ICoreDispatcher COM object used for activating device from UI thread.

func (e *Mfvideosrc) GetDispatcher() (interface{}, error) {
	return e.Element.GetProperty("dispatcher")
}

func (e *Mfvideosrc) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

