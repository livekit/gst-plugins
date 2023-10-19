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

type LspPlugInPluginsLv2GraphEqualizerX32Lr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GraphEqualizerX32Lr() (*LspPlugInPluginsLv2GraphEqualizerX32Lr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-graph-equalizer-x32-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX32Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GraphEqualizerX32LrWithName(name string) (*LspPlugInPluginsLv2GraphEqualizerX32Lr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-graph-equalizer-x32-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX32Lr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2GraphEqualizerX32Lrfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFftvL() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFftvR() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fsel (GstLspPlugInPluginsLv2GraphEqualizerX32Lrfsel)
//
// Band select

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// fvl-0 (bool)
//
// Filter visibility  Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl0() (bool, error) {
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
// Filter visibility  Left 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl1() (bool, error) {
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
// Filter visibility  Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl10() (bool, error) {
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
// Filter visibility  Left 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl11() (bool, error) {
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
// Filter visibility  Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl12() (bool, error) {
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
// Filter visibility  Left 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl13() (bool, error) {
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
// Filter visibility  Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl14() (bool, error) {
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
// Filter visibility  Left 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl15() (bool, error) {
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
// Filter visibility  Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl16() (bool, error) {
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
// Filter visibility  Left 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl17() (bool, error) {
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
// Filter visibility  Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl18() (bool, error) {
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
// Filter visibility  Left 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl19() (bool, error) {
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
// Filter visibility  Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl2() (bool, error) {
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
// Filter visibility  Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl20() (bool, error) {
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
// Filter visibility  Left 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl21() (bool, error) {
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
// Filter visibility  Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl22() (bool, error) {
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
// Filter visibility  Left 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl23() (bool, error) {
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
// Filter visibility  Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl24() (bool, error) {
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
// Filter visibility  Left 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl25() (bool, error) {
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
// Filter visibility  Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl26() (bool, error) {
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
// Filter visibility  Left 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl27() (bool, error) {
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
// Filter visibility  Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl28() (bool, error) {
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
// Filter visibility  Left 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl29() (bool, error) {
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
// Filter visibility  Left 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl3() (bool, error) {
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
// Filter visibility  Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl30() (bool, error) {
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
// Filter visibility  Left 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl31() (bool, error) {
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
// Filter visibility  Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl4() (bool, error) {
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
// Filter visibility  Left 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl5() (bool, error) {
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
// Filter visibility  Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl6() (bool, error) {
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
// Filter visibility  Left 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl7() (bool, error) {
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
// Filter visibility  Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl8() (bool, error) {
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
// Filter visibility  Left 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvl9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr0() (bool, error) {
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
// Filter visibility  Right 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr1() (bool, error) {
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
// Filter visibility  Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr10() (bool, error) {
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
// Filter visibility  Right 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr11() (bool, error) {
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
// Filter visibility  Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr12() (bool, error) {
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
// Filter visibility  Right 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr13() (bool, error) {
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
// Filter visibility  Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr14() (bool, error) {
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
// Filter visibility  Right 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr15() (bool, error) {
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
// Filter visibility  Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr16() (bool, error) {
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
// Filter visibility  Right 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr17() (bool, error) {
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
// Filter visibility  Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr18() (bool, error) {
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
// Filter visibility  Right 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr19() (bool, error) {
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
// Filter visibility  Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr2() (bool, error) {
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
// Filter visibility  Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr20() (bool, error) {
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
// Filter visibility  Right 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr21() (bool, error) {
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
// Filter visibility  Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr22() (bool, error) {
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
// Filter visibility  Right 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr23() (bool, error) {
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
// Filter visibility  Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr24() (bool, error) {
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
// Filter visibility  Right 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr25() (bool, error) {
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
// Filter visibility  Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr26() (bool, error) {
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
// Filter visibility  Right 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr27() (bool, error) {
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
// Filter visibility  Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr28() (bool, error) {
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
// Filter visibility  Right 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr29() (bool, error) {
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
// Filter visibility  Right 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr3() (bool, error) {
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
// Filter visibility  Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr30() (bool, error) {
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
// Filter visibility  Right 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr31() (bool, error) {
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
// Filter visibility  Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr4() (bool, error) {
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
// Filter visibility  Right 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr5() (bool, error) {
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
// Filter visibility  Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr6() (bool, error) {
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
// Filter visibility  Right 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr7() (bool, error) {
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
// Filter visibility  Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr8() (bool, error) {
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
// Filter visibility  Right 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetFvr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gl-0 (float32)
//
// Band gain Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl0(value float32) error {
	return e.Element.SetProperty("gl-0", value)
}

// gl-1 (float32)
//
// Band gain Left 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl1(value float32) error {
	return e.Element.SetProperty("gl-1", value)
}

// gl-10 (float32)
//
// Band gain Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl10(value float32) error {
	return e.Element.SetProperty("gl-10", value)
}

// gl-11 (float32)
//
// Band gain Left 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl11(value float32) error {
	return e.Element.SetProperty("gl-11", value)
}

// gl-12 (float32)
//
// Band gain Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl12(value float32) error {
	return e.Element.SetProperty("gl-12", value)
}

// gl-13 (float32)
//
// Band gain Left 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl13(value float32) error {
	return e.Element.SetProperty("gl-13", value)
}

// gl-14 (float32)
//
// Band gain Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl14(value float32) error {
	return e.Element.SetProperty("gl-14", value)
}

// gl-15 (float32)
//
// Band gain Left 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl15(value float32) error {
	return e.Element.SetProperty("gl-15", value)
}

// gl-16 (float32)
//
// Band gain Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl16() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl16(value float32) error {
	return e.Element.SetProperty("gl-16", value)
}

// gl-17 (float32)
//
// Band gain Left 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl17() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl17(value float32) error {
	return e.Element.SetProperty("gl-17", value)
}

// gl-18 (float32)
//
// Band gain Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl18() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl18(value float32) error {
	return e.Element.SetProperty("gl-18", value)
}

// gl-19 (float32)
//
// Band gain Left 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl19() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl19(value float32) error {
	return e.Element.SetProperty("gl-19", value)
}

// gl-2 (float32)
//
// Band gain Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl2(value float32) error {
	return e.Element.SetProperty("gl-2", value)
}

// gl-20 (float32)
//
// Band gain Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl20() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl20(value float32) error {
	return e.Element.SetProperty("gl-20", value)
}

// gl-21 (float32)
//
// Band gain Left 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl21() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl21(value float32) error {
	return e.Element.SetProperty("gl-21", value)
}

// gl-22 (float32)
//
// Band gain Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl22() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl22(value float32) error {
	return e.Element.SetProperty("gl-22", value)
}

// gl-23 (float32)
//
// Band gain Left 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl23() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl23(value float32) error {
	return e.Element.SetProperty("gl-23", value)
}

// gl-24 (float32)
//
// Band gain Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl24() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl24(value float32) error {
	return e.Element.SetProperty("gl-24", value)
}

// gl-25 (float32)
//
// Band gain Left 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl25() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl25(value float32) error {
	return e.Element.SetProperty("gl-25", value)
}

// gl-26 (float32)
//
// Band gain Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl26() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl26(value float32) error {
	return e.Element.SetProperty("gl-26", value)
}

// gl-27 (float32)
//
// Band gain Left 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl27() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl27(value float32) error {
	return e.Element.SetProperty("gl-27", value)
}

// gl-28 (float32)
//
// Band gain Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl28() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl28(value float32) error {
	return e.Element.SetProperty("gl-28", value)
}

// gl-29 (float32)
//
// Band gain Left 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl29() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl29(value float32) error {
	return e.Element.SetProperty("gl-29", value)
}

// gl-3 (float32)
//
// Band gain Left 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl3(value float32) error {
	return e.Element.SetProperty("gl-3", value)
}

// gl-30 (float32)
//
// Band gain Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl30() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl30(value float32) error {
	return e.Element.SetProperty("gl-30", value)
}

// gl-31 (float32)
//
// Band gain Left 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl31() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl31(value float32) error {
	return e.Element.SetProperty("gl-31", value)
}

// gl-4 (float32)
//
// Band gain Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl4(value float32) error {
	return e.Element.SetProperty("gl-4", value)
}

// gl-5 (float32)
//
// Band gain Left 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl5(value float32) error {
	return e.Element.SetProperty("gl-5", value)
}

// gl-6 (float32)
//
// Band gain Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl6(value float32) error {
	return e.Element.SetProperty("gl-6", value)
}

// gl-7 (float32)
//
// Band gain Left 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl7(value float32) error {
	return e.Element.SetProperty("gl-7", value)
}

// gl-8 (float32)
//
// Band gain Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl8(value float32) error {
	return e.Element.SetProperty("gl-8", value)
}

// gl-9 (float32)
//
// Band gain Left 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGl9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGl9(value float32) error {
	return e.Element.SetProperty("gl-9", value)
}

// gr-0 (float32)
//
// Band gain Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr0(value float32) error {
	return e.Element.SetProperty("gr-0", value)
}

// gr-1 (float32)
//
// Band gain Right 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr1(value float32) error {
	return e.Element.SetProperty("gr-1", value)
}

// gr-10 (float32)
//
// Band gain Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr10(value float32) error {
	return e.Element.SetProperty("gr-10", value)
}

// gr-11 (float32)
//
// Band gain Right 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr11(value float32) error {
	return e.Element.SetProperty("gr-11", value)
}

// gr-12 (float32)
//
// Band gain Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr12(value float32) error {
	return e.Element.SetProperty("gr-12", value)
}

// gr-13 (float32)
//
// Band gain Right 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr13(value float32) error {
	return e.Element.SetProperty("gr-13", value)
}

// gr-14 (float32)
//
// Band gain Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr14(value float32) error {
	return e.Element.SetProperty("gr-14", value)
}

// gr-15 (float32)
//
// Band gain Right 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr15(value float32) error {
	return e.Element.SetProperty("gr-15", value)
}

// gr-16 (float32)
//
// Band gain Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr16() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr16(value float32) error {
	return e.Element.SetProperty("gr-16", value)
}

// gr-17 (float32)
//
// Band gain Right 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr17() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr17(value float32) error {
	return e.Element.SetProperty("gr-17", value)
}

// gr-18 (float32)
//
// Band gain Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr18() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr18(value float32) error {
	return e.Element.SetProperty("gr-18", value)
}

// gr-19 (float32)
//
// Band gain Right 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr19() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr19(value float32) error {
	return e.Element.SetProperty("gr-19", value)
}

// gr-2 (float32)
//
// Band gain Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr2(value float32) error {
	return e.Element.SetProperty("gr-2", value)
}

// gr-20 (float32)
//
// Band gain Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr20() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr20(value float32) error {
	return e.Element.SetProperty("gr-20", value)
}

// gr-21 (float32)
//
// Band gain Right 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr21() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr21(value float32) error {
	return e.Element.SetProperty("gr-21", value)
}

// gr-22 (float32)
//
// Band gain Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr22() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr22(value float32) error {
	return e.Element.SetProperty("gr-22", value)
}

// gr-23 (float32)
//
// Band gain Right 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr23() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr23(value float32) error {
	return e.Element.SetProperty("gr-23", value)
}

// gr-24 (float32)
//
// Band gain Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr24() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr24(value float32) error {
	return e.Element.SetProperty("gr-24", value)
}

// gr-25 (float32)
//
// Band gain Right 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr25() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr25(value float32) error {
	return e.Element.SetProperty("gr-25", value)
}

// gr-26 (float32)
//
// Band gain Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr26() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr26(value float32) error {
	return e.Element.SetProperty("gr-26", value)
}

// gr-27 (float32)
//
// Band gain Right 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr27() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr27(value float32) error {
	return e.Element.SetProperty("gr-27", value)
}

// gr-28 (float32)
//
// Band gain Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr28() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr28(value float32) error {
	return e.Element.SetProperty("gr-28", value)
}

// gr-29 (float32)
//
// Band gain Right 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr29() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr29(value float32) error {
	return e.Element.SetProperty("gr-29", value)
}

// gr-3 (float32)
//
// Band gain Right 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr3(value float32) error {
	return e.Element.SetProperty("gr-3", value)
}

// gr-30 (float32)
//
// Band gain Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr30() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr30(value float32) error {
	return e.Element.SetProperty("gr-30", value)
}

// gr-31 (float32)
//
// Band gain Right 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr31() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr31(value float32) error {
	return e.Element.SetProperty("gr-31", value)
}

// gr-4 (float32)
//
// Band gain Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr4(value float32) error {
	return e.Element.SetProperty("gr-4", value)
}

// gr-5 (float32)
//
// Band gain Right 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr5(value float32) error {
	return e.Element.SetProperty("gr-5", value)
}

// gr-6 (float32)
//
// Band gain Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr6(value float32) error {
	return e.Element.SetProperty("gr-6", value)
}

// gr-7 (float32)
//
// Band gain Right 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr7(value float32) error {
	return e.Element.SetProperty("gr-7", value)
}

// gr-8 (float32)
//
// Band gain Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr8(value float32) error {
	return e.Element.SetProperty("gr-8", value)
}

// gr-9 (float32)
//
// Band gain Right 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetGr9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetGr9(value float32) error {
	return e.Element.SetProperty("gr-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetImr() (float32, error) {
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

// mode (GstLspPlugInPluginsLv2GraphEqualizerX32Lrmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// slope (GstLspPlugInPluginsLv2GraphEqualizerX32Lrslope)
//
// Filter slope

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetSlope() (interface{}, error) {
	return e.Element.GetProperty("slope")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetSlope(value interface{}) error {
	return e.Element.SetProperty("slope", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetSmr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel0(value bool) error {
	return e.Element.SetProperty("xel-0", value)
}

// xel-1 (bool)
//
// Band on Left 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel1(value bool) error {
	return e.Element.SetProperty("xel-1", value)
}

// xel-10 (bool)
//
// Band on Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel10(value bool) error {
	return e.Element.SetProperty("xel-10", value)
}

// xel-11 (bool)
//
// Band on Left 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel11(value bool) error {
	return e.Element.SetProperty("xel-11", value)
}

// xel-12 (bool)
//
// Band on Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel12(value bool) error {
	return e.Element.SetProperty("xel-12", value)
}

// xel-13 (bool)
//
// Band on Left 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel13(value bool) error {
	return e.Element.SetProperty("xel-13", value)
}

// xel-14 (bool)
//
// Band on Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel14(value bool) error {
	return e.Element.SetProperty("xel-14", value)
}

// xel-15 (bool)
//
// Band on Left 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel15(value bool) error {
	return e.Element.SetProperty("xel-15", value)
}

// xel-16 (bool)
//
// Band on Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel16(value bool) error {
	return e.Element.SetProperty("xel-16", value)
}

// xel-17 (bool)
//
// Band on Left 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel17(value bool) error {
	return e.Element.SetProperty("xel-17", value)
}

// xel-18 (bool)
//
// Band on Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel18(value bool) error {
	return e.Element.SetProperty("xel-18", value)
}

// xel-19 (bool)
//
// Band on Left 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel19(value bool) error {
	return e.Element.SetProperty("xel-19", value)
}

// xel-2 (bool)
//
// Band on Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel2(value bool) error {
	return e.Element.SetProperty("xel-2", value)
}

// xel-20 (bool)
//
// Band on Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel20(value bool) error {
	return e.Element.SetProperty("xel-20", value)
}

// xel-21 (bool)
//
// Band on Left 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel21(value bool) error {
	return e.Element.SetProperty("xel-21", value)
}

// xel-22 (bool)
//
// Band on Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel22(value bool) error {
	return e.Element.SetProperty("xel-22", value)
}

// xel-23 (bool)
//
// Band on Left 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel23(value bool) error {
	return e.Element.SetProperty("xel-23", value)
}

// xel-24 (bool)
//
// Band on Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel24(value bool) error {
	return e.Element.SetProperty("xel-24", value)
}

// xel-25 (bool)
//
// Band on Left 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel25(value bool) error {
	return e.Element.SetProperty("xel-25", value)
}

// xel-26 (bool)
//
// Band on Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel26(value bool) error {
	return e.Element.SetProperty("xel-26", value)
}

// xel-27 (bool)
//
// Band on Left 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel27(value bool) error {
	return e.Element.SetProperty("xel-27", value)
}

// xel-28 (bool)
//
// Band on Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel28(value bool) error {
	return e.Element.SetProperty("xel-28", value)
}

// xel-29 (bool)
//
// Band on Left 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel29(value bool) error {
	return e.Element.SetProperty("xel-29", value)
}

// xel-3 (bool)
//
// Band on Left 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel3(value bool) error {
	return e.Element.SetProperty("xel-3", value)
}

// xel-30 (bool)
//
// Band on Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel30(value bool) error {
	return e.Element.SetProperty("xel-30", value)
}

// xel-31 (bool)
//
// Band on Left 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xel-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel31(value bool) error {
	return e.Element.SetProperty("xel-31", value)
}

// xel-4 (bool)
//
// Band on Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel4(value bool) error {
	return e.Element.SetProperty("xel-4", value)
}

// xel-5 (bool)
//
// Band on Left 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel5(value bool) error {
	return e.Element.SetProperty("xel-5", value)
}

// xel-6 (bool)
//
// Band on Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel6(value bool) error {
	return e.Element.SetProperty("xel-6", value)
}

// xel-7 (bool)
//
// Band on Left 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel7(value bool) error {
	return e.Element.SetProperty("xel-7", value)
}

// xel-8 (bool)
//
// Band on Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel8(value bool) error {
	return e.Element.SetProperty("xel-8", value)
}

// xel-9 (bool)
//
// Band on Left 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXel9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXel9(value bool) error {
	return e.Element.SetProperty("xel-9", value)
}

// xer-0 (bool)
//
// Band on Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer0(value bool) error {
	return e.Element.SetProperty("xer-0", value)
}

// xer-1 (bool)
//
// Band on Right 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer1(value bool) error {
	return e.Element.SetProperty("xer-1", value)
}

// xer-10 (bool)
//
// Band on Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer10(value bool) error {
	return e.Element.SetProperty("xer-10", value)
}

// xer-11 (bool)
//
// Band on Right 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer11(value bool) error {
	return e.Element.SetProperty("xer-11", value)
}

// xer-12 (bool)
//
// Band on Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer12(value bool) error {
	return e.Element.SetProperty("xer-12", value)
}

// xer-13 (bool)
//
// Band on Right 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer13(value bool) error {
	return e.Element.SetProperty("xer-13", value)
}

// xer-14 (bool)
//
// Band on Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer14(value bool) error {
	return e.Element.SetProperty("xer-14", value)
}

// xer-15 (bool)
//
// Band on Right 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer15(value bool) error {
	return e.Element.SetProperty("xer-15", value)
}

// xer-16 (bool)
//
// Band on Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer16(value bool) error {
	return e.Element.SetProperty("xer-16", value)
}

// xer-17 (bool)
//
// Band on Right 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer17(value bool) error {
	return e.Element.SetProperty("xer-17", value)
}

// xer-18 (bool)
//
// Band on Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer18(value bool) error {
	return e.Element.SetProperty("xer-18", value)
}

// xer-19 (bool)
//
// Band on Right 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer19(value bool) error {
	return e.Element.SetProperty("xer-19", value)
}

// xer-2 (bool)
//
// Band on Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer2(value bool) error {
	return e.Element.SetProperty("xer-2", value)
}

// xer-20 (bool)
//
// Band on Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer20(value bool) error {
	return e.Element.SetProperty("xer-20", value)
}

// xer-21 (bool)
//
// Band on Right 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer21(value bool) error {
	return e.Element.SetProperty("xer-21", value)
}

// xer-22 (bool)
//
// Band on Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer22(value bool) error {
	return e.Element.SetProperty("xer-22", value)
}

// xer-23 (bool)
//
// Band on Right 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer23(value bool) error {
	return e.Element.SetProperty("xer-23", value)
}

// xer-24 (bool)
//
// Band on Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer24(value bool) error {
	return e.Element.SetProperty("xer-24", value)
}

// xer-25 (bool)
//
// Band on Right 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer25(value bool) error {
	return e.Element.SetProperty("xer-25", value)
}

// xer-26 (bool)
//
// Band on Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer26(value bool) error {
	return e.Element.SetProperty("xer-26", value)
}

// xer-27 (bool)
//
// Band on Right 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer27(value bool) error {
	return e.Element.SetProperty("xer-27", value)
}

// xer-28 (bool)
//
// Band on Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer28(value bool) error {
	return e.Element.SetProperty("xer-28", value)
}

// xer-29 (bool)
//
// Band on Right 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer29(value bool) error {
	return e.Element.SetProperty("xer-29", value)
}

// xer-3 (bool)
//
// Band on Right 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer3(value bool) error {
	return e.Element.SetProperty("xer-3", value)
}

// xer-30 (bool)
//
// Band on Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer30(value bool) error {
	return e.Element.SetProperty("xer-30", value)
}

// xer-31 (bool)
//
// Band on Right 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xer-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer31(value bool) error {
	return e.Element.SetProperty("xer-31", value)
}

// xer-4 (bool)
//
// Band on Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer4(value bool) error {
	return e.Element.SetProperty("xer-4", value)
}

// xer-5 (bool)
//
// Band on Right 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer5(value bool) error {
	return e.Element.SetProperty("xer-5", value)
}

// xer-6 (bool)
//
// Band on Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer6(value bool) error {
	return e.Element.SetProperty("xer-6", value)
}

// xer-7 (bool)
//
// Band on Right 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer7(value bool) error {
	return e.Element.SetProperty("xer-7", value)
}

// xer-8 (bool)
//
// Band on Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer8(value bool) error {
	return e.Element.SetProperty("xer-8", value)
}

// xer-9 (bool)
//
// Band on Right 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXer9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXer9(value bool) error {
	return e.Element.SetProperty("xer-9", value)
}

// xml-0 (bool)
//
// Band mute Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml0(value bool) error {
	return e.Element.SetProperty("xml-0", value)
}

// xml-1 (bool)
//
// Band mute Left 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml1(value bool) error {
	return e.Element.SetProperty("xml-1", value)
}

// xml-10 (bool)
//
// Band mute Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml10(value bool) error {
	return e.Element.SetProperty("xml-10", value)
}

// xml-11 (bool)
//
// Band mute Left 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml11(value bool) error {
	return e.Element.SetProperty("xml-11", value)
}

// xml-12 (bool)
//
// Band mute Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml12(value bool) error {
	return e.Element.SetProperty("xml-12", value)
}

// xml-13 (bool)
//
// Band mute Left 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml13(value bool) error {
	return e.Element.SetProperty("xml-13", value)
}

// xml-14 (bool)
//
// Band mute Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml14(value bool) error {
	return e.Element.SetProperty("xml-14", value)
}

// xml-15 (bool)
//
// Band mute Left 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml15(value bool) error {
	return e.Element.SetProperty("xml-15", value)
}

// xml-16 (bool)
//
// Band mute Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml16(value bool) error {
	return e.Element.SetProperty("xml-16", value)
}

// xml-17 (bool)
//
// Band mute Left 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml17(value bool) error {
	return e.Element.SetProperty("xml-17", value)
}

// xml-18 (bool)
//
// Band mute Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml18(value bool) error {
	return e.Element.SetProperty("xml-18", value)
}

// xml-19 (bool)
//
// Band mute Left 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml19(value bool) error {
	return e.Element.SetProperty("xml-19", value)
}

// xml-2 (bool)
//
// Band mute Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml2(value bool) error {
	return e.Element.SetProperty("xml-2", value)
}

// xml-20 (bool)
//
// Band mute Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml20(value bool) error {
	return e.Element.SetProperty("xml-20", value)
}

// xml-21 (bool)
//
// Band mute Left 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml21(value bool) error {
	return e.Element.SetProperty("xml-21", value)
}

// xml-22 (bool)
//
// Band mute Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml22(value bool) error {
	return e.Element.SetProperty("xml-22", value)
}

// xml-23 (bool)
//
// Band mute Left 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml23(value bool) error {
	return e.Element.SetProperty("xml-23", value)
}

// xml-24 (bool)
//
// Band mute Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml24(value bool) error {
	return e.Element.SetProperty("xml-24", value)
}

// xml-25 (bool)
//
// Band mute Left 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml25(value bool) error {
	return e.Element.SetProperty("xml-25", value)
}

// xml-26 (bool)
//
// Band mute Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml26(value bool) error {
	return e.Element.SetProperty("xml-26", value)
}

// xml-27 (bool)
//
// Band mute Left 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml27(value bool) error {
	return e.Element.SetProperty("xml-27", value)
}

// xml-28 (bool)
//
// Band mute Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml28(value bool) error {
	return e.Element.SetProperty("xml-28", value)
}

// xml-29 (bool)
//
// Band mute Left 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml29(value bool) error {
	return e.Element.SetProperty("xml-29", value)
}

// xml-3 (bool)
//
// Band mute Left 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml3(value bool) error {
	return e.Element.SetProperty("xml-3", value)
}

// xml-30 (bool)
//
// Band mute Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml30(value bool) error {
	return e.Element.SetProperty("xml-30", value)
}

// xml-31 (bool)
//
// Band mute Left 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml31(value bool) error {
	return e.Element.SetProperty("xml-31", value)
}

// xml-4 (bool)
//
// Band mute Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml4(value bool) error {
	return e.Element.SetProperty("xml-4", value)
}

// xml-5 (bool)
//
// Band mute Left 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml5(value bool) error {
	return e.Element.SetProperty("xml-5", value)
}

// xml-6 (bool)
//
// Band mute Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml6(value bool) error {
	return e.Element.SetProperty("xml-6", value)
}

// xml-7 (bool)
//
// Band mute Left 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml7(value bool) error {
	return e.Element.SetProperty("xml-7", value)
}

// xml-8 (bool)
//
// Band mute Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml8(value bool) error {
	return e.Element.SetProperty("xml-8", value)
}

// xml-9 (bool)
//
// Band mute Left 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXml9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXml9(value bool) error {
	return e.Element.SetProperty("xml-9", value)
}

// xmr-0 (bool)
//
// Band mute Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr0(value bool) error {
	return e.Element.SetProperty("xmr-0", value)
}

// xmr-1 (bool)
//
// Band mute Right 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr1(value bool) error {
	return e.Element.SetProperty("xmr-1", value)
}

// xmr-10 (bool)
//
// Band mute Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr10(value bool) error {
	return e.Element.SetProperty("xmr-10", value)
}

// xmr-11 (bool)
//
// Band mute Right 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr11(value bool) error {
	return e.Element.SetProperty("xmr-11", value)
}

// xmr-12 (bool)
//
// Band mute Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr12(value bool) error {
	return e.Element.SetProperty("xmr-12", value)
}

// xmr-13 (bool)
//
// Band mute Right 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr13(value bool) error {
	return e.Element.SetProperty("xmr-13", value)
}

// xmr-14 (bool)
//
// Band mute Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr14(value bool) error {
	return e.Element.SetProperty("xmr-14", value)
}

// xmr-15 (bool)
//
// Band mute Right 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr15(value bool) error {
	return e.Element.SetProperty("xmr-15", value)
}

// xmr-16 (bool)
//
// Band mute Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr16(value bool) error {
	return e.Element.SetProperty("xmr-16", value)
}

// xmr-17 (bool)
//
// Band mute Right 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr17(value bool) error {
	return e.Element.SetProperty("xmr-17", value)
}

// xmr-18 (bool)
//
// Band mute Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr18(value bool) error {
	return e.Element.SetProperty("xmr-18", value)
}

// xmr-19 (bool)
//
// Band mute Right 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr19(value bool) error {
	return e.Element.SetProperty("xmr-19", value)
}

// xmr-2 (bool)
//
// Band mute Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr2(value bool) error {
	return e.Element.SetProperty("xmr-2", value)
}

// xmr-20 (bool)
//
// Band mute Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr20(value bool) error {
	return e.Element.SetProperty("xmr-20", value)
}

// xmr-21 (bool)
//
// Band mute Right 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr21(value bool) error {
	return e.Element.SetProperty("xmr-21", value)
}

// xmr-22 (bool)
//
// Band mute Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr22(value bool) error {
	return e.Element.SetProperty("xmr-22", value)
}

// xmr-23 (bool)
//
// Band mute Right 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr23(value bool) error {
	return e.Element.SetProperty("xmr-23", value)
}

// xmr-24 (bool)
//
// Band mute Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr24(value bool) error {
	return e.Element.SetProperty("xmr-24", value)
}

// xmr-25 (bool)
//
// Band mute Right 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr25(value bool) error {
	return e.Element.SetProperty("xmr-25", value)
}

// xmr-26 (bool)
//
// Band mute Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr26(value bool) error {
	return e.Element.SetProperty("xmr-26", value)
}

// xmr-27 (bool)
//
// Band mute Right 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr27(value bool) error {
	return e.Element.SetProperty("xmr-27", value)
}

// xmr-28 (bool)
//
// Band mute Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr28(value bool) error {
	return e.Element.SetProperty("xmr-28", value)
}

// xmr-29 (bool)
//
// Band mute Right 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr29(value bool) error {
	return e.Element.SetProperty("xmr-29", value)
}

// xmr-3 (bool)
//
// Band mute Right 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr3(value bool) error {
	return e.Element.SetProperty("xmr-3", value)
}

// xmr-30 (bool)
//
// Band mute Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr30(value bool) error {
	return e.Element.SetProperty("xmr-30", value)
}

// xmr-31 (bool)
//
// Band mute Right 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr31(value bool) error {
	return e.Element.SetProperty("xmr-31", value)
}

// xmr-4 (bool)
//
// Band mute Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr4(value bool) error {
	return e.Element.SetProperty("xmr-4", value)
}

// xmr-5 (bool)
//
// Band mute Right 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr5(value bool) error {
	return e.Element.SetProperty("xmr-5", value)
}

// xmr-6 (bool)
//
// Band mute Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr6(value bool) error {
	return e.Element.SetProperty("xmr-6", value)
}

// xmr-7 (bool)
//
// Band mute Right 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr7(value bool) error {
	return e.Element.SetProperty("xmr-7", value)
}

// xmr-8 (bool)
//
// Band mute Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr8(value bool) error {
	return e.Element.SetProperty("xmr-8", value)
}

// xmr-9 (bool)
//
// Band mute Right 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXmr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXmr9(value bool) error {
	return e.Element.SetProperty("xmr-9", value)
}

// xsl-0 (bool)
//
// Band solo Left 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl0(value bool) error {
	return e.Element.SetProperty("xsl-0", value)
}

// xsl-1 (bool)
//
// Band solo Left 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl1(value bool) error {
	return e.Element.SetProperty("xsl-1", value)
}

// xsl-10 (bool)
//
// Band solo Left 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl10(value bool) error {
	return e.Element.SetProperty("xsl-10", value)
}

// xsl-11 (bool)
//
// Band solo Left 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl11(value bool) error {
	return e.Element.SetProperty("xsl-11", value)
}

// xsl-12 (bool)
//
// Band solo Left 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl12(value bool) error {
	return e.Element.SetProperty("xsl-12", value)
}

// xsl-13 (bool)
//
// Band solo Left 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl13(value bool) error {
	return e.Element.SetProperty("xsl-13", value)
}

// xsl-14 (bool)
//
// Band solo Left 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl14(value bool) error {
	return e.Element.SetProperty("xsl-14", value)
}

// xsl-15 (bool)
//
// Band solo Left 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl15(value bool) error {
	return e.Element.SetProperty("xsl-15", value)
}

// xsl-16 (bool)
//
// Band solo Left 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl16(value bool) error {
	return e.Element.SetProperty("xsl-16", value)
}

// xsl-17 (bool)
//
// Band solo Left 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl17(value bool) error {
	return e.Element.SetProperty("xsl-17", value)
}

// xsl-18 (bool)
//
// Band solo Left 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl18(value bool) error {
	return e.Element.SetProperty("xsl-18", value)
}

// xsl-19 (bool)
//
// Band solo Left 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl19(value bool) error {
	return e.Element.SetProperty("xsl-19", value)
}

// xsl-2 (bool)
//
// Band solo Left 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl2(value bool) error {
	return e.Element.SetProperty("xsl-2", value)
}

// xsl-20 (bool)
//
// Band solo Left 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl20(value bool) error {
	return e.Element.SetProperty("xsl-20", value)
}

// xsl-21 (bool)
//
// Band solo Left 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl21(value bool) error {
	return e.Element.SetProperty("xsl-21", value)
}

// xsl-22 (bool)
//
// Band solo Left 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl22(value bool) error {
	return e.Element.SetProperty("xsl-22", value)
}

// xsl-23 (bool)
//
// Band solo Left 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl23(value bool) error {
	return e.Element.SetProperty("xsl-23", value)
}

// xsl-24 (bool)
//
// Band solo Left 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl24(value bool) error {
	return e.Element.SetProperty("xsl-24", value)
}

// xsl-25 (bool)
//
// Band solo Left 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl25(value bool) error {
	return e.Element.SetProperty("xsl-25", value)
}

// xsl-26 (bool)
//
// Band solo Left 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl26(value bool) error {
	return e.Element.SetProperty("xsl-26", value)
}

// xsl-27 (bool)
//
// Band solo Left 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl27(value bool) error {
	return e.Element.SetProperty("xsl-27", value)
}

// xsl-28 (bool)
//
// Band solo Left 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl28(value bool) error {
	return e.Element.SetProperty("xsl-28", value)
}

// xsl-29 (bool)
//
// Band solo Left 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl29(value bool) error {
	return e.Element.SetProperty("xsl-29", value)
}

// xsl-3 (bool)
//
// Band solo Left 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl3(value bool) error {
	return e.Element.SetProperty("xsl-3", value)
}

// xsl-30 (bool)
//
// Band solo Left 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl30(value bool) error {
	return e.Element.SetProperty("xsl-30", value)
}

// xsl-31 (bool)
//
// Band solo Left 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl31(value bool) error {
	return e.Element.SetProperty("xsl-31", value)
}

// xsl-4 (bool)
//
// Band solo Left 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl4(value bool) error {
	return e.Element.SetProperty("xsl-4", value)
}

// xsl-5 (bool)
//
// Band solo Left 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl5(value bool) error {
	return e.Element.SetProperty("xsl-5", value)
}

// xsl-6 (bool)
//
// Band solo Left 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl6(value bool) error {
	return e.Element.SetProperty("xsl-6", value)
}

// xsl-7 (bool)
//
// Band solo Left 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl7(value bool) error {
	return e.Element.SetProperty("xsl-7", value)
}

// xsl-8 (bool)
//
// Band solo Left 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl8(value bool) error {
	return e.Element.SetProperty("xsl-8", value)
}

// xsl-9 (bool)
//
// Band solo Left 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsl9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsl9(value bool) error {
	return e.Element.SetProperty("xsl-9", value)
}

// xsr-0 (bool)
//
// Band solo Right 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr0(value bool) error {
	return e.Element.SetProperty("xsr-0", value)
}

// xsr-1 (bool)
//
// Band solo Right 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr1(value bool) error {
	return e.Element.SetProperty("xsr-1", value)
}

// xsr-10 (bool)
//
// Band solo Right 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr10(value bool) error {
	return e.Element.SetProperty("xsr-10", value)
}

// xsr-11 (bool)
//
// Band solo Right 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr11(value bool) error {
	return e.Element.SetProperty("xsr-11", value)
}

// xsr-12 (bool)
//
// Band solo Right 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr12(value bool) error {
	return e.Element.SetProperty("xsr-12", value)
}

// xsr-13 (bool)
//
// Band solo Right 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr13(value bool) error {
	return e.Element.SetProperty("xsr-13", value)
}

// xsr-14 (bool)
//
// Band solo Right 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr14(value bool) error {
	return e.Element.SetProperty("xsr-14", value)
}

// xsr-15 (bool)
//
// Band solo Right 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr15(value bool) error {
	return e.Element.SetProperty("xsr-15", value)
}

// xsr-16 (bool)
//
// Band solo Right 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr16(value bool) error {
	return e.Element.SetProperty("xsr-16", value)
}

// xsr-17 (bool)
//
// Band solo Right 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr17(value bool) error {
	return e.Element.SetProperty("xsr-17", value)
}

// xsr-18 (bool)
//
// Band solo Right 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr18(value bool) error {
	return e.Element.SetProperty("xsr-18", value)
}

// xsr-19 (bool)
//
// Band solo Right 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr19(value bool) error {
	return e.Element.SetProperty("xsr-19", value)
}

// xsr-2 (bool)
//
// Band solo Right 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr2(value bool) error {
	return e.Element.SetProperty("xsr-2", value)
}

// xsr-20 (bool)
//
// Band solo Right 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr20(value bool) error {
	return e.Element.SetProperty("xsr-20", value)
}

// xsr-21 (bool)
//
// Band solo Right 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr21(value bool) error {
	return e.Element.SetProperty("xsr-21", value)
}

// xsr-22 (bool)
//
// Band solo Right 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr22(value bool) error {
	return e.Element.SetProperty("xsr-22", value)
}

// xsr-23 (bool)
//
// Band solo Right 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr23(value bool) error {
	return e.Element.SetProperty("xsr-23", value)
}

// xsr-24 (bool)
//
// Band solo Right 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr24(value bool) error {
	return e.Element.SetProperty("xsr-24", value)
}

// xsr-25 (bool)
//
// Band solo Right 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr25(value bool) error {
	return e.Element.SetProperty("xsr-25", value)
}

// xsr-26 (bool)
//
// Band solo Right 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr26(value bool) error {
	return e.Element.SetProperty("xsr-26", value)
}

// xsr-27 (bool)
//
// Band solo Right 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr27(value bool) error {
	return e.Element.SetProperty("xsr-27", value)
}

// xsr-28 (bool)
//
// Band solo Right 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr28(value bool) error {
	return e.Element.SetProperty("xsr-28", value)
}

// xsr-29 (bool)
//
// Band solo Right 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr29(value bool) error {
	return e.Element.SetProperty("xsr-29", value)
}

// xsr-3 (bool)
//
// Band solo Right 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr3(value bool) error {
	return e.Element.SetProperty("xsr-3", value)
}

// xsr-30 (bool)
//
// Band solo Right 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr30(value bool) error {
	return e.Element.SetProperty("xsr-30", value)
}

// xsr-31 (bool)
//
// Band solo Right 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr31(value bool) error {
	return e.Element.SetProperty("xsr-31", value)
}

// xsr-4 (bool)
//
// Band solo Right 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr4(value bool) error {
	return e.Element.SetProperty("xsr-4", value)
}

// xsr-5 (bool)
//
// Band solo Right 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr5(value bool) error {
	return e.Element.SetProperty("xsr-5", value)
}

// xsr-6 (bool)
//
// Band solo Right 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr6(value bool) error {
	return e.Element.SetProperty("xsr-6", value)
}

// xsr-7 (bool)
//
// Band solo Right 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr7(value bool) error {
	return e.Element.SetProperty("xsr-7", value)
}

// xsr-8 (bool)
//
// Band solo Right 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr8(value bool) error {
	return e.Element.SetProperty("xsr-8", value)
}

// xsr-9 (bool)
//
// Band solo Right 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetXsr9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetXsr9(value bool) error {
	return e.Element.SetProperty("xsr-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Lr) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GraphEqualizerX32Lrfft string

const (
	LspPlugInPluginsLv2GraphEqualizerX32LrfftOff LspPlugInPluginsLv2GraphEqualizerX32Lrfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2GraphEqualizerX32LrfftPostEq LspPlugInPluginsLv2GraphEqualizerX32Lrfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2GraphEqualizerX32LrfftPreEq LspPlugInPluginsLv2GraphEqualizerX32Lrfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2GraphEqualizerX32Lrfsel string

const (
	LspPlugInPluginsLv2GraphEqualizerX32LrfselBandsLeft015 LspPlugInPluginsLv2GraphEqualizerX32Lrfsel = "Bands Left 0-15" // Bands Left 0-15 (0) – Bands Left 0-15
	LspPlugInPluginsLv2GraphEqualizerX32LrfselBandsRight015 LspPlugInPluginsLv2GraphEqualizerX32Lrfsel = "Bands Right 0-15" // Bands Right 0-15 (1) – Bands Right 0-15
	LspPlugInPluginsLv2GraphEqualizerX32LrfselBandsLeft1631 LspPlugInPluginsLv2GraphEqualizerX32Lrfsel = "Bands Left 16-31" // Bands Left 16-31 (2) – Bands Left 16-31
	LspPlugInPluginsLv2GraphEqualizerX32LrfselBandsRight1631 LspPlugInPluginsLv2GraphEqualizerX32Lrfsel = "Bands Right 16-31" // Bands Right 16-31 (3) – Bands Right 16-31
)

type LspPlugInPluginsLv2GraphEqualizerX32Lrmode string

const (
	LspPlugInPluginsLv2GraphEqualizerX32LrmodeIir LspPlugInPluginsLv2GraphEqualizerX32Lrmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2GraphEqualizerX32LrmodeFir LspPlugInPluginsLv2GraphEqualizerX32Lrmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2GraphEqualizerX32LrmodeFft LspPlugInPluginsLv2GraphEqualizerX32Lrmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2GraphEqualizerX32Lrslope string

const (
	LspPlugInPluginsLv2GraphEqualizerX32LrslopeBt48 LspPlugInPluginsLv2GraphEqualizerX32Lrslope = "BT48" // BT48 (0) – BT48
	LspPlugInPluginsLv2GraphEqualizerX32LrslopeMt48 LspPlugInPluginsLv2GraphEqualizerX32Lrslope = "MT48" // MT48 (1) – MT48
	LspPlugInPluginsLv2GraphEqualizerX32LrslopeBt72 LspPlugInPluginsLv2GraphEqualizerX32Lrslope = "BT72" // BT72 (2) – BT72
	LspPlugInPluginsLv2GraphEqualizerX32LrslopeMt72 LspPlugInPluginsLv2GraphEqualizerX32Lrslope = "MT72" // MT72 (3) – MT72
	LspPlugInPluginsLv2GraphEqualizerX32LrslopeBt96 LspPlugInPluginsLv2GraphEqualizerX32Lrslope = "BT96" // BT96 (4) – BT96
	LspPlugInPluginsLv2GraphEqualizerX32LrslopeMt96 LspPlugInPluginsLv2GraphEqualizerX32Lrslope = "MT96" // MT96 (5) – MT96
)

