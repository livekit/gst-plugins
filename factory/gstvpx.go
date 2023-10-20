// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decode VP9 video streams
type GstVP9Dec struct {
	*GstVPXDec
}

func NewVP9Dec() (*GstVP9Dec, error) {
	e, err := gst.NewElement("vp9dec")
	if err != nil {
		return nil, err
	}
	return &GstVP9Dec{GstVPXDec: &GstVPXDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVP9DecWithName(name string) (*GstVP9Dec, error) {
	e, err := gst.NewElementWithName("vp9dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVP9Dec{GstVPXDec: &GstVPXDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Encode VP9 video streams
type GstVP9Enc struct {
	*GstVPXEnc
}

func NewVP9Enc() (*GstVP9Enc, error) {
	e, err := gst.NewElement("vp9enc")
	if err != nil {
		return nil, err
	}
	return &GstVP9Enc{GstVPXEnc: &GstVPXEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewVP9EncWithName(name string) (*GstVP9Enc, error) {
	e, err := gst.NewElementWithName("vp9enc", name)
	if err != nil {
		return nil, err
	}
	return &GstVP9Enc{GstVPXEnc: &GstVPXEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Which adaptive quantization mode should be used
// Default: off (0), Min: 0, Max: 4
func (e *GstVP9Enc) SetAqMode(value GstVPXAQ) error {
	e.Element.SetArg("aq-mode", string(value))
	return nil
}

func (e *GstVP9Enc) GetAqMode() (GstVPXAQ, error) {
	var value GstVPXAQ
	var ok bool
	v, err := e.Element.GetProperty("aq-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXAQ)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXAQ")
	}
	return value, nil
}

// Whether encoded bitstream should allow parallel processing of video frames in the decoder (default is on)
// Default: true
func (e *GstVP9Enc) SetFrameParallelDecoding(value bool) error {
	return e.Element.SetProperty("frame-parallel-decoding", value)
}

func (e *GstVP9Enc) GetFrameParallelDecoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("frame-parallel-decoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether each row should be encoded using multiple threads
// Default: false
func (e *GstVP9Enc) SetRowMt(value bool) error {
	return e.Element.SetProperty("row-mt", value)
}

func (e *GstVP9Enc) GetRowMt() (bool, error) {
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

// Number of tile columns, log2
// Default: 6, Min: 0, Max: 6
func (e *GstVP9Enc) SetTileColumns(value int) error {
	return e.Element.SetProperty("tile-columns", value)
}

func (e *GstVP9Enc) GetTileColumns() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of tile rows, log2
// Default: 0, Min: 0, Max: 2
func (e *GstVP9Enc) SetTileRows(value int) error {
	return e.Element.SetProperty("tile-rows", value)
}

func (e *GstVP9Enc) GetTileRows() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Decode VP8 video streams
type GstVP8Dec struct {
	*GstVPXDec
}

func NewVP8Dec() (*GstVP8Dec, error) {
	e, err := gst.NewElement("vp8dec")
	if err != nil {
		return nil, err
	}
	return &GstVP8Dec{GstVPXDec: &GstVPXDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVP8DecWithName(name string) (*GstVP8Dec, error) {
	e, err := gst.NewElementWithName("vp8dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVP8Dec{GstVPXDec: &GstVPXDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Encode VP8 video streams
type GstVP8Enc struct {
	*GstVPXEnc
}

func NewVP8Enc() (*GstVP8Enc, error) {
	e, err := gst.NewElement("vp8enc")
	if err != nil {
		return nil, err
	}
	return &GstVP8Enc{GstVPXEnc: &GstVPXEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewVP8EncWithName(name string) (*GstVP8Enc, error) {
	e, err := gst.NewElementWithName("vp8enc", name)
	if err != nil {
		return nil, err
	}
	return &GstVP8Enc{GstVPXEnc: &GstVPXEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

type GstVPXEnc struct {
	*GstVideoEncoder
}

// Datarate overshoot (max) target (%%)
// Default: 100, Min: 0, Max: 1000
func (e *GstVPXEnc) SetOvershoot(value int) error {
	return e.Element.SetProperty("overshoot", value)
}

func (e *GstVPXEnc) GetOvershoot() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overshoot")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of threads to use
// Default: 0, Min: 0, Max: 64
func (e *GstVPXEnc) SetThreads(value int) error {
	return e.Element.SetProperty("threads", value)
}

func (e *GstVPXEnc) GetThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// GOP maximum bitrate (%% target)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetTwopassVbrMaxsection(value int) error {
	return e.Element.SetProperty("twopass-vbr-maxsection", value)
}

func (e *GstVPXEnc) GetTwopassVbrMaxsection() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("twopass-vbr-maxsection")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Error resilience flags
// Default: (none)
func (e *GstVPXEnc) SetErrorResilient(value GstVPXEncErFlags) error {
	e.Element.SetArg("error-resilient", fmt.Sprint(value))
	return nil
}

func (e *GstVPXEnc) GetErrorResilient() (GstVPXEncErFlags, error) {
	var value GstVPXEncErFlags
	var ok bool
	v, err := e.Element.GetProperty("error-resilient")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncErFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncErFlags")
	}
	return value, nil
}

// Maximum Intra frame bitrate
// Default: 0, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetMaxIntraBitrate(value int) error {
	return e.Element.SetProperty("max-intra-bitrate", value)
}

func (e *GstVPXEnc) GetMaxIntraBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-intra-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Allow spatial resampling
// Default: false
func (e *GstVPXEnc) SetResizeAllowed(value bool) error {
	return e.Element.SetProperty("resize-allowed", value)
}

func (e *GstVPXEnc) GetResizeAllowed() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("resize-allowed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Length of sequence that defines layer membership periodicity
// Default: 0, Min: 0, Max: 16
func (e *GstVPXEnc) SetTemporalScalabilityPeriodicity(value int) error {
	return e.Element.SetProperty("temporal-scalability-periodicity", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityPeriodicity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("temporal-scalability-periodicity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of token partitions
// Default: 1 (0)
func (e *GstVPXEnc) SetTokenPartitions(value GstVPXEncTokenPartitions) error {
	e.Element.SetArg("token-partitions", string(value))
	return nil
}

func (e *GstVPXEnc) GetTokenPartitions() (GstVPXEncTokenPartitions, error) {
	var value GstVPXEncTokenPartitions
	var ok bool
	v, err := e.Element.GetProperty("token-partitions")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncTokenPartitions)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncTokenPartitions")
	}
	return value, nil
}

// Client buffer size (ms)
// Default: 6000, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstVPXEnc) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Temporal resampling threshold (buf %%)
// Default: 0, Min: 0, Max: 100
func (e *GstVPXEnc) SetDropframeThreshold(value int) error {
	return e.Element.SetProperty("dropframe-threshold", value)
}

func (e *GstVPXEnc) GetDropframeThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("dropframe-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Fraction of one second that is the shortest interframe time - normally left as zero which will default to the framerate
// Default: 0/1, Min: 0/1, Max: 2147483647/1
func (e *GstVPXEnc) SetTimebase(value interface{}) error {
	return e.Element.SetProperty("timebase", value)
}

func (e *GstVPXEnc) GetTimebase() (interface{}, error) {
	return e.Element.GetProperty("timebase")
}

// Datarate undershoot (min) target (%%)
// Default: 100, Min: 0, Max: 1000
func (e *GstVPXEnc) SetUndershoot(value int) error {
	return e.Element.SetProperty("undershoot", value)
}

func (e *GstVPXEnc) GetUndershoot() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("undershoot")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// AltRef strength
// Default: 3, Min: 0, Max: 6
func (e *GstVPXEnc) SetArnrStrength(value int) error {
	return e.Element.SetProperty("arnr-strength", value)
}

func (e *GstVPXEnc) GetArnrStrength() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("arnr-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Multipass cache file. If stream caps reinited, multiple files will be created: file, file.1, file.2, ... and so on.
// Default: multipass.cache
func (e *GstVPXEnc) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

func (e *GstVPXEnc) GetMultipassCacheFile() (string, error) {
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

// Sequence defining coding layer sync flags

func (e *GstVPXEnc) SetTemporalScalabilityLayerSyncFlags(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("temporal-scalability-layer-sync-flags", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityLayerSyncFlags() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("temporal-scalability-layer-sync-flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Target bitrates (bits/sec) for coding layers (one per layer)

func (e *GstVPXEnc) SetTemporalScalabilityTargetBitrate(value interface{}) error {
	return e.Element.SetProperty("temporal-scalability-target-bitrate", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityTargetBitrate() (interface{}, error) {
	return e.Element.GetProperty("temporal-scalability-target-bitrate")
}

// Tuning
// Default: psnr (0)
func (e *GstVPXEnc) SetTuning(value GstVPXEncTuning) error {
	e.Element.SetArg("tuning", string(value))
	return nil
}

func (e *GstVPXEnc) GetTuning() (GstVPXEncTuning, error) {
	var value GstVPXEncTuning
	var ok bool
	v, err := e.Element.GetProperty("tuning")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncTuning)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncTuning")
	}
	return value, nil
}

// Initial client buffer size (ms)
// Default: 4000, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetBufferInitialSize(value int) error {
	return e.Element.SetProperty("buffer-initial-size", value)
}

func (e *GstVPXEnc) GetBufferInitialSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-initial-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Downscale threshold (buf %%)
// Default: 60, Min: 0, Max: 100
func (e *GstVPXEnc) SetResizeDownThreshold(value int) error {
	return e.Element.SetProperty("resize-down-threshold", value)
}

func (e *GstVPXEnc) GetResizeDownThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("resize-down-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// CPU used
// Default: 0, Min: -16, Max: 16
func (e *GstVPXEnc) SetCpuUsed(value int) error {
	return e.Element.SetProperty("cpu-used", value)
}

func (e *GstVPXEnc) GetCpuUsed() (int, error) {
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

// Constrained quality level
// Default: 10, Min: 0, Max: 63
func (e *GstVPXEnc) SetCqLevel(value int) error {
	return e.Element.SetProperty("cq-level", value)
}

func (e *GstVPXEnc) GetCqLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cq-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Horizontal scaling mode
// Default: normal (0)
func (e *GstVPXEnc) SetHorizontalScalingMode(value GstVPXEncScalingMode) error {
	e.Element.SetArg("horizontal-scaling-mode", string(value))
	return nil
}

func (e *GstVPXEnc) GetHorizontalScalingMode() (GstVPXEncScalingMode, error) {
	var value GstVPXEncScalingMode
	var ok bool
	v, err := e.Element.GetProperty("horizontal-scaling-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncScalingMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncScalingMode")
	}
	return value, nil
}

// Keyframe placement
// Default: auto (1)
func (e *GstVPXEnc) SetKeyframeMode(value GstVPXEncKfMode) error {
	e.Element.SetArg("keyframe-mode", string(value))
	return nil
}

func (e *GstVPXEnc) GetKeyframeMode() (GstVPXEncKfMode, error) {
	var value GstVPXEncKfMode
	var ok bool
	v, err := e.Element.GetProperty("keyframe-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncKfMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncKfMode")
	}
	return value, nil
}

// Maximum Quantizer (worst)
// Default: 63, Min: 0, Max: 63
func (e *GstVPXEnc) SetMaxQuantizer(value int) error {
	return e.Element.SetProperty("max-quantizer", value)
}

func (e *GstVPXEnc) GetMaxQuantizer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum Quantizer (best)
// Default: 4, Min: 0, Max: 63
func (e *GstVPXEnc) SetMinQuantizer(value int) error {
	return e.Element.SetProperty("min-quantizer", value)
}

func (e *GstVPXEnc) GetMinQuantizer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-quantizer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// AltRef type
// Default: 3, Min: 1, Max: 3
func (e *GstVPXEnc) SetArnrType(value int) error {
	return e.Element.SetProperty("arnr-type", value)
}

func (e *GstVPXEnc) GetArnrType() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("arnr-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Automatically generate AltRef frames
// Default: false
func (e *GstVPXEnc) SetAutoAltRef(value bool) error {
	return e.Element.SetProperty("auto-alt-ref", value)
}

func (e *GstVPXEnc) GetAutoAltRef() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-alt-ref")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Sequence defining coding layer flags

func (e *GstVPXEnc) SetTemporalScalabilityLayerFlags(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("temporal-scalability-layer-flags", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityLayerFlags() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("temporal-scalability-layer-flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Sequence defining coding layer membership

func (e *GstVPXEnc) SetTemporalScalabilityLayerId(value interface{}) error {
	return e.Element.SetProperty("temporal-scalability-layer-id", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityLayerId() (interface{}, error) {
	return e.Element.GetProperty("temporal-scalability-layer-id")
}

// Rate decimation factors for each layer

func (e *GstVPXEnc) SetTemporalScalabilityRateDecimator(value interface{}) error {
	return e.Element.SetProperty("temporal-scalability-rate-decimator", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityRateDecimator() (interface{}, error) {
	return e.Element.GetProperty("temporal-scalability-rate-decimator")
}

// CBR/VBR bias (0=CBR, 100=VBR)
// Default: 50, Min: 0, Max: 100
func (e *GstVPXEnc) SetTwopassVbrBias(value int) error {
	return e.Element.SetProperty("twopass-vbr-bias", value)
}

func (e *GstVPXEnc) GetTwopassVbrBias() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("twopass-vbr-bias")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Multipass encode mode
// Default: one-pass (0)
func (e *GstVPXEnc) SetMultipassMode(value GstVPXEncMultipassMode) error {
	e.Element.SetArg("multipass-mode", string(value))
	return nil
}

func (e *GstVPXEnc) GetMultipassMode() (GstVPXEncMultipassMode, error) {
	var value GstVPXEncMultipassMode
	var ok bool
	v, err := e.Element.GetProperty("multipass-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncMultipassMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncMultipassMode")
	}
	return value, nil
}

// Filter sharpness
// Default: 0, Min: 0, Max: 7
func (e *GstVPXEnc) SetSharpness(value int) error {
	return e.Element.SetProperty("sharpness", value)
}

func (e *GstVPXEnc) GetSharpness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sharpness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Target bitrate (in bits/sec) (0: auto - bitrate depends on resolution, see "bits-per-pixel" property for more info)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetTargetBitrate(value int) error {
	return e.Element.SetProperty("target-bitrate", value)
}

func (e *GstVPXEnc) GetTargetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("target-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Optimal client buffer size (ms)
// Default: 5000, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetBufferOptimalSize(value int) error {
	return e.Element.SetProperty("buffer-optimal-size", value)
}

func (e *GstVPXEnc) GetBufferOptimalSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-optimal-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Motion detection threshold. Recommendation is to set 100 for screen/window sharing
// Default: 1, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetStaticThreshold(value int) error {
	return e.Element.SetProperty("static-threshold", value)
}

func (e *GstVPXEnc) GetStaticThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("static-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// GOP minimum bitrate (%% target)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetTwopassVbrMinsection(value int) error {
	return e.Element.SetProperty("twopass-vbr-minsection", value)
}

func (e *GstVPXEnc) GetTwopassVbrMinsection() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("twopass-vbr-minsection")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Deadline per frame (usec, 0=best, 1=realtime)
// Default: 1000000, Min: 0, Max: 9223372036854775807
func (e *GstVPXEnc) SetDeadline(value int64) error {
	return e.Element.SetProperty("deadline", value)
}

func (e *GstVPXEnc) GetDeadline() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Maximum distance between keyframes (number of frames)
// Default: 128, Min: 0, Max: 2147483647
func (e *GstVPXEnc) SetKeyframeMaxDist(value int) error {
	return e.Element.SetProperty("keyframe-max-dist", value)
}

func (e *GstVPXEnc) GetKeyframeMaxDist() (int, error) {
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

// Rate control mode
// Default: vbr (0)
func (e *GstVPXEnc) SetEndUsage(value GstVPXEncEndUsage) error {
	e.Element.SetArg("end-usage", string(value))
	return nil
}

func (e *GstVPXEnc) GetEndUsage() (GstVPXEncEndUsage, error) {
	var value GstVPXEncEndUsage
	var ok bool
	v, err := e.Element.GetProperty("end-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncEndUsage)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncEndUsage")
	}
	return value, nil
}

// Maximum number of frames to lag
// Default: 0, Min: 0, Max: 25
func (e *GstVPXEnc) SetLagInFrames(value int) error {
	return e.Element.SetProperty("lag-in-frames", value)
}

func (e *GstVPXEnc) GetLagInFrames() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("lag-in-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Noise sensisivity (frames to blur)
// Default: 0, Min: 0, Max: 6
func (e *GstVPXEnc) SetNoiseSensitivity(value int) error {
	return e.Element.SetProperty("noise-sensitivity", value)
}

func (e *GstVPXEnc) GetNoiseSensitivity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("noise-sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Upscale threshold (buf %%)
// Default: 30, Min: 0, Max: 100
func (e *GstVPXEnc) SetResizeUpThreshold(value int) error {
	return e.Element.SetProperty("resize-up-threshold", value)
}

func (e *GstVPXEnc) GetResizeUpThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("resize-up-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of coding layers to use
// Default: 1, Min: 1, Max: 5
func (e *GstVPXEnc) SetTemporalScalabilityNumberLayers(value int) error {
	return e.Element.SetProperty("temporal-scalability-number-layers", value)
}

func (e *GstVPXEnc) GetTemporalScalabilityNumberLayers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("temporal-scalability-number-layers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Vertical scaling mode
// Default: normal (0)
func (e *GstVPXEnc) SetVerticalScalingMode(value GstVPXEncScalingMode) error {
	e.Element.SetArg("vertical-scaling-mode", string(value))
	return nil
}

func (e *GstVPXEnc) GetVerticalScalingMode() (GstVPXEncScalingMode, error) {
	var value GstVPXEncScalingMode
	var ok bool
	v, err := e.Element.GetProperty("vertical-scaling-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXEncScalingMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXEncScalingMode")
	}
	return value, nil
}

// AltRef maximum number of frames
// Default: 0, Min: 0, Max: 15
func (e *GstVPXEnc) SetArnrMaxframes(value int) error {
	return e.Element.SetProperty("arnr-maxframes", value)
}

func (e *GstVPXEnc) GetArnrMaxframes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("arnr-maxframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Factor to convert number of pixels to bitrate value (only has an effect if target-bitrate=0)
// Default: 0.0434, Min: 0, Max: 3.40282e+38
func (e *GstVPXEnc) SetBitsPerPixel(value float32) error {
	return e.Element.SetProperty("bits-per-pixel", value)
}

func (e *GstVPXEnc) GetBitsPerPixel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("bits-per-pixel")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

type GstVPXEncEndUsage string

const (
	GstVPXEncEndUsageVbr GstVPXEncEndUsage = "vbr" // Variable Bit Rate (VBR) mode
	GstVPXEncEndUsageCbr GstVPXEncEndUsage = "cbr" // Constant Bit Rate (CBR) mode
	GstVPXEncEndUsageCq  GstVPXEncEndUsage = "cq"  // Constant Quality Mode (CQ) mode
)

type GstVPXEncErFlags int

const (
	GstVPXEncErFlagsDefault    GstVPXEncErFlags = 0x00000001 // Default error resilience
	GstVPXEncErFlagsPartitions GstVPXEncErFlags = 0x00000002 // Allow partitions to be decoded independently
)

type GstVPXEncKfMode string

const (
	GstVPXEncKfModeAuto     GstVPXEncKfMode = "auto"     // Determine optimal placement automatically
	GstVPXEncKfModeDisabled GstVPXEncKfMode = "disabled" // Don't automatically place keyframes
)

type GstVPXEncMultipassMode string

const (
	GstVPXEncMultipassModeOnePass   GstVPXEncMultipassMode = "one-pass"   // One pass encoding (default)
	GstVPXEncMultipassModeFirstPass GstVPXEncMultipassMode = "first-pass" // First pass of multipass encoding
	GstVPXEncMultipassModeLastPass  GstVPXEncMultipassMode = "last-pass"  // Last pass of multipass encoding
)

type GstVPXAQ string

const (
	GstVPXAQOff           GstVPXAQ = "off"            // GST_VPX_AQ_OFF
	GstVPXAQVariance      GstVPXAQ = "variance"       // GST_VPX_AQ_VARIANCE
	GstVPXAQComplexity    GstVPXAQ = "complexity"     // GST_VPX_AQ_COMPLEXITY
	GstVPXAQCyclicRefresh GstVPXAQ = "cyclic-refresh" // GST_VPX_AQ_CYCLIC_REFRESH
	GstVPXAQEquator360    GstVPXAQ = "equator360"     // GST_VPX_AQ_EQUATOR360
	GstVPXAQPerceptual    GstVPXAQ = "perceptual"     // GST_VPX_AQ_PERCEPTUAL
	GstVPXAQPsnr          GstVPXAQ = "psnr"           // GST_VPX_AQ_PSNR
	GstVPXAQLookahead     GstVPXAQ = "lookahead"      // GST_VPX_AQ_LOOKAHEAD
)

type GstVPXDecPostProcessingFlags int

const (
	GstVPXDecPostProcessingFlagsDeblock      GstVPXDecPostProcessingFlags = 0x00000001 // Deblock
	GstVPXDecPostProcessingFlagsDemacroblock GstVPXDecPostProcessingFlags = 0x00000002 // Demacroblock
	GstVPXDecPostProcessingFlagsAddnoise     GstVPXDecPostProcessingFlags = 0x00000004 // Add noise
	GstVPXDecPostProcessingFlagsMfqe         GstVPXDecPostProcessingFlags = 0x00000008 // Multi-frame quality enhancement
)

type GstVPXEncTokenPartitions string

const (
	GstVPXEncTokenPartitions1 GstVPXEncTokenPartitions = "1" // One token partition
	GstVPXEncTokenPartitions2 GstVPXEncTokenPartitions = "2" // Two token partitions
	GstVPXEncTokenPartitions4 GstVPXEncTokenPartitions = "4" // Four token partitions
	GstVPXEncTokenPartitions8 GstVPXEncTokenPartitions = "8" // Eight token partitions
)

type GstVPXEncTuning string

const (
	GstVPXEncTuningPsnr GstVPXEncTuning = "psnr" // Tune for PSNR
	GstVPXEncTuningSsim GstVPXEncTuning = "ssim" // Tune for SSIM
)

type GstVPXDec struct {
	*GstVideoDecoder
}

// Maximum number of decoding threads
// Default: 0, Min: 0, Max: 16
func (e *GstVPXDec) SetThreads(value uint) error {
	return e.Element.SetProperty("threads", value)
}

func (e *GstVPXDec) GetThreads() (uint, error) {
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

// Deblocking level
// Default: 4, Min: 0, Max: 16
func (e *GstVPXDec) SetDeblockingLevel(value uint) error {
	return e.Element.SetProperty("deblocking-level", value)
}

func (e *GstVPXDec) GetDeblockingLevel() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("deblocking-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Noise level
// Default: 0, Min: 0, Max: 16
func (e *GstVPXDec) SetNoiseLevel(value uint) error {
	return e.Element.SetProperty("noise-level", value)
}

func (e *GstVPXDec) GetNoiseLevel() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("noise-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable post processing
// Default: false
func (e *GstVPXDec) SetPostProcessing(value bool) error {
	return e.Element.SetProperty("post-processing", value)
}

func (e *GstVPXDec) GetPostProcessing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-processing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Flags to control post processing
// Default: mfqe+demacroblock+deblock
func (e *GstVPXDec) SetPostProcessingFlags(value GstVPXDecPostProcessingFlags) error {
	e.Element.SetArg("post-processing-flags", fmt.Sprint(value))
	return nil
}

func (e *GstVPXDec) GetPostProcessingFlags() (GstVPXDecPostProcessingFlags, error) {
	var value GstVPXDecPostProcessingFlags
	var ok bool
	v, err := e.Element.GetProperty("post-processing-flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVPXDecPostProcessingFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVPXDecPostProcessingFlags")
	}
	return value, nil
}

type GstVPXEncScalingMode string

const (
	GstVPXEncScalingModeNormal GstVPXEncScalingMode = "normal" // Normal
	GstVPXEncScalingMode45     GstVPXEncScalingMode = "4:5"    // 4:5
	GstVPXEncScalingMode35     GstVPXEncScalingMode = "3:5"    // 3:5
	GstVPXEncScalingMode12     GstVPXEncScalingMode = "1:2"    // 1:2
)
