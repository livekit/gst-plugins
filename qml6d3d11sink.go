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

type Qml6D3D11Sink struct {
	*base.GstBaseSink
}

func NewQml6D3D11Sink() (*Qml6D3D11Sink, error) {
	e, err := gst.NewElement("qml6d3d11sink")
	if err != nil {
		return nil, err
	}
	return &Qml6D3D11Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewQml6D3D11SinkWithName(name string) (*Qml6D3D11Sink, error) {
	e, err := gst.NewElementWithName("qml6d3d11sink", name)
	if err != nil {
		return nil, err
	}
	return &Qml6D3D11Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Qml6D3D11Sink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qml6D3D11Sink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// widget (GstGpointer)
//
// The QQuickItem to place in the object hierarchy

func (e *Qml6D3D11Sink) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}

func (e *Qml6D3D11Sink) SetWidget(value interface{}) error {
	return e.Element.SetProperty("widget", value)
}

