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

	"github.com/livekit/gstplugins/base"
)

type Giosrc struct {
	*base.GstBaseSrc
}

func NewGiosrc() (*Giosrc, error) {
	e, err := gst.NewElement("giosrc")
	if err != nil {
		return nil, err
	}
	return &Giosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewGiosrcWithName(name string) (*Giosrc, error) {
	e, err := gst.NewElementWithName("giosrc", name)
	if err != nil {
		return nil, err
	}
	return &Giosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// file (GstGfile)
//
// GFile to read from.

func (e *Giosrc) GetFile() (interface{}, error) {
	return e.Element.GetProperty("file")
}

func (e *Giosrc) SetFile(value interface{}) error {
	return e.Element.SetProperty("file", value)
}

// is-growing (bool)
//
// Whether the file is currently growing. When activated EOS is never pushed
// and the user needs to handle it himself. This modes allows to keep reading
// the file while it is being written on file.

// location (string)
//
// URI location to read from

func (e *Giosrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Giosrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

