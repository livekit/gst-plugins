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

type LspPlugInPluginsLv2LatencyMeter struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2LatencyMeter() (*LspPlugInPluginsLv2LatencyMeter, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-latency-meter")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2LatencyMeter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2LatencyMeterWithName(name string) (*LspPlugInPluginsLv2LatencyMeter, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-latency-meter", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2LatencyMeter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// athr (float32)
//
// Absolute threshold

func (e *LspPlugInPluginsLv2LatencyMeter) GetAthr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("athr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetAthr(value float32) error {
	return e.Element.SetProperty("athr", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2LatencyMeter) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2LatencyMeter) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fback (bool)
//
// Feedback

func (e *LspPlugInPluginsLv2LatencyMeter) GetFback() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fback")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetFback(value bool) error {
	return e.Element.SetProperty("fback", value)
}

// gin (float32)
//
// Input Gain

func (e *LspPlugInPluginsLv2LatencyMeter) GetGin() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gin")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetGin(value float32) error {
	return e.Element.SetProperty("gin", value)
}

// gout (float32)
//
// Output Gain

func (e *LspPlugInPluginsLv2LatencyMeter) GetGout() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gout")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetGout(value float32) error {
	return e.Element.SetProperty("gout", value)
}

// ilvl (float32)
//
// Input Level

func (e *LspPlugInPluginsLv2LatencyMeter) GetIlvl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilvl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// l-v (float32)
//
// Latency Value

func (e *LspPlugInPluginsLv2LatencyMeter) GetLV() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("l-v")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// mlat (float32)
//
// Max expected latency

func (e *LspPlugInPluginsLv2LatencyMeter) GetMlat() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mlat")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetMlat(value float32) error {
	return e.Element.SetProperty("mlat", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2LatencyMeter) GetOutLatency() (int, error) {
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

// pthr (float32)
//
// Peak threshold

func (e *LspPlugInPluginsLv2LatencyMeter) GetPthr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pthr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetPthr(value float32) error {
	return e.Element.SetProperty("pthr", value)
}

// ttrig (bool)
//
// Trig a Latency measurement

func (e *LspPlugInPluginsLv2LatencyMeter) GetTtrig() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ttrig")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LatencyMeter) SetTtrig(value bool) error {
	return e.Element.SetProperty("ttrig", value)
}

