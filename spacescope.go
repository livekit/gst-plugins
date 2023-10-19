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

type Spacescope struct {
	Element *gst.Element
}

func NewSpacescope() (*Spacescope, error) {
	e, err := gst.NewElement("spacescope")
	if err != nil {
		return nil, err
	}
	return &Spacescope{Element: e}, nil
}

func NewSpacescopeWithName(name string) (*Spacescope, error) {
	e, err := gst.NewElementWithName("spacescope", name)
	if err != nil {
		return nil, err
	}
	return &Spacescope{Element: e}, nil
}

// ----- Properties -----

// style (GstSpaceScopeStyle)
//
// Drawing styles for the space scope display.

func (e *Spacescope) GetStyle() (GstSpaceScopeStyle, error) {
	var value GstSpaceScopeStyle
	var ok bool
	v, err := e.Element.GetProperty("style")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSpaceScopeStyle)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSpaceScopeStyle")
	}
	return value, nil
}

func (e *Spacescope) SetStyle(value GstSpaceScopeStyle) error {
	e.Element.SetArg("style", string(value))
	return nil
}

// ----- Constants -----

type GstSpaceScopeStyle string

const (
	GstSpaceScopeStyleDots GstSpaceScopeStyle = "dots" // dots (0) – draw dots (default)
	GstSpaceScopeStyleLines GstSpaceScopeStyle = "lines" // lines (1) – draw lines
	GstSpaceScopeStyleColorDots GstSpaceScopeStyle = "color-dots" // color-dots (2) – draw color dots
	GstSpaceScopeStyleColorLines GstSpaceScopeStyle = "color-lines" // color-lines (3) – draw color lines
)

