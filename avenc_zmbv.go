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

type AvencZmbv struct {
	Element *gst.Element
}

func NewAvencZmbv() (*AvencZmbv, error) {
	e, err := gst.NewElement("avenc_zmbv")
	if err != nil {
		return nil, err
	}
	return &AvencZmbv{Element: e}, nil
}

func NewAvencZmbvWithName(name string) (*AvencZmbv, error) {
	e, err := gst.NewElementWithName("avenc_zmbv", name)
	if err != nil {
		return nil, err
	}
	return &AvencZmbv{Element: e}, nil
}

// ----- Properties -----

// b-qfactor (float32)
//
// QP factor between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBQfactor() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("b-qfactor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetBQfactor(value float32) error {
	return e.Element.SetProperty("b-qfactor", value)
}

// b-qoffset (float32)
//
// QP offset between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBQoffset() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("b-qoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetBQoffset(value float32) error {
	return e.Element.SetProperty("b-qoffset", value)
}

// b-sensitivity (int)
//
// adjust sensitivity of b_frame_strategy 1 (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBSensitivity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("b-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBSensitivity(value int) error {
	return e.Element.SetProperty("b-sensitivity", value)
}

// b-strategy (int)
//
// strategy to choose between I/P/B-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBStrategy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("b-strategy")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBStrategy(value int) error {
	return e.Element.SetProperty("b-strategy", value)
}

// bidir-refine (int)
//
// refine the two motion vectors used in bidirectional macroblocks (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBidirRefine() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bidir-refine")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBidirRefine(value int) error {
	return e.Element.SetProperty("bidir-refine", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-tolerance (int)
//
// Set video bitrate tolerance (in bits/s). In 1-pass mode, bitrate tolerance specifies how far ratecontrol is willing to deviate from the target average bitrate value. This is not related to minimum/maximum bitrate. Lowering tolerance too much has an adverse effect on quality. (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBitrateTolerance() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBitrateTolerance(value int) error {
	return e.Element.SetProperty("bitrate-tolerance", value)
}

// brd-scale (int)
//
// downscale frames for dynamic B-frame decision (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBrdScale() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brd-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBrdScale(value int) error {
	return e.Element.SetProperty("brd-scale", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetBufsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bufsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// chroma-sample-location (GstAvcodeccontextChromaSampleLocationType)
//
// chroma sample location (Generic codec option, might have no effect)

func (e *AvencZmbv) GetChromaSampleLocation() (interface{}, error) {
	return e.Element.GetProperty("chroma-sample-location")
}

func (e *AvencZmbv) SetChromaSampleLocation(value interface{}) error {
	return e.Element.SetProperty("chroma-sample-location", value)
}

// chromaoffset (int)
//
// chroma QP offset from luma (Generic codec option, might have no effect)

func (e *AvencZmbv) GetChromaoffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("chromaoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetChromaoffset(value int) error {
	return e.Element.SetProperty("chromaoffset", value)
}

// cmp (GstAvcodeccontextCmpFunc)
//
// full-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencZmbv) GetCmp() (interface{}, error) {
	return e.Element.GetProperty("cmp")
}

func (e *AvencZmbv) SetCmp(value interface{}) error {
	return e.Element.SetProperty("cmp", value)
}

// coder (GstAvcodeccontextCoder)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetCoder() (interface{}, error) {
	return e.Element.GetProperty("coder")
}

func (e *AvencZmbv) SetCoder(value interface{}) error {
	return e.Element.SetProperty("coder", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetCompressionLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("compression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// context (int)
//
// context model (Generic codec option, might have no effect)

func (e *AvencZmbv) GetContext() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("context")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetContext(value int) error {
	return e.Element.SetProperty("context", value)
}

// dark-mask (float32)
//
// compresses dark areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencZmbv) GetDarkMask() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dark-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetDarkMask(value float32) error {
	return e.Element.SetProperty("dark-mask", value)
}

// dc (int)
//
// intra_dc_precision (Generic codec option, might have no effect)

func (e *AvencZmbv) GetDc() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("dc")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetDc(value int) error {
	return e.Element.SetProperty("dc", value)
}

// dct (GstAvcodeccontextDct)
//
// DCT algorithm (Generic codec option, might have no effect)

func (e *AvencZmbv) GetDct() (interface{}, error) {
	return e.Element.GetProperty("dct")
}

func (e *AvencZmbv) SetDct(value interface{}) error {
	return e.Element.SetProperty("dct", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencZmbv) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencZmbv) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dia-size (int)
//
// diamond type & size for motion estimation (Generic codec option, might have no effect)

func (e *AvencZmbv) GetDiaSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("dia-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetDiaSize(value int) error {
	return e.Element.SetProperty("dia-size", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencZmbv) GetDumpSeparator() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dump-separator")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *AvencZmbv) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencZmbv) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencZmbv) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencZmbv) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencZmbv) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// field-order (GstAvcodeccontextFieldOrder)
//
// Field order (Generic codec option, might have no effect)

func (e *AvencZmbv) GetFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("field-order")
}

func (e *AvencZmbv) SetFieldOrder(value interface{}) error {
	return e.Element.SetProperty("field-order", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencZmbv) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencZmbv) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetGlobalQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("global-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// gop-size (int)
//
// set the group of picture (GOP) size (Generic codec option, might have no effect)

func (e *AvencZmbv) GetGopSize() (int, error) {
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

func (e *AvencZmbv) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// i-qfactor (float32)
//
// QP factor between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetIQfactor() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("i-qfactor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetIQfactor(value float32) error {
	return e.Element.SetProperty("i-qfactor", value)
}

// i-qoffset (float32)
//
// QP offset between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetIQoffset() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("i-qoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetIQoffset(value float32) error {
	return e.Element.SetProperty("i-qoffset", value)
}

// idct (GstAvcodeccontextIdct)
//
// select IDCT implementation (Generic codec option, might have no effect)

func (e *AvencZmbv) GetIdct() (interface{}, error) {
	return e.Element.GetProperty("idct")
}

func (e *AvencZmbv) SetIdct(value interface{}) error {
	return e.Element.SetProperty("idct", value)
}

// ildctcmp (GstAvcodeccontextCmpFunc)
//
// interlaced DCT compare function (Generic codec option, might have no effect)

func (e *AvencZmbv) GetIldctcmp() (interface{}, error) {
	return e.Element.GetProperty("ildctcmp")
}

func (e *AvencZmbv) SetIldctcmp(value interface{}) error {
	return e.Element.SetProperty("ildctcmp", value)
}

// keyint-min (int)
//
// minimum interval between IDR-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetKeyintMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyint-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetKeyintMin(value int) error {
	return e.Element.SetProperty("keyint-min", value)
}

// last-pred (int)
//
// amount of motion predictors from the previous frame (Generic codec option, might have no effect)

func (e *AvencZmbv) GetLastPred() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("last-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetLastPred(value int) error {
	return e.Element.SetProperty("last-pred", value)
}

// lumi-mask (float32)
//
// compresses bright areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencZmbv) GetLumiMask() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lumi-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetLumiMask(value float32) error {
	return e.Element.SetProperty("lumi-mask", value)
}

// max-bframes (int)
//
// set maximum number of B-frames between non-B-frames (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMaxBframes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMaxBframes(value int) error {
	return e.Element.SetProperty("max-bframes", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMaxPixels() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-pixels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencZmbv) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMaxrate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("maxrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencZmbv) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// mbcmp (GstAvcodeccontextCmpFunc)
//
// macroblock compare function (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMbcmp() (interface{}, error) {
	return e.Element.GetProperty("mbcmp")
}

func (e *AvencZmbv) SetMbcmp(value interface{}) error {
	return e.Element.SetProperty("mbcmp", value)
}

// mbd (GstAvcodeccontextMbd)
//
// macroblock decision algorithm (high quality mode) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMbd() (interface{}, error) {
	return e.Element.GetProperty("mbd")
}

func (e *AvencZmbv) SetMbd(value interface{}) error {
	return e.Element.SetProperty("mbd", value)
}

// mblmax (int)
//
// maximum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMblmax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mblmax")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMblmax(value int) error {
	return e.Element.SetProperty("mblmax", value)
}

// mblmin (int)
//
// minimum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMblmin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mblmin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMblmin(value int) error {
	return e.Element.SetProperty("mblmin", value)
}

// me-range (int)
//
// limit motion vectors range (1023 for DivX player) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMeRange() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("me-range")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMeRange(value int) error {
	return e.Element.SetProperty("me-range", value)
}

// mepc (int)
//
// motion estimation bitrate penalty compensation (1.0 = 256) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMepc() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mepc")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMepc(value int) error {
	return e.Element.SetProperty("mepc", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMinrate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("minrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencZmbv) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// mpeg-quant (int)
//
// use MPEG quantizers instead of H.263 (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMpegQuant() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mpeg-quant")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMpegQuant(value int) error {
	return e.Element.SetProperty("mpeg-quant", value)
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *AvencZmbv) GetMultipassCacheFile() (string, error) {
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

func (e *AvencZmbv) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// mv0-threshold (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetMv0Threshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mv0-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetMv0Threshold(value int) error {
	return e.Element.SetProperty("mv0-threshold", value)
}

// nr (int)
//
// noise reduction (Generic codec option, might have no effect)

func (e *AvencZmbv) GetNr() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("nr")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetNr(value int) error {
	return e.Element.SetProperty("nr", value)
}

// nssew (int)
//
// nsse weight (Generic codec option, might have no effect)

func (e *AvencZmbv) GetNssew() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("nssew")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetNssew(value int) error {
	return e.Element.SetProperty("nssew", value)
}

// p-mask (float32)
//
// inter masking (Generic codec option, might have no effect)

func (e *AvencZmbv) GetPMask() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetPMask(value float32) error {
	return e.Element.SetProperty("p-mask", value)
}

// pass (GstLibAVEncPass)
//
// Encoding pass/type

func (e *AvencZmbv) GetPass() (interface{}, error) {
	return e.Element.GetProperty("pass")
}

func (e *AvencZmbv) SetPass(value interface{}) error {
	return e.Element.SetProperty("pass", value)
}

// pre-dia-size (int)
//
// diamond type & size for motion estimation pre-pass (Generic codec option, might have no effect)

func (e *AvencZmbv) GetPreDiaSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pre-dia-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetPreDiaSize(value int) error {
	return e.Element.SetProperty("pre-dia-size", value)
}

// precmp (GstAvcodeccontextCmpFunc)
//
// pre motion estimation compare function (Generic codec option, might have no effect)

func (e *AvencZmbv) GetPrecmp() (interface{}, error) {
	return e.Element.GetProperty("precmp")
}

func (e *AvencZmbv) SetPrecmp(value interface{}) error {
	return e.Element.SetProperty("precmp", value)
}

// pred (GstAvcodeccontextPred)
//
// prediction method (Generic codec option, might have no effect)

func (e *AvencZmbv) GetPred() (interface{}, error) {
	return e.Element.GetProperty("pred")
}

func (e *AvencZmbv) SetPred(value interface{}) error {
	return e.Element.SetProperty("pred", value)
}

// preme (int)
//
// pre motion estimation (Generic codec option, might have no effect)

func (e *AvencZmbv) GetPreme() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preme")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetPreme(value int) error {
	return e.Element.SetProperty("preme", value)
}

// ps (int)
//
// RTP payload size in bytes (Generic codec option, might have no effect)

func (e *AvencZmbv) GetPs() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetPs(value int) error {
	return e.Element.SetProperty("ps", value)
}

// qblur (float32)
//
// video quantizer scale blur (VBR) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetQblur() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qblur")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetQblur(value float32) error {
	return e.Element.SetProperty("qblur", value)
}

// qcomp (float32)
//
// video quantizer scale compression (VBR). Constant of ratecontrol equation. Recommended range for default rc_eq: 0.0-1.0 (Generic codec option, might have no effect)

func (e *AvencZmbv) GetQcomp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qcomp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetQcomp(value float32) error {
	return e.Element.SetProperty("qcomp", value)
}

// qdiff (int)
//
// maximum difference between the quantizer scales (VBR) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetQdiff() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qdiff")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetQdiff(value int) error {
	return e.Element.SetProperty("qdiff", value)
}

// qmax (int)
//
// maximum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetQmax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qmax")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetQmax(value int) error {
	return e.Element.SetProperty("qmax", value)
}

// qmin (int)
//
// minimum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetQmin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qmin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetQmin(value int) error {
	return e.Element.SetProperty("qmin", value)
}

// quantizer (float32)
//
// Constant Quantizer

func (e *AvencZmbv) GetQuantizer() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetQuantizer(value float32) error {
	return e.Element.SetProperty("quantizer", value)
}

// rc-init-occupancy (int)
//
// number of bits which should be loaded into the rc buffer before decoding starts (Generic codec option, might have no effect)

func (e *AvencZmbv) GetRcInitOccupancy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rc-init-occupancy")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetRcInitOccupancy(value int) error {
	return e.Element.SetProperty("rc-init-occupancy", value)
}

// rc-max-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetRcMaxVbvUse() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rc-max-vbv-use")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetRcMaxVbvUse(value float32) error {
	return e.Element.SetProperty("rc-max-vbv-use", value)
}

// rc-min-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetRcMinVbvUse() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rc-min-vbv-use")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetRcMinVbvUse(value float32) error {
	return e.Element.SetProperty("rc-min-vbv-use", value)
}

// refs (int)
//
// reference frames to consider for motion compensation (Generic codec option, might have no effect)

func (e *AvencZmbv) GetRefs() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("refs")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetRefs(value int) error {
	return e.Element.SetProperty("refs", value)
}

// sc-threshold (int)
//
// scene change threshold (Generic codec option, might have no effect)

func (e *AvencZmbv) GetScThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sc-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetScThreshold(value int) error {
	return e.Element.SetProperty("sc-threshold", value)
}

// scplx-mask (float32)
//
// spatial complexity masking (Generic codec option, might have no effect)

func (e *AvencZmbv) GetScplxMask() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scplx-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetScplxMask(value float32) error {
	return e.Element.SetProperty("scplx-mask", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSideDataOnlyPackets() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("side-data-only-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencZmbv) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// skip-exp (int)
//
// frame skip exponent (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSkipExp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("skip-exp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetSkipExp(value int) error {
	return e.Element.SetProperty("skip-exp", value)
}

// skip-factor (int)
//
// frame skip factor (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSkipFactor() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("skip-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetSkipFactor(value int) error {
	return e.Element.SetProperty("skip-factor", value)
}

// skip-threshold (int)
//
// frame skip threshold (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSkipThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("skip-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetSkipThreshold(value int) error {
	return e.Element.SetProperty("skip-threshold", value)
}

// skipcmp (GstAvcodeccontextCmpFunc)
//
// frame skip compare function (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSkipcmp() (interface{}, error) {
	return e.Element.GetProperty("skipcmp")
}

func (e *AvencZmbv) SetSkipcmp(value interface{}) error {
	return e.Element.SetProperty("skipcmp", value)
}

// slices (int)
//
// set the number of slices, used in parallelized encoding (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSlices() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("slices")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetSlices(value int) error {
	return e.Element.SetProperty("slices", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencZmbv) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencZmbv) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// subcmp (GstAvcodeccontextCmpFunc)
//
// sub-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSubcmp() (interface{}, error) {
	return e.Element.GetProperty("subcmp")
}

func (e *AvencZmbv) SetSubcmp(value interface{}) error {
	return e.Element.SetProperty("subcmp", value)
}

// subq (int)
//
// sub-pel motion estimation quality (Generic codec option, might have no effect)

func (e *AvencZmbv) GetSubq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("subq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetSubq(value int) error {
	return e.Element.SetProperty("subq", value)
}

// tcplx-mask (float32)
//
// temporal complexity masking (Generic codec option, might have no effect)

func (e *AvencZmbv) GetTcplxMask() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tcplx-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencZmbv) SetTcplxMask(value float32) error {
	return e.Element.SetProperty("tcplx-mask", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencZmbv) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencZmbv) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencZmbv) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencZmbv) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencZmbv) GetTicksPerFrame() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ticks-per-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// timecode-frame-start (int64)
//
// GOP timecode frame start number, in non-drop-frame format (Generic codec option, might have no effect)

func (e *AvencZmbv) GetTimecodeFrameStart() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timecode-frame-start")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencZmbv) SetTimecodeFrameStart(value int64) error {
	return e.Element.SetProperty("timecode-frame-start", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencZmbv) GetTrellis() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencZmbv) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type AvcodeccontextChromaSampleLocationType string

const (
	AvcodeccontextChromaSampleLocationTypeUnknown AvcodeccontextChromaSampleLocationType = "unknown" // unknown (0) – Unspecified
	AvcodeccontextChromaSampleLocationTypeLeft AvcodeccontextChromaSampleLocationType = "left" // left (1) – Left
	AvcodeccontextChromaSampleLocationTypeCenter AvcodeccontextChromaSampleLocationType = "center" // center (2) – Center
	AvcodeccontextChromaSampleLocationTypeTopleft AvcodeccontextChromaSampleLocationType = "topleft" // topleft (3) – Top-left
	AvcodeccontextChromaSampleLocationTypeTop AvcodeccontextChromaSampleLocationType = "top" // top (4) – Top
	AvcodeccontextChromaSampleLocationTypeBottomleft AvcodeccontextChromaSampleLocationType = "bottomleft" // bottomleft (5) – Bottom-left
	AvcodeccontextChromaSampleLocationTypeBottom AvcodeccontextChromaSampleLocationType = "bottom" // bottom (6) – Bottom
)

type AvcodeccontextFieldOrder string

const (
	AvcodeccontextFieldOrderUnknown AvcodeccontextFieldOrder = "unknown" // unknown (0) – Unspecified
	AvcodeccontextFieldOrderProgressive AvcodeccontextFieldOrder = "progressive" // progressive (1) – progressive
	AvcodeccontextFieldOrderTt AvcodeccontextFieldOrder = "tt" // tt (2) – tt
	AvcodeccontextFieldOrderBb AvcodeccontextFieldOrder = "bb" // bb (3) – bb
	AvcodeccontextFieldOrderTb AvcodeccontextFieldOrder = "tb" // tb (4) – tb
	AvcodeccontextFieldOrderBt AvcodeccontextFieldOrder = "bt" // bt (5) – bt
)

type GstLibAvencPass string

const (
	GstLibAvencPassCbr GstLibAvencPass = "cbr" // cbr (0) – Constant Bitrate Encoding
	GstLibAvencPassQuant GstLibAvencPass = "quant" // quant (2) – Constant Quantizer
	GstLibAvencPassPass1 GstLibAvencPass = "pass1" // pass1 (512) – VBR Encoding - Pass 1
	GstLibAvencPassPass2 GstLibAvencPass = "pass2" // pass2 (1024) – VBR Encoding - Pass 2
)

type AvcodeccontextErrDetect string

const (
	AvcodeccontextErrDetectCrccheck AvcodeccontextErrDetect = "crccheck" // crccheck (0x00000001) – verify embedded CRCs
	AvcodeccontextErrDetectBitstream AvcodeccontextErrDetect = "bitstream" // bitstream (0x00000002) – detect bitstream specification deviations
	AvcodeccontextErrDetectBuffer AvcodeccontextErrDetect = "buffer" // buffer (0x00000004) – detect improper bitstream length
	AvcodeccontextErrDetectExplode AvcodeccontextErrDetect = "explode" // explode (0x00000008) – abort decoding on minor error detection
	AvcodeccontextErrDetectIgnoreErr AvcodeccontextErrDetect = "ignore_err" // ignore_err (0x00008000) – ignore errors
	AvcodeccontextErrDetectCareful AvcodeccontextErrDetect = "careful" // careful (0x00010000) – consider things that violate the spec, are fast to check and have not been seen in the wild as errors
	AvcodeccontextErrDetectCompliant AvcodeccontextErrDetect = "compliant" // compliant (0x00030000) – consider all spec non compliancies as errors
	AvcodeccontextErrDetectAggressive AvcodeccontextErrDetect = "aggressive" // aggressive (0x00070000) – consider things that a sane encoder should not do as an error
)

type AvcodeccontextThreadType string

const (
	AvcodeccontextThreadTypeFrame AvcodeccontextThreadType = "frame" // frame (0x00000001) – frame
	AvcodeccontextThreadTypeSlice AvcodeccontextThreadType = "slice" // slice (0x00000002) – slice
)

type AvcodeccontextThreads string

const (
	AvcodeccontextThreadsAuto AvcodeccontextThreads = "auto" // auto (0) – autodetect a suitable number of threads to use
	AvcodeccontextThreadsUnknown AvcodeccontextThreads = "unknown" // unknown (1) – Unspecified
)

type AvcodeccontextCoder string

const (
	AvcodeccontextCoderVlc AvcodeccontextCoder = "vlc" // vlc (0) – variable length coder / Huffman coder
	AvcodeccontextCoderAc AvcodeccontextCoder = "ac" // ac (1) – arithmetic coder
	AvcodeccontextCoderRaw AvcodeccontextCoder = "raw" // raw (2) – raw (no encoding)
	AvcodeccontextCoderRle AvcodeccontextCoder = "rle" // rle (3) – run-length coder
)

type AvcodeccontextDct string

const (
	AvcodeccontextDctAuto AvcodeccontextDct = "auto" // auto (0) – autoselect a good one
	AvcodeccontextDctFastint AvcodeccontextDct = "fastint" // fastint (1) – fast integer
	AvcodeccontextDctInt AvcodeccontextDct = "int" // int (2) – accurate integer
	AvcodeccontextDctMmx AvcodeccontextDct = "mmx" // mmx (3) – mmx
	AvcodeccontextDctAltivec AvcodeccontextDct = "altivec" // altivec (5) – altivec
	AvcodeccontextDctFaan AvcodeccontextDct = "faan" // faan (6) – floating point AAN DCT
)

type AvcodeccontextExportSideData string

const (
	AvcodeccontextExportSideDataMvs AvcodeccontextExportSideData = "mvs" // mvs (0x00000001) – export motion vectors through frame side data
	AvcodeccontextExportSideDataPrft AvcodeccontextExportSideData = "prft" // prft (0x00000002) – export Producer Reference Time through packet side data
	AvcodeccontextExportSideDataVencParams AvcodeccontextExportSideData = "venc_params" // venc_params (0x00000004) – export video encoding parameters through frame side data
	AvcodeccontextExportSideDataFilmGrain AvcodeccontextExportSideData = "film_grain" // film_grain (0x00000008) – export film grain parameters through frame side data
)

type AvcodeccontextFlags string

const (
	AvcodeccontextFlagsUnaligned AvcodeccontextFlags = "unaligned" // unaligned (0x00000001) – allow decoders to produce unaligned output
	AvcodeccontextFlagsMv4 AvcodeccontextFlags = "mv4" // mv4 (0x00000004) – use four motion vectors per macroblock (MPEG-4)
	AvcodeccontextFlagsOutputCorrupt AvcodeccontextFlags = "output_corrupt" // output_corrupt (0x00000008) – Output even potentially corrupted frames
	AvcodeccontextFlagsQpel AvcodeccontextFlags = "qpel" // qpel (0x00000010) – use 1/4-pel motion compensation
	AvcodeccontextFlagsDropChanged AvcodeccontextFlags = "drop_changed" // drop_changed (0x00000020) – Drop frames whose parameters differ from first decoded frame
	AvcodeccontextFlagsLoop AvcodeccontextFlags = "loop" // loop (0x00000800) – use loop filter
	AvcodeccontextFlagsGray AvcodeccontextFlags = "gray" // gray (0x00002000) – only decode/encode grayscale
	AvcodeccontextFlagsPsnr AvcodeccontextFlags = "psnr" // psnr (0x00008000) – error[?] variables will be set during encoding
	AvcodeccontextFlagsTruncated AvcodeccontextFlags = "truncated" // truncated (0x00010000) – Input bitstream might be randomly truncated
	AvcodeccontextFlagsIldct AvcodeccontextFlags = "ildct" // ildct (0x00040000) – use interlaced DCT
	AvcodeccontextFlagsLowDelay AvcodeccontextFlags = "low_delay" // low_delay (0x00080000) – force low delay
	AvcodeccontextFlagsGlobalHeader AvcodeccontextFlags = "global_header" // global_header (0x00400000) – place global headers in extradata instead of every keyframe
	AvcodeccontextFlagsBitexact AvcodeccontextFlags = "bitexact" // bitexact (0x00800000) – use only bitexact functions (except (I)DCT)
	AvcodeccontextFlagsAic AvcodeccontextFlags = "aic" // aic (0x01000000) – H.263 advanced intra coding / MPEG-4 AC prediction
	AvcodeccontextFlagsIlme AvcodeccontextFlags = "ilme" // ilme (0x20000000) – interlaced motion estimation
	AvcodeccontextFlagsCgop AvcodeccontextFlags = "cgop" // cgop (0x80000000) – closed GOP
)

type AvcodeccontextIdct string

const (
	AvcodeccontextIdctAuto AvcodeccontextIdct = "auto" // auto (0) – auto
	AvcodeccontextIdctInt AvcodeccontextIdct = "int" // int (1) – int
	AvcodeccontextIdctSimple AvcodeccontextIdct = "simple" // simple (2) – simple
	AvcodeccontextIdctSimplemmx AvcodeccontextIdct = "simplemmx" // simplemmx (3) – simplemmx
	AvcodeccontextIdctArm AvcodeccontextIdct = "arm" // arm (7) – arm
	AvcodeccontextIdctAltivec AvcodeccontextIdct = "altivec" // altivec (8) – altivec
	AvcodeccontextIdctSimplearm AvcodeccontextIdct = "simplearm" // simplearm (10) – simplearm
	AvcodeccontextIdctXvid AvcodeccontextIdct = "xvid" // xvid (14) – xvid
	AvcodeccontextIdctSimplearmv5Te AvcodeccontextIdct = "simplearmv5te" // simplearmv5te (16) – simplearmv5te
	AvcodeccontextIdctSimplearmv6 AvcodeccontextIdct = "simplearmv6" // simplearmv6 (17) – simplearmv6
	AvcodeccontextIdctFaani AvcodeccontextIdct = "faani" // faani (20) – floating point AAN IDCT
	AvcodeccontextIdctSimpleneon AvcodeccontextIdct = "simpleneon" // simpleneon (22) – simpleneon
	AvcodeccontextIdctSimpleauto AvcodeccontextIdct = "simpleauto" // simpleauto (128) – simpleauto
)

type AvcodeccontextCmpFunc string

const (
	AvcodeccontextCmpFuncSad AvcodeccontextCmpFunc = "sad" // sad (0) – sum of absolute differences, fast
	AvcodeccontextCmpFuncSse AvcodeccontextCmpFunc = "sse" // sse (1) – sum of squared errors
	AvcodeccontextCmpFuncSatd AvcodeccontextCmpFunc = "satd" // satd (2) – sum of absolute Hadamard transformed differences
	AvcodeccontextCmpFuncDct AvcodeccontextCmpFunc = "dct" // dct (3) – sum of absolute DCT transformed differences
	AvcodeccontextCmpFuncPsnr AvcodeccontextCmpFunc = "psnr" // psnr (4) – sum of squared quantization errors (avoid, low quality)
	AvcodeccontextCmpFuncBit AvcodeccontextCmpFunc = "bit" // bit (5) – number of bits needed for the block
	AvcodeccontextCmpFuncRd AvcodeccontextCmpFunc = "rd" // rd (6) – rate distortion optimal, slow
	AvcodeccontextCmpFuncZero AvcodeccontextCmpFunc = "zero" // zero (7) – 0
	AvcodeccontextCmpFuncVsad AvcodeccontextCmpFunc = "vsad" // vsad (8) – sum of absolute vertical differences
	AvcodeccontextCmpFuncVsse AvcodeccontextCmpFunc = "vsse" // vsse (9) – sum of squared vertical differences
	AvcodeccontextCmpFuncNsse AvcodeccontextCmpFunc = "nsse" // nsse (10) – noise preserving sum of squared differences
	AvcodeccontextCmpFuncW53 AvcodeccontextCmpFunc = "w53" // w53 (11) – 5/3 wavelet, only used in snow
	AvcodeccontextCmpFuncW97 AvcodeccontextCmpFunc = "w97" // w97 (12) – 9/7 wavelet, only used in snow
	AvcodeccontextCmpFuncDctmax AvcodeccontextCmpFunc = "dctmax" // dctmax (13) – dctmax
	AvcodeccontextCmpFuncMsad AvcodeccontextCmpFunc = "msad" // msad (15) – sum of absolute differences, median predicted
	AvcodeccontextCmpFuncChroma AvcodeccontextCmpFunc = "chroma" // chroma (256) – chroma
)

type AvcodeccontextDebug string

const (
	AvcodeccontextDebugPict AvcodeccontextDebug = "pict" // pict (0x00000001) – picture info
	AvcodeccontextDebugRc AvcodeccontextDebug = "rc" // rc (0x00000002) – rate control
	AvcodeccontextDebugBitstream AvcodeccontextDebug = "bitstream" // bitstream (0x00000004) – bitstream
	AvcodeccontextDebugMbType AvcodeccontextDebug = "mb_type" // mb_type (0x00000008) – macroblock (MB) type
	AvcodeccontextDebugQp AvcodeccontextDebug = "qp" // qp (0x00000010) – per-block quantization parameter (QP)
	AvcodeccontextDebugDctCoeff AvcodeccontextDebug = "dct_coeff" // dct_coeff (0x00000040) – dct_coeff
	AvcodeccontextDebugSkip AvcodeccontextDebug = "skip" // skip (0x00000080) – skip
	AvcodeccontextDebugStartcode AvcodeccontextDebug = "startcode" // startcode (0x00000100) – startcode
	AvcodeccontextDebugEr AvcodeccontextDebug = "er" // er (0x00000400) – error recognition
	AvcodeccontextDebugMmco AvcodeccontextDebug = "mmco" // mmco (0x00000800) – memory management control operations (H.264)
	AvcodeccontextDebugBugs AvcodeccontextDebug = "bugs" // bugs (0x00001000) – bugs
	AvcodeccontextDebugBuffers AvcodeccontextDebug = "buffers" // buffers (0x00008000) – picture buffer allocations
	AvcodeccontextDebugThreadOps AvcodeccontextDebug = "thread_ops" // thread_ops (0x00010000) – threading operations
	AvcodeccontextDebugGreenMetadata AvcodeccontextDebug = "green_metadata" // green_metadata (0x00800000) – green_metadata
	AvcodeccontextDebugNomc AvcodeccontextDebug = "nomc" // nomc (0x01000000) – skip motion compensation
)

type AvcodeccontextFlags2 string

const (
	AvcodeccontextFlags2Fast AvcodeccontextFlags2 = "fast" // fast (0x00000001) – allow non-spec-compliant speedup tricks
	AvcodeccontextFlags2Noout AvcodeccontextFlags2 = "noout" // noout (0x00000004) – skip bitstream encoding
	AvcodeccontextFlags2LocalHeader AvcodeccontextFlags2 = "local_header" // local_header (0x00000008) – place global headers at every keyframe instead of in extradata
	AvcodeccontextFlags2Chunks AvcodeccontextFlags2 = "chunks" // chunks (0x00008000) – Frame data might be split into multiple chunks
	AvcodeccontextFlags2Ignorecrop AvcodeccontextFlags2 = "ignorecrop" // ignorecrop (0x00010000) – ignore cropping information from sps
	AvcodeccontextFlags2Showall AvcodeccontextFlags2 = "showall" // showall (0x00400000) – Show all frames before the first keyframe
	AvcodeccontextFlags2ExportMvs AvcodeccontextFlags2 = "export_mvs" // export_mvs (0x10000000) – export motion vectors through frame side data
	AvcodeccontextFlags2SkipManual AvcodeccontextFlags2 = "skip_manual" // skip_manual (0x20000000) – do not skip samples and export skip information as frame side data
	AvcodeccontextFlags2AssRoFlushNoop AvcodeccontextFlags2 = "ass_ro_flush_noop" // ass_ro_flush_noop (0x40000000) – do not reset ASS ReadOrder field on flush
)

type AvcodeccontextMbd string

const (
	AvcodeccontextMbdSimple AvcodeccontextMbd = "simple" // simple (0) – use mbcmp
	AvcodeccontextMbdBits AvcodeccontextMbd = "bits" // bits (1) – use fewest bits
	AvcodeccontextMbdRd AvcodeccontextMbd = "rd" // rd (2) – use best rate distortion
)

type AvcodeccontextPred string

const (
	AvcodeccontextPredLeft AvcodeccontextPred = "left" // left (0) – left
	AvcodeccontextPredPlane AvcodeccontextPred = "plane" // plane (1) – plane
	AvcodeccontextPredMedian AvcodeccontextPred = "median" // median (2) – median
)

type AvcodeccontextStrict string

const (
	AvcodeccontextStrictExperimental AvcodeccontextStrict = "experimental" // experimental (-2) – allow non-standardized experimental things
	AvcodeccontextStrictUnofficial AvcodeccontextStrict = "unofficial" // unofficial (-1) – allow unofficial extensions
	AvcodeccontextStrictNormal AvcodeccontextStrict = "normal" // normal (0) – normal
	AvcodeccontextStrictStrict AvcodeccontextStrict = "strict" // strict (1) – strictly conform to all the things in the spec no matter what the consequences
	AvcodeccontextStrictVery AvcodeccontextStrict = "very" // very (2) – strictly conform to a older more strict version of the spec or reference software
)

