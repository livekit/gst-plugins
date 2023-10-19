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

type LspPlugInPluginsLv2CompressorLr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2CompressorLr() (*LspPlugInPluginsLv2CompressorLr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-compressor-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2CompressorLrWithName(name string) (*LspPlugInPluginsLv2CompressorLr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-compressor-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2CompressorLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-l (float32)
//
// Attack level Left

func (e *LspPlugInPluginsLv2CompressorLr) GetAlL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetAlL(value float32) error {
	return e.Element.SetProperty("al-l", value)
}

// al-r (float32)
//
// Attack level Right

func (e *LspPlugInPluginsLv2CompressorLr) GetAlR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetAlR(value float32) error {
	return e.Element.SetProperty("al-r", value)
}

// at-l (float32)
//
// Attack time Left

func (e *LspPlugInPluginsLv2CompressorLr) GetAtL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetAtL(value float32) error {
	return e.Element.SetProperty("at-l", value)
}

// at-r (float32)
//
// Attack time Right

func (e *LspPlugInPluginsLv2CompressorLr) GetAtR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetAtR(value float32) error {
	return e.Element.SetProperty("at-r", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2CompressorLr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-l (float32)
//
// Dry gain Left

func (e *LspPlugInPluginsLv2CompressorLr) GetCdrL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetCdrL(value float32) error {
	return e.Element.SetProperty("cdr-l", value)
}

// cdr-r (float32)
//
// Dry gain Right

func (e *LspPlugInPluginsLv2CompressorLr) GetCdrR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetCdrR(value float32) error {
	return e.Element.SetProperty("cdr-r", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2CompressorLr) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-l (float32)
//
// Curve level meter Left

func (e *LspPlugInPluginsLv2CompressorLr) GetClmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetClmR() (float32, error) {
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

// cm-l (GstLspPlugInPluginsLv2CompressorLrcmL)
//
// Compression mode Left

func (e *LspPlugInPluginsLv2CompressorLr) GetCmL() (interface{}, error) {
	return e.Element.GetProperty("cm-l")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetCmL(value interface{}) error {
	return e.Element.SetProperty("cm-l", value)
}

// cm-r (GstLspPlugInPluginsLv2CompressorLrcmR)
//
// Compression mode Right

func (e *LspPlugInPluginsLv2CompressorLr) GetCmR() (interface{}, error) {
	return e.Element.GetProperty("cm-r")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetCmR(value interface{}) error {
	return e.Element.SetProperty("cm-r", value)
}

// cr-l (float32)
//
// Ratio Left

func (e *LspPlugInPluginsLv2CompressorLr) GetCrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetCrL(value float32) error {
	return e.Element.SetProperty("cr-l", value)
}

// cr-r (float32)
//
// Ratio Right

func (e *LspPlugInPluginsLv2CompressorLr) GetCrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetCrR(value float32) error {
	return e.Element.SetProperty("cr-r", value)
}

// cwt-l (float32)
//
// Wet gain Left

func (e *LspPlugInPluginsLv2CompressorLr) GetCwtL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetCwtL(value float32) error {
	return e.Element.SetProperty("cwt-l", value)
}

// cwt-r (float32)
//
// Wet gain Right

func (e *LspPlugInPluginsLv2CompressorLr) GetCwtR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetCwtR(value float32) error {
	return e.Element.SetProperty("cwt-r", value)
}

// elm-l (float32)
//
// Envelope level meter Left

func (e *LspPlugInPluginsLv2CompressorLr) GetElmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetElmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetElvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetElvL(value bool) error {
	return e.Element.SetProperty("elv-l", value)
}

// elv-r (bool)
//
// Envelope level visibility Right

func (e *LspPlugInPluginsLv2CompressorLr) GetElvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetElvR(value bool) error {
	return e.Element.SetProperty("elv-r", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2CompressorLr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2CompressorLr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv-l (bool)
//
// Gain reduction visibility Left

func (e *LspPlugInPluginsLv2CompressorLr) GetGrvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetGrvL(value bool) error {
	return e.Element.SetProperty("grv-l", value)
}

// grv-r (bool)
//
// Gain reduction visibility Right

func (e *LspPlugInPluginsLv2CompressorLr) GetGrvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetGrvR(value bool) error {
	return e.Element.SetProperty("grv-r", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2CompressorLr) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetIlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetIlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2CompressorLr) GetIlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// kn-l (float32)
//
// Knee Left

func (e *LspPlugInPluginsLv2CompressorLr) GetKnL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetKnL(value float32) error {
	return e.Element.SetProperty("kn-l", value)
}

// kn-r (float32)
//
// Knee Right

func (e *LspPlugInPluginsLv2CompressorLr) GetKnR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetKnR(value float32) error {
	return e.Element.SetProperty("kn-r", value)
}

// mk-l (float32)
//
// Makeup gain Left

func (e *LspPlugInPluginsLv2CompressorLr) GetMkL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetMkL(value float32) error {
	return e.Element.SetProperty("mk-l", value)
}

// mk-r (float32)
//
// Makeup gain Right

func (e *LspPlugInPluginsLv2CompressorLr) GetMkR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetMkR(value float32) error {
	return e.Element.SetProperty("mk-r", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2CompressorLr) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetOlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2CompressorLr) GetOlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2CompressorLr) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl-l (float32)
//
// Release level Left

func (e *LspPlugInPluginsLv2CompressorLr) GetRlL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-r (float32)
//
// Release level Right

func (e *LspPlugInPluginsLv2CompressorLr) GetRlR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-l (float32)
//
// Reduction level meter Left

func (e *LspPlugInPluginsLv2CompressorLr) GetRlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetRlmR() (float32, error) {
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

// rrl-l (float32)
//
// Relative release level Left

func (e *LspPlugInPluginsLv2CompressorLr) GetRrlL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetRrlL(value float32) error {
	return e.Element.SetProperty("rrl-l", value)
}

// rrl-r (float32)
//
// Relative release level Right

func (e *LspPlugInPluginsLv2CompressorLr) GetRrlR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2CompressorLr) SetRrlR(value float32) error {
	return e.Element.SetProperty("rrl-r", value)
}

// rt-l (float32)
//
// Release time Left

func (e *LspPlugInPluginsLv2CompressorLr) GetRtL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetRtL(value float32) error {
	return e.Element.SetProperty("rt-l", value)
}

// rt-r (float32)
//
// Release time Right

func (e *LspPlugInPluginsLv2CompressorLr) GetRtR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetRtR(value float32) error {
	return e.Element.SetProperty("rt-r", value)
}

// scl-l (bool)
//
// Sidechain listen Left

func (e *LspPlugInPluginsLv2CompressorLr) GetSclL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetSclL(value bool) error {
	return e.Element.SetProperty("scl-l", value)
}

// scl-r (bool)
//
// Sidechain listen Right

func (e *LspPlugInPluginsLv2CompressorLr) GetSclR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetSclR(value bool) error {
	return e.Element.SetProperty("scl-r", value)
}

// scm-l (GstLspPlugInPluginsLv2CompressorLrscmL)
//
// Sidechain mode Left

func (e *LspPlugInPluginsLv2CompressorLr) GetScmL() (interface{}, error) {
	return e.Element.GetProperty("scm-l")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetScmL(value interface{}) error {
	return e.Element.SetProperty("scm-l", value)
}

// scm-r (GstLspPlugInPluginsLv2CompressorLrscmR)
//
// Sidechain mode Right

func (e *LspPlugInPluginsLv2CompressorLr) GetScmR() (interface{}, error) {
	return e.Element.GetProperty("scm-r")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetScmR(value interface{}) error {
	return e.Element.SetProperty("scm-r", value)
}

// scp-l (float32)
//
// Sidechain preamp Left

func (e *LspPlugInPluginsLv2CompressorLr) GetScpL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetScpL(value float32) error {
	return e.Element.SetProperty("scp-l", value)
}

// scp-r (float32)
//
// Sidechain preamp Right

func (e *LspPlugInPluginsLv2CompressorLr) GetScpR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetScpR(value float32) error {
	return e.Element.SetProperty("scp-r", value)
}

// scr-l (float32)
//
// Sidechain reactivity Left

func (e *LspPlugInPluginsLv2CompressorLr) GetScrL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetScrL(value float32) error {
	return e.Element.SetProperty("scr-l", value)
}

// scr-r (float32)
//
// Sidechain reactivity Right

func (e *LspPlugInPluginsLv2CompressorLr) GetScrR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetScrR(value float32) error {
	return e.Element.SetProperty("scr-r", value)
}

// scs-l (GstLspPlugInPluginsLv2CompressorLrscsL)
//
// Sidechain source Left

func (e *LspPlugInPluginsLv2CompressorLr) GetScsL() (interface{}, error) {
	return e.Element.GetProperty("scs-l")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetScsL(value interface{}) error {
	return e.Element.SetProperty("scs-l", value)
}

// scs-r (GstLspPlugInPluginsLv2CompressorLrscsR)
//
// Sidechain source Right

func (e *LspPlugInPluginsLv2CompressorLr) GetScsR() (interface{}, error) {
	return e.Element.GetProperty("scs-r")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetScsR(value interface{}) error {
	return e.Element.SetProperty("scs-r", value)
}

// sct-l (GstLspPlugInPluginsLv2CompressorLrsctL)
//
// Sidechain type Left

func (e *LspPlugInPluginsLv2CompressorLr) GetSctL() (interface{}, error) {
	return e.Element.GetProperty("sct-l")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetSctL(value interface{}) error {
	return e.Element.SetProperty("sct-l", value)
}

// sct-r (GstLspPlugInPluginsLv2CompressorLrsctR)
//
// Sidechain type Right

func (e *LspPlugInPluginsLv2CompressorLr) GetSctR() (interface{}, error) {
	return e.Element.GetProperty("sct-r")
}

func (e *LspPlugInPluginsLv2CompressorLr) SetSctR(value interface{}) error {
	return e.Element.SetProperty("sct-r", value)
}

// sla-l (float32)
//
// Sidechain lookahead Left

func (e *LspPlugInPluginsLv2CompressorLr) GetSlaL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetSlaL(value float32) error {
	return e.Element.SetProperty("sla-l", value)
}

// sla-r (float32)
//
// Sidechain lookahead Right

func (e *LspPlugInPluginsLv2CompressorLr) GetSlaR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetSlaR(value float32) error {
	return e.Element.SetProperty("sla-r", value)
}

// slm-l (float32)
//
// Sidechain level meter Left

func (e *LspPlugInPluginsLv2CompressorLr) GetSlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetSlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) GetSlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetSlvL(value bool) error {
	return e.Element.SetProperty("slv-l", value)
}

// slv-r (bool)
//
// Sidechain level visibility Right

func (e *LspPlugInPluginsLv2CompressorLr) GetSlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2CompressorLr) SetSlvR(value bool) error {
	return e.Element.SetProperty("slv-r", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2CompressorLrscmR string

const (
	LspPlugInPluginsLv2CompressorLrscmRPeak LspPlugInPluginsLv2CompressorLrscmR = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2CompressorLrscmRRms LspPlugInPluginsLv2CompressorLrscmR = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2CompressorLrscmRLowPass LspPlugInPluginsLv2CompressorLrscmR = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2CompressorLrscmRUniform LspPlugInPluginsLv2CompressorLrscmR = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2CompressorLrscsL string

const (
	LspPlugInPluginsLv2CompressorLrscsLMiddle LspPlugInPluginsLv2CompressorLrscsL = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2CompressorLrscsLSide LspPlugInPluginsLv2CompressorLrscsL = "Side" // Side (1) – Side
	LspPlugInPluginsLv2CompressorLrscsLLeft LspPlugInPluginsLv2CompressorLrscsL = "Left" // Left (2) – Left
	LspPlugInPluginsLv2CompressorLrscsLRight LspPlugInPluginsLv2CompressorLrscsL = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2CompressorLrscsR string

const (
	LspPlugInPluginsLv2CompressorLrscsRMiddle LspPlugInPluginsLv2CompressorLrscsR = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2CompressorLrscsRSide LspPlugInPluginsLv2CompressorLrscsR = "Side" // Side (1) – Side
	LspPlugInPluginsLv2CompressorLrscsRLeft LspPlugInPluginsLv2CompressorLrscsR = "Left" // Left (2) – Left
	LspPlugInPluginsLv2CompressorLrscsRRight LspPlugInPluginsLv2CompressorLrscsR = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2CompressorLrsctL string

const (
	LspPlugInPluginsLv2CompressorLrsctLFeedForward LspPlugInPluginsLv2CompressorLrsctL = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2CompressorLrsctLFeedBack LspPlugInPluginsLv2CompressorLrsctL = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2CompressorLrsctR string

const (
	LspPlugInPluginsLv2CompressorLrsctRFeedForward LspPlugInPluginsLv2CompressorLrsctR = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2CompressorLrsctRFeedBack LspPlugInPluginsLv2CompressorLrsctR = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2CompressorLrcmL string

const (
	LspPlugInPluginsLv2CompressorLrcmLDown LspPlugInPluginsLv2CompressorLrcmL = "Down" // Down (0) – Down
	LspPlugInPluginsLv2CompressorLrcmLUp LspPlugInPluginsLv2CompressorLrcmL = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2CompressorLrcmR string

const (
	LspPlugInPluginsLv2CompressorLrcmRDown LspPlugInPluginsLv2CompressorLrcmR = "Down" // Down (0) – Down
	LspPlugInPluginsLv2CompressorLrcmRUp LspPlugInPluginsLv2CompressorLrcmR = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2CompressorLrscmL string

const (
	LspPlugInPluginsLv2CompressorLrscmLPeak LspPlugInPluginsLv2CompressorLrscmL = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2CompressorLrscmLRms LspPlugInPluginsLv2CompressorLrscmL = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2CompressorLrscmLLowPass LspPlugInPluginsLv2CompressorLrscmL = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2CompressorLrscmLUniform LspPlugInPluginsLv2CompressorLrscmL = "Uniform" // Uniform (3) – Uniform
)

