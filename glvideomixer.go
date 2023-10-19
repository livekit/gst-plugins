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
)

type Glvideomixer struct {
	Element *gst.Element
}

func NewGlvideomixer() (*Glvideomixer, error) {
	e, err := gst.NewElement("glvideomixer")
	if err != nil {
		return nil, err
	}
	return &Glvideomixer{Element: e}, nil
}

func NewGlvideomixerWithName(name string) (*Glvideomixer, error) {
	e, err := gst.NewElementWithName("glvideomixer", name)
	if err != nil {
		return nil, err
	}
	return &Glvideomixer{Element: e}, nil
}

// ----- Properties -----

// background (GstGLVideoMixerBackground)
//
// Background type

func (e *Glvideomixer) GetBackground() (interface{}, error) {
	return e.Element.GetProperty("background")
}

func (e *Glvideomixer) SetBackground(value interface{}) error {
	return e.Element.SetProperty("background", value)
}

