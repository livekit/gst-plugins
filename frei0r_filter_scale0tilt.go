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

type Frei0RFilterScale0Tilt struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterScale0Tilt() (*Frei0RFilterScale0Tilt, error) {
	e, err := gst.NewElement("frei0r-filter-scale0tilt")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterScale0Tilt{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterScale0TiltWithName(name string) (*Frei0RFilterScale0Tilt, error) {
	e, err := gst.NewElementWithName("frei0r-filter-scale0tilt", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterScale0Tilt{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// clip-bottom (float64)

func (e *Frei0RFilterScale0Tilt) GetClipBottom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetClipBottom(value float64) error {
	return e.Element.SetProperty("clip-bottom", value)
}

// clip-left (float64)

func (e *Frei0RFilterScale0Tilt) GetClipLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetClipLeft(value float64) error {
	return e.Element.SetProperty("clip-left", value)
}

// clip-right (float64)

func (e *Frei0RFilterScale0Tilt) GetClipRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetClipRight(value float64) error {
	return e.Element.SetProperty("clip-right", value)
}

// clip-top (float64)

func (e *Frei0RFilterScale0Tilt) GetClipTop() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("clip-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetClipTop(value float64) error {
	return e.Element.SetProperty("clip-top", value)
}

// scale-x (float64)

func (e *Frei0RFilterScale0Tilt) GetScaleX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetScaleX(value float64) error {
	return e.Element.SetProperty("scale-x", value)
}

// scale-y (float64)

func (e *Frei0RFilterScale0Tilt) GetScaleY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetScaleY(value float64) error {
	return e.Element.SetProperty("scale-y", value)
}

// tilt-x (float64)

func (e *Frei0RFilterScale0Tilt) GetTiltX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetTiltX(value float64) error {
	return e.Element.SetProperty("tilt-x", value)
}

// tilt-y (float64)

func (e *Frei0RFilterScale0Tilt) GetTiltY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("tilt-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterScale0Tilt) SetTiltY(value float64) error {
	return e.Element.SetProperty("tilt-y", value)
}

