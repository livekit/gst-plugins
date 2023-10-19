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

type Frei0RFilterNdviFilter struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterNdviFilter() (*Frei0RFilterNdviFilter, error) {
	e, err := gst.NewElement("frei0r-filter-ndvi-filter")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterNdviFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterNdviFilterWithName(name string) (*Frei0RFilterNdviFilter, error) {
	e, err := gst.NewElementWithName("frei0r-filter-ndvi-filter", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterNdviFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color-map (string)
//
// The color map to use. One of 'earth', 'grayscale', 'heat' or 'rainbow'.

func (e *Frei0RFilterNdviFilter) GetColorMap() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("color-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetColorMap(value string) error {
	return e.Element.SetProperty("color-map", value)
}

// index-calculation (string)
//
// The index calculation to use. One of 'ndvi' or 'vi'.

func (e *Frei0RFilterNdviFilter) GetIndexCalculation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("index-calculation")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetIndexCalculation(value string) error {
	return e.Element.SetProperty("index-calculation", value)
}

// legend (string)
//
// Control legend display. One of 'off' or 'bottom'.

func (e *Frei0RFilterNdviFilter) GetLegend() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("legend")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetLegend(value string) error {
	return e.Element.SetProperty("legend", value)
}

// levels (float64)
//
// The number of color levels to use in the false image (divided by 1000).

func (e *Frei0RFilterNdviFilter) GetLevels() (float64, error) {
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

func (e *Frei0RFilterNdviFilter) SetLevels(value float64) error {
	return e.Element.SetProperty("levels", value)
}

// nir-channel (string)
//
// The channel to use for the near-infrared component. One of 'r', 'g', or 'b'.

func (e *Frei0RFilterNdviFilter) GetNirChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("nir-channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetNirChannel(value string) error {
	return e.Element.SetProperty("nir-channel", value)
}

// nir-offset (float64)
//
// An offset to be applied to the near-infrared component (mapped to [-100%%, 100%%].

func (e *Frei0RFilterNdviFilter) GetNirOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("nir-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetNirOffset(value float64) error {
	return e.Element.SetProperty("nir-offset", value)
}

// nir-scale (float64)
//
// A scaling factor to be applied to the near-infrared component (divided by 10).

func (e *Frei0RFilterNdviFilter) GetNirScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("nir-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetNirScale(value float64) error {
	return e.Element.SetProperty("nir-scale", value)
}

// vis-offset (float64)
//
// An offset to be applied to the visible component (mapped to [-100%%, 100%%].

func (e *Frei0RFilterNdviFilter) GetVisOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("vis-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetVisOffset(value float64) error {
	return e.Element.SetProperty("vis-offset", value)
}

// vis-scale (float64)
//
// A scaling factor to be applied to the visible component (divided by 10).

func (e *Frei0RFilterNdviFilter) GetVisScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("vis-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetVisScale(value float64) error {
	return e.Element.SetProperty("vis-scale", value)
}

// visible-channel (string)
//
// The channel to use for the visible component. One of 'r', 'g', or 'b'.

func (e *Frei0RFilterNdviFilter) GetVisibleChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("visible-channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterNdviFilter) SetVisibleChannel(value string) error {
	return e.Element.SetProperty("visible-channel", value)
}

