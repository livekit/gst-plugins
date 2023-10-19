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

type X264Enc struct {
	Element *gst.Element
}

func NewX264Enc() (*X264Enc, error) {
	e, err := gst.NewElement("x264enc")
	if err != nil {
		return nil, err
	}
	return &X264Enc{Element: e}, nil
}

func NewX264EncWithName(name string) (*X264Enc, error) {
	e, err := gst.NewElementWithName("x264enc", name)
	if err != nil {
		return nil, err
	}
	return &X264Enc{Element: e}, nil
}

// ----- Properties -----

// analyse (GstX264EncAnalyse)
//
// Partitions to consider

func (e *X264Enc) GetAnalyse() (GstX264EncAnalyse, error) {
	var value GstX264EncAnalyse
	var ok bool
	v, err := e.Element.GetProperty("analyse")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncAnalyse)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncAnalyse")
	}
	return value, nil
}

func (e *X264Enc) SetAnalyse(value GstX264EncAnalyse) error {
	e.Element.SetArg("analyse", string(value))
	return nil
}

// aud (bool)
//
// Use AU (Access Unit) delimiter

func (e *X264Enc) GetAud() (bool, error) {
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

func (e *X264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

// b-adapt (bool)
//
// Automatically decide how many B-frames to use

func (e *X264Enc) GetBAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

// b-pyramid (bool)
//
// Keep some B-frames as references

func (e *X264Enc) GetBPyramid() (bool, error) {
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

func (e *X264Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

// bframes (uint)
//
// Number of B-frames between I and P

func (e *X264Enc) GetBframes() (uint, error) {
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

func (e *X264Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

// bitrate (uint)
//
// Bitrate in kbit/sec

func (e *X264Enc) GetBitrate() (uint, error) {
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

func (e *X264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// byte-stream (bool)
//
// Generate byte stream format of NALU

func (e *X264Enc) GetByteStream() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("byte-stream")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetByteStream(value bool) error {
	return e.Element.SetProperty("byte-stream", value)
}

// cabac (bool)
//
// Enable CABAC entropy coding

func (e *X264Enc) GetCabac() (bool, error) {
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

func (e *X264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

// dct8x8 (bool)
//
// Adaptive spatial transform size

func (e *X264Enc) GetDct8X8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dct8x8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetDct8X8(value bool) error {
	return e.Element.SetProperty("dct8x8", value)
}

// frame-packing (GstX264EncFramePacking)
//
// Set frame packing mode for Stereoscopic content

func (e *X264Enc) GetFramePacking() (GstX264EncFramePacking, error) {
	var value GstX264EncFramePacking
	var ok bool
	v, err := e.Element.GetProperty("frame-packing")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncFramePacking)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncFramePacking")
	}
	return value, nil
}

func (e *X264Enc) SetFramePacking(value GstX264EncFramePacking) error {
	e.Element.SetArg("frame-packing", string(value))
	return nil
}

// insert-vui (bool)
//
// Insert VUI NAL in stream

func (e *X264Enc) GetInsertVui() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("insert-vui")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetInsertVui(value bool) error {
	return e.Element.SetProperty("insert-vui", value)
}

// interlaced (bool)
//
// Interlaced material

func (e *X264Enc) GetInterlaced() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("interlaced")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetInterlaced(value bool) error {
	return e.Element.SetProperty("interlaced", value)
}

// intra-refresh (bool)
//
// Use Periodic Intra Refresh instead of IDR frames

func (e *X264Enc) GetIntraRefresh() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("intra-refresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetIntraRefresh(value bool) error {
	return e.Element.SetProperty("intra-refresh", value)
}

// ip-factor (float32)
//
// Quantizer factor between I- and P-frames

func (e *X264Enc) GetIpFactor() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ip-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *X264Enc) SetIpFactor(value float32) error {
	return e.Element.SetProperty("ip-factor", value)
}

// key-int-max (uint)
//
// Maximal distance between two key-frames (0 for automatic)

func (e *X264Enc) GetKeyIntMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("key-int-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetKeyIntMax(value uint) error {
	return e.Element.SetProperty("key-int-max", value)
}

// mb-tree (bool)
//
// Macroblock-Tree ratecontrol

func (e *X264Enc) GetMbTree() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mb-tree")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetMbTree(value bool) error {
	return e.Element.SetProperty("mb-tree", value)
}

// me (GstX264EncMe)
//
// Integer pixel motion estimation method

func (e *X264Enc) GetMe() (GstX264EncMe, error) {
	var value GstX264EncMe
	var ok bool
	v, err := e.Element.GetProperty("me")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncMe)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncMe")
	}
	return value, nil
}

func (e *X264Enc) SetMe(value GstX264EncMe) error {
	e.Element.SetArg("me", string(value))
	return nil
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *X264Enc) GetMultipassCacheFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("multipass-cache-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *X264Enc) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// noise-reduction (uint)
//
// Noise reduction strength

func (e *X264Enc) GetNoiseReduction() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("noise-reduction")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetNoiseReduction(value uint) error {
	return e.Element.SetProperty("noise-reduction", value)
}

// option-string (string)
//
// String of x264 options (overridden by element properties) in the format "key1=value1:key2=value2".

func (e *X264Enc) GetOptionString() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("option-string")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *X264Enc) SetOptionString(value string) error {
	return e.Element.SetProperty("option-string", value)
}

// pass (GstX264EncPass)
//
// Encoding pass/type

func (e *X264Enc) GetPass() (GstX264EncPass, error) {
	var value GstX264EncPass
	var ok bool
	v, err := e.Element.GetProperty("pass")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncPass)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncPass")
	}
	return value, nil
}

func (e *X264Enc) SetPass(value GstX264EncPass) error {
	e.Element.SetArg("pass", string(value))
	return nil
}

// pb-factor (float32)
//
// Quantizer factor between P- and B-frames

func (e *X264Enc) GetPbFactor() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pb-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *X264Enc) SetPbFactor(value float32) error {
	return e.Element.SetProperty("pb-factor", value)
}

// psy-tune (GstX264EncPsyTune)
//
// Preset name for psychovisual tuning options

func (e *X264Enc) GetPsyTune() (GstX264EncPsyTune, error) {
	var value GstX264EncPsyTune
	var ok bool
	v, err := e.Element.GetProperty("psy-tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncPsyTune)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncPsyTune")
	}
	return value, nil
}

func (e *X264Enc) SetPsyTune(value GstX264EncPsyTune) error {
	e.Element.SetArg("psy-tune", string(value))
	return nil
}

// qp-max (uint)
//
// Maximum quantizer

func (e *X264Enc) GetQpMax() (uint, error) {
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

func (e *X264Enc) SetQpMax(value uint) error {
	return e.Element.SetProperty("qp-max", value)
}

// qp-min (uint)
//
// Minimum quantizer

func (e *X264Enc) GetQpMin() (uint, error) {
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

func (e *X264Enc) SetQpMin(value uint) error {
	return e.Element.SetProperty("qp-min", value)
}

// qp-step (uint)
//
// Maximum quantizer difference between frames

func (e *X264Enc) GetQpStep() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-step")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetQpStep(value uint) error {
	return e.Element.SetProperty("qp-step", value)
}

// quantizer (uint)
//
// Constant quantizer or quality to apply

func (e *X264Enc) GetQuantizer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetQuantizer(value uint) error {
	return e.Element.SetProperty("quantizer", value)
}

// rc-lookahead (int)
//
// Number of frames for frametype lookahead

func (e *X264Enc) GetRcLookahead() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *X264Enc) SetRcLookahead(value int) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

// ref (uint)
//
// Number of reference frames

func (e *X264Enc) GetRef() (uint, error) {
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

func (e *X264Enc) SetRef(value uint) error {
	return e.Element.SetProperty("ref", value)
}

// sliced-threads (bool)
//
// Low latency but lower efficiency threading

func (e *X264Enc) GetSlicedThreads() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sliced-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetSlicedThreads(value bool) error {
	return e.Element.SetProperty("sliced-threads", value)
}

// speed-preset (GstX264EncPreset)
//
// Preset name for speed/quality tradeoff options (can affect decode compatibility - impose restrictions separately for your target decoder)

func (e *X264Enc) GetSpeedPreset() (GstX264EncPreset, error) {
	var value GstX264EncPreset
	var ok bool
	v, err := e.Element.GetProperty("speed-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncPreset")
	}
	return value, nil
}

func (e *X264Enc) SetSpeedPreset(value GstX264EncPreset) error {
	e.Element.SetArg("speed-preset", string(value))
	return nil
}

// sps-id (uint)
//
// SPS and PPS ID number

func (e *X264Enc) GetSpsId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sps-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetSpsId(value uint) error {
	return e.Element.SetProperty("sps-id", value)
}

// subme (uint)
//
// Subpixel motion estimation and partition decision quality: 1=fast, 10=best

func (e *X264Enc) GetSubme() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("subme")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetSubme(value uint) error {
	return e.Element.SetProperty("subme", value)
}

// sync-lookahead (int)
//
// Number of buffer frames for threaded lookahead (-1 for automatic)

func (e *X264Enc) GetSyncLookahead() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sync-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *X264Enc) SetSyncLookahead(value int) error {
	return e.Element.SetProperty("sync-lookahead", value)
}

// threads (uint)
//
// Number of threads used by the codec (0 for automatic)

func (e *X264Enc) GetThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetThreads(value uint) error {
	return e.Element.SetProperty("threads", value)
}

// trellis (bool)
//
// Enable trellis searched quantization

func (e *X264Enc) GetTrellis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetTrellis(value bool) error {
	return e.Element.SetProperty("trellis", value)
}

// tune (GstX264EncTune)
//
// Preset name for non-psychovisual tuning options

func (e *X264Enc) GetTune() (GstX264EncTune, error) {
	var value GstX264EncTune
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX264EncTune)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX264EncTune")
	}
	return value, nil
}

func (e *X264Enc) SetTune(value GstX264EncTune) error {
	e.Element.SetArg("tune", string(value))
	return nil
}

// vbv-buf-capacity (uint)
//
// Size of the VBV buffer in milliseconds

func (e *X264Enc) GetVbvBufCapacity() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buf-capacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *X264Enc) SetVbvBufCapacity(value uint) error {
	return e.Element.SetProperty("vbv-buf-capacity", value)
}

// weightb (bool)
//
// Weighted prediction for B-frames

func (e *X264Enc) GetWeightb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weightb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *X264Enc) SetWeightb(value bool) error {
	return e.Element.SetProperty("weightb", value)
}

// ----- Constants -----

type GstX264EncAnalyse string

const (
	GstX264EncAnalyseI4X4 GstX264EncAnalyse = "i4x4" // i4x4 (0x00000001) – i4x4
	GstX264EncAnalyseI8X8 GstX264EncAnalyse = "i8x8" // i8x8 (0x00000002) – i8x8
	GstX264EncAnalyseP8X8 GstX264EncAnalyse = "p8x8" // p8x8 (0x00000010) – p8x8
	GstX264EncAnalyseP4X4 GstX264EncAnalyse = "p4x4" // p4x4 (0x00000020) – p4x4
	GstX264EncAnalyseB8X8 GstX264EncAnalyse = "b8x8" // b8x8 (0x00000100) – b8x8
)

type GstX264EncFramePacking string

const (
	GstX264EncFramePackingAuto GstX264EncFramePacking = "auto" // auto (-1) – Automatic (use incoming video information)
	GstX264EncFramePackingCheckerboard GstX264EncFramePacking = "checkerboard" // checkerboard (0) – checkerboard - Left and Right pixels alternate in a checkerboard pattern
	GstX264EncFramePackingColumnInterleaved GstX264EncFramePacking = "column-interleaved" // column-interleaved (1) – column interleaved - Alternating pixel columns represent Left and Right views
	GstX264EncFramePackingRowInterleaved GstX264EncFramePacking = "row-interleaved" // row-interleaved (2) – row interleaved - Alternating pixel rows represent Left and Right views
	GstX264EncFramePackingSideBySide GstX264EncFramePacking = "side-by-side" // side-by-side (3) – side by side - The left half of the frame contains the Left eye view, the right half the Right eye view
	GstX264EncFramePackingTopBottom GstX264EncFramePacking = "top-bottom" // top-bottom (4) – top bottom - L is on top, R on bottom
	GstX264EncFramePackingFrameInterleaved GstX264EncFramePacking = "frame-interleaved" // frame-interleaved (5) – frame interleaved - Each frame contains either Left or Right view alternately
)

type GstX264EncMe string

const (
	GstX264EncMeDia GstX264EncMe = "dia" // dia (0) – dia
	GstX264EncMeHex GstX264EncMe = "hex" // hex (1) – hex
	GstX264EncMeUmh GstX264EncMe = "umh" // umh (2) – umh
	GstX264EncMeEsa GstX264EncMe = "esa" // esa (3) – esa
	GstX264EncMeTesa GstX264EncMe = "tesa" // tesa (4) – tesa
)

type GstX264EncPass string

const (
	GstX264EncPassCbr GstX264EncPass = "cbr" // cbr (0) – Constant Bitrate Encoding
	GstX264EncPassQuant GstX264EncPass = "quant" // quant (4) – Constant Quantizer
	GstX264EncPassQual GstX264EncPass = "qual" // qual (5) – Constant Quality
	GstX264EncPassPass1 GstX264EncPass = "pass1" // pass1 (17) – VBR Encoding - Pass 1
	GstX264EncPassPass2 GstX264EncPass = "pass2" // pass2 (18) – VBR Encoding - Pass 2
	GstX264EncPassPass3 GstX264EncPass = "pass3" // pass3 (19) – VBR Encoding - Pass 3
)

type GstX264EncPreset string

const (
	GstX264EncPresetNone GstX264EncPreset = "None" // None (0) – No preset
	GstX264EncPresetUltrafast GstX264EncPreset = "ultrafast" // ultrafast (1) – ultrafast
	GstX264EncPresetSuperfast GstX264EncPreset = "superfast" // superfast (2) – superfast
	GstX264EncPresetVeryfast GstX264EncPreset = "veryfast" // veryfast (3) – veryfast
	GstX264EncPresetFaster GstX264EncPreset = "faster" // faster (4) – faster
	GstX264EncPresetFast GstX264EncPreset = "fast" // fast (5) – fast
	GstX264EncPresetMedium GstX264EncPreset = "medium" // medium (6) – medium
	GstX264EncPresetSlow GstX264EncPreset = "slow" // slow (7) – slow
	GstX264EncPresetSlower GstX264EncPreset = "slower" // slower (8) – slower
	GstX264EncPresetVeryslow GstX264EncPreset = "veryslow" // veryslow (9) – veryslow
	GstX264EncPresetPlacebo GstX264EncPreset = "placebo" // placebo (10) – placebo
)

type GstX264EncPsyTune string

const (
	GstX264EncPsyTuneNone GstX264EncPsyTune = "none" // none (0) – No tuning
	GstX264EncPsyTuneFilm GstX264EncPsyTune = "film" // film (1) – Film
	GstX264EncPsyTuneAnimation GstX264EncPsyTune = "animation" // animation (2) – Animation
	GstX264EncPsyTuneGrain GstX264EncPsyTune = "grain" // grain (3) – Grain
	GstX264EncPsyTunePsnr GstX264EncPsyTune = "psnr" // psnr (4) – PSNR
	GstX264EncPsyTuneSsim GstX264EncPsyTune = "ssim" // ssim (5) – SSIM
)

type GstX264EncTune string

const (
	GstX264EncTuneStillimage GstX264EncTune = "stillimage" // stillimage (0x00000001) – Still image
	GstX264EncTuneFastdecode GstX264EncTune = "fastdecode" // fastdecode (0x00000002) – Fast decode
	GstX264EncTuneZerolatency GstX264EncTune = "zerolatency" // zerolatency (0x00000004) – Zero latency
)

