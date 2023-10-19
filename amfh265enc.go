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

type Amfh265Enc struct {
	Element *gst.Element
}

func NewAmfh265Enc() (*Amfh265Enc, error) {
	e, err := gst.NewElement("amfh265enc")
	if err != nil {
		return nil, err
	}
	return &Amfh265Enc{Element: e}, nil
}

func NewAmfh265EncWithName(name string) (*Amfh265Enc, error) {
	e, err := gst.NewElementWithName("amfh265enc", name)
	if err != nil {
		return nil, err
	}
	return &Amfh265Enc{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID (Locally Unique Identifier) of associated GPU

func (e *Amfh265Enc) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// aud (bool)
//
// Use AU (Access Unit) delimiter

func (e *Amfh265Enc) GetAud() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aud")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Amfh265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

// bitrate (uint)
//
// Target bitrate in kbit/sec (0: USAGE default)

func (e *Amfh265Enc) GetBitrate() (uint, error) {
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

func (e *Amfh265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// gop-size (uint)
//
// Number of pictures within a GOP

func (e *Amfh265Enc) GetGopSize() (uint, error) {
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

func (e *Amfh265Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

// max-bitrate (uint)
//
// Maximum bitrate in kbit/sec (0: USAGE default)

func (e *Amfh265Enc) GetMaxBitrate() (uint, error) {
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

func (e *Amfh265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-qp-i (int)
//
// Maximum allowed QP value for I frames (-1: USAGE default)

func (e *Amfh265Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh265Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

// max-qp-p (int)
//
// Maximum allowed QP value for P frames (-1: USAGE default)

func (e *Amfh265Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh265Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

// min-qp-i (int)
//
// Minimum allowed QP value for I frames (-1: USAGE default)

func (e *Amfh265Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh265Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

// min-qp-p (int)
//
// Minimum allowed QP value for P frames (-1: USAGE default)

func (e *Amfh265Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh265Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

// preset (GstAmfH265EncPreset)
//
// Preset

func (e *Amfh265Enc) GetPreset() (GstAmfH265EncPreset, error) {
	var value GstAmfH265EncPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH265EncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH265EncPreset")
	}
	return value, nil
}

func (e *Amfh265Enc) SetPreset(value GstAmfH265EncPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

// qp-i (uint)
//
// Constant QP for I frames

func (e *Amfh265Enc) GetQpI() (uint, error) {
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

func (e *Amfh265Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)
//
// Constant QP for P frames

func (e *Amfh265Enc) GetQpP() (uint, error) {
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

func (e *Amfh265Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// rate-control (GstAmfH265EncRateControl)
//
// Rate Control Method

func (e *Amfh265Enc) GetRateControl() (GstAmfH265EncRateControl, error) {
	var value GstAmfH265EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH265EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH265EncRateControl")
	}
	return value, nil
}

func (e *Amfh265Enc) SetRateControl(value GstAmfH265EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// ref-frames (uint)
//
// Number of reference frames

func (e *Amfh265Enc) GetRefFrames() (uint, error) {
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

func (e *Amfh265Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// usage (GstAmfH265EncUsage)
//
// Target usage

func (e *Amfh265Enc) GetUsage() (GstAmfH265EncUsage, error) {
	var value GstAmfH265EncUsage
	var ok bool
	v, err := e.Element.GetProperty("usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH265EncUsage)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH265EncUsage")
	}
	return value, nil
}

func (e *Amfh265Enc) SetUsage(value GstAmfH265EncUsage) error {
	e.Element.SetArg("usage", string(value))
	return nil
}

// ----- Constants -----

type GstAmfH265EncUsage string

const (
	GstAmfH265EncUsageTranscoding GstAmfH265EncUsage = "transcoding" // transcoding (0) – Transcoding
	GstAmfH265EncUsageUltraLowLatency GstAmfH265EncUsage = "ultra-low-latency" // ultra-low-latency (1) – Ultra Low Latency
	GstAmfH265EncUsageLowLatency GstAmfH265EncUsage = "low-latency" // low-latency (2) – Low Latency
	GstAmfH265EncUsageWebcam GstAmfH265EncUsage = "webcam" // webcam (3) – Webcam
)

type GstAmfH265EncPreset string

const (
	GstAmfH265EncPresetDefault GstAmfH265EncPreset = "default" // default (-1) – Default, depends on USAGE
	GstAmfH265EncPresetQuality GstAmfH265EncPreset = "quality" // quality (0) – Quality
	GstAmfH265EncPresetBalanced GstAmfH265EncPreset = "balanced" // balanced (5) – Balanced
	GstAmfH265EncPresetSpeed GstAmfH265EncPreset = "speed" // speed (10) – Speed
)

type GstAmfH265EncRateControl string

const (
	GstAmfH265EncRateControlDefault GstAmfH265EncRateControl = "default" // default (-1) – Default, depends on Usage
	GstAmfH265EncRateControlCqp GstAmfH265EncRateControl = "cqp" // cqp (0) – Constant QP
	GstAmfH265EncRateControlLcvbr GstAmfH265EncRateControl = "lcvbr" // lcvbr (1) – Latency Constrained VBR
	GstAmfH265EncRateControlVbr GstAmfH265EncRateControl = "vbr" // vbr (2) – Peak Constrained VBR
	GstAmfH265EncRateControlCbr GstAmfH265EncRateControl = "cbr" // cbr (3) – Constant Bitrate
)

