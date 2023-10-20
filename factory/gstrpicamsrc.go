// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Raspberry Pi camera module source
type GstRpiCamSrc struct {
	*base.GstPushSrc
}

func NewRpiCamSrc() (*GstRpiCamSrc, error) {
	e, err := gst.NewElement("rpicamsrc")
	if err != nil {
		return nil, err
	}
	return &GstRpiCamSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewRpiCamSrcWithName(name string) (*GstRpiCamSrc, error) {
	e, err := gst.NewElementWithName("rpicamsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRpiCamSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Width of the preview window (in pixels)
// Default: 1024, Min: 0, Max: 2048
func (e *GstRpiCamSrc) SetPreviewW(value int) error {
	return e.Element.SetProperty("preview-w", value)
}

func (e *GstRpiCamSrc) GetPreviewW() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preview-w")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Flags to control annotation of the output video
// Default: (none)
func (e *GstRpiCamSrc) SetAnnotationMode(value GstRpiCamSrcAnnotationMode) error {
	e.Element.SetArg("annotation-mode", fmt.Sprint(value))
	return nil
}

func (e *GstRpiCamSrc) GetAnnotationMode() (GstRpiCamSrcAnnotationMode, error) {
	var value GstRpiCamSrcAnnotationMode
	var ok bool
	v, err := e.Element.GetProperty("annotation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcAnnotationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcAnnotationMode")
	}
	return value, nil
}

// Interval (in frames) between I frames. -1 = automatic, 0 = single-keyframe
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRpiCamSrc) SetKeyframeInterval(value int) error {
	return e.Element.SetProperty("keyframe-interval", value)
}

func (e *GstRpiCamSrc) GetKeyframeInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Height of the preview window (in pixels)
// Default: 768, Min: 0, Max: 2048
func (e *GstRpiCamSrc) SetPreviewH(value int) error {
	return e.Element.SetProperty("preview-h", value)
}

func (e *GstRpiCamSrc) GetPreviewH() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preview-h")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Camera exposure metering mode to use
// Default: average (0)
func (e *GstRpiCamSrc) SetMeteringMode(value GstRpiCamSrcExposureMeteringMode) error {
	e.Element.SetArg("metering-mode", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetMeteringMode() (GstRpiCamSrcExposureMeteringMode, error) {
	var value GstRpiCamSrcExposureMeteringMode
	var ok bool
	v, err := e.Element.GetProperty("metering-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcExposureMeteringMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcExposureMeteringMode")
	}
	return value, nil
}

// Image capture sharpness
// Default: 0, Min: -100, Max: 100
func (e *GstRpiCamSrc) SetSharpness(value int) error {
	return e.Element.SetProperty("sharpness", value)
}

func (e *GstRpiCamSrc) GetSharpness() (int, error) {
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

// Set the annotation text colour, as the integer corresponding to a VUY value eg 0x8080FF = 8421631, -1 for default
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRpiCamSrc) SetAnnotationTextColour(value int) error {
	return e.Element.SetProperty("annotation-text-colour", value)
}

func (e *GstRpiCamSrc) GetAnnotationTextColour() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("annotation-text-colour")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Visual FX to apply to the image
// Default: none (0)
func (e *GstRpiCamSrc) SetImageEffect(value GstRpiCamSrcImageEffect) error {
	e.Element.SetArg("image-effect", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetImageEffect() (GstRpiCamSrcImageEffect, error) {
	var value GstRpiCamSrcImageEffect
	var ok bool
	v, err := e.Element.GetProperty("image-effect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcImageEffect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcImageEffect")
	}
	return value, nil
}

// Set to TRUE to insert SPS/PPS before each IDR packet
// Default: false
func (e *GstRpiCamSrc) SetInlineHeaders(value bool) error {
	return e.Element.SetProperty("inline-headers", value)
}

func (e *GstRpiCamSrc) GetInlineHeaders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("inline-headers")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Rotate captured image (0, 90, 180, 270 degrees)
// Default: 0, Min: 0, Max: 270
func (e *GstRpiCamSrc) SetRotation(value int) error {
	return e.Element.SetProperty("rotation", value)
}

func (e *GstRpiCamSrc) GetRotation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Manually set the camera sensor mode
// Default: automatic (0)
func (e *GstRpiCamSrc) SetSensorMode(value GstRpiCamSrcSensorMode) error {
	e.Element.SetArg("sensor-mode", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetSensorMode() (GstRpiCamSrcSensorMode, error) {
	var value GstRpiCamSrcSensorMode
	var ok bool
	v, err := e.Element.GetProperty("sensor-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcSensorMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcSensorMode")
	}
	return value, nil
}

// Set the annotation text background colour, as the integer corresponding to a VUY value eg 0x8080FF = 8421631, -1 for default
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRpiCamSrc) SetAnnotationTextBgColour(value int) error {
	return e.Element.SetProperty("annotation-text-bg-colour", value)
}

func (e *GstRpiCamSrc) GetAnnotationTextBgColour() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("annotation-text-bg-colour")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Display preview window overlay
// Default: true
func (e *GstRpiCamSrc) SetPreview(value bool) error {
	return e.Element.SetProperty("preview", value)
}

func (e *GstRpiCamSrc) GetPreview() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preview")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Display encoder output in the preview
// Default: true
func (e *GstRpiCamSrc) SetPreviewEncoded(value bool) error {
	return e.Element.SetProperty("preview-encoded", value)
}

func (e *GstRpiCamSrc) GetPreviewEncoded() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("preview-encoded")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Bitrate for encoding. 0 for VBR using quantisation-parameter
// Default: 17000000, Min: 0, Max: 25000000
func (e *GstRpiCamSrc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstRpiCamSrc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Camera exposure mode to use
// Default: auto (1)
func (e *GstRpiCamSrc) SetExposureMode(value GstRpiCamSrcExposureMode) error {
	e.Element.SetArg("exposure-mode", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetExposureMode() (GstRpiCamSrcExposureMode, error) {
	var value GstRpiCamSrcExposureMode
	var ok bool
	v, err := e.Element.GetProperty("exposure-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcExposureMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcExposureMode")
	}
	return value, nil
}

// Opacity to use for the preview window
// Default: 255, Min: 0, Max: 255
func (e *GstRpiCamSrc) SetPreviewOpacity(value int) error {
	return e.Element.SetProperty("preview-opacity", value)
}

func (e *GstRpiCamSrc) GetPreviewOpacity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preview-opacity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set a Quantisation Parameter approx 10-40 with bitrate=0 for VBR encoding. 0 = off
// Default: 0, Min: 0, Max: 2147483647
func (e *GstRpiCamSrc) SetQuantisationParameter(value int) error {
	return e.Element.SetProperty("quantisation-parameter", value)
}

func (e *GstRpiCamSrc) GetQuantisationParameter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quantisation-parameter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Image capture saturation
// Default: 0, Min: -100, Max: 100
func (e *GstRpiCamSrc) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstRpiCamSrc) GetSaturation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Enable or disable video stabilisation
// Default: false
func (e *GstRpiCamSrc) SetVideoStabilisation(value bool) error {
	return e.Element.SetProperty("video-stabilisation", value)
}

func (e *GstRpiCamSrc) GetVideoStabilisation() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("video-stabilisation")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set the size of annotation text (in pixels) (0 = Auto)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstRpiCamSrc) SetAnnotationTextSize(value int) error {
	return e.Element.SetProperty("annotation-text-size", value)
}

func (e *GstRpiCamSrc) GetAnnotationTextSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("annotation-text-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Manual AWB Gain for blue channel when awb-mode=off
// Default: 0, Min: 0, Max: 8
func (e *GstRpiCamSrc) SetAwbGainBlue(value float32) error {
	return e.Element.SetProperty("awb-gain-blue", value)
}

func (e *GstRpiCamSrc) GetAwbGainBlue() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("awb-gain-blue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Dynamic Range Control level
// Default: off (0)
func (e *GstRpiCamSrc) SetDrc(value GstRpiCamSrcDRCLevel) error {
	e.Element.SetArg("drc", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetDrc() (GstRpiCamSrcDRCLevel, error) {
	var value GstRpiCamSrcDRCLevel
	var ok bool
	v, err := e.Element.GetProperty("drc")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcDRCLevel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcDRCLevel")
	}
	return value, nil
}

// Image capture brightness
// Default: 50, Min: 0, Max: 100
func (e *GstRpiCamSrc) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstRpiCamSrc) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Which camera to use on a multi-camera system - 0 or 1
// Default: 0, Min: 0, Max: 1
func (e *GstRpiCamSrc) SetCameraNumber(value int) error {
	return e.Element.SetProperty("camera-number", value)
}

func (e *GstRpiCamSrc) GetCameraNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("camera-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Image capture contrast
// Default: 0, Min: -100, Max: 100
func (e *GstRpiCamSrc) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstRpiCamSrc) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Display preview window full screen
// Default: true
func (e *GstRpiCamSrc) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

func (e *GstRpiCamSrc) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Flip capture horizontally
// Default: false
func (e *GstRpiCamSrc) SetHflip(value bool) error {
	return e.Element.SetProperty("hflip", value)
}

func (e *GstRpiCamSrc) GetHflip() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hflip")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Text string to annotate onto video when annotation-mode flags include 'custom-text'

func (e *GstRpiCamSrc) SetAnnotationText(value string) error {
	return e.Element.SetProperty("annotation-text", value)
}

func (e *GstRpiCamSrc) GetAnnotationText() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("annotation-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Manual AWB Gain for red channel when awb-mode=off
// Default: 0, Min: 0, Max: 8
func (e *GstRpiCamSrc) SetAwbGainRed(value float32) error {
	return e.Element.SetProperty("awb-gain-red", value)
}

func (e *GstRpiCamSrc) GetAwbGainRed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("awb-gain-red")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// White Balance mode
// Default: auto (1)
func (e *GstRpiCamSrc) SetAwbMode(value GstRpiCamSrcAWBMode) error {
	e.Element.SetArg("awb-mode", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetAwbMode() (GstRpiCamSrcAWBMode, error) {
	var value GstRpiCamSrcAWBMode
	var ok bool
	v, err := e.Element.GetProperty("awb-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcAWBMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcAWBMode")
	}
	return value, nil
}

// ISO value to use (0 = Auto)
// Default: 0, Min: 0, Max: 3200
func (e *GstRpiCamSrc) SetIso(value int) error {
	return e.Element.SetProperty("iso", value)
}

func (e *GstRpiCamSrc) GetIso() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("iso")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Start X coordinate of the preview window (in pixels)
// Default: 0, Min: 0, Max: 2048
func (e *GstRpiCamSrc) SetPreviewX(value int) error {
	return e.Element.SetProperty("preview-x", value)
}

func (e *GstRpiCamSrc) GetPreviewX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preview-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Normalised region-of-interest X coord
// Default: 0, Min: 0, Max: 1
func (e *GstRpiCamSrc) SetRoiX(value float32) error {
	return e.Element.SetProperty("roi-x", value)
}

func (e *GstRpiCamSrc) GetRoiX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("roi-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Normalised region-of-interest Y coord
// Default: 0, Min: 0, Max: 1
func (e *GstRpiCamSrc) SetRoiY(value float32) error {
	return e.Element.SetProperty("roi-y", value)
}

func (e *GstRpiCamSrc) GetRoiY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("roi-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Flip capture vertically
// Default: false
func (e *GstRpiCamSrc) SetVflip(value bool) error {
	return e.Element.SetProperty("vflip", value)
}

func (e *GstRpiCamSrc) GetVflip() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vflip")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Exposure Value compensation
// Default: 0, Min: -10, Max: 10
func (e *GstRpiCamSrc) SetExposureCompensation(value int) error {
	return e.Element.SetProperty("exposure-compensation", value)
}

func (e *GstRpiCamSrc) GetExposureCompensation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("exposure-compensation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Type of Intra Refresh to use, -1 to disable intra refresh
// Default: none (-1)
func (e *GstRpiCamSrc) SetIntraRefreshType(value GstRpiCamSrcIntraRefreshType) error {
	e.Element.SetArg("intra-refresh-type", string(value))
	return nil
}

func (e *GstRpiCamSrc) GetIntraRefreshType() (GstRpiCamSrcIntraRefreshType, error) {
	var value GstRpiCamSrcIntraRefreshType
	var ok bool
	v, err := e.Element.GetProperty("intra-refresh-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRpiCamSrcIntraRefreshType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRpiCamSrcIntraRefreshType")
	}
	return value, nil
}

// Normalised region-of-interest H coord
// Default: 1, Min: 0, Max: 1
func (e *GstRpiCamSrc) SetRoiH(value float32) error {
	return e.Element.SetProperty("roi-h", value)
}

func (e *GstRpiCamSrc) GetRoiH() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("roi-h")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Use the camera STC for timestamping buffers
// Default: true
func (e *GstRpiCamSrc) SetUseStc(value bool) error {
	return e.Element.SetProperty("use-stc", value)
}

func (e *GstRpiCamSrc) GetUseStc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-stc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Start Y coordinate of the preview window (in pixels)
// Default: 0, Min: 0, Max: 2048
func (e *GstRpiCamSrc) SetPreviewY(value int) error {
	return e.Element.SetProperty("preview-y", value)
}

func (e *GstRpiCamSrc) GetPreviewY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("preview-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Normalised region-of-interest W coord
// Default: 1, Min: 0, Max: 1
func (e *GstRpiCamSrc) SetRoiW(value float32) error {
	return e.Element.SetProperty("roi-w", value)
}

func (e *GstRpiCamSrc) GetRoiW() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("roi-w")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Set a fixed shutter speed, in microseconds. (0 = Auto)
// Default: 0, Min: 0, Max: 6000000
func (e *GstRpiCamSrc) SetShutterSpeed(value int) error {
	return e.Element.SetProperty("shutter-speed", value)
}

func (e *GstRpiCamSrc) GetShutterSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("shutter-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstRpiCamSrcExposureMode string

const (
	GstRpiCamSrcExposureModeOff          GstRpiCamSrcExposureMode = "off"          // GST_RPI_CAM_SRC_EXPOSURE_MODE_OFF
	GstRpiCamSrcExposureModeAuto         GstRpiCamSrcExposureMode = "auto"         // GST_RPI_CAM_SRC_EXPOSURE_MODE_AUTO
	GstRpiCamSrcExposureModeNight        GstRpiCamSrcExposureMode = "night"        // GST_RPI_CAM_SRC_EXPOSURE_MODE_NIGHT
	GstRpiCamSrcExposureModeNightpreview GstRpiCamSrcExposureMode = "nightpreview" // GST_RPI_CAM_SRC_EXPOSURE_MODE_NIGHTPREVIEW
	GstRpiCamSrcExposureModeBacklight    GstRpiCamSrcExposureMode = "backlight"    // GST_RPI_CAM_SRC_EXPOSURE_MODE_BACKLIGHT
	GstRpiCamSrcExposureModeSpotlight    GstRpiCamSrcExposureMode = "spotlight"    // GST_RPI_CAM_SRC_EXPOSURE_MODE_SPOTLIGHT
	GstRpiCamSrcExposureModeSports       GstRpiCamSrcExposureMode = "sports"       // GST_RPI_CAM_SRC_EXPOSURE_MODE_SPORTS
	GstRpiCamSrcExposureModeSnow         GstRpiCamSrcExposureMode = "snow"         // GST_RPI_CAM_SRC_EXPOSURE_MODE_SNOW
	GstRpiCamSrcExposureModeBeach        GstRpiCamSrcExposureMode = "beach"        // GST_RPI_CAM_SRC_EXPOSURE_MODE_BEACH
	GstRpiCamSrcExposureModeVerylong     GstRpiCamSrcExposureMode = "verylong"     // GST_RPI_CAM_SRC_EXPOSURE_MODE_VERYLONG
	GstRpiCamSrcExposureModeFixedfps     GstRpiCamSrcExposureMode = "fixedfps"     // GST_RPI_CAM_SRC_EXPOSURE_MODE_FIXEDFPS
	GstRpiCamSrcExposureModeAntishake    GstRpiCamSrcExposureMode = "antishake"    // GST_RPI_CAM_SRC_EXPOSURE_MODE_ANTISHAKE
	GstRpiCamSrcExposureModeFireworks    GstRpiCamSrcExposureMode = "fireworks"    // GST_RPI_CAM_SRC_EXPOSURE_MODE_FIREWORKS
)

type GstRpiCamSrcImageEffect string

const (
	GstRpiCamSrcImageEffectNone          GstRpiCamSrcImageEffect = "none"          // GST_RPI_CAM_SRC_IMAGEFX_NONE
	GstRpiCamSrcImageEffectNegative      GstRpiCamSrcImageEffect = "negative"      // GST_RPI_CAM_SRC_IMAGEFX_NEGATIVE
	GstRpiCamSrcImageEffectSolarize      GstRpiCamSrcImageEffect = "solarize"      // GST_RPI_CAM_SRC_IMAGEFX_SOLARIZE
	GstRpiCamSrcImageEffectPosterize     GstRpiCamSrcImageEffect = "posterize"     // GST_RPI_CAM_SRC_IMAGEFX_POSTERIZE
	GstRpiCamSrcImageEffectWhiteboard    GstRpiCamSrcImageEffect = "whiteboard"    // GST_RPI_CAM_SRC_IMAGEFX_WHITEBOARD
	GstRpiCamSrcImageEffectBlackboard    GstRpiCamSrcImageEffect = "blackboard"    // GST_RPI_CAM_SRC_IMAGEFX_BLACKBOARD
	GstRpiCamSrcImageEffectSketch        GstRpiCamSrcImageEffect = "sketch"        // GST_RPI_CAM_SRC_IMAGEFX_SKETCH
	GstRpiCamSrcImageEffectDenoise       GstRpiCamSrcImageEffect = "denoise"       // GST_RPI_CAM_SRC_IMAGEFX_DENOISE
	GstRpiCamSrcImageEffectEmboss        GstRpiCamSrcImageEffect = "emboss"        // GST_RPI_CAM_SRC_IMAGEFX_EMBOSS
	GstRpiCamSrcImageEffectOilpaint      GstRpiCamSrcImageEffect = "oilpaint"      // GST_RPI_CAM_SRC_IMAGEFX_OILPAINT
	GstRpiCamSrcImageEffectHatch         GstRpiCamSrcImageEffect = "hatch"         // GST_RPI_CAM_SRC_IMAGEFX_HATCH
	GstRpiCamSrcImageEffectGpen          GstRpiCamSrcImageEffect = "gpen"          // GST_RPI_CAM_SRC_IMAGEFX_GPEN
	GstRpiCamSrcImageEffectPastel        GstRpiCamSrcImageEffect = "pastel"        // GST_RPI_CAM_SRC_IMAGEFX_PASTEL
	GstRpiCamSrcImageEffectWatercolour   GstRpiCamSrcImageEffect = "watercolour"   // GST_RPI_CAM_SRC_IMAGEFX_WATERCOLOUR
	GstRpiCamSrcImageEffectFilm          GstRpiCamSrcImageEffect = "film"          // GST_RPI_CAM_SRC_IMAGEFX_FILM
	GstRpiCamSrcImageEffectBlur          GstRpiCamSrcImageEffect = "blur"          // GST_RPI_CAM_SRC_IMAGEFX_BLUR
	GstRpiCamSrcImageEffectSaturation    GstRpiCamSrcImageEffect = "saturation"    // GST_RPI_CAM_SRC_IMAGEFX_SATURATION
	GstRpiCamSrcImageEffectColourswap    GstRpiCamSrcImageEffect = "colourswap"    // GST_RPI_CAM_SRC_IMAGEFX_COLOURSWAP
	GstRpiCamSrcImageEffectWashedout     GstRpiCamSrcImageEffect = "washedout"     // GST_RPI_CAM_SRC_IMAGEFX_WASHEDOUT
	GstRpiCamSrcImageEffectPosterise     GstRpiCamSrcImageEffect = "posterise"     // GST_RPI_CAM_SRC_IMAGEFX_POSTERISE
	GstRpiCamSrcImageEffectColourpoint   GstRpiCamSrcImageEffect = "colourpoint"   // GST_RPI_CAM_SRC_IMAGEFX_COLOURPOINT
	GstRpiCamSrcImageEffectColourbalance GstRpiCamSrcImageEffect = "colourbalance" // GST_RPI_CAM_SRC_IMAGEFX_COLOURBALANCE
	GstRpiCamSrcImageEffectCartoon       GstRpiCamSrcImageEffect = "cartoon"       // GST_RPI_CAM_SRC_IMAGEFX_CARTOON
)

type GstRpiCamSrcIntraRefreshType string

const (
	GstRpiCamSrcIntraRefreshTypeNone       GstRpiCamSrcIntraRefreshType = "none"        // GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_NONE
	GstRpiCamSrcIntraRefreshTypeCyclic     GstRpiCamSrcIntraRefreshType = "cyclic"      // GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_CYCLIC
	GstRpiCamSrcIntraRefreshTypeAdaptive   GstRpiCamSrcIntraRefreshType = "adaptive"    // GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_ADAPTIVE
	GstRpiCamSrcIntraRefreshTypeBoth       GstRpiCamSrcIntraRefreshType = "both"        // GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_BOTH
	GstRpiCamSrcIntraRefreshTypeCyclicRows GstRpiCamSrcIntraRefreshType = "cyclic-rows" // GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_CYCLIC_ROWS
)

type GstRpiCamSrcSensorMode string

const (
	GstRpiCamSrcSensorModeAutomatic     GstRpiCamSrcSensorMode = "automatic"      // Automatic
	GstRpiCamSrcSensorMode1920x1080     GstRpiCamSrcSensorMode = "1920x1080"      // 1920x1080 16:9 1-30fps
	GstRpiCamSrcSensorMode2592x1944Fast GstRpiCamSrcSensorMode = "2592x1944-fast" // 2592x1944 4:3 1-15fps / 3240x2464 15fps w/ v.2 board
	GstRpiCamSrcSensorMode2592x1944Slow GstRpiCamSrcSensorMode = "2592x1944-slow" // 2592x1944 4:3 0.1666-1fps / 3240x2464 15fps w/ v.2 board
	GstRpiCamSrcSensorMode1296x972      GstRpiCamSrcSensorMode = "1296x972"       // 1296x972 4:3 1-42fps
	GstRpiCamSrcSensorMode1296x730      GstRpiCamSrcSensorMode = "1296x730"       // 1296x730 16:9 1-49fps
	GstRpiCamSrcSensorMode640x480Slow   GstRpiCamSrcSensorMode = "640x480-slow"   // 640x480 4:3 42.1-60fps
	GstRpiCamSrcSensorMode640x480Fast   GstRpiCamSrcSensorMode = "640x480-fast"   // 640x480 4:3 60.1-90fps
)

type GstRpiCamSrcAWBMode string

const (
	GstRpiCamSrcAWBModeOff          GstRpiCamSrcAWBMode = "off"          // GST_RPI_CAM_SRC_AWB_MODE_OFF
	GstRpiCamSrcAWBModeAuto         GstRpiCamSrcAWBMode = "auto"         // GST_RPI_CAM_SRC_AWB_MODE_AUTO
	GstRpiCamSrcAWBModeSunlight     GstRpiCamSrcAWBMode = "sunlight"     // GST_RPI_CAM_SRC_AWB_MODE_SUNLIGHT
	GstRpiCamSrcAWBModeCloudy       GstRpiCamSrcAWBMode = "cloudy"       // GST_RPI_CAM_SRC_AWB_MODE_CLOUDY
	GstRpiCamSrcAWBModeShade        GstRpiCamSrcAWBMode = "shade"        // GST_RPI_CAM_SRC_AWB_MODE_SHADE
	GstRpiCamSrcAWBModeTungsten     GstRpiCamSrcAWBMode = "tungsten"     // GST_RPI_CAM_SRC_AWB_MODE_TUNGSTEN
	GstRpiCamSrcAWBModeFluorescent  GstRpiCamSrcAWBMode = "fluorescent"  // GST_RPI_CAM_SRC_AWB_MODE_FLUORESCENT
	GstRpiCamSrcAWBModeIncandescent GstRpiCamSrcAWBMode = "incandescent" // GST_RPI_CAM_SRC_AWB_MODE_INCANDESCENT
	GstRpiCamSrcAWBModeFlash        GstRpiCamSrcAWBMode = "flash"        // GST_RPI_CAM_SRC_AWB_MODE_FLASH
	GstRpiCamSrcAWBModeHorizon      GstRpiCamSrcAWBMode = "horizon"      // GST_RPI_CAM_SRC_AWB_MODE_HORIZON
)

type GstRpiCamSrcAnnotationMode int

const (
	GstRpiCamSrcAnnotationModeCustomText      GstRpiCamSrcAnnotationMode = 0x00000001 // GST_RPI_CAM_SRC_ANNOTATION_MODE_CUSTOM_TEXT
	GstRpiCamSrcAnnotationModeText            GstRpiCamSrcAnnotationMode = 0x00000002 // GST_RPI_CAM_SRC_ANNOTATION_MODE_TEXT
	GstRpiCamSrcAnnotationModeDate            GstRpiCamSrcAnnotationMode = 0x00000004 // GST_RPI_CAM_SRC_ANNOTATION_MODE_DATE
	GstRpiCamSrcAnnotationModeTime            GstRpiCamSrcAnnotationMode = 0x00000008 // GST_RPI_CAM_SRC_ANNOTATION_MODE_TIME
	GstRpiCamSrcAnnotationModeShutterSettings GstRpiCamSrcAnnotationMode = 0x00000010 // GST_RPI_CAM_SRC_ANNOTATION_MODE_SHUTTER_SETTINGS
	GstRpiCamSrcAnnotationModeCafSettings     GstRpiCamSrcAnnotationMode = 0x00000020 // GST_RPI_CAM_SRC_ANNOTATION_MODE_CAF_SETTINGS
	GstRpiCamSrcAnnotationModeGainSettings    GstRpiCamSrcAnnotationMode = 0x00000040 // GST_RPI_CAM_SRC_ANNOTATION_MODE_GAIN_SETTINGS
	GstRpiCamSrcAnnotationModeLensSettings    GstRpiCamSrcAnnotationMode = 0x00000080 // GST_RPI_CAM_SRC_ANNOTATION_MODE_LENS_SETTINGS
	GstRpiCamSrcAnnotationModeMotionSettings  GstRpiCamSrcAnnotationMode = 0x00000100 // GST_RPI_CAM_SRC_ANNOTATION_MODE_MOTION_SETTINGS
	GstRpiCamSrcAnnotationModeFrameNumber     GstRpiCamSrcAnnotationMode = 0x00000200 // GST_RPI_CAM_SRC_ANNOTATION_MODE_FRAME_NUMBER
	GstRpiCamSrcAnnotationModeBlackBackground GstRpiCamSrcAnnotationMode = 0x00000400 // GST_RPI_CAM_SRC_ANNOTATION_MODE_BLACK_BACKGROUND
)

type GstRpiCamSrcDRCLevel string

const (
	GstRpiCamSrcDRCLevelOff    GstRpiCamSrcDRCLevel = "off"    // GST_RPI_CAM_SRC_DRC_LEVEL_OFF
	GstRpiCamSrcDRCLevelLow    GstRpiCamSrcDRCLevel = "low"    // GST_RPI_CAM_SRC_DRC_LEVEL_LOW
	GstRpiCamSrcDRCLevelMedium GstRpiCamSrcDRCLevel = "medium" // GST_RPI_CAM_SRC_DRC_LEVEL_MEDIUM
	GstRpiCamSrcDRCLevelHigh   GstRpiCamSrcDRCLevel = "high"   // GST_RPI_CAM_SRC_DRC_LEVEL_HIGH
)

type GstRpiCamSrcExposureMeteringMode string

const (
	GstRpiCamSrcExposureMeteringModeAverage  GstRpiCamSrcExposureMeteringMode = "average"  // GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_AVERAGE
	GstRpiCamSrcExposureMeteringModeSpot     GstRpiCamSrcExposureMeteringMode = "spot"     // GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_SPOT
	GstRpiCamSrcExposureMeteringModeBacklist GstRpiCamSrcExposureMeteringMode = "backlist" // GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_BACKLIST
	GstRpiCamSrcExposureMeteringModeMatrix   GstRpiCamSrcExposureMeteringMode = "matrix"   // GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_MATRIX
)
