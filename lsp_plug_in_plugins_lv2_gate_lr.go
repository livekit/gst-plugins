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

type LspPlugInPluginsLv2GateLr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GateLr() (*LspPlugInPluginsLv2GateLr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-gate-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GateLrWithName(name string) (*LspPlugInPluginsLv2GateLr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-gate-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GateLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// at-l (float32)
//
// Attack Left

func (e *LspPlugInPluginsLv2GateLr) GetAtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetAtL(value float32) error {
	return e.Element.SetProperty("at-l", value)
}

// at-r (float32)
//
// Attack Right

func (e *LspPlugInPluginsLv2GateLr) GetAtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetAtR(value float32) error {
	return e.Element.SetProperty("at-r", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GateLr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-l (float32)
//
// Dry gain Left

func (e *LspPlugInPluginsLv2GateLr) GetCdrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetCdrL(value float32) error {
	return e.Element.SetProperty("cdr-l", value)
}

// cdr-r (float32)
//
// Dry gain Right

func (e *LspPlugInPluginsLv2GateLr) GetCdrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetCdrR(value float32) error {
	return e.Element.SetProperty("cdr-r", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2GateLr) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-l (float32)
//
// Curve level meter Left

func (e *LspPlugInPluginsLv2GateLr) GetClmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-r (float32)
//
// Curve level meter Right

func (e *LspPlugInPluginsLv2GateLr) GetClmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cwt-l (float32)
//
// Wet gain Left

func (e *LspPlugInPluginsLv2GateLr) GetCwtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetCwtL(value float32) error {
	return e.Element.SetProperty("cwt-l", value)
}

// cwt-r (float32)
//
// Wet gain Right

func (e *LspPlugInPluginsLv2GateLr) GetCwtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetCwtR(value float32) error {
	return e.Element.SetProperty("cwt-r", value)
}

// elm-l (float32)
//
// Envelope level meter Left

func (e *LspPlugInPluginsLv2GateLr) GetElmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-r (float32)
//
// Envelope level meter Right

func (e *LspPlugInPluginsLv2GateLr) GetElmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elv-l (bool)
//
// Envelope level visibility Left

func (e *LspPlugInPluginsLv2GateLr) GetElvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetElvL(value bool) error {
	return e.Element.SetProperty("elv-l", value)
}

// elv-r (bool)
//
// Envelope level visibility Right

func (e *LspPlugInPluginsLv2GateLr) GetElvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetElvR(value bool) error {
	return e.Element.SetProperty("elv-r", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2GateLr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GateLr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gh-l (bool)
//
// Hysteresis Left

func (e *LspPlugInPluginsLv2GateLr) GetGhL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gh-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGhL(value bool) error {
	return e.Element.SetProperty("gh-l", value)
}

// gh-r (bool)
//
// Hysteresis Right

func (e *LspPlugInPluginsLv2GateLr) GetGhR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gh-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGhR(value bool) error {
	return e.Element.SetProperty("gh-r", value)
}

// gr-l (float32)
//
// Reduction Left

func (e *LspPlugInPluginsLv2GateLr) GetGrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGrL(value float32) error {
	return e.Element.SetProperty("gr-l", value)
}

// gr-r (float32)
//
// Reduction Right

func (e *LspPlugInPluginsLv2GateLr) GetGrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGrR(value float32) error {
	return e.Element.SetProperty("gr-r", value)
}

// grv-l (bool)
//
// Gain reduction visibility Left

func (e *LspPlugInPluginsLv2GateLr) GetGrvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGrvL(value bool) error {
	return e.Element.SetProperty("grv-l", value)
}

// grv-r (bool)
//
// Gain reduction visibility Right

func (e *LspPlugInPluginsLv2GateLr) GetGrvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGrvR(value bool) error {
	return e.Element.SetProperty("grv-r", value)
}

// gt-l (float32)
//
// Threshold Left

func (e *LspPlugInPluginsLv2GateLr) GetGtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gt-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGtL(value float32) error {
	return e.Element.SetProperty("gt-l", value)
}

// gt-r (float32)
//
// Threshold Right

func (e *LspPlugInPluginsLv2GateLr) GetGtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGtR(value float32) error {
	return e.Element.SetProperty("gt-r", value)
}

// gz-l (float32)
//
// Zone size Left

func (e *LspPlugInPluginsLv2GateLr) GetGzL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gz-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGzL(value float32) error {
	return e.Element.SetProperty("gz-l", value)
}

// gz-r (float32)
//
// Zone size Right

func (e *LspPlugInPluginsLv2GateLr) GetGzR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gz-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetGzR(value float32) error {
	return e.Element.SetProperty("gz-r", value)
}

// gzs-l (float32)
//
// Zone start Left

func (e *LspPlugInPluginsLv2GateLr) GetGzsL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gzs-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// gzs-r (float32)
//
// Zone start Right

func (e *LspPlugInPluginsLv2GateLr) GetGzsR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gzs-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ht-l (float32)
//
// Hysteresis threshold Left

func (e *LspPlugInPluginsLv2GateLr) GetHtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ht-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetHtL(value float32) error {
	return e.Element.SetProperty("ht-l", value)
}

// ht-r (float32)
//
// Hysteresis threshold Right

func (e *LspPlugInPluginsLv2GateLr) GetHtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ht-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetHtR(value float32) error {
	return e.Element.SetProperty("ht-r", value)
}

// hts-l (float32)
//
// Hysteresis threshold start Left

func (e *LspPlugInPluginsLv2GateLr) GetHtsL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hts-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hts-r (float32)
//
// Hysteresis threshold start Right

func (e *LspPlugInPluginsLv2GateLr) GetHtsR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hts-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hz-l (float32)
//
// Hysteresis zone size Left

func (e *LspPlugInPluginsLv2GateLr) GetHzL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hz-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetHzL(value float32) error {
	return e.Element.SetProperty("hz-l", value)
}

// hz-r (float32)
//
// Hysteresis zone size Right

func (e *LspPlugInPluginsLv2GateLr) GetHzR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hz-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetHzR(value float32) error {
	return e.Element.SetProperty("hz-r", value)
}

// hzs-l (float32)
//
// Hysteresis zone start Left

func (e *LspPlugInPluginsLv2GateLr) GetHzsL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hzs-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// hzs-r (float32)
//
// Hysteresis zone start Right

func (e *LspPlugInPluginsLv2GateLr) GetHzsR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hzs-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2GateLr) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateLr) GetIlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateLr) GetIlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2GateLr) GetIlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// mk-l (float32)
//
// Makeup gain Left

func (e *LspPlugInPluginsLv2GateLr) GetMkL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetMkL(value float32) error {
	return e.Element.SetProperty("mk-l", value)
}

// mk-r (float32)
//
// Makeup gain Right

func (e *LspPlugInPluginsLv2GateLr) GetMkR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetMkR(value float32) error {
	return e.Element.SetProperty("mk-r", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2GateLr) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateLr) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2GateLr) GetOlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2GateLr) GetOlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GateLr) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GateLr) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2GateLr) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rlm-l (float32)
//
// Reduction level meter Left

func (e *LspPlugInPluginsLv2GateLr) GetRlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-r (float32)
//
// Reduction level meter Right

func (e *LspPlugInPluginsLv2GateLr) GetRlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rt-l (float32)
//
// Release Left

func (e *LspPlugInPluginsLv2GateLr) GetRtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetRtL(value float32) error {
	return e.Element.SetProperty("rt-l", value)
}

// rt-r (float32)
//
// Release Right

func (e *LspPlugInPluginsLv2GateLr) GetRtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetRtR(value float32) error {
	return e.Element.SetProperty("rt-r", value)
}

// scl-l (bool)
//
// Sidechain listen Left

func (e *LspPlugInPluginsLv2GateLr) GetSclL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetSclL(value bool) error {
	return e.Element.SetProperty("scl-l", value)
}

// scl-r (bool)
//
// Sidechain listen Right

func (e *LspPlugInPluginsLv2GateLr) GetSclR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetSclR(value bool) error {
	return e.Element.SetProperty("scl-r", value)
}

// scm-l (GstLspPlugInPluginsLv2GateLrscmL)
//
// Sidechain mode Left

func (e *LspPlugInPluginsLv2GateLr) GetScmL() (interface{}, error) {
	return e.Element.GetProperty("scm-l")
}

func (e *LspPlugInPluginsLv2GateLr) SetScmL(value interface{}) error {
	return e.Element.SetProperty("scm-l", value)
}

// scm-r (GstLspPlugInPluginsLv2GateLrscmR)
//
// Sidechain mode Right

func (e *LspPlugInPluginsLv2GateLr) GetScmR() (interface{}, error) {
	return e.Element.GetProperty("scm-r")
}

func (e *LspPlugInPluginsLv2GateLr) SetScmR(value interface{}) error {
	return e.Element.SetProperty("scm-r", value)
}

// scp-l (float32)
//
// Sidechain preamp Left

func (e *LspPlugInPluginsLv2GateLr) GetScpL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetScpL(value float32) error {
	return e.Element.SetProperty("scp-l", value)
}

// scp-r (float32)
//
// Sidechain preamp Right

func (e *LspPlugInPluginsLv2GateLr) GetScpR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetScpR(value float32) error {
	return e.Element.SetProperty("scp-r", value)
}

// scr-l (float32)
//
// Sidechain reactivity Left

func (e *LspPlugInPluginsLv2GateLr) GetScrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetScrL(value float32) error {
	return e.Element.SetProperty("scr-l", value)
}

// scr-r (float32)
//
// Sidechain reactivity Right

func (e *LspPlugInPluginsLv2GateLr) GetScrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetScrR(value float32) error {
	return e.Element.SetProperty("scr-r", value)
}

// scs-l (GstLspPlugInPluginsLv2GateLrscsL)
//
// Sidechain source Left

func (e *LspPlugInPluginsLv2GateLr) GetScsL() (interface{}, error) {
	return e.Element.GetProperty("scs-l")
}

func (e *LspPlugInPluginsLv2GateLr) SetScsL(value interface{}) error {
	return e.Element.SetProperty("scs-l", value)
}

// scs-r (GstLspPlugInPluginsLv2GateLrscsR)
//
// Sidechain source Right

func (e *LspPlugInPluginsLv2GateLr) GetScsR() (interface{}, error) {
	return e.Element.GetProperty("scs-r")
}

func (e *LspPlugInPluginsLv2GateLr) SetScsR(value interface{}) error {
	return e.Element.SetProperty("scs-r", value)
}

// sla-l (float32)
//
// Sidechain lookahead Left

func (e *LspPlugInPluginsLv2GateLr) GetSlaL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetSlaL(value float32) error {
	return e.Element.SetProperty("sla-l", value)
}

// sla-r (float32)
//
// Sidechain lookahead Right

func (e *LspPlugInPluginsLv2GateLr) GetSlaR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetSlaR(value float32) error {
	return e.Element.SetProperty("sla-r", value)
}

// slm-l (float32)
//
// Sidechain level meter Left

func (e *LspPlugInPluginsLv2GateLr) GetSlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slm-r (float32)
//
// Sidechain level meter Right

func (e *LspPlugInPluginsLv2GateLr) GetSlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slv-l (bool)
//
// Sidechain level visibility Left

func (e *LspPlugInPluginsLv2GateLr) GetSlvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetSlvL(value bool) error {
	return e.Element.SetProperty("slv-l", value)
}

// slv-r (bool)
//
// Sidechain level visibility Right

func (e *LspPlugInPluginsLv2GateLr) GetSlvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GateLr) SetSlvR(value bool) error {
	return e.Element.SetProperty("slv-r", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GateLrscsL string

const (
	LspPlugInPluginsLv2GateLrscsLMiddle LspPlugInPluginsLv2GateLrscsL = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2GateLrscsLSide LspPlugInPluginsLv2GateLrscsL = "Side" // Side (1) – Side
	LspPlugInPluginsLv2GateLrscsLLeft LspPlugInPluginsLv2GateLrscsL = "Left" // Left (2) – Left
	LspPlugInPluginsLv2GateLrscsLRight LspPlugInPluginsLv2GateLrscsL = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2GateLrscsR string

const (
	LspPlugInPluginsLv2GateLrscsRMiddle LspPlugInPluginsLv2GateLrscsR = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2GateLrscsRSide LspPlugInPluginsLv2GateLrscsR = "Side" // Side (1) – Side
	LspPlugInPluginsLv2GateLrscsRLeft LspPlugInPluginsLv2GateLrscsR = "Left" // Left (2) – Left
	LspPlugInPluginsLv2GateLrscsRRight LspPlugInPluginsLv2GateLrscsR = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2GateLrscmL string

const (
	LspPlugInPluginsLv2GateLrscmLPeak LspPlugInPluginsLv2GateLrscmL = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2GateLrscmLRms LspPlugInPluginsLv2GateLrscmL = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2GateLrscmLLowPass LspPlugInPluginsLv2GateLrscmL = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2GateLrscmLUniform LspPlugInPluginsLv2GateLrscmL = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2GateLrscmR string

const (
	LspPlugInPluginsLv2GateLrscmRPeak LspPlugInPluginsLv2GateLrscmR = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2GateLrscmRRms LspPlugInPluginsLv2GateLrscmR = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2GateLrscmRLowPass LspPlugInPluginsLv2GateLrscmR = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2GateLrscmRUniform LspPlugInPluginsLv2GateLrscmR = "Uniform" // Uniform (3) – Uniform
)

