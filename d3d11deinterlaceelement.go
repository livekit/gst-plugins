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

type D3D11Deinterlaceelement struct {
	*base.GstBaseTransform
}

func NewD3D11Deinterlaceelement() (*D3D11Deinterlaceelement, error) {
	e, err := gst.NewElement("d3d11deinterlaceelement")
	if err != nil {
		return nil, err
	}
	return &D3D11Deinterlaceelement{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewD3D11DeinterlaceelementWithName(name string) (*D3D11Deinterlaceelement, error) {
	e, err := gst.NewElementWithName("d3d11deinterlaceelement", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Deinterlaceelement{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// adapter (uint)
//
// DXGI Adapter index for creating device

func (e *D3D11Deinterlaceelement) GetAdapter() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// device-id (uint)
//
// DXGI Device ID

func (e *D3D11Deinterlaceelement) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// method (GstD3D11DeinterlaceMethod)
//
// Deinterlace Method. Use can set multiple methods as a flagset and element will select one of method automatically. If deinterlacing device failed to deinterlace with given mode, fallback might happen by the device

func (e *D3D11Deinterlaceelement) GetMethod() (GstD3D11DeinterlaceMethod, error) {
	var value GstD3D11DeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11DeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11DeinterlaceMethod")
	}
	return value, nil
}

func (e *D3D11Deinterlaceelement) SetMethod(value GstD3D11DeinterlaceMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// supported-methods (GstD3D11DeinterlaceMethod)
//
// Set of supported deinterlace methods by device

func (e *D3D11Deinterlaceelement) GetSupportedMethods() (GstD3D11DeinterlaceMethod, error) {
	var value GstD3D11DeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("supported-methods")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11DeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11DeinterlaceMethod")
	}
	return value, nil
}

// vendor-id (uint)
//
// DXGI Vendor ID

func (e *D3D11Deinterlaceelement) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// ----- Constants -----

type GstD3D11DeinterlaceMethod string

const (
	GstD3D11DeinterlaceMethodBlend GstD3D11DeinterlaceMethod = "blend" // blend (0x00000001) – Blend: Blending top/bottom field pictures into one frame. Framerate will be preserved (e.g., 60i -> 30p)
	GstD3D11DeinterlaceMethodBob GstD3D11DeinterlaceMethod = "bob" // bob (0x00000002) – Bob: Interpolating missing lines by using the adjacent lines. Framerate will be doubled (e,g, 60i -> 60p)
	GstD3D11DeinterlaceMethodAdaptive GstD3D11DeinterlaceMethod = "adaptive" // adaptive (0x00000004) – Adaptive: Interpolating missing lines by using spatial/temporal references. Framerate will be doubled (e,g, 60i -> 60p)
	GstD3D11DeinterlaceMethodMocomp GstD3D11DeinterlaceMethod = "mocomp" // mocomp (0x00000008) – Motion Compensation: Recreating missing lines by using motion vector. Framerate will be doubled (e,g, 60i -> 60p)
)

