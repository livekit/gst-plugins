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

type Glfilterbin struct {
	Element *gst.Element
}

func NewGlfilterbin() (*Glfilterbin, error) {
	e, err := gst.NewElement("glfilterbin")
	if err != nil {
		return nil, err
	}
	return &Glfilterbin{Element: e}, nil
}

func NewGlfilterbinWithName(name string) (*Glfilterbin, error) {
	e, err := gst.NewElementWithName("glfilterbin", name)
	if err != nil {
		return nil, err
	}
	return &Glfilterbin{Element: e}, nil
}

// ----- Properties -----

// filter (GstElement)
//
// The GL filter chain to use

func (e *Glfilterbin) GetFilter() (interface{}, error) {
	return e.Element.GetProperty("filter")
}

func (e *Glfilterbin) SetFilter(value interface{}) error {
	return e.Element.SetProperty("filter", value)
}

