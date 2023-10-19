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
)

type Rndbuffersize struct {
	Element *gst.Element
}

func NewRndbuffersize() (*Rndbuffersize, error) {
	e, err := gst.NewElement("rndbuffersize")
	if err != nil {
		return nil, err
	}
	return &Rndbuffersize{Element: e}, nil
}

func NewRndbuffersizeWithName(name string) (*Rndbuffersize, error) {
	e, err := gst.NewElementWithName("rndbuffersize", name)
	if err != nil {
		return nil, err
	}
	return &Rndbuffersize{Element: e}, nil
}

// ----- Properties -----

// max (int)
//
// maximum buffer size

func (e *Rndbuffersize) GetMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rndbuffersize) SetMax(value int) error {
	return e.Element.SetProperty("max", value)
}

// min (int)
//
// minimum buffer size

func (e *Rndbuffersize) GetMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rndbuffersize) SetMin(value int) error {
	return e.Element.SetProperty("min", value)
}

// seed (uint)
//
// seed for randomness (initialized when going from READY to PAUSED)

func (e *Rndbuffersize) GetSeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rndbuffersize) SetSeed(value uint) error {
	return e.Element.SetProperty("seed", value)
}

