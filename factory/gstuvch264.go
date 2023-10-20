// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// UVC H264 Encoding camera source
type GstUvcH264Src struct {
	*GstBaseCameraSrc
}

func NewUvcH264Src() (*GstUvcH264Src, error) {
	e, err := gst.NewElement("uvch264src")
	if err != nil {
		return nil, err
	}
	return &GstUvcH264Src{GstBaseCameraSrc: &GstBaseCameraSrc{Bin: &gst.Bin{Element: e}}}, nil
}

func NewUvcH264SrcWithName(name string) (*GstUvcH264Src, error) {
	e, err := gst.NewElementWithName("uvch264src", name)
	if err != nil {
		return nil, err
	}
	return &GstUvcH264Src{GstBaseCameraSrc: &GstBaseCameraSrc{Bin: &gst.Bin{Element: e}}}, nil
}

// The minimum Quantization step size for P frames (dynamic control)
// Default: 10, Min: -127, Max: 127
func (e *GstUvcH264Src) SetMinPframeQp(value int) error {
	return e.Element.SetProperty("min-pframe-qp", value)
}

func (e *GstUvcH264Src) GetMinPframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-pframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of clock samples to gather for the PTS synchronization (-1 = unlimited)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstUvcH264Src) SetNumClockSamples(value int) error {
	return e.Element.SetProperty("num-clock-samples", value)
}

func (e *GstUvcH264Src) GetNumClockSamples() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-clock-samples")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The peak bitrate in bits/second (dynamic control)
// Default: 3000000, Min: 0, Max: -1
func (e *GstUvcH264Src) SetPeakBitrate(value uint) error {
	return e.Element.SetProperty("peak-bitrate", value)
}

func (e *GstUvcH264Src) GetPeakBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("peak-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Slice units (static control)
// Default: 4, Min: 0, Max: 65535
func (e *GstUvcH264Src) SetSliceUnits(value uint) error {
	return e.Element.SetProperty("slice-units", value)
}

func (e *GstUvcH264Src) GetSliceUnits() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("slice-units")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Time between IDR frames in milliseconds (static control)
// Default: 10000, Min: 0, Max: 65535
func (e *GstUvcH264Src) SetIframePeriod(value uint) error {
	return e.Element.SetProperty("iframe-period", value)
}

func (e *GstUvcH264Src) GetIframePeriod() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("iframe-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Level IDC (dynamic control)
// Default: 40, Min: 0, Max: 255
func (e *GstUvcH264Src) SetLevelIdc(value uint) error {
	return e.Element.SetProperty("level-idc", value)
}

func (e *GstUvcH264Src) GetLevelIdc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("level-idc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Total number of Long-Term Reference frames (dynamic control)
// Default: 0, Min: 0, Max: 255
func (e *GstUvcH264Src) SetLtrBufferSize(value int) error {
	return e.Element.SetProperty("ltr-buffer-size", value)
}

func (e *GstUvcH264Src) GetLtrBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ltr-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The number of macroblocks per second for the maximum processing rate
// Default: 0, Min: 0, Max: -1
func (e *GstUvcH264Src) GetMaxMbps() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-mbps")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Size of the leaky bucket size in milliseconds (static control)
// Default: 1000, Min: 0, Max: 65535
func (e *GstUvcH264Src) SetLeakyBucketSize(value uint) error {
	return e.Element.SetProperty("leaky-bucket-size", value)
}

func (e *GstUvcH264Src) GetLeakyBucketSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("leaky-bucket-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The minimum Quantization step size for I frames (dynamic control)
// Default: 10, Min: -127, Max: 127
func (e *GstUvcH264Src) SetMinIframeQp(value int) error {
	return e.Element.SetProperty("min-iframe-qp", value)
}

func (e *GstUvcH264Src) GetMinIframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-iframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The average bitrate in bits/second (dynamic control)
// Default: 3000000, Min: 0, Max: -1
func (e *GstUvcH264Src) SetAverageBitrate(value uint) error {
	return e.Element.SetProperty("average-bitrate", value)
}

func (e *GstUvcH264Src) GetAverageBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("average-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable SEI picture timing (static control)
// Default: false
func (e *GstUvcH264Src) SetEnableSei(value bool) error {
	return e.Element.SetProperty("enable-sei", value)
}

func (e *GstUvcH264Src) GetEnableSei() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-sei")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Entropy (static control)
// Default: cavlc (0)
func (e *GstUvcH264Src) SetEntropy(value UvcH264Entropy) error {
	e.Element.SetArg("entropy", string(value))
	return nil
}

func (e *GstUvcH264Src) GetEntropy() (UvcH264Entropy, error) {
	var value UvcH264Entropy
	var ok bool
	v, err := e.Element.GetProperty("entropy")
	if err != nil {
		return value, err
	}
	value, ok = v.(UvcH264Entropy)
	if !ok {
		return value, fmt.Errorf("could not cast value to UvcH264Entropy")
	}
	return value, nil
}

// Initial bitrate in bits/second (static control)
// Default: 3000000, Min: 0, Max: -1
func (e *GstUvcH264Src) SetInitialBitrate(value uint) error {
	return e.Element.SetProperty("initial-bitrate", value)
}

func (e *GstUvcH264Src) GetInitialBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("initial-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The name of the jpeg decoder element
// Default: jpegdec
func (e *GstUvcH264Src) SetJpegDecoderName(value string) error {
	return e.Element.SetProperty("jpeg-decoder-name", value)
}

func (e *GstUvcH264Src) GetJpegDecoderName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("jpeg-decoder-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Number of buffers to output before sending EOS (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstUvcH264Src) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstUvcH264Src) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of B frames between the references frames (static control)
// Default: 0, Min: 0, Max: 255
func (e *GstUvcH264Src) SetNumReorderFrames(value uint) error {
	return e.Element.SetProperty("num-reorder-frames", value)
}

func (e *GstUvcH264Src) GetNumReorderFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-reorder-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate control mode (static & dynamic control)
// Default: cbr (1)
func (e *GstUvcH264Src) SetRateControl(value UvcH264RateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstUvcH264Src) GetRateControl() (UvcH264RateControl, error) {
	var value UvcH264RateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(UvcH264RateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to UvcH264RateControl")
	}
	return value, nil
}

// The minimum Quantization step size for B frames (dynamic control)
// Default: 10, Min: -127, Max: 127
func (e *GstUvcH264Src) SetMinBframeQp(value int) error {
	return e.Element.SetProperty("min-bframe-qp", value)
}

func (e *GstUvcH264Src) GetMinBframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-bframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Horizontal flipped image for non H.264 streams (static control)
// Default: false
func (e *GstUvcH264Src) SetPreviewFlipped(value bool) error {
	return e.Element.SetProperty("preview-flipped", value)
}

func (e *GstUvcH264Src) GetPreviewFlipped() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preview-flipped")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Device location
// Default: /dev/video0
func (e *GstUvcH264Src) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstUvcH264Src) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The minimum Quantization step size for B frames (dynamic control)
// Default: 46, Min: -127, Max: 127
func (e *GstUvcH264Src) SetMaxBframeQp(value int) error {
	return e.Element.SetProperty("max-bframe-qp", value)
}

func (e *GstUvcH264Src) GetMaxBframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The usage type (static control)
// Default: realtime (1)
func (e *GstUvcH264Src) SetUsageType(value UvcH264UsageType) error {
	e.Element.SetArg("usage-type", string(value))
	return nil
}

func (e *GstUvcH264Src) GetUsageType() (UvcH264UsageType, error) {
	var value UvcH264UsageType
	var ok bool
	v, err := e.Element.GetProperty("usage-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(UvcH264UsageType)
	if !ok {
		return value, fmt.Errorf("could not cast value to UvcH264UsageType")
	}
	return value, nil
}

// Number of LTR frames the device can control (dynamic control)
// Default: 0, Min: 0, Max: 255
func (e *GstUvcH264Src) SetLtrEncoderControl(value int) error {
	return e.Element.SetProperty("ltr-encoder-control", value)
}

func (e *GstUvcH264Src) GetLtrEncoderControl() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ltr-encoder-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The minimum Quantization step size for P frames (dynamic control)
// Default: 46, Min: -127, Max: 127
func (e *GstUvcH264Src) SetMaxPframeQp(value int) error {
	return e.Element.SetProperty("max-pframe-qp", value)
}

func (e *GstUvcH264Src) GetMaxPframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-pframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The name of the colorspace element
// Default: videoconvert
func (e *GstUvcH264Src) SetColorspaceName(value string) error {
	return e.Element.SetProperty("colorspace-name", value)
}

func (e *GstUvcH264Src) GetColorspaceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("colorspace-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Name of the device

func (e *GstUvcH264Src) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Fixed framerate (static & dynamic control)
// Default: false
func (e *GstUvcH264Src) SetFixedFramerate(value bool) error {
	return e.Element.SetProperty("fixed-framerate", value)
}

func (e *GstUvcH264Src) GetFixedFramerate() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fixed-framerate")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The minimum Quantization step size for I frames (dynamic control)
// Default: 46, Min: -127, Max: 127
func (e *GstUvcH264Src) SetMaxIframeQp(value int) error {
	return e.Element.SetProperty("max-iframe-qp", value)
}

func (e *GstUvcH264Src) GetMaxIframeQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-iframe-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Defines the unit of the slice-units property (static control)
// Default: slice/frame (3)
func (e *GstUvcH264Src) SetSliceMode(value UvcH264SliceMode) error {
	e.Element.SetArg("slice-mode", string(value))
	return nil
}

func (e *GstUvcH264Src) GetSliceMode() (UvcH264SliceMode, error) {
	var value UvcH264SliceMode
	var ok bool
	v, err := e.Element.GetProperty("slice-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(UvcH264SliceMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to UvcH264SliceMode")
	}
	return value, nil
}

// Demux UVC H264 auxiliary streams from MJPG images
type GstUvcH264MjpgDemux struct {
	*gst.Element
}

func NewUvcH264MjpgDemux() (*GstUvcH264MjpgDemux, error) {
	e, err := gst.NewElement("uvch264mjpgdemux")
	if err != nil {
		return nil, err
	}
	return &GstUvcH264MjpgDemux{Element: e}, nil
}

func NewUvcH264MjpgDemuxWithName(name string) (*GstUvcH264MjpgDemux, error) {
	e, err := gst.NewElementWithName("uvch264mjpgdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstUvcH264MjpgDemux{Element: e}, nil
}

// File descriptor of the v4l2 device
// Default: -1, Min: -1, Max: 2147483647
func (e *GstUvcH264MjpgDemux) SetDeviceFd(value int) error {
	return e.Element.SetProperty("device-fd", value)
}

func (e *GstUvcH264MjpgDemux) GetDeviceFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of clock samples to gather for the PTS synchronization (-1 = unlimited)
// Default: 32, Min: 0, Max: 2147483647
func (e *GstUvcH264MjpgDemux) SetNumClockSamples(value int) error {
	return e.Element.SetProperty("num-clock-samples", value)
}

func (e *GstUvcH264MjpgDemux) GetNumClockSamples() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-clock-samples")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type UvcH264UsageType string

const (
	UvcH264UsageTypeRealtime   UvcH264UsageType = "realtime"   // Realtime (video conferencing)
	UvcH264UsageTypeBroadcast  UvcH264UsageType = "broadcast"  // Broadcast
	UvcH264UsageTypeStorage    UvcH264UsageType = "storage"    // Storage
	UvcH264UsageTypeUcconfig0  UvcH264UsageType = "ucconfig0"  // UCConfig 0
	UvcH264UsageTypeUcconfig1  UvcH264UsageType = "ucconfig1"  // UCConfig 1
	UvcH264UsageTypeUcconfig2q UvcH264UsageType = "ucconfig2q" // UCConfig 2Q
	UvcH264UsageTypeUcconfig2s UvcH264UsageType = "ucconfig2s" // UCConfig 2S
	UvcH264UsageTypeUcconfig3  UvcH264UsageType = "ucconfig3"  // UCConfig 3
)

type UvcH264Entropy string

const (
	UvcH264EntropyCavlc UvcH264Entropy = "cavlc" // CAVLC
	UvcH264EntropyCabac UvcH264Entropy = "cabac" // CABAC
)

type UvcH264RateControl string

const (
	UvcH264RateControlCbr UvcH264RateControl = "cbr" // Constant bit rate
	UvcH264RateControlVbr UvcH264RateControl = "vbr" // Variable bit rate
	UvcH264RateControlQp  UvcH264RateControl = "qp"  // Constant QP
)

type UvcH264SliceMode string

const (
	UvcH264SliceModeIgnored    UvcH264SliceMode = "ignored"     // Ignored
	UvcH264SliceModeBitsslice  UvcH264SliceMode = "bits/slice"  // Bits per slice
	UvcH264SliceModeMBsslice   UvcH264SliceMode = "MBs/slice"   // MBs per Slice
	UvcH264SliceModeSliceframe UvcH264SliceMode = "slice/frame" // Slice Per Frame
)
