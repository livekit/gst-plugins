// source: gst-plugins-ugly

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// libx264-based H.264 video encoder
type GstX264Enc struct {
	*GstVideoEncoder
}

func NewX264Enc() (*GstX264Enc, error) {
	e, err := gst.NewElement("x264enc")
	if err != nil {
		return nil, err
	}
	return &GstX264Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewX264EncWithName(name string) (*GstX264Enc, error) {
	e, err := gst.NewElementWithName("x264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstX264Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Maximal distance between two key-frames (0 for automatic)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstX264Enc) SetKeyIntMax(value uint) error {
	return e.Element.SetProperty("key-int-max", value)
}

func (e *GstX264Enc) GetKeyIntMax() (uint, error) {
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

// String of x264 options (overridden by element properties) in the format "key1=value1:key2=value2".

func (e *GstX264Enc) SetOptionString(value string) error {
	return e.Element.SetProperty("option-string", value)
}

func (e *GstX264Enc) GetOptionString() (string, error) {
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

// Constant quantizer or quality to apply
// Default: 21, Min: 0, Max: 50
func (e *GstX264Enc) SetQuantizer(value uint) error {
	return e.Element.SetProperty("quantizer", value)
}

func (e *GstX264Enc) GetQuantizer() (uint, error) {
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

// Size of the VBV buffer in milliseconds
// Default: 600, Min: 0, Max: 10000
func (e *GstX264Enc) SetVbvBufCapacity(value uint) error {
	return e.Element.SetProperty("vbv-buf-capacity", value)
}

func (e *GstX264Enc) GetVbvBufCapacity() (uint, error) {
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

// Keep some B-frames as references
// Default: false
func (e *GstX264Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

func (e *GstX264Enc) GetBPyramid() (bool, error) {
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

// Insert VUI NAL in stream
// Default: true
func (e *GstX264Enc) SetInsertVui(value bool) error {
	return e.Element.SetProperty("insert-vui", value)
}

func (e *GstX264Enc) GetInsertVui() (bool, error) {
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

// Use Periodic Intra Refresh instead of IDR frames
// Default: false
func (e *GstX264Enc) SetIntraRefresh(value bool) error {
	return e.Element.SetProperty("intra-refresh", value)
}

func (e *GstX264Enc) GetIntraRefresh() (bool, error) {
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

// Quantizer factor between P- and B-frames
// Default: 1.3, Min: 0, Max: 2
func (e *GstX264Enc) SetPbFactor(value float32) error {
	return e.Element.SetProperty("pb-factor", value)
}

func (e *GstX264Enc) GetPbFactor() (float32, error) {
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

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstX264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstX264Enc) GetAud() (bool, error) {
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

// Bitrate in kbit/sec
// Default: 2048, Min: 1, Max: 2048000
func (e *GstX264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstX264Enc) GetBitrate() (uint, error) {
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

// Generate byte stream format of NALU
// Default: false
func (e *GstX264Enc) SetByteStream(value bool) error {
	return e.Element.SetProperty("byte-stream", value)
}

func (e *GstX264Enc) GetByteStream() (bool, error) {
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

// Noise reduction strength
// Default: 0, Min: 0, Max: 100000
func (e *GstX264Enc) SetNoiseReduction(value uint) error {
	return e.Element.SetProperty("noise-reduction", value)
}

func (e *GstX264Enc) GetNoiseReduction() (uint, error) {
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

// Preset name for speed/quality tradeoff options (can affect decode compatibility - impose restrictions separately for your target decoder)
// Default: medium (6)
func (e *GstX264Enc) SetSpeedPreset(value GstX264EncPreset) error {
	e.Element.SetArg("speed-preset", string(value))
	return nil
}

func (e *GstX264Enc) GetSpeedPreset() (GstX264EncPreset, error) {
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

// Automatically decide how many B-frames to use
// Default: true
func (e *GstX264Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstX264Enc) GetBAdapt() (bool, error) {
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

// Set frame packing mode for Stereoscopic content
// Default: auto (-1)
func (e *GstX264Enc) SetFramePacking(value GstX264EncFramePacking) error {
	e.Element.SetArg("frame-packing", string(value))
	return nil
}

func (e *GstX264Enc) GetFramePacking() (GstX264EncFramePacking, error) {
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

// Integer pixel motion estimation method
// Default: hex (1)
func (e *GstX264Enc) SetMe(value GstX264EncMe) error {
	e.Element.SetArg("me", string(value))
	return nil
}

func (e *GstX264Enc) GetMe() (GstX264EncMe, error) {
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

// Preset name for psychovisual tuning options
// Default: none (0)
func (e *GstX264Enc) SetPsyTune(value GstX264EncPsyTune) error {
	e.Element.SetArg("psy-tune", string(value))
	return nil
}

func (e *GstX264Enc) GetPsyTune() (GstX264EncPsyTune, error) {
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

// SPS and PPS ID number
// Default: 0, Min: 0, Max: 31
func (e *GstX264Enc) SetSpsId(value uint) error {
	return e.Element.SetProperty("sps-id", value)
}

func (e *GstX264Enc) GetSpsId() (uint, error) {
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

// Number of threads used by the codec (0 for automatic)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstX264Enc) SetThreads(value uint) error {
	return e.Element.SetProperty("threads", value)
}

func (e *GstX264Enc) GetThreads() (uint, error) {
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

// Enable trellis searched quantization
// Default: true
func (e *GstX264Enc) SetTrellis(value bool) error {
	return e.Element.SetProperty("trellis", value)
}

func (e *GstX264Enc) GetTrellis() (bool, error) {
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

// Partitions to consider
// Default: (none)
func (e *GstX264Enc) SetAnalyse(value GstX264EncAnalyse) error {
	e.Element.SetArg("analyse", fmt.Sprint(value))
	return nil
}

func (e *GstX264Enc) GetAnalyse() (GstX264EncAnalyse, error) {
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

// Enable CABAC entropy coding
// Default: true
func (e *GstX264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstX264Enc) GetCabac() (bool, error) {
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

// Quantizer factor between I- and P-frames
// Default: 1.4, Min: 0, Max: 2
func (e *GstX264Enc) SetIpFactor(value float32) error {
	return e.Element.SetProperty("ip-factor", value)
}

func (e *GstX264Enc) GetIpFactor() (float32, error) {
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

// Weighted prediction for B-frames
// Default: false
func (e *GstX264Enc) SetWeightb(value bool) error {
	return e.Element.SetProperty("weightb", value)
}

func (e *GstX264Enc) GetWeightb() (bool, error) {
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

// Interlaced material
// Default: false
func (e *GstX264Enc) SetInterlaced(value bool) error {
	return e.Element.SetProperty("interlaced", value)
}

func (e *GstX264Enc) GetInterlaced() (bool, error) {
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

// Maximum quantizer
// Default: 51, Min: 0, Max: 63
func (e *GstX264Enc) SetQpMax(value uint) error {
	return e.Element.SetProperty("qp-max", value)
}

func (e *GstX264Enc) GetQpMax() (uint, error) {
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

// Low latency but lower efficiency threading
// Default: false
func (e *GstX264Enc) SetSlicedThreads(value bool) error {
	return e.Element.SetProperty("sliced-threads", value)
}

func (e *GstX264Enc) GetSlicedThreads() (bool, error) {
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

// Encoding pass/type
// Default: cbr (0)
func (e *GstX264Enc) SetPass(value GstX264EncPass) error {
	e.Element.SetArg("pass", string(value))
	return nil
}

func (e *GstX264Enc) GetPass() (GstX264EncPass, error) {
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

// Minimum quantizer
// Default: 10, Min: 0, Max: 63
func (e *GstX264Enc) SetQpMin(value uint) error {
	return e.Element.SetProperty("qp-min", value)
}

func (e *GstX264Enc) GetQpMin() (uint, error) {
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

// Maximum quantizer difference between frames
// Default: 4, Min: 0, Max: 63
func (e *GstX264Enc) SetQpStep(value uint) error {
	return e.Element.SetProperty("qp-step", value)
}

func (e *GstX264Enc) GetQpStep() (uint, error) {
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

// Number of frames for frametype lookahead
// Default: 40, Min: 0, Max: 250
func (e *GstX264Enc) SetRcLookahead(value int) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstX264Enc) GetRcLookahead() (int, error) {
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

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 16
func (e *GstX264Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

func (e *GstX264Enc) GetBframes() (uint, error) {
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

// Adaptive spatial transform size
// Default: false
func (e *GstX264Enc) SetDct8x8(value bool) error {
	return e.Element.SetProperty("dct8x8", value)
}

func (e *GstX264Enc) GetDct8x8() (bool, error) {
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

// Macroblock-Tree ratecontrol
// Default: true
func (e *GstX264Enc) SetMbTree(value bool) error {
	return e.Element.SetProperty("mb-tree", value)
}

func (e *GstX264Enc) GetMbTree() (bool, error) {
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

// Preset name for non-psychovisual tuning options
// Default: (none)
func (e *GstX264Enc) SetTune(value GstX264EncTune) error {
	e.Element.SetArg("tune", fmt.Sprint(value))
	return nil
}

func (e *GstX264Enc) GetTune() (GstX264EncTune, error) {
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

// Number of buffer frames for threaded lookahead (-1 for automatic)
// Default: -1, Min: -1, Max: 250
func (e *GstX264Enc) SetSyncLookahead(value int) error {
	return e.Element.SetProperty("sync-lookahead", value)
}

func (e *GstX264Enc) GetSyncLookahead() (int, error) {
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

// Filename for multipass cache file
// Default: x264.log
func (e *GstX264Enc) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

func (e *GstX264Enc) GetMultipassCacheFile() (string, error) {
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

// Number of reference frames
// Default: 3, Min: 1, Max: 16
func (e *GstX264Enc) SetRef(value uint) error {
	return e.Element.SetProperty("ref", value)
}

func (e *GstX264Enc) GetRef() (uint, error) {
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

// Subpixel motion estimation and partition decision quality: 1=fast, 10=best
// Default: 1, Min: 1, Max: 10
func (e *GstX264Enc) SetSubme(value uint) error {
	return e.Element.SetProperty("subme", value)
}

func (e *GstX264Enc) GetSubme() (uint, error) {
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

type GstX264EncPreset string

const (
	GstX264EncPresetNone      GstX264EncPreset = "None"      // No preset
	GstX264EncPresetUltrafast GstX264EncPreset = "ultrafast" // ultrafast
	GstX264EncPresetSuperfast GstX264EncPreset = "superfast" // superfast
	GstX264EncPresetVeryfast  GstX264EncPreset = "veryfast"  // veryfast
	GstX264EncPresetFaster    GstX264EncPreset = "faster"    // faster
	GstX264EncPresetFast      GstX264EncPreset = "fast"      // fast
	GstX264EncPresetMedium    GstX264EncPreset = "medium"    // medium
	GstX264EncPresetSlow      GstX264EncPreset = "slow"      // slow
	GstX264EncPresetSlower    GstX264EncPreset = "slower"    // slower
	GstX264EncPresetVeryslow  GstX264EncPreset = "veryslow"  // veryslow
	GstX264EncPresetPlacebo   GstX264EncPreset = "placebo"   // placebo
)

type GstX264EncPsyTune string

const (
	GstX264EncPsyTuneNone      GstX264EncPsyTune = "none"      // No tuning
	GstX264EncPsyTuneFilm      GstX264EncPsyTune = "film"      // Film
	GstX264EncPsyTuneAnimation GstX264EncPsyTune = "animation" // Animation
	GstX264EncPsyTuneGrain     GstX264EncPsyTune = "grain"     // Grain
	GstX264EncPsyTunePsnr      GstX264EncPsyTune = "psnr"      // PSNR
	GstX264EncPsyTuneSsim      GstX264EncPsyTune = "ssim"      // SSIM
)

type GstX264EncTune int

const (
	GstX264EncTuneStillimage  GstX264EncTune = 0x00000001 // Still image
	GstX264EncTuneFastdecode  GstX264EncTune = 0x00000002 // Fast decode
	GstX264EncTuneZerolatency GstX264EncTune = 0x00000004 // Zero latency
)

type GstX264EncAnalyse int

const (
	GstX264EncAnalyseI4x4 GstX264EncAnalyse = 0x00000001 // i4x4
	GstX264EncAnalyseI8x8 GstX264EncAnalyse = 0x00000002 // i8x8
	GstX264EncAnalyseP8x8 GstX264EncAnalyse = 0x00000010 // p8x8
	GstX264EncAnalyseP4x4 GstX264EncAnalyse = 0x00000020 // p4x4
	GstX264EncAnalyseB8x8 GstX264EncAnalyse = 0x00000100 // b8x8
)

type GstX264EncFramePacking string

const (
	GstX264EncFramePackingAuto              GstX264EncFramePacking = "auto"               // Automatic (use incoming video information)
	GstX264EncFramePackingCheckerboard      GstX264EncFramePacking = "checkerboard"       // checkerboard - Left and Right pixels alternate in a checkerboard pattern
	GstX264EncFramePackingColumnInterleaved GstX264EncFramePacking = "column-interleaved" // column interleaved - Alternating pixel columns represent Left and Right views
	GstX264EncFramePackingRowInterleaved    GstX264EncFramePacking = "row-interleaved"    // row interleaved - Alternating pixel rows represent Left and Right views
	GstX264EncFramePackingSideBySide        GstX264EncFramePacking = "side-by-side"       // side by side - The left half of the frame contains the Left eye view, the right half the Right eye view
	GstX264EncFramePackingTopBottom         GstX264EncFramePacking = "top-bottom"         // top bottom - L is on top, R on bottom
	GstX264EncFramePackingFrameInterleaved  GstX264EncFramePacking = "frame-interleaved"  // frame interleaved - Each frame contains either Left or Right view alternately
)

type GstX264EncMe string

const (
	GstX264EncMeDia  GstX264EncMe = "dia"  // dia
	GstX264EncMeHex  GstX264EncMe = "hex"  // hex
	GstX264EncMeUmh  GstX264EncMe = "umh"  // umh
	GstX264EncMeEsa  GstX264EncMe = "esa"  // esa
	GstX264EncMeTesa GstX264EncMe = "tesa" // tesa
)

type GstX264EncPass string

const (
	GstX264EncPassCbr   GstX264EncPass = "cbr"   // Constant Bitrate Encoding
	GstX264EncPassQuant GstX264EncPass = "quant" // Constant Quantizer
	GstX264EncPassQual  GstX264EncPass = "qual"  // Constant Quality
	GstX264EncPassPass1 GstX264EncPass = "pass1" // VBR Encoding - Pass 1
	GstX264EncPassPass2 GstX264EncPass = "pass2" // VBR Encoding - Pass 2
	GstX264EncPassPass3 GstX264EncPass = "pass3" // VBR Encoding - Pass 3
)
