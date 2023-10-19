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

type LspPlugInPluginsLv2ParaEqualizerX32Ms struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX32Ms() (*LspPlugInPluginsLv2ParaEqualizerX32Ms, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x32-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX32MsWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX32Ms, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x32-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX32Msfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-m (bool)
//
// FFT visibility Mid

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFftvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFftvM(value bool) error {
	return e.Element.SetProperty("fftv-m", value)
}

// fftv-s (bool)
//
// FFT visibility Side

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFftvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFftvS(value bool) error {
	return e.Element.SetProperty("fftv-s", value)
}

// fm-0 (float32)
//
// Frequency Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm0(value float32) error {
	return e.Element.SetProperty("fm-0", value)
}

// fm-1 (float32)
//
// Frequency Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm1(value float32) error {
	return e.Element.SetProperty("fm-1", value)
}

// fm-10 (float32)
//
// Frequency Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm10(value float32) error {
	return e.Element.SetProperty("fm-10", value)
}

// fm-11 (float32)
//
// Frequency Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm11(value float32) error {
	return e.Element.SetProperty("fm-11", value)
}

// fm-12 (float32)
//
// Frequency Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm12(value float32) error {
	return e.Element.SetProperty("fm-12", value)
}

// fm-13 (float32)
//
// Frequency Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm13(value float32) error {
	return e.Element.SetProperty("fm-13", value)
}

// fm-14 (float32)
//
// Frequency Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm14(value float32) error {
	return e.Element.SetProperty("fm-14", value)
}

// fm-15 (float32)
//
// Frequency Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm15(value float32) error {
	return e.Element.SetProperty("fm-15", value)
}

// fm-16 (float32)
//
// Frequency Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm16(value float32) error {
	return e.Element.SetProperty("fm-16", value)
}

// fm-17 (float32)
//
// Frequency Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm17(value float32) error {
	return e.Element.SetProperty("fm-17", value)
}

// fm-18 (float32)
//
// Frequency Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm18(value float32) error {
	return e.Element.SetProperty("fm-18", value)
}

// fm-19 (float32)
//
// Frequency Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm19(value float32) error {
	return e.Element.SetProperty("fm-19", value)
}

// fm-2 (float32)
//
// Frequency Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm2(value float32) error {
	return e.Element.SetProperty("fm-2", value)
}

// fm-20 (float32)
//
// Frequency Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm20(value float32) error {
	return e.Element.SetProperty("fm-20", value)
}

// fm-21 (float32)
//
// Frequency Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm21(value float32) error {
	return e.Element.SetProperty("fm-21", value)
}

// fm-22 (float32)
//
// Frequency Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm22(value float32) error {
	return e.Element.SetProperty("fm-22", value)
}

// fm-23 (float32)
//
// Frequency Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm23(value float32) error {
	return e.Element.SetProperty("fm-23", value)
}

// fm-24 (float32)
//
// Frequency Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm24(value float32) error {
	return e.Element.SetProperty("fm-24", value)
}

// fm-25 (float32)
//
// Frequency Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm25(value float32) error {
	return e.Element.SetProperty("fm-25", value)
}

// fm-26 (float32)
//
// Frequency Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm26(value float32) error {
	return e.Element.SetProperty("fm-26", value)
}

// fm-27 (float32)
//
// Frequency Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm27(value float32) error {
	return e.Element.SetProperty("fm-27", value)
}

// fm-28 (float32)
//
// Frequency Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm28(value float32) error {
	return e.Element.SetProperty("fm-28", value)
}

// fm-29 (float32)
//
// Frequency Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm29(value float32) error {
	return e.Element.SetProperty("fm-29", value)
}

// fm-3 (float32)
//
// Frequency Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm3(value float32) error {
	return e.Element.SetProperty("fm-3", value)
}

// fm-30 (float32)
//
// Frequency Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm30(value float32) error {
	return e.Element.SetProperty("fm-30", value)
}

// fm-31 (float32)
//
// Frequency Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm31(value float32) error {
	return e.Element.SetProperty("fm-31", value)
}

// fm-4 (float32)
//
// Frequency Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm4(value float32) error {
	return e.Element.SetProperty("fm-4", value)
}

// fm-5 (float32)
//
// Frequency Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm5(value float32) error {
	return e.Element.SetProperty("fm-5", value)
}

// fm-6 (float32)
//
// Frequency Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm6(value float32) error {
	return e.Element.SetProperty("fm-6", value)
}

// fm-7 (float32)
//
// Frequency Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm7(value float32) error {
	return e.Element.SetProperty("fm-7", value)
}

// fm-8 (float32)
//
// Frequency Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm8(value float32) error {
	return e.Element.SetProperty("fm-8", value)
}

// fm-9 (float32)
//
// Frequency Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFm9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFm9(value float32) error {
	return e.Element.SetProperty("fm-9", value)
}

// fmm-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm0)
//
// Filter mode Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm0() (interface{}, error) {
	return e.Element.GetProperty("fmm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm0(value interface{}) error {
	return e.Element.SetProperty("fmm-0", value)
}

// fmm-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm1)
//
// Filter mode Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm1() (interface{}, error) {
	return e.Element.GetProperty("fmm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm1(value interface{}) error {
	return e.Element.SetProperty("fmm-1", value)
}

// fmm-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm10)
//
// Filter mode Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm10() (interface{}, error) {
	return e.Element.GetProperty("fmm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm10(value interface{}) error {
	return e.Element.SetProperty("fmm-10", value)
}

// fmm-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm11)
//
// Filter mode Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm11() (interface{}, error) {
	return e.Element.GetProperty("fmm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm11(value interface{}) error {
	return e.Element.SetProperty("fmm-11", value)
}

// fmm-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm12)
//
// Filter mode Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm12() (interface{}, error) {
	return e.Element.GetProperty("fmm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm12(value interface{}) error {
	return e.Element.SetProperty("fmm-12", value)
}

// fmm-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm13)
//
// Filter mode Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm13() (interface{}, error) {
	return e.Element.GetProperty("fmm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm13(value interface{}) error {
	return e.Element.SetProperty("fmm-13", value)
}

// fmm-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm14)
//
// Filter mode Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm14() (interface{}, error) {
	return e.Element.GetProperty("fmm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm14(value interface{}) error {
	return e.Element.SetProperty("fmm-14", value)
}

// fmm-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm15)
//
// Filter mode Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm15() (interface{}, error) {
	return e.Element.GetProperty("fmm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm15(value interface{}) error {
	return e.Element.SetProperty("fmm-15", value)
}

// fmm-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm16)
//
// Filter mode Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm16() (interface{}, error) {
	return e.Element.GetProperty("fmm-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm16(value interface{}) error {
	return e.Element.SetProperty("fmm-16", value)
}

// fmm-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm17)
//
// Filter mode Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm17() (interface{}, error) {
	return e.Element.GetProperty("fmm-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm17(value interface{}) error {
	return e.Element.SetProperty("fmm-17", value)
}

// fmm-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm18)
//
// Filter mode Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm18() (interface{}, error) {
	return e.Element.GetProperty("fmm-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm18(value interface{}) error {
	return e.Element.SetProperty("fmm-18", value)
}

// fmm-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm19)
//
// Filter mode Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm19() (interface{}, error) {
	return e.Element.GetProperty("fmm-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm19(value interface{}) error {
	return e.Element.SetProperty("fmm-19", value)
}

// fmm-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm2)
//
// Filter mode Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm2() (interface{}, error) {
	return e.Element.GetProperty("fmm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm2(value interface{}) error {
	return e.Element.SetProperty("fmm-2", value)
}

// fmm-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm20)
//
// Filter mode Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm20() (interface{}, error) {
	return e.Element.GetProperty("fmm-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm20(value interface{}) error {
	return e.Element.SetProperty("fmm-20", value)
}

// fmm-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm21)
//
// Filter mode Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm21() (interface{}, error) {
	return e.Element.GetProperty("fmm-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm21(value interface{}) error {
	return e.Element.SetProperty("fmm-21", value)
}

// fmm-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm22)
//
// Filter mode Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm22() (interface{}, error) {
	return e.Element.GetProperty("fmm-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm22(value interface{}) error {
	return e.Element.SetProperty("fmm-22", value)
}

// fmm-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm23)
//
// Filter mode Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm23() (interface{}, error) {
	return e.Element.GetProperty("fmm-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm23(value interface{}) error {
	return e.Element.SetProperty("fmm-23", value)
}

// fmm-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm24)
//
// Filter mode Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm24() (interface{}, error) {
	return e.Element.GetProperty("fmm-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm24(value interface{}) error {
	return e.Element.SetProperty("fmm-24", value)
}

// fmm-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm25)
//
// Filter mode Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm25() (interface{}, error) {
	return e.Element.GetProperty("fmm-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm25(value interface{}) error {
	return e.Element.SetProperty("fmm-25", value)
}

// fmm-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm26)
//
// Filter mode Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm26() (interface{}, error) {
	return e.Element.GetProperty("fmm-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm26(value interface{}) error {
	return e.Element.SetProperty("fmm-26", value)
}

// fmm-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm27)
//
// Filter mode Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm27() (interface{}, error) {
	return e.Element.GetProperty("fmm-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm27(value interface{}) error {
	return e.Element.SetProperty("fmm-27", value)
}

// fmm-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm28)
//
// Filter mode Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm28() (interface{}, error) {
	return e.Element.GetProperty("fmm-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm28(value interface{}) error {
	return e.Element.SetProperty("fmm-28", value)
}

// fmm-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm29)
//
// Filter mode Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm29() (interface{}, error) {
	return e.Element.GetProperty("fmm-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm29(value interface{}) error {
	return e.Element.SetProperty("fmm-29", value)
}

// fmm-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm3)
//
// Filter mode Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm3() (interface{}, error) {
	return e.Element.GetProperty("fmm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm3(value interface{}) error {
	return e.Element.SetProperty("fmm-3", value)
}

// fmm-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm30)
//
// Filter mode Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm30() (interface{}, error) {
	return e.Element.GetProperty("fmm-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm30(value interface{}) error {
	return e.Element.SetProperty("fmm-30", value)
}

// fmm-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm31)
//
// Filter mode Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm31() (interface{}, error) {
	return e.Element.GetProperty("fmm-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm31(value interface{}) error {
	return e.Element.SetProperty("fmm-31", value)
}

// fmm-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm4)
//
// Filter mode Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm4() (interface{}, error) {
	return e.Element.GetProperty("fmm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm4(value interface{}) error {
	return e.Element.SetProperty("fmm-4", value)
}

// fmm-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm5)
//
// Filter mode Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm5() (interface{}, error) {
	return e.Element.GetProperty("fmm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm5(value interface{}) error {
	return e.Element.SetProperty("fmm-5", value)
}

// fmm-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm6)
//
// Filter mode Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm6() (interface{}, error) {
	return e.Element.GetProperty("fmm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm6(value interface{}) error {
	return e.Element.SetProperty("fmm-6", value)
}

// fmm-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm7)
//
// Filter mode Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm7() (interface{}, error) {
	return e.Element.GetProperty("fmm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm7(value interface{}) error {
	return e.Element.SetProperty("fmm-7", value)
}

// fmm-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm8)
//
// Filter mode Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm8() (interface{}, error) {
	return e.Element.GetProperty("fmm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm8(value interface{}) error {
	return e.Element.SetProperty("fmm-8", value)
}

// fmm-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfmm9)
//
// Filter mode Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFmm9() (interface{}, error) {
	return e.Element.GetProperty("fmm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFmm9(value interface{}) error {
	return e.Element.SetProperty("fmm-9", value)
}

// fms-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms0)
//
// Filter mode Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms0() (interface{}, error) {
	return e.Element.GetProperty("fms-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms0(value interface{}) error {
	return e.Element.SetProperty("fms-0", value)
}

// fms-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms1)
//
// Filter mode Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms1() (interface{}, error) {
	return e.Element.GetProperty("fms-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms1(value interface{}) error {
	return e.Element.SetProperty("fms-1", value)
}

// fms-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms10)
//
// Filter mode Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms10() (interface{}, error) {
	return e.Element.GetProperty("fms-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms10(value interface{}) error {
	return e.Element.SetProperty("fms-10", value)
}

// fms-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms11)
//
// Filter mode Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms11() (interface{}, error) {
	return e.Element.GetProperty("fms-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms11(value interface{}) error {
	return e.Element.SetProperty("fms-11", value)
}

// fms-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms12)
//
// Filter mode Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms12() (interface{}, error) {
	return e.Element.GetProperty("fms-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms12(value interface{}) error {
	return e.Element.SetProperty("fms-12", value)
}

// fms-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms13)
//
// Filter mode Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms13() (interface{}, error) {
	return e.Element.GetProperty("fms-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms13(value interface{}) error {
	return e.Element.SetProperty("fms-13", value)
}

// fms-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms14)
//
// Filter mode Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms14() (interface{}, error) {
	return e.Element.GetProperty("fms-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms14(value interface{}) error {
	return e.Element.SetProperty("fms-14", value)
}

// fms-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms15)
//
// Filter mode Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms15() (interface{}, error) {
	return e.Element.GetProperty("fms-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms15(value interface{}) error {
	return e.Element.SetProperty("fms-15", value)
}

// fms-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms16)
//
// Filter mode Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms16() (interface{}, error) {
	return e.Element.GetProperty("fms-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms16(value interface{}) error {
	return e.Element.SetProperty("fms-16", value)
}

// fms-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms17)
//
// Filter mode Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms17() (interface{}, error) {
	return e.Element.GetProperty("fms-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms17(value interface{}) error {
	return e.Element.SetProperty("fms-17", value)
}

// fms-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms18)
//
// Filter mode Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms18() (interface{}, error) {
	return e.Element.GetProperty("fms-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms18(value interface{}) error {
	return e.Element.SetProperty("fms-18", value)
}

// fms-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms19)
//
// Filter mode Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms19() (interface{}, error) {
	return e.Element.GetProperty("fms-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms19(value interface{}) error {
	return e.Element.SetProperty("fms-19", value)
}

// fms-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms2)
//
// Filter mode Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms2() (interface{}, error) {
	return e.Element.GetProperty("fms-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms2(value interface{}) error {
	return e.Element.SetProperty("fms-2", value)
}

// fms-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms20)
//
// Filter mode Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms20() (interface{}, error) {
	return e.Element.GetProperty("fms-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms20(value interface{}) error {
	return e.Element.SetProperty("fms-20", value)
}

// fms-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms21)
//
// Filter mode Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms21() (interface{}, error) {
	return e.Element.GetProperty("fms-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms21(value interface{}) error {
	return e.Element.SetProperty("fms-21", value)
}

// fms-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms22)
//
// Filter mode Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms22() (interface{}, error) {
	return e.Element.GetProperty("fms-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms22(value interface{}) error {
	return e.Element.SetProperty("fms-22", value)
}

// fms-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms23)
//
// Filter mode Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms23() (interface{}, error) {
	return e.Element.GetProperty("fms-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms23(value interface{}) error {
	return e.Element.SetProperty("fms-23", value)
}

// fms-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms24)
//
// Filter mode Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms24() (interface{}, error) {
	return e.Element.GetProperty("fms-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms24(value interface{}) error {
	return e.Element.SetProperty("fms-24", value)
}

// fms-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms25)
//
// Filter mode Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms25() (interface{}, error) {
	return e.Element.GetProperty("fms-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms25(value interface{}) error {
	return e.Element.SetProperty("fms-25", value)
}

// fms-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms26)
//
// Filter mode Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms26() (interface{}, error) {
	return e.Element.GetProperty("fms-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms26(value interface{}) error {
	return e.Element.SetProperty("fms-26", value)
}

// fms-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms27)
//
// Filter mode Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms27() (interface{}, error) {
	return e.Element.GetProperty("fms-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms27(value interface{}) error {
	return e.Element.SetProperty("fms-27", value)
}

// fms-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms28)
//
// Filter mode Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms28() (interface{}, error) {
	return e.Element.GetProperty("fms-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms28(value interface{}) error {
	return e.Element.SetProperty("fms-28", value)
}

// fms-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms29)
//
// Filter mode Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms29() (interface{}, error) {
	return e.Element.GetProperty("fms-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms29(value interface{}) error {
	return e.Element.SetProperty("fms-29", value)
}

// fms-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms3)
//
// Filter mode Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms3() (interface{}, error) {
	return e.Element.GetProperty("fms-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms3(value interface{}) error {
	return e.Element.SetProperty("fms-3", value)
}

// fms-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms30)
//
// Filter mode Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms30() (interface{}, error) {
	return e.Element.GetProperty("fms-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms30(value interface{}) error {
	return e.Element.SetProperty("fms-30", value)
}

// fms-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms31)
//
// Filter mode Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms31() (interface{}, error) {
	return e.Element.GetProperty("fms-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms31(value interface{}) error {
	return e.Element.SetProperty("fms-31", value)
}

// fms-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms4)
//
// Filter mode Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms4() (interface{}, error) {
	return e.Element.GetProperty("fms-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms4(value interface{}) error {
	return e.Element.SetProperty("fms-4", value)
}

// fms-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms5)
//
// Filter mode Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms5() (interface{}, error) {
	return e.Element.GetProperty("fms-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms5(value interface{}) error {
	return e.Element.SetProperty("fms-5", value)
}

// fms-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms6)
//
// Filter mode Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms6() (interface{}, error) {
	return e.Element.GetProperty("fms-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms6(value interface{}) error {
	return e.Element.SetProperty("fms-6", value)
}

// fms-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms7)
//
// Filter mode Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms7() (interface{}, error) {
	return e.Element.GetProperty("fms-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms7(value interface{}) error {
	return e.Element.SetProperty("fms-7", value)
}

// fms-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms8)
//
// Filter mode Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms8() (interface{}, error) {
	return e.Element.GetProperty("fms-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms8(value interface{}) error {
	return e.Element.SetProperty("fms-8", value)
}

// fms-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfms9)
//
// Filter mode Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFms9() (interface{}, error) {
	return e.Element.GetProperty("fms-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFms9(value interface{}) error {
	return e.Element.SetProperty("fms-9", value)
}

// frqs-m (float32)
//
// Frequency shift Mid

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFrqsM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFrqsM(value float32) error {
	return e.Element.SetProperty("frqs-m", value)
}

// frqs-s (float32)
//
// Frequency shift Side

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFrqsS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFrqsS(value float32) error {
	return e.Element.SetProperty("frqs-s", value)
}

// fs-0 (float32)
//
// Frequency Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs0(value float32) error {
	return e.Element.SetProperty("fs-0", value)
}

// fs-1 (float32)
//
// Frequency Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs1(value float32) error {
	return e.Element.SetProperty("fs-1", value)
}

// fs-10 (float32)
//
// Frequency Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs10(value float32) error {
	return e.Element.SetProperty("fs-10", value)
}

// fs-11 (float32)
//
// Frequency Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs11(value float32) error {
	return e.Element.SetProperty("fs-11", value)
}

// fs-12 (float32)
//
// Frequency Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs12(value float32) error {
	return e.Element.SetProperty("fs-12", value)
}

// fs-13 (float32)
//
// Frequency Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs13(value float32) error {
	return e.Element.SetProperty("fs-13", value)
}

// fs-14 (float32)
//
// Frequency Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs14(value float32) error {
	return e.Element.SetProperty("fs-14", value)
}

// fs-15 (float32)
//
// Frequency Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs15(value float32) error {
	return e.Element.SetProperty("fs-15", value)
}

// fs-16 (float32)
//
// Frequency Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs16(value float32) error {
	return e.Element.SetProperty("fs-16", value)
}

// fs-17 (float32)
//
// Frequency Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs17(value float32) error {
	return e.Element.SetProperty("fs-17", value)
}

// fs-18 (float32)
//
// Frequency Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs18(value float32) error {
	return e.Element.SetProperty("fs-18", value)
}

// fs-19 (float32)
//
// Frequency Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs19(value float32) error {
	return e.Element.SetProperty("fs-19", value)
}

// fs-2 (float32)
//
// Frequency Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs2(value float32) error {
	return e.Element.SetProperty("fs-2", value)
}

// fs-20 (float32)
//
// Frequency Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs20(value float32) error {
	return e.Element.SetProperty("fs-20", value)
}

// fs-21 (float32)
//
// Frequency Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs21(value float32) error {
	return e.Element.SetProperty("fs-21", value)
}

// fs-22 (float32)
//
// Frequency Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs22(value float32) error {
	return e.Element.SetProperty("fs-22", value)
}

// fs-23 (float32)
//
// Frequency Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs23(value float32) error {
	return e.Element.SetProperty("fs-23", value)
}

// fs-24 (float32)
//
// Frequency Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs24(value float32) error {
	return e.Element.SetProperty("fs-24", value)
}

// fs-25 (float32)
//
// Frequency Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs25(value float32) error {
	return e.Element.SetProperty("fs-25", value)
}

// fs-26 (float32)
//
// Frequency Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs26(value float32) error {
	return e.Element.SetProperty("fs-26", value)
}

// fs-27 (float32)
//
// Frequency Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs27(value float32) error {
	return e.Element.SetProperty("fs-27", value)
}

// fs-28 (float32)
//
// Frequency Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs28(value float32) error {
	return e.Element.SetProperty("fs-28", value)
}

// fs-29 (float32)
//
// Frequency Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs29(value float32) error {
	return e.Element.SetProperty("fs-29", value)
}

// fs-3 (float32)
//
// Frequency Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs3(value float32) error {
	return e.Element.SetProperty("fs-3", value)
}

// fs-30 (float32)
//
// Frequency Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs30(value float32) error {
	return e.Element.SetProperty("fs-30", value)
}

// fs-31 (float32)
//
// Frequency Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fs-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs31(value float32) error {
	return e.Element.SetProperty("fs-31", value)
}

// fs-4 (float32)
//
// Frequency Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs4(value float32) error {
	return e.Element.SetProperty("fs-4", value)
}

// fs-5 (float32)
//
// Frequency Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs5(value float32) error {
	return e.Element.SetProperty("fs-5", value)
}

// fs-6 (float32)
//
// Frequency Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs6(value float32) error {
	return e.Element.SetProperty("fs-6", value)
}

// fs-7 (float32)
//
// Frequency Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs7(value float32) error {
	return e.Element.SetProperty("fs-7", value)
}

// fs-8 (float32)
//
// Frequency Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs8(value float32) error {
	return e.Element.SetProperty("fs-8", value)
}

// fs-9 (float32)
//
// Frequency Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFs9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFs9(value float32) error {
	return e.Element.SetProperty("fs-9", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX32Msfsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ftm-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm0)
//
// Filter type Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm0() (interface{}, error) {
	return e.Element.GetProperty("ftm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm0(value interface{}) error {
	return e.Element.SetProperty("ftm-0", value)
}

// ftm-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm1)
//
// Filter type Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm1() (interface{}, error) {
	return e.Element.GetProperty("ftm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm1(value interface{}) error {
	return e.Element.SetProperty("ftm-1", value)
}

// ftm-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm10)
//
// Filter type Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm10() (interface{}, error) {
	return e.Element.GetProperty("ftm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm10(value interface{}) error {
	return e.Element.SetProperty("ftm-10", value)
}

// ftm-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm11)
//
// Filter type Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm11() (interface{}, error) {
	return e.Element.GetProperty("ftm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm11(value interface{}) error {
	return e.Element.SetProperty("ftm-11", value)
}

// ftm-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm12)
//
// Filter type Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm12() (interface{}, error) {
	return e.Element.GetProperty("ftm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm12(value interface{}) error {
	return e.Element.SetProperty("ftm-12", value)
}

// ftm-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm13)
//
// Filter type Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm13() (interface{}, error) {
	return e.Element.GetProperty("ftm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm13(value interface{}) error {
	return e.Element.SetProperty("ftm-13", value)
}

// ftm-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm14)
//
// Filter type Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm14() (interface{}, error) {
	return e.Element.GetProperty("ftm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm14(value interface{}) error {
	return e.Element.SetProperty("ftm-14", value)
}

// ftm-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm15)
//
// Filter type Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm15() (interface{}, error) {
	return e.Element.GetProperty("ftm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm15(value interface{}) error {
	return e.Element.SetProperty("ftm-15", value)
}

// ftm-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm16)
//
// Filter type Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm16() (interface{}, error) {
	return e.Element.GetProperty("ftm-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm16(value interface{}) error {
	return e.Element.SetProperty("ftm-16", value)
}

// ftm-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm17)
//
// Filter type Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm17() (interface{}, error) {
	return e.Element.GetProperty("ftm-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm17(value interface{}) error {
	return e.Element.SetProperty("ftm-17", value)
}

// ftm-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm18)
//
// Filter type Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm18() (interface{}, error) {
	return e.Element.GetProperty("ftm-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm18(value interface{}) error {
	return e.Element.SetProperty("ftm-18", value)
}

// ftm-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm19)
//
// Filter type Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm19() (interface{}, error) {
	return e.Element.GetProperty("ftm-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm19(value interface{}) error {
	return e.Element.SetProperty("ftm-19", value)
}

// ftm-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm2)
//
// Filter type Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm2() (interface{}, error) {
	return e.Element.GetProperty("ftm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm2(value interface{}) error {
	return e.Element.SetProperty("ftm-2", value)
}

// ftm-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm20)
//
// Filter type Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm20() (interface{}, error) {
	return e.Element.GetProperty("ftm-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm20(value interface{}) error {
	return e.Element.SetProperty("ftm-20", value)
}

// ftm-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm21)
//
// Filter type Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm21() (interface{}, error) {
	return e.Element.GetProperty("ftm-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm21(value interface{}) error {
	return e.Element.SetProperty("ftm-21", value)
}

// ftm-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm22)
//
// Filter type Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm22() (interface{}, error) {
	return e.Element.GetProperty("ftm-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm22(value interface{}) error {
	return e.Element.SetProperty("ftm-22", value)
}

// ftm-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm23)
//
// Filter type Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm23() (interface{}, error) {
	return e.Element.GetProperty("ftm-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm23(value interface{}) error {
	return e.Element.SetProperty("ftm-23", value)
}

// ftm-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm24)
//
// Filter type Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm24() (interface{}, error) {
	return e.Element.GetProperty("ftm-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm24(value interface{}) error {
	return e.Element.SetProperty("ftm-24", value)
}

// ftm-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm25)
//
// Filter type Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm25() (interface{}, error) {
	return e.Element.GetProperty("ftm-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm25(value interface{}) error {
	return e.Element.SetProperty("ftm-25", value)
}

// ftm-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm26)
//
// Filter type Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm26() (interface{}, error) {
	return e.Element.GetProperty("ftm-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm26(value interface{}) error {
	return e.Element.SetProperty("ftm-26", value)
}

// ftm-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm27)
//
// Filter type Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm27() (interface{}, error) {
	return e.Element.GetProperty("ftm-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm27(value interface{}) error {
	return e.Element.SetProperty("ftm-27", value)
}

// ftm-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm28)
//
// Filter type Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm28() (interface{}, error) {
	return e.Element.GetProperty("ftm-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm28(value interface{}) error {
	return e.Element.SetProperty("ftm-28", value)
}

// ftm-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm29)
//
// Filter type Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm29() (interface{}, error) {
	return e.Element.GetProperty("ftm-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm29(value interface{}) error {
	return e.Element.SetProperty("ftm-29", value)
}

// ftm-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm3)
//
// Filter type Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm3() (interface{}, error) {
	return e.Element.GetProperty("ftm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm3(value interface{}) error {
	return e.Element.SetProperty("ftm-3", value)
}

// ftm-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm30)
//
// Filter type Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm30() (interface{}, error) {
	return e.Element.GetProperty("ftm-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm30(value interface{}) error {
	return e.Element.SetProperty("ftm-30", value)
}

// ftm-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm31)
//
// Filter type Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm31() (interface{}, error) {
	return e.Element.GetProperty("ftm-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm31(value interface{}) error {
	return e.Element.SetProperty("ftm-31", value)
}

// ftm-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm4)
//
// Filter type Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm4() (interface{}, error) {
	return e.Element.GetProperty("ftm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm4(value interface{}) error {
	return e.Element.SetProperty("ftm-4", value)
}

// ftm-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm5)
//
// Filter type Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm5() (interface{}, error) {
	return e.Element.GetProperty("ftm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm5(value interface{}) error {
	return e.Element.SetProperty("ftm-5", value)
}

// ftm-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm6)
//
// Filter type Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm6() (interface{}, error) {
	return e.Element.GetProperty("ftm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm6(value interface{}) error {
	return e.Element.SetProperty("ftm-6", value)
}

// ftm-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm7)
//
// Filter type Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm7() (interface{}, error) {
	return e.Element.GetProperty("ftm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm7(value interface{}) error {
	return e.Element.SetProperty("ftm-7", value)
}

// ftm-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm8)
//
// Filter type Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm8() (interface{}, error) {
	return e.Element.GetProperty("ftm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm8(value interface{}) error {
	return e.Element.SetProperty("ftm-8", value)
}

// ftm-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Msftm9)
//
// Filter type Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFtm9() (interface{}, error) {
	return e.Element.GetProperty("ftm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFtm9(value interface{}) error {
	return e.Element.SetProperty("ftm-9", value)
}

// fts-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts0)
//
// Filter type Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts0() (interface{}, error) {
	return e.Element.GetProperty("fts-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts0(value interface{}) error {
	return e.Element.SetProperty("fts-0", value)
}

// fts-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts1)
//
// Filter type Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts1() (interface{}, error) {
	return e.Element.GetProperty("fts-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts1(value interface{}) error {
	return e.Element.SetProperty("fts-1", value)
}

// fts-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts10)
//
// Filter type Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts10() (interface{}, error) {
	return e.Element.GetProperty("fts-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts10(value interface{}) error {
	return e.Element.SetProperty("fts-10", value)
}

// fts-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts11)
//
// Filter type Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts11() (interface{}, error) {
	return e.Element.GetProperty("fts-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts11(value interface{}) error {
	return e.Element.SetProperty("fts-11", value)
}

// fts-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts12)
//
// Filter type Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts12() (interface{}, error) {
	return e.Element.GetProperty("fts-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts12(value interface{}) error {
	return e.Element.SetProperty("fts-12", value)
}

// fts-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts13)
//
// Filter type Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts13() (interface{}, error) {
	return e.Element.GetProperty("fts-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts13(value interface{}) error {
	return e.Element.SetProperty("fts-13", value)
}

// fts-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts14)
//
// Filter type Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts14() (interface{}, error) {
	return e.Element.GetProperty("fts-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts14(value interface{}) error {
	return e.Element.SetProperty("fts-14", value)
}

// fts-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts15)
//
// Filter type Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts15() (interface{}, error) {
	return e.Element.GetProperty("fts-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts15(value interface{}) error {
	return e.Element.SetProperty("fts-15", value)
}

// fts-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts16)
//
// Filter type Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts16() (interface{}, error) {
	return e.Element.GetProperty("fts-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts16(value interface{}) error {
	return e.Element.SetProperty("fts-16", value)
}

// fts-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts17)
//
// Filter type Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts17() (interface{}, error) {
	return e.Element.GetProperty("fts-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts17(value interface{}) error {
	return e.Element.SetProperty("fts-17", value)
}

// fts-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts18)
//
// Filter type Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts18() (interface{}, error) {
	return e.Element.GetProperty("fts-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts18(value interface{}) error {
	return e.Element.SetProperty("fts-18", value)
}

// fts-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts19)
//
// Filter type Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts19() (interface{}, error) {
	return e.Element.GetProperty("fts-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts19(value interface{}) error {
	return e.Element.SetProperty("fts-19", value)
}

// fts-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts2)
//
// Filter type Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts2() (interface{}, error) {
	return e.Element.GetProperty("fts-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts2(value interface{}) error {
	return e.Element.SetProperty("fts-2", value)
}

// fts-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts20)
//
// Filter type Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts20() (interface{}, error) {
	return e.Element.GetProperty("fts-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts20(value interface{}) error {
	return e.Element.SetProperty("fts-20", value)
}

// fts-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts21)
//
// Filter type Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts21() (interface{}, error) {
	return e.Element.GetProperty("fts-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts21(value interface{}) error {
	return e.Element.SetProperty("fts-21", value)
}

// fts-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts22)
//
// Filter type Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts22() (interface{}, error) {
	return e.Element.GetProperty("fts-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts22(value interface{}) error {
	return e.Element.SetProperty("fts-22", value)
}

// fts-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts23)
//
// Filter type Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts23() (interface{}, error) {
	return e.Element.GetProperty("fts-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts23(value interface{}) error {
	return e.Element.SetProperty("fts-23", value)
}

// fts-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts24)
//
// Filter type Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts24() (interface{}, error) {
	return e.Element.GetProperty("fts-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts24(value interface{}) error {
	return e.Element.SetProperty("fts-24", value)
}

// fts-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts25)
//
// Filter type Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts25() (interface{}, error) {
	return e.Element.GetProperty("fts-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts25(value interface{}) error {
	return e.Element.SetProperty("fts-25", value)
}

// fts-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts26)
//
// Filter type Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts26() (interface{}, error) {
	return e.Element.GetProperty("fts-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts26(value interface{}) error {
	return e.Element.SetProperty("fts-26", value)
}

// fts-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts27)
//
// Filter type Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts27() (interface{}, error) {
	return e.Element.GetProperty("fts-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts27(value interface{}) error {
	return e.Element.SetProperty("fts-27", value)
}

// fts-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts28)
//
// Filter type Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts28() (interface{}, error) {
	return e.Element.GetProperty("fts-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts28(value interface{}) error {
	return e.Element.SetProperty("fts-28", value)
}

// fts-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts29)
//
// Filter type Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts29() (interface{}, error) {
	return e.Element.GetProperty("fts-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts29(value interface{}) error {
	return e.Element.SetProperty("fts-29", value)
}

// fts-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts3)
//
// Filter type Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts3() (interface{}, error) {
	return e.Element.GetProperty("fts-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts3(value interface{}) error {
	return e.Element.SetProperty("fts-3", value)
}

// fts-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts30)
//
// Filter type Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts30() (interface{}, error) {
	return e.Element.GetProperty("fts-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts30(value interface{}) error {
	return e.Element.SetProperty("fts-30", value)
}

// fts-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts31)
//
// Filter type Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts31() (interface{}, error) {
	return e.Element.GetProperty("fts-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts31(value interface{}) error {
	return e.Element.SetProperty("fts-31", value)
}

// fts-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts4)
//
// Filter type Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts4() (interface{}, error) {
	return e.Element.GetProperty("fts-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts4(value interface{}) error {
	return e.Element.SetProperty("fts-4", value)
}

// fts-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts5)
//
// Filter type Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts5() (interface{}, error) {
	return e.Element.GetProperty("fts-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts5(value interface{}) error {
	return e.Element.SetProperty("fts-5", value)
}

// fts-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts6)
//
// Filter type Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts6() (interface{}, error) {
	return e.Element.GetProperty("fts-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts6(value interface{}) error {
	return e.Element.SetProperty("fts-6", value)
}

// fts-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts7)
//
// Filter type Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts7() (interface{}, error) {
	return e.Element.GetProperty("fts-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts7(value interface{}) error {
	return e.Element.SetProperty("fts-7", value)
}

// fts-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts8)
//
// Filter type Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts8() (interface{}, error) {
	return e.Element.GetProperty("fts-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts8(value interface{}) error {
	return e.Element.SetProperty("fts-8", value)
}

// fts-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Msfts9)
//
// Filter type Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFts9() (interface{}, error) {
	return e.Element.GetProperty("fts-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetFts9(value interface{}) error {
	return e.Element.SetProperty("fts-9", value)
}

// fvm-0 (bool)
//
// Filter visibility Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm15() (bool, error) {
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

// fvm-16 (bool)
//
// Filter visibility Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-17 (bool)
//
// Filter visibility Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-18 (bool)
//
// Filter visibility Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-19 (bool)
//
// Filter visibility Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-19")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm2() (bool, error) {
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

// fvm-20 (bool)
//
// Filter visibility Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-21 (bool)
//
// Filter visibility Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-22 (bool)
//
// Filter visibility Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-23 (bool)
//
// Filter visibility Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-24 (bool)
//
// Filter visibility Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-25 (bool)
//
// Filter visibility Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-26 (bool)
//
// Filter visibility Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-27 (bool)
//
// Filter visibility Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-28 (bool)
//
// Filter visibility Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-29 (bool)
//
// Filter visibility Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-29")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm3() (bool, error) {
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

// fvm-30 (bool)
//
// Filter visibility Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvm-31 (bool)
//
// Filter visibility Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvm-31")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs15() (bool, error) {
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

// fvs-16 (bool)
//
// Filter visibility Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-17 (bool)
//
// Filter visibility Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-18 (bool)
//
// Filter visibility Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-19 (bool)
//
// Filter visibility Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-19")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs2() (bool, error) {
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

// fvs-20 (bool)
//
// Filter visibility Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-21 (bool)
//
// Filter visibility Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-22 (bool)
//
// Filter visibility Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-23 (bool)
//
// Filter visibility Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-24 (bool)
//
// Filter visibility Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-25 (bool)
//
// Filter visibility Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-26 (bool)
//
// Filter visibility Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-27 (bool)
//
// Filter visibility Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-28 (bool)
//
// Filter visibility Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-29 (bool)
//
// Filter visibility Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-29")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs3() (bool, error) {
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

// fvs-30 (bool)
//
// Filter visibility Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvs-31 (bool)
//
// Filter visibility Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvs-31")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetFvs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gain-m (float32)
//
// Mid gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGainM() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGainM(value float32) error {
	return e.Element.SetProperty("gain-m", value)
}

// gain-s (float32)
//
// Side gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGainS() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGainS(value float32) error {
	return e.Element.SetProperty("gain-s", value)
}

// gm-0 (float32)
//
// Gain Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm0(value float32) error {
	return e.Element.SetProperty("gm-0", value)
}

// gm-1 (float32)
//
// Gain Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm1(value float32) error {
	return e.Element.SetProperty("gm-1", value)
}

// gm-10 (float32)
//
// Gain Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm10(value float32) error {
	return e.Element.SetProperty("gm-10", value)
}

// gm-11 (float32)
//
// Gain Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm11(value float32) error {
	return e.Element.SetProperty("gm-11", value)
}

// gm-12 (float32)
//
// Gain Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm12(value float32) error {
	return e.Element.SetProperty("gm-12", value)
}

// gm-13 (float32)
//
// Gain Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm13(value float32) error {
	return e.Element.SetProperty("gm-13", value)
}

// gm-14 (float32)
//
// Gain Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm14(value float32) error {
	return e.Element.SetProperty("gm-14", value)
}

// gm-15 (float32)
//
// Gain Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm15(value float32) error {
	return e.Element.SetProperty("gm-15", value)
}

// gm-16 (float32)
//
// Gain Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm16(value float32) error {
	return e.Element.SetProperty("gm-16", value)
}

// gm-17 (float32)
//
// Gain Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm17(value float32) error {
	return e.Element.SetProperty("gm-17", value)
}

// gm-18 (float32)
//
// Gain Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm18(value float32) error {
	return e.Element.SetProperty("gm-18", value)
}

// gm-19 (float32)
//
// Gain Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm19(value float32) error {
	return e.Element.SetProperty("gm-19", value)
}

// gm-2 (float32)
//
// Gain Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm2(value float32) error {
	return e.Element.SetProperty("gm-2", value)
}

// gm-20 (float32)
//
// Gain Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm20(value float32) error {
	return e.Element.SetProperty("gm-20", value)
}

// gm-21 (float32)
//
// Gain Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm21(value float32) error {
	return e.Element.SetProperty("gm-21", value)
}

// gm-22 (float32)
//
// Gain Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm22(value float32) error {
	return e.Element.SetProperty("gm-22", value)
}

// gm-23 (float32)
//
// Gain Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm23(value float32) error {
	return e.Element.SetProperty("gm-23", value)
}

// gm-24 (float32)
//
// Gain Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm24(value float32) error {
	return e.Element.SetProperty("gm-24", value)
}

// gm-25 (float32)
//
// Gain Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm25(value float32) error {
	return e.Element.SetProperty("gm-25", value)
}

// gm-26 (float32)
//
// Gain Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm26(value float32) error {
	return e.Element.SetProperty("gm-26", value)
}

// gm-27 (float32)
//
// Gain Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm27(value float32) error {
	return e.Element.SetProperty("gm-27", value)
}

// gm-28 (float32)
//
// Gain Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm28(value float32) error {
	return e.Element.SetProperty("gm-28", value)
}

// gm-29 (float32)
//
// Gain Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm29(value float32) error {
	return e.Element.SetProperty("gm-29", value)
}

// gm-3 (float32)
//
// Gain Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm3(value float32) error {
	return e.Element.SetProperty("gm-3", value)
}

// gm-30 (float32)
//
// Gain Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm30(value float32) error {
	return e.Element.SetProperty("gm-30", value)
}

// gm-31 (float32)
//
// Gain Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm31(value float32) error {
	return e.Element.SetProperty("gm-31", value)
}

// gm-4 (float32)
//
// Gain Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm4(value float32) error {
	return e.Element.SetProperty("gm-4", value)
}

// gm-5 (float32)
//
// Gain Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm5(value float32) error {
	return e.Element.SetProperty("gm-5", value)
}

// gm-6 (float32)
//
// Gain Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm6(value float32) error {
	return e.Element.SetProperty("gm-6", value)
}

// gm-7 (float32)
//
// Gain Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm7(value float32) error {
	return e.Element.SetProperty("gm-7", value)
}

// gm-8 (float32)
//
// Gain Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm8(value float32) error {
	return e.Element.SetProperty("gm-8", value)
}

// gm-9 (float32)
//
// Gain Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGm9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGm9(value float32) error {
	return e.Element.SetProperty("gm-9", value)
}

// gs-0 (float32)
//
// Gain Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs0(value float32) error {
	return e.Element.SetProperty("gs-0", value)
}

// gs-1 (float32)
//
// Gain Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs1(value float32) error {
	return e.Element.SetProperty("gs-1", value)
}

// gs-10 (float32)
//
// Gain Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs10(value float32) error {
	return e.Element.SetProperty("gs-10", value)
}

// gs-11 (float32)
//
// Gain Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs11(value float32) error {
	return e.Element.SetProperty("gs-11", value)
}

// gs-12 (float32)
//
// Gain Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs12(value float32) error {
	return e.Element.SetProperty("gs-12", value)
}

// gs-13 (float32)
//
// Gain Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs13(value float32) error {
	return e.Element.SetProperty("gs-13", value)
}

// gs-14 (float32)
//
// Gain Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs14(value float32) error {
	return e.Element.SetProperty("gs-14", value)
}

// gs-15 (float32)
//
// Gain Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs15(value float32) error {
	return e.Element.SetProperty("gs-15", value)
}

// gs-16 (float32)
//
// Gain Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs16(value float32) error {
	return e.Element.SetProperty("gs-16", value)
}

// gs-17 (float32)
//
// Gain Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs17(value float32) error {
	return e.Element.SetProperty("gs-17", value)
}

// gs-18 (float32)
//
// Gain Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs18(value float32) error {
	return e.Element.SetProperty("gs-18", value)
}

// gs-19 (float32)
//
// Gain Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs19(value float32) error {
	return e.Element.SetProperty("gs-19", value)
}

// gs-2 (float32)
//
// Gain Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs2(value float32) error {
	return e.Element.SetProperty("gs-2", value)
}

// gs-20 (float32)
//
// Gain Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs20(value float32) error {
	return e.Element.SetProperty("gs-20", value)
}

// gs-21 (float32)
//
// Gain Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs21(value float32) error {
	return e.Element.SetProperty("gs-21", value)
}

// gs-22 (float32)
//
// Gain Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs22(value float32) error {
	return e.Element.SetProperty("gs-22", value)
}

// gs-23 (float32)
//
// Gain Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs23(value float32) error {
	return e.Element.SetProperty("gs-23", value)
}

// gs-24 (float32)
//
// Gain Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs24(value float32) error {
	return e.Element.SetProperty("gs-24", value)
}

// gs-25 (float32)
//
// Gain Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs25(value float32) error {
	return e.Element.SetProperty("gs-25", value)
}

// gs-26 (float32)
//
// Gain Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs26(value float32) error {
	return e.Element.SetProperty("gs-26", value)
}

// gs-27 (float32)
//
// Gain Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs27(value float32) error {
	return e.Element.SetProperty("gs-27", value)
}

// gs-28 (float32)
//
// Gain Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs28(value float32) error {
	return e.Element.SetProperty("gs-28", value)
}

// gs-29 (float32)
//
// Gain Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs29(value float32) error {
	return e.Element.SetProperty("gs-29", value)
}

// gs-3 (float32)
//
// Gain Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs3(value float32) error {
	return e.Element.SetProperty("gs-3", value)
}

// gs-30 (float32)
//
// Gain Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs30(value float32) error {
	return e.Element.SetProperty("gs-30", value)
}

// gs-31 (float32)
//
// Gain Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gs-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs31(value float32) error {
	return e.Element.SetProperty("gs-31", value)
}

// gs-4 (float32)
//
// Gain Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs4(value float32) error {
	return e.Element.SetProperty("gs-4", value)
}

// gs-5 (float32)
//
// Gain Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs5(value float32) error {
	return e.Element.SetProperty("gs-5", value)
}

// gs-6 (float32)
//
// Gain Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs6(value float32) error {
	return e.Element.SetProperty("gs-6", value)
}

// gs-7 (float32)
//
// Gain Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs7(value float32) error {
	return e.Element.SetProperty("gs-7", value)
}

// gs-8 (float32)
//
// Gain Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs8(value float32) error {
	return e.Element.SetProperty("gs-8", value)
}

// gs-9 (float32)
//
// Gain Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetGs9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetGs9(value float32) error {
	return e.Element.SetProperty("gs-9", value)
}

// huem-0 (float32)
//
// Hue Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem0(value float32) error {
	return e.Element.SetProperty("huem-0", value)
}

// huem-1 (float32)
//
// Hue Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem1(value float32) error {
	return e.Element.SetProperty("huem-1", value)
}

// huem-10 (float32)
//
// Hue Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem10(value float32) error {
	return e.Element.SetProperty("huem-10", value)
}

// huem-11 (float32)
//
// Hue Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem11(value float32) error {
	return e.Element.SetProperty("huem-11", value)
}

// huem-12 (float32)
//
// Hue Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem12(value float32) error {
	return e.Element.SetProperty("huem-12", value)
}

// huem-13 (float32)
//
// Hue Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem13(value float32) error {
	return e.Element.SetProperty("huem-13", value)
}

// huem-14 (float32)
//
// Hue Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem14(value float32) error {
	return e.Element.SetProperty("huem-14", value)
}

// huem-15 (float32)
//
// Hue Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem15(value float32) error {
	return e.Element.SetProperty("huem-15", value)
}

// huem-16 (float32)
//
// Hue Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem16(value float32) error {
	return e.Element.SetProperty("huem-16", value)
}

// huem-17 (float32)
//
// Hue Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem17(value float32) error {
	return e.Element.SetProperty("huem-17", value)
}

// huem-18 (float32)
//
// Hue Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem18(value float32) error {
	return e.Element.SetProperty("huem-18", value)
}

// huem-19 (float32)
//
// Hue Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem19(value float32) error {
	return e.Element.SetProperty("huem-19", value)
}

// huem-2 (float32)
//
// Hue Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem2(value float32) error {
	return e.Element.SetProperty("huem-2", value)
}

// huem-20 (float32)
//
// Hue Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem20(value float32) error {
	return e.Element.SetProperty("huem-20", value)
}

// huem-21 (float32)
//
// Hue Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem21(value float32) error {
	return e.Element.SetProperty("huem-21", value)
}

// huem-22 (float32)
//
// Hue Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem22(value float32) error {
	return e.Element.SetProperty("huem-22", value)
}

// huem-23 (float32)
//
// Hue Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem23(value float32) error {
	return e.Element.SetProperty("huem-23", value)
}

// huem-24 (float32)
//
// Hue Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem24(value float32) error {
	return e.Element.SetProperty("huem-24", value)
}

// huem-25 (float32)
//
// Hue Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem25(value float32) error {
	return e.Element.SetProperty("huem-25", value)
}

// huem-26 (float32)
//
// Hue Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem26(value float32) error {
	return e.Element.SetProperty("huem-26", value)
}

// huem-27 (float32)
//
// Hue Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem27(value float32) error {
	return e.Element.SetProperty("huem-27", value)
}

// huem-28 (float32)
//
// Hue Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem28(value float32) error {
	return e.Element.SetProperty("huem-28", value)
}

// huem-29 (float32)
//
// Hue Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem29(value float32) error {
	return e.Element.SetProperty("huem-29", value)
}

// huem-3 (float32)
//
// Hue Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem3(value float32) error {
	return e.Element.SetProperty("huem-3", value)
}

// huem-30 (float32)
//
// Hue Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem30(value float32) error {
	return e.Element.SetProperty("huem-30", value)
}

// huem-31 (float32)
//
// Hue Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huem-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem31(value float32) error {
	return e.Element.SetProperty("huem-31", value)
}

// huem-4 (float32)
//
// Hue Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem4(value float32) error {
	return e.Element.SetProperty("huem-4", value)
}

// huem-5 (float32)
//
// Hue Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem5(value float32) error {
	return e.Element.SetProperty("huem-5", value)
}

// huem-6 (float32)
//
// Hue Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem6(value float32) error {
	return e.Element.SetProperty("huem-6", value)
}

// huem-7 (float32)
//
// Hue Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem7(value float32) error {
	return e.Element.SetProperty("huem-7", value)
}

// huem-8 (float32)
//
// Hue Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem8(value float32) error {
	return e.Element.SetProperty("huem-8", value)
}

// huem-9 (float32)
//
// Hue Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHuem9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHuem9(value float32) error {
	return e.Element.SetProperty("huem-9", value)
}

// hues-0 (float32)
//
// Hue Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues0(value float32) error {
	return e.Element.SetProperty("hues-0", value)
}

// hues-1 (float32)
//
// Hue Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues1(value float32) error {
	return e.Element.SetProperty("hues-1", value)
}

// hues-10 (float32)
//
// Hue Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues10(value float32) error {
	return e.Element.SetProperty("hues-10", value)
}

// hues-11 (float32)
//
// Hue Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues11(value float32) error {
	return e.Element.SetProperty("hues-11", value)
}

// hues-12 (float32)
//
// Hue Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues12(value float32) error {
	return e.Element.SetProperty("hues-12", value)
}

// hues-13 (float32)
//
// Hue Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues13(value float32) error {
	return e.Element.SetProperty("hues-13", value)
}

// hues-14 (float32)
//
// Hue Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues14(value float32) error {
	return e.Element.SetProperty("hues-14", value)
}

// hues-15 (float32)
//
// Hue Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues15(value float32) error {
	return e.Element.SetProperty("hues-15", value)
}

// hues-16 (float32)
//
// Hue Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues16(value float32) error {
	return e.Element.SetProperty("hues-16", value)
}

// hues-17 (float32)
//
// Hue Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues17(value float32) error {
	return e.Element.SetProperty("hues-17", value)
}

// hues-18 (float32)
//
// Hue Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues18(value float32) error {
	return e.Element.SetProperty("hues-18", value)
}

// hues-19 (float32)
//
// Hue Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues19(value float32) error {
	return e.Element.SetProperty("hues-19", value)
}

// hues-2 (float32)
//
// Hue Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues2(value float32) error {
	return e.Element.SetProperty("hues-2", value)
}

// hues-20 (float32)
//
// Hue Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues20(value float32) error {
	return e.Element.SetProperty("hues-20", value)
}

// hues-21 (float32)
//
// Hue Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues21(value float32) error {
	return e.Element.SetProperty("hues-21", value)
}

// hues-22 (float32)
//
// Hue Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues22(value float32) error {
	return e.Element.SetProperty("hues-22", value)
}

// hues-23 (float32)
//
// Hue Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues23(value float32) error {
	return e.Element.SetProperty("hues-23", value)
}

// hues-24 (float32)
//
// Hue Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues24(value float32) error {
	return e.Element.SetProperty("hues-24", value)
}

// hues-25 (float32)
//
// Hue Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues25(value float32) error {
	return e.Element.SetProperty("hues-25", value)
}

// hues-26 (float32)
//
// Hue Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues26(value float32) error {
	return e.Element.SetProperty("hues-26", value)
}

// hues-27 (float32)
//
// Hue Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues27(value float32) error {
	return e.Element.SetProperty("hues-27", value)
}

// hues-28 (float32)
//
// Hue Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues28(value float32) error {
	return e.Element.SetProperty("hues-28", value)
}

// hues-29 (float32)
//
// Hue Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues29(value float32) error {
	return e.Element.SetProperty("hues-29", value)
}

// hues-3 (float32)
//
// Hue Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues3(value float32) error {
	return e.Element.SetProperty("hues-3", value)
}

// hues-30 (float32)
//
// Hue Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues30(value float32) error {
	return e.Element.SetProperty("hues-30", value)
}

// hues-31 (float32)
//
// Hue Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hues-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues31(value float32) error {
	return e.Element.SetProperty("hues-31", value)
}

// hues-4 (float32)
//
// Hue Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues4(value float32) error {
	return e.Element.SetProperty("hues-4", value)
}

// hues-5 (float32)
//
// Hue Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues5(value float32) error {
	return e.Element.SetProperty("hues-5", value)
}

// hues-6 (float32)
//
// Hue Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues6(value float32) error {
	return e.Element.SetProperty("hues-6", value)
}

// hues-7 (float32)
//
// Hue Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues7(value float32) error {
	return e.Element.SetProperty("hues-7", value)
}

// hues-8 (float32)
//
// Hue Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues8(value float32) error {
	return e.Element.SetProperty("hues-8", value)
}

// hues-9 (float32)
//
// Hue Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetHues9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetHues9(value float32) error {
	return e.Element.SetProperty("hues-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetImr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetLstn() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mode (GstLspPlugInPluginsLv2ParaEqualizerX32Msmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm0(value float32) error {
	return e.Element.SetProperty("qm-0", value)
}

// qm-1 (float32)
//
// Quality factor Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm1(value float32) error {
	return e.Element.SetProperty("qm-1", value)
}

// qm-10 (float32)
//
// Quality factor Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm10(value float32) error {
	return e.Element.SetProperty("qm-10", value)
}

// qm-11 (float32)
//
// Quality factor Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm11(value float32) error {
	return e.Element.SetProperty("qm-11", value)
}

// qm-12 (float32)
//
// Quality factor Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm12(value float32) error {
	return e.Element.SetProperty("qm-12", value)
}

// qm-13 (float32)
//
// Quality factor Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm13(value float32) error {
	return e.Element.SetProperty("qm-13", value)
}

// qm-14 (float32)
//
// Quality factor Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm14(value float32) error {
	return e.Element.SetProperty("qm-14", value)
}

// qm-15 (float32)
//
// Quality factor Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm15(value float32) error {
	return e.Element.SetProperty("qm-15", value)
}

// qm-16 (float32)
//
// Quality factor Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm16(value float32) error {
	return e.Element.SetProperty("qm-16", value)
}

// qm-17 (float32)
//
// Quality factor Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm17(value float32) error {
	return e.Element.SetProperty("qm-17", value)
}

// qm-18 (float32)
//
// Quality factor Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm18(value float32) error {
	return e.Element.SetProperty("qm-18", value)
}

// qm-19 (float32)
//
// Quality factor Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm19(value float32) error {
	return e.Element.SetProperty("qm-19", value)
}

// qm-2 (float32)
//
// Quality factor Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm2(value float32) error {
	return e.Element.SetProperty("qm-2", value)
}

// qm-20 (float32)
//
// Quality factor Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm20(value float32) error {
	return e.Element.SetProperty("qm-20", value)
}

// qm-21 (float32)
//
// Quality factor Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm21(value float32) error {
	return e.Element.SetProperty("qm-21", value)
}

// qm-22 (float32)
//
// Quality factor Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm22(value float32) error {
	return e.Element.SetProperty("qm-22", value)
}

// qm-23 (float32)
//
// Quality factor Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm23(value float32) error {
	return e.Element.SetProperty("qm-23", value)
}

// qm-24 (float32)
//
// Quality factor Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm24(value float32) error {
	return e.Element.SetProperty("qm-24", value)
}

// qm-25 (float32)
//
// Quality factor Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm25(value float32) error {
	return e.Element.SetProperty("qm-25", value)
}

// qm-26 (float32)
//
// Quality factor Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm26(value float32) error {
	return e.Element.SetProperty("qm-26", value)
}

// qm-27 (float32)
//
// Quality factor Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm27(value float32) error {
	return e.Element.SetProperty("qm-27", value)
}

// qm-28 (float32)
//
// Quality factor Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm28(value float32) error {
	return e.Element.SetProperty("qm-28", value)
}

// qm-29 (float32)
//
// Quality factor Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm29(value float32) error {
	return e.Element.SetProperty("qm-29", value)
}

// qm-3 (float32)
//
// Quality factor Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm3(value float32) error {
	return e.Element.SetProperty("qm-3", value)
}

// qm-30 (float32)
//
// Quality factor Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm30(value float32) error {
	return e.Element.SetProperty("qm-30", value)
}

// qm-31 (float32)
//
// Quality factor Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm31(value float32) error {
	return e.Element.SetProperty("qm-31", value)
}

// qm-4 (float32)
//
// Quality factor Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm4(value float32) error {
	return e.Element.SetProperty("qm-4", value)
}

// qm-5 (float32)
//
// Quality factor Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm5(value float32) error {
	return e.Element.SetProperty("qm-5", value)
}

// qm-6 (float32)
//
// Quality factor Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm6(value float32) error {
	return e.Element.SetProperty("qm-6", value)
}

// qm-7 (float32)
//
// Quality factor Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm7(value float32) error {
	return e.Element.SetProperty("qm-7", value)
}

// qm-8 (float32)
//
// Quality factor Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm8(value float32) error {
	return e.Element.SetProperty("qm-8", value)
}

// qm-9 (float32)
//
// Quality factor Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQm9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQm9(value float32) error {
	return e.Element.SetProperty("qm-9", value)
}

// qs-0 (float32)
//
// Quality factor Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs0(value float32) error {
	return e.Element.SetProperty("qs-0", value)
}

// qs-1 (float32)
//
// Quality factor Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs1(value float32) error {
	return e.Element.SetProperty("qs-1", value)
}

// qs-10 (float32)
//
// Quality factor Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs10(value float32) error {
	return e.Element.SetProperty("qs-10", value)
}

// qs-11 (float32)
//
// Quality factor Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs11(value float32) error {
	return e.Element.SetProperty("qs-11", value)
}

// qs-12 (float32)
//
// Quality factor Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs12(value float32) error {
	return e.Element.SetProperty("qs-12", value)
}

// qs-13 (float32)
//
// Quality factor Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs13(value float32) error {
	return e.Element.SetProperty("qs-13", value)
}

// qs-14 (float32)
//
// Quality factor Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs14(value float32) error {
	return e.Element.SetProperty("qs-14", value)
}

// qs-15 (float32)
//
// Quality factor Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs15(value float32) error {
	return e.Element.SetProperty("qs-15", value)
}

// qs-16 (float32)
//
// Quality factor Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs16(value float32) error {
	return e.Element.SetProperty("qs-16", value)
}

// qs-17 (float32)
//
// Quality factor Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs17(value float32) error {
	return e.Element.SetProperty("qs-17", value)
}

// qs-18 (float32)
//
// Quality factor Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs18(value float32) error {
	return e.Element.SetProperty("qs-18", value)
}

// qs-19 (float32)
//
// Quality factor Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs19(value float32) error {
	return e.Element.SetProperty("qs-19", value)
}

// qs-2 (float32)
//
// Quality factor Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs2(value float32) error {
	return e.Element.SetProperty("qs-2", value)
}

// qs-20 (float32)
//
// Quality factor Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs20(value float32) error {
	return e.Element.SetProperty("qs-20", value)
}

// qs-21 (float32)
//
// Quality factor Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs21(value float32) error {
	return e.Element.SetProperty("qs-21", value)
}

// qs-22 (float32)
//
// Quality factor Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs22(value float32) error {
	return e.Element.SetProperty("qs-22", value)
}

// qs-23 (float32)
//
// Quality factor Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs23(value float32) error {
	return e.Element.SetProperty("qs-23", value)
}

// qs-24 (float32)
//
// Quality factor Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs24(value float32) error {
	return e.Element.SetProperty("qs-24", value)
}

// qs-25 (float32)
//
// Quality factor Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs25(value float32) error {
	return e.Element.SetProperty("qs-25", value)
}

// qs-26 (float32)
//
// Quality factor Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs26(value float32) error {
	return e.Element.SetProperty("qs-26", value)
}

// qs-27 (float32)
//
// Quality factor Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs27(value float32) error {
	return e.Element.SetProperty("qs-27", value)
}

// qs-28 (float32)
//
// Quality factor Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs28(value float32) error {
	return e.Element.SetProperty("qs-28", value)
}

// qs-29 (float32)
//
// Quality factor Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs29(value float32) error {
	return e.Element.SetProperty("qs-29", value)
}

// qs-3 (float32)
//
// Quality factor Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs3(value float32) error {
	return e.Element.SetProperty("qs-3", value)
}

// qs-30 (float32)
//
// Quality factor Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs30(value float32) error {
	return e.Element.SetProperty("qs-30", value)
}

// qs-31 (float32)
//
// Quality factor Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qs-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs31(value float32) error {
	return e.Element.SetProperty("qs-31", value)
}

// qs-4 (float32)
//
// Quality factor Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs4(value float32) error {
	return e.Element.SetProperty("qs-4", value)
}

// qs-5 (float32)
//
// Quality factor Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs5(value float32) error {
	return e.Element.SetProperty("qs-5", value)
}

// qs-6 (float32)
//
// Quality factor Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs6(value float32) error {
	return e.Element.SetProperty("qs-6", value)
}

// qs-7 (float32)
//
// Quality factor Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs7(value float32) error {
	return e.Element.SetProperty("qs-7", value)
}

// qs-8 (float32)
//
// Quality factor Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs8(value float32) error {
	return e.Element.SetProperty("qs-8", value)
}

// qs-9 (float32)
//
// Quality factor Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetQs9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetQs9(value float32) error {
	return e.Element.SetProperty("qs-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sm-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm0)
//
// Filter slope Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm0() (interface{}, error) {
	return e.Element.GetProperty("sm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm0(value interface{}) error {
	return e.Element.SetProperty("sm-0", value)
}

// sm-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm1)
//
// Filter slope Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm1() (interface{}, error) {
	return e.Element.GetProperty("sm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm1(value interface{}) error {
	return e.Element.SetProperty("sm-1", value)
}

// sm-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm10)
//
// Filter slope Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm10() (interface{}, error) {
	return e.Element.GetProperty("sm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm10(value interface{}) error {
	return e.Element.SetProperty("sm-10", value)
}

// sm-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm11)
//
// Filter slope Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm11() (interface{}, error) {
	return e.Element.GetProperty("sm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm11(value interface{}) error {
	return e.Element.SetProperty("sm-11", value)
}

// sm-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm12)
//
// Filter slope Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm12() (interface{}, error) {
	return e.Element.GetProperty("sm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm12(value interface{}) error {
	return e.Element.SetProperty("sm-12", value)
}

// sm-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm13)
//
// Filter slope Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm13() (interface{}, error) {
	return e.Element.GetProperty("sm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm13(value interface{}) error {
	return e.Element.SetProperty("sm-13", value)
}

// sm-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm14)
//
// Filter slope Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm14() (interface{}, error) {
	return e.Element.GetProperty("sm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm14(value interface{}) error {
	return e.Element.SetProperty("sm-14", value)
}

// sm-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm15)
//
// Filter slope Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm15() (interface{}, error) {
	return e.Element.GetProperty("sm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm15(value interface{}) error {
	return e.Element.SetProperty("sm-15", value)
}

// sm-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm16)
//
// Filter slope Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm16() (interface{}, error) {
	return e.Element.GetProperty("sm-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm16(value interface{}) error {
	return e.Element.SetProperty("sm-16", value)
}

// sm-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm17)
//
// Filter slope Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm17() (interface{}, error) {
	return e.Element.GetProperty("sm-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm17(value interface{}) error {
	return e.Element.SetProperty("sm-17", value)
}

// sm-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm18)
//
// Filter slope Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm18() (interface{}, error) {
	return e.Element.GetProperty("sm-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm18(value interface{}) error {
	return e.Element.SetProperty("sm-18", value)
}

// sm-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm19)
//
// Filter slope Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm19() (interface{}, error) {
	return e.Element.GetProperty("sm-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm19(value interface{}) error {
	return e.Element.SetProperty("sm-19", value)
}

// sm-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm2)
//
// Filter slope Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm2() (interface{}, error) {
	return e.Element.GetProperty("sm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm2(value interface{}) error {
	return e.Element.SetProperty("sm-2", value)
}

// sm-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm20)
//
// Filter slope Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm20() (interface{}, error) {
	return e.Element.GetProperty("sm-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm20(value interface{}) error {
	return e.Element.SetProperty("sm-20", value)
}

// sm-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm21)
//
// Filter slope Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm21() (interface{}, error) {
	return e.Element.GetProperty("sm-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm21(value interface{}) error {
	return e.Element.SetProperty("sm-21", value)
}

// sm-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm22)
//
// Filter slope Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm22() (interface{}, error) {
	return e.Element.GetProperty("sm-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm22(value interface{}) error {
	return e.Element.SetProperty("sm-22", value)
}

// sm-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm23)
//
// Filter slope Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm23() (interface{}, error) {
	return e.Element.GetProperty("sm-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm23(value interface{}) error {
	return e.Element.SetProperty("sm-23", value)
}

// sm-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm24)
//
// Filter slope Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm24() (interface{}, error) {
	return e.Element.GetProperty("sm-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm24(value interface{}) error {
	return e.Element.SetProperty("sm-24", value)
}

// sm-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm25)
//
// Filter slope Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm25() (interface{}, error) {
	return e.Element.GetProperty("sm-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm25(value interface{}) error {
	return e.Element.SetProperty("sm-25", value)
}

// sm-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm26)
//
// Filter slope Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm26() (interface{}, error) {
	return e.Element.GetProperty("sm-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm26(value interface{}) error {
	return e.Element.SetProperty("sm-26", value)
}

// sm-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm27)
//
// Filter slope Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm27() (interface{}, error) {
	return e.Element.GetProperty("sm-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm27(value interface{}) error {
	return e.Element.SetProperty("sm-27", value)
}

// sm-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm28)
//
// Filter slope Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm28() (interface{}, error) {
	return e.Element.GetProperty("sm-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm28(value interface{}) error {
	return e.Element.SetProperty("sm-28", value)
}

// sm-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm29)
//
// Filter slope Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm29() (interface{}, error) {
	return e.Element.GetProperty("sm-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm29(value interface{}) error {
	return e.Element.SetProperty("sm-29", value)
}

// sm-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm3)
//
// Filter slope Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm3() (interface{}, error) {
	return e.Element.GetProperty("sm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm3(value interface{}) error {
	return e.Element.SetProperty("sm-3", value)
}

// sm-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm30)
//
// Filter slope Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm30() (interface{}, error) {
	return e.Element.GetProperty("sm-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm30(value interface{}) error {
	return e.Element.SetProperty("sm-30", value)
}

// sm-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm31)
//
// Filter slope Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm31() (interface{}, error) {
	return e.Element.GetProperty("sm-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm31(value interface{}) error {
	return e.Element.SetProperty("sm-31", value)
}

// sm-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm4)
//
// Filter slope Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm4() (interface{}, error) {
	return e.Element.GetProperty("sm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm4(value interface{}) error {
	return e.Element.SetProperty("sm-4", value)
}

// sm-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm5)
//
// Filter slope Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm5() (interface{}, error) {
	return e.Element.GetProperty("sm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm5(value interface{}) error {
	return e.Element.SetProperty("sm-5", value)
}

// sm-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm6)
//
// Filter slope Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm6() (interface{}, error) {
	return e.Element.GetProperty("sm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm6(value interface{}) error {
	return e.Element.SetProperty("sm-6", value)
}

// sm-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm7)
//
// Filter slope Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm7() (interface{}, error) {
	return e.Element.GetProperty("sm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm7(value interface{}) error {
	return e.Element.SetProperty("sm-7", value)
}

// sm-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm8)
//
// Filter slope Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm8() (interface{}, error) {
	return e.Element.GetProperty("sm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm8(value interface{}) error {
	return e.Element.SetProperty("sm-8", value)
}

// sm-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Mssm9)
//
// Filter slope Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSm9() (interface{}, error) {
	return e.Element.GetProperty("sm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSm9(value interface{}) error {
	return e.Element.SetProperty("sm-9", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSmr() (float32, error) {
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

// ss-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss0)
//
// Filter slope Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs0() (interface{}, error) {
	return e.Element.GetProperty("ss-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs0(value interface{}) error {
	return e.Element.SetProperty("ss-0", value)
}

// ss-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss1)
//
// Filter slope Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs1() (interface{}, error) {
	return e.Element.GetProperty("ss-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs1(value interface{}) error {
	return e.Element.SetProperty("ss-1", value)
}

// ss-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss10)
//
// Filter slope Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs10() (interface{}, error) {
	return e.Element.GetProperty("ss-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs10(value interface{}) error {
	return e.Element.SetProperty("ss-10", value)
}

// ss-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss11)
//
// Filter slope Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs11() (interface{}, error) {
	return e.Element.GetProperty("ss-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs11(value interface{}) error {
	return e.Element.SetProperty("ss-11", value)
}

// ss-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss12)
//
// Filter slope Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs12() (interface{}, error) {
	return e.Element.GetProperty("ss-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs12(value interface{}) error {
	return e.Element.SetProperty("ss-12", value)
}

// ss-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss13)
//
// Filter slope Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs13() (interface{}, error) {
	return e.Element.GetProperty("ss-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs13(value interface{}) error {
	return e.Element.SetProperty("ss-13", value)
}

// ss-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss14)
//
// Filter slope Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs14() (interface{}, error) {
	return e.Element.GetProperty("ss-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs14(value interface{}) error {
	return e.Element.SetProperty("ss-14", value)
}

// ss-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss15)
//
// Filter slope Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs15() (interface{}, error) {
	return e.Element.GetProperty("ss-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs15(value interface{}) error {
	return e.Element.SetProperty("ss-15", value)
}

// ss-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss16)
//
// Filter slope Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs16() (interface{}, error) {
	return e.Element.GetProperty("ss-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs16(value interface{}) error {
	return e.Element.SetProperty("ss-16", value)
}

// ss-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss17)
//
// Filter slope Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs17() (interface{}, error) {
	return e.Element.GetProperty("ss-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs17(value interface{}) error {
	return e.Element.SetProperty("ss-17", value)
}

// ss-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss18)
//
// Filter slope Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs18() (interface{}, error) {
	return e.Element.GetProperty("ss-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs18(value interface{}) error {
	return e.Element.SetProperty("ss-18", value)
}

// ss-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss19)
//
// Filter slope Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs19() (interface{}, error) {
	return e.Element.GetProperty("ss-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs19(value interface{}) error {
	return e.Element.SetProperty("ss-19", value)
}

// ss-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss2)
//
// Filter slope Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs2() (interface{}, error) {
	return e.Element.GetProperty("ss-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs2(value interface{}) error {
	return e.Element.SetProperty("ss-2", value)
}

// ss-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss20)
//
// Filter slope Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs20() (interface{}, error) {
	return e.Element.GetProperty("ss-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs20(value interface{}) error {
	return e.Element.SetProperty("ss-20", value)
}

// ss-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss21)
//
// Filter slope Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs21() (interface{}, error) {
	return e.Element.GetProperty("ss-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs21(value interface{}) error {
	return e.Element.SetProperty("ss-21", value)
}

// ss-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss22)
//
// Filter slope Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs22() (interface{}, error) {
	return e.Element.GetProperty("ss-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs22(value interface{}) error {
	return e.Element.SetProperty("ss-22", value)
}

// ss-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss23)
//
// Filter slope Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs23() (interface{}, error) {
	return e.Element.GetProperty("ss-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs23(value interface{}) error {
	return e.Element.SetProperty("ss-23", value)
}

// ss-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss24)
//
// Filter slope Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs24() (interface{}, error) {
	return e.Element.GetProperty("ss-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs24(value interface{}) error {
	return e.Element.SetProperty("ss-24", value)
}

// ss-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss25)
//
// Filter slope Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs25() (interface{}, error) {
	return e.Element.GetProperty("ss-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs25(value interface{}) error {
	return e.Element.SetProperty("ss-25", value)
}

// ss-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss26)
//
// Filter slope Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs26() (interface{}, error) {
	return e.Element.GetProperty("ss-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs26(value interface{}) error {
	return e.Element.SetProperty("ss-26", value)
}

// ss-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss27)
//
// Filter slope Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs27() (interface{}, error) {
	return e.Element.GetProperty("ss-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs27(value interface{}) error {
	return e.Element.SetProperty("ss-27", value)
}

// ss-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss28)
//
// Filter slope Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs28() (interface{}, error) {
	return e.Element.GetProperty("ss-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs28(value interface{}) error {
	return e.Element.SetProperty("ss-28", value)
}

// ss-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss29)
//
// Filter slope Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs29() (interface{}, error) {
	return e.Element.GetProperty("ss-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs29(value interface{}) error {
	return e.Element.SetProperty("ss-29", value)
}

// ss-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss3)
//
// Filter slope Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs3() (interface{}, error) {
	return e.Element.GetProperty("ss-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs3(value interface{}) error {
	return e.Element.SetProperty("ss-3", value)
}

// ss-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss30)
//
// Filter slope Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs30() (interface{}, error) {
	return e.Element.GetProperty("ss-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs30(value interface{}) error {
	return e.Element.SetProperty("ss-30", value)
}

// ss-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss31)
//
// Filter slope Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs31() (interface{}, error) {
	return e.Element.GetProperty("ss-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs31(value interface{}) error {
	return e.Element.SetProperty("ss-31", value)
}

// ss-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss4)
//
// Filter slope Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs4() (interface{}, error) {
	return e.Element.GetProperty("ss-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs4(value interface{}) error {
	return e.Element.SetProperty("ss-4", value)
}

// ss-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss5)
//
// Filter slope Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs5() (interface{}, error) {
	return e.Element.GetProperty("ss-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs5(value interface{}) error {
	return e.Element.SetProperty("ss-5", value)
}

// ss-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss6)
//
// Filter slope Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs6() (interface{}, error) {
	return e.Element.GetProperty("ss-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs6(value interface{}) error {
	return e.Element.SetProperty("ss-6", value)
}

// ss-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss7)
//
// Filter slope Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs7() (interface{}, error) {
	return e.Element.GetProperty("ss-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs7(value interface{}) error {
	return e.Element.SetProperty("ss-7", value)
}

// ss-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss8)
//
// Filter slope Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs8() (interface{}, error) {
	return e.Element.GetProperty("ss-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs8(value interface{}) error {
	return e.Element.SetProperty("ss-8", value)
}

// ss-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Msss9)
//
// Filter slope Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetSs9() (interface{}, error) {
	return e.Element.GetProperty("ss-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetSs9(value interface{}) error {
	return e.Element.SetProperty("ss-9", value)
}

// xmm-0 (bool)
//
// Filter mute Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm0(value bool) error {
	return e.Element.SetProperty("xmm-0", value)
}

// xmm-1 (bool)
//
// Filter mute Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm1(value bool) error {
	return e.Element.SetProperty("xmm-1", value)
}

// xmm-10 (bool)
//
// Filter mute Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm10(value bool) error {
	return e.Element.SetProperty("xmm-10", value)
}

// xmm-11 (bool)
//
// Filter mute Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm11(value bool) error {
	return e.Element.SetProperty("xmm-11", value)
}

// xmm-12 (bool)
//
// Filter mute Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm12(value bool) error {
	return e.Element.SetProperty("xmm-12", value)
}

// xmm-13 (bool)
//
// Filter mute Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm13(value bool) error {
	return e.Element.SetProperty("xmm-13", value)
}

// xmm-14 (bool)
//
// Filter mute Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm14(value bool) error {
	return e.Element.SetProperty("xmm-14", value)
}

// xmm-15 (bool)
//
// Filter mute Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm15(value bool) error {
	return e.Element.SetProperty("xmm-15", value)
}

// xmm-16 (bool)
//
// Filter mute Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm16(value bool) error {
	return e.Element.SetProperty("xmm-16", value)
}

// xmm-17 (bool)
//
// Filter mute Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm17(value bool) error {
	return e.Element.SetProperty("xmm-17", value)
}

// xmm-18 (bool)
//
// Filter mute Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm18(value bool) error {
	return e.Element.SetProperty("xmm-18", value)
}

// xmm-19 (bool)
//
// Filter mute Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm19(value bool) error {
	return e.Element.SetProperty("xmm-19", value)
}

// xmm-2 (bool)
//
// Filter mute Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm2(value bool) error {
	return e.Element.SetProperty("xmm-2", value)
}

// xmm-20 (bool)
//
// Filter mute Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm20(value bool) error {
	return e.Element.SetProperty("xmm-20", value)
}

// xmm-21 (bool)
//
// Filter mute Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm21(value bool) error {
	return e.Element.SetProperty("xmm-21", value)
}

// xmm-22 (bool)
//
// Filter mute Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm22(value bool) error {
	return e.Element.SetProperty("xmm-22", value)
}

// xmm-23 (bool)
//
// Filter mute Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm23(value bool) error {
	return e.Element.SetProperty("xmm-23", value)
}

// xmm-24 (bool)
//
// Filter mute Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm24(value bool) error {
	return e.Element.SetProperty("xmm-24", value)
}

// xmm-25 (bool)
//
// Filter mute Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm25(value bool) error {
	return e.Element.SetProperty("xmm-25", value)
}

// xmm-26 (bool)
//
// Filter mute Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm26(value bool) error {
	return e.Element.SetProperty("xmm-26", value)
}

// xmm-27 (bool)
//
// Filter mute Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm27(value bool) error {
	return e.Element.SetProperty("xmm-27", value)
}

// xmm-28 (bool)
//
// Filter mute Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm28(value bool) error {
	return e.Element.SetProperty("xmm-28", value)
}

// xmm-29 (bool)
//
// Filter mute Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm29(value bool) error {
	return e.Element.SetProperty("xmm-29", value)
}

// xmm-3 (bool)
//
// Filter mute Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm3(value bool) error {
	return e.Element.SetProperty("xmm-3", value)
}

// xmm-30 (bool)
//
// Filter mute Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm30(value bool) error {
	return e.Element.SetProperty("xmm-30", value)
}

// xmm-31 (bool)
//
// Filter mute Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm31(value bool) error {
	return e.Element.SetProperty("xmm-31", value)
}

// xmm-4 (bool)
//
// Filter mute Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm4(value bool) error {
	return e.Element.SetProperty("xmm-4", value)
}

// xmm-5 (bool)
//
// Filter mute Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm5(value bool) error {
	return e.Element.SetProperty("xmm-5", value)
}

// xmm-6 (bool)
//
// Filter mute Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm6(value bool) error {
	return e.Element.SetProperty("xmm-6", value)
}

// xmm-7 (bool)
//
// Filter mute Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm7(value bool) error {
	return e.Element.SetProperty("xmm-7", value)
}

// xmm-8 (bool)
//
// Filter mute Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm8(value bool) error {
	return e.Element.SetProperty("xmm-8", value)
}

// xmm-9 (bool)
//
// Filter mute Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXmm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXmm9(value bool) error {
	return e.Element.SetProperty("xmm-9", value)
}

// xms-0 (bool)
//
// Filter mute Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms0(value bool) error {
	return e.Element.SetProperty("xms-0", value)
}

// xms-1 (bool)
//
// Filter mute Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms1(value bool) error {
	return e.Element.SetProperty("xms-1", value)
}

// xms-10 (bool)
//
// Filter mute Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms10(value bool) error {
	return e.Element.SetProperty("xms-10", value)
}

// xms-11 (bool)
//
// Filter mute Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms11(value bool) error {
	return e.Element.SetProperty("xms-11", value)
}

// xms-12 (bool)
//
// Filter mute Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms12(value bool) error {
	return e.Element.SetProperty("xms-12", value)
}

// xms-13 (bool)
//
// Filter mute Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms13(value bool) error {
	return e.Element.SetProperty("xms-13", value)
}

// xms-14 (bool)
//
// Filter mute Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms14(value bool) error {
	return e.Element.SetProperty("xms-14", value)
}

// xms-15 (bool)
//
// Filter mute Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms15(value bool) error {
	return e.Element.SetProperty("xms-15", value)
}

// xms-16 (bool)
//
// Filter mute Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms16(value bool) error {
	return e.Element.SetProperty("xms-16", value)
}

// xms-17 (bool)
//
// Filter mute Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms17(value bool) error {
	return e.Element.SetProperty("xms-17", value)
}

// xms-18 (bool)
//
// Filter mute Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms18(value bool) error {
	return e.Element.SetProperty("xms-18", value)
}

// xms-19 (bool)
//
// Filter mute Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms19(value bool) error {
	return e.Element.SetProperty("xms-19", value)
}

// xms-2 (bool)
//
// Filter mute Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms2(value bool) error {
	return e.Element.SetProperty("xms-2", value)
}

// xms-20 (bool)
//
// Filter mute Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms20(value bool) error {
	return e.Element.SetProperty("xms-20", value)
}

// xms-21 (bool)
//
// Filter mute Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms21(value bool) error {
	return e.Element.SetProperty("xms-21", value)
}

// xms-22 (bool)
//
// Filter mute Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms22(value bool) error {
	return e.Element.SetProperty("xms-22", value)
}

// xms-23 (bool)
//
// Filter mute Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms23(value bool) error {
	return e.Element.SetProperty("xms-23", value)
}

// xms-24 (bool)
//
// Filter mute Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms24(value bool) error {
	return e.Element.SetProperty("xms-24", value)
}

// xms-25 (bool)
//
// Filter mute Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms25(value bool) error {
	return e.Element.SetProperty("xms-25", value)
}

// xms-26 (bool)
//
// Filter mute Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms26(value bool) error {
	return e.Element.SetProperty("xms-26", value)
}

// xms-27 (bool)
//
// Filter mute Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms27(value bool) error {
	return e.Element.SetProperty("xms-27", value)
}

// xms-28 (bool)
//
// Filter mute Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms28(value bool) error {
	return e.Element.SetProperty("xms-28", value)
}

// xms-29 (bool)
//
// Filter mute Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms29(value bool) error {
	return e.Element.SetProperty("xms-29", value)
}

// xms-3 (bool)
//
// Filter mute Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms3(value bool) error {
	return e.Element.SetProperty("xms-3", value)
}

// xms-30 (bool)
//
// Filter mute Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms30(value bool) error {
	return e.Element.SetProperty("xms-30", value)
}

// xms-31 (bool)
//
// Filter mute Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xms-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms31(value bool) error {
	return e.Element.SetProperty("xms-31", value)
}

// xms-4 (bool)
//
// Filter mute Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms4(value bool) error {
	return e.Element.SetProperty("xms-4", value)
}

// xms-5 (bool)
//
// Filter mute Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms5(value bool) error {
	return e.Element.SetProperty("xms-5", value)
}

// xms-6 (bool)
//
// Filter mute Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms6(value bool) error {
	return e.Element.SetProperty("xms-6", value)
}

// xms-7 (bool)
//
// Filter mute Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms7(value bool) error {
	return e.Element.SetProperty("xms-7", value)
}

// xms-8 (bool)
//
// Filter mute Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms8(value bool) error {
	return e.Element.SetProperty("xms-8", value)
}

// xms-9 (bool)
//
// Filter mute Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXms9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXms9(value bool) error {
	return e.Element.SetProperty("xms-9", value)
}

// xsm-0 (bool)
//
// Filter solo Mid 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm0(value bool) error {
	return e.Element.SetProperty("xsm-0", value)
}

// xsm-1 (bool)
//
// Filter solo Mid 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm1(value bool) error {
	return e.Element.SetProperty("xsm-1", value)
}

// xsm-10 (bool)
//
// Filter solo Mid 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm10(value bool) error {
	return e.Element.SetProperty("xsm-10", value)
}

// xsm-11 (bool)
//
// Filter solo Mid 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm11(value bool) error {
	return e.Element.SetProperty("xsm-11", value)
}

// xsm-12 (bool)
//
// Filter solo Mid 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm12(value bool) error {
	return e.Element.SetProperty("xsm-12", value)
}

// xsm-13 (bool)
//
// Filter solo Mid 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm13(value bool) error {
	return e.Element.SetProperty("xsm-13", value)
}

// xsm-14 (bool)
//
// Filter solo Mid 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm14(value bool) error {
	return e.Element.SetProperty("xsm-14", value)
}

// xsm-15 (bool)
//
// Filter solo Mid 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm15(value bool) error {
	return e.Element.SetProperty("xsm-15", value)
}

// xsm-16 (bool)
//
// Filter solo Mid 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm16(value bool) error {
	return e.Element.SetProperty("xsm-16", value)
}

// xsm-17 (bool)
//
// Filter solo Mid 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm17(value bool) error {
	return e.Element.SetProperty("xsm-17", value)
}

// xsm-18 (bool)
//
// Filter solo Mid 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm18(value bool) error {
	return e.Element.SetProperty("xsm-18", value)
}

// xsm-19 (bool)
//
// Filter solo Mid 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm19(value bool) error {
	return e.Element.SetProperty("xsm-19", value)
}

// xsm-2 (bool)
//
// Filter solo Mid 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm2(value bool) error {
	return e.Element.SetProperty("xsm-2", value)
}

// xsm-20 (bool)
//
// Filter solo Mid 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm20(value bool) error {
	return e.Element.SetProperty("xsm-20", value)
}

// xsm-21 (bool)
//
// Filter solo Mid 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm21(value bool) error {
	return e.Element.SetProperty("xsm-21", value)
}

// xsm-22 (bool)
//
// Filter solo Mid 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm22(value bool) error {
	return e.Element.SetProperty("xsm-22", value)
}

// xsm-23 (bool)
//
// Filter solo Mid 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm23(value bool) error {
	return e.Element.SetProperty("xsm-23", value)
}

// xsm-24 (bool)
//
// Filter solo Mid 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm24(value bool) error {
	return e.Element.SetProperty("xsm-24", value)
}

// xsm-25 (bool)
//
// Filter solo Mid 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm25(value bool) error {
	return e.Element.SetProperty("xsm-25", value)
}

// xsm-26 (bool)
//
// Filter solo Mid 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm26(value bool) error {
	return e.Element.SetProperty("xsm-26", value)
}

// xsm-27 (bool)
//
// Filter solo Mid 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm27(value bool) error {
	return e.Element.SetProperty("xsm-27", value)
}

// xsm-28 (bool)
//
// Filter solo Mid 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm28(value bool) error {
	return e.Element.SetProperty("xsm-28", value)
}

// xsm-29 (bool)
//
// Filter solo Mid 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm29(value bool) error {
	return e.Element.SetProperty("xsm-29", value)
}

// xsm-3 (bool)
//
// Filter solo Mid 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm3(value bool) error {
	return e.Element.SetProperty("xsm-3", value)
}

// xsm-30 (bool)
//
// Filter solo Mid 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm30(value bool) error {
	return e.Element.SetProperty("xsm-30", value)
}

// xsm-31 (bool)
//
// Filter solo Mid 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm31(value bool) error {
	return e.Element.SetProperty("xsm-31", value)
}

// xsm-4 (bool)
//
// Filter solo Mid 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm4(value bool) error {
	return e.Element.SetProperty("xsm-4", value)
}

// xsm-5 (bool)
//
// Filter solo Mid 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm5(value bool) error {
	return e.Element.SetProperty("xsm-5", value)
}

// xsm-6 (bool)
//
// Filter solo Mid 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm6(value bool) error {
	return e.Element.SetProperty("xsm-6", value)
}

// xsm-7 (bool)
//
// Filter solo Mid 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm7(value bool) error {
	return e.Element.SetProperty("xsm-7", value)
}

// xsm-8 (bool)
//
// Filter solo Mid 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm8(value bool) error {
	return e.Element.SetProperty("xsm-8", value)
}

// xsm-9 (bool)
//
// Filter solo Mid 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXsm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXsm9(value bool) error {
	return e.Element.SetProperty("xsm-9", value)
}

// xss-0 (bool)
//
// Filter solo Side 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss0(value bool) error {
	return e.Element.SetProperty("xss-0", value)
}

// xss-1 (bool)
//
// Filter solo Side 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss1(value bool) error {
	return e.Element.SetProperty("xss-1", value)
}

// xss-10 (bool)
//
// Filter solo Side 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss10(value bool) error {
	return e.Element.SetProperty("xss-10", value)
}

// xss-11 (bool)
//
// Filter solo Side 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss11(value bool) error {
	return e.Element.SetProperty("xss-11", value)
}

// xss-12 (bool)
//
// Filter solo Side 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss12(value bool) error {
	return e.Element.SetProperty("xss-12", value)
}

// xss-13 (bool)
//
// Filter solo Side 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss13(value bool) error {
	return e.Element.SetProperty("xss-13", value)
}

// xss-14 (bool)
//
// Filter solo Side 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss14(value bool) error {
	return e.Element.SetProperty("xss-14", value)
}

// xss-15 (bool)
//
// Filter solo Side 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss15(value bool) error {
	return e.Element.SetProperty("xss-15", value)
}

// xss-16 (bool)
//
// Filter solo Side 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss16(value bool) error {
	return e.Element.SetProperty("xss-16", value)
}

// xss-17 (bool)
//
// Filter solo Side 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss17(value bool) error {
	return e.Element.SetProperty("xss-17", value)
}

// xss-18 (bool)
//
// Filter solo Side 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss18(value bool) error {
	return e.Element.SetProperty("xss-18", value)
}

// xss-19 (bool)
//
// Filter solo Side 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss19(value bool) error {
	return e.Element.SetProperty("xss-19", value)
}

// xss-2 (bool)
//
// Filter solo Side 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss2(value bool) error {
	return e.Element.SetProperty("xss-2", value)
}

// xss-20 (bool)
//
// Filter solo Side 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss20(value bool) error {
	return e.Element.SetProperty("xss-20", value)
}

// xss-21 (bool)
//
// Filter solo Side 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss21(value bool) error {
	return e.Element.SetProperty("xss-21", value)
}

// xss-22 (bool)
//
// Filter solo Side 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss22(value bool) error {
	return e.Element.SetProperty("xss-22", value)
}

// xss-23 (bool)
//
// Filter solo Side 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss23(value bool) error {
	return e.Element.SetProperty("xss-23", value)
}

// xss-24 (bool)
//
// Filter solo Side 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss24(value bool) error {
	return e.Element.SetProperty("xss-24", value)
}

// xss-25 (bool)
//
// Filter solo Side 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss25(value bool) error {
	return e.Element.SetProperty("xss-25", value)
}

// xss-26 (bool)
//
// Filter solo Side 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss26(value bool) error {
	return e.Element.SetProperty("xss-26", value)
}

// xss-27 (bool)
//
// Filter solo Side 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss27(value bool) error {
	return e.Element.SetProperty("xss-27", value)
}

// xss-28 (bool)
//
// Filter solo Side 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss28(value bool) error {
	return e.Element.SetProperty("xss-28", value)
}

// xss-29 (bool)
//
// Filter solo Side 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss29(value bool) error {
	return e.Element.SetProperty("xss-29", value)
}

// xss-3 (bool)
//
// Filter solo Side 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss3(value bool) error {
	return e.Element.SetProperty("xss-3", value)
}

// xss-30 (bool)
//
// Filter solo Side 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss30(value bool) error {
	return e.Element.SetProperty("xss-30", value)
}

// xss-31 (bool)
//
// Filter solo Side 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xss-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss31(value bool) error {
	return e.Element.SetProperty("xss-31", value)
}

// xss-4 (bool)
//
// Filter solo Side 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss4(value bool) error {
	return e.Element.SetProperty("xss-4", value)
}

// xss-5 (bool)
//
// Filter solo Side 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss5(value bool) error {
	return e.Element.SetProperty("xss-5", value)
}

// xss-6 (bool)
//
// Filter solo Side 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss6(value bool) error {
	return e.Element.SetProperty("xss-6", value)
}

// xss-7 (bool)
//
// Filter solo Side 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss7(value bool) error {
	return e.Element.SetProperty("xss-7", value)
}

// xss-8 (bool)
//
// Filter solo Side 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss8(value bool) error {
	return e.Element.SetProperty("xss-8", value)
}

// xss-9 (bool)
//
// Filter solo Side 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetXss9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetXss9(value bool) error {
	return e.Element.SetProperty("xss-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Ms) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX32Mssm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm6X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm6X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm6X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm6X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms7ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23Off LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm23Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm23 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18Off LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts18Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts18 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8Off LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts8Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms27ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms27 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms31ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms31 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm13X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm13X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm13X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm13X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss1X1 LspPlugInPluginsLv2ParaEqualizerX32Msss1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss1X2 LspPlugInPluginsLv2ParaEqualizerX32Msss1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss1X3 LspPlugInPluginsLv2ParaEqualizerX32Msss1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss1X4 LspPlugInPluginsLv2ParaEqualizerX32Msss1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm18ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm18 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm9ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms2ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms20ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms20 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm25X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm25 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm25X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm25 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm25X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm25 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm25X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm25 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss18X1 LspPlugInPluginsLv2ParaEqualizerX32Msss18 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss18X2 LspPlugInPluginsLv2ParaEqualizerX32Msss18 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss18X3 LspPlugInPluginsLv2ParaEqualizerX32Msss18 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss18X4 LspPlugInPluginsLv2ParaEqualizerX32Msss18 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss20X1 LspPlugInPluginsLv2ParaEqualizerX32Msss20 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss20X2 LspPlugInPluginsLv2ParaEqualizerX32Msss20 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss20X3 LspPlugInPluginsLv2ParaEqualizerX32Msss20 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss20X4 LspPlugInPluginsLv2ParaEqualizerX32Msss20 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms11ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms23ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms23 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms4ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28Off LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts28Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts28 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm21X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm21 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm21X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm21 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm21X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm21 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm21X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm21 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm31X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm31 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm31X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm31 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm31X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm31 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm31X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm31 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm20ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm20 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm28ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm28 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms21ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms21 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24Off LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts24Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts24 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm4ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms1ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm9X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm9X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm9X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm9X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss14X1 LspPlugInPluginsLv2ParaEqualizerX32Msss14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss14X2 LspPlugInPluginsLv2ParaEqualizerX32Msss14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss14X3 LspPlugInPluginsLv2ParaEqualizerX32Msss14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss14X4 LspPlugInPluginsLv2ParaEqualizerX32Msss14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms9ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms10ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss7X1 LspPlugInPluginsLv2ParaEqualizerX32Msss7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss7X2 LspPlugInPluginsLv2ParaEqualizerX32Msss7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss7X3 LspPlugInPluginsLv2ParaEqualizerX32Msss7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss7X4 LspPlugInPluginsLv2ParaEqualizerX32Msss7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27Off LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm27Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm27 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm10X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm10X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm10X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm10X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss6X1 LspPlugInPluginsLv2ParaEqualizerX32Msss6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss6X2 LspPlugInPluginsLv2ParaEqualizerX32Msss6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss6X3 LspPlugInPluginsLv2ParaEqualizerX32Msss6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss6X4 LspPlugInPluginsLv2ParaEqualizerX32Msss6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16Off LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts16Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts16 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30Off LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts30Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts30 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6Off LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts6Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms0ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms17ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms17 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms6ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1Off LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts1Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm1ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm22ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm22 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm17X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm17 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm17X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm17 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm17X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm17 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm17X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm17 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss23X1 LspPlugInPluginsLv2ParaEqualizerX32Msss23 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss23X2 LspPlugInPluginsLv2ParaEqualizerX32Msss23 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss23X3 LspPlugInPluginsLv2ParaEqualizerX32Msss23 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss23X4 LspPlugInPluginsLv2ParaEqualizerX32Msss23 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms18ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms18 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16Off LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm16Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm16 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9Off LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm9Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm8X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm8X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm8X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm8X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6Off LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm6Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12Off LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts12Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm28X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm28 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm28X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm28 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm28X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm28 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm28X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm28 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss0X1 LspPlugInPluginsLv2ParaEqualizerX32Msss0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss0X2 LspPlugInPluginsLv2ParaEqualizerX32Msss0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss0X3 LspPlugInPluginsLv2ParaEqualizerX32Msss0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss0X4 LspPlugInPluginsLv2ParaEqualizerX32Msss0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm15ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm19ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm19 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms13ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms24ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms24 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26Off LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm26Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm26 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11Off LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts11Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17Off LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts17Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts17 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss26X1 LspPlugInPluginsLv2ParaEqualizerX32Msss26 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss26X2 LspPlugInPluginsLv2ParaEqualizerX32Msss26 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss26X3 LspPlugInPluginsLv2ParaEqualizerX32Msss26 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss26X4 LspPlugInPluginsLv2ParaEqualizerX32Msss26 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm29ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm29 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms3ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2Off LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm2Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24Off LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm24Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm24 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31Off LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm31Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm31 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm2X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm2X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm2X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm2X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm26X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm26 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm26X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm26 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm26X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm26 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm26X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm26 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss11X1 LspPlugInPluginsLv2ParaEqualizerX32Msss11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss11X2 LspPlugInPluginsLv2ParaEqualizerX32Msss11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss11X3 LspPlugInPluginsLv2ParaEqualizerX32Msss11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss11X4 LspPlugInPluginsLv2ParaEqualizerX32Msss11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm16ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm16 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm3ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms25ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms25 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms5ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20Off LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts20Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts20 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm11X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm11X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm11X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm11X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm22X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm22 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm22X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm22 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm22X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm22 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm22X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm22 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm29X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm29 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm29X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm29 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm29X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm29 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm29X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm29 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm2ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms29ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms29 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22Off LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm22Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm22 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19Off LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts19Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts19 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm30X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm30 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm30X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm30 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm30X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm30 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm30X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm30 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss16X1 LspPlugInPluginsLv2ParaEqualizerX32Msss16 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss16X2 LspPlugInPluginsLv2ParaEqualizerX32Msss16 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss16X3 LspPlugInPluginsLv2ParaEqualizerX32Msss16 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss16X4 LspPlugInPluginsLv2ParaEqualizerX32Msss16 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss4X1 LspPlugInPluginsLv2ParaEqualizerX32Msss4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss4X2 LspPlugInPluginsLv2ParaEqualizerX32Msss4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss4X3 LspPlugInPluginsLv2ParaEqualizerX32Msss4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss4X4 LspPlugInPluginsLv2ParaEqualizerX32Msss4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8Off LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm8Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21Off LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts21Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts21 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23Off LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts23Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts23 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31Off LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts31Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts31 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm10ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm7ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms8ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18Off LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm18Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm18 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss28X1 LspPlugInPluginsLv2ParaEqualizerX32Msss28 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss28X2 LspPlugInPluginsLv2ParaEqualizerX32Msss28 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss28X3 LspPlugInPluginsLv2ParaEqualizerX32Msss28 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss28X4 LspPlugInPluginsLv2ParaEqualizerX32Msss28 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm15X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm15X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm15X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm15X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm19X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm19 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm19X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm19 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm19X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm19 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm19X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm19 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm20X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm20 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm20X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm20 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm20X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm20 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm20X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm20 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm5X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm5X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm5X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm5X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfft string

const (
	LspPlugInPluginsLv2ParaEqualizerX32MsfftOff LspPlugInPluginsLv2ParaEqualizerX32Msfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32MsfftPostEq LspPlugInPluginsLv2ParaEqualizerX32Msfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX32MsfftPreEq LspPlugInPluginsLv2ParaEqualizerX32Msfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm12ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm8ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm27X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm27 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm27X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm27 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm27X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm27 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm27X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm27 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm18X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm18 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm18X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm18 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm18X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm18 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm18X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm18 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm27ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm27 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms26ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms26 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13Off LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm13Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14Off LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm14Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25Off LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts25Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts25 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm1X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm1X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm1X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm1X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm12X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm12X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm12X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm12X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm3X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm3X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm3X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm3X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm14ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm31ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm31 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11Off LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm11Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22Off LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts22Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts22 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss13X1 LspPlugInPluginsLv2ParaEqualizerX32Msss13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss13X2 LspPlugInPluginsLv2ParaEqualizerX32Msss13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss13X3 LspPlugInPluginsLv2ParaEqualizerX32Msss13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss13X4 LspPlugInPluginsLv2ParaEqualizerX32Msss13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss2X1 LspPlugInPluginsLv2ParaEqualizerX32Msss2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss2X2 LspPlugInPluginsLv2ParaEqualizerX32Msss2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss2X3 LspPlugInPluginsLv2ParaEqualizerX32Msss2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss2X4 LspPlugInPluginsLv2ParaEqualizerX32Msss2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss24X1 LspPlugInPluginsLv2ParaEqualizerX32Msss24 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss24X2 LspPlugInPluginsLv2ParaEqualizerX32Msss24 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss24X3 LspPlugInPluginsLv2ParaEqualizerX32Msss24 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss24X4 LspPlugInPluginsLv2ParaEqualizerX32Msss24 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14Off LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts14Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss8X1 LspPlugInPluginsLv2ParaEqualizerX32Msss8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss8X2 LspPlugInPluginsLv2ParaEqualizerX32Msss8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss8X3 LspPlugInPluginsLv2ParaEqualizerX32Msss8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss8X4 LspPlugInPluginsLv2ParaEqualizerX32Msss8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm6ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms12ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30Off LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm30Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm30 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5Off LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm5Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss3X1 LspPlugInPluginsLv2ParaEqualizerX32Msss3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss3X2 LspPlugInPluginsLv2ParaEqualizerX32Msss3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss3X3 LspPlugInPluginsLv2ParaEqualizerX32Msss3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss3X4 LspPlugInPluginsLv2ParaEqualizerX32Msss3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss12X1 LspPlugInPluginsLv2ParaEqualizerX32Msss12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss12X2 LspPlugInPluginsLv2ParaEqualizerX32Msss12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss12X3 LspPlugInPluginsLv2ParaEqualizerX32Msss12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss12X4 LspPlugInPluginsLv2ParaEqualizerX32Msss12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm11ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1Off LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm1Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20Off LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm20Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm20 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9Off LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts9Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25Off LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm25Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm25 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0Off LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts0Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2Off LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts2Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4Off LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts4Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm30ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm30 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms15ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms16ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms16 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12Off LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm12Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5Off LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts5Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msmode string

const (
	LspPlugInPluginsLv2ParaEqualizerX32MsmodeIir LspPlugInPluginsLv2ParaEqualizerX32Msmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX32MsmodeFir LspPlugInPluginsLv2ParaEqualizerX32Msmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX32MsmodeFft LspPlugInPluginsLv2ParaEqualizerX32Msmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm23X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm23 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm23X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm23 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm23X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm23 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm23X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm23 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm24X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm24 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm24X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm24 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm24X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm24 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm24X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm24 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss19X1 LspPlugInPluginsLv2ParaEqualizerX32Msss19 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss19X2 LspPlugInPluginsLv2ParaEqualizerX32Msss19 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss19X3 LspPlugInPluginsLv2ParaEqualizerX32Msss19 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss19X4 LspPlugInPluginsLv2ParaEqualizerX32Msss19 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss30X1 LspPlugInPluginsLv2ParaEqualizerX32Msss30 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss30X2 LspPlugInPluginsLv2ParaEqualizerX32Msss30 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss30X3 LspPlugInPluginsLv2ParaEqualizerX32Msss30 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss30X4 LspPlugInPluginsLv2ParaEqualizerX32Msss30 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm21ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm21 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29Off LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts29Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts29 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm16X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm16 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm16X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm16 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm16X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm16 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm16X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm16 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm4X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm4X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm4X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm4X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss25X1 LspPlugInPluginsLv2ParaEqualizerX32Msss25 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss25X2 LspPlugInPluginsLv2ParaEqualizerX32Msss25 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss25X3 LspPlugInPluginsLv2ParaEqualizerX32Msss25 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss25X4 LspPlugInPluginsLv2ParaEqualizerX32Msss25 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15Off LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts15Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26Off LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts26Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts26 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm7X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm7X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm7X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm7X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss15X1 LspPlugInPluginsLv2ParaEqualizerX32Msss15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss15X2 LspPlugInPluginsLv2ParaEqualizerX32Msss15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss15X3 LspPlugInPluginsLv2ParaEqualizerX32Msss15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss15X4 LspPlugInPluginsLv2ParaEqualizerX32Msss15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm14X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm14X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm14X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm14X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm5ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms19ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms19 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms30ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms30 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17Off LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm17Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm17 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7Off LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm7Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Mssm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Mssm0X1 LspPlugInPluginsLv2ParaEqualizerX32Mssm0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Mssm0X2 LspPlugInPluginsLv2ParaEqualizerX32Mssm0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Mssm0X3 LspPlugInPluginsLv2ParaEqualizerX32Mssm0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Mssm0X4 LspPlugInPluginsLv2ParaEqualizerX32Mssm0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss10X1 LspPlugInPluginsLv2ParaEqualizerX32Msss10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss10X2 LspPlugInPluginsLv2ParaEqualizerX32Msss10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss10X3 LspPlugInPluginsLv2ParaEqualizerX32Msss10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss10X4 LspPlugInPluginsLv2ParaEqualizerX32Msss10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss17X1 LspPlugInPluginsLv2ParaEqualizerX32Msss17 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss17X2 LspPlugInPluginsLv2ParaEqualizerX32Msss17 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss17X3 LspPlugInPluginsLv2ParaEqualizerX32Msss17 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss17X4 LspPlugInPluginsLv2ParaEqualizerX32Msss17 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm26ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm26 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms22ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms22 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15Off LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm15Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28Off LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm28Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm28 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm0ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0Off LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm0Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21Off LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm21Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm21 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3Off LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm3Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss9X1 LspPlugInPluginsLv2ParaEqualizerX32Msss9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss9X2 LspPlugInPluginsLv2ParaEqualizerX32Msss9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss9X3 LspPlugInPluginsLv2ParaEqualizerX32Msss9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss9X4 LspPlugInPluginsLv2ParaEqualizerX32Msss9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm13ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms28ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms28 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4Off LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm4Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13Off LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts13Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3Off LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts3Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss27X1 LspPlugInPluginsLv2ParaEqualizerX32Msss27 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss27X2 LspPlugInPluginsLv2ParaEqualizerX32Msss27 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss27X3 LspPlugInPluginsLv2ParaEqualizerX32Msss27 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss27X4 LspPlugInPluginsLv2ParaEqualizerX32Msss27 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss29X1 LspPlugInPluginsLv2ParaEqualizerX32Msss29 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss29X2 LspPlugInPluginsLv2ParaEqualizerX32Msss29 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss29X3 LspPlugInPluginsLv2ParaEqualizerX32Msss29 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss29X4 LspPlugInPluginsLv2ParaEqualizerX32Msss29 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm17ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm17 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm23ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm23 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm25ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm25 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29Off LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm29Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm29 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss21X1 LspPlugInPluginsLv2ParaEqualizerX32Msss21 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss21X2 LspPlugInPluginsLv2ParaEqualizerX32Msss21 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss21X3 LspPlugInPluginsLv2ParaEqualizerX32Msss21 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss21X4 LspPlugInPluginsLv2ParaEqualizerX32Msss21 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss22X1 LspPlugInPluginsLv2ParaEqualizerX32Msss22 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss22X2 LspPlugInPluginsLv2ParaEqualizerX32Msss22 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss22X3 LspPlugInPluginsLv2ParaEqualizerX32Msss22 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss22X4 LspPlugInPluginsLv2ParaEqualizerX32Msss22 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfmm24ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfmm24 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersMid07 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Mid 0-7" // Filters Mid 0-7 (0) – Filters Mid 0-7
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersSide07 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Side 0-7" // Filters Side 0-7 (1) – Filters Side 0-7
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersMid815 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Mid 8-15" // Filters Mid 8-15 (2) – Filters Mid 8-15
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersSide815 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Side 8-15" // Filters Side 8-15 (3) – Filters Side 8-15
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersMid1623 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Mid 16-23" // Filters Mid 16-23 (4) – Filters Mid 16-23
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersSide1623 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Side 16-23" // Filters Side 16-23 (5) – Filters Side 16-23
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersMid2431 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Mid 24-31" // Filters Mid 24-31 (6) – Filters Mid 24-31
	LspPlugInPluginsLv2ParaEqualizerX32MsfselFiltersSide2431 LspPlugInPluginsLv2ParaEqualizerX32Msfsel = "Filters Side 24-31" // Filters Side 24-31 (7) – Filters Side 24-31
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10Off LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts10Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27Off LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts27Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts27 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss31X1 LspPlugInPluginsLv2ParaEqualizerX32Msss31 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss31X2 LspPlugInPluginsLv2ParaEqualizerX32Msss31 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss31X3 LspPlugInPluginsLv2ParaEqualizerX32Msss31 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss31X4 LspPlugInPluginsLv2ParaEqualizerX32Msss31 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msss5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msss5X1 LspPlugInPluginsLv2ParaEqualizerX32Msss5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Msss5X2 LspPlugInPluginsLv2ParaEqualizerX32Msss5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Msss5X3 LspPlugInPluginsLv2ParaEqualizerX32Msss5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Msss5X4 LspPlugInPluginsLv2ParaEqualizerX32Msss5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfms14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14RlcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14RlcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14BwcBt LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14BwcMt LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14LrxBt LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14LrxMt LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Msfms14ApoDr LspPlugInPluginsLv2ParaEqualizerX32Msfms14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10Off LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm10Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msftm19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19Off LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19Bell LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19HiPass LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19LoPass LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19Notch LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19Resonance LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msftm19Allpass LspPlugInPluginsLv2ParaEqualizerX32Msftm19 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Msfts7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7Off LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7Bell LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7HiPass LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7HiShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7LoPass LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7LoShelf LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7Notch LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7Resonance LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Msfts7Allpass LspPlugInPluginsLv2ParaEqualizerX32Msfts7 = "Allpass" // Allpass (8) – Allpass
)

