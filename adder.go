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

type Adder struct {
	Element *gst.Element
}

func NewAdder() (*Adder, error) {
	e, err := gst.NewElement("adder")
	if err != nil {
		return nil, err
	}
	return &Adder{Element: e}, nil
}

func NewAdderWithName(name string) (*Adder, error) {
	e, err := gst.NewElementWithName("adder", name)
	if err != nil {
		return nil, err
	}
	return &Adder{Element: e}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// Set target format for mixing (NULL means ANY). Setting this property takes a reference to the supplied GstCaps object.

func (e *Adder) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Adder) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

