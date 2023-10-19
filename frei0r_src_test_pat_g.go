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

type Frei0RSrcTestPatG struct {
	*base.GstPushSrc
}

func NewFrei0RSrcTestPatG() (*Frei0RSrcTestPatG, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-g")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatG{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcTestPatGWithName(name string) (*Frei0RSrcTestPatG, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-g", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatG{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// aspect-type (float64)
//
// Pixel aspect ratio presets

func (e *Frei0RSrcTestPatG) GetAspectType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("aspect-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatG) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

// manual-aspect (float64)
//
// Manual pixel aspect ratio

func (e *Frei0RSrcTestPatG) GetManualAspect() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("manual-aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatG) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

// negative (bool)
//
// Polarity of image

func (e *Frei0RSrcTestPatG) GetNegative() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("negative")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatG) SetNegative(value bool) error {
	return e.Element.SetProperty("negative", value)
}

// size-1 (float64)
//
// Size of major features

func (e *Frei0RSrcTestPatG) GetSize1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatG) SetSize1(value float64) error {
	return e.Element.SetProperty("size-1", value)
}

// size-2 (float64)
//
// Size of minor features

func (e *Frei0RSrcTestPatG) GetSize2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("size-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatG) SetSize2(value float64) error {
	return e.Element.SetProperty("size-2", value)
}

// type (float64)
//
// Type of test pattern

func (e *Frei0RSrcTestPatG) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatG) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

