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

type LspPlugInPluginsLv2TriggerMidiMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2TriggerMidiMono() (*LspPlugInPluginsLv2TriggerMidiMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-trigger-midi-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerMidiMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2TriggerMidiMonoWithName(name string) (*LspPlugInPluginsLv2TriggerMidiMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-trigger-midi-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2TriggerMidiMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// asel (GstLspPlugInPluginsLv2TriggerMidiMonoasel)
//
// Area selector

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetAsel() (interface{}, error) {
	return e.Element.GetProperty("asel")
}

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetAsel(value interface{}) error {
	return e.Element.SetProperty("asel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// chan (GstLspPlugInPluginsLv2TriggerMidiMonochan)
//
// Channel

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetChan() (interface{}, error) {
	return e.Element.GetProperty("chan")
}

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetChan(value interface{}) error {
	return e.Element.SetProperty("chan", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// dl (float32)
//
// Detect level

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetDl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetDl(value float32) error {
	return e.Element.SetProperty("dl", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// dt (float32)
//
// Detect time

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetDt() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetDt(value float32) error {
	return e.Element.SetProperty("dt", value)
}

// dtr1 (float32)
//
// Dynamics range 1

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetDtr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetDtr1(value float32) error {
	return e.Element.SetProperty("dtr1", value)
}

// dtr2 (float32)
//
// Dynamics range 2

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetDtr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetDtr2(value float32) error {
	return e.Element.SetProperty("dtr2", value)
}

// dyna (float32)
//
// Dynamics

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetDyna() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetDyna(value float32) error {
	return e.Element.SetProperty("dyna", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// ism (float32)
//
// Input signal meter

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetIsm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetIsv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetIsv(value bool) error {
	return e.Element.SetProperty("isv", value)
}

// lstn (bool)
//
// Trigger listen

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetLstn() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mn (int)
//
// MIDI Note #

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetMn() (int, error) {
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

// mode (GstLspPlugInPluginsLv2TriggerMidiMonomode)
//
// Detection mode

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// note (GstLspPlugInPluginsLv2TriggerMidiMononote)
//
// Note

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetNote() (interface{}, error) {
	return e.Element.GetProperty("note")
}

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetNote(value interface{}) error {
	return e.Element.SetProperty("note", value)
}

// oct (GstLspPlugInPluginsLv2TriggerMidiMonooct)
//
// Octave

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetOct() (interface{}, error) {
	return e.Element.GetProperty("oct")
}

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetOct(value interface{}) error {
	return e.Element.SetProperty("oct", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// preamp (float32)
//
// Signal pre-amplification

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetPreamp() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetPreamp(value float32) error {
	return e.Element.SetProperty("preamp", value)
}

// react (float32)
//
// Reactivity

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl (float32)
//
// Release level

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetRl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetRrl() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetRrl(value float32) error {
	return e.Element.SetProperty("rrl", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// tfm (float32)
//
// Trigger function meter

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetTfm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetTfv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetTfv(value bool) error {
	return e.Element.SetProperty("tfv", value)
}

// tla (bool)
//
// Trigger activity

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetTla() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetTlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetTlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetTlv(value bool) error {
	return e.Element.SetProperty("tlv", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2TriggerMidiMono) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2TriggerMidiMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2TriggerMidiMonoasel string

const (
	LspPlugInPluginsLv2TriggerMidiMonoaselTrigger LspPlugInPluginsLv2TriggerMidiMonoasel = "Trigger" // Trigger (0) – Trigger
	LspPlugInPluginsLv2TriggerMidiMonoaselInstrument LspPlugInPluginsLv2TriggerMidiMonoasel = "Instrument" // Instrument (1) – Instrument
)

type LspPlugInPluginsLv2TriggerMidiMonochan string

const (
	LspPlugInPluginsLv2TriggerMidiMonochan01 LspPlugInPluginsLv2TriggerMidiMonochan = "01" // 01 (0) – 01
	LspPlugInPluginsLv2TriggerMidiMonochan02 LspPlugInPluginsLv2TriggerMidiMonochan = "02" // 02 (1) – 02
	LspPlugInPluginsLv2TriggerMidiMonochan03 LspPlugInPluginsLv2TriggerMidiMonochan = "03" // 03 (2) – 03
	LspPlugInPluginsLv2TriggerMidiMonochan04 LspPlugInPluginsLv2TriggerMidiMonochan = "04" // 04 (3) – 04
	LspPlugInPluginsLv2TriggerMidiMonochan05 LspPlugInPluginsLv2TriggerMidiMonochan = "05" // 05 (4) – 05
	LspPlugInPluginsLv2TriggerMidiMonochan06 LspPlugInPluginsLv2TriggerMidiMonochan = "06" // 06 (5) – 06
	LspPlugInPluginsLv2TriggerMidiMonochan07 LspPlugInPluginsLv2TriggerMidiMonochan = "07" // 07 (6) – 07
	LspPlugInPluginsLv2TriggerMidiMonochan08 LspPlugInPluginsLv2TriggerMidiMonochan = "08" // 08 (7) – 08
	LspPlugInPluginsLv2TriggerMidiMonochan09 LspPlugInPluginsLv2TriggerMidiMonochan = "09" // 09 (8) – 09
	LspPlugInPluginsLv2TriggerMidiMonochan10 LspPlugInPluginsLv2TriggerMidiMonochan = "10" // 10 (9) – 10
	LspPlugInPluginsLv2TriggerMidiMonochan11 LspPlugInPluginsLv2TriggerMidiMonochan = "11" // 11 (10) – 11
	LspPlugInPluginsLv2TriggerMidiMonochan12 LspPlugInPluginsLv2TriggerMidiMonochan = "12" // 12 (11) – 12
	LspPlugInPluginsLv2TriggerMidiMonochan13 LspPlugInPluginsLv2TriggerMidiMonochan = "13" // 13 (12) – 13
	LspPlugInPluginsLv2TriggerMidiMonochan14 LspPlugInPluginsLv2TriggerMidiMonochan = "14" // 14 (13) – 14
	LspPlugInPluginsLv2TriggerMidiMonochan15 LspPlugInPluginsLv2TriggerMidiMonochan = "15" // 15 (14) – 15
	LspPlugInPluginsLv2TriggerMidiMonochan16 LspPlugInPluginsLv2TriggerMidiMonochan = "16" // 16 (15) – 16
)

type LspPlugInPluginsLv2TriggerMidiMonomode string

const (
	LspPlugInPluginsLv2TriggerMidiMonomodePeak LspPlugInPluginsLv2TriggerMidiMonomode = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2TriggerMidiMonomodeRms LspPlugInPluginsLv2TriggerMidiMonomode = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2TriggerMidiMonomodeLowPass LspPlugInPluginsLv2TriggerMidiMonomode = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2TriggerMidiMonomodeUniform LspPlugInPluginsLv2TriggerMidiMonomode = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2TriggerMidiMononote string

const (
	LspPlugInPluginsLv2TriggerMidiMononoteC LspPlugInPluginsLv2TriggerMidiMononote = "C" // C (0) – C
	LspPlugInPluginsLv2TriggerMidiMononoteC1 LspPlugInPluginsLv2TriggerMidiMononote = "C#" // C# (1) – C#
	LspPlugInPluginsLv2TriggerMidiMononoteD LspPlugInPluginsLv2TriggerMidiMononote = "D" // D (2) – D
	LspPlugInPluginsLv2TriggerMidiMononoteD1 LspPlugInPluginsLv2TriggerMidiMononote = "D#" // D# (3) – D#
	LspPlugInPluginsLv2TriggerMidiMononoteE LspPlugInPluginsLv2TriggerMidiMononote = "E" // E (4) – E
	LspPlugInPluginsLv2TriggerMidiMononoteF LspPlugInPluginsLv2TriggerMidiMononote = "F" // F (5) – F
	LspPlugInPluginsLv2TriggerMidiMononoteF1 LspPlugInPluginsLv2TriggerMidiMononote = "F#" // F# (6) – F#
	LspPlugInPluginsLv2TriggerMidiMononoteG LspPlugInPluginsLv2TriggerMidiMononote = "G" // G (7) – G
	LspPlugInPluginsLv2TriggerMidiMononoteG1 LspPlugInPluginsLv2TriggerMidiMononote = "G#" // G# (8) – G#
	LspPlugInPluginsLv2TriggerMidiMononoteA LspPlugInPluginsLv2TriggerMidiMononote = "A" // A (9) – A
	LspPlugInPluginsLv2TriggerMidiMononoteA1 LspPlugInPluginsLv2TriggerMidiMononote = "A#" // A# (10) – A#
	LspPlugInPluginsLv2TriggerMidiMononoteB LspPlugInPluginsLv2TriggerMidiMononote = "B" // B (11) – B
)

type LspPlugInPluginsLv2TriggerMidiMonooct string

const (
	LspPlugInPluginsLv2TriggerMidiMonooct2 LspPlugInPluginsLv2TriggerMidiMonooct = "-2" // -2 (0) – -2
	LspPlugInPluginsLv2TriggerMidiMonooct1 LspPlugInPluginsLv2TriggerMidiMonooct = "-1" // -1 (1) – -1
	LspPlugInPluginsLv2TriggerMidiMonooct0 LspPlugInPluginsLv2TriggerMidiMonooct = "0" // 0 (2) – 0
	LspPlugInPluginsLv2TriggerMidiMonooct11 LspPlugInPluginsLv2TriggerMidiMonooct = "1" // 1 (3) – 1
	LspPlugInPluginsLv2TriggerMidiMonooct21 LspPlugInPluginsLv2TriggerMidiMonooct = "2" // 2 (4) – 2
	LspPlugInPluginsLv2TriggerMidiMonooct3 LspPlugInPluginsLv2TriggerMidiMonooct = "3" // 3 (5) – 3
	LspPlugInPluginsLv2TriggerMidiMonooct4 LspPlugInPluginsLv2TriggerMidiMonooct = "4" // 4 (6) – 4
	LspPlugInPluginsLv2TriggerMidiMonooct5 LspPlugInPluginsLv2TriggerMidiMonooct = "5" // 5 (7) – 5
	LspPlugInPluginsLv2TriggerMidiMonooct6 LspPlugInPluginsLv2TriggerMidiMonooct = "6" // 6 (8) – 6
	LspPlugInPluginsLv2TriggerMidiMonooct7 LspPlugInPluginsLv2TriggerMidiMonooct = "7" // 7 (9) – 7
	LspPlugInPluginsLv2TriggerMidiMonooct8 LspPlugInPluginsLv2TriggerMidiMonooct = "8" // 8 (10) – 8
)

