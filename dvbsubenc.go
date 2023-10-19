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

type Dvbsubenc struct {
	Element *gst.Element
}

func NewDvbsubenc() (*Dvbsubenc, error) {
	e, err := gst.NewElement("dvbsubenc")
	if err != nil {
		return nil, err
	}
	return &Dvbsubenc{Element: e}, nil
}

func NewDvbsubencWithName(name string) (*Dvbsubenc, error) {
	e, err := gst.NewElementWithName("dvbsubenc", name)
	if err != nil {
		return nil, err
	}
	return &Dvbsubenc{Element: e}, nil
}

// ----- Properties -----

// max-colours (int)
//
// Set the maximum number of colours to output into the DVB subpictures.
// Good choices are 4, 16 or 256 - as they correspond to the 2-bit, 4-bit
// and 8-bit palette modes that the DVB subpicture encoding supports.

func (e *Dvbsubenc) GetMaxColours() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-colours")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dvbsubenc) SetMaxColours(value int) error {
	return e.Element.SetProperty("max-colours", value)
}

// ts-offset (int64)
//
// Advance or delay the output subpicture time-line. This is a
// convenience property for setting the src pad offset.

func (e *Dvbsubenc) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Dvbsubenc) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

