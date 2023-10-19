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

type LspPlugInPluginsLv2ImpulseResponsesStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ImpulseResponsesStereo() (*LspPlugInPluginsLv2ImpulseResponsesStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-impulse-responses-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseResponsesStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ImpulseResponsesStereoWithName(name string) (*LspPlugInPluginsLv2ImpulseResponsesStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-impulse-responses-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseResponsesStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// ca-l (bool)
//
// Channel activity Left

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetCaL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// ca-r (bool)
//
// Channel activity Right

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetCaR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// cs-l (GstLspPlugInPluginsLv2ImpulseResponsesStereocsL)
//
// Channel source Left

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetCsL() (interface{}, error) {
	return e.Element.GetProperty("cs-l")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetCsL(value interface{}) error {
	return e.Element.SetProperty("cs-l", value)
}

// cs-r (GstLspPlugInPluginsLv2ImpulseResponsesStereocsR)
//
// Channel source Right

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetCsR() (interface{}, error) {
	return e.Element.GetProperty("cs-r")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetCsR(value interface{}) error {
	return e.Element.SetProperty("cs-r", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// eq-0 (float32)
//
// Band 50Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq0(value float32) error {
	return e.Element.SetProperty("eq-0", value)
}

// eq-1 (float32)
//
// Band 107Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq1(value float32) error {
	return e.Element.SetProperty("eq-1", value)
}

// eq-2 (float32)
//
// Band 227Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq2(value float32) error {
	return e.Element.SetProperty("eq-2", value)
}

// eq-3 (float32)
//
// Band 484Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq3(value float32) error {
	return e.Element.SetProperty("eq-3", value)
}

// eq-4 (float32)
//
// Band 1 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq4(value float32) error {
	return e.Element.SetProperty("eq-4", value)
}

// eq-5 (float32)
//
// Band 2.2 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq5(value float32) error {
	return e.Element.SetProperty("eq-5", value)
}

// eq-6 (float32)
//
// Band 4.7 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq6(value float32) error {
	return e.Element.SetProperty("eq-6", value)
}

// eq-7 (float32)
//
// Band 10 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetEq7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetEq7(value float32) error {
	return e.Element.SetProperty("eq-7", value)
}

// fft (GstLspPlugInPluginsLv2ImpulseResponsesStereofft)
//
// FFT size

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fsel (GstLspPlugInPluginsLv2ImpulseResponsesStereofsel)
//
// File selector

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hcf (float32)
//
// High-cut frequency

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetHcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetHcf(value float32) error {
	return e.Element.SetProperty("hcf", value)
}

// hcm (GstLspPlugInPluginsLv2ImpulseResponsesStereohcm)
//
// High-cut mode

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetHcm() (interface{}, error) {
	return e.Element.GetProperty("hcm")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetHcm(value interface{}) error {
	return e.Element.SetProperty("hcm", value)
}

// ifi0 (float32)
//
// Fade in 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfi0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIfi0(value float32) error {
	return e.Element.SetProperty("ifi0", value)
}

// ifi1 (float32)
//
// Fade in 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfi1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIfi1(value float32) error {
	return e.Element.SetProperty("ifi1", value)
}

// ifl0 (float32)
//
// Impulse length 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfl0() (float32, error) {
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
// Impulse length 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfl1() (float32, error) {
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

// ifo0 (float32)
//
// Fade out 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfo0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIfo0(value float32) error {
	return e.Element.SetProperty("ifo0", value)
}

// ifo1 (float32)
//
// Fade out 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfo1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIfo1(value float32) error {
	return e.Element.SetProperty("ifo1", value)
}

// ifs0 (int)
//
// Load status 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfs0() (int, error) {
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
// Load status 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIfs1() (int, error) {
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

// ihc0 (float32)
//
// Head cut 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIhc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIhc0(value float32) error {
	return e.Element.SetProperty("ihc0", value)
}

// ihc1 (float32)
//
// Head cut 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIhc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIhc1(value float32) error {
	return e.Element.SetProperty("ihc1", value)
}

// ils0 (bool)
//
// Impulse listen 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIls0() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIls0(value bool) error {
	return e.Element.SetProperty("ils0", value)
}

// ils1 (bool)
//
// Impulse listen 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetIls1() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetIls1(value bool) error {
	return e.Element.SetProperty("ils1", value)
}

// itc0 (float32)
//
// Tail cut 1

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetItc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetItc0(value float32) error {
	return e.Element.SetProperty("itc0", value)
}

// itc1 (float32)
//
// Tail cut 2

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetItc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetItc1(value float32) error {
	return e.Element.SetProperty("itc1", value)
}

// lcf (float32)
//
// Low-cut frequency

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetLcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetLcf(value float32) error {
	return e.Element.SetProperty("lcf", value)
}

// lcm (GstLspPlugInPluginsLv2ImpulseResponsesStereolcm)
//
// Low-cut mode

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetLcm() (interface{}, error) {
	return e.Element.GetProperty("lcm")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetLcm(value interface{}) error {
	return e.Element.SetProperty("lcm", value)
}

// mk-l (float32)
//
// Makeup gain Left

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetMkL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetMkL(value float32) error {
	return e.Element.SetProperty("mk-l", value)
}

// mk-r (float32)
//
// Makeup gain Right

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetMkR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetMkR(value float32) error {
	return e.Element.SetProperty("mk-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetOutLatency() (int, error) {
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

// pd-l (float32)
//
// Pre-delay Left

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetPdL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetPdL(value float32) error {
	return e.Element.SetProperty("pd-l", value)
}

// pd-r (float32)
//
// Pre-delay Right

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetPdR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pd-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetPdR(value float32) error {
	return e.Element.SetProperty("pd-r", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wpp (bool)
//
// Wet post-process

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) GetWpp() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesStereo) SetWpp(value bool) error {
	return e.Element.SetProperty("wpp", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ImpulseResponsesStereofft string

const (
	LspPlugInPluginsLv2ImpulseResponsesStereofft512 LspPlugInPluginsLv2ImpulseResponsesStereofft = "512" // 512 (0) – 512
	LspPlugInPluginsLv2ImpulseResponsesStereofft1024 LspPlugInPluginsLv2ImpulseResponsesStereofft = "1024" // 1024 (1) – 1024
	LspPlugInPluginsLv2ImpulseResponsesStereofft2048 LspPlugInPluginsLv2ImpulseResponsesStereofft = "2048" // 2048 (2) – 2048
	LspPlugInPluginsLv2ImpulseResponsesStereofft4096 LspPlugInPluginsLv2ImpulseResponsesStereofft = "4096" // 4096 (3) – 4096
	LspPlugInPluginsLv2ImpulseResponsesStereofft8192 LspPlugInPluginsLv2ImpulseResponsesStereofft = "8192" // 8192 (4) – 8192
	LspPlugInPluginsLv2ImpulseResponsesStereofft16384 LspPlugInPluginsLv2ImpulseResponsesStereofft = "16384" // 16384 (5) – 16384
	LspPlugInPluginsLv2ImpulseResponsesStereofft32767 LspPlugInPluginsLv2ImpulseResponsesStereofft = "32767" // 32767 (6) – 32767
	LspPlugInPluginsLv2ImpulseResponsesStereofft65536 LspPlugInPluginsLv2ImpulseResponsesStereofft = "65536" // 65536 (7) – 65536
)

type LspPlugInPluginsLv2ImpulseResponsesStereofsel string

const (
	LspPlugInPluginsLv2ImpulseResponsesStereofselFile1 LspPlugInPluginsLv2ImpulseResponsesStereofsel = "File 1" // File 1 (0) – File 1
	LspPlugInPluginsLv2ImpulseResponsesStereofselFile2 LspPlugInPluginsLv2ImpulseResponsesStereofsel = "File 2" // File 2 (1) – File 2
)

type LspPlugInPluginsLv2ImpulseResponsesStereohcm string

const (
	LspPlugInPluginsLv2ImpulseResponsesStereohcmOff LspPlugInPluginsLv2ImpulseResponsesStereohcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseResponsesStereohcm6DBoct LspPlugInPluginsLv2ImpulseResponsesStereohcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesStereohcm12DBoct LspPlugInPluginsLv2ImpulseResponsesStereohcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesStereohcm18DBoct LspPlugInPluginsLv2ImpulseResponsesStereohcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseResponsesStereolcm string

const (
	LspPlugInPluginsLv2ImpulseResponsesStereolcmOff LspPlugInPluginsLv2ImpulseResponsesStereolcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseResponsesStereolcm6DBoct LspPlugInPluginsLv2ImpulseResponsesStereolcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesStereolcm12DBoct LspPlugInPluginsLv2ImpulseResponsesStereolcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesStereolcm18DBoct LspPlugInPluginsLv2ImpulseResponsesStereolcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseResponsesStereocsL string

const (
	LspPlugInPluginsLv2ImpulseResponsesStereocsLNone LspPlugInPluginsLv2ImpulseResponsesStereocsL = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseResponsesStereocsLFile1Left LspPlugInPluginsLv2ImpulseResponsesStereocsL = "File 1 Left" // File 1 Left (1) – File 1 Left
	LspPlugInPluginsLv2ImpulseResponsesStereocsLFile1Right LspPlugInPluginsLv2ImpulseResponsesStereocsL = "File 1 Right" // File 1 Right (2) – File 1 Right
	LspPlugInPluginsLv2ImpulseResponsesStereocsLFile2Left LspPlugInPluginsLv2ImpulseResponsesStereocsL = "File 2 Left" // File 2 Left (3) – File 2 Left
	LspPlugInPluginsLv2ImpulseResponsesStereocsLFile2Right LspPlugInPluginsLv2ImpulseResponsesStereocsL = "File 2 Right" // File 2 Right (4) – File 2 Right
)

type LspPlugInPluginsLv2ImpulseResponsesStereocsR string

const (
	LspPlugInPluginsLv2ImpulseResponsesStereocsRNone LspPlugInPluginsLv2ImpulseResponsesStereocsR = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseResponsesStereocsRFile1Left LspPlugInPluginsLv2ImpulseResponsesStereocsR = "File 1 Left" // File 1 Left (1) – File 1 Left
	LspPlugInPluginsLv2ImpulseResponsesStereocsRFile1Right LspPlugInPluginsLv2ImpulseResponsesStereocsR = "File 1 Right" // File 1 Right (2) – File 1 Right
	LspPlugInPluginsLv2ImpulseResponsesStereocsRFile2Left LspPlugInPluginsLv2ImpulseResponsesStereocsR = "File 2 Left" // File 2 Left (3) – File 2 Left
	LspPlugInPluginsLv2ImpulseResponsesStereocsRFile2Right LspPlugInPluginsLv2ImpulseResponsesStereocsR = "File 2 Right" // File 2 Right (4) – File 2 Right
)

