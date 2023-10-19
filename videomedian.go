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

type Videomedian struct {
	*base.GstBaseTransform
}

func NewVideomedian() (*Videomedian, error) {
	e, err := gst.NewElement("videomedian")
	if err != nil {
		return nil, err
	}
	return &Videomedian{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideomedianWithName(name string) (*Videomedian, error) {
	e, err := gst.NewElementWithName("videomedian", name)
	if err != nil {
		return nil, err
	}
	return &Videomedian{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// filtersize (GstVideoMedianSize)
//
// The size of the filter

func (e *Videomedian) GetFiltersize() (GstVideoMedianSize, error) {
	var value GstVideoMedianSize
	var ok bool
	v, err := e.Element.GetProperty("filtersize")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoMedianSize)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoMedianSize")
	}
	return value, nil
}

func (e *Videomedian) SetFiltersize(value GstVideoMedianSize) error {
	e.Element.SetArg("filtersize", string(value))
	return nil
}

// lum-only (bool)
//
// Only apply filter on luminance

func (e *Videomedian) GetLumOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lum-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videomedian) SetLumOnly(value bool) error {
	return e.Element.SetProperty("lum-only", value)
}

// ----- Constants -----

type GstVideoMedianSize string

const (
	GstVideoMedianSize5 GstVideoMedianSize = "5" // 5 (5) – Median of 5 neighbour pixels
	GstVideoMedianSize9 GstVideoMedianSize = "9" // 9 (9) – Median of 9 neighbour pixels
)

