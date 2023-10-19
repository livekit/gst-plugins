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

type LspPlugInPluginsLv2CompressorStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2CompressorStereo() (*LspPlugInPluginsLv2CompressorStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-compressor-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2CompressorStereoWithName(name string) (*LspPlugInPluginsLv2CompressorStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-compressor-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al (float32)
//
// Attack level

func (e *LspPlugInPluginsLv2CompressorStereo) GetAl() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetAl(value float32) error {
	return e.Element.SetProperty("al", value)
}

// at (float32)
//
// Attack time

func (e *LspPlugInPluginsLv2CompressorStereo) GetAt() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2CompressorStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2CompressorStereo) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2CompressorStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2CompressorStereo) GetClm() (float32, error) {
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

// cm (GstLspPlugInPluginsLv2CompressorStereocm)
//
// Compression mode

func (e *LspPlugInPluginsLv2CompressorStereo) GetCm() (interface{}, error) {
	return e.Element.GetProperty("cm")
}

func (e *LspPlugInPluginsLv2CompressorStereo) SetCm(value interface{}) error {
	return e.Element.SetProperty("cm", value)
}

// cr (float32)
//
// Ratio

func (e *LspPlugInPluginsLv2CompressorStereo) GetCr() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetCr(value float32) error {
	return e.Element.SetProperty("cr", value)
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2CompressorStereo) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2CompressorStereo) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2CompressorStereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2CompressorStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2CompressorStereo) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2CompressorStereo) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetIlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetIlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2CompressorStereo) GetIlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// kn (float32)
//
// Knee

func (e *LspPlugInPluginsLv2CompressorStereo) GetKn() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetKn(value float32) error {
	return e.Element.SetProperty("kn", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2CompressorStereo) GetMk() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2CompressorStereo) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetOlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2CompressorStereo) GetOlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2CompressorStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2CompressorStereo) GetRl() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetRlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2CompressorStereo) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2CompressorStereo) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2CompressorStereoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2CompressorStereo) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2CompressorStereo) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2CompressorStereo) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2CompressorStereo) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// scs (GstLspPlugInPluginsLv2CompressorStereoscs)
//
// Sidechain source

func (e *LspPlugInPluginsLv2CompressorStereo) GetScs() (interface{}, error) {
	return e.Element.GetProperty("scs")
}

func (e *LspPlugInPluginsLv2CompressorStereo) SetScs(value interface{}) error {
	return e.Element.SetProperty("scs", value)
}

// sct (GstLspPlugInPluginsLv2CompressorStereosct)
//
// Sidechain type

func (e *LspPlugInPluginsLv2CompressorStereo) GetSct() (interface{}, error) {
	return e.Element.GetProperty("sct")
}

func (e *LspPlugInPluginsLv2CompressorStereo) SetSct(value interface{}) error {
	return e.Element.SetProperty("sct", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2CompressorStereo) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2CompressorStereo) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorStereo) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2CompressorStereoscs string

const (
	LspPlugInPluginsLv2CompressorStereoscsMiddle LspPlugInPluginsLv2CompressorStereoscs = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2CompressorStereoscsSide LspPlugInPluginsLv2CompressorStereoscs = "Side" // Side (1) – Side
	LspPlugInPluginsLv2CompressorStereoscsLeft LspPlugInPluginsLv2CompressorStereoscs = "Left" // Left (2) – Left
	LspPlugInPluginsLv2CompressorStereoscsRight LspPlugInPluginsLv2CompressorStereoscs = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2CompressorStereosct string

const (
	LspPlugInPluginsLv2CompressorStereosctFeedForward LspPlugInPluginsLv2CompressorStereosct = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2CompressorStereosctFeedBack LspPlugInPluginsLv2CompressorStereosct = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2CompressorStereocm string

const (
	LspPlugInPluginsLv2CompressorStereocmDownward LspPlugInPluginsLv2CompressorStereocm = "Downward" // Downward (0) – Downward
	LspPlugInPluginsLv2CompressorStereocmUpward LspPlugInPluginsLv2CompressorStereocm = "Upward" // Upward (1) – Upward
)

type LspPlugInPluginsLv2CompressorStereoscm string

const (
	LspPlugInPluginsLv2CompressorStereoscmPeak LspPlugInPluginsLv2CompressorStereoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2CompressorStereoscmRms LspPlugInPluginsLv2CompressorStereoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2CompressorStereoscmLowPass LspPlugInPluginsLv2CompressorStereoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2CompressorStereoscmUniform LspPlugInPluginsLv2CompressorStereoscm = "Uniform" // Uniform (3) – Uniform
)

