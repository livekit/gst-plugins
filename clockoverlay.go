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

type Clockoverlay struct {
	Element *gst.Element
}

func NewClockoverlay() (*Clockoverlay, error) {
	e, err := gst.NewElement("clockoverlay")
	if err != nil {
		return nil, err
	}
	return &Clockoverlay{Element: e}, nil
}

func NewClockoverlayWithName(name string) (*Clockoverlay, error) {
	e, err := gst.NewElementWithName("clockoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Clockoverlay{Element: e}, nil
}

// ----- Properties -----

// time-format (string)
//
// Format to use for time and date value, as in strftime.

func (e *Clockoverlay) GetTimeFormat() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("time-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Clockoverlay) SetTimeFormat(value string) error {
	return e.Element.SetProperty("time-format", value)
}

