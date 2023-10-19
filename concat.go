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

type Concat struct {
	Element *gst.Element
}

func NewConcat() (*Concat, error) {
	e, err := gst.NewElement("concat")
	if err != nil {
		return nil, err
	}
	return &Concat{Element: e}, nil
}

func NewConcatWithName(name string) (*Concat, error) {
	e, err := gst.NewElementWithName("concat", name)
	if err != nil {
		return nil, err
	}
	return &Concat{Element: e}, nil
}

// ----- Properties -----

// active-pad (GstPad)
//
// Currently active sink pad

func (e *Concat) GetActivePad() (interface{}, error) {
	return e.Element.GetProperty("active-pad")
}

// adjust-base (bool)
//
// Adjust the base value of segments to ensure they are adjacent

func (e *Concat) GetAdjustBase() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("adjust-base")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Concat) SetAdjustBase(value bool) error {
	return e.Element.SetProperty("adjust-base", value)
}

