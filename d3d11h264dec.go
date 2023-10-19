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
)

type D3D11H264Dec struct {
	Element *gst.Element
}

func NewD3D11H264Dec() (*D3D11H264Dec, error) {
	e, err := gst.NewElement("d3d11h264dec")
	if err != nil {
		return nil, err
	}
	return &D3D11H264Dec{Element: e}, nil
}

func NewD3D11H264DecWithName(name string) (*D3D11H264Dec, error) {
	e, err := gst.NewElementWithName("d3d11h264dec", name)
	if err != nil {
		return nil, err
	}
	return &D3D11H264Dec{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID for this element

func (e *D3D11H264Dec) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// device-id (uint)
//
// DXGI Device ID

func (e *D3D11H264Dec) GetDeviceId() (uint, error) {
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

// vendor-id (uint)
//
// DXGI Vendor ID

func (e *D3D11H264Dec) GetVendorId() (uint, error) {
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
