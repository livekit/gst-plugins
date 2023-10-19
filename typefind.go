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

type Typefind struct {
	Element *gst.Element
}

func NewTypefind() (*Typefind, error) {
	e, err := gst.NewElement("typefind")
	if err != nil {
		return nil, err
	}
	return &Typefind{Element: e}, nil
}

func NewTypefindWithName(name string) (*Typefind, error) {
	e, err := gst.NewElementWithName("typefind", name)
	if err != nil {
		return nil, err
	}
	return &Typefind{Element: e}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// detected capabilities in stream

func (e *Typefind) GetCaps() (*gst.Caps, error) {
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

// force-caps (GstCaps)
//
// force caps without doing a typefind

func (e *Typefind) GetForceCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("force-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Typefind) SetForceCaps(value *gst.Caps) error {
	return e.Element.SetProperty("force-caps", value)
}

// minimum (uint)
//
// minimum probability required to accept caps

func (e *Typefind) GetMinimum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("minimum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Typefind) SetMinimum(value uint) error {
	return e.Element.SetProperty("minimum", value)
}

