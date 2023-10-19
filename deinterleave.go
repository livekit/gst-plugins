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

type Deinterleave struct {
	Element *gst.Element
}

func NewDeinterleave() (*Deinterleave, error) {
	e, err := gst.NewElement("deinterleave")
	if err != nil {
		return nil, err
	}
	return &Deinterleave{Element: e}, nil
}

func NewDeinterleaveWithName(name string) (*Deinterleave, error) {
	e, err := gst.NewElementWithName("deinterleave", name)
	if err != nil {
		return nil, err
	}
	return &Deinterleave{Element: e}, nil
}

// ----- Properties -----

// keep-positions (bool)
//
// Keep positions: When enable the caps on the output buffers will
// contain the original channel positions. This can be used to correctly
// interleave the output again later but can also lead to unwanted effects
// if the output should be handled as Mono.

func (e *Deinterleave) GetKeepPositions() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-positions")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Deinterleave) SetKeepPositions(value bool) error {
	return e.Element.SetProperty("keep-positions", value)
}

