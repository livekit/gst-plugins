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

type AvencSpeedhq struct {
	Element *gst.Element
}

func NewAvencSpeedhq() (*AvencSpeedhq, error) {
	e, err := gst.NewElement("avenc_speedhq")
	if err != nil {
		return nil, err
	}
	return &AvencSpeedhq{Element: e}, nil
}

func NewAvencSpeedhqWithName(name string) (*AvencSpeedhq, error) {
	e, err := gst.NewElementWithName("avenc_speedhq", name)
	if err != nil {
		return nil, err
	}
	return &AvencSpeedhq{Element: e}, nil
}

// ----- Properties -----

// a53cc (bool)
//
// Use A53 Closed Captions (if available) (Private codec option)

func (e *AvencSpeedhq) GetA53Cc() (bool, error) {
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

func (e *AvencSpeedhq) SetA53Cc(value bool) error {
	return e.Element.SetProperty("a53cc", value)
}

// b-qfactor (float32)
//
// QP factor between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetBQfactor() (float32, error) {
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

func (e *AvencSpeedhq) SetBQfactor(value float32) error {
	return e.Element.SetProperty("b-qfactor", value)
}

// b-qoffset (float32)
//
// QP offset between P- and B-frames (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetBQoffset() (float32, error) {
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

func (e *AvencSpeedhq) SetBQoffset(value float32) error {
	return e.Element.SetProperty("b-qoffset", value)
}

// b-sensitivity (int)
//
// Adjust sensitivity of b_frame_strategy 1 (Private codec option)

func (e *AvencSpeedhq) GetBSensitivity() (int, error) {
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

func (e *AvencSpeedhq) SetBSensitivity(value int) error {
	return e.Element.SetProperty("b-sensitivity", value)
}

// b-strategy (int)
//
// Strategy to choose between I/P/B-frames (Private codec option)

func (e *AvencSpeedhq) GetBStrategy() (int, error) {
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

func (e *AvencSpeedhq) SetBStrategy(value int) error {
	return e.Element.SetProperty("b-strategy", value)
}

// bidir-refine (int)
//
// refine the two motion vectors used in bidirectional macroblocks (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetBidirRefine() (int, error) {
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

func (e *AvencSpeedhq) SetBidirRefine(value int) error {
	return e.Element.SetProperty("bidir-refine", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetBitrate() (int, error) {
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

func (e *AvencSpeedhq) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-tolerance (int)
//
// Set video bitrate tolerance (in bits/s). In 1-pass mode, bitrate tolerance specifies how far ratecontrol is willing to deviate from the target average bitrate value. This is not related to minimum/maximum bitrate. Lowering tolerance too much has an adverse effect on quality. (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetBitrateTolerance() (int, error) {
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

func (e *AvencSpeedhq) SetBitrateTolerance(value int) error {
	return e.Element.SetProperty("bitrate-tolerance", value)
}

// border-mask (float32)
//
// increase the quantizer for macroblocks close to borders (Private codec option)

func (e *AvencSpeedhq) GetBorderMask() (float32, error) {
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

func (e *AvencSpeedhq) SetBorderMask(value float32) error {
	return e.Element.SetProperty("border-mask", value)
}

// brd-scale (int)
//
// Downscale frames for dynamic B-frame decision (Private codec option)

func (e *AvencSpeedhq) GetBrdScale() (int, error) {
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

func (e *AvencSpeedhq) SetBrdScale(value int) error {
	return e.Element.SetProperty("brd-scale", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetBufsize() (int, error) {
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

func (e *AvencSpeedhq) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// chroma-elim-threshold (int)
//
// single coefficient elimination threshold for chrominance (negative values also consider dc coefficient) (Private codec option)

func (e *AvencSpeedhq) GetChromaElimThreshold() (int, error) {
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

func (e *AvencSpeedhq) SetChromaElimThreshold(value int) error {
	return e.Element.SetProperty("chroma-elim-threshold", value)
}

// chroma-sample-location (GstAvcodeccontextChromaSampleLocationType)
//
// chroma sample location (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetChromaSampleLocation() (interface{}, error) {
	return e.Element.GetProperty("chroma-sample-location")
}

func (e *AvencSpeedhq) SetChromaSampleLocation(value interface{}) error {
	return e.Element.SetProperty("chroma-sample-location", value)
}

// chromaoffset (int)
//
// chroma QP offset from luma (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetChromaoffset() (int, error) {
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

func (e *AvencSpeedhq) SetChromaoffset(value int) error {
	return e.Element.SetProperty("chromaoffset", value)
}

// cmp (GstAvcodeccontextCmpFunc)
//
// full-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetCmp() (interface{}, error) {
	return e.Element.GetProperty("cmp")
}

func (e *AvencSpeedhq) SetCmp(value interface{}) error {
	return e.Element.SetProperty("cmp", value)
}

// coder (GstAvcodeccontextCoder)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetCoder() (interface{}, error) {
	return e.Element.GetProperty("coder")
}

func (e *AvencSpeedhq) SetCoder(value interface{}) error {
	return e.Element.SetProperty("coder", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetCompressionLevel() (int, error) {
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

func (e *AvencSpeedhq) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// context (int)
//
// context model (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetContext() (int, error) {
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

func (e *AvencSpeedhq) SetContext(value int) error {
	return e.Element.SetProperty("context", value)
}

// dark-mask (float32)
//
// compresses dark areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetDarkMask() (float32, error) {
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

func (e *AvencSpeedhq) SetDarkMask(value float32) error {
	return e.Element.SetProperty("dark-mask", value)
}

// dc (int)
//
// intra_dc_precision (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetDc() (int, error) {
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

func (e *AvencSpeedhq) SetDc(value int) error {
	return e.Element.SetProperty("dc", value)
}

// dct (GstAvcodeccontextDct)
//
// DCT algorithm (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetDct() (interface{}, error) {
	return e.Element.GetProperty("dct")
}

func (e *AvencSpeedhq) SetDct(value interface{}) error {
	return e.Element.SetProperty("dct", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencSpeedhq) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dia-size (int)
//
// diamond type & size for motion estimation (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetDiaSize() (int, error) {
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

func (e *AvencSpeedhq) SetDiaSize(value int) error {
	return e.Element.SetProperty("dia-size", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetDumpSeparator() (string, error) {
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

func (e *AvencSpeedhq) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencSpeedhq) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// error-rate (int)
//
// Simulate errors in the bitstream to test error concealment. (Private codec option)

func (e *AvencSpeedhq) GetErrorRate() (int, error) {
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

func (e *AvencSpeedhq) SetErrorRate(value int) error {
	return e.Element.SetProperty("error-rate", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencSpeedhq) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// field-order (GstAvcodeccontextFieldOrder)
//
// Field order (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("field-order")
}

func (e *AvencSpeedhq) SetFieldOrder(value interface{}) error {
	return e.Element.SetProperty("field-order", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencSpeedhq) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencSpeedhq) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// force-duplicated-matrix (bool)
//
// Always write luma and chroma matrix for mjpeg, useful for rtp streaming. (Private codec option)

func (e *AvencSpeedhq) GetForceDuplicatedMatrix() (bool, error) {
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

func (e *AvencSpeedhq) SetForceDuplicatedMatrix(value bool) error {
	return e.Element.SetProperty("force-duplicated-matrix", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetGlobalQuality() (int, error) {
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

func (e *AvencSpeedhq) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// gop-size (int)
//
// set the group of picture (GOP) size (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetGopSize() (int, error) {
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

func (e *AvencSpeedhq) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// i-qfactor (float32)
//
// QP factor between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetIQfactor() (float32, error) {
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

func (e *AvencSpeedhq) SetIQfactor(value float32) error {
	return e.Element.SetProperty("i-qfactor", value)
}

// i-qoffset (float32)
//
// QP offset between P- and I-frames (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetIQoffset() (float32, error) {
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

func (e *AvencSpeedhq) SetIQoffset(value float32) error {
	return e.Element.SetProperty("i-qoffset", value)
}

// ibias (int)
//
// intra quant bias (Private codec option)

func (e *AvencSpeedhq) GetIbias() (int, error) {
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

func (e *AvencSpeedhq) SetIbias(value int) error {
	return e.Element.SetProperty("ibias", value)
}

// idct (GstAvcodeccontextIdct)
//
// select IDCT implementation (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetIdct() (interface{}, error) {
	return e.Element.GetProperty("idct")
}

func (e *AvencSpeedhq) SetIdct(value interface{}) error {
	return e.Element.SetProperty("idct", value)
}

// ildctcmp (GstAvcodeccontextCmpFunc)
//
// interlaced DCT compare function (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetIldctcmp() (interface{}, error) {
	return e.Element.GetProperty("ildctcmp")
}

func (e *AvencSpeedhq) SetIldctcmp(value interface{}) error {
	return e.Element.SetProperty("ildctcmp", value)
}

// intra-penalty (int)
//
// Penalty for intra blocks in block decision (Private codec option)

func (e *AvencSpeedhq) GetIntraPenalty() (int, error) {
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

func (e *AvencSpeedhq) SetIntraPenalty(value int) error {
	return e.Element.SetProperty("intra-penalty", value)
}

// keyint-min (int)
//
// minimum interval between IDR-frames (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetKeyintMin() (int, error) {
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

func (e *AvencSpeedhq) SetKeyintMin(value int) error {
	return e.Element.SetProperty("keyint-min", value)
}

// last-pred (int)
//
// amount of motion predictors from the previous frame (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetLastPred() (int, error) {
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

func (e *AvencSpeedhq) SetLastPred(value int) error {
	return e.Element.SetProperty("last-pred", value)
}

// lmax (int)
//
// maximum Lagrange factor (VBR) (Private codec option)

func (e *AvencSpeedhq) GetLmax() (int, error) {
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

func (e *AvencSpeedhq) SetLmax(value int) error {
	return e.Element.SetProperty("lmax", value)
}

// lmin (int)
//
// minimum Lagrange factor (VBR) (Private codec option)

func (e *AvencSpeedhq) GetLmin() (int, error) {
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

func (e *AvencSpeedhq) SetLmin(value int) error {
	return e.Element.SetProperty("lmin", value)
}

// luma-elim-threshold (int)
//
// single coefficient elimination threshold for luminance (negative values also consider dc coefficient) (Private codec option)

func (e *AvencSpeedhq) GetLumaElimThreshold() (int, error) {
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

func (e *AvencSpeedhq) SetLumaElimThreshold(value int) error {
	return e.Element.SetProperty("luma-elim-threshold", value)
}

// lumi-mask (float32)
//
// compresses bright areas stronger than medium ones (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetLumiMask() (float32, error) {
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

func (e *AvencSpeedhq) SetLumiMask(value float32) error {
	return e.Element.SetProperty("lumi-mask", value)
}

// max-bframes (int)
//
// set maximum number of B-frames between non-B-frames (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMaxBframes() (int, error) {
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

func (e *AvencSpeedhq) SetMaxBframes(value int) error {
	return e.Element.SetProperty("max-bframes", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMaxPixels() (int64, error) {
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

func (e *AvencSpeedhq) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMaxrate() (int64, error) {
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

func (e *AvencSpeedhq) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// mbcmp (GstAvcodeccontextCmpFunc)
//
// macroblock compare function (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMbcmp() (interface{}, error) {
	return e.Element.GetProperty("mbcmp")
}

func (e *AvencSpeedhq) SetMbcmp(value interface{}) error {
	return e.Element.SetProperty("mbcmp", value)
}

// mbd (GstAvcodeccontextMbd)
//
// macroblock decision algorithm (high quality mode) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMbd() (interface{}, error) {
	return e.Element.GetProperty("mbd")
}

func (e *AvencSpeedhq) SetMbd(value interface{}) error {
	return e.Element.SetProperty("mbd", value)
}

// mblmax (int)
//
// maximum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMblmax() (int, error) {
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

func (e *AvencSpeedhq) SetMblmax(value int) error {
	return e.Element.SetProperty("mblmax", value)
}

// mblmin (int)
//
// minimum macroblock Lagrange factor (VBR) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMblmin() (int, error) {
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

func (e *AvencSpeedhq) SetMblmin(value int) error {
	return e.Element.SetProperty("mblmin", value)
}

// me-range (int)
//
// limit motion vectors range (1023 for DivX player) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMeRange() (int, error) {
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

func (e *AvencSpeedhq) SetMeRange(value int) error {
	return e.Element.SetProperty("me-range", value)
}

// mepc (int)
//
// Motion estimation bitrate penalty compensation (1.0 = 256) (Private codec option)

func (e *AvencSpeedhq) GetMepc() (int, error) {
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

func (e *AvencSpeedhq) SetMepc(value int) error {
	return e.Element.SetProperty("mepc", value)
}

// mepre (int)
//
// pre motion estimation (Private codec option)

func (e *AvencSpeedhq) GetMepre() (int, error) {
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

func (e *AvencSpeedhq) SetMepre(value int) error {
	return e.Element.SetProperty("mepre", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMinrate() (int64, error) {
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

func (e *AvencSpeedhq) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// motion-est (GstSpeedhqEncoderMotionEst)
//
// motion estimation algorithm (Private codec option)

func (e *AvencSpeedhq) GetMotionEst() (interface{}, error) {
	return e.Element.GetProperty("motion-est")
}

func (e *AvencSpeedhq) SetMotionEst(value interface{}) error {
	return e.Element.SetProperty("motion-est", value)
}

// mpeg-quant (int)
//
// Use MPEG quantizers instead of H.263 (Private codec option)

func (e *AvencSpeedhq) GetMpegQuant() (int, error) {
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

func (e *AvencSpeedhq) SetMpegQuant(value int) error {
	return e.Element.SetProperty("mpeg-quant", value)
}

// mpv-flags (GstSpeedhqEncoderMpvFlags)
//
// Flags common for all mpegvideo-based encoders. (Private codec option)

func (e *AvencSpeedhq) GetMpvFlags() (interface{}, error) {
	return e.Element.GetProperty("mpv-flags")
}

func (e *AvencSpeedhq) SetMpvFlags(value interface{}) error {
	return e.Element.SetProperty("mpv-flags", value)
}

// multipass-cache-file (string)
//
// Filename for multipass cache file

func (e *AvencSpeedhq) GetMultipassCacheFile() (string, error) {
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

func (e *AvencSpeedhq) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// mv0-threshold (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetMv0Threshold() (int, error) {
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

func (e *AvencSpeedhq) SetMv0Threshold(value int) error {
	return e.Element.SetProperty("mv0-threshold", value)
}

// noise-reduction (int)
//
// Noise reduction (Private codec option)

func (e *AvencSpeedhq) GetNoiseReduction() (int, error) {
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

func (e *AvencSpeedhq) SetNoiseReduction(value int) error {
	return e.Element.SetProperty("noise-reduction", value)
}

// nr (int)
//
// noise reduction (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetNr() (int, error) {
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

func (e *AvencSpeedhq) SetNr(value int) error {
	return e.Element.SetProperty("nr", value)
}

// nssew (int)
//
// nsse weight (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetNssew() (int, error) {
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

func (e *AvencSpeedhq) SetNssew(value int) error {
	return e.Element.SetProperty("nssew", value)
}

// p-mask (float32)
//
// inter masking (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetPMask() (float32, error) {
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

func (e *AvencSpeedhq) SetPMask(value float32) error {
	return e.Element.SetProperty("p-mask", value)
}

// pass (GstLibAVEncPass)
//
// Encoding pass/type

func (e *AvencSpeedhq) GetPass() (interface{}, error) {
	return e.Element.GetProperty("pass")
}

func (e *AvencSpeedhq) SetPass(value interface{}) error {
	return e.Element.SetProperty("pass", value)
}

// pbias (int)
//
// inter quant bias (Private codec option)

func (e *AvencSpeedhq) GetPbias() (int, error) {
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

func (e *AvencSpeedhq) SetPbias(value int) error {
	return e.Element.SetProperty("pbias", value)
}

// pre-dia-size (int)
//
// diamond type & size for motion estimation pre-pass (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetPreDiaSize() (int, error) {
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

func (e *AvencSpeedhq) SetPreDiaSize(value int) error {
	return e.Element.SetProperty("pre-dia-size", value)
}

// precmp (GstAvcodeccontextCmpFunc)
//
// pre motion estimation compare function (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetPrecmp() (interface{}, error) {
	return e.Element.GetProperty("precmp")
}

func (e *AvencSpeedhq) SetPrecmp(value interface{}) error {
	return e.Element.SetProperty("precmp", value)
}

// pred (GstAvcodeccontextPred)
//
// prediction method (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetPred() (interface{}, error) {
	return e.Element.GetProperty("pred")
}

func (e *AvencSpeedhq) SetPred(value interface{}) error {
	return e.Element.SetProperty("pred", value)
}

// preme (int)
//
// pre motion estimation (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetPreme() (int, error) {
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

func (e *AvencSpeedhq) SetPreme(value int) error {
	return e.Element.SetProperty("preme", value)
}

// ps (int)
//
// RTP payload size in bytes (Private codec option)

func (e *AvencSpeedhq) GetPs() (int, error) {
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

func (e *AvencSpeedhq) SetPs(value int) error {
	return e.Element.SetProperty("ps", value)
}

// qblur (float32)
//
// video quantizer scale blur (VBR) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetQblur() (float32, error) {
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

func (e *AvencSpeedhq) SetQblur(value float32) error {
	return e.Element.SetProperty("qblur", value)
}

// qcomp (float32)
//
// video quantizer scale compression (VBR). Constant of ratecontrol equation. Recommended range for default rc_eq: 0.0-1.0 (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetQcomp() (float32, error) {
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

func (e *AvencSpeedhq) SetQcomp(value float32) error {
	return e.Element.SetProperty("qcomp", value)
}

// qdiff (int)
//
// maximum difference between the quantizer scales (VBR) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetQdiff() (int, error) {
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

func (e *AvencSpeedhq) SetQdiff(value int) error {
	return e.Element.SetProperty("qdiff", value)
}

// qmax (int)
//
// maximum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetQmax() (int, error) {
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

func (e *AvencSpeedhq) SetQmax(value int) error {
	return e.Element.SetProperty("qmax", value)
}

// qmin (int)
//
// minimum video quantizer scale (VBR) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetQmin() (int, error) {
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

func (e *AvencSpeedhq) SetQmin(value int) error {
	return e.Element.SetProperty("qmin", value)
}

// qsquish (float32)
//
// how to keep quantizer between qmin and qmax (0 = clip, 1 = use differentiable function) (Private codec option)

func (e *AvencSpeedhq) GetQsquish() (float32, error) {
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

func (e *AvencSpeedhq) SetQsquish(value float32) error {
	return e.Element.SetProperty("qsquish", value)
}

// quantizer (float32)
//
// Constant Quantizer

func (e *AvencSpeedhq) GetQuantizer() (float32, error) {
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

func (e *AvencSpeedhq) SetQuantizer(value float32) error {
	return e.Element.SetProperty("quantizer", value)
}

// quantizer-noise-shaping (int)
//
// (null) (Private codec option)

func (e *AvencSpeedhq) GetQuantizerNoiseShaping() (int, error) {
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

func (e *AvencSpeedhq) SetQuantizerNoiseShaping(value int) error {
	return e.Element.SetProperty("quantizer-noise-shaping", value)
}

// rc-buf-aggressivity (float32)
//
// currently useless (Private codec option)

func (e *AvencSpeedhq) GetRcBufAggressivity() (float32, error) {
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

func (e *AvencSpeedhq) SetRcBufAggressivity(value float32) error {
	return e.Element.SetProperty("rc-buf-aggressivity", value)
}

// rc-eq (string)
//
// Set rate control equation. When computing the expression, besides the standard functions defined in the section 'Expression Evaluation', the following functions are available: bits2qp(bits), qp2bits(qp). Also the following constants are available: iTex pTex tex mv fCode iCount mcVar var isI isP isB avgQP qComp avgIITex avgPITex avgPPTex avgBPTex avgTex. (Private codec option)

func (e *AvencSpeedhq) GetRcEq() (string, error) {
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

func (e *AvencSpeedhq) SetRcEq(value string) error {
	return e.Element.SetProperty("rc-eq", value)
}

// rc-init-cplx (float32)
//
// initial complexity for 1-pass encoding (Private codec option)

func (e *AvencSpeedhq) GetRcInitCplx() (float32, error) {
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

func (e *AvencSpeedhq) SetRcInitCplx(value float32) error {
	return e.Element.SetProperty("rc-init-cplx", value)
}

// rc-init-occupancy (int)
//
// number of bits which should be loaded into the rc buffer before decoding starts (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetRcInitOccupancy() (int, error) {
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

func (e *AvencSpeedhq) SetRcInitOccupancy(value int) error {
	return e.Element.SetProperty("rc-init-occupancy", value)
}

// rc-max-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetRcMaxVbvUse() (float32, error) {
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

func (e *AvencSpeedhq) SetRcMaxVbvUse(value float32) error {
	return e.Element.SetProperty("rc-max-vbv-use", value)
}

// rc-min-vbv-use (float32)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetRcMinVbvUse() (float32, error) {
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

func (e *AvencSpeedhq) SetRcMinVbvUse(value float32) error {
	return e.Element.SetProperty("rc-min-vbv-use", value)
}

// rc-qmod-amp (float32)
//
// experimental quantizer modulation (Private codec option)

func (e *AvencSpeedhq) GetRcQmodAmp() (float32, error) {
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

func (e *AvencSpeedhq) SetRcQmodAmp(value float32) error {
	return e.Element.SetProperty("rc-qmod-amp", value)
}

// rc-qmod-freq (int)
//
// experimental quantizer modulation (Private codec option)

func (e *AvencSpeedhq) GetRcQmodFreq() (int, error) {
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

func (e *AvencSpeedhq) SetRcQmodFreq(value int) error {
	return e.Element.SetProperty("rc-qmod-freq", value)
}

// rc-strategy (GstSpeedhqEncoderRcStrategy)
//
// ratecontrol method (Private codec option)

func (e *AvencSpeedhq) GetRcStrategy() (interface{}, error) {
	return e.Element.GetProperty("rc-strategy")
}

func (e *AvencSpeedhq) SetRcStrategy(value interface{}) error {
	return e.Element.SetProperty("rc-strategy", value)
}

// refs (int)
//
// reference frames to consider for motion compensation (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetRefs() (int, error) {
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

func (e *AvencSpeedhq) SetRefs(value int) error {
	return e.Element.SetProperty("refs", value)
}

// sc-threshold (int)
//
// Scene change threshold (Private codec option)

func (e *AvencSpeedhq) GetScThreshold() (int, error) {
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

func (e *AvencSpeedhq) SetScThreshold(value int) error {
	return e.Element.SetProperty("sc-threshold", value)
}

// scplx-mask (float32)
//
// spatial complexity masking (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetScplxMask() (float32, error) {
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

func (e *AvencSpeedhq) SetScplxMask(value float32) error {
	return e.Element.SetProperty("scplx-mask", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetSideDataOnlyPackets() (bool, error) {
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

func (e *AvencSpeedhq) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// skip-cmp (GstSpeedhqEncoderCmpFunc)
//
// Frame skip compare function (Private codec option)

func (e *AvencSpeedhq) GetSkipCmp() (interface{}, error) {
	return e.Element.GetProperty("skip-cmp")
}

func (e *AvencSpeedhq) SetSkipCmp(value interface{}) error {
	return e.Element.SetProperty("skip-cmp", value)
}

// skip-exp (int)
//
// Frame skip exponent (Private codec option)

func (e *AvencSpeedhq) GetSkipExp() (int, error) {
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

func (e *AvencSpeedhq) SetSkipExp(value int) error {
	return e.Element.SetProperty("skip-exp", value)
}

// skip-factor (int)
//
// Frame skip factor (Private codec option)

func (e *AvencSpeedhq) GetSkipFactor() (int, error) {
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

func (e *AvencSpeedhq) SetSkipFactor(value int) error {
	return e.Element.SetProperty("skip-factor", value)
}

// skip-threshold (int)
//
// Frame skip threshold (Private codec option)

func (e *AvencSpeedhq) GetSkipThreshold() (int, error) {
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

func (e *AvencSpeedhq) SetSkipThreshold(value int) error {
	return e.Element.SetProperty("skip-threshold", value)
}

// skipcmp (GstAvcodeccontextCmpFunc)
//
// frame skip compare function (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetSkipcmp() (interface{}, error) {
	return e.Element.GetProperty("skipcmp")
}

func (e *AvencSpeedhq) SetSkipcmp(value interface{}) error {
	return e.Element.SetProperty("skipcmp", value)
}

// slices (int)
//
// set the number of slices, used in parallelized encoding (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetSlices() (int, error) {
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

func (e *AvencSpeedhq) SetSlices(value int) error {
	return e.Element.SetProperty("slices", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencSpeedhq) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// subcmp (GstAvcodeccontextCmpFunc)
//
// sub-pel ME compare function (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetSubcmp() (interface{}, error) {
	return e.Element.GetProperty("subcmp")
}

func (e *AvencSpeedhq) SetSubcmp(value interface{}) error {
	return e.Element.SetProperty("subcmp", value)
}

// subq (int)
//
// sub-pel motion estimation quality (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetSubq() (int, error) {
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

func (e *AvencSpeedhq) SetSubq(value int) error {
	return e.Element.SetProperty("subq", value)
}

// tcplx-mask (float32)
//
// temporal complexity masking (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetTcplxMask() (float32, error) {
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

func (e *AvencSpeedhq) SetTcplxMask(value float32) error {
	return e.Element.SetProperty("tcplx-mask", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencSpeedhq) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencSpeedhq) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetTicksPerFrame() (int, error) {
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

func (e *AvencSpeedhq) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// timecode-frame-start (int64)
//
// GOP timecode frame start number, in non-drop-frame format (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetTimecodeFrameStart() (int64, error) {
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

func (e *AvencSpeedhq) SetTimecodeFrameStart(value int64) error {
	return e.Element.SetProperty("timecode-frame-start", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencSpeedhq) GetTrellis() (int, error) {
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

func (e *AvencSpeedhq) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type SpeedhqEncoderCmpFunc string

const (
	SpeedhqEncoderCmpFuncSad SpeedhqEncoderCmpFunc = "sad" // sad (0) – Sum of absolute differences, fast
	SpeedhqEncoderCmpFuncSse SpeedhqEncoderCmpFunc = "sse" // sse (1) – Sum of squared errors
	SpeedhqEncoderCmpFuncSatd SpeedhqEncoderCmpFunc = "satd" // satd (2) – Sum of absolute Hadamard transformed differences
	SpeedhqEncoderCmpFuncDct SpeedhqEncoderCmpFunc = "dct" // dct (3) – Sum of absolute DCT transformed differences
	SpeedhqEncoderCmpFuncPsnr SpeedhqEncoderCmpFunc = "psnr" // psnr (4) – Sum of squared quantization errors, low quality
	SpeedhqEncoderCmpFuncBit SpeedhqEncoderCmpFunc = "bit" // bit (5) – Number of bits needed for the block
	SpeedhqEncoderCmpFuncRd SpeedhqEncoderCmpFunc = "rd" // rd (6) – Rate distortion optimal, slow
	SpeedhqEncoderCmpFuncZero SpeedhqEncoderCmpFunc = "zero" // zero (7) – Zero
	SpeedhqEncoderCmpFuncVsad SpeedhqEncoderCmpFunc = "vsad" // vsad (8) – Sum of absolute vertical differences
	SpeedhqEncoderCmpFuncVsse SpeedhqEncoderCmpFunc = "vsse" // vsse (9) – Sum of squared vertical differences
	SpeedhqEncoderCmpFuncNsse SpeedhqEncoderCmpFunc = "nsse" // nsse (10) – Noise preserving sum of squared differences
	SpeedhqEncoderCmpFuncDctmax SpeedhqEncoderCmpFunc = "dctmax" // dctmax (13) – dctmax
	SpeedhqEncoderCmpFuncDct264 SpeedhqEncoderCmpFunc = "dct264" // dct264 (14) – dct264
	SpeedhqEncoderCmpFuncMsad SpeedhqEncoderCmpFunc = "msad" // msad (15) – Sum of absolute differences, median predicted
	SpeedhqEncoderCmpFuncChroma SpeedhqEncoderCmpFunc = "chroma" // chroma (256) – chroma
)

type SpeedhqEncoderMotionEst string

const (
	SpeedhqEncoderMotionEstZero SpeedhqEncoderMotionEst = "zero" // zero (0) – zero
	SpeedhqEncoderMotionEstEpzs SpeedhqEncoderMotionEst = "epzs" // epzs (1) – epzs
	SpeedhqEncoderMotionEstXone SpeedhqEncoderMotionEst = "xone" // xone (2) – xone
)

type SpeedhqEncoderMpvFlags string

const (
	SpeedhqEncoderMpvFlagsSkipRd SpeedhqEncoderMpvFlags = "skip_rd" // skip_rd (0x00000001) – RD optimal MB level residual skipping
	SpeedhqEncoderMpvFlagsStrictGop SpeedhqEncoderMpvFlags = "strict_gop" // strict_gop (0x00000002) – Strictly enforce gop size
	SpeedhqEncoderMpvFlagsQpRd SpeedhqEncoderMpvFlags = "qp_rd" // qp_rd (0x00000004) – Use rate distortion optimization for qp selection
	SpeedhqEncoderMpvFlagsCbpRd SpeedhqEncoderMpvFlags = "cbp_rd" // cbp_rd (0x00000008) – use rate distortion optimization for CBP
	SpeedhqEncoderMpvFlagsNaq SpeedhqEncoderMpvFlags = "naq" // naq (0x00000010) – normalize adaptive quantization
	SpeedhqEncoderMpvFlagsMv0 SpeedhqEncoderMpvFlags = "mv0" // mv0 (0x00000020) – always try a mb with mv=<0,0>
)

type SpeedhqEncoderRcStrategy string

const (
	SpeedhqEncoderRcStrategyFfmpeg SpeedhqEncoderRcStrategy = "ffmpeg" // ffmpeg (0) – deprecated, does nothing
)

