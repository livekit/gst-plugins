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

type LspPlugInPluginsLv2ParaEqualizerX16Lr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX16Lr() (*LspPlugInPluginsLv2ParaEqualizerX16Lr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x16-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX16LrWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX16Lr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x16-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX16Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFftvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFftvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fl-0 (float32)
//
// Frequency Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl0(value float32) error {
	return e.Element.SetProperty("fl-0", value)
}

// fl-1 (float32)
//
// Frequency Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl1(value float32) error {
	return e.Element.SetProperty("fl-1", value)
}

// fl-10 (float32)
//
// Frequency Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl10(value float32) error {
	return e.Element.SetProperty("fl-10", value)
}

// fl-11 (float32)
//
// Frequency Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl11(value float32) error {
	return e.Element.SetProperty("fl-11", value)
}

// fl-12 (float32)
//
// Frequency Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl12(value float32) error {
	return e.Element.SetProperty("fl-12", value)
}

// fl-13 (float32)
//
// Frequency Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl13(value float32) error {
	return e.Element.SetProperty("fl-13", value)
}

// fl-14 (float32)
//
// Frequency Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl14(value float32) error {
	return e.Element.SetProperty("fl-14", value)
}

// fl-15 (float32)
//
// Frequency Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl15(value float32) error {
	return e.Element.SetProperty("fl-15", value)
}

// fl-2 (float32)
//
// Frequency Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl2(value float32) error {
	return e.Element.SetProperty("fl-2", value)
}

// fl-3 (float32)
//
// Frequency Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl3(value float32) error {
	return e.Element.SetProperty("fl-3", value)
}

// fl-4 (float32)
//
// Frequency Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl4(value float32) error {
	return e.Element.SetProperty("fl-4", value)
}

// fl-5 (float32)
//
// Frequency Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl5(value float32) error {
	return e.Element.SetProperty("fl-5", value)
}

// fl-6 (float32)
//
// Frequency Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl6(value float32) error {
	return e.Element.SetProperty("fl-6", value)
}

// fl-7 (float32)
//
// Frequency Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl7(value float32) error {
	return e.Element.SetProperty("fl-7", value)
}

// fl-8 (float32)
//
// Frequency Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl8(value float32) error {
	return e.Element.SetProperty("fl-8", value)
}

// fl-9 (float32)
//
// Frequency Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFl9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFl9(value float32) error {
	return e.Element.SetProperty("fl-9", value)
}

// fml-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml0)
//
// Filter mode Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml0() (interface{}, error) {
	return e.Element.GetProperty("fml-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml0(value interface{}) error {
	return e.Element.SetProperty("fml-0", value)
}

// fml-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml1)
//
// Filter mode Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml1() (interface{}, error) {
	return e.Element.GetProperty("fml-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml1(value interface{}) error {
	return e.Element.SetProperty("fml-1", value)
}

// fml-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml10)
//
// Filter mode Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml10() (interface{}, error) {
	return e.Element.GetProperty("fml-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml10(value interface{}) error {
	return e.Element.SetProperty("fml-10", value)
}

// fml-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml11)
//
// Filter mode Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml11() (interface{}, error) {
	return e.Element.GetProperty("fml-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml11(value interface{}) error {
	return e.Element.SetProperty("fml-11", value)
}

// fml-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml12)
//
// Filter mode Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml12() (interface{}, error) {
	return e.Element.GetProperty("fml-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml12(value interface{}) error {
	return e.Element.SetProperty("fml-12", value)
}

// fml-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml13)
//
// Filter mode Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml13() (interface{}, error) {
	return e.Element.GetProperty("fml-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml13(value interface{}) error {
	return e.Element.SetProperty("fml-13", value)
}

// fml-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml14)
//
// Filter mode Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml14() (interface{}, error) {
	return e.Element.GetProperty("fml-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml14(value interface{}) error {
	return e.Element.SetProperty("fml-14", value)
}

// fml-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml15)
//
// Filter mode Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml15() (interface{}, error) {
	return e.Element.GetProperty("fml-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml15(value interface{}) error {
	return e.Element.SetProperty("fml-15", value)
}

// fml-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml2)
//
// Filter mode Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml2() (interface{}, error) {
	return e.Element.GetProperty("fml-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml2(value interface{}) error {
	return e.Element.SetProperty("fml-2", value)
}

// fml-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml3)
//
// Filter mode Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml3() (interface{}, error) {
	return e.Element.GetProperty("fml-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml3(value interface{}) error {
	return e.Element.SetProperty("fml-3", value)
}

// fml-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml4)
//
// Filter mode Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml4() (interface{}, error) {
	return e.Element.GetProperty("fml-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml4(value interface{}) error {
	return e.Element.SetProperty("fml-4", value)
}

// fml-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml5)
//
// Filter mode Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml5() (interface{}, error) {
	return e.Element.GetProperty("fml-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml5(value interface{}) error {
	return e.Element.SetProperty("fml-5", value)
}

// fml-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml6)
//
// Filter mode Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml6() (interface{}, error) {
	return e.Element.GetProperty("fml-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml6(value interface{}) error {
	return e.Element.SetProperty("fml-6", value)
}

// fml-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml7)
//
// Filter mode Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml7() (interface{}, error) {
	return e.Element.GetProperty("fml-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml7(value interface{}) error {
	return e.Element.SetProperty("fml-7", value)
}

// fml-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml8)
//
// Filter mode Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml8() (interface{}, error) {
	return e.Element.GetProperty("fml-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml8(value interface{}) error {
	return e.Element.SetProperty("fml-8", value)
}

// fml-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfml9)
//
// Filter mode Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFml9() (interface{}, error) {
	return e.Element.GetProperty("fml-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFml9(value interface{}) error {
	return e.Element.SetProperty("fml-9", value)
}

// fmr-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr0)
//
// Filter mode Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr0() (interface{}, error) {
	return e.Element.GetProperty("fmr-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr0(value interface{}) error {
	return e.Element.SetProperty("fmr-0", value)
}

// fmr-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr1)
//
// Filter mode Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr1() (interface{}, error) {
	return e.Element.GetProperty("fmr-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr1(value interface{}) error {
	return e.Element.SetProperty("fmr-1", value)
}

// fmr-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr10)
//
// Filter mode Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr10() (interface{}, error) {
	return e.Element.GetProperty("fmr-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr10(value interface{}) error {
	return e.Element.SetProperty("fmr-10", value)
}

// fmr-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr11)
//
// Filter mode Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr11() (interface{}, error) {
	return e.Element.GetProperty("fmr-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr11(value interface{}) error {
	return e.Element.SetProperty("fmr-11", value)
}

// fmr-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr12)
//
// Filter mode Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr12() (interface{}, error) {
	return e.Element.GetProperty("fmr-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr12(value interface{}) error {
	return e.Element.SetProperty("fmr-12", value)
}

// fmr-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr13)
//
// Filter mode Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr13() (interface{}, error) {
	return e.Element.GetProperty("fmr-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr13(value interface{}) error {
	return e.Element.SetProperty("fmr-13", value)
}

// fmr-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr14)
//
// Filter mode Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr14() (interface{}, error) {
	return e.Element.GetProperty("fmr-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr14(value interface{}) error {
	return e.Element.SetProperty("fmr-14", value)
}

// fmr-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr15)
//
// Filter mode Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr15() (interface{}, error) {
	return e.Element.GetProperty("fmr-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr15(value interface{}) error {
	return e.Element.SetProperty("fmr-15", value)
}

// fmr-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr2)
//
// Filter mode Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr2() (interface{}, error) {
	return e.Element.GetProperty("fmr-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr2(value interface{}) error {
	return e.Element.SetProperty("fmr-2", value)
}

// fmr-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr3)
//
// Filter mode Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr3() (interface{}, error) {
	return e.Element.GetProperty("fmr-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr3(value interface{}) error {
	return e.Element.SetProperty("fmr-3", value)
}

// fmr-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr4)
//
// Filter mode Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr4() (interface{}, error) {
	return e.Element.GetProperty("fmr-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr4(value interface{}) error {
	return e.Element.SetProperty("fmr-4", value)
}

// fmr-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr5)
//
// Filter mode Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr5() (interface{}, error) {
	return e.Element.GetProperty("fmr-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr5(value interface{}) error {
	return e.Element.SetProperty("fmr-5", value)
}

// fmr-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr6)
//
// Filter mode Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr6() (interface{}, error) {
	return e.Element.GetProperty("fmr-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr6(value interface{}) error {
	return e.Element.SetProperty("fmr-6", value)
}

// fmr-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr7)
//
// Filter mode Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr7() (interface{}, error) {
	return e.Element.GetProperty("fmr-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr7(value interface{}) error {
	return e.Element.SetProperty("fmr-7", value)
}

// fmr-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr8)
//
// Filter mode Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr8() (interface{}, error) {
	return e.Element.GetProperty("fmr-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr8(value interface{}) error {
	return e.Element.SetProperty("fmr-8", value)
}

// fmr-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfmr9)
//
// Filter mode Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFmr9() (interface{}, error) {
	return e.Element.GetProperty("fmr-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFmr9(value interface{}) error {
	return e.Element.SetProperty("fmr-9", value)
}

// fr-0 (float32)
//
// Frequency Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr0(value float32) error {
	return e.Element.SetProperty("fr-0", value)
}

// fr-1 (float32)
//
// Frequency Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr1(value float32) error {
	return e.Element.SetProperty("fr-1", value)
}

// fr-10 (float32)
//
// Frequency Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr10(value float32) error {
	return e.Element.SetProperty("fr-10", value)
}

// fr-11 (float32)
//
// Frequency Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr11(value float32) error {
	return e.Element.SetProperty("fr-11", value)
}

// fr-12 (float32)
//
// Frequency Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr12(value float32) error {
	return e.Element.SetProperty("fr-12", value)
}

// fr-13 (float32)
//
// Frequency Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr13(value float32) error {
	return e.Element.SetProperty("fr-13", value)
}

// fr-14 (float32)
//
// Frequency Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr14(value float32) error {
	return e.Element.SetProperty("fr-14", value)
}

// fr-15 (float32)
//
// Frequency Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr15(value float32) error {
	return e.Element.SetProperty("fr-15", value)
}

// fr-2 (float32)
//
// Frequency Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr2(value float32) error {
	return e.Element.SetProperty("fr-2", value)
}

// fr-3 (float32)
//
// Frequency Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr3(value float32) error {
	return e.Element.SetProperty("fr-3", value)
}

// fr-4 (float32)
//
// Frequency Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr4(value float32) error {
	return e.Element.SetProperty("fr-4", value)
}

// fr-5 (float32)
//
// Frequency Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr5(value float32) error {
	return e.Element.SetProperty("fr-5", value)
}

// fr-6 (float32)
//
// Frequency Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr6(value float32) error {
	return e.Element.SetProperty("fr-6", value)
}

// fr-7 (float32)
//
// Frequency Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr7(value float32) error {
	return e.Element.SetProperty("fr-7", value)
}

// fr-8 (float32)
//
// Frequency Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr8(value float32) error {
	return e.Element.SetProperty("fr-8", value)
}

// fr-9 (float32)
//
// Frequency Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFr9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFr9(value float32) error {
	return e.Element.SetProperty("fr-9", value)
}

// frqs-l (float32)
//
// Frequency shift Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFrqsL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frqs-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFrqsL(value float32) error {
	return e.Element.SetProperty("frqs-l", value)
}

// frqs-r (float32)
//
// Frequency shift Right

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFrqsR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("frqs-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFrqsR(value float32) error {
	return e.Element.SetProperty("frqs-r", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX16Lrfsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ftl-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl0)
//
// Filter type Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl0() (interface{}, error) {
	return e.Element.GetProperty("ftl-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl0(value interface{}) error {
	return e.Element.SetProperty("ftl-0", value)
}

// ftl-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl1)
//
// Filter type Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl1() (interface{}, error) {
	return e.Element.GetProperty("ftl-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl1(value interface{}) error {
	return e.Element.SetProperty("ftl-1", value)
}

// ftl-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl10)
//
// Filter type Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl10() (interface{}, error) {
	return e.Element.GetProperty("ftl-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl10(value interface{}) error {
	return e.Element.SetProperty("ftl-10", value)
}

// ftl-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl11)
//
// Filter type Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl11() (interface{}, error) {
	return e.Element.GetProperty("ftl-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl11(value interface{}) error {
	return e.Element.SetProperty("ftl-11", value)
}

// ftl-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl12)
//
// Filter type Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl12() (interface{}, error) {
	return e.Element.GetProperty("ftl-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl12(value interface{}) error {
	return e.Element.SetProperty("ftl-12", value)
}

// ftl-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl13)
//
// Filter type Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl13() (interface{}, error) {
	return e.Element.GetProperty("ftl-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl13(value interface{}) error {
	return e.Element.SetProperty("ftl-13", value)
}

// ftl-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl14)
//
// Filter type Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl14() (interface{}, error) {
	return e.Element.GetProperty("ftl-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl14(value interface{}) error {
	return e.Element.SetProperty("ftl-14", value)
}

// ftl-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl15)
//
// Filter type Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl15() (interface{}, error) {
	return e.Element.GetProperty("ftl-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl15(value interface{}) error {
	return e.Element.SetProperty("ftl-15", value)
}

// ftl-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl2)
//
// Filter type Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl2() (interface{}, error) {
	return e.Element.GetProperty("ftl-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl2(value interface{}) error {
	return e.Element.SetProperty("ftl-2", value)
}

// ftl-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl3)
//
// Filter type Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl3() (interface{}, error) {
	return e.Element.GetProperty("ftl-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl3(value interface{}) error {
	return e.Element.SetProperty("ftl-3", value)
}

// ftl-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl4)
//
// Filter type Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl4() (interface{}, error) {
	return e.Element.GetProperty("ftl-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl4(value interface{}) error {
	return e.Element.SetProperty("ftl-4", value)
}

// ftl-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl5)
//
// Filter type Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl5() (interface{}, error) {
	return e.Element.GetProperty("ftl-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl5(value interface{}) error {
	return e.Element.SetProperty("ftl-5", value)
}

// ftl-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl6)
//
// Filter type Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl6() (interface{}, error) {
	return e.Element.GetProperty("ftl-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl6(value interface{}) error {
	return e.Element.SetProperty("ftl-6", value)
}

// ftl-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl7)
//
// Filter type Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl7() (interface{}, error) {
	return e.Element.GetProperty("ftl-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl7(value interface{}) error {
	return e.Element.SetProperty("ftl-7", value)
}

// ftl-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl8)
//
// Filter type Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl8() (interface{}, error) {
	return e.Element.GetProperty("ftl-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl8(value interface{}) error {
	return e.Element.SetProperty("ftl-8", value)
}

// ftl-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftl9)
//
// Filter type Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtl9() (interface{}, error) {
	return e.Element.GetProperty("ftl-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtl9(value interface{}) error {
	return e.Element.SetProperty("ftl-9", value)
}

// ftr-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr0)
//
// Filter type Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr0() (interface{}, error) {
	return e.Element.GetProperty("ftr-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr0(value interface{}) error {
	return e.Element.SetProperty("ftr-0", value)
}

// ftr-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr1)
//
// Filter type Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr1() (interface{}, error) {
	return e.Element.GetProperty("ftr-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr1(value interface{}) error {
	return e.Element.SetProperty("ftr-1", value)
}

// ftr-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr10)
//
// Filter type Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr10() (interface{}, error) {
	return e.Element.GetProperty("ftr-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr10(value interface{}) error {
	return e.Element.SetProperty("ftr-10", value)
}

// ftr-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr11)
//
// Filter type Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr11() (interface{}, error) {
	return e.Element.GetProperty("ftr-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr11(value interface{}) error {
	return e.Element.SetProperty("ftr-11", value)
}

// ftr-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr12)
//
// Filter type Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr12() (interface{}, error) {
	return e.Element.GetProperty("ftr-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr12(value interface{}) error {
	return e.Element.SetProperty("ftr-12", value)
}

// ftr-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr13)
//
// Filter type Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr13() (interface{}, error) {
	return e.Element.GetProperty("ftr-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr13(value interface{}) error {
	return e.Element.SetProperty("ftr-13", value)
}

// ftr-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr14)
//
// Filter type Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr14() (interface{}, error) {
	return e.Element.GetProperty("ftr-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr14(value interface{}) error {
	return e.Element.SetProperty("ftr-14", value)
}

// ftr-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr15)
//
// Filter type Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr15() (interface{}, error) {
	return e.Element.GetProperty("ftr-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr15(value interface{}) error {
	return e.Element.SetProperty("ftr-15", value)
}

// ftr-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr2)
//
// Filter type Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr2() (interface{}, error) {
	return e.Element.GetProperty("ftr-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr2(value interface{}) error {
	return e.Element.SetProperty("ftr-2", value)
}

// ftr-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr3)
//
// Filter type Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr3() (interface{}, error) {
	return e.Element.GetProperty("ftr-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr3(value interface{}) error {
	return e.Element.SetProperty("ftr-3", value)
}

// ftr-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr4)
//
// Filter type Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr4() (interface{}, error) {
	return e.Element.GetProperty("ftr-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr4(value interface{}) error {
	return e.Element.SetProperty("ftr-4", value)
}

// ftr-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr5)
//
// Filter type Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr5() (interface{}, error) {
	return e.Element.GetProperty("ftr-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr5(value interface{}) error {
	return e.Element.SetProperty("ftr-5", value)
}

// ftr-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr6)
//
// Filter type Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr6() (interface{}, error) {
	return e.Element.GetProperty("ftr-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr6(value interface{}) error {
	return e.Element.SetProperty("ftr-6", value)
}

// ftr-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr7)
//
// Filter type Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr7() (interface{}, error) {
	return e.Element.GetProperty("ftr-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr7(value interface{}) error {
	return e.Element.SetProperty("ftr-7", value)
}

// ftr-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr8)
//
// Filter type Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr8() (interface{}, error) {
	return e.Element.GetProperty("ftr-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr8(value interface{}) error {
	return e.Element.SetProperty("ftr-8", value)
}

// ftr-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrftr9)
//
// Filter type Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFtr9() (interface{}, error) {
	return e.Element.GetProperty("ftr-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetFtr9(value interface{}) error {
	return e.Element.SetProperty("ftr-9", value)
}

// fvl-0 (bool)
//
// Filter visibility Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-1 (bool)
//
// Filter visibility Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-10 (bool)
//
// Filter visibility Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-11 (bool)
//
// Filter visibility Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-12 (bool)
//
// Filter visibility Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-13 (bool)
//
// Filter visibility Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-14 (bool)
//
// Filter visibility Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-15 (bool)
//
// Filter visibility Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-2 (bool)
//
// Filter visibility Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-3 (bool)
//
// Filter visibility Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-4 (bool)
//
// Filter visibility Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-5 (bool)
//
// Filter visibility Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-6 (bool)
//
// Filter visibility Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-7 (bool)
//
// Filter visibility Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-8 (bool)
//
// Filter visibility Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-9 (bool)
//
// Filter visibility Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvl9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-0 (bool)
//
// Filter visibility Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-1 (bool)
//
// Filter visibility Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-10 (bool)
//
// Filter visibility Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-11 (bool)
//
// Filter visibility Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-12 (bool)
//
// Filter visibility Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-13 (bool)
//
// Filter visibility Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-14 (bool)
//
// Filter visibility Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-15 (bool)
//
// Filter visibility Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-2 (bool)
//
// Filter visibility Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-3 (bool)
//
// Filter visibility Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-4 (bool)
//
// Filter visibility Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-5 (bool)
//
// Filter visibility Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-6 (bool)
//
// Filter visibility Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-7 (bool)
//
// Filter visibility Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-8 (bool)
//
// Filter visibility Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-9 (bool)
//
// Filter visibility Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetFvr9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-9")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl-0 (float32)
//
// Gain Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl0(value float32) error {
	return e.Element.SetProperty("gl-0", value)
}

// gl-1 (float32)
//
// Gain Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl1(value float32) error {
	return e.Element.SetProperty("gl-1", value)
}

// gl-10 (float32)
//
// Gain Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl10(value float32) error {
	return e.Element.SetProperty("gl-10", value)
}

// gl-11 (float32)
//
// Gain Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl11(value float32) error {
	return e.Element.SetProperty("gl-11", value)
}

// gl-12 (float32)
//
// Gain Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl12(value float32) error {
	return e.Element.SetProperty("gl-12", value)
}

// gl-13 (float32)
//
// Gain Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl13(value float32) error {
	return e.Element.SetProperty("gl-13", value)
}

// gl-14 (float32)
//
// Gain Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl14(value float32) error {
	return e.Element.SetProperty("gl-14", value)
}

// gl-15 (float32)
//
// Gain Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl15(value float32) error {
	return e.Element.SetProperty("gl-15", value)
}

// gl-2 (float32)
//
// Gain Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl2(value float32) error {
	return e.Element.SetProperty("gl-2", value)
}

// gl-3 (float32)
//
// Gain Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl3(value float32) error {
	return e.Element.SetProperty("gl-3", value)
}

// gl-4 (float32)
//
// Gain Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl4(value float32) error {
	return e.Element.SetProperty("gl-4", value)
}

// gl-5 (float32)
//
// Gain Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl5(value float32) error {
	return e.Element.SetProperty("gl-5", value)
}

// gl-6 (float32)
//
// Gain Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl6(value float32) error {
	return e.Element.SetProperty("gl-6", value)
}

// gl-7 (float32)
//
// Gain Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl7(value float32) error {
	return e.Element.SetProperty("gl-7", value)
}

// gl-8 (float32)
//
// Gain Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl8(value float32) error {
	return e.Element.SetProperty("gl-8", value)
}

// gl-9 (float32)
//
// Gain Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGl9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGl9(value float32) error {
	return e.Element.SetProperty("gl-9", value)
}

// gr-0 (float32)
//
// Gain Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr0(value float32) error {
	return e.Element.SetProperty("gr-0", value)
}

// gr-1 (float32)
//
// Gain Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr1(value float32) error {
	return e.Element.SetProperty("gr-1", value)
}

// gr-10 (float32)
//
// Gain Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr10(value float32) error {
	return e.Element.SetProperty("gr-10", value)
}

// gr-11 (float32)
//
// Gain Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr11(value float32) error {
	return e.Element.SetProperty("gr-11", value)
}

// gr-12 (float32)
//
// Gain Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr12(value float32) error {
	return e.Element.SetProperty("gr-12", value)
}

// gr-13 (float32)
//
// Gain Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr13(value float32) error {
	return e.Element.SetProperty("gr-13", value)
}

// gr-14 (float32)
//
// Gain Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr14(value float32) error {
	return e.Element.SetProperty("gr-14", value)
}

// gr-15 (float32)
//
// Gain Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr15(value float32) error {
	return e.Element.SetProperty("gr-15", value)
}

// gr-2 (float32)
//
// Gain Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr2(value float32) error {
	return e.Element.SetProperty("gr-2", value)
}

// gr-3 (float32)
//
// Gain Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr3(value float32) error {
	return e.Element.SetProperty("gr-3", value)
}

// gr-4 (float32)
//
// Gain Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr4(value float32) error {
	return e.Element.SetProperty("gr-4", value)
}

// gr-5 (float32)
//
// Gain Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr5(value float32) error {
	return e.Element.SetProperty("gr-5", value)
}

// gr-6 (float32)
//
// Gain Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr6(value float32) error {
	return e.Element.SetProperty("gr-6", value)
}

// gr-7 (float32)
//
// Gain Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr7(value float32) error {
	return e.Element.SetProperty("gr-7", value)
}

// gr-8 (float32)
//
// Gain Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr8(value float32) error {
	return e.Element.SetProperty("gr-8", value)
}

// gr-9 (float32)
//
// Gain Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetGr9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetGr9(value float32) error {
	return e.Element.SetProperty("gr-9", value)
}

// huel-0 (float32)
//
// Hue Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel0(value float32) error {
	return e.Element.SetProperty("huel-0", value)
}

// huel-1 (float32)
//
// Hue Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel1(value float32) error {
	return e.Element.SetProperty("huel-1", value)
}

// huel-10 (float32)
//
// Hue Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel10(value float32) error {
	return e.Element.SetProperty("huel-10", value)
}

// huel-11 (float32)
//
// Hue Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel11(value float32) error {
	return e.Element.SetProperty("huel-11", value)
}

// huel-12 (float32)
//
// Hue Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel12(value float32) error {
	return e.Element.SetProperty("huel-12", value)
}

// huel-13 (float32)
//
// Hue Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel13(value float32) error {
	return e.Element.SetProperty("huel-13", value)
}

// huel-14 (float32)
//
// Hue Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel14(value float32) error {
	return e.Element.SetProperty("huel-14", value)
}

// huel-15 (float32)
//
// Hue Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel15(value float32) error {
	return e.Element.SetProperty("huel-15", value)
}

// huel-2 (float32)
//
// Hue Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel2(value float32) error {
	return e.Element.SetProperty("huel-2", value)
}

// huel-3 (float32)
//
// Hue Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel3(value float32) error {
	return e.Element.SetProperty("huel-3", value)
}

// huel-4 (float32)
//
// Hue Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel4(value float32) error {
	return e.Element.SetProperty("huel-4", value)
}

// huel-5 (float32)
//
// Hue Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel5(value float32) error {
	return e.Element.SetProperty("huel-5", value)
}

// huel-6 (float32)
//
// Hue Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel6(value float32) error {
	return e.Element.SetProperty("huel-6", value)
}

// huel-7 (float32)
//
// Hue Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel7(value float32) error {
	return e.Element.SetProperty("huel-7", value)
}

// huel-8 (float32)
//
// Hue Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel8(value float32) error {
	return e.Element.SetProperty("huel-8", value)
}

// huel-9 (float32)
//
// Hue Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuel9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuel9(value float32) error {
	return e.Element.SetProperty("huel-9", value)
}

// huer-0 (float32)
//
// Hue Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer0(value float32) error {
	return e.Element.SetProperty("huer-0", value)
}

// huer-1 (float32)
//
// Hue Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer1(value float32) error {
	return e.Element.SetProperty("huer-1", value)
}

// huer-10 (float32)
//
// Hue Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer10(value float32) error {
	return e.Element.SetProperty("huer-10", value)
}

// huer-11 (float32)
//
// Hue Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer11(value float32) error {
	return e.Element.SetProperty("huer-11", value)
}

// huer-12 (float32)
//
// Hue Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer12(value float32) error {
	return e.Element.SetProperty("huer-12", value)
}

// huer-13 (float32)
//
// Hue Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer13(value float32) error {
	return e.Element.SetProperty("huer-13", value)
}

// huer-14 (float32)
//
// Hue Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer14(value float32) error {
	return e.Element.SetProperty("huer-14", value)
}

// huer-15 (float32)
//
// Hue Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer15(value float32) error {
	return e.Element.SetProperty("huer-15", value)
}

// huer-2 (float32)
//
// Hue Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer2(value float32) error {
	return e.Element.SetProperty("huer-2", value)
}

// huer-3 (float32)
//
// Hue Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer3(value float32) error {
	return e.Element.SetProperty("huer-3", value)
}

// huer-4 (float32)
//
// Hue Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer4(value float32) error {
	return e.Element.SetProperty("huer-4", value)
}

// huer-5 (float32)
//
// Hue Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer5(value float32) error {
	return e.Element.SetProperty("huer-5", value)
}

// huer-6 (float32)
//
// Hue Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer6(value float32) error {
	return e.Element.SetProperty("huer-6", value)
}

// huer-7 (float32)
//
// Hue Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer7(value float32) error {
	return e.Element.SetProperty("huer-7", value)
}

// huer-8 (float32)
//
// Hue Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer8(value float32) error {
	return e.Element.SetProperty("huer-8", value)
}

// huer-9 (float32)
//
// Hue Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetHuer9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetHuer9(value float32) error {
	return e.Element.SetProperty("huer-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetImr() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2ParaEqualizerX16Lrmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetOutLatency() (int, error) {
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

// ql-0 (float32)
//
// Quality factor Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl0(value float32) error {
	return e.Element.SetProperty("ql-0", value)
}

// ql-1 (float32)
//
// Quality factor Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl1(value float32) error {
	return e.Element.SetProperty("ql-1", value)
}

// ql-10 (float32)
//
// Quality factor Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl10(value float32) error {
	return e.Element.SetProperty("ql-10", value)
}

// ql-11 (float32)
//
// Quality factor Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl11(value float32) error {
	return e.Element.SetProperty("ql-11", value)
}

// ql-12 (float32)
//
// Quality factor Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl12(value float32) error {
	return e.Element.SetProperty("ql-12", value)
}

// ql-13 (float32)
//
// Quality factor Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl13(value float32) error {
	return e.Element.SetProperty("ql-13", value)
}

// ql-14 (float32)
//
// Quality factor Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl14(value float32) error {
	return e.Element.SetProperty("ql-14", value)
}

// ql-15 (float32)
//
// Quality factor Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl15(value float32) error {
	return e.Element.SetProperty("ql-15", value)
}

// ql-2 (float32)
//
// Quality factor Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl2(value float32) error {
	return e.Element.SetProperty("ql-2", value)
}

// ql-3 (float32)
//
// Quality factor Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl3(value float32) error {
	return e.Element.SetProperty("ql-3", value)
}

// ql-4 (float32)
//
// Quality factor Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl4(value float32) error {
	return e.Element.SetProperty("ql-4", value)
}

// ql-5 (float32)
//
// Quality factor Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl5(value float32) error {
	return e.Element.SetProperty("ql-5", value)
}

// ql-6 (float32)
//
// Quality factor Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl6(value float32) error {
	return e.Element.SetProperty("ql-6", value)
}

// ql-7 (float32)
//
// Quality factor Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl7(value float32) error {
	return e.Element.SetProperty("ql-7", value)
}

// ql-8 (float32)
//
// Quality factor Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl8(value float32) error {
	return e.Element.SetProperty("ql-8", value)
}

// ql-9 (float32)
//
// Quality factor Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQl9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQl9(value float32) error {
	return e.Element.SetProperty("ql-9", value)
}

// qr-0 (float32)
//
// Quality factor Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr0(value float32) error {
	return e.Element.SetProperty("qr-0", value)
}

// qr-1 (float32)
//
// Quality factor Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr1(value float32) error {
	return e.Element.SetProperty("qr-1", value)
}

// qr-10 (float32)
//
// Quality factor Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr10(value float32) error {
	return e.Element.SetProperty("qr-10", value)
}

// qr-11 (float32)
//
// Quality factor Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr11(value float32) error {
	return e.Element.SetProperty("qr-11", value)
}

// qr-12 (float32)
//
// Quality factor Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr12(value float32) error {
	return e.Element.SetProperty("qr-12", value)
}

// qr-13 (float32)
//
// Quality factor Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr13(value float32) error {
	return e.Element.SetProperty("qr-13", value)
}

// qr-14 (float32)
//
// Quality factor Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr14(value float32) error {
	return e.Element.SetProperty("qr-14", value)
}

// qr-15 (float32)
//
// Quality factor Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr15(value float32) error {
	return e.Element.SetProperty("qr-15", value)
}

// qr-2 (float32)
//
// Quality factor Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr2(value float32) error {
	return e.Element.SetProperty("qr-2", value)
}

// qr-3 (float32)
//
// Quality factor Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr3(value float32) error {
	return e.Element.SetProperty("qr-3", value)
}

// qr-4 (float32)
//
// Quality factor Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr4(value float32) error {
	return e.Element.SetProperty("qr-4", value)
}

// qr-5 (float32)
//
// Quality factor Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr5(value float32) error {
	return e.Element.SetProperty("qr-5", value)
}

// qr-6 (float32)
//
// Quality factor Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr6(value float32) error {
	return e.Element.SetProperty("qr-6", value)
}

// qr-7 (float32)
//
// Quality factor Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr7(value float32) error {
	return e.Element.SetProperty("qr-7", value)
}

// qr-8 (float32)
//
// Quality factor Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr8(value float32) error {
	return e.Element.SetProperty("qr-8", value)
}

// qr-9 (float32)
//
// Quality factor Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetQr9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetQr9(value float32) error {
	return e.Element.SetProperty("qr-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sl-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl0)
//
// Filter slope Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl0() (interface{}, error) {
	return e.Element.GetProperty("sl-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl0(value interface{}) error {
	return e.Element.SetProperty("sl-0", value)
}

// sl-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl1)
//
// Filter slope Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl1() (interface{}, error) {
	return e.Element.GetProperty("sl-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl1(value interface{}) error {
	return e.Element.SetProperty("sl-1", value)
}

// sl-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl10)
//
// Filter slope Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl10() (interface{}, error) {
	return e.Element.GetProperty("sl-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl10(value interface{}) error {
	return e.Element.SetProperty("sl-10", value)
}

// sl-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl11)
//
// Filter slope Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl11() (interface{}, error) {
	return e.Element.GetProperty("sl-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl11(value interface{}) error {
	return e.Element.SetProperty("sl-11", value)
}

// sl-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl12)
//
// Filter slope Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl12() (interface{}, error) {
	return e.Element.GetProperty("sl-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl12(value interface{}) error {
	return e.Element.SetProperty("sl-12", value)
}

// sl-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl13)
//
// Filter slope Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl13() (interface{}, error) {
	return e.Element.GetProperty("sl-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl13(value interface{}) error {
	return e.Element.SetProperty("sl-13", value)
}

// sl-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl14)
//
// Filter slope Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl14() (interface{}, error) {
	return e.Element.GetProperty("sl-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl14(value interface{}) error {
	return e.Element.SetProperty("sl-14", value)
}

// sl-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl15)
//
// Filter slope Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl15() (interface{}, error) {
	return e.Element.GetProperty("sl-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl15(value interface{}) error {
	return e.Element.SetProperty("sl-15", value)
}

// sl-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl2)
//
// Filter slope Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl2() (interface{}, error) {
	return e.Element.GetProperty("sl-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl2(value interface{}) error {
	return e.Element.SetProperty("sl-2", value)
}

// sl-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl3)
//
// Filter slope Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl3() (interface{}, error) {
	return e.Element.GetProperty("sl-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl3(value interface{}) error {
	return e.Element.SetProperty("sl-3", value)
}

// sl-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl4)
//
// Filter slope Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl4() (interface{}, error) {
	return e.Element.GetProperty("sl-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl4(value interface{}) error {
	return e.Element.SetProperty("sl-4", value)
}

// sl-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl5)
//
// Filter slope Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl5() (interface{}, error) {
	return e.Element.GetProperty("sl-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl5(value interface{}) error {
	return e.Element.SetProperty("sl-5", value)
}

// sl-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl6)
//
// Filter slope Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl6() (interface{}, error) {
	return e.Element.GetProperty("sl-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl6(value interface{}) error {
	return e.Element.SetProperty("sl-6", value)
}

// sl-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl7)
//
// Filter slope Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl7() (interface{}, error) {
	return e.Element.GetProperty("sl-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl7(value interface{}) error {
	return e.Element.SetProperty("sl-7", value)
}

// sl-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl8)
//
// Filter slope Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl8() (interface{}, error) {
	return e.Element.GetProperty("sl-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl8(value interface{}) error {
	return e.Element.SetProperty("sl-8", value)
}

// sl-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsl9)
//
// Filter slope Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSl9() (interface{}, error) {
	return e.Element.GetProperty("sl-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSl9(value interface{}) error {
	return e.Element.SetProperty("sl-9", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSmr() (float32, error) {
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

// sr-0 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr0)
//
// Filter slope Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr0() (interface{}, error) {
	return e.Element.GetProperty("sr-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr0(value interface{}) error {
	return e.Element.SetProperty("sr-0", value)
}

// sr-1 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr1)
//
// Filter slope Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr1() (interface{}, error) {
	return e.Element.GetProperty("sr-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr1(value interface{}) error {
	return e.Element.SetProperty("sr-1", value)
}

// sr-10 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr10)
//
// Filter slope Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr10() (interface{}, error) {
	return e.Element.GetProperty("sr-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr10(value interface{}) error {
	return e.Element.SetProperty("sr-10", value)
}

// sr-11 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr11)
//
// Filter slope Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr11() (interface{}, error) {
	return e.Element.GetProperty("sr-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr11(value interface{}) error {
	return e.Element.SetProperty("sr-11", value)
}

// sr-12 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr12)
//
// Filter slope Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr12() (interface{}, error) {
	return e.Element.GetProperty("sr-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr12(value interface{}) error {
	return e.Element.SetProperty("sr-12", value)
}

// sr-13 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr13)
//
// Filter slope Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr13() (interface{}, error) {
	return e.Element.GetProperty("sr-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr13(value interface{}) error {
	return e.Element.SetProperty("sr-13", value)
}

// sr-14 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr14)
//
// Filter slope Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr14() (interface{}, error) {
	return e.Element.GetProperty("sr-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr14(value interface{}) error {
	return e.Element.SetProperty("sr-14", value)
}

// sr-15 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr15)
//
// Filter slope Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr15() (interface{}, error) {
	return e.Element.GetProperty("sr-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr15(value interface{}) error {
	return e.Element.SetProperty("sr-15", value)
}

// sr-2 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr2)
//
// Filter slope Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr2() (interface{}, error) {
	return e.Element.GetProperty("sr-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr2(value interface{}) error {
	return e.Element.SetProperty("sr-2", value)
}

// sr-3 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr3)
//
// Filter slope Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr3() (interface{}, error) {
	return e.Element.GetProperty("sr-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr3(value interface{}) error {
	return e.Element.SetProperty("sr-3", value)
}

// sr-4 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr4)
//
// Filter slope Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr4() (interface{}, error) {
	return e.Element.GetProperty("sr-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr4(value interface{}) error {
	return e.Element.SetProperty("sr-4", value)
}

// sr-5 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr5)
//
// Filter slope Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr5() (interface{}, error) {
	return e.Element.GetProperty("sr-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr5(value interface{}) error {
	return e.Element.SetProperty("sr-5", value)
}

// sr-6 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr6)
//
// Filter slope Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr6() (interface{}, error) {
	return e.Element.GetProperty("sr-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr6(value interface{}) error {
	return e.Element.SetProperty("sr-6", value)
}

// sr-7 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr7)
//
// Filter slope Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr7() (interface{}, error) {
	return e.Element.GetProperty("sr-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr7(value interface{}) error {
	return e.Element.SetProperty("sr-7", value)
}

// sr-8 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr8)
//
// Filter slope Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr8() (interface{}, error) {
	return e.Element.GetProperty("sr-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr8(value interface{}) error {
	return e.Element.SetProperty("sr-8", value)
}

// sr-9 (GstLspPlugInPluginsLv2ParaEqualizerX16Lrsr9)
//
// Filter slope Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetSr9() (interface{}, error) {
	return e.Element.GetProperty("sr-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetSr9(value interface{}) error {
	return e.Element.SetProperty("sr-9", value)
}

// xml-0 (bool)
//
// Filter mute Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml0(value bool) error {
	return e.Element.SetProperty("xml-0", value)
}

// xml-1 (bool)
//
// Filter mute Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml1(value bool) error {
	return e.Element.SetProperty("xml-1", value)
}

// xml-10 (bool)
//
// Filter mute Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml10(value bool) error {
	return e.Element.SetProperty("xml-10", value)
}

// xml-11 (bool)
//
// Filter mute Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml11(value bool) error {
	return e.Element.SetProperty("xml-11", value)
}

// xml-12 (bool)
//
// Filter mute Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml12(value bool) error {
	return e.Element.SetProperty("xml-12", value)
}

// xml-13 (bool)
//
// Filter mute Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml13(value bool) error {
	return e.Element.SetProperty("xml-13", value)
}

// xml-14 (bool)
//
// Filter mute Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml14(value bool) error {
	return e.Element.SetProperty("xml-14", value)
}

// xml-15 (bool)
//
// Filter mute Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml15(value bool) error {
	return e.Element.SetProperty("xml-15", value)
}

// xml-2 (bool)
//
// Filter mute Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml2(value bool) error {
	return e.Element.SetProperty("xml-2", value)
}

// xml-3 (bool)
//
// Filter mute Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml3(value bool) error {
	return e.Element.SetProperty("xml-3", value)
}

// xml-4 (bool)
//
// Filter mute Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml4(value bool) error {
	return e.Element.SetProperty("xml-4", value)
}

// xml-5 (bool)
//
// Filter mute Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml5(value bool) error {
	return e.Element.SetProperty("xml-5", value)
}

// xml-6 (bool)
//
// Filter mute Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml6(value bool) error {
	return e.Element.SetProperty("xml-6", value)
}

// xml-7 (bool)
//
// Filter mute Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml7(value bool) error {
	return e.Element.SetProperty("xml-7", value)
}

// xml-8 (bool)
//
// Filter mute Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml8(value bool) error {
	return e.Element.SetProperty("xml-8", value)
}

// xml-9 (bool)
//
// Filter mute Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXml9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXml9(value bool) error {
	return e.Element.SetProperty("xml-9", value)
}

// xmr-0 (bool)
//
// Filter mute Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr0(value bool) error {
	return e.Element.SetProperty("xmr-0", value)
}

// xmr-1 (bool)
//
// Filter mute Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr1(value bool) error {
	return e.Element.SetProperty("xmr-1", value)
}

// xmr-10 (bool)
//
// Filter mute Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr10(value bool) error {
	return e.Element.SetProperty("xmr-10", value)
}

// xmr-11 (bool)
//
// Filter mute Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr11(value bool) error {
	return e.Element.SetProperty("xmr-11", value)
}

// xmr-12 (bool)
//
// Filter mute Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr12(value bool) error {
	return e.Element.SetProperty("xmr-12", value)
}

// xmr-13 (bool)
//
// Filter mute Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr13(value bool) error {
	return e.Element.SetProperty("xmr-13", value)
}

// xmr-14 (bool)
//
// Filter mute Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr14(value bool) error {
	return e.Element.SetProperty("xmr-14", value)
}

// xmr-15 (bool)
//
// Filter mute Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr15(value bool) error {
	return e.Element.SetProperty("xmr-15", value)
}

// xmr-2 (bool)
//
// Filter mute Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr2(value bool) error {
	return e.Element.SetProperty("xmr-2", value)
}

// xmr-3 (bool)
//
// Filter mute Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr3(value bool) error {
	return e.Element.SetProperty("xmr-3", value)
}

// xmr-4 (bool)
//
// Filter mute Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr4(value bool) error {
	return e.Element.SetProperty("xmr-4", value)
}

// xmr-5 (bool)
//
// Filter mute Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr5(value bool) error {
	return e.Element.SetProperty("xmr-5", value)
}

// xmr-6 (bool)
//
// Filter mute Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr6(value bool) error {
	return e.Element.SetProperty("xmr-6", value)
}

// xmr-7 (bool)
//
// Filter mute Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr7(value bool) error {
	return e.Element.SetProperty("xmr-7", value)
}

// xmr-8 (bool)
//
// Filter mute Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr8(value bool) error {
	return e.Element.SetProperty("xmr-8", value)
}

// xmr-9 (bool)
//
// Filter mute Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXmr9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXmr9(value bool) error {
	return e.Element.SetProperty("xmr-9", value)
}

// xsl-0 (bool)
//
// Filter solo Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl0(value bool) error {
	return e.Element.SetProperty("xsl-0", value)
}

// xsl-1 (bool)
//
// Filter solo Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl1(value bool) error {
	return e.Element.SetProperty("xsl-1", value)
}

// xsl-10 (bool)
//
// Filter solo Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl10(value bool) error {
	return e.Element.SetProperty("xsl-10", value)
}

// xsl-11 (bool)
//
// Filter solo Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl11(value bool) error {
	return e.Element.SetProperty("xsl-11", value)
}

// xsl-12 (bool)
//
// Filter solo Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl12(value bool) error {
	return e.Element.SetProperty("xsl-12", value)
}

// xsl-13 (bool)
//
// Filter solo Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl13(value bool) error {
	return e.Element.SetProperty("xsl-13", value)
}

// xsl-14 (bool)
//
// Filter solo Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl14(value bool) error {
	return e.Element.SetProperty("xsl-14", value)
}

// xsl-15 (bool)
//
// Filter solo Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl15(value bool) error {
	return e.Element.SetProperty("xsl-15", value)
}

// xsl-2 (bool)
//
// Filter solo Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl2(value bool) error {
	return e.Element.SetProperty("xsl-2", value)
}

// xsl-3 (bool)
//
// Filter solo Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl3(value bool) error {
	return e.Element.SetProperty("xsl-3", value)
}

// xsl-4 (bool)
//
// Filter solo Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl4(value bool) error {
	return e.Element.SetProperty("xsl-4", value)
}

// xsl-5 (bool)
//
// Filter solo Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl5(value bool) error {
	return e.Element.SetProperty("xsl-5", value)
}

// xsl-6 (bool)
//
// Filter solo Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl6(value bool) error {
	return e.Element.SetProperty("xsl-6", value)
}

// xsl-7 (bool)
//
// Filter solo Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl7(value bool) error {
	return e.Element.SetProperty("xsl-7", value)
}

// xsl-8 (bool)
//
// Filter solo Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl8(value bool) error {
	return e.Element.SetProperty("xsl-8", value)
}

// xsl-9 (bool)
//
// Filter solo Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsl9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsl9(value bool) error {
	return e.Element.SetProperty("xsl-9", value)
}

// xsr-0 (bool)
//
// Filter solo Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr0(value bool) error {
	return e.Element.SetProperty("xsr-0", value)
}

// xsr-1 (bool)
//
// Filter solo Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr1(value bool) error {
	return e.Element.SetProperty("xsr-1", value)
}

// xsr-10 (bool)
//
// Filter solo Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr10(value bool) error {
	return e.Element.SetProperty("xsr-10", value)
}

// xsr-11 (bool)
//
// Filter solo Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr11(value bool) error {
	return e.Element.SetProperty("xsr-11", value)
}

// xsr-12 (bool)
//
// Filter solo Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr12(value bool) error {
	return e.Element.SetProperty("xsr-12", value)
}

// xsr-13 (bool)
//
// Filter solo Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr13(value bool) error {
	return e.Element.SetProperty("xsr-13", value)
}

// xsr-14 (bool)
//
// Filter solo Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr14(value bool) error {
	return e.Element.SetProperty("xsr-14", value)
}

// xsr-15 (bool)
//
// Filter solo Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr15(value bool) error {
	return e.Element.SetProperty("xsr-15", value)
}

// xsr-2 (bool)
//
// Filter solo Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr2(value bool) error {
	return e.Element.SetProperty("xsr-2", value)
}

// xsr-3 (bool)
//
// Filter solo Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr3(value bool) error {
	return e.Element.SetProperty("xsr-3", value)
}

// xsr-4 (bool)
//
// Filter solo Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr4(value bool) error {
	return e.Element.SetProperty("xsr-4", value)
}

// xsr-5 (bool)
//
// Filter solo Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr5(value bool) error {
	return e.Element.SetProperty("xsr-5", value)
}

// xsr-6 (bool)
//
// Filter solo Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr6(value bool) error {
	return e.Element.SetProperty("xsr-6", value)
}

// xsr-7 (bool)
//
// Filter solo Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr7(value bool) error {
	return e.Element.SetProperty("xsr-7", value)
}

// xsr-8 (bool)
//
// Filter solo Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr8(value bool) error {
	return e.Element.SetProperty("xsr-8", value)
}

// xsr-9 (bool)
//
// Filter solo Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetXsr9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetXsr9(value bool) error {
	return e.Element.SetProperty("xsr-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX16Lr) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl5Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr4X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr4X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr4X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr4X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml7ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl6Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl9X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl9X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl9X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl9X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr6X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr6X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr6X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr6X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml12ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml13ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr9X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr9X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr9X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr9X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl4X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl4X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl4X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl4X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl7Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl3X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl3X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl3X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl3X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl2Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl9Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr7Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl12X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl12X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl12X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl12X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl7X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl7X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl7X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl7X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr1X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr1X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr1X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr1X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml6ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr11X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr11X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr11X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr11X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml10ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl13Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl1X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl1X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl1X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl1X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr5X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr5X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr5X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr5X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml0ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl14X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl14X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl14X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl14X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr0X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr0X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr0X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr0X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr12X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr12X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr12X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr12X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr13X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr13X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr13X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr13X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr14X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr14X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr14X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr14X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr9Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl12Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl8Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr10Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr2Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr10X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr10X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr10X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr10X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr2X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr2X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr2X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr2X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml11ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX16LrfselFiltersLeft07 LspPlugInPluginsLv2ParaEqualizerX16Lrfsel = "Filters Left 0-7" // Filters Left 0-7 (0) – Filters Left 0-7
	LspPlugInPluginsLv2ParaEqualizerX16LrfselFiltersRight07 LspPlugInPluginsLv2ParaEqualizerX16Lrfsel = "Filters Right 0-7" // Filters Right 0-7 (1) – Filters Right 0-7
	LspPlugInPluginsLv2ParaEqualizerX16LrfselFiltersLeft815 LspPlugInPluginsLv2ParaEqualizerX16Lrfsel = "Filters Left 8-15" // Filters Left 8-15 (2) – Filters Left 8-15
	LspPlugInPluginsLv2ParaEqualizerX16LrfselFiltersRight815 LspPlugInPluginsLv2ParaEqualizerX16Lrfsel = "Filters Right 8-15" // Filters Right 8-15 (3) – Filters Right 8-15
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl3Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr13Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrmode string

const (
	LspPlugInPluginsLv2ParaEqualizerX16LrmodeIir LspPlugInPluginsLv2ParaEqualizerX16Lrmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX16LrmodeFir LspPlugInPluginsLv2ParaEqualizerX16Lrmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX16LrmodeFft LspPlugInPluginsLv2ParaEqualizerX16Lrmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl13X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl13X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl13X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl13X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfft string

const (
	LspPlugInPluginsLv2ParaEqualizerX16LrfftOff LspPlugInPluginsLv2ParaEqualizerX16Lrfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16LrfftPostEq LspPlugInPluginsLv2ParaEqualizerX16Lrfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX16LrfftPreEq LspPlugInPluginsLv2ParaEqualizerX16Lrfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml15ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr15X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr15X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr15X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr15X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl0X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl0X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl0X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl0X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl10X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl10X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl10X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl10X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml5ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr0Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr5Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl15X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl15X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl15X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl15X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl8X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl8X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl8X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl8X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml2ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml3ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl0Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl1Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr14Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr15Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr4Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml14ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl11X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl11X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl11X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl11X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr3X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr3X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr3X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr3X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl2X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl2X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl2X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl2X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl5X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl5X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl5X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl5X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl4Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl11Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr12Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr6Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml9ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl10Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl14Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15Off LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftl15Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftl15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr11Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsl6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl6X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsl6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl6X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsl6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl6X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsl6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsl6X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsl6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr7X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr7X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr7X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr7X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml4ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml8ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr1Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr3Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8Off LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8Bell LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8HiPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8HiShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8LoPass LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8LoShelf LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8Notch LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8Resonance LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX16Lrftr8Allpass LspPlugInPluginsLv2ParaEqualizerX16Lrftr8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrsr8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr8X1 LspPlugInPluginsLv2ParaEqualizerX16Lrsr8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr8X2 LspPlugInPluginsLv2ParaEqualizerX16Lrsr8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr8X3 LspPlugInPluginsLv2ParaEqualizerX16Lrsr8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX16Lrsr8X4 LspPlugInPluginsLv2ParaEqualizerX16Lrsr8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfml1ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfml1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10RlcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10RlcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10BwcBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10BwcMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10LrxBt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10LrxMt LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10ApoDr LspPlugInPluginsLv2ParaEqualizerX16Lrfmr10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

