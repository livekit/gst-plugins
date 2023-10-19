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

type Sdpsrc struct {
	Element *gst.Element
}

func NewSdpsrc() (*Sdpsrc, error) {
	e, err := gst.NewElement("sdpsrc")
	if err != nil {
		return nil, err
	}
	return &Sdpsrc{Element: e}, nil
}

func NewSdpsrcWithName(name string) (*Sdpsrc, error) {
	e, err := gst.NewElementWithName("sdpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Sdpsrc{Element: e}, nil
}

// ----- Properties -----

// location (string)
//
// URI to SDP file (sdp:///path/to/file)

func (e *Sdpsrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Sdpsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// sdp (string)
//
// SDP description used instead of location

func (e *Sdpsrc) GetSdp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sdp")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Sdpsrc) SetSdp(value string) error {
	return e.Element.SetProperty("sdp", value)
}

