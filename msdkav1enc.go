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

type Msdkav1Enc struct {
	Element *gst.Element
}

func NewMsdkav1Enc() (*Msdkav1Enc, error) {
	e, err := gst.NewElement("msdkav1enc")
	if err != nil {
		return nil, err
	}
	return &Msdkav1Enc{Element: e}, nil
}

func NewMsdkav1EncWithName(name string) (*Msdkav1Enc, error) {
	e, err := gst.NewElementWithName("msdkav1enc", name)
	if err != nil {
		return nil, err
	}
	return &Msdkav1Enc{Element: e}, nil
}

// ----- Properties -----

// b-pyramid (bool)
//
// Enable B-Pyramid Reference structure

func (e *Msdkav1Enc) GetBPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkav1Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

// num-tile-cols (uint)
//
// number of columns for tiled encoding

func (e *Msdkav1Enc) GetNumTileCols() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-cols")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkav1Enc) SetNumTileCols(value uint) error {
	return e.Element.SetProperty("num-tile-cols", value)
}

// num-tile-rows (uint)
//
// number of rows for tiled encoding

func (e *Msdkav1Enc) GetNumTileRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkav1Enc) SetNumTileRows(value uint) error {
	return e.Element.SetProperty("num-tile-rows", value)
}

// p-pyramid (bool)
//
// Enable P-Pyramid Reference structure

func (e *Msdkav1Enc) GetPPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("p-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkav1Enc) SetPPyramid(value bool) error {
	return e.Element.SetProperty("p-pyramid", value)
}

