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

type AvencAliasPix struct {
	Element *gst.Element
}

func NewAvencAliasPix() (*AvencAliasPix, error) {
	e, err := gst.NewElement("avenc_alias_pix")
	if err != nil {
		return nil, err
	}
	return &AvencAliasPix{Element: e}, nil
}

func NewAvencAliasPixWithName(name string) (*AvencAliasPix, error) {
	e, err := gst.NewElementWithName("avenc_alias_pix", name)
	if err != nil {
		return nil, err
	}
	return &AvencAliasPix{Element: e}, nil
}

// ----- Properties -----

// b-qfactor (float32)
//
// QP factor between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBQfactor() (float32, error) {
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

func (e *AvencAliasPix) SetBQfactor(value float32) error {
	return e.Element.SetProperty("b-qfactor", value)
}

// b-qoffset (float32)
//
// QP offset between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBQoffset() (float32, error) {
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

func (e *AvencAliasPix) SetBQoffset(value float32) error {
	return e.Element.SetProperty("b-qoffset", value)
}

// b-sensitivity (int)
//
// adjust sensitivity of b_frame_strategy 1 (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBSensitivity() (int, error) {
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

func (e *AvencAliasPix) SetBSensitivity(value int) error {
	return e.Element.SetProperty("b-sensitivity", value)
}

// b-strategy (int)
//
// strategy to choose between I/P/B-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBStrategy() (int, error) {
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

func (e *AvencAliasPix) SetBStrategy(value int) error {
	return e.Element.SetProperty("b-strategy", value)
}

// bidir-refine (int)
//
// refine the two motion vectors used in bidirectional macroblocks (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBidirRefine() (int, error) {
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

func (e *AvencAliasPix) SetBidirRefine(value int) error {
	return e.Element.SetProperty("bidir-refine", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBitrate() (int, error) {
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

func (e *AvencAliasPix) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-tolerance (int)
//
// Set video bitrate tolerance (in bits/s). In 1-pass mode, bitrate tolerance specifies how far ratecontrol is willing to deviate from the target average bitrate value. This is not related to minimum/maximum bitrate. Lowering tolerance too much has an adverse effect on quality. (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBitrateTolerance() (int, error) {
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

func (e *AvencAliasPix) SetBitrateTolerance(value int) error {
	return e.Element.SetProperty("bitrate-tolerance", value)
}

// brd-scale (int)
//
// downscale frames for dynamic B-frame decision (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBrdScale() (int, error) {
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

func (e *AvencAliasPix) SetBrdScale(value int) error {
	return e.Element.SetProperty("brd-scale", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetBufsize() (int, error) {
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

func (e *AvencAliasPix) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// chroma-sample-location (GstAvcodeccontextChromaSampleLocationType)
//
// chroma sample location (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetChromaSampleLocation() (interface{}, error) {
	return e.Element.GetProperty("chroma-sample-location")
}

func (e *AvencAliasPix) SetChromaSampleLocation(value interface{}) error {
	return e.Element.SetProperty("chroma-sample-location", value)
}

// chromaoffset (int)
//
// chroma QP offset from luma (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetChromaoffset() (int, error) {
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

func (e *AvencAliasPix) SetChromaoffset(value int) error {
	return e.Element.SetProperty("chromaoffset", value)
}

// cmp (GstAvcodeccontextCmpFunc)
//
// full-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetCmp() (interface{}, error) {
	return e.Element.GetProperty("cmp")
}

func (e *AvencAliasPix) SetCmp(value interface{}) error {
	return e.Element.SetProperty("cmp", value)
}

// coder (GstAvcodeccontextCoder)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetCoder() (interface{}, error) {
	return e.Element.GetProperty("coder")
}

func (e *AvencAliasPix) SetCoder(value interface{}) error {
	return e.Element.SetProperty("coder", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetCompressionLevel() (int, error) {
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

func (e *AvencAliasPix) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// context (int)
//
// context model (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetContext() (int, error) {
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

func (e *AvencAliasPix) SetContext(value int) error {
	return e.Element.SetProperty("context", value)
}

// dark-mask (float32)
//
// compresses dark areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetDarkMask() (float32, error) {
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

func (e *AvencAliasPix) SetDarkMask(value float32) error {
	return e.Element.SetProperty("dark-mask", value)
}

// dc (int)
//
// intra_dc_precision (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetDc() (int, error) {
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

func (e *AvencAliasPix) SetDc(value int) error {
	return e.Element.SetProperty("dc", value)
}

// dct (GstAvcodeccontextDct)
//
// DCT algorithm (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetDct() (interface{}, error) {
	return e.Element.GetProperty("dct")
}

func (e *AvencAliasPix) SetDct(value interface{}) error {
	return e.Element.SetProperty("dct", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencAliasPix) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dia-size (int)
//
// diamond type & size for motion estimation (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetDiaSize() (int, error) {
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

func (e *AvencAliasPix) SetDiaSize(value int) error {
	return e.Element.SetProperty("dia-size", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetDumpSeparator() (string, error) {
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

func (e *AvencAliasPix) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencAliasPix) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencAliasPix) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// field-order (GstAvcodeccontextFieldOrder)
//
// Field order (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("field-order")
}

func (e *AvencAliasPix) SetFieldOrder(value interface{}) error {
	return e.Element.SetProperty("field-order", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencAliasPix) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencAliasPix) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetGlobalQuality() (int, error) {
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

func (e *AvencAliasPix) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// gop-size (int)
//
// set the group of picture (GOP) size (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetGopSize() (int, error) {
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

func (e *AvencAliasPix) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// i-qfactor (float32)
//
// QP factor between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetIQfactor() (float32, error) {
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

func (e *AvencAliasPix) SetIQfactor(value float32) error {
	return e.Element.SetProperty("i-qfactor", value)
}

// i-qoffset (float32)
//
// QP offset between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetIQoffset() (float32, error) {
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

func (e *AvencAliasPix) SetIQoffset(value float32) error {
	return e.Element.SetProperty("i-qoffset", value)
}

// idct (GstAvcodeccontextIdct)
//
// select IDCT implementation (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetIdct() (interface{}, error) {
	return e.Element.GetProperty("idct")
}

func (e *AvencAliasPix) SetIdct(value interface{}) error {
	return e.Element.SetProperty("idct", value)
}

// ildctcmp (GstAvcodeccontextCmpFunc)
//
// interlaced DCT compare function (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetIldctcmp() (interface{}, error) {
	return e.Element.GetProperty("ildctcmp")
}

func (e *AvencAliasPix) SetIldctcmp(value interface{}) error {
	return e.Element.SetProperty("ildctcmp", value)
}

// keyint-min (int)
//
// minimum interval between IDR-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetKeyintMin() (int, error) {
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

func (e *AvencAliasPix) SetKeyintMin(value int) error {
	return e.Element.SetProperty("keyint-min", value)
}

// last-pred (int)
//
// amount of motion predictors from the previous frame (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetLastPred() (int, error) {
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

func (e *AvencAliasPix) SetLastPred(value int) error {
	return e.Element.SetProperty("last-pred", value)
}

// lumi-mask (float32)
//
// compresses bright areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetLumiMask() (float32, error) {
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

func (e *AvencAliasPix) SetLumiMask(value float32) error {
	return e.Element.SetProperty("lumi-mask", value)
}

// max-bframes (int)
//
// set maximum number of B-frames between non-B-frames (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMaxBframes() (int, error) {
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

func (e *AvencAliasPix) SetMaxBframes(value int) error {
	return e.Element.SetProperty("max-bframes", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMaxPixels() (int64, error) {
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

func (e *AvencAliasPix) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMaxrate() (int64, error) {
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

func (e *AvencAliasPix) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// mbcmp (GstAvcodeccontextCmpFunc)
//
// macroblock compare function (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMbcmp() (interface{}, error) {
	return e.Element.GetProperty("mbcmp")
}

func (e *AvencAliasPix) SetMbcmp(value interface{}) error {
	return e.Element.SetProperty("mbcmp", value)
}

// mbd (GstAvcodeccontextMbd)
//
// macroblock decision algorithm (high quality mode) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMbd() (interface{}, error) {
	return e.Element.GetProperty("mbd")
}

func (e *AvencAliasPix) SetMbd(value interface{}) error {
	return e.Element.SetProperty("mbd", value)
}

// mblmax (int)
//
// maximum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMblmax() (int, error) {
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

func (e *AvencAliasPix) SetMblmax(value int) error {
	return e.Element.SetProperty("mblmax", value)
}

// mblmin (int)
//
// minimum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMblmin() (int, error) {
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

func (e *AvencAliasPix) SetMblmin(value int) error {
	return e.Element.SetProperty("mblmin", value)
}

// me-range (int)
//
// limit motion vectors range (1023 for DivX player) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMeRange() (int, error) {
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

func (e *AvencAliasPix) SetMeRange(value int) error {
	return e.Element.SetProperty("me-range", value)
}

// mepc (int)
//
// motion estimation bitrate penalty compensation (1.0 = 256) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMepc() (int, error) {
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

func (e *AvencAliasPix) SetMepc(value int) error {
	return e.Element.SetProperty("mepc", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMinrate() (int64, error) {
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

func (e *AvencAliasPix) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// mpeg-quant (int)
//
// use MPEG quantizers instead of H.263 (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMpegQuant() (int, error) {
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

func (e *AvencAliasPix) SetMpegQuant(value int) error {
	return e.Element.SetProperty("mpeg-quant", value)
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *AvencAliasPix) GetMultipassCacheFile() (string, error) {
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

func (e *AvencAliasPix) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// mv0-threshold (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetMv0Threshold() (int, error) {
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

func (e *AvencAliasPix) SetMv0Threshold(value int) error {
	return e.Element.SetProperty("mv0-threshold", value)
}

// nr (int)
//
// noise reduction (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetNr() (int, error) {
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

func (e *AvencAliasPix) SetNr(value int) error {
	return e.Element.SetProperty("nr", value)
}

// nssew (int)
//
// nsse weight (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetNssew() (int, error) {
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

func (e *AvencAliasPix) SetNssew(value int) error {
	return e.Element.SetProperty("nssew", value)
}

// p-mask (float32)
//
// inter masking (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetPMask() (float32, error) {
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

func (e *AvencAliasPix) SetPMask(value float32) error {
	return e.Element.SetProperty("p-mask", value)
}

// pass (GstLibAVEncPass)
//
// Encoding pass/type

func (e *AvencAliasPix) GetPass() (interface{}, error) {
	return e.Element.GetProperty("pass")
}

func (e *AvencAliasPix) SetPass(value interface{}) error {
	return e.Element.SetProperty("pass", value)
}

// pre-dia-size (int)
//
// diamond type & size for motion estimation pre-pass (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetPreDiaSize() (int, error) {
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

func (e *AvencAliasPix) SetPreDiaSize(value int) error {
	return e.Element.SetProperty("pre-dia-size", value)
}

// precmp (GstAvcodeccontextCmpFunc)
//
// pre motion estimation compare function (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetPrecmp() (interface{}, error) {
	return e.Element.GetProperty("precmp")
}

func (e *AvencAliasPix) SetPrecmp(value interface{}) error {
	return e.Element.SetProperty("precmp", value)
}

// pred (GstAvcodeccontextPred)
//
// prediction method (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetPred() (interface{}, error) {
	return e.Element.GetProperty("pred")
}

func (e *AvencAliasPix) SetPred(value interface{}) error {
	return e.Element.SetProperty("pred", value)
}

// preme (int)
//
// pre motion estimation (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetPreme() (int, error) {
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

func (e *AvencAliasPix) SetPreme(value int) error {
	return e.Element.SetProperty("preme", value)
}

// ps (int)
//
// RTP payload size in bytes (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetPs() (int, error) {
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

func (e *AvencAliasPix) SetPs(value int) error {
	return e.Element.SetProperty("ps", value)
}

// qblur (float32)
//
// video quantizer scale blur (VBR) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetQblur() (float32, error) {
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

func (e *AvencAliasPix) SetQblur(value float32) error {
	return e.Element.SetProperty("qblur", value)
}

// qcomp (float32)
//
// video quantizer scale compression (VBR). Constant of ratecontrol equation. Recommended range for default rc_eq: 0.0-1.0 (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetQcomp() (float32, error) {
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

func (e *AvencAliasPix) SetQcomp(value float32) error {
	return e.Element.SetProperty("qcomp", value)
}

// qdiff (int)
//
// maximum difference between the quantizer scales (VBR) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetQdiff() (int, error) {
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

func (e *AvencAliasPix) SetQdiff(value int) error {
	return e.Element.SetProperty("qdiff", value)
}

// qmax (int)
//
// maximum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetQmax() (int, error) {
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

func (e *AvencAliasPix) SetQmax(value int) error {
	return e.Element.SetProperty("qmax", value)
}

// qmin (int)
//
// minimum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetQmin() (int, error) {
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

func (e *AvencAliasPix) SetQmin(value int) error {
	return e.Element.SetProperty("qmin", value)
}

// quantizer (float32)
//
// Constant Quantizer

func (e *AvencAliasPix) GetQuantizer() (float32, error) {
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

func (e *AvencAliasPix) SetQuantizer(value float32) error {
	return e.Element.SetProperty("quantizer", value)
}

// rc-init-occupancy (int)
//
// number of bits which should be loaded into the rc buffer before decoding starts (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetRcInitOccupancy() (int, error) {
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

func (e *AvencAliasPix) SetRcInitOccupancy(value int) error {
	return e.Element.SetProperty("rc-init-occupancy", value)
}

// rc-max-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetRcMaxVbvUse() (float32, error) {
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

func (e *AvencAliasPix) SetRcMaxVbvUse(value float32) error {
	return e.Element.SetProperty("rc-max-vbv-use", value)
}

// rc-min-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetRcMinVbvUse() (float32, error) {
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

func (e *AvencAliasPix) SetRcMinVbvUse(value float32) error {
	return e.Element.SetProperty("rc-min-vbv-use", value)
}

// refs (int)
//
// reference frames to consider for motion compensation (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetRefs() (int, error) {
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

func (e *AvencAliasPix) SetRefs(value int) error {
	return e.Element.SetProperty("refs", value)
}

// sc-threshold (int)
//
// scene change threshold (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetScThreshold() (int, error) {
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

func (e *AvencAliasPix) SetScThreshold(value int) error {
	return e.Element.SetProperty("sc-threshold", value)
}

// scplx-mask (float32)
//
// spatial complexity masking (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetScplxMask() (float32, error) {
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

func (e *AvencAliasPix) SetScplxMask(value float32) error {
	return e.Element.SetProperty("scplx-mask", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSideDataOnlyPackets() (bool, error) {
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

func (e *AvencAliasPix) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// skip-exp (int)
//
// frame skip exponent (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSkipExp() (int, error) {
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

func (e *AvencAliasPix) SetSkipExp(value int) error {
	return e.Element.SetProperty("skip-exp", value)
}

// skip-factor (int)
//
// frame skip factor (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSkipFactor() (int, error) {
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

func (e *AvencAliasPix) SetSkipFactor(value int) error {
	return e.Element.SetProperty("skip-factor", value)
}

// skip-threshold (int)
//
// frame skip threshold (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSkipThreshold() (int, error) {
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

func (e *AvencAliasPix) SetSkipThreshold(value int) error {
	return e.Element.SetProperty("skip-threshold", value)
}

// skipcmp (GstAvcodeccontextCmpFunc)
//
// frame skip compare function (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSkipcmp() (interface{}, error) {
	return e.Element.GetProperty("skipcmp")
}

func (e *AvencAliasPix) SetSkipcmp(value interface{}) error {
	return e.Element.SetProperty("skipcmp", value)
}

// slices (int)
//
// set the number of slices, used in parallelized encoding (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSlices() (int, error) {
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

func (e *AvencAliasPix) SetSlices(value int) error {
	return e.Element.SetProperty("slices", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencAliasPix) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// subcmp (GstAvcodeccontextCmpFunc)
//
// sub-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSubcmp() (interface{}, error) {
	return e.Element.GetProperty("subcmp")
}

func (e *AvencAliasPix) SetSubcmp(value interface{}) error {
	return e.Element.SetProperty("subcmp", value)
}

// subq (int)
//
// sub-pel motion estimation quality (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetSubq() (int, error) {
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

func (e *AvencAliasPix) SetSubq(value int) error {
	return e.Element.SetProperty("subq", value)
}

// tcplx-mask (float32)
//
// temporal complexity masking (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetTcplxMask() (float32, error) {
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

func (e *AvencAliasPix) SetTcplxMask(value float32) error {
	return e.Element.SetProperty("tcplx-mask", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencAliasPix) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencAliasPix) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetTicksPerFrame() (int, error) {
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

func (e *AvencAliasPix) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// timecode-frame-start (int64)
//
// GOP timecode frame start number, in non-drop-frame format (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetTimecodeFrameStart() (int64, error) {
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

func (e *AvencAliasPix) SetTimecodeFrameStart(value int64) error {
	return e.Element.SetProperty("timecode-frame-start", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencAliasPix) GetTrellis() (int, error) {
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

func (e *AvencAliasPix) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

