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

type Amfh264Enc struct {
	Element *gst.Element
}

func NewAmfh264Enc() (*Amfh264Enc, error) {
	e, err := gst.NewElement("amfh264enc")
	if err != nil {
		return nil, err
	}
	return &Amfh264Enc{Element: e}, nil
}

func NewAmfh264EncWithName(name string) (*Amfh264Enc, error) {
	e, err := gst.NewElementWithName("amfh264enc", name)
	if err != nil {
		return nil, err
	}
	return &Amfh264Enc{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID (Locally Unique Identifier) of associated GPU

func (e *Amfh264Enc) GetAdapterLuid() (int64, error) {
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

func (e *Amfh264Enc) GetAud() (bool, error) {
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

func (e *Amfh264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

// bitrate (uint)
//
// Target bitrate in kbit/sec (0: USAGE default)

func (e *Amfh264Enc) GetBitrate() (uint, error) {
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

func (e *Amfh264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// cabac (bool)
//
// Enable CABAC entropy coding

func (e *Amfh264Enc) GetCabac() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cabac")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Amfh264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

// gop-size (int)
//
// Number of pictures within a GOP (-1: USAGE default)

func (e *Amfh264Enc) GetGopSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// max-bitrate (uint)
//
// Maximum bitrate in kbit/sec (0: USAGE default)

func (e *Amfh264Enc) GetMaxBitrate() (uint, error) {
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

func (e *Amfh264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-qp (int)
//
// Maximum allowed QP value (-1: USAGE default)

func (e *Amfh264Enc) GetMaxQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh264Enc) SetMaxQp(value int) error {
	return e.Element.SetProperty("max-qp", value)
}

// min-qp (int)
//
// Minimum allowed QP value (-1: USAGE default)

func (e *Amfh264Enc) GetMinQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Amfh264Enc) SetMinQp(value int) error {
	return e.Element.SetProperty("min-qp", value)
}

// preset (GstAmfH264EncPreset)
//
// Preset

func (e *Amfh264Enc) GetPreset() (GstAmfH264EncPreset, error) {
	var value GstAmfH264EncPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH264EncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH264EncPreset")
	}
	return value, nil
}

func (e *Amfh264Enc) SetPreset(value GstAmfH264EncPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

// qp-i (uint)
//
// Constant QP for I frames

func (e *Amfh264Enc) GetQpI() (uint, error) {
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

func (e *Amfh264Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)
//
// Constant QP for P frames

func (e *Amfh264Enc) GetQpP() (uint, error) {
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

func (e *Amfh264Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// rate-control (GstAmfH264EncRateControl)
//
// Rate Control Method

func (e *Amfh264Enc) GetRateControl() (GstAmfH264EncRateControl, error) {
	var value GstAmfH264EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH264EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH264EncRateControl")
	}
	return value, nil
}

func (e *Amfh264Enc) SetRateControl(value GstAmfH264EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// ref-frames (uint)
//
// Number of reference frames

func (e *Amfh264Enc) GetRefFrames() (uint, error) {
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

func (e *Amfh264Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

// usage (GstAmfH264EncUsage)
//
// Target usage

func (e *Amfh264Enc) GetUsage() (GstAmfH264EncUsage, error) {
	var value GstAmfH264EncUsage
	var ok bool
	v, err := e.Element.GetProperty("usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH264EncUsage)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH264EncUsage")
	}
	return value, nil
}

func (e *Amfh264Enc) SetUsage(value GstAmfH264EncUsage) error {
	e.Element.SetArg("usage", string(value))
	return nil
}

// ----- Constants -----

type GstAmfH264EncPreset string

const (
	GstAmfH264EncPresetDefault GstAmfH264EncPreset = "default" // default (-1) – Default, depends on USAGE
	GstAmfH264EncPresetBalanced GstAmfH264EncPreset = "balanced" // balanced (0) – Balanced
	GstAmfH264EncPresetSpeed GstAmfH264EncPreset = "speed" // speed (1) – Speed
	GstAmfH264EncPresetQuality GstAmfH264EncPreset = "quality" // quality (2) – Quality
)

type GstAmfH264EncRateControl string

const (
	GstAmfH264EncRateControlDefault GstAmfH264EncRateControl = "default" // default (-1) – Default, depends on Usage
	GstAmfH264EncRateControlCqp GstAmfH264EncRateControl = "cqp" // cqp (0) – Constant QP
	GstAmfH264EncRateControlCbr GstAmfH264EncRateControl = "cbr" // cbr (1) – Constant Bitrate
	GstAmfH264EncRateControlVbr GstAmfH264EncRateControl = "vbr" // vbr (2) – Peak Constrained VBR
	GstAmfH264EncRateControlLcvbr GstAmfH264EncRateControl = "lcvbr" // lcvbr (3) – Latency Constrained VBR
)

type GstAmfH264EncUsage string

const (
	GstAmfH264EncUsageTranscoding GstAmfH264EncUsage = "transcoding" // transcoding (0) – Transcoding
	GstAmfH264EncUsageUltraLowLatency GstAmfH264EncUsage = "ultra-low-latency" // ultra-low-latency (1) – Ultra Low Latency
	GstAmfH264EncUsageLowLatency GstAmfH264EncUsage = "low-latency" // low-latency (2) – Low Latency
	GstAmfH264EncUsageWebcam GstAmfH264EncUsage = "webcam" // webcam (3) – Webcam
)

