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

type D3D11Compositor struct {
	*base.GstAggregator
}

func NewD3D11Compositor() (*D3D11Compositor, error) {
	e, err := gst.NewElement("d3d11compositor")
	if err != nil {
		return nil, err
	}
	return &D3D11Compositor{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewD3D11CompositorWithName(name string) (*D3D11Compositor, error) {
	e, err := gst.NewElementWithName("d3d11compositor", name)
	if err != nil {
		return nil, err
	}
	return &D3D11Compositor{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// adapter (int)
//
// Adapter index for creating device (-1 for default)

func (e *D3D11Compositor) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *D3D11Compositor) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

// background (GstD3D11CompositorBackground)
//
// Background type

func (e *D3D11Compositor) GetBackground() (GstD3D11CompositorBackground, error) {
	var value GstD3D11CompositorBackground
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11CompositorBackground)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11CompositorBackground")
	}
	return value, nil
}

func (e *D3D11Compositor) SetBackground(value GstD3D11CompositorBackground) error {
	e.Element.SetArg("background", string(value))
	return nil
}

// ignore-inactive-pads (bool)
//
// Don't wait for inactive pads when live. An inactive pad
// is a pad that hasn't yet received a buffer, but that has
// been waited on at least once.

// ----- Constants -----

type GstD3D11CompositorBackground string

const (
	GstD3D11CompositorBackgroundChecker GstD3D11CompositorBackground = "checker" // checker (0) – Checker pattern
	GstD3D11CompositorBackgroundBlack GstD3D11CompositorBackground = "black" // black (1) – Black
	GstD3D11CompositorBackgroundWhite GstD3D11CompositorBackground = "white" // white (2) – White
	GstD3D11CompositorBackgroundTransparent GstD3D11CompositorBackground = "transparent" // transparent (3) – Transparent Background to enable further compositing
)

