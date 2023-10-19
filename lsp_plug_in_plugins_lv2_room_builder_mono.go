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

type LspPlugInPluginsLv2RoomBuilderMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2RoomBuilderMono() (*LspPlugInPluginsLv2RoomBuilderMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-room-builder-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2RoomBuilderMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2RoomBuilderMonoWithName(name string) (*LspPlugInPluginsLv2RoomBuilderMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-room-builder-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2RoomBuilderMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// ca0 (bool)
//
// Channel activity 0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCa0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCa1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCa2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCa3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCam0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCam0(value bool) error {
	return e.Element.SetProperty("cam0", value)
}

// cam1 (bool)
//
// Channel mute 1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCam1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCam1(value bool) error {
	return e.Element.SetProperty("cam1", value)
}

// cam2 (bool)
//
// Channel mute 2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCam2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCam2(value bool) error {
	return e.Element.SetProperty("cam2", value)
}

// cam3 (bool)
//
// Channel mute 3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCam3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCam3(value bool) error {
	return e.Element.SetProperty("cam3", value)
}

// com0 (float32)
//
// Channel Left/Right output mix 0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCom0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCom0(value float32) error {
	return e.Element.SetProperty("com0", value)
}

// com1 (float32)
//
// Channel Left/Right output mix 1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCom1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCom1(value float32) error {
	return e.Element.SetProperty("com1", value)
}

// com2 (float32)
//
// Channel Left/Right output mix 2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCom2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCom2(value float32) error {
	return e.Element.SetProperty("com2", value)
}

// com3 (float32)
//
// Channel Left/Right output mix 3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCom3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCom3(value float32) error {
	return e.Element.SetProperty("com3", value)
}

// cpitch (float32)
//
// Camera Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCpitch() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cpitch")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCpitch(value float32) error {
	return e.Element.SetProperty("cpitch", value)
}

// cposx (float32)
//
// Camera X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCposx() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cposx")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCposx(value float32) error {
	return e.Element.SetProperty("cposx", value)
}

// cposy (float32)
//
// Camera Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCposy() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cposy")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCposy(value float32) error {
	return e.Element.SetProperty("cposy", value)
}

// cposz (float32)
//
// Camera Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCposz() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cposz")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCposz(value float32) error {
	return e.Element.SetProperty("cposz", value)
}

// csel (GstLspPlugInPluginsLv2RoomBuilderMonocsel)
//
// Capture selector

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCsel() (interface{}, error) {
	return e.Element.GetProperty("csel")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCsel(value interface{}) error {
	return e.Element.SetProperty("csel", value)
}

// csf0 (GstLspPlugInPluginsLv2RoomBuilderMonocsf0)
//
// Channel source sample 0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCsf0() (interface{}, error) {
	return e.Element.GetProperty("csf0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCsf0(value interface{}) error {
	return e.Element.SetProperty("csf0", value)
}

// csf1 (GstLspPlugInPluginsLv2RoomBuilderMonocsf1)
//
// Channel source sample 1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCsf1() (interface{}, error) {
	return e.Element.GetProperty("csf1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCsf1(value interface{}) error {
	return e.Element.SetProperty("csf1", value)
}

// csf2 (GstLspPlugInPluginsLv2RoomBuilderMonocsf2)
//
// Channel source sample 2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCsf2() (interface{}, error) {
	return e.Element.GetProperty("csf2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCsf2(value interface{}) error {
	return e.Element.SetProperty("csf2", value)
}

// csf3 (GstLspPlugInPluginsLv2RoomBuilderMonocsf3)
//
// Channel source sample 3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCsf3() (interface{}, error) {
	return e.Element.GetProperty("csf3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCsf3(value interface{}) error {
	return e.Element.SetProperty("csf3", value)
}

// cst0 (GstLspPlugInPluginsLv2RoomBuilderMonocst0)
//
// Channel source track 0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCst0() (interface{}, error) {
	return e.Element.GetProperty("cst0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCst0(value interface{}) error {
	return e.Element.SetProperty("cst0", value)
}

// cst1 (GstLspPlugInPluginsLv2RoomBuilderMonocst1)
//
// Channel source track 1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCst1() (interface{}, error) {
	return e.Element.GetProperty("cst1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCst1(value interface{}) error {
	return e.Element.SetProperty("cst1", value)
}

// cst2 (GstLspPlugInPluginsLv2RoomBuilderMonocst2)
//
// Channel source track 2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCst2() (interface{}, error) {
	return e.Element.GetProperty("cst2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCst2(value interface{}) error {
	return e.Element.SetProperty("cst2", value)
}

// cst3 (GstLspPlugInPluginsLv2RoomBuilderMonocst3)
//
// Channel source track 3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCst3() (interface{}, error) {
	return e.Element.GetProperty("cst3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCst3(value interface{}) error {
	return e.Element.SetProperty("cst3", value)
}

// cyaw (float32)
//
// Camera Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetCyaw() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cyaw")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetCyaw(value float32) error {
	return e.Element.SetProperty("cyaw", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// editor (GstLspPlugInPluginsLv2RoomBuilderMonoeditor)
//
// Current editor

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEditor() (interface{}, error) {
	return e.Element.GetProperty("editor")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEditor(value interface{}) error {
	return e.Element.SetProperty("editor", value)
}

// eq-0 (float32)
//
// Band 50Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq0(value float32) error {
	return e.Element.SetProperty("eq-0", value)
}

// eq-1 (float32)
//
// Band 107Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq1(value float32) error {
	return e.Element.SetProperty("eq-1", value)
}

// eq-2 (float32)
//
// Band 227Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq2(value float32) error {
	return e.Element.SetProperty("eq-2", value)
}

// eq-3 (float32)
//
// Band 484Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq3(value float32) error {
	return e.Element.SetProperty("eq-3", value)
}

// eq-4 (float32)
//
// Band 1 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq4(value float32) error {
	return e.Element.SetProperty("eq-4", value)
}

// eq-5 (float32)
//
// Band 2.2 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq5(value float32) error {
	return e.Element.SetProperty("eq-5", value)
}

// eq-6 (float32)
//
// Band 4.7 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq6(value float32) error {
	return e.Element.SetProperty("eq-6", value)
}

// eq-7 (float32)
//
// Band 10 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetEq7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetEq7(value float32) error {
	return e.Element.SetProperty("eq-7", value)
}

// fft (GstLspPlugInPluginsLv2RoomBuilderMonofft)
//
// FFT size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hcf (float32)
//
// High-cut frequency

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetHcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetHcf(value float32) error {
	return e.Element.SetProperty("hcf", value)
}

// hcm (GstLspPlugInPluginsLv2RoomBuilderMonohcm)
//
// High-cut mode

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetHcm() (interface{}, error) {
	return e.Element.GetProperty("hcm")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetHcm(value interface{}) error {
	return e.Element.SetProperty("hcm", value)
}

// ifi-0 (float32)
//
// Fade in0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi0(value float32) error {
	return e.Element.SetProperty("ifi-0", value)
}

// ifi-1 (float32)
//
// Fade in1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi1(value float32) error {
	return e.Element.SetProperty("ifi-1", value)
}

// ifi-2 (float32)
//
// Fade in2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi2(value float32) error {
	return e.Element.SetProperty("ifi-2", value)
}

// ifi-3 (float32)
//
// Fade in3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi3(value float32) error {
	return e.Element.SetProperty("ifi-3", value)
}

// ifi-4 (float32)
//
// Fade in4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi4(value float32) error {
	return e.Element.SetProperty("ifi-4", value)
}

// ifi-5 (float32)
//
// Fade in5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi5(value float32) error {
	return e.Element.SetProperty("ifi-5", value)
}

// ifi-6 (float32)
//
// Fade in6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi6(value float32) error {
	return e.Element.SetProperty("ifi-6", value)
}

// ifi-7 (float32)
//
// Fade in7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfi7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfi7(value float32) error {
	return e.Element.SetProperty("ifi-7", value)
}

// ifl-0 (float32)
//
// Impulse length0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-1 (float32)
//
// Impulse length1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-2 (float32)
//
// Impulse length2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-3 (float32)
//
// Impulse length3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-4 (float32)
//
// Impulse length4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-5 (float32)
//
// Impulse length5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-6 (float32)
//
// Impulse length6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifl-7 (float32)
//
// Impulse length7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifo (GstLspPlugInPluginsLv2RoomBuilderMonoifo)
//
// Input 3D model orientation

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo() (interface{}, error) {
	return e.Element.GetProperty("ifo")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo(value interface{}) error {
	return e.Element.SetProperty("ifo", value)
}

// ifo-0 (float32)
//
// Fade out0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo0(value float32) error {
	return e.Element.SetProperty("ifo-0", value)
}

// ifo-1 (float32)
//
// Fade out1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo1(value float32) error {
	return e.Element.SetProperty("ifo-1", value)
}

// ifo-2 (float32)
//
// Fade out2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo2(value float32) error {
	return e.Element.SetProperty("ifo-2", value)
}

// ifo-3 (float32)
//
// Fade out3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo3(value float32) error {
	return e.Element.SetProperty("ifo-3", value)
}

// ifo-4 (float32)
//
// Fade out4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo4(value float32) error {
	return e.Element.SetProperty("ifo-4", value)
}

// ifo-5 (float32)
//
// Fade out5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo5(value float32) error {
	return e.Element.SetProperty("ifo-5", value)
}

// ifo-6 (float32)
//
// Fade out6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo6(value float32) error {
	return e.Element.SetProperty("ifo-6", value)
}

// ifo-7 (float32)
//
// Fade out7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfo7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIfo7(value float32) error {
	return e.Element.SetProperty("ifo-7", value)
}

// ifp (float32)
//
// File loading progress

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifs (int)
//
// Input 3D model load status

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-0 (int)
//
// Impulse status0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs0() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-1 (int)
//
// Impulse status1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs1() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-2 (int)
//
// Impulse status2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-3 (int)
//
// Impulse status3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs3() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-4 (int)
//
// Impulse status4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs4() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-5 (int)
//
// Impulse status5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs5() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-6 (int)
//
// Impulse status6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs6() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ifs-7 (int)
//
// Impulse status7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIfs7() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ifs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ihc-0 (float32)
//
// Head cut0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc0(value float32) error {
	return e.Element.SetProperty("ihc-0", value)
}

// ihc-1 (float32)
//
// Head cut1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc1(value float32) error {
	return e.Element.SetProperty("ihc-1", value)
}

// ihc-2 (float32)
//
// Head cut2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc2(value float32) error {
	return e.Element.SetProperty("ihc-2", value)
}

// ihc-3 (float32)
//
// Head cut3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc3(value float32) error {
	return e.Element.SetProperty("ihc-3", value)
}

// ihc-4 (float32)
//
// Head cut4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc4(value float32) error {
	return e.Element.SetProperty("ihc-4", value)
}

// ihc-5 (float32)
//
// Head cut5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc5(value float32) error {
	return e.Element.SetProperty("ihc-5", value)
}

// ihc-6 (float32)
//
// Head cut6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc6(value float32) error {
	return e.Element.SetProperty("ihc-6", value)
}

// ihc-7 (float32)
//
// Head cut7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIhc7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIhc7(value float32) error {
	return e.Element.SetProperty("ihc-7", value)
}

// ils-0 (bool)
//
// Impulse listen0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls0(value bool) error {
	return e.Element.SetProperty("ils-0", value)
}

// ils-1 (bool)
//
// Impulse listen1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls1(value bool) error {
	return e.Element.SetProperty("ils-1", value)
}

// ils-2 (bool)
//
// Impulse listen2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls2(value bool) error {
	return e.Element.SetProperty("ils-2", value)
}

// ils-3 (bool)
//
// Impulse listen3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls3(value bool) error {
	return e.Element.SetProperty("ils-3", value)
}

// ils-4 (bool)
//
// Impulse listen4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls4(value bool) error {
	return e.Element.SetProperty("ils-4", value)
}

// ils-5 (bool)
//
// Impulse listen5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls5(value bool) error {
	return e.Element.SetProperty("ils-5", value)
}

// ils-6 (bool)
//
// Impulse listen6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls6(value bool) error {
	return e.Element.SetProperty("ils-6", value)
}

// ils-7 (bool)
//
// Impulse listen7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIls7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIls7(value bool) error {
	return e.Element.SetProperty("ils-7", value)
}

// imkp-0 (float32)
//
// Impulse makeup gain0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp0(value float32) error {
	return e.Element.SetProperty("imkp-0", value)
}

// imkp-1 (float32)
//
// Impulse makeup gain1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp1(value float32) error {
	return e.Element.SetProperty("imkp-1", value)
}

// imkp-2 (float32)
//
// Impulse makeup gain2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp2(value float32) error {
	return e.Element.SetProperty("imkp-2", value)
}

// imkp-3 (float32)
//
// Impulse makeup gain3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp3(value float32) error {
	return e.Element.SetProperty("imkp-3", value)
}

// imkp-4 (float32)
//
// Impulse makeup gain4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp4(value float32) error {
	return e.Element.SetProperty("imkp-4", value)
}

// imkp-5 (float32)
//
// Impulse makeup gain5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp5(value float32) error {
	return e.Element.SetProperty("imkp-5", value)
}

// imkp-6 (float32)
//
// Impulse makeup gain6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp6(value float32) error {
	return e.Element.SetProperty("imkp-6", value)
}

// imkp-7 (float32)
//
// Impulse makeup gain7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetImkp7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imkp-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetImkp7(value float32) error {
	return e.Element.SetProperty("imkp-7", value)
}

// irv-0 (bool)
//
// Impulse reverse0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv0(value bool) error {
	return e.Element.SetProperty("irv-0", value)
}

// irv-1 (bool)
//
// Impulse reverse1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv1(value bool) error {
	return e.Element.SetProperty("irv-1", value)
}

// irv-2 (bool)
//
// Impulse reverse2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv2(value bool) error {
	return e.Element.SetProperty("irv-2", value)
}

// irv-3 (bool)
//
// Impulse reverse3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv3(value bool) error {
	return e.Element.SetProperty("irv-3", value)
}

// irv-4 (bool)
//
// Impulse reverse4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv4(value bool) error {
	return e.Element.SetProperty("irv-4", value)
}

// irv-5 (bool)
//
// Impulse reverse5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv5(value bool) error {
	return e.Element.SetProperty("irv-5", value)
}

// irv-6 (bool)
//
// Impulse reverse6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv6(value bool) error {
	return e.Element.SetProperty("irv-6", value)
}

// irv-7 (bool)
//
// Impulse reverse7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetIrv7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irv-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetIrv7(value bool) error {
	return e.Element.SetProperty("irv-7", value)
}

// itc-0 (float32)
//
// Tail cut0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc0(value float32) error {
	return e.Element.SetProperty("itc-0", value)
}

// itc-1 (float32)
//
// Tail cut1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc1(value float32) error {
	return e.Element.SetProperty("itc-1", value)
}

// itc-2 (float32)
//
// Tail cut2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc2(value float32) error {
	return e.Element.SetProperty("itc-2", value)
}

// itc-3 (float32)
//
// Tail cut3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc3(value float32) error {
	return e.Element.SetProperty("itc-3", value)
}

// itc-4 (float32)
//
// Tail cut4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc4(value float32) error {
	return e.Element.SetProperty("itc-4", value)
}

// itc-5 (float32)
//
// Tail cut5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc5(value float32) error {
	return e.Element.SetProperty("itc-5", value)
}

// itc-6 (float32)
//
// Tail cut6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc6(value float32) error {
	return e.Element.SetProperty("itc-6", value)
}

// itc-7 (float32)
//
// Tail cut7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetItc7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetItc7(value float32) error {
	return e.Element.SetProperty("itc-7", value)
}

// lcf (float32)
//
// Low-cut frequency

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetLcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetLcf(value float32) error {
	return e.Element.SetProperty("lcf", value)
}

// lcm (GstLspPlugInPluginsLv2RoomBuilderMonolcm)
//
// Low-cut mode

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetLcm() (interface{}, error) {
	return e.Element.GetProperty("lcm")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetLcm(value interface{}) error {
	return e.Element.SetProperty("lcm", value)
}

// mk0 (float32)
//
// Makeup gain 0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetMk0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetMk0(value float32) error {
	return e.Element.SetProperty("mk0", value)
}

// mk1 (float32)
//
// Makeup gain 1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetMk1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetMk1(value float32) error {
	return e.Element.SetProperty("mk1", value)
}

// mk2 (float32)
//
// Makeup gain 2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetMk2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetMk2(value float32) error {
	return e.Element.SetProperty("mk2", value)
}

// mk3 (float32)
//
// Makeup gain 3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetMk3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetMk3(value float32) error {
	return e.Element.SetProperty("mk3", value)
}

// normal (bool)
//
// Normalize rendered samples

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetNormal() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("normal")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetNormal(value bool) error {
	return e.Element.SetProperty("normal", value)
}

// ofc-0 (bool)
//
// Sample save command0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc0(value bool) error {
	return e.Element.SetProperty("ofc-0", value)
}

// ofc-1 (bool)
//
// Sample save command1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc1(value bool) error {
	return e.Element.SetProperty("ofc-1", value)
}

// ofc-2 (bool)
//
// Sample save command2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc2(value bool) error {
	return e.Element.SetProperty("ofc-2", value)
}

// ofc-3 (bool)
//
// Sample save command3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc3(value bool) error {
	return e.Element.SetProperty("ofc-3", value)
}

// ofc-4 (bool)
//
// Sample save command4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc4(value bool) error {
	return e.Element.SetProperty("ofc-4", value)
}

// ofc-5 (bool)
//
// Sample save command5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc5(value bool) error {
	return e.Element.SetProperty("ofc-5", value)
}

// ofc-6 (bool)
//
// Sample save command6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc6(value bool) error {
	return e.Element.SetProperty("ofc-6", value)
}

// ofc-7 (bool)
//
// Sample save command7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfc7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofc-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetOfc7(value bool) error {
	return e.Element.SetProperty("ofc-7", value)
}

// ofp-0 (float32)
//
// Sample saving progress0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-1 (float32)
//
// Sample saving progress1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-2 (float32)
//
// Sample saving progress2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-3 (float32)
//
// Sample saving progress3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-4 (float32)
//
// Sample saving progress4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-5 (float32)
//
// Sample saving progress5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-6 (float32)
//
// Sample saving progress6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofp-7 (float32)
//
// Sample saving progress7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfp7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ofp-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ofs-0 (int)
//
// Sample saving status0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs0() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-1 (int)
//
// Sample saving status1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs1() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-2 (int)
//
// Sample saving status2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-3 (int)
//
// Sample saving status3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs3() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-4 (int)
//
// Sample saving status4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs4() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-5 (int)
//
// Sample saving status5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs5() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-6 (int)
//
// Sample saving status6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs6() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ofs-7 (int)
//
// Sample saving status7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOfs7() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ofs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetP() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetP(value float32) error {
	return e.Element.SetProperty("p", value)
}

// pd (float32)
//
// Pre-delay

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetPd() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetPd(value float32) error {
	return e.Element.SetProperty("pd", value)
}

// pd0 (float32)
//
// Channel pre-delay 0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetPd0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetPd0(value float32) error {
	return e.Element.SetProperty("pd0", value)
}

// pd1 (float32)
//
// Channel pre-delay 1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetPd1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetPd1(value float32) error {
	return e.Element.SetProperty("pd1", value)
}

// pd2 (float32)
//
// Channel pre-delay 2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetPd2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetPd2(value float32) error {
	return e.Element.SetProperty("pd2", value)
}

// pd3 (float32)
//
// Channel pre-delay 3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetPd3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetPd3(value float32) error {
	return e.Element.SetProperty("pd3", value)
}

// prog (float32)
//
// Rendering progress

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetProg() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("prog")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// quality (float32)
//
// Quality factor

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

// render (bool)
//
// Launch/Stop rendering process

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetRender() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("render")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetRender(value bool) error {
	return e.Element.SetProperty("render", value)
}

// sca-0 (float32)
//
// Capture 0 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca0(value float32) error {
	return e.Element.SetProperty("sca-0", value)
}

// sca-1 (float32)
//
// Capture 1 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca1(value float32) error {
	return e.Element.SetProperty("sca-1", value)
}

// sca-2 (float32)
//
// Capture 2 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca2(value float32) error {
	return e.Element.SetProperty("sca-2", value)
}

// sca-3 (float32)
//
// Capture 3 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca3(value float32) error {
	return e.Element.SetProperty("sca-3", value)
}

// sca-4 (float32)
//
// Capture 4 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca4(value float32) error {
	return e.Element.SetProperty("sca-4", value)
}

// sca-5 (float32)
//
// Capture 5 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca5(value float32) error {
	return e.Element.SetProperty("sca-5", value)
}

// sca-6 (float32)
//
// Capture 6 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca6(value float32) error {
	return e.Element.SetProperty("sca-6", value)
}

// sca-7 (float32)
//
// Capture 7 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSca7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sca-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSca7(value float32) error {
	return e.Element.SetProperty("sca-7", value)
}

// scab-0 (float32)
//
// Capture 0 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab0(value float32) error {
	return e.Element.SetProperty("scab-0", value)
}

// scab-1 (float32)
//
// Capture 1 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab1(value float32) error {
	return e.Element.SetProperty("scab-1", value)
}

// scab-2 (float32)
//
// Capture 2 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab2(value float32) error {
	return e.Element.SetProperty("scab-2", value)
}

// scab-3 (float32)
//
// Capture 3 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab3(value float32) error {
	return e.Element.SetProperty("scab-3", value)
}

// scab-4 (float32)
//
// Capture 4 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab4(value float32) error {
	return e.Element.SetProperty("scab-4", value)
}

// scab-5 (float32)
//
// Capture 5 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab5(value float32) error {
	return e.Element.SetProperty("scab-5", value)
}

// scab-6 (float32)
//
// Capture 6 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab6(value float32) error {
	return e.Element.SetProperty("scab-6", value)
}

// scab-7 (float32)
//
// Capture 7 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScab7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scab-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScab7(value float32) error {
	return e.Element.SetProperty("scab-7", value)
}

// scap-0 (float32)
//
// Capture 0 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap0(value float32) error {
	return e.Element.SetProperty("scap-0", value)
}

// scap-1 (float32)
//
// Capture 1 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap1(value float32) error {
	return e.Element.SetProperty("scap-1", value)
}

// scap-2 (float32)
//
// Capture 2 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap2(value float32) error {
	return e.Element.SetProperty("scap-2", value)
}

// scap-3 (float32)
//
// Capture 3 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap3(value float32) error {
	return e.Element.SetProperty("scap-3", value)
}

// scap-4 (float32)
//
// Capture 4 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap4(value float32) error {
	return e.Element.SetProperty("scap-4", value)
}

// scap-5 (float32)
//
// Capture 5 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap5(value float32) error {
	return e.Element.SetProperty("scap-5", value)
}

// scap-6 (float32)
//
// Capture 6 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap6(value float32) error {
	return e.Element.SetProperty("scap-6", value)
}

// scap-7 (float32)
//
// Capture 7 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScap7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scap-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScap7(value float32) error {
	return e.Element.SetProperty("scap-7", value)
}

// scar-0 (float32)
//
// Capture 0 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar0(value float32) error {
	return e.Element.SetProperty("scar-0", value)
}

// scar-1 (float32)
//
// Capture 1 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar1(value float32) error {
	return e.Element.SetProperty("scar-1", value)
}

// scar-2 (float32)
//
// Capture 2 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar2(value float32) error {
	return e.Element.SetProperty("scar-2", value)
}

// scar-3 (float32)
//
// Capture 3 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar3(value float32) error {
	return e.Element.SetProperty("scar-3", value)
}

// scar-4 (float32)
//
// Capture 4 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar4(value float32) error {
	return e.Element.SetProperty("scar-4", value)
}

// scar-5 (float32)
//
// Capture 5 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar5(value float32) error {
	return e.Element.SetProperty("scar-5", value)
}

// scar-6 (float32)
//
// Capture 6 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar6(value float32) error {
	return e.Element.SetProperty("scar-6", value)
}

// scar-7 (float32)
//
// Capture 7 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScar7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scar-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScar7(value float32) error {
	return e.Element.SetProperty("scar-7", value)
}

// scay-0 (float32)
//
// Capture 0 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay0(value float32) error {
	return e.Element.SetProperty("scay-0", value)
}

// scay-1 (float32)
//
// Capture 1 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay1(value float32) error {
	return e.Element.SetProperty("scay-1", value)
}

// scay-2 (float32)
//
// Capture 2 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay2(value float32) error {
	return e.Element.SetProperty("scay-2", value)
}

// scay-3 (float32)
//
// Capture 3 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay3(value float32) error {
	return e.Element.SetProperty("scay-3", value)
}

// scay-4 (float32)
//
// Capture 4 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay4(value float32) error {
	return e.Element.SetProperty("scay-4", value)
}

// scay-5 (float32)
//
// Capture 5 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay5(value float32) error {
	return e.Element.SetProperty("scay-5", value)
}

// scay-6 (float32)
//
// Capture 6 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay6(value float32) error {
	return e.Element.SetProperty("scay-6", value)
}

// scay-7 (float32)
//
// Capture 7 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScay7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scay-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScay7(value float32) error {
	return e.Element.SetProperty("scay-7", value)
}

// sccf-0 (GstLspPlugInPluginsLv2RoomBuilderMonosccf0)
//
// Capture 0 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf0() (interface{}, error) {
	return e.Element.GetProperty("sccf-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf0(value interface{}) error {
	return e.Element.SetProperty("sccf-0", value)
}

// sccf-1 (GstLspPlugInPluginsLv2RoomBuilderMonosccf1)
//
// Capture 1 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf1() (interface{}, error) {
	return e.Element.GetProperty("sccf-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf1(value interface{}) error {
	return e.Element.SetProperty("sccf-1", value)
}

// sccf-2 (GstLspPlugInPluginsLv2RoomBuilderMonosccf2)
//
// Capture 2 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf2() (interface{}, error) {
	return e.Element.GetProperty("sccf-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf2(value interface{}) error {
	return e.Element.SetProperty("sccf-2", value)
}

// sccf-3 (GstLspPlugInPluginsLv2RoomBuilderMonosccf3)
//
// Capture 3 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf3() (interface{}, error) {
	return e.Element.GetProperty("sccf-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf3(value interface{}) error {
	return e.Element.SetProperty("sccf-3", value)
}

// sccf-4 (GstLspPlugInPluginsLv2RoomBuilderMonosccf4)
//
// Capture 4 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf4() (interface{}, error) {
	return e.Element.GetProperty("sccf-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf4(value interface{}) error {
	return e.Element.SetProperty("sccf-4", value)
}

// sccf-5 (GstLspPlugInPluginsLv2RoomBuilderMonosccf5)
//
// Capture 5 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf5() (interface{}, error) {
	return e.Element.GetProperty("sccf-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf5(value interface{}) error {
	return e.Element.SetProperty("sccf-5", value)
}

// sccf-6 (GstLspPlugInPluginsLv2RoomBuilderMonosccf6)
//
// Capture 6 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf6() (interface{}, error) {
	return e.Element.GetProperty("sccf-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf6(value interface{}) error {
	return e.Element.SetProperty("sccf-6", value)
}

// sccf-7 (GstLspPlugInPluginsLv2RoomBuilderMonosccf7)
//
// Capture 7 configuration

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccf7() (interface{}, error) {
	return e.Element.GetProperty("sccf-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccf7(value interface{}) error {
	return e.Element.SetProperty("sccf-7", value)
}

// sccs-0 (float32)
//
// Capture 0 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs0(value float32) error {
	return e.Element.SetProperty("sccs-0", value)
}

// sccs-1 (float32)
//
// Capture 1 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs1(value float32) error {
	return e.Element.SetProperty("sccs-1", value)
}

// sccs-2 (float32)
//
// Capture 2 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs2(value float32) error {
	return e.Element.SetProperty("sccs-2", value)
}

// sccs-3 (float32)
//
// Capture 3 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs3(value float32) error {
	return e.Element.SetProperty("sccs-3", value)
}

// sccs-4 (float32)
//
// Capture 4 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs4(value float32) error {
	return e.Element.SetProperty("sccs-4", value)
}

// sccs-5 (float32)
//
// Capture 5 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs5(value float32) error {
	return e.Element.SetProperty("sccs-5", value)
}

// sccs-6 (float32)
//
// Capture 6 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs6(value float32) error {
	return e.Element.SetProperty("sccs-6", value)
}

// sccs-7 (float32)
//
// Capture 7 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSccs7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sccs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSccs7(value float32) error {
	return e.Element.SetProperty("sccs-7", value)
}

// sce-0 (bool)
//
// Capture 0 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce0(value bool) error {
	return e.Element.SetProperty("sce-0", value)
}

// sce-1 (bool)
//
// Capture 1 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce1(value bool) error {
	return e.Element.SetProperty("sce-1", value)
}

// sce-2 (bool)
//
// Capture 2 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce2(value bool) error {
	return e.Element.SetProperty("sce-2", value)
}

// sce-3 (bool)
//
// Capture 3 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce3(value bool) error {
	return e.Element.SetProperty("sce-3", value)
}

// sce-4 (bool)
//
// Capture 4 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce4(value bool) error {
	return e.Element.SetProperty("sce-4", value)
}

// sce-5 (bool)
//
// Capture 5 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce5(value bool) error {
	return e.Element.SetProperty("sce-5", value)
}

// sce-6 (bool)
//
// Capture 6 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce6(value bool) error {
	return e.Element.SetProperty("sce-6", value)
}

// sce-7 (bool)
//
// Capture 7 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSce7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sce-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSce7(value bool) error {
	return e.Element.SetProperty("sce-7", value)
}

// scf-0 (GstLspPlugInPluginsLv2RoomBuilderMonoscf0)
//
// Capture 0 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf0() (interface{}, error) {
	return e.Element.GetProperty("scf-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf0(value interface{}) error {
	return e.Element.SetProperty("scf-0", value)
}

// scf-1 (GstLspPlugInPluginsLv2RoomBuilderMonoscf1)
//
// Capture 1 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf1() (interface{}, error) {
	return e.Element.GetProperty("scf-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf1(value interface{}) error {
	return e.Element.SetProperty("scf-1", value)
}

// scf-2 (GstLspPlugInPluginsLv2RoomBuilderMonoscf2)
//
// Capture 2 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf2() (interface{}, error) {
	return e.Element.GetProperty("scf-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf2(value interface{}) error {
	return e.Element.SetProperty("scf-2", value)
}

// scf-3 (GstLspPlugInPluginsLv2RoomBuilderMonoscf3)
//
// Capture 3 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf3() (interface{}, error) {
	return e.Element.GetProperty("scf-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf3(value interface{}) error {
	return e.Element.SetProperty("scf-3", value)
}

// scf-4 (GstLspPlugInPluginsLv2RoomBuilderMonoscf4)
//
// Capture 4 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf4() (interface{}, error) {
	return e.Element.GetProperty("scf-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf4(value interface{}) error {
	return e.Element.SetProperty("scf-4", value)
}

// scf-5 (GstLspPlugInPluginsLv2RoomBuilderMonoscf5)
//
// Capture 5 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf5() (interface{}, error) {
	return e.Element.GetProperty("scf-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf5(value interface{}) error {
	return e.Element.SetProperty("scf-5", value)
}

// scf-6 (GstLspPlugInPluginsLv2RoomBuilderMonoscf6)
//
// Capture 6 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf6() (interface{}, error) {
	return e.Element.GetProperty("scf-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf6(value interface{}) error {
	return e.Element.SetProperty("scf-6", value)
}

// scf-7 (GstLspPlugInPluginsLv2RoomBuilderMonoscf7)
//
// Capture 7 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScf7() (interface{}, error) {
	return e.Element.GetProperty("scf-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScf7(value interface{}) error {
	return e.Element.SetProperty("scf-7", value)
}

// sch-0 (float32)
//
// Capture 0 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch0(value float32) error {
	return e.Element.SetProperty("sch-0", value)
}

// sch-1 (float32)
//
// Capture 1 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch1(value float32) error {
	return e.Element.SetProperty("sch-1", value)
}

// sch-2 (float32)
//
// Capture 2 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch2(value float32) error {
	return e.Element.SetProperty("sch-2", value)
}

// sch-3 (float32)
//
// Capture 3 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch3(value float32) error {
	return e.Element.SetProperty("sch-3", value)
}

// sch-4 (float32)
//
// Capture 4 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch4(value float32) error {
	return e.Element.SetProperty("sch-4", value)
}

// sch-5 (float32)
//
// Capture 5 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch5(value float32) error {
	return e.Element.SetProperty("sch-5", value)
}

// sch-6 (float32)
//
// Capture 6 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch6(value float32) error {
	return e.Element.SetProperty("sch-6", value)
}

// sch-7 (float32)
//
// Capture 7 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSch7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sch-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSch7(value float32) error {
	return e.Element.SetProperty("sch-7", value)
}

// scl-0 (GstLspPlugInPluginsLv2RoomBuilderMonoscl0)
//
// Capture 0 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl0() (interface{}, error) {
	return e.Element.GetProperty("scl-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl0(value interface{}) error {
	return e.Element.SetProperty("scl-0", value)
}

// scl-1 (GstLspPlugInPluginsLv2RoomBuilderMonoscl1)
//
// Capture 1 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl1() (interface{}, error) {
	return e.Element.GetProperty("scl-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl1(value interface{}) error {
	return e.Element.SetProperty("scl-1", value)
}

// scl-2 (GstLspPlugInPluginsLv2RoomBuilderMonoscl2)
//
// Capture 2 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl2() (interface{}, error) {
	return e.Element.GetProperty("scl-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl2(value interface{}) error {
	return e.Element.SetProperty("scl-2", value)
}

// scl-3 (GstLspPlugInPluginsLv2RoomBuilderMonoscl3)
//
// Capture 3 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl3() (interface{}, error) {
	return e.Element.GetProperty("scl-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl3(value interface{}) error {
	return e.Element.SetProperty("scl-3", value)
}

// scl-4 (GstLspPlugInPluginsLv2RoomBuilderMonoscl4)
//
// Capture 4 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl4() (interface{}, error) {
	return e.Element.GetProperty("scl-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl4(value interface{}) error {
	return e.Element.SetProperty("scl-4", value)
}

// scl-5 (GstLspPlugInPluginsLv2RoomBuilderMonoscl5)
//
// Capture 5 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl5() (interface{}, error) {
	return e.Element.GetProperty("scl-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl5(value interface{}) error {
	return e.Element.SetProperty("scl-5", value)
}

// scl-6 (GstLspPlugInPluginsLv2RoomBuilderMonoscl6)
//
// Capture 6 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl6() (interface{}, error) {
	return e.Element.GetProperty("scl-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl6(value interface{}) error {
	return e.Element.SetProperty("scl-6", value)
}

// scl-7 (GstLspPlugInPluginsLv2RoomBuilderMonoscl7)
//
// Capture 7 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScl7() (interface{}, error) {
	return e.Element.GetProperty("scl-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScl7(value interface{}) error {
	return e.Element.SetProperty("scl-7", value)
}

// scmd-0 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd0)
//
// Capture 0 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd0() (interface{}, error) {
	return e.Element.GetProperty("scmd-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd0(value interface{}) error {
	return e.Element.SetProperty("scmd-0", value)
}

// scmd-1 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd1)
//
// Capture 1 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd1() (interface{}, error) {
	return e.Element.GetProperty("scmd-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd1(value interface{}) error {
	return e.Element.SetProperty("scmd-1", value)
}

// scmd-2 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd2)
//
// Capture 2 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd2() (interface{}, error) {
	return e.Element.GetProperty("scmd-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd2(value interface{}) error {
	return e.Element.SetProperty("scmd-2", value)
}

// scmd-3 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd3)
//
// Capture 3 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd3() (interface{}, error) {
	return e.Element.GetProperty("scmd-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd3(value interface{}) error {
	return e.Element.SetProperty("scmd-3", value)
}

// scmd-4 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd4)
//
// Capture 4 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd4() (interface{}, error) {
	return e.Element.GetProperty("scmd-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd4(value interface{}) error {
	return e.Element.SetProperty("scmd-4", value)
}

// scmd-5 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd5)
//
// Capture 5 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd5() (interface{}, error) {
	return e.Element.GetProperty("scmd-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd5(value interface{}) error {
	return e.Element.SetProperty("scmd-5", value)
}

// scmd-6 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd6)
//
// Capture 6 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd6() (interface{}, error) {
	return e.Element.GetProperty("scmd-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd6(value interface{}) error {
	return e.Element.SetProperty("scmd-6", value)
}

// scmd-7 (GstLspPlugInPluginsLv2RoomBuilderMonoscmd7)
//
// Capture 7 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScmd7() (interface{}, error) {
	return e.Element.GetProperty("scmd-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScmd7(value interface{}) error {
	return e.Element.SetProperty("scmd-7", value)
}

// scpx-0 (float32)
//
// Capture 0 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx0(value float32) error {
	return e.Element.SetProperty("scpx-0", value)
}

// scpx-1 (float32)
//
// Capture 1 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx1(value float32) error {
	return e.Element.SetProperty("scpx-1", value)
}

// scpx-2 (float32)
//
// Capture 2 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx2(value float32) error {
	return e.Element.SetProperty("scpx-2", value)
}

// scpx-3 (float32)
//
// Capture 3 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx3(value float32) error {
	return e.Element.SetProperty("scpx-3", value)
}

// scpx-4 (float32)
//
// Capture 4 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx4(value float32) error {
	return e.Element.SetProperty("scpx-4", value)
}

// scpx-5 (float32)
//
// Capture 5 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx5(value float32) error {
	return e.Element.SetProperty("scpx-5", value)
}

// scpx-6 (float32)
//
// Capture 6 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx6(value float32) error {
	return e.Element.SetProperty("scpx-6", value)
}

// scpx-7 (float32)
//
// Capture 7 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpx7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpx-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpx7(value float32) error {
	return e.Element.SetProperty("scpx-7", value)
}

// scpy-0 (float32)
//
// Capture 0 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy0(value float32) error {
	return e.Element.SetProperty("scpy-0", value)
}

// scpy-1 (float32)
//
// Capture 1 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy1(value float32) error {
	return e.Element.SetProperty("scpy-1", value)
}

// scpy-2 (float32)
//
// Capture 2 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy2(value float32) error {
	return e.Element.SetProperty("scpy-2", value)
}

// scpy-3 (float32)
//
// Capture 3 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy3(value float32) error {
	return e.Element.SetProperty("scpy-3", value)
}

// scpy-4 (float32)
//
// Capture 4 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy4(value float32) error {
	return e.Element.SetProperty("scpy-4", value)
}

// scpy-5 (float32)
//
// Capture 5 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy5(value float32) error {
	return e.Element.SetProperty("scpy-5", value)
}

// scpy-6 (float32)
//
// Capture 6 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy6(value float32) error {
	return e.Element.SetProperty("scpy-6", value)
}

// scpy-7 (float32)
//
// Capture 7 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpy7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpy-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpy7(value float32) error {
	return e.Element.SetProperty("scpy-7", value)
}

// scpz-0 (float32)
//
// Capture 0 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz0(value float32) error {
	return e.Element.SetProperty("scpz-0", value)
}

// scpz-1 (float32)
//
// Capture 1 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz1(value float32) error {
	return e.Element.SetProperty("scpz-1", value)
}

// scpz-2 (float32)
//
// Capture 2 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz2(value float32) error {
	return e.Element.SetProperty("scpz-2", value)
}

// scpz-3 (float32)
//
// Capture 3 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz3(value float32) error {
	return e.Element.SetProperty("scpz-3", value)
}

// scpz-4 (float32)
//
// Capture 4 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz4(value float32) error {
	return e.Element.SetProperty("scpz-4", value)
}

// scpz-5 (float32)
//
// Capture 5 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz5(value float32) error {
	return e.Element.SetProperty("scpz-5", value)
}

// scpz-6 (float32)
//
// Capture 6 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz6(value float32) error {
	return e.Element.SetProperty("scpz-6", value)
}

// scpz-7 (float32)
//
// Capture 7 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScpz7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scpz-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScpz7(value float32) error {
	return e.Element.SetProperty("scpz-7", value)
}

// scsd-0 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd0)
//
// Capture 0 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd0() (interface{}, error) {
	return e.Element.GetProperty("scsd-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd0(value interface{}) error {
	return e.Element.SetProperty("scsd-0", value)
}

// scsd-1 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd1)
//
// Capture 1 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd1() (interface{}, error) {
	return e.Element.GetProperty("scsd-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd1(value interface{}) error {
	return e.Element.SetProperty("scsd-1", value)
}

// scsd-2 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd2)
//
// Capture 2 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd2() (interface{}, error) {
	return e.Element.GetProperty("scsd-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd2(value interface{}) error {
	return e.Element.SetProperty("scsd-2", value)
}

// scsd-3 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd3)
//
// Capture 3 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd3() (interface{}, error) {
	return e.Element.GetProperty("scsd-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd3(value interface{}) error {
	return e.Element.SetProperty("scsd-3", value)
}

// scsd-4 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd4)
//
// Capture 4 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd4() (interface{}, error) {
	return e.Element.GetProperty("scsd-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd4(value interface{}) error {
	return e.Element.SetProperty("scsd-4", value)
}

// scsd-5 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd5)
//
// Capture 5 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd5() (interface{}, error) {
	return e.Element.GetProperty("scsd-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd5(value interface{}) error {
	return e.Element.SetProperty("scsd-5", value)
}

// scsd-6 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd6)
//
// Capture 6 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd6() (interface{}, error) {
	return e.Element.GetProperty("scsd-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd6(value interface{}) error {
	return e.Element.SetProperty("scsd-6", value)
}

// scsd-7 (GstLspPlugInPluginsLv2RoomBuilderMonoscsd7)
//
// Capture 7 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetScsd7() (interface{}, error) {
	return e.Element.GetProperty("scsd-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetScsd7(value interface{}) error {
	return e.Element.SetProperty("scsd-7", value)
}

// sdur-0 (float32)
//
// Impulse current duration0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-1 (float32)
//
// Impulse current duration1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-2 (float32)
//
// Impulse current duration2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-3 (float32)
//
// Impulse current duration3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-4 (float32)
//
// Impulse current duration4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-5 (float32)
//
// Impulse current duration5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-6 (float32)
//
// Impulse current duration6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sdur-7 (float32)
//
// Impulse current duration7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSdur7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sdur-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// shh-0 (float32)
//
// Source 0 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh0(value float32) error {
	return e.Element.SetProperty("shh-0", value)
}

// shh-1 (float32)
//
// Source 1 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh1(value float32) error {
	return e.Element.SetProperty("shh-1", value)
}

// shh-2 (float32)
//
// Source 2 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh2(value float32) error {
	return e.Element.SetProperty("shh-2", value)
}

// shh-3 (float32)
//
// Source 3 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh3(value float32) error {
	return e.Element.SetProperty("shh-3", value)
}

// shh-4 (float32)
//
// Source 4 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh4(value float32) error {
	return e.Element.SetProperty("shh-4", value)
}

// shh-5 (float32)
//
// Source 5 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh5(value float32) error {
	return e.Element.SetProperty("shh-5", value)
}

// shh-6 (float32)
//
// Source 6 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh6(value float32) error {
	return e.Element.SetProperty("shh-6", value)
}

// shh-7 (float32)
//
// Source 7 height

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetShh7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shh-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetShh7(value float32) error {
	return e.Element.SetProperty("shh-7", value)
}

// signal (GstLspPlugInPluginsLv2RoomBuilderMonosignal)
//
// Current signal processor

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSignal() (interface{}, error) {
	return e.Element.GetProperty("signal")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSignal(value interface{}) error {
	return e.Element.SetProperty("signal", value)
}

// smdur-0 (float32)
//
// Impulse max duration0

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-1 (float32)
//
// Impulse max duration1

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-2 (float32)
//
// Impulse max duration2

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-3 (float32)
//
// Impulse max duration3

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-4 (float32)
//
// Impulse max duration4

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-5 (float32)
//
// Impulse max duration5

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-6 (float32)
//
// Impulse max duration6

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smdur-7 (float32)
//
// Impulse max duration7

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSmdur7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smdur-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ssa-0 (float32)
//
// Source 0 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa0(value float32) error {
	return e.Element.SetProperty("ssa-0", value)
}

// ssa-1 (float32)
//
// Source 1 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa1(value float32) error {
	return e.Element.SetProperty("ssa-1", value)
}

// ssa-2 (float32)
//
// Source 2 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa2(value float32) error {
	return e.Element.SetProperty("ssa-2", value)
}

// ssa-3 (float32)
//
// Source 3 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa3(value float32) error {
	return e.Element.SetProperty("ssa-3", value)
}

// ssa-4 (float32)
//
// Source 4 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa4(value float32) error {
	return e.Element.SetProperty("ssa-4", value)
}

// ssa-5 (float32)
//
// Source 5 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa5(value float32) error {
	return e.Element.SetProperty("ssa-5", value)
}

// ssa-6 (float32)
//
// Source 6 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa6(value float32) error {
	return e.Element.SetProperty("ssa-6", value)
}

// ssa-7 (float32)
//
// Source 7 angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsa7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssa-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsa7(value float32) error {
	return e.Element.SetProperty("ssa-7", value)
}

// ssap-0 (float32)
//
// Source 0 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap0(value float32) error {
	return e.Element.SetProperty("ssap-0", value)
}

// ssap-1 (float32)
//
// Source 1 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap1(value float32) error {
	return e.Element.SetProperty("ssap-1", value)
}

// ssap-2 (float32)
//
// Source 2 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap2(value float32) error {
	return e.Element.SetProperty("ssap-2", value)
}

// ssap-3 (float32)
//
// Source 3 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap3(value float32) error {
	return e.Element.SetProperty("ssap-3", value)
}

// ssap-4 (float32)
//
// Source 4 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap4(value float32) error {
	return e.Element.SetProperty("ssap-4", value)
}

// ssap-5 (float32)
//
// Source 5 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap5(value float32) error {
	return e.Element.SetProperty("ssap-5", value)
}

// ssap-6 (float32)
//
// Source 6 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap6(value float32) error {
	return e.Element.SetProperty("ssap-6", value)
}

// ssap-7 (float32)
//
// Source 7 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsap7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssap-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsap7(value float32) error {
	return e.Element.SetProperty("ssap-7", value)
}

// ssar-0 (float32)
//
// Source 0 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar0(value float32) error {
	return e.Element.SetProperty("ssar-0", value)
}

// ssar-1 (float32)
//
// Source 1 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar1(value float32) error {
	return e.Element.SetProperty("ssar-1", value)
}

// ssar-2 (float32)
//
// Source 2 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar2(value float32) error {
	return e.Element.SetProperty("ssar-2", value)
}

// ssar-3 (float32)
//
// Source 3 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar3(value float32) error {
	return e.Element.SetProperty("ssar-3", value)
}

// ssar-4 (float32)
//
// Source 4 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar4(value float32) error {
	return e.Element.SetProperty("ssar-4", value)
}

// ssar-5 (float32)
//
// Source 5 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar5(value float32) error {
	return e.Element.SetProperty("ssar-5", value)
}

// ssar-6 (float32)
//
// Source 6 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar6(value float32) error {
	return e.Element.SetProperty("ssar-6", value)
}

// ssar-7 (float32)
//
// Source 7 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsar7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssar-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsar7(value float32) error {
	return e.Element.SetProperty("ssar-7", value)
}

// ssay-0 (float32)
//
// Source 0 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay0(value float32) error {
	return e.Element.SetProperty("ssay-0", value)
}

// ssay-1 (float32)
//
// Source 1 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay1(value float32) error {
	return e.Element.SetProperty("ssay-1", value)
}

// ssay-2 (float32)
//
// Source 2 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay2(value float32) error {
	return e.Element.SetProperty("ssay-2", value)
}

// ssay-3 (float32)
//
// Source 3 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay3(value float32) error {
	return e.Element.SetProperty("ssay-3", value)
}

// ssay-4 (float32)
//
// Source 4 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay4(value float32) error {
	return e.Element.SetProperty("ssay-4", value)
}

// ssay-5 (float32)
//
// Source 5 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay5(value float32) error {
	return e.Element.SetProperty("ssay-5", value)
}

// ssay-6 (float32)
//
// Source 6 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay6(value float32) error {
	return e.Element.SetProperty("ssay-6", value)
}

// ssay-7 (float32)
//
// Source 7 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsay7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssay-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsay7(value float32) error {
	return e.Element.SetProperty("ssay-7", value)
}

// sscf-0 (GstLspPlugInPluginsLv2RoomBuilderMonosscf0)
//
// Source 0 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf0() (interface{}, error) {
	return e.Element.GetProperty("sscf-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf0(value interface{}) error {
	return e.Element.SetProperty("sscf-0", value)
}

// sscf-1 (GstLspPlugInPluginsLv2RoomBuilderMonosscf1)
//
// Source 1 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf1() (interface{}, error) {
	return e.Element.GetProperty("sscf-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf1(value interface{}) error {
	return e.Element.SetProperty("sscf-1", value)
}

// sscf-2 (GstLspPlugInPluginsLv2RoomBuilderMonosscf2)
//
// Source 2 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf2() (interface{}, error) {
	return e.Element.GetProperty("sscf-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf2(value interface{}) error {
	return e.Element.SetProperty("sscf-2", value)
}

// sscf-3 (GstLspPlugInPluginsLv2RoomBuilderMonosscf3)
//
// Source 3 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf3() (interface{}, error) {
	return e.Element.GetProperty("sscf-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf3(value interface{}) error {
	return e.Element.SetProperty("sscf-3", value)
}

// sscf-4 (GstLspPlugInPluginsLv2RoomBuilderMonosscf4)
//
// Source 4 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf4() (interface{}, error) {
	return e.Element.GetProperty("sscf-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf4(value interface{}) error {
	return e.Element.SetProperty("sscf-4", value)
}

// sscf-5 (GstLspPlugInPluginsLv2RoomBuilderMonosscf5)
//
// Source 5 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf5() (interface{}, error) {
	return e.Element.GetProperty("sscf-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf5(value interface{}) error {
	return e.Element.SetProperty("sscf-5", value)
}

// sscf-6 (GstLspPlugInPluginsLv2RoomBuilderMonosscf6)
//
// Source 6 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf6() (interface{}, error) {
	return e.Element.GetProperty("sscf-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf6(value interface{}) error {
	return e.Element.SetProperty("sscf-6", value)
}

// sscf-7 (GstLspPlugInPluginsLv2RoomBuilderMonosscf7)
//
// Source 7 type

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscf7() (interface{}, error) {
	return e.Element.GetProperty("sscf-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscf7(value interface{}) error {
	return e.Element.SetProperty("sscf-7", value)
}

// sscv-0 (float32)
//
// Source 0 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv0(value float32) error {
	return e.Element.SetProperty("sscv-0", value)
}

// sscv-1 (float32)
//
// Source 1 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv1(value float32) error {
	return e.Element.SetProperty("sscv-1", value)
}

// sscv-2 (float32)
//
// Source 2 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv2(value float32) error {
	return e.Element.SetProperty("sscv-2", value)
}

// sscv-3 (float32)
//
// Source 3 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv3(value float32) error {
	return e.Element.SetProperty("sscv-3", value)
}

// sscv-4 (float32)
//
// Source 4 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv4(value float32) error {
	return e.Element.SetProperty("sscv-4", value)
}

// sscv-5 (float32)
//
// Source 5 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv5(value float32) error {
	return e.Element.SetProperty("sscv-5", value)
}

// sscv-6 (float32)
//
// Source 6 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv6(value float32) error {
	return e.Element.SetProperty("sscv-6", value)
}

// sscv-7 (float32)
//
// Source 7 curvature

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSscv7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sscv-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSscv7(value float32) error {
	return e.Element.SetProperty("sscv-7", value)
}

// sse-0 (bool)
//
// Source 0 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse0(value bool) error {
	return e.Element.SetProperty("sse-0", value)
}

// sse-1 (bool)
//
// Source 1 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse1(value bool) error {
	return e.Element.SetProperty("sse-1", value)
}

// sse-2 (bool)
//
// Source 2 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse2(value bool) error {
	return e.Element.SetProperty("sse-2", value)
}

// sse-3 (bool)
//
// Source 3 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse3(value bool) error {
	return e.Element.SetProperty("sse-3", value)
}

// sse-4 (bool)
//
// Source 4 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse4(value bool) error {
	return e.Element.SetProperty("sse-4", value)
}

// sse-5 (bool)
//
// Source 5 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse5(value bool) error {
	return e.Element.SetProperty("sse-5", value)
}

// sse-6 (bool)
//
// Source 6 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse6(value bool) error {
	return e.Element.SetProperty("sse-6", value)
}

// sse-7 (bool)
//
// Source 7 enable

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSse7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sse-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSse7(value bool) error {
	return e.Element.SetProperty("sse-7", value)
}

// ssel (GstLspPlugInPluginsLv2RoomBuilderMonossel)
//
// Source selector

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsel() (interface{}, error) {
	return e.Element.GetProperty("ssel")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsel(value interface{}) error {
	return e.Element.SetProperty("ssel", value)
}

// ssh-0 (float32)
//
// Source 0 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh0(value float32) error {
	return e.Element.SetProperty("ssh-0", value)
}

// ssh-1 (float32)
//
// Source 1 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh1(value float32) error {
	return e.Element.SetProperty("ssh-1", value)
}

// ssh-2 (float32)
//
// Source 2 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh2(value float32) error {
	return e.Element.SetProperty("ssh-2", value)
}

// ssh-3 (float32)
//
// Source 3 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh3(value float32) error {
	return e.Element.SetProperty("ssh-3", value)
}

// ssh-4 (float32)
//
// Source 4 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh4(value float32) error {
	return e.Element.SetProperty("ssh-4", value)
}

// ssh-5 (float32)
//
// Source 5 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh5(value float32) error {
	return e.Element.SetProperty("ssh-5", value)
}

// ssh-6 (float32)
//
// Source 6 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh6(value float32) error {
	return e.Element.SetProperty("ssh-6", value)
}

// ssh-7 (float32)
//
// Source 7 hue

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsh7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ssh-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsh7(value float32) error {
	return e.Element.SetProperty("ssh-7", value)
}

// ssph-0 (bool)
//
// Source 0 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph0(value bool) error {
	return e.Element.SetProperty("ssph-0", value)
}

// ssph-1 (bool)
//
// Source 1 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph1(value bool) error {
	return e.Element.SetProperty("ssph-1", value)
}

// ssph-2 (bool)
//
// Source 2 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph2(value bool) error {
	return e.Element.SetProperty("ssph-2", value)
}

// ssph-3 (bool)
//
// Source 3 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph3(value bool) error {
	return e.Element.SetProperty("ssph-3", value)
}

// ssph-4 (bool)
//
// Source 4 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph4(value bool) error {
	return e.Element.SetProperty("ssph-4", value)
}

// ssph-5 (bool)
//
// Source 5 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph5(value bool) error {
	return e.Element.SetProperty("ssph-5", value)
}

// ssph-6 (bool)
//
// Source 6 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph6(value bool) error {
	return e.Element.SetProperty("ssph-6", value)
}

// ssph-7 (bool)
//
// Source 7 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSsph7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssph-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSsph7(value bool) error {
	return e.Element.SetProperty("ssph-7", value)
}

// sspx-0 (float32)
//
// Source 0 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx0(value float32) error {
	return e.Element.SetProperty("sspx-0", value)
}

// sspx-1 (float32)
//
// Source 1 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx1(value float32) error {
	return e.Element.SetProperty("sspx-1", value)
}

// sspx-2 (float32)
//
// Source 2 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx2(value float32) error {
	return e.Element.SetProperty("sspx-2", value)
}

// sspx-3 (float32)
//
// Source 3 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx3(value float32) error {
	return e.Element.SetProperty("sspx-3", value)
}

// sspx-4 (float32)
//
// Source 4 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx4(value float32) error {
	return e.Element.SetProperty("sspx-4", value)
}

// sspx-5 (float32)
//
// Source 5 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx5(value float32) error {
	return e.Element.SetProperty("sspx-5", value)
}

// sspx-6 (float32)
//
// Source 6 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx6(value float32) error {
	return e.Element.SetProperty("sspx-6", value)
}

// sspx-7 (float32)
//
// Source 7 X position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspx7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspx-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspx7(value float32) error {
	return e.Element.SetProperty("sspx-7", value)
}

// sspy-0 (float32)
//
// Source 0 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy0(value float32) error {
	return e.Element.SetProperty("sspy-0", value)
}

// sspy-1 (float32)
//
// Source 1 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy1(value float32) error {
	return e.Element.SetProperty("sspy-1", value)
}

// sspy-2 (float32)
//
// Source 2 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy2(value float32) error {
	return e.Element.SetProperty("sspy-2", value)
}

// sspy-3 (float32)
//
// Source 3 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy3(value float32) error {
	return e.Element.SetProperty("sspy-3", value)
}

// sspy-4 (float32)
//
// Source 4 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy4(value float32) error {
	return e.Element.SetProperty("sspy-4", value)
}

// sspy-5 (float32)
//
// Source 5 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy5(value float32) error {
	return e.Element.SetProperty("sspy-5", value)
}

// sspy-6 (float32)
//
// Source 6 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy6(value float32) error {
	return e.Element.SetProperty("sspy-6", value)
}

// sspy-7 (float32)
//
// Source 7 Y position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspy7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspy-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspy7(value float32) error {
	return e.Element.SetProperty("sspy-7", value)
}

// sspz-0 (float32)
//
// Source 0 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz0(value float32) error {
	return e.Element.SetProperty("sspz-0", value)
}

// sspz-1 (float32)
//
// Source 1 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz1(value float32) error {
	return e.Element.SetProperty("sspz-1", value)
}

// sspz-2 (float32)
//
// Source 2 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz2(value float32) error {
	return e.Element.SetProperty("sspz-2", value)
}

// sspz-3 (float32)
//
// Source 3 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz3(value float32) error {
	return e.Element.SetProperty("sspz-3", value)
}

// sspz-4 (float32)
//
// Source 4 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz4(value float32) error {
	return e.Element.SetProperty("sspz-4", value)
}

// sspz-5 (float32)
//
// Source 5 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz5(value float32) error {
	return e.Element.SetProperty("sspz-5", value)
}

// sspz-6 (float32)
//
// Source 6 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz6(value float32) error {
	return e.Element.SetProperty("sspz-6", value)
}

// sspz-7 (float32)
//
// Source 7 Z position

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSspz7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sspz-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSspz7(value float32) error {
	return e.Element.SetProperty("sspz-7", value)
}

// sss-0 (float32)
//
// Source 0 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss0(value float32) error {
	return e.Element.SetProperty("sss-0", value)
}

// sss-1 (float32)
//
// Source 1 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss1(value float32) error {
	return e.Element.SetProperty("sss-1", value)
}

// sss-2 (float32)
//
// Source 2 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss2(value float32) error {
	return e.Element.SetProperty("sss-2", value)
}

// sss-3 (float32)
//
// Source 3 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss3(value float32) error {
	return e.Element.SetProperty("sss-3", value)
}

// sss-4 (float32)
//
// Source 4 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss4(value float32) error {
	return e.Element.SetProperty("sss-4", value)
}

// sss-5 (float32)
//
// Source 5 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss5(value float32) error {
	return e.Element.SetProperty("sss-5", value)
}

// sss-6 (float32)
//
// Source 6 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss6(value float32) error {
	return e.Element.SetProperty("sss-6", value)
}

// sss-7 (float32)
//
// Source 7 size

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetSss7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sss-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetSss7(value float32) error {
	return e.Element.SetProperty("sss-7", value)
}

// status (int)
//
// Render status

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetStatus() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("status")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// threads (float32)
//
// Number of threads for processing

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetThreads() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetThreads(value float32) error {
	return e.Element.SetProperty("threads", value)
}

// view (GstLspPlugInPluginsLv2RoomBuilderMonoview)
//
// Current view

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetView() (interface{}, error) {
	return e.Element.GetProperty("view")
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetView(value interface{}) error {
	return e.Element.SetProperty("view", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wpp (bool)
//
// Wet post-process

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetWpp() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetWpp(value bool) error {
	return e.Element.SetProperty("wpp", value)
}

// xscale (float32)
//
// Scene X scale

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetXscale() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("xscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetXscale(value float32) error {
	return e.Element.SetProperty("xscale", value)
}

// yscale (float32)
//
// Scene Y scale

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetYscale() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("yscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetYscale(value float32) error {
	return e.Element.SetProperty("yscale", value)
}

// zscale (float32)
//
// Scene Z scale

func (e *LspPlugInPluginsLv2RoomBuilderMono) GetZscale() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("zscale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2RoomBuilderMono) SetZscale(value float32) error {
	return e.Element.SetProperty("zscale", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2RoomBuilderMonocsf1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocsf1None LspPlugInPluginsLv2RoomBuilderMonocsf1 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample0 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample1 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample2 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample3 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample4 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample5 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample6 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderMonocsf1Sample7 LspPlugInPluginsLv2RoomBuilderMonocsf1 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderMonofft string

const (
	LspPlugInPluginsLv2RoomBuilderMonofft512 LspPlugInPluginsLv2RoomBuilderMonofft = "512" // 512 (0) – 512
	LspPlugInPluginsLv2RoomBuilderMonofft1024 LspPlugInPluginsLv2RoomBuilderMonofft = "1024" // 1024 (1) – 1024
	LspPlugInPluginsLv2RoomBuilderMonofft2048 LspPlugInPluginsLv2RoomBuilderMonofft = "2048" // 2048 (2) – 2048
	LspPlugInPluginsLv2RoomBuilderMonofft4096 LspPlugInPluginsLv2RoomBuilderMonofft = "4096" // 4096 (3) – 4096
	LspPlugInPluginsLv2RoomBuilderMonofft8192 LspPlugInPluginsLv2RoomBuilderMonofft = "8192" // 8192 (4) – 8192
	LspPlugInPluginsLv2RoomBuilderMonofft16384 LspPlugInPluginsLv2RoomBuilderMonofft = "16384" // 16384 (5) – 16384
	LspPlugInPluginsLv2RoomBuilderMonofft32767 LspPlugInPluginsLv2RoomBuilderMonofft = "32767" // 32767 (6) – 32767
	LspPlugInPluginsLv2RoomBuilderMonofft65536 LspPlugInPluginsLv2RoomBuilderMonofft = "65536" // 65536 (7) – 65536
)

type LspPlugInPluginsLv2RoomBuilderMonosccf2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf2Mono LspPlugInPluginsLv2RoomBuilderMonosccf2 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf2Xy LspPlugInPluginsLv2RoomBuilderMonosccf2 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf2Ab LspPlugInPluginsLv2RoomBuilderMonosccf2 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf2Ortf LspPlugInPluginsLv2RoomBuilderMonosccf2 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf2Ms LspPlugInPluginsLv2RoomBuilderMonosccf2 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonoscl2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl2Any LspPlugInPluginsLv2RoomBuilderMonoscl2 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl20 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl21 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl22 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl23 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl24 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl25 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl26 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl27 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl28 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl29 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl210 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl211 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl212 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl213 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl214 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl215 LspPlugInPluginsLv2RoomBuilderMonoscl2 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd0Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd0 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd0Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd0 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd0Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd0 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd0Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd0 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd08Directional LspPlugInPluginsLv2RoomBuilderMonoscmd0 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd0Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd0 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd2Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd2 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd28Directional LspPlugInPluginsLv2RoomBuilderMonoscsd2 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonossel string

const (
	LspPlugInPluginsLv2RoomBuilderMonossel0 LspPlugInPluginsLv2RoomBuilderMonossel = "0" // 0 (0) – 0
	LspPlugInPluginsLv2RoomBuilderMonossel1 LspPlugInPluginsLv2RoomBuilderMonossel = "1" // 1 (1) – 1
	LspPlugInPluginsLv2RoomBuilderMonossel2 LspPlugInPluginsLv2RoomBuilderMonossel = "2" // 2 (2) – 2
	LspPlugInPluginsLv2RoomBuilderMonossel3 LspPlugInPluginsLv2RoomBuilderMonossel = "3" // 3 (3) – 3
	LspPlugInPluginsLv2RoomBuilderMonossel4 LspPlugInPluginsLv2RoomBuilderMonossel = "4" // 4 (4) – 4
	LspPlugInPluginsLv2RoomBuilderMonossel5 LspPlugInPluginsLv2RoomBuilderMonossel = "5" // 5 (5) – 5
	LspPlugInPluginsLv2RoomBuilderMonossel6 LspPlugInPluginsLv2RoomBuilderMonossel = "6" // 6 (6) – 6
	LspPlugInPluginsLv2RoomBuilderMonossel7 LspPlugInPluginsLv2RoomBuilderMonossel = "7" // 7 (7) – 7
)

type LspPlugInPluginsLv2RoomBuilderMonolcm string

const (
	LspPlugInPluginsLv2RoomBuilderMonolcmOff LspPlugInPluginsLv2RoomBuilderMonolcm = "off" // off (0) – off
	LspPlugInPluginsLv2RoomBuilderMonolcm6DBoct LspPlugInPluginsLv2RoomBuilderMonolcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2RoomBuilderMonolcm12DBoct LspPlugInPluginsLv2RoomBuilderMonolcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2RoomBuilderMonolcm18DBoct LspPlugInPluginsLv2RoomBuilderMonolcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd5 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd5Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd5 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd5Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd5 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd5Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd5 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd5Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd5 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd58Directional LspPlugInPluginsLv2RoomBuilderMonoscmd5 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd5Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd5 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonosscf4 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf4Triangle LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf4Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf4Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf4Box LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf4Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf4Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf4Cone LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf4Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf4Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf4FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf4CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf4SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf4 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonoeditor string

const (
	LspPlugInPluginsLv2RoomBuilderMonoeditorSourceEditor LspPlugInPluginsLv2RoomBuilderMonoeditor = "Source editor" // Source editor (0) – Source editor
	LspPlugInPluginsLv2RoomBuilderMonoeditorCaptureEditor LspPlugInPluginsLv2RoomBuilderMonoeditor = "Capture editor" // Capture editor (1) – Capture editor
	LspPlugInPluginsLv2RoomBuilderMonoeditorObjectEditor LspPlugInPluginsLv2RoomBuilderMonoeditor = "Object editor" // Object editor (2) – Object editor
	LspPlugInPluginsLv2RoomBuilderMonoeditorMaterialEditor LspPlugInPluginsLv2RoomBuilderMonoeditor = "Material editor" // Material editor (3) – Material editor
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd6 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd6Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd6 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd6Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd6 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd6Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd6 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd6Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd6 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd68Directional LspPlugInPluginsLv2RoomBuilderMonoscmd6 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd6Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd6 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd4 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd4Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd4 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd48Directional LspPlugInPluginsLv2RoomBuilderMonoscsd4 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonosscf2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf2Triangle LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf2Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf2Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf2Box LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf2Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf2Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf2Cone LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf2Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf2Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf2FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf2CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf2SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf2 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonoview string

const (
	LspPlugInPluginsLv2RoomBuilderMonoviewRoomBrowser LspPlugInPluginsLv2RoomBuilderMonoview = "Room browser" // Room browser (0) – Room browser
	LspPlugInPluginsLv2RoomBuilderMonoviewSample0 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderMonoviewSample1 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderMonoviewSample2 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderMonoviewSample3 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderMonoviewSample4 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderMonoviewSample5 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderMonoviewSample6 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderMonoviewSample7 LspPlugInPluginsLv2RoomBuilderMonoview = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderMonoscf1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf1Any LspPlugInPluginsLv2RoomBuilderMonoscf1 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf10 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf11 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf12 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf13 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf14 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf15 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf16 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf17 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf18 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf19 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf110 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf111 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf112 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf113 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf114 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf115 LspPlugInPluginsLv2RoomBuilderMonoscf1 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonosscf0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf0Triangle LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf0Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf0Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf0Box LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf0Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf0Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf0Cone LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf0Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf0Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf0FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf0CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf0SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf0 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonosscf1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf1Triangle LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf1Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf1Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf1Box LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf1Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf1Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf1Cone LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf1Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf1Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf1FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf1CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf1SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf1 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonosscf5 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf5Triangle LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf5Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf5Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf5Box LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf5Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf5Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf5Cone LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf5Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf5Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf5FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf5CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf5SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf5 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonosscf7 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf7Triangle LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf7Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf7Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf7Box LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf7Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf7Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf7Cone LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf7Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf7Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf7FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf7CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf7SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf7 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonocsel string

const (
	LspPlugInPluginsLv2RoomBuilderMonocsel0 LspPlugInPluginsLv2RoomBuilderMonocsel = "0" // 0 (0) – 0
	LspPlugInPluginsLv2RoomBuilderMonocsel1 LspPlugInPluginsLv2RoomBuilderMonocsel = "1" // 1 (1) – 1
	LspPlugInPluginsLv2RoomBuilderMonocsel2 LspPlugInPluginsLv2RoomBuilderMonocsel = "2" // 2 (2) – 2
	LspPlugInPluginsLv2RoomBuilderMonocsel3 LspPlugInPluginsLv2RoomBuilderMonocsel = "3" // 3 (3) – 3
	LspPlugInPluginsLv2RoomBuilderMonocsel4 LspPlugInPluginsLv2RoomBuilderMonocsel = "4" // 4 (4) – 4
	LspPlugInPluginsLv2RoomBuilderMonocsel5 LspPlugInPluginsLv2RoomBuilderMonocsel = "5" // 5 (5) – 5
	LspPlugInPluginsLv2RoomBuilderMonocsel6 LspPlugInPluginsLv2RoomBuilderMonocsel = "6" // 6 (6) – 6
	LspPlugInPluginsLv2RoomBuilderMonocsel7 LspPlugInPluginsLv2RoomBuilderMonocsel = "7" // 7 (7) – 7
)

type LspPlugInPluginsLv2RoomBuilderMonohcm string

const (
	LspPlugInPluginsLv2RoomBuilderMonohcmOff LspPlugInPluginsLv2RoomBuilderMonohcm = "off" // off (0) – off
	LspPlugInPluginsLv2RoomBuilderMonohcm6DBoct LspPlugInPluginsLv2RoomBuilderMonohcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2RoomBuilderMonohcm12DBoct LspPlugInPluginsLv2RoomBuilderMonohcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2RoomBuilderMonohcm18DBoct LspPlugInPluginsLv2RoomBuilderMonohcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2RoomBuilderMonosccf1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf1Mono LspPlugInPluginsLv2RoomBuilderMonosccf1 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf1Xy LspPlugInPluginsLv2RoomBuilderMonosccf1 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf1Ab LspPlugInPluginsLv2RoomBuilderMonosccf1 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf1Ortf LspPlugInPluginsLv2RoomBuilderMonosccf1 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf1Ms LspPlugInPluginsLv2RoomBuilderMonosccf1 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonosccf6 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf6Mono LspPlugInPluginsLv2RoomBuilderMonosccf6 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf6Xy LspPlugInPluginsLv2RoomBuilderMonosccf6 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf6Ab LspPlugInPluginsLv2RoomBuilderMonosccf6 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf6Ortf LspPlugInPluginsLv2RoomBuilderMonosccf6 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf6Ms LspPlugInPluginsLv2RoomBuilderMonosccf6 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonoscl5 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl5Any LspPlugInPluginsLv2RoomBuilderMonoscl5 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl50 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl51 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl52 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl53 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl54 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl55 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl56 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl57 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl58 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl59 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl510 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl511 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl512 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl513 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl514 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl515 LspPlugInPluginsLv2RoomBuilderMonoscl5 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscl7 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl7Any LspPlugInPluginsLv2RoomBuilderMonoscl7 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl70 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl71 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl72 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl73 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl74 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl75 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl76 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl77 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl78 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl79 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl710 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl711 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl712 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl713 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl714 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl715 LspPlugInPluginsLv2RoomBuilderMonoscl7 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonocst0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocst0Left LspPlugInPluginsLv2RoomBuilderMonocst0 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderMonocst0Right LspPlugInPluginsLv2RoomBuilderMonocst0 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderMonoscf7 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf7Any LspPlugInPluginsLv2RoomBuilderMonoscf7 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf70 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf71 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf72 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf73 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf74 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf75 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf76 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf77 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf78 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf79 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf710 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf711 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf712 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf713 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf714 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf715 LspPlugInPluginsLv2RoomBuilderMonoscf7 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscl3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl3Any LspPlugInPluginsLv2RoomBuilderMonoscl3 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl30 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl31 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl32 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl33 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl34 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl35 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl36 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl37 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl38 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl39 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl310 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl311 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl312 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl313 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl314 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl315 LspPlugInPluginsLv2RoomBuilderMonoscl3 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscl6 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl6Any LspPlugInPluginsLv2RoomBuilderMonoscl6 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl60 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl61 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl62 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl63 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl64 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl65 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl66 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl67 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl68 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl69 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl610 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl611 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl612 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl613 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl614 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl615 LspPlugInPluginsLv2RoomBuilderMonoscl6 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonosscf3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf3Triangle LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf3Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf3Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf3Box LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf3Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf3Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf3Cone LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf3Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf3Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf3FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf3CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf3SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf3 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonosscf6 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosscf6Triangle LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderMonosscf6Tetrahedron LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf6Octahedron LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf6Box LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderMonosscf6Icosahedron LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderMonosscf6Cylinder LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderMonosscf6Cone LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderMonosscf6Octasphere LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderMonosscf6Icosphere LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderMonosscf6FlatSpot LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderMonosscf6CylindricalSpot LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderMonosscf6SphericalSpot LspPlugInPluginsLv2RoomBuilderMonosscf6 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderMonocsf2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocsf2None LspPlugInPluginsLv2RoomBuilderMonocsf2 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample0 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample1 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample2 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample3 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample4 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample5 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample6 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderMonocsf2Sample7 LspPlugInPluginsLv2RoomBuilderMonocsf2 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderMonoscf6 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf6Any LspPlugInPluginsLv2RoomBuilderMonoscf6 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf60 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf61 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf62 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf63 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf64 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf65 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf66 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf67 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf68 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf69 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf610 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf611 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf612 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf613 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf614 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf615 LspPlugInPluginsLv2RoomBuilderMonoscf6 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd0Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd0 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd08Directional LspPlugInPluginsLv2RoomBuilderMonoscsd0 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonocsf0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocsf0None LspPlugInPluginsLv2RoomBuilderMonocsf0 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample0 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample1 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample2 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample3 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample4 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample5 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample6 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderMonocsf0Sample7 LspPlugInPluginsLv2RoomBuilderMonocsf0 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderMonocsf3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocsf3None LspPlugInPluginsLv2RoomBuilderMonocsf3 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample0 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample1 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample2 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample3 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample4 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample5 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample6 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderMonocsf3Sample7 LspPlugInPluginsLv2RoomBuilderMonocsf3 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd1Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd1 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd1Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd1 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd1Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd1 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd1Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd1 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd18Directional LspPlugInPluginsLv2RoomBuilderMonoscmd1 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd1Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd1 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonocst1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocst1Left LspPlugInPluginsLv2RoomBuilderMonocst1 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderMonocst1Right LspPlugInPluginsLv2RoomBuilderMonocst1 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderMonosccf7 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf7Mono LspPlugInPluginsLv2RoomBuilderMonosccf7 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf7Xy LspPlugInPluginsLv2RoomBuilderMonosccf7 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf7Ab LspPlugInPluginsLv2RoomBuilderMonosccf7 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf7Ortf LspPlugInPluginsLv2RoomBuilderMonosccf7 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf7Ms LspPlugInPluginsLv2RoomBuilderMonosccf7 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonoscl0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl0Any LspPlugInPluginsLv2RoomBuilderMonoscl0 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl00 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl01 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl02 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl03 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl04 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl05 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl06 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl07 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl08 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl09 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl010 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl011 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl012 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl013 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl014 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl015 LspPlugInPluginsLv2RoomBuilderMonoscl0 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd4 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd4Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd4 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd4Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd4 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd4Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd4 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd4Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd4 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd48Directional LspPlugInPluginsLv2RoomBuilderMonoscmd4 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd4Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd4 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonosccf5 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf5Mono LspPlugInPluginsLv2RoomBuilderMonosccf5 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf5Xy LspPlugInPluginsLv2RoomBuilderMonosccf5 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf5Ab LspPlugInPluginsLv2RoomBuilderMonosccf5 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf5Ortf LspPlugInPluginsLv2RoomBuilderMonosccf5 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf5Ms LspPlugInPluginsLv2RoomBuilderMonosccf5 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonoscf4 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf4Any LspPlugInPluginsLv2RoomBuilderMonoscf4 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf40 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf41 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf42 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf43 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf44 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf45 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf46 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf47 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf48 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf49 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf410 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf411 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf412 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf413 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf414 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf415 LspPlugInPluginsLv2RoomBuilderMonoscf4 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd7 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd7Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd7 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd7Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd7 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd7Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd7 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd7Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd7 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd78Directional LspPlugInPluginsLv2RoomBuilderMonoscmd7 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd7Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd7 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd1Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd1 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd18Directional LspPlugInPluginsLv2RoomBuilderMonoscsd1 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd3Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd3 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd3Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd3 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd3Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd3 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd3Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd3 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd38Directional LspPlugInPluginsLv2RoomBuilderMonoscmd3 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd3Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd3 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonocst2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocst2Left LspPlugInPluginsLv2RoomBuilderMonocst2 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderMonocst2Right LspPlugInPluginsLv2RoomBuilderMonocst2 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderMonocst3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonocst3Left LspPlugInPluginsLv2RoomBuilderMonocst3 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderMonocst3Right LspPlugInPluginsLv2RoomBuilderMonocst3 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderMonoscf0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf0Any LspPlugInPluginsLv2RoomBuilderMonoscf0 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf00 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf01 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf02 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf03 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf04 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf05 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf06 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf07 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf08 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf09 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf010 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf011 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf012 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf013 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf014 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf015 LspPlugInPluginsLv2RoomBuilderMonoscf0 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscf2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf2Any LspPlugInPluginsLv2RoomBuilderMonoscf2 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf20 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf21 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf22 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf23 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf24 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf25 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf26 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf27 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf28 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf29 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf210 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf211 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf212 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf213 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf214 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf215 LspPlugInPluginsLv2RoomBuilderMonoscf2 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscf3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf3Any LspPlugInPluginsLv2RoomBuilderMonoscf3 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf30 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf31 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf32 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf33 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf34 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf35 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf36 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf37 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf38 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf39 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf310 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf311 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf312 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf313 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf314 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf315 LspPlugInPluginsLv2RoomBuilderMonoscf3 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscf5 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscf5Any LspPlugInPluginsLv2RoomBuilderMonoscf5 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscf50 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscf51 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscf52 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscf53 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscf54 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscf55 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscf56 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscf57 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscf58 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscf59 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscf510 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscf511 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscf512 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscf513 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscf514 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscf515 LspPlugInPluginsLv2RoomBuilderMonoscf5 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscmd2 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscmd2Cardioid LspPlugInPluginsLv2RoomBuilderMonoscmd2 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd2Supercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd2 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd2Hypercardioid LspPlugInPluginsLv2RoomBuilderMonoscmd2 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderMonoscmd2Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd2 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscmd28Directional LspPlugInPluginsLv2RoomBuilderMonoscmd2 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderMonoscmd2Omnidirectional LspPlugInPluginsLv2RoomBuilderMonoscmd2 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd3Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd3 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd38Directional LspPlugInPluginsLv2RoomBuilderMonoscsd3 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd5 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd5Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd5 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd58Directional LspPlugInPluginsLv2RoomBuilderMonoscsd5 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd6 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd6Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd6 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd68Directional LspPlugInPluginsLv2RoomBuilderMonoscsd6 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonosignal string

const (
	LspPlugInPluginsLv2RoomBuilderMonosignalConvolvers LspPlugInPluginsLv2RoomBuilderMonosignal = "Convolvers" // Convolvers (0) – Convolvers
	LspPlugInPluginsLv2RoomBuilderMonosignalIrEqualizer LspPlugInPluginsLv2RoomBuilderMonosignal = "IR Equalizer" // IR Equalizer (1) – IR Equalizer
)

type LspPlugInPluginsLv2RoomBuilderMonosccf0 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf0Mono LspPlugInPluginsLv2RoomBuilderMonosccf0 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf0Xy LspPlugInPluginsLv2RoomBuilderMonosccf0 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf0Ab LspPlugInPluginsLv2RoomBuilderMonosccf0 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf0Ortf LspPlugInPluginsLv2RoomBuilderMonosccf0 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf0Ms LspPlugInPluginsLv2RoomBuilderMonosccf0 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonosccf3 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf3Mono LspPlugInPluginsLv2RoomBuilderMonosccf3 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf3Xy LspPlugInPluginsLv2RoomBuilderMonosccf3 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf3Ab LspPlugInPluginsLv2RoomBuilderMonosccf3 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf3Ortf LspPlugInPluginsLv2RoomBuilderMonosccf3 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf3Ms LspPlugInPluginsLv2RoomBuilderMonosccf3 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonoscl1 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl1Any LspPlugInPluginsLv2RoomBuilderMonoscl1 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl10 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl11 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl12 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl13 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl14 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl15 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl16 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl17 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl18 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl19 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl110 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl111 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl112 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl113 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl114 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl115 LspPlugInPluginsLv2RoomBuilderMonoscl1 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderMonoscsd7 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscsd7Bidirectional LspPlugInPluginsLv2RoomBuilderMonoscsd7 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderMonoscsd78Directional LspPlugInPluginsLv2RoomBuilderMonoscsd7 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderMonoifo string

const (
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardYUp LspPlugInPluginsLv2RoomBuilderMonoifo = "+X forward, +Y up" // +X forward, +Y up (0) – +X forward, +Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardZUp LspPlugInPluginsLv2RoomBuilderMonoifo = "+X forward, +Z up" // +X forward, +Z up (1) – +X forward, +Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardYUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "+X forward, -Y up" // +X forward, -Y up (2) – +X forward, -Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardZUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "+X forward, -Z up" // +X forward, -Z up (3) – +X forward, -Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardYUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-X forward, +Y up" // -X forward, +Y up (4) – -X forward, +Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardZUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-X forward, +Z up" // -X forward, +Z up (5) – -X forward, +Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardYUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-X forward, -Y up" // -X forward, -Y up (6) – -X forward, -Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoXForwardZUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-X forward, -Z up" // -X forward, -Z up (7) – -X forward, -Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardXUp LspPlugInPluginsLv2RoomBuilderMonoifo = "+Y forward, +X up" // +Y forward, +X up (8) – +Y forward, +X up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardZUp LspPlugInPluginsLv2RoomBuilderMonoifo = "+Y forward, +Z up" // +Y forward, +Z up (9) – +Y forward, +Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardXUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "+Y forward, -X up" // +Y forward, -X up (10) – +Y forward, -X up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardZUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "+Y forward, -Z up" // +Y forward, -Z up (11) – +Y forward, -Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardXUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Y forward, +X up" // -Y forward, +X up (12) – -Y forward, +X up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardZUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Y forward, +Z up" // -Y forward, +Z up (13) – -Y forward, +Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardXUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Y forward, -X up" // -Y forward, -X up (14) – -Y forward, -X up
	LspPlugInPluginsLv2RoomBuilderMonoifoYForwardZUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Y forward, -Z up" // -Y forward, -Z up (15) – -Y forward, -Z up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardXUp LspPlugInPluginsLv2RoomBuilderMonoifo = "+Z forward, +X up" // +Z forward, +X up (16) – +Z forward, +X up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardYUp LspPlugInPluginsLv2RoomBuilderMonoifo = "+Z forward, +Y up" // +Z forward, +Y up (17) – +Z forward, +Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardXUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "+Z forward, -X up" // +Z forward, -X up (18) – +Z forward, -X up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardYUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "+Z forward, -Y up" // +Z forward, -Y up (19) – +Z forward, -Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardXUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Z forward, +X up" // -Z forward, +X up (20) – -Z forward, +X up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardYUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Z forward, +Y up" // -Z forward, +Y up (21) – -Z forward, +Y up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardXUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Z forward, -X up" // -Z forward, -X up (22) – -Z forward, -X up
	LspPlugInPluginsLv2RoomBuilderMonoifoZForwardYUp1 LspPlugInPluginsLv2RoomBuilderMonoifo = "-Z forward, -Y up" // -Z forward, -Y up (23) – -Z forward, -Y up
)

type LspPlugInPluginsLv2RoomBuilderMonosccf4 string

const (
	LspPlugInPluginsLv2RoomBuilderMonosccf4Mono LspPlugInPluginsLv2RoomBuilderMonosccf4 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderMonosccf4Xy LspPlugInPluginsLv2RoomBuilderMonosccf4 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderMonosccf4Ab LspPlugInPluginsLv2RoomBuilderMonosccf4 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderMonosccf4Ortf LspPlugInPluginsLv2RoomBuilderMonosccf4 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderMonosccf4Ms LspPlugInPluginsLv2RoomBuilderMonosccf4 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderMonoscl4 string

const (
	LspPlugInPluginsLv2RoomBuilderMonoscl4Any LspPlugInPluginsLv2RoomBuilderMonoscl4 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderMonoscl40 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderMonoscl41 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderMonoscl42 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderMonoscl43 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderMonoscl44 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderMonoscl45 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderMonoscl46 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderMonoscl47 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderMonoscl48 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderMonoscl49 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderMonoscl410 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderMonoscl411 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderMonoscl412 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderMonoscl413 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderMonoscl414 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderMonoscl415 LspPlugInPluginsLv2RoomBuilderMonoscl4 = "15" // 15 (16) – 15
)

