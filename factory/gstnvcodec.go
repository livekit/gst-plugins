// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Encode H.264 video streams using NVCODEC API CUDA Mode
type GstNvCudaH264Enc struct {
	*GstNvEncoder
}

func NewNvCudaH264Enc() (*GstNvCudaH264Enc, error) {
	e, err := gst.NewElement("nvcudah264enc")
	if err != nil {
		return nil, err
	}
	return &GstNvCudaH264Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewNvCudaH264EncWithName(name string) (*GstNvCudaH264Enc, error) {
	e, err := gst.NewElementWithName("nvcudah264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvCudaH264Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Minimum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstNvCudaH264Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvCudaH264Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvCudaH264Enc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enables Weighted Prediction
// Default: false
func (e *GstNvCudaH264Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvCudaH264Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstNvCudaH264Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstNvCudaH264Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvCudaH264Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvCudaH264Enc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvCudaH264Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvCudaH264Enc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// CUDA device ID of associated GPU
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvCudaH264Enc) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvCudaH264Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvCudaH264Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Insert sequence headers (SPS/PPS) per IDR
// Default: false
func (e *GstNvCudaH264Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

func (e *GstNvCudaH264Enc) GetRepeatSequenceHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat-sequence-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvCudaH264Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvCudaH264Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Zero latency operation (i.e., num_reorder_frames = 0)
// Default: false
func (e *GstNvCudaH264Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

func (e *GstNvCudaH264Enc) GetZeroReorderDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-reorder-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 30, Min: -1, Max: 2147483647
func (e *GstNvCudaH264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvCudaH264Enc) GetGopSize() (int, error) {
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

// Enable CABAC entropy coding
// Default: true
func (e *GstNvCudaH264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstNvCudaH264Enc) GetCabac() (bool, error) {
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

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvCudaH264Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvCudaH264Enc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstNvCudaH264Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvCudaH264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvCudaH264Enc) GetAud() (bool, error) {
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

// Constant QP value for B frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstNvCudaH264Enc) GetQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for I frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstNvCudaH264Enc) GetQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for P frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstNvCudaH264Enc) GetQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvCudaH264Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvCudaH264Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum Bitrate in kbit/sec (ignored in CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvCudaH264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvCudaH264Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 4
func (e *GstNvCudaH264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstNvCudaH264Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Bitrate in kbit/sec (0 = automatic)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvCudaH264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvCudaH264Enc) GetBitrate() (uint, error) {
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

// Maximum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstNvCudaH264Enc) GetMaxQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH264Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstNvCudaH264Enc) GetMinQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding Preset
// Default: default (0)
func (e *GstNvCudaH264Enc) SetPreset(value GstNvEncoderPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvCudaH264Enc) GetPreset() (GstNvEncoderPreset, error) {
	var value GstNvEncoderPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderPreset")
	}
	return value, nil
}

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvCudaH264Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvCudaH264Enc) GetBAdapt() (bool, error) {
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

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvCudaH264Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvCudaH264Enc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (1)
func (e *GstNvCudaH264Enc) SetRateControl(value GstNvEncoderRCMode) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstNvCudaH264Enc) GetRateControl() (GstNvEncoderRCMode, error) {
	var value GstNvEncoderRCMode
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderRCMode")
	}
	return value, nil
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvCudaH264Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvCudaH264Enc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encode H.265 video streams using NVCODEC API CUDA Mode
type GstNvCudaH265Enc struct {
	*GstNvEncoder
}

func NewNvCudaH265Enc() (*GstNvCudaH265Enc, error) {
	e, err := gst.NewElement("nvcudah265enc")
	if err != nil {
		return nil, err
	}
	return &GstNvCudaH265Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewNvCudaH265EncWithName(name string) (*GstNvCudaH265Enc, error) {
	e, err := gst.NewElementWithName("nvcudah265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvCudaH265Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Maximum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstNvCudaH265Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for B frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstNvCudaH265Enc) GetQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvCudaH265Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvCudaH265Enc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvCudaH265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvCudaH265Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvCudaH265Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvCudaH265Enc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Bitrate in kbit/sec (0 = automatic)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvCudaH265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvCudaH265Enc) GetBitrate() (uint, error) {
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

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 5
func (e *GstNvCudaH265Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstNvCudaH265Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvCudaH265Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvCudaH265Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvCudaH265Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvCudaH265Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstNvCudaH265Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstNvCudaH265Enc) GetMaxQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstNvCudaH265Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for I frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstNvCudaH265Enc) GetQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvCudaH265Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvCudaH265Enc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enables Weighted Prediction
// Default: false
func (e *GstNvCudaH265Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvCudaH265Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum Bitrate in kbit/sec (ignored in CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvCudaH265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvCudaH265Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (1)
func (e *GstNvCudaH265Enc) SetRateControl(value GstNvEncoderRCMode) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstNvCudaH265Enc) GetRateControl() (GstNvEncoderRCMode, error) {
	var value GstNvEncoderRCMode
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderRCMode")
	}
	return value, nil
}

// Minimum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstNvCudaH265Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvCudaH265Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvCudaH265Enc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Zero latency operation (i.e., num_reorder_frames = 0)
// Default: false
func (e *GstNvCudaH265Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

func (e *GstNvCudaH265Enc) GetZeroReorderDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-reorder-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvCudaH265Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvCudaH265Enc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvCudaH265Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvCudaH265Enc) GetBAdapt() (bool, error) {
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

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvCudaH265Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvCudaH265Enc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// CUDA device ID of associated GPU
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvCudaH265Enc) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 30, Min: -1, Max: 2147483647
func (e *GstNvCudaH265Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvCudaH265Enc) GetGopSize() (int, error) {
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

// Minimum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstNvCudaH265Enc) GetMinQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding Preset
// Default: default (0)
func (e *GstNvCudaH265Enc) SetPreset(value GstNvEncoderPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvCudaH265Enc) GetPreset() (GstNvEncoderPreset, error) {
	var value GstNvEncoderPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderPreset")
	}
	return value, nil
}

// Constant QP value for P frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvCudaH265Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstNvCudaH265Enc) GetQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvCudaH265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvCudaH265Enc) GetAud() (bool, error) {
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

// Insert sequence headers (SPS/PPS) per IDR, ignored if negotiated stream-format is "hvc1"
// Default: false
func (e *GstNvCudaH265Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

func (e *GstNvCudaH265Enc) GetRepeatSequenceHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat-sequence-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encode HEVC video streams using NVIDIA's hardware-accelerated NVENC encoder API
type GstNvH265Enc struct {
	*GstNvDevice0H265Enc
}

func NewNvH265Enc() (*GstNvH265Enc, error) {
	e, err := gst.NewElement("nvh265enc")
	if err != nil {
		return nil, err
	}
	return &GstNvH265Enc{GstNvDevice0H265Enc: &GstNvDevice0H265Enc{GstNvBaseEnc: &GstNvBaseEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}}, nil
}

func NewNvH265EncWithName(name string) (*GstNvH265Enc, error) {
	e, err := gst.NewElementWithName("nvh265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvH265Enc{GstNvDevice0H265Enc: &GstNvDevice0H265Enc{GstNvBaseEnc: &GstNvBaseEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}}, nil
}

// Weighted Prediction
// Default: false
func (e *GstNvH265Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvH265Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvH265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvH265Enc) GetAud() (bool, error) {
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

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvH265Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvH265Enc) GetBAdapt() (bool, error) {
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

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 5
func (e *GstNvH265Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

func (e *GstNvH265Enc) GetBframes() (uint, error) {
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

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvH265Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvH265Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvH265Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvH265Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvH265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvH265Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Resize video using CUDA
type GstCudaScale struct {
	*GstCudaBaseConvert
}

func NewCudaScale() (*GstCudaScale, error) {
	e, err := gst.NewElement("cudascale")
	if err != nil {
		return nil, err
	}
	return &GstCudaScale{GstCudaBaseConvert: &GstCudaBaseConvert{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCudaScaleWithName(name string) (*GstCudaScale, error) {
	e, err := gst.NewElementWithName("cudascale", name)
	if err != nil {
		return nil, err
	}
	return &GstCudaScale{GstCudaBaseConvert: &GstCudaBaseConvert{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Add black borders if necessary to keep the display aspect ratio
// Default: true
func (e *GstCudaScale) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

func (e *GstCudaScale) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encode H.264 video streams using NVCODEC API auto GPU select Mode
type GstNvAutoGpuH264Enc struct {
	*GstNvEncoder
}

func NewNvAutoGpuH264Enc() (*GstNvAutoGpuH264Enc, error) {
	e, err := gst.NewElement("nvautogpuh264enc")
	if err != nil {
		return nil, err
	}
	return &GstNvAutoGpuH264Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewNvAutoGpuH264EncWithName(name string) (*GstNvAutoGpuH264Enc, error) {
	e, err := gst.NewElementWithName("nvautogpuh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvAutoGpuH264Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Maximum Bitrate in kbit/sec (ignored in CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvAutoGpuH264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvAutoGpuH264Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstNvAutoGpuH264Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvAutoGpuH264Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvAutoGpuH264Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Zero latency operation (i.e., num_reorder_frames = 0)
// Default: false
func (e *GstNvAutoGpuH264Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

func (e *GstNvAutoGpuH264Enc) GetZeroReorderDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-reorder-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvAutoGpuH264Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvAutoGpuH264Enc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvAutoGpuH264Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvAutoGpuH264Enc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvAutoGpuH264Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvAutoGpuH264Enc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 30, Min: -1, Max: 2147483647
func (e *GstNvAutoGpuH264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvAutoGpuH264Enc) GetGopSize() (int, error) {
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

// Maximum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstNvAutoGpuH264Enc) GetMaxQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstNvAutoGpuH264Enc) GetMinQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for B frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstNvAutoGpuH264Enc) GetQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for I frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstNvAutoGpuH264Enc) GetQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (1)
func (e *GstNvAutoGpuH264Enc) SetRateControl(value GstNvEncoderRCMode) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstNvAutoGpuH264Enc) GetRateControl() (GstNvEncoderRCMode, error) {
	var value GstNvEncoderRCMode
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderRCMode")
	}
	return value, nil
}

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvAutoGpuH264Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvAutoGpuH264Enc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvAutoGpuH264Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvAutoGpuH264Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable CABAC entropy coding
// Default: true
func (e *GstNvAutoGpuH264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstNvAutoGpuH264Enc) GetCabac() (bool, error) {
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

// Maximum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstNvAutoGpuH264Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstNvAutoGpuH264Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding Preset
// Default: default (0)
func (e *GstNvAutoGpuH264Enc) SetPreset(value GstNvEncoderPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvAutoGpuH264Enc) GetPreset() (GstNvEncoderPreset, error) {
	var value GstNvEncoderPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderPreset")
	}
	return value, nil
}

// Constant QP value for P frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstNvAutoGpuH264Enc) GetQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvAutoGpuH264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvAutoGpuH264Enc) GetAud() (bool, error) {
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

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvAutoGpuH264Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvAutoGpuH264Enc) GetBAdapt() (bool, error) {
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

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 4
func (e *GstNvAutoGpuH264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstNvAutoGpuH264Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH264Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstNvAutoGpuH264Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Insert sequence headers (SPS/PPS) per IDR
// Default: false
func (e *GstNvAutoGpuH264Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

func (e *GstNvAutoGpuH264Enc) GetRepeatSequenceHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat-sequence-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvAutoGpuH264Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvAutoGpuH264Enc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvAutoGpuH264Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvAutoGpuH264Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enables Weighted Prediction
// Default: false
func (e *GstNvAutoGpuH264Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvAutoGpuH264Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) to use
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstNvAutoGpuH264Enc) SetAdapterLuid(value int64) error {
	return e.Element.SetProperty("adapter-luid", value)
}

func (e *GstNvAutoGpuH264Enc) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Bitrate in kbit/sec (0 = automatic)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvAutoGpuH264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvAutoGpuH264Enc) GetBitrate() (uint, error) {
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

// CUDA device ID to use
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvAutoGpuH264Enc) SetCudaDeviceId(value uint) error {
	return e.Element.SetProperty("cuda-device-id", value)
}

func (e *GstNvAutoGpuH264Enc) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvAutoGpuH264Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvAutoGpuH264Enc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// NVIDIA AV1 video decoder
type GstNvAV1Dec struct {
	*GstAV1Decoder
}

func NewNvAV1Dec() (*GstNvAV1Dec, error) {
	e, err := gst.NewElement("nvav1dec")
	if err != nil {
		return nil, err
	}
	return &GstNvAV1Dec{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvAV1DecWithName(name string) (*GstNvAV1Dec, error) {
	e, err := gst.NewElementWithName("nvav1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstNvAV1Dec{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Assigned CUDA device id
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvAV1Dec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// NVDEC video decoder
type Nvmpeg2videodec struct {
	*GstNvDec
}

func NewNvmpeg2videodec() (*Nvmpeg2videodec, error) {
	e, err := gst.NewElement("nvmpeg2videodec")
	if err != nil {
		return nil, err
	}
	return &Nvmpeg2videodec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvmpeg2videodecWithName(name string) (*Nvmpeg2videodec, error) {
	e, err := gst.NewElementWithName("nvmpeg2videodec", name)
	if err != nil {
		return nil, err
	}
	return &Nvmpeg2videodec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// NVIDIA VP8 video decoder
type GstNvVp8SLDec struct {
	*GstVp8Decoder
}

func NewNvVp8SLDec() (*GstNvVp8SLDec, error) {
	e, err := gst.NewElement("nvvp8sldec")
	if err != nil {
		return nil, err
	}
	return &GstNvVp8SLDec{GstVp8Decoder: &GstVp8Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvVp8SLDecWithName(name string) (*GstNvVp8SLDec, error) {
	e, err := gst.NewElementWithName("nvvp8sldec", name)
	if err != nil {
		return nil, err
	}
	return &GstNvVp8SLDec{GstVp8Decoder: &GstVp8Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Assigned CUDA device id
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvVp8SLDec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encode H.264 video streams using NVCODEC API Direct3D11 Mode
type GstNvD3D11H264Enc struct {
	*GstNvEncoder
}

func NewNvD3D11H264Enc() (*GstNvD3D11H264Enc, error) {
	e, err := gst.NewElement("nvd3d11h264enc")
	if err != nil {
		return nil, err
	}
	return &GstNvD3D11H264Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewNvD3D11H264EncWithName(name string) (*GstNvD3D11H264Enc, error) {
	e, err := gst.NewElementWithName("nvd3d11h264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvD3D11H264Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Maximum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstNvD3D11H264Enc) GetMaxQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Enables Weighted Prediction
// Default: false
func (e *GstNvD3D11H264Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvD3D11H264Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvD3D11H264Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvD3D11H264Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Zero latency operation (i.e., num_reorder_frames = 0)
// Default: false
func (e *GstNvD3D11H264Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

func (e *GstNvD3D11H264Enc) GetZeroReorderDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-reorder-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of associated GPU
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstNvD3D11H264Enc) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Constant QP value for P frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstNvD3D11H264Enc) GetQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Insert sequence headers (SPS/PPS) per IDR
// Default: false
func (e *GstNvD3D11H264Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

func (e *GstNvD3D11H264Enc) GetRepeatSequenceHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat-sequence-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvD3D11H264Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvD3D11H264Enc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvD3D11H264Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvD3D11H264Enc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum Bitrate in kbit/sec (ignored in CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvD3D11H264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvD3D11H264Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvD3D11H264Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvD3D11H264Enc) GetBAdapt() (bool, error) {
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

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvD3D11H264Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvD3D11H264Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvD3D11H264Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvD3D11H264Enc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvD3D11H264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvD3D11H264Enc) GetAud() (bool, error) {
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

// Enable CABAC entropy coding
// Default: true
func (e *GstNvD3D11H264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstNvD3D11H264Enc) GetCabac() (bool, error) {
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

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvD3D11H264Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvD3D11H264Enc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Bitrate in kbit/sec (0 = automatic)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvD3D11H264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvD3D11H264Enc) GetBitrate() (uint, error) {
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

// Maximum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstNvD3D11H264Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvD3D11H264Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvD3D11H264Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encoding Preset
// Default: default (0)
func (e *GstNvD3D11H264Enc) SetPreset(value GstNvEncoderPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvD3D11H264Enc) GetPreset() (GstNvEncoderPreset, error) {
	var value GstNvEncoderPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderPreset")
	}
	return value, nil
}

// Constant QP value for I frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstNvD3D11H264Enc) GetQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (1)
func (e *GstNvD3D11H264Enc) SetRateControl(value GstNvEncoderRCMode) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstNvD3D11H264Enc) GetRateControl() (GstNvEncoderRCMode, error) {
	var value GstNvEncoderRCMode
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderRCMode")
	}
	return value, nil
}

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 4
func (e *GstNvD3D11H264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstNvD3D11H264Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstNvD3D11H264Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstNvD3D11H264Enc) GetMinQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstNvD3D11H264Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvD3D11H264Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvD3D11H264Enc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Constant QP value for B frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstNvD3D11H264Enc) GetQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvD3D11H264Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvD3D11H264Enc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 30, Min: -1, Max: 2147483647
func (e *GstNvD3D11H264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvD3D11H264Enc) GetGopSize() (int, error) {
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

// Minimum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H264Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstNvD3D11H264Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encode H.264 video streams using NVIDIA's hardware-accelerated NVENC encoder API
type GstNvH264Enc struct {
	*GstNvDevice0H264Enc
}

func NewNvH264Enc() (*GstNvH264Enc, error) {
	e, err := gst.NewElement("nvh264enc")
	if err != nil {
		return nil, err
	}
	return &GstNvH264Enc{GstNvDevice0H264Enc: &GstNvDevice0H264Enc{GstNvBaseEnc: &GstNvBaseEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}}, nil
}

func NewNvH264EncWithName(name string) (*GstNvH264Enc, error) {
	e, err := gst.NewElementWithName("nvh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvH264Enc{GstNvDevice0H264Enc: &GstNvDevice0H264Enc{GstNvBaseEnc: &GstNvBaseEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}}, nil
}

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvH264Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvH264Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvH264Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvH264Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvH264Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvH264Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Weighted Prediction
// Default: false
func (e *GstNvH264Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvH264Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvH264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvH264Enc) GetAud() (bool, error) {
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

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvH264Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvH264Enc) GetBAdapt() (bool, error) {
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

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 4
func (e *GstNvH264Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

func (e *GstNvH264Enc) GetBframes() (uint, error) {
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

// NVIDIA H.264 video decoder
type GstNvH264SLDec struct {
	*GstH264Decoder
}

func NewNvH264SLDec() (*GstNvH264SLDec, error) {
	e, err := gst.NewElement("nvh264sldec")
	if err != nil {
		return nil, err
	}
	return &GstNvH264SLDec{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvH264SLDecWithName(name string) (*GstNvH264SLDec, error) {
	e, err := gst.NewElementWithName("nvh264sldec", name)
	if err != nil {
		return nil, err
	}
	return &GstNvH264SLDec{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Assigned CUDA device id
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvH264SLDec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// NVDEC video decoder
type Nvjpegdec struct {
	*GstNvDec
}

func NewNvjpegdec() (*Nvjpegdec, error) {
	e, err := gst.NewElement("nvjpegdec")
	if err != nil {
		return nil, err
	}
	return &Nvjpegdec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvjpegdecWithName(name string) (*Nvjpegdec, error) {
	e, err := gst.NewElementWithName("nvjpegdec", name)
	if err != nil {
		return nil, err
	}
	return &Nvjpegdec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// NVDEC video decoder
type Nvvp9dec struct {
	*GstNvDec
}

func NewNvvp9dec() (*Nvvp9dec, error) {
	e, err := gst.NewElement("nvvp9dec")
	if err != nil {
		return nil, err
	}
	return &Nvvp9dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvvp9decWithName(name string) (*Nvvp9dec, error) {
	e, err := gst.NewElementWithName("nvvp9dec", name)
	if err != nil {
		return nil, err
	}
	return &Nvvp9dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Resizes video and allow color conversion using CUDA
type GstCudaConvertScale struct {
	*GstCudaBaseConvert
}

func NewCudaConvertScale() (*GstCudaConvertScale, error) {
	e, err := gst.NewElement("cudaconvertscale")
	if err != nil {
		return nil, err
	}
	return &GstCudaConvertScale{GstCudaBaseConvert: &GstCudaBaseConvert{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCudaConvertScaleWithName(name string) (*GstCudaConvertScale, error) {
	e, err := gst.NewElementWithName("cudaconvertscale", name)
	if err != nil {
		return nil, err
	}
	return &GstCudaConvertScale{GstCudaBaseConvert: &GstCudaBaseConvert{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Add black borders if necessary to keep the display aspect ratio
// Default: true
func (e *GstCudaConvertScale) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

func (e *GstCudaConvertScale) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Downloads data from NVIDA GPU via CUDA APIs
type GstCudaDownload struct {
	*GstCudaMemoryCopy
}

func NewCudaDownload() (*GstCudaDownload, error) {
	e, err := gst.NewElement("cudadownload")
	if err != nil {
		return nil, err
	}
	return &GstCudaDownload{GstCudaMemoryCopy: &GstCudaMemoryCopy{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCudaDownloadWithName(name string) (*GstCudaDownload, error) {
	e, err := gst.NewElementWithName("cudadownload", name)
	if err != nil {
		return nil, err
	}
	return &GstCudaDownload{GstCudaMemoryCopy: &GstCudaMemoryCopy{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Uploads data into NVIDA GPU via CUDA APIs
type GstCudaUpload struct {
	*GstCudaMemoryCopy
}

func NewCudaUpload() (*GstCudaUpload, error) {
	e, err := gst.NewElement("cudaupload")
	if err != nil {
		return nil, err
	}
	return &GstCudaUpload{GstCudaMemoryCopy: &GstCudaMemoryCopy{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCudaUploadWithName(name string) (*GstCudaUpload, error) {
	e, err := gst.NewElementWithName("cudaupload", name)
	if err != nil {
		return nil, err
	}
	return &GstCudaUpload{GstCudaMemoryCopy: &GstCudaMemoryCopy{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// NVDEC video decoder
type Nvh264dec struct {
	*GstNvDec
}

func NewNvh264dec() (*Nvh264dec, error) {
	e, err := gst.NewElement("nvh264dec")
	if err != nil {
		return nil, err
	}
	return &Nvh264dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvh264decWithName(name string) (*Nvh264dec, error) {
	e, err := gst.NewElementWithName("nvh264dec", name)
	if err != nil {
		return nil, err
	}
	return &Nvh264dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// NVDEC video decoder
type Nvh265dec struct {
	*GstNvDec
}

func NewNvh265dec() (*Nvh265dec, error) {
	e, err := gst.NewElement("nvh265dec")
	if err != nil {
		return nil, err
	}
	return &Nvh265dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvh265decWithName(name string) (*Nvh265dec, error) {
	e, err := gst.NewElementWithName("nvh265dec", name)
	if err != nil {
		return nil, err
	}
	return &Nvh265dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// NVIDIA H.265 video decoder
type GstNvH265SLDec struct {
	*GstH265Decoder
}

func NewNvH265SLDec() (*GstNvH265SLDec, error) {
	e, err := gst.NewElement("nvh265sldec")
	if err != nil {
		return nil, err
	}
	return &GstNvH265SLDec{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvH265SLDecWithName(name string) (*GstNvH265SLDec, error) {
	e, err := gst.NewElementWithName("nvh265sldec", name)
	if err != nil {
		return nil, err
	}
	return &GstNvH265SLDec{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Assigned CUDA device id
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvH265SLDec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// NVDEC video decoder
type Nvmpeg4videodec struct {
	*GstNvDec
}

func NewNvmpeg4videodec() (*Nvmpeg4videodec, error) {
	e, err := gst.NewElement("nvmpeg4videodec")
	if err != nil {
		return nil, err
	}
	return &Nvmpeg4videodec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvmpeg4videodecWithName(name string) (*Nvmpeg4videodec, error) {
	e, err := gst.NewElementWithName("nvmpeg4videodec", name)
	if err != nil {
		return nil, err
	}
	return &Nvmpeg4videodec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// NVDEC video decoder
type Nvmpegvideodec struct {
	*GstNvDec
}

func NewNvmpegvideodec() (*Nvmpegvideodec, error) {
	e, err := gst.NewElement("nvmpegvideodec")
	if err != nil {
		return nil, err
	}
	return &Nvmpegvideodec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvmpegvideodecWithName(name string) (*Nvmpegvideodec, error) {
	e, err := gst.NewElementWithName("nvmpegvideodec", name)
	if err != nil {
		return nil, err
	}
	return &Nvmpegvideodec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Converts video from one colorspace to another using CUDA
type GstCudaConvert struct {
	*GstCudaBaseConvert
}

func NewCudaConvert() (*GstCudaConvert, error) {
	e, err := gst.NewElement("cudaconvert")
	if err != nil {
		return nil, err
	}
	return &GstCudaConvert{GstCudaBaseConvert: &GstCudaBaseConvert{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCudaConvertWithName(name string) (*GstCudaConvert, error) {
	e, err := gst.NewElementWithName("cudaconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstCudaConvert{GstCudaBaseConvert: &GstCudaBaseConvert{GstCudaBaseTransform: &GstCudaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Encode H.265 video streams using NVCODEC API auto GPU select Mode
type GstNvAutoGpuH265Enc struct {
	*GstNvEncoder
}

func NewNvAutoGpuH265Enc() (*GstNvAutoGpuH265Enc, error) {
	e, err := gst.NewElement("nvautogpuh265enc")
	if err != nil {
		return nil, err
	}
	return &GstNvAutoGpuH265Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewNvAutoGpuH265EncWithName(name string) (*GstNvAutoGpuH265Enc, error) {
	e, err := gst.NewElementWithName("nvautogpuh265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvAutoGpuH265Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Constant QP value for B frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstNvAutoGpuH265Enc) GetQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvAutoGpuH265Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvAutoGpuH265Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvAutoGpuH265Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvAutoGpuH265Enc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Constant QP value for I frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstNvAutoGpuH265Enc) GetQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (1)
func (e *GstNvAutoGpuH265Enc) SetRateControl(value GstNvEncoderRCMode) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstNvAutoGpuH265Enc) GetRateControl() (GstNvEncoderRCMode, error) {
	var value GstNvEncoderRCMode
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderRCMode")
	}
	return value, nil
}

// CUDA device ID to use
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvAutoGpuH265Enc) SetCudaDeviceId(value uint) error {
	return e.Element.SetProperty("cuda-device-id", value)
}

func (e *GstNvAutoGpuH265Enc) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstNvAutoGpuH265Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstNvAutoGpuH265Enc) GetMaxQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstNvAutoGpuH265Enc) GetMinQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding Preset
// Default: default (0)
func (e *GstNvAutoGpuH265Enc) SetPreset(value GstNvEncoderPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvAutoGpuH265Enc) GetPreset() (GstNvEncoderPreset, error) {
	var value GstNvEncoderPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderPreset")
	}
	return value, nil
}

// Insert sequence headers (SPS/PPS) per IDR, ignored if negotiated stream-format is "hvc1"
// Default: false
func (e *GstNvAutoGpuH265Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

func (e *GstNvAutoGpuH265Enc) GetRepeatSequenceHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat-sequence-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvAutoGpuH265Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvAutoGpuH265Enc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enables Weighted Prediction
// Default: false
func (e *GstNvAutoGpuH265Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvAutoGpuH265Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvAutoGpuH265Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvAutoGpuH265Enc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 5
func (e *GstNvAutoGpuH265Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstNvAutoGpuH265Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Zero latency operation (i.e., num_reorder_frames = 0)
// Default: false
func (e *GstNvAutoGpuH265Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

func (e *GstNvAutoGpuH265Enc) GetZeroReorderDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-reorder-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) to use
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstNvAutoGpuH265Enc) SetAdapterLuid(value int64) error {
	return e.Element.SetProperty("adapter-luid", value)
}

func (e *GstNvAutoGpuH265Enc) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Maximum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstNvAutoGpuH265Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvAutoGpuH265Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvAutoGpuH265Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 30, Min: -1, Max: 2147483647
func (e *GstNvAutoGpuH265Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvAutoGpuH265Enc) GetGopSize() (int, error) {
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

// Constant QP value for P frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstNvAutoGpuH265Enc) GetQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvAutoGpuH265Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvAutoGpuH265Enc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstNvAutoGpuH265Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvAutoGpuH265Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstNvAutoGpuH265Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum Bitrate in kbit/sec (ignored in CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvAutoGpuH265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvAutoGpuH265Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvAutoGpuH265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvAutoGpuH265Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvAutoGpuH265Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvAutoGpuH265Enc) GetBAdapt() (bool, error) {
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

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvAutoGpuH265Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvAutoGpuH265Enc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvAutoGpuH265Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvAutoGpuH265Enc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvAutoGpuH265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvAutoGpuH265Enc) GetAud() (bool, error) {
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

// Bitrate in kbit/sec (0 = automatic)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvAutoGpuH265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvAutoGpuH265Enc) GetBitrate() (uint, error) {
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

// Encode H.265 video streams using NVCODEC API Direct3D11 Mode
type GstNvD3D11H265Enc struct {
	*GstNvEncoder
}

func NewNvD3D11H265Enc() (*GstNvD3D11H265Enc, error) {
	e, err := gst.NewElement("nvd3d11h265enc")
	if err != nil {
		return nil, err
	}
	return &GstNvD3D11H265Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewNvD3D11H265EncWithName(name string) (*GstNvD3D11H265Enc, error) {
	e, err := gst.NewElementWithName("nvd3d11h265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstNvD3D11H265Enc{GstNvEncoder: &GstNvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Insert sequence headers (SPS/PPS) per IDR, ignored if negotiated stream-format is "hvc1"
// Default: false
func (e *GstNvD3D11H265Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

func (e *GstNvD3D11H265Enc) GetRepeatSequenceHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat-sequence-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstNvD3D11H265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstNvD3D11H265Enc) GetAud() (bool, error) {
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

// Encoding Preset
// Default: default (0)
func (e *GstNvD3D11H265Enc) SetPreset(value GstNvEncoderPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvD3D11H265Enc) GetPreset() (GstNvEncoderPreset, error) {
	var value GstNvEncoderPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderPreset")
	}
	return value, nil
}

// Enables Weighted Prediction
// Default: false
func (e *GstNvD3D11H265Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

func (e *GstNvD3D11H265Enc) GetWeightedPred() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("weighted-pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 30, Min: -1, Max: 2147483647
func (e *GstNvD3D11H265Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvD3D11H265Enc) GetGopSize() (int, error) {
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

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvD3D11H265Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvD3D11H265Enc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Constant QP value for P frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstNvD3D11H265Enc) GetQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (1)
func (e *GstNvD3D11H265Enc) SetRateControl(value GstNvEncoderRCMode) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstNvD3D11H265Enc) GetRateControl() (GstNvEncoderRCMode, error) {
	var value GstNvEncoderRCMode
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvEncoderRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvEncoderRCMode")
	}
	return value, nil
}

// Number of frames for frame type lookahead
// Default: 0, Min: 0, Max: 32
func (e *GstNvD3D11H265Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstNvD3D11H265Enc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable adaptive B-frame insert when lookahead is enabled
// Default: false
func (e *GstNvD3D11H265Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

func (e *GstNvD3D11H265Enc) GetBAdapt() (bool, error) {
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

// Number of B-frames between I and P
// Default: 0, Min: 0, Max: 5
func (e *GstNvD3D11H265Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstNvD3D11H265Enc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Bitrate in kbit/sec (0 = automatic)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvD3D11H265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvD3D11H265Enc) GetBitrate() (uint, error) {
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

// Maximum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstNvD3D11H265Enc) GetMaxQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for I frame, (-1 = disabled)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstNvD3D11H265Enc) GetMinQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvD3D11H265Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvD3D11H265Enc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvD3D11H265Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvD3D11H265Enc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvD3D11H265Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvD3D11H265Enc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstNvD3D11H265Enc) GetMinQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for I frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstNvD3D11H265Enc) GetQpI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvD3D11H265Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvD3D11H265Enc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of associated GPU
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstNvD3D11H265Enc) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Maximum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstNvD3D11H265Enc) GetMaxQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for P frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstNvD3D11H265Enc) GetMinQpP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum Bitrate in kbit/sec (ignored in CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvD3D11H265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvD3D11H265Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum QP value for B frame, (-1 = automatic)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstNvD3D11H265Enc) GetMaxQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for B frame (-1 = default)
// Default: -1, Min: -1, Max: 51
func (e *GstNvD3D11H265Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstNvD3D11H265Enc) GetQpB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvD3D11H265Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvD3D11H265Enc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Temporal Adaptive Quantization
// Default: false
func (e *GstNvD3D11H265Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

func (e *GstNvD3D11H265Enc) GetTemporalAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temporal-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in kbits (0 = NVENC default)
// Default: 0, Min: 0, Max: -1
func (e *GstNvD3D11H265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstNvD3D11H265Enc) GetVbvBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vbv-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Zero latency operation (i.e., num_reorder_frames = 0)
// Default: false
func (e *GstNvD3D11H265Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

func (e *GstNvD3D11H265Enc) GetZeroReorderDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-reorder-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// NVDEC video decoder
type Nvvp8dec struct {
	*GstNvDec
}

func NewNvvp8dec() (*Nvvp8dec, error) {
	e, err := gst.NewElement("nvvp8dec")
	if err != nil {
		return nil, err
	}
	return &Nvvp8dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvvp8decWithName(name string) (*Nvvp8dec, error) {
	e, err := gst.NewElementWithName("nvvp8dec", name)
	if err != nil {
		return nil, err
	}
	return &Nvvp8dec{GstNvDec: &GstNvDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// NVIDIA VP9 video decoder
type GstNvVp9SLDec struct {
	*GstVp9Decoder
}

func NewNvVp9SLDec() (*GstNvVp9SLDec, error) {
	e, err := gst.NewElement("nvvp9sldec")
	if err != nil {
		return nil, err
	}
	return &GstNvVp9SLDec{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewNvVp9SLDecWithName(name string) (*GstNvVp9SLDec, error) {
	e, err := gst.NewElementWithName("nvvp9sldec", name)
	if err != nil {
		return nil, err
	}
	return &GstNvVp9SLDec{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Assigned CUDA device id
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvVp9SLDec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstNvEncoder struct {
	*GstVideoEncoder
}

type GstNvEncoderRCMode string

const (
	GstNvEncoderRCModeCqp     GstNvEncoderRCMode = "cqp"       // Constant Quantization
	GstNvEncoderRCModeVbr     GstNvEncoderRCMode = "vbr"       // Variable Bit Rate
	GstNvEncoderRCModeCbr     GstNvEncoderRCMode = "cbr"       // Constant Bit Rate
	GstNvEncoderRCModeCbrLdHq GstNvEncoderRCMode = "cbr-ld-hq" // Low-Delay CBR, High Quality
	GstNvEncoderRCModeCbrHq   GstNvEncoderRCMode = "cbr-hq"    // CBR, High Quality (slower)
	GstNvEncoderRCModeVbrHq   GstNvEncoderRCMode = "vbr-hq"    // VBR, High Quality (slower)
)

type GstNvPreset string

const (
	GstNvPresetDefault      GstNvPreset = "default"        // Default
	GstNvPresetHp           GstNvPreset = "hp"             // High Performance
	GstNvPresetHq           GstNvPreset = "hq"             // High Quality
	GstNvPresetLowLatency   GstNvPreset = "low-latency"    // Low Latency
	GstNvPresetLowLatencyHq GstNvPreset = "low-latency-hq" // Low Latency, High Quality
	GstNvPresetLowLatencyHp GstNvPreset = "low-latency-hp" // Low Latency, High Performance
	GstNvPresetLossless     GstNvPreset = "lossless"       // Lossless
	GstNvPresetLosslessHp   GstNvPreset = "lossless-hp"    // Lossless, High Performance
)

type GstNvRCMode string

const (
	GstNvRCModeDefault  GstNvRCMode = "default"   // Default
	GstNvRCModeConstqp  GstNvRCMode = "constqp"   // Constant Quantization
	GstNvRCModeCbr      GstNvRCMode = "cbr"       // Constant Bit Rate
	GstNvRCModeVbr      GstNvRCMode = "vbr"       // Variable Bit Rate
	GstNvRCModeVbrMinqp GstNvRCMode = "vbr-minqp" // Variable Bit Rate (with minimum quantization parameter, DEPRECATED)
	GstNvRCModeCbrLdHq  GstNvRCMode = "cbr-ld-hq" // Low-Delay CBR, High Quality
	GstNvRCModeCbrHq    GstNvRCMode = "cbr-hq"    // CBR, High Quality (slower)
	GstNvRCModeVbrHq    GstNvRCMode = "vbr-hq"    // VBR, High Quality (slower)
)

type GstCudaBaseConvert struct {
	*GstCudaBaseTransform
}

type GstNvDevice0H264Enc struct {
	*GstNvBaseEnc
}

type GstNvDevice0H265Enc struct {
	*GstNvBaseEnc
}

type GstNvDec struct {
	*GstVideoDecoder
}

// Improves pipelining of decode with display, 0 means no delay (auto = -1)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstNvDec) SetMaxDisplayDelay(value int) error {
	return e.Element.SetProperty("max-display-delay", value)
}

func (e *GstNvDec) GetMaxDisplayDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-display-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Assigned CUDA device id
// Default: 0, Min: 0, Max: 2147483647
func (e *GstNvDec) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstNvEncoderPreset string

const (
	GstNvEncoderPresetDefault      GstNvEncoderPreset = "default"        // Default
	GstNvEncoderPresetHp           GstNvEncoderPreset = "hp"             // High Performance
	GstNvEncoderPresetHq           GstNvEncoderPreset = "hq"             // High Quality
	GstNvEncoderPresetLowLatency   GstNvEncoderPreset = "low-latency"    // Low Latency
	GstNvEncoderPresetLowLatencyHq GstNvEncoderPreset = "low-latency-hq" // Low Latency, High Quality
	GstNvEncoderPresetLowLatencyHp GstNvEncoderPreset = "low-latency-hp" // Low Latency, High Performance
	GstNvEncoderPresetLossless     GstNvEncoderPreset = "lossless"       // Lossless
	GstNvEncoderPresetLosslessHp   GstNvEncoderPreset = "lossless-hp"    // Lossless, High Performance
)

type GstCudaBaseTransform struct {
	*base.GstBaseTransform
}

// Set the GPU device to use for operations (-1 = auto)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstCudaBaseTransform) SetCudaDeviceId(value int) error {
	return e.Element.SetProperty("cuda-device-id", value)
}

func (e *GstCudaBaseTransform) GetCudaDeviceId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstCudaMemoryCopy struct {
	*GstCudaBaseTransform
}

type GstNvBaseEnc struct {
	*GstVideoEncoder
}

// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)
// Default: 0, Min: 0, Max: 15
func (e *GstNvBaseEnc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

func (e *GstNvBaseEnc) GetAqStrength() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("aq-strength")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum Bitrate in kbit/sec (ignored for CBR mode)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvBaseEnc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstNvBaseEnc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encoding Preset
// Default: default (0)
func (e *GstNvBaseEnc) SetPreset(value GstNvPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstNvBaseEnc) GetPreset() (GstNvPreset, error) {
	var value GstNvPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvPreset")
	}
	return value, nil
}

// Constant QP value for P frame, When >= 0, "qp-const-i" and "qp-const-b" should be also >= 0. Overwritten by "qp-const" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpConstP(value int) error {
	return e.Element.SetProperty("qp-const-p", value)
}

func (e *GstNvBaseEnc) GetQpConstP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-const-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for P frame, When >= 0, "qp-max-i" and "qp-max-b" should be also >= 0. Overwritten by "qp-max" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMaxP(value int) error {
	return e.Element.SetProperty("qp-max-p", value)
}

func (e *GstNvBaseEnc) GetQpMaxP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-max-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Rate Control Mode
// Default: default (0)
func (e *GstNvBaseEnc) SetRcMode(value GstNvRCMode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

func (e *GstNvBaseEnc) GetRcMode() (GstNvRCMode, error) {
	var value GstNvRCMode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNvRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNvRCMode")
	}
	return value, nil
}

// Constant QP value for I frame, When >= 0, "qp-const-p" and "qp-const-b" should be also >= 0. Overwritten by "qp-const" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpConstI(value int) error {
	return e.Element.SetProperty("qp-const-i", value)
}

func (e *GstNvBaseEnc) GetQpConstI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-const-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum QP value for B frame, When >= 0, "qp-max-i" and "qp-max-p" should be also >= 0. Overwritten by "qp-max" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMaxB(value int) error {
	return e.Element.SetProperty("qp-max-b", value)
}

func (e *GstNvBaseEnc) GetQpMaxB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-max-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for B frame, When >= 0, "qp-min-i" and "qp-min-p" should be also >= 0. Overwritten by "qp-min" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMinB(value int) error {
	return e.Element.SetProperty("qp-min-b", value)
}

func (e *GstNvBaseEnc) GetQpMinB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-min-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Target Constant Quality level for VBR mode (0 = automatic)
// Default: 0, Min: 0, Max: 51
func (e *GstNvBaseEnc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

func (e *GstNvBaseEnc) GetConstQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("const-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Constant quantizer (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpConst(value int) error {
	return e.Element.SetProperty("qp-const", value)
}

func (e *GstNvBaseEnc) GetQpConst() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-const")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP value for B frame, When >= 0, "qp-const-i" and "qp-const-p" should be also >= 0. Overwritten by "qp-const" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpConstB(value int) error {
	return e.Element.SetProperty("qp-const-b", value)
}

func (e *GstNvBaseEnc) GetQpConstB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-const-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum quantizer (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMin(value int) error {
	return e.Element.SetProperty("qp-min", value)
}

func (e *GstNvBaseEnc) GetQpMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zero latency operation (no reordering delay)
// Default: false
func (e *GstNvBaseEnc) SetZerolatency(value bool) error {
	return e.Element.SetProperty("zerolatency", value)
}

func (e *GstNvBaseEnc) GetZerolatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zerolatency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum QP value for I frame, When >= 0, "qp-max-p" and "qp-max-b" should be also >= 0. Overwritten by "qp-max" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMaxI(value int) error {
	return e.Element.SetProperty("qp-max-i", value)
}

func (e *GstNvBaseEnc) GetQpMaxI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-max-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for I frame, When >= 0, "qp-min-p" and "qp-min-b" should be also >= 0. Overwritten by "qp-min" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMinI(value int) error {
	return e.Element.SetProperty("qp-min-i", value)
}

func (e *GstNvBaseEnc) GetQpMinI() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-min-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Bitrate in kbit/sec (0 = from NVENC preset)
// Default: 0, Min: 0, Max: 2048000
func (e *GstNvBaseEnc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstNvBaseEnc) GetBitrate() (uint, error) {
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

// Get the GPU device to use for operations
// Default: 0, Min: 0, Max: -1
func (e *GstNvBaseEnc) GetCudaDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cuda-device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames between intra frames (-1 = infinite)
// Default: 75, Min: -1, Max: 2147483647
func (e *GstNvBaseEnc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstNvBaseEnc) GetGopSize() (int, error) {
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

// Enable adaptive I-frame insert when lookahead is enabled
// Default: false
func (e *GstNvBaseEnc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

func (e *GstNvBaseEnc) GetIAdapt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Automatic insertion of non-reference P-frames
// Default: false
func (e *GstNvBaseEnc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

func (e *GstNvBaseEnc) GetNonrefP() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("nonref-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum quantizer (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMax(value int) error {
	return e.Element.SetProperty("qp-max", value)
}

func (e *GstNvBaseEnc) GetQpMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum QP value for P frame, When >= 0, "qp-min-i" and "qp-min-b" should be also >= 0. Overwritten by "qp-min" (-1 = from NVENC preset)
// Default: -1, Min: -1, Max: 51
func (e *GstNvBaseEnc) SetQpMinP(value int) error {
	return e.Element.SetProperty("qp-min-p", value)
}

func (e *GstNvBaseEnc) GetQpMinP() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp-min-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Spatial Adaptive Quantization
// Default: false
func (e *GstNvBaseEnc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

func (e *GstNvBaseEnc) GetSpatialAq() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("spatial-aq")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimize GOP-to-GOP rate fluctuations
// Default: false
func (e *GstNvBaseEnc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

func (e *GstNvBaseEnc) GetStrictGop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-gop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
