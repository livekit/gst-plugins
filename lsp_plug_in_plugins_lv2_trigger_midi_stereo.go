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

type LspPlugInPluginsLv2TriggerMidiStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2TriggerMidiStereo() (*LspPlugInPluginsLv2TriggerMidiStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-trigger-midi-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerMidiStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2TriggerMidiStereoWithName(name string) (*LspPlugInPluginsLv2TriggerMidiStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-trigger-midi-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerMidiStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// asel (GstLspPlugInPluginsLv2TriggerMidiStereoasel)
//
// Area selector

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetAsel() (interface{}, error) {
	return e.Element.GetProperty("asel")
}

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetAsel(value interface{}) error {
	return e.Element.SetProperty("asel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// chan (GstLspPlugInPluginsLv2TriggerMidiStereochan)
//
// Channel

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetChan() (interface{}, error) {
	return e.Element.GetProperty("chan")
}

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetChan(value interface{}) error {
	return e.Element.SetProperty("chan", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// dl (float32)
//
// Detect level

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetDl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetDl(value float32) error {
	return e.Element.SetProperty("dl", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// dt (float32)
//
// Detect time

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetDt() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetDt(value float32) error {
	return e.Element.SetProperty("dt", value)
}

// dtr1 (float32)
//
// Dynamics range 1

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetDtr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetDtr1(value float32) error {
	return e.Element.SetProperty("dtr1", value)
}

// dtr2 (float32)
//
// Dynamics range 2

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetDtr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetDtr2(value float32) error {
	return e.Element.SetProperty("dtr2", value)
}

// dyna (float32)
//
// Dynamics

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetDyna() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetDyna(value float32) error {
	return e.Element.SetProperty("dyna", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// isml (float32)
//
// Input signal meter left

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetIsml() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetIsmr() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetIsvl() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetIsvl(value bool) error {
	return e.Element.SetProperty("isvl", value)
}

// isvr (bool)
//
// Input signal right display

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetIsvr() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetIsvr(value bool) error {
	return e.Element.SetProperty("isvr", value)
}

// lstn (bool)
//
// Trigger listen

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetLstn() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mn (int)
//
// MIDI Note #

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetMn() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mn")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// mode (GstLspPlugInPluginsLv2TriggerMidiStereomode)
//
// Detection mode

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// note (GstLspPlugInPluginsLv2TriggerMidiStereonote)
//
// Note

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetNote() (interface{}, error) {
	return e.Element.GetProperty("note")
}

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetNote(value interface{}) error {
	return e.Element.SetProperty("note", value)
}

// oct (GstLspPlugInPluginsLv2TriggerMidiStereooct)
//
// Octave

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetOct() (interface{}, error) {
	return e.Element.GetProperty("oct")
}

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetOct(value interface{}) error {
	return e.Element.SetProperty("oct", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// preamp (float32)
//
// Signal pre-amplification

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetPreamp() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetPreamp(value float32) error {
	return e.Element.SetProperty("preamp", value)
}

// react (float32)
//
// Reactivity

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetRl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// ssrc (GstLspPlugInPluginsLv2TriggerMidiStereossrc)
//
// Signal source

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetSsrc() (interface{}, error) {
	return e.Element.GetProperty("ssrc")
}

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetSsrc(value interface{}) error {
	return e.Element.SetProperty("ssrc", value)
}

// tfm (float32)
//
// Trigger function meter

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetTfm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetTfv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetTfv(value bool) error {
	return e.Element.SetProperty("tfv", value)
}

// tla (bool)
//
// Trigger activity

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetTla() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetTlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetTlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetTlv(value bool) error {
	return e.Element.SetProperty("tlv", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2TriggerMidiStereo) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2TriggerMidiStereossrc string

const (
	LspPlugInPluginsLv2TriggerMidiStereossrcMiddle LspPlugInPluginsLv2TriggerMidiStereossrc = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2TriggerMidiStereossrcSide LspPlugInPluginsLv2TriggerMidiStereossrc = "Side" // Side (1) – Side
	LspPlugInPluginsLv2TriggerMidiStereossrcLeft LspPlugInPluginsLv2TriggerMidiStereossrc = "Left" // Left (2) – Left
	LspPlugInPluginsLv2TriggerMidiStereossrcRight LspPlugInPluginsLv2TriggerMidiStereossrc = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2TriggerMidiStereoasel string

const (
	LspPlugInPluginsLv2TriggerMidiStereoaselTrigger LspPlugInPluginsLv2TriggerMidiStereoasel = "Trigger" // Trigger (0) – Trigger
	LspPlugInPluginsLv2TriggerMidiStereoaselInstrument LspPlugInPluginsLv2TriggerMidiStereoasel = "Instrument" // Instrument (1) – Instrument
)

type LspPlugInPluginsLv2TriggerMidiStereochan string

const (
	LspPlugInPluginsLv2TriggerMidiStereochan01 LspPlugInPluginsLv2TriggerMidiStereochan = "01" // 01 (0) – 01
	LspPlugInPluginsLv2TriggerMidiStereochan02 LspPlugInPluginsLv2TriggerMidiStereochan = "02" // 02 (1) – 02
	LspPlugInPluginsLv2TriggerMidiStereochan03 LspPlugInPluginsLv2TriggerMidiStereochan = "03" // 03 (2) – 03
	LspPlugInPluginsLv2TriggerMidiStereochan04 LspPlugInPluginsLv2TriggerMidiStereochan = "04" // 04 (3) – 04
	LspPlugInPluginsLv2TriggerMidiStereochan05 LspPlugInPluginsLv2TriggerMidiStereochan = "05" // 05 (4) – 05
	LspPlugInPluginsLv2TriggerMidiStereochan06 LspPlugInPluginsLv2TriggerMidiStereochan = "06" // 06 (5) – 06
	LspPlugInPluginsLv2TriggerMidiStereochan07 LspPlugInPluginsLv2TriggerMidiStereochan = "07" // 07 (6) – 07
	LspPlugInPluginsLv2TriggerMidiStereochan08 LspPlugInPluginsLv2TriggerMidiStereochan = "08" // 08 (7) – 08
	LspPlugInPluginsLv2TriggerMidiStereochan09 LspPlugInPluginsLv2TriggerMidiStereochan = "09" // 09 (8) – 09
	LspPlugInPluginsLv2TriggerMidiStereochan10 LspPlugInPluginsLv2TriggerMidiStereochan = "10" // 10 (9) – 10
	LspPlugInPluginsLv2TriggerMidiStereochan11 LspPlugInPluginsLv2TriggerMidiStereochan = "11" // 11 (10) – 11
	LspPlugInPluginsLv2TriggerMidiStereochan12 LspPlugInPluginsLv2TriggerMidiStereochan = "12" // 12 (11) – 12
	LspPlugInPluginsLv2TriggerMidiStereochan13 LspPlugInPluginsLv2TriggerMidiStereochan = "13" // 13 (12) – 13
	LspPlugInPluginsLv2TriggerMidiStereochan14 LspPlugInPluginsLv2TriggerMidiStereochan = "14" // 14 (13) – 14
	LspPlugInPluginsLv2TriggerMidiStereochan15 LspPlugInPluginsLv2TriggerMidiStereochan = "15" // 15 (14) – 15
	LspPlugInPluginsLv2TriggerMidiStereochan16 LspPlugInPluginsLv2TriggerMidiStereochan = "16" // 16 (15) – 16
)

type LspPlugInPluginsLv2TriggerMidiStereomode string

const (
	LspPlugInPluginsLv2TriggerMidiStereomodePeak LspPlugInPluginsLv2TriggerMidiStereomode = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2TriggerMidiStereomodeRms LspPlugInPluginsLv2TriggerMidiStereomode = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2TriggerMidiStereomodeLowPass LspPlugInPluginsLv2TriggerMidiStereomode = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2TriggerMidiStereomodeUniform LspPlugInPluginsLv2TriggerMidiStereomode = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2TriggerMidiStereonote string

const (
	LspPlugInPluginsLv2TriggerMidiStereonoteC LspPlugInPluginsLv2TriggerMidiStereonote = "C" // C (0) – C
	LspPlugInPluginsLv2TriggerMidiStereonoteC1 LspPlugInPluginsLv2TriggerMidiStereonote = "C#" // C# (1) – C#
	LspPlugInPluginsLv2TriggerMidiStereonoteD LspPlugInPluginsLv2TriggerMidiStereonote = "D" // D (2) – D
	LspPlugInPluginsLv2TriggerMidiStereonoteD1 LspPlugInPluginsLv2TriggerMidiStereonote = "D#" // D# (3) – D#
	LspPlugInPluginsLv2TriggerMidiStereonoteE LspPlugInPluginsLv2TriggerMidiStereonote = "E" // E (4) – E
	LspPlugInPluginsLv2TriggerMidiStereonoteF LspPlugInPluginsLv2TriggerMidiStereonote = "F" // F (5) – F
	LspPlugInPluginsLv2TriggerMidiStereonoteF1 LspPlugInPluginsLv2TriggerMidiStereonote = "F#" // F# (6) – F#
	LspPlugInPluginsLv2TriggerMidiStereonoteG LspPlugInPluginsLv2TriggerMidiStereonote = "G" // G (7) – G
	LspPlugInPluginsLv2TriggerMidiStereonoteG1 LspPlugInPluginsLv2TriggerMidiStereonote = "G#" // G# (8) – G#
	LspPlugInPluginsLv2TriggerMidiStereonoteA LspPlugInPluginsLv2TriggerMidiStereonote = "A" // A (9) – A
	LspPlugInPluginsLv2TriggerMidiStereonoteA1 LspPlugInPluginsLv2TriggerMidiStereonote = "A#" // A# (10) – A#
	LspPlugInPluginsLv2TriggerMidiStereonoteB LspPlugInPluginsLv2TriggerMidiStereonote = "B" // B (11) – B
)

type LspPlugInPluginsLv2TriggerMidiStereooct string

const (
	LspPlugInPluginsLv2TriggerMidiStereooct2 LspPlugInPluginsLv2TriggerMidiStereooct = "-2" // -2 (0) – -2
	LspPlugInPluginsLv2TriggerMidiStereooct1 LspPlugInPluginsLv2TriggerMidiStereooct = "-1" // -1 (1) – -1
	LspPlugInPluginsLv2TriggerMidiStereooct0 LspPlugInPluginsLv2TriggerMidiStereooct = "0" // 0 (2) – 0
	LspPlugInPluginsLv2TriggerMidiStereooct11 LspPlugInPluginsLv2TriggerMidiStereooct = "1" // 1 (3) – 1
	LspPlugInPluginsLv2TriggerMidiStereooct21 LspPlugInPluginsLv2TriggerMidiStereooct = "2" // 2 (4) – 2
	LspPlugInPluginsLv2TriggerMidiStereooct3 LspPlugInPluginsLv2TriggerMidiStereooct = "3" // 3 (5) – 3
	LspPlugInPluginsLv2TriggerMidiStereooct4 LspPlugInPluginsLv2TriggerMidiStereooct = "4" // 4 (6) – 4
	LspPlugInPluginsLv2TriggerMidiStereooct5 LspPlugInPluginsLv2TriggerMidiStereooct = "5" // 5 (7) – 5
	LspPlugInPluginsLv2TriggerMidiStereooct6 LspPlugInPluginsLv2TriggerMidiStereooct = "6" // 6 (8) – 6
	LspPlugInPluginsLv2TriggerMidiStereooct7 LspPlugInPluginsLv2TriggerMidiStereooct = "7" // 7 (9) – 7
	LspPlugInPluginsLv2TriggerMidiStereooct8 LspPlugInPluginsLv2TriggerMidiStereooct = "8" // 8 (10) – 8
)

