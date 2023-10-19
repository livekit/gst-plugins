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

type Cairooverlay struct {
	*base.GstBaseTransform
}

func NewCairooverlay() (*Cairooverlay, error) {
	e, err := gst.NewElement("cairooverlay")
	if err != nil {
		return nil, err
	}
	return &Cairooverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCairooverlayWithName(name string) (*Cairooverlay, error) {
	e, err := gst.NewElementWithName("cairooverlay", name)
	if err != nil {
		return nil, err
	}
	return &Cairooverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// draw-on-transparent-surface (bool)
//
// Let the draw signal work on a transparent surface and blend the results with the video at a later time

func (e *Cairooverlay) GetDrawOnTransparentSurface() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-on-transparent-surface")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cairooverlay) SetDrawOnTransparentSurface(value bool) error {
	return e.Element.SetProperty("draw-on-transparent-surface", value)
}

