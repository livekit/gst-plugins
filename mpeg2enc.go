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

type Mpeg2Enc struct {
	Element *gst.Element
}

func NewMpeg2Enc() (*Mpeg2Enc, error) {
	e, err := gst.NewElement("mpeg2enc")
	if err != nil {
		return nil, err
	}
	return &Mpeg2Enc{Element: e}, nil
}

func NewMpeg2EncWithName(name string) (*Mpeg2Enc, error) {
	e, err := gst.NewElementWithName("mpeg2enc", name)
	if err != nil {
		return nil, err
	}
	return &Mpeg2Enc{Element: e}, nil
}

// ----- Properties -----

// altscan-mpeg2 (bool)
//
// Alternate MPEG-2 block scanning. Disabling this might make buggy players play SVCD streams

func (e *Mpeg2Enc) GetAltscanMpeg2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("altscan-mpeg2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetAltscanMpeg2(value bool) error {
	return e.Element.SetProperty("altscan-mpeg2", value)
}

// aspect (GstMpeg2encAspect)
//
// Display aspect ratio

func (e *Mpeg2Enc) GetAspect() (interface{}, error) {
	return e.Element.GetProperty("aspect")
}

func (e *Mpeg2Enc) SetAspect(value interface{}) error {
	return e.Element.SetProperty("aspect", value)
}

// b-per-refframe (int)
//
// Number of B frames between each I/P frame

func (e *Mpeg2Enc) GetBPerRefframe() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("b-per-refframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetBPerRefframe(value int) error {
	return e.Element.SetProperty("b-per-refframe", value)
}

// bitrate (int)
//
// Compressed video bitrate (kbps)

func (e *Mpeg2Enc) GetBitrate() (int, error) {
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

func (e *Mpeg2Enc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bufsize (int)
//
// Target decoders video buffer size (kB) (default depends on format)

func (e *Mpeg2Enc) GetBufsize() (int, error) {
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

func (e *Mpeg2Enc) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// closed-gop (bool)
//
// All Group-of-Pictures are closed (for multi-angle DVDs)

func (e *Mpeg2Enc) GetClosedGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("closed-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetClosedGop(value bool) error {
	return e.Element.SetProperty("closed-gop", value)
}

// constraints (bool)
//
// Use strict video resolution and bitrate checks

func (e *Mpeg2Enc) GetConstraints() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("constraints")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetConstraints(value bool) error {
	return e.Element.SetProperty("constraints", value)
}

// correct-svcd-hds (bool)
//
// Force SVCD width to 480 instead of 540/720

func (e *Mpeg2Enc) GetCorrectSvcdHds() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("correct-svcd-hds")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetCorrectSvcdHds(value bool) error {
	return e.Element.SetProperty("correct-svcd-hds", value)
}

// disable-encode-retries (bool)
//
// Prevent the encoder from reencoding pictures in a second pass.

// dualprime (bool)
//
// Dual Prime Motion Estimation Mode for MPEG-2 I/P-frame only streams.  Quite some players do not support this.

func (e *Mpeg2Enc) GetDualprime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dualprime")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetDualprime(value bool) error {
	return e.Element.SetProperty("dualprime", value)
}

// dummy-svcd-sof (bool)
//
// Generate dummy SVCD scan-data (for vcdimager)

func (e *Mpeg2Enc) GetDummySvcdSof() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dummy-svcd-sof")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetDummySvcdSof(value bool) error {
	return e.Element.SetProperty("dummy-svcd-sof", value)
}

// force-b-b-p (bool)
//
// Force two B frames between I/P frames when closing GOP boundaries

func (e *Mpeg2Enc) GetForceBBP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-b-b-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetForceBBP(value bool) error {
	return e.Element.SetProperty("force-b-b-p", value)
}

// format (GstMpeg2encFormat)
//
// Encoding profile format

func (e *Mpeg2Enc) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

func (e *Mpeg2Enc) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

// framerate (GstMpeg2encFramerate)
//
// Output framerate

func (e *Mpeg2Enc) GetFramerate() (interface{}, error) {
	return e.Element.GetProperty("framerate")
}

func (e *Mpeg2Enc) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

// interlace-mode (GstMpeg2encInterlaceMode)
//
// MPEG-2 motion estimation and encoding modes

func (e *Mpeg2Enc) GetInterlaceMode() (interface{}, error) {
	return e.Element.GetProperty("interlace-mode")
}

func (e *Mpeg2Enc) SetInterlaceMode(value interface{}) error {
	return e.Element.SetProperty("interlace-mode", value)
}

// intra-dc-prec (int)
//
// Number of bits precision for DC (base colour) in MPEG-2 blocks

func (e *Mpeg2Enc) GetIntraDcPrec() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("intra-dc-prec")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetIntraDcPrec(value int) error {
	return e.Element.SetProperty("intra-dc-prec", value)
}

// keep-hf (bool)
//
// Maximize high-frequency resolution (for high-quality sources)

func (e *Mpeg2Enc) GetKeepHf() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-hf")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetKeepHf(value bool) error {
	return e.Element.SetProperty("keep-hf", value)
}

// max-gop-size (int)
//
// Maximal size per Group-of-Pictures (-1=default)

func (e *Mpeg2Enc) GetMaxGopSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetMaxGopSize(value int) error {
	return e.Element.SetProperty("max-gop-size", value)
}

// min-gop-size (int)
//
// Minimal size per Group-of-Pictures (-1=default)

func (e *Mpeg2Enc) GetMinGopSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetMinGopSize(value int) error {
	return e.Element.SetProperty("min-gop-size", value)
}

// motion-search-radius (int)
//
// Motion compensation search radius

func (e *Mpeg2Enc) GetMotionSearchRadius() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("motion-search-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetMotionSearchRadius(value int) error {
	return e.Element.SetProperty("motion-search-radius", value)
}

// non-video-bitrate (int)
//
// Assumed bitrate of non-video for sequence splitting (kbps)

func (e *Mpeg2Enc) GetNonVideoBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("non-video-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetNonVideoBitrate(value int) error {
	return e.Element.SetProperty("non-video-bitrate", value)
}

// norm (GstMpeg2encVideoNorm)
//
// Tag output for specific video norm

func (e *Mpeg2Enc) GetNorm() (interface{}, error) {
	return e.Element.GetProperty("norm")
}

func (e *Mpeg2Enc) SetNorm(value interface{}) error {
	return e.Element.SetProperty("norm", value)
}

// playback-field-order (GstMpeg2encPlaybackFieldOrders)
//
// Force specific playback field order

func (e *Mpeg2Enc) GetPlaybackFieldOrder() (interface{}, error) {
	return e.Element.GetProperty("playback-field-order")
}

func (e *Mpeg2Enc) SetPlaybackFieldOrder(value interface{}) error {
	return e.Element.SetProperty("playback-field-order", value)
}

// pulldown-3-2 (bool)
//
// Generate header flags for 3-2 pull down 24fps movies

func (e *Mpeg2Enc) GetPulldown32() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pulldown-3-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetPulldown32(value bool) error {
	return e.Element.SetProperty("pulldown-3-2", value)
}

// quant-matrix (GstMpeg2encQuantisationMatrix)
//
// Quantisation matrix to use for encoding

func (e *Mpeg2Enc) GetQuantMatrix() (interface{}, error) {
	return e.Element.GetProperty("quant-matrix")
}

func (e *Mpeg2Enc) SetQuantMatrix(value interface{}) error {
	return e.Element.SetProperty("quant-matrix", value)
}

// quant-reduction-max-var (float32)
//
// Maximal luma variance below which quantisation boost is used

func (e *Mpeg2Enc) GetQuantReductionMaxVar() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quant-reduction-max-var")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetQuantReductionMaxVar(value float32) error {
	return e.Element.SetProperty("quant-reduction-max-var", value)
}

// quantisation (int)
//
// Quantisation factor (-1=cbr, 0=default, 1=best, 31=worst)

func (e *Mpeg2Enc) GetQuantisation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quantisation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetQuantisation(value int) error {
	return e.Element.SetProperty("quantisation", value)
}

// quantisation-reduction (float32)
//
// Max. quantisation reduction for highly active blocks

func (e *Mpeg2Enc) GetQuantisationReduction() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quantisation-reduction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetQuantisationReduction(value float32) error {
	return e.Element.SetProperty("quantisation-reduction", value)
}

// reduce-hf (float32)
//
// How much to reduce high-frequency resolution (by increasing quantisation)

func (e *Mpeg2Enc) GetReduceHf() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("reduce-hf")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetReduceHf(value float32) error {
	return e.Element.SetProperty("reduce-hf", value)
}

// reduction-2x2 (int)
//
// Reduction factor for 2x2 subsampled candidate motion estimates (1=max. quality, 4=max. speed)

func (e *Mpeg2Enc) GetReduction2X2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("reduction-2x2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetReduction2X2(value int) error {
	return e.Element.SetProperty("reduction-2x2", value)
}

// reduction-4x4 (int)
//
// Reduction factor for 4x4 subsampled candidate motion estimates (1=max. quality, 4=max. speed)

func (e *Mpeg2Enc) GetReduction4X4() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("reduction-4x4")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetReduction4X4(value int) error {
	return e.Element.SetProperty("reduction-4x4", value)
}

// sequence-header-every-gop (bool)
//
// Include a sequence header in every GOP

func (e *Mpeg2Enc) GetSequenceHeaderEveryGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sequence-header-every-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetSequenceHeaderEveryGop(value bool) error {
	return e.Element.SetProperty("sequence-header-every-gop", value)
}

// sequence-length (int)
//
// Place a sequence boundary after each  MB (0=disable)

func (e *Mpeg2Enc) GetSequenceLength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sequence-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetSequenceLength(value int) error {
	return e.Element.SetProperty("sequence-length", value)
}

// unit-coeff-elim (int)
//
// How aggressively small-unit picture blocks should be skipped

func (e *Mpeg2Enc) GetUnitCoeffElim() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("unit-coeff-elim")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetUnitCoeffElim(value int) error {
	return e.Element.SetProperty("unit-coeff-elim", value)
}

// vcd-still-size (int)
//
// Size of VCD stills (in KB)

func (e *Mpeg2Enc) GetVcdStillSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("vcd-still-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mpeg2Enc) SetVcdStillSize(value int) error {
	return e.Element.SetProperty("vcd-still-size", value)
}

// ----- Constants -----

type GstMpeg2EncPlaybackFieldOrders string

const (
	GstMpeg2EncPlaybackFieldOrders0 GstMpeg2EncPlaybackFieldOrders = "0" // 0 (-1) – Unspecified
	GstMpeg2EncPlaybackFieldOrders1 GstMpeg2EncPlaybackFieldOrders = "1" // 1 (1) – Top-field first
	GstMpeg2EncPlaybackFieldOrders2 GstMpeg2EncPlaybackFieldOrders = "2" // 2 (2) – Bottom-field first
)

type GstMpeg2EncQuantisationMatrix string

const (
	GstMpeg2EncQuantisationMatrix9 GstMpeg2EncQuantisationMatrix = "9" // 9 (0) – Default
	GstMpeg2EncQuantisationMatrix1 GstMpeg2EncQuantisationMatrix = "1" // 1 (1) – High resolution
	GstMpeg2EncQuantisationMatrix2 GstMpeg2EncQuantisationMatrix = "2" // 2 (2) – KVCD
	GstMpeg2EncQuantisationMatrix3 GstMpeg2EncQuantisationMatrix = "3" // 3 (3) – TMPGEnc
)

type GstMpeg2EncVideoNorm string

const (
	GstMpeg2EncVideoNorm0 GstMpeg2EncVideoNorm = "0" // 0 (0) – Unspecified
	GstMpeg2EncVideoNormP GstMpeg2EncVideoNorm = "p" // p (112) – PAL
	GstMpeg2EncVideoNormN GstMpeg2EncVideoNorm = "n" // n (110) – NTSC
	GstMpeg2EncVideoNormS GstMpeg2EncVideoNorm = "s" // s (115) – SECAM
)

type GstMpeg2EncAspect string

const (
	GstMpeg2EncAspect0 GstMpeg2EncAspect = "0" // 0 (0) – Deduce from input
	GstMpeg2EncAspect1 GstMpeg2EncAspect = "1" // 1 (1) – 1:1
	GstMpeg2EncAspect2 GstMpeg2EncAspect = "2" // 2 (2) – 4:3
	GstMpeg2EncAspect3 GstMpeg2EncAspect = "3" // 3 (3) – 16:9
	GstMpeg2EncAspect4 GstMpeg2EncAspect = "4" // 4 (4) – 2.21:1
)

type GstMpeg2EncFormat string

const (
	GstMpeg2EncFormat0 GstMpeg2EncFormat = "0" // 0 (0) – Generic MPEG-1
	GstMpeg2EncFormat1 GstMpeg2EncFormat = "1" // 1 (1) – Standard VCD
	GstMpeg2EncFormat2 GstMpeg2EncFormat = "2" // 2 (2) – User VCD
	GstMpeg2EncFormat3 GstMpeg2EncFormat = "3" // 3 (3) – Generic MPEG-2
	GstMpeg2EncFormat4 GstMpeg2EncFormat = "4" // 4 (4) – Standard SVCD
	GstMpeg2EncFormat5 GstMpeg2EncFormat = "5" // 5 (5) – User SVCD
	GstMpeg2EncFormat6 GstMpeg2EncFormat = "6" // 6 (6) – VCD Stills sequences
	GstMpeg2EncFormat7 GstMpeg2EncFormat = "7" // 7 (7) – SVCD Stills sequences
	GstMpeg2EncFormat8 GstMpeg2EncFormat = "8" // 8 (8) – DVD MPEG-2 for dvdauthor
	GstMpeg2EncFormat9 GstMpeg2EncFormat = "9" // 9 (9) – DVD MPEG-2
	GstMpeg2EncFormat10 GstMpeg2EncFormat = "10" // 10 (10) – ATSC 480i
	GstMpeg2EncFormat11 GstMpeg2EncFormat = "11" // 11 (11) – ATSC 480p
	GstMpeg2EncFormat12 GstMpeg2EncFormat = "12" // 12 (12) – ATSC 720p
	GstMpeg2EncFormat13 GstMpeg2EncFormat = "13" // 13 (13) – ATSC 1080i
)

type GstMpeg2EncFramerate string

const (
	GstMpeg2EncFramerate0 GstMpeg2EncFramerate = "0" // 0 (0) – Same as input
	GstMpeg2EncFramerate1 GstMpeg2EncFramerate = "1" // 1 (1) – 24/1.001 (NTSC 3:2 pulldown converted film)
	GstMpeg2EncFramerate2 GstMpeg2EncFramerate = "2" // 2 (2) – 24 (native film)
	GstMpeg2EncFramerate3 GstMpeg2EncFramerate = "3" // 3 (3) – 25 (PAL/SECAM video)
	GstMpeg2EncFramerate4 GstMpeg2EncFramerate = "4" // 4 (4) – 30/1.001 (NTSC video)
	GstMpeg2EncFramerate5 GstMpeg2EncFramerate = "5" // 5 (5) – 30
	GstMpeg2EncFramerate6 GstMpeg2EncFramerate = "6" // 6 (6) – 50 (PAL/SECAM fields)
	GstMpeg2EncFramerate7 GstMpeg2EncFramerate = "7" // 7 (7) – 60/1.001 (NTSC fields)
	GstMpeg2EncFramerate8 GstMpeg2EncFramerate = "8" // 8 (8) – 60
)

type GstMpeg2EncInterlaceMode string

const (
	GstMpeg2EncInterlaceMode1 GstMpeg2EncInterlaceMode = "-1" // -1 (-1) – Format default mode
	GstMpeg2EncInterlaceMode0 GstMpeg2EncInterlaceMode = "0" // 0 (0) – Progressive
	GstMpeg2EncInterlaceMode11 GstMpeg2EncInterlaceMode = "1" // 1 (1) – Interlaced, per-frame encoding
	GstMpeg2EncInterlaceMode2 GstMpeg2EncInterlaceMode = "2" // 2 (2) – Interlaced, per-field-encoding
)

