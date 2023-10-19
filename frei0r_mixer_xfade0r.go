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

type Frei0RMixerXfade0R struct {
	Element *gst.Element
}

func NewFrei0RMixerXfade0R() (*Frei0RMixerXfade0R, error) {
	e, err := gst.NewElement("frei0r-mixer-xfade0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerXfade0R{Element: e}, nil
}

func NewFrei0RMixerXfade0RWithName(name string) (*Frei0RMixerXfade0R, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-xfade0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerXfade0R{Element: e}, nil
}

// ----- Properties -----

// fader (float64)
//
// the fader position

func (e *Frei0RMixerXfade0R) GetFader() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("fader")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerXfade0R) SetFader(value float64) error {
	return e.Element.SetProperty("fader", value)
}
