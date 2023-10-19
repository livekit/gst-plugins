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

type Bz2Enc struct {
	Element *gst.Element
}

func NewBz2Enc() (*Bz2Enc, error) {
	e, err := gst.NewElement("bz2enc")
	if err != nil {
		return nil, err
	}
	return &Bz2Enc{Element: e}, nil
}

func NewBz2EncWithName(name string) (*Bz2Enc, error) {
	e, err := gst.NewElementWithName("bz2enc", name)
	if err != nil {
		return nil, err
	}
	return &Bz2Enc{Element: e}, nil
}

// ----- Properties -----

// block-size (uint)
//
// Block size

func (e *Bz2Enc) GetBlockSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("block-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Bz2Enc) SetBlockSize(value uint) error {
	return e.Element.SetProperty("block-size", value)
}

// buffer-size (uint)
//
// Buffer size

func (e *Bz2Enc) GetBufferSize() (uint, error) {
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

func (e *Bz2Enc) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

