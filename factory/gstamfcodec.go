// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Encode H.265 video streams using AMF API
type GstAmfH265Enc struct {
	*GstAmfEncoder
}

func NewAmfH265Enc() (*GstAmfH265Enc, error) {
	e, err := gst.NewElement("amfh265enc")
	if err != nil {
		return nil, err
	}
	return &GstAmfH265Enc{GstAmfEncoder: &GstAmfEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewAmfH265EncWithName(name string) (*GstAmfH265Enc, error) {
	e, err := gst.NewElementWithName("amfh265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstAmfH265Enc{GstAmfEncoder: &GstAmfEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Number of reference frames
// Default: 1, Min: 1, Max: 16
func (e *GstAmfH265Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstAmfH265Enc) GetRefFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ref-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Target bitrate in kbit/sec (0: USAGE default)
// Default: 0, Min: 0, Max: 2147483
func (e *GstAmfH265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstAmfH265Enc) GetBitrate() (uint, error) {
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

// Maximum bitrate in kbit/sec (0: USAGE default)
// Default: 0, Min: 0, Max: 2147483
func (e *GstAmfH265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstAmfH265Enc) GetMaxBitrate() (uint, error) {
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
// Default: default (-1)
func (e *GstAmfH265Enc) SetRateControl(value GstAmfH265EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstAmfH265Enc) GetRateControl() (GstAmfH265EncRateControl, error) {
	var value GstAmfH265EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH265EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH265EncRateControl")
	}
	return value, nil
}

// Constant QP for I frames
// Default: 26, Min: 0, Max: 51
func (e *GstAmfH265Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstAmfH265Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant QP for P frames
// Default: 26, Min: 0, Max: 51
func (e *GstAmfH265Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstAmfH265Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Target usage
// Default: transcoding (0)
func (e *GstAmfH265Enc) SetUsage(value GstAmfH265EncUsage) error {
	e.Element.SetArg("usage", string(value))
	return nil
}

func (e *GstAmfH265Enc) GetUsage() (GstAmfH265EncUsage, error) {
	var value GstAmfH265EncUsage
	var ok bool
	v, err := e.Element.GetProperty("usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH265EncUsage)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH265EncUsage")
	}
	return value, nil
}

// Preset
// Default: default (-1)
func (e *GstAmfH265Enc) SetPreset(value GstAmfH265EncPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstAmfH265Enc) GetPreset() (GstAmfH265EncPreset, error) {
	var value GstAmfH265EncPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH265EncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH265EncPreset")
	}
	return value, nil
}

// Number of pictures within a GOP
// Default: 30, Min: 0, Max: 2147483647
func (e *GstAmfH265Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstAmfH265Enc) GetGopSize() (uint, error) {
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

// Minimum allowed QP value for I frames (-1: USAGE default)
// Default: -1, Min: -1, Max: 51
func (e *GstAmfH265Enc) SetMinQpI(value int) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstAmfH265Enc) GetMinQpI() (int, error) {
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

// Minimum allowed QP value for P frames (-1: USAGE default)
// Default: -1, Min: -1, Max: 51
func (e *GstAmfH265Enc) SetMinQpP(value int) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstAmfH265Enc) GetMinQpP() (int, error) {
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

// Maximum allowed QP value for P frames (-1: USAGE default)
// Default: -1, Min: -1, Max: 51
func (e *GstAmfH265Enc) SetMaxQpP(value int) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstAmfH265Enc) GetMaxQpP() (int, error) {
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

// DXGI Adapter LUID (Locally Unique Identifier) of associated GPU
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstAmfH265Enc) GetAdapterLuid() (int64, error) {
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

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstAmfH265Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstAmfH265Enc) GetAud() (bool, error) {
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

// Maximum allowed QP value for I frames (-1: USAGE default)
// Default: -1, Min: -1, Max: 51
func (e *GstAmfH265Enc) SetMaxQpI(value int) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstAmfH265Enc) GetMaxQpI() (int, error) {
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

// Encode H.264 video streams using AMF API
type GstAmfH264Enc struct {
	*GstAmfEncoder
}

func NewAmfH264Enc() (*GstAmfH264Enc, error) {
	e, err := gst.NewElement("amfh264enc")
	if err != nil {
		return nil, err
	}
	return &GstAmfH264Enc{GstAmfEncoder: &GstAmfEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewAmfH264EncWithName(name string) (*GstAmfH264Enc, error) {
	e, err := gst.NewElementWithName("amfh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstAmfH264Enc{GstAmfEncoder: &GstAmfEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Maximum bitrate in kbit/sec (0: USAGE default)
// Default: 0, Min: 0, Max: 2147483
func (e *GstAmfH264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstAmfH264Enc) GetMaxBitrate() (uint, error) {
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

// Constant QP for P frames
// Default: 22, Min: 0, Max: 51
func (e *GstAmfH264Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstAmfH264Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
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
// Default: default (-1)
func (e *GstAmfH264Enc) SetRateControl(value GstAmfH264EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstAmfH264Enc) GetRateControl() (GstAmfH264EncRateControl, error) {
	var value GstAmfH264EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH264EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH264EncRateControl")
	}
	return value, nil
}

// Use AU (Access Unit) delimiter
// Default: true
func (e *GstAmfH264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstAmfH264Enc) GetAud() (bool, error) {
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

// Target bitrate in kbit/sec (0: USAGE default)
// Default: 0, Min: 0, Max: 2147483
func (e *GstAmfH264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstAmfH264Enc) GetBitrate() (uint, error) {
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

// Enable CABAC entropy coding
// Default: true
func (e *GstAmfH264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstAmfH264Enc) GetCabac() (bool, error) {
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

// Target usage
// Default: transcoding (0)
func (e *GstAmfH264Enc) SetUsage(value GstAmfH264EncUsage) error {
	e.Element.SetArg("usage", string(value))
	return nil
}

func (e *GstAmfH264Enc) GetUsage() (GstAmfH264EncUsage, error) {
	var value GstAmfH264EncUsage
	var ok bool
	v, err := e.Element.GetProperty("usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH264EncUsage)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH264EncUsage")
	}
	return value, nil
}

// Number of pictures within a GOP (-1: USAGE default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstAmfH264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstAmfH264Enc) GetGopSize() (int, error) {
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

// Maximum allowed QP value (-1: USAGE default)
// Default: -1, Min: -1, Max: 51
func (e *GstAmfH264Enc) SetMaxQp(value int) error {
	return e.Element.SetProperty("max-qp", value)
}

func (e *GstAmfH264Enc) GetMaxQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Constant QP for I frames
// Default: 22, Min: 0, Max: 51
func (e *GstAmfH264Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstAmfH264Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of reference frames
// Default: 1, Min: 1, Max: 16
func (e *GstAmfH264Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstAmfH264Enc) GetRefFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ref-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of associated GPU
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstAmfH264Enc) GetAdapterLuid() (int64, error) {
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

// Minimum allowed QP value (-1: USAGE default)
// Default: -1, Min: -1, Max: 51
func (e *GstAmfH264Enc) SetMinQp(value int) error {
	return e.Element.SetProperty("min-qp", value)
}

func (e *GstAmfH264Enc) GetMinQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Preset
// Default: default (-1)
func (e *GstAmfH264Enc) SetPreset(value GstAmfH264EncPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstAmfH264Enc) GetPreset() (GstAmfH264EncPreset, error) {
	var value GstAmfH264EncPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmfH264EncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmfH264EncPreset")
	}
	return value, nil
}

type GstAmfH264EncRateControl string

const (
	GstAmfH264EncRateControlDefault GstAmfH264EncRateControl = "default" // Default, depends on Usage
	GstAmfH264EncRateControlCqp     GstAmfH264EncRateControl = "cqp"     // Constant QP
	GstAmfH264EncRateControlCbr     GstAmfH264EncRateControl = "cbr"     // Constant Bitrate
	GstAmfH264EncRateControlVbr     GstAmfH264EncRateControl = "vbr"     // Peak Constrained VBR
	GstAmfH264EncRateControlLcvbr   GstAmfH264EncRateControl = "lcvbr"   // Latency Constrained VBR
)

type GstAmfH264EncUsage string

const (
	GstAmfH264EncUsageTranscoding     GstAmfH264EncUsage = "transcoding"       // Transcoding
	GstAmfH264EncUsageUltraLowLatency GstAmfH264EncUsage = "ultra-low-latency" // Ultra Low Latency
	GstAmfH264EncUsageLowLatency      GstAmfH264EncUsage = "low-latency"       // Low Latency
	GstAmfH264EncUsageWebcam          GstAmfH264EncUsage = "webcam"            // Webcam
)

type GstAmfH265EncPreset string

const (
	GstAmfH265EncPresetDefault  GstAmfH265EncPreset = "default"  // Default, depends on USAGE
	GstAmfH265EncPresetQuality  GstAmfH265EncPreset = "quality"  // Quality
	GstAmfH265EncPresetBalanced GstAmfH265EncPreset = "balanced" // Balanced
	GstAmfH265EncPresetSpeed    GstAmfH265EncPreset = "speed"    // Speed
)

type GstAmfH265EncRateControl string

const (
	GstAmfH265EncRateControlDefault GstAmfH265EncRateControl = "default" // Default, depends on Usage
	GstAmfH265EncRateControlCqp     GstAmfH265EncRateControl = "cqp"     // Constant QP
	GstAmfH265EncRateControlLcvbr   GstAmfH265EncRateControl = "lcvbr"   // Latency Constrained VBR
	GstAmfH265EncRateControlVbr     GstAmfH265EncRateControl = "vbr"     // Peak Constrained VBR
	GstAmfH265EncRateControlCbr     GstAmfH265EncRateControl = "cbr"     // Constant Bitrate
)

type GstAmfH265EncUsage string

const (
	GstAmfH265EncUsageTranscoding     GstAmfH265EncUsage = "transcoding"       // Transcoding
	GstAmfH265EncUsageUltraLowLatency GstAmfH265EncUsage = "ultra-low-latency" // Ultra Low Latency
	GstAmfH265EncUsageLowLatency      GstAmfH265EncUsage = "low-latency"       // Low Latency
	GstAmfH265EncUsageWebcam          GstAmfH265EncUsage = "webcam"            // Webcam
)

type GstAmfEncoder struct {
	*GstVideoEncoder
}

type GstAmfH264EncPreset string

const (
	GstAmfH264EncPresetDefault  GstAmfH264EncPreset = "default"  // Default, depends on USAGE
	GstAmfH264EncPresetBalanced GstAmfH264EncPreset = "balanced" // Balanced
	GstAmfH264EncPresetSpeed    GstAmfH264EncPreset = "speed"    // Speed
	GstAmfH264EncPresetQuality  GstAmfH264EncPreset = "quality"  // Quality
)
