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

type LspPlugInPluginsLv2TriggerMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2TriggerMono() (*LspPlugInPluginsLv2TriggerMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-trigger-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2TriggerMonoWithName(name string) (*LspPlugInPluginsLv2TriggerMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-trigger-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// asel (GstLspPlugInPluginsLv2TriggerMonoasel)
//
// Area selector

func (e *LspPlugInPluginsLv2TriggerMono) GetAsel() (interface{}, error) {
	return e.Element.GetProperty("asel")
}

func (e *LspPlugInPluginsLv2TriggerMono) SetAsel(value interface{}) error {
	return e.Element.SetProperty("asel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2TriggerMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2TriggerMono) GetClear() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clear")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// dl (float32)
//
// Detect level

func (e *LspPlugInPluginsLv2TriggerMono) GetDl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetDl(value float32) error {
	return e.Element.SetProperty("dl", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2TriggerMono) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// dt (float32)
//
// Detect time

func (e *LspPlugInPluginsLv2TriggerMono) GetDt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetDt(value float32) error {
	return e.Element.SetProperty("dt", value)
}

// dtr1 (float32)
//
// Dynamics range 1

func (e *LspPlugInPluginsLv2TriggerMono) GetDtr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dtr1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetDtr1(value float32) error {
	return e.Element.SetProperty("dtr1", value)
}

// dtr2 (float32)
//
// Dynamics range 2

func (e *LspPlugInPluginsLv2TriggerMono) GetDtr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dtr2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetDtr2(value float32) error {
	return e.Element.SetProperty("dtr2", value)
}

// dyna (float32)
//
// Dynamics

func (e *LspPlugInPluginsLv2TriggerMono) GetDyna() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dyna")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetDyna(value float32) error {
	return e.Element.SetProperty("dyna", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2TriggerMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// ism (float32)
//
// Input signal meter

func (e *LspPlugInPluginsLv2TriggerMono) GetIsm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ism")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// isv (bool)
//
// Input signal display

func (e *LspPlugInPluginsLv2TriggerMono) GetIsv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("isv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetIsv(value bool) error {
	return e.Element.SetProperty("isv", value)
}

// lstn (bool)
//
// Trigger listen

func (e *LspPlugInPluginsLv2TriggerMono) GetLstn() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lstn")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mode (GstLspPlugInPluginsLv2TriggerMonomode)
//
// Detection mode

func (e *LspPlugInPluginsLv2TriggerMono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2TriggerMono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2TriggerMono) GetOutLatency() (int, error) {
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

// pause (bool)
//
// Pause graph analysis

func (e *LspPlugInPluginsLv2TriggerMono) GetPause() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pause")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// preamp (float32)
//
// Signal pre-amplification

func (e *LspPlugInPluginsLv2TriggerMono) GetPreamp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("preamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetPreamp(value float32) error {
	return e.Element.SetProperty("preamp", value)
}

// react (float32)
//
// Reactivity

func (e *LspPlugInPluginsLv2TriggerMono) GetReact() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("react")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2TriggerMono) GetRl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rrl (float32)
//
// Relative release level

func (e *LspPlugInPluginsLv2TriggerMono) GetRrl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2TriggerMono) GetRt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// tfm (float32)
//
// Trigger function meter

func (e *LspPlugInPluginsLv2TriggerMono) GetTfm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tfm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// tfv (bool)
//
// Trigger function display

func (e *LspPlugInPluginsLv2TriggerMono) GetTfv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("tfv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetTfv(value bool) error {
	return e.Element.SetProperty("tfv", value)
}

// tla (bool)
//
// Trigger activity

func (e *LspPlugInPluginsLv2TriggerMono) GetTla() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("tla")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// tlm (float32)
//
// Trigger level meter

func (e *LspPlugInPluginsLv2TriggerMono) GetTlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tlm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// tlv (bool)
//
// Trigger level display

func (e *LspPlugInPluginsLv2TriggerMono) GetTlv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("tlv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerMono) SetTlv(value bool) error {
	return e.Element.SetProperty("tlv", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2TriggerMono) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2TriggerMonoasel string

const (
	LspPlugInPluginsLv2TriggerMonoaselTrigger LspPlugInPluginsLv2TriggerMonoasel = "Trigger" // Trigger (0) – Trigger
	LspPlugInPluginsLv2TriggerMonoaselInstrument LspPlugInPluginsLv2TriggerMonoasel = "Instrument" // Instrument (1) – Instrument
)

type LspPlugInPluginsLv2TriggerMonomode string

const (
	LspPlugInPluginsLv2TriggerMonomodePeak LspPlugInPluginsLv2TriggerMonomode = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2TriggerMonomodeRms LspPlugInPluginsLv2TriggerMonomode = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2TriggerMonomodeLowPass LspPlugInPluginsLv2TriggerMonomode = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2TriggerMonomodeUniform LspPlugInPluginsLv2TriggerMonomode = "Uniform" // Uniform (3) – Uniform
)

