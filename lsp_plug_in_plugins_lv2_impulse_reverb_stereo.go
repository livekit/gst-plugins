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

type LspPlugInPluginsLv2ImpulseReverbStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ImpulseReverbStereo() (*LspPlugInPluginsLv2ImpulseReverbStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-impulse-reverb-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseReverbStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ImpulseReverbStereoWithName(name string) (*LspPlugInPluginsLv2ImpulseReverbStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-impulse-reverb-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseReverbStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// ca0 (bool)
//
// Channel activity 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCa0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCa1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCa2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCa3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCam0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCam0(value bool) error {
	return e.Element.SetProperty("cam0", value)
}

// cam1 (bool)
//
// Channel mute 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCam1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCam1(value bool) error {
	return e.Element.SetProperty("cam1", value)
}

// cam2 (bool)
//
// Channel mute 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCam2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCam2(value bool) error {
	return e.Element.SetProperty("cam2", value)
}

// cam3 (bool)
//
// Channel mute 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCam3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCam3(value bool) error {
	return e.Element.SetProperty("cam3", value)
}

// cim0 (float32)
//
// Left/Right input mix 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCim0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cim0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCim0(value float32) error {
	return e.Element.SetProperty("cim0", value)
}

// cim1 (float32)
//
// Left/Right input mix 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCim1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cim1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCim1(value float32) error {
	return e.Element.SetProperty("cim1", value)
}

// cim2 (float32)
//
// Left/Right input mix 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCim2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cim2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCim2(value float32) error {
	return e.Element.SetProperty("cim2", value)
}

// cim3 (float32)
//
// Left/Right input mix 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCim3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cim3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCim3(value float32) error {
	return e.Element.SetProperty("cim3", value)
}

// com0 (float32)
//
// Channel Left/Right output mix 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCom0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCom0(value float32) error {
	return e.Element.SetProperty("com0", value)
}

// com1 (float32)
//
// Channel Left/Right output mix 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCom1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCom1(value float32) error {
	return e.Element.SetProperty("com1", value)
}

// com2 (float32)
//
// Channel Left/Right output mix 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCom2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCom2(value float32) error {
	return e.Element.SetProperty("com2", value)
}

// com3 (float32)
//
// Channel Left/Right output mix 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCom3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCom3(value float32) error {
	return e.Element.SetProperty("com3", value)
}

// csf0 (GstLspPlugInPluginsLv2ImpulseReverbStereocsf0)
//
// Channel source file 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCsf0() (interface{}, error) {
	return e.Element.GetProperty("csf0")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCsf0(value interface{}) error {
	return e.Element.SetProperty("csf0", value)
}

// csf1 (GstLspPlugInPluginsLv2ImpulseReverbStereocsf1)
//
// Channel source file 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCsf1() (interface{}, error) {
	return e.Element.GetProperty("csf1")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCsf1(value interface{}) error {
	return e.Element.SetProperty("csf1", value)
}

// csf2 (GstLspPlugInPluginsLv2ImpulseReverbStereocsf2)
//
// Channel source file 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCsf2() (interface{}, error) {
	return e.Element.GetProperty("csf2")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCsf2(value interface{}) error {
	return e.Element.SetProperty("csf2", value)
}

// csf3 (GstLspPlugInPluginsLv2ImpulseReverbStereocsf3)
//
// Channel source file 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCsf3() (interface{}, error) {
	return e.Element.GetProperty("csf3")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCsf3(value interface{}) error {
	return e.Element.SetProperty("csf3", value)
}

// cst0 (GstLspPlugInPluginsLv2ImpulseReverbStereocst0)
//
// Channel source track 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCst0() (interface{}, error) {
	return e.Element.GetProperty("cst0")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCst0(value interface{}) error {
	return e.Element.SetProperty("cst0", value)
}

// cst1 (GstLspPlugInPluginsLv2ImpulseReverbStereocst1)
//
// Channel source track 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCst1() (interface{}, error) {
	return e.Element.GetProperty("cst1")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCst1(value interface{}) error {
	return e.Element.SetProperty("cst1", value)
}

// cst2 (GstLspPlugInPluginsLv2ImpulseReverbStereocst2)
//
// Channel source track 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCst2() (interface{}, error) {
	return e.Element.GetProperty("cst2")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCst2(value interface{}) error {
	return e.Element.SetProperty("cst2", value)
}

// cst3 (GstLspPlugInPluginsLv2ImpulseReverbStereocst3)
//
// Channel source track 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetCst3() (interface{}, error) {
	return e.Element.GetProperty("cst3")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetCst3(value interface{}) error {
	return e.Element.SetProperty("cst3", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// eq-0 (float32)
//
// Band 50Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq0(value float32) error {
	return e.Element.SetProperty("eq-0", value)
}

// eq-1 (float32)
//
// Band 107Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq1(value float32) error {
	return e.Element.SetProperty("eq-1", value)
}

// eq-2 (float32)
//
// Band 227Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq2(value float32) error {
	return e.Element.SetProperty("eq-2", value)
}

// eq-3 (float32)
//
// Band 484Hz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq3(value float32) error {
	return e.Element.SetProperty("eq-3", value)
}

// eq-4 (float32)
//
// Band 1 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq4(value float32) error {
	return e.Element.SetProperty("eq-4", value)
}

// eq-5 (float32)
//
// Band 2.2 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq5(value float32) error {
	return e.Element.SetProperty("eq-5", value)
}

// eq-6 (float32)
//
// Band 4.7 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq6(value float32) error {
	return e.Element.SetProperty("eq-6", value)
}

// eq-7 (float32)
//
// Band 10 kHz gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetEq7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetEq7(value float32) error {
	return e.Element.SetProperty("eq-7", value)
}

// fft (GstLspPlugInPluginsLv2ImpulseReverbStereofft)
//
// FFT size

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fsel (GstLspPlugInPluginsLv2ImpulseReverbStereofsel)
//
// File selector

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hcf (float32)
//
// High-cut frequency

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetHcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetHcf(value float32) error {
	return e.Element.SetProperty("hcf", value)
}

// hcm (GstLspPlugInPluginsLv2ImpulseReverbStereohcm)
//
// High-cut mode

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetHcm() (interface{}, error) {
	return e.Element.GetProperty("hcm")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetHcm(value interface{}) error {
	return e.Element.SetProperty("hcm", value)
}

// ifi0 (float32)
//
// Fade in 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfi0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfi0(value float32) error {
	return e.Element.SetProperty("ifi0", value)
}

// ifi1 (float32)
//
// Fade in 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfi1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfi1(value float32) error {
	return e.Element.SetProperty("ifi1", value)
}

// ifi2 (float32)
//
// Fade in 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfi2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfi2(value float32) error {
	return e.Element.SetProperty("ifi2", value)
}

// ifi3 (float32)
//
// Fade in 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfi3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfi3(value float32) error {
	return e.Element.SetProperty("ifi3", value)
}

// ifl0 (float32)
//
// Impulse length 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfo0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfo0(value float32) error {
	return e.Element.SetProperty("ifo0", value)
}

// ifo1 (float32)
//
// Fade out 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfo1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfo1(value float32) error {
	return e.Element.SetProperty("ifo1", value)
}

// ifo2 (float32)
//
// Fade out 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfo2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfo2(value float32) error {
	return e.Element.SetProperty("ifo2", value)
}

// ifo3 (float32)
//
// Fade out 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfo3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIfo3(value float32) error {
	return e.Element.SetProperty("ifo3", value)
}

// ifs0 (int)
//
// Load status 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfs0() (int, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfs1() (int, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfs2() (int, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIfs3() (int, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIhc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIhc0(value float32) error {
	return e.Element.SetProperty("ihc0", value)
}

// ihc1 (float32)
//
// Head cut 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIhc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIhc1(value float32) error {
	return e.Element.SetProperty("ihc1", value)
}

// ihc2 (float32)
//
// Head cut 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIhc2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIhc2(value float32) error {
	return e.Element.SetProperty("ihc2", value)
}

// ihc3 (float32)
//
// Head cut 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIhc3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIhc3(value float32) error {
	return e.Element.SetProperty("ihc3", value)
}

// ils0 (bool)
//
// Impulse listen 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIls0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIls0(value bool) error {
	return e.Element.SetProperty("ils0", value)
}

// ils1 (bool)
//
// Impulse listen 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIls1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIls1(value bool) error {
	return e.Element.SetProperty("ils1", value)
}

// ils2 (bool)
//
// Impulse listen 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIls2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIls2(value bool) error {
	return e.Element.SetProperty("ils2", value)
}

// ils3 (bool)
//
// Impulse listen 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIls3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIls3(value bool) error {
	return e.Element.SetProperty("ils3", value)
}

// irv0 (bool)
//
// Impulse reverse 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIrv0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIrv0(value bool) error {
	return e.Element.SetProperty("irv0", value)
}

// irv1 (bool)
//
// Impulse reverse 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIrv1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIrv1(value bool) error {
	return e.Element.SetProperty("irv1", value)
}

// irv2 (bool)
//
// Impulse reverse 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIrv2() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIrv2(value bool) error {
	return e.Element.SetProperty("irv2", value)
}

// irv3 (bool)
//
// Impulse reverse 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetIrv3() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetIrv3(value bool) error {
	return e.Element.SetProperty("irv3", value)
}

// itc0 (float32)
//
// Tail cut 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetItc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetItc0(value float32) error {
	return e.Element.SetProperty("itc0", value)
}

// itc1 (float32)
//
// Tail cut 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetItc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetItc1(value float32) error {
	return e.Element.SetProperty("itc1", value)
}

// itc2 (float32)
//
// Tail cut 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetItc2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetItc2(value float32) error {
	return e.Element.SetProperty("itc2", value)
}

// itc3 (float32)
//
// Tail cut 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetItc3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetItc3(value float32) error {
	return e.Element.SetProperty("itc3", value)
}

// lcf (float32)
//
// Low-cut frequency

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetLcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetLcf(value float32) error {
	return e.Element.SetProperty("lcf", value)
}

// lcm (GstLspPlugInPluginsLv2ImpulseReverbStereolcm)
//
// Low-cut mode

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetLcm() (interface{}, error) {
	return e.Element.GetProperty("lcm")
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetLcm(value interface{}) error {
	return e.Element.SetProperty("lcm", value)
}

// mk0 (float32)
//
// Makeup gain 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetMk0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetMk0(value float32) error {
	return e.Element.SetProperty("mk0", value)
}

// mk1 (float32)
//
// Makeup gain 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetMk1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetMk1(value float32) error {
	return e.Element.SetProperty("mk1", value)
}

// mk2 (float32)
//
// Makeup gain 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetMk2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetMk2(value float32) error {
	return e.Element.SetProperty("mk2", value)
}

// mk3 (float32)
//
// Makeup gain 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetMk3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetMk3(value float32) error {
	return e.Element.SetProperty("mk3", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetOutLatency() (int, error) {
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

// pd (float32)
//
// Pre-delay

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPd() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPd(value float32) error {
	return e.Element.SetProperty("pd", value)
}

// pd0 (float32)
//
// Channel pre-delay 0

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPd0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPd0(value float32) error {
	return e.Element.SetProperty("pd0", value)
}

// pd1 (float32)
//
// Channel pre-delay 1

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPd1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPd1(value float32) error {
	return e.Element.SetProperty("pd1", value)
}

// pd2 (float32)
//
// Channel pre-delay 2

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPd2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPd2(value float32) error {
	return e.Element.SetProperty("pd2", value)
}

// pd3 (float32)
//
// Channel pre-delay 3

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPd3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPd3(value float32) error {
	return e.Element.SetProperty("pd3", value)
}

// pl (float32)
//
// Left channel panorama

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPl(value float32) error {
	return e.Element.SetProperty("pl", value)
}

// pr (float32)
//
// Right channel panorama

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetPr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetPr(value float32) error {
	return e.Element.SetProperty("pr", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wpp (bool)
//
// Wet post-process

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) GetWpp() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseReverbStereo) SetWpp(value bool) error {
	return e.Element.SetProperty("wpp", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ImpulseReverbStereocsf3 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocsf3None LspPlugInPluginsLv2ImpulseReverbStereocsf3 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbStereocsf3File1 LspPlugInPluginsLv2ImpulseReverbStereocsf3 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbStereocsf3File2 LspPlugInPluginsLv2ImpulseReverbStereocsf3 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbStereocsf3File3 LspPlugInPluginsLv2ImpulseReverbStereocsf3 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbStereocsf3File4 LspPlugInPluginsLv2ImpulseReverbStereocsf3 = "File 4" // File 4 (4) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbStereocst0 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track1 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track2 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track3 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track4 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track5 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track6 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track7 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbStereocst0Track8 LspPlugInPluginsLv2ImpulseReverbStereocst0 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbStereocst2 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track1 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track2 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track3 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track4 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track5 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track6 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track7 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbStereocst2Track8 LspPlugInPluginsLv2ImpulseReverbStereocst2 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbStereofft string

const (
	LspPlugInPluginsLv2ImpulseReverbStereofft512 LspPlugInPluginsLv2ImpulseReverbStereofft = "512" // 512 (0) – 512
	LspPlugInPluginsLv2ImpulseReverbStereofft1024 LspPlugInPluginsLv2ImpulseReverbStereofft = "1024" // 1024 (1) – 1024
	LspPlugInPluginsLv2ImpulseReverbStereofft2048 LspPlugInPluginsLv2ImpulseReverbStereofft = "2048" // 2048 (2) – 2048
	LspPlugInPluginsLv2ImpulseReverbStereofft4096 LspPlugInPluginsLv2ImpulseReverbStereofft = "4096" // 4096 (3) – 4096
	LspPlugInPluginsLv2ImpulseReverbStereofft8192 LspPlugInPluginsLv2ImpulseReverbStereofft = "8192" // 8192 (4) – 8192
	LspPlugInPluginsLv2ImpulseReverbStereofft16384 LspPlugInPluginsLv2ImpulseReverbStereofft = "16384" // 16384 (5) – 16384
	LspPlugInPluginsLv2ImpulseReverbStereofft32767 LspPlugInPluginsLv2ImpulseReverbStereofft = "32767" // 32767 (6) – 32767
	LspPlugInPluginsLv2ImpulseReverbStereofft65536 LspPlugInPluginsLv2ImpulseReverbStereofft = "65536" // 65536 (7) – 65536
)

type LspPlugInPluginsLv2ImpulseReverbStereocsf1 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocsf1None LspPlugInPluginsLv2ImpulseReverbStereocsf1 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbStereocsf1File1 LspPlugInPluginsLv2ImpulseReverbStereocsf1 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbStereocsf1File2 LspPlugInPluginsLv2ImpulseReverbStereocsf1 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbStereocsf1File3 LspPlugInPluginsLv2ImpulseReverbStereocsf1 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbStereocsf1File4 LspPlugInPluginsLv2ImpulseReverbStereocsf1 = "File 4" // File 4 (4) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbStereocsf2 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocsf2None LspPlugInPluginsLv2ImpulseReverbStereocsf2 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbStereocsf2File1 LspPlugInPluginsLv2ImpulseReverbStereocsf2 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbStereocsf2File2 LspPlugInPluginsLv2ImpulseReverbStereocsf2 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbStereocsf2File3 LspPlugInPluginsLv2ImpulseReverbStereocsf2 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbStereocsf2File4 LspPlugInPluginsLv2ImpulseReverbStereocsf2 = "File 4" // File 4 (4) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbStereocst1 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track1 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track2 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track3 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track4 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track5 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track6 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track7 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbStereocst1Track8 LspPlugInPluginsLv2ImpulseReverbStereocst1 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbStereocst3 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track1 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 1" // Track 1 (0) – Track 1
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track2 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 2" // Track 2 (1) – Track 2
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track3 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 3" // Track 3 (2) – Track 3
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track4 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 4" // Track 4 (3) – Track 4
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track5 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 5" // Track 5 (4) – Track 5
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track6 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 6" // Track 6 (5) – Track 6
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track7 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 7" // Track 7 (6) – Track 7
	LspPlugInPluginsLv2ImpulseReverbStereocst3Track8 LspPlugInPluginsLv2ImpulseReverbStereocst3 = "Track 8" // Track 8 (7) – Track 8
)

type LspPlugInPluginsLv2ImpulseReverbStereofsel string

const (
	LspPlugInPluginsLv2ImpulseReverbStereofselFile1 LspPlugInPluginsLv2ImpulseReverbStereofsel = "File 1" // File 1 (0) – File 1
	LspPlugInPluginsLv2ImpulseReverbStereofselFile2 LspPlugInPluginsLv2ImpulseReverbStereofsel = "File 2" // File 2 (1) – File 2
	LspPlugInPluginsLv2ImpulseReverbStereofselFile3 LspPlugInPluginsLv2ImpulseReverbStereofsel = "File 3" // File 3 (2) – File 3
	LspPlugInPluginsLv2ImpulseReverbStereofselFile4 LspPlugInPluginsLv2ImpulseReverbStereofsel = "File 4" // File 4 (3) – File 4
)

type LspPlugInPluginsLv2ImpulseReverbStereohcm string

const (
	LspPlugInPluginsLv2ImpulseReverbStereohcmOff LspPlugInPluginsLv2ImpulseReverbStereohcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseReverbStereohcm6DBoct LspPlugInPluginsLv2ImpulseReverbStereohcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseReverbStereohcm12DBoct LspPlugInPluginsLv2ImpulseReverbStereohcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseReverbStereohcm18DBoct LspPlugInPluginsLv2ImpulseReverbStereohcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseReverbStereolcm string

const (
	LspPlugInPluginsLv2ImpulseReverbStereolcmOff LspPlugInPluginsLv2ImpulseReverbStereolcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseReverbStereolcm6DBoct LspPlugInPluginsLv2ImpulseReverbStereolcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseReverbStereolcm12DBoct LspPlugInPluginsLv2ImpulseReverbStereolcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseReverbStereolcm18DBoct LspPlugInPluginsLv2ImpulseReverbStereolcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseReverbStereocsf0 string

const (
	LspPlugInPluginsLv2ImpulseReverbStereocsf0None LspPlugInPluginsLv2ImpulseReverbStereocsf0 = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseReverbStereocsf0File1 LspPlugInPluginsLv2ImpulseReverbStereocsf0 = "File 1" // File 1 (1) – File 1
	LspPlugInPluginsLv2ImpulseReverbStereocsf0File2 LspPlugInPluginsLv2ImpulseReverbStereocsf0 = "File 2" // File 2 (2) – File 2
	LspPlugInPluginsLv2ImpulseReverbStereocsf0File3 LspPlugInPluginsLv2ImpulseReverbStereocsf0 = "File 3" // File 3 (3) – File 3
	LspPlugInPluginsLv2ImpulseReverbStereocsf0File4 LspPlugInPluginsLv2ImpulseReverbStereocsf0 = "File 4" // File 4 (4) – File 4
)

