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

type Radioactv struct {
	*base.GstBaseTransform
}

func NewRadioactv() (*Radioactv, error) {
	e, err := gst.NewElement("radioactv")
	if err != nil {
		return nil, err
	}
	return &Radioactv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRadioactvWithName(name string) (*Radioactv, error) {
	e, err := gst.NewElementWithName("radioactv", name)
	if err != nil {
		return nil, err
	}
	return &Radioactv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color (GstRadioacTvcolor)
//
// Color

func (e *Radioactv) GetColor() (GstRadioacTvcolor, error) {
	var value GstRadioacTvcolor
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRadioacTvcolor)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRadioacTvcolor")
	}
	return value, nil
}

func (e *Radioactv) SetColor(value GstRadioacTvcolor) error {
	e.Element.SetArg("color", string(value))
	return nil
}

// interval (uint)
//
// Snapshot interval (in strobe mode)

func (e *Radioactv) GetInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Radioactv) SetInterval(value uint) error {
	return e.Element.SetProperty("interval", value)
}

// mode (GstRadioacTvmode)
//
// Mode

func (e *Radioactv) GetMode() (GstRadioacTvmode, error) {
	var value GstRadioacTvmode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRadioacTvmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRadioacTvmode")
	}
	return value, nil
}

func (e *Radioactv) SetMode(value GstRadioacTvmode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// trigger (bool)
//
// Trigger (in trigger mode)

func (e *Radioactv) GetTrigger() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("trigger")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Radioactv) SetTrigger(value bool) error {
	return e.Element.SetProperty("trigger", value)
}

// ----- Constants -----

type GstRadioacTvcolor string

const (
	GstRadioacTvcolorRed GstRadioacTvcolor = "red" // red (0) – Red
	GstRadioacTvcolorGreen GstRadioacTvcolor = "green" // green (1) – Green
	GstRadioacTvcolorBlue GstRadioacTvcolor = "blue" // blue (2) – Blue
	GstRadioacTvcolorWhite GstRadioacTvcolor = "white" // white (3) – White
)

type GstRadioacTvmode string

const (
	GstRadioacTvmodeNormal GstRadioacTvmode = "normal" // normal (0) – Normal
	GstRadioacTvmodeStrobe1 GstRadioacTvmode = "strobe1" // strobe1 (1) – Strobe 1
	GstRadioacTvmodeStrobe2 GstRadioacTvmode = "strobe2" // strobe2 (2) – Strobe 2
	GstRadioacTvmodeTrigger GstRadioacTvmode = "trigger" // trigger (3) – Trigger
)

