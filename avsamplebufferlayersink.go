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

type Avsamplebufferlayersink struct {
	*base.GstBaseSink
}

func NewAvsamplebufferlayersink() (*Avsamplebufferlayersink, error) {
	e, err := gst.NewElement("avsamplebufferlayersink")
	if err != nil {
		return nil, err
	}
	return &Avsamplebufferlayersink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAvsamplebufferlayersinkWithName(name string) (*Avsamplebufferlayersink, error) {
	e, err := gst.NewElementWithName("avsamplebufferlayersink", name)
	if err != nil {
		return nil, err
	}
	return &Avsamplebufferlayersink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Avsamplebufferlayersink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avsamplebufferlayersink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// layer (GstGpointer)
//
// The CoreAnimation layer that can be placed in the render tree

func (e *Avsamplebufferlayersink) GetLayer() (interface{}, error) {
	return e.Element.GetProperty("layer")
}

