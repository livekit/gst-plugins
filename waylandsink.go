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

type Waylandsink struct {
	*base.GstBaseSink
}

func NewWaylandsink() (*Waylandsink, error) {
	e, err := gst.NewElement("waylandsink")
	if err != nil {
		return nil, err
	}
	return &Waylandsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewWaylandsinkWithName(name string) (*Waylandsink, error) {
	e, err := gst.NewElementWithName("waylandsink", name)
	if err != nil {
		return nil, err
	}
	return &Waylandsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// display (string)
//
// Wayland display name to connect to, if not supplied via the GstContext

func (e *Waylandsink) GetDisplay() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Waylandsink) SetDisplay(value string) error {
	return e.Element.SetProperty("display", value)
}

// drm-device (string)

func (e *Waylandsink) GetDrmDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("drm-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Waylandsink) SetDrmDevice(value string) error {
	return e.Element.SetProperty("drm-device", value)
}

// fullscreen (bool)
//
// Whether the surface should be made fullscreen

func (e *Waylandsink) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Waylandsink) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

// render-rectangle (GstValueArray)
//
// This helper installs the "render-rectangle" property into the
// class.

func (e *Waylandsink) GetRenderRectangle() (interface{}, error) {
	return e.Element.GetProperty("render-rectangle")
}

func (e *Waylandsink) SetRenderRectangle(value interface{}) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// rotate-method (GstVideoOrientationMethod)

func (e *Waylandsink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

func (e *Waylandsink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

