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

type Msdkmpeg2Dec struct {
	Element *gst.Element
}

func NewMsdkmpeg2Dec() (*Msdkmpeg2Dec, error) {
	e, err := gst.NewElement("msdkmpeg2dec")
	if err != nil {
		return nil, err
	}
	return &Msdkmpeg2Dec{Element: e}, nil
}

func NewMsdkmpeg2DecWithName(name string) (*Msdkmpeg2Dec, error) {
	e, err := gst.NewElementWithName("msdkmpeg2dec", name)
	if err != nil {
		return nil, err
	}
	return &Msdkmpeg2Dec{Element: e}, nil
}

// ----- Properties -----

// output-order (GstMsdkDecOutputOrder)
//
// Decoded frames output order

func (e *Msdkmpeg2Dec) GetOutputOrder() (interface{}, error) {
	return e.Element.GetProperty("output-order")
}

func (e *Msdkmpeg2Dec) SetOutputOrder(value interface{}) error {
	return e.Element.SetProperty("output-order", value)
}

