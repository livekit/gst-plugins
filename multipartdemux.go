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

type Multipartdemux struct {
	Element *gst.Element
}

func NewMultipartdemux() (*Multipartdemux, error) {
	e, err := gst.NewElement("multipartdemux")
	if err != nil {
		return nil, err
	}
	return &Multipartdemux{Element: e}, nil
}

func NewMultipartdemuxWithName(name string) (*Multipartdemux, error) {
	e, err := gst.NewElementWithName("multipartdemux", name)
	if err != nil {
		return nil, err
	}
	return &Multipartdemux{Element: e}, nil
}

// ----- Properties -----

// boundary (string)
//
// The boundary string separating data, automatic if NULL

func (e *Multipartdemux) GetBoundary() (string, error) {
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

func (e *Multipartdemux) SetBoundary(value string) error {
	return e.Element.SetProperty("boundary", value)
}

// single-stream (bool)
//
// Assume that there is only one stream whose content-type will
// not change and emit no-more-pads as soon as the first boundary
// content is parsed, decoded, and pads are linked.

func (e *Multipartdemux) GetSingleStream() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("single-stream")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multipartdemux) SetSingleStream(value bool) error {
	return e.Element.SetProperty("single-stream", value)
}

