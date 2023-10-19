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

type LspPlugInPluginsLv2DynaProcessorMs struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2DynaProcessorMs() (*LspPlugInPluginsLv2DynaProcessorMs, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-dyna-processor-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2DynaProcessorMsWithName(name string) (*LspPlugInPluginsLv2DynaProcessorMs, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-dyna-processor-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// ae0-m (bool)
//
// Attack enable 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe0M(value bool) error {
	return e.Element.SetProperty("ae0-m", value)
}

// ae0-s (bool)
//
// Attack enable 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe0S(value bool) error {
	return e.Element.SetProperty("ae0-s", value)
}

// ae1-m (bool)
//
// Attack enable 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe1M(value bool) error {
	return e.Element.SetProperty("ae1-m", value)
}

// ae1-s (bool)
//
// Attack enable 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe1S(value bool) error {
	return e.Element.SetProperty("ae1-s", value)
}

// ae2-m (bool)
//
// Attack enable 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe2M(value bool) error {
	return e.Element.SetProperty("ae2-m", value)
}

// ae2-s (bool)
//
// Attack enable 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe2S(value bool) error {
	return e.Element.SetProperty("ae2-s", value)
}

// ae3-m (bool)
//
// Attack enable 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe3M(value bool) error {
	return e.Element.SetProperty("ae3-m", value)
}

// ae3-s (bool)
//
// Attack enable 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAe3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAe3S(value bool) error {
	return e.Element.SetProperty("ae3-s", value)
}

// al0-m (float32)
//
// Attack level 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl0M(value float32) error {
	return e.Element.SetProperty("al0-m", value)
}

// al0-s (float32)
//
// Attack level 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl0S(value float32) error {
	return e.Element.SetProperty("al0-s", value)
}

// al1-m (float32)
//
// Attack level 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl1M(value float32) error {
	return e.Element.SetProperty("al1-m", value)
}

// al1-s (float32)
//
// Attack level 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl1S(value float32) error {
	return e.Element.SetProperty("al1-s", value)
}

// al2-m (float32)
//
// Attack level 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl2M(value float32) error {
	return e.Element.SetProperty("al2-m", value)
}

// al2-s (float32)
//
// Attack level 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl2S(value float32) error {
	return e.Element.SetProperty("al2-s", value)
}

// al3-m (float32)
//
// Attack level 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl3M(value float32) error {
	return e.Element.SetProperty("al3-m", value)
}

// al3-s (float32)
//
// Attack level 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAl3S(value float32) error {
	return e.Element.SetProperty("al3-s", value)
}

// at0-m (float32)
//
// Attack time 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt0M(value float32) error {
	return e.Element.SetProperty("at0-m", value)
}

// at0-s (float32)
//
// Attack time 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt0S(value float32) error {
	return e.Element.SetProperty("at0-s", value)
}

// at1-m (float32)
//
// Attack time 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt1M(value float32) error {
	return e.Element.SetProperty("at1-m", value)
}

// at1-s (float32)
//
// Attack time 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt1S(value float32) error {
	return e.Element.SetProperty("at1-s", value)
}

// at2-m (float32)
//
// Attack time 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt2M(value float32) error {
	return e.Element.SetProperty("at2-m", value)
}

// at2-s (float32)
//
// Attack time 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt2S(value float32) error {
	return e.Element.SetProperty("at2-s", value)
}

// at3-m (float32)
//
// Attack time 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt3M(value float32) error {
	return e.Element.SetProperty("at3-m", value)
}

// at3-s (float32)
//
// Attack time 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAt3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAt3S(value float32) error {
	return e.Element.SetProperty("at3-s", value)
}

// atd-m (float32)
//
// Attack time default Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAtdM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("atd-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAtdM(value float32) error {
	return e.Element.SetProperty("atd-m", value)
}

// atd-s (float32)
//
// Attack time default Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetAtdS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("atd-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetAtdS(value float32) error {
	return e.Element.SetProperty("atd-s", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-m (float32)
//
// Dry gain Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetCdrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetCdrM(value float32) error {
	return e.Element.SetProperty("cdr-m", value)
}

// cdr-s (float32)
//
// Dry gain Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetCdrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetCdrS(value float32) error {
	return e.Element.SetProperty("cdr-s", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-m (float32)
//
// Curve level meter Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetClmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetClmS() (float32, error) {
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

// cmv-m (bool)
//
// Curve modelling visibility Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetCmvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cmv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetCmvM(value bool) error {
	return e.Element.SetProperty("cmv-m", value)
}

// cmv-s (bool)
//
// Curve modelling visibility Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetCmvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cmv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetCmvS(value bool) error {
	return e.Element.SetProperty("cmv-s", value)
}

// cwt-m (float32)
//
// Wet gain Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetCwtM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetCwtM(value float32) error {
	return e.Element.SetProperty("cwt-m", value)
}

// cwt-s (float32)
//
// Wet gain Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetCwtS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetCwtS(value float32) error {
	return e.Element.SetProperty("cwt-s", value)
}

// elm-m (float32)
//
// Envelope level meter Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetElmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetElmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetElvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetElvM(value bool) error {
	return e.Element.SetProperty("elv-m", value)
}

// elv-s (bool)
//
// Envelope level visibility Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetElvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetElvS(value bool) error {
	return e.Element.SetProperty("elv-s", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl0-m (float32)
//
// Gain 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl0M(value float32) error {
	return e.Element.SetProperty("gl0-m", value)
}

// gl0-s (float32)
//
// Gain 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl0S(value float32) error {
	return e.Element.SetProperty("gl0-s", value)
}

// gl1-m (float32)
//
// Gain 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl1M(value float32) error {
	return e.Element.SetProperty("gl1-m", value)
}

// gl1-s (float32)
//
// Gain 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl1S(value float32) error {
	return e.Element.SetProperty("gl1-s", value)
}

// gl2-m (float32)
//
// Gain 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl2M(value float32) error {
	return e.Element.SetProperty("gl2-m", value)
}

// gl2-s (float32)
//
// Gain 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl2S(value float32) error {
	return e.Element.SetProperty("gl2-s", value)
}

// gl3-m (float32)
//
// Gain 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl3M(value float32) error {
	return e.Element.SetProperty("gl3-m", value)
}

// gl3-s (float32)
//
// Gain 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGl3S(value float32) error {
	return e.Element.SetProperty("gl3-s", value)
}

// grv-m (bool)
//
// Gain reduction visibility Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGrvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGrvM(value bool) error {
	return e.Element.SetProperty("grv-m", value)
}

// grv-s (bool)
//
// Gain reduction visibility Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetGrvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetGrvS(value bool) error {
	return e.Element.SetProperty("grv-s", value)
}

// hlr-m (float32)
//
// High-level ratio Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetHlrM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hlr-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetHlrM(value float32) error {
	return e.Element.SetProperty("hlr-m", value)
}

// hlr-s (float32)
//
// High-level ratio Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetHlrS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hlr-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetHlrS(value float32) error {
	return e.Element.SetProperty("hlr-s", value)
}

// ilm-m (float32)
//
// Input level meter Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetIlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetIlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetIlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetIlvM(value bool) error {
	return e.Element.SetProperty("ilv-m", value)
}

// ilv-s (bool)
//
// Input level visibility Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetIlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetIlvS(value bool) error {
	return e.Element.SetProperty("ilv-s", value)
}

// kn0-m (float32)
//
// Knee 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn0M(value float32) error {
	return e.Element.SetProperty("kn0-m", value)
}

// kn0-s (float32)
//
// Knee 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn0S(value float32) error {
	return e.Element.SetProperty("kn0-s", value)
}

// kn1-m (float32)
//
// Knee 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn1M(value float32) error {
	return e.Element.SetProperty("kn1-m", value)
}

// kn1-s (float32)
//
// Knee 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn1S(value float32) error {
	return e.Element.SetProperty("kn1-s", value)
}

// kn2-m (float32)
//
// Knee 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn2M(value float32) error {
	return e.Element.SetProperty("kn2-m", value)
}

// kn2-s (float32)
//
// Knee 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn2S(value float32) error {
	return e.Element.SetProperty("kn2-s", value)
}

// kn3-m (float32)
//
// Knee 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn3M(value float32) error {
	return e.Element.SetProperty("kn3-m", value)
}

// kn3-s (float32)
//
// Knee 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetKn3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetKn3S(value float32) error {
	return e.Element.SetProperty("kn3-s", value)
}

// llr-m (float32)
//
// Low-level ratio Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetLlrM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("llr-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetLlrM(value float32) error {
	return e.Element.SetProperty("llr-m", value)
}

// llr-s (float32)
//
// Low-level ratio Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetLlrS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("llr-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetLlrS(value float32) error {
	return e.Element.SetProperty("llr-s", value)
}

// msl (bool)
//
// Mid/Side listen

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetMsl() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetMsl(value bool) error {
	return e.Element.SetProperty("msl", value)
}

// olm-m (float32)
//
// Output level meter Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetOlvM(value bool) error {
	return e.Element.SetProperty("olv-m", value)
}

// olv-s (bool)
//
// Output level visibility Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetOlvS(value bool) error {
	return e.Element.SetProperty("olv-s", value)
}

// omk-m (float32)
//
// Overall makeup gain Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOmkM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("omk-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetOmkM(value float32) error {
	return e.Element.SetProperty("omk-m", value)
}

// omk-s (float32)
//
// Overall makeup gain Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOmkS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("omk-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetOmkS(value float32) error {
	return e.Element.SetProperty("omk-s", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// pe0-m (bool)
//
// Point enable 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe0M(value bool) error {
	return e.Element.SetProperty("pe0-m", value)
}

// pe0-s (bool)
//
// Point enable 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe0S(value bool) error {
	return e.Element.SetProperty("pe0-s", value)
}

// pe1-m (bool)
//
// Point enable 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe1M(value bool) error {
	return e.Element.SetProperty("pe1-m", value)
}

// pe1-s (bool)
//
// Point enable 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe1S(value bool) error {
	return e.Element.SetProperty("pe1-s", value)
}

// pe2-m (bool)
//
// Point enable 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe2M(value bool) error {
	return e.Element.SetProperty("pe2-m", value)
}

// pe2-s (bool)
//
// Point enable 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe2S(value bool) error {
	return e.Element.SetProperty("pe2-s", value)
}

// pe3-m (bool)
//
// Point enable 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe3M(value bool) error {
	return e.Element.SetProperty("pe3-m", value)
}

// pe3-s (bool)
//
// Point enable 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPe3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPe3S(value bool) error {
	return e.Element.SetProperty("pe3-s", value)
}

// psel (GstLspPlugInPluginsLv2DynaProcessorMspsel)
//
// Processor selector

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetPsel() (interface{}, error) {
	return e.Element.GetProperty("psel")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetPsel(value interface{}) error {
	return e.Element.SetProperty("psel", value)
}

// re0-m (bool)
//
// Release enable 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe0M(value bool) error {
	return e.Element.SetProperty("re0-m", value)
}

// re0-s (bool)
//
// Release enable 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe0S(value bool) error {
	return e.Element.SetProperty("re0-s", value)
}

// re1-m (bool)
//
// Release enable 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe1M(value bool) error {
	return e.Element.SetProperty("re1-m", value)
}

// re1-s (bool)
//
// Release enable 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe1S(value bool) error {
	return e.Element.SetProperty("re1-s", value)
}

// re2-m (bool)
//
// Release enable 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe2M(value bool) error {
	return e.Element.SetProperty("re2-m", value)
}

// re2-s (bool)
//
// Release enable 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe2S(value bool) error {
	return e.Element.SetProperty("re2-s", value)
}

// re3-m (bool)
//
// Release enable 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe3M(value bool) error {
	return e.Element.SetProperty("re3-m", value)
}

// re3-s (bool)
//
// Release enable 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRe3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRe3S(value bool) error {
	return e.Element.SetProperty("re3-s", value)
}

// rl0-m (float32)
//
// Relative level 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl0M(value float32) error {
	return e.Element.SetProperty("rl0-m", value)
}

// rl0-s (float32)
//
// Relative level 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl0S(value float32) error {
	return e.Element.SetProperty("rl0-s", value)
}

// rl1-m (float32)
//
// Relative level 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl1M(value float32) error {
	return e.Element.SetProperty("rl1-m", value)
}

// rl1-s (float32)
//
// Relative level 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl1S(value float32) error {
	return e.Element.SetProperty("rl1-s", value)
}

// rl2-m (float32)
//
// Relative level 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl2M(value float32) error {
	return e.Element.SetProperty("rl2-m", value)
}

// rl2-s (float32)
//
// Relative level 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl2S(value float32) error {
	return e.Element.SetProperty("rl2-s", value)
}

// rl3-m (float32)
//
// Relative level 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl3M(value float32) error {
	return e.Element.SetProperty("rl3-m", value)
}

// rl3-s (float32)
//
// Relative level 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRl3S(value float32) error {
	return e.Element.SetProperty("rl3-s", value)
}

// rlm-m (float32)
//
// Reduction level meter Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRlmS() (float32, error) {
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

// rt0-m (float32)
//
// Release time 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt0M(value float32) error {
	return e.Element.SetProperty("rt0-m", value)
}

// rt0-s (float32)
//
// Release time 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt0S(value float32) error {
	return e.Element.SetProperty("rt0-s", value)
}

// rt1-m (float32)
//
// Release time 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt1M(value float32) error {
	return e.Element.SetProperty("rt1-m", value)
}

// rt1-s (float32)
//
// Release time 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt1S(value float32) error {
	return e.Element.SetProperty("rt1-s", value)
}

// rt2-m (float32)
//
// Release time 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt2M(value float32) error {
	return e.Element.SetProperty("rt2-m", value)
}

// rt2-s (float32)
//
// Release time 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt2S(value float32) error {
	return e.Element.SetProperty("rt2-s", value)
}

// rt3-m (float32)
//
// Release time 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt3M(value float32) error {
	return e.Element.SetProperty("rt3-m", value)
}

// rt3-s (float32)
//
// Release time 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRt3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRt3S(value float32) error {
	return e.Element.SetProperty("rt3-s", value)
}

// rtd-m (float32)
//
// Release time default Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRtdM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rtd-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRtdM(value float32) error {
	return e.Element.SetProperty("rtd-m", value)
}

// rtd-s (float32)
//
// Release time default Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetRtdS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rtd-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetRtdS(value float32) error {
	return e.Element.SetProperty("rtd-s", value)
}

// scl-m (bool)
//
// Sidechain listen Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSclM() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSclM(value bool) error {
	return e.Element.SetProperty("scl-m", value)
}

// scl-s (bool)
//
// Sidechain listen Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSclS() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSclS(value bool) error {
	return e.Element.SetProperty("scl-s", value)
}

// scm-m (GstLspPlugInPluginsLv2DynaProcessorMsscmM)
//
// Sidechain mode Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScmM() (interface{}, error) {
	return e.Element.GetProperty("scm-m")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScmM(value interface{}) error {
	return e.Element.SetProperty("scm-m", value)
}

// scm-s (GstLspPlugInPluginsLv2DynaProcessorMsscmS)
//
// Sidechain mode Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScmS() (interface{}, error) {
	return e.Element.GetProperty("scm-s")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScmS(value interface{}) error {
	return e.Element.SetProperty("scm-s", value)
}

// scp-m (float32)
//
// Sidechain preamp Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScpM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScpM(value float32) error {
	return e.Element.SetProperty("scp-m", value)
}

// scp-s (float32)
//
// Sidechain preamp Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScpS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScpS(value float32) error {
	return e.Element.SetProperty("scp-s", value)
}

// scr-m (float32)
//
// Sidechain reactivity Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScrM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScrM(value float32) error {
	return e.Element.SetProperty("scr-m", value)
}

// scr-s (float32)
//
// Sidechain reactivity Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScrS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScrS(value float32) error {
	return e.Element.SetProperty("scr-s", value)
}

// scs-m (GstLspPlugInPluginsLv2DynaProcessorMsscsM)
//
// Sidechain source Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScsM() (interface{}, error) {
	return e.Element.GetProperty("scs-m")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScsM(value interface{}) error {
	return e.Element.SetProperty("scs-m", value)
}

// scs-s (GstLspPlugInPluginsLv2DynaProcessorMsscsS)
//
// Sidechain source Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetScsS() (interface{}, error) {
	return e.Element.GetProperty("scs-s")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetScsS(value interface{}) error {
	return e.Element.SetProperty("scs-s", value)
}

// sct-m (GstLspPlugInPluginsLv2DynaProcessorMssctM)
//
// Sidechain type Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSctM() (interface{}, error) {
	return e.Element.GetProperty("sct-m")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSctM(value interface{}) error {
	return e.Element.SetProperty("sct-m", value)
}

// sct-s (GstLspPlugInPluginsLv2DynaProcessorMssctS)
//
// Sidechain type Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSctS() (interface{}, error) {
	return e.Element.GetProperty("sct-s")
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSctS(value interface{}) error {
	return e.Element.SetProperty("sct-s", value)
}

// sla-m (float32)
//
// Sidechain lookahead Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSlaM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSlaM(value float32) error {
	return e.Element.SetProperty("sla-m", value)
}

// sla-s (float32)
//
// Sidechain lookahead Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSlaS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSlaS(value float32) error {
	return e.Element.SetProperty("sla-s", value)
}

// slm-m (float32)
//
// Sidechain level meter Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSlmM() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSlmS() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSlvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSlvM(value bool) error {
	return e.Element.SetProperty("slv-m", value)
}

// slv-s (bool)
//
// Sidechain level visibility Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetSlvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetSlvS(value bool) error {
	return e.Element.SetProperty("slv-s", value)
}

// tl0-m (float32)
//
// Threshold 0 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl0-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl0M(value float32) error {
	return e.Element.SetProperty("tl0-m", value)
}

// tl0-s (float32)
//
// Threshold 0 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl0-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl0S(value float32) error {
	return e.Element.SetProperty("tl0-s", value)
}

// tl1-m (float32)
//
// Threshold 1 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl1-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl1M(value float32) error {
	return e.Element.SetProperty("tl1-m", value)
}

// tl1-s (float32)
//
// Threshold 1 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl1-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl1S(value float32) error {
	return e.Element.SetProperty("tl1-s", value)
}

// tl2-m (float32)
//
// Threshold 2 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl2-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl2M(value float32) error {
	return e.Element.SetProperty("tl2-m", value)
}

// tl2-s (float32)
//
// Threshold 2 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl2-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl2S(value float32) error {
	return e.Element.SetProperty("tl2-s", value)
}

// tl3-m (float32)
//
// Threshold 3 Mid

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl3-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl3M(value float32) error {
	return e.Element.SetProperty("tl3-m", value)
}

// tl3-s (float32)
//
// Threshold 3 Side

func (e *LspPlugInPluginsLv2DynaProcessorMs) GetTl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl3-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorMs) SetTl3S(value float32) error {
	return e.Element.SetProperty("tl3-s", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2DynaProcessorMsscmM string

const (
	LspPlugInPluginsLv2DynaProcessorMsscmMPeak LspPlugInPluginsLv2DynaProcessorMsscmM = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2DynaProcessorMsscmMRms LspPlugInPluginsLv2DynaProcessorMsscmM = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2DynaProcessorMsscmMLowPass LspPlugInPluginsLv2DynaProcessorMsscmM = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2DynaProcessorMsscmMUniform LspPlugInPluginsLv2DynaProcessorMsscmM = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2DynaProcessorMsscmS string

const (
	LspPlugInPluginsLv2DynaProcessorMsscmSPeak LspPlugInPluginsLv2DynaProcessorMsscmS = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2DynaProcessorMsscmSRms LspPlugInPluginsLv2DynaProcessorMsscmS = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2DynaProcessorMsscmSLowPass LspPlugInPluginsLv2DynaProcessorMsscmS = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2DynaProcessorMsscmSUniform LspPlugInPluginsLv2DynaProcessorMsscmS = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2DynaProcessorMsscsM string

const (
	LspPlugInPluginsLv2DynaProcessorMsscsMMiddle LspPlugInPluginsLv2DynaProcessorMsscsM = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2DynaProcessorMsscsMSide LspPlugInPluginsLv2DynaProcessorMsscsM = "Side" // Side (1) – Side
	LspPlugInPluginsLv2DynaProcessorMsscsMLeft LspPlugInPluginsLv2DynaProcessorMsscsM = "Left" // Left (2) – Left
	LspPlugInPluginsLv2DynaProcessorMsscsMRight LspPlugInPluginsLv2DynaProcessorMsscsM = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2DynaProcessorMsscsS string

const (
	LspPlugInPluginsLv2DynaProcessorMsscsSMiddle LspPlugInPluginsLv2DynaProcessorMsscsS = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2DynaProcessorMsscsSSide LspPlugInPluginsLv2DynaProcessorMsscsS = "Side" // Side (1) – Side
	LspPlugInPluginsLv2DynaProcessorMsscsSLeft LspPlugInPluginsLv2DynaProcessorMsscsS = "Left" // Left (2) – Left
	LspPlugInPluginsLv2DynaProcessorMsscsSRight LspPlugInPluginsLv2DynaProcessorMsscsS = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2DynaProcessorMssctM string

const (
	LspPlugInPluginsLv2DynaProcessorMssctMFeedForward LspPlugInPluginsLv2DynaProcessorMssctM = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2DynaProcessorMssctMFeedBack LspPlugInPluginsLv2DynaProcessorMssctM = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2DynaProcessorMssctS string

const (
	LspPlugInPluginsLv2DynaProcessorMssctSFeedForward LspPlugInPluginsLv2DynaProcessorMssctS = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2DynaProcessorMssctSFeedBack LspPlugInPluginsLv2DynaProcessorMssctS = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2DynaProcessorMspsel string

const (
	LspPlugInPluginsLv2DynaProcessorMspselMiddle LspPlugInPluginsLv2DynaProcessorMspsel = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2DynaProcessorMspselSide LspPlugInPluginsLv2DynaProcessorMspsel = "Side" // Side (1) – Side
)

