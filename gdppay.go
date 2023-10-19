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

type Gdppay struct {
	Element *gst.Element
}

func NewGdppay() (*Gdppay, error) {
	e, err := gst.NewElement("gdppay")
	if err != nil {
		return nil, err
	}
	return &Gdppay{Element: e}, nil
}

func NewGdppayWithName(name string) (*Gdppay, error) {
	e, err := gst.NewElementWithName("gdppay", name)
	if err != nil {
		return nil, err
	}
	return &Gdppay{Element: e}, nil
}

// ----- Properties -----

// crc-header (bool)
//
// Calculate and store a CRC checksum on the header

func (e *Gdppay) GetCrcHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("crc-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Gdppay) SetCrcHeader(value bool) error {
	return e.Element.SetProperty("crc-header", value)
}

// crc-payload (bool)
//
// Calculate and store a CRC checksum on the payload

func (e *Gdppay) GetCrcPayload() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("crc-payload")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Gdppay) SetCrcPayload(value bool) error {
	return e.Element.SetProperty("crc-payload", value)
}

