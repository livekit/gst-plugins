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

type Frei0RFilterSopSat struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterSopSat() (*Frei0RFilterSopSat, error) {
	e, err := gst.NewElement("frei0r-filter-sop-sat")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSopSat{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterSopSatWithName(name string) (*Frei0RFilterSopSat, error) {
	e, err := gst.NewElementWithName("frei0r-filter-sop-sat", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterSopSat{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aoffset (float64)
//
// Offset of the alpha component

func (e *Frei0RFilterSopSat) GetAoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetAoffset(value float64) error {
	return e.Element.SetProperty("aoffset", value)
}

// apower (float64)
//
// Power (Gamma) of the alpha component

func (e *Frei0RFilterSopSat) GetApower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("apower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetApower(value float64) error {
	return e.Element.SetProperty("apower", value)
}

// aslope (float64)
//
// Slope of the alpha component

func (e *Frei0RFilterSopSat) GetAslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetAslope(value float64) error {
	return e.Element.SetProperty("aslope", value)
}

// boffset (float64)
//
// Offset of the blue color component

func (e *Frei0RFilterSopSat) GetBoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("boffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetBoffset(value float64) error {
	return e.Element.SetProperty("boffset", value)
}

// bpower (float64)
//
// Power (Gamma) of the blue color component

func (e *Frei0RFilterSopSat) GetBpower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bpower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetBpower(value float64) error {
	return e.Element.SetProperty("bpower", value)
}

// bslope (float64)
//
// Slope of the blue color component

func (e *Frei0RFilterSopSat) GetBslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetBslope(value float64) error {
	return e.Element.SetProperty("bslope", value)
}

// goffset (float64)
//
// Offset of the green color component

func (e *Frei0RFilterSopSat) GetGoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("goffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetGoffset(value float64) error {
	return e.Element.SetProperty("goffset", value)
}

// gpower (float64)
//
// Power (Gamma) of the green color component

func (e *Frei0RFilterSopSat) GetGpower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("gpower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetGpower(value float64) error {
	return e.Element.SetProperty("gpower", value)
}

// gslope (float64)
//
// Slope of the green color component

func (e *Frei0RFilterSopSat) GetGslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("gslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetGslope(value float64) error {
	return e.Element.SetProperty("gslope", value)
}

// roffset (float64)
//
// Offset of the red color component

func (e *Frei0RFilterSopSat) GetRoffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("roffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetRoffset(value float64) error {
	return e.Element.SetProperty("roffset", value)
}

// rpower (float64)
//
// Power (Gamma) of the red color component

func (e *Frei0RFilterSopSat) GetRpower() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rpower")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetRpower(value float64) error {
	return e.Element.SetProperty("rpower", value)
}

// rslope (float64)
//
// Slope of the red color component

func (e *Frei0RFilterSopSat) GetRslope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rslope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetRslope(value float64) error {
	return e.Element.SetProperty("rslope", value)
}

// saturation (float64)
//
// Overall saturation

func (e *Frei0RFilterSopSat) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterSopSat) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

