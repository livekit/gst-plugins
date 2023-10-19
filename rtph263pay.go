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

type Rtph263Pay struct {
	Element *gst.Element
}

func NewRtph263Pay() (*Rtph263Pay, error) {
	e, err := gst.NewElement("rtph263pay")
	if err != nil {
		return nil, err
	}
	return &Rtph263Pay{Element: e}, nil
}

func NewRtph263PayWithName(name string) (*Rtph263Pay, error) {
	e, err := gst.NewElementWithName("rtph263pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtph263Pay{Element: e}, nil
}

// ----- Properties -----

// modea-only (bool)
//
// Disable packetization modes B and C

func (e *Rtph263Pay) GetModeaOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("modea-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtph263Pay) SetModeaOnly(value bool) error {
	return e.Element.SetProperty("modea-only", value)
}

