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

type Frei0RFilterLevels struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterLevels() (*Frei0RFilterLevels, error) {
	e, err := gst.NewElement("frei0r-filter-levels")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLevels{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterLevelsWithName(name string) (*Frei0RFilterLevels, error) {
	e, err := gst.NewElementWithName("frei0r-filter-levels", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLevels{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// black-output (float64)
//
// Black output

func (e *Frei0RFilterLevels) GetBlackOutput() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("black-output")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetBlackOutput(value float64) error {
	return e.Element.SetProperty("black-output", value)
}

// channel (float64)
//
// Channel to adjust levels. 0%%=R, 10%%=G, 20%%=B, 30%%=Luma

func (e *Frei0RFilterLevels) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

// gamma (float64)
//
// Gamma

func (e *Frei0RFilterLevels) GetGamma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("gamma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetGamma(value float64) error {
	return e.Element.SetProperty("gamma", value)
}

// histogram-position (float64)
//
// Histogram position. 0%%=TL, 10%%=TR, 20%%=BL, 30%%=BR

func (e *Frei0RFilterLevels) GetHistogramPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("histogram-position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetHistogramPosition(value float64) error {
	return e.Element.SetProperty("histogram-position", value)
}

// input-black-level (float64)
//
// Input black level

func (e *Frei0RFilterLevels) GetInputBlackLevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("input-black-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetInputBlackLevel(value float64) error {
	return e.Element.SetProperty("input-black-level", value)
}

// input-white-level (float64)
//
// Input white level

func (e *Frei0RFilterLevels) GetInputWhiteLevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("input-white-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetInputWhiteLevel(value float64) error {
	return e.Element.SetProperty("input-white-level", value)
}

// show-histogram (bool)
//
// Show histogram

func (e *Frei0RFilterLevels) GetShowHistogram() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-histogram")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetShowHistogram(value bool) error {
	return e.Element.SetProperty("show-histogram", value)
}

// white-output (float64)
//
// White output

func (e *Frei0RFilterLevels) GetWhiteOutput() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("white-output")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLevels) SetWhiteOutput(value float64) error {
	return e.Element.SetProperty("white-output", value)
}

