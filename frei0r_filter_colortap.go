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

type Frei0RFilterColortap struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterColortap() (*Frei0RFilterColortap, error) {
	e, err := gst.NewElement("frei0r-filter-colortap")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColortap{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterColortapWithName(name string) (*Frei0RFilterColortap, error) {
	e, err := gst.NewElementWithName("frei0r-filter-colortap", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterColortap{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// table (string)
//
// Lookup table used to filter colors. One of: xpro, sepia, heat, red_green, old_photo, xray, esses, yellow_blue

func (e *Frei0RFilterColortap) GetTable() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("table")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterColortap) SetTable(value string) error {
	return e.Element.SetProperty("table", value)
}

