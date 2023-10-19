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

type Multifdsink struct {
	*base.GstBaseSink
}

func NewMultifdsink() (*Multifdsink, error) {
	e, err := gst.NewElement("multifdsink")
	if err != nil {
		return nil, err
	}
	return &Multifdsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewMultifdsinkWithName(name string) (*Multifdsink, error) {
	e, err := gst.NewElementWithName("multifdsink", name)
	if err != nil {
		return nil, err
	}
	return &Multifdsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// handle-read (bool)
//
// Handle client reads and discard the data

func (e *Multifdsink) GetHandleRead() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-read")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multifdsink) SetHandleRead(value bool) error {
	return e.Element.SetProperty("handle-read", value)
}

