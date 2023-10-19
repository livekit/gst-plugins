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

type Uvch264Src struct {
	Element *gst.Element
}

func NewUvch264Src() (*Uvch264Src, error) {
	e, err := gst.NewElement("uvch264src")
	if err != nil {
		return nil, err
	}
	return &Uvch264Src{Element: e}, nil
}

func NewUvch264SrcWithName(name string) (*Uvch264Src, error) {
	e, err := gst.NewElementWithName("uvch264src", name)
	if err != nil {
		return nil, err
	}
	return &Uvch264Src{Element: e}, nil
}

// ----- Properties -----

// average-bitrate (uint)
//
// The average bitrate in bits/second (dynamic control)

func (e *Uvch264Src) GetAverageBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("average-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetAverageBitrate(value uint) error {
	return e.Element.SetProperty("average-bitrate", value)
}

// colorspace-name (string)
//
// The name of the colorspace element

func (e *Uvch264Src) GetColorspaceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("colorspace-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uvch264Src) SetColorspaceName(value string) error {
	return e.Element.SetProperty("colorspace-name", value)
}

// device (string)
//
// Device location

func (e *Uvch264Src) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uvch264Src) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-name (string)
//
// Name of the device

func (e *Uvch264Src) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// enable-sei (bool)
//
// Enable SEI picture timing (static control)

func (e *Uvch264Src) GetEnableSei() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-sei")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Uvch264Src) SetEnableSei(value bool) error {
	return e.Element.SetProperty("enable-sei", value)
}

// entropy (GstUvcH264Entropy)
//
// Entropy (static control)

func (e *Uvch264Src) GetEntropy() (interface{}, error) {
	return e.Element.GetProperty("entropy")
}

func (e *Uvch264Src) SetEntropy(value interface{}) error {
	return e.Element.SetProperty("entropy", value)
}

// fixed-framerate (bool)
//
// Fixed framerate (static & dynamic control)

func (e *Uvch264Src) GetFixedFramerate() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fixed-framerate")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Uvch264Src) SetFixedFramerate(value bool) error {
	return e.Element.SetProperty("fixed-framerate", value)
}

// iframe-period (uint)
//
// Time between IDR frames in milliseconds (static control)

func (e *Uvch264Src) GetIframePeriod() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("iframe-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetIframePeriod(value uint) error {
	return e.Element.SetProperty("iframe-period", value)
}

// initial-bitrate (uint)
//
// Initial bitrate in bits/second (static control)

func (e *Uvch264Src) GetInitialBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("initial-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetInitialBitrate(value uint) error {
	return e.Element.SetProperty("initial-bitrate", value)
}

// jpeg-decoder-name (string)
//
// The name of the jpeg decoder element

func (e *Uvch264Src) GetJpegDecoderName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("jpeg-decoder-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uvch264Src) SetJpegDecoderName(value string) error {
	return e.Element.SetProperty("jpeg-decoder-name", value)
}

// leaky-bucket-size (uint)
//
// Size of the leaky bucket size in milliseconds (static control)

func (e *Uvch264Src) GetLeakyBucketSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("leaky-bucket-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetLeakyBucketSize(value uint) error {
	return e.Element.SetProperty("leaky-bucket-size", value)
}

// level-idc (uint)
//
// Level IDC (dynamic control)

func (e *Uvch264Src) GetLevelIdc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("level-idc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetLevelIdc(value uint) error {
	return e.Element.SetProperty("level-idc", value)
}

// ltr-buffer-size (int)
//
// Total number of Long-Term Reference frames (dynamic control)

func (e *Uvch264Src) GetLtrBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ltr-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetLtrBufferSize(value int) error {
	return e.Element.SetProperty("ltr-buffer-size", value)
}

// ltr-encoder-control (int)
//
// Number of LTR frames the device can control (dynamic control)

func (e *Uvch264Src) GetLtrEncoderControl() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ltr-encoder-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetLtrEncoderControl(value int) error {
	return e.Element.SetProperty("ltr-encoder-control", value)
}

// max-bframe-qp (int)
//
// The minimum Quantization step size for B frames (dynamic control)

func (e *Uvch264Src) GetMaxBframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetMaxBframeQp(value int) error {
	return e.Element.SetProperty("max-bframe-qp", value)
}

// max-iframe-qp (int)
//
// The minimum Quantization step size for I frames (dynamic control)

func (e *Uvch264Src) GetMaxIframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-iframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetMaxIframeQp(value int) error {
	return e.Element.SetProperty("max-iframe-qp", value)
}

// max-mbps (uint)
//
// The number of macroblocks per second for the maximum processing rate

func (e *Uvch264Src) GetMaxMbps() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-mbps")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// max-pframe-qp (int)
//
// The minimum Quantization step size for P frames (dynamic control)

func (e *Uvch264Src) GetMaxPframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-pframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetMaxPframeQp(value int) error {
	return e.Element.SetProperty("max-pframe-qp", value)
}

// min-bframe-qp (int)
//
// The minimum Quantization step size for B frames (dynamic control)

func (e *Uvch264Src) GetMinBframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-bframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetMinBframeQp(value int) error {
	return e.Element.SetProperty("min-bframe-qp", value)
}

// min-iframe-qp (int)
//
// The minimum Quantization step size for I frames (dynamic control)

func (e *Uvch264Src) GetMinIframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-iframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetMinIframeQp(value int) error {
	return e.Element.SetProperty("min-iframe-qp", value)
}

// min-pframe-qp (int)
//
// The minimum Quantization step size for P frames (dynamic control)

func (e *Uvch264Src) GetMinPframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-pframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetMinPframeQp(value int) error {
	return e.Element.SetProperty("min-pframe-qp", value)
}

// num-buffers (int)
//
// Number of buffers to output before sending EOS (-1 = unlimited)

func (e *Uvch264Src) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// num-clock-samples (int)
//
// Number of clock samples to gather for the PTS synchronization (-1 = unlimited)

func (e *Uvch264Src) GetNumClockSamples() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-clock-samples")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uvch264Src) SetNumClockSamples(value int) error {
	return e.Element.SetProperty("num-clock-samples", value)
}

// num-reorder-frames (uint)
//
// Number of B frames between the references frames (static control)

func (e *Uvch264Src) GetNumReorderFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-reorder-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetNumReorderFrames(value uint) error {
	return e.Element.SetProperty("num-reorder-frames", value)
}

// peak-bitrate (uint)
//
// The peak bitrate in bits/second (dynamic control)

func (e *Uvch264Src) GetPeakBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("peak-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetPeakBitrate(value uint) error {
	return e.Element.SetProperty("peak-bitrate", value)
}

// preview-flipped (bool)
//
// Horizontal flipped image for non H.264 streams (static control)

func (e *Uvch264Src) GetPreviewFlipped() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preview-flipped")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Uvch264Src) SetPreviewFlipped(value bool) error {
	return e.Element.SetProperty("preview-flipped", value)
}

// rate-control (GstUvcH264RateControl)
//
// Rate control mode (static & dynamic control)

func (e *Uvch264Src) GetRateControl() (interface{}, error) {
	return e.Element.GetProperty("rate-control")
}

func (e *Uvch264Src) SetRateControl(value interface{}) error {
	return e.Element.SetProperty("rate-control", value)
}

// slice-mode (GstUvcH264SliceMode)
//
// Defines the unit of the slice-units property (static control)

func (e *Uvch264Src) GetSliceMode() (interface{}, error) {
	return e.Element.GetProperty("slice-mode")
}

func (e *Uvch264Src) SetSliceMode(value interface{}) error {
	return e.Element.SetProperty("slice-mode", value)
}

// slice-units (uint)
//
// Slice units (static control)

func (e *Uvch264Src) GetSliceUnits() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("slice-units")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Uvch264Src) SetSliceUnits(value uint) error {
	return e.Element.SetProperty("slice-units", value)
}

// usage-type (GstUvcH264UsageType)
//
// The usage type (static control)

func (e *Uvch264Src) GetUsageType() (interface{}, error) {
	return e.Element.GetProperty("usage-type")
}

func (e *Uvch264Src) SetUsageType(value interface{}) error {
	return e.Element.SetProperty("usage-type", value)
}

// ----- Constants -----

type UvcH264SliceMode string

const (
	UvcH264SliceModeIgnored UvcH264SliceMode = "ignored" // ignored (0) – Ignored
	UvcH264SliceModeBitsslice UvcH264SliceMode = "bits/slice" // bits/slice (1) – Bits per slice
	UvcH264SliceModeMbsslice UvcH264SliceMode = "MBs/slice" // MBs/slice (2) – MBs per Slice
	UvcH264SliceModeSliceframe UvcH264SliceMode = "slice/frame" // slice/frame (3) – Slice Per Frame
)

type UvcH264UsageType string

const (
	UvcH264UsageTypeRealtime UvcH264UsageType = "realtime" // realtime (1) – Realtime (video conferencing)
	UvcH264UsageTypeBroadcast UvcH264UsageType = "broadcast" // broadcast (2) – Broadcast
	UvcH264UsageTypeStorage UvcH264UsageType = "storage" // storage (3) – Storage
	UvcH264UsageTypeUcconfig0 UvcH264UsageType = "ucconfig0" // ucconfig0 (4) – UCConfig 0
	UvcH264UsageTypeUcconfig1 UvcH264UsageType = "ucconfig1" // ucconfig1 (5) – UCConfig 1
	UvcH264UsageTypeUcconfig2Q UvcH264UsageType = "ucconfig2q" // ucconfig2q (6) – UCConfig 2Q
	UvcH264UsageTypeUcconfig2S UvcH264UsageType = "ucconfig2s" // ucconfig2s (7) – UCConfig 2S
	UvcH264UsageTypeUcconfig3 UvcH264UsageType = "ucconfig3" // ucconfig3 (8) – UCConfig 3
)

type UvcH264Entropy string

const (
	UvcH264EntropyCavlc UvcH264Entropy = "cavlc" // cavlc (0) – CAVLC
	UvcH264EntropyCabac UvcH264Entropy = "cabac" // cabac (1) – CABAC
)

type UvcH264RateControl string

const (
	UvcH264RateControlCbr UvcH264RateControl = "cbr" // cbr (1) – Constant bit rate
	UvcH264RateControlVbr UvcH264RateControl = "vbr" // vbr (2) – Variable bit rate
	UvcH264RateControlQp UvcH264RateControl = "qp" // qp (3) – Constant QP
)

