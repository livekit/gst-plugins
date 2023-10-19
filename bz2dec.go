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

type Bz2Dec struct {
	Element *gst.Element
}

func NewBz2Dec() (*Bz2Dec, error) {
	e, err := gst.NewElement("bz2dec")
	if err != nil {
		return nil, err
	}
	return &Bz2Dec{Element: e}, nil
}

func NewBz2DecWithName(name string) (*Bz2Dec, error) {
	e, err := gst.NewElementWithName("bz2dec", name)
	if err != nil {
		return nil, err
	}
	return &Bz2Dec{Element: e}, nil
}

// ----- Properties -----

// buffer-size (uint)
//
// Buffer size

func (e *Bz2Dec) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Bz2Dec) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

// first-buffer-size (uint)
//
// Size of first buffer (used to determine the mime type of the uncompressed data)

func (e *Bz2Dec) GetFirstBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("first-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Bz2Dec) SetFirstBufferSize(value uint) error {
	return e.Element.SetProperty("first-buffer-size", value)
}

