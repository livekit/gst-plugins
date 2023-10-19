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

type Cvsobel struct {
	*base.GstBaseTransform
}

func NewCvsobel() (*Cvsobel, error) {
	e, err := gst.NewElement("cvsobel")
	if err != nil {
		return nil, err
	}
	return &Cvsobel{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCvsobelWithName(name string) (*Cvsobel, error) {
	e, err := gst.NewElementWithName("cvsobel", name)
	if err != nil {
		return nil, err
	}
	return &Cvsobel{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aperture-size (int)
//
// Size of the extended Sobel Kernel (1, 3, 5 or 7)

func (e *Cvsobel) GetApertureSize() (int, error) {
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

func (e *Cvsobel) SetApertureSize(value int) error {
	return e.Element.SetProperty("aperture-size", value)
}

// mask (bool)
//
// Sets whether the detected derivative edges should be used as a mask on the original input or not

func (e *Cvsobel) GetMask() (bool, error) {
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

func (e *Cvsobel) SetMask(value bool) error {
	return e.Element.SetProperty("mask", value)
}

// x-order (int)
//
// Order of the derivative x

func (e *Cvsobel) GetXOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("x-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsobel) SetXOrder(value int) error {
	return e.Element.SetProperty("x-order", value)
}

// y-order (int)
//
// Order of the derivative y

func (e *Cvsobel) GetYOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("y-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsobel) SetYOrder(value int) error {
	return e.Element.SetProperty("y-order", value)
}

