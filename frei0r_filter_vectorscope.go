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

type Frei0RFilterVectorscope struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterVectorscope() (*Frei0RFilterVectorscope, error) {
	e, err := gst.NewElement("frei0r-filter-vectorscope")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterVectorscope{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterVectorscopeWithName(name string) (*Frei0RFilterVectorscope, error) {
	e, err := gst.NewElementWithName("frei0r-filter-vectorscope", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterVectorscope{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// mix (float64)
//
// The amount of source image mixed into background of display

func (e *Frei0RFilterVectorscope) GetMix() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("mix")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterVectorscope) SetMix(value float64) error {
	return e.Element.SetProperty("mix", value)
}

// overlay-sides (bool)
//
// If false, the sides of image are shown without overlay

func (e *Frei0RFilterVectorscope) GetOverlaySides() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("overlay-sides")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterVectorscope) SetOverlaySides(value bool) error {
	return e.Element.SetProperty("overlay-sides", value)
}

