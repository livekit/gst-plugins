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

type Multipartmux struct {
	Element *gst.Element
}

func NewMultipartmux() (*Multipartmux, error) {
	e, err := gst.NewElement("multipartmux")
	if err != nil {
		return nil, err
	}
	return &Multipartmux{Element: e}, nil
}

func NewMultipartmuxWithName(name string) (*Multipartmux, error) {
	e, err := gst.NewElementWithName("multipartmux", name)
	if err != nil {
		return nil, err
	}
	return &Multipartmux{Element: e}, nil
}

// ----- Properties -----

// boundary (string)
//
// Boundary string

func (e *Multipartmux) GetBoundary() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("boundary")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Multipartmux) SetBoundary(value string) error {
	return e.Element.SetProperty("boundary", value)
}

