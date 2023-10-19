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

type Agingtv struct {
	*base.GstBaseTransform
}

func NewAgingtv() (*Agingtv, error) {
	e, err := gst.NewElement("agingtv")
	if err != nil {
		return nil, err
	}
	return &Agingtv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAgingtvWithName(name string) (*Agingtv, error) {
	e, err := gst.NewElementWithName("agingtv", name)
	if err != nil {
		return nil, err
	}
	return &Agingtv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// color-aging (bool)
//
// Color Aging

func (e *Agingtv) GetColorAging() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("color-aging")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Agingtv) SetColorAging(value bool) error {
	return e.Element.SetProperty("color-aging", value)
}

// dusts (bool)
//
// Dusts

func (e *Agingtv) GetDusts() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dusts")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Agingtv) SetDusts(value bool) error {
	return e.Element.SetProperty("dusts", value)
}

// pits (bool)
//
// Pits

func (e *Agingtv) GetPits() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pits")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Agingtv) SetPits(value bool) error {
	return e.Element.SetProperty("pits", value)
}

// scratch-lines (uint)
//
// Number of scratch lines

func (e *Agingtv) GetScratchLines() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("scratch-lines")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Agingtv) SetScratchLines(value uint) error {
	return e.Element.SetProperty("scratch-lines", value)
}

