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

type Frei0RFilterPremultiplyOrUnpremultiply struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterPremultiplyOrUnpremultiply() (*Frei0RFilterPremultiplyOrUnpremultiply, error) {
	e, err := gst.NewElement("frei0r-filter-premultiply-or-unpremultiply")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPremultiplyOrUnpremultiply{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterPremultiplyOrUnpremultiplyWithName(name string) (*Frei0RFilterPremultiplyOrUnpremultiply, error) {
	e, err := gst.NewElementWithName("frei0r-filter-premultiply-or-unpremultiply", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterPremultiplyOrUnpremultiply{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// unpremultiply (bool)
//
// Whether to unpremultiply instead

func (e *Frei0RFilterPremultiplyOrUnpremultiply) GetUnpremultiply() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("unpremultiply")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterPremultiplyOrUnpremultiply) SetUnpremultiply(value bool) error {
	return e.Element.SetProperty("unpremultiply", value)
}

