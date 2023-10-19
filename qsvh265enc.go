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

type Qsvh265Enc struct {
	Element *gst.Element
}

func NewQsvh265Enc() (*Qsvh265Enc, error) {
	e, err := gst.NewElement("qsvh265enc")
	if err != nil {
		return nil, err
	}
	return &Qsvh265Enc{Element: e}, nil
}

func NewQsvh265EncWithName(name string) (*Qsvh265Enc, error) {
	e, err := gst.NewElementWithName("qsvh265enc", name)
	if err != nil {
		return nil, err
	}
	return &Qsvh265Enc{Element: e}, nil
}

// ----- Properties -----

// b-frames (uint)
//
// Number of B frames between I and P frames

func (e *Qsvh265Enc) GetBFrames() (uint, error) {
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

func (e *Qsvh265Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

// bitrate (uint)
//
// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")

func (e *Qsvh265Enc) GetBitrate() (uint, error) {
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

func (e *Qsvh265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// cc-insert (GstQsvH265EncSeiInsertMode)
//
// Closed Caption Insert mode. Only CEA-708 RAW format is supported for now

func (e *Qsvh265Enc) GetCcInsert() (GstQsvH265EncSeiInsertMode, error) {
	var value GstQsvH265EncSeiInsertMode
	var ok bool
	v, err := e.Element.GetProperty("cc-insert")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH265EncSeiInsertMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH265EncSeiInsertMode")
	}
	return value, nil
}

func (e *Qsvh265Enc) SetCcInsert(value GstQsvH265EncSeiInsertMode) error {
	e.Element.SetArg("cc-insert", string(value))
	return nil
}

// disable-hrd-conformance (bool)
//
// Allow NAL HRD non-conformant stream

func (e *Qsvh265Enc) GetDisableHrdConformance() (bool, error) {
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

func (e *Qsvh265Enc) SetDisableHrdConformance(value bool) error {
	return e.Element.SetProperty("disable-hrd-conformance", value)
}

// gop-size (uint)
//
// Number of pictures within a GOP (0: unspecified)

func (e *Qsvh265Enc) GetGopSize() (uint, error) {
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

func (e *Qsvh265Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

// icq-quality (uint)
//
// Intelligent Constant Quality for "icq" rate-control (0: default)

func (e *Qsvh265Enc) GetIcqQuality() (uint, error) {
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

func (e *Qsvh265Enc) SetIcqQuality(value uint) error {
	return e.Element.SetProperty("icq-quality", value)
}

// idr-interval (uint)
//
// IDR-frame interval in terms of I-frames. 0: only first I-frame is is an IDR frame, 1: every I-frame is an IDR frame, N: "N - 1" I-frames are inserted between IDR-frames

func (e *Qsvh265Enc) GetIdrInterval() (uint, error) {
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

func (e *Qsvh265Enc) SetIdrInterval(value uint) error {
	return e.Element.SetProperty("idr-interval", value)
}

// max-bitrate (uint)
//
// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")

func (e *Qsvh265Enc) GetMaxBitrate() (uint, error) {
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

func (e *Qsvh265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-qp-b (uint)
//
// Maximum allowed QP value for B-frame types (0: default)

func (e *Qsvh265Enc) GetMaxQpB() (uint, error) {
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

func (e *Qsvh265Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

// max-qp-i (uint)
//
// Maximum allowed QP value for I-frame types (0: default)

func (e *Qsvh265Enc) GetMaxQpI() (uint, error) {
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

func (e *Qsvh265Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

// max-qp-p (uint)
//
// Maximum allowed QP value for P-frame types (0: default)

func (e *Qsvh265Enc) GetMaxQpP() (uint, error) {
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

func (e *Qsvh265Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

// min-qp-b (uint)
//
// Minimum allowed QP value for B-frame types (0: default)

func (e *Qsvh265Enc) GetMinQpB() (uint, error) {
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

func (e *Qsvh265Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

// min-qp-i (uint)
//
// Minimum allowed QP value for I-frame types (0: default)

func (e *Qsvh265Enc) GetMinQpI() (uint, error) {
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

func (e *Qsvh265Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

// min-qp-p (uint)
//
// Minimum allowed QP value for P-frame types (0: default)

func (e *Qsvh265Enc) GetMinQpP() (uint, error) {
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

func (e *Qsvh265Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

// qp-b (uint)
//
// Constant quantizer for B frames (0: default)

func (e *Qsvh265Enc) GetQpB() (uint, error) {
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

func (e *Qsvh265Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

// qp-i (uint)
//
// Constant quantizer for I frames (0: default)

func (e *Qsvh265Enc) GetQpI() (uint, error) {
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

func (e *Qsvh265Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)
//
// Constant quantizer for P frames (0: default)

func (e *Qsvh265Enc) GetQpP() (uint, error) {
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

func (e *Qsvh265Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// qvbr-quality (uint)
//
// Quality level used for "qvbr" rate-control mode (0: default)

func (e *Qsvh265Enc) GetQvbrQuality() (uint, error) {
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

func (e *Qsvh265Enc) SetQvbrQuality(value uint) error {
	return e.Element.SetProperty("qvbr-quality", value)
}

// rate-control (GstQsvH265EncRateControl)
//
// Rate Control Method

func (e *Qsvh265Enc) GetRateControl() (GstQsvH265EncRateControl, error) {
	var value GstQsvH265EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH265EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH265EncRateControl")
	}
	return value, nil
}

func (e *Qsvh265Enc) SetRateControl(value GstQsvH265EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// ref-frames (uint)
//
// Number of reference frames (0: unspecified)

func (e *Qsvh265Enc) GetRefFrames() (uint, error) {
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

func (e *Qsvh265Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// ----- Constants -----

type GstQsvH265EncRateControl string

const (
	GstQsvH265EncRateControlCbr GstQsvH265EncRateControl = "cbr" // cbr (1) – Constant Bitrate
	GstQsvH265EncRateControlVbr GstQsvH265EncRateControl = "vbr" // vbr (2) – Variable Bitrate
	GstQsvH265EncRateControlCqp GstQsvH265EncRateControl = "cqp" // cqp (3) – Constant Quantizer
	GstQsvH265EncRateControlIcq GstQsvH265EncRateControl = "icq" // icq (9) – Intelligent CQP
	GstQsvH265EncRateControlVcm GstQsvH265EncRateControl = "vcm" // vcm (10) – Video Conferencing Mode (Non HRD compliant)
	GstQsvH265EncRateControlQvbr GstQsvH265EncRateControl = "qvbr" // qvbr (14) – VBR with CQP
)

type GstQsvH265EncSeiInsertMode string

const (
	GstQsvH265EncSeiInsertModeInsert GstQsvH265EncSeiInsertMode = "insert" // insert (0) – Insert SEI
	GstQsvH265EncSeiInsertModeInsertAndDrop GstQsvH265EncSeiInsertMode = "insert-and-drop" // insert-and-drop (1) – Insert SEI and remove corresponding meta from output buffer
	GstQsvH265EncSeiInsertModeDisabled GstQsvH265EncSeiInsertMode = "disabled" // disabled (2) – Disable SEI insertion
)

