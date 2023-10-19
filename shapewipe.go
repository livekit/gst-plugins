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

type Shapewipe struct {
	Element *gst.Element
}

func NewShapewipe() (*Shapewipe, error) {
	e, err := gst.NewElement("shapewipe")
	if err != nil {
		return nil, err
	}
	return &Shapewipe{Element: e}, nil
}

func NewShapewipeWithName(name string) (*Shapewipe, error) {
	e, err := gst.NewElementWithName("shapewipe", name)
	if err != nil {
		return nil, err
	}
	return &Shapewipe{Element: e}, nil
}

// ----- Properties -----

// border (float32)
//
// Border of the mask

func (e *Shapewipe) GetBorder() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("border")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Shapewipe) SetBorder(value float32) error {
	return e.Element.SetProperty("border", value)
}

// position (float32)
//
// Position of the mask

func (e *Shapewipe) GetPosition() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Shapewipe) SetPosition(value float32) error {
	return e.Element.SetProperty("position", value)
}

