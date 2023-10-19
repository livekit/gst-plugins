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

type Frei0RMixerCairoblend struct {
	Element *gst.Element
}

func NewFrei0RMixerCairoblend() (*Frei0RMixerCairoblend, error) {
	e, err := gst.NewElement("frei0r-mixer-cairoblend")
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerCairoblend{Element: e}, nil
}

func NewFrei0RMixerCairoblendWithName(name string) (*Frei0RMixerCairoblend, error) {
	e, err := gst.NewElementWithName("frei0r-mixer-cairoblend", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RMixerCairoblend{Element: e}, nil
}

// ----- Properties -----

// blend-mode (string)
//
// Blend mode used to compose image. Accepted values: 'normal', 'add', 'saturate', 'multiply', 'screen', 'overlay', 'darken', 'lighten', 'colordodge', 'colorburn', 'hardlight', 'softlight', 'difference', 'exclusion', 'hslhue', 'hslsaturation', 'hslcolor', 'hslluminosity'

func (e *Frei0RMixerCairoblend) GetBlendMode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("blend-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RMixerCairoblend) SetBlendMode(value string) error {
	return e.Element.SetProperty("blend-mode", value)
}

// opacity (float64)
//
// Opacity of composited image

func (e *Frei0RMixerCairoblend) GetOpacity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RMixerCairoblend) SetOpacity(value float64) error {
	return e.Element.SetProperty("opacity", value)
}

