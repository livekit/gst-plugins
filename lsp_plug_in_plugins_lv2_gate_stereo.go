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

type LspPlugInPluginsLv2GateStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GateStereo() (*LspPlugInPluginsLv2GateStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-gate-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GateStereoWithName(name string) (*LspPlugInPluginsLv2GateStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-gate-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// at (float32)
//
// Attack

func (e *LspPlugInPluginsLv2GateStereo) GetAt() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GateStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2GateStereo) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2GateStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2GateStereo) GetClm() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2GateStereo) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2GateStereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GateStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gh (bool)
//
// Hysteresis

func (e *LspPlugInPluginsLv2GateStereo) GetGh() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGh(value bool) error {
	return e.Element.SetProperty("gh", value)
}

// gr (float32)
//
// Reduction

func (e *LspPlugInPluginsLv2GateStereo) GetGr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGr(value float32) error {
	return e.Element.SetProperty("gr", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2GateStereo) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// gt (float32)
//
// Threshold

func (e *LspPlugInPluginsLv2GateStereo) GetGt() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGt(value float32) error {
	return e.Element.SetProperty("gt", value)
}

// gz (float32)
//
// Zone size

func (e *LspPlugInPluginsLv2GateStereo) GetGz() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetGz(value float32) error {
	return e.Element.SetProperty("gz", value)
}

// gzs (float32)
//
// Zone start

func (e *LspPlugInPluginsLv2GateStereo) GetGzs() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetHt() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetHt(value float32) error {
	return e.Element.SetProperty("ht", value)
}

// hts (float32)
//
// Hysteresis threshold start

func (e *LspPlugInPluginsLv2GateStereo) GetHts() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetHz() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetHz(value float32) error {
	return e.Element.SetProperty("hz", value)
}

// hzs (float32)
//
// Hysteresis zone start

func (e *LspPlugInPluginsLv2GateStereo) GetHzs() (float32, error) {
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

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2GateStereo) GetIlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm-r (float32)
//
// Input level meter Right

func (e *LspPlugInPluginsLv2GateStereo) GetIlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv-l (bool)
//
// Input level visibility Left

func (e *LspPlugInPluginsLv2GateStereo) GetIlvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateStereo) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2GateStereo) GetIlvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateStereo) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2GateStereo) GetMk() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2GateStereo) GetOlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olm-r (float32)
//
// Output level meter Right

func (e *LspPlugInPluginsLv2GateStereo) GetOlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olv-l (bool)
//
// Output level visibility Left

func (e *LspPlugInPluginsLv2GateStereo) GetOlvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateStereo) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2GateStereo) GetOlvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateStereo) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GateStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rlm (float32)
//
// Reduction level meter

func (e *LspPlugInPluginsLv2GateStereo) GetRlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2GateStereo) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2GateStereoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2GateStereo) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2GateStereo) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2GateStereo) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2GateStereo) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// scs (GstLspPlugInPluginsLv2GateStereoscs)
//
// Sidechain source

func (e *LspPlugInPluginsLv2GateStereo) GetScs() (interface{}, error) {
	return e.Element.GetProperty("scs")
}

func (e *LspPlugInPluginsLv2GateStereo) SetScs(value interface{}) error {
	return e.Element.SetProperty("scs", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2GateStereo) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2GateStereo) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateStereo) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GateStereoscm string

const (
	LspPlugInPluginsLv2GateStereoscmPeak LspPlugInPluginsLv2GateStereoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2GateStereoscmRms LspPlugInPluginsLv2GateStereoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2GateStereoscmLowPass LspPlugInPluginsLv2GateStereoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2GateStereoscmUniform LspPlugInPluginsLv2GateStereoscm = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2GateStereoscs string

const (
	LspPlugInPluginsLv2GateStereoscsMiddle LspPlugInPluginsLv2GateStereoscs = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2GateStereoscsSide LspPlugInPluginsLv2GateStereoscs = "Side" // Side (1) – Side
	LspPlugInPluginsLv2GateStereoscsLeft LspPlugInPluginsLv2GateStereoscs = "Left" // Left (2) – Left
	LspPlugInPluginsLv2GateStereoscsRight LspPlugInPluginsLv2GateStereoscs = "Right" // Right (3) – Right
)

