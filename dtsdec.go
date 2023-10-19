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

type Dtsdec struct {
	Element *gst.Element
}

func NewDtsdec() (*Dtsdec, error) {
	e, err := gst.NewElement("dtsdec")
	if err != nil {
		return nil, err
	}
	return &Dtsdec{Element: e}, nil
}

func NewDtsdecWithName(name string) (*Dtsdec, error) {
	e, err := gst.NewElementWithName("dtsdec", name)
	if err != nil {
		return nil, err
	}
	return &Dtsdec{Element: e}, nil
}

// ----- Properties -----

// drc (bool)
//
// Use Dynamic Range Compression

func (e *Dtsdec) GetDrc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dtsdec) SetDrc(value bool) error {
	return e.Element.SetProperty("drc", value)
}

