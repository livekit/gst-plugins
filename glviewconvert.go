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

type Glviewconvert struct {
	*base.GstBaseTransform
}

func NewGlviewconvert() (*Glviewconvert, error) {
	e, err := gst.NewElement("glviewconvert")
	if err != nil {
		return nil, err
	}
	return &Glviewconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGlviewconvertWithName(name string) (*Glviewconvert, error) {
	e, err := gst.NewElementWithName("glviewconvert", name)
	if err != nil {
		return nil, err
	}
	return &Glviewconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// downmix-mode (GstGLStereoDownmix)
//
// Output anaglyph type to generate when downmixing to mono

func (e *Glviewconvert) GetDownmixMode() (interface{}, error) {
	return e.Element.GetProperty("downmix-mode")
}

func (e *Glviewconvert) SetDownmixMode(value interface{}) error {
	return e.Element.SetProperty("downmix-mode", value)
}

// input-flags-override (GstVideoMultiviewFlags)
//
// Override any input information about multiview layout flags

func (e *Glviewconvert) GetInputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("input-flags-override")
}

func (e *Glviewconvert) SetInputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("input-flags-override", value)
}

// input-mode-override (GstVideoMultiviewFramePacking)
//
// Override any input information about multiview layout

func (e *Glviewconvert) GetInputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("input-mode-override")
}

func (e *Glviewconvert) SetInputModeOverride(value interface{}) error {
	return e.Element.SetProperty("input-mode-override", value)
}

// output-flags-override (GstVideoMultiviewFlags)
//
// Override automatic negotiation for output multiview layout flags

func (e *Glviewconvert) GetOutputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("output-flags-override")
}

func (e *Glviewconvert) SetOutputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("output-flags-override", value)
}

// output-mode-override (GstVideoMultiviewMode)
//
// Override automatic output mode selection for multiview layout

func (e *Glviewconvert) GetOutputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("output-mode-override")
}

func (e *Glviewconvert) SetOutputModeOverride(value interface{}) error {
	return e.Element.SetProperty("output-mode-override", value)
}

// ----- Constants -----

type GstGlstereoDownmix string

const (
	GstGlstereoDownmixGreenMagentaDubois GstGlstereoDownmix = "green-magenta-dubois" // green-magenta-dubois (0) – GST_GL_STEREO_DOWNMIX_ANAGLYPH_GREEN_MAGENTA_DUBOIS
	GstGlstereoDownmixRedCyanDubois GstGlstereoDownmix = "red-cyan-dubois" // red-cyan-dubois (1) – GST_GL_STEREO_DOWNMIX_ANAGLYPH_RED_CYAN_DUBOIS
	GstGlstereoDownmixAmberBlueDubois GstGlstereoDownmix = "amber-blue-dubois" // amber-blue-dubois (2) – GST_GL_STEREO_DOWNMIX_ANAGLYPH_AMBER_BLUE_DUBOIS
)

