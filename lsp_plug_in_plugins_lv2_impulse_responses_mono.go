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

type LspPlugInPluginsLv2ImpulseResponsesMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ImpulseResponsesMono() (*LspPlugInPluginsLv2ImpulseResponsesMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-impulse-responses-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseResponsesMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ImpulseResponsesMonoWithName(name string) (*LspPlugInPluginsLv2ImpulseResponsesMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-impulse-responses-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ImpulseResponsesMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// ca (bool)
//
// Channel activity

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetCa() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ca")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// cs (GstLspPlugInPluginsLv2ImpulseResponsesMonocs)
//
// Channel source

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetCs() (interface{}, error) {
	return e.Element.GetProperty("cs")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetCs(value interface{}) error {
	return e.Element.SetProperty("cs", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// eq-0 (float32)
//
// Band 50Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq0() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq0(value float32) error {
	return e.Element.SetProperty("eq-0", value)
}

// eq-1 (float32)
//
// Band 107Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq1() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq1(value float32) error {
	return e.Element.SetProperty("eq-1", value)
}

// eq-2 (float32)
//
// Band 227Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq2() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq2(value float32) error {
	return e.Element.SetProperty("eq-2", value)
}

// eq-3 (float32)
//
// Band 484Hz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq3() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq3(value float32) error {
	return e.Element.SetProperty("eq-3", value)
}

// eq-4 (float32)
//
// Band 1 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq4() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq4(value float32) error {
	return e.Element.SetProperty("eq-4", value)
}

// eq-5 (float32)
//
// Band 2.2 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq5() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq5(value float32) error {
	return e.Element.SetProperty("eq-5", value)
}

// eq-6 (float32)
//
// Band 4.7 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq6() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq6(value float32) error {
	return e.Element.SetProperty("eq-6", value)
}

// eq-7 (float32)
//
// Band 10 kHz gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetEq7() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetEq7(value float32) error {
	return e.Element.SetProperty("eq-7", value)
}

// fft (GstLspPlugInPluginsLv2ImpulseResponsesMonofft)
//
// FFT size

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hcf (float32)
//
// High-cut frequency

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetHcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetHcf(value float32) error {
	return e.Element.SetProperty("hcf", value)
}

// hcm (GstLspPlugInPluginsLv2ImpulseResponsesMonohcm)
//
// High-cut mode

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetHcm() (interface{}, error) {
	return e.Element.GetProperty("hcm")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetHcm(value interface{}) error {
	return e.Element.SetProperty("hcm", value)
}

// ifi (float32)
//
// Fade in

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetIfi() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifi")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetIfi(value float32) error {
	return e.Element.SetProperty("ifi", value)
}

// ifl (float32)
//
// Impulse length

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetIfl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ifo (float32)
//
// Fade out

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetIfo() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ifo")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetIfo(value float32) error {
	return e.Element.SetProperty("ifo", value)
}

// ifs (int)
//
// Load status

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetIfs() (int, error) {
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

// ihc (float32)
//
// Head cut

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetIhc() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ihc")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetIhc(value float32) error {
	return e.Element.SetProperty("ihc", value)
}

// ils (bool)
//
// Impulse listen

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetIls() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ils")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetIls(value bool) error {
	return e.Element.SetProperty("ils", value)
}

// itc (float32)
//
// Tail cut

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetItc() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("itc")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetItc(value float32) error {
	return e.Element.SetProperty("itc", value)
}

// lcf (float32)
//
// Low-cut frequency

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetLcf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetLcf(value float32) error {
	return e.Element.SetProperty("lcf", value)
}

// lcm (GstLspPlugInPluginsLv2ImpulseResponsesMonolcm)
//
// Low-cut mode

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetLcm() (interface{}, error) {
	return e.Element.GetProperty("lcm")
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetLcm(value interface{}) error {
	return e.Element.SetProperty("lcm", value)
}

// mk (float32)
//
// Makeup gain

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetMk() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetMk(value float32) error {
	return e.Element.SetProperty("mk", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetPd() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetPd(value float32) error {
	return e.Element.SetProperty("pd", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wpp (bool)
//
// Wet post-process

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) GetWpp() (bool, error) {
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

func (e *LspPlugInPluginsLv2ImpulseResponsesMono) SetWpp(value bool) error {
	return e.Element.SetProperty("wpp", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ImpulseResponsesMonofft string

const (
	LspPlugInPluginsLv2ImpulseResponsesMonofft512 LspPlugInPluginsLv2ImpulseResponsesMonofft = "512" // 512 (0) – 512
	LspPlugInPluginsLv2ImpulseResponsesMonofft1024 LspPlugInPluginsLv2ImpulseResponsesMonofft = "1024" // 1024 (1) – 1024
	LspPlugInPluginsLv2ImpulseResponsesMonofft2048 LspPlugInPluginsLv2ImpulseResponsesMonofft = "2048" // 2048 (2) – 2048
	LspPlugInPluginsLv2ImpulseResponsesMonofft4096 LspPlugInPluginsLv2ImpulseResponsesMonofft = "4096" // 4096 (3) – 4096
	LspPlugInPluginsLv2ImpulseResponsesMonofft8192 LspPlugInPluginsLv2ImpulseResponsesMonofft = "8192" // 8192 (4) – 8192
	LspPlugInPluginsLv2ImpulseResponsesMonofft16384 LspPlugInPluginsLv2ImpulseResponsesMonofft = "16384" // 16384 (5) – 16384
	LspPlugInPluginsLv2ImpulseResponsesMonofft32767 LspPlugInPluginsLv2ImpulseResponsesMonofft = "32767" // 32767 (6) – 32767
	LspPlugInPluginsLv2ImpulseResponsesMonofft65536 LspPlugInPluginsLv2ImpulseResponsesMonofft = "65536" // 65536 (7) – 65536
)

type LspPlugInPluginsLv2ImpulseResponsesMonohcm string

const (
	LspPlugInPluginsLv2ImpulseResponsesMonohcmOff LspPlugInPluginsLv2ImpulseResponsesMonohcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseResponsesMonohcm6DBoct LspPlugInPluginsLv2ImpulseResponsesMonohcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesMonohcm12DBoct LspPlugInPluginsLv2ImpulseResponsesMonohcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesMonohcm18DBoct LspPlugInPluginsLv2ImpulseResponsesMonohcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseResponsesMonolcm string

const (
	LspPlugInPluginsLv2ImpulseResponsesMonolcmOff LspPlugInPluginsLv2ImpulseResponsesMonolcm = "off" // off (0) – off
	LspPlugInPluginsLv2ImpulseResponsesMonolcm6DBoct LspPlugInPluginsLv2ImpulseResponsesMonolcm = "6 dB/oct" // 6 dB/oct (1) – 6 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesMonolcm12DBoct LspPlugInPluginsLv2ImpulseResponsesMonolcm = "12 dB/oct" // 12 dB/oct (2) – 12 dB/oct
	LspPlugInPluginsLv2ImpulseResponsesMonolcm18DBoct LspPlugInPluginsLv2ImpulseResponsesMonolcm = "18 dB/oct" // 18 dB/oct (3) – 18 dB/oct
)

type LspPlugInPluginsLv2ImpulseResponsesMonocs string

const (
	LspPlugInPluginsLv2ImpulseResponsesMonocsNone LspPlugInPluginsLv2ImpulseResponsesMonocs = "None" // None (0) – None
	LspPlugInPluginsLv2ImpulseResponsesMonocsLeft LspPlugInPluginsLv2ImpulseResponsesMonocs = "Left" // Left (1) – Left
	LspPlugInPluginsLv2ImpulseResponsesMonocsRight LspPlugInPluginsLv2ImpulseResponsesMonocs = "Right" // Right (2) – Right
)

