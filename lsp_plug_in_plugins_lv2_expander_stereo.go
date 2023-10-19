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

type LspPlugInPluginsLv2ExpanderStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ExpanderStereo() (*LspPlugInPluginsLv2ExpanderStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-expander-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ExpanderStereoWithName(name string) (*LspPlugInPluginsLv2ExpanderStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-expander-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al (float32)
//
// Attack level

func (e *LspPlugInPluginsLv2ExpanderStereo) GetAl() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetAl(value float32) error {
	return e.Element.SetProperty("al", value)
}

// at (float32)
//
// Attack time

func (e *LspPlugInPluginsLv2ExpanderStereo) GetAt() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ExpanderStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2ExpanderStereo) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2ExpanderStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2ExpanderStereo) GetClm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetCr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetCr(value float32) error {
	return e.Element.SetProperty("cr", value)
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2ExpanderStereo) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2ExpanderStereo) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// em (GstLspPlugInPluginsLv2ExpanderStereoem)
//
// Expander mode

func (e *LspPlugInPluginsLv2ExpanderStereo) GetEm() (interface{}, error) {
	return e.Element.GetProperty("em")
}

func (e *LspPlugInPluginsLv2ExpanderStereo) SetEm(value interface{}) error {
	return e.Element.SetProperty("em", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ExpanderStereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ExpanderStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2ExpanderStereo) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2ExpanderStereo) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetIlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetIlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2ExpanderStereo) GetIlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// kn (float32)
//
// Knee

func (e *LspPlugInPluginsLv2ExpanderStereo) GetKn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetKn(value float32) error {
	return e.Element.SetProperty("kn", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2ExpanderStereo) GetMk() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2ExpanderStereo) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetOlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2ExpanderStereo) GetOlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ExpanderStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2ExpanderStereo) GetRl() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetRlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2ExpanderStereo) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2ExpanderStereo) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2ExpanderStereoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2ExpanderStereo) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2ExpanderStereo) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2ExpanderStereo) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2ExpanderStereo) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// scs (GstLspPlugInPluginsLv2ExpanderStereoscs)
//
// Sidechain source

func (e *LspPlugInPluginsLv2ExpanderStereo) GetScs() (interface{}, error) {
	return e.Element.GetProperty("scs")
}

func (e *LspPlugInPluginsLv2ExpanderStereo) SetScs(value interface{}) error {
	return e.Element.SetProperty("scs", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2ExpanderStereo) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2ExpanderStereo) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderStereo) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ExpanderStereoem string

const (
	LspPlugInPluginsLv2ExpanderStereoemDown LspPlugInPluginsLv2ExpanderStereoem = "Down" // Down (0) – Down
	LspPlugInPluginsLv2ExpanderStereoemUp LspPlugInPluginsLv2ExpanderStereoem = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2ExpanderStereoscm string

const (
	LspPlugInPluginsLv2ExpanderStereoscmPeak LspPlugInPluginsLv2ExpanderStereoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2ExpanderStereoscmRms LspPlugInPluginsLv2ExpanderStereoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2ExpanderStereoscmLowPass LspPlugInPluginsLv2ExpanderStereoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2ExpanderStereoscmUniform LspPlugInPluginsLv2ExpanderStereoscm = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2ExpanderStereoscs string

const (
	LspPlugInPluginsLv2ExpanderStereoscsMiddle LspPlugInPluginsLv2ExpanderStereoscs = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2ExpanderStereoscsSide LspPlugInPluginsLv2ExpanderStereoscs = "Side" // Side (1) – Side
	LspPlugInPluginsLv2ExpanderStereoscsLeft LspPlugInPluginsLv2ExpanderStereoscs = "Left" // Left (2) – Left
	LspPlugInPluginsLv2ExpanderStereoscsRight LspPlugInPluginsLv2ExpanderStereoscs = "Right" // Right (3) – Right
)

