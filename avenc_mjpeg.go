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

type AvencMjpeg struct {
	Element *gst.Element
}

func NewAvencMjpeg() (*AvencMjpeg, error) {
	e, err := gst.NewElement("avenc_mjpeg")
	if err != nil {
		return nil, err
	}
	return &AvencMjpeg{Element: e}, nil
}

func NewAvencMjpegWithName(name string) (*AvencMjpeg, error) {
	e, err := gst.NewElementWithName("avenc_mjpeg", name)
	if err != nil {
		return nil, err
	}
	return &AvencMjpeg{Element: e}, nil
}

// ----- Properties -----

// a53cc (bool)
//
// Use A53 Closed Captions (if available) (Private codec option)

func (e *AvencMjpeg) GetA53Cc() (bool, error) {
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

func (e *AvencMjpeg) SetA53Cc(value bool) error {
	return e.Element.SetProperty("a53cc", value)
}

// b-qfactor (float32)
//
// QP factor between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetBQfactor() (float32, error) {
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

func (e *AvencMjpeg) SetBQfactor(value float32) error {
	return e.Element.SetProperty("b-qfactor", value)
}

// b-qoffset (float32)
//
// QP offset between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetBQoffset() (float32, error) {
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

func (e *AvencMjpeg) SetBQoffset(value float32) error {
	return e.Element.SetProperty("b-qoffset", value)
}

// b-sensitivity (int)
//
// Adjust sensitivity of b_frame_strategy 1 (Private codec option)

func (e *AvencMjpeg) GetBSensitivity() (int, error) {
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

func (e *AvencMjpeg) SetBSensitivity(value int) error {
	return e.Element.SetProperty("b-sensitivity", value)
}

// b-strategy (int)
//
// Strategy to choose between I/P/B-frames (Private codec option)

func (e *AvencMjpeg) GetBStrategy() (int, error) {
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

func (e *AvencMjpeg) SetBStrategy(value int) error {
	return e.Element.SetProperty("b-strategy", value)
}

// bidir-refine (int)
//
// refine the two motion vectors used in bidirectional macroblocks (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetBidirRefine() (int, error) {
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

func (e *AvencMjpeg) SetBidirRefine(value int) error {
	return e.Element.SetProperty("bidir-refine", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetBitrate() (int, error) {
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

func (e *AvencMjpeg) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-tolerance (int)
//
// Set video bitrate tolerance (in bits/s). In 1-pass mode, bitrate tolerance specifies how far ratecontrol is willing to deviate from the target average bitrate value. This is not related to minimum/maximum bitrate. Lowering tolerance too much has an adverse effect on quality. (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetBitrateTolerance() (int, error) {
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

func (e *AvencMjpeg) SetBitrateTolerance(value int) error {
	return e.Element.SetProperty("bitrate-tolerance", value)
}

// border-mask (float32)
//
// increase the quantizer for macroblocks close to borders (Private codec option)

func (e *AvencMjpeg) GetBorderMask() (float32, error) {
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

func (e *AvencMjpeg) SetBorderMask(value float32) error {
	return e.Element.SetProperty("border-mask", value)
}

// brd-scale (int)
//
// Downscale frames for dynamic B-frame decision (Private codec option)

func (e *AvencMjpeg) GetBrdScale() (int, error) {
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

func (e *AvencMjpeg) SetBrdScale(value int) error {
	return e.Element.SetProperty("brd-scale", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetBufsize() (int, error) {
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

func (e *AvencMjpeg) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// chroma-elim-threshold (int)
//
// single coefficient elimination threshold for chrominance (negative values also consider dc coefficient) (Private codec option)

func (e *AvencMjpeg) GetChromaElimThreshold() (int, error) {
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

func (e *AvencMjpeg) SetChromaElimThreshold(value int) error {
	return e.Element.SetProperty("chroma-elim-threshold", value)
}

// chroma-sample-location (GstAvcodeccontextChromaSampleLocationType)
//
// chroma sample location (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetChromaSampleLocation() (interface{}, error) {
	return e.Element.GetProperty("chroma-sample-location")
}

func (e *AvencMjpeg) SetChromaSampleLocation(value interface{}) error {
	return e.Element.SetProperty("chroma-sample-location", value)
}

// chromaoffset (int)
//
// chroma QP offset from luma (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetChromaoffset() (int, error) {
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

func (e *AvencMjpeg) SetChromaoffset(value int) error {
	return e.Element.SetProperty("chromaoffset", value)
}

// cmp (GstAvcodeccontextCmpFunc)
//
// full-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetCmp() (interface{}, error) {
	return e.Element.GetProperty("cmp")
}

func (e *AvencMjpeg) SetCmp(value interface{}) error {
	return e.Element.SetProperty("cmp", value)
}

// coder (GstAvcodeccontextCoder)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetCoder() (interface{}, error) {
	return e.Element.GetProperty("coder")
}

func (e *AvencMjpeg) SetCoder(value interface{}) error {
	return e.Element.SetProperty("coder", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetCompressionLevel() (int, error) {
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

func (e *AvencMjpeg) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// context (int)
//
// context model (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetContext() (int, error) {
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

func (e *AvencMjpeg) SetContext(value int) error {
	return e.Element.SetProperty("context", value)
}

// dark-mask (float32)
//
// compresses dark areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetDarkMask() (float32, error) {
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

func (e *AvencMjpeg) SetDarkMask(value float32) error {
	return e.Element.SetProperty("dark-mask", value)
}

// dc (int)
//
// intra_dc_precision (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetDc() (int, error) {
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

func (e *AvencMjpeg) SetDc(value int) error {
	return e.Element.SetProperty("dc", value)
}

// dct (GstAvcodeccontextDct)
//
// DCT algorithm (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetDct() (interface{}, error) {
	return e.Element.GetProperty("dct")
}

func (e *AvencMjpeg) SetDct(value interface{}) error {
	return e.Element.SetProperty("dct", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencMjpeg) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dia-size (int)
//
// diamond type & size for motion estimation (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetDiaSize() (int, error) {
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

func (e *AvencMjpeg) SetDiaSize(value int) error {
	return e.Element.SetProperty("dia-size", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetDumpSeparator() (string, error) {
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

func (e *AvencMjpeg) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencMjpeg) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// error-rate (int)
//
// Simulate errors in the bitstream to test error concealment. (Private codec option)

func (e *AvencMjpeg) GetErrorRate() (int, error) {
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

func (e *AvencMjpeg) SetErrorRate(value int) error {
	return e.Element.SetProperty("error-rate", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencMjpeg) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// field-order (GstAvcodeccontextFieldOrder)
//
// Field order (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("field-order")
}

func (e *AvencMjpeg) SetFieldOrder(value interface{}) error {
	return e.Element.SetProperty("field-order", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencMjpeg) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencMjpeg) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// force-duplicated-matrix (bool)
//
// Always write luma and chroma matrix for mjpeg, useful for rtp streaming. (Private codec option)

func (e *AvencMjpeg) GetForceDuplicatedMatrix() (bool, error) {
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

func (e *AvencMjpeg) SetForceDuplicatedMatrix(value bool) error {
	return e.Element.SetProperty("force-duplicated-matrix", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetGlobalQuality() (int, error) {
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

func (e *AvencMjpeg) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// gop-size (int)
//
// set the group of picture (GOP) size (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetGopSize() (int, error) {
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

func (e *AvencMjpeg) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// huffman (GstMjpegEncoderHuffman)
//
// Huffman table strategy (Private codec option)

func (e *AvencMjpeg) GetHuffman() (interface{}, error) {
	return e.Element.GetProperty("huffman")
}

func (e *AvencMjpeg) SetHuffman(value interface{}) error {
	return e.Element.SetProperty("huffman", value)
}

// i-qfactor (float32)
//
// QP factor between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetIQfactor() (float32, error) {
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

func (e *AvencMjpeg) SetIQfactor(value float32) error {
	return e.Element.SetProperty("i-qfactor", value)
}

// i-qoffset (float32)
//
// QP offset between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetIQoffset() (float32, error) {
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

func (e *AvencMjpeg) SetIQoffset(value float32) error {
	return e.Element.SetProperty("i-qoffset", value)
}

// ibias (int)
//
// intra quant bias (Private codec option)

func (e *AvencMjpeg) GetIbias() (int, error) {
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

func (e *AvencMjpeg) SetIbias(value int) error {
	return e.Element.SetProperty("ibias", value)
}

// idct (GstAvcodeccontextIdct)
//
// select IDCT implementation (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetIdct() (interface{}, error) {
	return e.Element.GetProperty("idct")
}

func (e *AvencMjpeg) SetIdct(value interface{}) error {
	return e.Element.SetProperty("idct", value)
}

// ildctcmp (GstAvcodeccontextCmpFunc)
//
// interlaced DCT compare function (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetIldctcmp() (interface{}, error) {
	return e.Element.GetProperty("ildctcmp")
}

func (e *AvencMjpeg) SetIldctcmp(value interface{}) error {
	return e.Element.SetProperty("ildctcmp", value)
}

// intra-penalty (int)
//
// Penalty for intra blocks in block decision (Private codec option)

func (e *AvencMjpeg) GetIntraPenalty() (int, error) {
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

func (e *AvencMjpeg) SetIntraPenalty(value int) error {
	return e.Element.SetProperty("intra-penalty", value)
}

// keyint-min (int)
//
// minimum interval between IDR-frames (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetKeyintMin() (int, error) {
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

func (e *AvencMjpeg) SetKeyintMin(value int) error {
	return e.Element.SetProperty("keyint-min", value)
}

// last-pred (int)
//
// amount of motion predictors from the previous frame (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetLastPred() (int, error) {
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

func (e *AvencMjpeg) SetLastPred(value int) error {
	return e.Element.SetProperty("last-pred", value)
}

// lmax (int)
//
// maximum Lagrange factor (VBR) (Private codec option)

func (e *AvencMjpeg) GetLmax() (int, error) {
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

func (e *AvencMjpeg) SetLmax(value int) error {
	return e.Element.SetProperty("lmax", value)
}

// lmin (int)
//
// minimum Lagrange factor (VBR) (Private codec option)

func (e *AvencMjpeg) GetLmin() (int, error) {
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

func (e *AvencMjpeg) SetLmin(value int) error {
	return e.Element.SetProperty("lmin", value)
}

// luma-elim-threshold (int)
//
// single coefficient elimination threshold for luminance (negative values also consider dc coefficient) (Private codec option)

func (e *AvencMjpeg) GetLumaElimThreshold() (int, error) {
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

func (e *AvencMjpeg) SetLumaElimThreshold(value int) error {
	return e.Element.SetProperty("luma-elim-threshold", value)
}

// lumi-mask (float32)
//
// compresses bright areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetLumiMask() (float32, error) {
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

func (e *AvencMjpeg) SetLumiMask(value float32) error {
	return e.Element.SetProperty("lumi-mask", value)
}

// max-bframes (int)
//
// set maximum number of B-frames between non-B-frames (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMaxBframes() (int, error) {
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

func (e *AvencMjpeg) SetMaxBframes(value int) error {
	return e.Element.SetProperty("max-bframes", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMaxPixels() (int64, error) {
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

func (e *AvencMjpeg) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMaxrate() (int64, error) {
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

func (e *AvencMjpeg) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// mbcmp (GstAvcodeccontextCmpFunc)
//
// macroblock compare function (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMbcmp() (interface{}, error) {
	return e.Element.GetProperty("mbcmp")
}

func (e *AvencMjpeg) SetMbcmp(value interface{}) error {
	return e.Element.SetProperty("mbcmp", value)
}

// mbd (GstAvcodeccontextMbd)
//
// macroblock decision algorithm (high quality mode) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMbd() (interface{}, error) {
	return e.Element.GetProperty("mbd")
}

func (e *AvencMjpeg) SetMbd(value interface{}) error {
	return e.Element.SetProperty("mbd", value)
}

// mblmax (int)
//
// maximum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMblmax() (int, error) {
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

func (e *AvencMjpeg) SetMblmax(value int) error {
	return e.Element.SetProperty("mblmax", value)
}

// mblmin (int)
//
// minimum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMblmin() (int, error) {
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

func (e *AvencMjpeg) SetMblmin(value int) error {
	return e.Element.SetProperty("mblmin", value)
}

// me-range (int)
//
// limit motion vectors range (1023 for DivX player) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMeRange() (int, error) {
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

func (e *AvencMjpeg) SetMeRange(value int) error {
	return e.Element.SetProperty("me-range", value)
}

// mepc (int)
//
// Motion estimation bitrate penalty compensation (1.0 = 256) (Private codec option)

func (e *AvencMjpeg) GetMepc() (int, error) {
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

func (e *AvencMjpeg) SetMepc(value int) error {
	return e.Element.SetProperty("mepc", value)
}

// mepre (int)
//
// pre motion estimation (Private codec option)

func (e *AvencMjpeg) GetMepre() (int, error) {
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

func (e *AvencMjpeg) SetMepre(value int) error {
	return e.Element.SetProperty("mepre", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMinrate() (int64, error) {
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

func (e *AvencMjpeg) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// motion-est (GstMjpegEncoderMotionEst)
//
// motion estimation algorithm (Private codec option)

func (e *AvencMjpeg) GetMotionEst() (interface{}, error) {
	return e.Element.GetProperty("motion-est")
}

func (e *AvencMjpeg) SetMotionEst(value interface{}) error {
	return e.Element.SetProperty("motion-est", value)
}

// mpeg-quant (int)
//
// Use MPEG quantizers instead of H.263 (Private codec option)

func (e *AvencMjpeg) GetMpegQuant() (int, error) {
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

func (e *AvencMjpeg) SetMpegQuant(value int) error {
	return e.Element.SetProperty("mpeg-quant", value)
}

// mpv-flags (GstMjpegEncoderMpvFlags)
//
// Flags common for all mpegvideo-based encoders. (Private codec option)

func (e *AvencMjpeg) GetMpvFlags() (interface{}, error) {
	return e.Element.GetProperty("mpv-flags")
}

func (e *AvencMjpeg) SetMpvFlags(value interface{}) error {
	return e.Element.SetProperty("mpv-flags", value)
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *AvencMjpeg) GetMultipassCacheFile() (string, error) {
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

func (e *AvencMjpeg) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// mv0-threshold (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetMv0Threshold() (int, error) {
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

func (e *AvencMjpeg) SetMv0Threshold(value int) error {
	return e.Element.SetProperty("mv0-threshold", value)
}

// noise-reduction (int)
//
// Noise reduction (Private codec option)

func (e *AvencMjpeg) GetNoiseReduction() (int, error) {
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

func (e *AvencMjpeg) SetNoiseReduction(value int) error {
	return e.Element.SetProperty("noise-reduction", value)
}

// nr (int)
//
// noise reduction (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetNr() (int, error) {
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

func (e *AvencMjpeg) SetNr(value int) error {
	return e.Element.SetProperty("nr", value)
}

// nssew (int)
//
// nsse weight (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetNssew() (int, error) {
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

func (e *AvencMjpeg) SetNssew(value int) error {
	return e.Element.SetProperty("nssew", value)
}

// p-mask (float32)
//
// inter masking (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetPMask() (float32, error) {
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

func (e *AvencMjpeg) SetPMask(value float32) error {
	return e.Element.SetProperty("p-mask", value)
}

// pass (GstLibAVEncPass)
//
// Encoding pass/type

func (e *AvencMjpeg) GetPass() (interface{}, error) {
	return e.Element.GetProperty("pass")
}

func (e *AvencMjpeg) SetPass(value interface{}) error {
	return e.Element.SetProperty("pass", value)
}

// pbias (int)
//
// inter quant bias (Private codec option)

func (e *AvencMjpeg) GetPbias() (int, error) {
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

func (e *AvencMjpeg) SetPbias(value int) error {
	return e.Element.SetProperty("pbias", value)
}

// pre-dia-size (int)
//
// diamond type & size for motion estimation pre-pass (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetPreDiaSize() (int, error) {
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

func (e *AvencMjpeg) SetPreDiaSize(value int) error {
	return e.Element.SetProperty("pre-dia-size", value)
}

// precmp (GstAvcodeccontextCmpFunc)
//
// pre motion estimation compare function (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetPrecmp() (interface{}, error) {
	return e.Element.GetProperty("precmp")
}

func (e *AvencMjpeg) SetPrecmp(value interface{}) error {
	return e.Element.SetProperty("precmp", value)
}

// pred (GstMjpegEncoderPred)
//
// Prediction method (Private codec option)

func (e *AvencMjpeg) GetPred() (interface{}, error) {
	return e.Element.GetProperty("pred")
}

func (e *AvencMjpeg) SetPred(value interface{}) error {
	return e.Element.SetProperty("pred", value)
}

// preme (int)
//
// pre motion estimation (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetPreme() (int, error) {
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

func (e *AvencMjpeg) SetPreme(value int) error {
	return e.Element.SetProperty("preme", value)
}

// ps (int)
//
// RTP payload size in bytes (Private codec option)

func (e *AvencMjpeg) GetPs() (int, error) {
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

func (e *AvencMjpeg) SetPs(value int) error {
	return e.Element.SetProperty("ps", value)
}

// qblur (float32)
//
// video quantizer scale blur (VBR) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetQblur() (float32, error) {
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

func (e *AvencMjpeg) SetQblur(value float32) error {
	return e.Element.SetProperty("qblur", value)
}

// qcomp (float32)
//
// video quantizer scale compression (VBR). Constant of ratecontrol equation. Recommended range for default rc_eq: 0.0-1.0 (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetQcomp() (float32, error) {
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

func (e *AvencMjpeg) SetQcomp(value float32) error {
	return e.Element.SetProperty("qcomp", value)
}

// qdiff (int)
//
// maximum difference between the quantizer scales (VBR) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetQdiff() (int, error) {
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

func (e *AvencMjpeg) SetQdiff(value int) error {
	return e.Element.SetProperty("qdiff", value)
}

// qmax (int)
//
// maximum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetQmax() (int, error) {
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

func (e *AvencMjpeg) SetQmax(value int) error {
	return e.Element.SetProperty("qmax", value)
}

// qmin (int)
//
// minimum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetQmin() (int, error) {
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

func (e *AvencMjpeg) SetQmin(value int) error {
	return e.Element.SetProperty("qmin", value)
}

// qsquish (float32)
//
// how to keep quantizer between qmin and qmax (0 = clip, 1 = use differentiable function) (Private codec option)

func (e *AvencMjpeg) GetQsquish() (float32, error) {
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

func (e *AvencMjpeg) SetQsquish(value float32) error {
	return e.Element.SetProperty("qsquish", value)
}

// quantizer (float32)
//
// Constant Quantizer

func (e *AvencMjpeg) GetQuantizer() (float32, error) {
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

func (e *AvencMjpeg) SetQuantizer(value float32) error {
	return e.Element.SetProperty("quantizer", value)
}

// quantizer-noise-shaping (int)
//
// (null) (Private codec option)

func (e *AvencMjpeg) GetQuantizerNoiseShaping() (int, error) {
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

func (e *AvencMjpeg) SetQuantizerNoiseShaping(value int) error {
	return e.Element.SetProperty("quantizer-noise-shaping", value)
}

// rc-buf-aggressivity (float32)
//
// currently useless (Private codec option)

func (e *AvencMjpeg) GetRcBufAggressivity() (float32, error) {
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

func (e *AvencMjpeg) SetRcBufAggressivity(value float32) error {
	return e.Element.SetProperty("rc-buf-aggressivity", value)
}

// rc-eq (string)
//
// Set rate control equation. When computing the expression, besides the standard functions defined in the section 'Expression Evaluation', the following functions are available: bits2qp(bits), qp2bits(qp). Also the following constants are available: iTex pTex tex mv fCode iCount mcVar var isI isP isB avgQP qComp avgIITex avgPITex avgPPTex avgBPTex avgTex. (Private codec option)

func (e *AvencMjpeg) GetRcEq() (string, error) {
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

func (e *AvencMjpeg) SetRcEq(value string) error {
	return e.Element.SetProperty("rc-eq", value)
}

// rc-init-cplx (float32)
//
// initial complexity for 1-pass encoding (Private codec option)

func (e *AvencMjpeg) GetRcInitCplx() (float32, error) {
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

func (e *AvencMjpeg) SetRcInitCplx(value float32) error {
	return e.Element.SetProperty("rc-init-cplx", value)
}

// rc-init-occupancy (int)
//
// number of bits which should be loaded into the rc buffer before decoding starts (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetRcInitOccupancy() (int, error) {
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

func (e *AvencMjpeg) SetRcInitOccupancy(value int) error {
	return e.Element.SetProperty("rc-init-occupancy", value)
}

// rc-max-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetRcMaxVbvUse() (float32, error) {
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

func (e *AvencMjpeg) SetRcMaxVbvUse(value float32) error {
	return e.Element.SetProperty("rc-max-vbv-use", value)
}

// rc-min-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetRcMinVbvUse() (float32, error) {
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

func (e *AvencMjpeg) SetRcMinVbvUse(value float32) error {
	return e.Element.SetProperty("rc-min-vbv-use", value)
}

// rc-qmod-amp (float32)
//
// experimental quantizer modulation (Private codec option)

func (e *AvencMjpeg) GetRcQmodAmp() (float32, error) {
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

func (e *AvencMjpeg) SetRcQmodAmp(value float32) error {
	return e.Element.SetProperty("rc-qmod-amp", value)
}

// rc-qmod-freq (int)
//
// experimental quantizer modulation (Private codec option)

func (e *AvencMjpeg) GetRcQmodFreq() (int, error) {
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

func (e *AvencMjpeg) SetRcQmodFreq(value int) error {
	return e.Element.SetProperty("rc-qmod-freq", value)
}

// rc-strategy (GstMjpegEncoderRcStrategy)
//
// ratecontrol method (Private codec option)

func (e *AvencMjpeg) GetRcStrategy() (interface{}, error) {
	return e.Element.GetProperty("rc-strategy")
}

func (e *AvencMjpeg) SetRcStrategy(value interface{}) error {
	return e.Element.SetProperty("rc-strategy", value)
}

// refs (int)
//
// reference frames to consider for motion compensation (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetRefs() (int, error) {
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

func (e *AvencMjpeg) SetRefs(value int) error {
	return e.Element.SetProperty("refs", value)
}

// sc-threshold (int)
//
// Scene change threshold (Private codec option)

func (e *AvencMjpeg) GetScThreshold() (int, error) {
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

func (e *AvencMjpeg) SetScThreshold(value int) error {
	return e.Element.SetProperty("sc-threshold", value)
}

// scplx-mask (float32)
//
// spatial complexity masking (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetScplxMask() (float32, error) {
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

func (e *AvencMjpeg) SetScplxMask(value float32) error {
	return e.Element.SetProperty("scplx-mask", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetSideDataOnlyPackets() (bool, error) {
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

func (e *AvencMjpeg) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// skip-cmp (GstMjpegEncoderCmpFunc)
//
// Frame skip compare function (Private codec option)

func (e *AvencMjpeg) GetSkipCmp() (interface{}, error) {
	return e.Element.GetProperty("skip-cmp")
}

func (e *AvencMjpeg) SetSkipCmp(value interface{}) error {
	return e.Element.SetProperty("skip-cmp", value)
}

// skip-exp (int)
//
// Frame skip exponent (Private codec option)

func (e *AvencMjpeg) GetSkipExp() (int, error) {
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

func (e *AvencMjpeg) SetSkipExp(value int) error {
	return e.Element.SetProperty("skip-exp", value)
}

// skip-factor (int)
//
// Frame skip factor (Private codec option)

func (e *AvencMjpeg) GetSkipFactor() (int, error) {
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

func (e *AvencMjpeg) SetSkipFactor(value int) error {
	return e.Element.SetProperty("skip-factor", value)
}

// skip-threshold (int)
//
// Frame skip threshold (Private codec option)

func (e *AvencMjpeg) GetSkipThreshold() (int, error) {
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

func (e *AvencMjpeg) SetSkipThreshold(value int) error {
	return e.Element.SetProperty("skip-threshold", value)
}

// skipcmp (GstAvcodeccontextCmpFunc)
//
// frame skip compare function (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetSkipcmp() (interface{}, error) {
	return e.Element.GetProperty("skipcmp")
}

func (e *AvencMjpeg) SetSkipcmp(value interface{}) error {
	return e.Element.SetProperty("skipcmp", value)
}

// slices (int)
//
// set the number of slices, used in parallelized encoding (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetSlices() (int, error) {
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

func (e *AvencMjpeg) SetSlices(value int) error {
	return e.Element.SetProperty("slices", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencMjpeg) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// subcmp (GstAvcodeccontextCmpFunc)
//
// sub-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetSubcmp() (interface{}, error) {
	return e.Element.GetProperty("subcmp")
}

func (e *AvencMjpeg) SetSubcmp(value interface{}) error {
	return e.Element.SetProperty("subcmp", value)
}

// subq (int)
//
// sub-pel motion estimation quality (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetSubq() (int, error) {
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

func (e *AvencMjpeg) SetSubq(value int) error {
	return e.Element.SetProperty("subq", value)
}

// tcplx-mask (float32)
//
// temporal complexity masking (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetTcplxMask() (float32, error) {
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

func (e *AvencMjpeg) SetTcplxMask(value float32) error {
	return e.Element.SetProperty("tcplx-mask", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencMjpeg) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencMjpeg) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetTicksPerFrame() (int, error) {
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

func (e *AvencMjpeg) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// timecode-frame-start (int64)
//
// GOP timecode frame start number, in non-drop-frame format (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetTimecodeFrameStart() (int64, error) {
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

func (e *AvencMjpeg) SetTimecodeFrameStart(value int64) error {
	return e.Element.SetProperty("timecode-frame-start", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencMjpeg) GetTrellis() (int, error) {
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

func (e *AvencMjpeg) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type MjpegEncoderMotionEst string

const (
	MjpegEncoderMotionEstZero MjpegEncoderMotionEst = "zero" // zero (0) – zero
	MjpegEncoderMotionEstEpzs MjpegEncoderMotionEst = "epzs" // epzs (1) – epzs
	MjpegEncoderMotionEstXone MjpegEncoderMotionEst = "xone" // xone (2) – xone
)

type MjpegEncoderMpvFlags string

const (
	MjpegEncoderMpvFlagsSkipRd MjpegEncoderMpvFlags = "skip_rd" // skip_rd (0x00000001) – RD optimal MB level residual skipping
	MjpegEncoderMpvFlagsStrictGop MjpegEncoderMpvFlags = "strict_gop" // strict_gop (0x00000002) – Strictly enforce gop size
	MjpegEncoderMpvFlagsQpRd MjpegEncoderMpvFlags = "qp_rd" // qp_rd (0x00000004) – Use rate distortion optimization for qp selection
	MjpegEncoderMpvFlagsCbpRd MjpegEncoderMpvFlags = "cbp_rd" // cbp_rd (0x00000008) – use rate distortion optimization for CBP
	MjpegEncoderMpvFlagsNaq MjpegEncoderMpvFlags = "naq" // naq (0x00000010) – normalize adaptive quantization
	MjpegEncoderMpvFlagsMv0 MjpegEncoderMpvFlags = "mv0" // mv0 (0x00000020) – always try a mb with mv=<0,0>
)

type MjpegEncoderPred string

const (
	MjpegEncoderPredLeft MjpegEncoderPred = "left" // left (1) – left
	MjpegEncoderPredPlane MjpegEncoderPred = "plane" // plane (2) – plane
	MjpegEncoderPredMedian MjpegEncoderPred = "median" // median (3) – median
)

type MjpegEncoderRcStrategy string

const (
	MjpegEncoderRcStrategyFfmpeg MjpegEncoderRcStrategy = "ffmpeg" // ffmpeg (0) – deprecated, does nothing
)

type MjpegEncoderCmpFunc string

const (
	MjpegEncoderCmpFuncSad MjpegEncoderCmpFunc = "sad" // sad (0) – Sum of absolute differences, fast
	MjpegEncoderCmpFuncSse MjpegEncoderCmpFunc = "sse" // sse (1) – Sum of squared errors
	MjpegEncoderCmpFuncSatd MjpegEncoderCmpFunc = "satd" // satd (2) – Sum of absolute Hadamard transformed differences
	MjpegEncoderCmpFuncDct MjpegEncoderCmpFunc = "dct" // dct (3) – Sum of absolute DCT transformed differences
	MjpegEncoderCmpFuncPsnr MjpegEncoderCmpFunc = "psnr" // psnr (4) – Sum of squared quantization errors, low quality
	MjpegEncoderCmpFuncBit MjpegEncoderCmpFunc = "bit" // bit (5) – Number of bits needed for the block
	MjpegEncoderCmpFuncRd MjpegEncoderCmpFunc = "rd" // rd (6) – Rate distortion optimal, slow
	MjpegEncoderCmpFuncZero MjpegEncoderCmpFunc = "zero" // zero (7) – Zero
	MjpegEncoderCmpFuncVsad MjpegEncoderCmpFunc = "vsad" // vsad (8) – Sum of absolute vertical differences
	MjpegEncoderCmpFuncVsse MjpegEncoderCmpFunc = "vsse" // vsse (9) – Sum of squared vertical differences
	MjpegEncoderCmpFuncNsse MjpegEncoderCmpFunc = "nsse" // nsse (10) – Noise preserving sum of squared differences
	MjpegEncoderCmpFuncDctmax MjpegEncoderCmpFunc = "dctmax" // dctmax (13) – dctmax
	MjpegEncoderCmpFuncDct264 MjpegEncoderCmpFunc = "dct264" // dct264 (14) – dct264
	MjpegEncoderCmpFuncMsad MjpegEncoderCmpFunc = "msad" // msad (15) – Sum of absolute differences, median predicted
	MjpegEncoderCmpFuncChroma MjpegEncoderCmpFunc = "chroma" // chroma (256) – chroma
)

type MjpegEncoderHuffman string

const (
	MjpegEncoderHuffmanDefault MjpegEncoderHuffman = "default" // default (0) – default
	MjpegEncoderHuffmanOptimal MjpegEncoderHuffman = "optimal" // optimal (1) – optimal
)

