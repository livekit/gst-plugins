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

type Retinex struct {
	*base.GstBaseTransform
}

func NewRetinex() (*Retinex, error) {
	e, err := gst.NewElement("retinex")
	if err != nil {
		return nil, err
	}
	return &Retinex{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRetinexWithName(name string) (*Retinex, error) {
	e, err := gst.NewElementWithName("retinex", name)
	if err != nil {
		return nil, err
	}
	return &Retinex{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// gain (int)
//
// Gain

func (e *Retinex) GetGain() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Retinex) SetGain(value int) error {
	return e.Element.SetProperty("gain", value)
}

// method (GstRetinexMethod)
//
// Retinex method to use

func (e *Retinex) GetMethod() (GstRetinexMethod, error) {
	var value GstRetinexMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRetinexMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRetinexMethod")
	}
	return value, nil
}

func (e *Retinex) SetMethod(value GstRetinexMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// offset (int)
//
// Offset

func (e *Retinex) GetOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Retinex) SetOffset(value int) error {
	return e.Element.SetProperty("offset", value)
}

// scales (int)
//
// Amount of gaussian filters (scales) used in multiscale retinex

func (e *Retinex) GetScales() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("scales")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Retinex) SetScales(value int) error {
	return e.Element.SetProperty("scales", value)
}

// sigma (float64)
//
// Sigma

func (e *Retinex) GetSigma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sigma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Retinex) SetSigma(value float64) error {
	return e.Element.SetProperty("sigma", value)
}

// ----- Constants -----

type GstRetinexMethod string

const (
	GstRetinexMethodBasic GstRetinexMethod = "basic" // basic (0) – Basic retinex restoration
	GstRetinexMethodMultiscale GstRetinexMethod = "multiscale" // multiscale (1) – Mutiscale retinex restoration
)

