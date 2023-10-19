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

type Qsvh264Enc struct {
	Element *gst.Element
}

func NewQsvh264Enc() (*Qsvh264Enc, error) {
	e, err := gst.NewElement("qsvh264enc")
	if err != nil {
		return nil, err
	}
	return &Qsvh264Enc{Element: e}, nil
}

func NewQsvh264EncWithName(name string) (*Qsvh264Enc, error) {
	e, err := gst.NewElementWithName("qsvh264enc", name)
	if err != nil {
		return nil, err
	}
	return &Qsvh264Enc{Element: e}, nil
}

// ----- Properties -----

// avbr-accuracy (uint)
//
// AVBR Accuracy in the unit of tenth of percent

func (e *Qsvh264Enc) GetAvbrAccuracy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("avbr-accuracy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetAvbrAccuracy(value uint) error {
	return e.Element.SetProperty("avbr-accuracy", value)
}

// avbr-convergence (uint)
//
// AVBR Convergence in the unit of 100 frames

func (e *Qsvh264Enc) GetAvbrConvergence() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("avbr-convergence")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetAvbrConvergence(value uint) error {
	return e.Element.SetProperty("avbr-convergence", value)
}

// b-frames (uint)
//
// Number of B frames between I and P frames

func (e *Qsvh264Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

// bitrate (uint)
//
// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp", "icq", and "la_icq")

func (e *Qsvh264Enc) GetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// cabac (GstQsvCodingOption)
//
// Enables CABAC entropy coding

func (e *Qsvh264Enc) GetCabac() (GstQsvCodingOption, error) {
	var value GstQsvCodingOption
	var ok bool
	v, err := e.Element.GetProperty("cabac")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvCodingOption)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvCodingOption")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetCabac(value GstQsvCodingOption) error {
	e.Element.SetArg("cabac", string(value))
	return nil
}

// cc-insert (GstQsvH264EncSeiInsertMode)
//
// Closed Caption Insert mode. Only CEA-708 RAW format is supported for now

func (e *Qsvh264Enc) GetCcInsert() (GstQsvH264EncSeiInsertMode, error) {
	var value GstQsvH264EncSeiInsertMode
	var ok bool
	v, err := e.Element.GetProperty("cc-insert")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264EncSeiInsertMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264EncSeiInsertMode")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetCcInsert(value GstQsvH264EncSeiInsertMode) error {
	e.Element.SetArg("cc-insert", string(value))
	return nil
}

// disable-hrd-conformance (bool)
//
// Allow NAL HRD non-conformant stream

func (e *Qsvh264Enc) GetDisableHrdConformance() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-hrd-conformance")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetDisableHrdConformance(value bool) error {
	return e.Element.SetProperty("disable-hrd-conformance", value)
}

// gop-size (uint)
//
// Number of pictures within a GOP (0: unspecified)

func (e *Qsvh264Enc) GetGopSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

// icq-quality (uint)
//
// Intelligent Constant Quality for "icq" rate-control (0: default)

func (e *Qsvh264Enc) GetIcqQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("icq-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetIcqQuality(value uint) error {
	return e.Element.SetProperty("icq-quality", value)
}

// idr-interval (uint)
//
// IDR-frame interval in terms of I-frames. 0: every I-frame is an IDR frame, N: "N" I-frames are inserted between IDR-frames

func (e *Qsvh264Enc) GetIdrInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("idr-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetIdrInterval(value uint) error {
	return e.Element.SetProperty("idr-interval", value)
}

// max-bitrate (uint)
//
// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp", "icq", and "la_icq")

func (e *Qsvh264Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-frame-size (uint)

func (e *Qsvh264Enc) GetMaxFrameSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxFrameSize(value uint) error {
	return e.Element.SetProperty("max-frame-size", value)
}

// max-frame-size-i (uint)

func (e *Qsvh264Enc) GetMaxFrameSizeI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-frame-size-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxFrameSizeI(value uint) error {
	return e.Element.SetProperty("max-frame-size-i", value)
}

// max-frame-size-p (uint)

func (e *Qsvh264Enc) GetMaxFrameSizeP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-frame-size-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxFrameSizeP(value uint) error {
	return e.Element.SetProperty("max-frame-size-p", value)
}

// max-qp-b (uint)
//
// Maximum allowed QP value for B-frame types (0: default)

func (e *Qsvh264Enc) GetMaxQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

// max-qp-i (uint)
//
// Maximum allowed QP value for I-frame types (0: default)

func (e *Qsvh264Enc) GetMaxQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

// max-qp-p (uint)
//
// Maximum allowed QP value for P-frame types (0: default)

func (e *Qsvh264Enc) GetMaxQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

// max-slice-size (uint)

func (e *Qsvh264Enc) GetMaxSliceSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-slice-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

// min-qp-b (uint)
//
// Minimum allowed QP value for B-frame types (0: default)

func (e *Qsvh264Enc) GetMinQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

// min-qp-i (uint)
//
// Minimum allowed QP value for I-frame types (0: default)

func (e *Qsvh264Enc) GetMinQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

// min-qp-p (uint)
//
// Minimum allowed QP value for P-frame types (0: default)

func (e *Qsvh264Enc) GetMinQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

// num-slice (uint)

func (e *Qsvh264Enc) GetNumSlice() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetNumSlice(value uint) error {
	return e.Element.SetProperty("num-slice", value)
}

// num-slice-b (uint)

func (e *Qsvh264Enc) GetNumSliceB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetNumSliceB(value uint) error {
	return e.Element.SetProperty("num-slice-b", value)
}

// num-slice-i (uint)

func (e *Qsvh264Enc) GetNumSliceI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetNumSliceI(value uint) error {
	return e.Element.SetProperty("num-slice-i", value)
}

// num-slice-p (uint)

func (e *Qsvh264Enc) GetNumSliceP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetNumSliceP(value uint) error {
	return e.Element.SetProperty("num-slice-p", value)
}

// qp-b (uint)
//
// Constant quantizer for B frames (0: default)

func (e *Qsvh264Enc) GetQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

// qp-i (uint)
//
// Constant quantizer for I frames (0: default)

func (e *Qsvh264Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)
//
// Constant quantizer for P frames (0: default)

func (e *Qsvh264Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// qvbr-quality (uint)
//
// Quality level used for "qvbr" rate-control mode (0: default)

func (e *Qsvh264Enc) GetQvbrQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qvbr-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetQvbrQuality(value uint) error {
	return e.Element.SetProperty("qvbr-quality", value)
}

// rate-control (GstQsvH264EncRateControl)
//
// Rate Control Method

func (e *Qsvh264Enc) GetRateControl() (GstQsvH264EncRateControl, error) {
	var value GstQsvH264EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264EncRateControl")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetRateControl(value GstQsvH264EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// rc-lookahead (uint)
//
// Number of frames to look ahead for Rate Control, used for "la_vbr", "la_icq", and "la_hrd" rate-control modes

func (e *Qsvh264Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

// rc-lookahead-ds (GstQsvH264EncRclookAheadDs)
//
// Downsampling method in look-ahead rate control

func (e *Qsvh264Enc) GetRcLookaheadDs() (GstQsvH264EncRclookAheadDs, error) {
	var value GstQsvH264EncRclookAheadDs
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead-ds")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264EncRclookAheadDs)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264EncRclookAheadDs")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetRcLookaheadDs(value GstQsvH264EncRclookAheadDs) error {
	e.Element.SetArg("rc-lookahead-ds", string(value))
	return nil
}

// ref-frames (uint)
//
// Number of reference frames (0: unspecified)

func (e *Qsvh264Enc) GetRefFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ref-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// trellis (GstQsvH264Trellis)

func (e *Qsvh264Enc) GetTrellis() (GstQsvH264Trellis, error) {
	var value GstQsvH264Trellis
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264Trellis)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264Trellis")
	}
	return value, nil
}

func (e *Qsvh264Enc) SetTrellis(value GstQsvH264Trellis) error {
	e.Element.SetArg("trellis", string(value))
	return nil
}

// ----- Constants -----

type GstQsvCodingOption string

const (
	GstQsvCodingOptionUnknown GstQsvCodingOption = "unknown" // unknown (0) – Unknown
	GstQsvCodingOptionOn GstQsvCodingOption = "on" // on (16) – On
	GstQsvCodingOptionOff GstQsvCodingOption = "off" // off (32) – Off
)

type GstQsvH264EncRclookAheadDs string

const (
	GstQsvH264EncRclookAheadDsUnknown GstQsvH264EncRclookAheadDs = "unknown" // unknown (0) – Unknown
	GstQsvH264EncRclookAheadDsOff GstQsvH264EncRclookAheadDs = "off" // off (1) – Do not use down sampling
	GstQsvH264EncRclookAheadDs2X GstQsvH264EncRclookAheadDs = "2x" // 2x (2) – Down sample frames two times before estimation
	GstQsvH264EncRclookAheadDs4X GstQsvH264EncRclookAheadDs = "4x" // 4x (3) – Down sample frames four times before estimation
)

type GstQsvH264EncRateControl string

const (
	GstQsvH264EncRateControlCbr GstQsvH264EncRateControl = "cbr" // cbr (1) – Constant Bitrate
	GstQsvH264EncRateControlVbr GstQsvH264EncRateControl = "vbr" // vbr (2) – Variable Bitrate
	GstQsvH264EncRateControlCqp GstQsvH264EncRateControl = "cqp" // cqp (3) – Constant Quantizer
	GstQsvH264EncRateControlAvbr GstQsvH264EncRateControl = "avbr" // avbr (4) – Average Variable Bitrate
	GstQsvH264EncRateControlLaVbr GstQsvH264EncRateControl = "la-vbr" // la-vbr (8) – VBR with look ahead (Non HRD compliant)
	GstQsvH264EncRateControlIcq GstQsvH264EncRateControl = "icq" // icq (9) – Intelligent CQP
	GstQsvH264EncRateControlVcm GstQsvH264EncRateControl = "vcm" // vcm (10) – Video Conferencing Mode (Non HRD compliant)
	GstQsvH264EncRateControlLaIcq GstQsvH264EncRateControl = "la-icq" // la-icq (11) – Intelligent CQP with LA (Non HRD compliant)
	GstQsvH264EncRateControlLaHrd GstQsvH264EncRateControl = "la-hrd" // la-hrd (13) – HRD compliant LA
	GstQsvH264EncRateControlQvbr GstQsvH264EncRateControl = "qvbr" // qvbr (14) – VBR with CQP
)

type GstQsvH264EncSeiInsertMode string

const (
	GstQsvH264EncSeiInsertModeInsert GstQsvH264EncSeiInsertMode = "insert" // insert (0) – Insert SEI
	GstQsvH264EncSeiInsertModeInsertAndDrop GstQsvH264EncSeiInsertMode = "insert-and-drop" // insert-and-drop (1) – Insert SEI and remove corresponding meta from output buffer
	GstQsvH264EncSeiInsertModeDisabled GstQsvH264EncSeiInsertMode = "disabled" // disabled (2) – Disable SEI insertion
)

type GstQsvH264Trellis string

const (
	GstQsvH264TrellisUnknown GstQsvH264Trellis = "unknown" // unknown (0x00000000) – Unknown
	GstQsvH264TrellisOff GstQsvH264Trellis = "off" // off (0x00000001) – Disable for all frame types
	GstQsvH264TrellisI GstQsvH264Trellis = "i" // i (0x00000002) – Enable for I frames
	GstQsvH264TrellisP GstQsvH264Trellis = "p" // p (0x00000004) – Enable for P frames
	GstQsvH264TrellisB GstQsvH264Trellis = "b" // b (0x00000008) – Enable for B frames
)

