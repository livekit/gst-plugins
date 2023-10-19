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

type Uvch264Mjpgdemux struct {
	Element *gst.Element
}

func NewUvch264Mjpgdemux() (*Uvch264Mjpgdemux, error) {
	e, err := gst.NewElement("uvch264mjpgdemux")
	if err != nil {
		return nil, err
	}
	return &Uvch264Mjpgdemux{Element: e}, nil
}

func NewUvch264MjpgdemuxWithName(name string) (*Uvch264Mjpgdemux, error) {
	e, err := gst.NewElementWithName("uvch264mjpgdemux", name)
	if err != nil {
		return nil, err
	}
	return &Uvch264Mjpgdemux{Element: e}, nil
}

// ----- Properties -----

// device-fd (int)
//
// File descriptor of the v4l2 device

func (e *Uvch264Mjpgdemux) GetDeviceFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Mjpgdemux) SetDeviceFd(value int) error {
	return e.Element.SetProperty("device-fd", value)
}

// num-clock-samples (int)
//
// Number of clock samples to gather for the PTS synchronization (-1 = unlimited)

func (e *Uvch264Mjpgdemux) GetNumClockSamples() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-clock-samples")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Mjpgdemux) SetNumClockSamples(value int) error {
	return e.Element.SetProperty("num-clock-samples", value)
}

