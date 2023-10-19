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

type Rippletv struct {
	*base.GstBaseTransform
}

func NewRippletv() (*Rippletv, error) {
	e, err := gst.NewElement("rippletv")
	if err != nil {
		return nil, err
	}
	return &Rippletv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRippletvWithName(name string) (*Rippletv, error) {
	e, err := gst.NewElementWithName("rippletv", name)
	if err != nil {
		return nil, err
	}
	return &Rippletv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// mode (GstRippleTvmode)
//
// Mode

func (e *Rippletv) GetMode() (GstRippleTvmode, error) {
	var value GstRippleTvmode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRippleTvmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRippleTvmode")
	}
	return value, nil
}

func (e *Rippletv) SetMode(value GstRippleTvmode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// reset (bool)
//
// Reset all current ripples

func (e *Rippletv) GetReset() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reset")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rippletv) SetReset(value bool) error {
	return e.Element.SetProperty("reset", value)
}

// ----- Constants -----

type GstRippleTvmode string

const (
	GstRippleTvmodeMotionDetection GstRippleTvmode = "motion-detection" // motion-detection (0) – Motion Detection
	GstRippleTvmodeRain GstRippleTvmode = "rain" // rain (1) – Rain
)

