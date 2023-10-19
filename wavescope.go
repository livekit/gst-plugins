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

type Wavescope struct {
	Element *gst.Element
}

func NewWavescope() (*Wavescope, error) {
	e, err := gst.NewElement("wavescope")
	if err != nil {
		return nil, err
	}
	return &Wavescope{Element: e}, nil
}

func NewWavescopeWithName(name string) (*Wavescope, error) {
	e, err := gst.NewElementWithName("wavescope", name)
	if err != nil {
		return nil, err
	}
	return &Wavescope{Element: e}, nil
}

// ----- Properties -----

// style (GstWaveScopeStyle)
//
// Drawing styles for the wave form display.

func (e *Wavescope) GetStyle() (GstWaveScopeStyle, error) {
	var value GstWaveScopeStyle
	var ok bool
	v, err := e.Element.GetProperty("style")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWaveScopeStyle)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWaveScopeStyle")
	}
	return value, nil
}

func (e *Wavescope) SetStyle(value GstWaveScopeStyle) error {
	e.Element.SetArg("style", string(value))
	return nil
}

// ----- Constants -----

type GstWaveScopeStyle string

const (
	GstWaveScopeStyleDots GstWaveScopeStyle = "dots" // dots (0) – draw dots (default)
	GstWaveScopeStyleLines GstWaveScopeStyle = "lines" // lines (1) – draw lines
	GstWaveScopeStyleColorDots GstWaveScopeStyle = "color-dots" // color-dots (2) – draw color dots
	GstWaveScopeStyleColorLines GstWaveScopeStyle = "color-lines" // color-lines (3) – draw color lines
)

