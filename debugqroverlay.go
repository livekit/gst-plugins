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

type Debugqroverlay struct {
	*base.GstBaseTransform
}

func NewDebugqroverlay() (*Debugqroverlay, error) {
	e, err := gst.NewElement("debugqroverlay")
	if err != nil {
		return nil, err
	}
	return &Debugqroverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDebugqroverlayWithName(name string) (*Debugqroverlay, error) {
	e, err := gst.NewElementWithName("debugqroverlay", name)
	if err != nil {
		return nil, err
	}
	return &Debugqroverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// extra-data-array (string)
//
// List of comma separated values that the extra data value will be  cycled from at each interval, example array structure : "240,480,720,960,1200,1440,1680,1920"

func (e *Debugqroverlay) GetExtraDataArray() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("extra-data-array")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Debugqroverlay) SetExtraDataArray(value string) error {
	return e.Element.SetProperty("extra-data-array", value)
}

// extra-data-interval-buffers (int64)
//
// Extra data append into the Qrcode at the first buffer of each  interval

func (e *Debugqroverlay) GetExtraDataIntervalBuffers() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("extra-data-interval-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Debugqroverlay) SetExtraDataIntervalBuffers(value int64) error {
	return e.Element.SetProperty("extra-data-interval-buffers", value)
}

// extra-data-name (string)
//
// Json key name for extra append data

func (e *Debugqroverlay) GetExtraDataName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("extra-data-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Debugqroverlay) SetExtraDataName(value string) error {
	return e.Element.SetProperty("extra-data-name", value)
}

// extra-data-span-buffers (int64)
//
// Numbers of consecutive buffers that the extra data will be inserted  (counting the first buffer)

func (e *Debugqroverlay) GetExtraDataSpanBuffers() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("extra-data-span-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Debugqroverlay) SetExtraDataSpanBuffers(value int64) error {
	return e.Element.SetProperty("extra-data-span-buffers", value)
}

