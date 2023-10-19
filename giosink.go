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

type Giosink struct {
	*base.GstBaseSink
}

func NewGiosink() (*Giosink, error) {
	e, err := gst.NewElement("giosink")
	if err != nil {
		return nil, err
	}
	return &Giosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGiosinkWithName(name string) (*Giosink, error) {
	e, err := gst.NewElementWithName("giosink", name)
	if err != nil {
		return nil, err
	}
	return &Giosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// file (GstGfile)
//
// GFile to write to.

func (e *Giosink) GetFile() (interface{}, error) {
	return e.Element.GetProperty("file")
}

func (e *Giosink) SetFile(value interface{}) error {
	return e.Element.SetProperty("file", value)
}

// location (string)
//
// URI location to write to

func (e *Giosink) GetLocation() (string, error) {
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

func (e *Giosink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

