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

type Qsvvp9Enc struct {
	Element *gst.Element
}

func NewQsvvp9Enc() (*Qsvvp9Enc, error) {
	e, err := gst.NewElement("qsvvp9enc")
	if err != nil {
		return nil, err
	}
	return &Qsvvp9Enc{Element: e}, nil
}

func NewQsvvp9EncWithName(name string) (*Qsvvp9Enc, error) {
	e, err := gst.NewElementWithName("qsvvp9enc", name)
	if err != nil {
		return nil, err
	}
	return &Qsvvp9Enc{Element: e}, nil
}

// ----- Properties -----

// bitrate (uint)
//
// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")

func (e *Qsvvp9Enc) GetBitrate() (uint, error) {
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

func (e *Qsvvp9Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// gop-size (uint)
//
// Number of pictures within a GOP (0: unspecified)

func (e *Qsvvp9Enc) GetGopSize() (uint, error) {
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

func (e *Qsvvp9Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

// icq-quality (uint)
//
// Intelligent Constant Quality for "icq" rate-control (0: default)

func (e *Qsvvp9Enc) GetIcqQuality() (uint, error) {
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

func (e *Qsvvp9Enc) SetIcqQuality(value uint) error {
	return e.Element.SetProperty("icq-quality", value)
}

// max-bitrate (uint)
//
// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")

func (e *Qsvvp9Enc) GetMaxBitrate() (uint, error) {
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

func (e *Qsvvp9Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// qp-i (uint)
//
// Constant quantizer for I frames (0: default)

func (e *Qsvvp9Enc) GetQpI() (uint, error) {
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

func (e *Qsvvp9Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)
//
// Constant quantizer for P frames (0: default)

func (e *Qsvvp9Enc) GetQpP() (uint, error) {
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

func (e *Qsvvp9Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// rate-control (GstQsvVp9EncRateControl)
//
// Rate Control Method

func (e *Qsvvp9Enc) GetRateControl() (GstQsvVp9EncRateControl, error) {
	var value GstQsvVp9EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvVp9EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvVp9EncRateControl")
	}
	return value, nil
}

func (e *Qsvvp9Enc) SetRateControl(value GstQsvVp9EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// ref-frames (uint)
//
// Number of reference frames (0: unspecified)

func (e *Qsvvp9Enc) GetRefFrames() (uint, error) {
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

func (e *Qsvvp9Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// ----- Constants -----

type GstQsvVp9EncRateControl string

const (
	GstQsvVp9EncRateControlCbr GstQsvVp9EncRateControl = "cbr" // cbr (1) – Constant Bitrate
	GstQsvVp9EncRateControlVbr GstQsvVp9EncRateControl = "vbr" // vbr (2) – Variable Bitrate
	GstQsvVp9EncRateControlCqp GstQsvVp9EncRateControl = "cqp" // cqp (3) – Constant Quantizer
	GstQsvVp9EncRateControlIcq GstQsvVp9EncRateControl = "icq" // icq (9) – Intelligent CQP
)

