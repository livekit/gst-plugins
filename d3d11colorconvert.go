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
	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type D3D11Colorconvert struct {
	*base.GstBaseTransform
}

func NewD3D11Colorconvert() (*D3D11Colorconvert, error) {
	e, err := gst.NewElement("d3d11colorconvert")
	if err != nil {
		return nil, err
	}
	return &D3D11Colorconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewD3D11ColorconvertWithName(name string) (*D3D11Colorconvert, error) {
	e, err := gst.NewElementWithName("d3d11colorconvert", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Colorconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// dest-alpha-mode (GstD3D11AlphaMode)
//
// Alpha mode to be applied to output stream. In case of "premultiplied",
// straight to premultiplied alpha conversion will be performed

func (e *D3D11Colorconvert) GetDestAlphaMode() (interface{}, error) {
	return e.Element.GetProperty("dest-alpha-mode")
}

func (e *D3D11Colorconvert) SetDestAlphaMode(value interface{}) error {
	return e.Element.SetProperty("dest-alpha-mode", value)
}

// gamma-mode (GstVideoGammaMode)
//
// Gamma conversion mode

func (e *D3D11Colorconvert) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

func (e *D3D11Colorconvert) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

// primaries-mode (GstVideoPrimariesMode)
//
// Primaries conversion mode

func (e *D3D11Colorconvert) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

func (e *D3D11Colorconvert) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

// src-alpha-mode (GstD3D11AlphaMode)
//
// Input stream's applied alpha mode. In case of "premultiplied",
// premultiplied to straight alpha conversion will be performed

func (e *D3D11Colorconvert) GetSrcAlphaMode() (interface{}, error) {
	return e.Element.GetProperty("src-alpha-mode")
}

func (e *D3D11Colorconvert) SetSrcAlphaMode(value interface{}) error {
	return e.Element.SetProperty("src-alpha-mode", value)
}

