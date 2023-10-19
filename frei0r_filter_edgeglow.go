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

type Frei0RFilterEdgeglow struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterEdgeglow() (*Frei0RFilterEdgeglow, error) {
	e, err := gst.NewElement("frei0r-filter-edgeglow")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterEdgeglow{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterEdgeglowWithName(name string) (*Frei0RFilterEdgeglow, error) {
	e, err := gst.NewElementWithName("frei0r-filter-edgeglow", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterEdgeglow{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// lredscale (float64)
//
// multiplier for downscaling non-edge brightness

func (e *Frei0RFilterEdgeglow) GetLredscale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lredscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterEdgeglow) SetLredscale(value float64) error {
	return e.Element.SetProperty("lredscale", value)
}

// lthresh (float64)
//
// threshold for edge lightening

func (e *Frei0RFilterEdgeglow) GetLthresh() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lthresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterEdgeglow) SetLthresh(value float64) error {
	return e.Element.SetProperty("lthresh", value)
}

// lupscale (float64)
//
// multiplier for upscaling edge brightness

func (e *Frei0RFilterEdgeglow) GetLupscale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("lupscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterEdgeglow) SetLupscale(value float64) error {
	return e.Element.SetProperty("lupscale", value)
}

