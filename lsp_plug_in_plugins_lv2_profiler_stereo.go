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

type LspPlugInPluginsLv2ProfilerStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ProfilerStereo() (*LspPlugInPluginsLv2ProfilerStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-profiler-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ProfilerStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ProfilerStereoWithName(name string) (*LspPlugInPluginsLv2ProfilerStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-profiler-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ProfilerStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ProfilerStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cala (float32)
//
// Amplitude

func (e *LspPlugInPluginsLv2ProfilerStereo) GetCala() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cala")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetCala(value float32) error {
	return e.Element.SetProperty("cala", value)
}

// calf (float32)
//
// Frequency

func (e *LspPlugInPluginsLv2ProfilerStereo) GetCalf() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("calf")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetCalf(value float32) error {
	return e.Element.SetProperty("calf", value)
}

// cals (bool)
//
// Calibration

func (e *LspPlugInPluginsLv2ProfilerStereo) GetCals() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cals")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetCals(value bool) error {
	return e.Element.SetProperty("cals", value)
}

// fbck (bool)
//
// Feedback

func (e *LspPlugInPluginsLv2ProfilerStereo) GetFbck() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fbck")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetFbck(value bool) error {
	return e.Element.SetProperty("fbck", value)
}

// ili-l (float32)
//
// Integration Time Left

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIliL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ili-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ili-r (float32)
//
// Integration Time Right

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIliR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ili-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv-l (float32)
//
// Input Level Left

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIlvL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv-r (float32)
//
// Input Level Right

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIlvR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// irfc (bool)
//
// Save file command

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIrfc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("irfc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetIrfc(value bool) error {
	return e.Element.SetProperty("irfc", value)
}

// irfp (float32)
//
// File saving progress

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIrfp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("irfp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// irfs (int)
//
// File saving status

func (e *LspPlugInPluginsLv2ProfilerStereo) GetIrfs() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("irfs")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// latt (bool)
//
// Trig a Latency measurement

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLatt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("latt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetLatt(value bool) error {
	return e.Element.SetProperty("latt", value)
}

// lint (bool)
//
// Trig a Linear measurement

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLint() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lint")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetLint(value bool) error {
	return e.Element.SetProperty("lint", value)
}

// ltda (float32)
//
// Absolute threshold

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLtda() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ltda")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetLtda(value float32) error {
	return e.Element.SetProperty("ltda", value)
}

// ltdm (float32)
//
// Max expected latency

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLtdm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ltdm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetLtdm(value float32) error {
	return e.Element.SetProperty("ltdm", value)
}

// ltdp (float32)
//
// Peak threshold

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLtdp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ltdp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetLtdp(value float32) error {
	return e.Element.SetProperty("ltdp", value)
}

// ltena (bool)
//
// Enable Latency Detection

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLtena() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ltena")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetLtena(value bool) error {
	return e.Element.SetProperty("ltena", value)
}

// lti-l (float32)
//
// Latency Value Left

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLtiL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lti-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// lti-r (float32)
//
// Latency Value Right

func (e *LspPlugInPluginsLv2ProfilerStereo) GetLtiR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lti-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// offc (float32)
//
// IR Time Offset

func (e *LspPlugInPluginsLv2ProfilerStereo) GetOffc() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("offc")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetOffc(value float32) error {
	return e.Element.SetProperty("offc", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ProfilerStereo) GetOutLatency() (int, error) {
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

// post (bool)
//
// Trig Post Processing

func (e *LspPlugInPluginsLv2ProfilerStereo) GetPost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetPost(value bool) error {
	return e.Element.SetProperty("post", value)
}

// rci-l (float32)
//
// Regression Line Correlation Left

func (e *LspPlugInPluginsLv2ProfilerStereo) GetRciL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rci-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rci-r (float32)
//
// Regression Line Correlation Right

func (e *LspPlugInPluginsLv2ProfilerStereo) GetRciR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rci-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rta-l (bool)
//
// Reverberation Time Accuracy Left

func (e *LspPlugInPluginsLv2ProfilerStereo) GetRtaL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rta-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// rta-r (bool)
//
// Reverberation Time Accuracy Right

func (e *LspPlugInPluginsLv2ProfilerStereo) GetRtaR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rta-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// rti-l (float32)
//
// Reverberation Time Left

func (e *LspPlugInPluginsLv2ProfilerStereo) GetRtiL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rti-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rti-r (float32)
//
// Reverberation Time Right

func (e *LspPlugInPluginsLv2ProfilerStereo) GetRtiR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rti-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// scra (GstLspPlugInPluginsLv2ProfilerStereoscra)
//
// RT Algorithm

func (e *LspPlugInPluginsLv2ProfilerStereo) GetScra() (interface{}, error) {
	return e.Element.GetProperty("scra")
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetScra(value interface{}) error {
	return e.Element.SetProperty("scra", value)
}

// scsv (GstLspPlugInPluginsLv2ProfilerStereoscsv)
//
// Save Mode

func (e *LspPlugInPluginsLv2ProfilerStereo) GetScsv() (interface{}, error) {
	return e.Element.GetProperty("scsv")
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetScsv(value interface{}) error {
	return e.Element.SetProperty("scsv", value)
}

// stld (GstLspPlugInPluginsLv2ProfilerStereostld)
//
// State LED

func (e *LspPlugInPluginsLv2ProfilerStereo) GetStld() (interface{}, error) {
	return e.Element.GetProperty("stld")
}

// tind (float32)
//
// Actual Signal Duration

func (e *LspPlugInPluginsLv2ProfilerStereo) GetTind() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tind")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// tsgl (float32)
//
// Duration

func (e *LspPlugInPluginsLv2ProfilerStereo) GetTsgl() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tsgl")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ProfilerStereo) SetTsgl(value float32) error {
	return e.Element.SetProperty("tsgl", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ProfilerStereoscra string

const (
	LspPlugInPluginsLv2ProfilerStereoscraEdt0 LspPlugInPluginsLv2ProfilerStereoscra = "EDT0" // EDT0 (0) – EDT0
	LspPlugInPluginsLv2ProfilerStereoscraEdt1 LspPlugInPluginsLv2ProfilerStereoscra = "EDT1" // EDT1 (1) – EDT1
	LspPlugInPluginsLv2ProfilerStereoscraRt10 LspPlugInPluginsLv2ProfilerStereoscra = "RT10" // RT10 (2) – RT10
	LspPlugInPluginsLv2ProfilerStereoscraRt20 LspPlugInPluginsLv2ProfilerStereoscra = "RT20" // RT20 (3) – RT20
	LspPlugInPluginsLv2ProfilerStereoscraRt30 LspPlugInPluginsLv2ProfilerStereoscra = "RT30" // RT30 (4) – RT30
)

type LspPlugInPluginsLv2ProfilerStereoscsv string

const (
	LspPlugInPluginsLv2ProfilerStereoscsvLtiAutoWav LspPlugInPluginsLv2ProfilerStereoscsv = "LTI Auto (*.wav)" // LTI Auto (*.wav) (0) – LTI Auto (*.wav)
	LspPlugInPluginsLv2ProfilerStereoscsvLtiRtWav LspPlugInPluginsLv2ProfilerStereoscsv = "LTI RT (*.wav)" // LTI RT (*.wav) (1) – LTI RT (*.wav)
	LspPlugInPluginsLv2ProfilerStereoscsvLtiCoarseWav LspPlugInPluginsLv2ProfilerStereoscsv = "LTI Coarse (*.wav)" // LTI Coarse (*.wav) (2) – LTI Coarse (*.wav)
	LspPlugInPluginsLv2ProfilerStereoscsvLtiAllWav LspPlugInPluginsLv2ProfilerStereoscsv = "LTI All (*.wav)" // LTI All (*.wav) (3) – LTI All (*.wav)
	LspPlugInPluginsLv2ProfilerStereoscsvAllInfoLspc LspPlugInPluginsLv2ProfilerStereoscsv = "All Info (*.lspc)" // All Info (*.lspc) (4) – All Info (*.lspc)
)

type LspPlugInPluginsLv2ProfilerStereostld string

const (
	LspPlugInPluginsLv2ProfilerStereostldIdle LspPlugInPluginsLv2ProfilerStereostld = "Idle" // Idle (0) – Idle
	LspPlugInPluginsLv2ProfilerStereostldCalibration LspPlugInPluginsLv2ProfilerStereostld = "Calibration" // Calibration (1) – Calibration
	LspPlugInPluginsLv2ProfilerStereostldLatencyDetection LspPlugInPluginsLv2ProfilerStereostld = "Latency Detection" // Latency Detection (2) – Latency Detection
	LspPlugInPluginsLv2ProfilerStereostldPreprocessing LspPlugInPluginsLv2ProfilerStereostld = "Preprocessing" // Preprocessing (3) – Preprocessing
	LspPlugInPluginsLv2ProfilerStereostldWaiting LspPlugInPluginsLv2ProfilerStereostld = "Waiting" // Waiting (4) – Waiting
	LspPlugInPluginsLv2ProfilerStereostldRecording LspPlugInPluginsLv2ProfilerStereostld = "Recording" // Recording (5) – Recording
	LspPlugInPluginsLv2ProfilerStereostldConvolving LspPlugInPluginsLv2ProfilerStereostld = "Convolving" // Convolving (6) – Convolving
	LspPlugInPluginsLv2ProfilerStereostldPostprocessing LspPlugInPluginsLv2ProfilerStereostld = "Postprocessing" // Postprocessing (7) – Postprocessing
	LspPlugInPluginsLv2ProfilerStereostldSaving LspPlugInPluginsLv2ProfilerStereostld = "Saving" // Saving (8) – Saving
)

