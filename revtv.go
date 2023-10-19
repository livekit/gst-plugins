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

type Revtv struct {
	*base.GstBaseTransform
}

func NewRevtv() (*Revtv, error) {
	e, err := gst.NewElement("revtv")
	if err != nil {
		return nil, err
	}
	return &Revtv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRevtvWithName(name string) (*Revtv, error) {
	e, err := gst.NewElementWithName("revtv", name)
	if err != nil {
		return nil, err
	}
	return &Revtv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// delay (int)
//
// Delay in frames between updates

func (e *Revtv) GetDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Revtv) SetDelay(value int) error {
	return e.Element.SetProperty("delay", value)
}

// gain (int)
//
// Control gain

func (e *Revtv) GetGain() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Revtv) SetGain(value int) error {
	return e.Element.SetProperty("gain", value)
}

// linespace (int)
//
// Control line spacing

func (e *Revtv) GetLinespace() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("linespace")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Revtv) SetLinespace(value int) error {
	return e.Element.SetProperty("linespace", value)
}

