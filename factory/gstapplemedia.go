// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Read and decode samples from AVFoundation assets using the AVFAssetReader API
type GstAVFAssetSrc struct {
	*gst.Element
}

func NewAVFAssetSrc() (*GstAVFAssetSrc, error) {
	e, err := gst.NewElement("avfassetsrc")
	if err != nil {
		return nil, err
	}
	return &GstAVFAssetSrc{Element: e}, nil
}

func NewAVFAssetSrcWithName(name string) (*GstAVFAssetSrc, error) {
	e, err := gst.NewElementWithName("avfassetsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAVFAssetSrc{Element: e}, nil
}

// URI of the asset to read
// Default: NULL
func (e *GstAVFAssetSrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstAVFAssetSrc) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// H.264 (HW only) encoder
type VtencH264Hw struct {
	*GstVideoEncoder
}

func NewVtencH264Hw() (*VtencH264Hw, error) {
	e, err := gst.NewElement("vtenc_h264_hw")
	if err != nil {
		return nil, err
	}
	return &VtencH264Hw{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewVtencH264HwWithName(name string) (*VtencH264Hw, error) {
	e, err := gst.NewElementWithName("vtenc_h264_hw", name)
	if err != nil {
		return nil, err
	}
	return &VtencH264Hw{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Whether to allow frame reordering or not
// Default: true
func (e *VtencH264Hw) SetAllowFrameReordering(value bool) error {
	return e.Element.SetProperty("allow-frame-reordering", value)
}

func (e *VtencH264Hw) GetAllowFrameReordering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-frame-reordering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Target video bitrate in kbps (0 = auto)
// Default: 0, Min: 0, Max: -1
func (e *VtencH264Hw) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *VtencH264Hw) GetBitrate() (uint, error) {
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

// Maximum number of frames between keyframes (0 = auto)
// Default: 0, Min: 0, Max: 2147483647
func (e *VtencH264Hw) SetMaxKeyframeInterval(value int) error {
	return e.Element.SetProperty("max-keyframe-interval", value)
}

func (e *VtencH264Hw) GetMaxKeyframeInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum number of nanoseconds between keyframes (0 = no limit)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *VtencH264Hw) SetMaxKeyframeIntervalDuration(value uint64) error {
	return e.Element.SetProperty("max-keyframe-interval-duration", value)
}

func (e *VtencH264Hw) GetMaxKeyframeIntervalDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The desired compression quality
// Default: 0.5, Min: 0, Max: 1
func (e *VtencH264Hw) SetQuality(value float64) error {
	return e.Element.SetProperty("quality", value)
}

func (e *VtencH264Hw) GetQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Configure the encoder for realtime output
// Default: false
func (e *VtencH264Hw) SetRealtime(value bool) error {
	return e.Element.SetProperty("realtime", value)
}

func (e *VtencH264Hw) GetRealtime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("realtime")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// AudioToolbox based audio decoder
type GstATDec struct {
	*GstAudioDecoder
}

func NewATDec() (*GstATDec, error) {
	e, err := gst.NewElement("atdec")
	if err != nil {
		return nil, err
	}
	return &GstATDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewATDecWithName(name string) (*GstATDec, error) {
	e, err := gst.NewElementWithName("atdec", name)
	if err != nil {
		return nil, err
	}
	return &GstATDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Reads frames from an iOS AVFoundation device
type GstAVFVideoSrc struct {
	*base.GstPushSrc
}

func NewAVFVideoSrc() (*GstAVFVideoSrc, error) {
	e, err := gst.NewElement("avfvideosrc")
	if err != nil {
		return nil, err
	}
	return &GstAVFVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewAVFVideoSrcWithName(name string) (*GstAVFVideoSrc, error) {
	e, err := gst.NewElementWithName("avfvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAVFVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// The name of the currently opened capture device
// Default: NULL
func (e *GstAVFVideoSrc) GetDeviceName() (string, error) {
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

// The general type of a video capture device
// Default: default (0)
func (e *GstAVFVideoSrc) SetDeviceType(value GstAVFVideoSourceDeviceType) error {
	e.Element.SetArg("device-type", string(value))
	return nil
}

func (e *GstAVFVideoSrc) GetDeviceType() (GstAVFVideoSourceDeviceType, error) {
	var value GstAVFVideoSourceDeviceType
	var ok bool
	v, err := e.Element.GetProperty("device-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAVFVideoSourceDeviceType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAVFVideoSourceDeviceType")
	}
	return value, nil
}

// Enable logging of statistics
// Default: false
func (e *GstAVFVideoSrc) SetDoStats(value bool) error {
	return e.Element.SetProperty("do-stats", value)
}

func (e *GstAVFVideoSrc) GetDoStats() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Last measured framerate, if statistics are enabled
// Default: 0, Min: -1, Max: 2147483647
func (e *GstAVFVideoSrc) GetFps() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The orientation of the video
// Default: default (0)
func (e *GstAVFVideoSrc) SetOrientation(value GstAVFVideoSourceOrientation) error {
	e.Element.SetArg("orientation", string(value))
	return nil
}

func (e *GstAVFVideoSrc) GetOrientation() (GstAVFVideoSourceOrientation, error) {
	var value GstAVFVideoSourceOrientation
	var ok bool
	v, err := e.Element.GetProperty("orientation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAVFVideoSourceOrientation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAVFVideoSourceOrientation")
	}
	return value, nil
}

// Enable cursor capture while capturing screen
// Default: false
func (e *GstAVFVideoSrc) SetCaptureScreenCursor(value bool) error {
	return e.Element.SetProperty("capture-screen-cursor", value)
}

func (e *GstAVFVideoSrc) GetCaptureScreenCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("capture-screen-cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable mouse clicks capture while capturing screen
// Default: false
func (e *GstAVFVideoSrc) SetCaptureScreenMouseClicks(value bool) error {
	return e.Element.SetProperty("capture-screen-mouse-clicks", value)
}

func (e *GstAVFVideoSrc) GetCaptureScreenMouseClicks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("capture-screen-mouse-clicks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The zero-based device index
// Default: -1, Min: -1, Max: 2147483647
func (e *GstAVFVideoSrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

func (e *GstAVFVideoSrc) GetDeviceIndex() (int, error) {
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

// The position of the capture device (front or back-facing)
// Default: default (0)
func (e *GstAVFVideoSrc) SetPosition(value GstAVFVideoSourcePosition) error {
	e.Element.SetArg("position", string(value))
	return nil
}

func (e *GstAVFVideoSrc) GetPosition() (GstAVFVideoSourcePosition, error) {
	var value GstAVFVideoSourcePosition
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAVFVideoSourcePosition)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAVFVideoSourcePosition")
	}
	return value, nil
}

// Enable screen capture functionality
// Default: false
func (e *GstAVFVideoSrc) SetCaptureScreen(value bool) error {
	return e.Element.SetProperty("capture-screen", value)
}

func (e *GstAVFVideoSrc) GetCaptureScreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("capture-screen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// A videosink based on AVSampleBuffers
type GstAVSampleVideoSink struct {
	*GstVideoSink
}

func NewAVSampleVideoSink() (*GstAVSampleVideoSink, error) {
	e, err := gst.NewElement("avsamplebufferlayersink")
	if err != nil {
		return nil, err
	}
	return &GstAVSampleVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewAVSampleVideoSinkWithName(name string) (*GstAVSampleVideoSink, error) {
	e, err := gst.NewElementWithName("avsamplebufferlayersink", name)
	if err != nil {
		return nil, err
	}
	return &GstAVSampleVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstAVSampleVideoSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstAVSampleVideoSink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The CoreAnimation layer that can be placed in the render tree

func (e *GstAVSampleVideoSink) GetLayer() (interface{}, error) {
	return e.Element.GetProperty("layer")
}

// Apple VideoToolbox Decoder
type GstVtdec struct {
	*GstVideoDecoder
}

func NewVtdec() (*GstVtdec, error) {
	e, err := gst.NewElement("vtdec")
	if err != nil {
		return nil, err
	}
	return &GstVtdec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewVtdecWithName(name string) (*GstVtdec, error) {
	e, err := gst.NewElementWithName("vtdec", name)
	if err != nil {
		return nil, err
	}
	return &GstVtdec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Apple VideoToolbox Decoder
type GstVtdecHw struct {
	*GstVtdec
}

func NewVtdecHw() (*GstVtdecHw, error) {
	e, err := gst.NewElement("vtdec_hw")
	if err != nil {
		return nil, err
	}
	return &GstVtdecHw{GstVtdec: &GstVtdec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVtdecHwWithName(name string) (*GstVtdecHw, error) {
	e, err := gst.NewElementWithName("vtdec_hw", name)
	if err != nil {
		return nil, err
	}
	return &GstVtdecHw{GstVtdec: &GstVtdec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// H.264 encoder
type VtencH264 struct {
	*GstVideoEncoder
}

func NewVtencH264() (*VtencH264, error) {
	e, err := gst.NewElement("vtenc_h264")
	if err != nil {
		return nil, err
	}
	return &VtencH264{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewVtencH264WithName(name string) (*VtencH264, error) {
	e, err := gst.NewElementWithName("vtenc_h264", name)
	if err != nil {
		return nil, err
	}
	return &VtencH264{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Configure the encoder for realtime output
// Default: false
func (e *VtencH264) SetRealtime(value bool) error {
	return e.Element.SetProperty("realtime", value)
}

func (e *VtencH264) GetRealtime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("realtime")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to allow frame reordering or not
// Default: true
func (e *VtencH264) SetAllowFrameReordering(value bool) error {
	return e.Element.SetProperty("allow-frame-reordering", value)
}

func (e *VtencH264) GetAllowFrameReordering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-frame-reordering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Target video bitrate in kbps (0 = auto)
// Default: 0, Min: 0, Max: -1
func (e *VtencH264) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *VtencH264) GetBitrate() (uint, error) {
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

// Maximum number of frames between keyframes (0 = auto)
// Default: 0, Min: 0, Max: 2147483647
func (e *VtencH264) SetMaxKeyframeInterval(value int) error {
	return e.Element.SetProperty("max-keyframe-interval", value)
}

func (e *VtencH264) GetMaxKeyframeInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum number of nanoseconds between keyframes (0 = no limit)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *VtencH264) SetMaxKeyframeIntervalDuration(value uint64) error {
	return e.Element.SetProperty("max-keyframe-interval-duration", value)
}

func (e *VtencH264) GetMaxKeyframeIntervalDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The desired compression quality
// Default: 0.5, Min: 0, Max: 1
func (e *VtencH264) SetQuality(value float64) error {
	return e.Element.SetProperty("quality", value)
}

func (e *VtencH264) GetQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Apple ProRes encoder
type VtencProres struct {
	*GstVideoEncoder
}

func NewVtencProres() (*VtencProres, error) {
	e, err := gst.NewElement("vtenc_prores")
	if err != nil {
		return nil, err
	}
	return &VtencProres{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewVtencProresWithName(name string) (*VtencProres, error) {
	e, err := gst.NewElementWithName("vtenc_prores", name)
	if err != nil {
		return nil, err
	}
	return &VtencProres{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Video alpha values (non opaque) need to be preserved
// Default: true
func (e *VtencProres) SetPreserveAlpha(value bool) error {
	return e.Element.SetProperty("preserve-alpha", value)
}

func (e *VtencProres) GetPreserveAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preserve-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The desired compression quality
// Default: 0.5, Min: 0, Max: 1
func (e *VtencProres) SetQuality(value float64) error {
	return e.Element.SetProperty("quality", value)
}

func (e *VtencProres) GetQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Configure the encoder for realtime output
// Default: false
func (e *VtencProres) SetRealtime(value bool) error {
	return e.Element.SetProperty("realtime", value)
}

func (e *VtencProres) GetRealtime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("realtime")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to allow frame reordering or not
// Default: true
func (e *VtencProres) SetAllowFrameReordering(value bool) error {
	return e.Element.SetProperty("allow-frame-reordering", value)
}

func (e *VtencProres) GetAllowFrameReordering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-frame-reordering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Target video bitrate in kbps (0 = auto)
// Default: 0, Min: 0, Max: -1
func (e *VtencProres) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *VtencProres) GetBitrate() (uint, error) {
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

// Maximum number of frames between keyframes (0 = auto)
// Default: 0, Min: 0, Max: 2147483647
func (e *VtencProres) SetMaxKeyframeInterval(value int) error {
	return e.Element.SetProperty("max-keyframe-interval", value)
}

func (e *VtencProres) GetMaxKeyframeInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum number of nanoseconds between keyframes (0 = no limit)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *VtencProres) SetMaxKeyframeIntervalDuration(value uint64) error {
	return e.Element.SetProperty("max-keyframe-interval-duration", value)
}

func (e *VtencProres) GetMaxKeyframeIntervalDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

type GstAVFVideoSourceOrientation string

const (
	GstAVFVideoSourceOrientationPortrait          GstAVFVideoSourceOrientation = "portrait"            // Indicates that video should be oriented vertically, top at the top.
	GstAVFVideoSourceOrientationPortratUpsideDown GstAVFVideoSourceOrientation = "portrat-upside-down" // Indicates that video should be oriented vertically, top at the bottom.
	GstAVFVideoSourceOrientationLandscapeRight    GstAVFVideoSourceOrientation = "landscape-right"     // Indicates that video should be oriented horizontally, top on the left.
	GstAVFVideoSourceOrientationLandscapeLeft     GstAVFVideoSourceOrientation = "landscape-left"      // Indicates that video should be oriented horizontally, top on the right.
	GstAVFVideoSourceOrientationDefault           GstAVFVideoSourceOrientation = "default"             // Default
)

type GstAVFVideoSourcePosition string

const (
	GstAVFVideoSourcePositionFront   GstAVFVideoSourcePosition = "front"   // Front-facing camera
	GstAVFVideoSourcePositionBack    GstAVFVideoSourcePosition = "back"    // Back-facing camera
	GstAVFVideoSourcePositionDefault GstAVFVideoSourcePosition = "default" // Default
)

type GstAVFVideoSourceDeviceType string

const (
	GstAVFVideoSourceDeviceTypeWideAngle GstAVFVideoSourceDeviceType = "wide-angle" // A built-in wide angle camera. These devices are suitable for general purpose use.
	GstAVFVideoSourceDeviceTypeTelephoto GstAVFVideoSourceDeviceType = "telephoto"  // A built-in camera device with a longer focal length than a wide-angle camera.
	GstAVFVideoSourceDeviceTypeDual      GstAVFVideoSourceDeviceType = "dual"       // A dual camera device, combining built-in wide-angle and telephoto cameras that work together as a single capture device.
	GstAVFVideoSourceDeviceTypeDefault   GstAVFVideoSourceDeviceType = "default"    // Default
)
