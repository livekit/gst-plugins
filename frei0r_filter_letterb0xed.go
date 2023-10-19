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

type Frei0RFilterLetterb0Xed struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterLetterb0Xed() (*Frei0RFilterLetterb0Xed, error) {
	e, err := gst.NewElement("frei0r-filter-letterb0xed")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLetterb0Xed{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterLetterb0XedWithName(name string) (*Frei0RFilterLetterb0Xed, error) {
	e, err := gst.NewElementWithName("frei0r-filter-letterb0xed", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterLetterb0Xed{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// border-width (float64)

func (e *Frei0RFilterLetterb0Xed) GetBorderWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("border-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterLetterb0Xed) SetBorderWidth(value float64) error {
	return e.Element.SetProperty("border-width", value)
}

// transparency (bool)

func (e *Frei0RFilterLetterb0Xed) GetTransparency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("transparency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterLetterb0Xed) SetTransparency(value bool) error {
	return e.Element.SetProperty("transparency", value)
}

