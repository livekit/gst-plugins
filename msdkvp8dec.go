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

type Msdkvp8Dec struct {
	Element *gst.Element
}

func NewMsdkvp8Dec() (*Msdkvp8Dec, error) {
	e, err := gst.NewElement("msdkvp8dec")
	if err != nil {
		return nil, err
	}
	return &Msdkvp8Dec{Element: e}, nil
}

func NewMsdkvp8DecWithName(name string) (*Msdkvp8Dec, error) {
	e, err := gst.NewElementWithName("msdkvp8dec", name)
	if err != nil {
		return nil, err
	}
	return &Msdkvp8Dec{Element: e}, nil
}

// ----- Properties -----

// output-order (GstMsdkDecOutputOrder)
//
// Decoded frames output order

func (e *Msdkvp8Dec) GetOutputOrder() (interface{}, error) {
	return e.Element.GetProperty("output-order")
}

func (e *Msdkvp8Dec) SetOutputOrder(value interface{}) error {
	return e.Element.SetProperty("output-order", value)
}

