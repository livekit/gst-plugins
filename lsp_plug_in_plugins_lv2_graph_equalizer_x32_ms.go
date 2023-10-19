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

type LspPlugInPluginsLv2GraphEqualizerX32Ms struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GraphEqualizerX32Ms() (*LspPlugInPluginsLv2GraphEqualizerX32Ms, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-graph-equalizer-x32-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX32Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GraphEqualizerX32MsWithName(name string) (*LspPlugInPluginsLv2GraphEqualizerX32Ms, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-graph-equalizer-x32-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX32Ms{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetBal() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2GraphEqualizerX32Msfft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-m (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFftvM() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetFftvM(value bool) error {
	return e.Element.SetProperty("fftv-m", value)
}

// fftv-s (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFftvS() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetFftvS(value bool) error {
	return e.Element.SetProperty("fftv-s", value)
}

// fsel (GstLspPlugInPluginsLv2GraphEqualizerX32Msfsel)
//
// Band select

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// fvm-0 (bool)
//
// Filter visibility  Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm0() (bool, error) {
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
// Filter visibility  Mid 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm1() (bool, error) {
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
// Filter visibility  Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm10() (bool, error) {
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
// Filter visibility  Mid 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm11() (bool, error) {
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
// Filter visibility  Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm12() (bool, error) {
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
// Filter visibility  Mid 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm13() (bool, error) {
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
// Filter visibility  Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm14() (bool, error) {
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
// Filter visibility  Mid 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm15() (bool, error) {
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
// Filter visibility  Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm16() (bool, error) {
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
// Filter visibility  Mid 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm17() (bool, error) {
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
// Filter visibility  Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm18() (bool, error) {
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
// Filter visibility  Mid 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm19() (bool, error) {
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
// Filter visibility  Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm2() (bool, error) {
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
// Filter visibility  Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm20() (bool, error) {
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
// Filter visibility  Mid 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm21() (bool, error) {
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
// Filter visibility  Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm22() (bool, error) {
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
// Filter visibility  Mid 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm23() (bool, error) {
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
// Filter visibility  Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm24() (bool, error) {
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
// Filter visibility  Mid 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm25() (bool, error) {
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
// Filter visibility  Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm26() (bool, error) {
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
// Filter visibility  Mid 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm27() (bool, error) {
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
// Filter visibility  Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm28() (bool, error) {
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
// Filter visibility  Mid 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm29() (bool, error) {
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
// Filter visibility  Mid 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm3() (bool, error) {
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
// Filter visibility  Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm30() (bool, error) {
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
// Filter visibility  Mid 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm31() (bool, error) {
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
// Filter visibility  Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm4() (bool, error) {
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
// Filter visibility  Mid 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm5() (bool, error) {
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
// Filter visibility  Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm6() (bool, error) {
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
// Filter visibility  Mid 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm7() (bool, error) {
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
// Filter visibility  Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm8() (bool, error) {
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
// Filter visibility  Mid 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs0() (bool, error) {
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
// Filter visibility  Side 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs1() (bool, error) {
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
// Filter visibility  Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs10() (bool, error) {
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
// Filter visibility  Side 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs11() (bool, error) {
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
// Filter visibility  Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs12() (bool, error) {
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
// Filter visibility  Side 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs13() (bool, error) {
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
// Filter visibility  Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs14() (bool, error) {
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
// Filter visibility  Side 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs15() (bool, error) {
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
// Filter visibility  Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs16() (bool, error) {
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
// Filter visibility  Side 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs17() (bool, error) {
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
// Filter visibility  Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs18() (bool, error) {
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
// Filter visibility  Side 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs19() (bool, error) {
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
// Filter visibility  Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs2() (bool, error) {
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
// Filter visibility  Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs20() (bool, error) {
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
// Filter visibility  Side 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs21() (bool, error) {
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
// Filter visibility  Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs22() (bool, error) {
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
// Filter visibility  Side 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs23() (bool, error) {
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
// Filter visibility  Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs24() (bool, error) {
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
// Filter visibility  Side 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs25() (bool, error) {
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
// Filter visibility  Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs26() (bool, error) {
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
// Filter visibility  Side 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs27() (bool, error) {
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
// Filter visibility  Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs28() (bool, error) {
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
// Filter visibility  Side 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs29() (bool, error) {
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
// Filter visibility  Side 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs3() (bool, error) {
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
// Filter visibility  Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs30() (bool, error) {
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
// Filter visibility  Side 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs31() (bool, error) {
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
// Filter visibility  Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs4() (bool, error) {
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
// Filter visibility  Side 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs5() (bool, error) {
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
// Filter visibility  Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs6() (bool, error) {
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
// Filter visibility  Side 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs7() (bool, error) {
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
// Filter visibility  Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs8() (bool, error) {
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
// Filter visibility  Side 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetFvs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// gain-m (float32)
//
// Mid gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGainM() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGainM(value float32) error {
	return e.Element.SetProperty("gain-m", value)
}

// gain-s (float32)
//
// Side gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGainS() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGainS(value float32) error {
	return e.Element.SetProperty("gain-s", value)
}

// gm-0 (float32)
//
// Band gain Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm0(value float32) error {
	return e.Element.SetProperty("gm-0", value)
}

// gm-1 (float32)
//
// Band gain Mid 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm1(value float32) error {
	return e.Element.SetProperty("gm-1", value)
}

// gm-10 (float32)
//
// Band gain Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm10(value float32) error {
	return e.Element.SetProperty("gm-10", value)
}

// gm-11 (float32)
//
// Band gain Mid 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm11(value float32) error {
	return e.Element.SetProperty("gm-11", value)
}

// gm-12 (float32)
//
// Band gain Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm12(value float32) error {
	return e.Element.SetProperty("gm-12", value)
}

// gm-13 (float32)
//
// Band gain Mid 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm13(value float32) error {
	return e.Element.SetProperty("gm-13", value)
}

// gm-14 (float32)
//
// Band gain Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm14(value float32) error {
	return e.Element.SetProperty("gm-14", value)
}

// gm-15 (float32)
//
// Band gain Mid 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm15(value float32) error {
	return e.Element.SetProperty("gm-15", value)
}

// gm-16 (float32)
//
// Band gain Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm16() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm16(value float32) error {
	return e.Element.SetProperty("gm-16", value)
}

// gm-17 (float32)
//
// Band gain Mid 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm17() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm17(value float32) error {
	return e.Element.SetProperty("gm-17", value)
}

// gm-18 (float32)
//
// Band gain Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm18() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm18(value float32) error {
	return e.Element.SetProperty("gm-18", value)
}

// gm-19 (float32)
//
// Band gain Mid 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm19() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm19(value float32) error {
	return e.Element.SetProperty("gm-19", value)
}

// gm-2 (float32)
//
// Band gain Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm2(value float32) error {
	return e.Element.SetProperty("gm-2", value)
}

// gm-20 (float32)
//
// Band gain Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm20() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm20(value float32) error {
	return e.Element.SetProperty("gm-20", value)
}

// gm-21 (float32)
//
// Band gain Mid 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm21() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm21(value float32) error {
	return e.Element.SetProperty("gm-21", value)
}

// gm-22 (float32)
//
// Band gain Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm22() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm22(value float32) error {
	return e.Element.SetProperty("gm-22", value)
}

// gm-23 (float32)
//
// Band gain Mid 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm23() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm23(value float32) error {
	return e.Element.SetProperty("gm-23", value)
}

// gm-24 (float32)
//
// Band gain Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm24() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm24(value float32) error {
	return e.Element.SetProperty("gm-24", value)
}

// gm-25 (float32)
//
// Band gain Mid 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm25() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm25(value float32) error {
	return e.Element.SetProperty("gm-25", value)
}

// gm-26 (float32)
//
// Band gain Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm26() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm26(value float32) error {
	return e.Element.SetProperty("gm-26", value)
}

// gm-27 (float32)
//
// Band gain Mid 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm27() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm27(value float32) error {
	return e.Element.SetProperty("gm-27", value)
}

// gm-28 (float32)
//
// Band gain Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm28() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm28(value float32) error {
	return e.Element.SetProperty("gm-28", value)
}

// gm-29 (float32)
//
// Band gain Mid 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm29() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm29(value float32) error {
	return e.Element.SetProperty("gm-29", value)
}

// gm-3 (float32)
//
// Band gain Mid 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm3(value float32) error {
	return e.Element.SetProperty("gm-3", value)
}

// gm-30 (float32)
//
// Band gain Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm30() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm30(value float32) error {
	return e.Element.SetProperty("gm-30", value)
}

// gm-31 (float32)
//
// Band gain Mid 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm31() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm31(value float32) error {
	return e.Element.SetProperty("gm-31", value)
}

// gm-4 (float32)
//
// Band gain Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm4(value float32) error {
	return e.Element.SetProperty("gm-4", value)
}

// gm-5 (float32)
//
// Band gain Mid 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm5(value float32) error {
	return e.Element.SetProperty("gm-5", value)
}

// gm-6 (float32)
//
// Band gain Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm6(value float32) error {
	return e.Element.SetProperty("gm-6", value)
}

// gm-7 (float32)
//
// Band gain Mid 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm7(value float32) error {
	return e.Element.SetProperty("gm-7", value)
}

// gm-8 (float32)
//
// Band gain Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm8(value float32) error {
	return e.Element.SetProperty("gm-8", value)
}

// gm-9 (float32)
//
// Band gain Mid 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGm9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGm9(value float32) error {
	return e.Element.SetProperty("gm-9", value)
}

// gs-0 (float32)
//
// Band gain Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs0(value float32) error {
	return e.Element.SetProperty("gs-0", value)
}

// gs-1 (float32)
//
// Band gain Side 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs1(value float32) error {
	return e.Element.SetProperty("gs-1", value)
}

// gs-10 (float32)
//
// Band gain Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs10(value float32) error {
	return e.Element.SetProperty("gs-10", value)
}

// gs-11 (float32)
//
// Band gain Side 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs11(value float32) error {
	return e.Element.SetProperty("gs-11", value)
}

// gs-12 (float32)
//
// Band gain Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs12(value float32) error {
	return e.Element.SetProperty("gs-12", value)
}

// gs-13 (float32)
//
// Band gain Side 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs13(value float32) error {
	return e.Element.SetProperty("gs-13", value)
}

// gs-14 (float32)
//
// Band gain Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs14(value float32) error {
	return e.Element.SetProperty("gs-14", value)
}

// gs-15 (float32)
//
// Band gain Side 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs15(value float32) error {
	return e.Element.SetProperty("gs-15", value)
}

// gs-16 (float32)
//
// Band gain Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs16() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs16(value float32) error {
	return e.Element.SetProperty("gs-16", value)
}

// gs-17 (float32)
//
// Band gain Side 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs17() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs17(value float32) error {
	return e.Element.SetProperty("gs-17", value)
}

// gs-18 (float32)
//
// Band gain Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs18() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs18(value float32) error {
	return e.Element.SetProperty("gs-18", value)
}

// gs-19 (float32)
//
// Band gain Side 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs19() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs19(value float32) error {
	return e.Element.SetProperty("gs-19", value)
}

// gs-2 (float32)
//
// Band gain Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs2(value float32) error {
	return e.Element.SetProperty("gs-2", value)
}

// gs-20 (float32)
//
// Band gain Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs20() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs20(value float32) error {
	return e.Element.SetProperty("gs-20", value)
}

// gs-21 (float32)
//
// Band gain Side 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs21() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs21(value float32) error {
	return e.Element.SetProperty("gs-21", value)
}

// gs-22 (float32)
//
// Band gain Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs22() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs22(value float32) error {
	return e.Element.SetProperty("gs-22", value)
}

// gs-23 (float32)
//
// Band gain Side 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs23() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs23(value float32) error {
	return e.Element.SetProperty("gs-23", value)
}

// gs-24 (float32)
//
// Band gain Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs24() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs24(value float32) error {
	return e.Element.SetProperty("gs-24", value)
}

// gs-25 (float32)
//
// Band gain Side 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs25() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs25(value float32) error {
	return e.Element.SetProperty("gs-25", value)
}

// gs-26 (float32)
//
// Band gain Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs26() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs26(value float32) error {
	return e.Element.SetProperty("gs-26", value)
}

// gs-27 (float32)
//
// Band gain Side 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs27() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs27(value float32) error {
	return e.Element.SetProperty("gs-27", value)
}

// gs-28 (float32)
//
// Band gain Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs28() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs28(value float32) error {
	return e.Element.SetProperty("gs-28", value)
}

// gs-29 (float32)
//
// Band gain Side 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs29() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs29(value float32) error {
	return e.Element.SetProperty("gs-29", value)
}

// gs-3 (float32)
//
// Band gain Side 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs3(value float32) error {
	return e.Element.SetProperty("gs-3", value)
}

// gs-30 (float32)
//
// Band gain Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs30() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs30(value float32) error {
	return e.Element.SetProperty("gs-30", value)
}

// gs-31 (float32)
//
// Band gain Side 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs31() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs31(value float32) error {
	return e.Element.SetProperty("gs-31", value)
}

// gs-4 (float32)
//
// Band gain Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs4(value float32) error {
	return e.Element.SetProperty("gs-4", value)
}

// gs-5 (float32)
//
// Band gain Side 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs5(value float32) error {
	return e.Element.SetProperty("gs-5", value)
}

// gs-6 (float32)
//
// Band gain Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs6(value float32) error {
	return e.Element.SetProperty("gs-6", value)
}

// gs-7 (float32)
//
// Band gain Side 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs7(value float32) error {
	return e.Element.SetProperty("gs-7", value)
}

// gs-8 (float32)
//
// Band gain Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs8(value float32) error {
	return e.Element.SetProperty("gs-8", value)
}

// gs-9 (float32)
//
// Band gain Side 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetGs9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetGs9(value float32) error {
	return e.Element.SetProperty("gs-9", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetIml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetImr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetLstn() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetLstn(value bool) error {
	return e.Element.SetProperty("lstn", value)
}

// mode (GstLspPlugInPluginsLv2GraphEqualizerX32Msmode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// slope (GstLspPlugInPluginsLv2GraphEqualizerX32Msslope)
//
// Filter slope

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetSlope() (interface{}, error) {
	return e.Element.GetProperty("slope")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetSlope(value interface{}) error {
	return e.Element.SetProperty("slope", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetSml() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetSmr() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem0(value bool) error {
	return e.Element.SetProperty("xem-0", value)
}

// xem-1 (bool)
//
// Band on Mid 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem1(value bool) error {
	return e.Element.SetProperty("xem-1", value)
}

// xem-10 (bool)
//
// Band on Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem10(value bool) error {
	return e.Element.SetProperty("xem-10", value)
}

// xem-11 (bool)
//
// Band on Mid 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem11(value bool) error {
	return e.Element.SetProperty("xem-11", value)
}

// xem-12 (bool)
//
// Band on Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem12(value bool) error {
	return e.Element.SetProperty("xem-12", value)
}

// xem-13 (bool)
//
// Band on Mid 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem13(value bool) error {
	return e.Element.SetProperty("xem-13", value)
}

// xem-14 (bool)
//
// Band on Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem14(value bool) error {
	return e.Element.SetProperty("xem-14", value)
}

// xem-15 (bool)
//
// Band on Mid 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem15(value bool) error {
	return e.Element.SetProperty("xem-15", value)
}

// xem-16 (bool)
//
// Band on Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem16(value bool) error {
	return e.Element.SetProperty("xem-16", value)
}

// xem-17 (bool)
//
// Band on Mid 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem17(value bool) error {
	return e.Element.SetProperty("xem-17", value)
}

// xem-18 (bool)
//
// Band on Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem18(value bool) error {
	return e.Element.SetProperty("xem-18", value)
}

// xem-19 (bool)
//
// Band on Mid 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem19(value bool) error {
	return e.Element.SetProperty("xem-19", value)
}

// xem-2 (bool)
//
// Band on Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem2(value bool) error {
	return e.Element.SetProperty("xem-2", value)
}

// xem-20 (bool)
//
// Band on Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem20(value bool) error {
	return e.Element.SetProperty("xem-20", value)
}

// xem-21 (bool)
//
// Band on Mid 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem21(value bool) error {
	return e.Element.SetProperty("xem-21", value)
}

// xem-22 (bool)
//
// Band on Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem22(value bool) error {
	return e.Element.SetProperty("xem-22", value)
}

// xem-23 (bool)
//
// Band on Mid 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem23(value bool) error {
	return e.Element.SetProperty("xem-23", value)
}

// xem-24 (bool)
//
// Band on Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem24(value bool) error {
	return e.Element.SetProperty("xem-24", value)
}

// xem-25 (bool)
//
// Band on Mid 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem25(value bool) error {
	return e.Element.SetProperty("xem-25", value)
}

// xem-26 (bool)
//
// Band on Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem26(value bool) error {
	return e.Element.SetProperty("xem-26", value)
}

// xem-27 (bool)
//
// Band on Mid 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem27(value bool) error {
	return e.Element.SetProperty("xem-27", value)
}

// xem-28 (bool)
//
// Band on Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem28(value bool) error {
	return e.Element.SetProperty("xem-28", value)
}

// xem-29 (bool)
//
// Band on Mid 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem29(value bool) error {
	return e.Element.SetProperty("xem-29", value)
}

// xem-3 (bool)
//
// Band on Mid 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem3(value bool) error {
	return e.Element.SetProperty("xem-3", value)
}

// xem-30 (bool)
//
// Band on Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem30(value bool) error {
	return e.Element.SetProperty("xem-30", value)
}

// xem-31 (bool)
//
// Band on Mid 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xem-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem31(value bool) error {
	return e.Element.SetProperty("xem-31", value)
}

// xem-4 (bool)
//
// Band on Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem4(value bool) error {
	return e.Element.SetProperty("xem-4", value)
}

// xem-5 (bool)
//
// Band on Mid 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem5(value bool) error {
	return e.Element.SetProperty("xem-5", value)
}

// xem-6 (bool)
//
// Band on Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem6(value bool) error {
	return e.Element.SetProperty("xem-6", value)
}

// xem-7 (bool)
//
// Band on Mid 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem7(value bool) error {
	return e.Element.SetProperty("xem-7", value)
}

// xem-8 (bool)
//
// Band on Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem8(value bool) error {
	return e.Element.SetProperty("xem-8", value)
}

// xem-9 (bool)
//
// Band on Mid 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXem9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXem9(value bool) error {
	return e.Element.SetProperty("xem-9", value)
}

// xes-0 (bool)
//
// Band on Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes0(value bool) error {
	return e.Element.SetProperty("xes-0", value)
}

// xes-1 (bool)
//
// Band on Side 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes1(value bool) error {
	return e.Element.SetProperty("xes-1", value)
}

// xes-10 (bool)
//
// Band on Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes10(value bool) error {
	return e.Element.SetProperty("xes-10", value)
}

// xes-11 (bool)
//
// Band on Side 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes11(value bool) error {
	return e.Element.SetProperty("xes-11", value)
}

// xes-12 (bool)
//
// Band on Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes12(value bool) error {
	return e.Element.SetProperty("xes-12", value)
}

// xes-13 (bool)
//
// Band on Side 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes13(value bool) error {
	return e.Element.SetProperty("xes-13", value)
}

// xes-14 (bool)
//
// Band on Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes14(value bool) error {
	return e.Element.SetProperty("xes-14", value)
}

// xes-15 (bool)
//
// Band on Side 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes15(value bool) error {
	return e.Element.SetProperty("xes-15", value)
}

// xes-16 (bool)
//
// Band on Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes16(value bool) error {
	return e.Element.SetProperty("xes-16", value)
}

// xes-17 (bool)
//
// Band on Side 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes17(value bool) error {
	return e.Element.SetProperty("xes-17", value)
}

// xes-18 (bool)
//
// Band on Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes18(value bool) error {
	return e.Element.SetProperty("xes-18", value)
}

// xes-19 (bool)
//
// Band on Side 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes19(value bool) error {
	return e.Element.SetProperty("xes-19", value)
}

// xes-2 (bool)
//
// Band on Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes2(value bool) error {
	return e.Element.SetProperty("xes-2", value)
}

// xes-20 (bool)
//
// Band on Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes20(value bool) error {
	return e.Element.SetProperty("xes-20", value)
}

// xes-21 (bool)
//
// Band on Side 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes21(value bool) error {
	return e.Element.SetProperty("xes-21", value)
}

// xes-22 (bool)
//
// Band on Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes22(value bool) error {
	return e.Element.SetProperty("xes-22", value)
}

// xes-23 (bool)
//
// Band on Side 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes23(value bool) error {
	return e.Element.SetProperty("xes-23", value)
}

// xes-24 (bool)
//
// Band on Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes24(value bool) error {
	return e.Element.SetProperty("xes-24", value)
}

// xes-25 (bool)
//
// Band on Side 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes25(value bool) error {
	return e.Element.SetProperty("xes-25", value)
}

// xes-26 (bool)
//
// Band on Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes26(value bool) error {
	return e.Element.SetProperty("xes-26", value)
}

// xes-27 (bool)
//
// Band on Side 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes27(value bool) error {
	return e.Element.SetProperty("xes-27", value)
}

// xes-28 (bool)
//
// Band on Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes28(value bool) error {
	return e.Element.SetProperty("xes-28", value)
}

// xes-29 (bool)
//
// Band on Side 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes29(value bool) error {
	return e.Element.SetProperty("xes-29", value)
}

// xes-3 (bool)
//
// Band on Side 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes3(value bool) error {
	return e.Element.SetProperty("xes-3", value)
}

// xes-30 (bool)
//
// Band on Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes30(value bool) error {
	return e.Element.SetProperty("xes-30", value)
}

// xes-31 (bool)
//
// Band on Side 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xes-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes31(value bool) error {
	return e.Element.SetProperty("xes-31", value)
}

// xes-4 (bool)
//
// Band on Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes4(value bool) error {
	return e.Element.SetProperty("xes-4", value)
}

// xes-5 (bool)
//
// Band on Side 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes5(value bool) error {
	return e.Element.SetProperty("xes-5", value)
}

// xes-6 (bool)
//
// Band on Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes6(value bool) error {
	return e.Element.SetProperty("xes-6", value)
}

// xes-7 (bool)
//
// Band on Side 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes7(value bool) error {
	return e.Element.SetProperty("xes-7", value)
}

// xes-8 (bool)
//
// Band on Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes8(value bool) error {
	return e.Element.SetProperty("xes-8", value)
}

// xes-9 (bool)
//
// Band on Side 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXes9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXes9(value bool) error {
	return e.Element.SetProperty("xes-9", value)
}

// xmm-0 (bool)
//
// Band mute Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm0(value bool) error {
	return e.Element.SetProperty("xmm-0", value)
}

// xmm-1 (bool)
//
// Band mute Mid 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm1(value bool) error {
	return e.Element.SetProperty("xmm-1", value)
}

// xmm-10 (bool)
//
// Band mute Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm10(value bool) error {
	return e.Element.SetProperty("xmm-10", value)
}

// xmm-11 (bool)
//
// Band mute Mid 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm11(value bool) error {
	return e.Element.SetProperty("xmm-11", value)
}

// xmm-12 (bool)
//
// Band mute Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm12(value bool) error {
	return e.Element.SetProperty("xmm-12", value)
}

// xmm-13 (bool)
//
// Band mute Mid 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm13(value bool) error {
	return e.Element.SetProperty("xmm-13", value)
}

// xmm-14 (bool)
//
// Band mute Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm14(value bool) error {
	return e.Element.SetProperty("xmm-14", value)
}

// xmm-15 (bool)
//
// Band mute Mid 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm15(value bool) error {
	return e.Element.SetProperty("xmm-15", value)
}

// xmm-16 (bool)
//
// Band mute Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm16(value bool) error {
	return e.Element.SetProperty("xmm-16", value)
}

// xmm-17 (bool)
//
// Band mute Mid 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm17(value bool) error {
	return e.Element.SetProperty("xmm-17", value)
}

// xmm-18 (bool)
//
// Band mute Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm18(value bool) error {
	return e.Element.SetProperty("xmm-18", value)
}

// xmm-19 (bool)
//
// Band mute Mid 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm19(value bool) error {
	return e.Element.SetProperty("xmm-19", value)
}

// xmm-2 (bool)
//
// Band mute Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm2(value bool) error {
	return e.Element.SetProperty("xmm-2", value)
}

// xmm-20 (bool)
//
// Band mute Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm20(value bool) error {
	return e.Element.SetProperty("xmm-20", value)
}

// xmm-21 (bool)
//
// Band mute Mid 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm21(value bool) error {
	return e.Element.SetProperty("xmm-21", value)
}

// xmm-22 (bool)
//
// Band mute Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm22(value bool) error {
	return e.Element.SetProperty("xmm-22", value)
}

// xmm-23 (bool)
//
// Band mute Mid 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm23(value bool) error {
	return e.Element.SetProperty("xmm-23", value)
}

// xmm-24 (bool)
//
// Band mute Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm24(value bool) error {
	return e.Element.SetProperty("xmm-24", value)
}

// xmm-25 (bool)
//
// Band mute Mid 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm25(value bool) error {
	return e.Element.SetProperty("xmm-25", value)
}

// xmm-26 (bool)
//
// Band mute Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm26(value bool) error {
	return e.Element.SetProperty("xmm-26", value)
}

// xmm-27 (bool)
//
// Band mute Mid 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm27(value bool) error {
	return e.Element.SetProperty("xmm-27", value)
}

// xmm-28 (bool)
//
// Band mute Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm28(value bool) error {
	return e.Element.SetProperty("xmm-28", value)
}

// xmm-29 (bool)
//
// Band mute Mid 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm29(value bool) error {
	return e.Element.SetProperty("xmm-29", value)
}

// xmm-3 (bool)
//
// Band mute Mid 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm3(value bool) error {
	return e.Element.SetProperty("xmm-3", value)
}

// xmm-30 (bool)
//
// Band mute Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm30(value bool) error {
	return e.Element.SetProperty("xmm-30", value)
}

// xmm-31 (bool)
//
// Band mute Mid 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm31(value bool) error {
	return e.Element.SetProperty("xmm-31", value)
}

// xmm-4 (bool)
//
// Band mute Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm4(value bool) error {
	return e.Element.SetProperty("xmm-4", value)
}

// xmm-5 (bool)
//
// Band mute Mid 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm5(value bool) error {
	return e.Element.SetProperty("xmm-5", value)
}

// xmm-6 (bool)
//
// Band mute Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm6(value bool) error {
	return e.Element.SetProperty("xmm-6", value)
}

// xmm-7 (bool)
//
// Band mute Mid 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm7(value bool) error {
	return e.Element.SetProperty("xmm-7", value)
}

// xmm-8 (bool)
//
// Band mute Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm8(value bool) error {
	return e.Element.SetProperty("xmm-8", value)
}

// xmm-9 (bool)
//
// Band mute Mid 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXmm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXmm9(value bool) error {
	return e.Element.SetProperty("xmm-9", value)
}

// xms-0 (bool)
//
// Band mute Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms0(value bool) error {
	return e.Element.SetProperty("xms-0", value)
}

// xms-1 (bool)
//
// Band mute Side 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms1(value bool) error {
	return e.Element.SetProperty("xms-1", value)
}

// xms-10 (bool)
//
// Band mute Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms10(value bool) error {
	return e.Element.SetProperty("xms-10", value)
}

// xms-11 (bool)
//
// Band mute Side 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms11(value bool) error {
	return e.Element.SetProperty("xms-11", value)
}

// xms-12 (bool)
//
// Band mute Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms12(value bool) error {
	return e.Element.SetProperty("xms-12", value)
}

// xms-13 (bool)
//
// Band mute Side 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms13(value bool) error {
	return e.Element.SetProperty("xms-13", value)
}

// xms-14 (bool)
//
// Band mute Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms14(value bool) error {
	return e.Element.SetProperty("xms-14", value)
}

// xms-15 (bool)
//
// Band mute Side 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms15(value bool) error {
	return e.Element.SetProperty("xms-15", value)
}

// xms-16 (bool)
//
// Band mute Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms16(value bool) error {
	return e.Element.SetProperty("xms-16", value)
}

// xms-17 (bool)
//
// Band mute Side 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms17(value bool) error {
	return e.Element.SetProperty("xms-17", value)
}

// xms-18 (bool)
//
// Band mute Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms18(value bool) error {
	return e.Element.SetProperty("xms-18", value)
}

// xms-19 (bool)
//
// Band mute Side 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms19(value bool) error {
	return e.Element.SetProperty("xms-19", value)
}

// xms-2 (bool)
//
// Band mute Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms2(value bool) error {
	return e.Element.SetProperty("xms-2", value)
}

// xms-20 (bool)
//
// Band mute Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms20(value bool) error {
	return e.Element.SetProperty("xms-20", value)
}

// xms-21 (bool)
//
// Band mute Side 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms21(value bool) error {
	return e.Element.SetProperty("xms-21", value)
}

// xms-22 (bool)
//
// Band mute Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms22(value bool) error {
	return e.Element.SetProperty("xms-22", value)
}

// xms-23 (bool)
//
// Band mute Side 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms23(value bool) error {
	return e.Element.SetProperty("xms-23", value)
}

// xms-24 (bool)
//
// Band mute Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms24(value bool) error {
	return e.Element.SetProperty("xms-24", value)
}

// xms-25 (bool)
//
// Band mute Side 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms25(value bool) error {
	return e.Element.SetProperty("xms-25", value)
}

// xms-26 (bool)
//
// Band mute Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms26(value bool) error {
	return e.Element.SetProperty("xms-26", value)
}

// xms-27 (bool)
//
// Band mute Side 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms27(value bool) error {
	return e.Element.SetProperty("xms-27", value)
}

// xms-28 (bool)
//
// Band mute Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms28(value bool) error {
	return e.Element.SetProperty("xms-28", value)
}

// xms-29 (bool)
//
// Band mute Side 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms29(value bool) error {
	return e.Element.SetProperty("xms-29", value)
}

// xms-3 (bool)
//
// Band mute Side 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms3(value bool) error {
	return e.Element.SetProperty("xms-3", value)
}

// xms-30 (bool)
//
// Band mute Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms30(value bool) error {
	return e.Element.SetProperty("xms-30", value)
}

// xms-31 (bool)
//
// Band mute Side 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms31(value bool) error {
	return e.Element.SetProperty("xms-31", value)
}

// xms-4 (bool)
//
// Band mute Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms4(value bool) error {
	return e.Element.SetProperty("xms-4", value)
}

// xms-5 (bool)
//
// Band mute Side 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms5(value bool) error {
	return e.Element.SetProperty("xms-5", value)
}

// xms-6 (bool)
//
// Band mute Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms6(value bool) error {
	return e.Element.SetProperty("xms-6", value)
}

// xms-7 (bool)
//
// Band mute Side 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms7(value bool) error {
	return e.Element.SetProperty("xms-7", value)
}

// xms-8 (bool)
//
// Band mute Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms8(value bool) error {
	return e.Element.SetProperty("xms-8", value)
}

// xms-9 (bool)
//
// Band mute Side 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXms9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXms9(value bool) error {
	return e.Element.SetProperty("xms-9", value)
}

// xsm-0 (bool)
//
// Band solo Mid 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm0(value bool) error {
	return e.Element.SetProperty("xsm-0", value)
}

// xsm-1 (bool)
//
// Band solo Mid 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm1(value bool) error {
	return e.Element.SetProperty("xsm-1", value)
}

// xsm-10 (bool)
//
// Band solo Mid 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm10(value bool) error {
	return e.Element.SetProperty("xsm-10", value)
}

// xsm-11 (bool)
//
// Band solo Mid 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm11(value bool) error {
	return e.Element.SetProperty("xsm-11", value)
}

// xsm-12 (bool)
//
// Band solo Mid 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm12(value bool) error {
	return e.Element.SetProperty("xsm-12", value)
}

// xsm-13 (bool)
//
// Band solo Mid 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm13(value bool) error {
	return e.Element.SetProperty("xsm-13", value)
}

// xsm-14 (bool)
//
// Band solo Mid 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm14(value bool) error {
	return e.Element.SetProperty("xsm-14", value)
}

// xsm-15 (bool)
//
// Band solo Mid 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm15(value bool) error {
	return e.Element.SetProperty("xsm-15", value)
}

// xsm-16 (bool)
//
// Band solo Mid 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm16(value bool) error {
	return e.Element.SetProperty("xsm-16", value)
}

// xsm-17 (bool)
//
// Band solo Mid 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm17(value bool) error {
	return e.Element.SetProperty("xsm-17", value)
}

// xsm-18 (bool)
//
// Band solo Mid 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm18(value bool) error {
	return e.Element.SetProperty("xsm-18", value)
}

// xsm-19 (bool)
//
// Band solo Mid 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm19(value bool) error {
	return e.Element.SetProperty("xsm-19", value)
}

// xsm-2 (bool)
//
// Band solo Mid 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm2(value bool) error {
	return e.Element.SetProperty("xsm-2", value)
}

// xsm-20 (bool)
//
// Band solo Mid 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm20(value bool) error {
	return e.Element.SetProperty("xsm-20", value)
}

// xsm-21 (bool)
//
// Band solo Mid 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm21(value bool) error {
	return e.Element.SetProperty("xsm-21", value)
}

// xsm-22 (bool)
//
// Band solo Mid 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm22(value bool) error {
	return e.Element.SetProperty("xsm-22", value)
}

// xsm-23 (bool)
//
// Band solo Mid 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm23(value bool) error {
	return e.Element.SetProperty("xsm-23", value)
}

// xsm-24 (bool)
//
// Band solo Mid 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm24(value bool) error {
	return e.Element.SetProperty("xsm-24", value)
}

// xsm-25 (bool)
//
// Band solo Mid 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm25(value bool) error {
	return e.Element.SetProperty("xsm-25", value)
}

// xsm-26 (bool)
//
// Band solo Mid 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm26(value bool) error {
	return e.Element.SetProperty("xsm-26", value)
}

// xsm-27 (bool)
//
// Band solo Mid 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm27(value bool) error {
	return e.Element.SetProperty("xsm-27", value)
}

// xsm-28 (bool)
//
// Band solo Mid 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm28(value bool) error {
	return e.Element.SetProperty("xsm-28", value)
}

// xsm-29 (bool)
//
// Band solo Mid 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm29(value bool) error {
	return e.Element.SetProperty("xsm-29", value)
}

// xsm-3 (bool)
//
// Band solo Mid 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm3(value bool) error {
	return e.Element.SetProperty("xsm-3", value)
}

// xsm-30 (bool)
//
// Band solo Mid 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm30(value bool) error {
	return e.Element.SetProperty("xsm-30", value)
}

// xsm-31 (bool)
//
// Band solo Mid 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm31(value bool) error {
	return e.Element.SetProperty("xsm-31", value)
}

// xsm-4 (bool)
//
// Band solo Mid 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm4(value bool) error {
	return e.Element.SetProperty("xsm-4", value)
}

// xsm-5 (bool)
//
// Band solo Mid 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm5(value bool) error {
	return e.Element.SetProperty("xsm-5", value)
}

// xsm-6 (bool)
//
// Band solo Mid 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm6(value bool) error {
	return e.Element.SetProperty("xsm-6", value)
}

// xsm-7 (bool)
//
// Band solo Mid 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm7(value bool) error {
	return e.Element.SetProperty("xsm-7", value)
}

// xsm-8 (bool)
//
// Band solo Mid 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm8(value bool) error {
	return e.Element.SetProperty("xsm-8", value)
}

// xsm-9 (bool)
//
// Band solo Mid 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXsm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXsm9(value bool) error {
	return e.Element.SetProperty("xsm-9", value)
}

// xss-0 (bool)
//
// Band solo Side 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss0(value bool) error {
	return e.Element.SetProperty("xss-0", value)
}

// xss-1 (bool)
//
// Band solo Side 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss1(value bool) error {
	return e.Element.SetProperty("xss-1", value)
}

// xss-10 (bool)
//
// Band solo Side 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss10(value bool) error {
	return e.Element.SetProperty("xss-10", value)
}

// xss-11 (bool)
//
// Band solo Side 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss11(value bool) error {
	return e.Element.SetProperty("xss-11", value)
}

// xss-12 (bool)
//
// Band solo Side 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss12(value bool) error {
	return e.Element.SetProperty("xss-12", value)
}

// xss-13 (bool)
//
// Band solo Side 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss13(value bool) error {
	return e.Element.SetProperty("xss-13", value)
}

// xss-14 (bool)
//
// Band solo Side 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss14(value bool) error {
	return e.Element.SetProperty("xss-14", value)
}

// xss-15 (bool)
//
// Band solo Side 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss15(value bool) error {
	return e.Element.SetProperty("xss-15", value)
}

// xss-16 (bool)
//
// Band solo Side 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss16() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss16(value bool) error {
	return e.Element.SetProperty("xss-16", value)
}

// xss-17 (bool)
//
// Band solo Side 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss17() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss17(value bool) error {
	return e.Element.SetProperty("xss-17", value)
}

// xss-18 (bool)
//
// Band solo Side 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss18() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss18(value bool) error {
	return e.Element.SetProperty("xss-18", value)
}

// xss-19 (bool)
//
// Band solo Side 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss19() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss19(value bool) error {
	return e.Element.SetProperty("xss-19", value)
}

// xss-2 (bool)
//
// Band solo Side 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss2(value bool) error {
	return e.Element.SetProperty("xss-2", value)
}

// xss-20 (bool)
//
// Band solo Side 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss20() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss20(value bool) error {
	return e.Element.SetProperty("xss-20", value)
}

// xss-21 (bool)
//
// Band solo Side 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss21() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss21(value bool) error {
	return e.Element.SetProperty("xss-21", value)
}

// xss-22 (bool)
//
// Band solo Side 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss22() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss22(value bool) error {
	return e.Element.SetProperty("xss-22", value)
}

// xss-23 (bool)
//
// Band solo Side 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss23() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss23(value bool) error {
	return e.Element.SetProperty("xss-23", value)
}

// xss-24 (bool)
//
// Band solo Side 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss24() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss24(value bool) error {
	return e.Element.SetProperty("xss-24", value)
}

// xss-25 (bool)
//
// Band solo Side 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss25() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss25(value bool) error {
	return e.Element.SetProperty("xss-25", value)
}

// xss-26 (bool)
//
// Band solo Side 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss26() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss26(value bool) error {
	return e.Element.SetProperty("xss-26", value)
}

// xss-27 (bool)
//
// Band solo Side 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss27() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss27(value bool) error {
	return e.Element.SetProperty("xss-27", value)
}

// xss-28 (bool)
//
// Band solo Side 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss28() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss28(value bool) error {
	return e.Element.SetProperty("xss-28", value)
}

// xss-29 (bool)
//
// Band solo Side 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss29() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss29(value bool) error {
	return e.Element.SetProperty("xss-29", value)
}

// xss-3 (bool)
//
// Band solo Side 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss3(value bool) error {
	return e.Element.SetProperty("xss-3", value)
}

// xss-30 (bool)
//
// Band solo Side 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss30() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss30(value bool) error {
	return e.Element.SetProperty("xss-30", value)
}

// xss-31 (bool)
//
// Band solo Side 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss31() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss31(value bool) error {
	return e.Element.SetProperty("xss-31", value)
}

// xss-4 (bool)
//
// Band solo Side 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss4(value bool) error {
	return e.Element.SetProperty("xss-4", value)
}

// xss-5 (bool)
//
// Band solo Side 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss5(value bool) error {
	return e.Element.SetProperty("xss-5", value)
}

// xss-6 (bool)
//
// Band solo Side 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss6(value bool) error {
	return e.Element.SetProperty("xss-6", value)
}

// xss-7 (bool)
//
// Band solo Side 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss7(value bool) error {
	return e.Element.SetProperty("xss-7", value)
}

// xss-8 (bool)
//
// Band solo Side 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss8(value bool) error {
	return e.Element.SetProperty("xss-8", value)
}

// xss-9 (bool)
//
// Band solo Side 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetXss9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetXss9(value bool) error {
	return e.Element.SetProperty("xss-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Ms) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GraphEqualizerX32Msmode string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MsmodeIir LspPlugInPluginsLv2GraphEqualizerX32Msmode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2GraphEqualizerX32MsmodeFir LspPlugInPluginsLv2GraphEqualizerX32Msmode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2GraphEqualizerX32MsmodeFft LspPlugInPluginsLv2GraphEqualizerX32Msmode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2GraphEqualizerX32Msslope string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MsslopeBt48 LspPlugInPluginsLv2GraphEqualizerX32Msslope = "BT48" // BT48 (0) – BT48
	LspPlugInPluginsLv2GraphEqualizerX32MsslopeMt48 LspPlugInPluginsLv2GraphEqualizerX32Msslope = "MT48" // MT48 (1) – MT48
	LspPlugInPluginsLv2GraphEqualizerX32MsslopeBt72 LspPlugInPluginsLv2GraphEqualizerX32Msslope = "BT72" // BT72 (2) – BT72
	LspPlugInPluginsLv2GraphEqualizerX32MsslopeMt72 LspPlugInPluginsLv2GraphEqualizerX32Msslope = "MT72" // MT72 (3) – MT72
	LspPlugInPluginsLv2GraphEqualizerX32MsslopeBt96 LspPlugInPluginsLv2GraphEqualizerX32Msslope = "BT96" // BT96 (4) – BT96
	LspPlugInPluginsLv2GraphEqualizerX32MsslopeMt96 LspPlugInPluginsLv2GraphEqualizerX32Msslope = "MT96" // MT96 (5) – MT96
)

type LspPlugInPluginsLv2GraphEqualizerX32Msfft string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MsfftOff LspPlugInPluginsLv2GraphEqualizerX32Msfft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2GraphEqualizerX32MsfftPostEq LspPlugInPluginsLv2GraphEqualizerX32Msfft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2GraphEqualizerX32MsfftPreEq LspPlugInPluginsLv2GraphEqualizerX32Msfft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2GraphEqualizerX32Msfsel string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MsfselBandsMiddle015 LspPlugInPluginsLv2GraphEqualizerX32Msfsel = "Bands Middle 0-15" // Bands Middle 0-15 (0) – Bands Middle 0-15
	LspPlugInPluginsLv2GraphEqualizerX32MsfselBandsSide015 LspPlugInPluginsLv2GraphEqualizerX32Msfsel = "Bands Side 0-15" // Bands Side 0-15 (1) – Bands Side 0-15
	LspPlugInPluginsLv2GraphEqualizerX32MsfselBandsMiddle1631 LspPlugInPluginsLv2GraphEqualizerX32Msfsel = "Bands Middle 16-31" // Bands Middle 16-31 (2) – Bands Middle 16-31
	LspPlugInPluginsLv2GraphEqualizerX32MsfselBandsSide1631 LspPlugInPluginsLv2GraphEqualizerX32Msfsel = "Bands Side 16-31" // Bands Side 16-31 (3) – Bands Side 16-31
)

