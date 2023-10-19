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

type LspPlugInPluginsLv2ParaEqualizerX16Stereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX16Stereo() (*LspPlugInPluginsLv2ParaEqualizerX16Stereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x16-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX16StereoWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX16Stereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x16-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// f-0 (float32)
//
// Frequency 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF0(value float32) error {
	return e.Element.SetProperty("f-0", value)
}

// f-1 (float32)
//
// Frequency 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF1(value float32) error {
	return e.Element.SetProperty("f-1", value)
}

// f-10 (float32)
//
// Frequency 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF10(value float32) error {
	return e.Element.SetProperty("f-10", value)
}

// f-11 (float32)
//
// Frequency 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF11(value float32) error {
	return e.Element.SetProperty("f-11", value)
}

// f-12 (float32)
//
// Frequency 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF12(value float32) error {
	return e.Element.SetProperty("f-12", value)
}

// f-13 (float32)
//
// Frequency 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF13(value float32) error {
	return e.Element.SetProperty("f-13", value)
}

// f-14 (float32)
//
// Frequency 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF14(value float32) error {
	return e.Element.SetProperty("f-14", value)
}

// f-15 (float32)
//
// Frequency 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF15(value float32) error {
	return e.Element.SetProperty("f-15", value)
}

// f-2 (float32)
//
// Frequency 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF2(value float32) error {
	return e.Element.SetProperty("f-2", value)
}

// f-3 (float32)
//
// Frequency 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF3(value float32) error {
	return e.Element.SetProperty("f-3", value)
}

// f-4 (float32)
//
// Frequency 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF4(value float32) error {
	return e.Element.SetProperty("f-4", value)
}

// f-5 (float32)
//
// Frequency 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF5(value float32) error {
	return e.Element.SetProperty("f-5", value)
}

// f-6 (float32)
//
// Frequency 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF6(value float32) error {
	return e.Element.SetProperty("f-6", value)
}

// f-7 (float32)
//
// Frequency 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF7(value float32) error {
	return e.Element.SetProperty("f-7", value)
}

// f-8 (float32)
//
// Frequency 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF8(value float32) error {
	return e.Element.SetProperty("f-8", value)
}

// f-9 (float32)
//
// Frequency 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetF9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetF9(value float32) error {
	return e.Element.SetProperty("f-9", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFftvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fftv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFftvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fftv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fm-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm0)
//
// Filter mode 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm0() (interface{}, error) {
	return e.Element.GetProperty("fm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm0(value interface{}) error {
	return e.Element.SetProperty("fm-0", value)
}

// fm-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm1)
//
// Filter mode 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm1() (interface{}, error) {
	return e.Element.GetProperty("fm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm1(value interface{}) error {
	return e.Element.SetProperty("fm-1", value)
}

// fm-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm10)
//
// Filter mode 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm10() (interface{}, error) {
	return e.Element.GetProperty("fm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm10(value interface{}) error {
	return e.Element.SetProperty("fm-10", value)
}

// fm-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm11)
//
// Filter mode 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm11() (interface{}, error) {
	return e.Element.GetProperty("fm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm11(value interface{}) error {
	return e.Element.SetProperty("fm-11", value)
}

// fm-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm12)
//
// Filter mode 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm12() (interface{}, error) {
	return e.Element.GetProperty("fm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm12(value interface{}) error {
	return e.Element.SetProperty("fm-12", value)
}

// fm-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm13)
//
// Filter mode 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm13() (interface{}, error) {
	return e.Element.GetProperty("fm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm13(value interface{}) error {
	return e.Element.SetProperty("fm-13", value)
}

// fm-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm14)
//
// Filter mode 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm14() (interface{}, error) {
	return e.Element.GetProperty("fm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm14(value interface{}) error {
	return e.Element.SetProperty("fm-14", value)
}

// fm-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm15)
//
// Filter mode 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm15() (interface{}, error) {
	return e.Element.GetProperty("fm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm15(value interface{}) error {
	return e.Element.SetProperty("fm-15", value)
}

// fm-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm2)
//
// Filter mode 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm2() (interface{}, error) {
	return e.Element.GetProperty("fm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm2(value interface{}) error {
	return e.Element.SetProperty("fm-2", value)
}

// fm-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm3)
//
// Filter mode 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm3() (interface{}, error) {
	return e.Element.GetProperty("fm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm3(value interface{}) error {
	return e.Element.SetProperty("fm-3", value)
}

// fm-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm4)
//
// Filter mode 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm4() (interface{}, error) {
	return e.Element.GetProperty("fm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm4(value interface{}) error {
	return e.Element.SetProperty("fm-4", value)
}

// fm-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm5)
//
// Filter mode 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm5() (interface{}, error) {
	return e.Element.GetProperty("fm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm5(value interface{}) error {
	return e.Element.SetProperty("fm-5", value)
}

// fm-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm6)
//
// Filter mode 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm6() (interface{}, error) {
	return e.Element.GetProperty("fm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm6(value interface{}) error {
	return e.Element.SetProperty("fm-6", value)
}

// fm-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm7)
//
// Filter mode 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm7() (interface{}, error) {
	return e.Element.GetProperty("fm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm7(value interface{}) error {
	return e.Element.SetProperty("fm-7", value)
}

// fm-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm8)
//
// Filter mode 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm8() (interface{}, error) {
	return e.Element.GetProperty("fm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm8(value interface{}) error {
	return e.Element.SetProperty("fm-8", value)
}

// fm-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofm9)
//
// Filter mode 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFm9() (interface{}, error) {
	return e.Element.GetProperty("fm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFm9(value interface{}) error {
	return e.Element.SetProperty("fm-9", value)
}

// frqs (float32)
//
// Frequency shift

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFrqs() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFrqs(value float32) error {
	return e.Element.SetProperty("frqs", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX16Stereofsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ft-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft0)
//
// Filter type 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt0() (interface{}, error) {
	return e.Element.GetProperty("ft-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt0(value interface{}) error {
	return e.Element.SetProperty("ft-0", value)
}

// ft-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft1)
//
// Filter type 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt1() (interface{}, error) {
	return e.Element.GetProperty("ft-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt1(value interface{}) error {
	return e.Element.SetProperty("ft-1", value)
}

// ft-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft10)
//
// Filter type 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt10() (interface{}, error) {
	return e.Element.GetProperty("ft-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt10(value interface{}) error {
	return e.Element.SetProperty("ft-10", value)
}

// ft-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft11)
//
// Filter type 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt11() (interface{}, error) {
	return e.Element.GetProperty("ft-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt11(value interface{}) error {
	return e.Element.SetProperty("ft-11", value)
}

// ft-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft12)
//
// Filter type 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt12() (interface{}, error) {
	return e.Element.GetProperty("ft-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt12(value interface{}) error {
	return e.Element.SetProperty("ft-12", value)
}

// ft-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft13)
//
// Filter type 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt13() (interface{}, error) {
	return e.Element.GetProperty("ft-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt13(value interface{}) error {
	return e.Element.SetProperty("ft-13", value)
}

// ft-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft14)
//
// Filter type 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt14() (interface{}, error) {
	return e.Element.GetProperty("ft-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt14(value interface{}) error {
	return e.Element.SetProperty("ft-14", value)
}

// ft-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft15)
//
// Filter type 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt15() (interface{}, error) {
	return e.Element.GetProperty("ft-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt15(value interface{}) error {
	return e.Element.SetProperty("ft-15", value)
}

// ft-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft2)
//
// Filter type 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt2() (interface{}, error) {
	return e.Element.GetProperty("ft-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt2(value interface{}) error {
	return e.Element.SetProperty("ft-2", value)
}

// ft-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft3)
//
// Filter type 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt3() (interface{}, error) {
	return e.Element.GetProperty("ft-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt3(value interface{}) error {
	return e.Element.SetProperty("ft-3", value)
}

// ft-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft4)
//
// Filter type 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt4() (interface{}, error) {
	return e.Element.GetProperty("ft-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt4(value interface{}) error {
	return e.Element.SetProperty("ft-4", value)
}

// ft-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft5)
//
// Filter type 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt5() (interface{}, error) {
	return e.Element.GetProperty("ft-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt5(value interface{}) error {
	return e.Element.SetProperty("ft-5", value)
}

// ft-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft6)
//
// Filter type 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt6() (interface{}, error) {
	return e.Element.GetProperty("ft-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt6(value interface{}) error {
	return e.Element.SetProperty("ft-6", value)
}

// ft-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft7)
//
// Filter type 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt7() (interface{}, error) {
	return e.Element.GetProperty("ft-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt7(value interface{}) error {
	return e.Element.SetProperty("ft-7", value)
}

// ft-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft8)
//
// Filter type 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt8() (interface{}, error) {
	return e.Element.GetProperty("ft-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt8(value interface{}) error {
	return e.Element.SetProperty("ft-8", value)
}

// ft-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereoft9)
//
// Filter type 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFt9() (interface{}, error) {
	return e.Element.GetProperty("ft-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetFt9(value interface{}) error {
	return e.Element.SetProperty("ft-9", value)
}

// fv-0 (bool)
//
// Filter visibility 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetFv9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG0(value float32) error {
	return e.Element.SetProperty("g-0", value)
}

// g-1 (float32)
//
// Gain 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG1(value float32) error {
	return e.Element.SetProperty("g-1", value)
}

// g-10 (float32)
//
// Gain 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG10(value float32) error {
	return e.Element.SetProperty("g-10", value)
}

// g-11 (float32)
//
// Gain 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG11(value float32) error {
	return e.Element.SetProperty("g-11", value)
}

// g-12 (float32)
//
// Gain 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG12(value float32) error {
	return e.Element.SetProperty("g-12", value)
}

// g-13 (float32)
//
// Gain 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG13(value float32) error {
	return e.Element.SetProperty("g-13", value)
}

// g-14 (float32)
//
// Gain 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG14(value float32) error {
	return e.Element.SetProperty("g-14", value)
}

// g-15 (float32)
//
// Gain 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG15(value float32) error {
	return e.Element.SetProperty("g-15", value)
}

// g-2 (float32)
//
// Gain 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG2(value float32) error {
	return e.Element.SetProperty("g-2", value)
}

// g-3 (float32)
//
// Gain 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG3(value float32) error {
	return e.Element.SetProperty("g-3", value)
}

// g-4 (float32)
//
// Gain 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG4(value float32) error {
	return e.Element.SetProperty("g-4", value)
}

// g-5 (float32)
//
// Gain 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG5(value float32) error {
	return e.Element.SetProperty("g-5", value)
}

// g-6 (float32)
//
// Gain 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG6(value float32) error {
	return e.Element.SetProperty("g-6", value)
}

// g-7 (float32)
//
// Gain 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG7(value float32) error {
	return e.Element.SetProperty("g-7", value)
}

// g-8 (float32)
//
// Gain 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG8(value float32) error {
	return e.Element.SetProperty("g-8", value)
}

// g-9 (float32)
//
// Gain 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetG9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetG9(value float32) error {
	return e.Element.SetProperty("g-9", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hue-0 (float32)
//
// Hue 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// hue-1 (float32)
//
// Hue 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue1(value float32) error {
	return e.Element.SetProperty("hue-1", value)
}

// hue-10 (float32)
//
// Hue 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue10(value float32) error {
	return e.Element.SetProperty("hue-10", value)
}

// hue-11 (float32)
//
// Hue 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue11(value float32) error {
	return e.Element.SetProperty("hue-11", value)
}

// hue-12 (float32)
//
// Hue 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue12(value float32) error {
	return e.Element.SetProperty("hue-12", value)
}

// hue-13 (float32)
//
// Hue 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue13(value float32) error {
	return e.Element.SetProperty("hue-13", value)
}

// hue-14 (float32)
//
// Hue 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue14(value float32) error {
	return e.Element.SetProperty("hue-14", value)
}

// hue-15 (float32)
//
// Hue 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue15(value float32) error {
	return e.Element.SetProperty("hue-15", value)
}

// hue-2 (float32)
//
// Hue 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue2(value float32) error {
	return e.Element.SetProperty("hue-2", value)
}

// hue-3 (float32)
//
// Hue 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue3(value float32) error {
	return e.Element.SetProperty("hue-3", value)
}

// hue-4 (float32)
//
// Hue 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue4(value float32) error {
	return e.Element.SetProperty("hue-4", value)
}

// hue-5 (float32)
//
// Hue 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue5(value float32) error {
	return e.Element.SetProperty("hue-5", value)
}

// hue-6 (float32)
//
// Hue 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue6(value float32) error {
	return e.Element.SetProperty("hue-6", value)
}

// hue-7 (float32)
//
// Hue 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue7(value float32) error {
	return e.Element.SetProperty("hue-7", value)
}

// hue-8 (float32)
//
// Hue 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue8(value float32) error {
	return e.Element.SetProperty("hue-8", value)
}

// hue-9 (float32)
//
// Hue 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetHue9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetHue9(value float32) error {
	return e.Element.SetProperty("hue-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetImr() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2ParaEqualizerX16Stereomode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ0(value float32) error {
	return e.Element.SetProperty("q-0", value)
}

// q-1 (float32)
//
// Quality factor 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ1(value float32) error {
	return e.Element.SetProperty("q-1", value)
}

// q-10 (float32)
//
// Quality factor 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ10(value float32) error {
	return e.Element.SetProperty("q-10", value)
}

// q-11 (float32)
//
// Quality factor 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ11(value float32) error {
	return e.Element.SetProperty("q-11", value)
}

// q-12 (float32)
//
// Quality factor 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ12(value float32) error {
	return e.Element.SetProperty("q-12", value)
}

// q-13 (float32)
//
// Quality factor 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ13(value float32) error {
	return e.Element.SetProperty("q-13", value)
}

// q-14 (float32)
//
// Quality factor 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ14(value float32) error {
	return e.Element.SetProperty("q-14", value)
}

// q-15 (float32)
//
// Quality factor 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ15(value float32) error {
	return e.Element.SetProperty("q-15", value)
}

// q-2 (float32)
//
// Quality factor 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ2(value float32) error {
	return e.Element.SetProperty("q-2", value)
}

// q-3 (float32)
//
// Quality factor 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ3(value float32) error {
	return e.Element.SetProperty("q-3", value)
}

// q-4 (float32)
//
// Quality factor 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ4(value float32) error {
	return e.Element.SetProperty("q-4", value)
}

// q-5 (float32)
//
// Quality factor 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ5(value float32) error {
	return e.Element.SetProperty("q-5", value)
}

// q-6 (float32)
//
// Quality factor 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ6(value float32) error {
	return e.Element.SetProperty("q-6", value)
}

// q-7 (float32)
//
// Quality factor 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ7(value float32) error {
	return e.Element.SetProperty("q-7", value)
}

// q-8 (float32)
//
// Quality factor 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ8(value float32) error {
	return e.Element.SetProperty("q-8", value)
}

// q-9 (float32)
//
// Quality factor 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetQ9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetQ9(value float32) error {
	return e.Element.SetProperty("q-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// s-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos0)
//
// Filter slope 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS0() (interface{}, error) {
	return e.Element.GetProperty("s-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS0(value interface{}) error {
	return e.Element.SetProperty("s-0", value)
}

// s-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos1)
//
// Filter slope 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS1() (interface{}, error) {
	return e.Element.GetProperty("s-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS1(value interface{}) error {
	return e.Element.SetProperty("s-1", value)
}

// s-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos10)
//
// Filter slope 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS10() (interface{}, error) {
	return e.Element.GetProperty("s-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS10(value interface{}) error {
	return e.Element.SetProperty("s-10", value)
}

// s-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos11)
//
// Filter slope 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS11() (interface{}, error) {
	return e.Element.GetProperty("s-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS11(value interface{}) error {
	return e.Element.SetProperty("s-11", value)
}

// s-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos12)
//
// Filter slope 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS12() (interface{}, error) {
	return e.Element.GetProperty("s-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS12(value interface{}) error {
	return e.Element.SetProperty("s-12", value)
}

// s-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos13)
//
// Filter slope 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS13() (interface{}, error) {
	return e.Element.GetProperty("s-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS13(value interface{}) error {
	return e.Element.SetProperty("s-13", value)
}

// s-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos14)
//
// Filter slope 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS14() (interface{}, error) {
	return e.Element.GetProperty("s-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS14(value interface{}) error {
	return e.Element.SetProperty("s-14", value)
}

// s-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos15)
//
// Filter slope 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS15() (interface{}, error) {
	return e.Element.GetProperty("s-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS15(value interface{}) error {
	return e.Element.SetProperty("s-15", value)
}

// s-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos2)
//
// Filter slope 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS2() (interface{}, error) {
	return e.Element.GetProperty("s-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS2(value interface{}) error {
	return e.Element.SetProperty("s-2", value)
}

// s-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos3)
//
// Filter slope 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS3() (interface{}, error) {
	return e.Element.GetProperty("s-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS3(value interface{}) error {
	return e.Element.SetProperty("s-3", value)
}

// s-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos4)
//
// Filter slope 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS4() (interface{}, error) {
	return e.Element.GetProperty("s-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS4(value interface{}) error {
	return e.Element.SetProperty("s-4", value)
}

// s-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos5)
//
// Filter slope 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS5() (interface{}, error) {
	return e.Element.GetProperty("s-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS5(value interface{}) error {
	return e.Element.SetProperty("s-5", value)
}

// s-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos6)
//
// Filter slope 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS6() (interface{}, error) {
	return e.Element.GetProperty("s-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS6(value interface{}) error {
	return e.Element.SetProperty("s-6", value)
}

// s-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos7)
//
// Filter slope 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS7() (interface{}, error) {
	return e.Element.GetProperty("s-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS7(value interface{}) error {
	return e.Element.SetProperty("s-7", value)
}

// s-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos8)
//
// Filter slope 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS8() (interface{}, error) {
	return e.Element.GetProperty("s-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS8(value interface{}) error {
	return e.Element.SetProperty("s-8", value)
}

// s-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Stereos9)
//
// Filter slope 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetS9() (interface{}, error) {
	return e.Element.GetProperty("s-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetS9(value interface{}) error {
	return e.Element.SetProperty("s-9", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetSmr() (float32, error) {
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

// xm-0 (bool)
//
// Filter mute 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm0(value bool) error {
	return e.Element.SetProperty("xm-0", value)
}

// xm-1 (bool)
//
// Filter mute 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm1(value bool) error {
	return e.Element.SetProperty("xm-1", value)
}

// xm-10 (bool)
//
// Filter mute 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm10(value bool) error {
	return e.Element.SetProperty("xm-10", value)
}

// xm-11 (bool)
//
// Filter mute 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm11(value bool) error {
	return e.Element.SetProperty("xm-11", value)
}

// xm-12 (bool)
//
// Filter mute 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm12(value bool) error {
	return e.Element.SetProperty("xm-12", value)
}

// xm-13 (bool)
//
// Filter mute 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm13(value bool) error {
	return e.Element.SetProperty("xm-13", value)
}

// xm-14 (bool)
//
// Filter mute 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm14(value bool) error {
	return e.Element.SetProperty("xm-14", value)
}

// xm-15 (bool)
//
// Filter mute 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm15(value bool) error {
	return e.Element.SetProperty("xm-15", value)
}

// xm-2 (bool)
//
// Filter mute 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm2(value bool) error {
	return e.Element.SetProperty("xm-2", value)
}

// xm-3 (bool)
//
// Filter mute 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm3(value bool) error {
	return e.Element.SetProperty("xm-3", value)
}

// xm-4 (bool)
//
// Filter mute 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm4(value bool) error {
	return e.Element.SetProperty("xm-4", value)
}

// xm-5 (bool)
//
// Filter mute 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm5(value bool) error {
	return e.Element.SetProperty("xm-5", value)
}

// xm-6 (bool)
//
// Filter mute 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm6(value bool) error {
	return e.Element.SetProperty("xm-6", value)
}

// xm-7 (bool)
//
// Filter mute 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm7(value bool) error {
	return e.Element.SetProperty("xm-7", value)
}

// xm-8 (bool)
//
// Filter mute 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm8(value bool) error {
	return e.Element.SetProperty("xm-8", value)
}

// xm-9 (bool)
//
// Filter mute 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXm9(value bool) error {
	return e.Element.SetProperty("xm-9", value)
}

// xs-0 (bool)
//
// Filter solo 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs0(value bool) error {
	return e.Element.SetProperty("xs-0", value)
}

// xs-1 (bool)
//
// Filter solo 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs1(value bool) error {
	return e.Element.SetProperty("xs-1", value)
}

// xs-10 (bool)
//
// Filter solo 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs10(value bool) error {
	return e.Element.SetProperty("xs-10", value)
}

// xs-11 (bool)
//
// Filter solo 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs11(value bool) error {
	return e.Element.SetProperty("xs-11", value)
}

// xs-12 (bool)
//
// Filter solo 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs12(value bool) error {
	return e.Element.SetProperty("xs-12", value)
}

// xs-13 (bool)
//
// Filter solo 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs13(value bool) error {
	return e.Element.SetProperty("xs-13", value)
}

// xs-14 (bool)
//
// Filter solo 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs14(value bool) error {
	return e.Element.SetProperty("xs-14", value)
}

// xs-15 (bool)
//
// Filter solo 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs15(value bool) error {
	return e.Element.SetProperty("xs-15", value)
}

// xs-2 (bool)
//
// Filter solo 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs2(value bool) error {
	return e.Element.SetProperty("xs-2", value)
}

// xs-3 (bool)
//
// Filter solo 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs3(value bool) error {
	return e.Element.SetProperty("xs-3", value)
}

// xs-4 (bool)
//
// Filter solo 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs4(value bool) error {
	return e.Element.SetProperty("xs-4", value)
}

// xs-5 (bool)
//
// Filter solo 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs5(value bool) error {
	return e.Element.SetProperty("xs-5", value)
}

// xs-6 (bool)
//
// Filter solo 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs6(value bool) error {
	return e.Element.SetProperty("xs-6", value)
}

// xs-7 (bool)
//
// Filter solo 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs7(value bool) error {
	return e.Element.SetProperty("xs-7", value)
}

// xs-8 (bool)
//
// Filter solo 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs8(value bool) error {
	return e.Element.SetProperty("xs-8", value)
}

// xs-9 (bool)
//
// Filter solo 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetXs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetXs9(value bool) error {
	return e.Element.SetProperty("xs-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Stereo) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm1ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm11ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft0Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos13X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos13X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos13X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos13X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos4X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos4X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos4X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos4X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm2ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft10Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft11Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft12Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft6Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos2X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos2X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos2X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos2X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos5X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos5X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos5X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos5X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm10ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm5ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm8ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft1Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft8Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm6ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX16StereofselFilters07 LspPlugInPluginsLv2ParaEqualizerX16Stereofsel = "Filters 0-7" // Filters 0-7 (0) – Filters 0-7
	LspPlugInPluginsLv2ParaEqualizerX16StereofselFilters815 LspPlugInPluginsLv2ParaEqualizerX16Stereofsel = "Filters 8-15" // Filters 8-15 (1) – Filters 8-15
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft2Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft5Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofft string

const (
	LspPlugInPluginsLv2ParaEqualizerX16StereofftOff LspPlugInPluginsLv2ParaEqualizerX16Stereofft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16StereofftPostEq LspPlugInPluginsLv2ParaEqualizerX16Stereofft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX16StereofftPreEq LspPlugInPluginsLv2ParaEqualizerX16Stereofft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm13ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm14ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm4ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft14Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft4Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos10X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos10X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos10X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos10X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm15ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft7Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereomode string

const (
	LspPlugInPluginsLv2ParaEqualizerX16StereomodeIir LspPlugInPluginsLv2ParaEqualizerX16Stereomode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX16StereomodeFir LspPlugInPluginsLv2ParaEqualizerX16Stereomode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX16StereomodeFft LspPlugInPluginsLv2ParaEqualizerX16Stereomode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos0X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos0X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos0X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos0X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos15X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos15X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos15X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos15X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos3X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos3X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos3X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos3X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos9X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos9X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos9X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos9X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm0ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm12ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft9Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos6X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos6X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos6X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos6X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos8X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos8X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos8X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos8X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm3ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm7ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9RlcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9RlcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9BwcBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9BwcMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9LrxBt LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9LrxMt LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Stereofm9ApoDr LspPlugInPluginsLv2ParaEqualizerX16Stereofm9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft13Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft15Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3Off LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3Bell LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3HiPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3HiShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3LoPass LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3LoShelf LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3Notch LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3Resonance LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Stereoft3Allpass LspPlugInPluginsLv2ParaEqualizerX16Stereoft3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos1X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos1X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos1X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos1X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos11X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos11X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos11X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos11X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos12X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos12X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos12X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos12X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos14X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos14X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos14X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos14X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Stereos7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Stereos7X1 LspPlugInPluginsLv2ParaEqualizerX16Stereos7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Stereos7X2 LspPlugInPluginsLv2ParaEqualizerX16Stereos7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Stereos7X3 LspPlugInPluginsLv2ParaEqualizerX16Stereos7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Stereos7X4 LspPlugInPluginsLv2ParaEqualizerX16Stereos7 = "x4" // x4 (3) – x4
)

