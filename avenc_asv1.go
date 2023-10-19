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

type AvencAsv1 struct {
	Element *gst.Element
}

func NewAvencAsv1() (*AvencAsv1, error) {
	e, err := gst.NewElement("avenc_asv1")
	if err != nil {
		return nil, err
	}
	return &AvencAsv1{Element: e}, nil
}

func NewAvencAsv1WithName(name string) (*AvencAsv1, error) {
	e, err := gst.NewElementWithName("avenc_asv1", name)
	if err != nil {
		return nil, err
	}
	return &AvencAsv1{Element: e}, nil
}

// ----- Properties -----

// b-qfactor (float32)
//
// QP factor between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBQfactor() (float32, error) {
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

func (e *AvencAsv1) SetBQfactor(value float32) error {
	return e.Element.SetProperty("b-qfactor", value)
}

// b-qoffset (float32)
//
// QP offset between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBQoffset() (float32, error) {
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

func (e *AvencAsv1) SetBQoffset(value float32) error {
	return e.Element.SetProperty("b-qoffset", value)
}

// b-sensitivity (int)
//
// adjust sensitivity of b_frame_strategy 1 (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBSensitivity() (int, error) {
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

func (e *AvencAsv1) SetBSensitivity(value int) error {
	return e.Element.SetProperty("b-sensitivity", value)
}

// b-strategy (int)
//
// strategy to choose between I/P/B-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBStrategy() (int, error) {
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

func (e *AvencAsv1) SetBStrategy(value int) error {
	return e.Element.SetProperty("b-strategy", value)
}

// bidir-refine (int)
//
// refine the two motion vectors used in bidirectional macroblocks (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBidirRefine() (int, error) {
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

func (e *AvencAsv1) SetBidirRefine(value int) error {
	return e.Element.SetProperty("bidir-refine", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBitrate() (int, error) {
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

func (e *AvencAsv1) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-tolerance (int)
//
// Set video bitrate tolerance (in bits/s). In 1-pass mode, bitrate tolerance specifies how far ratecontrol is willing to deviate from the target average bitrate value. This is not related to minimum/maximum bitrate. Lowering tolerance too much has an adverse effect on quality. (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBitrateTolerance() (int, error) {
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

func (e *AvencAsv1) SetBitrateTolerance(value int) error {
	return e.Element.SetProperty("bitrate-tolerance", value)
}

// brd-scale (int)
//
// downscale frames for dynamic B-frame decision (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBrdScale() (int, error) {
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

func (e *AvencAsv1) SetBrdScale(value int) error {
	return e.Element.SetProperty("brd-scale", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetBufsize() (int, error) {
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

func (e *AvencAsv1) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// chroma-sample-location (GstAvcodeccontextChromaSampleLocationType)
//
// chroma sample location (Generic codec option, might have no effect)

func (e *AvencAsv1) GetChromaSampleLocation() (interface{}, error) {
	return e.Element.GetProperty("chroma-sample-location")
}

func (e *AvencAsv1) SetChromaSampleLocation(value interface{}) error {
	return e.Element.SetProperty("chroma-sample-location", value)
}

// chromaoffset (int)
//
// chroma QP offset from luma (Generic codec option, might have no effect)

func (e *AvencAsv1) GetChromaoffset() (int, error) {
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

func (e *AvencAsv1) SetChromaoffset(value int) error {
	return e.Element.SetProperty("chromaoffset", value)
}

// cmp (GstAvcodeccontextCmpFunc)
//
// full-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencAsv1) GetCmp() (interface{}, error) {
	return e.Element.GetProperty("cmp")
}

func (e *AvencAsv1) SetCmp(value interface{}) error {
	return e.Element.SetProperty("cmp", value)
}

// coder (GstAvcodeccontextCoder)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetCoder() (interface{}, error) {
	return e.Element.GetProperty("coder")
}

func (e *AvencAsv1) SetCoder(value interface{}) error {
	return e.Element.SetProperty("coder", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetCompressionLevel() (int, error) {
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

func (e *AvencAsv1) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// context (int)
//
// context model (Generic codec option, might have no effect)

func (e *AvencAsv1) GetContext() (int, error) {
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

func (e *AvencAsv1) SetContext(value int) error {
	return e.Element.SetProperty("context", value)
}

// dark-mask (float32)
//
// compresses dark areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencAsv1) GetDarkMask() (float32, error) {
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

func (e *AvencAsv1) SetDarkMask(value float32) error {
	return e.Element.SetProperty("dark-mask", value)
}

// dc (int)
//
// intra_dc_precision (Generic codec option, might have no effect)

func (e *AvencAsv1) GetDc() (int, error) {
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

func (e *AvencAsv1) SetDc(value int) error {
	return e.Element.SetProperty("dc", value)
}

// dct (GstAvcodeccontextDct)
//
// DCT algorithm (Generic codec option, might have no effect)

func (e *AvencAsv1) GetDct() (interface{}, error) {
	return e.Element.GetProperty("dct")
}

func (e *AvencAsv1) SetDct(value interface{}) error {
	return e.Element.SetProperty("dct", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencAsv1) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencAsv1) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dia-size (int)
//
// diamond type & size for motion estimation (Generic codec option, might have no effect)

func (e *AvencAsv1) GetDiaSize() (int, error) {
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

func (e *AvencAsv1) SetDiaSize(value int) error {
	return e.Element.SetProperty("dia-size", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencAsv1) GetDumpSeparator() (string, error) {
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

func (e *AvencAsv1) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencAsv1) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencAsv1) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencAsv1) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencAsv1) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// field-order (GstAvcodeccontextFieldOrder)
//
// Field order (Generic codec option, might have no effect)

func (e *AvencAsv1) GetFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("field-order")
}

func (e *AvencAsv1) SetFieldOrder(value interface{}) error {
	return e.Element.SetProperty("field-order", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencAsv1) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencAsv1) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetGlobalQuality() (int, error) {
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

func (e *AvencAsv1) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// gop-size (int)
//
// set the group of picture (GOP) size (Generic codec option, might have no effect)

func (e *AvencAsv1) GetGopSize() (int, error) {
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

func (e *AvencAsv1) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// i-qfactor (float32)
//
// QP factor between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetIQfactor() (float32, error) {
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

func (e *AvencAsv1) SetIQfactor(value float32) error {
	return e.Element.SetProperty("i-qfactor", value)
}

// i-qoffset (float32)
//
// QP offset between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetIQoffset() (float32, error) {
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

func (e *AvencAsv1) SetIQoffset(value float32) error {
	return e.Element.SetProperty("i-qoffset", value)
}

// idct (GstAvcodeccontextIdct)
//
// select IDCT implementation (Generic codec option, might have no effect)

func (e *AvencAsv1) GetIdct() (interface{}, error) {
	return e.Element.GetProperty("idct")
}

func (e *AvencAsv1) SetIdct(value interface{}) error {
	return e.Element.SetProperty("idct", value)
}

// ildctcmp (GstAvcodeccontextCmpFunc)
//
// interlaced DCT compare function (Generic codec option, might have no effect)

func (e *AvencAsv1) GetIldctcmp() (interface{}, error) {
	return e.Element.GetProperty("ildctcmp")
}

func (e *AvencAsv1) SetIldctcmp(value interface{}) error {
	return e.Element.SetProperty("ildctcmp", value)
}

// keyint-min (int)
//
// minimum interval between IDR-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetKeyintMin() (int, error) {
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

func (e *AvencAsv1) SetKeyintMin(value int) error {
	return e.Element.SetProperty("keyint-min", value)
}

// last-pred (int)
//
// amount of motion predictors from the previous frame (Generic codec option, might have no effect)

func (e *AvencAsv1) GetLastPred() (int, error) {
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

func (e *AvencAsv1) SetLastPred(value int) error {
	return e.Element.SetProperty("last-pred", value)
}

// lumi-mask (float32)
//
// compresses bright areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencAsv1) GetLumiMask() (float32, error) {
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

func (e *AvencAsv1) SetLumiMask(value float32) error {
	return e.Element.SetProperty("lumi-mask", value)
}

// max-bframes (int)
//
// set maximum number of B-frames between non-B-frames (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMaxBframes() (int, error) {
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

func (e *AvencAsv1) SetMaxBframes(value int) error {
	return e.Element.SetProperty("max-bframes", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMaxPixels() (int64, error) {
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

func (e *AvencAsv1) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMaxrate() (int64, error) {
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

func (e *AvencAsv1) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// mbcmp (GstAvcodeccontextCmpFunc)
//
// macroblock compare function (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMbcmp() (interface{}, error) {
	return e.Element.GetProperty("mbcmp")
}

func (e *AvencAsv1) SetMbcmp(value interface{}) error {
	return e.Element.SetProperty("mbcmp", value)
}

// mbd (GstAvcodeccontextMbd)
//
// macroblock decision algorithm (high quality mode) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMbd() (interface{}, error) {
	return e.Element.GetProperty("mbd")
}

func (e *AvencAsv1) SetMbd(value interface{}) error {
	return e.Element.SetProperty("mbd", value)
}

// mblmax (int)
//
// maximum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMblmax() (int, error) {
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

func (e *AvencAsv1) SetMblmax(value int) error {
	return e.Element.SetProperty("mblmax", value)
}

// mblmin (int)
//
// minimum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMblmin() (int, error) {
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

func (e *AvencAsv1) SetMblmin(value int) error {
	return e.Element.SetProperty("mblmin", value)
}

// me-range (int)
//
// limit motion vectors range (1023 for DivX player) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMeRange() (int, error) {
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

func (e *AvencAsv1) SetMeRange(value int) error {
	return e.Element.SetProperty("me-range", value)
}

// mepc (int)
//
// motion estimation bitrate penalty compensation (1.0 = 256) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMepc() (int, error) {
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

func (e *AvencAsv1) SetMepc(value int) error {
	return e.Element.SetProperty("mepc", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMinrate() (int64, error) {
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

func (e *AvencAsv1) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// mpeg-quant (int)
//
// use MPEG quantizers instead of H.263 (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMpegQuant() (int, error) {
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

func (e *AvencAsv1) SetMpegQuant(value int) error {
	return e.Element.SetProperty("mpeg-quant", value)
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *AvencAsv1) GetMultipassCacheFile() (string, error) {
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

func (e *AvencAsv1) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// mv0-threshold (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetMv0Threshold() (int, error) {
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

func (e *AvencAsv1) SetMv0Threshold(value int) error {
	return e.Element.SetProperty("mv0-threshold", value)
}

// nr (int)
//
// noise reduction (Generic codec option, might have no effect)

func (e *AvencAsv1) GetNr() (int, error) {
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

func (e *AvencAsv1) SetNr(value int) error {
	return e.Element.SetProperty("nr", value)
}

// nssew (int)
//
// nsse weight (Generic codec option, might have no effect)

func (e *AvencAsv1) GetNssew() (int, error) {
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

func (e *AvencAsv1) SetNssew(value int) error {
	return e.Element.SetProperty("nssew", value)
}

// p-mask (float32)
//
// inter masking (Generic codec option, might have no effect)

func (e *AvencAsv1) GetPMask() (float32, error) {
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

func (e *AvencAsv1) SetPMask(value float32) error {
	return e.Element.SetProperty("p-mask", value)
}

// pass (GstLibAVEncPass)
//
// Encoding pass/type

func (e *AvencAsv1) GetPass() (interface{}, error) {
	return e.Element.GetProperty("pass")
}

func (e *AvencAsv1) SetPass(value interface{}) error {
	return e.Element.SetProperty("pass", value)
}

// pre-dia-size (int)
//
// diamond type & size for motion estimation pre-pass (Generic codec option, might have no effect)

func (e *AvencAsv1) GetPreDiaSize() (int, error) {
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

func (e *AvencAsv1) SetPreDiaSize(value int) error {
	return e.Element.SetProperty("pre-dia-size", value)
}

// precmp (GstAvcodeccontextCmpFunc)
//
// pre motion estimation compare function (Generic codec option, might have no effect)

func (e *AvencAsv1) GetPrecmp() (interface{}, error) {
	return e.Element.GetProperty("precmp")
}

func (e *AvencAsv1) SetPrecmp(value interface{}) error {
	return e.Element.SetProperty("precmp", value)
}

// pred (GstAvcodeccontextPred)
//
// prediction method (Generic codec option, might have no effect)

func (e *AvencAsv1) GetPred() (interface{}, error) {
	return e.Element.GetProperty("pred")
}

func (e *AvencAsv1) SetPred(value interface{}) error {
	return e.Element.SetProperty("pred", value)
}

// preme (int)
//
// pre motion estimation (Generic codec option, might have no effect)

func (e *AvencAsv1) GetPreme() (int, error) {
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

func (e *AvencAsv1) SetPreme(value int) error {
	return e.Element.SetProperty("preme", value)
}

// ps (int)
//
// RTP payload size in bytes (Generic codec option, might have no effect)

func (e *AvencAsv1) GetPs() (int, error) {
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

func (e *AvencAsv1) SetPs(value int) error {
	return e.Element.SetProperty("ps", value)
}

// qblur (float32)
//
// video quantizer scale blur (VBR) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetQblur() (float32, error) {
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

func (e *AvencAsv1) SetQblur(value float32) error {
	return e.Element.SetProperty("qblur", value)
}

// qcomp (float32)
//
// video quantizer scale compression (VBR). Constant of ratecontrol equation. Recommended range for default rc_eq: 0.0-1.0 (Generic codec option, might have no effect)

func (e *AvencAsv1) GetQcomp() (float32, error) {
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

func (e *AvencAsv1) SetQcomp(value float32) error {
	return e.Element.SetProperty("qcomp", value)
}

// qdiff (int)
//
// maximum difference between the quantizer scales (VBR) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetQdiff() (int, error) {
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

func (e *AvencAsv1) SetQdiff(value int) error {
	return e.Element.SetProperty("qdiff", value)
}

// qmax (int)
//
// maximum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetQmax() (int, error) {
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

func (e *AvencAsv1) SetQmax(value int) error {
	return e.Element.SetProperty("qmax", value)
}

// qmin (int)
//
// minimum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetQmin() (int, error) {
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

func (e *AvencAsv1) SetQmin(value int) error {
	return e.Element.SetProperty("qmin", value)
}

// quantizer (float32)
//
// Constant Quantizer

func (e *AvencAsv1) GetQuantizer() (float32, error) {
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

func (e *AvencAsv1) SetQuantizer(value float32) error {
	return e.Element.SetProperty("quantizer", value)
}

// rc-init-occupancy (int)
//
// number of bits which should be loaded into the rc buffer before decoding starts (Generic codec option, might have no effect)

func (e *AvencAsv1) GetRcInitOccupancy() (int, error) {
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

func (e *AvencAsv1) SetRcInitOccupancy(value int) error {
	return e.Element.SetProperty("rc-init-occupancy", value)
}

// rc-max-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetRcMaxVbvUse() (float32, error) {
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

func (e *AvencAsv1) SetRcMaxVbvUse(value float32) error {
	return e.Element.SetProperty("rc-max-vbv-use", value)
}

// rc-min-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetRcMinVbvUse() (float32, error) {
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

func (e *AvencAsv1) SetRcMinVbvUse(value float32) error {
	return e.Element.SetProperty("rc-min-vbv-use", value)
}

// refs (int)
//
// reference frames to consider for motion compensation (Generic codec option, might have no effect)

func (e *AvencAsv1) GetRefs() (int, error) {
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

func (e *AvencAsv1) SetRefs(value int) error {
	return e.Element.SetProperty("refs", value)
}

// sc-threshold (int)
//
// scene change threshold (Generic codec option, might have no effect)

func (e *AvencAsv1) GetScThreshold() (int, error) {
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

func (e *AvencAsv1) SetScThreshold(value int) error {
	return e.Element.SetProperty("sc-threshold", value)
}

// scplx-mask (float32)
//
// spatial complexity masking (Generic codec option, might have no effect)

func (e *AvencAsv1) GetScplxMask() (float32, error) {
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

func (e *AvencAsv1) SetScplxMask(value float32) error {
	return e.Element.SetProperty("scplx-mask", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSideDataOnlyPackets() (bool, error) {
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

func (e *AvencAsv1) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// skip-exp (int)
//
// frame skip exponent (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSkipExp() (int, error) {
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

func (e *AvencAsv1) SetSkipExp(value int) error {
	return e.Element.SetProperty("skip-exp", value)
}

// skip-factor (int)
//
// frame skip factor (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSkipFactor() (int, error) {
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

func (e *AvencAsv1) SetSkipFactor(value int) error {
	return e.Element.SetProperty("skip-factor", value)
}

// skip-threshold (int)
//
// frame skip threshold (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSkipThreshold() (int, error) {
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

func (e *AvencAsv1) SetSkipThreshold(value int) error {
	return e.Element.SetProperty("skip-threshold", value)
}

// skipcmp (GstAvcodeccontextCmpFunc)
//
// frame skip compare function (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSkipcmp() (interface{}, error) {
	return e.Element.GetProperty("skipcmp")
}

func (e *AvencAsv1) SetSkipcmp(value interface{}) error {
	return e.Element.SetProperty("skipcmp", value)
}

// slices (int)
//
// set the number of slices, used in parallelized encoding (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSlices() (int, error) {
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

func (e *AvencAsv1) SetSlices(value int) error {
	return e.Element.SetProperty("slices", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencAsv1) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencAsv1) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// subcmp (GstAvcodeccontextCmpFunc)
//
// sub-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSubcmp() (interface{}, error) {
	return e.Element.GetProperty("subcmp")
}

func (e *AvencAsv1) SetSubcmp(value interface{}) error {
	return e.Element.SetProperty("subcmp", value)
}

// subq (int)
//
// sub-pel motion estimation quality (Generic codec option, might have no effect)

func (e *AvencAsv1) GetSubq() (int, error) {
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

func (e *AvencAsv1) SetSubq(value int) error {
	return e.Element.SetProperty("subq", value)
}

// tcplx-mask (float32)
//
// temporal complexity masking (Generic codec option, might have no effect)

func (e *AvencAsv1) GetTcplxMask() (float32, error) {
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

func (e *AvencAsv1) SetTcplxMask(value float32) error {
	return e.Element.SetProperty("tcplx-mask", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencAsv1) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencAsv1) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencAsv1) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencAsv1) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAsv1) GetTicksPerFrame() (int, error) {
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

func (e *AvencAsv1) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// timecode-frame-start (int64)
//
// GOP timecode frame start number, in non-drop-frame format (Generic codec option, might have no effect)

func (e *AvencAsv1) GetTimecodeFrameStart() (int64, error) {
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

func (e *AvencAsv1) SetTimecodeFrameStart(value int64) error {
	return e.Element.SetProperty("timecode-frame-start", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencAsv1) GetTrellis() (int, error) {
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

func (e *AvencAsv1) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

