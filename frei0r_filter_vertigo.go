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

type Frei0RFilterVertigo struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterVertigo() (*Frei0RFilterVertigo, error) {
	e, err := gst.NewElement("frei0r-filter-vertigo")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterVertigo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterVertigoWithName(name string) (*Frei0RFilterVertigo, error) {
	e, err := gst.NewElementWithName("frei0r-filter-vertigo", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterVertigo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// phaseincrement (float64)
//
// Phase increment

func (e *Frei0RFilterVertigo) GetPhaseincrement() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("phaseincrement")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterVertigo) SetPhaseincrement(value float64) error {
	return e.Element.SetProperty("phaseincrement", value)
}

// zoomrate (float64)
//
// Zoomrate

func (e *Frei0RFilterVertigo) GetZoomrate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zoomrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterVertigo) SetZoomrate(value float64) error {
	return e.Element.SetProperty("zoomrate", value)
}

