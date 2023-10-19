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

type Qsvav1Enc struct {
	Element *gst.Element
}

func NewQsvav1Enc() (*Qsvav1Enc, error) {
	e, err := gst.NewElement("qsvav1enc")
	if err != nil {
		return nil, err
	}
	return &Qsvav1Enc{Element: e}, nil
}

func NewQsvav1EncWithName(name string) (*Qsvav1Enc, error) {
	e, err := gst.NewElementWithName("qsvav1enc", name)
	if err != nil {
		return nil, err
	}
	return &Qsvav1Enc{Element: e}, nil
}

// ----- Properties -----

// bitrate (uint)
//
// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")

func (e *Qsvav1Enc) GetBitrate() (uint, error) {
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

func (e *Qsvav1Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// gop-size (uint)
//
// Number of pictures within a GOP (0: unspecified)

func (e *Qsvav1Enc) GetGopSize() (uint, error) {
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

func (e *Qsvav1Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

// max-bitrate (uint)
//
// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")

func (e *Qsvav1Enc) GetMaxBitrate() (uint, error) {
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

func (e *Qsvav1Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// qp-i (uint)
//
// Constant quantizer for I frames (0: default)

func (e *Qsvav1Enc) GetQpI() (uint, error) {
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

func (e *Qsvav1Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)
//
// Constant quantizer for P frames (0: default)

func (e *Qsvav1Enc) GetQpP() (uint, error) {
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

func (e *Qsvav1Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// rate-control (GstQsvAv1EncRateControl)
//
// Rate Control Method

func (e *Qsvav1Enc) GetRateControl() (GstQsvAv1EncRateControl, error) {
	var value GstQsvAv1EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvAv1EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvAv1EncRateControl")
	}
	return value, nil
}

func (e *Qsvav1Enc) SetRateControl(value GstQsvAv1EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// ref-frames (uint)
//
// Number of reference frames (0: unspecified)

func (e *Qsvav1Enc) GetRefFrames() (uint, error) {
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

func (e *Qsvav1Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// ----- Constants -----

type GstQsvAv1EncRateControl string

const (
	GstQsvAv1EncRateControlCbr GstQsvAv1EncRateControl = "cbr" // cbr (1) – Constant Bitrate
	GstQsvAv1EncRateControlVbr GstQsvAv1EncRateControl = "vbr" // vbr (2) – Variable Bitrate
	GstQsvAv1EncRateControlCqp GstQsvAv1EncRateControl = "cqp" // cqp (3) – Constant Quantizer
)

