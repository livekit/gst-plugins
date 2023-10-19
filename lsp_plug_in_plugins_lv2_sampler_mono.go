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

type LspPlugInPluginsLv2SamplerMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2SamplerMono() (*LspPlugInPluginsLv2SamplerMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-sampler-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SamplerMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2SamplerMonoWithName(name string) (*LspPlugInPluginsLv2SamplerMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-sampler-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SamplerMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2SamplerMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2SamplerMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// chan (GstLspPlugInPluginsLv2SamplerMonochan)
//
// Channel

func (e *LspPlugInPluginsLv2SamplerMono) GetChan() (interface{}, error) {
	return e.Element.GetProperty("chan")
}

func (e *LspPlugInPluginsLv2SamplerMono) SetChan(value interface{}) error {
	return e.Element.SetProperty("chan", value)
}

// drft (float32)
//
// Time drifting

func (e *LspPlugInPluginsLv2SamplerMono) GetDrft() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("drft")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SamplerMono) SetDrft(value float32) error {
	return e.Element.SetProperty("drft", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2SamplerMono) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// dyna (float32)
//
// Dynamics

func (e *LspPlugInPluginsLv2SamplerMono) GetDyna() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerMono) SetDyna(value float32) error {
	return e.Element.SetProperty("dyna", value)
}

// fout (float32)
//
// Note-off fadeout

func (e *LspPlugInPluginsLv2SamplerMono) GetFout() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fout")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SamplerMono) SetFout(value float32) error {
	return e.Element.SetProperty("fout", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2SamplerMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// mn (int)
//
// MIDI Note #

func (e *LspPlugInPluginsLv2SamplerMono) GetMn() (int, error) {
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

// mute (bool)
//
// Forced mute

func (e *LspPlugInPluginsLv2SamplerMono) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SamplerMono) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// muting (bool)
//
// Mute on stop

func (e *LspPlugInPluginsLv2SamplerMono) GetMuting() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("muting")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SamplerMono) SetMuting(value bool) error {
	return e.Element.SetProperty("muting", value)
}

// noff (bool)
//
// Note-off handling

func (e *LspPlugInPluginsLv2SamplerMono) GetNoff() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("noff")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SamplerMono) SetNoff(value bool) error {
	return e.Element.SetProperty("noff", value)
}

// note (GstLspPlugInPluginsLv2SamplerMononote)
//
// Note

func (e *LspPlugInPluginsLv2SamplerMono) GetNote() (interface{}, error) {
	return e.Element.GetProperty("note")
}

func (e *LspPlugInPluginsLv2SamplerMono) SetNote(value interface{}) error {
	return e.Element.SetProperty("note", value)
}

// oct (GstLspPlugInPluginsLv2SamplerMonooct)
//
// Octave

func (e *LspPlugInPluginsLv2SamplerMono) GetOct() (interface{}, error) {
	return e.Element.GetProperty("oct")
}

func (e *LspPlugInPluginsLv2SamplerMono) SetOct(value interface{}) error {
	return e.Element.SetProperty("oct", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2SamplerMono) GetOutLatency() (int, error) {
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

// trg (bool)
//
// Instrument listen

func (e *LspPlugInPluginsLv2SamplerMono) GetTrg() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("trg")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SamplerMono) SetTrg(value bool) error {
	return e.Element.SetProperty("trg", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2SamplerMono) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2SamplerMonochan string

const (
	LspPlugInPluginsLv2SamplerMonochan01 LspPlugInPluginsLv2SamplerMonochan = "01" // 01 (0) – 01
	LspPlugInPluginsLv2SamplerMonochan02 LspPlugInPluginsLv2SamplerMonochan = "02" // 02 (1) – 02
	LspPlugInPluginsLv2SamplerMonochan03 LspPlugInPluginsLv2SamplerMonochan = "03" // 03 (2) – 03
	LspPlugInPluginsLv2SamplerMonochan04 LspPlugInPluginsLv2SamplerMonochan = "04" // 04 (3) – 04
	LspPlugInPluginsLv2SamplerMonochan05 LspPlugInPluginsLv2SamplerMonochan = "05" // 05 (4) – 05
	LspPlugInPluginsLv2SamplerMonochan06 LspPlugInPluginsLv2SamplerMonochan = "06" // 06 (5) – 06
	LspPlugInPluginsLv2SamplerMonochan07 LspPlugInPluginsLv2SamplerMonochan = "07" // 07 (6) – 07
	LspPlugInPluginsLv2SamplerMonochan08 LspPlugInPluginsLv2SamplerMonochan = "08" // 08 (7) – 08
	LspPlugInPluginsLv2SamplerMonochan09 LspPlugInPluginsLv2SamplerMonochan = "09" // 09 (8) – 09
	LspPlugInPluginsLv2SamplerMonochan10 LspPlugInPluginsLv2SamplerMonochan = "10" // 10 (9) – 10
	LspPlugInPluginsLv2SamplerMonochan11 LspPlugInPluginsLv2SamplerMonochan = "11" // 11 (10) – 11
	LspPlugInPluginsLv2SamplerMonochan12 LspPlugInPluginsLv2SamplerMonochan = "12" // 12 (11) – 12
	LspPlugInPluginsLv2SamplerMonochan13 LspPlugInPluginsLv2SamplerMonochan = "13" // 13 (12) – 13
	LspPlugInPluginsLv2SamplerMonochan14 LspPlugInPluginsLv2SamplerMonochan = "14" // 14 (13) – 14
	LspPlugInPluginsLv2SamplerMonochan15 LspPlugInPluginsLv2SamplerMonochan = "15" // 15 (14) – 15
	LspPlugInPluginsLv2SamplerMonochan16 LspPlugInPluginsLv2SamplerMonochan = "16" // 16 (15) – 16
)

type LspPlugInPluginsLv2SamplerMononote string

const (
	LspPlugInPluginsLv2SamplerMononoteC LspPlugInPluginsLv2SamplerMononote = "C" // C (0) – C
	LspPlugInPluginsLv2SamplerMononoteC1 LspPlugInPluginsLv2SamplerMononote = "C#" // C# (1) – C#
	LspPlugInPluginsLv2SamplerMononoteD LspPlugInPluginsLv2SamplerMononote = "D" // D (2) – D
	LspPlugInPluginsLv2SamplerMononoteD1 LspPlugInPluginsLv2SamplerMononote = "D#" // D# (3) – D#
	LspPlugInPluginsLv2SamplerMononoteE LspPlugInPluginsLv2SamplerMononote = "E" // E (4) – E
	LspPlugInPluginsLv2SamplerMononoteF LspPlugInPluginsLv2SamplerMononote = "F" // F (5) – F
	LspPlugInPluginsLv2SamplerMononoteF1 LspPlugInPluginsLv2SamplerMononote = "F#" // F# (6) – F#
	LspPlugInPluginsLv2SamplerMononoteG LspPlugInPluginsLv2SamplerMononote = "G" // G (7) – G
	LspPlugInPluginsLv2SamplerMononoteG1 LspPlugInPluginsLv2SamplerMononote = "G#" // G# (8) – G#
	LspPlugInPluginsLv2SamplerMononoteA LspPlugInPluginsLv2SamplerMononote = "A" // A (9) – A
	LspPlugInPluginsLv2SamplerMononoteA1 LspPlugInPluginsLv2SamplerMononote = "A#" // A# (10) – A#
	LspPlugInPluginsLv2SamplerMononoteB LspPlugInPluginsLv2SamplerMononote = "B" // B (11) – B
)

type LspPlugInPluginsLv2SamplerMonooct string

const (
	LspPlugInPluginsLv2SamplerMonooct2 LspPlugInPluginsLv2SamplerMonooct = "-2" // -2 (0) – -2
	LspPlugInPluginsLv2SamplerMonooct1 LspPlugInPluginsLv2SamplerMonooct = "-1" // -1 (1) – -1
	LspPlugInPluginsLv2SamplerMonooct0 LspPlugInPluginsLv2SamplerMonooct = "0" // 0 (2) – 0
	LspPlugInPluginsLv2SamplerMonooct11 LspPlugInPluginsLv2SamplerMonooct = "1" // 1 (3) – 1
	LspPlugInPluginsLv2SamplerMonooct21 LspPlugInPluginsLv2SamplerMonooct = "2" // 2 (4) – 2
	LspPlugInPluginsLv2SamplerMonooct3 LspPlugInPluginsLv2SamplerMonooct = "3" // 3 (5) – 3
	LspPlugInPluginsLv2SamplerMonooct4 LspPlugInPluginsLv2SamplerMonooct = "4" // 4 (6) – 4
	LspPlugInPluginsLv2SamplerMonooct5 LspPlugInPluginsLv2SamplerMonooct = "5" // 5 (7) – 5
	LspPlugInPluginsLv2SamplerMonooct6 LspPlugInPluginsLv2SamplerMonooct = "6" // 6 (8) – 6
	LspPlugInPluginsLv2SamplerMonooct7 LspPlugInPluginsLv2SamplerMonooct = "7" // 7 (9) – 7
	LspPlugInPluginsLv2SamplerMonooct8 LspPlugInPluginsLv2SamplerMonooct = "8" // 8 (10) – 8
)

