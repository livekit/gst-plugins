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

type Mfh265Enc struct {
	Element *gst.Element
}

func NewMfh265Enc() (*Mfh265Enc, error) {
	e, err := gst.NewElement("mfh265enc")
	if err != nil {
		return nil, err
	}
	return &Mfh265Enc{Element: e}, nil
}

func NewMfh265EncWithName(name string) (*Mfh265Enc, error) {
	e, err := gst.NewElementWithName("mfh265enc", name)
	if err != nil {
		return nil, err
	}
	return &Mfh265Enc{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID for this elemenet

func (e *Mfh265Enc) GetAdapterLuid() (int64, error) {
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

// bframes (uint)
//
// The maximum number of consecutive B frames

func (e *Mfh265Enc) GetBframes() (uint, error) {
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

func (e *Mfh265Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

// bitrate (uint)
//
// Bitrate in kbit/sec

func (e *Mfh265Enc) GetBitrate() (uint, error) {
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

func (e *Mfh265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// d3d11-aware (bool)
//
// Whether element supports Direct3D11 texture as an input or not

func (e *Mfh265Enc) GetD3D11Aware() (bool, error) {
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
//
// The number of pictures from one GOP header to the next. Depending on GPU vendor implementation, zero gop-size might produce only one keyframe at the beginning (-1 for automatic)

func (e *Mfh265Enc) GetGopSize() (int, error) {
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

func (e *Mfh265Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// low-latency (bool)
//
// Enable low latency encoding

func (e *Mfh265Enc) GetLowLatency() (bool, error) {
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

func (e *Mfh265Enc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

// max-bitrate (uint)
//
// The maximum bitrate applied when rc-mode is "pcvbr" in kbit/sec (0 = MFT default)

func (e *Mfh265Enc) GetMaxBitrate() (uint, error) {
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

func (e *Mfh265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-qp (uint)

func (e *Mfh265Enc) GetMaxQp() (uint, error) {
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

func (e *Mfh265Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

// min-qp (uint)

func (e *Mfh265Enc) GetMinQp() (uint, error) {
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

func (e *Mfh265Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

// qp (uint)
//
// QP applied when rc-mode is "qvbr"

func (e *Mfh265Enc) GetQp() (uint, error) {
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

func (e *Mfh265Enc) SetQp(value uint) error {
	return e.Element.SetProperty("qp", value)
}

// qp-b (uint)

func (e *Mfh265Enc) GetQpB() (uint, error) {
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

func (e *Mfh265Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

// qp-i (uint)

func (e *Mfh265Enc) GetQpI() (uint, error) {
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

func (e *Mfh265Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (uint)

func (e *Mfh265Enc) GetQpP() (uint, error) {
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

func (e *Mfh265Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

// quality-vs-speed (uint)
//
// Quality and speed tradeoff, [0, 33]: Low complexity, [34, 66]: Medium complexity, [67, 100]: High complexity

func (e *Mfh265Enc) GetQualityVsSpeed() (uint, error) {
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

func (e *Mfh265Enc) SetQualityVsSpeed(value uint) error {
	return e.Element.SetProperty("quality-vs-speed", value)
}

// rc-mode (GstMfh265EncRcmode)
//
// Rate Control Mode

func (e *Mfh265Enc) GetRcMode() (GstMfh265EncRcmode, error) {
	var value GstMfh265EncRcmode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMfh265EncRcmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMfh265EncRcmode")
	}
	return value, nil
}

func (e *Mfh265Enc) SetRcMode(value GstMfh265EncRcmode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

// ref (uint)

func (e *Mfh265Enc) GetRef() (uint, error) {
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

func (e *Mfh265Enc) SetRef(value uint) error {
	return e.Element.SetProperty("ref", value)
}

// vbv-buffer-size (uint)
//
// VBV(HRD) Buffer Size in bytes (0 = MFT default)

func (e *Mfh265Enc) GetVbvBufferSize() (uint, error) {
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

func (e *Mfh265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

// ----- Constants -----

type GstMfh265EncRcmode string

const (
	GstMfh265EncRcmodeCbr GstMfh265EncRcmode = "cbr" // cbr (0) – Constant bitrate
	GstMfh265EncRcmodeQvbr GstMfh265EncRcmode = "qvbr" // qvbr (1) – Quality-based variable bitrate
)

