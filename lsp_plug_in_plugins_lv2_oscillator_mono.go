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

type LspPlugInPluginsLv2OscillatorMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2OscillatorMono() (*LspPlugInPluginsLv2OscillatorMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-oscillator-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2OscillatorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2OscillatorMonoWithName(name string) (*LspPlugInPluginsLv2OscillatorMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-oscillator-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2OscillatorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2OscillatorMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2OscillatorMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// dcoff (float32)
//
// DC Offset

func (e *LspPlugInPluginsLv2OscillatorMono) GetDcoff() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dcoff")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetDcoff(value float32) error {
	return e.Element.SetProperty("dcoff", value)
}

// freq (float32)
//
// Frequency

func (e *LspPlugInPluginsLv2OscillatorMono) GetFreq() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetFreq(value float32) error {
	return e.Element.SetProperty("freq", value)
}

// gain (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2OscillatorMono) GetGain() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetGain(value float32) error {
	return e.Element.SetProperty("gain", value)
}

// iniph (float32)
//
// Initial Phase

func (e *LspPlugInPluginsLv2OscillatorMono) GetIniph() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("iniph")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetIniph(value float32) error {
	return e.Element.SetProperty("iniph", value)
}

// invps (bool)
//
// Invert Parabolic Signal

func (e *LspPlugInPluginsLv2OscillatorMono) GetInvps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetInvps(value bool) error {
	return e.Element.SetProperty("invps", value)
}

// invss (bool)
//
// Invert Squared Sinusoids

func (e *LspPlugInPluginsLv2OscillatorMono) GetInvss() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invss")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetInvss(value bool) error {
	return e.Element.SetProperty("invss", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2OscillatorMono) GetOutLatency() (int, error) {
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

// pwdth (float32)
//
// Width

func (e *LspPlugInPluginsLv2OscillatorMono) GetPwdth() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pwdth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetPwdth(value float32) error {
	return e.Element.SetProperty("pwdth", value)
}

// rdtrt (float32)
//
// Duty Ratio

func (e *LspPlugInPluginsLv2OscillatorMono) GetRdtrt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rdtrt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetRdtrt(value float32) error {
	return e.Element.SetProperty("rdtrt", value)
}

// scf (GstLspPlugInPluginsLv2OscillatorMonoscf)
//
// Function

func (e *LspPlugInPluginsLv2OscillatorMono) GetScf() (interface{}, error) {
	return e.Element.GetProperty("scf")
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetScf(value interface{}) error {
	return e.Element.SetProperty("scf", value)
}

// scm (GstLspPlugInPluginsLv2OscillatorMonoscm)
//
// Operation Mode

func (e *LspPlugInPluginsLv2OscillatorMono) GetScm() (interface{}, error) {
	return e.Element.GetProperty("scm")
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetScm(value interface{}) error {
	return e.Element.SetProperty("scm", value)
}

// scom (GstLspPlugInPluginsLv2OscillatorMonoscom)
//
// Oversampler Mode

func (e *LspPlugInPluginsLv2OscillatorMono) GetScom() (interface{}, error) {
	return e.Element.GetProperty("scom")
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetScom(value interface{}) error {
	return e.Element.SetProperty("scom", value)
}

// scr (GstLspPlugInPluginsLv2OscillatorMonoscr)
//
// DC Reference

func (e *LspPlugInPluginsLv2OscillatorMono) GetScr() (interface{}, error) {
	return e.Element.GetProperty("scr")
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetScr(value interface{}) error {
	return e.Element.SetProperty("scr", value)
}

// swdth (float32)
//
// Width

func (e *LspPlugInPluginsLv2OscillatorMono) GetSwdth() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("swdth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetSwdth(value float32) error {
	return e.Element.SetProperty("swdth", value)
}

// tflrt (float32)
//
// Fall Ratio

func (e *LspPlugInPluginsLv2OscillatorMono) GetTflrt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tflrt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetTflrt(value float32) error {
	return e.Element.SetProperty("tflrt", value)
}

// tnwrt (float32)
//
// Negative Width

func (e *LspPlugInPluginsLv2OscillatorMono) GetTnwrt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tnwrt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetTnwrt(value float32) error {
	return e.Element.SetProperty("tnwrt", value)
}

// tpwrt (float32)
//
// Positive Width

func (e *LspPlugInPluginsLv2OscillatorMono) GetTpwrt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tpwrt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetTpwrt(value float32) error {
	return e.Element.SetProperty("tpwrt", value)
}

// trsrt (float32)
//
// Raise Ratio

func (e *LspPlugInPluginsLv2OscillatorMono) GetTrsrt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("trsrt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2OscillatorMono) SetTrsrt(value float32) error {
	return e.Element.SetProperty("trsrt", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2OscillatorMonoscf string

const (
	LspPlugInPluginsLv2OscillatorMonoscfSine LspPlugInPluginsLv2OscillatorMonoscf = "Sine" // Sine (0) – Sine
	LspPlugInPluginsLv2OscillatorMonoscfCosine LspPlugInPluginsLv2OscillatorMonoscf = "Cosine" // Cosine (1) – Cosine
	LspPlugInPluginsLv2OscillatorMonoscfSquaredSine LspPlugInPluginsLv2OscillatorMonoscf = "Squared Sine" // Squared Sine (2) – Squared Sine
	LspPlugInPluginsLv2OscillatorMonoscfSquaredCosine LspPlugInPluginsLv2OscillatorMonoscf = "Squared Cosine" // Squared Cosine (3) – Squared Cosine
	LspPlugInPluginsLv2OscillatorMonoscfRectangular LspPlugInPluginsLv2OscillatorMonoscf = "Rectangular" // Rectangular (4) – Rectangular
	LspPlugInPluginsLv2OscillatorMonoscfSawtooth LspPlugInPluginsLv2OscillatorMonoscf = "Sawtooth" // Sawtooth (5) – Sawtooth
	LspPlugInPluginsLv2OscillatorMonoscfTrapezoid LspPlugInPluginsLv2OscillatorMonoscf = "Trapezoid" // Trapezoid (6) – Trapezoid
	LspPlugInPluginsLv2OscillatorMonoscfPulsetrain LspPlugInPluginsLv2OscillatorMonoscf = "Pulsetrain" // Pulsetrain (7) – Pulsetrain
	LspPlugInPluginsLv2OscillatorMonoscfParabolic LspPlugInPluginsLv2OscillatorMonoscf = "Parabolic" // Parabolic (8) – Parabolic
	LspPlugInPluginsLv2OscillatorMonoscfBandLimitedRectangular LspPlugInPluginsLv2OscillatorMonoscf = "Band Limited Rectangular" // Band Limited Rectangular (9) – Band Limited Rectangular
	LspPlugInPluginsLv2OscillatorMonoscfBandLimitedSawtooth LspPlugInPluginsLv2OscillatorMonoscf = "Band Limited Sawtooth" // Band Limited Sawtooth (10) – Band Limited Sawtooth
	LspPlugInPluginsLv2OscillatorMonoscfBandLimitedTrapezoid LspPlugInPluginsLv2OscillatorMonoscf = "Band Limited Trapezoid" // Band Limited Trapezoid (11) – Band Limited Trapezoid
	LspPlugInPluginsLv2OscillatorMonoscfBandLimitedPulsetrain LspPlugInPluginsLv2OscillatorMonoscf = "Band Limited Pulsetrain" // Band Limited Pulsetrain (12) – Band Limited Pulsetrain
	LspPlugInPluginsLv2OscillatorMonoscfBandLimitedParabolic LspPlugInPluginsLv2OscillatorMonoscf = "Band Limited Parabolic" // Band Limited Parabolic (13) – Band Limited Parabolic
)

type LspPlugInPluginsLv2OscillatorMonoscm string

const (
	LspPlugInPluginsLv2OscillatorMonoscmAdd LspPlugInPluginsLv2OscillatorMonoscm = "Add" // Add (0) – Add
	LspPlugInPluginsLv2OscillatorMonoscmMultiply LspPlugInPluginsLv2OscillatorMonoscm = "Multiply" // Multiply (1) – Multiply
	LspPlugInPluginsLv2OscillatorMonoscmReplace LspPlugInPluginsLv2OscillatorMonoscm = "Replace" // Replace (2) – Replace
)

type LspPlugInPluginsLv2OscillatorMonoscom string

const (
	LspPlugInPluginsLv2OscillatorMonoscomNone LspPlugInPluginsLv2OscillatorMonoscom = "None" // None (0) – None
	LspPlugInPluginsLv2OscillatorMonoscomX2 LspPlugInPluginsLv2OscillatorMonoscom = "x2" // x2 (1) – x2
	LspPlugInPluginsLv2OscillatorMonoscomX3 LspPlugInPluginsLv2OscillatorMonoscom = "x3" // x3 (2) – x3
	LspPlugInPluginsLv2OscillatorMonoscomX4 LspPlugInPluginsLv2OscillatorMonoscom = "x4" // x4 (3) – x4
	LspPlugInPluginsLv2OscillatorMonoscomX6 LspPlugInPluginsLv2OscillatorMonoscom = "x6" // x6 (4) – x6
	LspPlugInPluginsLv2OscillatorMonoscomX8 LspPlugInPluginsLv2OscillatorMonoscom = "x8" // x8 (5) – x8
)

type LspPlugInPluginsLv2OscillatorMonoscr string

const (
	LspPlugInPluginsLv2OscillatorMonoscrWaveDc LspPlugInPluginsLv2OscillatorMonoscr = "Wave DC" // Wave DC (0) – Wave DC
	LspPlugInPluginsLv2OscillatorMonoscrZeroDc LspPlugInPluginsLv2OscillatorMonoscr = "Zero DC" // Zero DC (1) – Zero DC
)

