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

type LspPlugInPluginsLv2ImpulseReverbMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ImpulseReverbMono() (*LspPlugInPluginsLv2ImpulseReverbMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-impulse-reverb-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseReverbMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ImpulseReverbMonoWithName(name string) (*LspPlugInPluginsLv2ImpulseReverbMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-impulse-reverb-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseReverbMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// ca0 (bool)
//
// Channel activity 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCa0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// ca1 (bool)
//
// Channel activity 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCa1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// ca2 (bool)
//
// Channel activity 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCa2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// ca3 (bool)
//
// Channel activity 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCa3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// cam0 (bool)
//
// Channel mute 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCam0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cam0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCam0(value bool) error {
	return e.Element.SetProperty("cam0", value)
}

// cam1 (bool)
//
// Channel mute 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCam1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cam1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCam1(value bool) error {
	return e.Element.SetProperty("cam1", value)
}

// cam2 (bool)
//
// Channel mute 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCam2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cam2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCam2(value bool) error {
	return e.Element.SetProperty("cam2", value)
}

// cam3 (bool)
//
// Channel mute 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCam3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cam3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCam3(value bool) error {
	return e.Element.SetProperty("cam3", value)
}

// com0 (float32)
//
// Channel Left/Right output mix 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCom0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("com0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCom0(value float32) error {
	return e.Element.SetProperty("com0", value)
}

// com1 (float32)
//
// Channel Left/Right output mix 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCom1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("com1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCom1(value float32) error {
	return e.Element.SetProperty("com1", value)
}

// com2 (float32)
//
// Channel Left/Right output mix 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCom2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("com2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCom2(value float32) error {
	return e.Element.SetProperty("com2", value)
}

// com3 (float32)
//
// Channel Left/Right output mix 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCom3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("com3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCom3(value float32) error {
	return e.Element.SetProperty("com3", value)
}

// csf0 (GstLspPlugInPluginsLv2ImpulseReverbMonocsf0)
//
// Channel source file 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCsf0() (interface{}, error) {
	return e.Element.GetProperty("csf0")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCsf0(value interface{}) error {
	return e.Element.SetProperty("csf0", value)
}

// csf1 (GstLspPlugInPluginsLv2ImpulseReverbMonocsf1)
//
// Channel source file 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCsf1() (interface{}, error) {
	return e.Element.GetProperty("csf1")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCsf1(value interface{}) error {
	return e.Element.SetProperty("csf1", value)
}

// csf2 (GstLspPlugInPluginsLv2ImpulseReverbMonocsf2)
//
// Channel source file 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCsf2() (interface{}, error) {
	return e.Element.GetProperty("csf2")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCsf2(value interface{}) error {
	return e.Element.SetProperty("csf2", value)
}

// csf3 (GstLspPlugInPluginsLv2ImpulseReverbMonocsf3)
//
// Channel source file 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCsf3() (interface{}, error) {
	return e.Element.GetProperty("csf3")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCsf3(value interface{}) error {
	return e.Element.SetProperty("csf3", value)
}

// cst0 (GstLspPlugInPluginsLv2ImpulseReverbMonocst0)
//
// Channel source track 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCst0() (interface{}, error) {
	return e.Element.GetProperty("cst0")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCst0(value interface{}) error {
	return e.Element.SetProperty("cst0", value)
}

// cst1 (GstLspPlugInPluginsLv2ImpulseReverbMonocst1)
//
// Channel source track 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCst1() (interface{}, error) {
	return e.Element.GetProperty("cst1")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCst1(value interface{}) error {
	return e.Element.SetProperty("cst1", value)
}

// cst2 (GstLspPlugInPluginsLv2ImpulseReverbMonocst2)
//
// Channel source track 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCst2() (interface{}, error) {
	return e.Element.GetProperty("cst2")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCst2(value interface{}) error {
	return e.Element.SetProperty("cst2", value)
}

// cst3 (GstLspPlugInPluginsLv2ImpulseReverbMonocst3)
//
// Channel source track 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetCst3() (interface{}, error) {
	return e.Element.GetProperty("cst3")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetCst3(value interface{}) error {
	return e.Element.SetProperty("cst3", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetDry() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dry")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// eq-0 (float32)
//
// Band 50Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq0(value float32) error {
	return e.Element.SetProperty("eq-0", value)
}

// eq-1 (float32)
//
// Band 107Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq1(value float32) error {
	return e.Element.SetProperty("eq-1", value)
}

// eq-2 (float32)
//
// Band 227Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq2(value float32) error {
	return e.Element.SetProperty("eq-2", value)
}

// eq-3 (float32)
//
// Band 484Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq3(value float32) error {
	return e.Element.SetProperty("eq-3", value)
}

// eq-4 (float32)
//
// Band 1 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq4(value float32) error {
	return e.Element.SetProperty("eq-4", value)
}

// eq-5 (float32)
//
// Band 2.2 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq5(value float32) error {
	return e.Element.SetProperty("eq-5", value)
}

// eq-6 (float32)
//
// Band 4.7 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq6(value float32) error {
	return e.Element.SetProperty("eq-6", value)
}

// eq-7 (float32)
//
// Band 10 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetEq7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("eq-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetEq7(value float32) error {
	return e.Element.SetProperty("eq-7", value)
}

// fft (GstLspPlugInPluginsLv2ImpulseReverbMonofft)
//
// FFT size

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fsel (GstLspPlugInPluginsLv2ImpulseReverbMonofsel)
//
// File selector

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hcf (float32)
//
// High-cut frequency

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetHcf() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hcf")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetHcf(value float32) error {
	return e.Element.SetProperty("hcf", value)
}

// hcm (GstLspPlugInPluginsLv2ImpulseReverbMonohcm)
//
// High-cut mode

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetHcm() (interface{}, error) {
	return e.Element.GetProperty("hcm")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetHcm(value interface{}) error {
	return e.Element.SetProperty("hcm", value)
}

// ifi0 (float32)
//
// Fade in 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfi0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfi0(value float32) error {
	return e.Element.SetProperty("ifi0", value)
}

// ifi1 (float32)
//
// Fade in 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfi1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfi1(value float32) error {
	return e.Element.SetProperty("ifi1", value)
}

// ifi2 (float32)
//
// Fade in 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfi2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfi2(value float32) error {
	return e.Element.SetProperty("ifi2", value)
}

// ifi3 (float32)
//
// Fade in 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfi3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfi3(value float32) error {
	return e.Element.SetProperty("ifi3", value)
}

// ifl0 (float32)
//
// Impulse length 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl1 (float32)
//
// Impulse length 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl2 (float32)
//
// Impulse length 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl3 (float32)
//
// Impulse length 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifo0 (float32)
//
// Fade out 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfo0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfo0(value float32) error {
	return e.Element.SetProperty("ifo0", value)
}

// ifo1 (float32)
//
// Fade out 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfo1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfo1(value float32) error {
	return e.Element.SetProperty("ifo1", value)
}

// ifo2 (float32)
//
// Fade out 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfo2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfo2(value float32) error {
	return e.Element.SetProperty("ifo2", value)
}

// ifo3 (float32)
//
// Fade out 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfo3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIfo3(value float32) error {
	return e.Element.SetProperty("ifo3", value)
}

// ifs0 (int)
//
// Load status 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfs0() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs0")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs1 (int)
//
// Load status 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfs1() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs1")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs2 (int)
//
// Load status 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfs2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs3 (int)
//
// Load status 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIfs3() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs3")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ihc0 (float32)
//
// Head cut 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIhc0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIhc0(value float32) error {
	return e.Element.SetProperty("ihc0", value)
}

// ihc1 (float32)
//
// Head cut 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIhc1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIhc1(value float32) error {
	return e.Element.SetProperty("ihc1", value)
}

// ihc2 (float32)
//
// Head cut 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIhc2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIhc2(value float32) error {
	return e.Element.SetProperty("ihc2", value)
}

// ihc3 (float32)
//
// Head cut 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIhc3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIhc3(value float32) error {
	return e.Element.SetProperty("ihc3", value)
}

// ils0 (bool)
//
// Impulse listen 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIls0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIls0(value bool) error {
	return e.Element.SetProperty("ils0", value)
}

// ils1 (bool)
//
// Impulse listen 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIls1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIls1(value bool) error {
	return e.Element.SetProperty("ils1", value)
}

// ils2 (bool)
//
// Impulse listen 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIls2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIls2(value bool) error {
	return e.Element.SetProperty("ils2", value)
}

// ils3 (bool)
//
// Impulse listen 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIls3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIls3(value bool) error {
	return e.Element.SetProperty("ils3", value)
}

// irv0 (bool)
//
// Impulse reverse 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIrv0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIrv0(value bool) error {
	return e.Element.SetProperty("irv0", value)
}

// irv1 (bool)
//
// Impulse reverse 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIrv1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIrv1(value bool) error {
	return e.Element.SetProperty("irv1", value)
}

// irv2 (bool)
//
// Impulse reverse 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIrv2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIrv2(value bool) error {
	return e.Element.SetProperty("irv2", value)
}

// irv3 (bool)
//
// Impulse reverse 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetIrv3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetIrv3(value bool) error {
	return e.Element.SetProperty("irv3", value)
}

// itc0 (float32)
//
// Tail cut 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetItc0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetItc0(value float32) error {
	return e.Element.SetProperty("itc0", value)
}

// itc1 (float32)
//
// Tail cut 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetItc1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetItc1(value float32) error {
	return e.Element.SetProperty("itc1", value)
}

// itc2 (float32)
//
// Tail cut 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetItc2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetItc2(value float32) error {
	return e.Element.SetProperty("itc2", value)
}

// itc3 (float32)
//
// Tail cut 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetItc3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetItc3(value float32) error {
	return e.Element.SetProperty("itc3", value)
}

// lcf (float32)
//
// Low-cut frequency

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetLcf() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lcf")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetLcf(value float32) error {
	return e.Element.SetProperty("lcf", value)
}

// lcm (GstLspPlugInPluginsLv2ImpulseReverbMonolcm)
//
// Low-cut mode

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetLcm() (interface{}, error) {
	return e.Element.GetProperty("lcm")
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetLcm(value interface{}) error {
	return e.Element.SetProperty("lcm", value)
}

// mk0 (float32)
//
// Makeup gain 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetMk0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetMk0(value float32) error {
	return e.Element.SetProperty("mk0", value)
}

// mk1 (float32)
//
// Makeup gain 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetMk1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetMk1(value float32) error {
	return e.Element.SetProperty("mk1", value)
}

// mk2 (float32)
//
// Makeup gain 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetMk2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetMk2(value float32) error {
	return e.Element.SetProperty("mk2", value)
}

// mk3 (float32)
//
// Makeup gain 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetMk3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetMk3(value float32) error {
	return e.Element.SetProperty("mk3", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetOutLatency() (int, error) {
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

// p (float32)
//
// Panorama

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetP() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetP(value float32) error {
	return e.Element.SetProperty("p", value)
}

// pd (float32)
//
// Pre-delay

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetPd() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetPd(value float32) error {
	return e.Element.SetProperty("pd", value)
}

// pd0 (float32)
//
// Channel pre-delay 0

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetPd0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetPd0(value float32) error {
	return e.Element.SetProperty("pd0", value)
}

// pd1 (float32)
//
// Channel pre-delay 1

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetPd1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetPd1(value float32) error {
	return e.Element.SetProperty("pd1", value)
}

// pd2 (float32)
//
// Channel pre-delay 2

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetPd2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetPd2(value float32) error {
	return e.Element.SetProperty("pd2", value)
}

// pd3 (float32)
//
// Channel pre-delay 3

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetPd3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetPd3(value float32) error {
	return e.Element.SetProperty("pd3", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetWet() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("wet")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wpp (bool)
//
// Wet post-process

func (e *LspPlugInPluginsLv2ImpulseReverbMono) GetWpp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wpp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbMono) SetWpp(value bool) error {
	return e.Element.SetProperty("wpp", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ImpulseReverbMonocsf2 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocsf2None LspPlugInPluginsLv2ImpulseReverbMonocsf2 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbMonocsf2File1 LspPlugInPluginsLv2ImpulseReverbMonocsf2 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbMonocsf2File2 LspPlugInPluginsLv2ImpulseReverbMonocsf2 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbMonocsf2File3 LspPlugInPluginsLv2ImpulseReverbMonocsf2 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbMonocsf2File4 LspPlugInPluginsLv2ImpulseReverbMonocsf2 = "File 4" // File 4 (4) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbMonocst0 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track1 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track2 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track3 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track4 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track5 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track6 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track7 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbMonocst0Track8 LspPlugInPluginsLv2ImpulseReverbMonocst0 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbMonocst1 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track1 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track2 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track3 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track4 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track5 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track6 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track7 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbMonocst1Track8 LspPlugInPluginsLv2ImpulseReverbMonocst1 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbMonocst2 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track1 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track2 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track3 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track4 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track5 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track6 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track7 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbMonocst2Track8 LspPlugInPluginsLv2ImpulseReverbMonocst2 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbMonocst3 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track1 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track2 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track3 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track4 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track5 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track6 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track7 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbMonocst3Track8 LspPlugInPluginsLv2ImpulseReverbMonocst3 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbMonofft string

const (
	LspPlugInPluginsLv2ImpulseReverbMonofft512 LspPlugInPluginsLv2ImpulseReverbMonofft = "512" // 512 (0) – 512
	LspPlugInPluginsLv2ImpulseReverbMonofft1024 LspPlugInPluginsLv2ImpulseReverbMonofft = "1024" // 1024 (1) – 1024
	LspPlugInPluginsLv2ImpulseReverbMonofft2048 LspPlugInPluginsLv2ImpulseReverbMonofft = "2048" // 2048 (2) – 2048
	LspPlugInPluginsLv2ImpulseReverbMonofft4096 LspPlugInPluginsLv2ImpulseReverbMonofft = "4096" // 4096 (3) – 4096
	LspPlugInPluginsLv2ImpulseReverbMonofft8192 LspPlugInPluginsLv2ImpulseReverbMonofft = "8192" // 8192 (4) – 8192
	LspPlugInPluginsLv2ImpulseReverbMonofft16384 LspPlugInPluginsLv2ImpulseReverbMonofft = "16384" // 16384 (5) – 16384
	LspPlugInPluginsLv2ImpulseReverbMonofft32767 LspPlugInPluginsLv2ImpulseReverbMonofft = "32767" // 32767 (6) – 32767
	LspPlugInPluginsLv2ImpulseReverbMonofft65536 LspPlugInPluginsLv2ImpulseReverbMonofft = "65536" // 65536 (7) – 65536
)

type LspPlugInPluginsLv2ImpulseReverbMonofsel string

const (
	LspPlugInPluginsLv2ImpulseReverbMonofselFile1 LspPlugInPluginsLv2ImpulseReverbMonofsel = "File 1" // File 1 (0) – File 1
	LspPlugInPluginsLv2ImpulseReverbMonofselFile2 LspPlugInPluginsLv2ImpulseReverbMonofsel = "File 2" // File 2 (1) – File 2
	LspPlugInPluginsLv2ImpulseReverbMonofselFile3 LspPlugInPluginsLv2ImpulseReverbMonofsel = "File 3" // File 3 (2) – File 3
	LspPlugInPluginsLv2ImpulseReverbMonofselFile4 LspPlugInPluginsLv2ImpulseReverbMonofsel = "File 4" // File 4 (3) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbMonocsf1 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocsf1None LspPlugInPluginsLv2ImpulseReverbMonocsf1 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbMonocsf1File1 LspPlugInPluginsLv2ImpulseReverbMonocsf1 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbMonocsf1File2 LspPlugInPluginsLv2ImpulseReverbMonocsf1 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbMonocsf1File3 LspPlugInPluginsLv2ImpulseReverbMonocsf1 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbMonocsf1File4 LspPlugInPluginsLv2ImpulseReverbMonocsf1 = "File 4" // File 4 (4) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbMonocsf3 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocsf3None LspPlugInPluginsLv2ImpulseReverbMonocsf3 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbMonocsf3File1 LspPlugInPluginsLv2ImpulseReverbMonocsf3 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbMonocsf3File2 LspPlugInPluginsLv2ImpulseReverbMonocsf3 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbMonocsf3File3 LspPlugInPluginsLv2ImpulseReverbMonocsf3 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbMonocsf3File4 LspPlugInPluginsLv2ImpulseReverbMonocsf3 = "File 4" // File 4 (4) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbMonohcm string

const (
	LspPlugInPluginsLv2ImpulseReverbMonohcmOff LspPlugInPluginsLv2ImpulseReverbMonohcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseReverbMonohcm6DBoct LspPlugInPluginsLv2ImpulseReverbMonohcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseReverbMonohcm12DBoct LspPlugInPluginsLv2ImpulseReverbMonohcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseReverbMonohcm18DBoct LspPlugInPluginsLv2ImpulseReverbMonohcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseReverbMonolcm string

const (
	LspPlugInPluginsLv2ImpulseReverbMonolcmOff LspPlugInPluginsLv2ImpulseReverbMonolcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseReverbMonolcm6DBoct LspPlugInPluginsLv2ImpulseReverbMonolcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseReverbMonolcm12DBoct LspPlugInPluginsLv2ImpulseReverbMonolcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseReverbMonolcm18DBoct LspPlugInPluginsLv2ImpulseReverbMonolcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseReverbMonocsf0 string

const (
	LspPlugInPluginsLv2ImpulseReverbMonocsf0None LspPlugInPluginsLv2ImpulseReverbMonocsf0 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbMonocsf0File1 LspPlugInPluginsLv2ImpulseReverbMonocsf0 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbMonocsf0File2 LspPlugInPluginsLv2ImpulseReverbMonocsf0 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbMonocsf0File3 LspPlugInPluginsLv2ImpulseReverbMonocsf0 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbMonocsf0File4 LspPlugInPluginsLv2ImpulseReverbMonocsf0 = "File 4" // File 4 (4) – File 4
)

