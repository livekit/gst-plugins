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

type LspPlugInPluginsLv2DynaProcessorStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2DynaProcessorStereo() (*LspPlugInPluginsLv2DynaProcessorStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-dyna-processor-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2DynaProcessorStereoWithName(name string) (*LspPlugInPluginsLv2DynaProcessorStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-dyna-processor-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// ae0 (bool)
//
// Attack enable 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAe0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAe0(value bool) error {
	return e.Element.SetProperty("ae0", value)
}

// ae1 (bool)
//
// Attack enable 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAe1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAe1(value bool) error {
	return e.Element.SetProperty("ae1", value)
}

// ae2 (bool)
//
// Attack enable 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAe2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAe2(value bool) error {
	return e.Element.SetProperty("ae2", value)
}

// ae3 (bool)
//
// Attack enable 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAe3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAe3(value bool) error {
	return e.Element.SetProperty("ae3", value)
}

// al0 (float32)
//
// Attack level 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAl0(value float32) error {
	return e.Element.SetProperty("al0", value)
}

// al1 (float32)
//
// Attack level 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAl1(value float32) error {
	return e.Element.SetProperty("al1", value)
}

// al2 (float32)
//
// Attack level 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAl2(value float32) error {
	return e.Element.SetProperty("al2", value)
}

// al3 (float32)
//
// Attack level 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAl3(value float32) error {
	return e.Element.SetProperty("al3", value)
}

// at0 (float32)
//
// Attack time 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAt0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAt0(value float32) error {
	return e.Element.SetProperty("at0", value)
}

// at1 (float32)
//
// Attack time 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAt1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAt1(value float32) error {
	return e.Element.SetProperty("at1", value)
}

// at2 (float32)
//
// Attack time 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAt2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAt2(value float32) error {
	return e.Element.SetProperty("at2", value)
}

// at3 (float32)
//
// Attack time 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAt3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAt3(value float32) error {
	return e.Element.SetProperty("at3", value)
}

// atd (float32)
//
// Attack time default

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetAtd() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("atd")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetAtd(value float32) error {
	return e.Element.SetProperty("atd", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetClm() (float32, error) {
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

// cmv (bool)
//
// Curve modelling visibility

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetCmv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cmv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetCmv(value bool) error {
	return e.Element.SetProperty("cmv", value)
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl0 (float32)
//
// Gain 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGl0(value float32) error {
	return e.Element.SetProperty("gl0", value)
}

// gl1 (float32)
//
// Gain 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGl1(value float32) error {
	return e.Element.SetProperty("gl1", value)
}

// gl2 (float32)
//
// Gain 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGl2(value float32) error {
	return e.Element.SetProperty("gl2", value)
}

// gl3 (float32)
//
// Gain 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGl3(value float32) error {
	return e.Element.SetProperty("gl3", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// hlr (float32)
//
// High-level ratio

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetHlr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hlr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetHlr(value float32) error {
	return e.Element.SetProperty("hlr", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetIlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetIlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetIlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// kn0 (float32)
//
// Knee 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetKn0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetKn0(value float32) error {
	return e.Element.SetProperty("kn0", value)
}

// kn1 (float32)
//
// Knee 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetKn1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetKn1(value float32) error {
	return e.Element.SetProperty("kn1", value)
}

// kn2 (float32)
//
// Knee 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetKn2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetKn2(value float32) error {
	return e.Element.SetProperty("kn2", value)
}

// kn3 (float32)
//
// Knee 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetKn3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetKn3(value float32) error {
	return e.Element.SetProperty("kn3", value)
}

// llr (float32)
//
// Low-level ratio

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetLlr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("llr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetLlr(value float32) error {
	return e.Element.SetProperty("llr", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetOlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetOlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// omk (float32)
//
// Overall makeup gain

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetOmk() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("omk")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetOmk(value float32) error {
	return e.Element.SetProperty("omk", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// pe0 (bool)
//
// Point enable 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetPe0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetPe0(value bool) error {
	return e.Element.SetProperty("pe0", value)
}

// pe1 (bool)
//
// Point enable 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetPe1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetPe1(value bool) error {
	return e.Element.SetProperty("pe1", value)
}

// pe2 (bool)
//
// Point enable 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetPe2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetPe2(value bool) error {
	return e.Element.SetProperty("pe2", value)
}

// pe3 (bool)
//
// Point enable 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetPe3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetPe3(value bool) error {
	return e.Element.SetProperty("pe3", value)
}

// re0 (bool)
//
// Release enable 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRe0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRe0(value bool) error {
	return e.Element.SetProperty("re0", value)
}

// re1 (bool)
//
// Release enable 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRe1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRe1(value bool) error {
	return e.Element.SetProperty("re1", value)
}

// re2 (bool)
//
// Release enable 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRe2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRe2(value bool) error {
	return e.Element.SetProperty("re2", value)
}

// re3 (bool)
//
// Release enable 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRe3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRe3(value bool) error {
	return e.Element.SetProperty("re3", value)
}

// rl0 (float32)
//
// Relative level 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRl0(value float32) error {
	return e.Element.SetProperty("rl0", value)
}

// rl1 (float32)
//
// Relative level 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRl1(value float32) error {
	return e.Element.SetProperty("rl1", value)
}

// rl2 (float32)
//
// Relative level 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRl2(value float32) error {
	return e.Element.SetProperty("rl2", value)
}

// rl3 (float32)
//
// Relative level 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRl3(value float32) error {
	return e.Element.SetProperty("rl3", value)
}

// rlm (float32)
//
// Reduction level meter

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRlm() (float32, error) {
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

// rt0 (float32)
//
// Release time 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRt0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRt0(value float32) error {
	return e.Element.SetProperty("rt0", value)
}

// rt1 (float32)
//
// Release time 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRt1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRt1(value float32) error {
	return e.Element.SetProperty("rt1", value)
}

// rt2 (float32)
//
// Release time 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRt2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRt2(value float32) error {
	return e.Element.SetProperty("rt2", value)
}

// rt3 (float32)
//
// Release time 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRt3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRt3(value float32) error {
	return e.Element.SetProperty("rt3", value)
}

// rtd (float32)
//
// Release time default

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetRtd() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rtd")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetRtd(value float32) error {
	return e.Element.SetProperty("rtd", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2DynaProcessorStereoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// scs (GstLspPlugInPluginsLv2DynaProcessorStereoscs)
//
// Sidechain source

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetScs() (interface{}, error) {
	return e.Element.GetProperty("scs")
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetScs(value interface{}) error {
	return e.Element.SetProperty("scs", value)
}

// sct (GstLspPlugInPluginsLv2DynaProcessorStereosct)
//
// Sidechain type

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetSct() (interface{}, error) {
	return e.Element.GetProperty("sct")
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetSct(value interface{}) error {
	return e.Element.SetProperty("sct", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// tl0 (float32)
//
// Threshold 0

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetTl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetTl0(value float32) error {
	return e.Element.SetProperty("tl0", value)
}

// tl1 (float32)
//
// Threshold 1

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetTl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetTl1(value float32) error {
	return e.Element.SetProperty("tl1", value)
}

// tl2 (float32)
//
// Threshold 2

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetTl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetTl2(value float32) error {
	return e.Element.SetProperty("tl2", value)
}

// tl3 (float32)
//
// Threshold 3

func (e *LspPlugInPluginsLv2DynaProcessorStereo) GetTl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorStereo) SetTl3(value float32) error {
	return e.Element.SetProperty("tl3", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2DynaProcessorStereosct string

const (
	LspPlugInPluginsLv2DynaProcessorStereosctFeedForward LspPlugInPluginsLv2DynaProcessorStereosct = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2DynaProcessorStereosctFeedBack LspPlugInPluginsLv2DynaProcessorStereosct = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2DynaProcessorStereoscm string

const (
	LspPlugInPluginsLv2DynaProcessorStereoscmPeak LspPlugInPluginsLv2DynaProcessorStereoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2DynaProcessorStereoscmRms LspPlugInPluginsLv2DynaProcessorStereoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2DynaProcessorStereoscmLowPass LspPlugInPluginsLv2DynaProcessorStereoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2DynaProcessorStereoscmUniform LspPlugInPluginsLv2DynaProcessorStereoscm = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2DynaProcessorStereoscs string

const (
	LspPlugInPluginsLv2DynaProcessorStereoscsMiddle LspPlugInPluginsLv2DynaProcessorStereoscs = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2DynaProcessorStereoscsSide LspPlugInPluginsLv2DynaProcessorStereoscs = "Side" // Side (1) – Side
	LspPlugInPluginsLv2DynaProcessorStereoscsLeft LspPlugInPluginsLv2DynaProcessorStereoscs = "Left" // Left (2) – Left
	LspPlugInPluginsLv2DynaProcessorStereoscsRight LspPlugInPluginsLv2DynaProcessorStereoscs = "Right" // Right (3) – Right
)

