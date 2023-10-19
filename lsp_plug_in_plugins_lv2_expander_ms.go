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

type LspPlugInPluginsLv2ExpanderMs struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ExpanderMs() (*LspPlugInPluginsLv2ExpanderMs, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-expander-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ExpanderMsWithName(name string) (*LspPlugInPluginsLv2ExpanderMs, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-expander-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-m (float32)
//
// Attack level Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetAlM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetAlM(value float32) error {
	return e.Element.SetProperty("al-m", value)
}

// al-s (float32)
//
// Attack level Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetAlS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetAlS(value float32) error {
	return e.Element.SetProperty("al-s", value)
}

// at-m (float32)
//
// Attack time Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetAtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetAtM(value float32) error {
	return e.Element.SetProperty("at-m", value)
}

// at-s (float32)
//
// Attack time Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetAtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetAtS(value float32) error {
	return e.Element.SetProperty("at-s", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ExpanderMs) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-m (float32)
//
// Dry gain Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetCdrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetCdrM(value float32) error {
	return e.Element.SetProperty("cdr-m", value)
}

// cdr-s (float32)
//
// Dry gain Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetCdrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetCdrS(value float32) error {
	return e.Element.SetProperty("cdr-s", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2ExpanderMs) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-m (float32)
//
// Curve level meter Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetClmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetClmS() (float32, error) {
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

// cr-m (float32)
//
// Ratio Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetCrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetCrM(value float32) error {
	return e.Element.SetProperty("cr-m", value)
}

// cr-s (float32)
//
// Ratio Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetCrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetCrS(value float32) error {
	return e.Element.SetProperty("cr-s", value)
}

// cwt-m (float32)
//
// Wet gain Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetCwtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetCwtM(value float32) error {
	return e.Element.SetProperty("cwt-m", value)
}

// cwt-s (float32)
//
// Wet gain Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetCwtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetCwtS(value float32) error {
	return e.Element.SetProperty("cwt-s", value)
}

// elm-m (float32)
//
// Envelope level meter Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetElmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetElmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetElvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetElvM(value bool) error {
	return e.Element.SetProperty("elv-m", value)
}

// elv-s (bool)
//
// Envelope level visibility Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetElvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetElvS(value bool) error {
	return e.Element.SetProperty("elv-s", value)
}

// em-m (GstLspPlugInPluginsLv2ExpanderMsemM)
//
// Expander mode Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetEmM() (interface{}, error) {
	return e.Element.GetProperty("em-m")
}

func (e *LspPlugInPluginsLv2ExpanderMs) SetEmM(value interface{}) error {
	return e.Element.SetProperty("em-m", value)
}

// em-s (GstLspPlugInPluginsLv2ExpanderMsemS)
//
// Expander mode Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetEmS() (interface{}, error) {
	return e.Element.GetProperty("em-s")
}

func (e *LspPlugInPluginsLv2ExpanderMs) SetEmS(value interface{}) error {
	return e.Element.SetProperty("em-s", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ExpanderMs) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ExpanderMs) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv-m (bool)
//
// Gain reduction visibility Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetGrvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetGrvM(value bool) error {
	return e.Element.SetProperty("grv-m", value)
}

// grv-s (bool)
//
// Gain reduction visibility Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetGrvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetGrvS(value bool) error {
	return e.Element.SetProperty("grv-s", value)
}

// ilm-m (float32)
//
// Input level meter Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetIlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetIlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetIlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetIlvM(value bool) error {
	return e.Element.SetProperty("ilv-m", value)
}

// ilv-s (bool)
//
// Input level visibility Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetIlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetIlvS(value bool) error {
	return e.Element.SetProperty("ilv-s", value)
}

// kn-m (float32)
//
// Knee Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetKnM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetKnM(value float32) error {
	return e.Element.SetProperty("kn-m", value)
}

// kn-s (float32)
//
// Knee Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetKnS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetKnS(value float32) error {
	return e.Element.SetProperty("kn-s", value)
}

// mk-m (float32)
//
// Makeup gain Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetMkM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetMkM(value float32) error {
	return e.Element.SetProperty("mk-m", value)
}

// mk-s (float32)
//
// Makeup gain Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetMkS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetMkS(value float32) error {
	return e.Element.SetProperty("mk-s", value)
}

// msl (bool)
//
// Mid/Side listen

func (e *LspPlugInPluginsLv2ExpanderMs) GetMsl() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetMsl(value bool) error {
	return e.Element.SetProperty("msl", value)
}

// olm-m (float32)
//
// Output level meter Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetOlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetOlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetOlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetOlvM(value bool) error {
	return e.Element.SetProperty("olv-m", value)
}

// olv-s (bool)
//
// Output level visibility Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetOlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetOlvS(value bool) error {
	return e.Element.SetProperty("olv-s", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ExpanderMs) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl-m (float32)
//
// Release level Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetRlM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetRlS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetRlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetRlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetRrlM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetRrlM(value float32) error {
	return e.Element.SetProperty("rrl-m", value)
}

// rrl-s (float32)
//
// Relative release level Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetRrlS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetRrlS(value float32) error {
	return e.Element.SetProperty("rrl-s", value)
}

// rt-m (float32)
//
// Release time Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetRtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetRtM(value float32) error {
	return e.Element.SetProperty("rt-m", value)
}

// rt-s (float32)
//
// Release time Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetRtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetRtS(value float32) error {
	return e.Element.SetProperty("rt-s", value)
}

// scl-m (bool)
//
// Sidechain listen Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetSclM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetSclM(value bool) error {
	return e.Element.SetProperty("scl-m", value)
}

// scl-s (bool)
//
// Sidechain listen Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetSclS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetSclS(value bool) error {
	return e.Element.SetProperty("scl-s", value)
}

// scm-m (GstLspPlugInPluginsLv2ExpanderMsscmM)
//
// Sidechain mode Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetScmM() (interface{}, error) {
	return e.Element.GetProperty("scm-m")
}

func (e *LspPlugInPluginsLv2ExpanderMs) SetScmM(value interface{}) error {
	return e.Element.SetProperty("scm-m", value)
}

// scm-s (GstLspPlugInPluginsLv2ExpanderMsscmS)
//
// Sidechain mode Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetScmS() (interface{}, error) {
	return e.Element.GetProperty("scm-s")
}

func (e *LspPlugInPluginsLv2ExpanderMs) SetScmS(value interface{}) error {
	return e.Element.SetProperty("scm-s", value)
}

// scp-m (float32)
//
// Sidechain preamp Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetScpM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetScpM(value float32) error {
	return e.Element.SetProperty("scp-m", value)
}

// scp-s (float32)
//
// Sidechain preamp Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetScpS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetScpS(value float32) error {
	return e.Element.SetProperty("scp-s", value)
}

// scr-m (float32)
//
// Sidechain reactivity Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetScrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetScrM(value float32) error {
	return e.Element.SetProperty("scr-m", value)
}

// scr-s (float32)
//
// Sidechain reactivity Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetScrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetScrS(value float32) error {
	return e.Element.SetProperty("scr-s", value)
}

// scs-m (GstLspPlugInPluginsLv2ExpanderMsscsM)
//
// Sidechain source Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetScsM() (interface{}, error) {
	return e.Element.GetProperty("scs-m")
}

func (e *LspPlugInPluginsLv2ExpanderMs) SetScsM(value interface{}) error {
	return e.Element.SetProperty("scs-m", value)
}

// scs-s (GstLspPlugInPluginsLv2ExpanderMsscsS)
//
// Sidechain source Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetScsS() (interface{}, error) {
	return e.Element.GetProperty("scs-s")
}

func (e *LspPlugInPluginsLv2ExpanderMs) SetScsS(value interface{}) error {
	return e.Element.SetProperty("scs-s", value)
}

// sla-m (float32)
//
// Sidechain lookahead Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetSlaM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetSlaM(value float32) error {
	return e.Element.SetProperty("sla-m", value)
}

// sla-s (float32)
//
// Sidechain lookahead Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetSlaS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetSlaS(value float32) error {
	return e.Element.SetProperty("sla-s", value)
}

// slm-m (float32)
//
// Sidechain level meter Mid

func (e *LspPlugInPluginsLv2ExpanderMs) GetSlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetSlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) GetSlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetSlvM(value bool) error {
	return e.Element.SetProperty("slv-m", value)
}

// slv-s (bool)
//
// Sidechain level visibility Side

func (e *LspPlugInPluginsLv2ExpanderMs) GetSlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ExpanderMs) SetSlvS(value bool) error {
	return e.Element.SetProperty("slv-s", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ExpanderMsscsS string

const (
	LspPlugInPluginsLv2ExpanderMsscsSMiddle LspPlugInPluginsLv2ExpanderMsscsS = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2ExpanderMsscsSSide LspPlugInPluginsLv2ExpanderMsscsS = "Side" // Side (1) – Side
	LspPlugInPluginsLv2ExpanderMsscsSLeft LspPlugInPluginsLv2ExpanderMsscsS = "Left" // Left (2) – Left
	LspPlugInPluginsLv2ExpanderMsscsSRight LspPlugInPluginsLv2ExpanderMsscsS = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2ExpanderMsemM string

const (
	LspPlugInPluginsLv2ExpanderMsemMDown LspPlugInPluginsLv2ExpanderMsemM = "Down" // Down (0) – Down
	LspPlugInPluginsLv2ExpanderMsemMUp LspPlugInPluginsLv2ExpanderMsemM = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2ExpanderMsemS string

const (
	LspPlugInPluginsLv2ExpanderMsemSDown LspPlugInPluginsLv2ExpanderMsemS = "Down" // Down (0) – Down
	LspPlugInPluginsLv2ExpanderMsemSUp LspPlugInPluginsLv2ExpanderMsemS = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2ExpanderMsscmM string

const (
	LspPlugInPluginsLv2ExpanderMsscmMPeak LspPlugInPluginsLv2ExpanderMsscmM = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2ExpanderMsscmMRms LspPlugInPluginsLv2ExpanderMsscmM = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2ExpanderMsscmMLowPass LspPlugInPluginsLv2ExpanderMsscmM = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2ExpanderMsscmMUniform LspPlugInPluginsLv2ExpanderMsscmM = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2ExpanderMsscmS string

const (
	LspPlugInPluginsLv2ExpanderMsscmSPeak LspPlugInPluginsLv2ExpanderMsscmS = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2ExpanderMsscmSRms LspPlugInPluginsLv2ExpanderMsscmS = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2ExpanderMsscmSLowPass LspPlugInPluginsLv2ExpanderMsscmS = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2ExpanderMsscmSUniform LspPlugInPluginsLv2ExpanderMsscmS = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2ExpanderMsscsM string

const (
	LspPlugInPluginsLv2ExpanderMsscsMMiddle LspPlugInPluginsLv2ExpanderMsscsM = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2ExpanderMsscsMSide LspPlugInPluginsLv2ExpanderMsscsM = "Side" // Side (1) – Side
	LspPlugInPluginsLv2ExpanderMsscsMLeft LspPlugInPluginsLv2ExpanderMsscsM = "Left" // Left (2) – Left
	LspPlugInPluginsLv2ExpanderMsscsMRight LspPlugInPluginsLv2ExpanderMsscsM = "Right" // Right (3) – Right
)

