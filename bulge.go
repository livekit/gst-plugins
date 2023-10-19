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

type Bulge struct {
	*base.GstBaseTransform
}

func NewBulge() (*Bulge, error) {
	e, err := gst.NewElement("bulge")
	if err != nil {
		return nil, err
	}
	return &Bulge{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewBulgeWithName(name string) (*Bulge, error) {
	e, err := gst.NewElementWithName("bulge", name)
	if err != nil {
		return nil, err
	}
	return &Bulge{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// zoom (float64)
//
// Zoom of the bulge effect

func (e *Bulge) GetZoom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Bulge) SetZoom(value float64) error {
	return e.Element.SetProperty("zoom", value)
}

