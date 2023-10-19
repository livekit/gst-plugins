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

type Frei0RFilterGlow struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterGlow() (*Frei0RFilterGlow, error) {
	e, err := gst.NewElement("frei0r-filter-glow")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterGlow{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterGlowWithName(name string) (*Frei0RFilterGlow, error) {
	e, err := gst.NewElementWithName("frei0r-filter-glow", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterGlow{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// blur (float64)
//
// Blur of the glow

func (e *Frei0RFilterGlow) GetBlur() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterGlow) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

