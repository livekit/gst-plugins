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

type Optv struct {
	*base.GstBaseTransform
}

func NewOptv() (*Optv, error) {
	e, err := gst.NewElement("optv")
	if err != nil {
		return nil, err
	}
	return &Optv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewOptvWithName(name string) (*Optv, error) {
	e, err := gst.NewElementWithName("optv", name)
	if err != nil {
		return nil, err
	}
	return &Optv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// mode (GstOpTvmode)
//
// Mode

func (e *Optv) GetMode() (GstOpTvmode, error) {
	var value GstOpTvmode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpTvmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpTvmode")
	}
	return value, nil
}

func (e *Optv) SetMode(value GstOpTvmode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// speed (int)
//
// Effect speed

func (e *Optv) GetSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Optv) SetSpeed(value int) error {
	return e.Element.SetProperty("speed", value)
}

// threshold (uint)
//
// Luma threshold

func (e *Optv) GetThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Optv) SetThreshold(value uint) error {
	return e.Element.SetProperty("threshold", value)
}

// ----- Constants -----

type GstOpTvmode string

const (
	GstOpTvmodeMaelstrom GstOpTvmode = "maelstrom" // maelstrom (0) – Maelstrom
	GstOpTvmodeRadiation GstOpTvmode = "radiation" // radiation (1) – Radiation
	GstOpTvmodeHorizontalStripes GstOpTvmode = "horizontal-stripes" // horizontal-stripes (2) – Horizontal Stripes
	GstOpTvmodeVerticalStripes GstOpTvmode = "vertical-stripes" // vertical-stripes (3) – Vertical Stripes
)

