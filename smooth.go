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

type Smooth struct {
	*base.GstBaseTransform
}

func NewSmooth() (*Smooth, error) {
	e, err := gst.NewElement("smooth")
	if err != nil {
		return nil, err
	}
	return &Smooth{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSmoothWithName(name string) (*Smooth, error) {
	e, err := gst.NewElementWithName("smooth", name)
	if err != nil {
		return nil, err
	}
	return &Smooth{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// active (bool)
//
// process video

func (e *Smooth) GetActive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("active")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Smooth) SetActive(value bool) error {
	return e.Element.SetProperty("active", value)
}

// filter-size (int)
//
// size of media filter

func (e *Smooth) GetFilterSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("filter-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Smooth) SetFilterSize(value int) error {
	return e.Element.SetProperty("filter-size", value)
}

// luma-only (bool)
//
// only filter luma part

func (e *Smooth) GetLumaOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("luma-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Smooth) SetLumaOnly(value bool) error {
	return e.Element.SetProperty("luma-only", value)
}

// tolerance (int)
//
// contrast tolerance for smoothing

func (e *Smooth) GetTolerance() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Smooth) SetTolerance(value int) error {
	return e.Element.SetProperty("tolerance", value)
}

