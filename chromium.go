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

type Chromium struct {
	*base.GstBaseTransform
}

func NewChromium() (*Chromium, error) {
	e, err := gst.NewElement("chromium")
	if err != nil {
		return nil, err
	}
	return &Chromium{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewChromiumWithName(name string) (*Chromium, error) {
	e, err := gst.NewElementWithName("chromium", name)
	if err != nil {
		return nil, err
	}
	return &Chromium{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// edge-a (uint)
//
// First edge parameter

func (e *Chromium) GetEdgeA() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("edge-a")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Chromium) SetEdgeA(value uint) error {
	return e.Element.SetProperty("edge-a", value)
}

// edge-b (uint)
//
// Second edge parameter

func (e *Chromium) GetEdgeB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("edge-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Chromium) SetEdgeB(value uint) error {
	return e.Element.SetProperty("edge-b", value)
}

