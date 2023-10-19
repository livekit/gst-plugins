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

type LspPlugInPluginsLv2ParaEqualizerX32Mono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX32Mono() (*LspPlugInPluginsLv2ParaEqualizerX32Mono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x32-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Mono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX32MonoWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX32Mono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x32-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Mono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// f-0 (float32)
//
// Frequency 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF0(value float32) error {
	return e.Element.SetProperty("f-0", value)
}

// f-1 (float32)
//
// Frequency 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF1(value float32) error {
	return e.Element.SetProperty("f-1", value)
}

// f-10 (float32)
//
// Frequency 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF10(value float32) error {
	return e.Element.SetProperty("f-10", value)
}

// f-11 (float32)
//
// Frequency 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF11(value float32) error {
	return e.Element.SetProperty("f-11", value)
}

// f-12 (float32)
//
// Frequency 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF12(value float32) error {
	return e.Element.SetProperty("f-12", value)
}

// f-13 (float32)
//
// Frequency 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF13(value float32) error {
	return e.Element.SetProperty("f-13", value)
}

// f-14 (float32)
//
// Frequency 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF14(value float32) error {
	return e.Element.SetProperty("f-14", value)
}

// f-15 (float32)
//
// Frequency 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF15(value float32) error {
	return e.Element.SetProperty("f-15", value)
}

// f-16 (float32)
//
// Frequency 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF16(value float32) error {
	return e.Element.SetProperty("f-16", value)
}

// f-17 (float32)
//
// Frequency 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF17(value float32) error {
	return e.Element.SetProperty("f-17", value)
}

// f-18 (float32)
//
// Frequency 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF18(value float32) error {
	return e.Element.SetProperty("f-18", value)
}

// f-19 (float32)
//
// Frequency 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF19(value float32) error {
	return e.Element.SetProperty("f-19", value)
}

// f-2 (float32)
//
// Frequency 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF2(value float32) error {
	return e.Element.SetProperty("f-2", value)
}

// f-20 (float32)
//
// Frequency 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF20(value float32) error {
	return e.Element.SetProperty("f-20", value)
}

// f-21 (float32)
//
// Frequency 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF21(value float32) error {
	return e.Element.SetProperty("f-21", value)
}

// f-22 (float32)
//
// Frequency 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF22(value float32) error {
	return e.Element.SetProperty("f-22", value)
}

// f-23 (float32)
//
// Frequency 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF23(value float32) error {
	return e.Element.SetProperty("f-23", value)
}

// f-24 (float32)
//
// Frequency 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF24(value float32) error {
	return e.Element.SetProperty("f-24", value)
}

// f-25 (float32)
//
// Frequency 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF25(value float32) error {
	return e.Element.SetProperty("f-25", value)
}

// f-26 (float32)
//
// Frequency 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF26(value float32) error {
	return e.Element.SetProperty("f-26", value)
}

// f-27 (float32)
//
// Frequency 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF27(value float32) error {
	return e.Element.SetProperty("f-27", value)
}

// f-28 (float32)
//
// Frequency 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF28(value float32) error {
	return e.Element.SetProperty("f-28", value)
}

// f-29 (float32)
//
// Frequency 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF29(value float32) error {
	return e.Element.SetProperty("f-29", value)
}

// f-3 (float32)
//
// Frequency 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF3(value float32) error {
	return e.Element.SetProperty("f-3", value)
}

// f-30 (float32)
//
// Frequency 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF30(value float32) error {
	return e.Element.SetProperty("f-30", value)
}

// f-31 (float32)
//
// Frequency 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF31(value float32) error {
	return e.Element.SetProperty("f-31", value)
}

// f-4 (float32)
//
// Frequency 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF4(value float32) error {
	return e.Element.SetProperty("f-4", value)
}

// f-5 (float32)
//
// Frequency 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF5(value float32) error {
	return e.Element.SetProperty("f-5", value)
}

// f-6 (float32)
//
// Frequency 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF6(value float32) error {
	return e.Element.SetProperty("f-6", value)
}

// f-7 (float32)
//
// Frequency 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF7(value float32) error {
	return e.Element.SetProperty("f-7", value)
}

// f-8 (float32)
//
// Frequency 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF8(value float32) error {
	return e.Element.SetProperty("f-8", value)
}

// f-9 (float32)
//
// Frequency 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetF9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("f-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetF9(value float32) error {
	return e.Element.SetProperty("f-9", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX32Monofft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fm-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm0)
//
// Filter mode 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm0() (interface{}, error) {
	return e.Element.GetProperty("fm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm0(value interface{}) error {
	return e.Element.SetProperty("fm-0", value)
}

// fm-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm1)
//
// Filter mode 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm1() (interface{}, error) {
	return e.Element.GetProperty("fm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm1(value interface{}) error {
	return e.Element.SetProperty("fm-1", value)
}

// fm-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm10)
//
// Filter mode 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm10() (interface{}, error) {
	return e.Element.GetProperty("fm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm10(value interface{}) error {
	return e.Element.SetProperty("fm-10", value)
}

// fm-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm11)
//
// Filter mode 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm11() (interface{}, error) {
	return e.Element.GetProperty("fm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm11(value interface{}) error {
	return e.Element.SetProperty("fm-11", value)
}

// fm-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm12)
//
// Filter mode 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm12() (interface{}, error) {
	return e.Element.GetProperty("fm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm12(value interface{}) error {
	return e.Element.SetProperty("fm-12", value)
}

// fm-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm13)
//
// Filter mode 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm13() (interface{}, error) {
	return e.Element.GetProperty("fm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm13(value interface{}) error {
	return e.Element.SetProperty("fm-13", value)
}

// fm-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm14)
//
// Filter mode 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm14() (interface{}, error) {
	return e.Element.GetProperty("fm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm14(value interface{}) error {
	return e.Element.SetProperty("fm-14", value)
}

// fm-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm15)
//
// Filter mode 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm15() (interface{}, error) {
	return e.Element.GetProperty("fm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm15(value interface{}) error {
	return e.Element.SetProperty("fm-15", value)
}

// fm-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm16)
//
// Filter mode 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm16() (interface{}, error) {
	return e.Element.GetProperty("fm-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm16(value interface{}) error {
	return e.Element.SetProperty("fm-16", value)
}

// fm-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm17)
//
// Filter mode 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm17() (interface{}, error) {
	return e.Element.GetProperty("fm-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm17(value interface{}) error {
	return e.Element.SetProperty("fm-17", value)
}

// fm-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm18)
//
// Filter mode 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm18() (interface{}, error) {
	return e.Element.GetProperty("fm-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm18(value interface{}) error {
	return e.Element.SetProperty("fm-18", value)
}

// fm-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm19)
//
// Filter mode 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm19() (interface{}, error) {
	return e.Element.GetProperty("fm-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm19(value interface{}) error {
	return e.Element.SetProperty("fm-19", value)
}

// fm-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm2)
//
// Filter mode 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm2() (interface{}, error) {
	return e.Element.GetProperty("fm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm2(value interface{}) error {
	return e.Element.SetProperty("fm-2", value)
}

// fm-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm20)
//
// Filter mode 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm20() (interface{}, error) {
	return e.Element.GetProperty("fm-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm20(value interface{}) error {
	return e.Element.SetProperty("fm-20", value)
}

// fm-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm21)
//
// Filter mode 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm21() (interface{}, error) {
	return e.Element.GetProperty("fm-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm21(value interface{}) error {
	return e.Element.SetProperty("fm-21", value)
}

// fm-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm22)
//
// Filter mode 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm22() (interface{}, error) {
	return e.Element.GetProperty("fm-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm22(value interface{}) error {
	return e.Element.SetProperty("fm-22", value)
}

// fm-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm23)
//
// Filter mode 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm23() (interface{}, error) {
	return e.Element.GetProperty("fm-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm23(value interface{}) error {
	return e.Element.SetProperty("fm-23", value)
}

// fm-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm24)
//
// Filter mode 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm24() (interface{}, error) {
	return e.Element.GetProperty("fm-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm24(value interface{}) error {
	return e.Element.SetProperty("fm-24", value)
}

// fm-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm25)
//
// Filter mode 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm25() (interface{}, error) {
	return e.Element.GetProperty("fm-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm25(value interface{}) error {
	return e.Element.SetProperty("fm-25", value)
}

// fm-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm26)
//
// Filter mode 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm26() (interface{}, error) {
	return e.Element.GetProperty("fm-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm26(value interface{}) error {
	return e.Element.SetProperty("fm-26", value)
}

// fm-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm27)
//
// Filter mode 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm27() (interface{}, error) {
	return e.Element.GetProperty("fm-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm27(value interface{}) error {
	return e.Element.SetProperty("fm-27", value)
}

// fm-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm28)
//
// Filter mode 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm28() (interface{}, error) {
	return e.Element.GetProperty("fm-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm28(value interface{}) error {
	return e.Element.SetProperty("fm-28", value)
}

// fm-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm29)
//
// Filter mode 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm29() (interface{}, error) {
	return e.Element.GetProperty("fm-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm29(value interface{}) error {
	return e.Element.SetProperty("fm-29", value)
}

// fm-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm3)
//
// Filter mode 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm3() (interface{}, error) {
	return e.Element.GetProperty("fm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm3(value interface{}) error {
	return e.Element.SetProperty("fm-3", value)
}

// fm-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm30)
//
// Filter mode 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm30() (interface{}, error) {
	return e.Element.GetProperty("fm-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm30(value interface{}) error {
	return e.Element.SetProperty("fm-30", value)
}

// fm-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm31)
//
// Filter mode 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm31() (interface{}, error) {
	return e.Element.GetProperty("fm-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm31(value interface{}) error {
	return e.Element.SetProperty("fm-31", value)
}

// fm-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm4)
//
// Filter mode 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm4() (interface{}, error) {
	return e.Element.GetProperty("fm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm4(value interface{}) error {
	return e.Element.SetProperty("fm-4", value)
}

// fm-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm5)
//
// Filter mode 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm5() (interface{}, error) {
	return e.Element.GetProperty("fm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm5(value interface{}) error {
	return e.Element.SetProperty("fm-5", value)
}

// fm-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm6)
//
// Filter mode 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm6() (interface{}, error) {
	return e.Element.GetProperty("fm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm6(value interface{}) error {
	return e.Element.SetProperty("fm-6", value)
}

// fm-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm7)
//
// Filter mode 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm7() (interface{}, error) {
	return e.Element.GetProperty("fm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm7(value interface{}) error {
	return e.Element.SetProperty("fm-7", value)
}

// fm-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm8)
//
// Filter mode 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm8() (interface{}, error) {
	return e.Element.GetProperty("fm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm8(value interface{}) error {
	return e.Element.SetProperty("fm-8", value)
}

// fm-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Monofm9)
//
// Filter mode 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFm9() (interface{}, error) {
	return e.Element.GetProperty("fm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFm9(value interface{}) error {
	return e.Element.SetProperty("fm-9", value)
}

// frqs (float32)
//
// Frequency shift

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFrqs() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frqs")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFrqs(value float32) error {
	return e.Element.SetProperty("frqs", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX32Monofsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ft-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft0)
//
// Filter type 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt0() (interface{}, error) {
	return e.Element.GetProperty("ft-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt0(value interface{}) error {
	return e.Element.SetProperty("ft-0", value)
}

// ft-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft1)
//
// Filter type 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt1() (interface{}, error) {
	return e.Element.GetProperty("ft-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt1(value interface{}) error {
	return e.Element.SetProperty("ft-1", value)
}

// ft-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft10)
//
// Filter type 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt10() (interface{}, error) {
	return e.Element.GetProperty("ft-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt10(value interface{}) error {
	return e.Element.SetProperty("ft-10", value)
}

// ft-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft11)
//
// Filter type 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt11() (interface{}, error) {
	return e.Element.GetProperty("ft-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt11(value interface{}) error {
	return e.Element.SetProperty("ft-11", value)
}

// ft-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft12)
//
// Filter type 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt12() (interface{}, error) {
	return e.Element.GetProperty("ft-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt12(value interface{}) error {
	return e.Element.SetProperty("ft-12", value)
}

// ft-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft13)
//
// Filter type 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt13() (interface{}, error) {
	return e.Element.GetProperty("ft-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt13(value interface{}) error {
	return e.Element.SetProperty("ft-13", value)
}

// ft-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft14)
//
// Filter type 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt14() (interface{}, error) {
	return e.Element.GetProperty("ft-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt14(value interface{}) error {
	return e.Element.SetProperty("ft-14", value)
}

// ft-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft15)
//
// Filter type 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt15() (interface{}, error) {
	return e.Element.GetProperty("ft-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt15(value interface{}) error {
	return e.Element.SetProperty("ft-15", value)
}

// ft-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft16)
//
// Filter type 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt16() (interface{}, error) {
	return e.Element.GetProperty("ft-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt16(value interface{}) error {
	return e.Element.SetProperty("ft-16", value)
}

// ft-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft17)
//
// Filter type 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt17() (interface{}, error) {
	return e.Element.GetProperty("ft-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt17(value interface{}) error {
	return e.Element.SetProperty("ft-17", value)
}

// ft-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft18)
//
// Filter type 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt18() (interface{}, error) {
	return e.Element.GetProperty("ft-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt18(value interface{}) error {
	return e.Element.SetProperty("ft-18", value)
}

// ft-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft19)
//
// Filter type 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt19() (interface{}, error) {
	return e.Element.GetProperty("ft-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt19(value interface{}) error {
	return e.Element.SetProperty("ft-19", value)
}

// ft-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft2)
//
// Filter type 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt2() (interface{}, error) {
	return e.Element.GetProperty("ft-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt2(value interface{}) error {
	return e.Element.SetProperty("ft-2", value)
}

// ft-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft20)
//
// Filter type 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt20() (interface{}, error) {
	return e.Element.GetProperty("ft-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt20(value interface{}) error {
	return e.Element.SetProperty("ft-20", value)
}

// ft-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft21)
//
// Filter type 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt21() (interface{}, error) {
	return e.Element.GetProperty("ft-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt21(value interface{}) error {
	return e.Element.SetProperty("ft-21", value)
}

// ft-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft22)
//
// Filter type 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt22() (interface{}, error) {
	return e.Element.GetProperty("ft-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt22(value interface{}) error {
	return e.Element.SetProperty("ft-22", value)
}

// ft-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft23)
//
// Filter type 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt23() (interface{}, error) {
	return e.Element.GetProperty("ft-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt23(value interface{}) error {
	return e.Element.SetProperty("ft-23", value)
}

// ft-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft24)
//
// Filter type 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt24() (interface{}, error) {
	return e.Element.GetProperty("ft-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt24(value interface{}) error {
	return e.Element.SetProperty("ft-24", value)
}

// ft-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft25)
//
// Filter type 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt25() (interface{}, error) {
	return e.Element.GetProperty("ft-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt25(value interface{}) error {
	return e.Element.SetProperty("ft-25", value)
}

// ft-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft26)
//
// Filter type 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt26() (interface{}, error) {
	return e.Element.GetProperty("ft-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt26(value interface{}) error {
	return e.Element.SetProperty("ft-26", value)
}

// ft-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft27)
//
// Filter type 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt27() (interface{}, error) {
	return e.Element.GetProperty("ft-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt27(value interface{}) error {
	return e.Element.SetProperty("ft-27", value)
}

// ft-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft28)
//
// Filter type 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt28() (interface{}, error) {
	return e.Element.GetProperty("ft-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt28(value interface{}) error {
	return e.Element.SetProperty("ft-28", value)
}

// ft-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft29)
//
// Filter type 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt29() (interface{}, error) {
	return e.Element.GetProperty("ft-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt29(value interface{}) error {
	return e.Element.SetProperty("ft-29", value)
}

// ft-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft3)
//
// Filter type 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt3() (interface{}, error) {
	return e.Element.GetProperty("ft-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt3(value interface{}) error {
	return e.Element.SetProperty("ft-3", value)
}

// ft-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft30)
//
// Filter type 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt30() (interface{}, error) {
	return e.Element.GetProperty("ft-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt30(value interface{}) error {
	return e.Element.SetProperty("ft-30", value)
}

// ft-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft31)
//
// Filter type 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt31() (interface{}, error) {
	return e.Element.GetProperty("ft-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt31(value interface{}) error {
	return e.Element.SetProperty("ft-31", value)
}

// ft-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft4)
//
// Filter type 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt4() (interface{}, error) {
	return e.Element.GetProperty("ft-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt4(value interface{}) error {
	return e.Element.SetProperty("ft-4", value)
}

// ft-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft5)
//
// Filter type 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt5() (interface{}, error) {
	return e.Element.GetProperty("ft-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt5(value interface{}) error {
	return e.Element.SetProperty("ft-5", value)
}

// ft-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft6)
//
// Filter type 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt6() (interface{}, error) {
	return e.Element.GetProperty("ft-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt6(value interface{}) error {
	return e.Element.SetProperty("ft-6", value)
}

// ft-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft7)
//
// Filter type 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt7() (interface{}, error) {
	return e.Element.GetProperty("ft-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt7(value interface{}) error {
	return e.Element.SetProperty("ft-7", value)
}

// ft-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft8)
//
// Filter type 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt8() (interface{}, error) {
	return e.Element.GetProperty("ft-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt8(value interface{}) error {
	return e.Element.SetProperty("ft-8", value)
}

// ft-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Monoft9)
//
// Filter type 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFt9() (interface{}, error) {
	return e.Element.GetProperty("ft-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetFt9(value interface{}) error {
	return e.Element.SetProperty("ft-9", value)
}

// fv-0 (bool)
//
// Filter visibility 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-1 (bool)
//
// Filter visibility 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-10 (bool)
//
// Filter visibility 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-11 (bool)
//
// Filter visibility 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-12 (bool)
//
// Filter visibility 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-13 (bool)
//
// Filter visibility 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-14 (bool)
//
// Filter visibility 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-15 (bool)
//
// Filter visibility 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-16 (bool)
//
// Filter visibility 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-17 (bool)
//
// Filter visibility 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-18 (bool)
//
// Filter visibility 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-19 (bool)
//
// Filter visibility 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-2 (bool)
//
// Filter visibility 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-20 (bool)
//
// Filter visibility 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-21 (bool)
//
// Filter visibility 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-22 (bool)
//
// Filter visibility 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-23 (bool)
//
// Filter visibility 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-24 (bool)
//
// Filter visibility 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-25 (bool)
//
// Filter visibility 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-26 (bool)
//
// Filter visibility 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-27 (bool)
//
// Filter visibility 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-28 (bool)
//
// Filter visibility 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-29 (bool)
//
// Filter visibility 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-3 (bool)
//
// Filter visibility 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-30 (bool)
//
// Filter visibility 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-31 (bool)
//
// Filter visibility 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-4 (bool)
//
// Filter visibility 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-5 (bool)
//
// Filter visibility 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-6 (bool)
//
// Filter visibility 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-7 (bool)
//
// Filter visibility 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-8 (bool)
//
// Filter visibility 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-9 (bool)
//
// Filter visibility 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetFv9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// g-0 (float32)
//
// Gain 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG0(value float32) error {
	return e.Element.SetProperty("g-0", value)
}

// g-1 (float32)
//
// Gain 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG1(value float32) error {
	return e.Element.SetProperty("g-1", value)
}

// g-10 (float32)
//
// Gain 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG10(value float32) error {
	return e.Element.SetProperty("g-10", value)
}

// g-11 (float32)
//
// Gain 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG11(value float32) error {
	return e.Element.SetProperty("g-11", value)
}

// g-12 (float32)
//
// Gain 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG12(value float32) error {
	return e.Element.SetProperty("g-12", value)
}

// g-13 (float32)
//
// Gain 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG13(value float32) error {
	return e.Element.SetProperty("g-13", value)
}

// g-14 (float32)
//
// Gain 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG14(value float32) error {
	return e.Element.SetProperty("g-14", value)
}

// g-15 (float32)
//
// Gain 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG15(value float32) error {
	return e.Element.SetProperty("g-15", value)
}

// g-16 (float32)
//
// Gain 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG16(value float32) error {
	return e.Element.SetProperty("g-16", value)
}

// g-17 (float32)
//
// Gain 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG17(value float32) error {
	return e.Element.SetProperty("g-17", value)
}

// g-18 (float32)
//
// Gain 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG18(value float32) error {
	return e.Element.SetProperty("g-18", value)
}

// g-19 (float32)
//
// Gain 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG19(value float32) error {
	return e.Element.SetProperty("g-19", value)
}

// g-2 (float32)
//
// Gain 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG2(value float32) error {
	return e.Element.SetProperty("g-2", value)
}

// g-20 (float32)
//
// Gain 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG20(value float32) error {
	return e.Element.SetProperty("g-20", value)
}

// g-21 (float32)
//
// Gain 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG21(value float32) error {
	return e.Element.SetProperty("g-21", value)
}

// g-22 (float32)
//
// Gain 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG22(value float32) error {
	return e.Element.SetProperty("g-22", value)
}

// g-23 (float32)
//
// Gain 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG23(value float32) error {
	return e.Element.SetProperty("g-23", value)
}

// g-24 (float32)
//
// Gain 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG24(value float32) error {
	return e.Element.SetProperty("g-24", value)
}

// g-25 (float32)
//
// Gain 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG25(value float32) error {
	return e.Element.SetProperty("g-25", value)
}

// g-26 (float32)
//
// Gain 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG26(value float32) error {
	return e.Element.SetProperty("g-26", value)
}

// g-27 (float32)
//
// Gain 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG27(value float32) error {
	return e.Element.SetProperty("g-27", value)
}

// g-28 (float32)
//
// Gain 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG28(value float32) error {
	return e.Element.SetProperty("g-28", value)
}

// g-29 (float32)
//
// Gain 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG29(value float32) error {
	return e.Element.SetProperty("g-29", value)
}

// g-3 (float32)
//
// Gain 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG3(value float32) error {
	return e.Element.SetProperty("g-3", value)
}

// g-30 (float32)
//
// Gain 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG30(value float32) error {
	return e.Element.SetProperty("g-30", value)
}

// g-31 (float32)
//
// Gain 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG31(value float32) error {
	return e.Element.SetProperty("g-31", value)
}

// g-4 (float32)
//
// Gain 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG4(value float32) error {
	return e.Element.SetProperty("g-4", value)
}

// g-5 (float32)
//
// Gain 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG5(value float32) error {
	return e.Element.SetProperty("g-5", value)
}

// g-6 (float32)
//
// Gain 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG6(value float32) error {
	return e.Element.SetProperty("g-6", value)
}

// g-7 (float32)
//
// Gain 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG7(value float32) error {
	return e.Element.SetProperty("g-7", value)
}

// g-8 (float32)
//
// Gain 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG8(value float32) error {
	return e.Element.SetProperty("g-8", value)
}

// g-9 (float32)
//
// Gain 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetG9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetG9(value float32) error {
	return e.Element.SetProperty("g-9", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hue-0 (float32)
//
// Hue 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// hue-1 (float32)
//
// Hue 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue1(value float32) error {
	return e.Element.SetProperty("hue-1", value)
}

// hue-10 (float32)
//
// Hue 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue10(value float32) error {
	return e.Element.SetProperty("hue-10", value)
}

// hue-11 (float32)
//
// Hue 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue11(value float32) error {
	return e.Element.SetProperty("hue-11", value)
}

// hue-12 (float32)
//
// Hue 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue12(value float32) error {
	return e.Element.SetProperty("hue-12", value)
}

// hue-13 (float32)
//
// Hue 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue13(value float32) error {
	return e.Element.SetProperty("hue-13", value)
}

// hue-14 (float32)
//
// Hue 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue14(value float32) error {
	return e.Element.SetProperty("hue-14", value)
}

// hue-15 (float32)
//
// Hue 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue15(value float32) error {
	return e.Element.SetProperty("hue-15", value)
}

// hue-16 (float32)
//
// Hue 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue16(value float32) error {
	return e.Element.SetProperty("hue-16", value)
}

// hue-17 (float32)
//
// Hue 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue17(value float32) error {
	return e.Element.SetProperty("hue-17", value)
}

// hue-18 (float32)
//
// Hue 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue18(value float32) error {
	return e.Element.SetProperty("hue-18", value)
}

// hue-19 (float32)
//
// Hue 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue19(value float32) error {
	return e.Element.SetProperty("hue-19", value)
}

// hue-2 (float32)
//
// Hue 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue2(value float32) error {
	return e.Element.SetProperty("hue-2", value)
}

// hue-20 (float32)
//
// Hue 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue20(value float32) error {
	return e.Element.SetProperty("hue-20", value)
}

// hue-21 (float32)
//
// Hue 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue21(value float32) error {
	return e.Element.SetProperty("hue-21", value)
}

// hue-22 (float32)
//
// Hue 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue22(value float32) error {
	return e.Element.SetProperty("hue-22", value)
}

// hue-23 (float32)
//
// Hue 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue23(value float32) error {
	return e.Element.SetProperty("hue-23", value)
}

// hue-24 (float32)
//
// Hue 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue24(value float32) error {
	return e.Element.SetProperty("hue-24", value)
}

// hue-25 (float32)
//
// Hue 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue25(value float32) error {
	return e.Element.SetProperty("hue-25", value)
}

// hue-26 (float32)
//
// Hue 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue26(value float32) error {
	return e.Element.SetProperty("hue-26", value)
}

// hue-27 (float32)
//
// Hue 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue27(value float32) error {
	return e.Element.SetProperty("hue-27", value)
}

// hue-28 (float32)
//
// Hue 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue28(value float32) error {
	return e.Element.SetProperty("hue-28", value)
}

// hue-29 (float32)
//
// Hue 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue29(value float32) error {
	return e.Element.SetProperty("hue-29", value)
}

// hue-3 (float32)
//
// Hue 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue3(value float32) error {
	return e.Element.SetProperty("hue-3", value)
}

// hue-30 (float32)
//
// Hue 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue30(value float32) error {
	return e.Element.SetProperty("hue-30", value)
}

// hue-31 (float32)
//
// Hue 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue31(value float32) error {
	return e.Element.SetProperty("hue-31", value)
}

// hue-4 (float32)
//
// Hue 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue4(value float32) error {
	return e.Element.SetProperty("hue-4", value)
}

// hue-5 (float32)
//
// Hue 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue5(value float32) error {
	return e.Element.SetProperty("hue-5", value)
}

// hue-6 (float32)
//
// Hue 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue6(value float32) error {
	return e.Element.SetProperty("hue-6", value)
}

// hue-7 (float32)
//
// Hue 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue7(value float32) error {
	return e.Element.SetProperty("hue-7", value)
}

// hue-8 (float32)
//
// Hue 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue8(value float32) error {
	return e.Element.SetProperty("hue-8", value)
}

// hue-9 (float32)
//
// Hue 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetHue9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetHue9(value float32) error {
	return e.Element.SetProperty("hue-9", value)
}

// im (float32)
//
// Input signal meter

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetIm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("im")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// mode (GstLspPlugInPluginsLv2ParaEqualizerX32Monomode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetOutLatency() (int, error) {
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

// q-0 (float32)
//
// Quality factor 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ0(value float32) error {
	return e.Element.SetProperty("q-0", value)
}

// q-1 (float32)
//
// Quality factor 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ1(value float32) error {
	return e.Element.SetProperty("q-1", value)
}

// q-10 (float32)
//
// Quality factor 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ10(value float32) error {
	return e.Element.SetProperty("q-10", value)
}

// q-11 (float32)
//
// Quality factor 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ11(value float32) error {
	return e.Element.SetProperty("q-11", value)
}

// q-12 (float32)
//
// Quality factor 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ12(value float32) error {
	return e.Element.SetProperty("q-12", value)
}

// q-13 (float32)
//
// Quality factor 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ13(value float32) error {
	return e.Element.SetProperty("q-13", value)
}

// q-14 (float32)
//
// Quality factor 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ14(value float32) error {
	return e.Element.SetProperty("q-14", value)
}

// q-15 (float32)
//
// Quality factor 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ15(value float32) error {
	return e.Element.SetProperty("q-15", value)
}

// q-16 (float32)
//
// Quality factor 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ16(value float32) error {
	return e.Element.SetProperty("q-16", value)
}

// q-17 (float32)
//
// Quality factor 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ17(value float32) error {
	return e.Element.SetProperty("q-17", value)
}

// q-18 (float32)
//
// Quality factor 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ18(value float32) error {
	return e.Element.SetProperty("q-18", value)
}

// q-19 (float32)
//
// Quality factor 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ19(value float32) error {
	return e.Element.SetProperty("q-19", value)
}

// q-2 (float32)
//
// Quality factor 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ2(value float32) error {
	return e.Element.SetProperty("q-2", value)
}

// q-20 (float32)
//
// Quality factor 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ20(value float32) error {
	return e.Element.SetProperty("q-20", value)
}

// q-21 (float32)
//
// Quality factor 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ21(value float32) error {
	return e.Element.SetProperty("q-21", value)
}

// q-22 (float32)
//
// Quality factor 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ22(value float32) error {
	return e.Element.SetProperty("q-22", value)
}

// q-23 (float32)
//
// Quality factor 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ23(value float32) error {
	return e.Element.SetProperty("q-23", value)
}

// q-24 (float32)
//
// Quality factor 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ24(value float32) error {
	return e.Element.SetProperty("q-24", value)
}

// q-25 (float32)
//
// Quality factor 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ25(value float32) error {
	return e.Element.SetProperty("q-25", value)
}

// q-26 (float32)
//
// Quality factor 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ26(value float32) error {
	return e.Element.SetProperty("q-26", value)
}

// q-27 (float32)
//
// Quality factor 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ27(value float32) error {
	return e.Element.SetProperty("q-27", value)
}

// q-28 (float32)
//
// Quality factor 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ28(value float32) error {
	return e.Element.SetProperty("q-28", value)
}

// q-29 (float32)
//
// Quality factor 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ29(value float32) error {
	return e.Element.SetProperty("q-29", value)
}

// q-3 (float32)
//
// Quality factor 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ3(value float32) error {
	return e.Element.SetProperty("q-3", value)
}

// q-30 (float32)
//
// Quality factor 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ30(value float32) error {
	return e.Element.SetProperty("q-30", value)
}

// q-31 (float32)
//
// Quality factor 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ31(value float32) error {
	return e.Element.SetProperty("q-31", value)
}

// q-4 (float32)
//
// Quality factor 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ4(value float32) error {
	return e.Element.SetProperty("q-4", value)
}

// q-5 (float32)
//
// Quality factor 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ5(value float32) error {
	return e.Element.SetProperty("q-5", value)
}

// q-6 (float32)
//
// Quality factor 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ6(value float32) error {
	return e.Element.SetProperty("q-6", value)
}

// q-7 (float32)
//
// Quality factor 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ7(value float32) error {
	return e.Element.SetProperty("q-7", value)
}

// q-8 (float32)
//
// Quality factor 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ8(value float32) error {
	return e.Element.SetProperty("q-8", value)
}

// q-9 (float32)
//
// Quality factor 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetQ9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("q-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetQ9(value float32) error {
	return e.Element.SetProperty("q-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetReact() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("react")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// s-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos0)
//
// Filter slope 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS0() (interface{}, error) {
	return e.Element.GetProperty("s-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS0(value interface{}) error {
	return e.Element.SetProperty("s-0", value)
}

// s-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos1)
//
// Filter slope 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS1() (interface{}, error) {
	return e.Element.GetProperty("s-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS1(value interface{}) error {
	return e.Element.SetProperty("s-1", value)
}

// s-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos10)
//
// Filter slope 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS10() (interface{}, error) {
	return e.Element.GetProperty("s-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS10(value interface{}) error {
	return e.Element.SetProperty("s-10", value)
}

// s-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos11)
//
// Filter slope 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS11() (interface{}, error) {
	return e.Element.GetProperty("s-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS11(value interface{}) error {
	return e.Element.SetProperty("s-11", value)
}

// s-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos12)
//
// Filter slope 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS12() (interface{}, error) {
	return e.Element.GetProperty("s-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS12(value interface{}) error {
	return e.Element.SetProperty("s-12", value)
}

// s-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos13)
//
// Filter slope 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS13() (interface{}, error) {
	return e.Element.GetProperty("s-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS13(value interface{}) error {
	return e.Element.SetProperty("s-13", value)
}

// s-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos14)
//
// Filter slope 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS14() (interface{}, error) {
	return e.Element.GetProperty("s-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS14(value interface{}) error {
	return e.Element.SetProperty("s-14", value)
}

// s-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos15)
//
// Filter slope 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS15() (interface{}, error) {
	return e.Element.GetProperty("s-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS15(value interface{}) error {
	return e.Element.SetProperty("s-15", value)
}

// s-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos16)
//
// Filter slope 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS16() (interface{}, error) {
	return e.Element.GetProperty("s-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS16(value interface{}) error {
	return e.Element.SetProperty("s-16", value)
}

// s-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos17)
//
// Filter slope 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS17() (interface{}, error) {
	return e.Element.GetProperty("s-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS17(value interface{}) error {
	return e.Element.SetProperty("s-17", value)
}

// s-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos18)
//
// Filter slope 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS18() (interface{}, error) {
	return e.Element.GetProperty("s-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS18(value interface{}) error {
	return e.Element.SetProperty("s-18", value)
}

// s-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos19)
//
// Filter slope 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS19() (interface{}, error) {
	return e.Element.GetProperty("s-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS19(value interface{}) error {
	return e.Element.SetProperty("s-19", value)
}

// s-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos2)
//
// Filter slope 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS2() (interface{}, error) {
	return e.Element.GetProperty("s-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS2(value interface{}) error {
	return e.Element.SetProperty("s-2", value)
}

// s-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos20)
//
// Filter slope 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS20() (interface{}, error) {
	return e.Element.GetProperty("s-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS20(value interface{}) error {
	return e.Element.SetProperty("s-20", value)
}

// s-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos21)
//
// Filter slope 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS21() (interface{}, error) {
	return e.Element.GetProperty("s-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS21(value interface{}) error {
	return e.Element.SetProperty("s-21", value)
}

// s-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos22)
//
// Filter slope 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS22() (interface{}, error) {
	return e.Element.GetProperty("s-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS22(value interface{}) error {
	return e.Element.SetProperty("s-22", value)
}

// s-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos23)
//
// Filter slope 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS23() (interface{}, error) {
	return e.Element.GetProperty("s-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS23(value interface{}) error {
	return e.Element.SetProperty("s-23", value)
}

// s-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos24)
//
// Filter slope 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS24() (interface{}, error) {
	return e.Element.GetProperty("s-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS24(value interface{}) error {
	return e.Element.SetProperty("s-24", value)
}

// s-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos25)
//
// Filter slope 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS25() (interface{}, error) {
	return e.Element.GetProperty("s-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS25(value interface{}) error {
	return e.Element.SetProperty("s-25", value)
}

// s-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos26)
//
// Filter slope 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS26() (interface{}, error) {
	return e.Element.GetProperty("s-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS26(value interface{}) error {
	return e.Element.SetProperty("s-26", value)
}

// s-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos27)
//
// Filter slope 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS27() (interface{}, error) {
	return e.Element.GetProperty("s-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS27(value interface{}) error {
	return e.Element.SetProperty("s-27", value)
}

// s-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos28)
//
// Filter slope 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS28() (interface{}, error) {
	return e.Element.GetProperty("s-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS28(value interface{}) error {
	return e.Element.SetProperty("s-28", value)
}

// s-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos29)
//
// Filter slope 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS29() (interface{}, error) {
	return e.Element.GetProperty("s-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS29(value interface{}) error {
	return e.Element.SetProperty("s-29", value)
}

// s-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos3)
//
// Filter slope 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS3() (interface{}, error) {
	return e.Element.GetProperty("s-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS3(value interface{}) error {
	return e.Element.SetProperty("s-3", value)
}

// s-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos30)
//
// Filter slope 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS30() (interface{}, error) {
	return e.Element.GetProperty("s-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS30(value interface{}) error {
	return e.Element.SetProperty("s-30", value)
}

// s-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos31)
//
// Filter slope 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS31() (interface{}, error) {
	return e.Element.GetProperty("s-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS31(value interface{}) error {
	return e.Element.SetProperty("s-31", value)
}

// s-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos4)
//
// Filter slope 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS4() (interface{}, error) {
	return e.Element.GetProperty("s-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS4(value interface{}) error {
	return e.Element.SetProperty("s-4", value)
}

// s-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos5)
//
// Filter slope 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS5() (interface{}, error) {
	return e.Element.GetProperty("s-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS5(value interface{}) error {
	return e.Element.SetProperty("s-5", value)
}

// s-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos6)
//
// Filter slope 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS6() (interface{}, error) {
	return e.Element.GetProperty("s-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS6(value interface{}) error {
	return e.Element.SetProperty("s-6", value)
}

// s-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos7)
//
// Filter slope 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS7() (interface{}, error) {
	return e.Element.GetProperty("s-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS7(value interface{}) error {
	return e.Element.SetProperty("s-7", value)
}

// s-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos8)
//
// Filter slope 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS8() (interface{}, error) {
	return e.Element.GetProperty("s-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS8(value interface{}) error {
	return e.Element.SetProperty("s-8", value)
}

// s-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Monos9)
//
// Filter slope 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetS9() (interface{}, error) {
	return e.Element.GetProperty("s-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetS9(value interface{}) error {
	return e.Element.SetProperty("s-9", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetShift() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shift")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sm (float32)
//
// Output signal meter

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetSm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// xm-0 (bool)
//
// Filter mute 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm0(value bool) error {
	return e.Element.SetProperty("xm-0", value)
}

// xm-1 (bool)
//
// Filter mute 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm1(value bool) error {
	return e.Element.SetProperty("xm-1", value)
}

// xm-10 (bool)
//
// Filter mute 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm10(value bool) error {
	return e.Element.SetProperty("xm-10", value)
}

// xm-11 (bool)
//
// Filter mute 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm11(value bool) error {
	return e.Element.SetProperty("xm-11", value)
}

// xm-12 (bool)
//
// Filter mute 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm12(value bool) error {
	return e.Element.SetProperty("xm-12", value)
}

// xm-13 (bool)
//
// Filter mute 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm13(value bool) error {
	return e.Element.SetProperty("xm-13", value)
}

// xm-14 (bool)
//
// Filter mute 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm14(value bool) error {
	return e.Element.SetProperty("xm-14", value)
}

// xm-15 (bool)
//
// Filter mute 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm15(value bool) error {
	return e.Element.SetProperty("xm-15", value)
}

// xm-16 (bool)
//
// Filter mute 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm16(value bool) error {
	return e.Element.SetProperty("xm-16", value)
}

// xm-17 (bool)
//
// Filter mute 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm17(value bool) error {
	return e.Element.SetProperty("xm-17", value)
}

// xm-18 (bool)
//
// Filter mute 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm18(value bool) error {
	return e.Element.SetProperty("xm-18", value)
}

// xm-19 (bool)
//
// Filter mute 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm19(value bool) error {
	return e.Element.SetProperty("xm-19", value)
}

// xm-2 (bool)
//
// Filter mute 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm2(value bool) error {
	return e.Element.SetProperty("xm-2", value)
}

// xm-20 (bool)
//
// Filter mute 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm20(value bool) error {
	return e.Element.SetProperty("xm-20", value)
}

// xm-21 (bool)
//
// Filter mute 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm21(value bool) error {
	return e.Element.SetProperty("xm-21", value)
}

// xm-22 (bool)
//
// Filter mute 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm22(value bool) error {
	return e.Element.SetProperty("xm-22", value)
}

// xm-23 (bool)
//
// Filter mute 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm23(value bool) error {
	return e.Element.SetProperty("xm-23", value)
}

// xm-24 (bool)
//
// Filter mute 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm24(value bool) error {
	return e.Element.SetProperty("xm-24", value)
}

// xm-25 (bool)
//
// Filter mute 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm25(value bool) error {
	return e.Element.SetProperty("xm-25", value)
}

// xm-26 (bool)
//
// Filter mute 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm26(value bool) error {
	return e.Element.SetProperty("xm-26", value)
}

// xm-27 (bool)
//
// Filter mute 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm27(value bool) error {
	return e.Element.SetProperty("xm-27", value)
}

// xm-28 (bool)
//
// Filter mute 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm28(value bool) error {
	return e.Element.SetProperty("xm-28", value)
}

// xm-29 (bool)
//
// Filter mute 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm29(value bool) error {
	return e.Element.SetProperty("xm-29", value)
}

// xm-3 (bool)
//
// Filter mute 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm3(value bool) error {
	return e.Element.SetProperty("xm-3", value)
}

// xm-30 (bool)
//
// Filter mute 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm30(value bool) error {
	return e.Element.SetProperty("xm-30", value)
}

// xm-31 (bool)
//
// Filter mute 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm31(value bool) error {
	return e.Element.SetProperty("xm-31", value)
}

// xm-4 (bool)
//
// Filter mute 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm4(value bool) error {
	return e.Element.SetProperty("xm-4", value)
}

// xm-5 (bool)
//
// Filter mute 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm5(value bool) error {
	return e.Element.SetProperty("xm-5", value)
}

// xm-6 (bool)
//
// Filter mute 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm6(value bool) error {
	return e.Element.SetProperty("xm-6", value)
}

// xm-7 (bool)
//
// Filter mute 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm7(value bool) error {
	return e.Element.SetProperty("xm-7", value)
}

// xm-8 (bool)
//
// Filter mute 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm8(value bool) error {
	return e.Element.SetProperty("xm-8", value)
}

// xm-9 (bool)
//
// Filter mute 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXm9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXm9(value bool) error {
	return e.Element.SetProperty("xm-9", value)
}

// xs-0 (bool)
//
// Filter solo 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs0(value bool) error {
	return e.Element.SetProperty("xs-0", value)
}

// xs-1 (bool)
//
// Filter solo 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs1(value bool) error {
	return e.Element.SetProperty("xs-1", value)
}

// xs-10 (bool)
//
// Filter solo 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs10(value bool) error {
	return e.Element.SetProperty("xs-10", value)
}

// xs-11 (bool)
//
// Filter solo 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs11(value bool) error {
	return e.Element.SetProperty("xs-11", value)
}

// xs-12 (bool)
//
// Filter solo 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs12(value bool) error {
	return e.Element.SetProperty("xs-12", value)
}

// xs-13 (bool)
//
// Filter solo 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs13(value bool) error {
	return e.Element.SetProperty("xs-13", value)
}

// xs-14 (bool)
//
// Filter solo 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs14(value bool) error {
	return e.Element.SetProperty("xs-14", value)
}

// xs-15 (bool)
//
// Filter solo 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs15(value bool) error {
	return e.Element.SetProperty("xs-15", value)
}

// xs-16 (bool)
//
// Filter solo 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs16(value bool) error {
	return e.Element.SetProperty("xs-16", value)
}

// xs-17 (bool)
//
// Filter solo 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs17(value bool) error {
	return e.Element.SetProperty("xs-17", value)
}

// xs-18 (bool)
//
// Filter solo 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs18(value bool) error {
	return e.Element.SetProperty("xs-18", value)
}

// xs-19 (bool)
//
// Filter solo 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs19(value bool) error {
	return e.Element.SetProperty("xs-19", value)
}

// xs-2 (bool)
//
// Filter solo 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs2(value bool) error {
	return e.Element.SetProperty("xs-2", value)
}

// xs-20 (bool)
//
// Filter solo 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs20(value bool) error {
	return e.Element.SetProperty("xs-20", value)
}

// xs-21 (bool)
//
// Filter solo 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs21(value bool) error {
	return e.Element.SetProperty("xs-21", value)
}

// xs-22 (bool)
//
// Filter solo 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs22(value bool) error {
	return e.Element.SetProperty("xs-22", value)
}

// xs-23 (bool)
//
// Filter solo 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs23(value bool) error {
	return e.Element.SetProperty("xs-23", value)
}

// xs-24 (bool)
//
// Filter solo 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs24(value bool) error {
	return e.Element.SetProperty("xs-24", value)
}

// xs-25 (bool)
//
// Filter solo 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs25(value bool) error {
	return e.Element.SetProperty("xs-25", value)
}

// xs-26 (bool)
//
// Filter solo 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs26(value bool) error {
	return e.Element.SetProperty("xs-26", value)
}

// xs-27 (bool)
//
// Filter solo 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs27(value bool) error {
	return e.Element.SetProperty("xs-27", value)
}

// xs-28 (bool)
//
// Filter solo 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs28(value bool) error {
	return e.Element.SetProperty("xs-28", value)
}

// xs-29 (bool)
//
// Filter solo 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs29(value bool) error {
	return e.Element.SetProperty("xs-29", value)
}

// xs-3 (bool)
//
// Filter solo 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs3(value bool) error {
	return e.Element.SetProperty("xs-3", value)
}

// xs-30 (bool)
//
// Filter solo 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs30(value bool) error {
	return e.Element.SetProperty("xs-30", value)
}

// xs-31 (bool)
//
// Filter solo 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs31(value bool) error {
	return e.Element.SetProperty("xs-31", value)
}

// xs-4 (bool)
//
// Filter solo 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs4(value bool) error {
	return e.Element.SetProperty("xs-4", value)
}

// xs-5 (bool)
//
// Filter solo 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs5(value bool) error {
	return e.Element.SetProperty("xs-5", value)
}

// xs-6 (bool)
//
// Filter solo 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs6(value bool) error {
	return e.Element.SetProperty("xs-6", value)
}

// xs-7 (bool)
//
// Filter solo 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs7(value bool) error {
	return e.Element.SetProperty("xs-7", value)
}

// xs-8 (bool)
//
// Filter solo 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs8(value bool) error {
	return e.Element.SetProperty("xs-8", value)
}

// xs-9 (bool)
//
// Filter solo 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetXs9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetXs9(value bool) error {
	return e.Element.SetProperty("xs-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) GetZoom() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Mono) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX32Monoft13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13Off LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft13Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm23ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm23 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm24ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm24 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0Off LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft0Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9Off LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft9Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos29X1 LspPlugInPluginsLv2ParaEqualizerX32Monos29 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos29X2 LspPlugInPluginsLv2ParaEqualizerX32Monos29 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos29X3 LspPlugInPluginsLv2ParaEqualizerX32Monos29 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos29X4 LspPlugInPluginsLv2ParaEqualizerX32Monos29 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos4X1 LspPlugInPluginsLv2ParaEqualizerX32Monos4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos4X2 LspPlugInPluginsLv2ParaEqualizerX32Monos4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos4X3 LspPlugInPluginsLv2ParaEqualizerX32Monos4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos4X4 LspPlugInPluginsLv2ParaEqualizerX32Monos4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos6X1 LspPlugInPluginsLv2ParaEqualizerX32Monos6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos6X2 LspPlugInPluginsLv2ParaEqualizerX32Monos6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos6X3 LspPlugInPluginsLv2ParaEqualizerX32Monos6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos6X4 LspPlugInPluginsLv2ParaEqualizerX32Monos6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm1ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm19ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm19 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm22ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm22 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6Off LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft6Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos15X1 LspPlugInPluginsLv2ParaEqualizerX32Monos15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos15X2 LspPlugInPluginsLv2ParaEqualizerX32Monos15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos15X3 LspPlugInPluginsLv2ParaEqualizerX32Monos15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos15X4 LspPlugInPluginsLv2ParaEqualizerX32Monos15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos16X1 LspPlugInPluginsLv2ParaEqualizerX32Monos16 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos16X2 LspPlugInPluginsLv2ParaEqualizerX32Monos16 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos16X3 LspPlugInPluginsLv2ParaEqualizerX32Monos16 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos16X4 LspPlugInPluginsLv2ParaEqualizerX32Monos16 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm8ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21Off LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft21Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft21 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26Off LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft26Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft26 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm20ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm20 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm21ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm21 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm29ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm29 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm4ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14Off LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft14Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18Off LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft18Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft18 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos23X1 LspPlugInPluginsLv2ParaEqualizerX32Monos23 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos23X2 LspPlugInPluginsLv2ParaEqualizerX32Monos23 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos23X3 LspPlugInPluginsLv2ParaEqualizerX32Monos23 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos23X4 LspPlugInPluginsLv2ParaEqualizerX32Monos23 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos24X1 LspPlugInPluginsLv2ParaEqualizerX32Monos24 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos24X2 LspPlugInPluginsLv2ParaEqualizerX32Monos24 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos24X3 LspPlugInPluginsLv2ParaEqualizerX32Monos24 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos24X4 LspPlugInPluginsLv2ParaEqualizerX32Monos24 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm14ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm5ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX32MonofselFilters07 LspPlugInPluginsLv2ParaEqualizerX32Monofsel = "Filters 0-7" // Filters 0-7 (0) – Filters 0-7
	LspPlugInPluginsLv2ParaEqualizerX32MonofselFilters815 LspPlugInPluginsLv2ParaEqualizerX32Monofsel = "Filters 8-15" // Filters 8-15 (1) – Filters 8-15
	LspPlugInPluginsLv2ParaEqualizerX32MonofselFilters1623 LspPlugInPluginsLv2ParaEqualizerX32Monofsel = "Filters 16-23" // Filters 16-23 (2) – Filters 16-23
	LspPlugInPluginsLv2ParaEqualizerX32MonofselFilters2431 LspPlugInPluginsLv2ParaEqualizerX32Monofsel = "Filters 24-31" // Filters 24-31 (3) – Filters 24-31
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19Off LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft19Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft19 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos26X1 LspPlugInPluginsLv2ParaEqualizerX32Monos26 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos26X2 LspPlugInPluginsLv2ParaEqualizerX32Monos26 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos26X3 LspPlugInPluginsLv2ParaEqualizerX32Monos26 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos26X4 LspPlugInPluginsLv2ParaEqualizerX32Monos26 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos7X1 LspPlugInPluginsLv2ParaEqualizerX32Monos7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos7X2 LspPlugInPluginsLv2ParaEqualizerX32Monos7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos7X3 LspPlugInPluginsLv2ParaEqualizerX32Monos7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos7X4 LspPlugInPluginsLv2ParaEqualizerX32Monos7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12Off LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft12Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4Off LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft4Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7Off LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft7Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos1X1 LspPlugInPluginsLv2ParaEqualizerX32Monos1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos1X2 LspPlugInPluginsLv2ParaEqualizerX32Monos1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos1X3 LspPlugInPluginsLv2ParaEqualizerX32Monos1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos1X4 LspPlugInPluginsLv2ParaEqualizerX32Monos1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos28X1 LspPlugInPluginsLv2ParaEqualizerX32Monos28 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos28X2 LspPlugInPluginsLv2ParaEqualizerX32Monos28 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos28X3 LspPlugInPluginsLv2ParaEqualizerX32Monos28 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos28X4 LspPlugInPluginsLv2ParaEqualizerX32Monos28 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11Off LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft11Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20Off LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft20Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft20 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm0ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm10ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm25ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm25 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm26ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm26 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm28ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm28 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm31ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm31 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22Off LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft22Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft22 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31Off LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft31Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft31 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos19X1 LspPlugInPluginsLv2ParaEqualizerX32Monos19 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos19X2 LspPlugInPluginsLv2ParaEqualizerX32Monos19 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos19X3 LspPlugInPluginsLv2ParaEqualizerX32Monos19 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos19X4 LspPlugInPluginsLv2ParaEqualizerX32Monos19 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos21X1 LspPlugInPluginsLv2ParaEqualizerX32Monos21 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos21X2 LspPlugInPluginsLv2ParaEqualizerX32Monos21 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos21X3 LspPlugInPluginsLv2ParaEqualizerX32Monos21 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos21X4 LspPlugInPluginsLv2ParaEqualizerX32Monos21 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos30X1 LspPlugInPluginsLv2ParaEqualizerX32Monos30 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos30X2 LspPlugInPluginsLv2ParaEqualizerX32Monos30 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos30X3 LspPlugInPluginsLv2ParaEqualizerX32Monos30 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos30X4 LspPlugInPluginsLv2ParaEqualizerX32Monos30 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm16ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm16 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos2X1 LspPlugInPluginsLv2ParaEqualizerX32Monos2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos2X2 LspPlugInPluginsLv2ParaEqualizerX32Monos2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos2X3 LspPlugInPluginsLv2ParaEqualizerX32Monos2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos2X4 LspPlugInPluginsLv2ParaEqualizerX32Monos2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos22X1 LspPlugInPluginsLv2ParaEqualizerX32Monos22 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos22X2 LspPlugInPluginsLv2ParaEqualizerX32Monos22 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos22X3 LspPlugInPluginsLv2ParaEqualizerX32Monos22 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos22X4 LspPlugInPluginsLv2ParaEqualizerX32Monos22 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm11ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm18ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm18 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm7ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27Off LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft27Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft27 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30Off LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft30Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft30 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8Off LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft8Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos5X1 LspPlugInPluginsLv2ParaEqualizerX32Monos5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos5X2 LspPlugInPluginsLv2ParaEqualizerX32Monos5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos5X3 LspPlugInPluginsLv2ParaEqualizerX32Monos5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos5X4 LspPlugInPluginsLv2ParaEqualizerX32Monos5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofft string

const (
	LspPlugInPluginsLv2ParaEqualizerX32MonofftOff LspPlugInPluginsLv2ParaEqualizerX32Monofft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32MonofftPostEq LspPlugInPluginsLv2ParaEqualizerX32Monofft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX32MonofftPreEq LspPlugInPluginsLv2ParaEqualizerX32Monofft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm27ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm27 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17Off LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft17Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft17 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos18X1 LspPlugInPluginsLv2ParaEqualizerX32Monos18 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos18X2 LspPlugInPluginsLv2ParaEqualizerX32Monos18 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos18X3 LspPlugInPluginsLv2ParaEqualizerX32Monos18 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos18X4 LspPlugInPluginsLv2ParaEqualizerX32Monos18 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos20X1 LspPlugInPluginsLv2ParaEqualizerX32Monos20 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos20X2 LspPlugInPluginsLv2ParaEqualizerX32Monos20 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos20X3 LspPlugInPluginsLv2ParaEqualizerX32Monos20 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos20X4 LspPlugInPluginsLv2ParaEqualizerX32Monos20 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos3X1 LspPlugInPluginsLv2ParaEqualizerX32Monos3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos3X2 LspPlugInPluginsLv2ParaEqualizerX32Monos3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos3X3 LspPlugInPluginsLv2ParaEqualizerX32Monos3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos3X4 LspPlugInPluginsLv2ParaEqualizerX32Monos3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5Off LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft5Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm12ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm15ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm9ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23Off LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft23Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft23 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24Off LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft24Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft24 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3Off LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft3Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm17ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm17 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25Off LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft25Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft25 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos10X1 LspPlugInPluginsLv2ParaEqualizerX32Monos10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos10X2 LspPlugInPluginsLv2ParaEqualizerX32Monos10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos10X3 LspPlugInPluginsLv2ParaEqualizerX32Monos10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos10X4 LspPlugInPluginsLv2ParaEqualizerX32Monos10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos13X1 LspPlugInPluginsLv2ParaEqualizerX32Monos13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos13X2 LspPlugInPluginsLv2ParaEqualizerX32Monos13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos13X3 LspPlugInPluginsLv2ParaEqualizerX32Monos13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos13X4 LspPlugInPluginsLv2ParaEqualizerX32Monos13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos9X1 LspPlugInPluginsLv2ParaEqualizerX32Monos9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos9X2 LspPlugInPluginsLv2ParaEqualizerX32Monos9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos9X3 LspPlugInPluginsLv2ParaEqualizerX32Monos9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos9X4 LspPlugInPluginsLv2ParaEqualizerX32Monos9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16Off LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft16Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft16 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2Off LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft2Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos0X1 LspPlugInPluginsLv2ParaEqualizerX32Monos0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos0X2 LspPlugInPluginsLv2ParaEqualizerX32Monos0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos0X3 LspPlugInPluginsLv2ParaEqualizerX32Monos0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos0X4 LspPlugInPluginsLv2ParaEqualizerX32Monos0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos17X1 LspPlugInPluginsLv2ParaEqualizerX32Monos17 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos17X2 LspPlugInPluginsLv2ParaEqualizerX32Monos17 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos17X3 LspPlugInPluginsLv2ParaEqualizerX32Monos17 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos17X4 LspPlugInPluginsLv2ParaEqualizerX32Monos17 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm3ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm30ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm30 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm6ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29Off LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft29Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft29 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos25X1 LspPlugInPluginsLv2ParaEqualizerX32Monos25 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos25X2 LspPlugInPluginsLv2ParaEqualizerX32Monos25 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos25X3 LspPlugInPluginsLv2ParaEqualizerX32Monos25 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos25X4 LspPlugInPluginsLv2ParaEqualizerX32Monos25 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos31X1 LspPlugInPluginsLv2ParaEqualizerX32Monos31 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos31X2 LspPlugInPluginsLv2ParaEqualizerX32Monos31 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos31X3 LspPlugInPluginsLv2ParaEqualizerX32Monos31 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos31X4 LspPlugInPluginsLv2ParaEqualizerX32Monos31 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monomode string

const (
	LspPlugInPluginsLv2ParaEqualizerX32MonomodeIir LspPlugInPluginsLv2ParaEqualizerX32Monomode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX32MonomodeFir LspPlugInPluginsLv2ParaEqualizerX32Monomode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX32MonomodeFft LspPlugInPluginsLv2ParaEqualizerX32Monomode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos11X1 LspPlugInPluginsLv2ParaEqualizerX32Monos11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos11X2 LspPlugInPluginsLv2ParaEqualizerX32Monos11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos11X3 LspPlugInPluginsLv2ParaEqualizerX32Monos11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos11X4 LspPlugInPluginsLv2ParaEqualizerX32Monos11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm13ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monofm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2RlcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2RlcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2BwcBt LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2BwcMt LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2LrxBt LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2LrxMt LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Monofm2ApoDr LspPlugInPluginsLv2ParaEqualizerX32Monofm2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1Off LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft1Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10Off LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft10Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15Off LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft15Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monoft28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28Off LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28Bell LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28HiPass LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28HiShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28LoPass LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28LoShelf LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28Notch LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28Resonance LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Monoft28Allpass LspPlugInPluginsLv2ParaEqualizerX32Monoft28 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos12X1 LspPlugInPluginsLv2ParaEqualizerX32Monos12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos12X2 LspPlugInPluginsLv2ParaEqualizerX32Monos12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos12X3 LspPlugInPluginsLv2ParaEqualizerX32Monos12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos12X4 LspPlugInPluginsLv2ParaEqualizerX32Monos12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos14X1 LspPlugInPluginsLv2ParaEqualizerX32Monos14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos14X2 LspPlugInPluginsLv2ParaEqualizerX32Monos14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos14X3 LspPlugInPluginsLv2ParaEqualizerX32Monos14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos14X4 LspPlugInPluginsLv2ParaEqualizerX32Monos14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos27X1 LspPlugInPluginsLv2ParaEqualizerX32Monos27 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos27X2 LspPlugInPluginsLv2ParaEqualizerX32Monos27 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos27X3 LspPlugInPluginsLv2ParaEqualizerX32Monos27 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos27X4 LspPlugInPluginsLv2ParaEqualizerX32Monos27 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Monos8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Monos8X1 LspPlugInPluginsLv2ParaEqualizerX32Monos8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Monos8X2 LspPlugInPluginsLv2ParaEqualizerX32Monos8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Monos8X3 LspPlugInPluginsLv2ParaEqualizerX32Monos8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Monos8X4 LspPlugInPluginsLv2ParaEqualizerX32Monos8 = "x4" // x4 (3) – x4
)

