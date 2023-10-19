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

type Scaletempo struct {
	*base.GstBaseTransform
}

func NewScaletempo() (*Scaletempo, error) {
	e, err := gst.NewElement("scaletempo")
	if err != nil {
		return nil, err
	}
	return &Scaletempo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewScaletempoWithName(name string) (*Scaletempo, error) {
	e, err := gst.NewElementWithName("scaletempo", name)
	if err != nil {
		return nil, err
	}
	return &Scaletempo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// overlap (float64)
//
// Percentage of stride to overlap

func (e *Scaletempo) GetOverlap() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("overlap")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Scaletempo) SetOverlap(value float64) error {
	return e.Element.SetProperty("overlap", value)
}

// rate (float64)
//
// Current playback rate

func (e *Scaletempo) GetRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// search (uint)
//
// Length in milliseconds to search for best overlap position

func (e *Scaletempo) GetSearch() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("search")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Scaletempo) SetSearch(value uint) error {
	return e.Element.SetProperty("search", value)
}

// stride (uint)
//
// Length in milliseconds to output each stride

func (e *Scaletempo) GetStride() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("stride")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Scaletempo) SetStride(value uint) error {
	return e.Element.SetProperty("stride", value)
}

