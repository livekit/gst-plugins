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

type Qmlgloverlay struct {
	*base.GstBaseTransform
}

func NewQmlgloverlay() (*Qmlgloverlay, error) {
	e, err := gst.NewElement("qmlgloverlay")
	if err != nil {
		return nil, err
	}
	return &Qmlgloverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewQmlgloverlayWithName(name string) (*Qmlgloverlay, error) {
	e, err := gst.NewElementWithName("qmlgloverlay", name)
	if err != nil {
		return nil, err
	}
	return &Qmlgloverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// context (GstGLContext)
//
// Get OpenGL context

func (e *Qmlgloverlay) GetContext() (interface{}, error) {
	return e.Element.GetProperty("context")
}

// qml-scene (string)
//
// The contents of the QML scene

func (e *Qmlgloverlay) GetQmlScene() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("qml-scene")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Qmlgloverlay) SetQmlScene(value string) error {
	return e.Element.SetProperty("qml-scene", value)
}

// qos (bool)
//
// Handle Quality-of-Service events

func (e *Qmlgloverlay) GetQos() (bool, error) {
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

func (e *Qmlgloverlay) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// root-item (GstGpointer)
//
// The root QQuickItem from the qml-scene used to render

func (e *Qmlgloverlay) GetRootItem() (interface{}, error) {
	return e.Element.GetProperty("root-item")
}

// widget (GstGpointer)
//
// The QQuickItem to place the input video in the object hierarchy

func (e *Qmlgloverlay) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}

func (e *Qmlgloverlay) SetWidget(value interface{}) error {
	return e.Element.SetProperty("widget", value)
}

