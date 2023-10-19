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
)

type Fdkaacenc struct {
	Element *gst.Element
}

func NewFdkaacenc() (*Fdkaacenc, error) {
	e, err := gst.NewElement("fdkaacenc")
	if err != nil {
		return nil, err
	}
	return &Fdkaacenc{Element: e}, nil
}

func NewFdkaacencWithName(name string) (*Fdkaacenc, error) {
	e, err := gst.NewElementWithName("fdkaacenc", name)
	if err != nil {
		return nil, err
	}
	return &Fdkaacenc{Element: e}, nil
}

// ----- Properties -----

// afterburner (bool)
//
// Afterburner - Quality Parameter.

func (e *Fdkaacenc) GetAfterburner() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("afterburner")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fdkaacenc) SetAfterburner(value bool) error {
	return e.Element.SetProperty("afterburner", value)
}

// bitrate (int)
//
// Target Audio Bitrate. Only applicable if rate-control=cbr. (0 = fixed value based on sample rate and channel count)

func (e *Fdkaacenc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fdkaacenc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// peak-bitrate (int)
//
// Peak Bitrate to adjust maximum bits per audio frame.

func (e *Fdkaacenc) GetPeakBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("peak-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fdkaacenc) SetPeakBitrate(value int) error {
	return e.Element.SetProperty("peak-bitrate", value)
}

// rate-control (GstFdkAacRateControl)
//
// Rate Control.

func (e *Fdkaacenc) GetRateControl() (GstFdkAacRateControl, error) {
	var value GstFdkAacRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFdkAacRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFdkAacRateControl")
	}
	return value, nil
}

func (e *Fdkaacenc) SetRateControl(value GstFdkAacRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// vbr-preset (GstFdkAacVbrPreset)
//
// AAC Variable Bitrate configurations.

func (e *Fdkaacenc) GetVbrPreset() (GstFdkAacVbrPreset, error) {
	var value GstFdkAacVbrPreset
	var ok bool
	v, err := e.Element.GetProperty("vbr-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFdkAacVbrPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFdkAacVbrPreset")
	}
	return value, nil
}

func (e *Fdkaacenc) SetVbrPreset(value GstFdkAacVbrPreset) error {
	e.Element.SetArg("vbr-preset", string(value))
	return nil
}

// ----- Constants -----

type GstFdkAacRateControl string

const (
	GstFdkAacRateControlCbr GstFdkAacRateControl = "cbr" // cbr (0) – Constant Bitrate
	GstFdkAacRateControlVbr GstFdkAacRateControl = "vbr" // vbr (1) – Variable Bitrate
)

type GstFdkAacVbrPreset string

const (
	GstFdkAacVbrPresetVeryLow GstFdkAacVbrPreset = "very-low" // very-low (1) – Very Low Variable Bitrate
	GstFdkAacVbrPresetLow GstFdkAacVbrPreset = "low" // low (2) – Low Variable Bitrate
	GstFdkAacVbrPresetMedium GstFdkAacVbrPreset = "medium" // medium (3) – Medium Variable Bitrate
	GstFdkAacVbrPresetHigh GstFdkAacVbrPreset = "high" // high (4) – High Variable Bitrate
	GstFdkAacVbrPresetVeryHigh GstFdkAacVbrPreset = "very-high" // very-high (5) – Very High Variable Bitrate
)

