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

type LspPlugInPluginsLv2CompDelayX2Stereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2CompDelayX2Stereo() (*LspPlugInPluginsLv2CompDelayX2Stereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-comp-delay-x2-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompDelayX2Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2CompDelayX2StereoWithName(name string) (*LspPlugInPluginsLv2CompDelayX2Stereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-comp-delay-x2-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompDelayX2Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetBypass() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bypass")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cm-l (float32)
//
// Centimeters L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetCmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetCmL(value float32) error {
	return e.Element.SetProperty("cm-l", value)
}

// cm-r (float32)
//
// Centimeters R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetCmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetCmR(value float32) error {
	return e.Element.SetProperty("cm-r", value)
}

// d-d-l (float32)
//
// Delay distance L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDDL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("d-d-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// d-d-r (float32)
//
// Delay distance R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDDR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("d-d-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// d-s-l (int)
//
// Delay samples L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDSL() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("d-s-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// d-s-r (int)
//
// Delay samples R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDSR() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("d-s-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// d-t-l (float32)
//
// Delay time L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDTL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("d-t-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// d-t-r (float32)
//
// Delay time R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDTR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("d-t-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// dry-l (float32)
//
// Dry amount L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDryL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dry-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetDryL(value float32) error {
	return e.Element.SetProperty("dry-l", value)
}

// dry-r (float32)
//
// Dry amount R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetDryR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dry-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetDryR(value float32) error {
	return e.Element.SetProperty("dry-r", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetGOut() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-out")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// m-l (int)
//
// Meters L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetML() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("m-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetML(value int) error {
	return e.Element.SetProperty("m-l", value)
}

// m-r (int)
//
// Meters R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetMR() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("m-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetMR(value int) error {
	return e.Element.SetProperty("m-r", value)
}

// mode-l (GstLspPlugInPluginsLv2CompDelayX2StereomodeL)
//
// Mode L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetModeL() (interface{}, error) {
	return e.Element.GetProperty("mode-l")
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetModeL(value interface{}) error {
	return e.Element.SetProperty("mode-l", value)
}

// mode-r (GstLspPlugInPluginsLv2CompDelayX2StereomodeR)
//
// Mode R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetModeR() (interface{}, error) {
	return e.Element.GetProperty("mode-r")
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetModeR(value interface{}) error {
	return e.Element.SetProperty("mode-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetOutLatency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("out-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ramp-l (bool)
//
// Ramping L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetRampL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ramp-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetRampL(value bool) error {
	return e.Element.SetProperty("ramp-l", value)
}

// ramp-r (bool)
//
// Ramping R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetRampR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ramp-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetRampR(value bool) error {
	return e.Element.SetProperty("ramp-r", value)
}

// samp-l (int)
//
// Samples L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetSampL() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samp-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetSampL(value int) error {
	return e.Element.SetProperty("samp-l", value)
}

// samp-r (int)
//
// Samples R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetSampR() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samp-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetSampR(value int) error {
	return e.Element.SetProperty("samp-r", value)
}

// t-l (float32)
//
// Temperature L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetTL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("t-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetTL(value float32) error {
	return e.Element.SetProperty("t-l", value)
}

// t-r (float32)
//
// Temperature R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetTR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("t-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetTR(value float32) error {
	return e.Element.SetProperty("t-r", value)
}

// time-l (float32)
//
// Time L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetTimeL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("time-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetTimeL(value float32) error {
	return e.Element.SetProperty("time-l", value)
}

// time-r (float32)
//
// Time R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetTimeR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("time-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetTimeR(value float32) error {
	return e.Element.SetProperty("time-r", value)
}

// wet-l (float32)
//
// Wet amount L

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetWetL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("wet-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetWetL(value float32) error {
	return e.Element.SetProperty("wet-l", value)
}

// wet-r (float32)
//
// Wet amount R

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) GetWetR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("wet-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayX2Stereo) SetWetR(value float32) error {
	return e.Element.SetProperty("wet-r", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2CompDelayX2StereomodeL string

const (
	LspPlugInPluginsLv2CompDelayX2StereomodeLSamples LspPlugInPluginsLv2CompDelayX2StereomodeL = "Samples" // Samples (0) – Samples
	LspPlugInPluginsLv2CompDelayX2StereomodeLDistance LspPlugInPluginsLv2CompDelayX2StereomodeL = "Distance" // Distance (1) – Distance
	LspPlugInPluginsLv2CompDelayX2StereomodeLTime LspPlugInPluginsLv2CompDelayX2StereomodeL = "Time" // Time (2) – Time
)

type LspPlugInPluginsLv2CompDelayX2StereomodeR string

const (
	LspPlugInPluginsLv2CompDelayX2StereomodeRSamples LspPlugInPluginsLv2CompDelayX2StereomodeR = "Samples" // Samples (0) – Samples
	LspPlugInPluginsLv2CompDelayX2StereomodeRDistance LspPlugInPluginsLv2CompDelayX2StereomodeR = "Distance" // Distance (1) – Distance
	LspPlugInPluginsLv2CompDelayX2StereomodeRTime LspPlugInPluginsLv2CompDelayX2StereomodeR = "Time" // Time (2) – Time
)

