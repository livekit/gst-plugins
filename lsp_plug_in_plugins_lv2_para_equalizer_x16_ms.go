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

type LspPlugInPluginsLv2ParaEqualizerX16Ms struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX16Ms() (*LspPlugInPluginsLv2ParaEqualizerX16Ms, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x16-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX16MsWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX16Ms, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x16-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetBal() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("bal")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX16Msfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-m (bool)
//
// FFT visibility Mid

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFftvM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fftv-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFftvM(value bool) error {
	return e.Element.SetProperty("fftv-m", value)
}

// fftv-s (bool)
//
// FFT visibility Side

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFftvS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fftv-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFftvS(value bool) error {
	return e.Element.SetProperty("fftv-s", value)
}

// fm-0 (float32)
//
// Frequency Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm0(value float32) error {
	return e.Element.SetProperty("fm-0", value)
}

// fm-1 (float32)
//
// Frequency Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm1(value float32) error {
	return e.Element.SetProperty("fm-1", value)
}

// fm-10 (float32)
//
// Frequency Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm10(value float32) error {
	return e.Element.SetProperty("fm-10", value)
}

// fm-11 (float32)
//
// Frequency Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm11(value float32) error {
	return e.Element.SetProperty("fm-11", value)
}

// fm-12 (float32)
//
// Frequency Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm12(value float32) error {
	return e.Element.SetProperty("fm-12", value)
}

// fm-13 (float32)
//
// Frequency Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm13(value float32) error {
	return e.Element.SetProperty("fm-13", value)
}

// fm-14 (float32)
//
// Frequency Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm14(value float32) error {
	return e.Element.SetProperty("fm-14", value)
}

// fm-15 (float32)
//
// Frequency Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm15(value float32) error {
	return e.Element.SetProperty("fm-15", value)
}

// fm-2 (float32)
//
// Frequency Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm2(value float32) error {
	return e.Element.SetProperty("fm-2", value)
}

// fm-3 (float32)
//
// Frequency Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm3(value float32) error {
	return e.Element.SetProperty("fm-3", value)
}

// fm-4 (float32)
//
// Frequency Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm4(value float32) error {
	return e.Element.SetProperty("fm-4", value)
}

// fm-5 (float32)
//
// Frequency Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm5(value float32) error {
	return e.Element.SetProperty("fm-5", value)
}

// fm-6 (float32)
//
// Frequency Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm6(value float32) error {
	return e.Element.SetProperty("fm-6", value)
}

// fm-7 (float32)
//
// Frequency Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm7(value float32) error {
	return e.Element.SetProperty("fm-7", value)
}

// fm-8 (float32)
//
// Frequency Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm8(value float32) error {
	return e.Element.SetProperty("fm-8", value)
}

// fm-9 (float32)
//
// Frequency Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFm9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFm9(value float32) error {
	return e.Element.SetProperty("fm-9", value)
}

// fmm-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm0)
//
// Filter mode Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm0() (interface{}, error) {
	return e.Element.GetProperty("fmm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm0(value interface{}) error {
	return e.Element.SetProperty("fmm-0", value)
}

// fmm-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm1)
//
// Filter mode Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm1() (interface{}, error) {
	return e.Element.GetProperty("fmm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm1(value interface{}) error {
	return e.Element.SetProperty("fmm-1", value)
}

// fmm-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm10)
//
// Filter mode Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm10() (interface{}, error) {
	return e.Element.GetProperty("fmm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm10(value interface{}) error {
	return e.Element.SetProperty("fmm-10", value)
}

// fmm-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm11)
//
// Filter mode Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm11() (interface{}, error) {
	return e.Element.GetProperty("fmm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm11(value interface{}) error {
	return e.Element.SetProperty("fmm-11", value)
}

// fmm-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm12)
//
// Filter mode Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm12() (interface{}, error) {
	return e.Element.GetProperty("fmm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm12(value interface{}) error {
	return e.Element.SetProperty("fmm-12", value)
}

// fmm-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm13)
//
// Filter mode Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm13() (interface{}, error) {
	return e.Element.GetProperty("fmm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm13(value interface{}) error {
	return e.Element.SetProperty("fmm-13", value)
}

// fmm-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm14)
//
// Filter mode Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm14() (interface{}, error) {
	return e.Element.GetProperty("fmm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm14(value interface{}) error {
	return e.Element.SetProperty("fmm-14", value)
}

// fmm-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm15)
//
// Filter mode Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm15() (interface{}, error) {
	return e.Element.GetProperty("fmm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm15(value interface{}) error {
	return e.Element.SetProperty("fmm-15", value)
}

// fmm-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm2)
//
// Filter mode Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm2() (interface{}, error) {
	return e.Element.GetProperty("fmm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm2(value interface{}) error {
	return e.Element.SetProperty("fmm-2", value)
}

// fmm-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm3)
//
// Filter mode Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm3() (interface{}, error) {
	return e.Element.GetProperty("fmm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm3(value interface{}) error {
	return e.Element.SetProperty("fmm-3", value)
}

// fmm-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm4)
//
// Filter mode Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm4() (interface{}, error) {
	return e.Element.GetProperty("fmm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm4(value interface{}) error {
	return e.Element.SetProperty("fmm-4", value)
}

// fmm-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm5)
//
// Filter mode Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm5() (interface{}, error) {
	return e.Element.GetProperty("fmm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm5(value interface{}) error {
	return e.Element.SetProperty("fmm-5", value)
}

// fmm-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm6)
//
// Filter mode Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm6() (interface{}, error) {
	return e.Element.GetProperty("fmm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm6(value interface{}) error {
	return e.Element.SetProperty("fmm-6", value)
}

// fmm-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm7)
//
// Filter mode Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm7() (interface{}, error) {
	return e.Element.GetProperty("fmm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm7(value interface{}) error {
	return e.Element.SetProperty("fmm-7", value)
}

// fmm-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm8)
//
// Filter mode Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm8() (interface{}, error) {
	return e.Element.GetProperty("fmm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm8(value interface{}) error {
	return e.Element.SetProperty("fmm-8", value)
}

// fmm-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfmm9)
//
// Filter mode Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFmm9() (interface{}, error) {
	return e.Element.GetProperty("fmm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFmm9(value interface{}) error {
	return e.Element.SetProperty("fmm-9", value)
}

// fms-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms0)
//
// Filter mode Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms0() (interface{}, error) {
	return e.Element.GetProperty("fms-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms0(value interface{}) error {
	return e.Element.SetProperty("fms-0", value)
}

// fms-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms1)
//
// Filter mode Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms1() (interface{}, error) {
	return e.Element.GetProperty("fms-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms1(value interface{}) error {
	return e.Element.SetProperty("fms-1", value)
}

// fms-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms10)
//
// Filter mode Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms10() (interface{}, error) {
	return e.Element.GetProperty("fms-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms10(value interface{}) error {
	return e.Element.SetProperty("fms-10", value)
}

// fms-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms11)
//
// Filter mode Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms11() (interface{}, error) {
	return e.Element.GetProperty("fms-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms11(value interface{}) error {
	return e.Element.SetProperty("fms-11", value)
}

// fms-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms12)
//
// Filter mode Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms12() (interface{}, error) {
	return e.Element.GetProperty("fms-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms12(value interface{}) error {
	return e.Element.SetProperty("fms-12", value)
}

// fms-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms13)
//
// Filter mode Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms13() (interface{}, error) {
	return e.Element.GetProperty("fms-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms13(value interface{}) error {
	return e.Element.SetProperty("fms-13", value)
}

// fms-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms14)
//
// Filter mode Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms14() (interface{}, error) {
	return e.Element.GetProperty("fms-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms14(value interface{}) error {
	return e.Element.SetProperty("fms-14", value)
}

// fms-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms15)
//
// Filter mode Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms15() (interface{}, error) {
	return e.Element.GetProperty("fms-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms15(value interface{}) error {
	return e.Element.SetProperty("fms-15", value)
}

// fms-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms2)
//
// Filter mode Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms2() (interface{}, error) {
	return e.Element.GetProperty("fms-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms2(value interface{}) error {
	return e.Element.SetProperty("fms-2", value)
}

// fms-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms3)
//
// Filter mode Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms3() (interface{}, error) {
	return e.Element.GetProperty("fms-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms3(value interface{}) error {
	return e.Element.SetProperty("fms-3", value)
}

// fms-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms4)
//
// Filter mode Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms4() (interface{}, error) {
	return e.Element.GetProperty("fms-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms4(value interface{}) error {
	return e.Element.SetProperty("fms-4", value)
}

// fms-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms5)
//
// Filter mode Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms5() (interface{}, error) {
	return e.Element.GetProperty("fms-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms5(value interface{}) error {
	return e.Element.SetProperty("fms-5", value)
}

// fms-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms6)
//
// Filter mode Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms6() (interface{}, error) {
	return e.Element.GetProperty("fms-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms6(value interface{}) error {
	return e.Element.SetProperty("fms-6", value)
}

// fms-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms7)
//
// Filter mode Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms7() (interface{}, error) {
	return e.Element.GetProperty("fms-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms7(value interface{}) error {
	return e.Element.SetProperty("fms-7", value)
}

// fms-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms8)
//
// Filter mode Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms8() (interface{}, error) {
	return e.Element.GetProperty("fms-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms8(value interface{}) error {
	return e.Element.SetProperty("fms-8", value)
}

// fms-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfms9)
//
// Filter mode Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFms9() (interface{}, error) {
	return e.Element.GetProperty("fms-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFms9(value interface{}) error {
	return e.Element.SetProperty("fms-9", value)
}

// frqs-m (float32)
//
// Frequency shift Mid

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFrqsM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frqs-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFrqsM(value float32) error {
	return e.Element.SetProperty("frqs-m", value)
}

// frqs-s (float32)
//
// Frequency shift Side

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFrqsS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frqs-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFrqsS(value float32) error {
	return e.Element.SetProperty("frqs-s", value)
}

// fs-0 (float32)
//
// Frequency Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs0(value float32) error {
	return e.Element.SetProperty("fs-0", value)
}

// fs-1 (float32)
//
// Frequency Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs1(value float32) error {
	return e.Element.SetProperty("fs-1", value)
}

// fs-10 (float32)
//
// Frequency Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs10(value float32) error {
	return e.Element.SetProperty("fs-10", value)
}

// fs-11 (float32)
//
// Frequency Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs11(value float32) error {
	return e.Element.SetProperty("fs-11", value)
}

// fs-12 (float32)
//
// Frequency Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs12(value float32) error {
	return e.Element.SetProperty("fs-12", value)
}

// fs-13 (float32)
//
// Frequency Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs13(value float32) error {
	return e.Element.SetProperty("fs-13", value)
}

// fs-14 (float32)
//
// Frequency Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs14(value float32) error {
	return e.Element.SetProperty("fs-14", value)
}

// fs-15 (float32)
//
// Frequency Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs15(value float32) error {
	return e.Element.SetProperty("fs-15", value)
}

// fs-2 (float32)
//
// Frequency Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs2(value float32) error {
	return e.Element.SetProperty("fs-2", value)
}

// fs-3 (float32)
//
// Frequency Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs3(value float32) error {
	return e.Element.SetProperty("fs-3", value)
}

// fs-4 (float32)
//
// Frequency Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs4(value float32) error {
	return e.Element.SetProperty("fs-4", value)
}

// fs-5 (float32)
//
// Frequency Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs5(value float32) error {
	return e.Element.SetProperty("fs-5", value)
}

// fs-6 (float32)
//
// Frequency Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs6(value float32) error {
	return e.Element.SetProperty("fs-6", value)
}

// fs-7 (float32)
//
// Frequency Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs7(value float32) error {
	return e.Element.SetProperty("fs-7", value)
}

// fs-8 (float32)
//
// Frequency Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs8(value float32) error {
	return e.Element.SetProperty("fs-8", value)
}

// fs-9 (float32)
//
// Frequency Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFs9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFs9(value float32) error {
	return e.Element.SetProperty("fs-9", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX16Msfsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ftm-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm0)
//
// Filter type Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm0() (interface{}, error) {
	return e.Element.GetProperty("ftm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm0(value interface{}) error {
	return e.Element.SetProperty("ftm-0", value)
}

// ftm-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm1)
//
// Filter type Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm1() (interface{}, error) {
	return e.Element.GetProperty("ftm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm1(value interface{}) error {
	return e.Element.SetProperty("ftm-1", value)
}

// ftm-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm10)
//
// Filter type Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm10() (interface{}, error) {
	return e.Element.GetProperty("ftm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm10(value interface{}) error {
	return e.Element.SetProperty("ftm-10", value)
}

// ftm-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm11)
//
// Filter type Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm11() (interface{}, error) {
	return e.Element.GetProperty("ftm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm11(value interface{}) error {
	return e.Element.SetProperty("ftm-11", value)
}

// ftm-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm12)
//
// Filter type Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm12() (interface{}, error) {
	return e.Element.GetProperty("ftm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm12(value interface{}) error {
	return e.Element.SetProperty("ftm-12", value)
}

// ftm-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm13)
//
// Filter type Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm13() (interface{}, error) {
	return e.Element.GetProperty("ftm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm13(value interface{}) error {
	return e.Element.SetProperty("ftm-13", value)
}

// ftm-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm14)
//
// Filter type Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm14() (interface{}, error) {
	return e.Element.GetProperty("ftm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm14(value interface{}) error {
	return e.Element.SetProperty("ftm-14", value)
}

// ftm-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm15)
//
// Filter type Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm15() (interface{}, error) {
	return e.Element.GetProperty("ftm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm15(value interface{}) error {
	return e.Element.SetProperty("ftm-15", value)
}

// ftm-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm2)
//
// Filter type Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm2() (interface{}, error) {
	return e.Element.GetProperty("ftm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm2(value interface{}) error {
	return e.Element.SetProperty("ftm-2", value)
}

// ftm-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm3)
//
// Filter type Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm3() (interface{}, error) {
	return e.Element.GetProperty("ftm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm3(value interface{}) error {
	return e.Element.SetProperty("ftm-3", value)
}

// ftm-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm4)
//
// Filter type Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm4() (interface{}, error) {
	return e.Element.GetProperty("ftm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm4(value interface{}) error {
	return e.Element.SetProperty("ftm-4", value)
}

// ftm-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm5)
//
// Filter type Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm5() (interface{}, error) {
	return e.Element.GetProperty("ftm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm5(value interface{}) error {
	return e.Element.SetProperty("ftm-5", value)
}

// ftm-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm6)
//
// Filter type Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm6() (interface{}, error) {
	return e.Element.GetProperty("ftm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm6(value interface{}) error {
	return e.Element.SetProperty("ftm-6", value)
}

// ftm-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm7)
//
// Filter type Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm7() (interface{}, error) {
	return e.Element.GetProperty("ftm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm7(value interface{}) error {
	return e.Element.SetProperty("ftm-7", value)
}

// ftm-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm8)
//
// Filter type Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm8() (interface{}, error) {
	return e.Element.GetProperty("ftm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm8(value interface{}) error {
	return e.Element.SetProperty("ftm-8", value)
}

// ftm-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Msftm9)
//
// Filter type Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFtm9() (interface{}, error) {
	return e.Element.GetProperty("ftm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFtm9(value interface{}) error {
	return e.Element.SetProperty("ftm-9", value)
}

// fts-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts0)
//
// Filter type Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts0() (interface{}, error) {
	return e.Element.GetProperty("fts-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts0(value interface{}) error {
	return e.Element.SetProperty("fts-0", value)
}

// fts-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts1)
//
// Filter type Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts1() (interface{}, error) {
	return e.Element.GetProperty("fts-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts1(value interface{}) error {
	return e.Element.SetProperty("fts-1", value)
}

// fts-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts10)
//
// Filter type Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts10() (interface{}, error) {
	return e.Element.GetProperty("fts-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts10(value interface{}) error {
	return e.Element.SetProperty("fts-10", value)
}

// fts-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts11)
//
// Filter type Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts11() (interface{}, error) {
	return e.Element.GetProperty("fts-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts11(value interface{}) error {
	return e.Element.SetProperty("fts-11", value)
}

// fts-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts12)
//
// Filter type Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts12() (interface{}, error) {
	return e.Element.GetProperty("fts-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts12(value interface{}) error {
	return e.Element.SetProperty("fts-12", value)
}

// fts-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts13)
//
// Filter type Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts13() (interface{}, error) {
	return e.Element.GetProperty("fts-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts13(value interface{}) error {
	return e.Element.SetProperty("fts-13", value)
}

// fts-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts14)
//
// Filter type Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts14() (interface{}, error) {
	return e.Element.GetProperty("fts-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts14(value interface{}) error {
	return e.Element.SetProperty("fts-14", value)
}

// fts-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts15)
//
// Filter type Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts15() (interface{}, error) {
	return e.Element.GetProperty("fts-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts15(value interface{}) error {
	return e.Element.SetProperty("fts-15", value)
}

// fts-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts2)
//
// Filter type Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts2() (interface{}, error) {
	return e.Element.GetProperty("fts-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts2(value interface{}) error {
	return e.Element.SetProperty("fts-2", value)
}

// fts-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts3)
//
// Filter type Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts3() (interface{}, error) {
	return e.Element.GetProperty("fts-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts3(value interface{}) error {
	return e.Element.SetProperty("fts-3", value)
}

// fts-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts4)
//
// Filter type Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts4() (interface{}, error) {
	return e.Element.GetProperty("fts-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts4(value interface{}) error {
	return e.Element.SetProperty("fts-4", value)
}

// fts-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts5)
//
// Filter type Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts5() (interface{}, error) {
	return e.Element.GetProperty("fts-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts5(value interface{}) error {
	return e.Element.SetProperty("fts-5", value)
}

// fts-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts6)
//
// Filter type Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts6() (interface{}, error) {
	return e.Element.GetProperty("fts-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts6(value interface{}) error {
	return e.Element.SetProperty("fts-6", value)
}

// fts-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts7)
//
// Filter type Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts7() (interface{}, error) {
	return e.Element.GetProperty("fts-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts7(value interface{}) error {
	return e.Element.SetProperty("fts-7", value)
}

// fts-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts8)
//
// Filter type Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts8() (interface{}, error) {
	return e.Element.GetProperty("fts-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts8(value interface{}) error {
	return e.Element.SetProperty("fts-8", value)
}

// fts-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Msfts9)
//
// Filter type Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFts9() (interface{}, error) {
	return e.Element.GetProperty("fts-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetFts9(value interface{}) error {
	return e.Element.SetProperty("fts-9", value)
}

// fvm-0 (bool)
//
// Filter visibility Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-1 (bool)
//
// Filter visibility Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-10 (bool)
//
// Filter visibility Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-11 (bool)
//
// Filter visibility Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-12 (bool)
//
// Filter visibility Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-13 (bool)
//
// Filter visibility Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-14 (bool)
//
// Filter visibility Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-15 (bool)
//
// Filter visibility Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-2 (bool)
//
// Filter visibility Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-3 (bool)
//
// Filter visibility Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-4 (bool)
//
// Filter visibility Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-5 (bool)
//
// Filter visibility Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-6 (bool)
//
// Filter visibility Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-7 (bool)
//
// Filter visibility Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-8 (bool)
//
// Filter visibility Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-9 (bool)
//
// Filter visibility Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvm9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-0 (bool)
//
// Filter visibility Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-1 (bool)
//
// Filter visibility Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-10 (bool)
//
// Filter visibility Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-11 (bool)
//
// Filter visibility Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-12 (bool)
//
// Filter visibility Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-13 (bool)
//
// Filter visibility Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-14 (bool)
//
// Filter visibility Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-15 (bool)
//
// Filter visibility Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-2 (bool)
//
// Filter visibility Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-3 (bool)
//
// Filter visibility Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-4 (bool)
//
// Filter visibility Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-5 (bool)
//
// Filter visibility Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-6 (bool)
//
// Filter visibility Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-7 (bool)
//
// Filter visibility Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-8 (bool)
//
// Filter visibility Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-9 (bool)
//
// Filter visibility Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetFvs9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gain-m (float32)
//
// Mid gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGainM() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gain-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGainM(value float32) error {
	return e.Element.SetProperty("gain-m", value)
}

// gain-s (float32)
//
// Side gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGainS() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gain-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGainS(value float32) error {
	return e.Element.SetProperty("gain-s", value)
}

// gm-0 (float32)
//
// Gain Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm0(value float32) error {
	return e.Element.SetProperty("gm-0", value)
}

// gm-1 (float32)
//
// Gain Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm1(value float32) error {
	return e.Element.SetProperty("gm-1", value)
}

// gm-10 (float32)
//
// Gain Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm10(value float32) error {
	return e.Element.SetProperty("gm-10", value)
}

// gm-11 (float32)
//
// Gain Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm11(value float32) error {
	return e.Element.SetProperty("gm-11", value)
}

// gm-12 (float32)
//
// Gain Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm12(value float32) error {
	return e.Element.SetProperty("gm-12", value)
}

// gm-13 (float32)
//
// Gain Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm13(value float32) error {
	return e.Element.SetProperty("gm-13", value)
}

// gm-14 (float32)
//
// Gain Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm14(value float32) error {
	return e.Element.SetProperty("gm-14", value)
}

// gm-15 (float32)
//
// Gain Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm15(value float32) error {
	return e.Element.SetProperty("gm-15", value)
}

// gm-2 (float32)
//
// Gain Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm2(value float32) error {
	return e.Element.SetProperty("gm-2", value)
}

// gm-3 (float32)
//
// Gain Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm3(value float32) error {
	return e.Element.SetProperty("gm-3", value)
}

// gm-4 (float32)
//
// Gain Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm4(value float32) error {
	return e.Element.SetProperty("gm-4", value)
}

// gm-5 (float32)
//
// Gain Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm5(value float32) error {
	return e.Element.SetProperty("gm-5", value)
}

// gm-6 (float32)
//
// Gain Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm6(value float32) error {
	return e.Element.SetProperty("gm-6", value)
}

// gm-7 (float32)
//
// Gain Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm7(value float32) error {
	return e.Element.SetProperty("gm-7", value)
}

// gm-8 (float32)
//
// Gain Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm8(value float32) error {
	return e.Element.SetProperty("gm-8", value)
}

// gm-9 (float32)
//
// Gain Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGm9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGm9(value float32) error {
	return e.Element.SetProperty("gm-9", value)
}

// gs-0 (float32)
//
// Gain Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs0(value float32) error {
	return e.Element.SetProperty("gs-0", value)
}

// gs-1 (float32)
//
// Gain Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs1(value float32) error {
	return e.Element.SetProperty("gs-1", value)
}

// gs-10 (float32)
//
// Gain Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs10(value float32) error {
	return e.Element.SetProperty("gs-10", value)
}

// gs-11 (float32)
//
// Gain Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs11(value float32) error {
	return e.Element.SetProperty("gs-11", value)
}

// gs-12 (float32)
//
// Gain Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs12(value float32) error {
	return e.Element.SetProperty("gs-12", value)
}

// gs-13 (float32)
//
// Gain Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs13(value float32) error {
	return e.Element.SetProperty("gs-13", value)
}

// gs-14 (float32)
//
// Gain Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs14(value float32) error {
	return e.Element.SetProperty("gs-14", value)
}

// gs-15 (float32)
//
// Gain Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs15(value float32) error {
	return e.Element.SetProperty("gs-15", value)
}

// gs-2 (float32)
//
// Gain Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs2(value float32) error {
	return e.Element.SetProperty("gs-2", value)
}

// gs-3 (float32)
//
// Gain Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs3(value float32) error {
	return e.Element.SetProperty("gs-3", value)
}

// gs-4 (float32)
//
// Gain Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs4(value float32) error {
	return e.Element.SetProperty("gs-4", value)
}

// gs-5 (float32)
//
// Gain Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs5(value float32) error {
	return e.Element.SetProperty("gs-5", value)
}

// gs-6 (float32)
//
// Gain Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs6(value float32) error {
	return e.Element.SetProperty("gs-6", value)
}

// gs-7 (float32)
//
// Gain Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs7(value float32) error {
	return e.Element.SetProperty("gs-7", value)
}

// gs-8 (float32)
//
// Gain Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs8(value float32) error {
	return e.Element.SetProperty("gs-8", value)
}

// gs-9 (float32)
//
// Gain Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetGs9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetGs9(value float32) error {
	return e.Element.SetProperty("gs-9", value)
}

// huem-0 (float32)
//
// Hue Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem0(value float32) error {
	return e.Element.SetProperty("huem-0", value)
}

// huem-1 (float32)
//
// Hue Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem1(value float32) error {
	return e.Element.SetProperty("huem-1", value)
}

// huem-10 (float32)
//
// Hue Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem10(value float32) error {
	return e.Element.SetProperty("huem-10", value)
}

// huem-11 (float32)
//
// Hue Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem11(value float32) error {
	return e.Element.SetProperty("huem-11", value)
}

// huem-12 (float32)
//
// Hue Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem12(value float32) error {
	return e.Element.SetProperty("huem-12", value)
}

// huem-13 (float32)
//
// Hue Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem13(value float32) error {
	return e.Element.SetProperty("huem-13", value)
}

// huem-14 (float32)
//
// Hue Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem14(value float32) error {
	return e.Element.SetProperty("huem-14", value)
}

// huem-15 (float32)
//
// Hue Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem15(value float32) error {
	return e.Element.SetProperty("huem-15", value)
}

// huem-2 (float32)
//
// Hue Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem2(value float32) error {
	return e.Element.SetProperty("huem-2", value)
}

// huem-3 (float32)
//
// Hue Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem3(value float32) error {
	return e.Element.SetProperty("huem-3", value)
}

// huem-4 (float32)
//
// Hue Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem4(value float32) error {
	return e.Element.SetProperty("huem-4", value)
}

// huem-5 (float32)
//
// Hue Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem5(value float32) error {
	return e.Element.SetProperty("huem-5", value)
}

// huem-6 (float32)
//
// Hue Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem6(value float32) error {
	return e.Element.SetProperty("huem-6", value)
}

// huem-7 (float32)
//
// Hue Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem7(value float32) error {
	return e.Element.SetProperty("huem-7", value)
}

// huem-8 (float32)
//
// Hue Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem8(value float32) error {
	return e.Element.SetProperty("huem-8", value)
}

// huem-9 (float32)
//
// Hue Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHuem9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHuem9(value float32) error {
	return e.Element.SetProperty("huem-9", value)
}

// hues-0 (float32)
//
// Hue Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues0(value float32) error {
	return e.Element.SetProperty("hues-0", value)
}

// hues-1 (float32)
//
// Hue Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues1(value float32) error {
	return e.Element.SetProperty("hues-1", value)
}

// hues-10 (float32)
//
// Hue Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues10(value float32) error {
	return e.Element.SetProperty("hues-10", value)
}

// hues-11 (float32)
//
// Hue Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues11(value float32) error {
	return e.Element.SetProperty("hues-11", value)
}

// hues-12 (float32)
//
// Hue Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues12(value float32) error {
	return e.Element.SetProperty("hues-12", value)
}

// hues-13 (float32)
//
// Hue Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues13(value float32) error {
	return e.Element.SetProperty("hues-13", value)
}

// hues-14 (float32)
//
// Hue Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues14(value float32) error {
	return e.Element.SetProperty("hues-14", value)
}

// hues-15 (float32)
//
// Hue Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues15(value float32) error {
	return e.Element.SetProperty("hues-15", value)
}

// hues-2 (float32)
//
// Hue Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues2(value float32) error {
	return e.Element.SetProperty("hues-2", value)
}

// hues-3 (float32)
//
// Hue Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues3(value float32) error {
	return e.Element.SetProperty("hues-3", value)
}

// hues-4 (float32)
//
// Hue Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues4(value float32) error {
	return e.Element.SetProperty("hues-4", value)
}

// hues-5 (float32)
//
// Hue Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues5(value float32) error {
	return e.Element.SetProperty("hues-5", value)
}

// hues-6 (float32)
//
// Hue Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues6(value float32) error {
	return e.Element.SetProperty("hues-6", value)
}

// hues-7 (float32)
//
// Hue Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues7(value float32) error {
	return e.Element.SetProperty("hues-7", value)
}

// hues-8 (float32)
//
// Hue Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues8(value float32) error {
	return e.Element.SetProperty("hues-8", value)
}

// hues-9 (float32)
//
// Hue Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetHues9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetHues9(value float32) error {
	return e.Element.SetProperty("hues-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetIml() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("iml")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// imr (float32)
//
// Input signal meter Right

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetImr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// lstn (bool)
//
// Mid/Side listen

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetLstn() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lstn")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mode (GstLspPlugInPluginsLv2ParaEqualizerX16Msmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetOutLatency() (int, error) {
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

// qm-0 (float32)
//
// Quality factor Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm0(value float32) error {
	return e.Element.SetProperty("qm-0", value)
}

// qm-1 (float32)
//
// Quality factor Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm1(value float32) error {
	return e.Element.SetProperty("qm-1", value)
}

// qm-10 (float32)
//
// Quality factor Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm10(value float32) error {
	return e.Element.SetProperty("qm-10", value)
}

// qm-11 (float32)
//
// Quality factor Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm11(value float32) error {
	return e.Element.SetProperty("qm-11", value)
}

// qm-12 (float32)
//
// Quality factor Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm12(value float32) error {
	return e.Element.SetProperty("qm-12", value)
}

// qm-13 (float32)
//
// Quality factor Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm13(value float32) error {
	return e.Element.SetProperty("qm-13", value)
}

// qm-14 (float32)
//
// Quality factor Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm14(value float32) error {
	return e.Element.SetProperty("qm-14", value)
}

// qm-15 (float32)
//
// Quality factor Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm15(value float32) error {
	return e.Element.SetProperty("qm-15", value)
}

// qm-2 (float32)
//
// Quality factor Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm2(value float32) error {
	return e.Element.SetProperty("qm-2", value)
}

// qm-3 (float32)
//
// Quality factor Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm3(value float32) error {
	return e.Element.SetProperty("qm-3", value)
}

// qm-4 (float32)
//
// Quality factor Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm4(value float32) error {
	return e.Element.SetProperty("qm-4", value)
}

// qm-5 (float32)
//
// Quality factor Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm5(value float32) error {
	return e.Element.SetProperty("qm-5", value)
}

// qm-6 (float32)
//
// Quality factor Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm6(value float32) error {
	return e.Element.SetProperty("qm-6", value)
}

// qm-7 (float32)
//
// Quality factor Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm7(value float32) error {
	return e.Element.SetProperty("qm-7", value)
}

// qm-8 (float32)
//
// Quality factor Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm8(value float32) error {
	return e.Element.SetProperty("qm-8", value)
}

// qm-9 (float32)
//
// Quality factor Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQm9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQm9(value float32) error {
	return e.Element.SetProperty("qm-9", value)
}

// qs-0 (float32)
//
// Quality factor Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs0(value float32) error {
	return e.Element.SetProperty("qs-0", value)
}

// qs-1 (float32)
//
// Quality factor Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs1(value float32) error {
	return e.Element.SetProperty("qs-1", value)
}

// qs-10 (float32)
//
// Quality factor Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs10(value float32) error {
	return e.Element.SetProperty("qs-10", value)
}

// qs-11 (float32)
//
// Quality factor Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs11(value float32) error {
	return e.Element.SetProperty("qs-11", value)
}

// qs-12 (float32)
//
// Quality factor Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs12(value float32) error {
	return e.Element.SetProperty("qs-12", value)
}

// qs-13 (float32)
//
// Quality factor Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs13(value float32) error {
	return e.Element.SetProperty("qs-13", value)
}

// qs-14 (float32)
//
// Quality factor Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs14(value float32) error {
	return e.Element.SetProperty("qs-14", value)
}

// qs-15 (float32)
//
// Quality factor Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs15(value float32) error {
	return e.Element.SetProperty("qs-15", value)
}

// qs-2 (float32)
//
// Quality factor Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs2(value float32) error {
	return e.Element.SetProperty("qs-2", value)
}

// qs-3 (float32)
//
// Quality factor Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs3(value float32) error {
	return e.Element.SetProperty("qs-3", value)
}

// qs-4 (float32)
//
// Quality factor Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs4(value float32) error {
	return e.Element.SetProperty("qs-4", value)
}

// qs-5 (float32)
//
// Quality factor Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs5(value float32) error {
	return e.Element.SetProperty("qs-5", value)
}

// qs-6 (float32)
//
// Quality factor Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs6(value float32) error {
	return e.Element.SetProperty("qs-6", value)
}

// qs-7 (float32)
//
// Quality factor Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs7(value float32) error {
	return e.Element.SetProperty("qs-7", value)
}

// qs-8 (float32)
//
// Quality factor Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs8(value float32) error {
	return e.Element.SetProperty("qs-8", value)
}

// qs-9 (float32)
//
// Quality factor Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetQs9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetQs9(value float32) error {
	return e.Element.SetProperty("qs-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sm-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm0)
//
// Filter slope Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm0() (interface{}, error) {
	return e.Element.GetProperty("sm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm0(value interface{}) error {
	return e.Element.SetProperty("sm-0", value)
}

// sm-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm1)
//
// Filter slope Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm1() (interface{}, error) {
	return e.Element.GetProperty("sm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm1(value interface{}) error {
	return e.Element.SetProperty("sm-1", value)
}

// sm-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm10)
//
// Filter slope Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm10() (interface{}, error) {
	return e.Element.GetProperty("sm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm10(value interface{}) error {
	return e.Element.SetProperty("sm-10", value)
}

// sm-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm11)
//
// Filter slope Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm11() (interface{}, error) {
	return e.Element.GetProperty("sm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm11(value interface{}) error {
	return e.Element.SetProperty("sm-11", value)
}

// sm-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm12)
//
// Filter slope Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm12() (interface{}, error) {
	return e.Element.GetProperty("sm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm12(value interface{}) error {
	return e.Element.SetProperty("sm-12", value)
}

// sm-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm13)
//
// Filter slope Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm13() (interface{}, error) {
	return e.Element.GetProperty("sm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm13(value interface{}) error {
	return e.Element.SetProperty("sm-13", value)
}

// sm-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm14)
//
// Filter slope Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm14() (interface{}, error) {
	return e.Element.GetProperty("sm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm14(value interface{}) error {
	return e.Element.SetProperty("sm-14", value)
}

// sm-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm15)
//
// Filter slope Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm15() (interface{}, error) {
	return e.Element.GetProperty("sm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm15(value interface{}) error {
	return e.Element.SetProperty("sm-15", value)
}

// sm-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm2)
//
// Filter slope Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm2() (interface{}, error) {
	return e.Element.GetProperty("sm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm2(value interface{}) error {
	return e.Element.SetProperty("sm-2", value)
}

// sm-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm3)
//
// Filter slope Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm3() (interface{}, error) {
	return e.Element.GetProperty("sm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm3(value interface{}) error {
	return e.Element.SetProperty("sm-3", value)
}

// sm-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm4)
//
// Filter slope Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm4() (interface{}, error) {
	return e.Element.GetProperty("sm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm4(value interface{}) error {
	return e.Element.SetProperty("sm-4", value)
}

// sm-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm5)
//
// Filter slope Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm5() (interface{}, error) {
	return e.Element.GetProperty("sm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm5(value interface{}) error {
	return e.Element.SetProperty("sm-5", value)
}

// sm-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm6)
//
// Filter slope Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm6() (interface{}, error) {
	return e.Element.GetProperty("sm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm6(value interface{}) error {
	return e.Element.SetProperty("sm-6", value)
}

// sm-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm7)
//
// Filter slope Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm7() (interface{}, error) {
	return e.Element.GetProperty("sm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm7(value interface{}) error {
	return e.Element.SetProperty("sm-7", value)
}

// sm-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm8)
//
// Filter slope Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm8() (interface{}, error) {
	return e.Element.GetProperty("sm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm8(value interface{}) error {
	return e.Element.SetProperty("sm-8", value)
}

// sm-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Mssm9)
//
// Filter slope Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSm9() (interface{}, error) {
	return e.Element.GetProperty("sm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSm9(value interface{}) error {
	return e.Element.SetProperty("sm-9", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSml() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sml")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smr (float32)
//
// Output signal meter Right

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSmr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ss-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss0)
//
// Filter slope Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs0() (interface{}, error) {
	return e.Element.GetProperty("ss-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs0(value interface{}) error {
	return e.Element.SetProperty("ss-0", value)
}

// ss-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss1)
//
// Filter slope Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs1() (interface{}, error) {
	return e.Element.GetProperty("ss-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs1(value interface{}) error {
	return e.Element.SetProperty("ss-1", value)
}

// ss-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss10)
//
// Filter slope Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs10() (interface{}, error) {
	return e.Element.GetProperty("ss-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs10(value interface{}) error {
	return e.Element.SetProperty("ss-10", value)
}

// ss-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss11)
//
// Filter slope Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs11() (interface{}, error) {
	return e.Element.GetProperty("ss-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs11(value interface{}) error {
	return e.Element.SetProperty("ss-11", value)
}

// ss-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss12)
//
// Filter slope Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs12() (interface{}, error) {
	return e.Element.GetProperty("ss-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs12(value interface{}) error {
	return e.Element.SetProperty("ss-12", value)
}

// ss-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss13)
//
// Filter slope Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs13() (interface{}, error) {
	return e.Element.GetProperty("ss-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs13(value interface{}) error {
	return e.Element.SetProperty("ss-13", value)
}

// ss-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss14)
//
// Filter slope Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs14() (interface{}, error) {
	return e.Element.GetProperty("ss-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs14(value interface{}) error {
	return e.Element.SetProperty("ss-14", value)
}

// ss-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss15)
//
// Filter slope Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs15() (interface{}, error) {
	return e.Element.GetProperty("ss-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs15(value interface{}) error {
	return e.Element.SetProperty("ss-15", value)
}

// ss-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss2)
//
// Filter slope Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs2() (interface{}, error) {
	return e.Element.GetProperty("ss-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs2(value interface{}) error {
	return e.Element.SetProperty("ss-2", value)
}

// ss-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss3)
//
// Filter slope Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs3() (interface{}, error) {
	return e.Element.GetProperty("ss-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs3(value interface{}) error {
	return e.Element.SetProperty("ss-3", value)
}

// ss-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss4)
//
// Filter slope Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs4() (interface{}, error) {
	return e.Element.GetProperty("ss-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs4(value interface{}) error {
	return e.Element.SetProperty("ss-4", value)
}

// ss-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss5)
//
// Filter slope Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs5() (interface{}, error) {
	return e.Element.GetProperty("ss-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs5(value interface{}) error {
	return e.Element.SetProperty("ss-5", value)
}

// ss-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss6)
//
// Filter slope Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs6() (interface{}, error) {
	return e.Element.GetProperty("ss-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs6(value interface{}) error {
	return e.Element.SetProperty("ss-6", value)
}

// ss-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss7)
//
// Filter slope Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs7() (interface{}, error) {
	return e.Element.GetProperty("ss-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs7(value interface{}) error {
	return e.Element.SetProperty("ss-7", value)
}

// ss-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss8)
//
// Filter slope Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs8() (interface{}, error) {
	return e.Element.GetProperty("ss-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs8(value interface{}) error {
	return e.Element.SetProperty("ss-8", value)
}

// ss-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Msss9)
//
// Filter slope Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetSs9() (interface{}, error) {
	return e.Element.GetProperty("ss-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetSs9(value interface{}) error {
	return e.Element.SetProperty("ss-9", value)
}

// xmm-0 (bool)
//
// Filter mute Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm0(value bool) error {
	return e.Element.SetProperty("xmm-0", value)
}

// xmm-1 (bool)
//
// Filter mute Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm1(value bool) error {
	return e.Element.SetProperty("xmm-1", value)
}

// xmm-10 (bool)
//
// Filter mute Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm10(value bool) error {
	return e.Element.SetProperty("xmm-10", value)
}

// xmm-11 (bool)
//
// Filter mute Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm11(value bool) error {
	return e.Element.SetProperty("xmm-11", value)
}

// xmm-12 (bool)
//
// Filter mute Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm12(value bool) error {
	return e.Element.SetProperty("xmm-12", value)
}

// xmm-13 (bool)
//
// Filter mute Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm13(value bool) error {
	return e.Element.SetProperty("xmm-13", value)
}

// xmm-14 (bool)
//
// Filter mute Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm14(value bool) error {
	return e.Element.SetProperty("xmm-14", value)
}

// xmm-15 (bool)
//
// Filter mute Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm15(value bool) error {
	return e.Element.SetProperty("xmm-15", value)
}

// xmm-2 (bool)
//
// Filter mute Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm2(value bool) error {
	return e.Element.SetProperty("xmm-2", value)
}

// xmm-3 (bool)
//
// Filter mute Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm3(value bool) error {
	return e.Element.SetProperty("xmm-3", value)
}

// xmm-4 (bool)
//
// Filter mute Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm4(value bool) error {
	return e.Element.SetProperty("xmm-4", value)
}

// xmm-5 (bool)
//
// Filter mute Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm5(value bool) error {
	return e.Element.SetProperty("xmm-5", value)
}

// xmm-6 (bool)
//
// Filter mute Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm6(value bool) error {
	return e.Element.SetProperty("xmm-6", value)
}

// xmm-7 (bool)
//
// Filter mute Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm7(value bool) error {
	return e.Element.SetProperty("xmm-7", value)
}

// xmm-8 (bool)
//
// Filter mute Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm8(value bool) error {
	return e.Element.SetProperty("xmm-8", value)
}

// xmm-9 (bool)
//
// Filter mute Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXmm9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXmm9(value bool) error {
	return e.Element.SetProperty("xmm-9", value)
}

// xms-0 (bool)
//
// Filter mute Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms0(value bool) error {
	return e.Element.SetProperty("xms-0", value)
}

// xms-1 (bool)
//
// Filter mute Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms1(value bool) error {
	return e.Element.SetProperty("xms-1", value)
}

// xms-10 (bool)
//
// Filter mute Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms10(value bool) error {
	return e.Element.SetProperty("xms-10", value)
}

// xms-11 (bool)
//
// Filter mute Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms11(value bool) error {
	return e.Element.SetProperty("xms-11", value)
}

// xms-12 (bool)
//
// Filter mute Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms12(value bool) error {
	return e.Element.SetProperty("xms-12", value)
}

// xms-13 (bool)
//
// Filter mute Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms13(value bool) error {
	return e.Element.SetProperty("xms-13", value)
}

// xms-14 (bool)
//
// Filter mute Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms14(value bool) error {
	return e.Element.SetProperty("xms-14", value)
}

// xms-15 (bool)
//
// Filter mute Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms15(value bool) error {
	return e.Element.SetProperty("xms-15", value)
}

// xms-2 (bool)
//
// Filter mute Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms2(value bool) error {
	return e.Element.SetProperty("xms-2", value)
}

// xms-3 (bool)
//
// Filter mute Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms3(value bool) error {
	return e.Element.SetProperty("xms-3", value)
}

// xms-4 (bool)
//
// Filter mute Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms4(value bool) error {
	return e.Element.SetProperty("xms-4", value)
}

// xms-5 (bool)
//
// Filter mute Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms5(value bool) error {
	return e.Element.SetProperty("xms-5", value)
}

// xms-6 (bool)
//
// Filter mute Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms6(value bool) error {
	return e.Element.SetProperty("xms-6", value)
}

// xms-7 (bool)
//
// Filter mute Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms7(value bool) error {
	return e.Element.SetProperty("xms-7", value)
}

// xms-8 (bool)
//
// Filter mute Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms8(value bool) error {
	return e.Element.SetProperty("xms-8", value)
}

// xms-9 (bool)
//
// Filter mute Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXms9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXms9(value bool) error {
	return e.Element.SetProperty("xms-9", value)
}

// xsm-0 (bool)
//
// Filter solo Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm0(value bool) error {
	return e.Element.SetProperty("xsm-0", value)
}

// xsm-1 (bool)
//
// Filter solo Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm1(value bool) error {
	return e.Element.SetProperty("xsm-1", value)
}

// xsm-10 (bool)
//
// Filter solo Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm10(value bool) error {
	return e.Element.SetProperty("xsm-10", value)
}

// xsm-11 (bool)
//
// Filter solo Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm11(value bool) error {
	return e.Element.SetProperty("xsm-11", value)
}

// xsm-12 (bool)
//
// Filter solo Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm12(value bool) error {
	return e.Element.SetProperty("xsm-12", value)
}

// xsm-13 (bool)
//
// Filter solo Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm13(value bool) error {
	return e.Element.SetProperty("xsm-13", value)
}

// xsm-14 (bool)
//
// Filter solo Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm14(value bool) error {
	return e.Element.SetProperty("xsm-14", value)
}

// xsm-15 (bool)
//
// Filter solo Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm15(value bool) error {
	return e.Element.SetProperty("xsm-15", value)
}

// xsm-2 (bool)
//
// Filter solo Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm2(value bool) error {
	return e.Element.SetProperty("xsm-2", value)
}

// xsm-3 (bool)
//
// Filter solo Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm3(value bool) error {
	return e.Element.SetProperty("xsm-3", value)
}

// xsm-4 (bool)
//
// Filter solo Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm4(value bool) error {
	return e.Element.SetProperty("xsm-4", value)
}

// xsm-5 (bool)
//
// Filter solo Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm5(value bool) error {
	return e.Element.SetProperty("xsm-5", value)
}

// xsm-6 (bool)
//
// Filter solo Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm6(value bool) error {
	return e.Element.SetProperty("xsm-6", value)
}

// xsm-7 (bool)
//
// Filter solo Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm7(value bool) error {
	return e.Element.SetProperty("xsm-7", value)
}

// xsm-8 (bool)
//
// Filter solo Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm8(value bool) error {
	return e.Element.SetProperty("xsm-8", value)
}

// xsm-9 (bool)
//
// Filter solo Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXsm9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXsm9(value bool) error {
	return e.Element.SetProperty("xsm-9", value)
}

// xss-0 (bool)
//
// Filter solo Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss0(value bool) error {
	return e.Element.SetProperty("xss-0", value)
}

// xss-1 (bool)
//
// Filter solo Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss1(value bool) error {
	return e.Element.SetProperty("xss-1", value)
}

// xss-10 (bool)
//
// Filter solo Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss10(value bool) error {
	return e.Element.SetProperty("xss-10", value)
}

// xss-11 (bool)
//
// Filter solo Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss11(value bool) error {
	return e.Element.SetProperty("xss-11", value)
}

// xss-12 (bool)
//
// Filter solo Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss12(value bool) error {
	return e.Element.SetProperty("xss-12", value)
}

// xss-13 (bool)
//
// Filter solo Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss13(value bool) error {
	return e.Element.SetProperty("xss-13", value)
}

// xss-14 (bool)
//
// Filter solo Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss14(value bool) error {
	return e.Element.SetProperty("xss-14", value)
}

// xss-15 (bool)
//
// Filter solo Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss15(value bool) error {
	return e.Element.SetProperty("xss-15", value)
}

// xss-2 (bool)
//
// Filter solo Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss2(value bool) error {
	return e.Element.SetProperty("xss-2", value)
}

// xss-3 (bool)
//
// Filter solo Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss3(value bool) error {
	return e.Element.SetProperty("xss-3", value)
}

// xss-4 (bool)
//
// Filter solo Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss4(value bool) error {
	return e.Element.SetProperty("xss-4", value)
}

// xss-5 (bool)
//
// Filter solo Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss5(value bool) error {
	return e.Element.SetProperty("xss-5", value)
}

// xss-6 (bool)
//
// Filter solo Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss6(value bool) error {
	return e.Element.SetProperty("xss-6", value)
}

// xss-7 (bool)
//
// Filter solo Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss7(value bool) error {
	return e.Element.SetProperty("xss-7", value)
}

// xss-8 (bool)
//
// Filter solo Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss8(value bool) error {
	return e.Element.SetProperty("xss-8", value)
}

// xss-9 (bool)
//
// Filter solo Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetXss9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetXss9(value bool) error {
	return e.Element.SetProperty("xss-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Ms) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX16Msfms6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms6ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11Off LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm11Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2Off LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts2Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm2X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm2X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm2X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm2X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm10ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm13ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm8ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms0ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm5X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm5X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm5X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm5X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss4X1 LspPlugInPluginsLv2ParaEqualizerX16Msss4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss4X2 LspPlugInPluginsLv2ParaEqualizerX16Msss4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss4X3 LspPlugInPluginsLv2ParaEqualizerX16Msss4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss4X4 LspPlugInPluginsLv2ParaEqualizerX16Msss4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm10X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm10X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm10X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm10X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm14X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm14X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm14X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm14X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfft string

const (
	LspPlugInPluginsLv2ParaEqualizerX16MsfftOff LspPlugInPluginsLv2ParaEqualizerX16Msfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16MsfftPostEq LspPlugInPluginsLv2ParaEqualizerX16Msfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX16MsfftPreEq LspPlugInPluginsLv2ParaEqualizerX16Msfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm11ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8Off LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm8Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8Off LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts8Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm15ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms9ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msmode string

const (
	LspPlugInPluginsLv2ParaEqualizerX16MsmodeIir LspPlugInPluginsLv2ParaEqualizerX16Msmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX16MsmodeFir LspPlugInPluginsLv2ParaEqualizerX16Msmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX16MsmodeFft LspPlugInPluginsLv2ParaEqualizerX16Msmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm15X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm15X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm15X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm15X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss9X1 LspPlugInPluginsLv2ParaEqualizerX16Msss9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss9X2 LspPlugInPluginsLv2ParaEqualizerX16Msss9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss9X3 LspPlugInPluginsLv2ParaEqualizerX16Msss9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss9X4 LspPlugInPluginsLv2ParaEqualizerX16Msss9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm2ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm3ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13Off LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm13Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm9X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm9X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm9X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm9X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12Off LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm12Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3Off LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm3Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7Off LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts7Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm0ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms3ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms5ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10Off LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm10Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7Off LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm7Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9Off LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm9Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15Off LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts15Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6Off LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts6Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm4X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm4X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm4X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm4X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss1X1 LspPlugInPluginsLv2ParaEqualizerX16Msss1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss1X2 LspPlugInPluginsLv2ParaEqualizerX16Msss1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss1X3 LspPlugInPluginsLv2ParaEqualizerX16Msss1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss1X4 LspPlugInPluginsLv2ParaEqualizerX16Msss1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss6X1 LspPlugInPluginsLv2ParaEqualizerX16Msss6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss6X2 LspPlugInPluginsLv2ParaEqualizerX16Msss6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss6X3 LspPlugInPluginsLv2ParaEqualizerX16Msss6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss6X4 LspPlugInPluginsLv2ParaEqualizerX16Msss6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm12ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms13ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1Off LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm1Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1Off LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts1Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9Off LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts9Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm0X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm0X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm0X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm0X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss3X1 LspPlugInPluginsLv2ParaEqualizerX16Msss3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss3X2 LspPlugInPluginsLv2ParaEqualizerX16Msss3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss3X3 LspPlugInPluginsLv2ParaEqualizerX16Msss3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss3X4 LspPlugInPluginsLv2ParaEqualizerX16Msss3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms2ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0Off LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm0Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6Off LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm6Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11Off LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts11Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm6X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm6X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm6X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm6X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss0X1 LspPlugInPluginsLv2ParaEqualizerX16Msss0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss0X2 LspPlugInPluginsLv2ParaEqualizerX16Msss0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss0X3 LspPlugInPluginsLv2ParaEqualizerX16Msss0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss0X4 LspPlugInPluginsLv2ParaEqualizerX16Msss0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss2X1 LspPlugInPluginsLv2ParaEqualizerX16Msss2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss2X2 LspPlugInPluginsLv2ParaEqualizerX16Msss2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss2X3 LspPlugInPluginsLv2ParaEqualizerX16Msss2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss2X4 LspPlugInPluginsLv2ParaEqualizerX16Msss2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm7ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms15ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0Off LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts0Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10Off LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts10Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5Off LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts5Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm7X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm7X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm7X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm7X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss8X1 LspPlugInPluginsLv2ParaEqualizerX16Msss8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss8X2 LspPlugInPluginsLv2ParaEqualizerX16Msss8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss8X3 LspPlugInPluginsLv2ParaEqualizerX16Msss8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss8X4 LspPlugInPluginsLv2ParaEqualizerX16Msss8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm9ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2Off LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm2Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5Off LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm5Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4Off LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts4Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13Off LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts13Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14Off LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts14Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm11X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm11X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm11X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm11X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss13X1 LspPlugInPluginsLv2ParaEqualizerX16Msss13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss13X2 LspPlugInPluginsLv2ParaEqualizerX16Msss13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss13X3 LspPlugInPluginsLv2ParaEqualizerX16Msss13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss13X4 LspPlugInPluginsLv2ParaEqualizerX16Msss13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms1ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms4ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14Off LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm14Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12Off LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts12Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss15X1 LspPlugInPluginsLv2ParaEqualizerX16Msss15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss15X2 LspPlugInPluginsLv2ParaEqualizerX16Msss15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss15X3 LspPlugInPluginsLv2ParaEqualizerX16Msss15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss15X4 LspPlugInPluginsLv2ParaEqualizerX16Msss15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss7X1 LspPlugInPluginsLv2ParaEqualizerX16Msss7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss7X2 LspPlugInPluginsLv2ParaEqualizerX16Msss7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss7X3 LspPlugInPluginsLv2ParaEqualizerX16Msss7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss7X4 LspPlugInPluginsLv2ParaEqualizerX16Msss7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms14ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX16MsfselFiltersMiddle07 LspPlugInPluginsLv2ParaEqualizerX16Msfsel = "Filters Middle 0-7" // Filters Middle 0-7 (0) – Filters Middle 0-7
	LspPlugInPluginsLv2ParaEqualizerX16MsfselFiltersSide07 LspPlugInPluginsLv2ParaEqualizerX16Msfsel = "Filters Side 0-7" // Filters Side 0-7 (1) – Filters Side 0-7
	LspPlugInPluginsLv2ParaEqualizerX16MsfselFiltersMiddle815 LspPlugInPluginsLv2ParaEqualizerX16Msfsel = "Filters Middle 8-15" // Filters Middle 8-15 (2) – Filters Middle 8-15
	LspPlugInPluginsLv2ParaEqualizerX16MsfselFiltersSide815 LspPlugInPluginsLv2ParaEqualizerX16Msfsel = "Filters Side 8-15" // Filters Side 8-15 (3) – Filters Side 8-15
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms12ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15Off LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm15Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Msftm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4Off LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4Bell LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4HiPass LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4LoPass LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4Notch LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4Resonance LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msftm4Allpass LspPlugInPluginsLv2ParaEqualizerX16Msftm4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm1X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm1X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm1X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm1X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm14ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm4ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm5ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm6ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm12X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm12X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm12X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm12X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss10X1 LspPlugInPluginsLv2ParaEqualizerX16Msss10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss10X2 LspPlugInPluginsLv2ParaEqualizerX16Msss10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss10X3 LspPlugInPluginsLv2ParaEqualizerX16Msss10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss10X4 LspPlugInPluginsLv2ParaEqualizerX16Msss10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms11ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfts3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3Off LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3Bell LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3HiPass LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3HiShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3LoPass LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3LoShelf LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3Notch LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3Resonance LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Msfts3Allpass LspPlugInPluginsLv2ParaEqualizerX16Msfts3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm8X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm8X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm8X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm8X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss11X1 LspPlugInPluginsLv2ParaEqualizerX16Msss11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss11X2 LspPlugInPluginsLv2ParaEqualizerX16Msss11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss11X3 LspPlugInPluginsLv2ParaEqualizerX16Msss11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss11X4 LspPlugInPluginsLv2ParaEqualizerX16Msss11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss12X1 LspPlugInPluginsLv2ParaEqualizerX16Msss12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss12X2 LspPlugInPluginsLv2ParaEqualizerX16Msss12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss12X3 LspPlugInPluginsLv2ParaEqualizerX16Msss12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss12X4 LspPlugInPluginsLv2ParaEqualizerX16Msss12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss14X1 LspPlugInPluginsLv2ParaEqualizerX16Msss14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss14X2 LspPlugInPluginsLv2ParaEqualizerX16Msss14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss14X3 LspPlugInPluginsLv2ParaEqualizerX16Msss14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss14X4 LspPlugInPluginsLv2ParaEqualizerX16Msss14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms10ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms7ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm13X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm13X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm13X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm13X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Mssm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Mssm3X1 LspPlugInPluginsLv2ParaEqualizerX16Mssm3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Mssm3X2 LspPlugInPluginsLv2ParaEqualizerX16Mssm3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Mssm3X3 LspPlugInPluginsLv2ParaEqualizerX16Mssm3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Mssm3X4 LspPlugInPluginsLv2ParaEqualizerX16Mssm3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfmm1ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfmm1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msfms8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8RlcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8RlcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8BwcBt LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8BwcMt LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8LrxBt LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8LrxMt LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Msfms8ApoDr LspPlugInPluginsLv2ParaEqualizerX16Msfms8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Msss5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Msss5X1 LspPlugInPluginsLv2ParaEqualizerX16Msss5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Msss5X2 LspPlugInPluginsLv2ParaEqualizerX16Msss5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Msss5X3 LspPlugInPluginsLv2ParaEqualizerX16Msss5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Msss5X4 LspPlugInPluginsLv2ParaEqualizerX16Msss5 = "x4" // x4 (3) – x4
)

