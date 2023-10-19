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

type Splitfilesrc struct {
	*base.GstBaseSrc
}

func NewSplitfilesrc() (*Splitfilesrc, error) {
	e, err := gst.NewElement("splitfilesrc")
	if err != nil {
		return nil, err
	}
	return &Splitfilesrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewSplitfilesrcWithName(name string) (*Splitfilesrc, error) {
	e, err := gst.NewElementWithName("splitfilesrc", name)
	if err != nil {
		return nil, err
	}
	return &Splitfilesrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// location (string)
//
// Wildcard pattern to match file names of the input files. If the location is an absolute path or contains directory components, only the base file name part will be considered for pattern matching. The results will be sorted.

func (e *Splitfilesrc) GetLocation() (string, error) {
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

func (e *Splitfilesrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

