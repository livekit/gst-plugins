// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Microsoft Media Foundation MP3 Decoder
type GstMFMp3Dec struct {
	*GstMFAudioDecoder
}

func NewMFMp3Dec() (*GstMFMp3Dec, error) {
	e, err := gst.NewElement("mfmp3dec")
	if err != nil {
		return nil, err
	}
	return &GstMFMp3Dec{GstMFAudioDecoder: &GstMFAudioDecoder{GstAudioDecoder: &GstAudioDecoder{Element: e}}}, nil
}

func NewMFMp3DecWithName(name string) (*GstMFMp3Dec, error) {
	e, err := gst.NewElementWithName("mfmp3dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMFMp3Dec{GstMFAudioDecoder: &GstMFAudioDecoder{GstAudioDecoder: &GstAudioDecoder{Element: e}}}, nil
}

// Microsoft Media Foundation MP3 Encoder
type GstMFMp3Enc struct {
	*GstMFAudioEncoder
}

func NewMFMp3Enc() (*GstMFMp3Enc, error) {
	e, err := gst.NewElement("mfmp3enc")
	if err != nil {
		return nil, err
	}
	return &GstMFMp3Enc{GstMFAudioEncoder: &GstMFAudioEncoder{GstAudioEncoder: &GstAudioEncoder{Element: e}}}, nil
}

func NewMFMp3EncWithName(name string) (*GstMFMp3Enc, error) {
	e, err := gst.NewElementWithName("mfmp3enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMFMp3Enc{GstMFAudioEncoder: &GstMFAudioEncoder{GstAudioEncoder: &GstAudioEncoder{Element: e}}}, nil
}

// Bitrate in bit/sec, (0 = auto), valid values are { 0, 8000, 16000, 18000, 20000, 24000, 32000, 40000, 48000, 56000, 64000, 80000, 96000, 112000, 128000, 160000, 192000, 224000, 256000, 320000 }
// Default: 0, Min: 0, Max: 320000
func (e *GstMFMp3Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMFMp3Enc) GetBitrate() (uint, error) {
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

// Capture video stream through Windows Media Foundation
type GstMFVideoSrc struct {
	*base.GstPushSrc
}

func NewMFVideoSrc() (*GstMFVideoSrc, error) {
	e, err := gst.NewElement("mfvideosrc")
	if err != nil {
		return nil, err
	}
	return &GstMFVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewMFVideoSrcWithName(name string) (*GstMFVideoSrc, error) {
	e, err := gst.NewElementWithName("mfvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstMFVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// ICoreDispatcher COM object to use. In order for application to ask permission of capture device, device activation should be running on UI thread via ICoreDispatcher. This element will increase the reference count of given ICoreDispatcher and release it after use. Therefore, caller does not need to consider additional reference count management

func (e *GstMFVideoSrc) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

// The zero-based device index
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMFVideoSrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

func (e *GstMFVideoSrc) GetDeviceIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The human-readable device name
// Default: NULL
func (e *GstMFVideoSrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

func (e *GstMFVideoSrc) GetDeviceName() (string, error) {
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

// The device path
// Default: NULL
func (e *GstMFVideoSrc) SetDevicePath(value string) error {
	return e.Element.SetProperty("device-path", value)
}

func (e *GstMFVideoSrc) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Microsoft Media Foundation VP9 Encoder
type GstMFVP9Enc struct {
	*GstMFVideoEncoder
}

func NewMFVP9Enc() (*GstMFVP9Enc, error) {
	e, err := gst.NewElement("mfvp9enc")
	if err != nil {
		return nil, err
	}
	return &GstMFVP9Enc{GstMFVideoEncoder: &GstMFVideoEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMFVP9EncWithName(name string) (*GstMFVP9Enc, error) {
	e, err := gst.NewElementWithName("mfvp9enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMFVP9Enc{GstMFVideoEncoder: &GstMFVideoEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Enable low latency encoding
// Default: false
func (e *GstMFVP9Enc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstMFVP9Enc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The maximum bitrate applied when rc-mode is "pcvbr" in kbit/sec (0 = MFT default)
// Default: 0, Min: 0, Max: 4194303
func (e *GstMFVP9Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstMFVP9Enc) GetMaxBitrate() (uint, error) {
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

// Rate Control Mode
// Default: cbr (0)
func (e *GstMFVP9Enc) SetRcMode(value GstMFVP9EncRCMode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

func (e *GstMFVP9Enc) GetRcMode() (GstMFVP9EncRCMode, error) {
	var value GstMFVP9EncRCMode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMFVP9EncRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMFVP9EncRCMode")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstMFVP9Enc) GetAdapterLuid() (int64, error) {
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

// Bitrate in kbit/sec
// Default: 2048, Min: 1, Max: 4194303
func (e *GstMFVP9Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMFVP9Enc) GetBitrate() (uint, error) {
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

// Whether device can support Direct3D11 interop
// Default: true
func (e *GstMFVP9Enc) GetD3d11Aware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("d3d11-aware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The number of pictures from one GOP header to the next. Depending on GPU vendor implementation, zero gop-size might produce only one keyframe at the beginning (-1 for automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMFVP9Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstMFVP9Enc) GetGopSize() (int, error) {
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

// Microsoft Media Foundation AAC Decoder
type GstMFAacDec struct {
	*GstMFAudioDecoder
}

func NewMFAacDec() (*GstMFAacDec, error) {
	e, err := gst.NewElement("mfaacdec")
	if err != nil {
		return nil, err
	}
	return &GstMFAacDec{GstMFAudioDecoder: &GstMFAudioDecoder{GstAudioDecoder: &GstAudioDecoder{Element: e}}}, nil
}

func NewMFAacDecWithName(name string) (*GstMFAacDec, error) {
	e, err := gst.NewElementWithName("mfaacdec", name)
	if err != nil {
		return nil, err
	}
	return &GstMFAacDec{GstMFAudioDecoder: &GstMFAudioDecoder{GstAudioDecoder: &GstAudioDecoder{Element: e}}}, nil
}

// Microsoft Media Foundation AAC Encoder
type GstMFAacEnc struct {
	*GstMFAudioEncoder
}

func NewMFAacEnc() (*GstMFAacEnc, error) {
	e, err := gst.NewElement("mfaacenc")
	if err != nil {
		return nil, err
	}
	return &GstMFAacEnc{GstMFAudioEncoder: &GstMFAudioEncoder{GstAudioEncoder: &GstAudioEncoder{Element: e}}}, nil
}

func NewMFAacEncWithName(name string) (*GstMFAacEnc, error) {
	e, err := gst.NewElementWithName("mfaacenc", name)
	if err != nil {
		return nil, err
	}
	return &GstMFAacEnc{GstMFAudioEncoder: &GstMFAudioEncoder{GstAudioEncoder: &GstAudioEncoder{Element: e}}}, nil
}

// Bitrate in bit/sec, (0 = auto), valid values are { 0, 8000, 12000, 16000, 24000, 32000, 48000, 64000, 96000, 128000, 160000, 192000, 256000, 320000, 480000, 512000, 576000, 640000, 720000, 768000, 960000, 1152000 }
// Default: 0, Min: 0, Max: 1152000
func (e *GstMFAacEnc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMFAacEnc) GetBitrate() (uint, error) {
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

// Microsoft Media Foundation H.264 Encoder
type GstMFH264Enc struct {
	*GstMFVideoEncoder
}

func NewMFH264Enc() (*GstMFH264Enc, error) {
	e, err := gst.NewElement("mfh264enc")
	if err != nil {
		return nil, err
	}
	return &GstMFH264Enc{GstMFVideoEncoder: &GstMFVideoEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMFH264EncWithName(name string) (*GstMFH264Enc, error) {
	e, err := gst.NewElementWithName("mfh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMFH264Enc{GstMFVideoEncoder: &GstMFVideoEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstMFH264Enc) GetAdapterLuid() (int64, error) {
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

// Enable CABAC entropy coding
// Default: true
func (e *GstMFH264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstMFH264Enc) GetCabac() (bool, error) {
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

// Whether device can support Direct3D11 interop
// Default: true
func (e *GstMFH264Enc) GetD3d11Aware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("d3d11-aware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The number of pictures from one GOP header to the next. Depending on GPU vendor implementation, zero gop-size might produce only one keyframe at the beginning (-1 for automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMFH264Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstMFH264Enc) GetGopSize() (int, error) {
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

// The maximum bitrate applied when rc-mode is "pcvbr" in kbit/sec
// Default: 0, Min: 0, Max: 4194303
func (e *GstMFH264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstMFH264Enc) GetMaxBitrate() (uint, error) {
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

// QP applied to B frames
// Default: 26, Min: 0, Max: 51
func (e *GstMFH264Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstMFH264Enc) GetQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Quality and speed tradeoff, [0, 33]: Low complexity, [34, 66]: Medium complexity, [67, 100]: High complexity
// Default: 50, Min: 0, Max: 100
func (e *GstMFH264Enc) SetQualityVsSpeed(value uint) error {
	return e.Element.SetProperty("quality-vs-speed", value)
}

func (e *GstMFH264Enc) GetQualityVsSpeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("quality-vs-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate Control Mode
// Default: uvbr (2)
func (e *GstMFH264Enc) SetRcMode(value GstMFH264EncRCMode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

func (e *GstMFH264Enc) GetRcMode() (GstMFH264EncRCMode, error) {
	var value GstMFH264EncRCMode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMFH264EncRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMFH264EncRCMode")
	}
	return value, nil
}

// The number of reference frames
// Default: 2, Min: 0, Max: 16
func (e *GstMFH264Enc) SetRef(value uint) error {
	return e.Element.SetProperty("ref", value)
}

func (e *GstMFH264Enc) GetRef() (uint, error) {
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

// Enable low latency encoding
// Default: false
func (e *GstMFH264Enc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstMFH264Enc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The minimum allowed QP applied to all rc-mode
// Default: 0, Min: 0, Max: 51
func (e *GstMFH264Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

func (e *GstMFH264Enc) GetMinQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// QP applied to I frames
// Default: 26, Min: 0, Max: 51
func (e *GstMFH264Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstMFH264Enc) GetQpI() (uint, error) {
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

// VBV(HRD) Buffer Size in bytes (0 = MFT default)
// Default: 0, Min: 0, Max: -2
func (e *GstMFH264Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstMFH264Enc) GetVbvBufferSize() (uint, error) {
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

// The maximum allowed QP applied to all rc-mode
// Default: 51, Min: 0, Max: 51
func (e *GstMFH264Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

func (e *GstMFH264Enc) GetMaxQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Adaptive Mode
// Default: none (0)
func (e *GstMFH264Enc) SetAdaptiveMode(value GstMFH264EncAdaptiveMode) error {
	e.Element.SetArg("adaptive-mode", string(value))
	return nil
}

func (e *GstMFH264Enc) GetAdaptiveMode() (GstMFH264EncAdaptiveMode, error) {
	var value GstMFH264EncAdaptiveMode
	var ok bool
	v, err := e.Element.GetProperty("adaptive-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMFH264EncAdaptiveMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMFH264EncAdaptiveMode")
	}
	return value, nil
}

// The maximum number of consecutive B frames
// Default: 0, Min: 0, Max: 2
func (e *GstMFH264Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

func (e *GstMFH264Enc) GetBframes() (uint, error) {
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

// Bitrate in kbit/sec
// Default: 2048, Min: 1, Max: 4194303
func (e *GstMFH264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMFH264Enc) GetBitrate() (uint, error) {
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

// QP applied when rc-mode is "qvbr"
// Default: 24, Min: 16, Max: 51
func (e *GstMFH264Enc) SetQp(value uint) error {
	return e.Element.SetProperty("qp", value)
}

func (e *GstMFH264Enc) GetQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// QP applied to P frames
// Default: 26, Min: 0, Max: 51
func (e *GstMFH264Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstMFH264Enc) GetQpP() (uint, error) {
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

// Microsoft Media Foundation H.265 Encoder
type GstMFH265Enc struct {
	*GstMFVideoEncoder
}

func NewMFH265Enc() (*GstMFH265Enc, error) {
	e, err := gst.NewElement("mfh265enc")
	if err != nil {
		return nil, err
	}
	return &GstMFH265Enc{GstMFVideoEncoder: &GstMFVideoEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMFH265EncWithName(name string) (*GstMFH265Enc, error) {
	e, err := gst.NewElementWithName("mfh265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMFH265Enc{GstMFVideoEncoder: &GstMFVideoEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// The maximum bitrate applied when rc-mode is "pcvbr" in kbit/sec (0 = MFT default)
// Default: 0, Min: 0, Max: 4194303
func (e *GstMFH265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstMFH265Enc) GetMaxBitrate() (uint, error) {
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

// QP applied when rc-mode is "qvbr"
// Default: 24, Min: 16, Max: 51
func (e *GstMFH265Enc) SetQp(value uint) error {
	return e.Element.SetProperty("qp", value)
}

func (e *GstMFH265Enc) GetQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// QP applied to B frames
// Default: 26, Min: 0, Max: 51
func (e *GstMFH265Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstMFH265Enc) GetQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The number of reference frames
// Default: 2, Min: 1, Max: 16
func (e *GstMFH265Enc) SetRef(value uint) error {
	return e.Element.SetProperty("ref", value)
}

func (e *GstMFH265Enc) GetRef() (uint, error) {
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

// The maximum number of consecutive B frames
// Default: 0, Min: 0, Max: 2
func (e *GstMFH265Enc) SetBframes(value uint) error {
	return e.Element.SetProperty("bframes", value)
}

func (e *GstMFH265Enc) GetBframes() (uint, error) {
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

// Whether device can support Direct3D11 interop
// Default: true
func (e *GstMFH265Enc) GetD3d11Aware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("d3d11-aware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// QP applied to P frames
// Default: 26, Min: 0, Max: 51
func (e *GstMFH265Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstMFH265Enc) GetQpP() (uint, error) {
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

// Rate Control Mode
// Default: cbr (0)
func (e *GstMFH265Enc) SetRcMode(value GstMFH265EncRCMode) error {
	e.Element.SetArg("rc-mode", string(value))
	return nil
}

func (e *GstMFH265Enc) GetRcMode() (GstMFH265EncRCMode, error) {
	var value GstMFH265EncRCMode
	var ok bool
	v, err := e.Element.GetProperty("rc-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMFH265EncRCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMFH265EncRCMode")
	}
	return value, nil
}

// Quality and speed tradeoff, [0, 33]: Low complexity, [34, 66]: Medium complexity, [67, 100]: High complexity
// Default: 50, Min: 0, Max: 100
func (e *GstMFH265Enc) SetQualityVsSpeed(value uint) error {
	return e.Element.SetProperty("quality-vs-speed", value)
}

func (e *GstMFH265Enc) GetQualityVsSpeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("quality-vs-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// VBV(HRD) Buffer Size in bytes (0 = MFT default)
// Default: 0, Min: 0, Max: -2
func (e *GstMFH265Enc) SetVbvBufferSize(value uint) error {
	return e.Element.SetProperty("vbv-buffer-size", value)
}

func (e *GstMFH265Enc) GetVbvBufferSize() (uint, error) {
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

// Bitrate in kbit/sec
// Default: 2048, Min: 1, Max: 4194303
func (e *GstMFH265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMFH265Enc) GetBitrate() (uint, error) {
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

// The number of pictures from one GOP header to the next. Depending on GPU vendor implementation, zero gop-size might produce only one keyframe at the beginning (-1 for automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMFH265Enc) SetGopSize(value int) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstMFH265Enc) GetGopSize() (int, error) {
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

// The maximum allowed QP applied to all rc-mode
// Default: 51, Min: 0, Max: 51
func (e *GstMFH265Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

func (e *GstMFH265Enc) GetMaxQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// QP applied to I frames
// Default: 26, Min: 0, Max: 51
func (e *GstMFH265Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstMFH265Enc) GetQpI() (uint, error) {
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

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstMFH265Enc) GetAdapterLuid() (int64, error) {
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

// Enable low latency encoding
// Default: false
func (e *GstMFH265Enc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstMFH265Enc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The minimum allowed QP applied to all rc-mode
// Default: 0, Min: 0, Max: 51
func (e *GstMFH265Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

func (e *GstMFH265Enc) GetMinQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstMFAudioDecoder struct {
	*GstAudioDecoder
}

type GstMFAudioEncoder struct {
	*GstAudioEncoder
}

type GstMFH264EncAdaptiveMode string

const (
	GstMFH264EncAdaptiveModeNone      GstMFH264EncAdaptiveMode = "none"      // None
	GstMFH264EncAdaptiveModeFramerate GstMFH264EncAdaptiveMode = "framerate" // Adaptively change the frame rate
)

type GstMFH264EncRCMode string

const (
	GstMFH264EncRCModeCbr   GstMFH264EncRCMode = "cbr"   // Constant bitrate
	GstMFH264EncRCModePcvbr GstMFH264EncRCMode = "pcvbr" // Peak Constrained variable bitrate
	GstMFH264EncRCModeUvbr  GstMFH264EncRCMode = "uvbr"  // Unconstrained variable bitrate
	GstMFH264EncRCModeQvbr  GstMFH264EncRCMode = "qvbr"  // Quality-based variable bitrate
)

type GstMFH265EncRCMode string

const (
	GstMFH265EncRCModeCbr  GstMFH265EncRCMode = "cbr"  // Constant bitrate
	GstMFH265EncRCModeQvbr GstMFH265EncRCMode = "qvbr" // Quality-based variable bitrate
)

type GstMFVP9EncRCMode string

const (
	GstMFVP9EncRCModeCbr  GstMFVP9EncRCMode = "cbr"  // Constant bitrate
	GstMFVP9EncRCModeQvbr GstMFVP9EncRCMode = "qvbr" // Quality-based variable bitrate
)

type GstMFVideoEncoder struct {
	*GstVideoEncoder
}
