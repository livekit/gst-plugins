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

type Frei0RFilterRgbsplit0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterRgbsplit0R() (*Frei0RFilterRgbsplit0R, error) {
	e, err := gst.NewElement("frei0r-filter-rgbsplit0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterRgbsplit0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterRgbsplit0RWithName(name string) (*Frei0RFilterRgbsplit0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-rgbsplit0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterRgbsplit0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// horizontal-split-distance (float64)
//
// How far should layers be moved horizontally from each other

func (e *Frei0RFilterRgbsplit0R) GetHorizontalSplitDistance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("horizontal-split-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterRgbsplit0R) SetHorizontalSplitDistance(value float64) error {
	return e.Element.SetProperty("horizontal-split-distance", value)
}

// vertical-split-distance (float64)
//
// How far should layers be moved vertically from each other

func (e *Frei0RFilterRgbsplit0R) GetVerticalSplitDistance() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("vertical-split-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterRgbsplit0R) SetVerticalSplitDistance(value float64) error {
	return e.Element.SetProperty("vertical-split-distance", value)
}

