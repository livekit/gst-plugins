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

type Iqa struct {
	*base.GstAggregator
}

func NewIqa() (*Iqa, error) {
	e, err := gst.NewElement("iqa")
	if err != nil {
		return nil, err
	}
	return &Iqa{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewIqaWithName(name string) (*Iqa, error) {
	e, err := gst.NewElementWithName("iqa", name)
	if err != nil {
		return nil, err
	}
	return &Iqa{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// do-dssim (bool)
//
// Run structural similarity checks

func (e *Iqa) GetDoDssim() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-dssim")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Iqa) SetDoDssim(value bool) error {
	return e.Element.SetProperty("do-dssim", value)
}

// dssim-error-threshold (float64)
//
// dssim value over which the element will post an error message on the bus. A value < 0.0 means 'disabled'.

func (e *Iqa) GetDssimErrorThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dssim-error-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Iqa) SetDssimErrorThreshold(value float64) error {
	return e.Element.SetProperty("dssim-error-threshold", value)
}

// mode (GstIqaMode)
//
// Controls the frame comparison mode.

func (e *Iqa) GetMode() (GstIqaMode, error) {
	var value GstIqaMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstIqaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstIqaMode")
	}
	return value, nil
}

func (e *Iqa) SetMode(value GstIqaMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ----- Constants -----

type GstIqaMode string

const (
	GstIqaModeStrict GstIqaMode = "strict" // strict (0x00000002) – Strict comparison of frames.
)

