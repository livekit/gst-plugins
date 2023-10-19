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

type AvencH263 struct {
	Element *gst.Element
}

func NewAvencH263() (*AvencH263, error) {
	e, err := gst.NewElement("avenc_h263")
	if err != nil {
		return nil, err
	}
	return &AvencH263{Element: e}, nil
}

func NewAvencH263WithName(name string) (*AvencH263, error) {
	e, err := gst.NewElementWithName("avenc_h263", name)
	if err != nil {
		return nil, err
	}
	return &AvencH263{Element: e}, nil
}

// ----- Properties -----

// a53cc (bool)
//
// Use A53 Closed Captions (if available) (Private codec option)

func (e *AvencH263) GetA53Cc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("a53cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencH263) SetA53Cc(value bool) error {
	return e.Element.SetProperty("a53cc", value)
}

// b-qfactor (float32)
//
// QP factor between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencH263) GetBQfactor() (float32, error) {
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

func (e *AvencH263) SetBQfactor(value float32) error {
	return e.Element.SetProperty("b-qfactor", value)
}

// b-qoffset (float32)
//
// QP offset between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencH263) GetBQoffset() (float32, error) {
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

func (e *AvencH263) SetBQoffset(value float32) error {
	return e.Element.SetProperty("b-qoffset", value)
}

// b-sensitivity (int)
//
// Adjust sensitivity of b_frame_strategy 1 (Private codec option)

func (e *AvencH263) GetBSensitivity() (int, error) {
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

func (e *AvencH263) SetBSensitivity(value int) error {
	return e.Element.SetProperty("b-sensitivity", value)
}

// b-strategy (int)
//
// Strategy to choose between I/P/B-frames (Private codec option)

func (e *AvencH263) GetBStrategy() (int, error) {
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

func (e *AvencH263) SetBStrategy(value int) error {
	return e.Element.SetProperty("b-strategy", value)
}

// bidir-refine (int)
//
// refine the two motion vectors used in bidirectional macroblocks (Generic codec option, might have no effect)

func (e *AvencH263) GetBidirRefine() (int, error) {
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

func (e *AvencH263) SetBidirRefine(value int) error {
	return e.Element.SetProperty("bidir-refine", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencH263) GetBitrate() (int, error) {
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

func (e *AvencH263) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-tolerance (int)
//
// Set video bitrate tolerance (in bits/s). In 1-pass mode, bitrate tolerance specifies how far ratecontrol is willing to deviate from the target average bitrate value. This is not related to minimum/maximum bitrate. Lowering tolerance too much has an adverse effect on quality. (Generic codec option, might have no effect)

func (e *AvencH263) GetBitrateTolerance() (int, error) {
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

func (e *AvencH263) SetBitrateTolerance(value int) error {
	return e.Element.SetProperty("bitrate-tolerance", value)
}

// border-mask (float32)
//
// increase the quantizer for macroblocks close to borders (Private codec option)

func (e *AvencH263) GetBorderMask() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("border-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencH263) SetBorderMask(value float32) error {
	return e.Element.SetProperty("border-mask", value)
}

// brd-scale (int)
//
// Downscale frames for dynamic B-frame decision (Private codec option)

func (e *AvencH263) GetBrdScale() (int, error) {
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

func (e *AvencH263) SetBrdScale(value int) error {
	return e.Element.SetProperty("brd-scale", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencH263) GetBufsize() (int, error) {
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

func (e *AvencH263) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// chroma-elim-threshold (int)
//
// single coefficient elimination threshold for chrominance (negative values also consider dc coefficient) (Private codec option)

func (e *AvencH263) GetChromaElimThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("chroma-elim-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetChromaElimThreshold(value int) error {
	return e.Element.SetProperty("chroma-elim-threshold", value)
}

// chroma-sample-location (GstAvcodeccontextChromaSampleLocationType)
//
// chroma sample location (Generic codec option, might have no effect)

func (e *AvencH263) GetChromaSampleLocation() (interface{}, error) {
	return e.Element.GetProperty("chroma-sample-location")
}

func (e *AvencH263) SetChromaSampleLocation(value interface{}) error {
	return e.Element.SetProperty("chroma-sample-location", value)
}

// chromaoffset (int)
//
// chroma QP offset from luma (Generic codec option, might have no effect)

func (e *AvencH263) GetChromaoffset() (int, error) {
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

func (e *AvencH263) SetChromaoffset(value int) error {
	return e.Element.SetProperty("chromaoffset", value)
}

// cmp (GstAvcodeccontextCmpFunc)
//
// full-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencH263) GetCmp() (interface{}, error) {
	return e.Element.GetProperty("cmp")
}

func (e *AvencH263) SetCmp(value interface{}) error {
	return e.Element.SetProperty("cmp", value)
}

// coder (GstAvcodeccontextCoder)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetCoder() (interface{}, error) {
	return e.Element.GetProperty("coder")
}

func (e *AvencH263) SetCoder(value interface{}) error {
	return e.Element.SetProperty("coder", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetCompressionLevel() (int, error) {
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

func (e *AvencH263) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// context (int)
//
// context model (Generic codec option, might have no effect)

func (e *AvencH263) GetContext() (int, error) {
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

func (e *AvencH263) SetContext(value int) error {
	return e.Element.SetProperty("context", value)
}

// dark-mask (float32)
//
// compresses dark areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencH263) GetDarkMask() (float32, error) {
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

func (e *AvencH263) SetDarkMask(value float32) error {
	return e.Element.SetProperty("dark-mask", value)
}

// dc (int)
//
// intra_dc_precision (Generic codec option, might have no effect)

func (e *AvencH263) GetDc() (int, error) {
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

func (e *AvencH263) SetDc(value int) error {
	return e.Element.SetProperty("dc", value)
}

// dct (GstAvcodeccontextDct)
//
// DCT algorithm (Generic codec option, might have no effect)

func (e *AvencH263) GetDct() (interface{}, error) {
	return e.Element.GetProperty("dct")
}

func (e *AvencH263) SetDct(value interface{}) error {
	return e.Element.SetProperty("dct", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencH263) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencH263) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dia-size (int)
//
// diamond type & size for motion estimation (Generic codec option, might have no effect)

func (e *AvencH263) GetDiaSize() (int, error) {
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

func (e *AvencH263) SetDiaSize(value int) error {
	return e.Element.SetProperty("dia-size", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencH263) GetDumpSeparator() (string, error) {
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

func (e *AvencH263) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencH263) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencH263) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// error-rate (int)
//
// Simulate errors in the bitstream to test error concealment. (Private codec option)

func (e *AvencH263) GetErrorRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("error-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetErrorRate(value int) error {
	return e.Element.SetProperty("error-rate", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencH263) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencH263) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// field-order (GstAvcodeccontextFieldOrder)
//
// Field order (Generic codec option, might have no effect)

func (e *AvencH263) GetFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("field-order")
}

func (e *AvencH263) SetFieldOrder(value interface{}) error {
	return e.Element.SetProperty("field-order", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencH263) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencH263) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// force-duplicated-matrix (bool)
//
// Always write luma and chroma matrix for mjpeg, useful for rtp streaming. (Private codec option)

func (e *AvencH263) GetForceDuplicatedMatrix() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-duplicated-matrix")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencH263) SetForceDuplicatedMatrix(value bool) error {
	return e.Element.SetProperty("force-duplicated-matrix", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetGlobalQuality() (int, error) {
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

func (e *AvencH263) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// gop-size (int)
//
// set the group of picture (GOP) size (Generic codec option, might have no effect)

func (e *AvencH263) GetGopSize() (int, error) {
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

func (e *AvencH263) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// i-qfactor (float32)
//
// QP factor between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencH263) GetIQfactor() (float32, error) {
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

func (e *AvencH263) SetIQfactor(value float32) error {
	return e.Element.SetProperty("i-qfactor", value)
}

// i-qoffset (float32)
//
// QP offset between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencH263) GetIQoffset() (float32, error) {
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

func (e *AvencH263) SetIQoffset(value float32) error {
	return e.Element.SetProperty("i-qoffset", value)
}

// ibias (int)
//
// intra quant bias (Private codec option)

func (e *AvencH263) GetIbias() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ibias")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetIbias(value int) error {
	return e.Element.SetProperty("ibias", value)
}

// idct (GstAvcodeccontextIdct)
//
// select IDCT implementation (Generic codec option, might have no effect)

func (e *AvencH263) GetIdct() (interface{}, error) {
	return e.Element.GetProperty("idct")
}

func (e *AvencH263) SetIdct(value interface{}) error {
	return e.Element.SetProperty("idct", value)
}

// ildctcmp (GstAvcodeccontextCmpFunc)
//
// interlaced DCT compare function (Generic codec option, might have no effect)

func (e *AvencH263) GetIldctcmp() (interface{}, error) {
	return e.Element.GetProperty("ildctcmp")
}

func (e *AvencH263) SetIldctcmp(value interface{}) error {
	return e.Element.SetProperty("ildctcmp", value)
}

// intra-penalty (int)
//
// Penalty for intra blocks in block decision (Private codec option)

func (e *AvencH263) GetIntraPenalty() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("intra-penalty")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetIntraPenalty(value int) error {
	return e.Element.SetProperty("intra-penalty", value)
}

// keyint-min (int)
//
// minimum interval between IDR-frames (Generic codec option, might have no effect)

func (e *AvencH263) GetKeyintMin() (int, error) {
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

func (e *AvencH263) SetKeyintMin(value int) error {
	return e.Element.SetProperty("keyint-min", value)
}

// last-pred (int)
//
// amount of motion predictors from the previous frame (Generic codec option, might have no effect)

func (e *AvencH263) GetLastPred() (int, error) {
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

func (e *AvencH263) SetLastPred(value int) error {
	return e.Element.SetProperty("last-pred", value)
}

// lmax (int)
//
// maximum Lagrange factor (VBR) (Private codec option)

func (e *AvencH263) GetLmax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("lmax")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetLmax(value int) error {
	return e.Element.SetProperty("lmax", value)
}

// lmin (int)
//
// minimum Lagrange factor (VBR) (Private codec option)

func (e *AvencH263) GetLmin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("lmin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetLmin(value int) error {
	return e.Element.SetProperty("lmin", value)
}

// luma-elim-threshold (int)
//
// single coefficient elimination threshold for luminance (negative values also consider dc coefficient) (Private codec option)

func (e *AvencH263) GetLumaElimThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("luma-elim-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetLumaElimThreshold(value int) error {
	return e.Element.SetProperty("luma-elim-threshold", value)
}

// lumi-mask (float32)
//
// compresses bright areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencH263) GetLumiMask() (float32, error) {
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

func (e *AvencH263) SetLumiMask(value float32) error {
	return e.Element.SetProperty("lumi-mask", value)
}

// max-bframes (int)
//
// set maximum number of B-frames between non-B-frames (Generic codec option, might have no effect)

func (e *AvencH263) GetMaxBframes() (int, error) {
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

func (e *AvencH263) SetMaxBframes(value int) error {
	return e.Element.SetProperty("max-bframes", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencH263) GetMaxPixels() (int64, error) {
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

func (e *AvencH263) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencH263) GetMaxrate() (int64, error) {
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

func (e *AvencH263) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// mb-info (int)
//
// emit macroblock info for RFC 2190 packetization, the parameter value is the maximum payload size (Private codec option)

func (e *AvencH263) GetMbInfo() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mb-info")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetMbInfo(value int) error {
	return e.Element.SetProperty("mb-info", value)
}

// mbcmp (GstAvcodeccontextCmpFunc)
//
// macroblock compare function (Generic codec option, might have no effect)

func (e *AvencH263) GetMbcmp() (interface{}, error) {
	return e.Element.GetProperty("mbcmp")
}

func (e *AvencH263) SetMbcmp(value interface{}) error {
	return e.Element.SetProperty("mbcmp", value)
}

// mbd (GstAvcodeccontextMbd)
//
// macroblock decision algorithm (high quality mode) (Generic codec option, might have no effect)

func (e *AvencH263) GetMbd() (interface{}, error) {
	return e.Element.GetProperty("mbd")
}

func (e *AvencH263) SetMbd(value interface{}) error {
	return e.Element.SetProperty("mbd", value)
}

// mblmax (int)
//
// maximum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencH263) GetMblmax() (int, error) {
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

func (e *AvencH263) SetMblmax(value int) error {
	return e.Element.SetProperty("mblmax", value)
}

// mblmin (int)
//
// minimum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencH263) GetMblmin() (int, error) {
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

func (e *AvencH263) SetMblmin(value int) error {
	return e.Element.SetProperty("mblmin", value)
}

// me-range (int)
//
// limit motion vectors range (1023 for DivX player) (Generic codec option, might have no effect)

func (e *AvencH263) GetMeRange() (int, error) {
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

func (e *AvencH263) SetMeRange(value int) error {
	return e.Element.SetProperty("me-range", value)
}

// mepc (int)
//
// Motion estimation bitrate penalty compensation (1.0 = 256) (Private codec option)

func (e *AvencH263) GetMepc() (int, error) {
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

func (e *AvencH263) SetMepc(value int) error {
	return e.Element.SetProperty("mepc", value)
}

// mepre (int)
//
// pre motion estimation (Private codec option)

func (e *AvencH263) GetMepre() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mepre")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetMepre(value int) error {
	return e.Element.SetProperty("mepre", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencH263) GetMinrate() (int64, error) {
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

func (e *AvencH263) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// motion-est (GstH263EncoderMotionEst)
//
// motion estimation algorithm (Private codec option)

func (e *AvencH263) GetMotionEst() (interface{}, error) {
	return e.Element.GetProperty("motion-est")
}

func (e *AvencH263) SetMotionEst(value interface{}) error {
	return e.Element.SetProperty("motion-est", value)
}

// mpeg-quant (int)
//
// Use MPEG quantizers instead of H.263 (Private codec option)

func (e *AvencH263) GetMpegQuant() (int, error) {
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

func (e *AvencH263) SetMpegQuant(value int) error {
	return e.Element.SetProperty("mpeg-quant", value)
}

// mpv-flags (GstH263EncoderMpvFlags)
//
// Flags common for all mpegvideo-based encoders. (Private codec option)

func (e *AvencH263) GetMpvFlags() (interface{}, error) {
	return e.Element.GetProperty("mpv-flags")
}

func (e *AvencH263) SetMpvFlags(value interface{}) error {
	return e.Element.SetProperty("mpv-flags", value)
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *AvencH263) GetMultipassCacheFile() (string, error) {
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

func (e *AvencH263) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// mv0-threshold (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetMv0Threshold() (int, error) {
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

func (e *AvencH263) SetMv0Threshold(value int) error {
	return e.Element.SetProperty("mv0-threshold", value)
}

// noise-reduction (int)
//
// Noise reduction (Private codec option)

func (e *AvencH263) GetNoiseReduction() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("noise-reduction")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetNoiseReduction(value int) error {
	return e.Element.SetProperty("noise-reduction", value)
}

// nr (int)
//
// noise reduction (Generic codec option, might have no effect)

func (e *AvencH263) GetNr() (int, error) {
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

func (e *AvencH263) SetNr(value int) error {
	return e.Element.SetProperty("nr", value)
}

// nssew (int)
//
// nsse weight (Generic codec option, might have no effect)

func (e *AvencH263) GetNssew() (int, error) {
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

func (e *AvencH263) SetNssew(value int) error {
	return e.Element.SetProperty("nssew", value)
}

// obmc (bool)
//
// use overlapped block motion compensation. (Private codec option)

func (e *AvencH263) GetObmc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("obmc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencH263) SetObmc(value bool) error {
	return e.Element.SetProperty("obmc", value)
}

// p-mask (float32)
//
// inter masking (Generic codec option, might have no effect)

func (e *AvencH263) GetPMask() (float32, error) {
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

func (e *AvencH263) SetPMask(value float32) error {
	return e.Element.SetProperty("p-mask", value)
}

// pass (GstLibAVEncPass)
//
// Encoding pass/type

func (e *AvencH263) GetPass() (interface{}, error) {
	return e.Element.GetProperty("pass")
}

func (e *AvencH263) SetPass(value interface{}) error {
	return e.Element.SetProperty("pass", value)
}

// pbias (int)
//
// inter quant bias (Private codec option)

func (e *AvencH263) GetPbias() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pbias")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetPbias(value int) error {
	return e.Element.SetProperty("pbias", value)
}

// pre-dia-size (int)
//
// diamond type & size for motion estimation pre-pass (Generic codec option, might have no effect)

func (e *AvencH263) GetPreDiaSize() (int, error) {
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

func (e *AvencH263) SetPreDiaSize(value int) error {
	return e.Element.SetProperty("pre-dia-size", value)
}

// precmp (GstAvcodeccontextCmpFunc)
//
// pre motion estimation compare function (Generic codec option, might have no effect)

func (e *AvencH263) GetPrecmp() (interface{}, error) {
	return e.Element.GetProperty("precmp")
}

func (e *AvencH263) SetPrecmp(value interface{}) error {
	return e.Element.SetProperty("precmp", value)
}

// pred (GstAvcodeccontextPred)
//
// prediction method (Generic codec option, might have no effect)

func (e *AvencH263) GetPred() (interface{}, error) {
	return e.Element.GetProperty("pred")
}

func (e *AvencH263) SetPred(value interface{}) error {
	return e.Element.SetProperty("pred", value)
}

// preme (int)
//
// pre motion estimation (Generic codec option, might have no effect)

func (e *AvencH263) GetPreme() (int, error) {
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

func (e *AvencH263) SetPreme(value int) error {
	return e.Element.SetProperty("preme", value)
}

// ps (int)
//
// RTP payload size in bytes (Private codec option)

func (e *AvencH263) GetPs() (int, error) {
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

func (e *AvencH263) SetPs(value int) error {
	return e.Element.SetProperty("ps", value)
}

// qblur (float32)
//
// video quantizer scale blur (VBR) (Generic codec option, might have no effect)

func (e *AvencH263) GetQblur() (float32, error) {
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

func (e *AvencH263) SetQblur(value float32) error {
	return e.Element.SetProperty("qblur", value)
}

// qcomp (float32)
//
// video quantizer scale compression (VBR). Constant of ratecontrol equation. Recommended range for default rc_eq: 0.0-1.0 (Generic codec option, might have no effect)

func (e *AvencH263) GetQcomp() (float32, error) {
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

func (e *AvencH263) SetQcomp(value float32) error {
	return e.Element.SetProperty("qcomp", value)
}

// qdiff (int)
//
// maximum difference between the quantizer scales (VBR) (Generic codec option, might have no effect)

func (e *AvencH263) GetQdiff() (int, error) {
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

func (e *AvencH263) SetQdiff(value int) error {
	return e.Element.SetProperty("qdiff", value)
}

// qmax (int)
//
// maximum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencH263) GetQmax() (int, error) {
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

func (e *AvencH263) SetQmax(value int) error {
	return e.Element.SetProperty("qmax", value)
}

// qmin (int)
//
// minimum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencH263) GetQmin() (int, error) {
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

func (e *AvencH263) SetQmin(value int) error {
	return e.Element.SetProperty("qmin", value)
}

// qsquish (float32)
//
// how to keep quantizer between qmin and qmax (0 = clip, 1 = use differentiable function) (Private codec option)

func (e *AvencH263) GetQsquish() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("qsquish")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencH263) SetQsquish(value float32) error {
	return e.Element.SetProperty("qsquish", value)
}

// quantizer (float32)
//
// Constant Quantizer

func (e *AvencH263) GetQuantizer() (float32, error) {
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

func (e *AvencH263) SetQuantizer(value float32) error {
	return e.Element.SetProperty("quantizer", value)
}

// quantizer-noise-shaping (int)
//
// (null) (Private codec option)

func (e *AvencH263) GetQuantizerNoiseShaping() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quantizer-noise-shaping")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetQuantizerNoiseShaping(value int) error {
	return e.Element.SetProperty("quantizer-noise-shaping", value)
}

// rc-buf-aggressivity (float32)
//
// currently useless (Private codec option)

func (e *AvencH263) GetRcBufAggressivity() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rc-buf-aggressivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencH263) SetRcBufAggressivity(value float32) error {
	return e.Element.SetProperty("rc-buf-aggressivity", value)
}

// rc-eq (string)
//
// Set rate control equation. When computing the expression, besides the standard functions defined in the section 'Expression Evaluation', the following functions are available: bits2qp(bits), qp2bits(qp). Also the following constants are available: iTex pTex tex mv fCode iCount mcVar var isI isP isB avgQP qComp avgIITex avgPITex avgPPTex avgBPTex avgTex. (Private codec option)

func (e *AvencH263) GetRcEq() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("rc-eq")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *AvencH263) SetRcEq(value string) error {
	return e.Element.SetProperty("rc-eq", value)
}

// rc-init-cplx (float32)
//
// initial complexity for 1-pass encoding (Private codec option)

func (e *AvencH263) GetRcInitCplx() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rc-init-cplx")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencH263) SetRcInitCplx(value float32) error {
	return e.Element.SetProperty("rc-init-cplx", value)
}

// rc-init-occupancy (int)
//
// number of bits which should be loaded into the rc buffer before decoding starts (Generic codec option, might have no effect)

func (e *AvencH263) GetRcInitOccupancy() (int, error) {
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

func (e *AvencH263) SetRcInitOccupancy(value int) error {
	return e.Element.SetProperty("rc-init-occupancy", value)
}

// rc-max-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetRcMaxVbvUse() (float32, error) {
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

func (e *AvencH263) SetRcMaxVbvUse(value float32) error {
	return e.Element.SetProperty("rc-max-vbv-use", value)
}

// rc-min-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetRcMinVbvUse() (float32, error) {
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

func (e *AvencH263) SetRcMinVbvUse(value float32) error {
	return e.Element.SetProperty("rc-min-vbv-use", value)
}

// rc-qmod-amp (float32)
//
// experimental quantizer modulation (Private codec option)

func (e *AvencH263) GetRcQmodAmp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rc-qmod-amp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencH263) SetRcQmodAmp(value float32) error {
	return e.Element.SetProperty("rc-qmod-amp", value)
}

// rc-qmod-freq (int)
//
// experimental quantizer modulation (Private codec option)

func (e *AvencH263) GetRcQmodFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rc-qmod-freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencH263) SetRcQmodFreq(value int) error {
	return e.Element.SetProperty("rc-qmod-freq", value)
}

// rc-strategy (GstH263EncoderRcStrategy)
//
// ratecontrol method (Private codec option)

func (e *AvencH263) GetRcStrategy() (interface{}, error) {
	return e.Element.GetProperty("rc-strategy")
}

func (e *AvencH263) SetRcStrategy(value interface{}) error {
	return e.Element.SetProperty("rc-strategy", value)
}

// refs (int)
//
// reference frames to consider for motion compensation (Generic codec option, might have no effect)

func (e *AvencH263) GetRefs() (int, error) {
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

func (e *AvencH263) SetRefs(value int) error {
	return e.Element.SetProperty("refs", value)
}

// sc-threshold (int)
//
// Scene change threshold (Private codec option)

func (e *AvencH263) GetScThreshold() (int, error) {
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

func (e *AvencH263) SetScThreshold(value int) error {
	return e.Element.SetProperty("sc-threshold", value)
}

// scplx-mask (float32)
//
// spatial complexity masking (Generic codec option, might have no effect)

func (e *AvencH263) GetScplxMask() (float32, error) {
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

func (e *AvencH263) SetScplxMask(value float32) error {
	return e.Element.SetProperty("scplx-mask", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetSideDataOnlyPackets() (bool, error) {
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

func (e *AvencH263) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// skip-cmp (GstH263EncoderCmpFunc)
//
// Frame skip compare function (Private codec option)

func (e *AvencH263) GetSkipCmp() (interface{}, error) {
	return e.Element.GetProperty("skip-cmp")
}

func (e *AvencH263) SetSkipCmp(value interface{}) error {
	return e.Element.SetProperty("skip-cmp", value)
}

// skip-exp (int)
//
// Frame skip exponent (Private codec option)

func (e *AvencH263) GetSkipExp() (int, error) {
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

func (e *AvencH263) SetSkipExp(value int) error {
	return e.Element.SetProperty("skip-exp", value)
}

// skip-factor (int)
//
// Frame skip factor (Private codec option)

func (e *AvencH263) GetSkipFactor() (int, error) {
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

func (e *AvencH263) SetSkipFactor(value int) error {
	return e.Element.SetProperty("skip-factor", value)
}

// skip-threshold (int)
//
// Frame skip threshold (Private codec option)

func (e *AvencH263) GetSkipThreshold() (int, error) {
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

func (e *AvencH263) SetSkipThreshold(value int) error {
	return e.Element.SetProperty("skip-threshold", value)
}

// skipcmp (GstAvcodeccontextCmpFunc)
//
// frame skip compare function (Generic codec option, might have no effect)

func (e *AvencH263) GetSkipcmp() (interface{}, error) {
	return e.Element.GetProperty("skipcmp")
}

func (e *AvencH263) SetSkipcmp(value interface{}) error {
	return e.Element.SetProperty("skipcmp", value)
}

// slices (int)
//
// set the number of slices, used in parallelized encoding (Generic codec option, might have no effect)

func (e *AvencH263) GetSlices() (int, error) {
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

func (e *AvencH263) SetSlices(value int) error {
	return e.Element.SetProperty("slices", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencH263) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencH263) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// subcmp (GstAvcodeccontextCmpFunc)
//
// sub-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencH263) GetSubcmp() (interface{}, error) {
	return e.Element.GetProperty("subcmp")
}

func (e *AvencH263) SetSubcmp(value interface{}) error {
	return e.Element.SetProperty("subcmp", value)
}

// subq (int)
//
// sub-pel motion estimation quality (Generic codec option, might have no effect)

func (e *AvencH263) GetSubq() (int, error) {
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

func (e *AvencH263) SetSubq(value int) error {
	return e.Element.SetProperty("subq", value)
}

// tcplx-mask (float32)
//
// temporal complexity masking (Generic codec option, might have no effect)

func (e *AvencH263) GetTcplxMask() (float32, error) {
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

func (e *AvencH263) SetTcplxMask(value float32) error {
	return e.Element.SetProperty("tcplx-mask", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencH263) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencH263) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencH263) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencH263) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencH263) GetTicksPerFrame() (int, error) {
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

func (e *AvencH263) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// timecode-frame-start (int64)
//
// GOP timecode frame start number, in non-drop-frame format (Generic codec option, might have no effect)

func (e *AvencH263) GetTimecodeFrameStart() (int64, error) {
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

func (e *AvencH263) SetTimecodeFrameStart(value int64) error {
	return e.Element.SetProperty("timecode-frame-start", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencH263) GetTrellis() (int, error) {
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

func (e *AvencH263) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type H263EncoderMpvFlags string

const (
	H263EncoderMpvFlagsSkipRd H263EncoderMpvFlags = "skip_rd" // skip_rd (0x00000001) – RD optimal MB level residual skipping
	H263EncoderMpvFlagsStrictGop H263EncoderMpvFlags = "strict_gop" // strict_gop (0x00000002) – Strictly enforce gop size
	H263EncoderMpvFlagsQpRd H263EncoderMpvFlags = "qp_rd" // qp_rd (0x00000004) – Use rate distortion optimization for qp selection
	H263EncoderMpvFlagsCbpRd H263EncoderMpvFlags = "cbp_rd" // cbp_rd (0x00000008) – use rate distortion optimization for CBP
	H263EncoderMpvFlagsNaq H263EncoderMpvFlags = "naq" // naq (0x00000010) – normalize adaptive quantization
	H263EncoderMpvFlagsMv0 H263EncoderMpvFlags = "mv0" // mv0 (0x00000020) – always try a mb with mv=<0,0>
)

type H263EncoderRcStrategy string

const (
	H263EncoderRcStrategyFfmpeg H263EncoderRcStrategy = "ffmpeg" // ffmpeg (0) – deprecated, does nothing
)

type H263EncoderCmpFunc string

const (
	H263EncoderCmpFuncSad H263EncoderCmpFunc = "sad" // sad (0) – Sum of absolute differences, fast
	H263EncoderCmpFuncSse H263EncoderCmpFunc = "sse" // sse (1) – Sum of squared errors
	H263EncoderCmpFuncSatd H263EncoderCmpFunc = "satd" // satd (2) – Sum of absolute Hadamard transformed differences
	H263EncoderCmpFuncDct H263EncoderCmpFunc = "dct" // dct (3) – Sum of absolute DCT transformed differences
	H263EncoderCmpFuncPsnr H263EncoderCmpFunc = "psnr" // psnr (4) – Sum of squared quantization errors, low quality
	H263EncoderCmpFuncBit H263EncoderCmpFunc = "bit" // bit (5) – Number of bits needed for the block
	H263EncoderCmpFuncRd H263EncoderCmpFunc = "rd" // rd (6) – Rate distortion optimal, slow
	H263EncoderCmpFuncZero H263EncoderCmpFunc = "zero" // zero (7) – Zero
	H263EncoderCmpFuncVsad H263EncoderCmpFunc = "vsad" // vsad (8) – Sum of absolute vertical differences
	H263EncoderCmpFuncVsse H263EncoderCmpFunc = "vsse" // vsse (9) – Sum of squared vertical differences
	H263EncoderCmpFuncNsse H263EncoderCmpFunc = "nsse" // nsse (10) – Noise preserving sum of squared differences
	H263EncoderCmpFuncDctmax H263EncoderCmpFunc = "dctmax" // dctmax (13) – dctmax
	H263EncoderCmpFuncDct264 H263EncoderCmpFunc = "dct264" // dct264 (14) – dct264
	H263EncoderCmpFuncMsad H263EncoderCmpFunc = "msad" // msad (15) – Sum of absolute differences, median predicted
	H263EncoderCmpFuncChroma H263EncoderCmpFunc = "chroma" // chroma (256) – chroma
)

type H263EncoderMotionEst string

const (
	H263EncoderMotionEstZero H263EncoderMotionEst = "zero" // zero (0) – zero
	H263EncoderMotionEstEpzs H263EncoderMotionEst = "epzs" // epzs (1) – epzs
	H263EncoderMotionEstXone H263EncoderMotionEst = "xone" // xone (2) – xone
)

