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

type Cvsmooth struct {
	*base.GstBaseTransform
}

func NewCvsmooth() (*Cvsmooth, error) {
	e, err := gst.NewElement("cvsmooth")
	if err != nil {
		return nil, err
	}
	return &Cvsmooth{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCvsmoothWithName(name string) (*Cvsmooth, error) {
	e, err := gst.NewElementWithName("cvsmooth", name)
	if err != nil {
		return nil, err
	}
	return &Cvsmooth{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color (float64)
//
// If type is gaussian, this means the standard deviation.If type is bilateral, this means the color-sigma. If zero, Default values are used.

func (e *Cvsmooth) GetColor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Cvsmooth) SetColor(value float64) error {
	return e.Element.SetProperty("color", value)
}

// height (int)
//
// Height of the area to blur (in pixels).

func (e *Cvsmooth) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsmooth) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// kernel-height (int)
//
// The gaussian kernel height (must be positive and odd).

func (e *Cvsmooth) GetKernelHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kernel-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsmooth) SetKernelHeight(value int) error {
	return e.Element.SetProperty("kernel-height", value)
}

// kernel-width (int)
//
// The gaussian kernel width (must be positive and odd).If type is median, this means the aperture linear size.Check OpenCV docs: http://docs.opencv.org/2.4/modules/imgproc/doc/filtering.htm

func (e *Cvsmooth) GetKernelWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kernel-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsmooth) SetKernelWidth(value int) error {
	return e.Element.SetProperty("kernel-width", value)
}

// position-x (int)
//
// Starting x position for blur (in pixels).

func (e *Cvsmooth) GetPositionX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("position-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsmooth) SetPositionX(value int) error {
	return e.Element.SetProperty("position-x", value)
}

// position-y (int)
//
// Starting y position for blur (in pixels).

func (e *Cvsmooth) GetPositionY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("position-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsmooth) SetPositionY(value int) error {
	return e.Element.SetProperty("position-y", value)
}

// spatial (float64)
//
// Only used in bilateral type, means the spatial-sigma.

func (e *Cvsmooth) GetSpatial() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spatial")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Cvsmooth) SetSpatial(value float64) error {
	return e.Element.SetProperty("spatial", value)
}

// type (GstCvSmoothTypeType)
//
// Smooth Type

func (e *Cvsmooth) GetType() (GstCvSmoothTypeType, error) {
	var value GstCvSmoothTypeType
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCvSmoothTypeType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCvSmoothTypeType")
	}
	return value, nil
}

func (e *Cvsmooth) SetType(value GstCvSmoothTypeType) error {
	e.Element.SetArg("type", string(value))
	return nil
}

// width (int)
//
// Width of the area to blur (in pixels).

func (e *Cvsmooth) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cvsmooth) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

// ----- Constants -----

type GstCvSmoothTypeType string

const (
	GstCvSmoothTypeTypeBlur GstCvSmoothTypeType = "blur" // blur (1) – CV Blur
	GstCvSmoothTypeTypeGaussian GstCvSmoothTypeType = "gaussian" // gaussian (2) – CV Gaussian
	GstCvSmoothTypeTypeMedian GstCvSmoothTypeType = "median" // median (3) – CV Median
	GstCvSmoothTypeTypeBilateral GstCvSmoothTypeType = "bilateral" // bilateral (4) – CV Bilateral
)

