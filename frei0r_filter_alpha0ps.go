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

type Frei0RFilterAlpha0Ps struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterAlpha0Ps() (*Frei0RFilterAlpha0Ps, error) {
	e, err := gst.NewElement("frei0r-filter-alpha0ps")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterAlpha0Ps{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterAlpha0PsWithName(name string) (*Frei0RFilterAlpha0Ps, error) {
	e, err := gst.NewElementWithName("frei0r-filter-alpha0ps", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterAlpha0Ps{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// display (float64)

func (e *Frei0RFilterAlpha0Ps) GetDisplay() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlpha0Ps) SetDisplay(value float64) error {
	return e.Element.SetProperty("display", value)
}

// display-input-alpha (bool)

func (e *Frei0RFilterAlpha0Ps) GetDisplayInputAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display-input-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterAlpha0Ps) SetDisplayInputAlpha(value bool) error {
	return e.Element.SetProperty("display-input-alpha", value)
}

// invert (bool)

func (e *Frei0RFilterAlpha0Ps) GetInvert() (bool, error) {
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

func (e *Frei0RFilterAlpha0Ps) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

// operation (float64)

func (e *Frei0RFilterAlpha0Ps) GetOperation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("operation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlpha0Ps) SetOperation(value float64) error {
	return e.Element.SetProperty("operation", value)
}

// shrink-grow-blur-amount (float64)

func (e *Frei0RFilterAlpha0Ps) GetShrinkGrowBlurAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shrink-grow-blur-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlpha0Ps) SetShrinkGrowBlurAmount(value float64) error {
	return e.Element.SetProperty("shrink-grow-blur-amount", value)
}

// threshold (float64)

func (e *Frei0RFilterAlpha0Ps) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterAlpha0Ps) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

