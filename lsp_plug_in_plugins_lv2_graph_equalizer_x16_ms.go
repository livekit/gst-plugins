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

type LspPlugInPluginsLv2GraphEqualizerX16Ms struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GraphEqualizerX16Ms() (*LspPlugInPluginsLv2GraphEqualizerX16Ms, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-graph-equalizer-x16-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX16Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GraphEqualizerX16MsWithName(name string) (*LspPlugInPluginsLv2GraphEqualizerX16Ms, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-graph-equalizer-x16-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX16Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2GraphEqualizerX16Msfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-m (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFftvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetFftvM(value bool) error {
	return e.Element.SetProperty("fftv-m", value)
}

// fftv-s (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFftvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetFftvS(value bool) error {
	return e.Element.SetProperty("fftv-s", value)
}

// fsel (GstLspPlugInPluginsLv2GraphEqualizerX16Msfsel)
//
// Band select

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// fvm-0 (bool)
//
// Filter visibility  Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm0() (bool, error) {
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
// Filter visibility  Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm1() (bool, error) {
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
// Filter visibility  Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm10() (bool, error) {
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
// Filter visibility  Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm11() (bool, error) {
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
// Filter visibility  Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm12() (bool, error) {
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
// Filter visibility  Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm13() (bool, error) {
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
// Filter visibility  Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm14() (bool, error) {
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
// Filter visibility  Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm15() (bool, error) {
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
// Filter visibility  Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm2() (bool, error) {
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
// Filter visibility  Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm3() (bool, error) {
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
// Filter visibility  Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm4() (bool, error) {
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
// Filter visibility  Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm5() (bool, error) {
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
// Filter visibility  Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm6() (bool, error) {
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
// Filter visibility  Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm7() (bool, error) {
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
// Filter visibility  Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm8() (bool, error) {
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
// Filter visibility  Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvm9() (bool, error) {
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
// Filter visibility  Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs0() (bool, error) {
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
// Filter visibility  Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs1() (bool, error) {
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
// Filter visibility  Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs10() (bool, error) {
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
// Filter visibility  Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs11() (bool, error) {
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
// Filter visibility  Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs12() (bool, error) {
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
// Filter visibility  Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs13() (bool, error) {
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
// Filter visibility  Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs14() (bool, error) {
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
// Filter visibility  Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs15() (bool, error) {
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
// Filter visibility  Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs2() (bool, error) {
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
// Filter visibility  Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs3() (bool, error) {
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
// Filter visibility  Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs4() (bool, error) {
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
// Filter visibility  Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs5() (bool, error) {
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
// Filter visibility  Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs6() (bool, error) {
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
// Filter visibility  Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs7() (bool, error) {
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
// Filter visibility  Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs8() (bool, error) {
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
// Filter visibility  Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetFvs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gain-m (float32)
//
// Mid gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGainM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGainM(value float32) error {
	return e.Element.SetProperty("gain-m", value)
}

// gain-s (float32)
//
// Side gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGainS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGainS(value float32) error {
	return e.Element.SetProperty("gain-s", value)
}

// gm-0 (float32)
//
// Band gain Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm0(value float32) error {
	return e.Element.SetProperty("gm-0", value)
}

// gm-1 (float32)
//
// Band gain Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm1(value float32) error {
	return e.Element.SetProperty("gm-1", value)
}

// gm-10 (float32)
//
// Band gain Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm10(value float32) error {
	return e.Element.SetProperty("gm-10", value)
}

// gm-11 (float32)
//
// Band gain Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm11(value float32) error {
	return e.Element.SetProperty("gm-11", value)
}

// gm-12 (float32)
//
// Band gain Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm12(value float32) error {
	return e.Element.SetProperty("gm-12", value)
}

// gm-13 (float32)
//
// Band gain Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm13(value float32) error {
	return e.Element.SetProperty("gm-13", value)
}

// gm-14 (float32)
//
// Band gain Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm14(value float32) error {
	return e.Element.SetProperty("gm-14", value)
}

// gm-15 (float32)
//
// Band gain Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm15(value float32) error {
	return e.Element.SetProperty("gm-15", value)
}

// gm-2 (float32)
//
// Band gain Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm2(value float32) error {
	return e.Element.SetProperty("gm-2", value)
}

// gm-3 (float32)
//
// Band gain Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm3(value float32) error {
	return e.Element.SetProperty("gm-3", value)
}

// gm-4 (float32)
//
// Band gain Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm4(value float32) error {
	return e.Element.SetProperty("gm-4", value)
}

// gm-5 (float32)
//
// Band gain Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm5(value float32) error {
	return e.Element.SetProperty("gm-5", value)
}

// gm-6 (float32)
//
// Band gain Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm6(value float32) error {
	return e.Element.SetProperty("gm-6", value)
}

// gm-7 (float32)
//
// Band gain Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm7(value float32) error {
	return e.Element.SetProperty("gm-7", value)
}

// gm-8 (float32)
//
// Band gain Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm8(value float32) error {
	return e.Element.SetProperty("gm-8", value)
}

// gm-9 (float32)
//
// Band gain Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGm9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGm9(value float32) error {
	return e.Element.SetProperty("gm-9", value)
}

// gs-0 (float32)
//
// Band gain Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs0(value float32) error {
	return e.Element.SetProperty("gs-0", value)
}

// gs-1 (float32)
//
// Band gain Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs1(value float32) error {
	return e.Element.SetProperty("gs-1", value)
}

// gs-10 (float32)
//
// Band gain Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs10(value float32) error {
	return e.Element.SetProperty("gs-10", value)
}

// gs-11 (float32)
//
// Band gain Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs11(value float32) error {
	return e.Element.SetProperty("gs-11", value)
}

// gs-12 (float32)
//
// Band gain Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs12(value float32) error {
	return e.Element.SetProperty("gs-12", value)
}

// gs-13 (float32)
//
// Band gain Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs13(value float32) error {
	return e.Element.SetProperty("gs-13", value)
}

// gs-14 (float32)
//
// Band gain Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs14(value float32) error {
	return e.Element.SetProperty("gs-14", value)
}

// gs-15 (float32)
//
// Band gain Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs15(value float32) error {
	return e.Element.SetProperty("gs-15", value)
}

// gs-2 (float32)
//
// Band gain Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs2(value float32) error {
	return e.Element.SetProperty("gs-2", value)
}

// gs-3 (float32)
//
// Band gain Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs3(value float32) error {
	return e.Element.SetProperty("gs-3", value)
}

// gs-4 (float32)
//
// Band gain Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs4(value float32) error {
	return e.Element.SetProperty("gs-4", value)
}

// gs-5 (float32)
//
// Band gain Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs5(value float32) error {
	return e.Element.SetProperty("gs-5", value)
}

// gs-6 (float32)
//
// Band gain Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs6(value float32) error {
	return e.Element.SetProperty("gs-6", value)
}

// gs-7 (float32)
//
// Band gain Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs7(value float32) error {
	return e.Element.SetProperty("gs-7", value)
}

// gs-8 (float32)
//
// Band gain Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs8(value float32) error {
	return e.Element.SetProperty("gs-8", value)
}

// gs-9 (float32)
//
// Band gain Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetGs9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetGs9(value float32) error {
	return e.Element.SetProperty("gs-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetImr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetLstn() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mode (GstLspPlugInPluginsLv2GraphEqualizerX16Msmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// slope (GstLspPlugInPluginsLv2GraphEqualizerX16Msslope)
//
// Filter slope

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetSlope() (interface{}, error) {
	return e.Element.GetProperty("slope")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetSlope(value interface{}) error {
	return e.Element.SetProperty("slope", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetSmr() (float32, error) {
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

// xem-0 (bool)
//
// Band on Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem0(value bool) error {
	return e.Element.SetProperty("xem-0", value)
}

// xem-1 (bool)
//
// Band on Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem1(value bool) error {
	return e.Element.SetProperty("xem-1", value)
}

// xem-10 (bool)
//
// Band on Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem10(value bool) error {
	return e.Element.SetProperty("xem-10", value)
}

// xem-11 (bool)
//
// Band on Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem11(value bool) error {
	return e.Element.SetProperty("xem-11", value)
}

// xem-12 (bool)
//
// Band on Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem12(value bool) error {
	return e.Element.SetProperty("xem-12", value)
}

// xem-13 (bool)
//
// Band on Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem13(value bool) error {
	return e.Element.SetProperty("xem-13", value)
}

// xem-14 (bool)
//
// Band on Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem14(value bool) error {
	return e.Element.SetProperty("xem-14", value)
}

// xem-15 (bool)
//
// Band on Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem15(value bool) error {
	return e.Element.SetProperty("xem-15", value)
}

// xem-2 (bool)
//
// Band on Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem2(value bool) error {
	return e.Element.SetProperty("xem-2", value)
}

// xem-3 (bool)
//
// Band on Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem3(value bool) error {
	return e.Element.SetProperty("xem-3", value)
}

// xem-4 (bool)
//
// Band on Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem4(value bool) error {
	return e.Element.SetProperty("xem-4", value)
}

// xem-5 (bool)
//
// Band on Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem5(value bool) error {
	return e.Element.SetProperty("xem-5", value)
}

// xem-6 (bool)
//
// Band on Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem6(value bool) error {
	return e.Element.SetProperty("xem-6", value)
}

// xem-7 (bool)
//
// Band on Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem7(value bool) error {
	return e.Element.SetProperty("xem-7", value)
}

// xem-8 (bool)
//
// Band on Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem8(value bool) error {
	return e.Element.SetProperty("xem-8", value)
}

// xem-9 (bool)
//
// Band on Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXem9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXem9(value bool) error {
	return e.Element.SetProperty("xem-9", value)
}

// xes-0 (bool)
//
// Band on Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes0(value bool) error {
	return e.Element.SetProperty("xes-0", value)
}

// xes-1 (bool)
//
// Band on Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes1(value bool) error {
	return e.Element.SetProperty("xes-1", value)
}

// xes-10 (bool)
//
// Band on Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes10(value bool) error {
	return e.Element.SetProperty("xes-10", value)
}

// xes-11 (bool)
//
// Band on Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes11(value bool) error {
	return e.Element.SetProperty("xes-11", value)
}

// xes-12 (bool)
//
// Band on Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes12(value bool) error {
	return e.Element.SetProperty("xes-12", value)
}

// xes-13 (bool)
//
// Band on Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes13(value bool) error {
	return e.Element.SetProperty("xes-13", value)
}

// xes-14 (bool)
//
// Band on Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes14(value bool) error {
	return e.Element.SetProperty("xes-14", value)
}

// xes-15 (bool)
//
// Band on Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes15(value bool) error {
	return e.Element.SetProperty("xes-15", value)
}

// xes-2 (bool)
//
// Band on Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes2(value bool) error {
	return e.Element.SetProperty("xes-2", value)
}

// xes-3 (bool)
//
// Band on Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes3(value bool) error {
	return e.Element.SetProperty("xes-3", value)
}

// xes-4 (bool)
//
// Band on Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes4(value bool) error {
	return e.Element.SetProperty("xes-4", value)
}

// xes-5 (bool)
//
// Band on Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes5(value bool) error {
	return e.Element.SetProperty("xes-5", value)
}

// xes-6 (bool)
//
// Band on Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes6(value bool) error {
	return e.Element.SetProperty("xes-6", value)
}

// xes-7 (bool)
//
// Band on Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes7(value bool) error {
	return e.Element.SetProperty("xes-7", value)
}

// xes-8 (bool)
//
// Band on Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes8(value bool) error {
	return e.Element.SetProperty("xes-8", value)
}

// xes-9 (bool)
//
// Band on Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXes9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXes9(value bool) error {
	return e.Element.SetProperty("xes-9", value)
}

// xmm-0 (bool)
//
// Band mute Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm0(value bool) error {
	return e.Element.SetProperty("xmm-0", value)
}

// xmm-1 (bool)
//
// Band mute Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm1(value bool) error {
	return e.Element.SetProperty("xmm-1", value)
}

// xmm-10 (bool)
//
// Band mute Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm10(value bool) error {
	return e.Element.SetProperty("xmm-10", value)
}

// xmm-11 (bool)
//
// Band mute Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm11(value bool) error {
	return e.Element.SetProperty("xmm-11", value)
}

// xmm-12 (bool)
//
// Band mute Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm12(value bool) error {
	return e.Element.SetProperty("xmm-12", value)
}

// xmm-13 (bool)
//
// Band mute Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm13(value bool) error {
	return e.Element.SetProperty("xmm-13", value)
}

// xmm-14 (bool)
//
// Band mute Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm14(value bool) error {
	return e.Element.SetProperty("xmm-14", value)
}

// xmm-15 (bool)
//
// Band mute Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm15(value bool) error {
	return e.Element.SetProperty("xmm-15", value)
}

// xmm-2 (bool)
//
// Band mute Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm2(value bool) error {
	return e.Element.SetProperty("xmm-2", value)
}

// xmm-3 (bool)
//
// Band mute Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm3(value bool) error {
	return e.Element.SetProperty("xmm-3", value)
}

// xmm-4 (bool)
//
// Band mute Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm4(value bool) error {
	return e.Element.SetProperty("xmm-4", value)
}

// xmm-5 (bool)
//
// Band mute Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm5(value bool) error {
	return e.Element.SetProperty("xmm-5", value)
}

// xmm-6 (bool)
//
// Band mute Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm6(value bool) error {
	return e.Element.SetProperty("xmm-6", value)
}

// xmm-7 (bool)
//
// Band mute Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm7(value bool) error {
	return e.Element.SetProperty("xmm-7", value)
}

// xmm-8 (bool)
//
// Band mute Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm8(value bool) error {
	return e.Element.SetProperty("xmm-8", value)
}

// xmm-9 (bool)
//
// Band mute Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXmm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXmm9(value bool) error {
	return e.Element.SetProperty("xmm-9", value)
}

// xms-0 (bool)
//
// Band mute Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms0(value bool) error {
	return e.Element.SetProperty("xms-0", value)
}

// xms-1 (bool)
//
// Band mute Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms1(value bool) error {
	return e.Element.SetProperty("xms-1", value)
}

// xms-10 (bool)
//
// Band mute Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms10(value bool) error {
	return e.Element.SetProperty("xms-10", value)
}

// xms-11 (bool)
//
// Band mute Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms11(value bool) error {
	return e.Element.SetProperty("xms-11", value)
}

// xms-12 (bool)
//
// Band mute Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms12(value bool) error {
	return e.Element.SetProperty("xms-12", value)
}

// xms-13 (bool)
//
// Band mute Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms13(value bool) error {
	return e.Element.SetProperty("xms-13", value)
}

// xms-14 (bool)
//
// Band mute Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms14(value bool) error {
	return e.Element.SetProperty("xms-14", value)
}

// xms-15 (bool)
//
// Band mute Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms15(value bool) error {
	return e.Element.SetProperty("xms-15", value)
}

// xms-2 (bool)
//
// Band mute Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms2(value bool) error {
	return e.Element.SetProperty("xms-2", value)
}

// xms-3 (bool)
//
// Band mute Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms3(value bool) error {
	return e.Element.SetProperty("xms-3", value)
}

// xms-4 (bool)
//
// Band mute Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms4(value bool) error {
	return e.Element.SetProperty("xms-4", value)
}

// xms-5 (bool)
//
// Band mute Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms5(value bool) error {
	return e.Element.SetProperty("xms-5", value)
}

// xms-6 (bool)
//
// Band mute Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms6(value bool) error {
	return e.Element.SetProperty("xms-6", value)
}

// xms-7 (bool)
//
// Band mute Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms7(value bool) error {
	return e.Element.SetProperty("xms-7", value)
}

// xms-8 (bool)
//
// Band mute Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms8(value bool) error {
	return e.Element.SetProperty("xms-8", value)
}

// xms-9 (bool)
//
// Band mute Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXms9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXms9(value bool) error {
	return e.Element.SetProperty("xms-9", value)
}

// xsm-0 (bool)
//
// Band solo Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm0(value bool) error {
	return e.Element.SetProperty("xsm-0", value)
}

// xsm-1 (bool)
//
// Band solo Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm1(value bool) error {
	return e.Element.SetProperty("xsm-1", value)
}

// xsm-10 (bool)
//
// Band solo Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm10(value bool) error {
	return e.Element.SetProperty("xsm-10", value)
}

// xsm-11 (bool)
//
// Band solo Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm11(value bool) error {
	return e.Element.SetProperty("xsm-11", value)
}

// xsm-12 (bool)
//
// Band solo Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm12(value bool) error {
	return e.Element.SetProperty("xsm-12", value)
}

// xsm-13 (bool)
//
// Band solo Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm13(value bool) error {
	return e.Element.SetProperty("xsm-13", value)
}

// xsm-14 (bool)
//
// Band solo Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm14(value bool) error {
	return e.Element.SetProperty("xsm-14", value)
}

// xsm-15 (bool)
//
// Band solo Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm15(value bool) error {
	return e.Element.SetProperty("xsm-15", value)
}

// xsm-2 (bool)
//
// Band solo Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm2(value bool) error {
	return e.Element.SetProperty("xsm-2", value)
}

// xsm-3 (bool)
//
// Band solo Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm3(value bool) error {
	return e.Element.SetProperty("xsm-3", value)
}

// xsm-4 (bool)
//
// Band solo Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm4(value bool) error {
	return e.Element.SetProperty("xsm-4", value)
}

// xsm-5 (bool)
//
// Band solo Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm5(value bool) error {
	return e.Element.SetProperty("xsm-5", value)
}

// xsm-6 (bool)
//
// Band solo Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm6(value bool) error {
	return e.Element.SetProperty("xsm-6", value)
}

// xsm-7 (bool)
//
// Band solo Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm7(value bool) error {
	return e.Element.SetProperty("xsm-7", value)
}

// xsm-8 (bool)
//
// Band solo Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm8(value bool) error {
	return e.Element.SetProperty("xsm-8", value)
}

// xsm-9 (bool)
//
// Band solo Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXsm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXsm9(value bool) error {
	return e.Element.SetProperty("xsm-9", value)
}

// xss-0 (bool)
//
// Band solo Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss0(value bool) error {
	return e.Element.SetProperty("xss-0", value)
}

// xss-1 (bool)
//
// Band solo Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss1(value bool) error {
	return e.Element.SetProperty("xss-1", value)
}

// xss-10 (bool)
//
// Band solo Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss10(value bool) error {
	return e.Element.SetProperty("xss-10", value)
}

// xss-11 (bool)
//
// Band solo Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss11(value bool) error {
	return e.Element.SetProperty("xss-11", value)
}

// xss-12 (bool)
//
// Band solo Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss12(value bool) error {
	return e.Element.SetProperty("xss-12", value)
}

// xss-13 (bool)
//
// Band solo Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss13(value bool) error {
	return e.Element.SetProperty("xss-13", value)
}

// xss-14 (bool)
//
// Band solo Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss14(value bool) error {
	return e.Element.SetProperty("xss-14", value)
}

// xss-15 (bool)
//
// Band solo Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss15(value bool) error {
	return e.Element.SetProperty("xss-15", value)
}

// xss-2 (bool)
//
// Band solo Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss2(value bool) error {
	return e.Element.SetProperty("xss-2", value)
}

// xss-3 (bool)
//
// Band solo Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss3(value bool) error {
	return e.Element.SetProperty("xss-3", value)
}

// xss-4 (bool)
//
// Band solo Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss4(value bool) error {
	return e.Element.SetProperty("xss-4", value)
}

// xss-5 (bool)
//
// Band solo Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss5(value bool) error {
	return e.Element.SetProperty("xss-5", value)
}

// xss-6 (bool)
//
// Band solo Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss6(value bool) error {
	return e.Element.SetProperty("xss-6", value)
}

// xss-7 (bool)
//
// Band solo Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss7(value bool) error {
	return e.Element.SetProperty("xss-7", value)
}

// xss-8 (bool)
//
// Band solo Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss8(value bool) error {
	return e.Element.SetProperty("xss-8", value)
}

// xss-9 (bool)
//
// Band solo Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetXss9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetXss9(value bool) error {
	return e.Element.SetProperty("xss-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Ms) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GraphEqualizerX16Msfsel string

const (
	LspPlugInPluginsLv2GraphEqualizerX16MsfselBandsMiddle LspPlugInPluginsLv2GraphEqualizerX16Msfsel = "Bands Middle" // Bands Middle (0) – Bands Middle
	LspPlugInPluginsLv2GraphEqualizerX16MsfselBandsSide LspPlugInPluginsLv2GraphEqualizerX16Msfsel = "Bands Side" // Bands Side (1) – Bands Side
)

type LspPlugInPluginsLv2GraphEqualizerX16Msmode string

const (
	LspPlugInPluginsLv2GraphEqualizerX16MsmodeIir LspPlugInPluginsLv2GraphEqualizerX16Msmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2GraphEqualizerX16MsmodeFir LspPlugInPluginsLv2GraphEqualizerX16Msmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2GraphEqualizerX16MsmodeFft LspPlugInPluginsLv2GraphEqualizerX16Msmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2GraphEqualizerX16Msslope string

const (
	LspPlugInPluginsLv2GraphEqualizerX16MsslopeBt48 LspPlugInPluginsLv2GraphEqualizerX16Msslope = "BT48" // BT48 (0) – BT48
	LspPlugInPluginsLv2GraphEqualizerX16MsslopeMt48 LspPlugInPluginsLv2GraphEqualizerX16Msslope = "MT48" // MT48 (1) – MT48
	LspPlugInPluginsLv2GraphEqualizerX16MsslopeBt72 LspPlugInPluginsLv2GraphEqualizerX16Msslope = "BT72" // BT72 (2) – BT72
	LspPlugInPluginsLv2GraphEqualizerX16MsslopeMt72 LspPlugInPluginsLv2GraphEqualizerX16Msslope = "MT72" // MT72 (3) – MT72
	LspPlugInPluginsLv2GraphEqualizerX16MsslopeBt96 LspPlugInPluginsLv2GraphEqualizerX16Msslope = "BT96" // BT96 (4) – BT96
	LspPlugInPluginsLv2GraphEqualizerX16MsslopeMt96 LspPlugInPluginsLv2GraphEqualizerX16Msslope = "MT96" // MT96 (5) – MT96
)

type LspPlugInPluginsLv2GraphEqualizerX16Msfft string

const (
	LspPlugInPluginsLv2GraphEqualizerX16MsfftOff LspPlugInPluginsLv2GraphEqualizerX16Msfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2GraphEqualizerX16MsfftPostEq LspPlugInPluginsLv2GraphEqualizerX16Msfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2GraphEqualizerX16MsfftPreEq LspPlugInPluginsLv2GraphEqualizerX16Msfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

