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

type Chopmydata struct {
	Element *gst.Element
}

func NewChopmydata() (*Chopmydata, error) {
	e, err := gst.NewElement("chopmydata")
	if err != nil {
		return nil, err
	}
	return &Chopmydata{Element: e}, nil
}

func NewChopmydataWithName(name string) (*Chopmydata, error) {
	e, err := gst.NewElementWithName("chopmydata", name)
	if err != nil {
		return nil, err
	}
	return &Chopmydata{Element: e}, nil
}

// ----- Properties -----

// max-size (int)
//
// Maximum size of outgoing buffers

func (e *Chopmydata) GetMaxSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Chopmydata) SetMaxSize(value int) error {
	return e.Element.SetProperty("max-size", value)
}

// min-size (int)
//
// Minimum size of outgoing buffers

func (e *Chopmydata) GetMinSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Chopmydata) SetMinSize(value int) error {
	return e.Element.SetProperty("min-size", value)
}

// step-size (int)
//
// Step increment for random buffer sizes

func (e *Chopmydata) GetStepSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("step-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Chopmydata) SetStepSize(value int) error {
	return e.Element.SetProperty("step-size", value)
}

