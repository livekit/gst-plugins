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

type Gtkglsink struct {
	*base.GstBaseSink
}

func NewGtkglsink() (*Gtkglsink, error) {
	e, err := gst.NewElement("gtkglsink")
	if err != nil {
		return nil, err
	}
	return &Gtkglsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGtkglsinkWithName(name string) (*Gtkglsink, error) {
	e, err := gst.NewElementWithName("gtkglsink", name)
	if err != nil {
		return nil, err
	}
	return &Gtkglsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// rotate-method (GstVideoOrientationMethod)
//
// Rotation method GstVideoOrientationMethod used to render the media

func (e *Gtkglsink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

func (e *Gtkglsink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

