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

type D3D11Scale struct {
	*base.GstBaseTransform
}

func NewD3D11Scale() (*D3D11Scale, error) {
	e, err := gst.NewElement("d3d11scale")
	if err != nil {
		return nil, err
	}
	return &D3D11Scale{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewD3D11ScaleWithName(name string) (*D3D11Scale, error) {
	e, err := gst.NewElementWithName("d3d11scale", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Scale{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// add-borders (bool)
//
// Add borders if necessary to keep the display aspect ratio

func (e *D3D11Scale) GetAddBorders() (bool, error) {
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

func (e *D3D11Scale) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

// border-color (uint64)
//
// Border color to use in ARGB64 format

func (e *D3D11Scale) GetBorderColor() (uint64, error) {
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

func (e *D3D11Scale) SetBorderColor(value uint64) error {
	return e.Element.SetProperty("border-color", value)
}

