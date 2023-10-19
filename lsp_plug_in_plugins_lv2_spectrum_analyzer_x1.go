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

type LspPlugInPluginsLv2SpectrumAnalyzerX1 struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2SpectrumAnalyzerX1() (*LspPlugInPluginsLv2SpectrumAnalyzerX1, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-spectrum-analyzer-x1")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SpectrumAnalyzerX1{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2SpectrumAnalyzerX1WithName(name string) (*LspPlugInPluginsLv2SpectrumAnalyzerX1, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-spectrum-analyzer-x1", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SpectrumAnalyzerX1{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// chn (GstLspPlugInPluginsLv2SpectrumAnalyzerX1Chn)
//
// Channel

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetChn() (interface{}, error) {
	return e.Element.GetProperty("chn")
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetChn(value interface{}) error {
	return e.Element.SetProperty("chn", value)
}

// env (GstLspPlugInPluginsLv2SpectrumAnalyzerX1Env)
//
// FFT Envelope

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetEnv() (interface{}, error) {
	return e.Element.GetProperty("env")
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetEnv(value interface{}) error {
	return e.Element.SetProperty("env", value)
}

// freeze (bool)
//
// Analyzer freeze

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetFreeze() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("freeze")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetFreeze(value bool) error {
	return e.Element.SetProperty("freeze", value)
}

// freq (float32)
//
// Frequency

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetFreq() (float32, error) {
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

// frz-0 (bool)
//
// Freeze 0

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetFrz0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("frz-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetFrz0(value bool) error {
	return e.Element.SetProperty("frz-0", value)
}

// hue-0 (float32)
//
// Hue 0

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetHue0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// lvl (float32)
//
// Level

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetLvl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lvl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// mode (GstLspPlugInPluginsLv2SpectrumAnalyzerX1Mode)
//
// Analyzer mode

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// on-0 (bool)
//
// Analyse 0

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetOn0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("on-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetOn0(value bool) error {
	return e.Element.SetProperty("on-0", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetOutLatency() (int, error) {
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

// pamp (float32)
//
// Preamp gain

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetPamp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetPamp(value float32) error {
	return e.Element.SetProperty("pamp", value)
}

// react (float32)
//
// Reactivity

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// sel (float32)
//
// Selector

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetSel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetSel(value float32) error {
	return e.Element.SetProperty("sel", value)
}

// sh-0 (float32)
//
// Shift gain 0

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetSh0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sh-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetSh0(value float32) error {
	return e.Element.SetProperty("sh-0", value)
}

// solo-0 (bool)
//
// Solo 0

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetSolo0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("solo-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetSolo0(value bool) error {
	return e.Element.SetProperty("solo-0", value)
}

// splog (bool)
//
// Spectralizer logarithmic scale

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetSplog() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("splog")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetSplog(value bool) error {
	return e.Element.SetProperty("splog", value)
}

// spm (GstLspPlugInPluginsLv2SpectrumAnalyzerX1Spm)
//
// Spectralizer mode

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetSpm() (interface{}, error) {
	return e.Element.GetProperty("spm")
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetSpm(value interface{}) error {
	return e.Element.SetProperty("spm", value)
}

// tol (GstLspPlugInPluginsLv2SpectrumAnalyzerX1Tol)
//
// FFT Tolerance

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetTol() (interface{}, error) {
	return e.Element.GetProperty("tol")
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetTol(value interface{}) error {
	return e.Element.SetProperty("tol", value)
}

// wnd (GstLspPlugInPluginsLv2SpectrumAnalyzerX1Wnd)
//
// FFT Window

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetWnd() (interface{}, error) {
	return e.Element.GetProperty("wnd")
}

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetWnd(value interface{}) error {
	return e.Element.SetProperty("wnd", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2SpectrumAnalyzerX1) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2SpectrumAnalyzerX1Spm string

const (
	LspPlugInPluginsLv2SpectrumAnalyzerX1SpmRainbow LspPlugInPluginsLv2SpectrumAnalyzerX1Spm = "Rainbow" // Rainbow (0) – Rainbow
	LspPlugInPluginsLv2SpectrumAnalyzerX1SpmFog LspPlugInPluginsLv2SpectrumAnalyzerX1Spm = "Fog" // Fog (1) – Fog
	LspPlugInPluginsLv2SpectrumAnalyzerX1SpmColor LspPlugInPluginsLv2SpectrumAnalyzerX1Spm = "Color" // Color (2) – Color
	LspPlugInPluginsLv2SpectrumAnalyzerX1SpmLightning LspPlugInPluginsLv2SpectrumAnalyzerX1Spm = "Lightning" // Lightning (3) – Lightning
	LspPlugInPluginsLv2SpectrumAnalyzerX1SpmLightness LspPlugInPluginsLv2SpectrumAnalyzerX1Spm = "Lightness" // Lightness (4) – Lightness
)

type LspPlugInPluginsLv2SpectrumAnalyzerX1Tol string

const (
	LspPlugInPluginsLv2SpectrumAnalyzerX1Tol1024 LspPlugInPluginsLv2SpectrumAnalyzerX1Tol = "1024" // 1024 (0) – 1024
	LspPlugInPluginsLv2SpectrumAnalyzerX1Tol2048 LspPlugInPluginsLv2SpectrumAnalyzerX1Tol = "2048" // 2048 (1) – 2048
	LspPlugInPluginsLv2SpectrumAnalyzerX1Tol4096 LspPlugInPluginsLv2SpectrumAnalyzerX1Tol = "4096" // 4096 (2) – 4096
	LspPlugInPluginsLv2SpectrumAnalyzerX1Tol8192 LspPlugInPluginsLv2SpectrumAnalyzerX1Tol = "8192" // 8192 (3) – 8192
	LspPlugInPluginsLv2SpectrumAnalyzerX1Tol16384 LspPlugInPluginsLv2SpectrumAnalyzerX1Tol = "16384" // 16384 (4) – 16384
)

type LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd string

const (
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndHann LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Hann" // Hann (0) – Hann
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndHamming LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Hamming" // Hamming (1) – Hamming
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndBlackman LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Blackman" // Blackman (2) – Blackman
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndLanczos LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Lanczos" // Lanczos (3) – Lanczos
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndGaussian LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Gaussian" // Gaussian (4) – Gaussian
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndPoisson LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Poisson" // Poisson (5) – Poisson
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndParzen LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Parzen" // Parzen (6) – Parzen
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndTukey LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Tukey" // Tukey (7) – Tukey
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndWelch LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Welch" // Welch (8) – Welch
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndNuttall LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Nuttall" // Nuttall (9) – Nuttall
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndBlackmanNuttall LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Blackman-Nuttall" // Blackman-Nuttall (10) – Blackman-Nuttall
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndBlackmanHarris LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Blackman-Harris" // Blackman-Harris (11) – Blackman-Harris
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndHannPoisson LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Hann-Poisson" // Hann-Poisson (12) – Hann-Poisson
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndBartlettHann LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Bartlett-Hann" // Bartlett-Hann (13) – Bartlett-Hann
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndBartlettFejer LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Bartlett-Fejer" // Bartlett-Fejer (14) – Bartlett-Fejer
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndTriangular LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Triangular" // Triangular (15) – Triangular
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndRectangular LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Rectangular" // Rectangular (16) – Rectangular
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndFlatTop LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Flat top" // Flat top (17) – Flat top
	LspPlugInPluginsLv2SpectrumAnalyzerX1WndCosine LspPlugInPluginsLv2SpectrumAnalyzerX1Wnd = "Cosine" // Cosine (18) – Cosine
)

type LspPlugInPluginsLv2SpectrumAnalyzerX1Chn string

const (
	LspPlugInPluginsLv2SpectrumAnalyzerX1Chn0 LspPlugInPluginsLv2SpectrumAnalyzerX1Chn = "0" // 0 (0) – 0
)

type LspPlugInPluginsLv2SpectrumAnalyzerX1Env string

const (
	LspPlugInPluginsLv2SpectrumAnalyzerX1EnvPurpleNoise LspPlugInPluginsLv2SpectrumAnalyzerX1Env = "Purple noise" // Purple noise (0) – Purple noise
	LspPlugInPluginsLv2SpectrumAnalyzerX1EnvBlueNoise LspPlugInPluginsLv2SpectrumAnalyzerX1Env = "Blue noise" // Blue noise (1) – Blue noise
	LspPlugInPluginsLv2SpectrumAnalyzerX1EnvWhiteNoise LspPlugInPluginsLv2SpectrumAnalyzerX1Env = "White noise" // White noise (2) – White noise
	LspPlugInPluginsLv2SpectrumAnalyzerX1EnvPinkNoise LspPlugInPluginsLv2SpectrumAnalyzerX1Env = "Pink noise" // Pink noise (3) – Pink noise
	LspPlugInPluginsLv2SpectrumAnalyzerX1EnvBrownNoise LspPlugInPluginsLv2SpectrumAnalyzerX1Env = "Brown noise" // Brown noise (4) – Brown noise
)

type LspPlugInPluginsLv2SpectrumAnalyzerX1Mode string

const (
	LspPlugInPluginsLv2SpectrumAnalyzerX1ModeAnalyzer LspPlugInPluginsLv2SpectrumAnalyzerX1Mode = "Analyzer" // Analyzer (0) – Analyzer
	LspPlugInPluginsLv2SpectrumAnalyzerX1ModeMastering LspPlugInPluginsLv2SpectrumAnalyzerX1Mode = "Mastering" // Mastering (1) – Mastering
	LspPlugInPluginsLv2SpectrumAnalyzerX1ModeSpectralizer LspPlugInPluginsLv2SpectrumAnalyzerX1Mode = "Spectralizer" // Spectralizer (2) – Spectralizer
)

