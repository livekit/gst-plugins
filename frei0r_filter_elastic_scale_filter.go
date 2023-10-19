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

type Frei0RFilterElasticScaleFilter struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterElasticScaleFilter() (*Frei0RFilterElasticScaleFilter, error) {
	e, err := gst.NewElement("frei0r-filter-elastic-scale-filter")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterElasticScaleFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterElasticScaleFilterWithName(name string) (*Frei0RFilterElasticScaleFilter, error) {
	e, err := gst.NewElementWithName("frei0r-filter-elastic-scale-filter", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterElasticScaleFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// center (float64)
//
// Horizontal center position of the linear area

func (e *Frei0RFilterElasticScaleFilter) GetCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterElasticScaleFilter) SetCenter(value float64) error {
	return e.Element.SetProperty("center", value)
}

// linear-scale-factor (float64)
//
// Amount how much the linear area is scaled

func (e *Frei0RFilterElasticScaleFilter) GetLinearScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("linear-scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterElasticScaleFilter) SetLinearScaleFactor(value float64) error {
	return e.Element.SetProperty("linear-scale-factor", value)
}

// linear-width (float64)
//
// Width of the linear area

func (e *Frei0RFilterElasticScaleFilter) GetLinearWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("linear-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterElasticScaleFilter) SetLinearWidth(value float64) error {
	return e.Element.SetProperty("linear-width", value)
}

// non-linear-scale-factor (float64)
//
// Amount how much the outer left and outer right areas are scaled non linearly

func (e *Frei0RFilterElasticScaleFilter) GetNonLinearScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("non-linear-scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterElasticScaleFilter) SetNonLinearScaleFactor(value float64) error {
	return e.Element.SetProperty("non-linear-scale-factor", value)
}

