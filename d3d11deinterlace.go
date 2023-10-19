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

type D3D11Deinterlace struct {
	Element *gst.Element
}

func NewD3D11Deinterlace() (*D3D11Deinterlace, error) {
	e, err := gst.NewElement("d3d11deinterlace")
	if err != nil {
		return nil, err
	}
	return &D3D11Deinterlace{Element: e}, nil
}

func NewD3D11DeinterlaceWithName(name string) (*D3D11Deinterlace, error) {
	e, err := gst.NewElementWithName("d3d11deinterlace", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Deinterlace{Element: e}, nil
}

// ----- Properties -----

// adapter (uint)
//
// DXGI Adapter index for creating device

func (e *D3D11Deinterlace) GetAdapter() (uint, error) {
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

func (e *D3D11Deinterlace) GetDeviceId() (uint, error) {
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

func (e *D3D11Deinterlace) GetMethod() (interface{}, error) {
	return e.Element.GetProperty("method")
}

func (e *D3D11Deinterlace) SetMethod(value interface{}) error {
	return e.Element.SetProperty("method", value)
}

// qos (bool)
//
// Handle Quality-of-Service events

func (e *D3D11Deinterlace) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3D11Deinterlace) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// supported-methods (GstD3D11DeinterlaceMethod)
//
// Set of supported deinterlace methods by device

func (e *D3D11Deinterlace) GetSupportedMethods() (interface{}, error) {
	return e.Element.GetProperty("supported-methods")
}

// vendor-id (uint)
//
// DXGI Vendor ID

func (e *D3D11Deinterlace) GetVendorId() (uint, error) {
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

