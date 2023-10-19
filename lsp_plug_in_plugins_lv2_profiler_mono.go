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

type LspPlugInPluginsLv2ProfilerMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ProfilerMono() (*LspPlugInPluginsLv2ProfilerMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-profiler-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ProfilerMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ProfilerMonoWithName(name string) (*LspPlugInPluginsLv2ProfilerMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-profiler-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ProfilerMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ProfilerMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cala (float32)
//
// Amplitude

func (e *LspPlugInPluginsLv2ProfilerMono) GetCala() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetCala(value float32) error {
	return e.Element.SetProperty("cala", value)
}

// calf (float32)
//
// Frequency

func (e *LspPlugInPluginsLv2ProfilerMono) GetCalf() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetCalf(value float32) error {
	return e.Element.SetProperty("calf", value)
}

// cals (bool)
//
// Calibration

func (e *LspPlugInPluginsLv2ProfilerMono) GetCals() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetCals(value bool) error {
	return e.Element.SetProperty("cals", value)
}

// fbck (bool)
//
// Feedback

func (e *LspPlugInPluginsLv2ProfilerMono) GetFbck() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetFbck(value bool) error {
	return e.Element.SetProperty("fbck", value)
}

// ili (float32)
//
// Integration Time

func (e *LspPlugInPluginsLv2ProfilerMono) GetIli() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ili")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv (float32)
//
// Input Level

func (e *LspPlugInPluginsLv2ProfilerMono) GetIlv() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilv")
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

func (e *LspPlugInPluginsLv2ProfilerMono) GetIrfc() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetIrfc(value bool) error {
	return e.Element.SetProperty("irfc", value)
}

// irfp (float32)
//
// File saving progress

func (e *LspPlugInPluginsLv2ProfilerMono) GetIrfp() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) GetIrfs() (int, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) GetLatt() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetLatt(value bool) error {
	return e.Element.SetProperty("latt", value)
}

// lint (bool)
//
// Trig a Linear measurement

func (e *LspPlugInPluginsLv2ProfilerMono) GetLint() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetLint(value bool) error {
	return e.Element.SetProperty("lint", value)
}

// ltda (float32)
//
// Absolute threshold

func (e *LspPlugInPluginsLv2ProfilerMono) GetLtda() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetLtda(value float32) error {
	return e.Element.SetProperty("ltda", value)
}

// ltdm (float32)
//
// Max expected latency

func (e *LspPlugInPluginsLv2ProfilerMono) GetLtdm() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetLtdm(value float32) error {
	return e.Element.SetProperty("ltdm", value)
}

// ltdp (float32)
//
// Peak threshold

func (e *LspPlugInPluginsLv2ProfilerMono) GetLtdp() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetLtdp(value float32) error {
	return e.Element.SetProperty("ltdp", value)
}

// ltena (bool)
//
// Enable Latency Detection

func (e *LspPlugInPluginsLv2ProfilerMono) GetLtena() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetLtena(value bool) error {
	return e.Element.SetProperty("ltena", value)
}

// lti (float32)
//
// Latency Value

func (e *LspPlugInPluginsLv2ProfilerMono) GetLti() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lti")
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

func (e *LspPlugInPluginsLv2ProfilerMono) GetOffc() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetOffc(value float32) error {
	return e.Element.SetProperty("offc", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ProfilerMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) GetPost() (bool, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetPost(value bool) error {
	return e.Element.SetProperty("post", value)
}

// rci (float32)
//
// Regression Line Correlation

func (e *LspPlugInPluginsLv2ProfilerMono) GetRci() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rci")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rta (bool)
//
// Reverberation Time Accuracy

func (e *LspPlugInPluginsLv2ProfilerMono) GetRta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// rti (float32)
//
// Reverberation Time

func (e *LspPlugInPluginsLv2ProfilerMono) GetRti() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rti")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// scra (GstLspPlugInPluginsLv2ProfilerMonoscra)
//
// RT Algorithm

func (e *LspPlugInPluginsLv2ProfilerMono) GetScra() (interface{}, error) {
	return e.Element.GetProperty("scra")
}

func (e *LspPlugInPluginsLv2ProfilerMono) SetScra(value interface{}) error {
	return e.Element.SetProperty("scra", value)
}

// scsv (GstLspPlugInPluginsLv2ProfilerMonoscsv)
//
// Save Mode

func (e *LspPlugInPluginsLv2ProfilerMono) GetScsv() (interface{}, error) {
	return e.Element.GetProperty("scsv")
}

func (e *LspPlugInPluginsLv2ProfilerMono) SetScsv(value interface{}) error {
	return e.Element.SetProperty("scsv", value)
}

// stld (GstLspPlugInPluginsLv2ProfilerMonostld)
//
// State LED

func (e *LspPlugInPluginsLv2ProfilerMono) GetStld() (interface{}, error) {
	return e.Element.GetProperty("stld")
}

// tind (float32)
//
// Actual Signal Duration

func (e *LspPlugInPluginsLv2ProfilerMono) GetTind() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) GetTsgl() (float32, error) {
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

func (e *LspPlugInPluginsLv2ProfilerMono) SetTsgl(value float32) error {
	return e.Element.SetProperty("tsgl", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ProfilerMonoscsv string

const (
	LspPlugInPluginsLv2ProfilerMonoscsvLtiAutoWav LspPlugInPluginsLv2ProfilerMonoscsv = "LTI Auto (*.wav)" // LTI Auto (*.wav) (0) – LTI Auto (*.wav)
	LspPlugInPluginsLv2ProfilerMonoscsvLtiRtWav LspPlugInPluginsLv2ProfilerMonoscsv = "LTI RT (*.wav)" // LTI RT (*.wav) (1) – LTI RT (*.wav)
	LspPlugInPluginsLv2ProfilerMonoscsvLtiCoarseWav LspPlugInPluginsLv2ProfilerMonoscsv = "LTI Coarse (*.wav)" // LTI Coarse (*.wav) (2) – LTI Coarse (*.wav)
	LspPlugInPluginsLv2ProfilerMonoscsvLtiAllWav LspPlugInPluginsLv2ProfilerMonoscsv = "LTI All (*.wav)" // LTI All (*.wav) (3) – LTI All (*.wav)
	LspPlugInPluginsLv2ProfilerMonoscsvAllInfoLspc LspPlugInPluginsLv2ProfilerMonoscsv = "All Info (*.lspc)" // All Info (*.lspc) (4) – All Info (*.lspc)
)

type LspPlugInPluginsLv2ProfilerMonostld string

const (
	LspPlugInPluginsLv2ProfilerMonostldIdle LspPlugInPluginsLv2ProfilerMonostld = "Idle" // Idle (0) – Idle
	LspPlugInPluginsLv2ProfilerMonostldCalibration LspPlugInPluginsLv2ProfilerMonostld = "Calibration" // Calibration (1) – Calibration
	LspPlugInPluginsLv2ProfilerMonostldLatencyDetection LspPlugInPluginsLv2ProfilerMonostld = "Latency Detection" // Latency Detection (2) – Latency Detection
	LspPlugInPluginsLv2ProfilerMonostldPreprocessing LspPlugInPluginsLv2ProfilerMonostld = "Preprocessing" // Preprocessing (3) – Preprocessing
	LspPlugInPluginsLv2ProfilerMonostldWaiting LspPlugInPluginsLv2ProfilerMonostld = "Waiting" // Waiting (4) – Waiting
	LspPlugInPluginsLv2ProfilerMonostldRecording LspPlugInPluginsLv2ProfilerMonostld = "Recording" // Recording (5) – Recording
	LspPlugInPluginsLv2ProfilerMonostldConvolving LspPlugInPluginsLv2ProfilerMonostld = "Convolving" // Convolving (6) – Convolving
	LspPlugInPluginsLv2ProfilerMonostldPostprocessing LspPlugInPluginsLv2ProfilerMonostld = "Postprocessing" // Postprocessing (7) – Postprocessing
	LspPlugInPluginsLv2ProfilerMonostldSaving LspPlugInPluginsLv2ProfilerMonostld = "Saving" // Saving (8) – Saving
)

type LspPlugInPluginsLv2ProfilerMonoscra string

const (
	LspPlugInPluginsLv2ProfilerMonoscraEdt0 LspPlugInPluginsLv2ProfilerMonoscra = "EDT0" // EDT0 (0) – EDT0
	LspPlugInPluginsLv2ProfilerMonoscraEdt1 LspPlugInPluginsLv2ProfilerMonoscra = "EDT1" // EDT1 (1) – EDT1
	LspPlugInPluginsLv2ProfilerMonoscraRt10 LspPlugInPluginsLv2ProfilerMonoscra = "RT10" // RT10 (2) – RT10
	LspPlugInPluginsLv2ProfilerMonoscraRt20 LspPlugInPluginsLv2ProfilerMonoscra = "RT20" // RT20 (3) – RT20
	LspPlugInPluginsLv2ProfilerMonoscraRt30 LspPlugInPluginsLv2ProfilerMonoscra = "RT30" // RT30 (4) – RT30
)

