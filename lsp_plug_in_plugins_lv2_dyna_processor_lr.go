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

type LspPlugInPluginsLv2DynaProcessorLr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2DynaProcessorLr() (*LspPlugInPluginsLv2DynaProcessorLr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-dyna-processor-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2DynaProcessorLrWithName(name string) (*LspPlugInPluginsLv2DynaProcessorLr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-dyna-processor-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2DynaProcessorLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// ae0-l (bool)
//
// Attack enable 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe0L(value bool) error {
	return e.Element.SetProperty("ae0-l", value)
}

// ae0-r (bool)
//
// Attack enable 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe0R(value bool) error {
	return e.Element.SetProperty("ae0-r", value)
}

// ae1-l (bool)
//
// Attack enable 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe1L(value bool) error {
	return e.Element.SetProperty("ae1-l", value)
}

// ae1-r (bool)
//
// Attack enable 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe1R(value bool) error {
	return e.Element.SetProperty("ae1-r", value)
}

// ae2-l (bool)
//
// Attack enable 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe2L(value bool) error {
	return e.Element.SetProperty("ae2-l", value)
}

// ae2-r (bool)
//
// Attack enable 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe2R(value bool) error {
	return e.Element.SetProperty("ae2-r", value)
}

// ae3-l (bool)
//
// Attack enable 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe3L(value bool) error {
	return e.Element.SetProperty("ae3-l", value)
}

// ae3-r (bool)
//
// Attack enable 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAe3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ae3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAe3R(value bool) error {
	return e.Element.SetProperty("ae3-r", value)
}

// al0-l (float32)
//
// Attack level 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl0L(value float32) error {
	return e.Element.SetProperty("al0-l", value)
}

// al0-r (float32)
//
// Attack level 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl0R(value float32) error {
	return e.Element.SetProperty("al0-r", value)
}

// al1-l (float32)
//
// Attack level 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl1L(value float32) error {
	return e.Element.SetProperty("al1-l", value)
}

// al1-r (float32)
//
// Attack level 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl1R(value float32) error {
	return e.Element.SetProperty("al1-r", value)
}

// al2-l (float32)
//
// Attack level 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl2L(value float32) error {
	return e.Element.SetProperty("al2-l", value)
}

// al2-r (float32)
//
// Attack level 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl2R(value float32) error {
	return e.Element.SetProperty("al2-r", value)
}

// al3-l (float32)
//
// Attack level 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl3L(value float32) error {
	return e.Element.SetProperty("al3-l", value)
}

// al3-r (float32)
//
// Attack level 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAl3R(value float32) error {
	return e.Element.SetProperty("al3-r", value)
}

// at0-l (float32)
//
// Attack time 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt0L(value float32) error {
	return e.Element.SetProperty("at0-l", value)
}

// at0-r (float32)
//
// Attack time 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt0R(value float32) error {
	return e.Element.SetProperty("at0-r", value)
}

// at1-l (float32)
//
// Attack time 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt1L(value float32) error {
	return e.Element.SetProperty("at1-l", value)
}

// at1-r (float32)
//
// Attack time 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt1R(value float32) error {
	return e.Element.SetProperty("at1-r", value)
}

// at2-l (float32)
//
// Attack time 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt2L(value float32) error {
	return e.Element.SetProperty("at2-l", value)
}

// at2-r (float32)
//
// Attack time 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt2R(value float32) error {
	return e.Element.SetProperty("at2-r", value)
}

// at3-l (float32)
//
// Attack time 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt3L(value float32) error {
	return e.Element.SetProperty("at3-l", value)
}

// at3-r (float32)
//
// Attack time 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAt3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAt3R(value float32) error {
	return e.Element.SetProperty("at3-r", value)
}

// atd-l (float32)
//
// Attack time default Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAtdL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("atd-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAtdL(value float32) error {
	return e.Element.SetProperty("atd-l", value)
}

// atd-r (float32)
//
// Attack time default Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetAtdR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("atd-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetAtdR(value float32) error {
	return e.Element.SetProperty("atd-r", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-l (float32)
//
// Dry gain Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetCdrL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetCdrL(value float32) error {
	return e.Element.SetProperty("cdr-l", value)
}

// cdr-r (float32)
//
// Dry gain Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetCdrR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetCdrR(value float32) error {
	return e.Element.SetProperty("cdr-r", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-l (float32)
//
// Curve level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetClmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetClmR() (float32, error) {
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

// cmv-l (bool)
//
// Curve modelling visibility Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetCmvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cmv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetCmvL(value bool) error {
	return e.Element.SetProperty("cmv-l", value)
}

// cmv-r (bool)
//
// Curve modelling visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetCmvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cmv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetCmvR(value bool) error {
	return e.Element.SetProperty("cmv-r", value)
}

// cwt-l (float32)
//
// Wet gain Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetCwtL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetCwtL(value float32) error {
	return e.Element.SetProperty("cwt-l", value)
}

// cwt-r (float32)
//
// Wet gain Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetCwtR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetCwtR(value float32) error {
	return e.Element.SetProperty("cwt-r", value)
}

// elm-l (float32)
//
// Envelope level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetElmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetElmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetElvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetElvL(value bool) error {
	return e.Element.SetProperty("elv-l", value)
}

// elv-r (bool)
//
// Envelope level visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetElvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetElvR(value bool) error {
	return e.Element.SetProperty("elv-r", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl0-l (float32)
//
// Gain 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl0L(value float32) error {
	return e.Element.SetProperty("gl0-l", value)
}

// gl0-r (float32)
//
// Gain 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl0R(value float32) error {
	return e.Element.SetProperty("gl0-r", value)
}

// gl1-l (float32)
//
// Gain 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl1L(value float32) error {
	return e.Element.SetProperty("gl1-l", value)
}

// gl1-r (float32)
//
// Gain 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl1R(value float32) error {
	return e.Element.SetProperty("gl1-r", value)
}

// gl2-l (float32)
//
// Gain 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl2L(value float32) error {
	return e.Element.SetProperty("gl2-l", value)
}

// gl2-r (float32)
//
// Gain 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl2R(value float32) error {
	return e.Element.SetProperty("gl2-r", value)
}

// gl3-l (float32)
//
// Gain 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl3L(value float32) error {
	return e.Element.SetProperty("gl3-l", value)
}

// gl3-r (float32)
//
// Gain 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGl3R(value float32) error {
	return e.Element.SetProperty("gl3-r", value)
}

// grv-l (bool)
//
// Gain reduction visibility Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGrvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGrvL(value bool) error {
	return e.Element.SetProperty("grv-l", value)
}

// grv-r (bool)
//
// Gain reduction visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetGrvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetGrvR(value bool) error {
	return e.Element.SetProperty("grv-r", value)
}

// hlr-l (float32)
//
// High-level ratio Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetHlrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hlr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetHlrL(value float32) error {
	return e.Element.SetProperty("hlr-l", value)
}

// hlr-r (float32)
//
// High-level ratio Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetHlrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hlr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetHlrR(value float32) error {
	return e.Element.SetProperty("hlr-r", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetIlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetIlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetIlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// kn0-l (float32)
//
// Knee 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn0L(value float32) error {
	return e.Element.SetProperty("kn0-l", value)
}

// kn0-r (float32)
//
// Knee 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn0R(value float32) error {
	return e.Element.SetProperty("kn0-r", value)
}

// kn1-l (float32)
//
// Knee 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn1L(value float32) error {
	return e.Element.SetProperty("kn1-l", value)
}

// kn1-r (float32)
//
// Knee 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn1R(value float32) error {
	return e.Element.SetProperty("kn1-r", value)
}

// kn2-l (float32)
//
// Knee 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn2L(value float32) error {
	return e.Element.SetProperty("kn2-l", value)
}

// kn2-r (float32)
//
// Knee 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn2R(value float32) error {
	return e.Element.SetProperty("kn2-r", value)
}

// kn3-l (float32)
//
// Knee 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn3L(value float32) error {
	return e.Element.SetProperty("kn3-l", value)
}

// kn3-r (float32)
//
// Knee 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetKn3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetKn3R(value float32) error {
	return e.Element.SetProperty("kn3-r", value)
}

// llr-l (float32)
//
// Low-level ratio Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetLlrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("llr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetLlrL(value float32) error {
	return e.Element.SetProperty("llr-l", value)
}

// llr-r (float32)
//
// Low-level ratio Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetLlrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("llr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetLlrR(value float32) error {
	return e.Element.SetProperty("llr-r", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// omk-l (float32)
//
// Overall makeup gain Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOmkL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("omk-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetOmkL(value float32) error {
	return e.Element.SetProperty("omk-l", value)
}

// omk-r (float32)
//
// Overall makeup gain Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOmkR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("omk-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetOmkR(value float32) error {
	return e.Element.SetProperty("omk-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// pe0-l (bool)
//
// Point enable 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe0L(value bool) error {
	return e.Element.SetProperty("pe0-l", value)
}

// pe0-r (bool)
//
// Point enable 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe0R(value bool) error {
	return e.Element.SetProperty("pe0-r", value)
}

// pe1-l (bool)
//
// Point enable 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe1L(value bool) error {
	return e.Element.SetProperty("pe1-l", value)
}

// pe1-r (bool)
//
// Point enable 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe1R(value bool) error {
	return e.Element.SetProperty("pe1-r", value)
}

// pe2-l (bool)
//
// Point enable 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe2L(value bool) error {
	return e.Element.SetProperty("pe2-l", value)
}

// pe2-r (bool)
//
// Point enable 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe2R(value bool) error {
	return e.Element.SetProperty("pe2-r", value)
}

// pe3-l (bool)
//
// Point enable 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe3L(value bool) error {
	return e.Element.SetProperty("pe3-l", value)
}

// pe3-r (bool)
//
// Point enable 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPe3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pe3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPe3R(value bool) error {
	return e.Element.SetProperty("pe3-r", value)
}

// psel (GstLspPlugInPluginsLv2DynaProcessorLrpsel)
//
// Processor selector

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetPsel() (interface{}, error) {
	return e.Element.GetProperty("psel")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetPsel(value interface{}) error {
	return e.Element.SetProperty("psel", value)
}

// re0-l (bool)
//
// Release enable 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe0L(value bool) error {
	return e.Element.SetProperty("re0-l", value)
}

// re0-r (bool)
//
// Release enable 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe0R(value bool) error {
	return e.Element.SetProperty("re0-r", value)
}

// re1-l (bool)
//
// Release enable 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe1L(value bool) error {
	return e.Element.SetProperty("re1-l", value)
}

// re1-r (bool)
//
// Release enable 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe1R(value bool) error {
	return e.Element.SetProperty("re1-r", value)
}

// re2-l (bool)
//
// Release enable 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe2L(value bool) error {
	return e.Element.SetProperty("re2-l", value)
}

// re2-r (bool)
//
// Release enable 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe2R(value bool) error {
	return e.Element.SetProperty("re2-r", value)
}

// re3-l (bool)
//
// Release enable 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe3L(value bool) error {
	return e.Element.SetProperty("re3-l", value)
}

// re3-r (bool)
//
// Release enable 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRe3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("re3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRe3R(value bool) error {
	return e.Element.SetProperty("re3-r", value)
}

// rl0-l (float32)
//
// Relative level 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl0L(value float32) error {
	return e.Element.SetProperty("rl0-l", value)
}

// rl0-r (float32)
//
// Relative level 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl0R(value float32) error {
	return e.Element.SetProperty("rl0-r", value)
}

// rl1-l (float32)
//
// Relative level 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl1L(value float32) error {
	return e.Element.SetProperty("rl1-l", value)
}

// rl1-r (float32)
//
// Relative level 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl1R(value float32) error {
	return e.Element.SetProperty("rl1-r", value)
}

// rl2-l (float32)
//
// Relative level 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl2L(value float32) error {
	return e.Element.SetProperty("rl2-l", value)
}

// rl2-r (float32)
//
// Relative level 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl2R(value float32) error {
	return e.Element.SetProperty("rl2-r", value)
}

// rl3-l (float32)
//
// Relative level 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl3L(value float32) error {
	return e.Element.SetProperty("rl3-l", value)
}

// rl3-r (float32)
//
// Relative level 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRl3R(value float32) error {
	return e.Element.SetProperty("rl3-r", value)
}

// rlm-l (float32)
//
// Reduction level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRlmR() (float32, error) {
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

// rt0-l (float32)
//
// Release time 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt0L(value float32) error {
	return e.Element.SetProperty("rt0-l", value)
}

// rt0-r (float32)
//
// Release time 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt0R(value float32) error {
	return e.Element.SetProperty("rt0-r", value)
}

// rt1-l (float32)
//
// Release time 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt1L(value float32) error {
	return e.Element.SetProperty("rt1-l", value)
}

// rt1-r (float32)
//
// Release time 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt1R(value float32) error {
	return e.Element.SetProperty("rt1-r", value)
}

// rt2-l (float32)
//
// Release time 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt2L(value float32) error {
	return e.Element.SetProperty("rt2-l", value)
}

// rt2-r (float32)
//
// Release time 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt2R(value float32) error {
	return e.Element.SetProperty("rt2-r", value)
}

// rt3-l (float32)
//
// Release time 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt3L(value float32) error {
	return e.Element.SetProperty("rt3-l", value)
}

// rt3-r (float32)
//
// Release time 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRt3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRt3R(value float32) error {
	return e.Element.SetProperty("rt3-r", value)
}

// rtd-l (float32)
//
// Release time default Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRtdL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rtd-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRtdL(value float32) error {
	return e.Element.SetProperty("rtd-l", value)
}

// rtd-r (float32)
//
// Release time default Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetRtdR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rtd-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetRtdR(value float32) error {
	return e.Element.SetProperty("rtd-r", value)
}

// scl-l (bool)
//
// Sidechain listen Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSclL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSclL(value bool) error {
	return e.Element.SetProperty("scl-l", value)
}

// scl-r (bool)
//
// Sidechain listen Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSclR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSclR(value bool) error {
	return e.Element.SetProperty("scl-r", value)
}

// scm-l (GstLspPlugInPluginsLv2DynaProcessorLrscmL)
//
// Sidechain mode Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScmL() (interface{}, error) {
	return e.Element.GetProperty("scm-l")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScmL(value interface{}) error {
	return e.Element.SetProperty("scm-l", value)
}

// scm-r (GstLspPlugInPluginsLv2DynaProcessorLrscmR)
//
// Sidechain mode Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScmR() (interface{}, error) {
	return e.Element.GetProperty("scm-r")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScmR(value interface{}) error {
	return e.Element.SetProperty("scm-r", value)
}

// scp-l (float32)
//
// Sidechain preamp Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScpL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScpL(value float32) error {
	return e.Element.SetProperty("scp-l", value)
}

// scp-r (float32)
//
// Sidechain preamp Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScpR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScpR(value float32) error {
	return e.Element.SetProperty("scp-r", value)
}

// scr-l (float32)
//
// Sidechain reactivity Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScrL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScrL(value float32) error {
	return e.Element.SetProperty("scr-l", value)
}

// scr-r (float32)
//
// Sidechain reactivity Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScrR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScrR(value float32) error {
	return e.Element.SetProperty("scr-r", value)
}

// scs-l (GstLspPlugInPluginsLv2DynaProcessorLrscsL)
//
// Sidechain source Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScsL() (interface{}, error) {
	return e.Element.GetProperty("scs-l")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScsL(value interface{}) error {
	return e.Element.SetProperty("scs-l", value)
}

// scs-r (GstLspPlugInPluginsLv2DynaProcessorLrscsR)
//
// Sidechain source Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetScsR() (interface{}, error) {
	return e.Element.GetProperty("scs-r")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetScsR(value interface{}) error {
	return e.Element.SetProperty("scs-r", value)
}

// sct-l (GstLspPlugInPluginsLv2DynaProcessorLrsctL)
//
// Sidechain type Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSctL() (interface{}, error) {
	return e.Element.GetProperty("sct-l")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSctL(value interface{}) error {
	return e.Element.SetProperty("sct-l", value)
}

// sct-r (GstLspPlugInPluginsLv2DynaProcessorLrsctR)
//
// Sidechain type Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSctR() (interface{}, error) {
	return e.Element.GetProperty("sct-r")
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSctR(value interface{}) error {
	return e.Element.SetProperty("sct-r", value)
}

// sla-l (float32)
//
// Sidechain lookahead Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSlaL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSlaL(value float32) error {
	return e.Element.SetProperty("sla-l", value)
}

// sla-r (float32)
//
// Sidechain lookahead Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSlaR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSlaR(value float32) error {
	return e.Element.SetProperty("sla-r", value)
}

// slm-l (float32)
//
// Sidechain level meter Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSlvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSlvL(value bool) error {
	return e.Element.SetProperty("slv-l", value)
}

// slv-r (bool)
//
// Sidechain level visibility Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetSlvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetSlvR(value bool) error {
	return e.Element.SetProperty("slv-r", value)
}

// tl0-l (float32)
//
// Threshold 0 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl0-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl0L(value float32) error {
	return e.Element.SetProperty("tl0-l", value)
}

// tl0-r (float32)
//
// Threshold 0 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl0-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl0R(value float32) error {
	return e.Element.SetProperty("tl0-r", value)
}

// tl1-l (float32)
//
// Threshold 1 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl1-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl1L(value float32) error {
	return e.Element.SetProperty("tl1-l", value)
}

// tl1-r (float32)
//
// Threshold 1 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl1R(value float32) error {
	return e.Element.SetProperty("tl1-r", value)
}

// tl2-l (float32)
//
// Threshold 2 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl2-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl2L(value float32) error {
	return e.Element.SetProperty("tl2-l", value)
}

// tl2-r (float32)
//
// Threshold 2 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl2R(value float32) error {
	return e.Element.SetProperty("tl2-r", value)
}

// tl3-l (float32)
//
// Threshold 3 Left

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl3-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl3L(value float32) error {
	return e.Element.SetProperty("tl3-l", value)
}

// tl3-r (float32)
//
// Threshold 3 Right

func (e *LspPlugInPluginsLv2DynaProcessorLr) GetTl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tl3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2DynaProcessorLr) SetTl3R(value float32) error {
	return e.Element.SetProperty("tl3-r", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2DynaProcessorLrscmR string

const (
	LspPlugInPluginsLv2DynaProcessorLrscmRPeak LspPlugInPluginsLv2DynaProcessorLrscmR = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2DynaProcessorLrscmRRms LspPlugInPluginsLv2DynaProcessorLrscmR = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2DynaProcessorLrscmRLowPass LspPlugInPluginsLv2DynaProcessorLrscmR = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2DynaProcessorLrscmRUniform LspPlugInPluginsLv2DynaProcessorLrscmR = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2DynaProcessorLrscsL string

const (
	LspPlugInPluginsLv2DynaProcessorLrscsLMiddle LspPlugInPluginsLv2DynaProcessorLrscsL = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2DynaProcessorLrscsLSide LspPlugInPluginsLv2DynaProcessorLrscsL = "Side" // Side (1) – Side
	LspPlugInPluginsLv2DynaProcessorLrscsLLeft LspPlugInPluginsLv2DynaProcessorLrscsL = "Left" // Left (2) – Left
	LspPlugInPluginsLv2DynaProcessorLrscsLRight LspPlugInPluginsLv2DynaProcessorLrscsL = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2DynaProcessorLrscsR string

const (
	LspPlugInPluginsLv2DynaProcessorLrscsRMiddle LspPlugInPluginsLv2DynaProcessorLrscsR = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2DynaProcessorLrscsRSide LspPlugInPluginsLv2DynaProcessorLrscsR = "Side" // Side (1) – Side
	LspPlugInPluginsLv2DynaProcessorLrscsRLeft LspPlugInPluginsLv2DynaProcessorLrscsR = "Left" // Left (2) – Left
	LspPlugInPluginsLv2DynaProcessorLrscsRRight LspPlugInPluginsLv2DynaProcessorLrscsR = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2DynaProcessorLrsctL string

const (
	LspPlugInPluginsLv2DynaProcessorLrsctLFeedForward LspPlugInPluginsLv2DynaProcessorLrsctL = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2DynaProcessorLrsctLFeedBack LspPlugInPluginsLv2DynaProcessorLrsctL = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2DynaProcessorLrsctR string

const (
	LspPlugInPluginsLv2DynaProcessorLrsctRFeedForward LspPlugInPluginsLv2DynaProcessorLrsctR = "Feed-forward" // Feed-forward (0) – Feed-forward
	LspPlugInPluginsLv2DynaProcessorLrsctRFeedBack LspPlugInPluginsLv2DynaProcessorLrsctR = "Feed-back" // Feed-back (1) – Feed-back
)

type LspPlugInPluginsLv2DynaProcessorLrpsel string

const (
	LspPlugInPluginsLv2DynaProcessorLrpselLeft LspPlugInPluginsLv2DynaProcessorLrpsel = "Left" // Left (0) – Left
	LspPlugInPluginsLv2DynaProcessorLrpselRight LspPlugInPluginsLv2DynaProcessorLrpsel = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2DynaProcessorLrscmL string

const (
	LspPlugInPluginsLv2DynaProcessorLrscmLPeak LspPlugInPluginsLv2DynaProcessorLrscmL = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2DynaProcessorLrscmLRms LspPlugInPluginsLv2DynaProcessorLrscmL = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2DynaProcessorLrscmLLowPass LspPlugInPluginsLv2DynaProcessorLrscmL = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2DynaProcessorLrscmLUniform LspPlugInPluginsLv2DynaProcessorLrscmL = "Uniform" // Uniform (3) – Uniform
)

