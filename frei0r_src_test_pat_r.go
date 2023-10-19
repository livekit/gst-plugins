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

type Frei0RSrcTestPatR struct {
	*base.GstPushSrc
}

func NewFrei0RSrcTestPatR() (*Frei0RSrcTestPatR, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-r")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatR{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcTestPatRWithName(name string) (*Frei0RSrcTestPatR, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatR{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// amplitude (float64)
//
// Amplitude (contrast) of the pattern

func (e *Frei0RSrcTestPatR) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatR) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

// aspect-type (float64)
//
// Pixel aspect ratio presets

func (e *Frei0RSrcTestPatR) GetAspectType() (float64, error) {
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

func (e *Frei0RSrcTestPatR) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

// channel (float64)
//
// Into which color channel to draw

func (e *Frei0RSrcTestPatR) GetChannel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatR) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

// freq-1 (float64)
//
// Pattern 7 H frequency

func (e *Frei0RSrcTestPatR) GetFreq1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("freq-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatR) SetFreq1(value float64) error {
	return e.Element.SetProperty("freq-1", value)
}

// freq-2 (float64)
//
// Pattern 7 V frequency

func (e *Frei0RSrcTestPatR) GetFreq2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("freq-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatR) SetFreq2(value float64) error {
	return e.Element.SetProperty("freq-2", value)
}

// lin-p-swp (bool)
//
// Use linear period sweep

func (e *Frei0RSrcTestPatR) GetLinPSwp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lin-p-swp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RSrcTestPatR) SetLinPSwp(value bool) error {
	return e.Element.SetProperty("lin-p-swp", value)
}

// manual-aspect (float64)
//
// Manual pixel aspect ratio

func (e *Frei0RSrcTestPatR) GetManualAspect() (float64, error) {
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

func (e *Frei0RSrcTestPatR) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

// type (float64)
//
// Type of test pattern

func (e *Frei0RSrcTestPatR) GetType() (float64, error) {
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

func (e *Frei0RSrcTestPatR) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

