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

type Frei0RFilterDither struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterDither() (*Frei0RFilterDither, error) {
	e, err := gst.NewElement("frei0r-filter-dither")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterDither{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterDitherWithName(name string) (*Frei0RFilterDither, error) {
	e, err := gst.NewElementWithName("frei0r-filter-dither", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterDither{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// levels (float64)
//
// Number of values per channel

func (e *Frei0RFilterDither) GetLevels() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("levels")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDither) SetLevels(value float64) error {
	return e.Element.SetProperty("levels", value)
}

// matrixid (float64)
//
// Id of matrix used for dithering

func (e *Frei0RFilterDither) GetMatrixid() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("matrixid")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterDither) SetMatrixid(value float64) error {
	return e.Element.SetProperty("matrixid", value)
}

