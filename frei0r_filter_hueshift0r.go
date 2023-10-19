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

type Frei0RFilterHueshift0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterHueshift0R() (*Frei0RFilterHueshift0R, error) {
	e, err := gst.NewElement("frei0r-filter-hueshift0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterHueshift0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterHueshift0RWithName(name string) (*Frei0RFilterHueshift0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-hueshift0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterHueshift0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// hue (float64)
//
// The shift value

func (e *Frei0RFilterHueshift0R) GetHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterHueshift0R) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

