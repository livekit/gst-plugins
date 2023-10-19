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

type LspPlugInPluginsLv2ParaEqualizerX16Mono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX16Mono() (*LspPlugInPluginsLv2ParaEqualizerX16Mono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x16-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Mono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX16MonoWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX16Mono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x16-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Mono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// f-0 (float32)
//
// Frequency 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF0(value float32) error {
	return e.Element.SetProperty("f-0", value)
}

// f-1 (float32)
//
// Frequency 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF1(value float32) error {
	return e.Element.SetProperty("f-1", value)
}

// f-10 (float32)
//
// Frequency 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF10(value float32) error {
	return e.Element.SetProperty("f-10", value)
}

// f-11 (float32)
//
// Frequency 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF11(value float32) error {
	return e.Element.SetProperty("f-11", value)
}

// f-12 (float32)
//
// Frequency 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF12(value float32) error {
	return e.Element.SetProperty("f-12", value)
}

// f-13 (float32)
//
// Frequency 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF13(value float32) error {
	return e.Element.SetProperty("f-13", value)
}

// f-14 (float32)
//
// Frequency 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF14(value float32) error {
	return e.Element.SetProperty("f-14", value)
}

// f-15 (float32)
//
// Frequency 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF15(value float32) error {
	return e.Element.SetProperty("f-15", value)
}

// f-2 (float32)
//
// Frequency 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF2(value float32) error {
	return e.Element.SetProperty("f-2", value)
}

// f-3 (float32)
//
// Frequency 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF3(value float32) error {
	return e.Element.SetProperty("f-3", value)
}

// f-4 (float32)
//
// Frequency 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF4(value float32) error {
	return e.Element.SetProperty("f-4", value)
}

// f-5 (float32)
//
// Frequency 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF5(value float32) error {
	return e.Element.SetProperty("f-5", value)
}

// f-6 (float32)
//
// Frequency 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF6(value float32) error {
	return e.Element.SetProperty("f-6", value)
}

// f-7 (float32)
//
// Frequency 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF7(value float32) error {
	return e.Element.SetProperty("f-7", value)
}

// f-8 (float32)
//
// Frequency 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF8(value float32) error {
	return e.Element.SetProperty("f-8", value)
}

// f-9 (float32)
//
// Frequency 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetF9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetF9(value float32) error {
	return e.Element.SetProperty("f-9", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX16Monofft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fm-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm0)
//
// Filter mode 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm0() (interface{}, error) {
	return e.Element.GetProperty("fm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm0(value interface{}) error {
	return e.Element.SetProperty("fm-0", value)
}

// fm-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm1)
//
// Filter mode 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm1() (interface{}, error) {
	return e.Element.GetProperty("fm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm1(value interface{}) error {
	return e.Element.SetProperty("fm-1", value)
}

// fm-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm10)
//
// Filter mode 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm10() (interface{}, error) {
	return e.Element.GetProperty("fm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm10(value interface{}) error {
	return e.Element.SetProperty("fm-10", value)
}

// fm-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm11)
//
// Filter mode 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm11() (interface{}, error) {
	return e.Element.GetProperty("fm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm11(value interface{}) error {
	return e.Element.SetProperty("fm-11", value)
}

// fm-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm12)
//
// Filter mode 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm12() (interface{}, error) {
	return e.Element.GetProperty("fm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm12(value interface{}) error {
	return e.Element.SetProperty("fm-12", value)
}

// fm-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm13)
//
// Filter mode 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm13() (interface{}, error) {
	return e.Element.GetProperty("fm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm13(value interface{}) error {
	return e.Element.SetProperty("fm-13", value)
}

// fm-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm14)
//
// Filter mode 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm14() (interface{}, error) {
	return e.Element.GetProperty("fm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm14(value interface{}) error {
	return e.Element.SetProperty("fm-14", value)
}

// fm-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm15)
//
// Filter mode 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm15() (interface{}, error) {
	return e.Element.GetProperty("fm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm15(value interface{}) error {
	return e.Element.SetProperty("fm-15", value)
}

// fm-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm2)
//
// Filter mode 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm2() (interface{}, error) {
	return e.Element.GetProperty("fm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm2(value interface{}) error {
	return e.Element.SetProperty("fm-2", value)
}

// fm-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm3)
//
// Filter mode 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm3() (interface{}, error) {
	return e.Element.GetProperty("fm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm3(value interface{}) error {
	return e.Element.SetProperty("fm-3", value)
}

// fm-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm4)
//
// Filter mode 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm4() (interface{}, error) {
	return e.Element.GetProperty("fm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm4(value interface{}) error {
	return e.Element.SetProperty("fm-4", value)
}

// fm-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm5)
//
// Filter mode 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm5() (interface{}, error) {
	return e.Element.GetProperty("fm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm5(value interface{}) error {
	return e.Element.SetProperty("fm-5", value)
}

// fm-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm6)
//
// Filter mode 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm6() (interface{}, error) {
	return e.Element.GetProperty("fm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm6(value interface{}) error {
	return e.Element.SetProperty("fm-6", value)
}

// fm-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm7)
//
// Filter mode 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm7() (interface{}, error) {
	return e.Element.GetProperty("fm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm7(value interface{}) error {
	return e.Element.SetProperty("fm-7", value)
}

// fm-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm8)
//
// Filter mode 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm8() (interface{}, error) {
	return e.Element.GetProperty("fm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm8(value interface{}) error {
	return e.Element.SetProperty("fm-8", value)
}

// fm-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Monofm9)
//
// Filter mode 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFm9() (interface{}, error) {
	return e.Element.GetProperty("fm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFm9(value interface{}) error {
	return e.Element.SetProperty("fm-9", value)
}

// frqs (float32)
//
// Frequency shift

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFrqs() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFrqs(value float32) error {
	return e.Element.SetProperty("frqs", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX16Monofsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ft-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft0)
//
// Filter type 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt0() (interface{}, error) {
	return e.Element.GetProperty("ft-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt0(value interface{}) error {
	return e.Element.SetProperty("ft-0", value)
}

// ft-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft1)
//
// Filter type 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt1() (interface{}, error) {
	return e.Element.GetProperty("ft-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt1(value interface{}) error {
	return e.Element.SetProperty("ft-1", value)
}

// ft-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft10)
//
// Filter type 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt10() (interface{}, error) {
	return e.Element.GetProperty("ft-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt10(value interface{}) error {
	return e.Element.SetProperty("ft-10", value)
}

// ft-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft11)
//
// Filter type 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt11() (interface{}, error) {
	return e.Element.GetProperty("ft-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt11(value interface{}) error {
	return e.Element.SetProperty("ft-11", value)
}

// ft-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft12)
//
// Filter type 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt12() (interface{}, error) {
	return e.Element.GetProperty("ft-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt12(value interface{}) error {
	return e.Element.SetProperty("ft-12", value)
}

// ft-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft13)
//
// Filter type 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt13() (interface{}, error) {
	return e.Element.GetProperty("ft-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt13(value interface{}) error {
	return e.Element.SetProperty("ft-13", value)
}

// ft-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft14)
//
// Filter type 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt14() (interface{}, error) {
	return e.Element.GetProperty("ft-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt14(value interface{}) error {
	return e.Element.SetProperty("ft-14", value)
}

// ft-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft15)
//
// Filter type 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt15() (interface{}, error) {
	return e.Element.GetProperty("ft-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt15(value interface{}) error {
	return e.Element.SetProperty("ft-15", value)
}

// ft-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft2)
//
// Filter type 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt2() (interface{}, error) {
	return e.Element.GetProperty("ft-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt2(value interface{}) error {
	return e.Element.SetProperty("ft-2", value)
}

// ft-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft3)
//
// Filter type 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt3() (interface{}, error) {
	return e.Element.GetProperty("ft-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt3(value interface{}) error {
	return e.Element.SetProperty("ft-3", value)
}

// ft-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft4)
//
// Filter type 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt4() (interface{}, error) {
	return e.Element.GetProperty("ft-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt4(value interface{}) error {
	return e.Element.SetProperty("ft-4", value)
}

// ft-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft5)
//
// Filter type 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt5() (interface{}, error) {
	return e.Element.GetProperty("ft-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt5(value interface{}) error {
	return e.Element.SetProperty("ft-5", value)
}

// ft-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft6)
//
// Filter type 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt6() (interface{}, error) {
	return e.Element.GetProperty("ft-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt6(value interface{}) error {
	return e.Element.SetProperty("ft-6", value)
}

// ft-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft7)
//
// Filter type 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt7() (interface{}, error) {
	return e.Element.GetProperty("ft-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt7(value interface{}) error {
	return e.Element.SetProperty("ft-7", value)
}

// ft-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft8)
//
// Filter type 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt8() (interface{}, error) {
	return e.Element.GetProperty("ft-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt8(value interface{}) error {
	return e.Element.SetProperty("ft-8", value)
}

// ft-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Monoft9)
//
// Filter type 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFt9() (interface{}, error) {
	return e.Element.GetProperty("ft-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetFt9(value interface{}) error {
	return e.Element.SetProperty("ft-9", value)
}

// fv-0 (bool)
//
// Filter visibility 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv15() (bool, error) {
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

// fv-2 (bool)
//
// Filter visibility 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv2() (bool, error) {
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

// fv-3 (bool)
//
// Filter visibility 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv3() (bool, error) {
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

// fv-4 (bool)
//
// Filter visibility 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetFv9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG0(value float32) error {
	return e.Element.SetProperty("g-0", value)
}

// g-1 (float32)
//
// Gain 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG1(value float32) error {
	return e.Element.SetProperty("g-1", value)
}

// g-10 (float32)
//
// Gain 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG10(value float32) error {
	return e.Element.SetProperty("g-10", value)
}

// g-11 (float32)
//
// Gain 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG11(value float32) error {
	return e.Element.SetProperty("g-11", value)
}

// g-12 (float32)
//
// Gain 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG12(value float32) error {
	return e.Element.SetProperty("g-12", value)
}

// g-13 (float32)
//
// Gain 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG13(value float32) error {
	return e.Element.SetProperty("g-13", value)
}

// g-14 (float32)
//
// Gain 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG14(value float32) error {
	return e.Element.SetProperty("g-14", value)
}

// g-15 (float32)
//
// Gain 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG15(value float32) error {
	return e.Element.SetProperty("g-15", value)
}

// g-2 (float32)
//
// Gain 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG2(value float32) error {
	return e.Element.SetProperty("g-2", value)
}

// g-3 (float32)
//
// Gain 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG3(value float32) error {
	return e.Element.SetProperty("g-3", value)
}

// g-4 (float32)
//
// Gain 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG4(value float32) error {
	return e.Element.SetProperty("g-4", value)
}

// g-5 (float32)
//
// Gain 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG5(value float32) error {
	return e.Element.SetProperty("g-5", value)
}

// g-6 (float32)
//
// Gain 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG6(value float32) error {
	return e.Element.SetProperty("g-6", value)
}

// g-7 (float32)
//
// Gain 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG7(value float32) error {
	return e.Element.SetProperty("g-7", value)
}

// g-8 (float32)
//
// Gain 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG8(value float32) error {
	return e.Element.SetProperty("g-8", value)
}

// g-9 (float32)
//
// Gain 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetG9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetG9(value float32) error {
	return e.Element.SetProperty("g-9", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hue-0 (float32)
//
// Hue 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// hue-1 (float32)
//
// Hue 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue1(value float32) error {
	return e.Element.SetProperty("hue-1", value)
}

// hue-10 (float32)
//
// Hue 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue10(value float32) error {
	return e.Element.SetProperty("hue-10", value)
}

// hue-11 (float32)
//
// Hue 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue11(value float32) error {
	return e.Element.SetProperty("hue-11", value)
}

// hue-12 (float32)
//
// Hue 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue12(value float32) error {
	return e.Element.SetProperty("hue-12", value)
}

// hue-13 (float32)
//
// Hue 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue13(value float32) error {
	return e.Element.SetProperty("hue-13", value)
}

// hue-14 (float32)
//
// Hue 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue14(value float32) error {
	return e.Element.SetProperty("hue-14", value)
}

// hue-15 (float32)
//
// Hue 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue15(value float32) error {
	return e.Element.SetProperty("hue-15", value)
}

// hue-2 (float32)
//
// Hue 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue2(value float32) error {
	return e.Element.SetProperty("hue-2", value)
}

// hue-3 (float32)
//
// Hue 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue3(value float32) error {
	return e.Element.SetProperty("hue-3", value)
}

// hue-4 (float32)
//
// Hue 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue4(value float32) error {
	return e.Element.SetProperty("hue-4", value)
}

// hue-5 (float32)
//
// Hue 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue5(value float32) error {
	return e.Element.SetProperty("hue-5", value)
}

// hue-6 (float32)
//
// Hue 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue6(value float32) error {
	return e.Element.SetProperty("hue-6", value)
}

// hue-7 (float32)
//
// Hue 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue7(value float32) error {
	return e.Element.SetProperty("hue-7", value)
}

// hue-8 (float32)
//
// Hue 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue8(value float32) error {
	return e.Element.SetProperty("hue-8", value)
}

// hue-9 (float32)
//
// Hue 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetHue9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetHue9(value float32) error {
	return e.Element.SetProperty("hue-9", value)
}

// im (float32)
//
// Input signal meter

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetIm() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2ParaEqualizerX16Monomode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ0(value float32) error {
	return e.Element.SetProperty("q-0", value)
}

// q-1 (float32)
//
// Quality factor 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ1(value float32) error {
	return e.Element.SetProperty("q-1", value)
}

// q-10 (float32)
//
// Quality factor 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ10(value float32) error {
	return e.Element.SetProperty("q-10", value)
}

// q-11 (float32)
//
// Quality factor 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ11(value float32) error {
	return e.Element.SetProperty("q-11", value)
}

// q-12 (float32)
//
// Quality factor 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ12(value float32) error {
	return e.Element.SetProperty("q-12", value)
}

// q-13 (float32)
//
// Quality factor 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ13(value float32) error {
	return e.Element.SetProperty("q-13", value)
}

// q-14 (float32)
//
// Quality factor 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ14(value float32) error {
	return e.Element.SetProperty("q-14", value)
}

// q-15 (float32)
//
// Quality factor 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ15(value float32) error {
	return e.Element.SetProperty("q-15", value)
}

// q-2 (float32)
//
// Quality factor 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ2(value float32) error {
	return e.Element.SetProperty("q-2", value)
}

// q-3 (float32)
//
// Quality factor 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ3(value float32) error {
	return e.Element.SetProperty("q-3", value)
}

// q-4 (float32)
//
// Quality factor 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ4(value float32) error {
	return e.Element.SetProperty("q-4", value)
}

// q-5 (float32)
//
// Quality factor 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ5(value float32) error {
	return e.Element.SetProperty("q-5", value)
}

// q-6 (float32)
//
// Quality factor 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ6(value float32) error {
	return e.Element.SetProperty("q-6", value)
}

// q-7 (float32)
//
// Quality factor 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ7(value float32) error {
	return e.Element.SetProperty("q-7", value)
}

// q-8 (float32)
//
// Quality factor 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ8(value float32) error {
	return e.Element.SetProperty("q-8", value)
}

// q-9 (float32)
//
// Quality factor 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetQ9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetQ9(value float32) error {
	return e.Element.SetProperty("q-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// s-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos0)
//
// Filter slope 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS0() (interface{}, error) {
	return e.Element.GetProperty("s-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS0(value interface{}) error {
	return e.Element.SetProperty("s-0", value)
}

// s-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos1)
//
// Filter slope 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS1() (interface{}, error) {
	return e.Element.GetProperty("s-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS1(value interface{}) error {
	return e.Element.SetProperty("s-1", value)
}

// s-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos10)
//
// Filter slope 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS10() (interface{}, error) {
	return e.Element.GetProperty("s-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS10(value interface{}) error {
	return e.Element.SetProperty("s-10", value)
}

// s-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos11)
//
// Filter slope 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS11() (interface{}, error) {
	return e.Element.GetProperty("s-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS11(value interface{}) error {
	return e.Element.SetProperty("s-11", value)
}

// s-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos12)
//
// Filter slope 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS12() (interface{}, error) {
	return e.Element.GetProperty("s-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS12(value interface{}) error {
	return e.Element.SetProperty("s-12", value)
}

// s-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos13)
//
// Filter slope 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS13() (interface{}, error) {
	return e.Element.GetProperty("s-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS13(value interface{}) error {
	return e.Element.SetProperty("s-13", value)
}

// s-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos14)
//
// Filter slope 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS14() (interface{}, error) {
	return e.Element.GetProperty("s-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS14(value interface{}) error {
	return e.Element.SetProperty("s-14", value)
}

// s-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos15)
//
// Filter slope 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS15() (interface{}, error) {
	return e.Element.GetProperty("s-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS15(value interface{}) error {
	return e.Element.SetProperty("s-15", value)
}

// s-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos2)
//
// Filter slope 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS2() (interface{}, error) {
	return e.Element.GetProperty("s-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS2(value interface{}) error {
	return e.Element.SetProperty("s-2", value)
}

// s-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos3)
//
// Filter slope 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS3() (interface{}, error) {
	return e.Element.GetProperty("s-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS3(value interface{}) error {
	return e.Element.SetProperty("s-3", value)
}

// s-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos4)
//
// Filter slope 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS4() (interface{}, error) {
	return e.Element.GetProperty("s-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS4(value interface{}) error {
	return e.Element.SetProperty("s-4", value)
}

// s-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos5)
//
// Filter slope 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS5() (interface{}, error) {
	return e.Element.GetProperty("s-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS5(value interface{}) error {
	return e.Element.SetProperty("s-5", value)
}

// s-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos6)
//
// Filter slope 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS6() (interface{}, error) {
	return e.Element.GetProperty("s-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS6(value interface{}) error {
	return e.Element.SetProperty("s-6", value)
}

// s-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos7)
//
// Filter slope 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS7() (interface{}, error) {
	return e.Element.GetProperty("s-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS7(value interface{}) error {
	return e.Element.SetProperty("s-7", value)
}

// s-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos8)
//
// Filter slope 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS8() (interface{}, error) {
	return e.Element.GetProperty("s-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS8(value interface{}) error {
	return e.Element.SetProperty("s-8", value)
}

// s-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Monos9)
//
// Filter slope 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetS9() (interface{}, error) {
	return e.Element.GetProperty("s-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetS9(value interface{}) error {
	return e.Element.SetProperty("s-9", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sm (float32)
//
// Output signal meter

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetSm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm0(value bool) error {
	return e.Element.SetProperty("xm-0", value)
}

// xm-1 (bool)
//
// Filter mute 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm1(value bool) error {
	return e.Element.SetProperty("xm-1", value)
}

// xm-10 (bool)
//
// Filter mute 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm10(value bool) error {
	return e.Element.SetProperty("xm-10", value)
}

// xm-11 (bool)
//
// Filter mute 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm11(value bool) error {
	return e.Element.SetProperty("xm-11", value)
}

// xm-12 (bool)
//
// Filter mute 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm12(value bool) error {
	return e.Element.SetProperty("xm-12", value)
}

// xm-13 (bool)
//
// Filter mute 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm13(value bool) error {
	return e.Element.SetProperty("xm-13", value)
}

// xm-14 (bool)
//
// Filter mute 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm14(value bool) error {
	return e.Element.SetProperty("xm-14", value)
}

// xm-15 (bool)
//
// Filter mute 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm15(value bool) error {
	return e.Element.SetProperty("xm-15", value)
}

// xm-2 (bool)
//
// Filter mute 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm2(value bool) error {
	return e.Element.SetProperty("xm-2", value)
}

// xm-3 (bool)
//
// Filter mute 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm3(value bool) error {
	return e.Element.SetProperty("xm-3", value)
}

// xm-4 (bool)
//
// Filter mute 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm4(value bool) error {
	return e.Element.SetProperty("xm-4", value)
}

// xm-5 (bool)
//
// Filter mute 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm5(value bool) error {
	return e.Element.SetProperty("xm-5", value)
}

// xm-6 (bool)
//
// Filter mute 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm6(value bool) error {
	return e.Element.SetProperty("xm-6", value)
}

// xm-7 (bool)
//
// Filter mute 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm7(value bool) error {
	return e.Element.SetProperty("xm-7", value)
}

// xm-8 (bool)
//
// Filter mute 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm8(value bool) error {
	return e.Element.SetProperty("xm-8", value)
}

// xm-9 (bool)
//
// Filter mute 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXm9(value bool) error {
	return e.Element.SetProperty("xm-9", value)
}

// xs-0 (bool)
//
// Filter solo 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs0(value bool) error {
	return e.Element.SetProperty("xs-0", value)
}

// xs-1 (bool)
//
// Filter solo 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs1(value bool) error {
	return e.Element.SetProperty("xs-1", value)
}

// xs-10 (bool)
//
// Filter solo 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs10(value bool) error {
	return e.Element.SetProperty("xs-10", value)
}

// xs-11 (bool)
//
// Filter solo 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs11(value bool) error {
	return e.Element.SetProperty("xs-11", value)
}

// xs-12 (bool)
//
// Filter solo 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs12(value bool) error {
	return e.Element.SetProperty("xs-12", value)
}

// xs-13 (bool)
//
// Filter solo 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs13(value bool) error {
	return e.Element.SetProperty("xs-13", value)
}

// xs-14 (bool)
//
// Filter solo 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs14(value bool) error {
	return e.Element.SetProperty("xs-14", value)
}

// xs-15 (bool)
//
// Filter solo 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs15(value bool) error {
	return e.Element.SetProperty("xs-15", value)
}

// xs-2 (bool)
//
// Filter solo 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs2(value bool) error {
	return e.Element.SetProperty("xs-2", value)
}

// xs-3 (bool)
//
// Filter solo 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs3(value bool) error {
	return e.Element.SetProperty("xs-3", value)
}

// xs-4 (bool)
//
// Filter solo 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs4(value bool) error {
	return e.Element.SetProperty("xs-4", value)
}

// xs-5 (bool)
//
// Filter solo 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs5(value bool) error {
	return e.Element.SetProperty("xs-5", value)
}

// xs-6 (bool)
//
// Filter solo 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs6(value bool) error {
	return e.Element.SetProperty("xs-6", value)
}

// xs-7 (bool)
//
// Filter solo 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs7(value bool) error {
	return e.Element.SetProperty("xs-7", value)
}

// xs-8 (bool)
//
// Filter solo 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs8(value bool) error {
	return e.Element.SetProperty("xs-8", value)
}

// xs-9 (bool)
//
// Filter solo 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetXs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetXs9(value bool) error {
	return e.Element.SetProperty("xs-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Mono) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX16Monofm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm1ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm7ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10Off LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft10Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12Off LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft12Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13Off LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft13Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos8X1 LspPlugInPluginsLv2ParaEqualizerX16Monos8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos8X2 LspPlugInPluginsLv2ParaEqualizerX16Monos8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos8X3 LspPlugInPluginsLv2ParaEqualizerX16Monos8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos8X4 LspPlugInPluginsLv2ParaEqualizerX16Monos8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofft string

const (
	LspPlugInPluginsLv2ParaEqualizerX16MonofftOff LspPlugInPluginsLv2ParaEqualizerX16Monofft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16MonofftPostEq LspPlugInPluginsLv2ParaEqualizerX16Monofft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX16MonofftPreEq LspPlugInPluginsLv2ParaEqualizerX16Monofft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2Off LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft2Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4Off LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft4Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos15X1 LspPlugInPluginsLv2ParaEqualizerX16Monos15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos15X2 LspPlugInPluginsLv2ParaEqualizerX16Monos15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos15X3 LspPlugInPluginsLv2ParaEqualizerX16Monos15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos15X4 LspPlugInPluginsLv2ParaEqualizerX16Monos15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm3ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1Off LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft1Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11Off LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft11Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3Off LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft3Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8Off LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft8Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos12X1 LspPlugInPluginsLv2ParaEqualizerX16Monos12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos12X2 LspPlugInPluginsLv2ParaEqualizerX16Monos12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos12X3 LspPlugInPluginsLv2ParaEqualizerX16Monos12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos12X4 LspPlugInPluginsLv2ParaEqualizerX16Monos12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos14X1 LspPlugInPluginsLv2ParaEqualizerX16Monos14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos14X2 LspPlugInPluginsLv2ParaEqualizerX16Monos14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos14X3 LspPlugInPluginsLv2ParaEqualizerX16Monos14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos14X4 LspPlugInPluginsLv2ParaEqualizerX16Monos14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm10ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm2ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm9ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monomode string

const (
	LspPlugInPluginsLv2ParaEqualizerX16MonomodeIir LspPlugInPluginsLv2ParaEqualizerX16Monomode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX16MonomodeFir LspPlugInPluginsLv2ParaEqualizerX16Monomode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX16MonomodeFft LspPlugInPluginsLv2ParaEqualizerX16Monomode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos11X1 LspPlugInPluginsLv2ParaEqualizerX16Monos11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos11X2 LspPlugInPluginsLv2ParaEqualizerX16Monos11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos11X3 LspPlugInPluginsLv2ParaEqualizerX16Monos11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos11X4 LspPlugInPluginsLv2ParaEqualizerX16Monos11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos5X1 LspPlugInPluginsLv2ParaEqualizerX16Monos5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos5X2 LspPlugInPluginsLv2ParaEqualizerX16Monos5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos5X3 LspPlugInPluginsLv2ParaEqualizerX16Monos5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos5X4 LspPlugInPluginsLv2ParaEqualizerX16Monos5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm11ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm8ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX16MonofselFilters07 LspPlugInPluginsLv2ParaEqualizerX16Monofsel = "Filters 0-7" // Filters 0-7 (0) – Filters 0-7
	LspPlugInPluginsLv2ParaEqualizerX16MonofselFilters815 LspPlugInPluginsLv2ParaEqualizerX16Monofsel = "Filters 8-15" // Filters 8-15 (1) – Filters 8-15
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5Off LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft5Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9Off LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft9Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos1X1 LspPlugInPluginsLv2ParaEqualizerX16Monos1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos1X2 LspPlugInPluginsLv2ParaEqualizerX16Monos1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos1X3 LspPlugInPluginsLv2ParaEqualizerX16Monos1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos1X4 LspPlugInPluginsLv2ParaEqualizerX16Monos1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos2X1 LspPlugInPluginsLv2ParaEqualizerX16Monos2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos2X2 LspPlugInPluginsLv2ParaEqualizerX16Monos2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos2X3 LspPlugInPluginsLv2ParaEqualizerX16Monos2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos2X4 LspPlugInPluginsLv2ParaEqualizerX16Monos2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos3X1 LspPlugInPluginsLv2ParaEqualizerX16Monos3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos3X2 LspPlugInPluginsLv2ParaEqualizerX16Monos3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos3X3 LspPlugInPluginsLv2ParaEqualizerX16Monos3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos3X4 LspPlugInPluginsLv2ParaEqualizerX16Monos3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm14ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm15ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm4ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm6ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6Off LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft6Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos13X1 LspPlugInPluginsLv2ParaEqualizerX16Monos13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos13X2 LspPlugInPluginsLv2ParaEqualizerX16Monos13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos13X3 LspPlugInPluginsLv2ParaEqualizerX16Monos13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos13X4 LspPlugInPluginsLv2ParaEqualizerX16Monos13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos9X1 LspPlugInPluginsLv2ParaEqualizerX16Monos9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos9X2 LspPlugInPluginsLv2ParaEqualizerX16Monos9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos9X3 LspPlugInPluginsLv2ParaEqualizerX16Monos9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos9X4 LspPlugInPluginsLv2ParaEqualizerX16Monos9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm0ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm5ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0Off LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft0Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15Off LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft15Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7Off LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft7Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos4X1 LspPlugInPluginsLv2ParaEqualizerX16Monos4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos4X2 LspPlugInPluginsLv2ParaEqualizerX16Monos4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos4X3 LspPlugInPluginsLv2ParaEqualizerX16Monos4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos4X4 LspPlugInPluginsLv2ParaEqualizerX16Monos4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos7X1 LspPlugInPluginsLv2ParaEqualizerX16Monos7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos7X2 LspPlugInPluginsLv2ParaEqualizerX16Monos7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos7X3 LspPlugInPluginsLv2ParaEqualizerX16Monos7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos7X4 LspPlugInPluginsLv2ParaEqualizerX16Monos7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm12ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monofm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13RlcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13RlcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13BwcBt LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13BwcMt LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13LrxBt LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13LrxMt LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Monofm13ApoDr LspPlugInPluginsLv2ParaEqualizerX16Monofm13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Monoft14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14Off LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14Bell LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14HiPass LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14HiShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14LoPass LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14LoShelf LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14Notch LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14Resonance LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Monoft14Allpass LspPlugInPluginsLv2ParaEqualizerX16Monoft14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos0X1 LspPlugInPluginsLv2ParaEqualizerX16Monos0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos0X2 LspPlugInPluginsLv2ParaEqualizerX16Monos0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos0X3 LspPlugInPluginsLv2ParaEqualizerX16Monos0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos0X4 LspPlugInPluginsLv2ParaEqualizerX16Monos0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos10X1 LspPlugInPluginsLv2ParaEqualizerX16Monos10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos10X2 LspPlugInPluginsLv2ParaEqualizerX16Monos10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos10X3 LspPlugInPluginsLv2ParaEqualizerX16Monos10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos10X4 LspPlugInPluginsLv2ParaEqualizerX16Monos10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Monos6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Monos6X1 LspPlugInPluginsLv2ParaEqualizerX16Monos6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Monos6X2 LspPlugInPluginsLv2ParaEqualizerX16Monos6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Monos6X3 LspPlugInPluginsLv2ParaEqualizerX16Monos6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Monos6X4 LspPlugInPluginsLv2ParaEqualizerX16Monos6 = "x4" // x4 (3) – x4
)

