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

type Vulkansink struct {
	*base.GstBaseSink
}

func NewVulkansink() (*Vulkansink, error) {
	e, err := gst.NewElement("vulkansink")
	if err != nil {
		return nil, err
	}
	return &Vulkansink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewVulkansinkWithName(name string) (*Vulkansink, error) {
	e, err := gst.NewElementWithName("vulkansink", name)
	if err != nil {
		return nil, err
	}
	return &Vulkansink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// device (GstVulkanDevice)
//
// Vulkan device

func (e *Vulkansink) GetDevice() (interface{}, error) {
	return e.Element.GetProperty("device")
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Vulkansink) GetForceAspectRatio() (bool, error) {
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

func (e *Vulkansink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// pixel-aspect-ratio (GstFraction)
//
// The pixel aspect ratio of the device

func (e *Vulkansink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

func (e *Vulkansink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

