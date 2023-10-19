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

type Waterripple struct {
	*base.GstBaseTransform
}

func NewWaterripple() (*Waterripple, error) {
	e, err := gst.NewElement("waterripple")
	if err != nil {
		return nil, err
	}
	return &Waterripple{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewWaterrippleWithName(name string) (*Waterripple, error) {
	e, err := gst.NewElementWithName("waterripple", name)
	if err != nil {
		return nil, err
	}
	return &Waterripple{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amplitude (float64)
//
// amplitude

func (e *Waterripple) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Waterripple) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

// phase (float64)
//
// phase

func (e *Waterripple) GetPhase() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("phase")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Waterripple) SetPhase(value float64) error {
	return e.Element.SetProperty("phase", value)
}

// wavelength (float64)
//
// wavelength

func (e *Waterripple) GetWavelength() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("wavelength")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Waterripple) SetWavelength(value float64) error {
	return e.Element.SetProperty("wavelength", value)
}

