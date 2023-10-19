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

type LspPlugInPluginsLv2CompressorMs struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2CompressorMs() (*LspPlugInPluginsLv2CompressorMs, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-compressor-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2CompressorMsWithName(name string) (*LspPlugInPluginsLv2CompressorMs, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-compressor-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-m (float32)
//
// Attack level Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetAlM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetAlM(value float32) error {
	return e.Element.SetProperty("al-m", value)
}

// al-s (float32)
//
// Attack level Side

func (e *LspPlugInPluginsLv2CompressorMs) GetAlS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetAlS(value float32) error {
	return e.Element.SetProperty("al-s", value)
}

// at-m (float32)
//
// Attack time Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetAtM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetAtM(value float32) error {
	return e.Element.SetProperty("at-m", value)
}

// at-s (float32)
//
// Attack time Side

func (e *LspPlugInPluginsLv2CompressorMs) GetAtS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetAtS(value float32) error {
	return e.Element.SetProperty("at-s", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2CompressorMs) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMs) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-m (float32)
//
// Dry gain Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetCdrM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCdrM(value float32) error {
	return e.Element.SetProperty("cdr-m", value)
}

// cdr-s (float32)
//
// Dry gain Side

func (e *LspPlugInPluginsLv2CompressorMs) GetCdrS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCdrS(value float32) error {
	return e.Element.SetProperty("cdr-s", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2CompressorMs) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMs) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-m (float32)
//
// Curve level meter Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetClmM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-s (float32)
//
// Curve level meter Side

func (e *LspPlugInPluginsLv2CompressorMs) GetClmS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cm-m (GstLspPlugInPluginsLv2CompressorMscmM)
//
// Compression mode Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetCmM() (interface{}, error) {
	return e.Element.GetProperty("cm-m")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCmM(value interface{}) error {
	return e.Element.SetProperty("cm-m", value)
}

// cm-s (GstLspPlugInPluginsLv2CompressorMscmS)
//
// Compression mode Side

func (e *LspPlugInPluginsLv2CompressorMs) GetCmS() (interface{}, error) {
	return e.Element.GetProperty("cm-s")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCmS(value interface{}) error {
	return e.Element.SetProperty("cm-s", value)
}

// cr-m (float32)
//
// Ratio Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetCrM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCrM(value float32) error {
	return e.Element.SetProperty("cr-m", value)
}

// cr-s (float32)
//
// Ratio Side

func (e *LspPlugInPluginsLv2CompressorMs) GetCrS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCrS(value float32) error {
	return e.Element.SetProperty("cr-s", value)
}

// cwt-m (float32)
//
// Wet gain Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetCwtM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCwtM(value float32) error {
	return e.Element.SetProperty("cwt-m", value)
}

// cwt-s (float32)
//
// Wet gain Side

func (e *LspPlugInPluginsLv2CompressorMs) GetCwtS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetCwtS(value float32) error {
	return e.Element.SetProperty("cwt-s", value)
}

// elm-m (float32)
//
// Envelope level meter Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetElmM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-s (float32)
//
// Envelope level meter Side

func (e *LspPlugInPluginsLv2CompressorMs) GetElmS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elv-m (bool)
//
// Envelope level visibility Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetElvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetElvM(value bool) error {
	return e.Element.SetProperty("elv-m", value)
}

// elv-s (bool)
//
// Envelope level visibility Side

func (e *LspPlugInPluginsLv2CompressorMs) GetElvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetElvS(value bool) error {
	return e.Element.SetProperty("elv-s", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2CompressorMs) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMs) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2CompressorMs) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorMs) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv-m (bool)
//
// Gain reduction visibility Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetGrvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetGrvM(value bool) error {
	return e.Element.SetProperty("grv-m", value)
}

// grv-s (bool)
//
// Gain reduction visibility Side

func (e *LspPlugInPluginsLv2CompressorMs) GetGrvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetGrvS(value bool) error {
	return e.Element.SetProperty("grv-s", value)
}

// ilm-m (float32)
//
// Input level meter Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetIlmM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm-s (float32)
//
// Input level meter Side

func (e *LspPlugInPluginsLv2CompressorMs) GetIlmS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv-m (bool)
//
// Input level visibility Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetIlvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetIlvM(value bool) error {
	return e.Element.SetProperty("ilv-m", value)
}

// ilv-s (bool)
//
// Input level visibility Side

func (e *LspPlugInPluginsLv2CompressorMs) GetIlvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetIlvS(value bool) error {
	return e.Element.SetProperty("ilv-s", value)
}

// kn-m (float32)
//
// Knee Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetKnM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetKnM(value float32) error {
	return e.Element.SetProperty("kn-m", value)
}

// kn-s (float32)
//
// Knee Side

func (e *LspPlugInPluginsLv2CompressorMs) GetKnS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetKnS(value float32) error {
	return e.Element.SetProperty("kn-s", value)
}

// mk-m (float32)
//
// Makeup gain Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetMkM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetMkM(value float32) error {
	return e.Element.SetProperty("mk-m", value)
}

// mk-s (float32)
//
// Makeup gain Side

func (e *LspPlugInPluginsLv2CompressorMs) GetMkS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetMkS(value float32) error {
	return e.Element.SetProperty("mk-s", value)
}

// msl (bool)
//
// Mid/Side listen

func (e *LspPlugInPluginsLv2CompressorMs) GetMsl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("msl")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetMsl(value bool) error {
	return e.Element.SetProperty("msl", value)
}

// olm-m (float32)
//
// Output level meter Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetOlmM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olm-s (float32)
//
// Output level meter Side

func (e *LspPlugInPluginsLv2CompressorMs) GetOlmS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olv-m (bool)
//
// Output level visibility Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetOlvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetOlvM(value bool) error {
	return e.Element.SetProperty("olv-m", value)
}

// olv-s (bool)
//
// Output level visibility Side

func (e *LspPlugInPluginsLv2CompressorMs) GetOlvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetOlvS(value bool) error {
	return e.Element.SetProperty("olv-s", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2CompressorMs) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2CompressorMs) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorMs) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl-m (float32)
//
// Release level Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetRlM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-s (float32)
//
// Release level Side

func (e *LspPlugInPluginsLv2CompressorMs) GetRlS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-m (float32)
//
// Reduction level meter Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetRlmM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-s (float32)
//
// Reduction level meter Side

func (e *LspPlugInPluginsLv2CompressorMs) GetRlmS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rrl-m (float32)
//
// Relative release level Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetRrlM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetRrlM(value float32) error {
	return e.Element.SetProperty("rrl-m", value)
}

// rrl-s (float32)
//
// Relative release level Side

func (e *LspPlugInPluginsLv2CompressorMs) GetRrlS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetRrlS(value float32) error {
	return e.Element.SetProperty("rrl-s", value)
}

// rt-m (float32)
//
// Release time Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetRtM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetRtM(value float32) error {
	return e.Element.SetProperty("rt-m", value)
}

// rt-s (float32)
//
// Release time Side

func (e *LspPlugInPluginsLv2CompressorMs) GetRtS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetRtS(value float32) error {
	return e.Element.SetProperty("rt-s", value)
}

// scl-m (bool)
//
// Sidechain listen Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetSclM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSclM(value bool) error {
	return e.Element.SetProperty("scl-m", value)
}

// scl-s (bool)
//
// Sidechain listen Side

func (e *LspPlugInPluginsLv2CompressorMs) GetSclS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSclS(value bool) error {
	return e.Element.SetProperty("scl-s", value)
}

// scm-m (GstLspPlugInPluginsLv2CompressorMsscmM)
//
// Sidechain mode Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetScmM() (interface{}, error) {
	return e.Element.GetProperty("scm-m")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScmM(value interface{}) error {
	return e.Element.SetProperty("scm-m", value)
}

// scm-s (GstLspPlugInPluginsLv2CompressorMsscmS)
//
// Sidechain mode Side

func (e *LspPlugInPluginsLv2CompressorMs) GetScmS() (interface{}, error) {
	return e.Element.GetProperty("scm-s")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScmS(value interface{}) error {
	return e.Element.SetProperty("scm-s", value)
}

// scp-m (float32)
//
// Sidechain preamp Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetScpM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScpM(value float32) error {
	return e.Element.SetProperty("scp-m", value)
}

// scp-s (float32)
//
// Sidechain preamp Side

func (e *LspPlugInPluginsLv2CompressorMs) GetScpS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScpS(value float32) error {
	return e.Element.SetProperty("scp-s", value)
}

// scr-m (float32)
//
// Sidechain reactivity Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetScrM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScrM(value float32) error {
	return e.Element.SetProperty("scr-m", value)
}

// scr-s (float32)
//
// Sidechain reactivity Side

func (e *LspPlugInPluginsLv2CompressorMs) GetScrS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScrS(value float32) error {
	return e.Element.SetProperty("scr-s", value)
}

// scs-m (GstLspPlugInPluginsLv2CompressorMsscsM)
//
// Sidechain source Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetScsM() (interface{}, error) {
	return e.Element.GetProperty("scs-m")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScsM(value interface{}) error {
	return e.Element.SetProperty("scs-m", value)
}

// scs-s (GstLspPlugInPluginsLv2CompressorMsscsS)
//
// Sidechain source Side

func (e *LspPlugInPluginsLv2CompressorMs) GetScsS() (interface{}, error) {
	return e.Element.GetProperty("scs-s")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetScsS(value interface{}) error {
	return e.Element.SetProperty("scs-s", value)
}

// sct-m (GstLspPlugInPluginsLv2CompressorMssctM)
//
// Sidechain type Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetSctM() (interface{}, error) {
	return e.Element.GetProperty("sct-m")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSctM(value interface{}) error {
	return e.Element.SetProperty("sct-m", value)
}

// sct-s (GstLspPlugInPluginsLv2CompressorMssctS)
//
// Sidechain type Side

func (e *LspPlugInPluginsLv2CompressorMs) GetSctS() (interface{}, error) {
	return e.Element.GetProperty("sct-s")
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSctS(value interface{}) error {
	return e.Element.SetProperty("sct-s", value)
}

// sla-m (float32)
//
// Sidechain lookahead Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetSlaM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSlaM(value float32) error {
	return e.Element.SetProperty("sla-m", value)
}

// sla-s (float32)
//
// Sidechain lookahead Side

func (e *LspPlugInPluginsLv2CompressorMs) GetSlaS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSlaS(value float32) error {
	return e.Element.SetProperty("sla-s", value)
}

// slm-m (float32)
//
// Sidechain level meter Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetSlmM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slm-s (float32)
//
// Sidechain level meter Side

func (e *LspPlugInPluginsLv2CompressorMs) GetSlmS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slv-m (bool)
//
// Sidechain level visibility Mid

func (e *LspPlugInPluginsLv2CompressorMs) GetSlvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSlvM(value bool) error {
	return e.Element.SetProperty("slv-m", value)
}

// slv-s (bool)
//
// Sidechain level visibility Side

func (e *LspPlugInPluginsLv2CompressorMs) GetSlvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorMs) SetSlvS(value bool) error {
	return e.Element.SetProperty("slv-s", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2CompressorMscmM string

const (
	LspPlugInPluginsLv2CompressorMscmMDown LspPlugInPluginsLv2CompressorMscmM = "Down" // Down (0) – Down
	LspPlugInPluginsLv2CompressorMscmMUp LspPlugInPluginsLv2CompressorMscmM = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2CompressorMscmS string

const (
	LspPlugInPluginsLv2CompressorMscmSDown LspPlugInPluginsLv2CompressorMscmS = "Down" // Down (0) – Down
	LspPlugInPluginsLv2CompressorMscmSUp LspPlugInPluginsLv2CompressorMscmS = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2CompressorMsscmM string

const (
	LspPlugInPluginsLv2CompressorMsscmMPeak LspPlugInPluginsLv2CompressorMsscmM = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2CompressorMsscmMRms LspPlugInPluginsLv2CompressorMsscmM = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2CompressorMsscmMLowPass LspPlugInPluginsLv2CompressorMsscmM = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2CompressorMsscmMUniform LspPlugInPluginsLv2CompressorMsscmM = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2CompressorMsscmS string

const (
	LspPlugInPluginsLv2CompressorMsscmSPeak LspPlugInPluginsLv2CompressorMsscmS = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2CompressorMsscmSRms LspPlugInPluginsLv2CompressorMsscmS = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2CompressorMsscmSLowPass LspPlugInPluginsLv2CompressorMsscmS = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2CompressorMsscmSUniform LspPlugInPluginsLv2CompressorMsscmS = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2CompressorMsscsM string

const (
	LspPlugInPluginsLv2CompressorMsscsMMiddle LspPlugInPluginsLv2CompressorMsscsM = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2CompressorMsscsMSide LspPlugInPluginsLv2CompressorMsscsM = "Side" // Side (1) – Side
	LspPlugInPluginsLv2CompressorMsscsMLeft LspPlugInPluginsLv2CompressorMsscsM = "Left" // Left (2) – Left
	LspPlugInPluginsLv2CompressorMsscsMRight LspPlugInPluginsLv2CompressorMsscsM = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2CompressorMsscsS string

const (
	LspPlugInPluginsLv2CompressorMsscsSMiddle LspPlugInPluginsLv2CompressorMsscsS = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2CompressorMsscsSSide LspPlugInPluginsLv2CompressorMsscsS = "Side" // Side (1) – Side
	LspPlugInPluginsLv2CompressorMsscsSLeft LspPlugInPluginsLv2CompressorMsscsS = "Left" // Left (2) – Left
	LspPlugInPluginsLv2CompressorMsscsSRight LspPlugInPluginsLv2CompressorMsscsS = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2CompressorMssctM string

const (
	LspPlugInPluginsLv2CompressorMssctMFeedForward LspPlugInPluginsLv2CompressorMssctM = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2CompressorMssctMFeedBack LspPlugInPluginsLv2CompressorMssctM = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2CompressorMssctS string

const (
	LspPlugInPluginsLv2CompressorMssctSFeedForward LspPlugInPluginsLv2CompressorMssctS = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2CompressorMssctSFeedBack LspPlugInPluginsLv2CompressorMssctS = "Feed-back" // Feed-back (1) – Feed-back
)

