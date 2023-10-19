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

type GleffectsHeat struct {
	*base.GstBaseTransform
}

func NewGleffectsHeat() (*GleffectsHeat, error) {
	e, err := gst.NewElement("gleffects_heat")
	if err != nil {
		return nil, err
	}
	return &GleffectsHeat{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGleffectsHeatWithName(name string) (*GleffectsHeat, error) {
	e, err := gst.NewElementWithName("gleffects_heat", name)
	if err != nil {
		return nil, err
	}
	return &GleffectsHeat{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// hswap (bool)
//
// Switch video texture left to right, useful with webcams

func (e *GleffectsHeat) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *GleffectsHeat) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}
