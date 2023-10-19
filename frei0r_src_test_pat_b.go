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

type Frei0RSrcTestPatB struct {
	*base.GstPushSrc
}

func NewFrei0RSrcTestPatB() (*Frei0RSrcTestPatB, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-b")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatB{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcTestPatBWithName(name string) (*Frei0RSrcTestPatB, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-b", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatB{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// aspect-type (float64)
//
// 7 choices, pixel aspect ratio

func (e *Frei0RSrcTestPatB) GetAspectType() (float64, error) {
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

func (e *Frei0RSrcTestPatB) SetAspectType(value float64) error {
	return e.Element.SetProperty("aspect-type", value)
}

// manual-aspect (float64)
//
// Manual pixel aspect ratio (Aspect type 6)

func (e *Frei0RSrcTestPatB) GetManualAspect() (float64, error) {
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

func (e *Frei0RSrcTestPatB) SetManualAspect(value float64) error {
	return e.Element.SetProperty("manual-aspect", value)
}

// type (float64)
//
// 8 choices, select test pattern

func (e *Frei0RSrcTestPatB) GetType() (float64, error) {
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

func (e *Frei0RSrcTestPatB) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

