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

type Frei0RSrcTestPatC struct {
	*base.GstPushSrc
}

func NewFrei0RSrcTestPatC() (*Frei0RSrcTestPatC, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-c")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatC{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcTestPatCWithName(name string) (*Frei0RSrcTestPatC, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-c", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatC{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// color-space (float64)

func (e *Frei0RSrcTestPatC) GetColorSpace() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color-space")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatC) SetColorSpace(value float64) error {
	return e.Element.SetProperty("color-space", value)
}

// cross-section (float64)

func (e *Frei0RSrcTestPatC) GetCrossSection() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("cross-section")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatC) SetCrossSection(value float64) error {
	return e.Element.SetProperty("cross-section", value)
}

// fullscreen (bool)

func (e *Frei0RSrcTestPatC) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatC) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

// third-axis-value (float64)

func (e *Frei0RSrcTestPatC) GetThirdAxisValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("third-axis-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatC) SetThirdAxisValue(value float64) error {
	return e.Element.SetProperty("third-axis-value", value)
}

