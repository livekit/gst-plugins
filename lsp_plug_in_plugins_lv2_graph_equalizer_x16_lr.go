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

type LspPlugInPluginsLv2GraphEqualizerX16Lr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GraphEqualizerX16Lr() (*LspPlugInPluginsLv2GraphEqualizerX16Lr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-graph-equalizer-x16-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX16Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GraphEqualizerX16LrWithName(name string) (*LspPlugInPluginsLv2GraphEqualizerX16Lr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-graph-equalizer-x16-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX16Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2GraphEqualizerX16Lrfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFftvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFftvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fsel (GstLspPlugInPluginsLv2GraphEqualizerX16Lrfsel)
//
// Band select

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// fvl-0 (bool)
//
// Filter visibility  Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl0() (bool, error) {
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
// Filter visibility  Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl1() (bool, error) {
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
// Filter visibility  Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl10() (bool, error) {
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
// Filter visibility  Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl11() (bool, error) {
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
// Filter visibility  Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl12() (bool, error) {
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
// Filter visibility  Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl13() (bool, error) {
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
// Filter visibility  Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl14() (bool, error) {
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
// Filter visibility  Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl15() (bool, error) {
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
// Filter visibility  Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl2() (bool, error) {
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
// Filter visibility  Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl3() (bool, error) {
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
// Filter visibility  Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl4() (bool, error) {
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
// Filter visibility  Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl5() (bool, error) {
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
// Filter visibility  Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl6() (bool, error) {
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
// Filter visibility  Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl7() (bool, error) {
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
// Filter visibility  Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl8() (bool, error) {
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
// Filter visibility  Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvl9() (bool, error) {
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
// Filter visibility  Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr0() (bool, error) {
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
// Filter visibility  Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr1() (bool, error) {
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
// Filter visibility  Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr10() (bool, error) {
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
// Filter visibility  Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr11() (bool, error) {
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
// Filter visibility  Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr12() (bool, error) {
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
// Filter visibility  Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr13() (bool, error) {
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
// Filter visibility  Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr14() (bool, error) {
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
// Filter visibility  Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr15() (bool, error) {
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
// Filter visibility  Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr2() (bool, error) {
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
// Filter visibility  Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr3() (bool, error) {
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
// Filter visibility  Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr4() (bool, error) {
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
// Filter visibility  Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr5() (bool, error) {
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
// Filter visibility  Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr6() (bool, error) {
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
// Filter visibility  Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr7() (bool, error) {
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
// Filter visibility  Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr8() (bool, error) {
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
// Filter visibility  Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetFvr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl-0 (float32)
//
// Band gain Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl0(value float32) error {
	return e.Element.SetProperty("gl-0", value)
}

// gl-1 (float32)
//
// Band gain Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl1(value float32) error {
	return e.Element.SetProperty("gl-1", value)
}

// gl-10 (float32)
//
// Band gain Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl10(value float32) error {
	return e.Element.SetProperty("gl-10", value)
}

// gl-11 (float32)
//
// Band gain Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl11(value float32) error {
	return e.Element.SetProperty("gl-11", value)
}

// gl-12 (float32)
//
// Band gain Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl12(value float32) error {
	return e.Element.SetProperty("gl-12", value)
}

// gl-13 (float32)
//
// Band gain Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl13(value float32) error {
	return e.Element.SetProperty("gl-13", value)
}

// gl-14 (float32)
//
// Band gain Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl14(value float32) error {
	return e.Element.SetProperty("gl-14", value)
}

// gl-15 (float32)
//
// Band gain Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl15(value float32) error {
	return e.Element.SetProperty("gl-15", value)
}

// gl-2 (float32)
//
// Band gain Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl2(value float32) error {
	return e.Element.SetProperty("gl-2", value)
}

// gl-3 (float32)
//
// Band gain Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl3(value float32) error {
	return e.Element.SetProperty("gl-3", value)
}

// gl-4 (float32)
//
// Band gain Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl4(value float32) error {
	return e.Element.SetProperty("gl-4", value)
}

// gl-5 (float32)
//
// Band gain Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl5(value float32) error {
	return e.Element.SetProperty("gl-5", value)
}

// gl-6 (float32)
//
// Band gain Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl6(value float32) error {
	return e.Element.SetProperty("gl-6", value)
}

// gl-7 (float32)
//
// Band gain Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl7(value float32) error {
	return e.Element.SetProperty("gl-7", value)
}

// gl-8 (float32)
//
// Band gain Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl8(value float32) error {
	return e.Element.SetProperty("gl-8", value)
}

// gl-9 (float32)
//
// Band gain Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGl9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGl9(value float32) error {
	return e.Element.SetProperty("gl-9", value)
}

// gr-0 (float32)
//
// Band gain Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr0(value float32) error {
	return e.Element.SetProperty("gr-0", value)
}

// gr-1 (float32)
//
// Band gain Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr1(value float32) error {
	return e.Element.SetProperty("gr-1", value)
}

// gr-10 (float32)
//
// Band gain Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr10(value float32) error {
	return e.Element.SetProperty("gr-10", value)
}

// gr-11 (float32)
//
// Band gain Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr11(value float32) error {
	return e.Element.SetProperty("gr-11", value)
}

// gr-12 (float32)
//
// Band gain Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr12(value float32) error {
	return e.Element.SetProperty("gr-12", value)
}

// gr-13 (float32)
//
// Band gain Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr13(value float32) error {
	return e.Element.SetProperty("gr-13", value)
}

// gr-14 (float32)
//
// Band gain Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr14(value float32) error {
	return e.Element.SetProperty("gr-14", value)
}

// gr-15 (float32)
//
// Band gain Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr15(value float32) error {
	return e.Element.SetProperty("gr-15", value)
}

// gr-2 (float32)
//
// Band gain Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr2(value float32) error {
	return e.Element.SetProperty("gr-2", value)
}

// gr-3 (float32)
//
// Band gain Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr3(value float32) error {
	return e.Element.SetProperty("gr-3", value)
}

// gr-4 (float32)
//
// Band gain Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr4(value float32) error {
	return e.Element.SetProperty("gr-4", value)
}

// gr-5 (float32)
//
// Band gain Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr5(value float32) error {
	return e.Element.SetProperty("gr-5", value)
}

// gr-6 (float32)
//
// Band gain Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr6(value float32) error {
	return e.Element.SetProperty("gr-6", value)
}

// gr-7 (float32)
//
// Band gain Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr7(value float32) error {
	return e.Element.SetProperty("gr-7", value)
}

// gr-8 (float32)
//
// Band gain Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr8(value float32) error {
	return e.Element.SetProperty("gr-8", value)
}

// gr-9 (float32)
//
// Band gain Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetGr9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetGr9(value float32) error {
	return e.Element.SetProperty("gr-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetImr() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2GraphEqualizerX16Lrmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetOutLatency() (int, error) {
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

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// slope (GstLspPlugInPluginsLv2GraphEqualizerX16Lrslope)
//
// Filter slope

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetSlope() (interface{}, error) {
	return e.Element.GetProperty("slope")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetSlope(value interface{}) error {
	return e.Element.SetProperty("slope", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetSmr() (float32, error) {
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

// xel-0 (bool)
//
// Band on Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel0(value bool) error {
	return e.Element.SetProperty("xel-0", value)
}

// xel-1 (bool)
//
// Band on Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel1(value bool) error {
	return e.Element.SetProperty("xel-1", value)
}

// xel-10 (bool)
//
// Band on Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel10(value bool) error {
	return e.Element.SetProperty("xel-10", value)
}

// xel-11 (bool)
//
// Band on Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel11(value bool) error {
	return e.Element.SetProperty("xel-11", value)
}

// xel-12 (bool)
//
// Band on Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel12(value bool) error {
	return e.Element.SetProperty("xel-12", value)
}

// xel-13 (bool)
//
// Band on Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel13(value bool) error {
	return e.Element.SetProperty("xel-13", value)
}

// xel-14 (bool)
//
// Band on Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel14(value bool) error {
	return e.Element.SetProperty("xel-14", value)
}

// xel-15 (bool)
//
// Band on Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel15(value bool) error {
	return e.Element.SetProperty("xel-15", value)
}

// xel-2 (bool)
//
// Band on Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel2(value bool) error {
	return e.Element.SetProperty("xel-2", value)
}

// xel-3 (bool)
//
// Band on Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel3(value bool) error {
	return e.Element.SetProperty("xel-3", value)
}

// xel-4 (bool)
//
// Band on Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel4(value bool) error {
	return e.Element.SetProperty("xel-4", value)
}

// xel-5 (bool)
//
// Band on Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel5(value bool) error {
	return e.Element.SetProperty("xel-5", value)
}

// xel-6 (bool)
//
// Band on Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel6(value bool) error {
	return e.Element.SetProperty("xel-6", value)
}

// xel-7 (bool)
//
// Band on Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel7(value bool) error {
	return e.Element.SetProperty("xel-7", value)
}

// xel-8 (bool)
//
// Band on Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel8(value bool) error {
	return e.Element.SetProperty("xel-8", value)
}

// xel-9 (bool)
//
// Band on Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXel9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXel9(value bool) error {
	return e.Element.SetProperty("xel-9", value)
}

// xer-0 (bool)
//
// Band on Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer0(value bool) error {
	return e.Element.SetProperty("xer-0", value)
}

// xer-1 (bool)
//
// Band on Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer1(value bool) error {
	return e.Element.SetProperty("xer-1", value)
}

// xer-10 (bool)
//
// Band on Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer10(value bool) error {
	return e.Element.SetProperty("xer-10", value)
}

// xer-11 (bool)
//
// Band on Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer11(value bool) error {
	return e.Element.SetProperty("xer-11", value)
}

// xer-12 (bool)
//
// Band on Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer12(value bool) error {
	return e.Element.SetProperty("xer-12", value)
}

// xer-13 (bool)
//
// Band on Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer13(value bool) error {
	return e.Element.SetProperty("xer-13", value)
}

// xer-14 (bool)
//
// Band on Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer14(value bool) error {
	return e.Element.SetProperty("xer-14", value)
}

// xer-15 (bool)
//
// Band on Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer15(value bool) error {
	return e.Element.SetProperty("xer-15", value)
}

// xer-2 (bool)
//
// Band on Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer2(value bool) error {
	return e.Element.SetProperty("xer-2", value)
}

// xer-3 (bool)
//
// Band on Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer3(value bool) error {
	return e.Element.SetProperty("xer-3", value)
}

// xer-4 (bool)
//
// Band on Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer4(value bool) error {
	return e.Element.SetProperty("xer-4", value)
}

// xer-5 (bool)
//
// Band on Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer5(value bool) error {
	return e.Element.SetProperty("xer-5", value)
}

// xer-6 (bool)
//
// Band on Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer6(value bool) error {
	return e.Element.SetProperty("xer-6", value)
}

// xer-7 (bool)
//
// Band on Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer7(value bool) error {
	return e.Element.SetProperty("xer-7", value)
}

// xer-8 (bool)
//
// Band on Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer8(value bool) error {
	return e.Element.SetProperty("xer-8", value)
}

// xer-9 (bool)
//
// Band on Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXer9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXer9(value bool) error {
	return e.Element.SetProperty("xer-9", value)
}

// xml-0 (bool)
//
// Band mute Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml0(value bool) error {
	return e.Element.SetProperty("xml-0", value)
}

// xml-1 (bool)
//
// Band mute Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml1(value bool) error {
	return e.Element.SetProperty("xml-1", value)
}

// xml-10 (bool)
//
// Band mute Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml10(value bool) error {
	return e.Element.SetProperty("xml-10", value)
}

// xml-11 (bool)
//
// Band mute Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml11(value bool) error {
	return e.Element.SetProperty("xml-11", value)
}

// xml-12 (bool)
//
// Band mute Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml12(value bool) error {
	return e.Element.SetProperty("xml-12", value)
}

// xml-13 (bool)
//
// Band mute Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml13(value bool) error {
	return e.Element.SetProperty("xml-13", value)
}

// xml-14 (bool)
//
// Band mute Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml14(value bool) error {
	return e.Element.SetProperty("xml-14", value)
}

// xml-15 (bool)
//
// Band mute Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml15(value bool) error {
	return e.Element.SetProperty("xml-15", value)
}

// xml-2 (bool)
//
// Band mute Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml2(value bool) error {
	return e.Element.SetProperty("xml-2", value)
}

// xml-3 (bool)
//
// Band mute Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml3(value bool) error {
	return e.Element.SetProperty("xml-3", value)
}

// xml-4 (bool)
//
// Band mute Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml4(value bool) error {
	return e.Element.SetProperty("xml-4", value)
}

// xml-5 (bool)
//
// Band mute Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml5(value bool) error {
	return e.Element.SetProperty("xml-5", value)
}

// xml-6 (bool)
//
// Band mute Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml6(value bool) error {
	return e.Element.SetProperty("xml-6", value)
}

// xml-7 (bool)
//
// Band mute Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml7(value bool) error {
	return e.Element.SetProperty("xml-7", value)
}

// xml-8 (bool)
//
// Band mute Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml8(value bool) error {
	return e.Element.SetProperty("xml-8", value)
}

// xml-9 (bool)
//
// Band mute Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXml9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXml9(value bool) error {
	return e.Element.SetProperty("xml-9", value)
}

// xmr-0 (bool)
//
// Band mute Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr0(value bool) error {
	return e.Element.SetProperty("xmr-0", value)
}

// xmr-1 (bool)
//
// Band mute Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr1(value bool) error {
	return e.Element.SetProperty("xmr-1", value)
}

// xmr-10 (bool)
//
// Band mute Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr10(value bool) error {
	return e.Element.SetProperty("xmr-10", value)
}

// xmr-11 (bool)
//
// Band mute Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr11(value bool) error {
	return e.Element.SetProperty("xmr-11", value)
}

// xmr-12 (bool)
//
// Band mute Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr12(value bool) error {
	return e.Element.SetProperty("xmr-12", value)
}

// xmr-13 (bool)
//
// Band mute Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr13(value bool) error {
	return e.Element.SetProperty("xmr-13", value)
}

// xmr-14 (bool)
//
// Band mute Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr14(value bool) error {
	return e.Element.SetProperty("xmr-14", value)
}

// xmr-15 (bool)
//
// Band mute Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr15(value bool) error {
	return e.Element.SetProperty("xmr-15", value)
}

// xmr-2 (bool)
//
// Band mute Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr2(value bool) error {
	return e.Element.SetProperty("xmr-2", value)
}

// xmr-3 (bool)
//
// Band mute Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr3(value bool) error {
	return e.Element.SetProperty("xmr-3", value)
}

// xmr-4 (bool)
//
// Band mute Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr4(value bool) error {
	return e.Element.SetProperty("xmr-4", value)
}

// xmr-5 (bool)
//
// Band mute Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr5(value bool) error {
	return e.Element.SetProperty("xmr-5", value)
}

// xmr-6 (bool)
//
// Band mute Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr6(value bool) error {
	return e.Element.SetProperty("xmr-6", value)
}

// xmr-7 (bool)
//
// Band mute Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr7(value bool) error {
	return e.Element.SetProperty("xmr-7", value)
}

// xmr-8 (bool)
//
// Band mute Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr8(value bool) error {
	return e.Element.SetProperty("xmr-8", value)
}

// xmr-9 (bool)
//
// Band mute Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXmr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXmr9(value bool) error {
	return e.Element.SetProperty("xmr-9", value)
}

// xsl-0 (bool)
//
// Band solo Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl0(value bool) error {
	return e.Element.SetProperty("xsl-0", value)
}

// xsl-1 (bool)
//
// Band solo Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl1(value bool) error {
	return e.Element.SetProperty("xsl-1", value)
}

// xsl-10 (bool)
//
// Band solo Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl10(value bool) error {
	return e.Element.SetProperty("xsl-10", value)
}

// xsl-11 (bool)
//
// Band solo Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl11(value bool) error {
	return e.Element.SetProperty("xsl-11", value)
}

// xsl-12 (bool)
//
// Band solo Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl12(value bool) error {
	return e.Element.SetProperty("xsl-12", value)
}

// xsl-13 (bool)
//
// Band solo Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl13(value bool) error {
	return e.Element.SetProperty("xsl-13", value)
}

// xsl-14 (bool)
//
// Band solo Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl14(value bool) error {
	return e.Element.SetProperty("xsl-14", value)
}

// xsl-15 (bool)
//
// Band solo Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl15(value bool) error {
	return e.Element.SetProperty("xsl-15", value)
}

// xsl-2 (bool)
//
// Band solo Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl2(value bool) error {
	return e.Element.SetProperty("xsl-2", value)
}

// xsl-3 (bool)
//
// Band solo Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl3(value bool) error {
	return e.Element.SetProperty("xsl-3", value)
}

// xsl-4 (bool)
//
// Band solo Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl4(value bool) error {
	return e.Element.SetProperty("xsl-4", value)
}

// xsl-5 (bool)
//
// Band solo Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl5(value bool) error {
	return e.Element.SetProperty("xsl-5", value)
}

// xsl-6 (bool)
//
// Band solo Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl6(value bool) error {
	return e.Element.SetProperty("xsl-6", value)
}

// xsl-7 (bool)
//
// Band solo Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl7(value bool) error {
	return e.Element.SetProperty("xsl-7", value)
}

// xsl-8 (bool)
//
// Band solo Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl8(value bool) error {
	return e.Element.SetProperty("xsl-8", value)
}

// xsl-9 (bool)
//
// Band solo Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsl9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsl9(value bool) error {
	return e.Element.SetProperty("xsl-9", value)
}

// xsr-0 (bool)
//
// Band solo Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr0(value bool) error {
	return e.Element.SetProperty("xsr-0", value)
}

// xsr-1 (bool)
//
// Band solo Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr1(value bool) error {
	return e.Element.SetProperty("xsr-1", value)
}

// xsr-10 (bool)
//
// Band solo Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr10(value bool) error {
	return e.Element.SetProperty("xsr-10", value)
}

// xsr-11 (bool)
//
// Band solo Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr11(value bool) error {
	return e.Element.SetProperty("xsr-11", value)
}

// xsr-12 (bool)
//
// Band solo Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr12(value bool) error {
	return e.Element.SetProperty("xsr-12", value)
}

// xsr-13 (bool)
//
// Band solo Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr13(value bool) error {
	return e.Element.SetProperty("xsr-13", value)
}

// xsr-14 (bool)
//
// Band solo Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr14(value bool) error {
	return e.Element.SetProperty("xsr-14", value)
}

// xsr-15 (bool)
//
// Band solo Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr15(value bool) error {
	return e.Element.SetProperty("xsr-15", value)
}

// xsr-2 (bool)
//
// Band solo Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr2(value bool) error {
	return e.Element.SetProperty("xsr-2", value)
}

// xsr-3 (bool)
//
// Band solo Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr3(value bool) error {
	return e.Element.SetProperty("xsr-3", value)
}

// xsr-4 (bool)
//
// Band solo Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr4(value bool) error {
	return e.Element.SetProperty("xsr-4", value)
}

// xsr-5 (bool)
//
// Band solo Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr5(value bool) error {
	return e.Element.SetProperty("xsr-5", value)
}

// xsr-6 (bool)
//
// Band solo Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr6(value bool) error {
	return e.Element.SetProperty("xsr-6", value)
}

// xsr-7 (bool)
//
// Band solo Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr7(value bool) error {
	return e.Element.SetProperty("xsr-7", value)
}

// xsr-8 (bool)
//
// Band solo Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr8(value bool) error {
	return e.Element.SetProperty("xsr-8", value)
}

// xsr-9 (bool)
//
// Band solo Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetXsr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetXsr9(value bool) error {
	return e.Element.SetProperty("xsr-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Lr) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GraphEqualizerX16Lrfft string

const (
	LspPlugInPluginsLv2GraphEqualizerX16LrfftOff LspPlugInPluginsLv2GraphEqualizerX16Lrfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2GraphEqualizerX16LrfftPostEq LspPlugInPluginsLv2GraphEqualizerX16Lrfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2GraphEqualizerX16LrfftPreEq LspPlugInPluginsLv2GraphEqualizerX16Lrfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2GraphEqualizerX16Lrfsel string

const (
	LspPlugInPluginsLv2GraphEqualizerX16LrfselBandsLeft LspPlugInPluginsLv2GraphEqualizerX16Lrfsel = "Bands Left" // Bands Left (0) – Bands Left
	LspPlugInPluginsLv2GraphEqualizerX16LrfselBandsRight LspPlugInPluginsLv2GraphEqualizerX16Lrfsel = "Bands Right" // Bands Right (1) – Bands Right
)

type LspPlugInPluginsLv2GraphEqualizerX16Lrmode string

const (
	LspPlugInPluginsLv2GraphEqualizerX16LrmodeIir LspPlugInPluginsLv2GraphEqualizerX16Lrmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2GraphEqualizerX16LrmodeFir LspPlugInPluginsLv2GraphEqualizerX16Lrmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2GraphEqualizerX16LrmodeFft LspPlugInPluginsLv2GraphEqualizerX16Lrmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2GraphEqualizerX16Lrslope string

const (
	LspPlugInPluginsLv2GraphEqualizerX16LrslopeBt48 LspPlugInPluginsLv2GraphEqualizerX16Lrslope = "BT48" // BT48 (0) – BT48
	LspPlugInPluginsLv2GraphEqualizerX16LrslopeMt48 LspPlugInPluginsLv2GraphEqualizerX16Lrslope = "MT48" // MT48 (1) – MT48
	LspPlugInPluginsLv2GraphEqualizerX16LrslopeBt72 LspPlugInPluginsLv2GraphEqualizerX16Lrslope = "BT72" // BT72 (2) – BT72
	LspPlugInPluginsLv2GraphEqualizerX16LrslopeMt72 LspPlugInPluginsLv2GraphEqualizerX16Lrslope = "MT72" // MT72 (3) – MT72
	LspPlugInPluginsLv2GraphEqualizerX16LrslopeBt96 LspPlugInPluginsLv2GraphEqualizerX16Lrslope = "BT96" // BT96 (4) – BT96
	LspPlugInPluginsLv2GraphEqualizerX16LrslopeMt96 LspPlugInPluginsLv2GraphEqualizerX16Lrslope = "MT96" // MT96 (5) – MT96
)

