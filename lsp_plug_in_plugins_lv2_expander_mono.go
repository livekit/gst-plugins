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

type LspPlugInPluginsLv2ExpanderMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ExpanderMono() (*LspPlugInPluginsLv2ExpanderMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-expander-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ExpanderMonoWithName(name string) (*LspPlugInPluginsLv2ExpanderMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-expander-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al (float32)
//
// Attack level

func (e *LspPlugInPluginsLv2ExpanderMono) GetAl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderMono) SetAl(value float32) error {
	return e.Element.SetProperty("al", value)
}

// at (float32)
//
// Attack time

func (e *LspPlugInPluginsLv2ExpanderMono) GetAt() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ExpanderMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2ExpanderMono) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2ExpanderMono) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2ExpanderMono) GetClm() (float32, error) {
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

// cr (float32)
//
// Ratio

func (e *LspPlugInPluginsLv2ExpanderMono) GetCr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderMono) SetCr(value float32) error {
	return e.Element.SetProperty("cr", value)
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2ExpanderMono) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2ExpanderMono) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// em (GstLspPlugInPluginsLv2ExpanderMonoem)
//
// Expander mode

func (e *LspPlugInPluginsLv2ExpanderMono) GetEm() (interface{}, error) {
	return e.Element.GetProperty("em")
}

func (e *LspPlugInPluginsLv2ExpanderMono) SetEm(value interface{}) error {
	return e.Element.SetProperty("em", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ExpanderMono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ExpanderMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2ExpanderMono) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// ilm (float32)
//
// Input level meter

func (e *LspPlugInPluginsLv2ExpanderMono) GetIlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) GetIlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetIlv(value bool) error {
	return e.Element.SetProperty("ilv", value)
}

// kn (float32)
//
// Knee

func (e *LspPlugInPluginsLv2ExpanderMono) GetKn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderMono) SetKn(value float32) error {
	return e.Element.SetProperty("kn", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2ExpanderMono) GetMk() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// olm (float32)
//
// Output level meter

func (e *LspPlugInPluginsLv2ExpanderMono) GetOlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) GetOlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetOlv(value bool) error {
	return e.Element.SetProperty("olv", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ExpanderMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2ExpanderMono) GetRl() (float32, error) {
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

// rlm (float32)
//
// Reduction level meter

func (e *LspPlugInPluginsLv2ExpanderMono) GetRlm() (float32, error) {
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

// rrl (float32)
//
// Relative release level

func (e *LspPlugInPluginsLv2ExpanderMono) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2ExpanderMono) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2ExpanderMono) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2ExpanderMonoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2ExpanderMono) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2ExpanderMono) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2ExpanderMono) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2ExpanderMono) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2ExpanderMono) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2ExpanderMono) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMono) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ExpanderMonoem string

const (
	LspPlugInPluginsLv2ExpanderMonoemDown LspPlugInPluginsLv2ExpanderMonoem = "Down" // Down (0) – Down
	LspPlugInPluginsLv2ExpanderMonoemUp LspPlugInPluginsLv2ExpanderMonoem = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2ExpanderMonoscm string

const (
	LspPlugInPluginsLv2ExpanderMonoscmPeak LspPlugInPluginsLv2ExpanderMonoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2ExpanderMonoscmRms LspPlugInPluginsLv2ExpanderMonoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2ExpanderMonoscmLowPass LspPlugInPluginsLv2ExpanderMonoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2ExpanderMonoscmUniform LspPlugInPluginsLv2ExpanderMonoscm = "Uniform" // Uniform (3) – Uniform
)

