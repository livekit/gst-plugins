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

type LspPlugInPluginsLv2CompDelayMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2CompDelayMono() (*LspPlugInPluginsLv2CompDelayMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-comp-delay-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompDelayMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2CompDelayMonoWithName(name string) (*LspPlugInPluginsLv2CompDelayMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-comp-delay-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompDelayMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2CompDelayMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompDelayMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cm (float32)
//
// Centimeters

func (e *LspPlugInPluginsLv2CompDelayMono) GetCm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetCm(value float32) error {
	return e.Element.SetProperty("cm", value)
}

// d-d (float32)
//
// Delay distance

func (e *LspPlugInPluginsLv2CompDelayMono) GetDD() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("d-d")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// d-s (int)
//
// Delay samples

func (e *LspPlugInPluginsLv2CompDelayMono) GetDS() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("d-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// d-t (float32)
//
// Delay time

func (e *LspPlugInPluginsLv2CompDelayMono) GetDT() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("d-t")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2CompDelayMono) GetDry() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dry")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2CompDelayMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompDelayMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// m (int)
//
// Meters

func (e *LspPlugInPluginsLv2CompDelayMono) GetM() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("m")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetM(value int) error {
	return e.Element.SetProperty("m", value)
}

// mode (GstLspPlugInPluginsLv2CompDelayMonomode)
//
// Mode

func (e *LspPlugInPluginsLv2CompDelayMono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2CompDelayMono) GetOutLatency() (int, error) {
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

// ramp (bool)
//
// Ramping

func (e *LspPlugInPluginsLv2CompDelayMono) GetRamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ramp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetRamp(value bool) error {
	return e.Element.SetProperty("ramp", value)
}

// samp (int)
//
// Samples

func (e *LspPlugInPluginsLv2CompDelayMono) GetSamp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetSamp(value int) error {
	return e.Element.SetProperty("samp", value)
}

// t (float32)
//
// Temperature

func (e *LspPlugInPluginsLv2CompDelayMono) GetT() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("t")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetT(value float32) error {
	return e.Element.SetProperty("t", value)
}

// time (float32)
//
// Time

func (e *LspPlugInPluginsLv2CompDelayMono) GetTime() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("time")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetTime(value float32) error {
	return e.Element.SetProperty("time", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2CompDelayMono) GetWet() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("wet")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompDelayMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2CompDelayMonomode string

const (
	LspPlugInPluginsLv2CompDelayMonomodeSamples LspPlugInPluginsLv2CompDelayMonomode = "Samples" // Samples (0) – Samples
	LspPlugInPluginsLv2CompDelayMonomodeDistance LspPlugInPluginsLv2CompDelayMonomode = "Distance" // Distance (1) – Distance
	LspPlugInPluginsLv2CompDelayMonomodeTime LspPlugInPluginsLv2CompDelayMonomode = "Time" // Time (2) – Time
)

