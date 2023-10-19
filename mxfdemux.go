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

type Mxfdemux struct {
	Element *gst.Element
}

func NewMxfdemux() (*Mxfdemux, error) {
	e, err := gst.NewElement("mxfdemux")
	if err != nil {
		return nil, err
	}
	return &Mxfdemux{Element: e}, nil
}

func NewMxfdemuxWithName(name string) (*Mxfdemux, error) {
	e, err := gst.NewElementWithName("mxfdemux", name)
	if err != nil {
		return nil, err
	}
	return &Mxfdemux{Element: e}, nil
}

// ----- Properties -----

// max-drift (uint64)
//
// Maximum number of nanoseconds by which tracks can differ

func (e *Mxfdemux) GetMaxDrift() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-drift")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Mxfdemux) SetMaxDrift(value uint64) error {
	return e.Element.SetProperty("max-drift", value)
}

// package (string)
//
// Material or Source package to use for playback

func (e *Mxfdemux) GetPackage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("package")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Mxfdemux) SetPackage(value string) error {
	return e.Element.SetProperty("package", value)
}

// structure (GstStructure)
//
// Structural metadata of the MXF file

func (e *Mxfdemux) GetStructure() (interface{}, error) {
	return e.Element.GetProperty("structure")
}

