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

type Id3Mux struct {
	Element *gst.Element
}

func NewId3Mux() (*Id3Mux, error) {
	e, err := gst.NewElement("id3mux")
	if err != nil {
		return nil, err
	}
	return &Id3Mux{Element: e}, nil
}

func NewId3MuxWithName(name string) (*Id3Mux, error) {
	e, err := gst.NewElementWithName("id3mux", name)
	if err != nil {
		return nil, err
	}
	return &Id3Mux{Element: e}, nil
}

// ----- Properties -----

// v2-version (int)
//
// Set version (3 for id3v2.3, 4 for id3v2.4) of id3v2 tags

func (e *Id3Mux) GetV2Version() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("v2-version")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Id3Mux) SetV2Version(value int) error {
	return e.Element.SetProperty("v2-version", value)
}

// write-v1 (bool)
//
// Write an id3v1 tag at the end of the file

func (e *Id3Mux) GetWriteV1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("write-v1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Id3Mux) SetWriteV1(value bool) error {
	return e.Element.SetProperty("write-v1", value)
}

// write-v2 (bool)
//
// Write an id3v2 tag at the start of the file

func (e *Id3Mux) GetWriteV2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("write-v2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Id3Mux) SetWriteV2(value bool) error {
	return e.Element.SetProperty("write-v2", value)
}

