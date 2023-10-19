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

type Msdkvp9Dec struct {
	Element *gst.Element
}

func NewMsdkvp9Dec() (*Msdkvp9Dec, error) {
	e, err := gst.NewElement("msdkvp9dec")
	if err != nil {
		return nil, err
	}
	return &Msdkvp9Dec{Element: e}, nil
}

func NewMsdkvp9DecWithName(name string) (*Msdkvp9Dec, error) {
	e, err := gst.NewElementWithName("msdkvp9dec", name)
	if err != nil {
		return nil, err
	}
	return &Msdkvp9Dec{Element: e}, nil
}

// ----- Properties -----

// output-order (GstMsdkDecOutputOrder)
//
// Decoded frames output order

func (e *Msdkvp9Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

func (e *Msdkvp9Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

// ----- Constants -----

type GstMsdkDecOutputOrder string

const (
	GstMsdkDecOutputOrderDisplay GstMsdkDecOutputOrder = "display" // display (0) – Output frames in Display order
	GstMsdkDecOutputOrderDecoded GstMsdkDecOutputOrder = "decoded" // decoded (1) – Output frames in Decoded order
)

