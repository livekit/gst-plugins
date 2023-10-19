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

type Cvlaplace struct {
	*base.GstBaseTransform
}

func NewCvlaplace() (*Cvlaplace, error) {
	e, err := gst.NewElement("cvlaplace")
	if err != nil {
		return nil, err
	}
	return &Cvlaplace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCvlaplaceWithName(name string) (*Cvlaplace, error) {
	e, err := gst.NewElementWithName("cvlaplace", name)
	if err != nil {
		return nil, err
	}
	return &Cvlaplace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aperture-size (int)
//
// Size of the extended Laplace Kernel (1, 3, 5 or 7)

func (e *Cvlaplace) GetApertureSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("aperture-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvlaplace) SetApertureSize(value int) error {
	return e.Element.SetProperty("aperture-size", value)
}

// mask (bool)
//
// Sets whether the detected edges should be used as a mask on the original input or not

func (e *Cvlaplace) GetMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cvlaplace) SetMask(value bool) error {
	return e.Element.SetProperty("mask", value)
}

// scale (float64)
//
// Scale factor

func (e *Cvlaplace) GetScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Cvlaplace) SetScale(value float64) error {
	return e.Element.SetProperty("scale", value)
}

// shift (float64)
//
// Value added to the scaled source array elements

func (e *Cvlaplace) GetShift() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shift")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Cvlaplace) SetShift(value float64) error {
	return e.Element.SetProperty("shift", value)
}

