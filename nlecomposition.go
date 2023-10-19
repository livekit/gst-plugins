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

type Nlecomposition struct {
	Element *gst.Element
}

func NewNlecomposition() (*Nlecomposition, error) {
	e, err := gst.NewElement("nlecomposition")
	if err != nil {
		return nil, err
	}
	return &Nlecomposition{Element: e}, nil
}

func NewNlecompositionWithName(name string) (*Nlecomposition, error) {
	e, err := gst.NewElementWithName("nlecomposition", name)
	if err != nil {
		return nil, err
	}
	return &Nlecomposition{Element: e}, nil
}

// ----- Properties -----

// drop-tags (bool)
//
// Whether the composition should drop tags from its children

func (e *Nlecomposition) GetDropTags() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-tags")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Nlecomposition) SetDropTags(value bool) error {
	return e.Element.SetProperty("drop-tags", value)
}

// id (string)
//
// The stream-id of the composition

func (e *Nlecomposition) GetId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Nlecomposition) SetId(value string) error {
	return e.Element.SetProperty("id", value)
}

