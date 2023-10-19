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

type Dewarp struct {
	*base.GstBaseTransform
}

func NewDewarp() (*Dewarp, error) {
	e, err := gst.NewElement("dewarp")
	if err != nil {
		return nil, err
	}
	return &Dewarp{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDewarpWithName(name string) (*Dewarp, error) {
	e, err := gst.NewElementWithName("dewarp", name)
	if err != nil {
		return nil, err
	}
	return &Dewarp{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// display-mode (GstDewarpDisplayMode)
//
// How to display the dewarped image

func (e *Dewarp) GetDisplayMode() (GstDewarpDisplayMode, error) {
	var value GstDewarpDisplayMode
	var ok bool
	v, err := e.Element.GetProperty("display-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDewarpDisplayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDewarpDisplayMode")
	}
	return value, nil
}

func (e *Dewarp) SetDisplayMode(value GstDewarpDisplayMode) error {
	e.Element.SetArg("display-mode", string(value))
	return nil
}

// inner-radius (float64)
//
// Inner radius of the fisheye image donut. If outer radius <= inner radius the element will work in passthrough mode

func (e *Dewarp) GetInnerRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("inner-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dewarp) SetInnerRadius(value float64) error {
	return e.Element.SetProperty("inner-radius", value)
}

// interpolation-method (GstDewarpInterpolationMode)
//
// Interpolation method to use

func (e *Dewarp) GetInterpolationMethod() (GstDewarpInterpolationMode, error) {
	var value GstDewarpInterpolationMode
	var ok bool
	v, err := e.Element.GetProperty("interpolation-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDewarpInterpolationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDewarpInterpolationMode")
	}
	return value, nil
}

func (e *Dewarp) SetInterpolationMethod(value GstDewarpInterpolationMode) error {
	e.Element.SetArg("interpolation-method", string(value))
	return nil
}

// outer-radius (float64)
//
// Outer radius of the fisheye image donut. If outer radius <= inner radius the element will work in passthrough mode

func (e *Dewarp) GetOuterRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("outer-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dewarp) SetOuterRadius(value float64) error {
	return e.Element.SetProperty("outer-radius", value)
}

// x-center (float64)
//
// X axis center of the fisheye image

func (e *Dewarp) GetXCenter() (float64, error) {
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

func (e *Dewarp) SetXCenter(value float64) error {
	return e.Element.SetProperty("x-center", value)
}

// x-remap-correction (float64)
//
// Correction factor for remapping on x axis. A correction is needed if the fisheye image is not inside a circle

func (e *Dewarp) GetXRemapCorrection() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-remap-correction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dewarp) SetXRemapCorrection(value float64) error {
	return e.Element.SetProperty("x-remap-correction", value)
}

// y-center (float64)
//
// Y axis center of the fisheye image

func (e *Dewarp) GetYCenter() (float64, error) {
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

func (e *Dewarp) SetYCenter(value float64) error {
	return e.Element.SetProperty("y-center", value)
}

// y-remap-correction (float64)
//
// Correction factor for remapping on y axis. A correction is needed if the fisheye image is not inside a circle

func (e *Dewarp) GetYRemapCorrection() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-remap-correction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Dewarp) SetYRemapCorrection(value float64) error {
	return e.Element.SetProperty("y-remap-correction", value)
}

// ----- Constants -----

type GstDewarpInterpolationMode string

const (
	GstDewarpInterpolationModeNearest GstDewarpInterpolationMode = "nearest" // nearest (0) – A nearest-neighbor interpolation
	GstDewarpInterpolationModeBilinear GstDewarpInterpolationMode = "bilinear" // bilinear (1) – A bilinear interpolation
	GstDewarpInterpolationModeBicubic GstDewarpInterpolationMode = "bicubic" // bicubic (2) – A bicubic interpolation over 4x4 pixel neighborhood
	GstDewarpInterpolationModeLanczos GstDewarpInterpolationMode = "Lanczos" // Lanczos (3) – A Lanczos interpolation over 8x8 pixel neighborhood
)

type GstDewarpDisplayMode string

const (
	GstDewarpDisplayModeSinglePanorama GstDewarpDisplayMode = "single-panorama" // single-panorama (0) – Single panorama image
	GstDewarpDisplayModeDoublePanorama GstDewarpDisplayMode = "double-panorama" // double-panorama (1) – Dewarped image is split in two images displayed one below the other
	GstDewarpDisplayModeQuadView GstDewarpDisplayMode = "quad-view" // quad-view (2) – Dewarped image is split in four images dysplayed as a quad view
)

