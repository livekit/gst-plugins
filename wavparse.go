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

type Wavparse struct {
	Element *gst.Element
}

func NewWavparse() (*Wavparse, error) {
	e, err := gst.NewElement("wavparse")
	if err != nil {
		return nil, err
	}
	return &Wavparse{Element: e}, nil
}

func NewWavparseWithName(name string) (*Wavparse, error) {
	e, err := gst.NewElementWithName("wavparse", name)
	if err != nil {
		return nil, err
	}
	return &Wavparse{Element: e}, nil
}

// ----- Properties -----

// ignore-length (bool)
//
// This selects whether the length found in a data chunk
// should be ignored. This may be useful for streamed audio
// where the length is unknown until the end of streaming,
// and various software/hardware just puts some random value
// in there and hopes it doesn't break too much.

func (e *Wavparse) GetIgnoreLength() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wavparse) SetIgnoreLength(value bool) error {
	return e.Element.SetProperty("ignore-length", value)
}

