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

type Splitmuxsrc struct {
	Element *gst.Element
}

func NewSplitmuxsrc() (*Splitmuxsrc, error) {
	e, err := gst.NewElement("splitmuxsrc")
	if err != nil {
		return nil, err
	}
	return &Splitmuxsrc{Element: e}, nil
}

func NewSplitmuxsrcWithName(name string) (*Splitmuxsrc, error) {
	e, err := gst.NewElementWithName("splitmuxsrc", name)
	if err != nil {
		return nil, err
	}
	return &Splitmuxsrc{Element: e}, nil
}

// ----- Properties -----

// location (string)
//
// Glob pattern for the location of the files to read

func (e *Splitmuxsrc) GetLocation() (string, error) {
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

func (e *Splitmuxsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

