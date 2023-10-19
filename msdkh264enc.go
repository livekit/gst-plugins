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

type Msdkh264Enc struct {
	Element *gst.Element
}

func NewMsdkh264Enc() (*Msdkh264Enc, error) {
	e, err := gst.NewElement("msdkh264enc")
	if err != nil {
		return nil, err
	}
	return &Msdkh264Enc{Element: e}, nil
}

func NewMsdkh264EncWithName(name string) (*Msdkh264Enc, error) {
	e, err := gst.NewElementWithName("msdkh264enc", name)
	if err != nil {
		return nil, err
	}
	return &Msdkh264Enc{Element: e}, nil
}

// ----- Properties -----

// b-pyramid (bool)
//
// Enable B-Pyramid Reference structure

func (e *Msdkh264Enc) GetBPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

// cabac (bool)
//
// Enable CABAC entropy coding

func (e *Msdkh264Enc) GetCabac() (bool, error) {
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

func (e *Msdkh264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

// dblk-idc (uint)
//
// Option of disable deblocking idc

func (e *Msdkh264Enc) GetDblkIdc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dblk-idc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetDblkIdc(value uint) error {
	return e.Element.SetProperty("dblk-idc", value)
}

// frame-packing (GstMsdkH264EncFramePacking)
//
// Set frame packing mode for Stereoscopic content

func (e *Msdkh264Enc) GetFramePacking() (GstMsdkH264EncFramePacking, error) {
	var value GstMsdkH264EncFramePacking
	var ok bool
	v, err := e.Element.GetProperty("frame-packing")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkH264EncFramePacking)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkH264EncFramePacking")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetFramePacking(value GstMsdkH264EncFramePacking) error {
	e.Element.SetArg("frame-packing", string(value))
	return nil
}

// intra-refresh-type (GstMsdkEncIntraRefreshType)
//
// Set intra refresh type

func (e *Msdkh264Enc) GetIntraRefreshType() (interface{}, error) {
	return e.Element.GetProperty("intra-refresh-type")
}

func (e *Msdkh264Enc) SetIntraRefreshType(value interface{}) error {
	return e.Element.SetProperty("intra-refresh-type", value)
}

// low-power (bool)
//
// Enable low power mode (DEPRECATED, use tune instead)

func (e *Msdkh264Enc) GetLowPower() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-power")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetLowPower(value bool) error {
	return e.Element.SetProperty("low-power", value)
}

// max-qp (uint)
//
// Maximum quantizer scale for I/P/B frames

func (e *Msdkh264Enc) GetMaxQp() (uint, error) {
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

func (e *Msdkh264Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

// max-qp-b (uint)

func (e *Msdkh264Enc) GetMaxQpB() (uint, error) {
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

func (e *Msdkh264Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

// max-qp-i (uint)

func (e *Msdkh264Enc) GetMaxQpI() (uint, error) {
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

func (e *Msdkh264Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

// max-qp-p (uint)

func (e *Msdkh264Enc) GetMaxQpP() (uint, error) {
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

func (e *Msdkh264Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

// max-slice-size (uint)
//
// Maximum slice size in bytes (if enabled MSDK will ignore the control over num-slices)

func (e *Msdkh264Enc) GetMaxSliceSize() (uint, error) {
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

func (e *Msdkh264Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

// min-qp (uint)
//
// Minimal quantizer scale for I/P/B frames

func (e *Msdkh264Enc) GetMinQp() (uint, error) {
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

func (e *Msdkh264Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

// min-qp-b (uint)

func (e *Msdkh264Enc) GetMinQpB() (uint, error) {
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

func (e *Msdkh264Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

// min-qp-i (uint)

func (e *Msdkh264Enc) GetMinQpI() (uint, error) {
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

func (e *Msdkh264Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

// min-qp-p (uint)

func (e *Msdkh264Enc) GetMinQpP() (uint, error) {
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

func (e *Msdkh264Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

// p-pyramid (bool)
//
// Enable P-Pyramid Reference structure

func (e *Msdkh264Enc) GetPPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("p-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetPPyramid(value bool) error {
	return e.Element.SetProperty("p-pyramid", value)
}

// pic-timing-sei (bool)

func (e *Msdkh264Enc) GetPicTimingSei() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pic-timing-sei")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetPicTimingSei(value bool) error {
	return e.Element.SetProperty("pic-timing-sei", value)
}

// rc-lookahead-ds (GstMsdkEncRclookAheadDownsampling)
//
// Down sampling mode in look ahead bitrate control

func (e *Msdkh264Enc) GetRcLookaheadDs() (GstMsdkEncRclookAheadDownsampling, error) {
	var value GstMsdkEncRclookAheadDownsampling
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead-ds")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncRclookAheadDownsampling)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncRclookAheadDownsampling")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetRcLookaheadDs(value GstMsdkEncRclookAheadDownsampling) error {
	e.Element.SetArg("rc-lookahead-ds", string(value))
	return nil
}

// trellis (GstMsdkEncTrellisQuantization)
//
// Enable Trellis Quantization

func (e *Msdkh264Enc) GetTrellis() (GstMsdkEncTrellisQuantization, error) {
	var value GstMsdkEncTrellisQuantization
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTrellisQuantization)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTrellisQuantization")
	}
	return value, nil
}

func (e *Msdkh264Enc) SetTrellis(value GstMsdkEncTrellisQuantization) error {
	e.Element.SetArg("trellis", string(value))
	return nil
}

// tune (GstMsdkEncTuneMode)
//
// Encoder tuning option

func (e *Msdkh264Enc) GetTune() (interface{}, error) {
	return e.Element.GetProperty("tune")
}

func (e *Msdkh264Enc) SetTune(value interface{}) error {
	return e.Element.SetProperty("tune", value)
}

// ----- Constants -----

type GstMsdkEncTrellisQuantization string

const (
	GstMsdkEncTrellisQuantizationNone GstMsdkEncTrellisQuantization = "None" // None (0x00000000) – Disable for all frames
	GstMsdkEncTrellisQuantizationI GstMsdkEncTrellisQuantization = "i" // i (0x00000002) – Enable for I frames
	GstMsdkEncTrellisQuantizationP GstMsdkEncTrellisQuantization = "p" // p (0x00000004) – Enable for P frames
	GstMsdkEncTrellisQuantizationB GstMsdkEncTrellisQuantization = "b" // b (0x00000008) – Enable for B frames
)

type GstMsdkH264EncFramePacking string

const (
	GstMsdkH264EncFramePackingNone GstMsdkH264EncFramePacking = "none" // none (-1) – None (default)
	GstMsdkH264EncFramePackingSideBySide GstMsdkH264EncFramePacking = "side-by-side" // side-by-side (3) – Side by Side
	GstMsdkH264EncFramePackingTopBottom GstMsdkH264EncFramePacking = "top-bottom" // top-bottom (7) – Top Bottom
)

type GstMsdkEncRclookAheadDownsampling string

const (
	GstMsdkEncRclookAheadDownsamplingDefault GstMsdkEncRclookAheadDownsampling = "default" // default (0) – SDK desides what to do
	GstMsdkEncRclookAheadDownsamplingOff GstMsdkEncRclookAheadDownsampling = "off" // off (1) – No downsampling
	GstMsdkEncRclookAheadDownsampling2X GstMsdkEncRclookAheadDownsampling = "2x" // 2x (2) – Down sample 2-times before estimation
	GstMsdkEncRclookAheadDownsampling4X GstMsdkEncRclookAheadDownsampling = "4x" // 4x (3) – Down sample 4-times before estimation
)

