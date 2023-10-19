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

type LspPlugInPluginsLv2SamplerStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2SamplerStereo() (*LspPlugInPluginsLv2SamplerStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-sampler-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SamplerStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2SamplerStereoWithName(name string) (*LspPlugInPluginsLv2SamplerStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-sampler-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SamplerStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2SamplerStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// chan (GstLspPlugInPluginsLv2SamplerStereochan)
//
// Channel

func (e *LspPlugInPluginsLv2SamplerStereo) GetChan() (interface{}, error) {
	return e.Element.GetProperty("chan")
}

func (e *LspPlugInPluginsLv2SamplerStereo) SetChan(value interface{}) error {
	return e.Element.SetProperty("chan", value)
}

// drft (float32)
//
// Time drifting

func (e *LspPlugInPluginsLv2SamplerStereo) GetDrft() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetDrft(value float32) error {
	return e.Element.SetProperty("drft", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2SamplerStereo) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// dyna (float32)
//
// Dynamics

func (e *LspPlugInPluginsLv2SamplerStereo) GetDyna() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetDyna(value float32) error {
	return e.Element.SetProperty("dyna", value)
}

// fout (float32)
//
// Note-off fadeout

func (e *LspPlugInPluginsLv2SamplerStereo) GetFout() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetFout(value float32) error {
	return e.Element.SetProperty("fout", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2SamplerStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// mn (int)
//
// MIDI Note #

func (e *LspPlugInPluginsLv2SamplerStereo) GetMn() (int, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) GetMute() (bool, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// muting (bool)
//
// Mute on stop

func (e *LspPlugInPluginsLv2SamplerStereo) GetMuting() (bool, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetMuting(value bool) error {
	return e.Element.SetProperty("muting", value)
}

// noff (bool)
//
// Note-off handling

func (e *LspPlugInPluginsLv2SamplerStereo) GetNoff() (bool, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetNoff(value bool) error {
	return e.Element.SetProperty("noff", value)
}

// note (GstLspPlugInPluginsLv2SamplerStereonote)
//
// Note

func (e *LspPlugInPluginsLv2SamplerStereo) GetNote() (interface{}, error) {
	return e.Element.GetProperty("note")
}

func (e *LspPlugInPluginsLv2SamplerStereo) SetNote(value interface{}) error {
	return e.Element.SetProperty("note", value)
}

// oct (GstLspPlugInPluginsLv2SamplerStereooct)
//
// Octave

func (e *LspPlugInPluginsLv2SamplerStereo) GetOct() (interface{}, error) {
	return e.Element.GetProperty("oct")
}

func (e *LspPlugInPluginsLv2SamplerStereo) SetOct(value interface{}) error {
	return e.Element.SetProperty("oct", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2SamplerStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) GetTrg() (bool, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetTrg(value bool) error {
	return e.Element.SetProperty("trg", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2SamplerStereo) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2SamplerStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2SamplerStereochan string

const (
	LspPlugInPluginsLv2SamplerStereochan01 LspPlugInPluginsLv2SamplerStereochan = "01" // 01 (0) – 01
	LspPlugInPluginsLv2SamplerStereochan02 LspPlugInPluginsLv2SamplerStereochan = "02" // 02 (1) – 02
	LspPlugInPluginsLv2SamplerStereochan03 LspPlugInPluginsLv2SamplerStereochan = "03" // 03 (2) – 03
	LspPlugInPluginsLv2SamplerStereochan04 LspPlugInPluginsLv2SamplerStereochan = "04" // 04 (3) – 04
	LspPlugInPluginsLv2SamplerStereochan05 LspPlugInPluginsLv2SamplerStereochan = "05" // 05 (4) – 05
	LspPlugInPluginsLv2SamplerStereochan06 LspPlugInPluginsLv2SamplerStereochan = "06" // 06 (5) – 06
	LspPlugInPluginsLv2SamplerStereochan07 LspPlugInPluginsLv2SamplerStereochan = "07" // 07 (6) – 07
	LspPlugInPluginsLv2SamplerStereochan08 LspPlugInPluginsLv2SamplerStereochan = "08" // 08 (7) – 08
	LspPlugInPluginsLv2SamplerStereochan09 LspPlugInPluginsLv2SamplerStereochan = "09" // 09 (8) – 09
	LspPlugInPluginsLv2SamplerStereochan10 LspPlugInPluginsLv2SamplerStereochan = "10" // 10 (9) – 10
	LspPlugInPluginsLv2SamplerStereochan11 LspPlugInPluginsLv2SamplerStereochan = "11" // 11 (10) – 11
	LspPlugInPluginsLv2SamplerStereochan12 LspPlugInPluginsLv2SamplerStereochan = "12" // 12 (11) – 12
	LspPlugInPluginsLv2SamplerStereochan13 LspPlugInPluginsLv2SamplerStereochan = "13" // 13 (12) – 13
	LspPlugInPluginsLv2SamplerStereochan14 LspPlugInPluginsLv2SamplerStereochan = "14" // 14 (13) – 14
	LspPlugInPluginsLv2SamplerStereochan15 LspPlugInPluginsLv2SamplerStereochan = "15" // 15 (14) – 15
	LspPlugInPluginsLv2SamplerStereochan16 LspPlugInPluginsLv2SamplerStereochan = "16" // 16 (15) – 16
)

type LspPlugInPluginsLv2SamplerStereonote string

const (
	LspPlugInPluginsLv2SamplerStereonoteC LspPlugInPluginsLv2SamplerStereonote = "C" // C (0) – C
	LspPlugInPluginsLv2SamplerStereonoteC1 LspPlugInPluginsLv2SamplerStereonote = "C#" // C# (1) – C#
	LspPlugInPluginsLv2SamplerStereonoteD LspPlugInPluginsLv2SamplerStereonote = "D" // D (2) – D
	LspPlugInPluginsLv2SamplerStereonoteD1 LspPlugInPluginsLv2SamplerStereonote = "D#" // D# (3) – D#
	LspPlugInPluginsLv2SamplerStereonoteE LspPlugInPluginsLv2SamplerStereonote = "E" // E (4) – E
	LspPlugInPluginsLv2SamplerStereonoteF LspPlugInPluginsLv2SamplerStereonote = "F" // F (5) – F
	LspPlugInPluginsLv2SamplerStereonoteF1 LspPlugInPluginsLv2SamplerStereonote = "F#" // F# (6) – F#
	LspPlugInPluginsLv2SamplerStereonoteG LspPlugInPluginsLv2SamplerStereonote = "G" // G (7) – G
	LspPlugInPluginsLv2SamplerStereonoteG1 LspPlugInPluginsLv2SamplerStereonote = "G#" // G# (8) – G#
	LspPlugInPluginsLv2SamplerStereonoteA LspPlugInPluginsLv2SamplerStereonote = "A" // A (9) – A
	LspPlugInPluginsLv2SamplerStereonoteA1 LspPlugInPluginsLv2SamplerStereonote = "A#" // A# (10) – A#
	LspPlugInPluginsLv2SamplerStereonoteB LspPlugInPluginsLv2SamplerStereonote = "B" // B (11) – B
)

type LspPlugInPluginsLv2SamplerStereooct string

const (
	LspPlugInPluginsLv2SamplerStereooct2 LspPlugInPluginsLv2SamplerStereooct = "-2" // -2 (0) – -2
	LspPlugInPluginsLv2SamplerStereooct1 LspPlugInPluginsLv2SamplerStereooct = "-1" // -1 (1) – -1
	LspPlugInPluginsLv2SamplerStereooct0 LspPlugInPluginsLv2SamplerStereooct = "0" // 0 (2) – 0
	LspPlugInPluginsLv2SamplerStereooct11 LspPlugInPluginsLv2SamplerStereooct = "1" // 1 (3) – 1
	LspPlugInPluginsLv2SamplerStereooct21 LspPlugInPluginsLv2SamplerStereooct = "2" // 2 (4) – 2
	LspPlugInPluginsLv2SamplerStereooct3 LspPlugInPluginsLv2SamplerStereooct = "3" // 3 (5) – 3
	LspPlugInPluginsLv2SamplerStereooct4 LspPlugInPluginsLv2SamplerStereooct = "4" // 4 (6) – 4
	LspPlugInPluginsLv2SamplerStereooct5 LspPlugInPluginsLv2SamplerStereooct = "5" // 5 (7) – 5
	LspPlugInPluginsLv2SamplerStereooct6 LspPlugInPluginsLv2SamplerStereooct = "6" // 6 (8) – 6
	LspPlugInPluginsLv2SamplerStereooct7 LspPlugInPluginsLv2SamplerStereooct = "7" // 7 (9) – 7
	LspPlugInPluginsLv2SamplerStereooct8 LspPlugInPluginsLv2SamplerStereooct = "8" // 8 (10) – 8
)

