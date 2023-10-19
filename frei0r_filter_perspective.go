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

type Frei0RFilterPerspective struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterPerspective() (*Frei0RFilterPerspective, error) {
	e, err := gst.NewElement("frei0r-filter-perspective")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPerspective{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterPerspectiveWithName(name string) (*Frei0RFilterPerspective, error) {
	e, err := gst.NewElementWithName("frei0r-filter-perspective", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPerspective{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bottom-left-Y (float64)

func (e *Frei0RFilterPerspective) GetBottomLeftY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-left-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetBottomLeftY(value float64) error {
	return e.Element.SetProperty("bottom-left-Y", value)
}

// bottom-left-x (float64)

func (e *Frei0RFilterPerspective) GetBottomLeftX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-left-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetBottomLeftX(value float64) error {
	return e.Element.SetProperty("bottom-left-x", value)
}

// bottom-right-Y (float64)

func (e *Frei0RFilterPerspective) GetBottomRightY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-right-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetBottomRightY(value float64) error {
	return e.Element.SetProperty("bottom-right-Y", value)
}

// bottom-right-x (float64)

func (e *Frei0RFilterPerspective) GetBottomRightX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bottom-right-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetBottomRightX(value float64) error {
	return e.Element.SetProperty("bottom-right-x", value)
}

// top-left-Y (float64)

func (e *Frei0RFilterPerspective) GetTopLeftY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-left-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetTopLeftY(value float64) error {
	return e.Element.SetProperty("top-left-Y", value)
}

// top-left-x (float64)

func (e *Frei0RFilterPerspective) GetTopLeftX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-left-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetTopLeftX(value float64) error {
	return e.Element.SetProperty("top-left-x", value)
}

// top-right-Y (float64)

func (e *Frei0RFilterPerspective) GetTopRightY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-right-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetTopRightY(value float64) error {
	return e.Element.SetProperty("top-right-Y", value)
}

// top-right-x (float64)

func (e *Frei0RFilterPerspective) GetTopRightX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("top-right-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterPerspective) SetTopRightX(value float64) error {
	return e.Element.SetProperty("top-right-x", value)
}

