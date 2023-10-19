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

type LspPlugInPluginsLv2TriggerStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2TriggerStereo() (*LspPlugInPluginsLv2TriggerStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-trigger-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2TriggerStereoWithName(name string) (*LspPlugInPluginsLv2TriggerStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-trigger-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// asel (GstLspPlugInPluginsLv2TriggerStereoasel)
//
// Area selector

func (e *LspPlugInPluginsLv2TriggerStereo) GetAsel() (interface{}, error) {
	return e.Element.GetProperty("asel")
}

func (e *LspPlugInPluginsLv2TriggerStereo) SetAsel(value interface{}) error {
	return e.Element.SetProperty("asel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2TriggerStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2TriggerStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// dl (float32)
//
// Detect level

func (e *LspPlugInPluginsLv2TriggerStereo) GetDl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetDl(value float32) error {
	return e.Element.SetProperty("dl", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2TriggerStereo) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// dt (float32)
//
// Detect time

func (e *LspPlugInPluginsLv2TriggerStereo) GetDt() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetDt(value float32) error {
	return e.Element.SetProperty("dt", value)
}

// dtr1 (float32)
//
// Dynamics range 1

func (e *LspPlugInPluginsLv2TriggerStereo) GetDtr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetDtr1(value float32) error {
	return e.Element.SetProperty("dtr1", value)
}

// dtr2 (float32)
//
// Dynamics range 2

func (e *LspPlugInPluginsLv2TriggerStereo) GetDtr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetDtr2(value float32) error {
	return e.Element.SetProperty("dtr2", value)
}

// dyna (float32)
//
// Dynamics

func (e *LspPlugInPluginsLv2TriggerStereo) GetDyna() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetDyna(value float32) error {
	return e.Element.SetProperty("dyna", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2TriggerStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// isml (float32)
//
// Input signal meter left

func (e *LspPlugInPluginsLv2TriggerStereo) GetIsml() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("isml")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ismr (float32)
//
// Input signal meter right

func (e *LspPlugInPluginsLv2TriggerStereo) GetIsmr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ismr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// isvl (bool)
//
// Input signal left display

func (e *LspPlugInPluginsLv2TriggerStereo) GetIsvl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("isvl")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerStereo) SetIsvl(value bool) error {
	return e.Element.SetProperty("isvl", value)
}

// isvr (bool)
//
// Input signal right display

func (e *LspPlugInPluginsLv2TriggerStereo) GetIsvr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("isvr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2TriggerStereo) SetIsvr(value bool) error {
	return e.Element.SetProperty("isvr", value)
}

// lstn (bool)
//
// Trigger listen

func (e *LspPlugInPluginsLv2TriggerStereo) GetLstn() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mode (GstLspPlugInPluginsLv2TriggerStereomode)
//
// Detection mode

func (e *LspPlugInPluginsLv2TriggerStereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2TriggerStereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2TriggerStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// preamp (float32)
//
// Signal pre-amplification

func (e *LspPlugInPluginsLv2TriggerStereo) GetPreamp() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetPreamp(value float32) error {
	return e.Element.SetProperty("preamp", value)
}

// react (float32)
//
// Reactivity

func (e *LspPlugInPluginsLv2TriggerStereo) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2TriggerStereo) GetRl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2TriggerStereo) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// ssrc (GstLspPlugInPluginsLv2TriggerStereossrc)
//
// Signal source

func (e *LspPlugInPluginsLv2TriggerStereo) GetSsrc() (interface{}, error) {
	return e.Element.GetProperty("ssrc")
}

func (e *LspPlugInPluginsLv2TriggerStereo) SetSsrc(value interface{}) error {
	return e.Element.SetProperty("ssrc", value)
}

// tfm (float32)
//
// Trigger function meter

func (e *LspPlugInPluginsLv2TriggerStereo) GetTfm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) GetTfv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetTfv(value bool) error {
	return e.Element.SetProperty("tfv", value)
}

// tla (bool)
//
// Trigger activity

func (e *LspPlugInPluginsLv2TriggerStereo) GetTla() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) GetTlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) GetTlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetTlv(value bool) error {
	return e.Element.SetProperty("tlv", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2TriggerStereo) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2TriggerStereoasel string

const (
	LspPlugInPluginsLv2TriggerStereoaselTrigger LspPlugInPluginsLv2TriggerStereoasel = "Trigger" // Trigger (0) – Trigger
	LspPlugInPluginsLv2TriggerStereoaselInstrument LspPlugInPluginsLv2TriggerStereoasel = "Instrument" // Instrument (1) – Instrument
)

type LspPlugInPluginsLv2TriggerStereomode string

const (
	LspPlugInPluginsLv2TriggerStereomodePeak LspPlugInPluginsLv2TriggerStereomode = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2TriggerStereomodeRms LspPlugInPluginsLv2TriggerStereomode = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2TriggerStereomodeLowPass LspPlugInPluginsLv2TriggerStereomode = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2TriggerStereomodeUniform LspPlugInPluginsLv2TriggerStereomode = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2TriggerStereossrc string

const (
	LspPlugInPluginsLv2TriggerStereossrcMiddle LspPlugInPluginsLv2TriggerStereossrc = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2TriggerStereossrcSide LspPlugInPluginsLv2TriggerStereossrc = "Side" // Side (1) – Side
	LspPlugInPluginsLv2TriggerStereossrcLeft LspPlugInPluginsLv2TriggerStereossrc = "Left" // Left (2) – Left
	LspPlugInPluginsLv2TriggerStereossrcRight LspPlugInPluginsLv2TriggerStereossrc = "Right" // Right (3) – Right
)

