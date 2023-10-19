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

type Frei0RMixerBlend struct {
	Element *gst.Element
}

func NewFrei0RMixerBlend() (*Frei0RMixerBlend, error) {
	e, err := gst.NewElement("frei0r-mixer-blend")
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerBlend{Element: e}, nil
}

func NewFrei0RMixerBlendWithName(name string) (*Frei0RMixerBlend, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-blend", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerBlend{Element: e}, nil
}

// ----- Properties -----

// blend (float64)
//
// blend factor

func (e *Frei0RMixerBlend) GetBlend() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("blend")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerBlend) SetBlend(value float64) error {
	return e.Element.SetProperty("blend", value)
}

