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

type Frei0RFilterLensCorrection struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterLensCorrection() (*Frei0RFilterLensCorrection, error) {
	e, err := gst.NewElement("frei0r-filter-lens-correction")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLensCorrection{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterLensCorrectionWithName(name string) (*Frei0RFilterLensCorrection, error) {
	e, err := gst.NewElementWithName("frei0r-filter-lens-correction", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLensCorrection{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// brightness (float64)

func (e *Frei0RFilterLensCorrection) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLensCorrection) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

// correction-near-center (float64)

func (e *Frei0RFilterLensCorrection) GetCorrectionNearCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("correction-near-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLensCorrection) SetCorrectionNearCenter(value float64) error {
	return e.Element.SetProperty("correction-near-center", value)
}

// correction-near-edges (float64)

func (e *Frei0RFilterLensCorrection) GetCorrectionNearEdges() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("correction-near-edges")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLensCorrection) SetCorrectionNearEdges(value float64) error {
	return e.Element.SetProperty("correction-near-edges", value)
}

// x-center (float64)

func (e *Frei0RFilterLensCorrection) GetXCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLensCorrection) SetXCenter(value float64) error {
	return e.Element.SetProperty("x-center", value)
}

// y-center (float64)

func (e *Frei0RFilterLensCorrection) GetYCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLensCorrection) SetYCenter(value float64) error {
	return e.Element.SetProperty("y-center", value)
}
