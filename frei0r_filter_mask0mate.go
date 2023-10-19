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

type Frei0RFilterMask0Mate struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterMask0Mate() (*Frei0RFilterMask0Mate, error) {
	e, err := gst.NewElement("frei0r-filter-mask0mate")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterMask0Mate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterMask0MateWithName(name string) (*Frei0RFilterMask0Mate, error) {
	e, err := gst.NewElementWithName("frei0r-filter-mask0mate", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterMask0Mate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// blur (float64)
//
// Blur the outline of the mask

func (e *Frei0RFilterMask0Mate) GetBlur() (float64, error) {
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

func (e *Frei0RFilterMask0Mate) SetBlur(value float64) error {
	return e.Element.SetProperty("blur", value)
}

// bottom (float64)

func (e *Frei0RFilterMask0Mate) GetBottom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterMask0Mate) SetBottom(value float64) error {
	return e.Element.SetProperty("bottom", value)
}

// invert (bool)
//
// Invert the mask, creates a hole in the frame.

func (e *Frei0RFilterMask0Mate) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterMask0Mate) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

// left (float64)

func (e *Frei0RFilterMask0Mate) GetLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterMask0Mate) SetLeft(value float64) error {
	return e.Element.SetProperty("left", value)
}

// right (float64)

func (e *Frei0RFilterMask0Mate) GetRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterMask0Mate) SetRight(value float64) error {
	return e.Element.SetProperty("right", value)
}

// top (float64)

func (e *Frei0RFilterMask0Mate) GetTop() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterMask0Mate) SetTop(value float64) error {
	return e.Element.SetProperty("top", value)
}

