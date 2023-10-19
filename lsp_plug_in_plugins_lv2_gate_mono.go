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

type LspPlugInPluginsLv2GateMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GateMono() (*LspPlugInPluginsLv2GateMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-gate-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GateMonoWithName(name string) (*LspPlugInPluginsLv2GateMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-gate-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// at (float32)
//
// Attack

func (e *LspPlugInPluginsLv2GateMono) GetAt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GateMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2GateMono) GetCdr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2GateMono) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2GateMono) GetClm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2GateMono) GetCwt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2GateMono) GetElm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elv (bool)
//
// Envelope level visibility

func (e *LspPlugInPluginsLv2GateMono) GetElv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2GateMono) GetGIn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-in")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GateMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gh (bool)
//
// Hysteresis

func (e *LspPlugInPluginsLv2GateMono) GetGh() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gh")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetGh(value bool) error {
	return e.Element.SetProperty("gh", value)
}

// gr (float32)
//
// Reduction

func (e *LspPlugInPluginsLv2GateMono) GetGr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetGr(value float32) error {
	return e.Element.SetProperty("gr", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2GateMono) GetGrv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// gt (float32)
//
// Threshold

func (e *LspPlugInPluginsLv2GateMono) GetGt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetGt(value float32) error {
	return e.Element.SetProperty("gt", value)
}

// gz (float32)
//
// Zone size

func (e *LspPlugInPluginsLv2GateMono) GetGz() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gz")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetGz(value float32) error {
	return e.Element.SetProperty("gz", value)
}

// gzs (float32)
//
// Zone start

func (e *LspPlugInPluginsLv2GateMono) GetGzs() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gzs")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ht (float32)
//
// Hysteresis threshold

func (e *LspPlugInPluginsLv2GateMono) GetHt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ht")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetHt(value float32) error {
	return e.Element.SetProperty("ht", value)
}

// hts (float32)
//
// Hysteresis threshold start

func (e *LspPlugInPluginsLv2GateMono) GetHts() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hts")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hz (float32)
//
// Hysteresis zone size

func (e *LspPlugInPluginsLv2GateMono) GetHz() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hz")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetHz(value float32) error {
	return e.Element.SetProperty("hz", value)
}

// hzs (float32)
//
// Hysteresis zone start

func (e *LspPlugInPluginsLv2GateMono) GetHzs() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hzs")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm (float32)
//
// Input level meter

func (e *LspPlugInPluginsLv2GateMono) GetIlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv (bool)
//
// Input level visibility

func (e *LspPlugInPluginsLv2GateMono) GetIlv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetIlv(value bool) error {
	return e.Element.SetProperty("ilv", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2GateMono) GetMk() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// olm (float32)
//
// Output level meter

func (e *LspPlugInPluginsLv2GateMono) GetOlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olv (bool)
//
// Output level visibility

func (e *LspPlugInPluginsLv2GateMono) GetOlv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetOlv(value bool) error {
	return e.Element.SetProperty("olv", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GateMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GateMono) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rlm (float32)
//
// Reduction level meter

func (e *LspPlugInPluginsLv2GateMono) GetRlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rt (float32)
//
// Release

func (e *LspPlugInPluginsLv2GateMono) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMono) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2GateMono) GetScl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2GateMonoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2GateMono) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2GateMono) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2GateMono) GetScp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2GateMono) GetScr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2GateMono) GetSla() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2GateMono) GetSlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slv (bool)
//
// Sidechain level visibility

func (e *LspPlugInPluginsLv2GateMono) GetSlv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMono) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GateMonoscm string

const (
	LspPlugInPluginsLv2GateMonoscmPeak LspPlugInPluginsLv2GateMonoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2GateMonoscmRms LspPlugInPluginsLv2GateMonoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2GateMonoscmLowPass LspPlugInPluginsLv2GateMonoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2GateMonoscmUniform LspPlugInPluginsLv2GateMonoscm = "Uniform" // Uniform (3) – Uniform
)

