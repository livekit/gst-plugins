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
	"github.com/tinyzimmer/go-gst/gst"
)

type Autoconvert struct {
	Element *gst.Element
}

func NewAutoconvert() (*Autoconvert, error) {
	e, err := gst.NewElement("autoconvert")
	if err != nil {
		return nil, err
	}
	return &Autoconvert{Element: e}, nil
}

func NewAutoconvertWithName(name string) (*Autoconvert, error) {
	e, err := gst.NewElementWithName("autoconvert", name)
	if err != nil {
		return nil, err
	}
	return &Autoconvert{Element: e}, nil
}

// ----- Properties -----

// factories (GstGpointer)
//
// GList of GstElementFactory objects to pick from (the element takes ownership of the list (NULL means it will go through all possible elements), can only be set once

func (e *Autoconvert) GetFactories() (interface{}, error) {
	return e.Element.GetProperty("factories")
}

func (e *Autoconvert) SetFactories(value interface{}) error {
	return e.Element.SetProperty("factories", value)
}

