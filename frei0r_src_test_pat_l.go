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

type Frei0RSrcTestPatL struct {
	*base.GstPushSrc
}

func NewFrei0RSrcTestPatL() (*Frei0RSrcTestPatL, error) {
	e, err := gst.NewElement("frei0r-src-test-pat-l")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatL{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcTestPatLWithName(name string) (*Frei0RSrcTestPatL, error) {
	e, err := gst.NewElementWithName("frei0r-src-test-pat-l", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcTestPatL{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// channel (float64)
//
// Into which color channel to draw

func (e *Frei0RSrcTestPatL) GetChannel() (float64, error) {
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

func (e *Frei0RSrcTestPatL) SetChannel(value float64) error {
	return e.Element.SetProperty("channel", value)
}

// type (float64)
//
// Type of test pattern

func (e *Frei0RSrcTestPatL) GetType() (float64, error) {
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

func (e *Frei0RSrcTestPatL) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}

