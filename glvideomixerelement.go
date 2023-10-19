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
	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Glvideomixerelement struct {
	*base.GstAggregator
}

func NewGlvideomixerelement() (*Glvideomixerelement, error) {
	e, err := gst.NewElement("glvideomixerelement")
	if err != nil {
		return nil, err
	}
	return &Glvideomixerelement{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewGlvideomixerelementWithName(name string) (*Glvideomixerelement, error) {
	e, err := gst.NewElementWithName("glvideomixerelement", name)
	if err != nil {
		return nil, err
	}
	return &Glvideomixerelement{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// background (GstGLVideoMixerBackground)
//
// Background type

func (e *Glvideomixerelement) GetBackground() (interface{}, error) {
	return e.Element.GetProperty("background")
}

func (e *Glvideomixerelement) SetBackground(value interface{}) error {
	return e.Element.SetProperty("background", value)
}

// ----- Constants -----

type GstGlvideoMixerBackground string

const (
	GstGlvideoMixerBackgroundChecker GstGlvideoMixerBackground = "checker" // checker (0) – Checker pattern
	GstGlvideoMixerBackgroundBlack GstGlvideoMixerBackground = "black" // black (1) – Black
	GstGlvideoMixerBackgroundWhite GstGlvideoMixerBackground = "white" // white (2) – White
	GstGlvideoMixerBackgroundTransparent GstGlvideoMixerBackground = "transparent" // transparent (3) – Transparent Background to enable further compositing
)

