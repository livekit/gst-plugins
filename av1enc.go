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
)

type Av1Enc struct {
	Element *gst.Element
}

func NewAv1Enc() (*Av1Enc, error) {
	e, err := gst.NewElement("av1enc")
	if err != nil {
		return nil, err
	}
	return &Av1Enc{Element: e}, nil
}

func NewAv1EncWithName(name string) (*Av1Enc, error) {
	e, err := gst.NewElementWithName("av1enc", name)
	if err != nil {
		return nil, err
	}
	return &Av1Enc{Element: e}, nil
}

// ----- Properties -----

// buf-initial-sz (uint)
//
// Decoder buffer initial size, expressed in units of time (milliseconds)

func (e *Av1Enc) GetBufInitialSz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buf-initial-sz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetBufInitialSz(value uint) error {
	return e.Element.SetProperty("buf-initial-sz", value)
}

// buf-optimal-sz (uint)
//
// Decoder buffer optimal size, expressed in units of time (milliseconds)

func (e *Av1Enc) GetBufOptimalSz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buf-optimal-sz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetBufOptimalSz(value uint) error {
	return e.Element.SetProperty("buf-optimal-sz", value)
}

// buf-sz (uint)
//
// Decoder buffer size, expressed in units of time (milliseconds)

func (e *Av1Enc) GetBufSz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buf-sz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetBufSz(value uint) error {
	return e.Element.SetProperty("buf-sz", value)
}

// cpu-used (int)
//
// CPU Used. A Value greater than 0 will increase encoder speed at the expense of quality.

func (e *Av1Enc) GetCpuUsed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cpu-used")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Av1Enc) SetCpuUsed(value int) error {
	return e.Element.SetProperty("cpu-used", value)
}

// drop-frame (uint)
//
// Temporal resampling configuration, drop frames as a strategy to meet its target data rate Set to zero (0) to disable this feature.

func (e *Av1Enc) GetDropFrame() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("drop-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetDropFrame(value uint) error {
	return e.Element.SetProperty("drop-frame", value)
}

// enc-pass (GstAv1EncEncPass)

func (e *Av1Enc) GetEncPass() (GstAv1EncEncPass, error) {
	var value GstAv1EncEncPass
	var ok bool
	v, err := e.Element.GetProperty("enc-pass")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAv1EncEncPass)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAv1EncEncPass")
	}
	return value, nil
}

func (e *Av1Enc) SetEncPass(value GstAv1EncEncPass) error {
	e.Element.SetArg("enc-pass", string(value))
	return nil
}

// end-usage (GstAv1EncEndUsageMode)
//
// Rate control algorithm to use, indicates the end usage of this stream

func (e *Av1Enc) GetEndUsage() (GstAv1EncEndUsageMode, error) {
	var value GstAv1EncEndUsageMode
	var ok bool
	v, err := e.Element.GetProperty("end-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAv1EncEndUsageMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAv1EncEndUsageMode")
	}
	return value, nil
}

func (e *Av1Enc) SetEndUsage(value GstAv1EncEndUsageMode) error {
	e.Element.SetArg("end-usage", string(value))
	return nil
}

// keyframe-max-dist (int)

func (e *Av1Enc) GetKeyframeMaxDist() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-max-dist")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Av1Enc) SetKeyframeMaxDist(value int) error {
	return e.Element.SetProperty("keyframe-max-dist", value)
}

// keyframe-mode (GstAv1EncKfmode)

func (e *Av1Enc) GetKeyframeMode() (GstAv1EncKfmode, error) {
	var value GstAv1EncKfmode
	var ok bool
	v, err := e.Element.GetProperty("keyframe-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAv1EncKfmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAv1EncKfmode")
	}
	return value, nil
}

func (e *Av1Enc) SetKeyframeMode(value GstAv1EncKfmode) error {
	e.Element.SetArg("keyframe-mode", string(value))
	return nil
}

// lag-in-frames (uint)

func (e *Av1Enc) GetLagInFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lag-in-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetLagInFrames(value uint) error {
	return e.Element.SetProperty("lag-in-frames", value)
}

// max-quantizer (uint)
//
// Maximum (worst quality) quantizer

func (e *Av1Enc) GetMaxQuantizer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetMaxQuantizer(value uint) error {
	return e.Element.SetProperty("max-quantizer", value)
}

// min-quantizer (uint)
//
// Minimum (best quality) quantizer

func (e *Av1Enc) GetMinQuantizer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetMinQuantizer(value uint) error {
	return e.Element.SetProperty("min-quantizer", value)
}

// overshoot-pct (uint)
//
// Rate control adaptation overshoot control

func (e *Av1Enc) GetOvershootPct() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("overshoot-pct")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetOvershootPct(value uint) error {
	return e.Element.SetProperty("overshoot-pct", value)
}

// resize-denominator (uint)
//
// Frame resize denominator, assuming 8 as the numerator

func (e *Av1Enc) GetResizeDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("resize-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetResizeDenominator(value uint) error {
	return e.Element.SetProperty("resize-denominator", value)
}

// resize-kf-denominator (uint)
//
// Frame resize keyframe denominator, assuming 8 as the numerator

func (e *Av1Enc) GetResizeKfDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("resize-kf-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetResizeKfDenominator(value uint) error {
	return e.Element.SetProperty("resize-kf-denominator", value)
}

// resize-mode (GstAv1EncResizeMode)
//
// Frame resize mode

func (e *Av1Enc) GetResizeMode() (GstAv1EncResizeMode, error) {
	var value GstAv1EncResizeMode
	var ok bool
	v, err := e.Element.GetProperty("resize-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAv1EncResizeMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAv1EncResizeMode")
	}
	return value, nil
}

func (e *Av1Enc) SetResizeMode(value GstAv1EncResizeMode) error {
	e.Element.SetArg("resize-mode", string(value))
	return nil
}

// row-mt (bool)
//
// Enable row based multi-threading

func (e *Av1Enc) GetRowMt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("row-mt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Av1Enc) SetRowMt(value bool) error {
	return e.Element.SetProperty("row-mt", value)
}

// superres-denominator (uint)
//
// Frame super-resolution denominator, used only by SUPERRES_FIXED mode

func (e *Av1Enc) GetSuperresDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetSuperresDenominator(value uint) error {
	return e.Element.SetProperty("superres-denominator", value)
}

// superres-kf-denominator (uint)
//
// Keyframe super-resolution denominator

func (e *Av1Enc) GetSuperresKfDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-kf-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetSuperresKfDenominator(value uint) error {
	return e.Element.SetProperty("superres-kf-denominator", value)
}

// superres-kf-qthresh (uint)
//
// Keyframe super-resolution qindex threshold, used only by SUPERRES_QTHRESH mode

func (e *Av1Enc) GetSuperresKfQthresh() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-kf-qthresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetSuperresKfQthresh(value uint) error {
	return e.Element.SetProperty("superres-kf-qthresh", value)
}

// superres-mode (GstAv1EncSuperresMode)
//
// It integrates upscaling after the encode/decode process

func (e *Av1Enc) GetSuperresMode() (GstAv1EncSuperresMode, error) {
	var value GstAv1EncSuperresMode
	var ok bool
	v, err := e.Element.GetProperty("superres-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAv1EncSuperresMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAv1EncSuperresMode")
	}
	return value, nil
}

func (e *Av1Enc) SetSuperresMode(value GstAv1EncSuperresMode) error {
	e.Element.SetArg("superres-mode", string(value))
	return nil
}

// superres-qthresh (uint)
//
// Frame super-resolution qindex threshold, used only by SUPERRES_QTHRESH mode

func (e *Av1Enc) GetSuperresQthresh() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-qthresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetSuperresQthresh(value uint) error {
	return e.Element.SetProperty("superres-qthresh", value)
}

// target-bitrate (uint)
//
// Target bitrate, in kilobits per second

func (e *Av1Enc) GetTargetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetTargetBitrate(value uint) error {
	return e.Element.SetProperty("target-bitrate", value)
}

// threads (uint)
//
// Max number of threads to use encoding, set to 0 determine the approximate number of threads that the system schedule

func (e *Av1Enc) GetThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetThreads(value uint) error {
	return e.Element.SetProperty("threads", value)
}

// tile-columns (uint)
//
// Partition into separate vertical tile columns from image frame which can enable parallel encoding

func (e *Av1Enc) GetTileColumns() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tile-columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetTileColumns(value uint) error {
	return e.Element.SetProperty("tile-columns", value)
}

// tile-rows (uint)
//
// Partition into separate horizontal tile rows from image frame which can enable parallel encoding

func (e *Av1Enc) GetTileRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetTileRows(value uint) error {
	return e.Element.SetProperty("tile-rows", value)
}

// undershoot-pct (uint)
//
// Rate control adaptation undershoot control

func (e *Av1Enc) GetUndershootPct() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("undershoot-pct")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Av1Enc) SetUndershootPct(value uint) error {
	return e.Element.SetProperty("undershoot-pct", value)
}

// usage-profile (GstAv1EncUsageProfile)

func (e *Av1Enc) GetUsageProfile() (GstAv1EncUsageProfile, error) {
	var value GstAv1EncUsageProfile
	var ok bool
	v, err := e.Element.GetProperty("usage-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAv1EncUsageProfile)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAv1EncUsageProfile")
	}
	return value, nil
}

func (e *Av1Enc) SetUsageProfile(value GstAv1EncUsageProfile) error {
	e.Element.SetArg("usage-profile", string(value))
	return nil
}

// ----- Constants -----

type GstAv1EncSuperresMode string

const (
	GstAv1EncSuperresModeNone GstAv1EncSuperresMode = "none" // none (0) – No frame superres allowed
	GstAv1EncSuperresModeFixed GstAv1EncSuperresMode = "fixed" // fixed (1) – All frames are coded at the specified scale and super-resolved
	GstAv1EncSuperresModeRandom GstAv1EncSuperresMode = "random" // random (2) – All frames are coded at a random scale and super-resolved
	GstAv1EncSuperresModeQthresh GstAv1EncSuperresMode = "qthresh" // qthresh (3) – Superres scale for a frame is determined based on q_index
)

type GstAv1EncUsageProfile string

const (
	GstAv1EncUsageProfileGoodQuality GstAv1EncUsageProfile = "good-quality" // good-quality (0) – Good Quality profile
	GstAv1EncUsageProfileRealtime GstAv1EncUsageProfile = "realtime" // realtime (1) – Realtime profile
	GstAv1EncUsageProfileAllIntra GstAv1EncUsageProfile = "all-intra" // all-intra (2) – All Intra profile
)

type GstAv1EncEncPass string

const (
	GstAv1EncEncPassOnePass GstAv1EncEncPass = "one-pass" // one-pass (0) – Single pass mode
	GstAv1EncEncPassFirstPass GstAv1EncEncPass = "first-pass" // first-pass (1) – First pass of multi-pass mode
	GstAv1EncEncPassSecondPass GstAv1EncEncPass = "second-pass" // second-pass (2) – Second pass of multi-pass mode
	GstAv1EncEncPassThirdPass GstAv1EncEncPass = "third-pass" // third-pass (3) – Third pass of multi-pass mode
)

type GstAv1EncEndUsageMode string

const (
	GstAv1EncEndUsageModeVbr GstAv1EncEndUsageMode = "vbr" // vbr (0) – Variable Bit Rate Mode
	GstAv1EncEndUsageModeCbr GstAv1EncEndUsageMode = "cbr" // cbr (1) – Constant Bit Rate Mode
	GstAv1EncEndUsageModeCq GstAv1EncEndUsageMode = "cq" // cq (2) – Constrained Quality Mode
	GstAv1EncEndUsageModeQ GstAv1EncEndUsageMode = "q" // q (3) – Constant Quality Mode
)

type GstAv1EncKfmode string

const (
	GstAv1EncKfmodeAuto GstAv1EncKfmode = "auto" // auto (1) – Encoder determines optimal keyframe placement automatically
	GstAv1EncKfmodeDisabled GstAv1EncKfmode = "disabled" // disabled (0) – Encoder does not place keyframes
)

type GstAv1EncResizeMode string

const (
	GstAv1EncResizeModeNone GstAv1EncResizeMode = "none" // none (0) – No frame resizing allowed
	GstAv1EncResizeModeFixed GstAv1EncResizeMode = "fixed" // fixed (1) – All frames are coded at the specified scale
	GstAv1EncResizeModeRandom GstAv1EncResizeMode = "random" // random (2) – All frames are coded at a random scale
)

