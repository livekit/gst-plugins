// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// OpenH264 video decoder
type GstOpenh264Dec struct {
	*GstVideoDecoder
}

func NewOpenh264Dec() (*GstOpenh264Dec, error) {
	e, err := gst.NewElement("openh264dec")
	if err != nil {
		return nil, err
	}
	return &GstOpenh264Dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewOpenh264DecWithName(name string) (*GstOpenh264Dec, error) {
	e, err := gst.NewElementWithName("openh264dec", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenh264Dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// OpenH264 video encoder
type GstOpenh264Enc struct {
	*GstVideoEncoder
}

func NewOpenh264Enc() (*GstOpenh264Enc, error) {
	e, err := gst.NewElement("openh264enc")
	if err != nil {
		return nil, err
	}
	return &GstOpenh264Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewOpenh264EncWithName(name string) (*GstOpenh264Enc, error) {
	e, err := gst.NewElementWithName("openh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenh264Enc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// The maximum size of one slice (in bytes).
// Default: 1500000, Min: 0, Max: -1
func (e *GstOpenh264Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

func (e *GstOpenh264Enc) GetMaxSliceSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-slice-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum quantizer
// Default: 51, Min: 0, Max: 51
func (e *GstOpenh264Enc) SetQpMax(value uint) error {
	return e.Element.SetProperty("qp-max", value)
}

func (e *GstOpenh264Enc) GetQpMax() (uint, error) {
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

// Type of video content
// Default: camera (0)
func (e *GstOpenh264Enc) SetUsageType(value EUsageType) error {
	e.Element.SetArg("usage-type", string(value))
	return nil
}

func (e *GstOpenh264Enc) GetUsageType() (EUsageType, error) {
	var value EUsageType
	var ok bool
	v, err := e.Element.GetProperty("usage-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(EUsageType)
	if !ok {
		return value, fmt.Errorf("could not cast value to EUsageType")
	}
	return value, nil
}

// Adaptive quantization
// Default: true
func (e *GstOpenh264Enc) SetAdaptiveQuantization(value bool) error {
	return e.Element.SetProperty("adaptive-quantization", value)
}

func (e *GstOpenh264Enc) GetAdaptiveQuantization() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("adaptive-quantization")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Background detection
// Default: true
func (e *GstOpenh264Enc) SetBackgroundDetection(value bool) error {
	return e.Element.SetProperty("background-detection", value)
}

func (e *GstOpenh264Enc) GetBackgroundDetection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("background-detection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Skip frames to reach target bitrate
// Default: false
func (e *GstOpenh264Enc) SetEnableFrameSkip(value bool) error {
	return e.Element.SetProperty("enable-frame-skip", value)
}

func (e *GstOpenh264Enc) GetEnableFrameSkip() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-frame-skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Denoise control
// Default: false
func (e *GstOpenh264Enc) SetEnableDenoise(value bool) error {
	return e.Element.SetProperty("enable-denoise", value)
}

func (e *GstOpenh264Enc) GetEnableDenoise() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-denoise")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of frames between intra frames
// Default: 90, Min: 0, Max: -1
func (e *GstOpenh264Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstOpenh264Enc) GetGopSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The number of threads.
// Default: 0, Min: 0, Max: -1
func (e *GstOpenh264Enc) SetMultiThread(value uint) error {
	return e.Element.SetProperty("multi-thread", value)
}

func (e *GstOpenh264Enc) GetMultiThread() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("multi-thread")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Scene change detection
// Default: true
func (e *GstOpenh264Enc) SetSceneChangeDetection(value bool) error {
	return e.Element.SetProperty("scene-change-detection", value)
}

func (e *GstOpenh264Enc) GetSceneChangeDetection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scene-change-detection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Bitrate (in bits per second)
// Default: 128000, Min: 0, Max: -1
func (e *GstOpenh264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstOpenh264Enc) GetBitrate() (uint, error) {
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

// Complexity
// Default: medium (1)
func (e *GstOpenh264Enc) SetComplexity(value GstOpenh264encComplexity) error {
	e.Element.SetArg("complexity", string(value))
	return nil
}

func (e *GstOpenh264Enc) GetComplexity() (GstOpenh264encComplexity, error) {
	var value GstOpenh264encComplexity
	var ok bool
	v, err := e.Element.GetProperty("complexity")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenh264encComplexity)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenh264encComplexity")
	}
	return value, nil
}

// Deblocking mode
// Default: on (0)
func (e *GstOpenh264Enc) SetDeblocking(value GstOpenh264encDeblockingModes) error {
	e.Element.SetArg("deblocking", string(value))
	return nil
}

func (e *GstOpenh264Enc) GetDeblocking() (GstOpenh264encDeblockingModes, error) {
	var value GstOpenh264encDeblockingModes
	var ok bool
	v, err := e.Element.GetProperty("deblocking")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenh264encDeblockingModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenh264encDeblockingModes")
	}
	return value, nil
}

// Rate control mode
// Default: quality (0)
func (e *GstOpenh264Enc) SetRateControl(value RCMODES) error {
	return e.Element.SetProperty("rate-control", value)
}

func (e *GstOpenh264Enc) GetRateControl() (RCMODES, error) {
	var value RCMODES
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(RCMODES)
	if !ok {
		return value, fmt.Errorf("could not cast value to RCMODES")
	}
	return value, nil
}

// Slice mode
// Default: n-slices (1)
func (e *GstOpenh264Enc) SetSliceMode(value GstOpenh264EncSliceModes) error {
	e.Element.SetArg("slice-mode", string(value))
	return nil
}

func (e *GstOpenh264Enc) GetSliceMode() (GstOpenh264EncSliceModes, error) {
	var value GstOpenh264EncSliceModes
	var ok bool
	v, err := e.Element.GetProperty("slice-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenh264EncSliceModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenh264EncSliceModes")
	}
	return value, nil
}

// Maximum Bitrate (in bits per second)
// Default: 0, Min: 0, Max: -1
func (e *GstOpenh264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstOpenh264Enc) GetMaxBitrate() (uint, error) {
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

// The number of slices (needs slice-mode=n-slices)
// Default: 1, Min: 0, Max: -1
func (e *GstOpenh264Enc) SetNumSlices(value uint) error {
	return e.Element.SetProperty("num-slices", value)
}

func (e *GstOpenh264Enc) GetNumSlices() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slices")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum quantizer
// Default: 0, Min: 0, Max: 51
func (e *GstOpenh264Enc) SetQpMin(value uint) error {
	return e.Element.SetProperty("qp-min", value)
}

func (e *GstOpenh264Enc) GetQpMin() (uint, error) {
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

type GstOpenh264encComplexity string

const (
	GstOpenh264encComplexityLow    GstOpenh264encComplexity = "low"    // Low complexity / high speed encoding
	GstOpenh264encComplexityMedium GstOpenh264encComplexity = "medium" // Medium complexity / medium speed encoding
	GstOpenh264encComplexityHigh   GstOpenh264encComplexity = "high"   // High complexity / low speed encoding
)

type GstOpenh264encDeblockingModes string

const (
	GstOpenh264encDeblockingModesOn                 GstOpenh264encDeblockingModes = "on"                   // Deblocking on
	GstOpenh264encDeblockingModesOff                GstOpenh264encDeblockingModes = "off"                  // Deblocking off
	GstOpenh264encDeblockingModesNotSliceBoundaries GstOpenh264encDeblockingModes = "not-slice-boundaries" // Deblocking on, except for slice boundaries
)

type RCMODES string

const (
	RCMODESQuality RCMODES = "quality" // Quality mode
	RCMODESBitrate RCMODES = "bitrate" // Bitrate mode
	RCMODESBuffer  RCMODES = "buffer"  // No bitrate control, just using buffer status
	RCMODESOff     RCMODES = "off"     // Rate control off mode
)

type EUsageType string

const (
	EUsageTypeCamera EUsageType = "camera" // video from camera
	EUsageTypeScreen EUsageType = "screen" // screen content
)

type GstOpenh264EncSliceModes string

const (
	GstOpenh264EncSliceModesNSlices GstOpenh264EncSliceModes = "n-slices" // Fixed number of slices
	GstOpenh264EncSliceModesAuto    GstOpenh264EncSliceModes = "auto"     // Number of slices equal to number of threads
)
