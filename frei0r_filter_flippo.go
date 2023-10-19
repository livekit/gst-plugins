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

type Frei0RFilterFlippo struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterFlippo() (*Frei0RFilterFlippo, error) {
	e, err := gst.NewElement("frei0r-filter-flippo")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterFlippo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterFlippoWithName(name string) (*Frei0RFilterFlippo, error) {
	e, err := gst.NewElementWithName("frei0r-filter-flippo", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterFlippo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// x-axis (bool)
//
// Flipping on the horizontal axis

func (e *Frei0RFilterFlippo) GetXAxis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("x-axis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterFlippo) SetXAxis(value bool) error {
	return e.Element.SetProperty("x-axis", value)
}

// y-axis (bool)
//
// Flipping on the vertical axis

func (e *Frei0RFilterFlippo) GetYAxis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("y-axis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterFlippo) SetYAxis(value bool) error {
	return e.Element.SetProperty("y-axis", value)
}

