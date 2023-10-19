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

type Ksvideosrc struct {
	*base.GstPushSrc
}

func NewKsvideosrc() (*Ksvideosrc, error) {
	e, err := gst.NewElement("ksvideosrc")
	if err != nil {
		return nil, err
	}
	return &Ksvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewKsvideosrcWithName(name string) (*Ksvideosrc, error) {
	e, err := gst.NewElementWithName("ksvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Ksvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// device-index (int)
//
// The zero-based device index

func (e *Ksvideosrc) GetDeviceIndex() (int, error) {
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

func (e *Ksvideosrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

// device-name (string)
//
// The human-readable device name

func (e *Ksvideosrc) GetDeviceName() (string, error) {
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

func (e *Ksvideosrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

// device-path (string)
//
// The device path

func (e *Ksvideosrc) GetDevicePath() (string, error) {
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

func (e *Ksvideosrc) SetDevicePath(value string) error {
	return e.Element.SetProperty("device-path", value)
}

// do-stats (bool)
//
// Enable logging of statistics

func (e *Ksvideosrc) GetDoStats() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ksvideosrc) SetDoStats(value bool) error {
	return e.Element.SetProperty("do-stats", value)
}

// enable-quirks (bool)
//
// Enable driver-specific quirks

func (e *Ksvideosrc) GetEnableQuirks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-quirks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ksvideosrc) SetEnableQuirks(value bool) error {
	return e.Element.SetProperty("enable-quirks", value)
}

// fps (int)
//
// Last measured framerate, if statistics are enabled

func (e *Ksvideosrc) GetFps() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

