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

type LspPlugInPluginsLv2ParaEqualizerX32Lr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ParaEqualizerX32Lr() (*LspPlugInPluginsLv2ParaEqualizerX32Lr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-para-equalizer-x32-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ParaEqualizerX32LrWithName(name string) (*LspPlugInPluginsLv2ParaEqualizerX32Lr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-para-equalizer-x32-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ParaEqualizerX32Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFftvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFftvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fl-0 (float32)
//
// Frequency Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl0(value float32) error {
	return e.Element.SetProperty("fl-0", value)
}

// fl-1 (float32)
//
// Frequency Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl1(value float32) error {
	return e.Element.SetProperty("fl-1", value)
}

// fl-10 (float32)
//
// Frequency Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl10(value float32) error {
	return e.Element.SetProperty("fl-10", value)
}

// fl-11 (float32)
//
// Frequency Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl11(value float32) error {
	return e.Element.SetProperty("fl-11", value)
}

// fl-12 (float32)
//
// Frequency Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl12(value float32) error {
	return e.Element.SetProperty("fl-12", value)
}

// fl-13 (float32)
//
// Frequency Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl13(value float32) error {
	return e.Element.SetProperty("fl-13", value)
}

// fl-14 (float32)
//
// Frequency Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl14(value float32) error {
	return e.Element.SetProperty("fl-14", value)
}

// fl-15 (float32)
//
// Frequency Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl15(value float32) error {
	return e.Element.SetProperty("fl-15", value)
}

// fl-16 (float32)
//
// Frequency Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl16(value float32) error {
	return e.Element.SetProperty("fl-16", value)
}

// fl-17 (float32)
//
// Frequency Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl17(value float32) error {
	return e.Element.SetProperty("fl-17", value)
}

// fl-18 (float32)
//
// Frequency Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl18(value float32) error {
	return e.Element.SetProperty("fl-18", value)
}

// fl-19 (float32)
//
// Frequency Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl19(value float32) error {
	return e.Element.SetProperty("fl-19", value)
}

// fl-2 (float32)
//
// Frequency Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl2(value float32) error {
	return e.Element.SetProperty("fl-2", value)
}

// fl-20 (float32)
//
// Frequency Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl20(value float32) error {
	return e.Element.SetProperty("fl-20", value)
}

// fl-21 (float32)
//
// Frequency Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl21(value float32) error {
	return e.Element.SetProperty("fl-21", value)
}

// fl-22 (float32)
//
// Frequency Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl22(value float32) error {
	return e.Element.SetProperty("fl-22", value)
}

// fl-23 (float32)
//
// Frequency Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl23(value float32) error {
	return e.Element.SetProperty("fl-23", value)
}

// fl-24 (float32)
//
// Frequency Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl24(value float32) error {
	return e.Element.SetProperty("fl-24", value)
}

// fl-25 (float32)
//
// Frequency Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl25(value float32) error {
	return e.Element.SetProperty("fl-25", value)
}

// fl-26 (float32)
//
// Frequency Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl26(value float32) error {
	return e.Element.SetProperty("fl-26", value)
}

// fl-27 (float32)
//
// Frequency Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl27(value float32) error {
	return e.Element.SetProperty("fl-27", value)
}

// fl-28 (float32)
//
// Frequency Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl28(value float32) error {
	return e.Element.SetProperty("fl-28", value)
}

// fl-29 (float32)
//
// Frequency Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl29(value float32) error {
	return e.Element.SetProperty("fl-29", value)
}

// fl-3 (float32)
//
// Frequency Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl3(value float32) error {
	return e.Element.SetProperty("fl-3", value)
}

// fl-30 (float32)
//
// Frequency Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl30(value float32) error {
	return e.Element.SetProperty("fl-30", value)
}

// fl-31 (float32)
//
// Frequency Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fl-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl31(value float32) error {
	return e.Element.SetProperty("fl-31", value)
}

// fl-4 (float32)
//
// Frequency Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl4(value float32) error {
	return e.Element.SetProperty("fl-4", value)
}

// fl-5 (float32)
//
// Frequency Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl5(value float32) error {
	return e.Element.SetProperty("fl-5", value)
}

// fl-6 (float32)
//
// Frequency Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl6(value float32) error {
	return e.Element.SetProperty("fl-6", value)
}

// fl-7 (float32)
//
// Frequency Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl7(value float32) error {
	return e.Element.SetProperty("fl-7", value)
}

// fl-8 (float32)
//
// Frequency Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl8(value float32) error {
	return e.Element.SetProperty("fl-8", value)
}

// fl-9 (float32)
//
// Frequency Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFl9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFl9(value float32) error {
	return e.Element.SetProperty("fl-9", value)
}

// fml-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml0)
//
// Filter mode Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml0() (interface{}, error) {
	return e.Element.GetProperty("fml-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml0(value interface{}) error {
	return e.Element.SetProperty("fml-0", value)
}

// fml-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml1)
//
// Filter mode Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml1() (interface{}, error) {
	return e.Element.GetProperty("fml-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml1(value interface{}) error {
	return e.Element.SetProperty("fml-1", value)
}

// fml-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml10)
//
// Filter mode Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml10() (interface{}, error) {
	return e.Element.GetProperty("fml-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml10(value interface{}) error {
	return e.Element.SetProperty("fml-10", value)
}

// fml-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml11)
//
// Filter mode Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml11() (interface{}, error) {
	return e.Element.GetProperty("fml-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml11(value interface{}) error {
	return e.Element.SetProperty("fml-11", value)
}

// fml-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml12)
//
// Filter mode Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml12() (interface{}, error) {
	return e.Element.GetProperty("fml-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml12(value interface{}) error {
	return e.Element.SetProperty("fml-12", value)
}

// fml-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml13)
//
// Filter mode Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml13() (interface{}, error) {
	return e.Element.GetProperty("fml-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml13(value interface{}) error {
	return e.Element.SetProperty("fml-13", value)
}

// fml-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml14)
//
// Filter mode Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml14() (interface{}, error) {
	return e.Element.GetProperty("fml-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml14(value interface{}) error {
	return e.Element.SetProperty("fml-14", value)
}

// fml-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml15)
//
// Filter mode Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml15() (interface{}, error) {
	return e.Element.GetProperty("fml-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml15(value interface{}) error {
	return e.Element.SetProperty("fml-15", value)
}

// fml-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml16)
//
// Filter mode Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml16() (interface{}, error) {
	return e.Element.GetProperty("fml-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml16(value interface{}) error {
	return e.Element.SetProperty("fml-16", value)
}

// fml-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml17)
//
// Filter mode Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml17() (interface{}, error) {
	return e.Element.GetProperty("fml-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml17(value interface{}) error {
	return e.Element.SetProperty("fml-17", value)
}

// fml-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml18)
//
// Filter mode Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml18() (interface{}, error) {
	return e.Element.GetProperty("fml-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml18(value interface{}) error {
	return e.Element.SetProperty("fml-18", value)
}

// fml-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml19)
//
// Filter mode Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml19() (interface{}, error) {
	return e.Element.GetProperty("fml-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml19(value interface{}) error {
	return e.Element.SetProperty("fml-19", value)
}

// fml-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml2)
//
// Filter mode Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml2() (interface{}, error) {
	return e.Element.GetProperty("fml-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml2(value interface{}) error {
	return e.Element.SetProperty("fml-2", value)
}

// fml-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml20)
//
// Filter mode Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml20() (interface{}, error) {
	return e.Element.GetProperty("fml-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml20(value interface{}) error {
	return e.Element.SetProperty("fml-20", value)
}

// fml-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml21)
//
// Filter mode Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml21() (interface{}, error) {
	return e.Element.GetProperty("fml-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml21(value interface{}) error {
	return e.Element.SetProperty("fml-21", value)
}

// fml-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml22)
//
// Filter mode Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml22() (interface{}, error) {
	return e.Element.GetProperty("fml-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml22(value interface{}) error {
	return e.Element.SetProperty("fml-22", value)
}

// fml-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml23)
//
// Filter mode Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml23() (interface{}, error) {
	return e.Element.GetProperty("fml-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml23(value interface{}) error {
	return e.Element.SetProperty("fml-23", value)
}

// fml-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml24)
//
// Filter mode Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml24() (interface{}, error) {
	return e.Element.GetProperty("fml-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml24(value interface{}) error {
	return e.Element.SetProperty("fml-24", value)
}

// fml-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml25)
//
// Filter mode Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml25() (interface{}, error) {
	return e.Element.GetProperty("fml-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml25(value interface{}) error {
	return e.Element.SetProperty("fml-25", value)
}

// fml-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml26)
//
// Filter mode Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml26() (interface{}, error) {
	return e.Element.GetProperty("fml-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml26(value interface{}) error {
	return e.Element.SetProperty("fml-26", value)
}

// fml-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml27)
//
// Filter mode Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml27() (interface{}, error) {
	return e.Element.GetProperty("fml-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml27(value interface{}) error {
	return e.Element.SetProperty("fml-27", value)
}

// fml-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml28)
//
// Filter mode Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml28() (interface{}, error) {
	return e.Element.GetProperty("fml-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml28(value interface{}) error {
	return e.Element.SetProperty("fml-28", value)
}

// fml-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml29)
//
// Filter mode Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml29() (interface{}, error) {
	return e.Element.GetProperty("fml-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml29(value interface{}) error {
	return e.Element.SetProperty("fml-29", value)
}

// fml-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml3)
//
// Filter mode Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml3() (interface{}, error) {
	return e.Element.GetProperty("fml-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml3(value interface{}) error {
	return e.Element.SetProperty("fml-3", value)
}

// fml-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml30)
//
// Filter mode Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml30() (interface{}, error) {
	return e.Element.GetProperty("fml-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml30(value interface{}) error {
	return e.Element.SetProperty("fml-30", value)
}

// fml-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml31)
//
// Filter mode Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml31() (interface{}, error) {
	return e.Element.GetProperty("fml-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml31(value interface{}) error {
	return e.Element.SetProperty("fml-31", value)
}

// fml-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml4)
//
// Filter mode Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml4() (interface{}, error) {
	return e.Element.GetProperty("fml-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml4(value interface{}) error {
	return e.Element.SetProperty("fml-4", value)
}

// fml-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml5)
//
// Filter mode Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml5() (interface{}, error) {
	return e.Element.GetProperty("fml-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml5(value interface{}) error {
	return e.Element.SetProperty("fml-5", value)
}

// fml-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml6)
//
// Filter mode Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml6() (interface{}, error) {
	return e.Element.GetProperty("fml-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml6(value interface{}) error {
	return e.Element.SetProperty("fml-6", value)
}

// fml-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml7)
//
// Filter mode Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml7() (interface{}, error) {
	return e.Element.GetProperty("fml-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml7(value interface{}) error {
	return e.Element.SetProperty("fml-7", value)
}

// fml-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml8)
//
// Filter mode Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml8() (interface{}, error) {
	return e.Element.GetProperty("fml-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml8(value interface{}) error {
	return e.Element.SetProperty("fml-8", value)
}

// fml-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfml9)
//
// Filter mode Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFml9() (interface{}, error) {
	return e.Element.GetProperty("fml-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFml9(value interface{}) error {
	return e.Element.SetProperty("fml-9", value)
}

// fmr-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr0)
//
// Filter mode Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr0() (interface{}, error) {
	return e.Element.GetProperty("fmr-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr0(value interface{}) error {
	return e.Element.SetProperty("fmr-0", value)
}

// fmr-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr1)
//
// Filter mode Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr1() (interface{}, error) {
	return e.Element.GetProperty("fmr-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr1(value interface{}) error {
	return e.Element.SetProperty("fmr-1", value)
}

// fmr-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr10)
//
// Filter mode Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr10() (interface{}, error) {
	return e.Element.GetProperty("fmr-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr10(value interface{}) error {
	return e.Element.SetProperty("fmr-10", value)
}

// fmr-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr11)
//
// Filter mode Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr11() (interface{}, error) {
	return e.Element.GetProperty("fmr-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr11(value interface{}) error {
	return e.Element.SetProperty("fmr-11", value)
}

// fmr-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr12)
//
// Filter mode Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr12() (interface{}, error) {
	return e.Element.GetProperty("fmr-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr12(value interface{}) error {
	return e.Element.SetProperty("fmr-12", value)
}

// fmr-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr13)
//
// Filter mode Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr13() (interface{}, error) {
	return e.Element.GetProperty("fmr-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr13(value interface{}) error {
	return e.Element.SetProperty("fmr-13", value)
}

// fmr-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr14)
//
// Filter mode Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr14() (interface{}, error) {
	return e.Element.GetProperty("fmr-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr14(value interface{}) error {
	return e.Element.SetProperty("fmr-14", value)
}

// fmr-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr15)
//
// Filter mode Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr15() (interface{}, error) {
	return e.Element.GetProperty("fmr-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr15(value interface{}) error {
	return e.Element.SetProperty("fmr-15", value)
}

// fmr-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr16)
//
// Filter mode Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr16() (interface{}, error) {
	return e.Element.GetProperty("fmr-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr16(value interface{}) error {
	return e.Element.SetProperty("fmr-16", value)
}

// fmr-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr17)
//
// Filter mode Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr17() (interface{}, error) {
	return e.Element.GetProperty("fmr-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr17(value interface{}) error {
	return e.Element.SetProperty("fmr-17", value)
}

// fmr-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr18)
//
// Filter mode Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr18() (interface{}, error) {
	return e.Element.GetProperty("fmr-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr18(value interface{}) error {
	return e.Element.SetProperty("fmr-18", value)
}

// fmr-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr19)
//
// Filter mode Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr19() (interface{}, error) {
	return e.Element.GetProperty("fmr-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr19(value interface{}) error {
	return e.Element.SetProperty("fmr-19", value)
}

// fmr-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr2)
//
// Filter mode Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr2() (interface{}, error) {
	return e.Element.GetProperty("fmr-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr2(value interface{}) error {
	return e.Element.SetProperty("fmr-2", value)
}

// fmr-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr20)
//
// Filter mode Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr20() (interface{}, error) {
	return e.Element.GetProperty("fmr-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr20(value interface{}) error {
	return e.Element.SetProperty("fmr-20", value)
}

// fmr-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr21)
//
// Filter mode Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr21() (interface{}, error) {
	return e.Element.GetProperty("fmr-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr21(value interface{}) error {
	return e.Element.SetProperty("fmr-21", value)
}

// fmr-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr22)
//
// Filter mode Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr22() (interface{}, error) {
	return e.Element.GetProperty("fmr-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr22(value interface{}) error {
	return e.Element.SetProperty("fmr-22", value)
}

// fmr-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr23)
//
// Filter mode Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr23() (interface{}, error) {
	return e.Element.GetProperty("fmr-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr23(value interface{}) error {
	return e.Element.SetProperty("fmr-23", value)
}

// fmr-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr24)
//
// Filter mode Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr24() (interface{}, error) {
	return e.Element.GetProperty("fmr-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr24(value interface{}) error {
	return e.Element.SetProperty("fmr-24", value)
}

// fmr-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr25)
//
// Filter mode Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr25() (interface{}, error) {
	return e.Element.GetProperty("fmr-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr25(value interface{}) error {
	return e.Element.SetProperty("fmr-25", value)
}

// fmr-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr26)
//
// Filter mode Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr26() (interface{}, error) {
	return e.Element.GetProperty("fmr-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr26(value interface{}) error {
	return e.Element.SetProperty("fmr-26", value)
}

// fmr-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr27)
//
// Filter mode Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr27() (interface{}, error) {
	return e.Element.GetProperty("fmr-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr27(value interface{}) error {
	return e.Element.SetProperty("fmr-27", value)
}

// fmr-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr28)
//
// Filter mode Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr28() (interface{}, error) {
	return e.Element.GetProperty("fmr-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr28(value interface{}) error {
	return e.Element.SetProperty("fmr-28", value)
}

// fmr-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr29)
//
// Filter mode Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr29() (interface{}, error) {
	return e.Element.GetProperty("fmr-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr29(value interface{}) error {
	return e.Element.SetProperty("fmr-29", value)
}

// fmr-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr3)
//
// Filter mode Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr3() (interface{}, error) {
	return e.Element.GetProperty("fmr-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr3(value interface{}) error {
	return e.Element.SetProperty("fmr-3", value)
}

// fmr-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr30)
//
// Filter mode Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr30() (interface{}, error) {
	return e.Element.GetProperty("fmr-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr30(value interface{}) error {
	return e.Element.SetProperty("fmr-30", value)
}

// fmr-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr31)
//
// Filter mode Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr31() (interface{}, error) {
	return e.Element.GetProperty("fmr-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr31(value interface{}) error {
	return e.Element.SetProperty("fmr-31", value)
}

// fmr-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr4)
//
// Filter mode Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr4() (interface{}, error) {
	return e.Element.GetProperty("fmr-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr4(value interface{}) error {
	return e.Element.SetProperty("fmr-4", value)
}

// fmr-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr5)
//
// Filter mode Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr5() (interface{}, error) {
	return e.Element.GetProperty("fmr-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr5(value interface{}) error {
	return e.Element.SetProperty("fmr-5", value)
}

// fmr-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr6)
//
// Filter mode Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr6() (interface{}, error) {
	return e.Element.GetProperty("fmr-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr6(value interface{}) error {
	return e.Element.SetProperty("fmr-6", value)
}

// fmr-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr7)
//
// Filter mode Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr7() (interface{}, error) {
	return e.Element.GetProperty("fmr-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr7(value interface{}) error {
	return e.Element.SetProperty("fmr-7", value)
}

// fmr-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr8)
//
// Filter mode Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr8() (interface{}, error) {
	return e.Element.GetProperty("fmr-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr8(value interface{}) error {
	return e.Element.SetProperty("fmr-8", value)
}

// fmr-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfmr9)
//
// Filter mode Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFmr9() (interface{}, error) {
	return e.Element.GetProperty("fmr-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFmr9(value interface{}) error {
	return e.Element.SetProperty("fmr-9", value)
}

// fr-0 (float32)
//
// Frequency Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr0(value float32) error {
	return e.Element.SetProperty("fr-0", value)
}

// fr-1 (float32)
//
// Frequency Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr1(value float32) error {
	return e.Element.SetProperty("fr-1", value)
}

// fr-10 (float32)
//
// Frequency Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr10(value float32) error {
	return e.Element.SetProperty("fr-10", value)
}

// fr-11 (float32)
//
// Frequency Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr11(value float32) error {
	return e.Element.SetProperty("fr-11", value)
}

// fr-12 (float32)
//
// Frequency Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr12(value float32) error {
	return e.Element.SetProperty("fr-12", value)
}

// fr-13 (float32)
//
// Frequency Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr13(value float32) error {
	return e.Element.SetProperty("fr-13", value)
}

// fr-14 (float32)
//
// Frequency Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr14(value float32) error {
	return e.Element.SetProperty("fr-14", value)
}

// fr-15 (float32)
//
// Frequency Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr15(value float32) error {
	return e.Element.SetProperty("fr-15", value)
}

// fr-16 (float32)
//
// Frequency Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr16(value float32) error {
	return e.Element.SetProperty("fr-16", value)
}

// fr-17 (float32)
//
// Frequency Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr17(value float32) error {
	return e.Element.SetProperty("fr-17", value)
}

// fr-18 (float32)
//
// Frequency Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr18(value float32) error {
	return e.Element.SetProperty("fr-18", value)
}

// fr-19 (float32)
//
// Frequency Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr19(value float32) error {
	return e.Element.SetProperty("fr-19", value)
}

// fr-2 (float32)
//
// Frequency Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr2(value float32) error {
	return e.Element.SetProperty("fr-2", value)
}

// fr-20 (float32)
//
// Frequency Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr20(value float32) error {
	return e.Element.SetProperty("fr-20", value)
}

// fr-21 (float32)
//
// Frequency Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr21(value float32) error {
	return e.Element.SetProperty("fr-21", value)
}

// fr-22 (float32)
//
// Frequency Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr22(value float32) error {
	return e.Element.SetProperty("fr-22", value)
}

// fr-23 (float32)
//
// Frequency Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr23(value float32) error {
	return e.Element.SetProperty("fr-23", value)
}

// fr-24 (float32)
//
// Frequency Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr24(value float32) error {
	return e.Element.SetProperty("fr-24", value)
}

// fr-25 (float32)
//
// Frequency Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr25(value float32) error {
	return e.Element.SetProperty("fr-25", value)
}

// fr-26 (float32)
//
// Frequency Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr26(value float32) error {
	return e.Element.SetProperty("fr-26", value)
}

// fr-27 (float32)
//
// Frequency Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr27(value float32) error {
	return e.Element.SetProperty("fr-27", value)
}

// fr-28 (float32)
//
// Frequency Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr28(value float32) error {
	return e.Element.SetProperty("fr-28", value)
}

// fr-29 (float32)
//
// Frequency Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr29(value float32) error {
	return e.Element.SetProperty("fr-29", value)
}

// fr-3 (float32)
//
// Frequency Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr3(value float32) error {
	return e.Element.SetProperty("fr-3", value)
}

// fr-30 (float32)
//
// Frequency Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr30(value float32) error {
	return e.Element.SetProperty("fr-30", value)
}

// fr-31 (float32)
//
// Frequency Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fr-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr31(value float32) error {
	return e.Element.SetProperty("fr-31", value)
}

// fr-4 (float32)
//
// Frequency Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr4(value float32) error {
	return e.Element.SetProperty("fr-4", value)
}

// fr-5 (float32)
//
// Frequency Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr5(value float32) error {
	return e.Element.SetProperty("fr-5", value)
}

// fr-6 (float32)
//
// Frequency Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr6(value float32) error {
	return e.Element.SetProperty("fr-6", value)
}

// fr-7 (float32)
//
// Frequency Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr7(value float32) error {
	return e.Element.SetProperty("fr-7", value)
}

// fr-8 (float32)
//
// Frequency Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr8(value float32) error {
	return e.Element.SetProperty("fr-8", value)
}

// fr-9 (float32)
//
// Frequency Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFr9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFr9(value float32) error {
	return e.Element.SetProperty("fr-9", value)
}

// frqs-l (float32)
//
// Frequency shift Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFrqsL() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFrqsL(value float32) error {
	return e.Element.SetProperty("frqs-l", value)
}

// frqs-r (float32)
//
// Frequency shift Right

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFrqsR() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFrqsR(value float32) error {
	return e.Element.SetProperty("frqs-r", value)
}

// fsel (GstLspPlugInPluginsLv2ParaEqualizerX32Lrfsel)
//
// Filter select

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// ftl-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl0)
//
// Filter type Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl0() (interface{}, error) {
	return e.Element.GetProperty("ftl-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl0(value interface{}) error {
	return e.Element.SetProperty("ftl-0", value)
}

// ftl-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl1)
//
// Filter type Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl1() (interface{}, error) {
	return e.Element.GetProperty("ftl-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl1(value interface{}) error {
	return e.Element.SetProperty("ftl-1", value)
}

// ftl-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl10)
//
// Filter type Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl10() (interface{}, error) {
	return e.Element.GetProperty("ftl-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl10(value interface{}) error {
	return e.Element.SetProperty("ftl-10", value)
}

// ftl-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl11)
//
// Filter type Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl11() (interface{}, error) {
	return e.Element.GetProperty("ftl-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl11(value interface{}) error {
	return e.Element.SetProperty("ftl-11", value)
}

// ftl-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl12)
//
// Filter type Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl12() (interface{}, error) {
	return e.Element.GetProperty("ftl-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl12(value interface{}) error {
	return e.Element.SetProperty("ftl-12", value)
}

// ftl-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl13)
//
// Filter type Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl13() (interface{}, error) {
	return e.Element.GetProperty("ftl-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl13(value interface{}) error {
	return e.Element.SetProperty("ftl-13", value)
}

// ftl-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl14)
//
// Filter type Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl14() (interface{}, error) {
	return e.Element.GetProperty("ftl-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl14(value interface{}) error {
	return e.Element.SetProperty("ftl-14", value)
}

// ftl-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl15)
//
// Filter type Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl15() (interface{}, error) {
	return e.Element.GetProperty("ftl-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl15(value interface{}) error {
	return e.Element.SetProperty("ftl-15", value)
}

// ftl-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl16)
//
// Filter type Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl16() (interface{}, error) {
	return e.Element.GetProperty("ftl-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl16(value interface{}) error {
	return e.Element.SetProperty("ftl-16", value)
}

// ftl-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl17)
//
// Filter type Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl17() (interface{}, error) {
	return e.Element.GetProperty("ftl-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl17(value interface{}) error {
	return e.Element.SetProperty("ftl-17", value)
}

// ftl-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl18)
//
// Filter type Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl18() (interface{}, error) {
	return e.Element.GetProperty("ftl-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl18(value interface{}) error {
	return e.Element.SetProperty("ftl-18", value)
}

// ftl-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl19)
//
// Filter type Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl19() (interface{}, error) {
	return e.Element.GetProperty("ftl-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl19(value interface{}) error {
	return e.Element.SetProperty("ftl-19", value)
}

// ftl-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl2)
//
// Filter type Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl2() (interface{}, error) {
	return e.Element.GetProperty("ftl-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl2(value interface{}) error {
	return e.Element.SetProperty("ftl-2", value)
}

// ftl-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl20)
//
// Filter type Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl20() (interface{}, error) {
	return e.Element.GetProperty("ftl-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl20(value interface{}) error {
	return e.Element.SetProperty("ftl-20", value)
}

// ftl-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl21)
//
// Filter type Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl21() (interface{}, error) {
	return e.Element.GetProperty("ftl-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl21(value interface{}) error {
	return e.Element.SetProperty("ftl-21", value)
}

// ftl-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl22)
//
// Filter type Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl22() (interface{}, error) {
	return e.Element.GetProperty("ftl-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl22(value interface{}) error {
	return e.Element.SetProperty("ftl-22", value)
}

// ftl-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl23)
//
// Filter type Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl23() (interface{}, error) {
	return e.Element.GetProperty("ftl-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl23(value interface{}) error {
	return e.Element.SetProperty("ftl-23", value)
}

// ftl-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl24)
//
// Filter type Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl24() (interface{}, error) {
	return e.Element.GetProperty("ftl-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl24(value interface{}) error {
	return e.Element.SetProperty("ftl-24", value)
}

// ftl-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl25)
//
// Filter type Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl25() (interface{}, error) {
	return e.Element.GetProperty("ftl-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl25(value interface{}) error {
	return e.Element.SetProperty("ftl-25", value)
}

// ftl-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl26)
//
// Filter type Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl26() (interface{}, error) {
	return e.Element.GetProperty("ftl-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl26(value interface{}) error {
	return e.Element.SetProperty("ftl-26", value)
}

// ftl-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl27)
//
// Filter type Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl27() (interface{}, error) {
	return e.Element.GetProperty("ftl-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl27(value interface{}) error {
	return e.Element.SetProperty("ftl-27", value)
}

// ftl-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl28)
//
// Filter type Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl28() (interface{}, error) {
	return e.Element.GetProperty("ftl-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl28(value interface{}) error {
	return e.Element.SetProperty("ftl-28", value)
}

// ftl-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl29)
//
// Filter type Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl29() (interface{}, error) {
	return e.Element.GetProperty("ftl-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl29(value interface{}) error {
	return e.Element.SetProperty("ftl-29", value)
}

// ftl-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl3)
//
// Filter type Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl3() (interface{}, error) {
	return e.Element.GetProperty("ftl-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl3(value interface{}) error {
	return e.Element.SetProperty("ftl-3", value)
}

// ftl-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl30)
//
// Filter type Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl30() (interface{}, error) {
	return e.Element.GetProperty("ftl-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl30(value interface{}) error {
	return e.Element.SetProperty("ftl-30", value)
}

// ftl-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl31)
//
// Filter type Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl31() (interface{}, error) {
	return e.Element.GetProperty("ftl-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl31(value interface{}) error {
	return e.Element.SetProperty("ftl-31", value)
}

// ftl-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl4)
//
// Filter type Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl4() (interface{}, error) {
	return e.Element.GetProperty("ftl-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl4(value interface{}) error {
	return e.Element.SetProperty("ftl-4", value)
}

// ftl-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl5)
//
// Filter type Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl5() (interface{}, error) {
	return e.Element.GetProperty("ftl-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl5(value interface{}) error {
	return e.Element.SetProperty("ftl-5", value)
}

// ftl-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl6)
//
// Filter type Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl6() (interface{}, error) {
	return e.Element.GetProperty("ftl-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl6(value interface{}) error {
	return e.Element.SetProperty("ftl-6", value)
}

// ftl-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl7)
//
// Filter type Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl7() (interface{}, error) {
	return e.Element.GetProperty("ftl-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl7(value interface{}) error {
	return e.Element.SetProperty("ftl-7", value)
}

// ftl-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl8)
//
// Filter type Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl8() (interface{}, error) {
	return e.Element.GetProperty("ftl-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl8(value interface{}) error {
	return e.Element.SetProperty("ftl-8", value)
}

// ftl-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftl9)
//
// Filter type Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtl9() (interface{}, error) {
	return e.Element.GetProperty("ftl-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtl9(value interface{}) error {
	return e.Element.SetProperty("ftl-9", value)
}

// ftr-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr0)
//
// Filter type Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr0() (interface{}, error) {
	return e.Element.GetProperty("ftr-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr0(value interface{}) error {
	return e.Element.SetProperty("ftr-0", value)
}

// ftr-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr1)
//
// Filter type Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr1() (interface{}, error) {
	return e.Element.GetProperty("ftr-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr1(value interface{}) error {
	return e.Element.SetProperty("ftr-1", value)
}

// ftr-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr10)
//
// Filter type Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr10() (interface{}, error) {
	return e.Element.GetProperty("ftr-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr10(value interface{}) error {
	return e.Element.SetProperty("ftr-10", value)
}

// ftr-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr11)
//
// Filter type Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr11() (interface{}, error) {
	return e.Element.GetProperty("ftr-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr11(value interface{}) error {
	return e.Element.SetProperty("ftr-11", value)
}

// ftr-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr12)
//
// Filter type Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr12() (interface{}, error) {
	return e.Element.GetProperty("ftr-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr12(value interface{}) error {
	return e.Element.SetProperty("ftr-12", value)
}

// ftr-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr13)
//
// Filter type Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr13() (interface{}, error) {
	return e.Element.GetProperty("ftr-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr13(value interface{}) error {
	return e.Element.SetProperty("ftr-13", value)
}

// ftr-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr14)
//
// Filter type Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr14() (interface{}, error) {
	return e.Element.GetProperty("ftr-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr14(value interface{}) error {
	return e.Element.SetProperty("ftr-14", value)
}

// ftr-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr15)
//
// Filter type Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr15() (interface{}, error) {
	return e.Element.GetProperty("ftr-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr15(value interface{}) error {
	return e.Element.SetProperty("ftr-15", value)
}

// ftr-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr16)
//
// Filter type Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr16() (interface{}, error) {
	return e.Element.GetProperty("ftr-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr16(value interface{}) error {
	return e.Element.SetProperty("ftr-16", value)
}

// ftr-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr17)
//
// Filter type Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr17() (interface{}, error) {
	return e.Element.GetProperty("ftr-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr17(value interface{}) error {
	return e.Element.SetProperty("ftr-17", value)
}

// ftr-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr18)
//
// Filter type Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr18() (interface{}, error) {
	return e.Element.GetProperty("ftr-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr18(value interface{}) error {
	return e.Element.SetProperty("ftr-18", value)
}

// ftr-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr19)
//
// Filter type Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr19() (interface{}, error) {
	return e.Element.GetProperty("ftr-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr19(value interface{}) error {
	return e.Element.SetProperty("ftr-19", value)
}

// ftr-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr2)
//
// Filter type Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr2() (interface{}, error) {
	return e.Element.GetProperty("ftr-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr2(value interface{}) error {
	return e.Element.SetProperty("ftr-2", value)
}

// ftr-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr20)
//
// Filter type Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr20() (interface{}, error) {
	return e.Element.GetProperty("ftr-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr20(value interface{}) error {
	return e.Element.SetProperty("ftr-20", value)
}

// ftr-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr21)
//
// Filter type Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr21() (interface{}, error) {
	return e.Element.GetProperty("ftr-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr21(value interface{}) error {
	return e.Element.SetProperty("ftr-21", value)
}

// ftr-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr22)
//
// Filter type Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr22() (interface{}, error) {
	return e.Element.GetProperty("ftr-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr22(value interface{}) error {
	return e.Element.SetProperty("ftr-22", value)
}

// ftr-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr23)
//
// Filter type Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr23() (interface{}, error) {
	return e.Element.GetProperty("ftr-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr23(value interface{}) error {
	return e.Element.SetProperty("ftr-23", value)
}

// ftr-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr24)
//
// Filter type Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr24() (interface{}, error) {
	return e.Element.GetProperty("ftr-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr24(value interface{}) error {
	return e.Element.SetProperty("ftr-24", value)
}

// ftr-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr25)
//
// Filter type Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr25() (interface{}, error) {
	return e.Element.GetProperty("ftr-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr25(value interface{}) error {
	return e.Element.SetProperty("ftr-25", value)
}

// ftr-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr26)
//
// Filter type Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr26() (interface{}, error) {
	return e.Element.GetProperty("ftr-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr26(value interface{}) error {
	return e.Element.SetProperty("ftr-26", value)
}

// ftr-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr27)
//
// Filter type Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr27() (interface{}, error) {
	return e.Element.GetProperty("ftr-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr27(value interface{}) error {
	return e.Element.SetProperty("ftr-27", value)
}

// ftr-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr28)
//
// Filter type Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr28() (interface{}, error) {
	return e.Element.GetProperty("ftr-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr28(value interface{}) error {
	return e.Element.SetProperty("ftr-28", value)
}

// ftr-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr29)
//
// Filter type Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr29() (interface{}, error) {
	return e.Element.GetProperty("ftr-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr29(value interface{}) error {
	return e.Element.SetProperty("ftr-29", value)
}

// ftr-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr3)
//
// Filter type Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr3() (interface{}, error) {
	return e.Element.GetProperty("ftr-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr3(value interface{}) error {
	return e.Element.SetProperty("ftr-3", value)
}

// ftr-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr30)
//
// Filter type Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr30() (interface{}, error) {
	return e.Element.GetProperty("ftr-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr30(value interface{}) error {
	return e.Element.SetProperty("ftr-30", value)
}

// ftr-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr31)
//
// Filter type Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr31() (interface{}, error) {
	return e.Element.GetProperty("ftr-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr31(value interface{}) error {
	return e.Element.SetProperty("ftr-31", value)
}

// ftr-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr4)
//
// Filter type Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr4() (interface{}, error) {
	return e.Element.GetProperty("ftr-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr4(value interface{}) error {
	return e.Element.SetProperty("ftr-4", value)
}

// ftr-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr5)
//
// Filter type Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr5() (interface{}, error) {
	return e.Element.GetProperty("ftr-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr5(value interface{}) error {
	return e.Element.SetProperty("ftr-5", value)
}

// ftr-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr6)
//
// Filter type Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr6() (interface{}, error) {
	return e.Element.GetProperty("ftr-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr6(value interface{}) error {
	return e.Element.SetProperty("ftr-6", value)
}

// ftr-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr7)
//
// Filter type Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr7() (interface{}, error) {
	return e.Element.GetProperty("ftr-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr7(value interface{}) error {
	return e.Element.SetProperty("ftr-7", value)
}

// ftr-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr8)
//
// Filter type Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr8() (interface{}, error) {
	return e.Element.GetProperty("ftr-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr8(value interface{}) error {
	return e.Element.SetProperty("ftr-8", value)
}

// ftr-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrftr9)
//
// Filter type Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFtr9() (interface{}, error) {
	return e.Element.GetProperty("ftr-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetFtr9(value interface{}) error {
	return e.Element.SetProperty("ftr-9", value)
}

// fvl-0 (bool)
//
// Filter visibility Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl15() (bool, error) {
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

// fvl-16 (bool)
//
// Filter visibility Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-17 (bool)
//
// Filter visibility Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-18 (bool)
//
// Filter visibility Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-19 (bool)
//
// Filter visibility Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-19")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl2() (bool, error) {
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

// fvl-20 (bool)
//
// Filter visibility Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-21 (bool)
//
// Filter visibility Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-22 (bool)
//
// Filter visibility Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-23 (bool)
//
// Filter visibility Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-24 (bool)
//
// Filter visibility Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-25 (bool)
//
// Filter visibility Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-26 (bool)
//
// Filter visibility Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-27 (bool)
//
// Filter visibility Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-28 (bool)
//
// Filter visibility Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-29 (bool)
//
// Filter visibility Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-29")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl3() (bool, error) {
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

// fvl-30 (bool)
//
// Filter visibility Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvl-31 (bool)
//
// Filter visibility Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvl-31")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvl9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr15() (bool, error) {
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

// fvr-16 (bool)
//
// Filter visibility Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-17 (bool)
//
// Filter visibility Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-18 (bool)
//
// Filter visibility Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-19 (bool)
//
// Filter visibility Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-19")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr2() (bool, error) {
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

// fvr-20 (bool)
//
// Filter visibility Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-21 (bool)
//
// Filter visibility Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-22 (bool)
//
// Filter visibility Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-23 (bool)
//
// Filter visibility Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-24 (bool)
//
// Filter visibility Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-25 (bool)
//
// Filter visibility Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-26 (bool)
//
// Filter visibility Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-27 (bool)
//
// Filter visibility Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-28 (bool)
//
// Filter visibility Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-29 (bool)
//
// Filter visibility Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-29")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr3() (bool, error) {
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

// fvr-30 (bool)
//
// Filter visibility Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fvr-31 (bool)
//
// Filter visibility Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fvr-31")
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetFvr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl-0 (float32)
//
// Gain Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl0(value float32) error {
	return e.Element.SetProperty("gl-0", value)
}

// gl-1 (float32)
//
// Gain Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl1(value float32) error {
	return e.Element.SetProperty("gl-1", value)
}

// gl-10 (float32)
//
// Gain Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl10(value float32) error {
	return e.Element.SetProperty("gl-10", value)
}

// gl-11 (float32)
//
// Gain Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl11(value float32) error {
	return e.Element.SetProperty("gl-11", value)
}

// gl-12 (float32)
//
// Gain Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl12(value float32) error {
	return e.Element.SetProperty("gl-12", value)
}

// gl-13 (float32)
//
// Gain Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl13(value float32) error {
	return e.Element.SetProperty("gl-13", value)
}

// gl-14 (float32)
//
// Gain Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl14(value float32) error {
	return e.Element.SetProperty("gl-14", value)
}

// gl-15 (float32)
//
// Gain Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl15(value float32) error {
	return e.Element.SetProperty("gl-15", value)
}

// gl-16 (float32)
//
// Gain Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl16(value float32) error {
	return e.Element.SetProperty("gl-16", value)
}

// gl-17 (float32)
//
// Gain Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl17(value float32) error {
	return e.Element.SetProperty("gl-17", value)
}

// gl-18 (float32)
//
// Gain Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl18(value float32) error {
	return e.Element.SetProperty("gl-18", value)
}

// gl-19 (float32)
//
// Gain Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl19(value float32) error {
	return e.Element.SetProperty("gl-19", value)
}

// gl-2 (float32)
//
// Gain Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl2(value float32) error {
	return e.Element.SetProperty("gl-2", value)
}

// gl-20 (float32)
//
// Gain Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl20(value float32) error {
	return e.Element.SetProperty("gl-20", value)
}

// gl-21 (float32)
//
// Gain Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl21(value float32) error {
	return e.Element.SetProperty("gl-21", value)
}

// gl-22 (float32)
//
// Gain Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl22(value float32) error {
	return e.Element.SetProperty("gl-22", value)
}

// gl-23 (float32)
//
// Gain Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl23(value float32) error {
	return e.Element.SetProperty("gl-23", value)
}

// gl-24 (float32)
//
// Gain Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl24(value float32) error {
	return e.Element.SetProperty("gl-24", value)
}

// gl-25 (float32)
//
// Gain Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl25(value float32) error {
	return e.Element.SetProperty("gl-25", value)
}

// gl-26 (float32)
//
// Gain Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl26(value float32) error {
	return e.Element.SetProperty("gl-26", value)
}

// gl-27 (float32)
//
// Gain Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl27(value float32) error {
	return e.Element.SetProperty("gl-27", value)
}

// gl-28 (float32)
//
// Gain Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl28(value float32) error {
	return e.Element.SetProperty("gl-28", value)
}

// gl-29 (float32)
//
// Gain Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl29(value float32) error {
	return e.Element.SetProperty("gl-29", value)
}

// gl-3 (float32)
//
// Gain Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl3(value float32) error {
	return e.Element.SetProperty("gl-3", value)
}

// gl-30 (float32)
//
// Gain Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl30(value float32) error {
	return e.Element.SetProperty("gl-30", value)
}

// gl-31 (float32)
//
// Gain Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gl-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl31(value float32) error {
	return e.Element.SetProperty("gl-31", value)
}

// gl-4 (float32)
//
// Gain Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl4(value float32) error {
	return e.Element.SetProperty("gl-4", value)
}

// gl-5 (float32)
//
// Gain Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl5(value float32) error {
	return e.Element.SetProperty("gl-5", value)
}

// gl-6 (float32)
//
// Gain Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl6(value float32) error {
	return e.Element.SetProperty("gl-6", value)
}

// gl-7 (float32)
//
// Gain Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl7(value float32) error {
	return e.Element.SetProperty("gl-7", value)
}

// gl-8 (float32)
//
// Gain Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl8(value float32) error {
	return e.Element.SetProperty("gl-8", value)
}

// gl-9 (float32)
//
// Gain Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGl9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGl9(value float32) error {
	return e.Element.SetProperty("gl-9", value)
}

// gr-0 (float32)
//
// Gain Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr0(value float32) error {
	return e.Element.SetProperty("gr-0", value)
}

// gr-1 (float32)
//
// Gain Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr1(value float32) error {
	return e.Element.SetProperty("gr-1", value)
}

// gr-10 (float32)
//
// Gain Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr10(value float32) error {
	return e.Element.SetProperty("gr-10", value)
}

// gr-11 (float32)
//
// Gain Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr11(value float32) error {
	return e.Element.SetProperty("gr-11", value)
}

// gr-12 (float32)
//
// Gain Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr12(value float32) error {
	return e.Element.SetProperty("gr-12", value)
}

// gr-13 (float32)
//
// Gain Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr13(value float32) error {
	return e.Element.SetProperty("gr-13", value)
}

// gr-14 (float32)
//
// Gain Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr14(value float32) error {
	return e.Element.SetProperty("gr-14", value)
}

// gr-15 (float32)
//
// Gain Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr15(value float32) error {
	return e.Element.SetProperty("gr-15", value)
}

// gr-16 (float32)
//
// Gain Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr16(value float32) error {
	return e.Element.SetProperty("gr-16", value)
}

// gr-17 (float32)
//
// Gain Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr17(value float32) error {
	return e.Element.SetProperty("gr-17", value)
}

// gr-18 (float32)
//
// Gain Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr18(value float32) error {
	return e.Element.SetProperty("gr-18", value)
}

// gr-19 (float32)
//
// Gain Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr19(value float32) error {
	return e.Element.SetProperty("gr-19", value)
}

// gr-2 (float32)
//
// Gain Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr2(value float32) error {
	return e.Element.SetProperty("gr-2", value)
}

// gr-20 (float32)
//
// Gain Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr20(value float32) error {
	return e.Element.SetProperty("gr-20", value)
}

// gr-21 (float32)
//
// Gain Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr21(value float32) error {
	return e.Element.SetProperty("gr-21", value)
}

// gr-22 (float32)
//
// Gain Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr22(value float32) error {
	return e.Element.SetProperty("gr-22", value)
}

// gr-23 (float32)
//
// Gain Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr23(value float32) error {
	return e.Element.SetProperty("gr-23", value)
}

// gr-24 (float32)
//
// Gain Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr24(value float32) error {
	return e.Element.SetProperty("gr-24", value)
}

// gr-25 (float32)
//
// Gain Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr25(value float32) error {
	return e.Element.SetProperty("gr-25", value)
}

// gr-26 (float32)
//
// Gain Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr26(value float32) error {
	return e.Element.SetProperty("gr-26", value)
}

// gr-27 (float32)
//
// Gain Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr27(value float32) error {
	return e.Element.SetProperty("gr-27", value)
}

// gr-28 (float32)
//
// Gain Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr28(value float32) error {
	return e.Element.SetProperty("gr-28", value)
}

// gr-29 (float32)
//
// Gain Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr29(value float32) error {
	return e.Element.SetProperty("gr-29", value)
}

// gr-3 (float32)
//
// Gain Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr3(value float32) error {
	return e.Element.SetProperty("gr-3", value)
}

// gr-30 (float32)
//
// Gain Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr30(value float32) error {
	return e.Element.SetProperty("gr-30", value)
}

// gr-31 (float32)
//
// Gain Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gr-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr31(value float32) error {
	return e.Element.SetProperty("gr-31", value)
}

// gr-4 (float32)
//
// Gain Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr4(value float32) error {
	return e.Element.SetProperty("gr-4", value)
}

// gr-5 (float32)
//
// Gain Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr5(value float32) error {
	return e.Element.SetProperty("gr-5", value)
}

// gr-6 (float32)
//
// Gain Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr6(value float32) error {
	return e.Element.SetProperty("gr-6", value)
}

// gr-7 (float32)
//
// Gain Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr7(value float32) error {
	return e.Element.SetProperty("gr-7", value)
}

// gr-8 (float32)
//
// Gain Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr8(value float32) error {
	return e.Element.SetProperty("gr-8", value)
}

// gr-9 (float32)
//
// Gain Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetGr9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetGr9(value float32) error {
	return e.Element.SetProperty("gr-9", value)
}

// huel-0 (float32)
//
// Hue Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel0(value float32) error {
	return e.Element.SetProperty("huel-0", value)
}

// huel-1 (float32)
//
// Hue Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel1(value float32) error {
	return e.Element.SetProperty("huel-1", value)
}

// huel-10 (float32)
//
// Hue Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel10(value float32) error {
	return e.Element.SetProperty("huel-10", value)
}

// huel-11 (float32)
//
// Hue Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel11(value float32) error {
	return e.Element.SetProperty("huel-11", value)
}

// huel-12 (float32)
//
// Hue Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel12(value float32) error {
	return e.Element.SetProperty("huel-12", value)
}

// huel-13 (float32)
//
// Hue Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel13(value float32) error {
	return e.Element.SetProperty("huel-13", value)
}

// huel-14 (float32)
//
// Hue Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel14(value float32) error {
	return e.Element.SetProperty("huel-14", value)
}

// huel-15 (float32)
//
// Hue Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel15(value float32) error {
	return e.Element.SetProperty("huel-15", value)
}

// huel-16 (float32)
//
// Hue Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel16(value float32) error {
	return e.Element.SetProperty("huel-16", value)
}

// huel-17 (float32)
//
// Hue Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel17(value float32) error {
	return e.Element.SetProperty("huel-17", value)
}

// huel-18 (float32)
//
// Hue Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel18(value float32) error {
	return e.Element.SetProperty("huel-18", value)
}

// huel-19 (float32)
//
// Hue Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel19(value float32) error {
	return e.Element.SetProperty("huel-19", value)
}

// huel-2 (float32)
//
// Hue Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel2(value float32) error {
	return e.Element.SetProperty("huel-2", value)
}

// huel-20 (float32)
//
// Hue Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel20(value float32) error {
	return e.Element.SetProperty("huel-20", value)
}

// huel-21 (float32)
//
// Hue Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel21(value float32) error {
	return e.Element.SetProperty("huel-21", value)
}

// huel-22 (float32)
//
// Hue Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel22(value float32) error {
	return e.Element.SetProperty("huel-22", value)
}

// huel-23 (float32)
//
// Hue Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel23(value float32) error {
	return e.Element.SetProperty("huel-23", value)
}

// huel-24 (float32)
//
// Hue Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel24(value float32) error {
	return e.Element.SetProperty("huel-24", value)
}

// huel-25 (float32)
//
// Hue Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel25(value float32) error {
	return e.Element.SetProperty("huel-25", value)
}

// huel-26 (float32)
//
// Hue Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel26(value float32) error {
	return e.Element.SetProperty("huel-26", value)
}

// huel-27 (float32)
//
// Hue Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel27(value float32) error {
	return e.Element.SetProperty("huel-27", value)
}

// huel-28 (float32)
//
// Hue Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel28(value float32) error {
	return e.Element.SetProperty("huel-28", value)
}

// huel-29 (float32)
//
// Hue Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel29(value float32) error {
	return e.Element.SetProperty("huel-29", value)
}

// huel-3 (float32)
//
// Hue Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel3(value float32) error {
	return e.Element.SetProperty("huel-3", value)
}

// huel-30 (float32)
//
// Hue Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel30(value float32) error {
	return e.Element.SetProperty("huel-30", value)
}

// huel-31 (float32)
//
// Hue Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huel-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel31(value float32) error {
	return e.Element.SetProperty("huel-31", value)
}

// huel-4 (float32)
//
// Hue Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel4(value float32) error {
	return e.Element.SetProperty("huel-4", value)
}

// huel-5 (float32)
//
// Hue Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel5(value float32) error {
	return e.Element.SetProperty("huel-5", value)
}

// huel-6 (float32)
//
// Hue Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel6(value float32) error {
	return e.Element.SetProperty("huel-6", value)
}

// huel-7 (float32)
//
// Hue Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel7(value float32) error {
	return e.Element.SetProperty("huel-7", value)
}

// huel-8 (float32)
//
// Hue Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel8(value float32) error {
	return e.Element.SetProperty("huel-8", value)
}

// huel-9 (float32)
//
// Hue Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuel9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuel9(value float32) error {
	return e.Element.SetProperty("huel-9", value)
}

// huer-0 (float32)
//
// Hue Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer0(value float32) error {
	return e.Element.SetProperty("huer-0", value)
}

// huer-1 (float32)
//
// Hue Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer1(value float32) error {
	return e.Element.SetProperty("huer-1", value)
}

// huer-10 (float32)
//
// Hue Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer10(value float32) error {
	return e.Element.SetProperty("huer-10", value)
}

// huer-11 (float32)
//
// Hue Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer11(value float32) error {
	return e.Element.SetProperty("huer-11", value)
}

// huer-12 (float32)
//
// Hue Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer12(value float32) error {
	return e.Element.SetProperty("huer-12", value)
}

// huer-13 (float32)
//
// Hue Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer13(value float32) error {
	return e.Element.SetProperty("huer-13", value)
}

// huer-14 (float32)
//
// Hue Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer14(value float32) error {
	return e.Element.SetProperty("huer-14", value)
}

// huer-15 (float32)
//
// Hue Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer15(value float32) error {
	return e.Element.SetProperty("huer-15", value)
}

// huer-16 (float32)
//
// Hue Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer16(value float32) error {
	return e.Element.SetProperty("huer-16", value)
}

// huer-17 (float32)
//
// Hue Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer17(value float32) error {
	return e.Element.SetProperty("huer-17", value)
}

// huer-18 (float32)
//
// Hue Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer18(value float32) error {
	return e.Element.SetProperty("huer-18", value)
}

// huer-19 (float32)
//
// Hue Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer19(value float32) error {
	return e.Element.SetProperty("huer-19", value)
}

// huer-2 (float32)
//
// Hue Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer2(value float32) error {
	return e.Element.SetProperty("huer-2", value)
}

// huer-20 (float32)
//
// Hue Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer20(value float32) error {
	return e.Element.SetProperty("huer-20", value)
}

// huer-21 (float32)
//
// Hue Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer21(value float32) error {
	return e.Element.SetProperty("huer-21", value)
}

// huer-22 (float32)
//
// Hue Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer22(value float32) error {
	return e.Element.SetProperty("huer-22", value)
}

// huer-23 (float32)
//
// Hue Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer23(value float32) error {
	return e.Element.SetProperty("huer-23", value)
}

// huer-24 (float32)
//
// Hue Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer24(value float32) error {
	return e.Element.SetProperty("huer-24", value)
}

// huer-25 (float32)
//
// Hue Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer25(value float32) error {
	return e.Element.SetProperty("huer-25", value)
}

// huer-26 (float32)
//
// Hue Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer26(value float32) error {
	return e.Element.SetProperty("huer-26", value)
}

// huer-27 (float32)
//
// Hue Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer27(value float32) error {
	return e.Element.SetProperty("huer-27", value)
}

// huer-28 (float32)
//
// Hue Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer28(value float32) error {
	return e.Element.SetProperty("huer-28", value)
}

// huer-29 (float32)
//
// Hue Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer29(value float32) error {
	return e.Element.SetProperty("huer-29", value)
}

// huer-3 (float32)
//
// Hue Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer3(value float32) error {
	return e.Element.SetProperty("huer-3", value)
}

// huer-30 (float32)
//
// Hue Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer30(value float32) error {
	return e.Element.SetProperty("huer-30", value)
}

// huer-31 (float32)
//
// Hue Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("huer-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer31(value float32) error {
	return e.Element.SetProperty("huer-31", value)
}

// huer-4 (float32)
//
// Hue Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer4(value float32) error {
	return e.Element.SetProperty("huer-4", value)
}

// huer-5 (float32)
//
// Hue Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer5(value float32) error {
	return e.Element.SetProperty("huer-5", value)
}

// huer-6 (float32)
//
// Hue Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer6(value float32) error {
	return e.Element.SetProperty("huer-6", value)
}

// huer-7 (float32)
//
// Hue Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer7(value float32) error {
	return e.Element.SetProperty("huer-7", value)
}

// huer-8 (float32)
//
// Hue Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer8(value float32) error {
	return e.Element.SetProperty("huer-8", value)
}

// huer-9 (float32)
//
// Hue Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetHuer9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetHuer9(value float32) error {
	return e.Element.SetProperty("huer-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetImr() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2ParaEqualizerX32Lrmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl0(value float32) error {
	return e.Element.SetProperty("ql-0", value)
}

// ql-1 (float32)
//
// Quality factor Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl1(value float32) error {
	return e.Element.SetProperty("ql-1", value)
}

// ql-10 (float32)
//
// Quality factor Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl10(value float32) error {
	return e.Element.SetProperty("ql-10", value)
}

// ql-11 (float32)
//
// Quality factor Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl11(value float32) error {
	return e.Element.SetProperty("ql-11", value)
}

// ql-12 (float32)
//
// Quality factor Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl12(value float32) error {
	return e.Element.SetProperty("ql-12", value)
}

// ql-13 (float32)
//
// Quality factor Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl13(value float32) error {
	return e.Element.SetProperty("ql-13", value)
}

// ql-14 (float32)
//
// Quality factor Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl14(value float32) error {
	return e.Element.SetProperty("ql-14", value)
}

// ql-15 (float32)
//
// Quality factor Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl15(value float32) error {
	return e.Element.SetProperty("ql-15", value)
}

// ql-16 (float32)
//
// Quality factor Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl16(value float32) error {
	return e.Element.SetProperty("ql-16", value)
}

// ql-17 (float32)
//
// Quality factor Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl17(value float32) error {
	return e.Element.SetProperty("ql-17", value)
}

// ql-18 (float32)
//
// Quality factor Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl18(value float32) error {
	return e.Element.SetProperty("ql-18", value)
}

// ql-19 (float32)
//
// Quality factor Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl19(value float32) error {
	return e.Element.SetProperty("ql-19", value)
}

// ql-2 (float32)
//
// Quality factor Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl2(value float32) error {
	return e.Element.SetProperty("ql-2", value)
}

// ql-20 (float32)
//
// Quality factor Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl20(value float32) error {
	return e.Element.SetProperty("ql-20", value)
}

// ql-21 (float32)
//
// Quality factor Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl21(value float32) error {
	return e.Element.SetProperty("ql-21", value)
}

// ql-22 (float32)
//
// Quality factor Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl22(value float32) error {
	return e.Element.SetProperty("ql-22", value)
}

// ql-23 (float32)
//
// Quality factor Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl23(value float32) error {
	return e.Element.SetProperty("ql-23", value)
}

// ql-24 (float32)
//
// Quality factor Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl24(value float32) error {
	return e.Element.SetProperty("ql-24", value)
}

// ql-25 (float32)
//
// Quality factor Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl25(value float32) error {
	return e.Element.SetProperty("ql-25", value)
}

// ql-26 (float32)
//
// Quality factor Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl26(value float32) error {
	return e.Element.SetProperty("ql-26", value)
}

// ql-27 (float32)
//
// Quality factor Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl27(value float32) error {
	return e.Element.SetProperty("ql-27", value)
}

// ql-28 (float32)
//
// Quality factor Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl28(value float32) error {
	return e.Element.SetProperty("ql-28", value)
}

// ql-29 (float32)
//
// Quality factor Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl29(value float32) error {
	return e.Element.SetProperty("ql-29", value)
}

// ql-3 (float32)
//
// Quality factor Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl3(value float32) error {
	return e.Element.SetProperty("ql-3", value)
}

// ql-30 (float32)
//
// Quality factor Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl30(value float32) error {
	return e.Element.SetProperty("ql-30", value)
}

// ql-31 (float32)
//
// Quality factor Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ql-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl31(value float32) error {
	return e.Element.SetProperty("ql-31", value)
}

// ql-4 (float32)
//
// Quality factor Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl4(value float32) error {
	return e.Element.SetProperty("ql-4", value)
}

// ql-5 (float32)
//
// Quality factor Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl5(value float32) error {
	return e.Element.SetProperty("ql-5", value)
}

// ql-6 (float32)
//
// Quality factor Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl6(value float32) error {
	return e.Element.SetProperty("ql-6", value)
}

// ql-7 (float32)
//
// Quality factor Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl7(value float32) error {
	return e.Element.SetProperty("ql-7", value)
}

// ql-8 (float32)
//
// Quality factor Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl8(value float32) error {
	return e.Element.SetProperty("ql-8", value)
}

// ql-9 (float32)
//
// Quality factor Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQl9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQl9(value float32) error {
	return e.Element.SetProperty("ql-9", value)
}

// qr-0 (float32)
//
// Quality factor Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr0(value float32) error {
	return e.Element.SetProperty("qr-0", value)
}

// qr-1 (float32)
//
// Quality factor Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr1(value float32) error {
	return e.Element.SetProperty("qr-1", value)
}

// qr-10 (float32)
//
// Quality factor Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr10() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr10(value float32) error {
	return e.Element.SetProperty("qr-10", value)
}

// qr-11 (float32)
//
// Quality factor Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr11() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr11(value float32) error {
	return e.Element.SetProperty("qr-11", value)
}

// qr-12 (float32)
//
// Quality factor Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr12() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr12(value float32) error {
	return e.Element.SetProperty("qr-12", value)
}

// qr-13 (float32)
//
// Quality factor Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr13() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr13(value float32) error {
	return e.Element.SetProperty("qr-13", value)
}

// qr-14 (float32)
//
// Quality factor Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr14() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr14(value float32) error {
	return e.Element.SetProperty("qr-14", value)
}

// qr-15 (float32)
//
// Quality factor Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr15() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr15(value float32) error {
	return e.Element.SetProperty("qr-15", value)
}

// qr-16 (float32)
//
// Quality factor Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr16(value float32) error {
	return e.Element.SetProperty("qr-16", value)
}

// qr-17 (float32)
//
// Quality factor Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr17(value float32) error {
	return e.Element.SetProperty("qr-17", value)
}

// qr-18 (float32)
//
// Quality factor Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr18(value float32) error {
	return e.Element.SetProperty("qr-18", value)
}

// qr-19 (float32)
//
// Quality factor Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr19(value float32) error {
	return e.Element.SetProperty("qr-19", value)
}

// qr-2 (float32)
//
// Quality factor Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr2(value float32) error {
	return e.Element.SetProperty("qr-2", value)
}

// qr-20 (float32)
//
// Quality factor Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr20(value float32) error {
	return e.Element.SetProperty("qr-20", value)
}

// qr-21 (float32)
//
// Quality factor Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr21(value float32) error {
	return e.Element.SetProperty("qr-21", value)
}

// qr-22 (float32)
//
// Quality factor Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr22(value float32) error {
	return e.Element.SetProperty("qr-22", value)
}

// qr-23 (float32)
//
// Quality factor Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr23(value float32) error {
	return e.Element.SetProperty("qr-23", value)
}

// qr-24 (float32)
//
// Quality factor Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr24(value float32) error {
	return e.Element.SetProperty("qr-24", value)
}

// qr-25 (float32)
//
// Quality factor Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr25(value float32) error {
	return e.Element.SetProperty("qr-25", value)
}

// qr-26 (float32)
//
// Quality factor Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr26(value float32) error {
	return e.Element.SetProperty("qr-26", value)
}

// qr-27 (float32)
//
// Quality factor Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr27(value float32) error {
	return e.Element.SetProperty("qr-27", value)
}

// qr-28 (float32)
//
// Quality factor Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr28(value float32) error {
	return e.Element.SetProperty("qr-28", value)
}

// qr-29 (float32)
//
// Quality factor Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr29(value float32) error {
	return e.Element.SetProperty("qr-29", value)
}

// qr-3 (float32)
//
// Quality factor Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr3(value float32) error {
	return e.Element.SetProperty("qr-3", value)
}

// qr-30 (float32)
//
// Quality factor Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr30(value float32) error {
	return e.Element.SetProperty("qr-30", value)
}

// qr-31 (float32)
//
// Quality factor Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qr-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr31(value float32) error {
	return e.Element.SetProperty("qr-31", value)
}

// qr-4 (float32)
//
// Quality factor Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr4(value float32) error {
	return e.Element.SetProperty("qr-4", value)
}

// qr-5 (float32)
//
// Quality factor Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr5(value float32) error {
	return e.Element.SetProperty("qr-5", value)
}

// qr-6 (float32)
//
// Quality factor Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr6(value float32) error {
	return e.Element.SetProperty("qr-6", value)
}

// qr-7 (float32)
//
// Quality factor Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr7(value float32) error {
	return e.Element.SetProperty("qr-7", value)
}

// qr-8 (float32)
//
// Quality factor Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr8() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr8(value float32) error {
	return e.Element.SetProperty("qr-8", value)
}

// qr-9 (float32)
//
// Quality factor Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetQr9() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetQr9(value float32) error {
	return e.Element.SetProperty("qr-9", value)
}

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sl-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl0)
//
// Filter slope Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl0() (interface{}, error) {
	return e.Element.GetProperty("sl-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl0(value interface{}) error {
	return e.Element.SetProperty("sl-0", value)
}

// sl-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl1)
//
// Filter slope Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl1() (interface{}, error) {
	return e.Element.GetProperty("sl-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl1(value interface{}) error {
	return e.Element.SetProperty("sl-1", value)
}

// sl-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl10)
//
// Filter slope Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl10() (interface{}, error) {
	return e.Element.GetProperty("sl-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl10(value interface{}) error {
	return e.Element.SetProperty("sl-10", value)
}

// sl-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl11)
//
// Filter slope Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl11() (interface{}, error) {
	return e.Element.GetProperty("sl-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl11(value interface{}) error {
	return e.Element.SetProperty("sl-11", value)
}

// sl-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl12)
//
// Filter slope Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl12() (interface{}, error) {
	return e.Element.GetProperty("sl-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl12(value interface{}) error {
	return e.Element.SetProperty("sl-12", value)
}

// sl-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl13)
//
// Filter slope Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl13() (interface{}, error) {
	return e.Element.GetProperty("sl-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl13(value interface{}) error {
	return e.Element.SetProperty("sl-13", value)
}

// sl-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl14)
//
// Filter slope Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl14() (interface{}, error) {
	return e.Element.GetProperty("sl-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl14(value interface{}) error {
	return e.Element.SetProperty("sl-14", value)
}

// sl-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl15)
//
// Filter slope Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl15() (interface{}, error) {
	return e.Element.GetProperty("sl-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl15(value interface{}) error {
	return e.Element.SetProperty("sl-15", value)
}

// sl-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl16)
//
// Filter slope Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl16() (interface{}, error) {
	return e.Element.GetProperty("sl-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl16(value interface{}) error {
	return e.Element.SetProperty("sl-16", value)
}

// sl-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl17)
//
// Filter slope Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl17() (interface{}, error) {
	return e.Element.GetProperty("sl-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl17(value interface{}) error {
	return e.Element.SetProperty("sl-17", value)
}

// sl-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl18)
//
// Filter slope Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl18() (interface{}, error) {
	return e.Element.GetProperty("sl-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl18(value interface{}) error {
	return e.Element.SetProperty("sl-18", value)
}

// sl-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl19)
//
// Filter slope Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl19() (interface{}, error) {
	return e.Element.GetProperty("sl-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl19(value interface{}) error {
	return e.Element.SetProperty("sl-19", value)
}

// sl-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl2)
//
// Filter slope Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl2() (interface{}, error) {
	return e.Element.GetProperty("sl-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl2(value interface{}) error {
	return e.Element.SetProperty("sl-2", value)
}

// sl-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl20)
//
// Filter slope Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl20() (interface{}, error) {
	return e.Element.GetProperty("sl-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl20(value interface{}) error {
	return e.Element.SetProperty("sl-20", value)
}

// sl-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl21)
//
// Filter slope Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl21() (interface{}, error) {
	return e.Element.GetProperty("sl-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl21(value interface{}) error {
	return e.Element.SetProperty("sl-21", value)
}

// sl-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl22)
//
// Filter slope Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl22() (interface{}, error) {
	return e.Element.GetProperty("sl-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl22(value interface{}) error {
	return e.Element.SetProperty("sl-22", value)
}

// sl-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl23)
//
// Filter slope Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl23() (interface{}, error) {
	return e.Element.GetProperty("sl-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl23(value interface{}) error {
	return e.Element.SetProperty("sl-23", value)
}

// sl-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl24)
//
// Filter slope Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl24() (interface{}, error) {
	return e.Element.GetProperty("sl-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl24(value interface{}) error {
	return e.Element.SetProperty("sl-24", value)
}

// sl-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl25)
//
// Filter slope Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl25() (interface{}, error) {
	return e.Element.GetProperty("sl-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl25(value interface{}) error {
	return e.Element.SetProperty("sl-25", value)
}

// sl-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl26)
//
// Filter slope Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl26() (interface{}, error) {
	return e.Element.GetProperty("sl-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl26(value interface{}) error {
	return e.Element.SetProperty("sl-26", value)
}

// sl-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl27)
//
// Filter slope Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl27() (interface{}, error) {
	return e.Element.GetProperty("sl-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl27(value interface{}) error {
	return e.Element.SetProperty("sl-27", value)
}

// sl-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl28)
//
// Filter slope Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl28() (interface{}, error) {
	return e.Element.GetProperty("sl-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl28(value interface{}) error {
	return e.Element.SetProperty("sl-28", value)
}

// sl-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl29)
//
// Filter slope Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl29() (interface{}, error) {
	return e.Element.GetProperty("sl-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl29(value interface{}) error {
	return e.Element.SetProperty("sl-29", value)
}

// sl-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl3)
//
// Filter slope Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl3() (interface{}, error) {
	return e.Element.GetProperty("sl-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl3(value interface{}) error {
	return e.Element.SetProperty("sl-3", value)
}

// sl-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl30)
//
// Filter slope Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl30() (interface{}, error) {
	return e.Element.GetProperty("sl-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl30(value interface{}) error {
	return e.Element.SetProperty("sl-30", value)
}

// sl-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl31)
//
// Filter slope Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl31() (interface{}, error) {
	return e.Element.GetProperty("sl-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl31(value interface{}) error {
	return e.Element.SetProperty("sl-31", value)
}

// sl-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl4)
//
// Filter slope Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl4() (interface{}, error) {
	return e.Element.GetProperty("sl-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl4(value interface{}) error {
	return e.Element.SetProperty("sl-4", value)
}

// sl-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl5)
//
// Filter slope Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl5() (interface{}, error) {
	return e.Element.GetProperty("sl-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl5(value interface{}) error {
	return e.Element.SetProperty("sl-5", value)
}

// sl-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl6)
//
// Filter slope Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl6() (interface{}, error) {
	return e.Element.GetProperty("sl-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl6(value interface{}) error {
	return e.Element.SetProperty("sl-6", value)
}

// sl-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl7)
//
// Filter slope Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl7() (interface{}, error) {
	return e.Element.GetProperty("sl-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl7(value interface{}) error {
	return e.Element.SetProperty("sl-7", value)
}

// sl-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl8)
//
// Filter slope Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl8() (interface{}, error) {
	return e.Element.GetProperty("sl-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl8(value interface{}) error {
	return e.Element.SetProperty("sl-8", value)
}

// sl-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsl9)
//
// Filter slope Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSl9() (interface{}, error) {
	return e.Element.GetProperty("sl-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSl9(value interface{}) error {
	return e.Element.SetProperty("sl-9", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSmr() (float32, error) {
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

// sr-0 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr0)
//
// Filter slope Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr0() (interface{}, error) {
	return e.Element.GetProperty("sr-0")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr0(value interface{}) error {
	return e.Element.SetProperty("sr-0", value)
}

// sr-1 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr1)
//
// Filter slope Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr1() (interface{}, error) {
	return e.Element.GetProperty("sr-1")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr1(value interface{}) error {
	return e.Element.SetProperty("sr-1", value)
}

// sr-10 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr10)
//
// Filter slope Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr10() (interface{}, error) {
	return e.Element.GetProperty("sr-10")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr10(value interface{}) error {
	return e.Element.SetProperty("sr-10", value)
}

// sr-11 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr11)
//
// Filter slope Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr11() (interface{}, error) {
	return e.Element.GetProperty("sr-11")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr11(value interface{}) error {
	return e.Element.SetProperty("sr-11", value)
}

// sr-12 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr12)
//
// Filter slope Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr12() (interface{}, error) {
	return e.Element.GetProperty("sr-12")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr12(value interface{}) error {
	return e.Element.SetProperty("sr-12", value)
}

// sr-13 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr13)
//
// Filter slope Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr13() (interface{}, error) {
	return e.Element.GetProperty("sr-13")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr13(value interface{}) error {
	return e.Element.SetProperty("sr-13", value)
}

// sr-14 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr14)
//
// Filter slope Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr14() (interface{}, error) {
	return e.Element.GetProperty("sr-14")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr14(value interface{}) error {
	return e.Element.SetProperty("sr-14", value)
}

// sr-15 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr15)
//
// Filter slope Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr15() (interface{}, error) {
	return e.Element.GetProperty("sr-15")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr15(value interface{}) error {
	return e.Element.SetProperty("sr-15", value)
}

// sr-16 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr16)
//
// Filter slope Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr16() (interface{}, error) {
	return e.Element.GetProperty("sr-16")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr16(value interface{}) error {
	return e.Element.SetProperty("sr-16", value)
}

// sr-17 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr17)
//
// Filter slope Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr17() (interface{}, error) {
	return e.Element.GetProperty("sr-17")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr17(value interface{}) error {
	return e.Element.SetProperty("sr-17", value)
}

// sr-18 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr18)
//
// Filter slope Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr18() (interface{}, error) {
	return e.Element.GetProperty("sr-18")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr18(value interface{}) error {
	return e.Element.SetProperty("sr-18", value)
}

// sr-19 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr19)
//
// Filter slope Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr19() (interface{}, error) {
	return e.Element.GetProperty("sr-19")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr19(value interface{}) error {
	return e.Element.SetProperty("sr-19", value)
}

// sr-2 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr2)
//
// Filter slope Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr2() (interface{}, error) {
	return e.Element.GetProperty("sr-2")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr2(value interface{}) error {
	return e.Element.SetProperty("sr-2", value)
}

// sr-20 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr20)
//
// Filter slope Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr20() (interface{}, error) {
	return e.Element.GetProperty("sr-20")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr20(value interface{}) error {
	return e.Element.SetProperty("sr-20", value)
}

// sr-21 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr21)
//
// Filter slope Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr21() (interface{}, error) {
	return e.Element.GetProperty("sr-21")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr21(value interface{}) error {
	return e.Element.SetProperty("sr-21", value)
}

// sr-22 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr22)
//
// Filter slope Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr22() (interface{}, error) {
	return e.Element.GetProperty("sr-22")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr22(value interface{}) error {
	return e.Element.SetProperty("sr-22", value)
}

// sr-23 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr23)
//
// Filter slope Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr23() (interface{}, error) {
	return e.Element.GetProperty("sr-23")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr23(value interface{}) error {
	return e.Element.SetProperty("sr-23", value)
}

// sr-24 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr24)
//
// Filter slope Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr24() (interface{}, error) {
	return e.Element.GetProperty("sr-24")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr24(value interface{}) error {
	return e.Element.SetProperty("sr-24", value)
}

// sr-25 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr25)
//
// Filter slope Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr25() (interface{}, error) {
	return e.Element.GetProperty("sr-25")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr25(value interface{}) error {
	return e.Element.SetProperty("sr-25", value)
}

// sr-26 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr26)
//
// Filter slope Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr26() (interface{}, error) {
	return e.Element.GetProperty("sr-26")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr26(value interface{}) error {
	return e.Element.SetProperty("sr-26", value)
}

// sr-27 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr27)
//
// Filter slope Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr27() (interface{}, error) {
	return e.Element.GetProperty("sr-27")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr27(value interface{}) error {
	return e.Element.SetProperty("sr-27", value)
}

// sr-28 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr28)
//
// Filter slope Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr28() (interface{}, error) {
	return e.Element.GetProperty("sr-28")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr28(value interface{}) error {
	return e.Element.SetProperty("sr-28", value)
}

// sr-29 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr29)
//
// Filter slope Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr29() (interface{}, error) {
	return e.Element.GetProperty("sr-29")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr29(value interface{}) error {
	return e.Element.SetProperty("sr-29", value)
}

// sr-3 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr3)
//
// Filter slope Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr3() (interface{}, error) {
	return e.Element.GetProperty("sr-3")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr3(value interface{}) error {
	return e.Element.SetProperty("sr-3", value)
}

// sr-30 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr30)
//
// Filter slope Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr30() (interface{}, error) {
	return e.Element.GetProperty("sr-30")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr30(value interface{}) error {
	return e.Element.SetProperty("sr-30", value)
}

// sr-31 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr31)
//
// Filter slope Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr31() (interface{}, error) {
	return e.Element.GetProperty("sr-31")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr31(value interface{}) error {
	return e.Element.SetProperty("sr-31", value)
}

// sr-4 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr4)
//
// Filter slope Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr4() (interface{}, error) {
	return e.Element.GetProperty("sr-4")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr4(value interface{}) error {
	return e.Element.SetProperty("sr-4", value)
}

// sr-5 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr5)
//
// Filter slope Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr5() (interface{}, error) {
	return e.Element.GetProperty("sr-5")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr5(value interface{}) error {
	return e.Element.SetProperty("sr-5", value)
}

// sr-6 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr6)
//
// Filter slope Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr6() (interface{}, error) {
	return e.Element.GetProperty("sr-6")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr6(value interface{}) error {
	return e.Element.SetProperty("sr-6", value)
}

// sr-7 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr7)
//
// Filter slope Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr7() (interface{}, error) {
	return e.Element.GetProperty("sr-7")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr7(value interface{}) error {
	return e.Element.SetProperty("sr-7", value)
}

// sr-8 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr8)
//
// Filter slope Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr8() (interface{}, error) {
	return e.Element.GetProperty("sr-8")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr8(value interface{}) error {
	return e.Element.SetProperty("sr-8", value)
}

// sr-9 (GstLspPlugInPluginsLv2ParaEqualizerX32Lrsr9)
//
// Filter slope Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetSr9() (interface{}, error) {
	return e.Element.GetProperty("sr-9")
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetSr9(value interface{}) error {
	return e.Element.SetProperty("sr-9", value)
}

// xml-0 (bool)
//
// Filter mute Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml0(value bool) error {
	return e.Element.SetProperty("xml-0", value)
}

// xml-1 (bool)
//
// Filter mute Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml1(value bool) error {
	return e.Element.SetProperty("xml-1", value)
}

// xml-10 (bool)
//
// Filter mute Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml10(value bool) error {
	return e.Element.SetProperty("xml-10", value)
}

// xml-11 (bool)
//
// Filter mute Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml11(value bool) error {
	return e.Element.SetProperty("xml-11", value)
}

// xml-12 (bool)
//
// Filter mute Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml12(value bool) error {
	return e.Element.SetProperty("xml-12", value)
}

// xml-13 (bool)
//
// Filter mute Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml13(value bool) error {
	return e.Element.SetProperty("xml-13", value)
}

// xml-14 (bool)
//
// Filter mute Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml14(value bool) error {
	return e.Element.SetProperty("xml-14", value)
}

// xml-15 (bool)
//
// Filter mute Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml15(value bool) error {
	return e.Element.SetProperty("xml-15", value)
}

// xml-16 (bool)
//
// Filter mute Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml16(value bool) error {
	return e.Element.SetProperty("xml-16", value)
}

// xml-17 (bool)
//
// Filter mute Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml17(value bool) error {
	return e.Element.SetProperty("xml-17", value)
}

// xml-18 (bool)
//
// Filter mute Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml18(value bool) error {
	return e.Element.SetProperty("xml-18", value)
}

// xml-19 (bool)
//
// Filter mute Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml19(value bool) error {
	return e.Element.SetProperty("xml-19", value)
}

// xml-2 (bool)
//
// Filter mute Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml2(value bool) error {
	return e.Element.SetProperty("xml-2", value)
}

// xml-20 (bool)
//
// Filter mute Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml20(value bool) error {
	return e.Element.SetProperty("xml-20", value)
}

// xml-21 (bool)
//
// Filter mute Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml21(value bool) error {
	return e.Element.SetProperty("xml-21", value)
}

// xml-22 (bool)
//
// Filter mute Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml22(value bool) error {
	return e.Element.SetProperty("xml-22", value)
}

// xml-23 (bool)
//
// Filter mute Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml23(value bool) error {
	return e.Element.SetProperty("xml-23", value)
}

// xml-24 (bool)
//
// Filter mute Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml24(value bool) error {
	return e.Element.SetProperty("xml-24", value)
}

// xml-25 (bool)
//
// Filter mute Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml25(value bool) error {
	return e.Element.SetProperty("xml-25", value)
}

// xml-26 (bool)
//
// Filter mute Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml26(value bool) error {
	return e.Element.SetProperty("xml-26", value)
}

// xml-27 (bool)
//
// Filter mute Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml27(value bool) error {
	return e.Element.SetProperty("xml-27", value)
}

// xml-28 (bool)
//
// Filter mute Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml28(value bool) error {
	return e.Element.SetProperty("xml-28", value)
}

// xml-29 (bool)
//
// Filter mute Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml29(value bool) error {
	return e.Element.SetProperty("xml-29", value)
}

// xml-3 (bool)
//
// Filter mute Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml3(value bool) error {
	return e.Element.SetProperty("xml-3", value)
}

// xml-30 (bool)
//
// Filter mute Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml30(value bool) error {
	return e.Element.SetProperty("xml-30", value)
}

// xml-31 (bool)
//
// Filter mute Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xml-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml31(value bool) error {
	return e.Element.SetProperty("xml-31", value)
}

// xml-4 (bool)
//
// Filter mute Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml4(value bool) error {
	return e.Element.SetProperty("xml-4", value)
}

// xml-5 (bool)
//
// Filter mute Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml5(value bool) error {
	return e.Element.SetProperty("xml-5", value)
}

// xml-6 (bool)
//
// Filter mute Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml6(value bool) error {
	return e.Element.SetProperty("xml-6", value)
}

// xml-7 (bool)
//
// Filter mute Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml7(value bool) error {
	return e.Element.SetProperty("xml-7", value)
}

// xml-8 (bool)
//
// Filter mute Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml8(value bool) error {
	return e.Element.SetProperty("xml-8", value)
}

// xml-9 (bool)
//
// Filter mute Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXml9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXml9(value bool) error {
	return e.Element.SetProperty("xml-9", value)
}

// xmr-0 (bool)
//
// Filter mute Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr0(value bool) error {
	return e.Element.SetProperty("xmr-0", value)
}

// xmr-1 (bool)
//
// Filter mute Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr1(value bool) error {
	return e.Element.SetProperty("xmr-1", value)
}

// xmr-10 (bool)
//
// Filter mute Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr10(value bool) error {
	return e.Element.SetProperty("xmr-10", value)
}

// xmr-11 (bool)
//
// Filter mute Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr11(value bool) error {
	return e.Element.SetProperty("xmr-11", value)
}

// xmr-12 (bool)
//
// Filter mute Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr12(value bool) error {
	return e.Element.SetProperty("xmr-12", value)
}

// xmr-13 (bool)
//
// Filter mute Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr13(value bool) error {
	return e.Element.SetProperty("xmr-13", value)
}

// xmr-14 (bool)
//
// Filter mute Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr14(value bool) error {
	return e.Element.SetProperty("xmr-14", value)
}

// xmr-15 (bool)
//
// Filter mute Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr15(value bool) error {
	return e.Element.SetProperty("xmr-15", value)
}

// xmr-16 (bool)
//
// Filter mute Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr16(value bool) error {
	return e.Element.SetProperty("xmr-16", value)
}

// xmr-17 (bool)
//
// Filter mute Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr17(value bool) error {
	return e.Element.SetProperty("xmr-17", value)
}

// xmr-18 (bool)
//
// Filter mute Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr18(value bool) error {
	return e.Element.SetProperty("xmr-18", value)
}

// xmr-19 (bool)
//
// Filter mute Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr19(value bool) error {
	return e.Element.SetProperty("xmr-19", value)
}

// xmr-2 (bool)
//
// Filter mute Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr2(value bool) error {
	return e.Element.SetProperty("xmr-2", value)
}

// xmr-20 (bool)
//
// Filter mute Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr20(value bool) error {
	return e.Element.SetProperty("xmr-20", value)
}

// xmr-21 (bool)
//
// Filter mute Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr21(value bool) error {
	return e.Element.SetProperty("xmr-21", value)
}

// xmr-22 (bool)
//
// Filter mute Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr22(value bool) error {
	return e.Element.SetProperty("xmr-22", value)
}

// xmr-23 (bool)
//
// Filter mute Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr23(value bool) error {
	return e.Element.SetProperty("xmr-23", value)
}

// xmr-24 (bool)
//
// Filter mute Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr24(value bool) error {
	return e.Element.SetProperty("xmr-24", value)
}

// xmr-25 (bool)
//
// Filter mute Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr25(value bool) error {
	return e.Element.SetProperty("xmr-25", value)
}

// xmr-26 (bool)
//
// Filter mute Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr26(value bool) error {
	return e.Element.SetProperty("xmr-26", value)
}

// xmr-27 (bool)
//
// Filter mute Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr27(value bool) error {
	return e.Element.SetProperty("xmr-27", value)
}

// xmr-28 (bool)
//
// Filter mute Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr28(value bool) error {
	return e.Element.SetProperty("xmr-28", value)
}

// xmr-29 (bool)
//
// Filter mute Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr29(value bool) error {
	return e.Element.SetProperty("xmr-29", value)
}

// xmr-3 (bool)
//
// Filter mute Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr3(value bool) error {
	return e.Element.SetProperty("xmr-3", value)
}

// xmr-30 (bool)
//
// Filter mute Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr30(value bool) error {
	return e.Element.SetProperty("xmr-30", value)
}

// xmr-31 (bool)
//
// Filter mute Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xmr-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr31(value bool) error {
	return e.Element.SetProperty("xmr-31", value)
}

// xmr-4 (bool)
//
// Filter mute Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr4(value bool) error {
	return e.Element.SetProperty("xmr-4", value)
}

// xmr-5 (bool)
//
// Filter mute Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr5(value bool) error {
	return e.Element.SetProperty("xmr-5", value)
}

// xmr-6 (bool)
//
// Filter mute Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr6(value bool) error {
	return e.Element.SetProperty("xmr-6", value)
}

// xmr-7 (bool)
//
// Filter mute Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr7(value bool) error {
	return e.Element.SetProperty("xmr-7", value)
}

// xmr-8 (bool)
//
// Filter mute Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr8(value bool) error {
	return e.Element.SetProperty("xmr-8", value)
}

// xmr-9 (bool)
//
// Filter mute Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXmr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXmr9(value bool) error {
	return e.Element.SetProperty("xmr-9", value)
}

// xsl-0 (bool)
//
// Filter solo Left 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl0(value bool) error {
	return e.Element.SetProperty("xsl-0", value)
}

// xsl-1 (bool)
//
// Filter solo Left 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl1(value bool) error {
	return e.Element.SetProperty("xsl-1", value)
}

// xsl-10 (bool)
//
// Filter solo Left 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl10(value bool) error {
	return e.Element.SetProperty("xsl-10", value)
}

// xsl-11 (bool)
//
// Filter solo Left 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl11(value bool) error {
	return e.Element.SetProperty("xsl-11", value)
}

// xsl-12 (bool)
//
// Filter solo Left 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl12(value bool) error {
	return e.Element.SetProperty("xsl-12", value)
}

// xsl-13 (bool)
//
// Filter solo Left 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl13(value bool) error {
	return e.Element.SetProperty("xsl-13", value)
}

// xsl-14 (bool)
//
// Filter solo Left 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl14(value bool) error {
	return e.Element.SetProperty("xsl-14", value)
}

// xsl-15 (bool)
//
// Filter solo Left 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl15(value bool) error {
	return e.Element.SetProperty("xsl-15", value)
}

// xsl-16 (bool)
//
// Filter solo Left 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl16(value bool) error {
	return e.Element.SetProperty("xsl-16", value)
}

// xsl-17 (bool)
//
// Filter solo Left 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl17(value bool) error {
	return e.Element.SetProperty("xsl-17", value)
}

// xsl-18 (bool)
//
// Filter solo Left 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl18(value bool) error {
	return e.Element.SetProperty("xsl-18", value)
}

// xsl-19 (bool)
//
// Filter solo Left 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl19(value bool) error {
	return e.Element.SetProperty("xsl-19", value)
}

// xsl-2 (bool)
//
// Filter solo Left 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl2(value bool) error {
	return e.Element.SetProperty("xsl-2", value)
}

// xsl-20 (bool)
//
// Filter solo Left 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl20(value bool) error {
	return e.Element.SetProperty("xsl-20", value)
}

// xsl-21 (bool)
//
// Filter solo Left 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl21(value bool) error {
	return e.Element.SetProperty("xsl-21", value)
}

// xsl-22 (bool)
//
// Filter solo Left 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl22(value bool) error {
	return e.Element.SetProperty("xsl-22", value)
}

// xsl-23 (bool)
//
// Filter solo Left 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl23(value bool) error {
	return e.Element.SetProperty("xsl-23", value)
}

// xsl-24 (bool)
//
// Filter solo Left 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl24(value bool) error {
	return e.Element.SetProperty("xsl-24", value)
}

// xsl-25 (bool)
//
// Filter solo Left 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl25(value bool) error {
	return e.Element.SetProperty("xsl-25", value)
}

// xsl-26 (bool)
//
// Filter solo Left 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl26(value bool) error {
	return e.Element.SetProperty("xsl-26", value)
}

// xsl-27 (bool)
//
// Filter solo Left 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl27(value bool) error {
	return e.Element.SetProperty("xsl-27", value)
}

// xsl-28 (bool)
//
// Filter solo Left 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl28(value bool) error {
	return e.Element.SetProperty("xsl-28", value)
}

// xsl-29 (bool)
//
// Filter solo Left 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl29(value bool) error {
	return e.Element.SetProperty("xsl-29", value)
}

// xsl-3 (bool)
//
// Filter solo Left 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl3(value bool) error {
	return e.Element.SetProperty("xsl-3", value)
}

// xsl-30 (bool)
//
// Filter solo Left 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl30(value bool) error {
	return e.Element.SetProperty("xsl-30", value)
}

// xsl-31 (bool)
//
// Filter solo Left 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsl-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl31(value bool) error {
	return e.Element.SetProperty("xsl-31", value)
}

// xsl-4 (bool)
//
// Filter solo Left 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl4(value bool) error {
	return e.Element.SetProperty("xsl-4", value)
}

// xsl-5 (bool)
//
// Filter solo Left 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl5(value bool) error {
	return e.Element.SetProperty("xsl-5", value)
}

// xsl-6 (bool)
//
// Filter solo Left 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl6(value bool) error {
	return e.Element.SetProperty("xsl-6", value)
}

// xsl-7 (bool)
//
// Filter solo Left 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl7(value bool) error {
	return e.Element.SetProperty("xsl-7", value)
}

// xsl-8 (bool)
//
// Filter solo Left 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl8(value bool) error {
	return e.Element.SetProperty("xsl-8", value)
}

// xsl-9 (bool)
//
// Filter solo Left 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsl9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsl9(value bool) error {
	return e.Element.SetProperty("xsl-9", value)
}

// xsr-0 (bool)
//
// Filter solo Right 0

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr0(value bool) error {
	return e.Element.SetProperty("xsr-0", value)
}

// xsr-1 (bool)
//
// Filter solo Right 1

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr1(value bool) error {
	return e.Element.SetProperty("xsr-1", value)
}

// xsr-10 (bool)
//
// Filter solo Right 10

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr10(value bool) error {
	return e.Element.SetProperty("xsr-10", value)
}

// xsr-11 (bool)
//
// Filter solo Right 11

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr11(value bool) error {
	return e.Element.SetProperty("xsr-11", value)
}

// xsr-12 (bool)
//
// Filter solo Right 12

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr12(value bool) error {
	return e.Element.SetProperty("xsr-12", value)
}

// xsr-13 (bool)
//
// Filter solo Right 13

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr13(value bool) error {
	return e.Element.SetProperty("xsr-13", value)
}

// xsr-14 (bool)
//
// Filter solo Right 14

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr14(value bool) error {
	return e.Element.SetProperty("xsr-14", value)
}

// xsr-15 (bool)
//
// Filter solo Right 15

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr15() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr15(value bool) error {
	return e.Element.SetProperty("xsr-15", value)
}

// xsr-16 (bool)
//
// Filter solo Right 16

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr16(value bool) error {
	return e.Element.SetProperty("xsr-16", value)
}

// xsr-17 (bool)
//
// Filter solo Right 17

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr17(value bool) error {
	return e.Element.SetProperty("xsr-17", value)
}

// xsr-18 (bool)
//
// Filter solo Right 18

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr18(value bool) error {
	return e.Element.SetProperty("xsr-18", value)
}

// xsr-19 (bool)
//
// Filter solo Right 19

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr19(value bool) error {
	return e.Element.SetProperty("xsr-19", value)
}

// xsr-2 (bool)
//
// Filter solo Right 2

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr2(value bool) error {
	return e.Element.SetProperty("xsr-2", value)
}

// xsr-20 (bool)
//
// Filter solo Right 20

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr20(value bool) error {
	return e.Element.SetProperty("xsr-20", value)
}

// xsr-21 (bool)
//
// Filter solo Right 21

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr21(value bool) error {
	return e.Element.SetProperty("xsr-21", value)
}

// xsr-22 (bool)
//
// Filter solo Right 22

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr22(value bool) error {
	return e.Element.SetProperty("xsr-22", value)
}

// xsr-23 (bool)
//
// Filter solo Right 23

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr23(value bool) error {
	return e.Element.SetProperty("xsr-23", value)
}

// xsr-24 (bool)
//
// Filter solo Right 24

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr24(value bool) error {
	return e.Element.SetProperty("xsr-24", value)
}

// xsr-25 (bool)
//
// Filter solo Right 25

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr25(value bool) error {
	return e.Element.SetProperty("xsr-25", value)
}

// xsr-26 (bool)
//
// Filter solo Right 26

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr26(value bool) error {
	return e.Element.SetProperty("xsr-26", value)
}

// xsr-27 (bool)
//
// Filter solo Right 27

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr27(value bool) error {
	return e.Element.SetProperty("xsr-27", value)
}

// xsr-28 (bool)
//
// Filter solo Right 28

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr28(value bool) error {
	return e.Element.SetProperty("xsr-28", value)
}

// xsr-29 (bool)
//
// Filter solo Right 29

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr29(value bool) error {
	return e.Element.SetProperty("xsr-29", value)
}

// xsr-3 (bool)
//
// Filter solo Right 3

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr3(value bool) error {
	return e.Element.SetProperty("xsr-3", value)
}

// xsr-30 (bool)
//
// Filter solo Right 30

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr30(value bool) error {
	return e.Element.SetProperty("xsr-30", value)
}

// xsr-31 (bool)
//
// Filter solo Right 31

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xsr-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr31(value bool) error {
	return e.Element.SetProperty("xsr-31", value)
}

// xsr-4 (bool)
//
// Filter solo Right 4

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr4(value bool) error {
	return e.Element.SetProperty("xsr-4", value)
}

// xsr-5 (bool)
//
// Filter solo Right 5

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr5(value bool) error {
	return e.Element.SetProperty("xsr-5", value)
}

// xsr-6 (bool)
//
// Filter solo Right 6

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr6(value bool) error {
	return e.Element.SetProperty("xsr-6", value)
}

// xsr-7 (bool)
//
// Filter solo Right 7

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr7(value bool) error {
	return e.Element.SetProperty("xsr-7", value)
}

// xsr-8 (bool)
//
// Filter solo Right 8

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr8(value bool) error {
	return e.Element.SetProperty("xsr-8", value)
}

// xsr-9 (bool)
//
// Filter solo Right 9

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetXsr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetXsr9(value bool) error {
	return e.Element.SetProperty("xsr-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2ParaEqualizerX32Lr) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl19Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl19 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr8Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr18X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr18 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr18X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr18 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr18X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr18 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr18X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr18 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr9Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl19X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl19 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl19X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl19 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl19X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl19 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl19X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl19 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr13X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr13X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr13X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr13X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml10ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml4ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr18 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl26Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl26 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr5Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr3X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr3X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr3X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr3X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr10X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr10X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr10X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr10X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr19X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr19 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr19X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr19 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr19X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr19 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr19X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr19 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr2X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr2X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr2X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr2X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml0ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr21 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr12Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl1X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl1X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl1X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl1X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl13X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl13 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl13X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl13 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl13X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl13 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl13X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl13 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr3Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl9X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl9X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl9X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl9X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr11X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr11X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr11X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr11X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml16ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml16 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml8ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl4Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr19Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr19 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr24Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr24 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr16Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr16 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr7Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr21X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr21 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr21X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr21 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr21X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr21 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr21X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr21 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml21ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml21 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml27ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml27 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr30 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl31Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl31 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr6X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr6X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr6X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr6X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml6ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl2X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl2 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl2X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl2 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl2X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl2 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl2X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl2 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl21X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl21 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl21X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl21 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl21X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl21 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl21X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl21 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl6X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl6 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl6X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl6 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl6X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl6 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl6X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl6 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl26X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl26 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl26X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl26 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl26X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl26 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl26X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl26 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl3X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl3 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl3X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl3 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl3X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl3 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl3X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl3 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml28ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml28 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml29ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml29 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl24Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl24 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr29Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr29 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl16X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl16 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl16X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl16 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl16X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl16 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl16X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl16 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl30X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl30 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl30X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl30 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl30X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl30 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl30X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl30 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl8X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl8X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl8X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl8X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml15ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml31ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml31 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr4 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr6 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl28X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl28 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl28X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl28 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl28X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl28 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl28X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl28 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr26Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr26 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrmode string

const (
	LspPlugInPluginsLv2ParaEqualizerX32LrmodeIir LspPlugInPluginsLv2ParaEqualizerX32Lrmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2ParaEqualizerX32LrmodeFir LspPlugInPluginsLv2ParaEqualizerX32Lrmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2ParaEqualizerX32LrmodeFft LspPlugInPluginsLv2ParaEqualizerX32Lrmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr7X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr7X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr7X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr7X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl21Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl21 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl22Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl22 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl30Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl30 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl7Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl7 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr18Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr18 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr25X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr25 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr25X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr25 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr25X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr25 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr25X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr25 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl1Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl16Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl16 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr20Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr20 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr0X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr0X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr0X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr0X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr16X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr16 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr16X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr16 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr16X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr16 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr16X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr16 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr17X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr17 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr17X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr17 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr17X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr17 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr17X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr17 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml2ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml2 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml30ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml30 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl6Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr28Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr28 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr4Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr4 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml23ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml23 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr31 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl27Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl27 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr11Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr6Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr6 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml26ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml26 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr10 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl25X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl25 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl25X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl25 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl25X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl25 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl25X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl25 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr4X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr4X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr4X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr4X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl15Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl25Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl25 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl8Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl8 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml14ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml20ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml20 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr14 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr1Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr1 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr13Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr14X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr14X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr14X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr14X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr24 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr26 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr27 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl28Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl28 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl20Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl20 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl9Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl9 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml18ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml18 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr15 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr23 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl11Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl11 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl13Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl13 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr8X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr8 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr8X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr8 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr8X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr8 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr8X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr8 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml7ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml7 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfsel string

const (
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersLeft07 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Left 0-7" // Filters Left 0-7 (0) – Filters Left 0-7
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersRight07 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Right 0-7" // Filters Right 0-7 (1) – Filters Right 0-7
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersLeft815 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Left 8-15" // Filters Left 8-15 (2) – Filters Left 8-15
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersRight815 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Right 8-15" // Filters Right 8-15 (3) – Filters Right 8-15
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersLeft1623 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Left 16-23" // Filters Left 16-23 (4) – Filters Left 16-23
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersRight1623 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Right 16-23" // Filters Right 16-23 (5) – Filters Right 16-23
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersLeft2431 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Left 24-31" // Filters Left 24-31 (6) – Filters Left 24-31
	LspPlugInPluginsLv2ParaEqualizerX32LrfselFiltersRight2431 LspPlugInPluginsLv2ParaEqualizerX32Lrfsel = "Filters Right 24-31" // Filters Right 24-31 (7) – Filters Right 24-31
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr0Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl29X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl29 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl29X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl29 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl29X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl29 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl29X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl29 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr15X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr15X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr15X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr15X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr8 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr21Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr21 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr26 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr26X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr26 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr26X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr26 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr26X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr26 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr26X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr26 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl11X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl11 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl11X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl11 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl11X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl11 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl11X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl11 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl27X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl27 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl27X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl27 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl27X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl27 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl27X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl27 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr30X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr30 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr30X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr30 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr30X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr30 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr30X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr30 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr19 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr22 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr17Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr17 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr23Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr23 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr22X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr22 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr22X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr22 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr22X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr22 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr22X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr22 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml9ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml9 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr27Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr27 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl17X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl17 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl17X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl17 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl17X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl17 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl17X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl17 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl24X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl24 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl24X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl24 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl24X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl24 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl24X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl24 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl7 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl7X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl7 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl7X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl7 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl7X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl7 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl7X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl7 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr12X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr12X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr12X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr12X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml1ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml1 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr16 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl18Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl18 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr14Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl31X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl31 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl31X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl31 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl31X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl31 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl31X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl31 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml13ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml13 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr17 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl14X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl14 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl14X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl14 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl14X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl14 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl14X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl14 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr31X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr31 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr31X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr31 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr31X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr31 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr31X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr31 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml3ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml3 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl10Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl14Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl14 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl29Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl29 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl3Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl3 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl12X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl12 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl12X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl12 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl12X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl12 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl12X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl12 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr1 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr1X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr1 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr1X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr1 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr1X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr1 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr1X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr1 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr24X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr24 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr24X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr24 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr24X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr24 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr24X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr24 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr28 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl0Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl0 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl12Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl12 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr25Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr25 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl0X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl0 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl0X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl0 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl0X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl0 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl0X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl0 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml11ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml11 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml19ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml19 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml25ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml25 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl22X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl22 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl22X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl22 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl22X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl22 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl22X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl22 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr20X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr20 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr20X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr20 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr20X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr20 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr20X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr20 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr20 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl23Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl23 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr30Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr30 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl10X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl10 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl10X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl10 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl10X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl10 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl10X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl10 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl4 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl4X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl4 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl4X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl4 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl4X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl4 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl4X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl4 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl15X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl15 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl15X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl15 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl15X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl15 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl15X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl15 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr5X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr5X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr5X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr5X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml24ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml24 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr0 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl5Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl5 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr15Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr15 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr22Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr22 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr28 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr28X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr28 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr28X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr28 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr28X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr28 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr28X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr28 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml22ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml22 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl17Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl17 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl18 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl18X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl18 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl18X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl18 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl18X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl18 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl18X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl18 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl23X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl23 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl23X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl23 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl23X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl23 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl23X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl23 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl5X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl5 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl5X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl5 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl5X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl5 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl5X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl5 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr10Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr10 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsl20 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl20X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsl20 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl20X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsl20 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl20X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsl20 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsl20X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsl20 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr29X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr29 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr29X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr29 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr29X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr29 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr29X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr29 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfft string

const (
	LspPlugInPluginsLv2ParaEqualizerX32LrfftOff LspPlugInPluginsLv2ParaEqualizerX32Lrfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32LrfftPostEq LspPlugInPluginsLv2ParaEqualizerX32Lrfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2ParaEqualizerX32LrfftPreEq LspPlugInPluginsLv2ParaEqualizerX32Lrfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml12ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml12 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml5ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml5 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr29 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2Off LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftl2Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftl2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr9 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr9X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr9 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr9X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr9 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr9X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr9 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr9X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr9 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr27 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr27X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr27 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr27X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr27 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr27X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr27 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr27X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr27 = "x4" // x4 (3) – x4
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfml17ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfml17 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25RlcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "RLC (BT)" // RLC (BT) (0) – RLC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25RlcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "RLC (MT)" // RLC (MT) (1) – RLC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25BwcBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "BWC (BT)" // BWC (BT) (2) – BWC (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25BwcMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "BWC (MT)" // BWC (MT) (3) – BWC (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25LrxBt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "LRX (BT)" // LRX (BT) (4) – LRX (BT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25LrxMt LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "LRX (MT)" // LRX (MT) (5) – LRX (MT)
	LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25ApoDr LspPlugInPluginsLv2ParaEqualizerX32Lrfmr25 = "APO (DR)" // APO (DR) (6) – APO (DR)
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr2Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr2 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31Off LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31Bell LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Bell" // Bell (1) – Bell
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31HiPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Hi-pass" // Hi-pass (2) – Hi-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31HiShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Hi-shelf" // Hi-shelf (3) – Hi-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31LoPass LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Lo-pass" // Lo-pass (4) – Lo-pass
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31LoShelf LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Lo-shelf" // Lo-shelf (5) – Lo-shelf
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31Notch LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Notch" // Notch (6) – Notch
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31Resonance LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Resonance" // Resonance (7) – Resonance
	LspPlugInPluginsLv2ParaEqualizerX32Lrftr31Allpass LspPlugInPluginsLv2ParaEqualizerX32Lrftr31 = "Allpass" // Allpass (8) – Allpass
)

type LspPlugInPluginsLv2ParaEqualizerX32Lrsr23 string

const (
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr23X1 LspPlugInPluginsLv2ParaEqualizerX32Lrsr23 = "x1" // x1 (0) – x1
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr23X2 LspPlugInPluginsLv2ParaEqualizerX32Lrsr23 = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr23X3 LspPlugInPluginsLv2ParaEqualizerX32Lrsr23 = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2ParaEqualizerX32Lrsr23X4 LspPlugInPluginsLv2ParaEqualizerX32Lrsr23 = "x4" // x4 (3) – x4
)

