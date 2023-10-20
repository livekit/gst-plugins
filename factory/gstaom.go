// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decode AV1 video streams
type GstAV1Dec struct {
	*GstVideoDecoder
}

func NewAV1Dec() (*GstAV1Dec, error) {
	e, err := gst.NewElement("av1dec")
	if err != nil {
		return nil, err
	}
	return &GstAV1Dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewAV1DecWithName(name string) (*GstAV1Dec, error) {
	e, err := gst.NewElementWithName("av1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstAV1Dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Encode AV1 video streams
type GstAV1Enc struct {
	*GstVideoEncoder
}

func NewAV1Enc() (*GstAV1Enc, error) {
	e, err := gst.NewElement("av1enc")
	if err != nil {
		return nil, err
	}
	return &GstAV1Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewAV1EncWithName(name string) (*GstAV1Enc, error) {
	e, err := gst.NewElementWithName("av1enc", name)
	if err != nil {
		return nil, err
	}
	return &GstAV1Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Maximum (worst quality) quantizer
// Default: 0, Min: 0, Max: -1
func (e *GstAV1Enc) SetMaxQuantizer(value uint) error {
	return e.Element.SetProperty("max-quantizer", value)
}

func (e *GstAV1Enc) GetMaxQuantizer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Frame resize denominator, assuming 8 as the numerator
// Default: 8, Min: 8, Max: 16
func (e *GstAV1Enc) SetResizeDenominator(value uint) error {
	return e.Element.SetProperty("resize-denominator", value)
}

func (e *GstAV1Enc) GetResizeDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("resize-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable row based multi-threading
// Default: true
func (e *GstAV1Enc) SetRowMt(value bool) error {
	return e.Element.SetProperty("row-mt", value)
}

func (e *GstAV1Enc) GetRowMt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("row-mt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Partition into separate horizontal tile rows from image frame which can enable parallel encoding
// Default: 0, Min: 0, Max: 6
func (e *GstAV1Enc) SetTileRows(value uint) error {
	return e.Element.SetProperty("tile-rows", value)
}

func (e *GstAV1Enc) GetTileRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Usage profile is used to guide the default config for the encoder
// Default: good-quality (0)
func (e *GstAV1Enc) SetUsageProfile(value GstAV1EncUsageProfile) error {
	e.Element.SetArg("usage-profile", string(value))
	return nil
}

func (e *GstAV1Enc) GetUsageProfile() (GstAV1EncUsageProfile, error) {
	var value GstAV1EncUsageProfile
	var ok bool
	v, err := e.Element.GetProperty("usage-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAV1EncUsageProfile)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAV1EncUsageProfile")
	}
	return value, nil
}

// CPU Used. A Value greater than 0 will increase encoder speed at the expense of quality.
// Default: 0, Min: 0, Max: 10
func (e *GstAV1Enc) SetCpuUsed(value int) error {
	return e.Element.SetProperty("cpu-used", value)
}

func (e *GstAV1Enc) GetCpuUsed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cpu-used")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum (best quality) quantizer
// Default: 0, Min: 0, Max: -1
func (e *GstAV1Enc) SetMinQuantizer(value uint) error {
	return e.Element.SetProperty("min-quantizer", value)
}

func (e *GstAV1Enc) GetMinQuantizer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Frame super-resolution denominator, used only by SUPERRES_FIXED mode
// Default: 8, Min: 8, Max: 16
func (e *GstAV1Enc) SetSuperresDenominator(value uint) error {
	return e.Element.SetProperty("superres-denominator", value)
}

func (e *GstAV1Enc) GetSuperresDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Keyframe super-resolution denominator
// Default: 8, Min: 8, Max: 16
func (e *GstAV1Enc) SetSuperresKfDenominator(value uint) error {
	return e.Element.SetProperty("superres-kf-denominator", value)
}

func (e *GstAV1Enc) GetSuperresKfDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-kf-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// It integrates upscaling after the encode/decode process
// Default: none (0)
func (e *GstAV1Enc) SetSuperresMode(value GstAV1EncSuperresMode) error {
	e.Element.SetArg("superres-mode", string(value))
	return nil
}

func (e *GstAV1Enc) GetSuperresMode() (GstAV1EncSuperresMode, error) {
	var value GstAV1EncSuperresMode
	var ok bool
	v, err := e.Element.GetProperty("superres-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAV1EncSuperresMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAV1EncSuperresMode")
	}
	return value, nil
}

// Target bitrate, in kilobits per second
// Default: 256, Min: 1, Max: -1
func (e *GstAV1Enc) SetTargetBitrate(value uint) error {
	return e.Element.SetProperty("target-bitrate", value)
}

func (e *GstAV1Enc) GetTargetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Temporal resampling configuration, drop frames as a strategy to meet its target data rate Set to zero (0) to disable this feature.
// Default: 0, Min: 0, Max: -1
func (e *GstAV1Enc) SetDropFrame(value uint) error {
	return e.Element.SetProperty("drop-frame", value)
}

func (e *GstAV1Enc) GetDropFrame() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("drop-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum distance between keyframes (number of frames)
// Default: 30, Min: 0, Max: 2147483647
func (e *GstAV1Enc) SetKeyframeMaxDist(value int) error {
	return e.Element.SetProperty("keyframe-max-dist", value)
}

func (e *GstAV1Enc) GetKeyframeMaxDist() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-max-dist")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Determines whether keyframes are placed automatically by the encoder
// Default: auto (1)
func (e *GstAV1Enc) SetKeyframeMode(value GstAV1EncKFMode) error {
	e.Element.SetArg("keyframe-mode", string(value))
	return nil
}

func (e *GstAV1Enc) GetKeyframeMode() (GstAV1EncKFMode, error) {
	var value GstAV1EncKFMode
	var ok bool
	v, err := e.Element.GetProperty("keyframe-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAV1EncKFMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAV1EncKFMode")
	}
	return value, nil
}

// Keyframe super-resolution qindex threshold, used only by SUPERRES_QTHRESH mode
// Default: 63, Min: 1, Max: 63
func (e *GstAV1Enc) SetSuperresKfQthresh(value uint) error {
	return e.Element.SetProperty("superres-kf-qthresh", value)
}

func (e *GstAV1Enc) GetSuperresKfQthresh() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-kf-qthresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max number of threads to use encoding, set to 0 determine the approximate number of threads that the system schedule
// Default: 0, Min: 0, Max: -1
func (e *GstAV1Enc) SetThreads(value uint) error {
	return e.Element.SetProperty("threads", value)
}

func (e *GstAV1Enc) GetThreads() (uint, error) {
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

// Partition into separate vertical tile columns from image frame which can enable parallel encoding
// Default: 0, Min: 0, Max: 6
func (e *GstAV1Enc) SetTileColumns(value uint) error {
	return e.Element.SetProperty("tile-columns", value)
}

func (e *GstAV1Enc) GetTileColumns() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tile-columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate control adaptation undershoot control
// Default: 25, Min: 0, Max: 1000
func (e *GstAV1Enc) SetUndershootPct(value uint) error {
	return e.Element.SetProperty("undershoot-pct", value)
}

func (e *GstAV1Enc) GetUndershootPct() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("undershoot-pct")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Decoder buffer optimal size, expressed in units of time (milliseconds)
// Default: 5000, Min: 0, Max: -1
func (e *GstAV1Enc) SetBufOptimalSz(value uint) error {
	return e.Element.SetProperty("buf-optimal-sz", value)
}

func (e *GstAV1Enc) GetBufOptimalSz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buf-optimal-sz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Current phase for multi-pass encoding or @GST_AV1_ENC_ONE_PASS for single pass
// Default: one-pass (0)
func (e *GstAV1Enc) SetEncPass(value GstAV1EncEncPass) error {
	e.Element.SetArg("enc-pass", string(value))
	return nil
}

func (e *GstAV1Enc) GetEncPass() (GstAV1EncEncPass, error) {
	var value GstAV1EncEncPass
	var ok bool
	v, err := e.Element.GetProperty("enc-pass")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAV1EncEncPass)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAV1EncEncPass")
	}
	return value, nil
}

// Rate control adaptation overshoot control
// Default: 25, Min: 0, Max: 1000
func (e *GstAV1Enc) SetOvershootPct(value uint) error {
	return e.Element.SetProperty("overshoot-pct", value)
}

func (e *GstAV1Enc) GetOvershootPct() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("overshoot-pct")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Frame resize mode
// Default: none (0)
func (e *GstAV1Enc) SetResizeMode(value GstAV1EncResizeMode) error {
	e.Element.SetArg("resize-mode", string(value))
	return nil
}

func (e *GstAV1Enc) GetResizeMode() (GstAV1EncResizeMode, error) {
	var value GstAV1EncResizeMode
	var ok bool
	v, err := e.Element.GetProperty("resize-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAV1EncResizeMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAV1EncResizeMode")
	}
	return value, nil
}

// Decoder buffer initial size, expressed in units of time (milliseconds)
// Default: 4000, Min: 0, Max: -1
func (e *GstAV1Enc) SetBufInitialSz(value uint) error {
	return e.Element.SetProperty("buf-initial-sz", value)
}

func (e *GstAV1Enc) GetBufInitialSz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buf-initial-sz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Frame super-resolution qindex threshold, used only by SUPERRES_QTHRESH mode
// Default: 63, Min: 1, Max: 63
func (e *GstAV1Enc) SetSuperresQthresh(value uint) error {
	return e.Element.SetProperty("superres-qthresh", value)
}

func (e *GstAV1Enc) GetSuperresQthresh() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("superres-qthresh")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Decoder buffer size, expressed in units of time (milliseconds)
// Default: 6000, Min: 0, Max: -1
func (e *GstAV1Enc) SetBufSz(value uint) error {
	return e.Element.SetProperty("buf-sz", value)
}

func (e *GstAV1Enc) GetBufSz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buf-sz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Frame resize keyframe denominator, assuming 8 as the numerator
// Default: 8, Min: 8, Max: 16
func (e *GstAV1Enc) SetResizeKfDenominator(value uint) error {
	return e.Element.SetProperty("resize-kf-denominator", value)
}

func (e *GstAV1Enc) GetResizeKfDenominator() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("resize-kf-denominator")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate control algorithm to use, indicates the end usage of this stream
// Default: vbr (0)
func (e *GstAV1Enc) SetEndUsage(value GstAV1EncEndUsageMode) error {
	e.Element.SetArg("end-usage", string(value))
	return nil
}

func (e *GstAV1Enc) GetEndUsage() (GstAV1EncEndUsageMode, error) {
	var value GstAV1EncEndUsageMode
	var ok bool
	v, err := e.Element.GetProperty("end-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAV1EncEndUsageMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAV1EncEndUsageMode")
	}
	return value, nil
}

// Maximum number of future frames the encoder is allowed to consume before producing the current output frame. Set value to 0 for disabling lagged encoding.
// Default: 0, Min: 0, Max: -1
func (e *GstAV1Enc) SetLagInFrames(value uint) error {
	return e.Element.SetProperty("lag-in-frames", value)
}

func (e *GstAV1Enc) GetLagInFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lag-in-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstAV1EncKFMode string

const (
	GstAV1EncKFModeAuto     GstAV1EncKFMode = "auto"     // Encoder determines optimal keyframe placement automatically
	GstAV1EncKFModeDisabled GstAV1EncKFMode = "disabled" // Encoder does not place keyframes
)

type GstAV1EncResizeMode string

const (
	GstAV1EncResizeModeNone   GstAV1EncResizeMode = "none"   // No frame resizing allowed
	GstAV1EncResizeModeFixed  GstAV1EncResizeMode = "fixed"  // All frames are coded at the specified scale
	GstAV1EncResizeModeRandom GstAV1EncResizeMode = "random" // All frames are coded at a random scale
)

type GstAV1EncSuperresMode string

const (
	GstAV1EncSuperresModeNone    GstAV1EncSuperresMode = "none"    // No frame superres allowed
	GstAV1EncSuperresModeFixed   GstAV1EncSuperresMode = "fixed"   // All frames are coded at the specified scale and super-resolved
	GstAV1EncSuperresModeRandom  GstAV1EncSuperresMode = "random"  // All frames are coded at a random scale and super-resolved
	GstAV1EncSuperresModeQthresh GstAV1EncSuperresMode = "qthresh" // Superres scale for a frame is determined based on q_index
)

type GstAV1EncUsageProfile string

const (
	GstAV1EncUsageProfileGoodQuality GstAV1EncUsageProfile = "good-quality" // Good Quality profile
	GstAV1EncUsageProfileRealtime    GstAV1EncUsageProfile = "realtime"     // Realtime profile
	GstAV1EncUsageProfileAllIntra    GstAV1EncUsageProfile = "all-intra"    // All Intra profile
)

type GstAV1EncEncPass string

const (
	GstAV1EncEncPassOnePass    GstAV1EncEncPass = "one-pass"    // Single pass mode
	GstAV1EncEncPassFirstPass  GstAV1EncEncPass = "first-pass"  // First pass of multi-pass mode
	GstAV1EncEncPassSecondPass GstAV1EncEncPass = "second-pass" // Second pass of multi-pass mode
	GstAV1EncEncPassThirdPass  GstAV1EncEncPass = "third-pass"  // Third pass of multi-pass mode
)

type GstAV1EncEndUsageMode string

const (
	GstAV1EncEndUsageModeVbr GstAV1EncEndUsageMode = "vbr" // Variable Bit Rate Mode
	GstAV1EncEndUsageModeCbr GstAV1EncEndUsageMode = "cbr" // Constant Bit Rate Mode
	GstAV1EncEndUsageModeCq  GstAV1EncEndUsageMode = "cq"  // Constrained Quality Mode
	GstAV1EncEndUsageModeQ   GstAV1EncEndUsageMode = "q"   // Constant Quality Mode
)
