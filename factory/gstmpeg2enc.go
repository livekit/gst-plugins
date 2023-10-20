// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// High-quality MPEG-1/2 video encoder
type GstMpeg2enc struct {
	*GstVideoEncoder
}

func NewMpeg2enc() (*GstMpeg2enc, error) {
	e, err := gst.NewElement("mpeg2enc")
	if err != nil {
		return nil, err
	}
	return &GstMpeg2enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewMpeg2encWithName(name string) (*GstMpeg2enc, error) {
	e, err := gst.NewElementWithName("mpeg2enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMpeg2enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// How much to reduce high-frequency resolution (by increasing quantisation)
// Default: 0, Min: 0, Max: 2
func (e *GstMpeg2enc) SetReduceHf(value float32) error {
	return e.Element.SetProperty("reduce-hf", value)
}

func (e *GstMpeg2enc) GetReduceHf() (float32, error) {
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

// Include a sequence header in every GOP
// Default: false
func (e *GstMpeg2enc) SetSequenceHeaderEveryGop(value bool) error {
	return e.Element.SetProperty("sequence-header-every-gop", value)
}

func (e *GstMpeg2enc) GetSequenceHeaderEveryGop() (bool, error) {
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

// Encoding profile format
// Default: 0 (0)
func (e *GstMpeg2enc) SetFormat(value GstMpeg2encFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

func (e *GstMpeg2enc) GetFormat() (GstMpeg2encFormat, error) {
	var value GstMpeg2encFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encFormat")
	}
	return value, nil
}

// MPEG-2 motion estimation and encoding modes
// Default: -1 (-1)
func (e *GstMpeg2enc) SetInterlaceMode(value GstMpeg2encInterlaceMode) error {
	e.Element.SetArg("interlace-mode", string(value))
	return nil
}

func (e *GstMpeg2enc) GetInterlaceMode() (GstMpeg2encInterlaceMode, error) {
	var value GstMpeg2encInterlaceMode
	var ok bool
	v, err := e.Element.GetProperty("interlace-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encInterlaceMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encInterlaceMode")
	}
	return value, nil
}

// Number of bits precision for DC (base colour) in MPEG-2 blocks
// Default: 9, Min: 8, Max: 11
func (e *GstMpeg2enc) SetIntraDcPrec(value int) error {
	return e.Element.SetProperty("intra-dc-prec", value)
}

func (e *GstMpeg2enc) GetIntraDcPrec() (int, error) {
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

// Generate header flags for 3-2 pull down 24fps movies
// Default: false
func (e *GstMpeg2enc) SetPulldown32(value bool) error {
	return e.Element.SetProperty("pulldown-3-2", value)
}

func (e *GstMpeg2enc) GetPulldown32() (bool, error) {
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

// Quantisation matrix to use for encoding
// Default: 9 (0)
func (e *GstMpeg2enc) SetQuantMatrix(value GstMpeg2encQuantisationMatrix) error {
	e.Element.SetArg("quant-matrix", string(value))
	return nil
}

func (e *GstMpeg2enc) GetQuantMatrix() (GstMpeg2encQuantisationMatrix, error) {
	var value GstMpeg2encQuantisationMatrix
	var ok bool
	v, err := e.Element.GetProperty("quant-matrix")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encQuantisationMatrix)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encQuantisationMatrix")
	}
	return value, nil
}

// Max. quantisation reduction for highly active blocks
// Default: 0, Min: -4, Max: 10
func (e *GstMpeg2enc) SetQuantisationReduction(value float32) error {
	return e.Element.SetProperty("quantisation-reduction", value)
}

func (e *GstMpeg2enc) GetQuantisationReduction() (float32, error) {
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

// Compressed video bitrate (kbps)
// Default: 1125, Min: 0, Max: 40000
func (e *GstMpeg2enc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMpeg2enc) GetBitrate() (int, error) {
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

// Target decoders video buffer size (kB) (default depends on format)
// Default: 0, Min: 20, Max: 4000
func (e *GstMpeg2enc) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

func (e *GstMpeg2enc) GetBufsize() (int, error) {
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

// All Group-of-Pictures are closed (for multi-angle DVDs)
// Default: false
func (e *GstMpeg2enc) SetClosedGop(value bool) error {
	return e.Element.SetProperty("closed-gop", value)
}

func (e *GstMpeg2enc) GetClosedGop() (bool, error) {
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

// Prevent the encoder from reencoding pictures in a second pass. This can vastly improve performance, but potentially affect reaching bitrate targets.
// Default: false
func (e *GstMpeg2enc) SetDisableEncodeRetries(value bool) error {
	return e.Element.SetProperty("disable-encode-retries", value)
}

func (e *GstMpeg2enc) GetDisableEncodeRetries() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-encode-retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Motion compensation search radius
// Default: 0, Min: 0, Max: 32
func (e *GstMpeg2enc) SetMotionSearchRadius(value int) error {
	return e.Element.SetProperty("motion-search-radius", value)
}

func (e *GstMpeg2enc) GetMotionSearchRadius() (int, error) {
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

// Quantisation factor (-1=cbr, 0=default, 1=best, 31=worst)
// Default: 0, Min: -1, Max: 31
func (e *GstMpeg2enc) SetQuantisation(value int) error {
	return e.Element.SetProperty("quantisation", value)
}

func (e *GstMpeg2enc) GetQuantisation() (int, error) {
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

// Use strict video resolution and bitrate checks
// Default: true
func (e *GstMpeg2enc) SetConstraints(value bool) error {
	return e.Element.SetProperty("constraints", value)
}

func (e *GstMpeg2enc) GetConstraints() (bool, error) {
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

// Force SVCD width to 480 instead of 540/720
// Default: false
func (e *GstMpeg2enc) SetCorrectSvcdHds(value bool) error {
	return e.Element.SetProperty("correct-svcd-hds", value)
}

func (e *GstMpeg2enc) GetCorrectSvcdHds() (bool, error) {
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

// Generate dummy SVCD scan-data (for vcdimager)
// Default: true
func (e *GstMpeg2enc) SetDummySvcdSof(value bool) error {
	return e.Element.SetProperty("dummy-svcd-sof", value)
}

func (e *GstMpeg2enc) GetDummySvcdSof() (bool, error) {
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

// Force two B frames between I/P frames when closing GOP boundaries
// Default: false
func (e *GstMpeg2enc) SetForceBBP(value bool) error {
	return e.Element.SetProperty("force-b-b-p", value)
}

func (e *GstMpeg2enc) GetForceBBP() (bool, error) {
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

// Maximal size per Group-of-Pictures (-1=default)
// Default: -1, Min: -1, Max: 250
func (e *GstMpeg2enc) SetMaxGopSize(value int) error {
	return e.Element.SetProperty("max-gop-size", value)
}

func (e *GstMpeg2enc) GetMaxGopSize() (int, error) {
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

// Tag output for specific video norm
// Default: 0 (0)
func (e *GstMpeg2enc) SetNorm(value GstMpeg2encVideoNorm) error {
	e.Element.SetArg("norm", string(value))
	return nil
}

func (e *GstMpeg2enc) GetNorm() (GstMpeg2encVideoNorm, error) {
	var value GstMpeg2encVideoNorm
	var ok bool
	v, err := e.Element.GetProperty("norm")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encVideoNorm)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encVideoNorm")
	}
	return value, nil
}

// Force specific playback field order
// Default: 0 (-1)
func (e *GstMpeg2enc) SetPlaybackFieldOrder(value GstMpeg2encPlaybackFieldOrders) error {
	e.Element.SetArg("playback-field-order", string(value))
	return nil
}

func (e *GstMpeg2enc) GetPlaybackFieldOrder() (GstMpeg2encPlaybackFieldOrders, error) {
	var value GstMpeg2encPlaybackFieldOrders
	var ok bool
	v, err := e.Element.GetProperty("playback-field-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encPlaybackFieldOrders)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encPlaybackFieldOrders")
	}
	return value, nil
}

// Maximal luma variance below which quantisation boost is used
// Default: 100, Min: 0, Max: 2500
func (e *GstMpeg2enc) SetQuantReductionMaxVar(value float32) error {
	return e.Element.SetProperty("quant-reduction-max-var", value)
}

func (e *GstMpeg2enc) GetQuantReductionMaxVar() (float32, error) {
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

// How aggressively small-unit picture blocks should be skipped
// Default: 0, Min: -40, Max: 40
func (e *GstMpeg2enc) SetUnitCoeffElim(value int) error {
	return e.Element.SetProperty("unit-coeff-elim", value)
}

func (e *GstMpeg2enc) GetUnitCoeffElim() (int, error) {
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

// Alternate MPEG-2 block scanning. Disabling this might make buggy players play SVCD streams
// Default: true
func (e *GstMpeg2enc) SetAltscanMpeg2(value bool) error {
	return e.Element.SetProperty("altscan-mpeg2", value)
}

func (e *GstMpeg2enc) GetAltscanMpeg2() (bool, error) {
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

// Assumed bitrate of non-video for sequence splitting (kbps)
// Default: 0, Min: 0, Max: 10000
func (e *GstMpeg2enc) SetNonVideoBitrate(value int) error {
	return e.Element.SetProperty("non-video-bitrate", value)
}

func (e *GstMpeg2enc) GetNonVideoBitrate() (int, error) {
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

// Reduction factor for 4x4 subsampled candidate motion estimates (1=max. quality, 4=max. speed)
// Default: 2, Min: 1, Max: 4
func (e *GstMpeg2enc) SetReduction4x4(value int) error {
	return e.Element.SetProperty("reduction-4x4", value)
}

func (e *GstMpeg2enc) GetReduction4x4() (int, error) {
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

// Size of VCD stills (in KB)
// Default: 0, Min: 0, Max: 512
func (e *GstMpeg2enc) SetVcdStillSize(value int) error {
	return e.Element.SetProperty("vcd-still-size", value)
}

func (e *GstMpeg2enc) GetVcdStillSize() (int, error) {
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

// Reduction factor for 2x2 subsampled candidate motion estimates (1=max. quality, 4=max. speed)
// Default: 3, Min: 1, Max: 4
func (e *GstMpeg2enc) SetReduction2x2(value int) error {
	return e.Element.SetProperty("reduction-2x2", value)
}

func (e *GstMpeg2enc) GetReduction2x2() (int, error) {
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

// Display aspect ratio
// Default: 0 (0)
func (e *GstMpeg2enc) SetAspect(value GstMpeg2encAspect) error {
	e.Element.SetArg("aspect", string(value))
	return nil
}

func (e *GstMpeg2enc) GetAspect() (GstMpeg2encAspect, error) {
	var value GstMpeg2encAspect
	var ok bool
	v, err := e.Element.GetProperty("aspect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encAspect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encAspect")
	}
	return value, nil
}

// Dual Prime Motion Estimation Mode for MPEG-2 I/P-frame only streams.  Quite some players do not support this.
// Default: false
func (e *GstMpeg2enc) SetDualprime(value bool) error {
	return e.Element.SetProperty("dualprime", value)
}

func (e *GstMpeg2enc) GetDualprime() (bool, error) {
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

// Output framerate
// Default: 0 (0)
func (e *GstMpeg2enc) SetFramerate(value GstMpeg2encFramerate) error {
	e.Element.SetArg("framerate", string(value))
	return nil
}

func (e *GstMpeg2enc) GetFramerate() (GstMpeg2encFramerate, error) {
	var value GstMpeg2encFramerate
	var ok bool
	v, err := e.Element.GetProperty("framerate")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMpeg2encFramerate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMpeg2encFramerate")
	}
	return value, nil
}

// Maximize high-frequency resolution (for high-quality sources)
// Default: false
func (e *GstMpeg2enc) SetKeepHf(value bool) error {
	return e.Element.SetProperty("keep-hf", value)
}

func (e *GstMpeg2enc) GetKeepHf() (bool, error) {
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

// Place a sequence boundary after each <num> MB (0=disable)
// Default: 0, Min: 0, Max: 10240
func (e *GstMpeg2enc) SetSequenceLength(value int) error {
	return e.Element.SetProperty("sequence-length", value)
}

func (e *GstMpeg2enc) GetSequenceLength() (int, error) {
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

// Number of B frames between each I/P frame
// Default: 0, Min: 0, Max: 2
func (e *GstMpeg2enc) SetBPerRefframe(value int) error {
	return e.Element.SetProperty("b-per-refframe", value)
}

func (e *GstMpeg2enc) GetBPerRefframe() (int, error) {
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

// Minimal size per Group-of-Pictures (-1=default)
// Default: -1, Min: -1, Max: 250
func (e *GstMpeg2enc) SetMinGopSize(value int) error {
	return e.Element.SetProperty("min-gop-size", value)
}

func (e *GstMpeg2enc) GetMinGopSize() (int, error) {
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

type GstMpeg2encFormat string

const (
	GstMpeg2encFormat0  GstMpeg2encFormat = "0"  // Generic MPEG-1
	GstMpeg2encFormat1  GstMpeg2encFormat = "1"  // Standard VCD
	GstMpeg2encFormat2  GstMpeg2encFormat = "2"  // User VCD
	GstMpeg2encFormat3  GstMpeg2encFormat = "3"  // Generic MPEG-2
	GstMpeg2encFormat4  GstMpeg2encFormat = "4"  // Standard SVCD
	GstMpeg2encFormat5  GstMpeg2encFormat = "5"  // User SVCD
	GstMpeg2encFormat6  GstMpeg2encFormat = "6"  // VCD Stills sequences
	GstMpeg2encFormat7  GstMpeg2encFormat = "7"  // SVCD Stills sequences
	GstMpeg2encFormat8  GstMpeg2encFormat = "8"  // DVD MPEG-2 for dvdauthor
	GstMpeg2encFormat9  GstMpeg2encFormat = "9"  // DVD MPEG-2
	GstMpeg2encFormat10 GstMpeg2encFormat = "10" // ATSC 480i
	GstMpeg2encFormat11 GstMpeg2encFormat = "11" // ATSC 480p
	GstMpeg2encFormat12 GstMpeg2encFormat = "12" // ATSC 720p
	GstMpeg2encFormat13 GstMpeg2encFormat = "13" // ATSC 1080i
)

type GstMpeg2encFramerate string

const (
	GstMpeg2encFramerate0 GstMpeg2encFramerate = "0" // Same as input
	GstMpeg2encFramerate1 GstMpeg2encFramerate = "1" // 24/1.001 (NTSC 3:2 pulldown converted film)
	GstMpeg2encFramerate2 GstMpeg2encFramerate = "2" // 24 (native film)
	GstMpeg2encFramerate3 GstMpeg2encFramerate = "3" // 25 (PAL/SECAM video)
	GstMpeg2encFramerate4 GstMpeg2encFramerate = "4" // 30/1.001 (NTSC video)
	GstMpeg2encFramerate5 GstMpeg2encFramerate = "5" // 30
	GstMpeg2encFramerate6 GstMpeg2encFramerate = "6" // 50 (PAL/SECAM fields)
	GstMpeg2encFramerate7 GstMpeg2encFramerate = "7" // 60/1.001 (NTSC fields)
	GstMpeg2encFramerate8 GstMpeg2encFramerate = "8" // 60
)

type GstMpeg2encInterlaceMode string

const (
	GstMpeg2encInterlaceModeMinus1 GstMpeg2encInterlaceMode = "-1" // Format default mode
	GstMpeg2encInterlaceMode0      GstMpeg2encInterlaceMode = "0"  // Progressive
	GstMpeg2encInterlaceMode1      GstMpeg2encInterlaceMode = "1"  // Interlaced, per-frame encoding
	GstMpeg2encInterlaceMode2      GstMpeg2encInterlaceMode = "2"  // Interlaced, per-field-encoding
)

type GstMpeg2encPlaybackFieldOrders string

const (
	GstMpeg2encPlaybackFieldOrders0 GstMpeg2encPlaybackFieldOrders = "0" // Unspecified
	GstMpeg2encPlaybackFieldOrders1 GstMpeg2encPlaybackFieldOrders = "1" // Top-field first
	GstMpeg2encPlaybackFieldOrders2 GstMpeg2encPlaybackFieldOrders = "2" // Bottom-field first
)

type GstMpeg2encQuantisationMatrix string

const (
	GstMpeg2encQuantisationMatrix9 GstMpeg2encQuantisationMatrix = "9" // Default
	GstMpeg2encQuantisationMatrix1 GstMpeg2encQuantisationMatrix = "1" // High resolution
	GstMpeg2encQuantisationMatrix2 GstMpeg2encQuantisationMatrix = "2" // KVCD
	GstMpeg2encQuantisationMatrix3 GstMpeg2encQuantisationMatrix = "3" // TMPGEnc
)

type GstMpeg2encVideoNorm string

const (
	GstMpeg2encVideoNorm0 GstMpeg2encVideoNorm = "0" // Unspecified
	GstMpeg2encVideoNormP GstMpeg2encVideoNorm = "p" // PAL
	GstMpeg2encVideoNormN GstMpeg2encVideoNorm = "n" // NTSC
	GstMpeg2encVideoNormS GstMpeg2encVideoNorm = "s" // SECAM
)

type GstMpeg2encAspect string

const (
	GstMpeg2encAspect0 GstMpeg2encAspect = "0" // Deduce from input
	GstMpeg2encAspect1 GstMpeg2encAspect = "1" // 1:1
	GstMpeg2encAspect2 GstMpeg2encAspect = "2" // 4:3
	GstMpeg2encAspect3 GstMpeg2encAspect = "3" // 16:9
	GstMpeg2encAspect4 GstMpeg2encAspect = "4" // 2.21:1
)
