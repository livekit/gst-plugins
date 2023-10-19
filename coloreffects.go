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

type Coloreffects struct {
	*base.GstBaseTransform
}

func NewColoreffects() (*Coloreffects, error) {
	e, err := gst.NewElement("coloreffects")
	if err != nil {
		return nil, err
	}
	return &Coloreffects{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewColoreffectsWithName(name string) (*Coloreffects, error) {
	e, err := gst.NewElementWithName("coloreffects", name)
	if err != nil {
		return nil, err
	}
	return &Coloreffects{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// preset (GstColorEffectsPreset)
//
// Color effect preset to use

func (e *Coloreffects) GetPreset() (GstColorEffectsPreset, error) {
	var value GstColorEffectsPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstColorEffectsPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstColorEffectsPreset")
	}
	return value, nil
}

func (e *Coloreffects) SetPreset(value GstColorEffectsPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

// ----- Constants -----

type GstColorEffectsPreset string

const (
	GstColorEffectsPresetNone GstColorEffectsPreset = "none" // none (0) – Do nothing preset
	GstColorEffectsPresetHeat GstColorEffectsPreset = "heat" // heat (1) – Fake heat camera toning
	GstColorEffectsPresetSepia GstColorEffectsPreset = "sepia" // sepia (2) – Sepia toning
	GstColorEffectsPresetXray GstColorEffectsPreset = "xray" // xray (3) – Invert and slightly shade to blue
	GstColorEffectsPresetXpro GstColorEffectsPreset = "xpro" // xpro (4) – Cross processing toning
	GstColorEffectsPresetYellowblue GstColorEffectsPreset = "yellowblue" // yellowblue (5) – Yellow foreground Blue background color filter
)

