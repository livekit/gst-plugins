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

type LspPlugInPluginsLv2RoomBuilderStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2RoomBuilderStereo() (*LspPlugInPluginsLv2RoomBuilderStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-room-builder-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2RoomBuilderStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2RoomBuilderStereoWithName(name string) (*LspPlugInPluginsLv2RoomBuilderStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-room-builder-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2RoomBuilderStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// ca0 (bool)
//
// Channel activity 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCa0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCa1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCa2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCa3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCam0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCam0(value bool) error {
	return e.Element.SetProperty("cam0", value)
}

// cam1 (bool)
//
// Channel mute 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCam1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCam1(value bool) error {
	return e.Element.SetProperty("cam1", value)
}

// cam2 (bool)
//
// Channel mute 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCam2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCam2(value bool) error {
	return e.Element.SetProperty("cam2", value)
}

// cam3 (bool)
//
// Channel mute 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCam3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCam3(value bool) error {
	return e.Element.SetProperty("cam3", value)
}

// cim0 (float32)
//
// Left/Right input mix 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCim0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCim0(value float32) error {
	return e.Element.SetProperty("cim0", value)
}

// cim1 (float32)
//
// Left/Right input mix 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCim1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCim1(value float32) error {
	return e.Element.SetProperty("cim1", value)
}

// cim2 (float32)
//
// Left/Right input mix 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCim2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCim2(value float32) error {
	return e.Element.SetProperty("cim2", value)
}

// cim3 (float32)
//
// Left/Right input mix 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCim3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCim3(value float32) error {
	return e.Element.SetProperty("cim3", value)
}

// com0 (float32)
//
// Channel Left/Right output mix 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCom0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCom0(value float32) error {
	return e.Element.SetProperty("com0", value)
}

// com1 (float32)
//
// Channel Left/Right output mix 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCom1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCom1(value float32) error {
	return e.Element.SetProperty("com1", value)
}

// com2 (float32)
//
// Channel Left/Right output mix 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCom2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCom2(value float32) error {
	return e.Element.SetProperty("com2", value)
}

// com3 (float32)
//
// Channel Left/Right output mix 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCom3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCom3(value float32) error {
	return e.Element.SetProperty("com3", value)
}

// cpitch (float32)
//
// Camera Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCpitch() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCpitch(value float32) error {
	return e.Element.SetProperty("cpitch", value)
}

// cposx (float32)
//
// Camera X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCposx() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCposx(value float32) error {
	return e.Element.SetProperty("cposx", value)
}

// cposy (float32)
//
// Camera Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCposy() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCposy(value float32) error {
	return e.Element.SetProperty("cposy", value)
}

// cposz (float32)
//
// Camera Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCposz() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCposz(value float32) error {
	return e.Element.SetProperty("cposz", value)
}

// csel (GstLspPlugInPluginsLv2RoomBuilderStereocsel)
//
// Capture selector

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCsel() (interface{}, error) {
	return e.Element.GetProperty("csel")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCsel(value interface{}) error {
	return e.Element.SetProperty("csel", value)
}

// csf0 (GstLspPlugInPluginsLv2RoomBuilderStereocsf0)
//
// Channel source sample 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCsf0() (interface{}, error) {
	return e.Element.GetProperty("csf0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCsf0(value interface{}) error {
	return e.Element.SetProperty("csf0", value)
}

// csf1 (GstLspPlugInPluginsLv2RoomBuilderStereocsf1)
//
// Channel source sample 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCsf1() (interface{}, error) {
	return e.Element.GetProperty("csf1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCsf1(value interface{}) error {
	return e.Element.SetProperty("csf1", value)
}

// csf2 (GstLspPlugInPluginsLv2RoomBuilderStereocsf2)
//
// Channel source sample 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCsf2() (interface{}, error) {
	return e.Element.GetProperty("csf2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCsf2(value interface{}) error {
	return e.Element.SetProperty("csf2", value)
}

// csf3 (GstLspPlugInPluginsLv2RoomBuilderStereocsf3)
//
// Channel source sample 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCsf3() (interface{}, error) {
	return e.Element.GetProperty("csf3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCsf3(value interface{}) error {
	return e.Element.SetProperty("csf3", value)
}

// cst0 (GstLspPlugInPluginsLv2RoomBuilderStereocst0)
//
// Channel source track 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCst0() (interface{}, error) {
	return e.Element.GetProperty("cst0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCst0(value interface{}) error {
	return e.Element.SetProperty("cst0", value)
}

// cst1 (GstLspPlugInPluginsLv2RoomBuilderStereocst1)
//
// Channel source track 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCst1() (interface{}, error) {
	return e.Element.GetProperty("cst1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCst1(value interface{}) error {
	return e.Element.SetProperty("cst1", value)
}

// cst2 (GstLspPlugInPluginsLv2RoomBuilderStereocst2)
//
// Channel source track 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCst2() (interface{}, error) {
	return e.Element.GetProperty("cst2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCst2(value interface{}) error {
	return e.Element.SetProperty("cst2", value)
}

// cst3 (GstLspPlugInPluginsLv2RoomBuilderStereocst3)
//
// Channel source track 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCst3() (interface{}, error) {
	return e.Element.GetProperty("cst3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCst3(value interface{}) error {
	return e.Element.SetProperty("cst3", value)
}

// cyaw (float32)
//
// Camera Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetCyaw() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetCyaw(value float32) error {
	return e.Element.SetProperty("cyaw", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// editor (GstLspPlugInPluginsLv2RoomBuilderStereoeditor)
//
// Current editor

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEditor() (interface{}, error) {
	return e.Element.GetProperty("editor")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEditor(value interface{}) error {
	return e.Element.SetProperty("editor", value)
}

// eq-0 (float32)
//
// Band 50Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq0(value float32) error {
	return e.Element.SetProperty("eq-0", value)
}

// eq-1 (float32)
//
// Band 107Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq1(value float32) error {
	return e.Element.SetProperty("eq-1", value)
}

// eq-2 (float32)
//
// Band 227Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq2(value float32) error {
	return e.Element.SetProperty("eq-2", value)
}

// eq-3 (float32)
//
// Band 484Hz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq3(value float32) error {
	return e.Element.SetProperty("eq-3", value)
}

// eq-4 (float32)
//
// Band 1 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq4(value float32) error {
	return e.Element.SetProperty("eq-4", value)
}

// eq-5 (float32)
//
// Band 2.2 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq5(value float32) error {
	return e.Element.SetProperty("eq-5", value)
}

// eq-6 (float32)
//
// Band 4.7 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq6(value float32) error {
	return e.Element.SetProperty("eq-6", value)
}

// eq-7 (float32)
//
// Band 10 kHz gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetEq7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetEq7(value float32) error {
	return e.Element.SetProperty("eq-7", value)
}

// fft (GstLspPlugInPluginsLv2RoomBuilderStereofft)
//
// FFT size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hcf (float32)
//
// High-cut frequency

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetHcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetHcf(value float32) error {
	return e.Element.SetProperty("hcf", value)
}

// hcm (GstLspPlugInPluginsLv2RoomBuilderStereohcm)
//
// High-cut mode

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetHcm() (interface{}, error) {
	return e.Element.GetProperty("hcm")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetHcm(value interface{}) error {
	return e.Element.SetProperty("hcm", value)
}

// ifi-0 (float32)
//
// Fade in0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi0(value float32) error {
	return e.Element.SetProperty("ifi-0", value)
}

// ifi-1 (float32)
//
// Fade in1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi1(value float32) error {
	return e.Element.SetProperty("ifi-1", value)
}

// ifi-2 (float32)
//
// Fade in2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi2(value float32) error {
	return e.Element.SetProperty("ifi-2", value)
}

// ifi-3 (float32)
//
// Fade in3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi3(value float32) error {
	return e.Element.SetProperty("ifi-3", value)
}

// ifi-4 (float32)
//
// Fade in4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi4(value float32) error {
	return e.Element.SetProperty("ifi-4", value)
}

// ifi-5 (float32)
//
// Fade in5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi5(value float32) error {
	return e.Element.SetProperty("ifi-5", value)
}

// ifi-6 (float32)
//
// Fade in6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi6(value float32) error {
	return e.Element.SetProperty("ifi-6", value)
}

// ifi-7 (float32)
//
// Fade in7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfi7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfi7(value float32) error {
	return e.Element.SetProperty("ifi-7", value)
}

// ifl-0 (float32)
//
// Impulse length0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfl7() (float32, error) {
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

// ifo (GstLspPlugInPluginsLv2RoomBuilderStereoifo)
//
// Input 3D model orientation

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo() (interface{}, error) {
	return e.Element.GetProperty("ifo")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo(value interface{}) error {
	return e.Element.SetProperty("ifo", value)
}

// ifo-0 (float32)
//
// Fade out0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo0(value float32) error {
	return e.Element.SetProperty("ifo-0", value)
}

// ifo-1 (float32)
//
// Fade out1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo1(value float32) error {
	return e.Element.SetProperty("ifo-1", value)
}

// ifo-2 (float32)
//
// Fade out2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo2(value float32) error {
	return e.Element.SetProperty("ifo-2", value)
}

// ifo-3 (float32)
//
// Fade out3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo3(value float32) error {
	return e.Element.SetProperty("ifo-3", value)
}

// ifo-4 (float32)
//
// Fade out4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo4(value float32) error {
	return e.Element.SetProperty("ifo-4", value)
}

// ifo-5 (float32)
//
// Fade out5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo5(value float32) error {
	return e.Element.SetProperty("ifo-5", value)
}

// ifo-6 (float32)
//
// Fade out6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo6(value float32) error {
	return e.Element.SetProperty("ifo-6", value)
}

// ifo-7 (float32)
//
// Fade out7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfo7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIfo7(value float32) error {
	return e.Element.SetProperty("ifo-7", value)
}

// ifp (float32)
//
// File loading progress

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfp() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs0() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs1() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs2() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs3() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs4() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs5() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs6() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIfs7() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc0(value float32) error {
	return e.Element.SetProperty("ihc-0", value)
}

// ihc-1 (float32)
//
// Head cut1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc1(value float32) error {
	return e.Element.SetProperty("ihc-1", value)
}

// ihc-2 (float32)
//
// Head cut2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc2(value float32) error {
	return e.Element.SetProperty("ihc-2", value)
}

// ihc-3 (float32)
//
// Head cut3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc3(value float32) error {
	return e.Element.SetProperty("ihc-3", value)
}

// ihc-4 (float32)
//
// Head cut4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc4(value float32) error {
	return e.Element.SetProperty("ihc-4", value)
}

// ihc-5 (float32)
//
// Head cut5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc5(value float32) error {
	return e.Element.SetProperty("ihc-5", value)
}

// ihc-6 (float32)
//
// Head cut6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc6(value float32) error {
	return e.Element.SetProperty("ihc-6", value)
}

// ihc-7 (float32)
//
// Head cut7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIhc7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIhc7(value float32) error {
	return e.Element.SetProperty("ihc-7", value)
}

// ils-0 (bool)
//
// Impulse listen0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls0(value bool) error {
	return e.Element.SetProperty("ils-0", value)
}

// ils-1 (bool)
//
// Impulse listen1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls1(value bool) error {
	return e.Element.SetProperty("ils-1", value)
}

// ils-2 (bool)
//
// Impulse listen2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls2(value bool) error {
	return e.Element.SetProperty("ils-2", value)
}

// ils-3 (bool)
//
// Impulse listen3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls3(value bool) error {
	return e.Element.SetProperty("ils-3", value)
}

// ils-4 (bool)
//
// Impulse listen4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls4() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls4(value bool) error {
	return e.Element.SetProperty("ils-4", value)
}

// ils-5 (bool)
//
// Impulse listen5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls5() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls5(value bool) error {
	return e.Element.SetProperty("ils-5", value)
}

// ils-6 (bool)
//
// Impulse listen6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls6() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls6(value bool) error {
	return e.Element.SetProperty("ils-6", value)
}

// ils-7 (bool)
//
// Impulse listen7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIls7() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIls7(value bool) error {
	return e.Element.SetProperty("ils-7", value)
}

// imkp-0 (float32)
//
// Impulse makeup gain0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp0(value float32) error {
	return e.Element.SetProperty("imkp-0", value)
}

// imkp-1 (float32)
//
// Impulse makeup gain1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp1(value float32) error {
	return e.Element.SetProperty("imkp-1", value)
}

// imkp-2 (float32)
//
// Impulse makeup gain2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp2(value float32) error {
	return e.Element.SetProperty("imkp-2", value)
}

// imkp-3 (float32)
//
// Impulse makeup gain3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp3(value float32) error {
	return e.Element.SetProperty("imkp-3", value)
}

// imkp-4 (float32)
//
// Impulse makeup gain4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp4(value float32) error {
	return e.Element.SetProperty("imkp-4", value)
}

// imkp-5 (float32)
//
// Impulse makeup gain5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp5(value float32) error {
	return e.Element.SetProperty("imkp-5", value)
}

// imkp-6 (float32)
//
// Impulse makeup gain6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp6(value float32) error {
	return e.Element.SetProperty("imkp-6", value)
}

// imkp-7 (float32)
//
// Impulse makeup gain7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetImkp7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetImkp7(value float32) error {
	return e.Element.SetProperty("imkp-7", value)
}

// irv-0 (bool)
//
// Impulse reverse0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv0(value bool) error {
	return e.Element.SetProperty("irv-0", value)
}

// irv-1 (bool)
//
// Impulse reverse1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv1(value bool) error {
	return e.Element.SetProperty("irv-1", value)
}

// irv-2 (bool)
//
// Impulse reverse2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv2(value bool) error {
	return e.Element.SetProperty("irv-2", value)
}

// irv-3 (bool)
//
// Impulse reverse3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv3(value bool) error {
	return e.Element.SetProperty("irv-3", value)
}

// irv-4 (bool)
//
// Impulse reverse4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv4() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv4(value bool) error {
	return e.Element.SetProperty("irv-4", value)
}

// irv-5 (bool)
//
// Impulse reverse5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv5() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv5(value bool) error {
	return e.Element.SetProperty("irv-5", value)
}

// irv-6 (bool)
//
// Impulse reverse6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv6() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv6(value bool) error {
	return e.Element.SetProperty("irv-6", value)
}

// irv-7 (bool)
//
// Impulse reverse7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetIrv7() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetIrv7(value bool) error {
	return e.Element.SetProperty("irv-7", value)
}

// itc-0 (float32)
//
// Tail cut0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc0(value float32) error {
	return e.Element.SetProperty("itc-0", value)
}

// itc-1 (float32)
//
// Tail cut1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc1(value float32) error {
	return e.Element.SetProperty("itc-1", value)
}

// itc-2 (float32)
//
// Tail cut2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc2(value float32) error {
	return e.Element.SetProperty("itc-2", value)
}

// itc-3 (float32)
//
// Tail cut3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc3(value float32) error {
	return e.Element.SetProperty("itc-3", value)
}

// itc-4 (float32)
//
// Tail cut4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc4(value float32) error {
	return e.Element.SetProperty("itc-4", value)
}

// itc-5 (float32)
//
// Tail cut5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc5(value float32) error {
	return e.Element.SetProperty("itc-5", value)
}

// itc-6 (float32)
//
// Tail cut6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc6(value float32) error {
	return e.Element.SetProperty("itc-6", value)
}

// itc-7 (float32)
//
// Tail cut7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetItc7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetItc7(value float32) error {
	return e.Element.SetProperty("itc-7", value)
}

// lcf (float32)
//
// Low-cut frequency

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetLcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetLcf(value float32) error {
	return e.Element.SetProperty("lcf", value)
}

// lcm (GstLspPlugInPluginsLv2RoomBuilderStereolcm)
//
// Low-cut mode

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetLcm() (interface{}, error) {
	return e.Element.GetProperty("lcm")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetLcm(value interface{}) error {
	return e.Element.SetProperty("lcm", value)
}

// mk0 (float32)
//
// Makeup gain 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetMk0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetMk0(value float32) error {
	return e.Element.SetProperty("mk0", value)
}

// mk1 (float32)
//
// Makeup gain 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetMk1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetMk1(value float32) error {
	return e.Element.SetProperty("mk1", value)
}

// mk2 (float32)
//
// Makeup gain 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetMk2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetMk2(value float32) error {
	return e.Element.SetProperty("mk2", value)
}

// mk3 (float32)
//
// Makeup gain 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetMk3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetMk3(value float32) error {
	return e.Element.SetProperty("mk3", value)
}

// normal (bool)
//
// Normalize rendered samples

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetNormal() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetNormal(value bool) error {
	return e.Element.SetProperty("normal", value)
}

// ofc-0 (bool)
//
// Sample save command0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc0(value bool) error {
	return e.Element.SetProperty("ofc-0", value)
}

// ofc-1 (bool)
//
// Sample save command1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc1(value bool) error {
	return e.Element.SetProperty("ofc-1", value)
}

// ofc-2 (bool)
//
// Sample save command2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc2(value bool) error {
	return e.Element.SetProperty("ofc-2", value)
}

// ofc-3 (bool)
//
// Sample save command3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc3(value bool) error {
	return e.Element.SetProperty("ofc-3", value)
}

// ofc-4 (bool)
//
// Sample save command4

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc4() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc4(value bool) error {
	return e.Element.SetProperty("ofc-4", value)
}

// ofc-5 (bool)
//
// Sample save command5

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc5() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc5(value bool) error {
	return e.Element.SetProperty("ofc-5", value)
}

// ofc-6 (bool)
//
// Sample save command6

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc6() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc6(value bool) error {
	return e.Element.SetProperty("ofc-6", value)
}

// ofc-7 (bool)
//
// Sample save command7

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfc7() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetOfc7(value bool) error {
	return e.Element.SetProperty("ofc-7", value)
}

// ofp-0 (float32)
//
// Sample saving progress0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfp7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs0() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs1() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs2() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs3() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs4() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs5() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs6() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOfs7() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPd() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPd(value float32) error {
	return e.Element.SetProperty("pd", value)
}

// pd0 (float32)
//
// Channel pre-delay 0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPd0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPd0(value float32) error {
	return e.Element.SetProperty("pd0", value)
}

// pd1 (float32)
//
// Channel pre-delay 1

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPd1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPd1(value float32) error {
	return e.Element.SetProperty("pd1", value)
}

// pd2 (float32)
//
// Channel pre-delay 2

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPd2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPd2(value float32) error {
	return e.Element.SetProperty("pd2", value)
}

// pd3 (float32)
//
// Channel pre-delay 3

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPd3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPd3(value float32) error {
	return e.Element.SetProperty("pd3", value)
}

// pl (float32)
//
// Left channel panorama

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPl() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPl(value float32) error {
	return e.Element.SetProperty("pl", value)
}

// pr (float32)
//
// Right channel panorama

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetPr() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetPr(value float32) error {
	return e.Element.SetProperty("pr", value)
}

// prog (float32)
//
// Rendering progress

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetProg() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetQuality() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

// render (bool)
//
// Launch/Stop rendering process

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetRender() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetRender(value bool) error {
	return e.Element.SetProperty("render", value)
}

// sca-0 (float32)
//
// Capture 0 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca0(value float32) error {
	return e.Element.SetProperty("sca-0", value)
}

// sca-1 (float32)
//
// Capture 1 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca1(value float32) error {
	return e.Element.SetProperty("sca-1", value)
}

// sca-2 (float32)
//
// Capture 2 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca2(value float32) error {
	return e.Element.SetProperty("sca-2", value)
}

// sca-3 (float32)
//
// Capture 3 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca3(value float32) error {
	return e.Element.SetProperty("sca-3", value)
}

// sca-4 (float32)
//
// Capture 4 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca4(value float32) error {
	return e.Element.SetProperty("sca-4", value)
}

// sca-5 (float32)
//
// Capture 5 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca5(value float32) error {
	return e.Element.SetProperty("sca-5", value)
}

// sca-6 (float32)
//
// Capture 6 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca6(value float32) error {
	return e.Element.SetProperty("sca-6", value)
}

// sca-7 (float32)
//
// Capture 7 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSca7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSca7(value float32) error {
	return e.Element.SetProperty("sca-7", value)
}

// scab-0 (float32)
//
// Capture 0 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab0(value float32) error {
	return e.Element.SetProperty("scab-0", value)
}

// scab-1 (float32)
//
// Capture 1 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab1(value float32) error {
	return e.Element.SetProperty("scab-1", value)
}

// scab-2 (float32)
//
// Capture 2 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab2(value float32) error {
	return e.Element.SetProperty("scab-2", value)
}

// scab-3 (float32)
//
// Capture 3 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab3(value float32) error {
	return e.Element.SetProperty("scab-3", value)
}

// scab-4 (float32)
//
// Capture 4 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab4(value float32) error {
	return e.Element.SetProperty("scab-4", value)
}

// scab-5 (float32)
//
// Capture 5 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab5(value float32) error {
	return e.Element.SetProperty("scab-5", value)
}

// scab-6 (float32)
//
// Capture 6 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab6(value float32) error {
	return e.Element.SetProperty("scab-6", value)
}

// scab-7 (float32)
//
// Capture 7 AB distance

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScab7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScab7(value float32) error {
	return e.Element.SetProperty("scab-7", value)
}

// scap-0 (float32)
//
// Capture 0 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap0(value float32) error {
	return e.Element.SetProperty("scap-0", value)
}

// scap-1 (float32)
//
// Capture 1 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap1(value float32) error {
	return e.Element.SetProperty("scap-1", value)
}

// scap-2 (float32)
//
// Capture 2 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap2(value float32) error {
	return e.Element.SetProperty("scap-2", value)
}

// scap-3 (float32)
//
// Capture 3 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap3(value float32) error {
	return e.Element.SetProperty("scap-3", value)
}

// scap-4 (float32)
//
// Capture 4 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap4(value float32) error {
	return e.Element.SetProperty("scap-4", value)
}

// scap-5 (float32)
//
// Capture 5 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap5(value float32) error {
	return e.Element.SetProperty("scap-5", value)
}

// scap-6 (float32)
//
// Capture 6 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap6(value float32) error {
	return e.Element.SetProperty("scap-6", value)
}

// scap-7 (float32)
//
// Capture 7 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScap7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScap7(value float32) error {
	return e.Element.SetProperty("scap-7", value)
}

// scar-0 (float32)
//
// Capture 0 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar0(value float32) error {
	return e.Element.SetProperty("scar-0", value)
}

// scar-1 (float32)
//
// Capture 1 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar1(value float32) error {
	return e.Element.SetProperty("scar-1", value)
}

// scar-2 (float32)
//
// Capture 2 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar2(value float32) error {
	return e.Element.SetProperty("scar-2", value)
}

// scar-3 (float32)
//
// Capture 3 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar3(value float32) error {
	return e.Element.SetProperty("scar-3", value)
}

// scar-4 (float32)
//
// Capture 4 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar4(value float32) error {
	return e.Element.SetProperty("scar-4", value)
}

// scar-5 (float32)
//
// Capture 5 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar5(value float32) error {
	return e.Element.SetProperty("scar-5", value)
}

// scar-6 (float32)
//
// Capture 6 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar6(value float32) error {
	return e.Element.SetProperty("scar-6", value)
}

// scar-7 (float32)
//
// Capture 7 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScar7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScar7(value float32) error {
	return e.Element.SetProperty("scar-7", value)
}

// scay-0 (float32)
//
// Capture 0 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay0(value float32) error {
	return e.Element.SetProperty("scay-0", value)
}

// scay-1 (float32)
//
// Capture 1 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay1(value float32) error {
	return e.Element.SetProperty("scay-1", value)
}

// scay-2 (float32)
//
// Capture 2 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay2(value float32) error {
	return e.Element.SetProperty("scay-2", value)
}

// scay-3 (float32)
//
// Capture 3 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay3(value float32) error {
	return e.Element.SetProperty("scay-3", value)
}

// scay-4 (float32)
//
// Capture 4 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay4(value float32) error {
	return e.Element.SetProperty("scay-4", value)
}

// scay-5 (float32)
//
// Capture 5 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay5(value float32) error {
	return e.Element.SetProperty("scay-5", value)
}

// scay-6 (float32)
//
// Capture 6 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay6(value float32) error {
	return e.Element.SetProperty("scay-6", value)
}

// scay-7 (float32)
//
// Capture 7 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScay7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScay7(value float32) error {
	return e.Element.SetProperty("scay-7", value)
}

// sccf-0 (GstLspPlugInPluginsLv2RoomBuilderStereosccf0)
//
// Capture 0 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf0() (interface{}, error) {
	return e.Element.GetProperty("sccf-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf0(value interface{}) error {
	return e.Element.SetProperty("sccf-0", value)
}

// sccf-1 (GstLspPlugInPluginsLv2RoomBuilderStereosccf1)
//
// Capture 1 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf1() (interface{}, error) {
	return e.Element.GetProperty("sccf-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf1(value interface{}) error {
	return e.Element.SetProperty("sccf-1", value)
}

// sccf-2 (GstLspPlugInPluginsLv2RoomBuilderStereosccf2)
//
// Capture 2 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf2() (interface{}, error) {
	return e.Element.GetProperty("sccf-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf2(value interface{}) error {
	return e.Element.SetProperty("sccf-2", value)
}

// sccf-3 (GstLspPlugInPluginsLv2RoomBuilderStereosccf3)
//
// Capture 3 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf3() (interface{}, error) {
	return e.Element.GetProperty("sccf-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf3(value interface{}) error {
	return e.Element.SetProperty("sccf-3", value)
}

// sccf-4 (GstLspPlugInPluginsLv2RoomBuilderStereosccf4)
//
// Capture 4 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf4() (interface{}, error) {
	return e.Element.GetProperty("sccf-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf4(value interface{}) error {
	return e.Element.SetProperty("sccf-4", value)
}

// sccf-5 (GstLspPlugInPluginsLv2RoomBuilderStereosccf5)
//
// Capture 5 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf5() (interface{}, error) {
	return e.Element.GetProperty("sccf-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf5(value interface{}) error {
	return e.Element.SetProperty("sccf-5", value)
}

// sccf-6 (GstLspPlugInPluginsLv2RoomBuilderStereosccf6)
//
// Capture 6 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf6() (interface{}, error) {
	return e.Element.GetProperty("sccf-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf6(value interface{}) error {
	return e.Element.SetProperty("sccf-6", value)
}

// sccf-7 (GstLspPlugInPluginsLv2RoomBuilderStereosccf7)
//
// Capture 7 configuration

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccf7() (interface{}, error) {
	return e.Element.GetProperty("sccf-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccf7(value interface{}) error {
	return e.Element.SetProperty("sccf-7", value)
}

// sccs-0 (float32)
//
// Capture 0 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs0(value float32) error {
	return e.Element.SetProperty("sccs-0", value)
}

// sccs-1 (float32)
//
// Capture 1 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs1(value float32) error {
	return e.Element.SetProperty("sccs-1", value)
}

// sccs-2 (float32)
//
// Capture 2 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs2(value float32) error {
	return e.Element.SetProperty("sccs-2", value)
}

// sccs-3 (float32)
//
// Capture 3 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs3(value float32) error {
	return e.Element.SetProperty("sccs-3", value)
}

// sccs-4 (float32)
//
// Capture 4 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs4(value float32) error {
	return e.Element.SetProperty("sccs-4", value)
}

// sccs-5 (float32)
//
// Capture 5 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs5(value float32) error {
	return e.Element.SetProperty("sccs-5", value)
}

// sccs-6 (float32)
//
// Capture 6 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs6(value float32) error {
	return e.Element.SetProperty("sccs-6", value)
}

// sccs-7 (float32)
//
// Capture 7 capsule size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSccs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSccs7(value float32) error {
	return e.Element.SetProperty("sccs-7", value)
}

// sce-0 (bool)
//
// Capture 0 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce0(value bool) error {
	return e.Element.SetProperty("sce-0", value)
}

// sce-1 (bool)
//
// Capture 1 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce1(value bool) error {
	return e.Element.SetProperty("sce-1", value)
}

// sce-2 (bool)
//
// Capture 2 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce2(value bool) error {
	return e.Element.SetProperty("sce-2", value)
}

// sce-3 (bool)
//
// Capture 3 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce3(value bool) error {
	return e.Element.SetProperty("sce-3", value)
}

// sce-4 (bool)
//
// Capture 4 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce4() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce4(value bool) error {
	return e.Element.SetProperty("sce-4", value)
}

// sce-5 (bool)
//
// Capture 5 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce5() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce5(value bool) error {
	return e.Element.SetProperty("sce-5", value)
}

// sce-6 (bool)
//
// Capture 6 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce6() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce6(value bool) error {
	return e.Element.SetProperty("sce-6", value)
}

// sce-7 (bool)
//
// Capture 7 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSce7() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSce7(value bool) error {
	return e.Element.SetProperty("sce-7", value)
}

// scf-0 (GstLspPlugInPluginsLv2RoomBuilderStereoscf0)
//
// Capture 0 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf0() (interface{}, error) {
	return e.Element.GetProperty("scf-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf0(value interface{}) error {
	return e.Element.SetProperty("scf-0", value)
}

// scf-1 (GstLspPlugInPluginsLv2RoomBuilderStereoscf1)
//
// Capture 1 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf1() (interface{}, error) {
	return e.Element.GetProperty("scf-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf1(value interface{}) error {
	return e.Element.SetProperty("scf-1", value)
}

// scf-2 (GstLspPlugInPluginsLv2RoomBuilderStereoscf2)
//
// Capture 2 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf2() (interface{}, error) {
	return e.Element.GetProperty("scf-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf2(value interface{}) error {
	return e.Element.SetProperty("scf-2", value)
}

// scf-3 (GstLspPlugInPluginsLv2RoomBuilderStereoscf3)
//
// Capture 3 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf3() (interface{}, error) {
	return e.Element.GetProperty("scf-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf3(value interface{}) error {
	return e.Element.SetProperty("scf-3", value)
}

// scf-4 (GstLspPlugInPluginsLv2RoomBuilderStereoscf4)
//
// Capture 4 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf4() (interface{}, error) {
	return e.Element.GetProperty("scf-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf4(value interface{}) error {
	return e.Element.SetProperty("scf-4", value)
}

// scf-5 (GstLspPlugInPluginsLv2RoomBuilderStereoscf5)
//
// Capture 5 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf5() (interface{}, error) {
	return e.Element.GetProperty("scf-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf5(value interface{}) error {
	return e.Element.SetProperty("scf-5", value)
}

// scf-6 (GstLspPlugInPluginsLv2RoomBuilderStereoscf6)
//
// Capture 6 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf6() (interface{}, error) {
	return e.Element.GetProperty("scf-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf6(value interface{}) error {
	return e.Element.SetProperty("scf-6", value)
}

// scf-7 (GstLspPlugInPluginsLv2RoomBuilderStereoscf7)
//
// Capture 7 first reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScf7() (interface{}, error) {
	return e.Element.GetProperty("scf-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScf7(value interface{}) error {
	return e.Element.SetProperty("scf-7", value)
}

// sch-0 (float32)
//
// Capture 0 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch0(value float32) error {
	return e.Element.SetProperty("sch-0", value)
}

// sch-1 (float32)
//
// Capture 1 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch1(value float32) error {
	return e.Element.SetProperty("sch-1", value)
}

// sch-2 (float32)
//
// Capture 2 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch2(value float32) error {
	return e.Element.SetProperty("sch-2", value)
}

// sch-3 (float32)
//
// Capture 3 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch3(value float32) error {
	return e.Element.SetProperty("sch-3", value)
}

// sch-4 (float32)
//
// Capture 4 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch4(value float32) error {
	return e.Element.SetProperty("sch-4", value)
}

// sch-5 (float32)
//
// Capture 5 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch5(value float32) error {
	return e.Element.SetProperty("sch-5", value)
}

// sch-6 (float32)
//
// Capture 6 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch6(value float32) error {
	return e.Element.SetProperty("sch-6", value)
}

// sch-7 (float32)
//
// Capture 7 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSch7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSch7(value float32) error {
	return e.Element.SetProperty("sch-7", value)
}

// scl-0 (GstLspPlugInPluginsLv2RoomBuilderStereoscl0)
//
// Capture 0 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl0() (interface{}, error) {
	return e.Element.GetProperty("scl-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl0(value interface{}) error {
	return e.Element.SetProperty("scl-0", value)
}

// scl-1 (GstLspPlugInPluginsLv2RoomBuilderStereoscl1)
//
// Capture 1 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl1() (interface{}, error) {
	return e.Element.GetProperty("scl-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl1(value interface{}) error {
	return e.Element.SetProperty("scl-1", value)
}

// scl-2 (GstLspPlugInPluginsLv2RoomBuilderStereoscl2)
//
// Capture 2 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl2() (interface{}, error) {
	return e.Element.GetProperty("scl-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl2(value interface{}) error {
	return e.Element.SetProperty("scl-2", value)
}

// scl-3 (GstLspPlugInPluginsLv2RoomBuilderStereoscl3)
//
// Capture 3 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl3() (interface{}, error) {
	return e.Element.GetProperty("scl-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl3(value interface{}) error {
	return e.Element.SetProperty("scl-3", value)
}

// scl-4 (GstLspPlugInPluginsLv2RoomBuilderStereoscl4)
//
// Capture 4 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl4() (interface{}, error) {
	return e.Element.GetProperty("scl-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl4(value interface{}) error {
	return e.Element.SetProperty("scl-4", value)
}

// scl-5 (GstLspPlugInPluginsLv2RoomBuilderStereoscl5)
//
// Capture 5 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl5() (interface{}, error) {
	return e.Element.GetProperty("scl-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl5(value interface{}) error {
	return e.Element.SetProperty("scl-5", value)
}

// scl-6 (GstLspPlugInPluginsLv2RoomBuilderStereoscl6)
//
// Capture 6 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl6() (interface{}, error) {
	return e.Element.GetProperty("scl-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl6(value interface{}) error {
	return e.Element.SetProperty("scl-6", value)
}

// scl-7 (GstLspPlugInPluginsLv2RoomBuilderStereoscl7)
//
// Capture 7 last reflection order

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScl7() (interface{}, error) {
	return e.Element.GetProperty("scl-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScl7(value interface{}) error {
	return e.Element.SetProperty("scl-7", value)
}

// scmd-0 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd0)
//
// Capture 0 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd0() (interface{}, error) {
	return e.Element.GetProperty("scmd-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd0(value interface{}) error {
	return e.Element.SetProperty("scmd-0", value)
}

// scmd-1 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd1)
//
// Capture 1 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd1() (interface{}, error) {
	return e.Element.GetProperty("scmd-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd1(value interface{}) error {
	return e.Element.SetProperty("scmd-1", value)
}

// scmd-2 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd2)
//
// Capture 2 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd2() (interface{}, error) {
	return e.Element.GetProperty("scmd-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd2(value interface{}) error {
	return e.Element.SetProperty("scmd-2", value)
}

// scmd-3 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd3)
//
// Capture 3 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd3() (interface{}, error) {
	return e.Element.GetProperty("scmd-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd3(value interface{}) error {
	return e.Element.SetProperty("scmd-3", value)
}

// scmd-4 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd4)
//
// Capture 4 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd4() (interface{}, error) {
	return e.Element.GetProperty("scmd-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd4(value interface{}) error {
	return e.Element.SetProperty("scmd-4", value)
}

// scmd-5 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd5)
//
// Capture 5 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd5() (interface{}, error) {
	return e.Element.GetProperty("scmd-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd5(value interface{}) error {
	return e.Element.SetProperty("scmd-5", value)
}

// scmd-6 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd6)
//
// Capture 6 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd6() (interface{}, error) {
	return e.Element.GetProperty("scmd-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd6(value interface{}) error {
	return e.Element.SetProperty("scmd-6", value)
}

// scmd-7 (GstLspPlugInPluginsLv2RoomBuilderStereoscmd7)
//
// Capture 7 microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScmd7() (interface{}, error) {
	return e.Element.GetProperty("scmd-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScmd7(value interface{}) error {
	return e.Element.SetProperty("scmd-7", value)
}

// scpx-0 (float32)
//
// Capture 0 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx0(value float32) error {
	return e.Element.SetProperty("scpx-0", value)
}

// scpx-1 (float32)
//
// Capture 1 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx1(value float32) error {
	return e.Element.SetProperty("scpx-1", value)
}

// scpx-2 (float32)
//
// Capture 2 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx2(value float32) error {
	return e.Element.SetProperty("scpx-2", value)
}

// scpx-3 (float32)
//
// Capture 3 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx3(value float32) error {
	return e.Element.SetProperty("scpx-3", value)
}

// scpx-4 (float32)
//
// Capture 4 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx4(value float32) error {
	return e.Element.SetProperty("scpx-4", value)
}

// scpx-5 (float32)
//
// Capture 5 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx5(value float32) error {
	return e.Element.SetProperty("scpx-5", value)
}

// scpx-6 (float32)
//
// Capture 6 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx6(value float32) error {
	return e.Element.SetProperty("scpx-6", value)
}

// scpx-7 (float32)
//
// Capture 7 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpx7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpx7(value float32) error {
	return e.Element.SetProperty("scpx-7", value)
}

// scpy-0 (float32)
//
// Capture 0 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy0(value float32) error {
	return e.Element.SetProperty("scpy-0", value)
}

// scpy-1 (float32)
//
// Capture 1 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy1(value float32) error {
	return e.Element.SetProperty("scpy-1", value)
}

// scpy-2 (float32)
//
// Capture 2 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy2(value float32) error {
	return e.Element.SetProperty("scpy-2", value)
}

// scpy-3 (float32)
//
// Capture 3 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy3(value float32) error {
	return e.Element.SetProperty("scpy-3", value)
}

// scpy-4 (float32)
//
// Capture 4 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy4(value float32) error {
	return e.Element.SetProperty("scpy-4", value)
}

// scpy-5 (float32)
//
// Capture 5 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy5(value float32) error {
	return e.Element.SetProperty("scpy-5", value)
}

// scpy-6 (float32)
//
// Capture 6 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy6(value float32) error {
	return e.Element.SetProperty("scpy-6", value)
}

// scpy-7 (float32)
//
// Capture 7 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpy7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpy7(value float32) error {
	return e.Element.SetProperty("scpy-7", value)
}

// scpz-0 (float32)
//
// Capture 0 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz0(value float32) error {
	return e.Element.SetProperty("scpz-0", value)
}

// scpz-1 (float32)
//
// Capture 1 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz1(value float32) error {
	return e.Element.SetProperty("scpz-1", value)
}

// scpz-2 (float32)
//
// Capture 2 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz2(value float32) error {
	return e.Element.SetProperty("scpz-2", value)
}

// scpz-3 (float32)
//
// Capture 3 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz3(value float32) error {
	return e.Element.SetProperty("scpz-3", value)
}

// scpz-4 (float32)
//
// Capture 4 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz4(value float32) error {
	return e.Element.SetProperty("scpz-4", value)
}

// scpz-5 (float32)
//
// Capture 5 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz5(value float32) error {
	return e.Element.SetProperty("scpz-5", value)
}

// scpz-6 (float32)
//
// Capture 6 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz6(value float32) error {
	return e.Element.SetProperty("scpz-6", value)
}

// scpz-7 (float32)
//
// Capture 7 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScpz7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScpz7(value float32) error {
	return e.Element.SetProperty("scpz-7", value)
}

// scsd-0 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd0)
//
// Capture 0 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd0() (interface{}, error) {
	return e.Element.GetProperty("scsd-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd0(value interface{}) error {
	return e.Element.SetProperty("scsd-0", value)
}

// scsd-1 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd1)
//
// Capture 1 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd1() (interface{}, error) {
	return e.Element.GetProperty("scsd-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd1(value interface{}) error {
	return e.Element.SetProperty("scsd-1", value)
}

// scsd-2 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd2)
//
// Capture 2 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd2() (interface{}, error) {
	return e.Element.GetProperty("scsd-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd2(value interface{}) error {
	return e.Element.SetProperty("scsd-2", value)
}

// scsd-3 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd3)
//
// Capture 3 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd3() (interface{}, error) {
	return e.Element.GetProperty("scsd-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd3(value interface{}) error {
	return e.Element.SetProperty("scsd-3", value)
}

// scsd-4 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd4)
//
// Capture 4 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd4() (interface{}, error) {
	return e.Element.GetProperty("scsd-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd4(value interface{}) error {
	return e.Element.SetProperty("scsd-4", value)
}

// scsd-5 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd5)
//
// Capture 5 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd5() (interface{}, error) {
	return e.Element.GetProperty("scsd-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd5(value interface{}) error {
	return e.Element.SetProperty("scsd-5", value)
}

// scsd-6 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd6)
//
// Capture 6 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd6() (interface{}, error) {
	return e.Element.GetProperty("scsd-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd6(value interface{}) error {
	return e.Element.SetProperty("scsd-6", value)
}

// scsd-7 (GstLspPlugInPluginsLv2RoomBuilderStereoscsd7)
//
// Capture 7 side microphone direction

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetScsd7() (interface{}, error) {
	return e.Element.GetProperty("scsd-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetScsd7(value interface{}) error {
	return e.Element.SetProperty("scsd-7", value)
}

// sdur-0 (float32)
//
// Impulse current duration0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSdur7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh0(value float32) error {
	return e.Element.SetProperty("shh-0", value)
}

// shh-1 (float32)
//
// Source 1 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh1(value float32) error {
	return e.Element.SetProperty("shh-1", value)
}

// shh-2 (float32)
//
// Source 2 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh2(value float32) error {
	return e.Element.SetProperty("shh-2", value)
}

// shh-3 (float32)
//
// Source 3 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh3(value float32) error {
	return e.Element.SetProperty("shh-3", value)
}

// shh-4 (float32)
//
// Source 4 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh4(value float32) error {
	return e.Element.SetProperty("shh-4", value)
}

// shh-5 (float32)
//
// Source 5 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh5(value float32) error {
	return e.Element.SetProperty("shh-5", value)
}

// shh-6 (float32)
//
// Source 6 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh6(value float32) error {
	return e.Element.SetProperty("shh-6", value)
}

// shh-7 (float32)
//
// Source 7 height

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetShh7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetShh7(value float32) error {
	return e.Element.SetProperty("shh-7", value)
}

// signal (GstLspPlugInPluginsLv2RoomBuilderStereosignal)
//
// Current signal processor

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSignal() (interface{}, error) {
	return e.Element.GetProperty("signal")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSignal(value interface{}) error {
	return e.Element.SetProperty("signal", value)
}

// smdur-0 (float32)
//
// Impulse max duration0

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSmdur7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa0(value float32) error {
	return e.Element.SetProperty("ssa-0", value)
}

// ssa-1 (float32)
//
// Source 1 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa1(value float32) error {
	return e.Element.SetProperty("ssa-1", value)
}

// ssa-2 (float32)
//
// Source 2 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa2(value float32) error {
	return e.Element.SetProperty("ssa-2", value)
}

// ssa-3 (float32)
//
// Source 3 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa3(value float32) error {
	return e.Element.SetProperty("ssa-3", value)
}

// ssa-4 (float32)
//
// Source 4 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa4(value float32) error {
	return e.Element.SetProperty("ssa-4", value)
}

// ssa-5 (float32)
//
// Source 5 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa5(value float32) error {
	return e.Element.SetProperty("ssa-5", value)
}

// ssa-6 (float32)
//
// Source 6 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa6(value float32) error {
	return e.Element.SetProperty("ssa-6", value)
}

// ssa-7 (float32)
//
// Source 7 angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsa7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsa7(value float32) error {
	return e.Element.SetProperty("ssa-7", value)
}

// ssap-0 (float32)
//
// Source 0 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap0(value float32) error {
	return e.Element.SetProperty("ssap-0", value)
}

// ssap-1 (float32)
//
// Source 1 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap1(value float32) error {
	return e.Element.SetProperty("ssap-1", value)
}

// ssap-2 (float32)
//
// Source 2 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap2(value float32) error {
	return e.Element.SetProperty("ssap-2", value)
}

// ssap-3 (float32)
//
// Source 3 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap3(value float32) error {
	return e.Element.SetProperty("ssap-3", value)
}

// ssap-4 (float32)
//
// Source 4 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap4(value float32) error {
	return e.Element.SetProperty("ssap-4", value)
}

// ssap-5 (float32)
//
// Source 5 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap5(value float32) error {
	return e.Element.SetProperty("ssap-5", value)
}

// ssap-6 (float32)
//
// Source 6 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap6(value float32) error {
	return e.Element.SetProperty("ssap-6", value)
}

// ssap-7 (float32)
//
// Source 7 Pitch angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsap7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsap7(value float32) error {
	return e.Element.SetProperty("ssap-7", value)
}

// ssar-0 (float32)
//
// Source 0 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar0(value float32) error {
	return e.Element.SetProperty("ssar-0", value)
}

// ssar-1 (float32)
//
// Source 1 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar1(value float32) error {
	return e.Element.SetProperty("ssar-1", value)
}

// ssar-2 (float32)
//
// Source 2 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar2(value float32) error {
	return e.Element.SetProperty("ssar-2", value)
}

// ssar-3 (float32)
//
// Source 3 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar3(value float32) error {
	return e.Element.SetProperty("ssar-3", value)
}

// ssar-4 (float32)
//
// Source 4 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar4(value float32) error {
	return e.Element.SetProperty("ssar-4", value)
}

// ssar-5 (float32)
//
// Source 5 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar5(value float32) error {
	return e.Element.SetProperty("ssar-5", value)
}

// ssar-6 (float32)
//
// Source 6 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar6(value float32) error {
	return e.Element.SetProperty("ssar-6", value)
}

// ssar-7 (float32)
//
// Source 7 Roll angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsar7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsar7(value float32) error {
	return e.Element.SetProperty("ssar-7", value)
}

// ssay-0 (float32)
//
// Source 0 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay0(value float32) error {
	return e.Element.SetProperty("ssay-0", value)
}

// ssay-1 (float32)
//
// Source 1 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay1(value float32) error {
	return e.Element.SetProperty("ssay-1", value)
}

// ssay-2 (float32)
//
// Source 2 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay2(value float32) error {
	return e.Element.SetProperty("ssay-2", value)
}

// ssay-3 (float32)
//
// Source 3 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay3(value float32) error {
	return e.Element.SetProperty("ssay-3", value)
}

// ssay-4 (float32)
//
// Source 4 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay4(value float32) error {
	return e.Element.SetProperty("ssay-4", value)
}

// ssay-5 (float32)
//
// Source 5 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay5(value float32) error {
	return e.Element.SetProperty("ssay-5", value)
}

// ssay-6 (float32)
//
// Source 6 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay6(value float32) error {
	return e.Element.SetProperty("ssay-6", value)
}

// ssay-7 (float32)
//
// Source 7 Yaw angle

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsay7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsay7(value float32) error {
	return e.Element.SetProperty("ssay-7", value)
}

// sscf-0 (GstLspPlugInPluginsLv2RoomBuilderStereosscf0)
//
// Source 0 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf0() (interface{}, error) {
	return e.Element.GetProperty("sscf-0")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf0(value interface{}) error {
	return e.Element.SetProperty("sscf-0", value)
}

// sscf-1 (GstLspPlugInPluginsLv2RoomBuilderStereosscf1)
//
// Source 1 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf1() (interface{}, error) {
	return e.Element.GetProperty("sscf-1")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf1(value interface{}) error {
	return e.Element.SetProperty("sscf-1", value)
}

// sscf-2 (GstLspPlugInPluginsLv2RoomBuilderStereosscf2)
//
// Source 2 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf2() (interface{}, error) {
	return e.Element.GetProperty("sscf-2")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf2(value interface{}) error {
	return e.Element.SetProperty("sscf-2", value)
}

// sscf-3 (GstLspPlugInPluginsLv2RoomBuilderStereosscf3)
//
// Source 3 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf3() (interface{}, error) {
	return e.Element.GetProperty("sscf-3")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf3(value interface{}) error {
	return e.Element.SetProperty("sscf-3", value)
}

// sscf-4 (GstLspPlugInPluginsLv2RoomBuilderStereosscf4)
//
// Source 4 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf4() (interface{}, error) {
	return e.Element.GetProperty("sscf-4")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf4(value interface{}) error {
	return e.Element.SetProperty("sscf-4", value)
}

// sscf-5 (GstLspPlugInPluginsLv2RoomBuilderStereosscf5)
//
// Source 5 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf5() (interface{}, error) {
	return e.Element.GetProperty("sscf-5")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf5(value interface{}) error {
	return e.Element.SetProperty("sscf-5", value)
}

// sscf-6 (GstLspPlugInPluginsLv2RoomBuilderStereosscf6)
//
// Source 6 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf6() (interface{}, error) {
	return e.Element.GetProperty("sscf-6")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf6(value interface{}) error {
	return e.Element.SetProperty("sscf-6", value)
}

// sscf-7 (GstLspPlugInPluginsLv2RoomBuilderStereosscf7)
//
// Source 7 type

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscf7() (interface{}, error) {
	return e.Element.GetProperty("sscf-7")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscf7(value interface{}) error {
	return e.Element.SetProperty("sscf-7", value)
}

// sscv-0 (float32)
//
// Source 0 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv0(value float32) error {
	return e.Element.SetProperty("sscv-0", value)
}

// sscv-1 (float32)
//
// Source 1 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv1(value float32) error {
	return e.Element.SetProperty("sscv-1", value)
}

// sscv-2 (float32)
//
// Source 2 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv2(value float32) error {
	return e.Element.SetProperty("sscv-2", value)
}

// sscv-3 (float32)
//
// Source 3 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv3(value float32) error {
	return e.Element.SetProperty("sscv-3", value)
}

// sscv-4 (float32)
//
// Source 4 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv4(value float32) error {
	return e.Element.SetProperty("sscv-4", value)
}

// sscv-5 (float32)
//
// Source 5 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv5(value float32) error {
	return e.Element.SetProperty("sscv-5", value)
}

// sscv-6 (float32)
//
// Source 6 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv6(value float32) error {
	return e.Element.SetProperty("sscv-6", value)
}

// sscv-7 (float32)
//
// Source 7 curvature

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSscv7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSscv7(value float32) error {
	return e.Element.SetProperty("sscv-7", value)
}

// sse-0 (bool)
//
// Source 0 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse0(value bool) error {
	return e.Element.SetProperty("sse-0", value)
}

// sse-1 (bool)
//
// Source 1 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse1(value bool) error {
	return e.Element.SetProperty("sse-1", value)
}

// sse-2 (bool)
//
// Source 2 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse2(value bool) error {
	return e.Element.SetProperty("sse-2", value)
}

// sse-3 (bool)
//
// Source 3 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse3(value bool) error {
	return e.Element.SetProperty("sse-3", value)
}

// sse-4 (bool)
//
// Source 4 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse4() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse4(value bool) error {
	return e.Element.SetProperty("sse-4", value)
}

// sse-5 (bool)
//
// Source 5 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse5() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse5(value bool) error {
	return e.Element.SetProperty("sse-5", value)
}

// sse-6 (bool)
//
// Source 6 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse6() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse6(value bool) error {
	return e.Element.SetProperty("sse-6", value)
}

// sse-7 (bool)
//
// Source 7 enable

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSse7() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSse7(value bool) error {
	return e.Element.SetProperty("sse-7", value)
}

// ssel (GstLspPlugInPluginsLv2RoomBuilderStereossel)
//
// Source selector

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsel() (interface{}, error) {
	return e.Element.GetProperty("ssel")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsel(value interface{}) error {
	return e.Element.SetProperty("ssel", value)
}

// ssh-0 (float32)
//
// Source 0 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh0(value float32) error {
	return e.Element.SetProperty("ssh-0", value)
}

// ssh-1 (float32)
//
// Source 1 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh1(value float32) error {
	return e.Element.SetProperty("ssh-1", value)
}

// ssh-2 (float32)
//
// Source 2 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh2(value float32) error {
	return e.Element.SetProperty("ssh-2", value)
}

// ssh-3 (float32)
//
// Source 3 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh3(value float32) error {
	return e.Element.SetProperty("ssh-3", value)
}

// ssh-4 (float32)
//
// Source 4 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh4(value float32) error {
	return e.Element.SetProperty("ssh-4", value)
}

// ssh-5 (float32)
//
// Source 5 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh5(value float32) error {
	return e.Element.SetProperty("ssh-5", value)
}

// ssh-6 (float32)
//
// Source 6 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh6(value float32) error {
	return e.Element.SetProperty("ssh-6", value)
}

// ssh-7 (float32)
//
// Source 7 hue

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsh7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsh7(value float32) error {
	return e.Element.SetProperty("ssh-7", value)
}

// ssph-0 (bool)
//
// Source 0 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph0() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph0(value bool) error {
	return e.Element.SetProperty("ssph-0", value)
}

// ssph-1 (bool)
//
// Source 1 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph1() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph1(value bool) error {
	return e.Element.SetProperty("ssph-1", value)
}

// ssph-2 (bool)
//
// Source 2 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph2() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph2(value bool) error {
	return e.Element.SetProperty("ssph-2", value)
}

// ssph-3 (bool)
//
// Source 3 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph3() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph3(value bool) error {
	return e.Element.SetProperty("ssph-3", value)
}

// ssph-4 (bool)
//
// Source 4 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph4() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph4(value bool) error {
	return e.Element.SetProperty("ssph-4", value)
}

// ssph-5 (bool)
//
// Source 5 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph5() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph5(value bool) error {
	return e.Element.SetProperty("ssph-5", value)
}

// ssph-6 (bool)
//
// Source 6 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph6() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph6(value bool) error {
	return e.Element.SetProperty("ssph-6", value)
}

// ssph-7 (bool)
//
// Source 7 phase invert

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSsph7() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSsph7(value bool) error {
	return e.Element.SetProperty("ssph-7", value)
}

// sspx-0 (float32)
//
// Source 0 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx0(value float32) error {
	return e.Element.SetProperty("sspx-0", value)
}

// sspx-1 (float32)
//
// Source 1 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx1(value float32) error {
	return e.Element.SetProperty("sspx-1", value)
}

// sspx-2 (float32)
//
// Source 2 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx2(value float32) error {
	return e.Element.SetProperty("sspx-2", value)
}

// sspx-3 (float32)
//
// Source 3 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx3(value float32) error {
	return e.Element.SetProperty("sspx-3", value)
}

// sspx-4 (float32)
//
// Source 4 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx4(value float32) error {
	return e.Element.SetProperty("sspx-4", value)
}

// sspx-5 (float32)
//
// Source 5 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx5(value float32) error {
	return e.Element.SetProperty("sspx-5", value)
}

// sspx-6 (float32)
//
// Source 6 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx6(value float32) error {
	return e.Element.SetProperty("sspx-6", value)
}

// sspx-7 (float32)
//
// Source 7 X position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspx7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspx7(value float32) error {
	return e.Element.SetProperty("sspx-7", value)
}

// sspy-0 (float32)
//
// Source 0 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy0(value float32) error {
	return e.Element.SetProperty("sspy-0", value)
}

// sspy-1 (float32)
//
// Source 1 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy1(value float32) error {
	return e.Element.SetProperty("sspy-1", value)
}

// sspy-2 (float32)
//
// Source 2 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy2(value float32) error {
	return e.Element.SetProperty("sspy-2", value)
}

// sspy-3 (float32)
//
// Source 3 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy3(value float32) error {
	return e.Element.SetProperty("sspy-3", value)
}

// sspy-4 (float32)
//
// Source 4 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy4(value float32) error {
	return e.Element.SetProperty("sspy-4", value)
}

// sspy-5 (float32)
//
// Source 5 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy5(value float32) error {
	return e.Element.SetProperty("sspy-5", value)
}

// sspy-6 (float32)
//
// Source 6 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy6(value float32) error {
	return e.Element.SetProperty("sspy-6", value)
}

// sspy-7 (float32)
//
// Source 7 Y position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspy7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspy7(value float32) error {
	return e.Element.SetProperty("sspy-7", value)
}

// sspz-0 (float32)
//
// Source 0 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz0(value float32) error {
	return e.Element.SetProperty("sspz-0", value)
}

// sspz-1 (float32)
//
// Source 1 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz1(value float32) error {
	return e.Element.SetProperty("sspz-1", value)
}

// sspz-2 (float32)
//
// Source 2 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz2(value float32) error {
	return e.Element.SetProperty("sspz-2", value)
}

// sspz-3 (float32)
//
// Source 3 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz3(value float32) error {
	return e.Element.SetProperty("sspz-3", value)
}

// sspz-4 (float32)
//
// Source 4 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz4(value float32) error {
	return e.Element.SetProperty("sspz-4", value)
}

// sspz-5 (float32)
//
// Source 5 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz5(value float32) error {
	return e.Element.SetProperty("sspz-5", value)
}

// sspz-6 (float32)
//
// Source 6 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz6(value float32) error {
	return e.Element.SetProperty("sspz-6", value)
}

// sspz-7 (float32)
//
// Source 7 Z position

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSspz7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSspz7(value float32) error {
	return e.Element.SetProperty("sspz-7", value)
}

// sss-0 (float32)
//
// Source 0 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss0() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss0(value float32) error {
	return e.Element.SetProperty("sss-0", value)
}

// sss-1 (float32)
//
// Source 1 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss1() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss1(value float32) error {
	return e.Element.SetProperty("sss-1", value)
}

// sss-2 (float32)
//
// Source 2 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss2() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss2(value float32) error {
	return e.Element.SetProperty("sss-2", value)
}

// sss-3 (float32)
//
// Source 3 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss3() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss3(value float32) error {
	return e.Element.SetProperty("sss-3", value)
}

// sss-4 (float32)
//
// Source 4 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss4() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss4(value float32) error {
	return e.Element.SetProperty("sss-4", value)
}

// sss-5 (float32)
//
// Source 5 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss5() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss5(value float32) error {
	return e.Element.SetProperty("sss-5", value)
}

// sss-6 (float32)
//
// Source 6 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss6() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss6(value float32) error {
	return e.Element.SetProperty("sss-6", value)
}

// sss-7 (float32)
//
// Source 7 size

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetSss7() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetSss7(value float32) error {
	return e.Element.SetProperty("sss-7", value)
}

// status (int)
//
// Render status

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetStatus() (int, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetThreads() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetThreads(value float32) error {
	return e.Element.SetProperty("threads", value)
}

// view (GstLspPlugInPluginsLv2RoomBuilderStereoview)
//
// Current view

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetView() (interface{}, error) {
	return e.Element.GetProperty("view")
}

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetView(value interface{}) error {
	return e.Element.SetProperty("view", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wpp (bool)
//
// Wet post-process

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetWpp() (bool, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetWpp(value bool) error {
	return e.Element.SetProperty("wpp", value)
}

// xscale (float32)
//
// Scene X scale

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetXscale() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetXscale(value float32) error {
	return e.Element.SetProperty("xscale", value)
}

// yscale (float32)
//
// Scene Y scale

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetYscale() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetYscale(value float32) error {
	return e.Element.SetProperty("yscale", value)
}

// zscale (float32)
//
// Scene Z scale

func (e *LspPlugInPluginsLv2RoomBuilderStereo) GetZscale() (float32, error) {
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

func (e *LspPlugInPluginsLv2RoomBuilderStereo) SetZscale(value float32) error {
	return e.Element.SetProperty("zscale", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2RoomBuilderStereoscl3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl3Any LspPlugInPluginsLv2RoomBuilderStereoscl3 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl30 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl31 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl32 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl33 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl34 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl35 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl36 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl37 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl38 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl39 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl310 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl311 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl312 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl313 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl314 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl315 LspPlugInPluginsLv2RoomBuilderStereoscl3 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd4 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd4Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd4 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd48Directional LspPlugInPluginsLv2RoomBuilderStereoscsd4 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereosscf2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf2Triangle LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf2Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf2Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf2Box LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf2Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf2Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf2Cone LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf2Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf2Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf2FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf2CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf2SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf2 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereocsf0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocsf0None LspPlugInPluginsLv2RoomBuilderStereocsf0 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample0 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample1 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample2 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample3 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample4 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample5 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample6 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderStereocsf0Sample7 LspPlugInPluginsLv2RoomBuilderStereocsf0 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderStereocst3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocst3Left LspPlugInPluginsLv2RoomBuilderStereocst3 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderStereocst3Right LspPlugInPluginsLv2RoomBuilderStereocst3 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderStereoscf5 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf5Any LspPlugInPluginsLv2RoomBuilderStereoscf5 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf50 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf51 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf52 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf53 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf54 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf55 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf56 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf57 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf58 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf59 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf510 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf511 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf512 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf513 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf514 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf515 LspPlugInPluginsLv2RoomBuilderStereoscf5 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd5 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd5Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd5 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd58Directional LspPlugInPluginsLv2RoomBuilderStereoscsd5 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereoscl6 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl6Any LspPlugInPluginsLv2RoomBuilderStereoscl6 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl60 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl61 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl62 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl63 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl64 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl65 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl66 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl67 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl68 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl69 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl610 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl611 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl612 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl613 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl614 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl615 LspPlugInPluginsLv2RoomBuilderStereoscl6 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereossel string

const (
	LspPlugInPluginsLv2RoomBuilderStereossel0 LspPlugInPluginsLv2RoomBuilderStereossel = "0" // 0 (0) – 0
	LspPlugInPluginsLv2RoomBuilderStereossel1 LspPlugInPluginsLv2RoomBuilderStereossel = "1" // 1 (1) – 1
	LspPlugInPluginsLv2RoomBuilderStereossel2 LspPlugInPluginsLv2RoomBuilderStereossel = "2" // 2 (2) – 2
	LspPlugInPluginsLv2RoomBuilderStereossel3 LspPlugInPluginsLv2RoomBuilderStereossel = "3" // 3 (3) – 3
	LspPlugInPluginsLv2RoomBuilderStereossel4 LspPlugInPluginsLv2RoomBuilderStereossel = "4" // 4 (4) – 4
	LspPlugInPluginsLv2RoomBuilderStereossel5 LspPlugInPluginsLv2RoomBuilderStereossel = "5" // 5 (5) – 5
	LspPlugInPluginsLv2RoomBuilderStereossel6 LspPlugInPluginsLv2RoomBuilderStereossel = "6" // 6 (6) – 6
	LspPlugInPluginsLv2RoomBuilderStereossel7 LspPlugInPluginsLv2RoomBuilderStereossel = "7" // 7 (7) – 7
)

type LspPlugInPluginsLv2RoomBuilderStereocst0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocst0Left LspPlugInPluginsLv2RoomBuilderStereocst0 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderStereocst0Right LspPlugInPluginsLv2RoomBuilderStereocst0 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderStereoscl4 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl4Any LspPlugInPluginsLv2RoomBuilderStereoscl4 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl40 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl41 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl42 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl43 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl44 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl45 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl46 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl47 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl48 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl49 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl410 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl411 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl412 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl413 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl414 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl415 LspPlugInPluginsLv2RoomBuilderStereoscl4 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd2Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd2 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd28Directional LspPlugInPluginsLv2RoomBuilderStereoscsd2 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd6 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd6Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd6 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd68Directional LspPlugInPluginsLv2RoomBuilderStereoscsd6 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereocst2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocst2Left LspPlugInPluginsLv2RoomBuilderStereocst2 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderStereocst2Right LspPlugInPluginsLv2RoomBuilderStereocst2 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd5 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd5Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd5 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd5Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd5 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd5Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd5 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd5Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd5 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd58Directional LspPlugInPluginsLv2RoomBuilderStereoscmd5 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd5Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd5 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereosscf3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf3Triangle LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf3Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf3Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf3Box LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf3Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf3Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf3Cone LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf3Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf3Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf3FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf3CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf3SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf3 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereosscf5 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf5Triangle LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf5Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf5Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf5Box LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf5Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf5Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf5Cone LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf5Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf5Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf5FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf5CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf5SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf5 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd1Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd1 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd1Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd1 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd1Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd1 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd1Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd1 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd18Directional LspPlugInPluginsLv2RoomBuilderStereoscmd1 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd1Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd1 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd3Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd3 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd3Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd3 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd3Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd3 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd3Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd3 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd38Directional LspPlugInPluginsLv2RoomBuilderStereoscmd3 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd3Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd3 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereofft string

const (
	LspPlugInPluginsLv2RoomBuilderStereofft512 LspPlugInPluginsLv2RoomBuilderStereofft = "512" // 512 (0) – 512
	LspPlugInPluginsLv2RoomBuilderStereofft1024 LspPlugInPluginsLv2RoomBuilderStereofft = "1024" // 1024 (1) – 1024
	LspPlugInPluginsLv2RoomBuilderStereofft2048 LspPlugInPluginsLv2RoomBuilderStereofft = "2048" // 2048 (2) – 2048
	LspPlugInPluginsLv2RoomBuilderStereofft4096 LspPlugInPluginsLv2RoomBuilderStereofft = "4096" // 4096 (3) – 4096
	LspPlugInPluginsLv2RoomBuilderStereofft8192 LspPlugInPluginsLv2RoomBuilderStereofft = "8192" // 8192 (4) – 8192
	LspPlugInPluginsLv2RoomBuilderStereofft16384 LspPlugInPluginsLv2RoomBuilderStereofft = "16384" // 16384 (5) – 16384
	LspPlugInPluginsLv2RoomBuilderStereofft32767 LspPlugInPluginsLv2RoomBuilderStereofft = "32767" // 32767 (6) – 32767
	LspPlugInPluginsLv2RoomBuilderStereofft65536 LspPlugInPluginsLv2RoomBuilderStereofft = "65536" // 65536 (7) – 65536
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd3Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd3 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd38Directional LspPlugInPluginsLv2RoomBuilderStereoscsd3 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereocsf3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocsf3None LspPlugInPluginsLv2RoomBuilderStereocsf3 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample0 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample1 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample2 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample3 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample4 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample5 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample6 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderStereocsf3Sample7 LspPlugInPluginsLv2RoomBuilderStereocsf3 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderStereoscf4 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf4Any LspPlugInPluginsLv2RoomBuilderStereoscf4 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf40 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf41 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf42 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf43 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf44 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf45 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf46 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf47 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf48 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf49 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf410 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf411 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf412 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf413 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf414 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf415 LspPlugInPluginsLv2RoomBuilderStereoscf4 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd6 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd6Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd6 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd6Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd6 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd6Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd6 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd6Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd6 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd68Directional LspPlugInPluginsLv2RoomBuilderStereoscmd6 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd6Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd6 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereocsel string

const (
	LspPlugInPluginsLv2RoomBuilderStereocsel0 LspPlugInPluginsLv2RoomBuilderStereocsel = "0" // 0 (0) – 0
	LspPlugInPluginsLv2RoomBuilderStereocsel1 LspPlugInPluginsLv2RoomBuilderStereocsel = "1" // 1 (1) – 1
	LspPlugInPluginsLv2RoomBuilderStereocsel2 LspPlugInPluginsLv2RoomBuilderStereocsel = "2" // 2 (2) – 2
	LspPlugInPluginsLv2RoomBuilderStereocsel3 LspPlugInPluginsLv2RoomBuilderStereocsel = "3" // 3 (3) – 3
	LspPlugInPluginsLv2RoomBuilderStereocsel4 LspPlugInPluginsLv2RoomBuilderStereocsel = "4" // 4 (4) – 4
	LspPlugInPluginsLv2RoomBuilderStereocsel5 LspPlugInPluginsLv2RoomBuilderStereocsel = "5" // 5 (5) – 5
	LspPlugInPluginsLv2RoomBuilderStereocsel6 LspPlugInPluginsLv2RoomBuilderStereocsel = "6" // 6 (6) – 6
	LspPlugInPluginsLv2RoomBuilderStereocsel7 LspPlugInPluginsLv2RoomBuilderStereocsel = "7" // 7 (7) – 7
)

type LspPlugInPluginsLv2RoomBuilderStereosccf0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf0Mono LspPlugInPluginsLv2RoomBuilderStereosccf0 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf0Xy LspPlugInPluginsLv2RoomBuilderStereosccf0 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf0Ab LspPlugInPluginsLv2RoomBuilderStereosccf0 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf0Ortf LspPlugInPluginsLv2RoomBuilderStereosccf0 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf0Ms LspPlugInPluginsLv2RoomBuilderStereosccf0 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereoscf1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf1Any LspPlugInPluginsLv2RoomBuilderStereoscf1 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf10 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf11 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf12 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf13 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf14 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf15 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf16 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf17 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf18 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf19 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf110 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf111 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf112 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf113 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf114 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf115 LspPlugInPluginsLv2RoomBuilderStereoscf1 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscf3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf3Any LspPlugInPluginsLv2RoomBuilderStereoscf3 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf30 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf31 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf32 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf33 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf34 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf35 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf36 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf37 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf38 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf39 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf310 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf311 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf312 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf313 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf314 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf315 LspPlugInPluginsLv2RoomBuilderStereoscf3 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd0Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd0 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd0Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd0 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd0Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd0 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd0Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd0 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd08Directional LspPlugInPluginsLv2RoomBuilderStereoscmd0 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd0Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd0 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereosscf4 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf4Triangle LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf4Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf4Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf4Box LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf4Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf4Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf4Cone LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf4Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf4Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf4FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf4CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf4SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf4 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereosccf4 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf4Mono LspPlugInPluginsLv2RoomBuilderStereosccf4 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf4Xy LspPlugInPluginsLv2RoomBuilderStereosccf4 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf4Ab LspPlugInPluginsLv2RoomBuilderStereosccf4 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf4Ortf LspPlugInPluginsLv2RoomBuilderStereosccf4 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf4Ms LspPlugInPluginsLv2RoomBuilderStereosccf4 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereosccf6 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf6Mono LspPlugInPluginsLv2RoomBuilderStereosccf6 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf6Xy LspPlugInPluginsLv2RoomBuilderStereosccf6 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf6Ab LspPlugInPluginsLv2RoomBuilderStereosccf6 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf6Ortf LspPlugInPluginsLv2RoomBuilderStereosccf6 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf6Ms LspPlugInPluginsLv2RoomBuilderStereosccf6 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereoeditor string

const (
	LspPlugInPluginsLv2RoomBuilderStereoeditorSourceEditor LspPlugInPluginsLv2RoomBuilderStereoeditor = "Source editor" // Source editor (0) – Source editor
	LspPlugInPluginsLv2RoomBuilderStereoeditorCaptureEditor LspPlugInPluginsLv2RoomBuilderStereoeditor = "Capture editor" // Capture editor (1) – Capture editor
	LspPlugInPluginsLv2RoomBuilderStereoeditorObjectEditor LspPlugInPluginsLv2RoomBuilderStereoeditor = "Object editor" // Object editor (2) – Object editor
	LspPlugInPluginsLv2RoomBuilderStereoeditorMaterialEditor LspPlugInPluginsLv2RoomBuilderStereoeditor = "Material editor" // Material editor (3) – Material editor
)

type LspPlugInPluginsLv2RoomBuilderStereocsf1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocsf1None LspPlugInPluginsLv2RoomBuilderStereocsf1 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample0 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample1 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample2 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample3 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample4 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample5 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample6 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderStereocsf1Sample7 LspPlugInPluginsLv2RoomBuilderStereocsf1 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderStereocsf2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocsf2None LspPlugInPluginsLv2RoomBuilderStereocsf2 = "None" // None (0) – None
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample0 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample1 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample2 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample3 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample4 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample5 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample6 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderStereocsf2Sample7 LspPlugInPluginsLv2RoomBuilderStereocsf2 = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderStereoscl0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl0Any LspPlugInPluginsLv2RoomBuilderStereoscl0 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl00 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl01 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl02 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl03 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl04 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl05 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl06 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl07 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl08 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl09 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl010 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl011 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl012 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl013 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl014 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl015 LspPlugInPluginsLv2RoomBuilderStereoscl0 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscl5 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl5Any LspPlugInPluginsLv2RoomBuilderStereoscl5 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl50 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl51 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl52 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl53 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl54 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl55 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl56 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl57 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl58 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl59 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl510 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl511 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl512 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl513 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl514 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl515 LspPlugInPluginsLv2RoomBuilderStereoscl5 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd2Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd2 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd2Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd2 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd2Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd2 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd2Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd2 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd28Directional LspPlugInPluginsLv2RoomBuilderStereoscmd2 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd2Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd2 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereosscf7 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf7Triangle LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf7Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf7Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf7Box LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf7Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf7Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf7Cone LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf7Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf7Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf7FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf7CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf7SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf7 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereohcm string

const (
	LspPlugInPluginsLv2RoomBuilderStereohcmOff LspPlugInPluginsLv2RoomBuilderStereohcm = "off" // off (0) – off
	LspPlugInPluginsLv2RoomBuilderStereohcm6DBoct LspPlugInPluginsLv2RoomBuilderStereohcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2RoomBuilderStereohcm12DBoct LspPlugInPluginsLv2RoomBuilderStereohcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2RoomBuilderStereohcm18DBoct LspPlugInPluginsLv2RoomBuilderStereohcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2RoomBuilderStereolcm string

const (
	LspPlugInPluginsLv2RoomBuilderStereolcmOff LspPlugInPluginsLv2RoomBuilderStereolcm = "off" // off (0) – off
	LspPlugInPluginsLv2RoomBuilderStereolcm6DBoct LspPlugInPluginsLv2RoomBuilderStereolcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2RoomBuilderStereolcm12DBoct LspPlugInPluginsLv2RoomBuilderStereolcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2RoomBuilderStereolcm18DBoct LspPlugInPluginsLv2RoomBuilderStereolcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2RoomBuilderStereoscf2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf2Any LspPlugInPluginsLv2RoomBuilderStereoscf2 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf20 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf21 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf22 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf23 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf24 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf25 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf26 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf27 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf28 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf29 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf210 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf211 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf212 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf213 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf214 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf215 LspPlugInPluginsLv2RoomBuilderStereoscf2 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscl7 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl7Any LspPlugInPluginsLv2RoomBuilderStereoscl7 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl70 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl71 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl72 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl73 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl74 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl75 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl76 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl77 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl78 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl79 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl710 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl711 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl712 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl713 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl714 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl715 LspPlugInPluginsLv2RoomBuilderStereoscl7 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd0Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd0 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd08Directional LspPlugInPluginsLv2RoomBuilderStereoscsd0 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd1Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd1 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd18Directional LspPlugInPluginsLv2RoomBuilderStereoscsd1 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereoifo string

const (
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardYUp LspPlugInPluginsLv2RoomBuilderStereoifo = "+X forward, +Y up" // +X forward, +Y up (0) – +X forward, +Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardZUp LspPlugInPluginsLv2RoomBuilderStereoifo = "+X forward, +Z up" // +X forward, +Z up (1) – +X forward, +Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardYUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "+X forward, -Y up" // +X forward, -Y up (2) – +X forward, -Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardZUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "+X forward, -Z up" // +X forward, -Z up (3) – +X forward, -Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardYUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-X forward, +Y up" // -X forward, +Y up (4) – -X forward, +Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardZUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-X forward, +Z up" // -X forward, +Z up (5) – -X forward, +Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardYUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-X forward, -Y up" // -X forward, -Y up (6) – -X forward, -Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoXForwardZUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-X forward, -Z up" // -X forward, -Z up (7) – -X forward, -Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardXUp LspPlugInPluginsLv2RoomBuilderStereoifo = "+Y forward, +X up" // +Y forward, +X up (8) – +Y forward, +X up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardZUp LspPlugInPluginsLv2RoomBuilderStereoifo = "+Y forward, +Z up" // +Y forward, +Z up (9) – +Y forward, +Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardXUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "+Y forward, -X up" // +Y forward, -X up (10) – +Y forward, -X up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardZUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "+Y forward, -Z up" // +Y forward, -Z up (11) – +Y forward, -Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardXUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Y forward, +X up" // -Y forward, +X up (12) – -Y forward, +X up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardZUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Y forward, +Z up" // -Y forward, +Z up (13) – -Y forward, +Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardXUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Y forward, -X up" // -Y forward, -X up (14) – -Y forward, -X up
	LspPlugInPluginsLv2RoomBuilderStereoifoYForwardZUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Y forward, -Z up" // -Y forward, -Z up (15) – -Y forward, -Z up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardXUp LspPlugInPluginsLv2RoomBuilderStereoifo = "+Z forward, +X up" // +Z forward, +X up (16) – +Z forward, +X up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardYUp LspPlugInPluginsLv2RoomBuilderStereoifo = "+Z forward, +Y up" // +Z forward, +Y up (17) – +Z forward, +Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardXUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "+Z forward, -X up" // +Z forward, -X up (18) – +Z forward, -X up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardYUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "+Z forward, -Y up" // +Z forward, -Y up (19) – +Z forward, -Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardXUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Z forward, +X up" // -Z forward, +X up (20) – -Z forward, +X up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardYUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Z forward, +Y up" // -Z forward, +Y up (21) – -Z forward, +Y up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardXUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Z forward, -X up" // -Z forward, -X up (22) – -Z forward, -X up
	LspPlugInPluginsLv2RoomBuilderStereoifoZForwardYUp1 LspPlugInPluginsLv2RoomBuilderStereoifo = "-Z forward, -Y up" // -Z forward, -Y up (23) – -Z forward, -Y up
)

type LspPlugInPluginsLv2RoomBuilderStereosccf7 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf7Mono LspPlugInPluginsLv2RoomBuilderStereosccf7 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf7Xy LspPlugInPluginsLv2RoomBuilderStereosccf7 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf7Ab LspPlugInPluginsLv2RoomBuilderStereosccf7 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf7Ortf LspPlugInPluginsLv2RoomBuilderStereosccf7 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf7Ms LspPlugInPluginsLv2RoomBuilderStereosccf7 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereocst1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereocst1Left LspPlugInPluginsLv2RoomBuilderStereocst1 = "Left" // Left (0) – Left
	LspPlugInPluginsLv2RoomBuilderStereocst1Right LspPlugInPluginsLv2RoomBuilderStereocst1 = "Right" // Right (1) – Right
)

type LspPlugInPluginsLv2RoomBuilderStereosscf1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf1Triangle LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf1Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf1Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf1Box LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf1Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf1Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf1Cone LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf1Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf1Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf1FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf1CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf1SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf1 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereoscf6 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf6Any LspPlugInPluginsLv2RoomBuilderStereoscf6 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf60 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf61 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf62 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf63 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf64 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf65 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf66 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf67 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf68 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf69 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf610 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf611 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf612 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf613 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf614 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf615 LspPlugInPluginsLv2RoomBuilderStereoscf6 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscf7 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf7Any LspPlugInPluginsLv2RoomBuilderStereoscf7 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf70 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf71 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf72 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf73 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf74 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf75 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf76 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf77 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf78 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf79 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf710 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf711 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf712 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf713 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf714 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf715 LspPlugInPluginsLv2RoomBuilderStereoscf7 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd4 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd4Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd4 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd4Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd4 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd4Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd4 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd4Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd4 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd48Directional LspPlugInPluginsLv2RoomBuilderStereoscmd4 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd4Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd4 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereoscmd7 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscmd7Cardioid LspPlugInPluginsLv2RoomBuilderStereoscmd7 = "Cardioid" // Cardioid (0) – Cardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd7Supercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd7 = "Supercardioid" // Supercardioid (1) – Supercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd7Hypercardioid LspPlugInPluginsLv2RoomBuilderStereoscmd7 = "Hypercardioid" // Hypercardioid (2) – Hypercardioid
	LspPlugInPluginsLv2RoomBuilderStereoscmd7Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd7 = "Bidirectional" // Bidirectional (3) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscmd78Directional LspPlugInPluginsLv2RoomBuilderStereoscmd7 = "8-directional" // 8-directional (4) – 8-directional
	LspPlugInPluginsLv2RoomBuilderStereoscmd7Omnidirectional LspPlugInPluginsLv2RoomBuilderStereoscmd7 = "Omnidirectional" // Omnidirectional (5) – Omnidirectional
)

type LspPlugInPluginsLv2RoomBuilderStereosignal string

const (
	LspPlugInPluginsLv2RoomBuilderStereosignalConvolvers LspPlugInPluginsLv2RoomBuilderStereosignal = "Convolvers" // Convolvers (0) – Convolvers
	LspPlugInPluginsLv2RoomBuilderStereosignalIrEqualizer LspPlugInPluginsLv2RoomBuilderStereosignal = "IR Equalizer" // IR Equalizer (1) – IR Equalizer
)

type LspPlugInPluginsLv2RoomBuilderStereoview string

const (
	LspPlugInPluginsLv2RoomBuilderStereoviewRoomBrowser LspPlugInPluginsLv2RoomBuilderStereoview = "Room browser" // Room browser (0) – Room browser
	LspPlugInPluginsLv2RoomBuilderStereoviewSample0 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 0" // Sample 0 (1) – Sample 0
	LspPlugInPluginsLv2RoomBuilderStereoviewSample1 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 1" // Sample 1 (2) – Sample 1
	LspPlugInPluginsLv2RoomBuilderStereoviewSample2 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 2" // Sample 2 (3) – Sample 2
	LspPlugInPluginsLv2RoomBuilderStereoviewSample3 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 3" // Sample 3 (4) – Sample 3
	LspPlugInPluginsLv2RoomBuilderStereoviewSample4 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 4" // Sample 4 (5) – Sample 4
	LspPlugInPluginsLv2RoomBuilderStereoviewSample5 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 5" // Sample 5 (6) – Sample 5
	LspPlugInPluginsLv2RoomBuilderStereoviewSample6 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 6" // Sample 6 (7) – Sample 6
	LspPlugInPluginsLv2RoomBuilderStereoviewSample7 LspPlugInPluginsLv2RoomBuilderStereoview = "Sample 7" // Sample 7 (8) – Sample 7
)

type LspPlugInPluginsLv2RoomBuilderStereosccf1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf1Mono LspPlugInPluginsLv2RoomBuilderStereosccf1 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf1Xy LspPlugInPluginsLv2RoomBuilderStereosccf1 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf1Ab LspPlugInPluginsLv2RoomBuilderStereosccf1 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf1Ortf LspPlugInPluginsLv2RoomBuilderStereosccf1 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf1Ms LspPlugInPluginsLv2RoomBuilderStereosccf1 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereoscf0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscf0Any LspPlugInPluginsLv2RoomBuilderStereoscf0 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscf00 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscf01 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscf02 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscf03 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscf04 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscf05 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscf06 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscf07 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscf08 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscf09 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscf010 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscf011 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscf012 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscf013 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscf014 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscf015 LspPlugInPluginsLv2RoomBuilderStereoscf0 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereoscl1 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl1Any LspPlugInPluginsLv2RoomBuilderStereoscl1 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl10 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl11 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl12 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl13 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl14 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl15 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl16 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl17 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl18 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl19 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl110 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl111 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl112 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl113 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl114 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl115 LspPlugInPluginsLv2RoomBuilderStereoscl1 = "15" // 15 (16) – 15
)

type LspPlugInPluginsLv2RoomBuilderStereosccf3 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf3Mono LspPlugInPluginsLv2RoomBuilderStereosccf3 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf3Xy LspPlugInPluginsLv2RoomBuilderStereosccf3 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf3Ab LspPlugInPluginsLv2RoomBuilderStereosccf3 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf3Ortf LspPlugInPluginsLv2RoomBuilderStereosccf3 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf3Ms LspPlugInPluginsLv2RoomBuilderStereosccf3 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereosccf5 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf5Mono LspPlugInPluginsLv2RoomBuilderStereosccf5 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf5Xy LspPlugInPluginsLv2RoomBuilderStereosccf5 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf5Ab LspPlugInPluginsLv2RoomBuilderStereosccf5 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf5Ortf LspPlugInPluginsLv2RoomBuilderStereosccf5 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf5Ms LspPlugInPluginsLv2RoomBuilderStereosccf5 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereoscsd7 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscsd7Bidirectional LspPlugInPluginsLv2RoomBuilderStereoscsd7 = "Bidirectional" // Bidirectional (0) – Bidirectional
	LspPlugInPluginsLv2RoomBuilderStereoscsd78Directional LspPlugInPluginsLv2RoomBuilderStereoscsd7 = "8-directional" // 8-directional (1) – 8-directional
)

type LspPlugInPluginsLv2RoomBuilderStereosscf0 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf0Triangle LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf0Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf0Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf0Box LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf0Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf0Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf0Cone LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf0Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf0Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf0FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf0CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf0SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf0 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereosscf6 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosscf6Triangle LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Triangle" // Triangle (0) – Triangle
	LspPlugInPluginsLv2RoomBuilderStereosscf6Tetrahedron LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Tetrahedron" // Tetrahedron (1) – Tetrahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf6Octahedron LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Octahedron" // Octahedron (2) – Octahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf6Box LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Box" // Box (3) – Box
	LspPlugInPluginsLv2RoomBuilderStereosscf6Icosahedron LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Icosahedron" // Icosahedron (4) – Icosahedron
	LspPlugInPluginsLv2RoomBuilderStereosscf6Cylinder LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Cylinder" // Cylinder (5) – Cylinder
	LspPlugInPluginsLv2RoomBuilderStereosscf6Cone LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Cone" // Cone (6) – Cone
	LspPlugInPluginsLv2RoomBuilderStereosscf6Octasphere LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Octasphere" // Octasphere (7) – Octasphere
	LspPlugInPluginsLv2RoomBuilderStereosscf6Icosphere LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Icosphere" // Icosphere (8) – Icosphere
	LspPlugInPluginsLv2RoomBuilderStereosscf6FlatSpot LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Flat spot" // Flat spot (9) – Flat spot
	LspPlugInPluginsLv2RoomBuilderStereosscf6CylindricalSpot LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Cylindrical spot" // Cylindrical spot (10) – Cylindrical spot
	LspPlugInPluginsLv2RoomBuilderStereosscf6SphericalSpot LspPlugInPluginsLv2RoomBuilderStereosscf6 = "Spherical spot" // Spherical spot (11) – Spherical spot
)

type LspPlugInPluginsLv2RoomBuilderStereosccf2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereosccf2Mono LspPlugInPluginsLv2RoomBuilderStereosccf2 = "Mono" // Mono (0) – Mono
	LspPlugInPluginsLv2RoomBuilderStereosccf2Xy LspPlugInPluginsLv2RoomBuilderStereosccf2 = "XY" // XY (1) – XY
	LspPlugInPluginsLv2RoomBuilderStereosccf2Ab LspPlugInPluginsLv2RoomBuilderStereosccf2 = "AB" // AB (2) – AB
	LspPlugInPluginsLv2RoomBuilderStereosccf2Ortf LspPlugInPluginsLv2RoomBuilderStereosccf2 = "ORTF" // ORTF (3) – ORTF
	LspPlugInPluginsLv2RoomBuilderStereosccf2Ms LspPlugInPluginsLv2RoomBuilderStereosccf2 = "MS" // MS (4) – MS
)

type LspPlugInPluginsLv2RoomBuilderStereoscl2 string

const (
	LspPlugInPluginsLv2RoomBuilderStereoscl2Any LspPlugInPluginsLv2RoomBuilderStereoscl2 = "Any" // Any (0) – Any
	LspPlugInPluginsLv2RoomBuilderStereoscl20 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "0" // 0 (1) – 0
	LspPlugInPluginsLv2RoomBuilderStereoscl21 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "1" // 1 (2) – 1
	LspPlugInPluginsLv2RoomBuilderStereoscl22 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "2" // 2 (3) – 2
	LspPlugInPluginsLv2RoomBuilderStereoscl23 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "3" // 3 (4) – 3
	LspPlugInPluginsLv2RoomBuilderStereoscl24 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "4" // 4 (5) – 4
	LspPlugInPluginsLv2RoomBuilderStereoscl25 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "5" // 5 (6) – 5
	LspPlugInPluginsLv2RoomBuilderStereoscl26 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "6" // 6 (7) – 6
	LspPlugInPluginsLv2RoomBuilderStereoscl27 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "7" // 7 (8) – 7
	LspPlugInPluginsLv2RoomBuilderStereoscl28 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "8" // 8 (9) – 8
	LspPlugInPluginsLv2RoomBuilderStereoscl29 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "9" // 9 (10) – 9
	LspPlugInPluginsLv2RoomBuilderStereoscl210 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "10" // 10 (11) – 10
	LspPlugInPluginsLv2RoomBuilderStereoscl211 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "11" // 11 (12) – 11
	LspPlugInPluginsLv2RoomBuilderStereoscl212 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "12" // 12 (13) – 12
	LspPlugInPluginsLv2RoomBuilderStereoscl213 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "13" // 13 (14) – 13
	LspPlugInPluginsLv2RoomBuilderStereoscl214 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "14" // 14 (15) – 14
	LspPlugInPluginsLv2RoomBuilderStereoscl215 LspPlugInPluginsLv2RoomBuilderStereoscl2 = "15" // 15 (16) – 15
)

