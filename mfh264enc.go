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

type Mfh264Enc struct {
	Element *gst.Element
}

func NewMfh264Enc() (*Mfh264Enc, error) {
	e, err := gst.NewElement("mfh264enc")
	if err != nil {
		return nil, err
	}
	return &Mfh264Enc{Element: e}, nil
}

func NewMfh264EncWithName(name string) (*Mfh264Enc, error) {
	e, err := gst.NewElementWithName("mfh264enc", name)
	if err != nil {
		return nil, err
	}
	return &Mfh264Enc{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID for this elemenet

func (e *Mfh264Enc) GetAdapterLuid() (int64, error) {
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

// adaptive-mode (GstMfh264EncAdaptiveMode)

func (e *Mfh264Enc) GetAdaptiveMode() (GstMfh264EncAdaptiveMode, error) {
	var value GstMfh264EncAdaptiveMode
	var ok bool
	v, err := e.Element.GetProperty("adaptive-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMfh264EncAdaptiveMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMfh264EncAdaptiveMode")
	}
	return value, nil
}

func (e *Mfh264Enc) SetAdaptiveMode(value GstMfh264EncAdaptiveMode) error {
	e.Element.SetArg("adaptive-mode", string(value))
	return nil
}

// bframes (uint)

func (e *Mfh264Enc) GetBframes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

// bitrate (uint)
//
// Bitrate in kbit/sec

func (e *Mfh264Enc) GetBitrate() (uint, error) {
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

func (e *Mfh264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// cabac (bool)

func (e *Mfh264Enc) GetCabac() (bool, error) {
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

func (e *Mfh264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

// d3d11-aware (bool)
//
// Whether element supports Direct3D11 texture as an input or not

func (e *Mfh264Enc) GetD3D11Aware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("d3d11-aware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// gop-size (int)

func (e *Mfh264Enc) GetGopSize() (int, error) {
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

func (e *Mfh264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// low-latency (bool)

func (e *Mfh264Enc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mfh264Enc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// max-bitrate (uint)

func (e *Mfh264Enc) GetMaxBitrate() (uint, error) {
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

func (e *Mfh264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-qp (uint)

func (e *Mfh264Enc) GetMaxQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

// min-qp (uint)

func (e *Mfh264Enc) GetMinQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

// qp (uint)

func (e *Mfh264Enc) GetQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetQp(value uint) error {
	return e.Element.SetProperty("qp", value)
}

// qp-b (uint)

func (e *Mfh264Enc) GetQpB() (uint, error) {
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

func (e *Mfh264Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

// qp-i (uint)

func (e *Mfh264Enc) GetQpI() (uint, error) {
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

func (e *Mfh264Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)

func (e *Mfh264Enc) GetQpP() (uint, error) {
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

func (e *Mfh264Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// quality-vs-speed (uint)

func (e *Mfh264Enc) GetQualityVsSpeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("quality-vs-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetQualityVsSpeed(value uint) error {
	return e.Element.SetProperty("quality-vs-speed", value)
}

// rc-mode (GstMfh264EncRcmode)

func (e *Mfh264Enc) GetRcMode() (GstMfh264EncRcmode, error) {
	var value GstMfh264EncRcmode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMfh264EncRcmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMfh264EncRcmode")
	}
	return value, nil
}

func (e *Mfh264Enc) SetRcMode(value GstMfh264EncRcmode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

// ref (uint)

func (e *Mfh264Enc) GetRef() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ref")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetRef(value uint) error {
	return e.Element.SetProperty("ref", value)
}

// vbv-buffer-size (uint)

func (e *Mfh264Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Mfh264Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

// ----- Constants -----

type GstMfh264EncAdaptiveMode string

const (
	GstMfh264EncAdaptiveModeNone GstMfh264EncAdaptiveMode = "none" // none (0) – None
	GstMfh264EncAdaptiveModeFramerate GstMfh264EncAdaptiveMode = "framerate" // framerate (1) – Adaptively change the frame rate
)

type GstMfh264EncRcmode string

const (
	GstMfh264EncRcmodeCbr GstMfh264EncRcmode = "cbr" // cbr (0) – Constant bitrate
	GstMfh264EncRcmodePcvbr GstMfh264EncRcmode = "pcvbr" // pcvbr (1) – Peak Constrained variable bitrate
	GstMfh264EncRcmodeUvbr GstMfh264EncRcmode = "uvbr" // uvbr (2) – Unconstrained variable bitrate
	GstMfh264EncRcmodeQvbr GstMfh264EncRcmode = "qvbr" // qvbr (3) – Quality-based variable bitrate
)

