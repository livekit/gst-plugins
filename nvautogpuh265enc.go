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

type Nvautogpuh265Enc struct {
	Element *gst.Element
}

func NewNvautogpuh265Enc() (*Nvautogpuh265Enc, error) {
	e, err := gst.NewElement("nvautogpuh265enc")
	if err != nil {
		return nil, err
	}
	return &Nvautogpuh265Enc{Element: e}, nil
}

func NewNvautogpuh265EncWithName(name string) (*Nvautogpuh265Enc, error) {
	e, err := gst.NewElementWithName("nvautogpuh265enc", name)
	if err != nil {
		return nil, err
	}
	return &Nvautogpuh265Enc{Element: e}, nil
}

// ----- Properties -----

// adapter-luid (int64)
//
// DXGI Adapter LUID (Locally Unique Identifier) to use

func (e *Nvautogpuh265Enc) GetAdapterLuid() (int64, error) {
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

func (e *Nvautogpuh265Enc) SetAdapterLuid(value int64) error {
	return e.Element.SetProperty("adapter-luid", value)
}

// aq-strength (uint)
//
// Adaptive Quantization Strength when spatial-aq is enabled from 1 (low) to 15 (aggressive), (0 = autoselect)

func (e *Nvautogpuh265Enc) GetAqStrength() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetAqStrength(value uint) error {
	return e.Element.SetProperty("aq-strength", value)
}

// aud (bool)
//
// Use AU (Access Unit) delimiter

func (e *Nvautogpuh265Enc) GetAud() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

// b-adapt (bool)
//
// Enable adaptive B-frame insert when lookahead is enabled

func (e *Nvautogpuh265Enc) GetBAdapt() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetBAdapt(value bool) error {
	return e.Element.SetProperty("b-adapt", value)
}

// b-frames (uint)
//
// Number of B-frames between I and P

func (e *Nvautogpuh265Enc) GetBFrames() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

// bitrate (uint)
//
// Bitrate in kbit/sec (0 = automatic)

func (e *Nvautogpuh265Enc) GetBitrate() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// const-quality (float64)
//
// Target Constant Quality level for VBR mode (0 = automatic)

func (e *Nvautogpuh265Enc) GetConstQuality() (float64, error) {
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

func (e *Nvautogpuh265Enc) SetConstQuality(value float64) error {
	return e.Element.SetProperty("const-quality", value)
}

// cuda-device-id (uint)
//
// CUDA device ID to use

func (e *Nvautogpuh265Enc) GetCudaDeviceId() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetCudaDeviceId(value uint) error {
	return e.Element.SetProperty("cuda-device-id", value)
}

// gop-size (int)
//
// Number of frames between intra frames (-1 = infinite)

func (e *Nvautogpuh265Enc) GetGopSize() (int, error) {
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

func (e *Nvautogpuh265Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

// i-adapt (bool)
//
// Enable adaptive I-frame insert when lookahead is enabled

func (e *Nvautogpuh265Enc) GetIAdapt() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetIAdapt(value bool) error {
	return e.Element.SetProperty("i-adapt", value)
}

// max-bitrate (uint)
//
// Maximum Bitrate in kbit/sec (ignored in CBR mode)

func (e *Nvautogpuh265Enc) GetMaxBitrate() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-qp-b (int)
//
// Maximum QP value for B frame, (-1 = automatic)

func (e *Nvautogpuh265Enc) GetMaxQpB() (int, error) {
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

func (e *Nvautogpuh265Enc) SetMaxQpB(value int) error {
	return e.Element.SetProperty("max-qp-b", value)
}

// max-qp-i (int)
//
// Maximum QP value for I frame, (-1 = disabled)

func (e *Nvautogpuh265Enc) GetMaxQpI() (int, error) {
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

func (e *Nvautogpuh265Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

// max-qp-p (int)
//
// Maximum QP value for P frame, (-1 = automatic)

func (e *Nvautogpuh265Enc) GetMaxQpP() (int, error) {
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

func (e *Nvautogpuh265Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

// min-qp-b (int)
//
// Minimum QP value for B frame, (-1 = automatic)

func (e *Nvautogpuh265Enc) GetMinQpB() (int, error) {
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

func (e *Nvautogpuh265Enc) SetMinQpB(value int) error {
	return e.Element.SetProperty("min-qp-b", value)
}

// min-qp-i (int)
//
// Minimum QP value for I frame, (-1 = disabled)

func (e *Nvautogpuh265Enc) GetMinQpI() (int, error) {
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

func (e *Nvautogpuh265Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

// min-qp-p (int)
//
// Minimum QP value for P frame, (-1 = automatic)

func (e *Nvautogpuh265Enc) GetMinQpP() (int, error) {
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

func (e *Nvautogpuh265Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

// nonref-p (bool)
//
// Automatic insertion of non-reference P-frames

func (e *Nvautogpuh265Enc) GetNonrefP() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetNonrefP(value bool) error {
	return e.Element.SetProperty("nonref-p", value)
}

// preset (GstNvEncoderPreset)
//
// Encoding Preset

func (e *Nvautogpuh265Enc) GetPreset() (interface{}, error) {
	return e.Element.GetProperty("preset")
}

func (e *Nvautogpuh265Enc) SetPreset(value interface{}) error {
	return e.Element.SetProperty("preset", value)
}

// qp-b (int)
//
// Constant QP value for B frame (-1 = default)

func (e *Nvautogpuh265Enc) GetQpB() (int, error) {
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

func (e *Nvautogpuh265Enc) SetQpB(value int) error {
	return e.Element.SetProperty("qp-b", value)
}

// qp-i (int)
//
// Constant QP value for I frame (-1 = default)

func (e *Nvautogpuh265Enc) GetQpI() (int, error) {
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

func (e *Nvautogpuh265Enc) SetQpI(value int) error {
	return e.Element.SetProperty("qp-i", value)
}

// qp-p (int)
//
// Constant QP value for P frame (-1 = default)

func (e *Nvautogpuh265Enc) GetQpP() (int, error) {
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

func (e *Nvautogpuh265Enc) SetQpP(value int) error {
	return e.Element.SetProperty("qp-p", value)
}

// rate-control (GstNvEncoderRCMode)
//
// Rate Control Method

func (e *Nvautogpuh265Enc) GetRateControl() (interface{}, error) {
	return e.Element.GetProperty("rate-control")
}

func (e *Nvautogpuh265Enc) SetRateControl(value interface{}) error {
	return e.Element.SetProperty("rate-control", value)
}

// rc-lookahead (uint)
//
// Number of frames for frame type lookahead

func (e *Nvautogpuh265Enc) GetRcLookahead() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

// repeat-sequence-header (bool)
//
// Insert sequence headers (SPS/PPS) per IDR, ignored if negotiated stream-format is "hvc1"

func (e *Nvautogpuh265Enc) GetRepeatSequenceHeader() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetRepeatSequenceHeader(value bool) error {
	return e.Element.SetProperty("repeat-sequence-header", value)
}

// spatial-aq (bool)
//
// Spatial Adaptive Quantization

func (e *Nvautogpuh265Enc) GetSpatialAq() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetSpatialAq(value bool) error {
	return e.Element.SetProperty("spatial-aq", value)
}

// strict-gop (bool)
//
// Minimize GOP-to-GOP rate fluctuations

func (e *Nvautogpuh265Enc) GetStrictGop() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetStrictGop(value bool) error {
	return e.Element.SetProperty("strict-gop", value)
}

// temporal-aq (bool)
//
// Temporal Adaptive Quantization

func (e *Nvautogpuh265Enc) GetTemporalAq() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetTemporalAq(value bool) error {
	return e.Element.SetProperty("temporal-aq", value)
}

// vbv-buffer-size (uint)
//
// VBV(HRD) Buffer Size in kbits (0 = NVENC default)

func (e *Nvautogpuh265Enc) GetVbvBufferSize() (uint, error) {
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

func (e *Nvautogpuh265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

// weighted-pred (bool)
//
// Enables Weighted Prediction

func (e *Nvautogpuh265Enc) GetWeightedPred() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetWeightedPred(value bool) error {
	return e.Element.SetProperty("weighted-pred", value)
}

// zero-reorder-delay (bool)
//
// Zero latency operation (i.e., num_reorder_frames = 0)

func (e *Nvautogpuh265Enc) GetZeroReorderDelay() (bool, error) {
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

func (e *Nvautogpuh265Enc) SetZeroReorderDelay(value bool) error {
	return e.Element.SetProperty("zero-reorder-delay", value)
}

