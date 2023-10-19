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

type Frei0RFilterCartoon struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterCartoon() (*Frei0RFilterCartoon, error) {
	e, err := gst.NewElement("frei0r-filter-cartoon")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCartoon{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterCartoonWithName(name string) (*Frei0RFilterCartoon, error) {
	e, err := gst.NewElementWithName("frei0r-filter-cartoon", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCartoon{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// diffspace (float64)
//
// difference space: a value from 0 to 256 (mapped to [0,1])

func (e *Frei0RFilterCartoon) GetDiffspace() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("diffspace")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCartoon) SetDiffspace(value float64) error {
	return e.Element.SetProperty("diffspace", value)
}

// triplevel (float64)
//
// level of trip: mapped to [0,1] asymptotical

func (e *Frei0RFilterCartoon) GetTriplevel() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("triplevel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCartoon) SetTriplevel(value float64) error {
	return e.Element.SetProperty("triplevel", value)
}

