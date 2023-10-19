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

type LspPlugInPluginsLv2GateMs struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GateMs() (*LspPlugInPluginsLv2GateMs, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-gate-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GateMsWithName(name string) (*LspPlugInPluginsLv2GateMs, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-gate-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// at-m (float32)
//
// Attack Mid

func (e *LspPlugInPluginsLv2GateMs) GetAtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetAtM(value float32) error {
	return e.Element.SetProperty("at-m", value)
}

// at-s (float32)
//
// Attack Side

func (e *LspPlugInPluginsLv2GateMs) GetAtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetAtS(value float32) error {
	return e.Element.SetProperty("at-s", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GateMs) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-m (float32)
//
// Dry gain Mid

func (e *LspPlugInPluginsLv2GateMs) GetCdrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetCdrM(value float32) error {
	return e.Element.SetProperty("cdr-m", value)
}

// cdr-s (float32)
//
// Dry gain Side

func (e *LspPlugInPluginsLv2GateMs) GetCdrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetCdrS(value float32) error {
	return e.Element.SetProperty("cdr-s", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2GateMs) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-m (float32)
//
// Curve level meter Mid

func (e *LspPlugInPluginsLv2GateMs) GetClmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetClmS() (float32, error) {
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

// cwt-m (float32)
//
// Wet gain Mid

func (e *LspPlugInPluginsLv2GateMs) GetCwtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetCwtM(value float32) error {
	return e.Element.SetProperty("cwt-m", value)
}

// cwt-s (float32)
//
// Wet gain Side

func (e *LspPlugInPluginsLv2GateMs) GetCwtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetCwtS(value float32) error {
	return e.Element.SetProperty("cwt-s", value)
}

// elm-m (float32)
//
// Envelope level meter Mid

func (e *LspPlugInPluginsLv2GateMs) GetElmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetElmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetElvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetElvM(value bool) error {
	return e.Element.SetProperty("elv-m", value)
}

// elv-s (bool)
//
// Envelope level visibility Side

func (e *LspPlugInPluginsLv2GateMs) GetElvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetElvS(value bool) error {
	return e.Element.SetProperty("elv-s", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2GateMs) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GateMs) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gh-m (bool)
//
// Hysteresis Mid

func (e *LspPlugInPluginsLv2GateMs) GetGhM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gh-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGhM(value bool) error {
	return e.Element.SetProperty("gh-m", value)
}

// gh-s (bool)
//
// Hysteresis Side

func (e *LspPlugInPluginsLv2GateMs) GetGhS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gh-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGhS(value bool) error {
	return e.Element.SetProperty("gh-s", value)
}

// gr-m (float32)
//
// Reduction Mid

func (e *LspPlugInPluginsLv2GateMs) GetGrM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGrM(value float32) error {
	return e.Element.SetProperty("gr-m", value)
}

// gr-s (float32)
//
// Reduction Side

func (e *LspPlugInPluginsLv2GateMs) GetGrS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGrS(value float32) error {
	return e.Element.SetProperty("gr-s", value)
}

// grv-m (bool)
//
// Gain reduction visibility Mid

func (e *LspPlugInPluginsLv2GateMs) GetGrvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetGrvM(value bool) error {
	return e.Element.SetProperty("grv-m", value)
}

// grv-s (bool)
//
// Gain reduction visibility Side

func (e *LspPlugInPluginsLv2GateMs) GetGrvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetGrvS(value bool) error {
	return e.Element.SetProperty("grv-s", value)
}

// gt-m (float32)
//
// Threshold Mid

func (e *LspPlugInPluginsLv2GateMs) GetGtM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gt-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGtM(value float32) error {
	return e.Element.SetProperty("gt-m", value)
}

// gt-s (float32)
//
// Threshold Side

func (e *LspPlugInPluginsLv2GateMs) GetGtS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gt-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGtS(value float32) error {
	return e.Element.SetProperty("gt-s", value)
}

// gz-m (float32)
//
// Zone size Mid

func (e *LspPlugInPluginsLv2GateMs) GetGzM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gz-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGzM(value float32) error {
	return e.Element.SetProperty("gz-m", value)
}

// gz-s (float32)
//
// Zone size Side

func (e *LspPlugInPluginsLv2GateMs) GetGzS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gz-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetGzS(value float32) error {
	return e.Element.SetProperty("gz-s", value)
}

// gzs-m (float32)
//
// Zone start Mid

func (e *LspPlugInPluginsLv2GateMs) GetGzsM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gzs-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// gzs-s (float32)
//
// Zone start Side

func (e *LspPlugInPluginsLv2GateMs) GetGzsS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gzs-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ht-m (float32)
//
// Hysteresis threshold Mid

func (e *LspPlugInPluginsLv2GateMs) GetHtM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ht-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetHtM(value float32) error {
	return e.Element.SetProperty("ht-m", value)
}

// ht-s (float32)
//
// Hysteresis threshold Side

func (e *LspPlugInPluginsLv2GateMs) GetHtS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ht-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetHtS(value float32) error {
	return e.Element.SetProperty("ht-s", value)
}

// hts-m (float32)
//
// Hysteresis threshold start Mid

func (e *LspPlugInPluginsLv2GateMs) GetHtsM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hts-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hts-s (float32)
//
// Hysteresis threshold start Side

func (e *LspPlugInPluginsLv2GateMs) GetHtsS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hts-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hz-m (float32)
//
// Hysteresis zone size Mid

func (e *LspPlugInPluginsLv2GateMs) GetHzM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hz-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetHzM(value float32) error {
	return e.Element.SetProperty("hz-m", value)
}

// hz-s (float32)
//
// Hysteresis zone size Side

func (e *LspPlugInPluginsLv2GateMs) GetHzS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hz-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateMs) SetHzS(value float32) error {
	return e.Element.SetProperty("hz-s", value)
}

// hzs-m (float32)
//
// Hysteresis zone start Mid

func (e *LspPlugInPluginsLv2GateMs) GetHzsM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hzs-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hzs-s (float32)
//
// Hysteresis zone start Side

func (e *LspPlugInPluginsLv2GateMs) GetHzsS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hzs-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm-m (float32)
//
// Input level meter Mid

func (e *LspPlugInPluginsLv2GateMs) GetIlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetIlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetIlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetIlvM(value bool) error {
	return e.Element.SetProperty("ilv-m", value)
}

// ilv-s (bool)
//
// Input level visibility Side

func (e *LspPlugInPluginsLv2GateMs) GetIlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetIlvS(value bool) error {
	return e.Element.SetProperty("ilv-s", value)
}

// mk-m (float32)
//
// Makeup gain Mid

func (e *LspPlugInPluginsLv2GateMs) GetMkM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetMkM(value float32) error {
	return e.Element.SetProperty("mk-m", value)
}

// mk-s (float32)
//
// Makeup gain Side

func (e *LspPlugInPluginsLv2GateMs) GetMkS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetMkS(value float32) error {
	return e.Element.SetProperty("mk-s", value)
}

// msl (bool)
//
// Mid/Side listen

func (e *LspPlugInPluginsLv2GateMs) GetMsl() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetMsl(value bool) error {
	return e.Element.SetProperty("msl", value)
}

// olm-m (float32)
//
// Output level meter Mid

func (e *LspPlugInPluginsLv2GateMs) GetOlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetOlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetOlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetOlvM(value bool) error {
	return e.Element.SetProperty("olv-m", value)
}

// olv-s (bool)
//
// Output level visibility Side

func (e *LspPlugInPluginsLv2GateMs) GetOlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetOlvS(value bool) error {
	return e.Element.SetProperty("olv-s", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GateMs) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rlm-m (float32)
//
// Reduction level meter Mid

func (e *LspPlugInPluginsLv2GateMs) GetRlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetRlmS() (float32, error) {
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

// rt-m (float32)
//
// Release Mid

func (e *LspPlugInPluginsLv2GateMs) GetRtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetRtM(value float32) error {
	return e.Element.SetProperty("rt-m", value)
}

// rt-s (float32)
//
// Release Side

func (e *LspPlugInPluginsLv2GateMs) GetRtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetRtS(value float32) error {
	return e.Element.SetProperty("rt-s", value)
}

// scl-m (bool)
//
// Sidechain listen Mid

func (e *LspPlugInPluginsLv2GateMs) GetSclM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetSclM(value bool) error {
	return e.Element.SetProperty("scl-m", value)
}

// scl-s (bool)
//
// Sidechain listen Side

func (e *LspPlugInPluginsLv2GateMs) GetSclS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetSclS(value bool) error {
	return e.Element.SetProperty("scl-s", value)
}

// scm-m (GstLspPlugInPluginsLv2GateMsscmM)
//
// Sidechain mode Mid

func (e *LspPlugInPluginsLv2GateMs) GetScmM() (interface{}, error) {
	return e.Element.GetProperty("scm-m")
}

func (e *LspPlugInPluginsLv2GateMs) SetScmM(value interface{}) error {
	return e.Element.SetProperty("scm-m", value)
}

// scm-s (GstLspPlugInPluginsLv2GateMsscmS)
//
// Sidechain mode Side

func (e *LspPlugInPluginsLv2GateMs) GetScmS() (interface{}, error) {
	return e.Element.GetProperty("scm-s")
}

func (e *LspPlugInPluginsLv2GateMs) SetScmS(value interface{}) error {
	return e.Element.SetProperty("scm-s", value)
}

// scp-m (float32)
//
// Sidechain preamp Mid

func (e *LspPlugInPluginsLv2GateMs) GetScpM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetScpM(value float32) error {
	return e.Element.SetProperty("scp-m", value)
}

// scp-s (float32)
//
// Sidechain preamp Side

func (e *LspPlugInPluginsLv2GateMs) GetScpS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetScpS(value float32) error {
	return e.Element.SetProperty("scp-s", value)
}

// scr-m (float32)
//
// Sidechain reactivity Mid

func (e *LspPlugInPluginsLv2GateMs) GetScrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetScrM(value float32) error {
	return e.Element.SetProperty("scr-m", value)
}

// scr-s (float32)
//
// Sidechain reactivity Side

func (e *LspPlugInPluginsLv2GateMs) GetScrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetScrS(value float32) error {
	return e.Element.SetProperty("scr-s", value)
}

// scs-m (GstLspPlugInPluginsLv2GateMsscsM)
//
// Sidechain source Mid

func (e *LspPlugInPluginsLv2GateMs) GetScsM() (interface{}, error) {
	return e.Element.GetProperty("scs-m")
}

func (e *LspPlugInPluginsLv2GateMs) SetScsM(value interface{}) error {
	return e.Element.SetProperty("scs-m", value)
}

// scs-s (GstLspPlugInPluginsLv2GateMsscsS)
//
// Sidechain source Side

func (e *LspPlugInPluginsLv2GateMs) GetScsS() (interface{}, error) {
	return e.Element.GetProperty("scs-s")
}

func (e *LspPlugInPluginsLv2GateMs) SetScsS(value interface{}) error {
	return e.Element.SetProperty("scs-s", value)
}

// sla-m (float32)
//
// Sidechain lookahead Mid

func (e *LspPlugInPluginsLv2GateMs) GetSlaM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetSlaM(value float32) error {
	return e.Element.SetProperty("sla-m", value)
}

// sla-s (float32)
//
// Sidechain lookahead Side

func (e *LspPlugInPluginsLv2GateMs) GetSlaS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetSlaS(value float32) error {
	return e.Element.SetProperty("sla-s", value)
}

// slm-m (float32)
//
// Sidechain level meter Mid

func (e *LspPlugInPluginsLv2GateMs) GetSlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetSlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateMs) GetSlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetSlvM(value bool) error {
	return e.Element.SetProperty("slv-m", value)
}

// slv-s (bool)
//
// Sidechain level visibility Side

func (e *LspPlugInPluginsLv2GateMs) GetSlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateMs) SetSlvS(value bool) error {
	return e.Element.SetProperty("slv-s", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GateMsscmS string

const (
	LspPlugInPluginsLv2GateMsscmSPeak LspPlugInPluginsLv2GateMsscmS = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2GateMsscmSRms LspPlugInPluginsLv2GateMsscmS = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2GateMsscmSLowPass LspPlugInPluginsLv2GateMsscmS = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2GateMsscmSUniform LspPlugInPluginsLv2GateMsscmS = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2GateMsscsM string

const (
	LspPlugInPluginsLv2GateMsscsMMiddle LspPlugInPluginsLv2GateMsscsM = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2GateMsscsMSide LspPlugInPluginsLv2GateMsscsM = "Side" // Side (1) – Side
	LspPlugInPluginsLv2GateMsscsMLeft LspPlugInPluginsLv2GateMsscsM = "Left" // Left (2) – Left
	LspPlugInPluginsLv2GateMsscsMRight LspPlugInPluginsLv2GateMsscsM = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2GateMsscsS string

const (
	LspPlugInPluginsLv2GateMsscsSMiddle LspPlugInPluginsLv2GateMsscsS = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2GateMsscsSSide LspPlugInPluginsLv2GateMsscsS = "Side" // Side (1) – Side
	LspPlugInPluginsLv2GateMsscsSLeft LspPlugInPluginsLv2GateMsscsS = "Left" // Left (2) – Left
	LspPlugInPluginsLv2GateMsscsSRight LspPlugInPluginsLv2GateMsscsS = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2GateMsscmM string

const (
	LspPlugInPluginsLv2GateMsscmMPeak LspPlugInPluginsLv2GateMsscmM = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2GateMsscmMRms LspPlugInPluginsLv2GateMsscmM = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2GateMsscmMLowPass LspPlugInPluginsLv2GateMsscmM = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2GateMsscmMUniform LspPlugInPluginsLv2GateMsscmM = "Uniform" // Uniform (3) – Uniform
)

