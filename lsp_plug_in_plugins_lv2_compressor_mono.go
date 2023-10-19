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

type LspPlugInPluginsLv2CompressorMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2CompressorMono() (*LspPlugInPluginsLv2CompressorMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-compressor-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2CompressorMonoWithName(name string) (*LspPlugInPluginsLv2CompressorMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-compressor-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al (float32)
//
// Attack level

func (e *LspPlugInPluginsLv2CompressorMono) GetAl() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetAl(value float32) error {
	return e.Element.SetProperty("al", value)
}

// at (float32)
//
// Attack time

func (e *LspPlugInPluginsLv2CompressorMono) GetAt() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2CompressorMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2CompressorMono) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2CompressorMono) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2CompressorMono) GetClm() (float32, error) {
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

// cm (GstLspPlugInPluginsLv2CompressorMonocm)
//
// Compression mode

func (e *LspPlugInPluginsLv2CompressorMono) GetCm() (interface{}, error) {
	return e.Element.GetProperty("cm")
}

func (e *LspPlugInPluginsLv2CompressorMono) SetCm(value interface{}) error {
	return e.Element.SetProperty("cm", value)
}

// cr (float32)
//
// Ratio

func (e *LspPlugInPluginsLv2CompressorMono) GetCr() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetCr(value float32) error {
	return e.Element.SetProperty("cr", value)
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2CompressorMono) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2CompressorMono) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2CompressorMono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2CompressorMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2CompressorMono) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// ilm (float32)
//
// Input level meter

func (e *LspPlugInPluginsLv2CompressorMono) GetIlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetIlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetIlv(value bool) error {
	return e.Element.SetProperty("ilv", value)
}

// kn (float32)
//
// Knee

func (e *LspPlugInPluginsLv2CompressorMono) GetKn() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetKn(value float32) error {
	return e.Element.SetProperty("kn", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2CompressorMono) GetMk() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// olm (float32)
//
// Output level meter

func (e *LspPlugInPluginsLv2CompressorMono) GetOlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetOlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetOlv(value bool) error {
	return e.Element.SetProperty("olv", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2CompressorMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2CompressorMono) GetRl() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetRlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2CompressorMono) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2CompressorMono) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2CompressorMonoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2CompressorMono) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2CompressorMono) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2CompressorMono) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2CompressorMono) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// sct (GstLspPlugInPluginsLv2CompressorMonosct)
//
// Sidechain type

func (e *LspPlugInPluginsLv2CompressorMono) GetSct() (interface{}, error) {
	return e.Element.GetProperty("sct")
}

func (e *LspPlugInPluginsLv2CompressorMono) SetSct(value interface{}) error {
	return e.Element.SetProperty("sct", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2CompressorMono) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2CompressorMono) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMono) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2CompressorMonocm string

const (
	LspPlugInPluginsLv2CompressorMonocmDownward LspPlugInPluginsLv2CompressorMonocm = "Downward" // Downward (0) – Downward
	LspPlugInPluginsLv2CompressorMonocmUpward LspPlugInPluginsLv2CompressorMonocm = "Upward" // Upward (1) – Upward
)

type LspPlugInPluginsLv2CompressorMonoscm string

const (
	LspPlugInPluginsLv2CompressorMonoscmPeak LspPlugInPluginsLv2CompressorMonoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2CompressorMonoscmRms LspPlugInPluginsLv2CompressorMonoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2CompressorMonoscmLowPass LspPlugInPluginsLv2CompressorMonoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2CompressorMonoscmUniform LspPlugInPluginsLv2CompressorMonoscm = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2CompressorMonosct string

const (
	LspPlugInPluginsLv2CompressorMonosctFeedForward LspPlugInPluginsLv2CompressorMonosct = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2CompressorMonosctFeedBack LspPlugInPluginsLv2CompressorMonosct = "Feed-back" // Feed-back (1) – Feed-back
)

