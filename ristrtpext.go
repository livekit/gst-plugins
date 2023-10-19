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

type Ristrtpext struct {
	Element *gst.Element
}

func NewRistrtpext() (*Ristrtpext, error) {
	e, err := gst.NewElement("ristrtpext")
	if err != nil {
		return nil, err
	}
	return &Ristrtpext{Element: e}, nil
}

func NewRistrtpextWithName(name string) (*Ristrtpext, error) {
	e, err := gst.NewElementWithName("ristrtpext", name)
	if err != nil {
		return nil, err
	}
	return &Ristrtpext{Element: e}, nil
}

// ----- Properties -----

// drop-null-ts-packets (bool)
//
// Drop null MPEG-TS packet and replace them with a custom header extension.

func (e *Ristrtpext) GetDropNullTsPackets() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-null-ts-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ristrtpext) SetDropNullTsPackets(value bool) error {
	return e.Element.SetProperty("drop-null-ts-packets", value)
}

// sequence-number-extension (bool)
//
// Add sequence number extension to packets.

func (e *Ristrtpext) GetSequenceNumberExtension() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sequence-number-extension")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ristrtpext) SetSequenceNumberExtension(value bool) error {
	return e.Element.SetProperty("sequence-number-extension", value)
}

