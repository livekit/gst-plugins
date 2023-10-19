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

type Openh264Enc struct {
	Element *gst.Element
}

func NewOpenh264Enc() (*Openh264Enc, error) {
	e, err := gst.NewElement("openh264enc")
	if err != nil {
		return nil, err
	}
	return &Openh264Enc{Element: e}, nil
}

func NewOpenh264EncWithName(name string) (*Openh264Enc, error) {
	e, err := gst.NewElementWithName("openh264enc", name)
	if err != nil {
		return nil, err
	}
	return &Openh264Enc{Element: e}, nil
}

// ----- Properties -----

// adaptive-quantization (bool)
//
// Adaptive quantization

func (e *Openh264Enc) GetAdaptiveQuantization() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("adaptive-quantization")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openh264Enc) SetAdaptiveQuantization(value bool) error {
	return e.Element.SetProperty("adaptive-quantization", value)
}

// background-detection (bool)
//
// Background detection

func (e *Openh264Enc) GetBackgroundDetection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("background-detection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openh264Enc) SetBackgroundDetection(value bool) error {
	return e.Element.SetProperty("background-detection", value)
}

// bitrate (uint)
//
// Bitrate (in bits per second)

func (e *Openh264Enc) GetBitrate() (uint, error) {
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

func (e *Openh264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// complexity (GstOpenh264EncComplexity)
//
// Complexity

func (e *Openh264Enc) GetComplexity() (GstOpenh264EncComplexity, error) {
	var value GstOpenh264EncComplexity
	var ok bool
	v, err := e.Element.GetProperty("complexity")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenh264EncComplexity)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenh264EncComplexity")
	}
	return value, nil
}

func (e *Openh264Enc) SetComplexity(value GstOpenh264EncComplexity) error {
	e.Element.SetArg("complexity", string(value))
	return nil
}

// deblocking (GstOpenh264EncDeblockingModes)
//
// Deblocking mode

func (e *Openh264Enc) GetDeblocking() (GstOpenh264EncDeblockingModes, error) {
	var value GstOpenh264EncDeblockingModes
	var ok bool
	v, err := e.Element.GetProperty("deblocking")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenh264EncDeblockingModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenh264EncDeblockingModes")
	}
	return value, nil
}

func (e *Openh264Enc) SetDeblocking(value GstOpenh264EncDeblockingModes) error {
	e.Element.SetArg("deblocking", string(value))
	return nil
}

// enable-denoise (bool)
//
// Denoise control

func (e *Openh264Enc) GetEnableDenoise() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-denoise")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openh264Enc) SetEnableDenoise(value bool) error {
	return e.Element.SetProperty("enable-denoise", value)
}

// enable-frame-skip (bool)
//
// Skip frames to reach target bitrate

func (e *Openh264Enc) GetEnableFrameSkip() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-frame-skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openh264Enc) SetEnableFrameSkip(value bool) error {
	return e.Element.SetProperty("enable-frame-skip", value)
}

// gop-size (uint)
//
// Number of frames between intra frames

func (e *Openh264Enc) GetGopSize() (uint, error) {
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

func (e *Openh264Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

// max-bitrate (uint)
//
// Maximum Bitrate (in bits per second)

func (e *Openh264Enc) GetMaxBitrate() (uint, error) {
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

func (e *Openh264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-slice-size (uint)
//
// The maximum size of one slice (in bytes).

func (e *Openh264Enc) GetMaxSliceSize() (uint, error) {
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

func (e *Openh264Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

// multi-thread (uint)
//
// The number of threads.

func (e *Openh264Enc) GetMultiThread() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("multi-thread")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openh264Enc) SetMultiThread(value uint) error {
	return e.Element.SetProperty("multi-thread", value)
}

// num-slices (uint)
//
// The number of slices (needs slice-mode=n-slices)

func (e *Openh264Enc) GetNumSlices() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slices")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openh264Enc) SetNumSlices(value uint) error {
	return e.Element.SetProperty("num-slices", value)
}

// qp-max (uint)
//
// Maximum quantizer

func (e *Openh264Enc) GetQpMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openh264Enc) SetQpMax(value uint) error {
	return e.Element.SetProperty("qp-max", value)
}

// qp-min (uint)
//
// Minimum quantizer

func (e *Openh264Enc) GetQpMin() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openh264Enc) SetQpMin(value uint) error {
	return e.Element.SetProperty("qp-min", value)
}

// rate-control (GstRcModes)
//
// Rate control mode

func (e *Openh264Enc) GetRateControl() (interface{}, error) {
	return e.Element.GetProperty("rate-control")
}

func (e *Openh264Enc) SetRateControl(value interface{}) error {
	return e.Element.SetProperty("rate-control", value)
}

// scene-change-detection (bool)
//
// Scene change detection

func (e *Openh264Enc) GetSceneChangeDetection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scene-change-detection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Openh264Enc) SetSceneChangeDetection(value bool) error {
	return e.Element.SetProperty("scene-change-detection", value)
}

// slice-mode (GstOpenh264EncSliceModes)
//
// Slice mode

func (e *Openh264Enc) GetSliceMode() (GstOpenh264EncSliceModes, error) {
	var value GstOpenh264EncSliceModes
	var ok bool
	v, err := e.Element.GetProperty("slice-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenh264EncSliceModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenh264EncSliceModes")
	}
	return value, nil
}

func (e *Openh264Enc) SetSliceMode(value GstOpenh264EncSliceModes) error {
	e.Element.SetArg("slice-mode", string(value))
	return nil
}

// usage-type (GstEusageType)
//
// Type of video content

func (e *Openh264Enc) GetUsageType() (interface{}, error) {
	return e.Element.GetProperty("usage-type")
}

func (e *Openh264Enc) SetUsageType(value interface{}) error {
	return e.Element.SetProperty("usage-type", value)
}

// ----- Constants -----

type EusageType string

const (
	EusageTypeCamera EusageType = "camera" // camera (0) – video from camera
	EusageTypeScreen EusageType = "screen" // screen (1) – screen content
)

type GstOpenh264EncSliceModes string

const (
	GstOpenh264EncSliceModesNSlices GstOpenh264EncSliceModes = "n-slices" // n-slices (1) – Fixed number of slices
	GstOpenh264EncSliceModesAuto GstOpenh264EncSliceModes = "auto" // auto (5) – Number of slices equal to number of threads
)

type GstOpenh264EncComplexity string

const (
	GstOpenh264EncComplexityLow GstOpenh264EncComplexity = "low" // low (0) – Low complexity / high speed encoding
	GstOpenh264EncComplexityMedium GstOpenh264EncComplexity = "medium" // medium (1) – Medium complexity / medium speed encoding
	GstOpenh264EncComplexityHigh GstOpenh264EncComplexity = "high" // high (2) – High complexity / low speed encoding
)

type GstOpenh264EncDeblockingModes string

const (
	GstOpenh264EncDeblockingModesOn GstOpenh264EncDeblockingModes = "on" // on (0) – Deblocking on
	GstOpenh264EncDeblockingModesOff GstOpenh264EncDeblockingModes = "off" // off (1) – Deblocking off
	GstOpenh264EncDeblockingModesNotSliceBoundaries GstOpenh264EncDeblockingModes = "not-slice-boundaries" // not-slice-boundaries (2) – Deblocking on, except for slice boundaries
)

type RcModes string

const (
	RcModesQuality RcModes = "quality" // quality (0) – Quality mode
	RcModesBitrate RcModes = "bitrate" // bitrate (1) – Bitrate mode
	RcModesBuffer RcModes = "buffer" // buffer (2) – No bitrate control, just using buffer status
	RcModesOff RcModes = "off" // off (-1) – Rate control off mode
)

