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

type Frei0RFilterMedians struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterMedians() (*Frei0RFilterMedians, error) {
	e, err := gst.NewElement("frei0r-filter-medians")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterMedians{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterMediansWithName(name string) (*Frei0RFilterMedians, error) {
	e, err := gst.NewElementWithName("frei0r-filter-medians", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterMedians{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// size (float64)
//
// Size for 'var size' type filter

func (e *Frei0RFilterMedians) GetSize() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterMedians) SetSize(value float64) error {
	return e.Element.SetProperty("size", value)
}

// type (string)
//
// Choose type of median: Cross5, Square3x3, Bilevel, Diamond3x3, Square5x5, Temp3, Temp5, ArceBI, ML3D, ML3dEX, VarSize

func (e *Frei0RFilterMedians) GetType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterMedians) SetType(value string) error {
	return e.Element.SetProperty("type", value)
}

