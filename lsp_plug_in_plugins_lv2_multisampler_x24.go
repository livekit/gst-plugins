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

type LspPlugInPluginsLv2MultisamplerX24 struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MultisamplerX24() (*LspPlugInPluginsLv2MultisamplerX24, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-multisampler-x24")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MultisamplerX24{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MultisamplerX24WithName(name string) (*LspPlugInPluginsLv2MultisamplerX24, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-multisampler-x24", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MultisamplerX24{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MultisamplerX24) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2MultisamplerX24) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// fout (float32)
//
// Note-off fadeout

func (e *LspPlugInPluginsLv2MultisamplerX24) GetFout() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetFout(value float32) error {
	return e.Element.SetProperty("fout", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MultisamplerX24) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// iact-0 (bool)
//
// Instrument activity 0

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-1 (bool)
//
// Instrument activity 1

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-10 (bool)
//
// Instrument activity 10

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-11 (bool)
//
// Instrument activity 11

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-12 (bool)
//
// Instrument activity 12

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-13 (bool)
//
// Instrument activity 13

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-14 (bool)
//
// Instrument activity 14

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-15 (bool)
//
// Instrument activity 15

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-16 (bool)
//
// Instrument activity 16

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-17 (bool)
//
// Instrument activity 17

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-18 (bool)
//
// Instrument activity 18

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-19 (bool)
//
// Instrument activity 19

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-2 (bool)
//
// Instrument activity 2

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-20 (bool)
//
// Instrument activity 20

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-21 (bool)
//
// Instrument activity 21

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-22 (bool)
//
// Instrument activity 22

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-23 (bool)
//
// Instrument activity 23

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-3 (bool)
//
// Instrument activity 3

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-4 (bool)
//
// Instrument activity 4

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-5 (bool)
//
// Instrument activity 5

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-6 (bool)
//
// Instrument activity 6

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-7 (bool)
//
// Instrument activity 7

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-8 (bool)
//
// Instrument activity 8

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-9 (bool)
//
// Instrument activity 9

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIact9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// imix-0 (float32)
//
// Instrument mix gain 0

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix0(value float32) error {
	return e.Element.SetProperty("imix-0", value)
}

// imix-1 (float32)
//
// Instrument mix gain 1

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix1(value float32) error {
	return e.Element.SetProperty("imix-1", value)
}

// imix-10 (float32)
//
// Instrument mix gain 10

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix10(value float32) error {
	return e.Element.SetProperty("imix-10", value)
}

// imix-11 (float32)
//
// Instrument mix gain 11

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix11(value float32) error {
	return e.Element.SetProperty("imix-11", value)
}

// imix-12 (float32)
//
// Instrument mix gain 12

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix12(value float32) error {
	return e.Element.SetProperty("imix-12", value)
}

// imix-13 (float32)
//
// Instrument mix gain 13

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix13(value float32) error {
	return e.Element.SetProperty("imix-13", value)
}

// imix-14 (float32)
//
// Instrument mix gain 14

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix14(value float32) error {
	return e.Element.SetProperty("imix-14", value)
}

// imix-15 (float32)
//
// Instrument mix gain 15

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix15(value float32) error {
	return e.Element.SetProperty("imix-15", value)
}

// imix-16 (float32)
//
// Instrument mix gain 16

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix16(value float32) error {
	return e.Element.SetProperty("imix-16", value)
}

// imix-17 (float32)
//
// Instrument mix gain 17

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix17(value float32) error {
	return e.Element.SetProperty("imix-17", value)
}

// imix-18 (float32)
//
// Instrument mix gain 18

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix18(value float32) error {
	return e.Element.SetProperty("imix-18", value)
}

// imix-19 (float32)
//
// Instrument mix gain 19

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix19(value float32) error {
	return e.Element.SetProperty("imix-19", value)
}

// imix-2 (float32)
//
// Instrument mix gain 2

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix2(value float32) error {
	return e.Element.SetProperty("imix-2", value)
}

// imix-20 (float32)
//
// Instrument mix gain 20

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix20(value float32) error {
	return e.Element.SetProperty("imix-20", value)
}

// imix-21 (float32)
//
// Instrument mix gain 21

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix21(value float32) error {
	return e.Element.SetProperty("imix-21", value)
}

// imix-22 (float32)
//
// Instrument mix gain 22

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix22(value float32) error {
	return e.Element.SetProperty("imix-22", value)
}

// imix-23 (float32)
//
// Instrument mix gain 23

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix23(value float32) error {
	return e.Element.SetProperty("imix-23", value)
}

// imix-3 (float32)
//
// Instrument mix gain 3

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix3(value float32) error {
	return e.Element.SetProperty("imix-3", value)
}

// imix-4 (float32)
//
// Instrument mix gain 4

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix4(value float32) error {
	return e.Element.SetProperty("imix-4", value)
}

// imix-5 (float32)
//
// Instrument mix gain 5

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix5(value float32) error {
	return e.Element.SetProperty("imix-5", value)
}

// imix-6 (float32)
//
// Instrument mix gain 6

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix6(value float32) error {
	return e.Element.SetProperty("imix-6", value)
}

// imix-7 (float32)
//
// Instrument mix gain 7

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix7(value float32) error {
	return e.Element.SetProperty("imix-7", value)
}

// imix-8 (float32)
//
// Instrument mix gain 8

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix8(value float32) error {
	return e.Element.SetProperty("imix-8", value)
}

// imix-9 (float32)
//
// Instrument mix gain 9

func (e *LspPlugInPluginsLv2MultisamplerX24) GetImix9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetImix9(value float32) error {
	return e.Element.SetProperty("imix-9", value)
}

// ion-0 (bool)
//
// Instrument on 0

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon0(value bool) error {
	return e.Element.SetProperty("ion-0", value)
}

// ion-1 (bool)
//
// Instrument on 1

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon1(value bool) error {
	return e.Element.SetProperty("ion-1", value)
}

// ion-10 (bool)
//
// Instrument on 10

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon10(value bool) error {
	return e.Element.SetProperty("ion-10", value)
}

// ion-11 (bool)
//
// Instrument on 11

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon11(value bool) error {
	return e.Element.SetProperty("ion-11", value)
}

// ion-12 (bool)
//
// Instrument on 12

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon12(value bool) error {
	return e.Element.SetProperty("ion-12", value)
}

// ion-13 (bool)
//
// Instrument on 13

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon13(value bool) error {
	return e.Element.SetProperty("ion-13", value)
}

// ion-14 (bool)
//
// Instrument on 14

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon14(value bool) error {
	return e.Element.SetProperty("ion-14", value)
}

// ion-15 (bool)
//
// Instrument on 15

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon15(value bool) error {
	return e.Element.SetProperty("ion-15", value)
}

// ion-16 (bool)
//
// Instrument on 16

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon16(value bool) error {
	return e.Element.SetProperty("ion-16", value)
}

// ion-17 (bool)
//
// Instrument on 17

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon17(value bool) error {
	return e.Element.SetProperty("ion-17", value)
}

// ion-18 (bool)
//
// Instrument on 18

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon18(value bool) error {
	return e.Element.SetProperty("ion-18", value)
}

// ion-19 (bool)
//
// Instrument on 19

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon19(value bool) error {
	return e.Element.SetProperty("ion-19", value)
}

// ion-2 (bool)
//
// Instrument on 2

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon2(value bool) error {
	return e.Element.SetProperty("ion-2", value)
}

// ion-20 (bool)
//
// Instrument on 20

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon20(value bool) error {
	return e.Element.SetProperty("ion-20", value)
}

// ion-21 (bool)
//
// Instrument on 21

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon21(value bool) error {
	return e.Element.SetProperty("ion-21", value)
}

// ion-22 (bool)
//
// Instrument on 22

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon22(value bool) error {
	return e.Element.SetProperty("ion-22", value)
}

// ion-23 (bool)
//
// Instrument on 23

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon23(value bool) error {
	return e.Element.SetProperty("ion-23", value)
}

// ion-3 (bool)
//
// Instrument on 3

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon3(value bool) error {
	return e.Element.SetProperty("ion-3", value)
}

// ion-4 (bool)
//
// Instrument on 4

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon4(value bool) error {
	return e.Element.SetProperty("ion-4", value)
}

// ion-5 (bool)
//
// Instrument on 5

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon5(value bool) error {
	return e.Element.SetProperty("ion-5", value)
}

// ion-6 (bool)
//
// Instrument on 6

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon6(value bool) error {
	return e.Element.SetProperty("ion-6", value)
}

// ion-7 (bool)
//
// Instrument on 7

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon7(value bool) error {
	return e.Element.SetProperty("ion-7", value)
}

// ion-8 (bool)
//
// Instrument on 8

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon8(value bool) error {
	return e.Element.SetProperty("ion-8", value)
}

// ion-9 (bool)
//
// Instrument on 9

func (e *LspPlugInPluginsLv2MultisamplerX24) GetIon9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetIon9(value bool) error {
	return e.Element.SetProperty("ion-9", value)
}

// msel (GstLspPlugInPluginsLv2MultisamplerX24Msel)
//
// Area selector

func (e *LspPlugInPluginsLv2MultisamplerX24) GetMsel() (interface{}, error) {
	return e.Element.GetProperty("msel")
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetMsel(value interface{}) error {
	return e.Element.SetProperty("msel", value)
}

// mute (bool)
//
// Forced mute

func (e *LspPlugInPluginsLv2MultisamplerX24) GetMute() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// muting (bool)
//
// Mute on stop

func (e *LspPlugInPluginsLv2MultisamplerX24) GetMuting() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetMuting(value bool) error {
	return e.Element.SetProperty("muting", value)
}

// noff (bool)
//
// Note-off handling

func (e *LspPlugInPluginsLv2MultisamplerX24) GetNoff() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetNoff(value bool) error {
	return e.Element.SetProperty("noff", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2MultisamplerX24) GetOutLatency() (int, error) {
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

// panl-0 (float32)
//
// Instrument panorama left 0

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl0(value float32) error {
	return e.Element.SetProperty("panl-0", value)
}

// panl-1 (float32)
//
// Instrument panorama left 1

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl1(value float32) error {
	return e.Element.SetProperty("panl-1", value)
}

// panl-10 (float32)
//
// Instrument panorama left 10

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl10(value float32) error {
	return e.Element.SetProperty("panl-10", value)
}

// panl-11 (float32)
//
// Instrument panorama left 11

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl11(value float32) error {
	return e.Element.SetProperty("panl-11", value)
}

// panl-12 (float32)
//
// Instrument panorama left 12

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl12(value float32) error {
	return e.Element.SetProperty("panl-12", value)
}

// panl-13 (float32)
//
// Instrument panorama left 13

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl13(value float32) error {
	return e.Element.SetProperty("panl-13", value)
}

// panl-14 (float32)
//
// Instrument panorama left 14

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl14(value float32) error {
	return e.Element.SetProperty("panl-14", value)
}

// panl-15 (float32)
//
// Instrument panorama left 15

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl15(value float32) error {
	return e.Element.SetProperty("panl-15", value)
}

// panl-16 (float32)
//
// Instrument panorama left 16

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl16(value float32) error {
	return e.Element.SetProperty("panl-16", value)
}

// panl-17 (float32)
//
// Instrument panorama left 17

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl17(value float32) error {
	return e.Element.SetProperty("panl-17", value)
}

// panl-18 (float32)
//
// Instrument panorama left 18

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl18(value float32) error {
	return e.Element.SetProperty("panl-18", value)
}

// panl-19 (float32)
//
// Instrument panorama left 19

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl19(value float32) error {
	return e.Element.SetProperty("panl-19", value)
}

// panl-2 (float32)
//
// Instrument panorama left 2

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl2(value float32) error {
	return e.Element.SetProperty("panl-2", value)
}

// panl-20 (float32)
//
// Instrument panorama left 20

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl20(value float32) error {
	return e.Element.SetProperty("panl-20", value)
}

// panl-21 (float32)
//
// Instrument panorama left 21

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl21(value float32) error {
	return e.Element.SetProperty("panl-21", value)
}

// panl-22 (float32)
//
// Instrument panorama left 22

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl22(value float32) error {
	return e.Element.SetProperty("panl-22", value)
}

// panl-23 (float32)
//
// Instrument panorama left 23

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl23(value float32) error {
	return e.Element.SetProperty("panl-23", value)
}

// panl-3 (float32)
//
// Instrument panorama left 3

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl3(value float32) error {
	return e.Element.SetProperty("panl-3", value)
}

// panl-4 (float32)
//
// Instrument panorama left 4

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl4(value float32) error {
	return e.Element.SetProperty("panl-4", value)
}

// panl-5 (float32)
//
// Instrument panorama left 5

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl5(value float32) error {
	return e.Element.SetProperty("panl-5", value)
}

// panl-6 (float32)
//
// Instrument panorama left 6

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl6(value float32) error {
	return e.Element.SetProperty("panl-6", value)
}

// panl-7 (float32)
//
// Instrument panorama left 7

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl7(value float32) error {
	return e.Element.SetProperty("panl-7", value)
}

// panl-8 (float32)
//
// Instrument panorama left 8

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl8(value float32) error {
	return e.Element.SetProperty("panl-8", value)
}

// panl-9 (float32)
//
// Instrument panorama left 9

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanl9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanl9(value float32) error {
	return e.Element.SetProperty("panl-9", value)
}

// panr-0 (float32)
//
// Instrument manorama right 0

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr0(value float32) error {
	return e.Element.SetProperty("panr-0", value)
}

// panr-1 (float32)
//
// Instrument manorama right 1

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr1(value float32) error {
	return e.Element.SetProperty("panr-1", value)
}

// panr-10 (float32)
//
// Instrument manorama right 10

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr10(value float32) error {
	return e.Element.SetProperty("panr-10", value)
}

// panr-11 (float32)
//
// Instrument manorama right 11

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr11(value float32) error {
	return e.Element.SetProperty("panr-11", value)
}

// panr-12 (float32)
//
// Instrument manorama right 12

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr12(value float32) error {
	return e.Element.SetProperty("panr-12", value)
}

// panr-13 (float32)
//
// Instrument manorama right 13

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr13(value float32) error {
	return e.Element.SetProperty("panr-13", value)
}

// panr-14 (float32)
//
// Instrument manorama right 14

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr14(value float32) error {
	return e.Element.SetProperty("panr-14", value)
}

// panr-15 (float32)
//
// Instrument manorama right 15

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr15(value float32) error {
	return e.Element.SetProperty("panr-15", value)
}

// panr-16 (float32)
//
// Instrument manorama right 16

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr16(value float32) error {
	return e.Element.SetProperty("panr-16", value)
}

// panr-17 (float32)
//
// Instrument manorama right 17

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr17(value float32) error {
	return e.Element.SetProperty("panr-17", value)
}

// panr-18 (float32)
//
// Instrument manorama right 18

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr18(value float32) error {
	return e.Element.SetProperty("panr-18", value)
}

// panr-19 (float32)
//
// Instrument manorama right 19

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr19(value float32) error {
	return e.Element.SetProperty("panr-19", value)
}

// panr-2 (float32)
//
// Instrument manorama right 2

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr2(value float32) error {
	return e.Element.SetProperty("panr-2", value)
}

// panr-20 (float32)
//
// Instrument manorama right 20

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr20(value float32) error {
	return e.Element.SetProperty("panr-20", value)
}

// panr-21 (float32)
//
// Instrument manorama right 21

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr21(value float32) error {
	return e.Element.SetProperty("panr-21", value)
}

// panr-22 (float32)
//
// Instrument manorama right 22

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr22(value float32) error {
	return e.Element.SetProperty("panr-22", value)
}

// panr-23 (float32)
//
// Instrument manorama right 23

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr23(value float32) error {
	return e.Element.SetProperty("panr-23", value)
}

// panr-3 (float32)
//
// Instrument manorama right 3

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr3(value float32) error {
	return e.Element.SetProperty("panr-3", value)
}

// panr-4 (float32)
//
// Instrument manorama right 4

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr4(value float32) error {
	return e.Element.SetProperty("panr-4", value)
}

// panr-5 (float32)
//
// Instrument manorama right 5

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr5(value float32) error {
	return e.Element.SetProperty("panr-5", value)
}

// panr-6 (float32)
//
// Instrument manorama right 6

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr6(value float32) error {
	return e.Element.SetProperty("panr-6", value)
}

// panr-7 (float32)
//
// Instrument manorama right 7

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr7(value float32) error {
	return e.Element.SetProperty("panr-7", value)
}

// panr-8 (float32)
//
// Instrument manorama right 8

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr8(value float32) error {
	return e.Element.SetProperty("panr-8", value)
}

// panr-9 (float32)
//
// Instrument manorama right 9

func (e *LspPlugInPluginsLv2MultisamplerX24) GetPanr9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX24) SetPanr9(value float32) error {
	return e.Element.SetProperty("panr-9", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2MultisamplerX24) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX24) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MultisamplerX24Msel string

const (
	LspPlugInPluginsLv2MultisamplerX24MselInstruments LspPlugInPluginsLv2MultisamplerX24Msel = "Instruments" // Instruments (0) – Instruments
	LspPlugInPluginsLv2MultisamplerX24MselMixer011 LspPlugInPluginsLv2MultisamplerX24Msel = "Mixer 0-11" // Mixer 0-11 (1) – Mixer 0-11
	LspPlugInPluginsLv2MultisamplerX24MselMixer1223 LspPlugInPluginsLv2MultisamplerX24Msel = "Mixer 12-23" // Mixer 12-23 (2) – Mixer 12-23
)

