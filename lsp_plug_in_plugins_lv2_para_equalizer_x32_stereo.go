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

type LspPlugInPluginsLv2ParaEqualizerX32Stereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX32Stereo() (*LspPlugInPluginsLv2ParaEqualizerX32Stereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x32-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX32StereoWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX32Stereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x32-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// f-0 (float32)
//
// Frequency 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF0(value float32) error {
	return e.Element.SetProperty("f-0", value)
}

// f-1 (float32)
//
// Frequency 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF1(value float32) error {
	return e.Element.SetProperty("f-1", value)
}

// f-10 (float32)
//
// Frequency 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF10(value float32) error {
	return e.Element.SetProperty("f-10", value)
}

// f-11 (float32)
//
// Frequency 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF11(value float32) error {
	return e.Element.SetProperty("f-11", value)
}

// f-12 (float32)
//
// Frequency 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF12(value float32) error {
	return e.Element.SetProperty("f-12", value)
}

// f-13 (float32)
//
// Frequency 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF13(value float32) error {
	return e.Element.SetProperty("f-13", value)
}

// f-14 (float32)
//
// Frequency 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF14(value float32) error {
	return e.Element.SetProperty("f-14", value)
}

// f-15 (float32)
//
// Frequency 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF15(value float32) error {
	return e.Element.SetProperty("f-15", value)
}

// f-16 (float32)
//
// Frequency 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF16() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF16(value float32) error {
	return e.Element.SetProperty("f-16", value)
}

// f-17 (float32)
//
// Frequency 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF17() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF17(value float32) error {
	return e.Element.SetProperty("f-17", value)
}

// f-18 (float32)
//
// Frequency 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF18() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF18(value float32) error {
	return e.Element.SetProperty("f-18", value)
}

// f-19 (float32)
//
// Frequency 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF19() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF19(value float32) error {
	return e.Element.SetProperty("f-19", value)
}

// f-2 (float32)
//
// Frequency 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF2(value float32) error {
	return e.Element.SetProperty("f-2", value)
}

// f-20 (float32)
//
// Frequency 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF20() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF20(value float32) error {
	return e.Element.SetProperty("f-20", value)
}

// f-21 (float32)
//
// Frequency 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF21() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF21(value float32) error {
	return e.Element.SetProperty("f-21", value)
}

// f-22 (float32)
//
// Frequency 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF22() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF22(value float32) error {
	return e.Element.SetProperty("f-22", value)
}

// f-23 (float32)
//
// Frequency 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF23() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF23(value float32) error {
	return e.Element.SetProperty("f-23", value)
}

// f-24 (float32)
//
// Frequency 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF24() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF24(value float32) error {
	return e.Element.SetProperty("f-24", value)
}

// f-25 (float32)
//
// Frequency 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF25() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF25(value float32) error {
	return e.Element.SetProperty("f-25", value)
}

// f-26 (float32)
//
// Frequency 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF26() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF26(value float32) error {
	return e.Element.SetProperty("f-26", value)
}

// f-27 (float32)
//
// Frequency 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF27() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF27(value float32) error {
	return e.Element.SetProperty("f-27", value)
}

// f-28 (float32)
//
// Frequency 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF28() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF28(value float32) error {
	return e.Element.SetProperty("f-28", value)
}

// f-29 (float32)
//
// Frequency 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF29() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF29(value float32) error {
	return e.Element.SetProperty("f-29", value)
}

// f-3 (float32)
//
// Frequency 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF3(value float32) error {
	return e.Element.SetProperty("f-3", value)
}

// f-30 (float32)
//
// Frequency 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF30() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF30(value float32) error {
	return e.Element.SetProperty("f-30", value)
}

// f-31 (float32)
//
// Frequency 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF31() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF31(value float32) error {
	return e.Element.SetProperty("f-31", value)
}

// f-4 (float32)
//
// Frequency 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF4(value float32) error {
	return e.Element.SetProperty("f-4", value)
}

// f-5 (float32)
//
// Frequency 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF5(value float32) error {
	return e.Element.SetProperty("f-5", value)
}

// f-6 (float32)
//
// Frequency 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF6(value float32) error {
	return e.Element.SetProperty("f-6", value)
}

// f-7 (float32)
//
// Frequency 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF7(value float32) error {
	return e.Element.SetProperty("f-7", value)
}

// f-8 (float32)
//
// Frequency 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF8(value float32) error {
	return e.Element.SetProperty("f-8", value)
}

// f-9 (float32)
//
// Frequency 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetF9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetF9(value float32) error {
	return e.Element.SetProperty("f-9", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFftvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFftvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fm-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm0)
//
// Filter mode 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm0() (interface{}, error) {
	return e.Element.GetProperty("fm-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm0(value interface{}) error {
	return e.Element.SetProperty("fm-0", value)
}

// fm-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm1)
//
// Filter mode 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm1() (interface{}, error) {
	return e.Element.GetProperty("fm-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm1(value interface{}) error {
	return e.Element.SetProperty("fm-1", value)
}

// fm-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm10)
//
// Filter mode 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm10() (interface{}, error) {
	return e.Element.GetProperty("fm-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm10(value interface{}) error {
	return e.Element.SetProperty("fm-10", value)
}

// fm-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm11)
//
// Filter mode 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm11() (interface{}, error) {
	return e.Element.GetProperty("fm-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm11(value interface{}) error {
	return e.Element.SetProperty("fm-11", value)
}

// fm-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm12)
//
// Filter mode 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm12() (interface{}, error) {
	return e.Element.GetProperty("fm-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm12(value interface{}) error {
	return e.Element.SetProperty("fm-12", value)
}

// fm-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm13)
//
// Filter mode 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm13() (interface{}, error) {
	return e.Element.GetProperty("fm-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm13(value interface{}) error {
	return e.Element.SetProperty("fm-13", value)
}

// fm-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm14)
//
// Filter mode 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm14() (interface{}, error) {
	return e.Element.GetProperty("fm-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm14(value interface{}) error {
	return e.Element.SetProperty("fm-14", value)
}

// fm-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm15)
//
// Filter mode 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm15() (interface{}, error) {
	return e.Element.GetProperty("fm-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm15(value interface{}) error {
	return e.Element.SetProperty("fm-15", value)
}

// fm-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm16)
//
// Filter mode 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm16() (interface{}, error) {
	return e.Element.GetProperty("fm-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm16(value interface{}) error {
	return e.Element.SetProperty("fm-16", value)
}

// fm-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm17)
//
// Filter mode 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm17() (interface{}, error) {
	return e.Element.GetProperty("fm-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm17(value interface{}) error {
	return e.Element.SetProperty("fm-17", value)
}

// fm-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm18)
//
// Filter mode 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm18() (interface{}, error) {
	return e.Element.GetProperty("fm-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm18(value interface{}) error {
	return e.Element.SetProperty("fm-18", value)
}

// fm-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm19)
//
// Filter mode 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm19() (interface{}, error) {
	return e.Element.GetProperty("fm-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm19(value interface{}) error {
	return e.Element.SetProperty("fm-19", value)
}

// fm-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm2)
//
// Filter mode 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm2() (interface{}, error) {
	return e.Element.GetProperty("fm-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm2(value interface{}) error {
	return e.Element.SetProperty("fm-2", value)
}

// fm-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm20)
//
// Filter mode 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm20() (interface{}, error) {
	return e.Element.GetProperty("fm-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm20(value interface{}) error {
	return e.Element.SetProperty("fm-20", value)
}

// fm-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm21)
//
// Filter mode 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm21() (interface{}, error) {
	return e.Element.GetProperty("fm-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm21(value interface{}) error {
	return e.Element.SetProperty("fm-21", value)
}

// fm-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm22)
//
// Filter mode 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm22() (interface{}, error) {
	return e.Element.GetProperty("fm-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm22(value interface{}) error {
	return e.Element.SetProperty("fm-22", value)
}

// fm-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm23)
//
// Filter mode 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm23() (interface{}, error) {
	return e.Element.GetProperty("fm-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm23(value interface{}) error {
	return e.Element.SetProperty("fm-23", value)
}

// fm-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm24)
//
// Filter mode 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm24() (interface{}, error) {
	return e.Element.GetProperty("fm-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm24(value interface{}) error {
	return e.Element.SetProperty("fm-24", value)
}

// fm-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm25)
//
// Filter mode 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm25() (interface{}, error) {
	return e.Element.GetProperty("fm-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm25(value interface{}) error {
	return e.Element.SetProperty("fm-25", value)
}

// fm-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm26)
//
// Filter mode 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm26() (interface{}, error) {
	return e.Element.GetProperty("fm-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm26(value interface{}) error {
	return e.Element.SetProperty("fm-26", value)
}

// fm-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm27)
//
// Filter mode 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm27() (interface{}, error) {
	return e.Element.GetProperty("fm-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm27(value interface{}) error {
	return e.Element.SetProperty("fm-27", value)
}

// fm-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm28)
//
// Filter mode 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm28() (interface{}, error) {
	return e.Element.GetProperty("fm-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm28(value interface{}) error {
	return e.Element.SetProperty("fm-28", value)
}

// fm-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm29)
//
// Filter mode 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm29() (interface{}, error) {
	return e.Element.GetProperty("fm-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm29(value interface{}) error {
	return e.Element.SetProperty("fm-29", value)
}

// fm-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm3)
//
// Filter mode 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm3() (interface{}, error) {
	return e.Element.GetProperty("fm-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm3(value interface{}) error {
	return e.Element.SetProperty("fm-3", value)
}

// fm-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm30)
//
// Filter mode 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm30() (interface{}, error) {
	return e.Element.GetProperty("fm-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm30(value interface{}) error {
	return e.Element.SetProperty("fm-30", value)
}

// fm-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm31)
//
// Filter mode 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm31() (interface{}, error) {
	return e.Element.GetProperty("fm-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm31(value interface{}) error {
	return e.Element.SetProperty("fm-31", value)
}

// fm-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm4)
//
// Filter mode 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm4() (interface{}, error) {
	return e.Element.GetProperty("fm-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm4(value interface{}) error {
	return e.Element.SetProperty("fm-4", value)
}

// fm-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm5)
//
// Filter mode 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm5() (interface{}, error) {
	return e.Element.GetProperty("fm-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm5(value interface{}) error {
	return e.Element.SetProperty("fm-5", value)
}

// fm-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm6)
//
// Filter mode 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm6() (interface{}, error) {
	return e.Element.GetProperty("fm-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm6(value interface{}) error {
	return e.Element.SetProperty("fm-6", value)
}

// fm-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm7)
//
// Filter mode 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm7() (interface{}, error) {
	return e.Element.GetProperty("fm-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm7(value interface{}) error {
	return e.Element.SetProperty("fm-7", value)
}

// fm-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm8)
//
// Filter mode 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm8() (interface{}, error) {
	return e.Element.GetProperty("fm-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm8(value interface{}) error {
	return e.Element.SetProperty("fm-8", value)
}

// fm-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofm9)
//
// Filter mode 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFm9() (interface{}, error) {
	return e.Element.GetProperty("fm-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFm9(value interface{}) error {
	return e.Element.SetProperty("fm-9", value)
}

// frqs (float32)
//
// Frequency shift

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFrqs() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFrqs(value float32) error {
	return e.Element.SetProperty("frqs", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX32Stereofsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ft-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft0)
//
// Filter type 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt0() (interface{}, error) {
	return e.Element.GetProperty("ft-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt0(value interface{}) error {
	return e.Element.SetProperty("ft-0", value)
}

// ft-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft1)
//
// Filter type 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt1() (interface{}, error) {
	return e.Element.GetProperty("ft-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt1(value interface{}) error {
	return e.Element.SetProperty("ft-1", value)
}

// ft-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft10)
//
// Filter type 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt10() (interface{}, error) {
	return e.Element.GetProperty("ft-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt10(value interface{}) error {
	return e.Element.SetProperty("ft-10", value)
}

// ft-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft11)
//
// Filter type 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt11() (interface{}, error) {
	return e.Element.GetProperty("ft-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt11(value interface{}) error {
	return e.Element.SetProperty("ft-11", value)
}

// ft-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft12)
//
// Filter type 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt12() (interface{}, error) {
	return e.Element.GetProperty("ft-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt12(value interface{}) error {
	return e.Element.SetProperty("ft-12", value)
}

// ft-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft13)
//
// Filter type 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt13() (interface{}, error) {
	return e.Element.GetProperty("ft-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt13(value interface{}) error {
	return e.Element.SetProperty("ft-13", value)
}

// ft-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft14)
//
// Filter type 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt14() (interface{}, error) {
	return e.Element.GetProperty("ft-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt14(value interface{}) error {
	return e.Element.SetProperty("ft-14", value)
}

// ft-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft15)
//
// Filter type 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt15() (interface{}, error) {
	return e.Element.GetProperty("ft-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt15(value interface{}) error {
	return e.Element.SetProperty("ft-15", value)
}

// ft-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft16)
//
// Filter type 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt16() (interface{}, error) {
	return e.Element.GetProperty("ft-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt16(value interface{}) error {
	return e.Element.SetProperty("ft-16", value)
}

// ft-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft17)
//
// Filter type 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt17() (interface{}, error) {
	return e.Element.GetProperty("ft-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt17(value interface{}) error {
	return e.Element.SetProperty("ft-17", value)
}

// ft-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft18)
//
// Filter type 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt18() (interface{}, error) {
	return e.Element.GetProperty("ft-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt18(value interface{}) error {
	return e.Element.SetProperty("ft-18", value)
}

// ft-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft19)
//
// Filter type 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt19() (interface{}, error) {
	return e.Element.GetProperty("ft-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt19(value interface{}) error {
	return e.Element.SetProperty("ft-19", value)
}

// ft-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft2)
//
// Filter type 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt2() (interface{}, error) {
	return e.Element.GetProperty("ft-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt2(value interface{}) error {
	return e.Element.SetProperty("ft-2", value)
}

// ft-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft20)
//
// Filter type 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt20() (interface{}, error) {
	return e.Element.GetProperty("ft-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt20(value interface{}) error {
	return e.Element.SetProperty("ft-20", value)
}

// ft-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft21)
//
// Filter type 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt21() (interface{}, error) {
	return e.Element.GetProperty("ft-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt21(value interface{}) error {
	return e.Element.SetProperty("ft-21", value)
}

// ft-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft22)
//
// Filter type 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt22() (interface{}, error) {
	return e.Element.GetProperty("ft-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt22(value interface{}) error {
	return e.Element.SetProperty("ft-22", value)
}

// ft-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft23)
//
// Filter type 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt23() (interface{}, error) {
	return e.Element.GetProperty("ft-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt23(value interface{}) error {
	return e.Element.SetProperty("ft-23", value)
}

// ft-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft24)
//
// Filter type 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt24() (interface{}, error) {
	return e.Element.GetProperty("ft-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt24(value interface{}) error {
	return e.Element.SetProperty("ft-24", value)
}

// ft-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft25)
//
// Filter type 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt25() (interface{}, error) {
	return e.Element.GetProperty("ft-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt25(value interface{}) error {
	return e.Element.SetProperty("ft-25", value)
}

// ft-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft26)
//
// Filter type 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt26() (interface{}, error) {
	return e.Element.GetProperty("ft-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt26(value interface{}) error {
	return e.Element.SetProperty("ft-26", value)
}

// ft-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft27)
//
// Filter type 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt27() (interface{}, error) {
	return e.Element.GetProperty("ft-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt27(value interface{}) error {
	return e.Element.SetProperty("ft-27", value)
}

// ft-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft28)
//
// Filter type 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt28() (interface{}, error) {
	return e.Element.GetProperty("ft-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt28(value interface{}) error {
	return e.Element.SetProperty("ft-28", value)
}

// ft-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft29)
//
// Filter type 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt29() (interface{}, error) {
	return e.Element.GetProperty("ft-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt29(value interface{}) error {
	return e.Element.SetProperty("ft-29", value)
}

// ft-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft3)
//
// Filter type 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt3() (interface{}, error) {
	return e.Element.GetProperty("ft-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt3(value interface{}) error {
	return e.Element.SetProperty("ft-3", value)
}

// ft-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft30)
//
// Filter type 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt30() (interface{}, error) {
	return e.Element.GetProperty("ft-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt30(value interface{}) error {
	return e.Element.SetProperty("ft-30", value)
}

// ft-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft31)
//
// Filter type 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt31() (interface{}, error) {
	return e.Element.GetProperty("ft-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt31(value interface{}) error {
	return e.Element.SetProperty("ft-31", value)
}

// ft-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft4)
//
// Filter type 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt4() (interface{}, error) {
	return e.Element.GetProperty("ft-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt4(value interface{}) error {
	return e.Element.SetProperty("ft-4", value)
}

// ft-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft5)
//
// Filter type 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt5() (interface{}, error) {
	return e.Element.GetProperty("ft-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt5(value interface{}) error {
	return e.Element.SetProperty("ft-5", value)
}

// ft-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft6)
//
// Filter type 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt6() (interface{}, error) {
	return e.Element.GetProperty("ft-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt6(value interface{}) error {
	return e.Element.SetProperty("ft-6", value)
}

// ft-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft7)
//
// Filter type 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt7() (interface{}, error) {
	return e.Element.GetProperty("ft-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt7(value interface{}) error {
	return e.Element.SetProperty("ft-7", value)
}

// ft-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft8)
//
// Filter type 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt8() (interface{}, error) {
	return e.Element.GetProperty("ft-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt8(value interface{}) error {
	return e.Element.SetProperty("ft-8", value)
}

// ft-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereoft9)
//
// Filter type 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFt9() (interface{}, error) {
	return e.Element.GetProperty("ft-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetFt9(value interface{}) error {
	return e.Element.SetProperty("ft-9", value)
}

// fv-0 (bool)
//
// Filter visibility 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv16() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv17() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv18() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv19() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv20() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv21() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv22() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv23() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv24() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv25() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv26() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv27() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv28() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv29() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv30() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv31() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetFv9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG0(value float32) error {
	return e.Element.SetProperty("g-0", value)
}

// g-1 (float32)
//
// Gain 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG1(value float32) error {
	return e.Element.SetProperty("g-1", value)
}

// g-10 (float32)
//
// Gain 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG10(value float32) error {
	return e.Element.SetProperty("g-10", value)
}

// g-11 (float32)
//
// Gain 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG11(value float32) error {
	return e.Element.SetProperty("g-11", value)
}

// g-12 (float32)
//
// Gain 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG12(value float32) error {
	return e.Element.SetProperty("g-12", value)
}

// g-13 (float32)
//
// Gain 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG13(value float32) error {
	return e.Element.SetProperty("g-13", value)
}

// g-14 (float32)
//
// Gain 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG14(value float32) error {
	return e.Element.SetProperty("g-14", value)
}

// g-15 (float32)
//
// Gain 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG15(value float32) error {
	return e.Element.SetProperty("g-15", value)
}

// g-16 (float32)
//
// Gain 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG16() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG16(value float32) error {
	return e.Element.SetProperty("g-16", value)
}

// g-17 (float32)
//
// Gain 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG17() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG17(value float32) error {
	return e.Element.SetProperty("g-17", value)
}

// g-18 (float32)
//
// Gain 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG18() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG18(value float32) error {
	return e.Element.SetProperty("g-18", value)
}

// g-19 (float32)
//
// Gain 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG19() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG19(value float32) error {
	return e.Element.SetProperty("g-19", value)
}

// g-2 (float32)
//
// Gain 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG2(value float32) error {
	return e.Element.SetProperty("g-2", value)
}

// g-20 (float32)
//
// Gain 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG20() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG20(value float32) error {
	return e.Element.SetProperty("g-20", value)
}

// g-21 (float32)
//
// Gain 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG21() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG21(value float32) error {
	return e.Element.SetProperty("g-21", value)
}

// g-22 (float32)
//
// Gain 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG22() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG22(value float32) error {
	return e.Element.SetProperty("g-22", value)
}

// g-23 (float32)
//
// Gain 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG23() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG23(value float32) error {
	return e.Element.SetProperty("g-23", value)
}

// g-24 (float32)
//
// Gain 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG24() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG24(value float32) error {
	return e.Element.SetProperty("g-24", value)
}

// g-25 (float32)
//
// Gain 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG25() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG25(value float32) error {
	return e.Element.SetProperty("g-25", value)
}

// g-26 (float32)
//
// Gain 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG26() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG26(value float32) error {
	return e.Element.SetProperty("g-26", value)
}

// g-27 (float32)
//
// Gain 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG27() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG27(value float32) error {
	return e.Element.SetProperty("g-27", value)
}

// g-28 (float32)
//
// Gain 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG28() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG28(value float32) error {
	return e.Element.SetProperty("g-28", value)
}

// g-29 (float32)
//
// Gain 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG29() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG29(value float32) error {
	return e.Element.SetProperty("g-29", value)
}

// g-3 (float32)
//
// Gain 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG3(value float32) error {
	return e.Element.SetProperty("g-3", value)
}

// g-30 (float32)
//
// Gain 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG30() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG30(value float32) error {
	return e.Element.SetProperty("g-30", value)
}

// g-31 (float32)
//
// Gain 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG31() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG31(value float32) error {
	return e.Element.SetProperty("g-31", value)
}

// g-4 (float32)
//
// Gain 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG4(value float32) error {
	return e.Element.SetProperty("g-4", value)
}

// g-5 (float32)
//
// Gain 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG5(value float32) error {
	return e.Element.SetProperty("g-5", value)
}

// g-6 (float32)
//
// Gain 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG6(value float32) error {
	return e.Element.SetProperty("g-6", value)
}

// g-7 (float32)
//
// Gain 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG7(value float32) error {
	return e.Element.SetProperty("g-7", value)
}

// g-8 (float32)
//
// Gain 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG8(value float32) error {
	return e.Element.SetProperty("g-8", value)
}

// g-9 (float32)
//
// Gain 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetG9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetG9(value float32) error {
	return e.Element.SetProperty("g-9", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hue-0 (float32)
//
// Hue 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// hue-1 (float32)
//
// Hue 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue1(value float32) error {
	return e.Element.SetProperty("hue-1", value)
}

// hue-10 (float32)
//
// Hue 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue10(value float32) error {
	return e.Element.SetProperty("hue-10", value)
}

// hue-11 (float32)
//
// Hue 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue11(value float32) error {
	return e.Element.SetProperty("hue-11", value)
}

// hue-12 (float32)
//
// Hue 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue12(value float32) error {
	return e.Element.SetProperty("hue-12", value)
}

// hue-13 (float32)
//
// Hue 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue13(value float32) error {
	return e.Element.SetProperty("hue-13", value)
}

// hue-14 (float32)
//
// Hue 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue14(value float32) error {
	return e.Element.SetProperty("hue-14", value)
}

// hue-15 (float32)
//
// Hue 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue15(value float32) error {
	return e.Element.SetProperty("hue-15", value)
}

// hue-16 (float32)
//
// Hue 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue16() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue16(value float32) error {
	return e.Element.SetProperty("hue-16", value)
}

// hue-17 (float32)
//
// Hue 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue17() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue17(value float32) error {
	return e.Element.SetProperty("hue-17", value)
}

// hue-18 (float32)
//
// Hue 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue18() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue18(value float32) error {
	return e.Element.SetProperty("hue-18", value)
}

// hue-19 (float32)
//
// Hue 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue19() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue19(value float32) error {
	return e.Element.SetProperty("hue-19", value)
}

// hue-2 (float32)
//
// Hue 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue2(value float32) error {
	return e.Element.SetProperty("hue-2", value)
}

// hue-20 (float32)
//
// Hue 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue20() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue20(value float32) error {
	return e.Element.SetProperty("hue-20", value)
}

// hue-21 (float32)
//
// Hue 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue21() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue21(value float32) error {
	return e.Element.SetProperty("hue-21", value)
}

// hue-22 (float32)
//
// Hue 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue22() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue22(value float32) error {
	return e.Element.SetProperty("hue-22", value)
}

// hue-23 (float32)
//
// Hue 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue23() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue23(value float32) error {
	return e.Element.SetProperty("hue-23", value)
}

// hue-24 (float32)
//
// Hue 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue24() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue24(value float32) error {
	return e.Element.SetProperty("hue-24", value)
}

// hue-25 (float32)
//
// Hue 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue25() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue25(value float32) error {
	return e.Element.SetProperty("hue-25", value)
}

// hue-26 (float32)
//
// Hue 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue26() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue26(value float32) error {
	return e.Element.SetProperty("hue-26", value)
}

// hue-27 (float32)
//
// Hue 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue27() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue27(value float32) error {
	return e.Element.SetProperty("hue-27", value)
}

// hue-28 (float32)
//
// Hue 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue28() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue28(value float32) error {
	return e.Element.SetProperty("hue-28", value)
}

// hue-29 (float32)
//
// Hue 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue29() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue29(value float32) error {
	return e.Element.SetProperty("hue-29", value)
}

// hue-3 (float32)
//
// Hue 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue3(value float32) error {
	return e.Element.SetProperty("hue-3", value)
}

// hue-30 (float32)
//
// Hue 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue30() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue30(value float32) error {
	return e.Element.SetProperty("hue-30", value)
}

// hue-31 (float32)
//
// Hue 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue31() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue31(value float32) error {
	return e.Element.SetProperty("hue-31", value)
}

// hue-4 (float32)
//
// Hue 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue4(value float32) error {
	return e.Element.SetProperty("hue-4", value)
}

// hue-5 (float32)
//
// Hue 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue5(value float32) error {
	return e.Element.SetProperty("hue-5", value)
}

// hue-6 (float32)
//
// Hue 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue6(value float32) error {
	return e.Element.SetProperty("hue-6", value)
}

// hue-7 (float32)
//
// Hue 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue7(value float32) error {
	return e.Element.SetProperty("hue-7", value)
}

// hue-8 (float32)
//
// Hue 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue8(value float32) error {
	return e.Element.SetProperty("hue-8", value)
}

// hue-9 (float32)
//
// Hue 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetHue9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetHue9(value float32) error {
	return e.Element.SetProperty("hue-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetImr() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2ParaEqualizerX32Stereomode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ0(value float32) error {
	return e.Element.SetProperty("q-0", value)
}

// q-1 (float32)
//
// Quality factor 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ1(value float32) error {
	return e.Element.SetProperty("q-1", value)
}

// q-10 (float32)
//
// Quality factor 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ10(value float32) error {
	return e.Element.SetProperty("q-10", value)
}

// q-11 (float32)
//
// Quality factor 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ11(value float32) error {
	return e.Element.SetProperty("q-11", value)
}

// q-12 (float32)
//
// Quality factor 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ12(value float32) error {
	return e.Element.SetProperty("q-12", value)
}

// q-13 (float32)
//
// Quality factor 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ13(value float32) error {
	return e.Element.SetProperty("q-13", value)
}

// q-14 (float32)
//
// Quality factor 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ14(value float32) error {
	return e.Element.SetProperty("q-14", value)
}

// q-15 (float32)
//
// Quality factor 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ15(value float32) error {
	return e.Element.SetProperty("q-15", value)
}

// q-16 (float32)
//
// Quality factor 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ16() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ16(value float32) error {
	return e.Element.SetProperty("q-16", value)
}

// q-17 (float32)
//
// Quality factor 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ17() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ17(value float32) error {
	return e.Element.SetProperty("q-17", value)
}

// q-18 (float32)
//
// Quality factor 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ18() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ18(value float32) error {
	return e.Element.SetProperty("q-18", value)
}

// q-19 (float32)
//
// Quality factor 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ19() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ19(value float32) error {
	return e.Element.SetProperty("q-19", value)
}

// q-2 (float32)
//
// Quality factor 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ2(value float32) error {
	return e.Element.SetProperty("q-2", value)
}

// q-20 (float32)
//
// Quality factor 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ20() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ20(value float32) error {
	return e.Element.SetProperty("q-20", value)
}

// q-21 (float32)
//
// Quality factor 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ21() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ21(value float32) error {
	return e.Element.SetProperty("q-21", value)
}

// q-22 (float32)
//
// Quality factor 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ22() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ22(value float32) error {
	return e.Element.SetProperty("q-22", value)
}

// q-23 (float32)
//
// Quality factor 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ23() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ23(value float32) error {
	return e.Element.SetProperty("q-23", value)
}

// q-24 (float32)
//
// Quality factor 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ24() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ24(value float32) error {
	return e.Element.SetProperty("q-24", value)
}

// q-25 (float32)
//
// Quality factor 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ25() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ25(value float32) error {
	return e.Element.SetProperty("q-25", value)
}

// q-26 (float32)
//
// Quality factor 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ26() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ26(value float32) error {
	return e.Element.SetProperty("q-26", value)
}

// q-27 (float32)
//
// Quality factor 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ27() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ27(value float32) error {
	return e.Element.SetProperty("q-27", value)
}

// q-28 (float32)
//
// Quality factor 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ28() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ28(value float32) error {
	return e.Element.SetProperty("q-28", value)
}

// q-29 (float32)
//
// Quality factor 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ29() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ29(value float32) error {
	return e.Element.SetProperty("q-29", value)
}

// q-3 (float32)
//
// Quality factor 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ3(value float32) error {
	return e.Element.SetProperty("q-3", value)
}

// q-30 (float32)
//
// Quality factor 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ30() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ30(value float32) error {
	return e.Element.SetProperty("q-30", value)
}

// q-31 (float32)
//
// Quality factor 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ31() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ31(value float32) error {
	return e.Element.SetProperty("q-31", value)
}

// q-4 (float32)
//
// Quality factor 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ4(value float32) error {
	return e.Element.SetProperty("q-4", value)
}

// q-5 (float32)
//
// Quality factor 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ5(value float32) error {
	return e.Element.SetProperty("q-5", value)
}

// q-6 (float32)
//
// Quality factor 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ6(value float32) error {
	return e.Element.SetProperty("q-6", value)
}

// q-7 (float32)
//
// Quality factor 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ7(value float32) error {
	return e.Element.SetProperty("q-7", value)
}

// q-8 (float32)
//
// Quality factor 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ8(value float32) error {
	return e.Element.SetProperty("q-8", value)
}

// q-9 (float32)
//
// Quality factor 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetQ9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetQ9(value float32) error {
	return e.Element.SetProperty("q-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// s-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos0)
//
// Filter slope 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS0() (interface{}, error) {
	return e.Element.GetProperty("s-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS0(value interface{}) error {
	return e.Element.SetProperty("s-0", value)
}

// s-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos1)
//
// Filter slope 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS1() (interface{}, error) {
	return e.Element.GetProperty("s-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS1(value interface{}) error {
	return e.Element.SetProperty("s-1", value)
}

// s-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos10)
//
// Filter slope 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS10() (interface{}, error) {
	return e.Element.GetProperty("s-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS10(value interface{}) error {
	return e.Element.SetProperty("s-10", value)
}

// s-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos11)
//
// Filter slope 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS11() (interface{}, error) {
	return e.Element.GetProperty("s-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS11(value interface{}) error {
	return e.Element.SetProperty("s-11", value)
}

// s-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos12)
//
// Filter slope 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS12() (interface{}, error) {
	return e.Element.GetProperty("s-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS12(value interface{}) error {
	return e.Element.SetProperty("s-12", value)
}

// s-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos13)
//
// Filter slope 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS13() (interface{}, error) {
	return e.Element.GetProperty("s-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS13(value interface{}) error {
	return e.Element.SetProperty("s-13", value)
}

// s-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos14)
//
// Filter slope 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS14() (interface{}, error) {
	return e.Element.GetProperty("s-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS14(value interface{}) error {
	return e.Element.SetProperty("s-14", value)
}

// s-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos15)
//
// Filter slope 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS15() (interface{}, error) {
	return e.Element.GetProperty("s-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS15(value interface{}) error {
	return e.Element.SetProperty("s-15", value)
}

// s-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos16)
//
// Filter slope 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS16() (interface{}, error) {
	return e.Element.GetProperty("s-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS16(value interface{}) error {
	return e.Element.SetProperty("s-16", value)
}

// s-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos17)
//
// Filter slope 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS17() (interface{}, error) {
	return e.Element.GetProperty("s-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS17(value interface{}) error {
	return e.Element.SetProperty("s-17", value)
}

// s-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos18)
//
// Filter slope 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS18() (interface{}, error) {
	return e.Element.GetProperty("s-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS18(value interface{}) error {
	return e.Element.SetProperty("s-18", value)
}

// s-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos19)
//
// Filter slope 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS19() (interface{}, error) {
	return e.Element.GetProperty("s-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS19(value interface{}) error {
	return e.Element.SetProperty("s-19", value)
}

// s-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos2)
//
// Filter slope 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS2() (interface{}, error) {
	return e.Element.GetProperty("s-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS2(value interface{}) error {
	return e.Element.SetProperty("s-2", value)
}

// s-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos20)
//
// Filter slope 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS20() (interface{}, error) {
	return e.Element.GetProperty("s-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS20(value interface{}) error {
	return e.Element.SetProperty("s-20", value)
}

// s-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos21)
//
// Filter slope 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS21() (interface{}, error) {
	return e.Element.GetProperty("s-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS21(value interface{}) error {
	return e.Element.SetProperty("s-21", value)
}

// s-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos22)
//
// Filter slope 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS22() (interface{}, error) {
	return e.Element.GetProperty("s-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS22(value interface{}) error {
	return e.Element.SetProperty("s-22", value)
}

// s-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos23)
//
// Filter slope 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS23() (interface{}, error) {
	return e.Element.GetProperty("s-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS23(value interface{}) error {
	return e.Element.SetProperty("s-23", value)
}

// s-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos24)
//
// Filter slope 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS24() (interface{}, error) {
	return e.Element.GetProperty("s-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS24(value interface{}) error {
	return e.Element.SetProperty("s-24", value)
}

// s-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos25)
//
// Filter slope 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS25() (interface{}, error) {
	return e.Element.GetProperty("s-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS25(value interface{}) error {
	return e.Element.SetProperty("s-25", value)
}

// s-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos26)
//
// Filter slope 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS26() (interface{}, error) {
	return e.Element.GetProperty("s-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS26(value interface{}) error {
	return e.Element.SetProperty("s-26", value)
}

// s-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos27)
//
// Filter slope 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS27() (interface{}, error) {
	return e.Element.GetProperty("s-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS27(value interface{}) error {
	return e.Element.SetProperty("s-27", value)
}

// s-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos28)
//
// Filter slope 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS28() (interface{}, error) {
	return e.Element.GetProperty("s-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS28(value interface{}) error {
	return e.Element.SetProperty("s-28", value)
}

// s-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos29)
//
// Filter slope 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS29() (interface{}, error) {
	return e.Element.GetProperty("s-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS29(value interface{}) error {
	return e.Element.SetProperty("s-29", value)
}

// s-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos3)
//
// Filter slope 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS3() (interface{}, error) {
	return e.Element.GetProperty("s-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS3(value interface{}) error {
	return e.Element.SetProperty("s-3", value)
}

// s-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos30)
//
// Filter slope 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS30() (interface{}, error) {
	return e.Element.GetProperty("s-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS30(value interface{}) error {
	return e.Element.SetProperty("s-30", value)
}

// s-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos31)
//
// Filter slope 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS31() (interface{}, error) {
	return e.Element.GetProperty("s-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS31(value interface{}) error {
	return e.Element.SetProperty("s-31", value)
}

// s-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos4)
//
// Filter slope 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS4() (interface{}, error) {
	return e.Element.GetProperty("s-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS4(value interface{}) error {
	return e.Element.SetProperty("s-4", value)
}

// s-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos5)
//
// Filter slope 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS5() (interface{}, error) {
	return e.Element.GetProperty("s-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS5(value interface{}) error {
	return e.Element.SetProperty("s-5", value)
}

// s-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos6)
//
// Filter slope 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS6() (interface{}, error) {
	return e.Element.GetProperty("s-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS6(value interface{}) error {
	return e.Element.SetProperty("s-6", value)
}

// s-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos7)
//
// Filter slope 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS7() (interface{}, error) {
	return e.Element.GetProperty("s-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS7(value interface{}) error {
	return e.Element.SetProperty("s-7", value)
}

// s-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos8)
//
// Filter slope 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS8() (interface{}, error) {
	return e.Element.GetProperty("s-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS8(value interface{}) error {
	return e.Element.SetProperty("s-8", value)
}

// s-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Stereos9)
//
// Filter slope 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetS9() (interface{}, error) {
	return e.Element.GetProperty("s-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetS9(value interface{}) error {
	return e.Element.SetProperty("s-9", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetSmr() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm0(value bool) error {
	return e.Element.SetProperty("xm-0", value)
}

// xm-1 (bool)
//
// Filter mute 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm1(value bool) error {
	return e.Element.SetProperty("xm-1", value)
}

// xm-10 (bool)
//
// Filter mute 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm10(value bool) error {
	return e.Element.SetProperty("xm-10", value)
}

// xm-11 (bool)
//
// Filter mute 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm11(value bool) error {
	return e.Element.SetProperty("xm-11", value)
}

// xm-12 (bool)
//
// Filter mute 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm12(value bool) error {
	return e.Element.SetProperty("xm-12", value)
}

// xm-13 (bool)
//
// Filter mute 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm13(value bool) error {
	return e.Element.SetProperty("xm-13", value)
}

// xm-14 (bool)
//
// Filter mute 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm14(value bool) error {
	return e.Element.SetProperty("xm-14", value)
}

// xm-15 (bool)
//
// Filter mute 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm15(value bool) error {
	return e.Element.SetProperty("xm-15", value)
}

// xm-16 (bool)
//
// Filter mute 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm16() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm16(value bool) error {
	return e.Element.SetProperty("xm-16", value)
}

// xm-17 (bool)
//
// Filter mute 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm17() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm17(value bool) error {
	return e.Element.SetProperty("xm-17", value)
}

// xm-18 (bool)
//
// Filter mute 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm18() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm18(value bool) error {
	return e.Element.SetProperty("xm-18", value)
}

// xm-19 (bool)
//
// Filter mute 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm19() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm19(value bool) error {
	return e.Element.SetProperty("xm-19", value)
}

// xm-2 (bool)
//
// Filter mute 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm2(value bool) error {
	return e.Element.SetProperty("xm-2", value)
}

// xm-20 (bool)
//
// Filter mute 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm20() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm20(value bool) error {
	return e.Element.SetProperty("xm-20", value)
}

// xm-21 (bool)
//
// Filter mute 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm21() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm21(value bool) error {
	return e.Element.SetProperty("xm-21", value)
}

// xm-22 (bool)
//
// Filter mute 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm22() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm22(value bool) error {
	return e.Element.SetProperty("xm-22", value)
}

// xm-23 (bool)
//
// Filter mute 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm23() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm23(value bool) error {
	return e.Element.SetProperty("xm-23", value)
}

// xm-24 (bool)
//
// Filter mute 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm24() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm24(value bool) error {
	return e.Element.SetProperty("xm-24", value)
}

// xm-25 (bool)
//
// Filter mute 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm25() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm25(value bool) error {
	return e.Element.SetProperty("xm-25", value)
}

// xm-26 (bool)
//
// Filter mute 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm26() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm26(value bool) error {
	return e.Element.SetProperty("xm-26", value)
}

// xm-27 (bool)
//
// Filter mute 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm27() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm27(value bool) error {
	return e.Element.SetProperty("xm-27", value)
}

// xm-28 (bool)
//
// Filter mute 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm28() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm28(value bool) error {
	return e.Element.SetProperty("xm-28", value)
}

// xm-29 (bool)
//
// Filter mute 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm29() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm29(value bool) error {
	return e.Element.SetProperty("xm-29", value)
}

// xm-3 (bool)
//
// Filter mute 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm3(value bool) error {
	return e.Element.SetProperty("xm-3", value)
}

// xm-30 (bool)
//
// Filter mute 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm30() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm30(value bool) error {
	return e.Element.SetProperty("xm-30", value)
}

// xm-31 (bool)
//
// Filter mute 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm31() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm31(value bool) error {
	return e.Element.SetProperty("xm-31", value)
}

// xm-4 (bool)
//
// Filter mute 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm4(value bool) error {
	return e.Element.SetProperty("xm-4", value)
}

// xm-5 (bool)
//
// Filter mute 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm5(value bool) error {
	return e.Element.SetProperty("xm-5", value)
}

// xm-6 (bool)
//
// Filter mute 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm6(value bool) error {
	return e.Element.SetProperty("xm-6", value)
}

// xm-7 (bool)
//
// Filter mute 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm7(value bool) error {
	return e.Element.SetProperty("xm-7", value)
}

// xm-8 (bool)
//
// Filter mute 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm8(value bool) error {
	return e.Element.SetProperty("xm-8", value)
}

// xm-9 (bool)
//
// Filter mute 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXm9(value bool) error {
	return e.Element.SetProperty("xm-9", value)
}

// xs-0 (bool)
//
// Filter solo 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs0(value bool) error {
	return e.Element.SetProperty("xs-0", value)
}

// xs-1 (bool)
//
// Filter solo 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs1(value bool) error {
	return e.Element.SetProperty("xs-1", value)
}

// xs-10 (bool)
//
// Filter solo 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs10(value bool) error {
	return e.Element.SetProperty("xs-10", value)
}

// xs-11 (bool)
//
// Filter solo 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs11(value bool) error {
	return e.Element.SetProperty("xs-11", value)
}

// xs-12 (bool)
//
// Filter solo 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs12(value bool) error {
	return e.Element.SetProperty("xs-12", value)
}

// xs-13 (bool)
//
// Filter solo 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs13(value bool) error {
	return e.Element.SetProperty("xs-13", value)
}

// xs-14 (bool)
//
// Filter solo 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs14(value bool) error {
	return e.Element.SetProperty("xs-14", value)
}

// xs-15 (bool)
//
// Filter solo 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs15(value bool) error {
	return e.Element.SetProperty("xs-15", value)
}

// xs-16 (bool)
//
// Filter solo 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs16() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs16(value bool) error {
	return e.Element.SetProperty("xs-16", value)
}

// xs-17 (bool)
//
// Filter solo 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs17() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs17(value bool) error {
	return e.Element.SetProperty("xs-17", value)
}

// xs-18 (bool)
//
// Filter solo 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs18() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs18(value bool) error {
	return e.Element.SetProperty("xs-18", value)
}

// xs-19 (bool)
//
// Filter solo 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs19() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs19(value bool) error {
	return e.Element.SetProperty("xs-19", value)
}

// xs-2 (bool)
//
// Filter solo 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs2(value bool) error {
	return e.Element.SetProperty("xs-2", value)
}

// xs-20 (bool)
//
// Filter solo 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs20() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs20(value bool) error {
	return e.Element.SetProperty("xs-20", value)
}

// xs-21 (bool)
//
// Filter solo 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs21() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs21(value bool) error {
	return e.Element.SetProperty("xs-21", value)
}

// xs-22 (bool)
//
// Filter solo 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs22() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs22(value bool) error {
	return e.Element.SetProperty("xs-22", value)
}

// xs-23 (bool)
//
// Filter solo 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs23() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs23(value bool) error {
	return e.Element.SetProperty("xs-23", value)
}

// xs-24 (bool)
//
// Filter solo 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs24() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs24(value bool) error {
	return e.Element.SetProperty("xs-24", value)
}

// xs-25 (bool)
//
// Filter solo 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs25() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs25(value bool) error {
	return e.Element.SetProperty("xs-25", value)
}

// xs-26 (bool)
//
// Filter solo 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs26() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs26(value bool) error {
	return e.Element.SetProperty("xs-26", value)
}

// xs-27 (bool)
//
// Filter solo 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs27() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs27(value bool) error {
	return e.Element.SetProperty("xs-27", value)
}

// xs-28 (bool)
//
// Filter solo 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs28() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs28(value bool) error {
	return e.Element.SetProperty("xs-28", value)
}

// xs-29 (bool)
//
// Filter solo 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs29() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs29(value bool) error {
	return e.Element.SetProperty("xs-29", value)
}

// xs-3 (bool)
//
// Filter solo 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs3(value bool) error {
	return e.Element.SetProperty("xs-3", value)
}

// xs-30 (bool)
//
// Filter solo 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs30() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs30(value bool) error {
	return e.Element.SetProperty("xs-30", value)
}

// xs-31 (bool)
//
// Filter solo 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs31() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs31(value bool) error {
	return e.Element.SetProperty("xs-31", value)
}

// xs-4 (bool)
//
// Filter solo 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs4(value bool) error {
	return e.Element.SetProperty("xs-4", value)
}

// xs-5 (bool)
//
// Filter solo 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs5(value bool) error {
	return e.Element.SetProperty("xs-5", value)
}

// xs-6 (bool)
//
// Filter solo 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs6(value bool) error {
	return e.Element.SetProperty("xs-6", value)
}

// xs-7 (bool)
//
// Filter solo 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs7(value bool) error {
	return e.Element.SetProperty("xs-7", value)
}

// xs-8 (bool)
//
// Filter solo 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs8(value bool) error {
	return e.Element.SetProperty("xs-8", value)
}

// xs-9 (bool)
//
// Filter solo 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetXs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetXs9(value bool) error {
	return e.Element.SetProperty("xs-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Stereo) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm27ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm27 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX32StereofselFilters07 LspPlugInPluginsLv2ParaEqualizerX32Stereofsel = "Filters 0-7" // Filters 0-7 (0) – Filters 0-7
	LspPlugInPluginsLv2ParaEqualizerX32StereofselFilters815 LspPlugInPluginsLv2ParaEqualizerX32Stereofsel = "Filters 8-15" // Filters 8-15 (1) – Filters 8-15
	LspPlugInPluginsLv2ParaEqualizerX32StereofselFilters1623 LspPlugInPluginsLv2ParaEqualizerX32Stereofsel = "Filters 16-23" // Filters 16-23 (2) – Filters 16-23
	LspPlugInPluginsLv2ParaEqualizerX32StereofselFilters2431 LspPlugInPluginsLv2ParaEqualizerX32Stereofsel = "Filters 24-31" // Filters 24-31 (3) – Filters 24-31
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft3Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft5Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos15X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos15X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos15X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos15X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos25X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos25 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos25X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos25 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos25X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos25 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos25X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos25 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos8X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos8X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos8X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos8X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm10ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft16Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft16 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft18Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft18 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft7Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos22X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos22 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos22X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos22 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos22X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos22 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos22X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos22 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm24ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm24 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereomode string

const (
	LspPlugInPluginsLv2ParaEqualizerX32StereomodeIir LspPlugInPluginsLv2ParaEqualizerX32Stereomode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX32StereomodeFir LspPlugInPluginsLv2ParaEqualizerX32Stereomode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX32StereomodeFft LspPlugInPluginsLv2ParaEqualizerX32Stereomode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos2X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos2X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos2X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos2X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos26X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos26 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos26X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos26 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos26X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos26 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos26X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos26 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos6X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos6X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos6X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos6X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm29ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm29 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm20ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm20 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm22ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm22 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft13Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft30Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft30 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos13X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos13X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos13X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos13X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos18X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos18 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos18X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos18 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos18X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos18 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos18X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos18 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm0ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos0X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos0X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos0X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos0X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos17X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos17 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos17X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos17 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos17X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos17 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos17X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos17 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm4ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm25ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm25 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft11Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft25Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft25 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft26Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft26 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos27X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos27 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos27X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos27 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos27X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos27 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos27X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos27 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos28X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos28 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos28X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos28 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos28X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos28 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos28X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos28 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos3X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos3X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos3X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos3X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm2ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos31X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos31 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos31X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos31 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos31X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos31 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos31X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos31 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos4X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos4X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos4X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos4X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos30X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos30 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos30X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos30 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos30X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos30 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos30X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos30 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm14ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm19ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm19 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft0Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft12Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft19Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft19 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos11X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos11X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos11X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos11X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos20X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos20 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos20X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos20 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos20X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos20 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos20X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos20 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm1ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos9X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos9X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos9X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos9X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos29X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos29 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos29X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos29 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos29X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos29 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos29X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos29 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos14X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos14X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos14X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos14X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm30ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm30 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft29Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft29 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft4Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos23X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos23 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos23X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos23 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos23X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos23 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos23X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos23 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft1Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft27Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft27 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft31Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft31 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft8Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos1X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos1X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos1X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos1X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos16X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos16 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos16X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos16 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos16X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos16 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos16X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos16 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos19X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos19 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos19X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos19 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos19X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos19 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos19X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos19 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm16ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm16 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft17Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft17 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft6Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft9Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos21X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos21 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos21X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos21 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos21X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos21 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos21X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos21 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos24X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos24 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos24X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos24 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos24X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos24 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos24X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos24 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm26ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm26 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm23ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm23 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm28ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm28 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft14Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft2Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft21Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft21 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos5X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos5X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos5X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos5X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm15ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm17ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm17 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm21ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm21 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm3ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm5ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm6ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft20Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft20 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft22Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft22 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm11ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos10X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos10X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos10X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos10X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos12X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos12X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos12X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos12X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereos7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereos7X1 LspPlugInPluginsLv2ParaEqualizerX32Stereos7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Stereos7X2 LspPlugInPluginsLv2ParaEqualizerX32Stereos7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Stereos7X3 LspPlugInPluginsLv2ParaEqualizerX32Stereos7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Stereos7X4 LspPlugInPluginsLv2ParaEqualizerX32Stereos7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft23Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft23 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm7ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft10Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft24Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft24 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft28Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft28 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofft string

const (
	LspPlugInPluginsLv2ParaEqualizerX32StereofftOff LspPlugInPluginsLv2ParaEqualizerX32Stereofft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32StereofftPostEq LspPlugInPluginsLv2ParaEqualizerX32Stereofft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX32StereofftPreEq LspPlugInPluginsLv2ParaEqualizerX32Stereofft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm18ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm18 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm31ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm31 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm13ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm8ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm9ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15Off LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15Bell LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15HiPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15HiShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15LoPass LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15LoShelf LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15Notch LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15Resonance LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Stereoft15Allpass LspPlugInPluginsLv2ParaEqualizerX32Stereoft15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12RlcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12RlcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12BwcBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12BwcMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12LrxBt LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12LrxMt LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Stereofm12ApoDr LspPlugInPluginsLv2ParaEqualizerX32Stereofm12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

