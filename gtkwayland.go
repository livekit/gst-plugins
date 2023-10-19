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

type Gtkwayland struct {
	*base.GstBaseSink
}

func NewGtkwayland() (*Gtkwayland, error) {
	e, err := gst.NewElement("gtkwayland")
	if err != nil {
		return nil, err
	}
	return &Gtkwayland{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGtkwaylandWithName(name string) (*Gtkwayland, error) {
	e, err := gst.NewElementWithName("gtkwayland", name)
	if err != nil {
		return nil, err
	}
	return &Gtkwayland{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// drm-device (string)

func (e *Gtkwayland) GetDrmDevice() (string, error) {
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

func (e *Gtkwayland) SetDrmDevice(value string) error {
	return e.Element.SetProperty("drm-device", value)
}

// rotate-method (GstVideoOrientationMethod)
//
// rotate method

func (e *Gtkwayland) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

func (e *Gtkwayland) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

// widget (GstGtkWidget)
//
// The GtkWidget to place in the widget hierarchy (must only be get from the GTK main thread)

func (e *Gtkwayland) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}

