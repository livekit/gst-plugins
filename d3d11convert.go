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

type D3D11Convert struct {
	*base.GstBaseTransform
}

func NewD3D11Convert() (*D3D11Convert, error) {
	e, err := gst.NewElement("d3d11convert")
	if err != nil {
		return nil, err
	}
	return &D3D11Convert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewD3D11ConvertWithName(name string) (*D3D11Convert, error) {
	e, err := gst.NewElementWithName("d3d11convert", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Convert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// add-borders (bool)
//
// Add borders if necessary to keep the display aspect ratio

func (e *D3D11Convert) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Convert) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

// border-color (uint64)
//
// Border color to use in ARGB64 format

func (e *D3D11Convert) GetBorderColor() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("border-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *D3D11Convert) SetBorderColor(value uint64) error {
	return e.Element.SetProperty("border-color", value)
}

// dest-alpha-mode (GstD3D11AlphaMode)
//
// Alpha mode to be applied to output stream. In case of "premultiplied",
// straight to premultiplied alpha conversion will be performed

func (e *D3D11Convert) GetDestAlphaMode() (interface{}, error) {
	return e.Element.GetProperty("dest-alpha-mode")
}

func (e *D3D11Convert) SetDestAlphaMode(value interface{}) error {
	return e.Element.SetProperty("dest-alpha-mode", value)
}

// gamma-mode (GstVideoGammaMode)
//
// Gamma conversion mode

func (e *D3D11Convert) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

func (e *D3D11Convert) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

// primaries-mode (GstVideoPrimariesMode)
//
// Primaries conversion mode

func (e *D3D11Convert) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

func (e *D3D11Convert) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

// src-alpha-mode (GstD3D11AlphaMode)
//
// Input stream's applied alpha mode. In case of "premultiplied",
// premultiplied to straight alpha conversion will be performed

func (e *D3D11Convert) GetSrcAlphaMode() (interface{}, error) {
	return e.Element.GetProperty("src-alpha-mode")
}

func (e *D3D11Convert) SetSrcAlphaMode(value interface{}) error {
	return e.Element.SetProperty("src-alpha-mode", value)
}

