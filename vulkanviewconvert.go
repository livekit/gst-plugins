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

type Vulkanviewconvert struct {
	*base.GstBaseTransform
}

func NewVulkanviewconvert() (*Vulkanviewconvert, error) {
	e, err := gst.NewElement("vulkanviewconvert")
	if err != nil {
		return nil, err
	}
	return &Vulkanviewconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVulkanviewconvertWithName(name string) (*Vulkanviewconvert, error) {
	e, err := gst.NewElementWithName("vulkanviewconvert", name)
	if err != nil {
		return nil, err
	}
	return &Vulkanviewconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// downmix-mode (GstVulkanStereoDownmix)
//
// Output anaglyph type to generate when downmixing to mono

func (e *Vulkanviewconvert) GetDownmixMode() (GstVulkanStereoDownmix, error) {
	var value GstVulkanStereoDownmix
	var ok bool
	v, err := e.Element.GetProperty("downmix-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVulkanStereoDownmix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVulkanStereoDownmix")
	}
	return value, nil
}

func (e *Vulkanviewconvert) SetDownmixMode(value GstVulkanStereoDownmix) error {
	e.Element.SetArg("downmix-mode", string(value))
	return nil
}

// input-flags-override (GstVideoMultiviewFlags)
//
// Override any input information about multiview layout flags

func (e *Vulkanviewconvert) GetInputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("input-flags-override")
}

func (e *Vulkanviewconvert) SetInputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("input-flags-override", value)
}

// input-mode-override (GstVideoMultiviewMode)
//
// Override any input information about multiview layout

func (e *Vulkanviewconvert) GetInputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("input-mode-override")
}

func (e *Vulkanviewconvert) SetInputModeOverride(value interface{}) error {
	return e.Element.SetProperty("input-mode-override", value)
}

// output-flags-override (GstVideoMultiviewFlags)
//
// Override automatic negotiation for output multiview layout flags

func (e *Vulkanviewconvert) GetOutputFlagsOverride() (interface{}, error) {
	return e.Element.GetProperty("output-flags-override")
}

func (e *Vulkanviewconvert) SetOutputFlagsOverride(value interface{}) error {
	return e.Element.SetProperty("output-flags-override", value)
}

// output-mode-override (GstVideoMultiviewMode)
//
// Override automatic output mode selection for multiview layout

func (e *Vulkanviewconvert) GetOutputModeOverride() (interface{}, error) {
	return e.Element.GetProperty("output-mode-override")
}

func (e *Vulkanviewconvert) SetOutputModeOverride(value interface{}) error {
	return e.Element.SetProperty("output-mode-override", value)
}

// ----- Constants -----

type GstVulkanStereoDownmix string

const (
	GstVulkanStereoDownmixGreenMagentaDubois GstVulkanStereoDownmix = "green-magenta-dubois" // green-magenta-dubois (0) – GST_VULKAN_STEREO_DOWNMIX_ANAGLYPH_GREEN_MAGENTA_DUBOIS
	GstVulkanStereoDownmixRedCyanDubois GstVulkanStereoDownmix = "red-cyan-dubois" // red-cyan-dubois (1) – GST_VULKAN_STEREO_DOWNMIX_ANAGLYPH_RED_CYAN_DUBOIS
	GstVulkanStereoDownmixAmberBlueDubois GstVulkanStereoDownmix = "amber-blue-dubois" // amber-blue-dubois (2) – GST_VULKAN_STEREO_DOWNMIX_ANAGLYPH_AMBER_BLUE_DUBOIS
)

