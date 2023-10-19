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

type LspPlugInPluginsLv2DynaProcessorMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2DynaProcessorMono() (*LspPlugInPluginsLv2DynaProcessorMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-dyna-processor-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2DynaProcessorMonoWithName(name string) (*LspPlugInPluginsLv2DynaProcessorMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-dyna-processor-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// ae0 (bool)
//
// Attack enable 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAe0() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAe0(value bool) error {
	return e.Element.SetProperty("ae0", value)
}

// ae1 (bool)
//
// Attack enable 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAe1() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAe1(value bool) error {
	return e.Element.SetProperty("ae1", value)
}

// ae2 (bool)
//
// Attack enable 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAe2() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAe2(value bool) error {
	return e.Element.SetProperty("ae2", value)
}

// ae3 (bool)
//
// Attack enable 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAe3() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAe3(value bool) error {
	return e.Element.SetProperty("ae3", value)
}

// al0 (float32)
//
// Attack level 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAl0(value float32) error {
	return e.Element.SetProperty("al0", value)
}

// al1 (float32)
//
// Attack level 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAl1(value float32) error {
	return e.Element.SetProperty("al1", value)
}

// al2 (float32)
//
// Attack level 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAl2(value float32) error {
	return e.Element.SetProperty("al2", value)
}

// al3 (float32)
//
// Attack level 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAl3(value float32) error {
	return e.Element.SetProperty("al3", value)
}

// at0 (float32)
//
// Attack time 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAt0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAt0(value float32) error {
	return e.Element.SetProperty("at0", value)
}

// at1 (float32)
//
// Attack time 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAt1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAt1(value float32) error {
	return e.Element.SetProperty("at1", value)
}

// at2 (float32)
//
// Attack time 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAt2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAt2(value float32) error {
	return e.Element.SetProperty("at2", value)
}

// at3 (float32)
//
// Attack time 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAt3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAt3(value float32) error {
	return e.Element.SetProperty("at3", value)
}

// atd (float32)
//
// Attack time default

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetAtd() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetAtd(value float32) error {
	return e.Element.SetProperty("atd", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetCdr() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetCdr(value float32) error {
	return e.Element.SetProperty("cdr", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm (float32)
//
// Curve level meter

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetClm() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetCmv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetCmv(value bool) error {
	return e.Element.SetProperty("cmv", value)
}

// cwt (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetCwt() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetCwt(value float32) error {
	return e.Element.SetProperty("cwt", value)
}

// elm (float32)
//
// Envelope level meter

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetElm() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetElv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetElv(value bool) error {
	return e.Element.SetProperty("elv", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl0 (float32)
//
// Gain 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGl0(value float32) error {
	return e.Element.SetProperty("gl0", value)
}

// gl1 (float32)
//
// Gain 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGl1(value float32) error {
	return e.Element.SetProperty("gl1", value)
}

// gl2 (float32)
//
// Gain 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGl2(value float32) error {
	return e.Element.SetProperty("gl2", value)
}

// gl3 (float32)
//
// Gain 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGl3(value float32) error {
	return e.Element.SetProperty("gl3", value)
}

// grv (bool)
//
// Gain reduction visibility

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetGrv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetGrv(value bool) error {
	return e.Element.SetProperty("grv", value)
}

// hlr (float32)
//
// High-level ratio

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetHlr() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetHlr(value float32) error {
	return e.Element.SetProperty("hlr", value)
}

// ilm (float32)
//
// Input level meter

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetIlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv (bool)
//
// Input level visibility

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetIlv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetIlv(value bool) error {
	return e.Element.SetProperty("ilv", value)
}

// kn0 (float32)
//
// Knee 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetKn0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetKn0(value float32) error {
	return e.Element.SetProperty("kn0", value)
}

// kn1 (float32)
//
// Knee 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetKn1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetKn1(value float32) error {
	return e.Element.SetProperty("kn1", value)
}

// kn2 (float32)
//
// Knee 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetKn2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetKn2(value float32) error {
	return e.Element.SetProperty("kn2", value)
}

// kn3 (float32)
//
// Knee 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetKn3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetKn3(value float32) error {
	return e.Element.SetProperty("kn3", value)
}

// llr (float32)
//
// Low-level ratio

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetLlr() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetLlr(value float32) error {
	return e.Element.SetProperty("llr", value)
}

// olm (float32)
//
// Output level meter

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetOlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olv (bool)
//
// Output level visibility

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetOlv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetOlv(value bool) error {
	return e.Element.SetProperty("olv", value)
}

// omk (float32)
//
// Overall makeup gain

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetOmk() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetOmk(value float32) error {
	return e.Element.SetProperty("omk", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// pe0 (bool)
//
// Point enable 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetPe0() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetPe0(value bool) error {
	return e.Element.SetProperty("pe0", value)
}

// pe1 (bool)
//
// Point enable 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetPe1() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetPe1(value bool) error {
	return e.Element.SetProperty("pe1", value)
}

// pe2 (bool)
//
// Point enable 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetPe2() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetPe2(value bool) error {
	return e.Element.SetProperty("pe2", value)
}

// pe3 (bool)
//
// Point enable 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetPe3() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetPe3(value bool) error {
	return e.Element.SetProperty("pe3", value)
}

// re0 (bool)
//
// Release enable 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRe0() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRe0(value bool) error {
	return e.Element.SetProperty("re0", value)
}

// re1 (bool)
//
// Release enable 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRe1() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRe1(value bool) error {
	return e.Element.SetProperty("re1", value)
}

// re2 (bool)
//
// Release enable 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRe2() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRe2(value bool) error {
	return e.Element.SetProperty("re2", value)
}

// re3 (bool)
//
// Release enable 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRe3() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRe3(value bool) error {
	return e.Element.SetProperty("re3", value)
}

// rl0 (float32)
//
// Relative level 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRl0(value float32) error {
	return e.Element.SetProperty("rl0", value)
}

// rl1 (float32)
//
// Relative level 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRl1(value float32) error {
	return e.Element.SetProperty("rl1", value)
}

// rl2 (float32)
//
// Relative level 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRl2(value float32) error {
	return e.Element.SetProperty("rl2", value)
}

// rl3 (float32)
//
// Relative level 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRl3(value float32) error {
	return e.Element.SetProperty("rl3", value)
}

// rlm (float32)
//
// Reduction level meter

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRt0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRt0(value float32) error {
	return e.Element.SetProperty("rt0", value)
}

// rt1 (float32)
//
// Release time 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRt1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRt1(value float32) error {
	return e.Element.SetProperty("rt1", value)
}

// rt2 (float32)
//
// Release time 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRt2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRt2(value float32) error {
	return e.Element.SetProperty("rt2", value)
}

// rt3 (float32)
//
// Release time 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRt3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRt3(value float32) error {
	return e.Element.SetProperty("rt3", value)
}

// rtd (float32)
//
// Release time default

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetRtd() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetRtd(value float32) error {
	return e.Element.SetProperty("rtd", value)
}

// scl (bool)
//
// Sidechain listen

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetScl() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetScl(value bool) error {
	return e.Element.SetProperty("scl", value)
}

// scm (GstLspPlugInPluginsLv2DynaProcessorMonoscm)
//
// Sidechain mode

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// scr (float32)
//
// Sidechain reactivity

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetScr() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetScr(value float32) error {
	return e.Element.SetProperty("scr", value)
}

// sct (GstLspPlugInPluginsLv2DynaProcessorMonosct)
//
// Sidechain type

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetSct() (interface{}, error) {
	return e.Element.GetProperty("sct")
}

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetSct(value interface{}) error {
	return e.Element.SetProperty("sct", value)
}

// sla (float32)
//
// Sidechain lookahead

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetSla() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetSla(value float32) error {
	return e.Element.SetProperty("sla", value)
}

// slm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetSlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetSlv() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetSlv(value bool) error {
	return e.Element.SetProperty("slv", value)
}

// tl0 (float32)
//
// Threshold 0

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetTl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetTl0(value float32) error {
	return e.Element.SetProperty("tl0", value)
}

// tl1 (float32)
//
// Threshold 1

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetTl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetTl1(value float32) error {
	return e.Element.SetProperty("tl1", value)
}

// tl2 (float32)
//
// Threshold 2

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetTl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetTl2(value float32) error {
	return e.Element.SetProperty("tl2", value)
}

// tl3 (float32)
//
// Threshold 3

func (e *LspPlugInPluginsLv2DynaProcessorMono) GetTl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMono) SetTl3(value float32) error {
	return e.Element.SetProperty("tl3", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2DynaProcessorMonoscm string

const (
	LspPlugInPluginsLv2DynaProcessorMonoscmPeak LspPlugInPluginsLv2DynaProcessorMonoscm = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2DynaProcessorMonoscmRms LspPlugInPluginsLv2DynaProcessorMonoscm = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2DynaProcessorMonoscmLowPass LspPlugInPluginsLv2DynaProcessorMonoscm = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2DynaProcessorMonoscmUniform LspPlugInPluginsLv2DynaProcessorMonoscm = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2DynaProcessorMonosct string

const (
	LspPlugInPluginsLv2DynaProcessorMonosctFeedForward LspPlugInPluginsLv2DynaProcessorMonosct = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2DynaProcessorMonosctFeedBack LspPlugInPluginsLv2DynaProcessorMonosct = "Feed-back" // Feed-back (1) – Feed-back
)

